# 第 4 阶段：DLL 注入（LoadLibrary 路径）

> 目标：完成第一条真正“可见效果”的跨进程注入实验。  
> 效果示例：target 加载你的 DLL 后写日志或弹出消息框。  
> 建议用时：半天到 1 天。

---

## 1. 原理（先背熟）

```text
1. 注入器 OpenProcess(target)
2. VirtualAllocEx 在 target 中分配内存
3. WriteProcessMemory 写入 DLL 的绝对路径（UTF-16）
4. 获取 LoadLibraryW 地址
5. CreateRemoteThread(
     start = LoadLibraryW,
     param = 远程路径字符串地址
   )
6. 目标进程加载 DLL，DllMain 执行
```

关键认知：

- 被执行的是目标进程里的 `LoadLibraryW`  
- 路径字符串必须位于目标进程地址空间  
- DLL 文件必须是目标能加载的（路径存在、位数匹配、依赖齐全）

---

## 2. 准备 Payload DLL

你需要一个很简单的 DLL。任选一种方式：

### 方式 A：用 C/C++（最常见）

`payload.c` 概念示例：

```c
#include <windows.h>

BOOL WINAPI DllMain(HINSTANCE h, DWORD reason, LPVOID reserved) {
    if (reason == DLL_PROCESS_ATTACH) {
        // 教学用途：写文件比 MessageBox 更不容易打断实验
        // 也可改成 MessageBoxW(NULL, L"injected", L"lab", MB_OK);
        HANDLE f = CreateFileW(
            L"C:\\Users\\Public\\dll-inject-lab.txt",
            GENERIC_WRITE, FILE_SHARE_READ, NULL,
            CREATE_ALWAYS, FILE_ATTRIBUTE_NORMAL, NULL);
        if (f != INVALID_HANDLE_VALUE) {
            const char *msg = "dll loaded in target\r\n";
            DWORD written;
            WriteFile(f, msg, (DWORD)lstrlenA(msg), &written, NULL);
            CloseHandle(f);
        }
    }
    return TRUE;
}
```

编译成 64 位 DLL（工具链按你环境选择 MSVC 或 MinGW）。

### 方式 B：暂时用现成教学 DLL

若你还不会编 DLL，先把“如何编译 DLL”单独学半天，再回来。  
**不要**从不明来源下载 DLL。

### 检查

- DLL 是 x64  
- 依赖尽量少  
- 绝对路径待会能传给注入器  

---

## 3. Go 注入器要点

### 3.1 路径必须是绝对路径

推荐：

```text
C:\Users\Public\payload.dll
```

而不是相对路径 `.\payload.dll`（目标进程当前目录未必和你一样）。

### 3.2 使用 UTF-16

`LoadLibraryW` 要宽字符：

```go
pathUTF16, err := windows.UTF16FromString(`C:\Users\Public\payload.dll`)
// 写入的是 UTF-16LE 字节序列，含结尾 0
```

### 3.3 获取 LoadLibraryW 地址

教学简化：

```go
hKernel32, err := windows.GetModuleHandle(windows.StringToUTF16Ptr("kernel32.dll"))
pLoadLibraryW, err := windows.GetProcAddress(hKernel32, "LoadLibraryW")
```

说明：这是经典实验假设。先跑通，再在笔记里写下它的前提与局限。

---

## 4. 注入器流程伪代码（按此实现）

```text
pid <- 命令行参数
dllPath <- 命令行参数（绝对路径）

access = PROCESS_CREATE_THREAD | PROCESS_VM_OPERATION |
         PROCESS_VM_WRITE | PROCESS_VM_READ | PROCESS_QUERY_INFORMATION

hProcess = OpenProcess(access, false, pid)

pathUTF16 = UTF16FromString(dllPath)
size = len(pathUTF16) * 2

remote = VirtualAllocEx(hProcess, size, PAGE_READWRITE)
WriteProcessMemory(hProcess, remote, pathUTF16 bytes)

pLoadLibraryW = GetProcAddress(GetModuleHandle(kernel32), LoadLibraryW)

hThread = CreateRemoteThread(hProcess, pLoadLibraryW, remote)
WaitForSingleObject(hThread, 10s)
exitCode = GetExitCodeThread(hThread)   // LoadLibrary 返回 HMODULE

// 观察 exitCode 是否非 0
// 用 Process Hacker 看 Modules
// 检查 C:\Users\Public\dll-inject-lab.txt 是否生成
```

---

## 5. 验证成功的 4 个信号

1. `CreateRemoteThread` 成功返回线程句柄  
2. 线程退出码非 0（作为 `HMODULE` 的教学观察）  
3. Process Hacker 的 Modules 列表出现 `payload.dll`  
4. 你设计的副作用出现（日志文件 / 消息框）  

任意一个不满足，就还没真正成功。

---

## 6. 分步排错表

| 步骤 | 失败表现 | 排查 |
|------|----------|------|
| OpenProcess | 拒绝访问 | 权限、PID、目标是否受保护 |
| VirtualAllocEx | NULL | 权限是否含 `PROCESS_VM_OPERATION` |
| WriteProcessMemory | 失败 | 路径 buffer、大小、写权限 |
| CreateRemoteThread | 失败 | 是否含 `PROCESS_CREATE_THREAD`、入口是否有效 |
| 线程成功但无模块 | 路径错 / DLL 位数错 / 依赖缺失 | 手动在目标同环境下 `LoadLibrary` 测试 |
| 目标崩溃 | DllMain 写崩了 | 简化 DllMain，只写文件 |

手动验证 DLL 可加载（在同架构小程序中）非常有用。

---

## 7. 练习作业

### 必做

- [ ] target 常驻运行  
- [ ] 注入器接收 `pid` 与 `dllPath` 两个参数  
- [ ] 成功看到模块与日志文件  
- [ ] 记录完整命令与输出  

### 选做

- [ ] DLL 改为追加写入，多次注入观察  
- [ ] 注入后用工具卸载模块？（仅研究，注意风险）  
- [ ] 对比 `LoadLibraryA` 与 `LoadLibraryW` 路径编码差异  
- [ ] 把每次 API 返回值都写进 JSON 实验日志  

---

## 8. 安全与法律提醒（再写一次）

本实验仅用于：

- 学习 Windows 机制  
- 理解恶意软件常见手法  
- 在**自有实验环境**中验证  

不要：

- 对未授权目标注入  
- 把 DLL 做成窃密/键盘记录等功能  
- 分享“一键注入任意进程”的攻击用法  

---

## 9. 完成后去哪里

→ [06-debug-and-defense.md](./06-debug-and-defense.md)  
→ 回看 [resources.md](./resources.md) 补体系化资料  
→ 用 [lab-checklist.md](./lab-checklist.md) 做阶段验收
