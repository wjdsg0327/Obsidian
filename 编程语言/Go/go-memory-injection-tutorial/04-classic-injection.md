# 第 3 阶段：经典 CreateRemoteThread 注入

> 目标：理解“远程线程执行”模型，而不是追求复杂 payload。  
> 建议：先读懂，再决定是否自己实现；入门更推荐先完成 DLL 注入章节。  
> 建议用时：概念 1 小时；动手 2～6 小时（视基础而定）。

---

## 1. 这一招在做什么

经典流程：

```text
OpenProcess
VirtualAllocEx          // 在目标里准备空间
WriteProcessMemory      // 写入要执行的内容或参数
CreateRemoteThread      // 让目标新线程从指定地址开始执行
WaitForSingleObject     // 可选：等待线程结束
GetExitCodeThread       // 可选：取返回值
CloseHandle
```

一句话：

> 你不是“魔法改了 CPU”，而是请操作系统在目标进程里**新建一个线程**，入口由你指定。

---

## 2. 两种教学级 payload 策略

### 策略 A：远程调用已有导出函数（更稳）

例如让远程线程去执行 `LoadLibraryW`：

- 入口地址：目标进程中 `kernel32!LoadLibraryW`  
- 参数：远程内存里的 DLL 路径字符串  

这就是下一章的 DLL 注入，也是最推荐的第一条“能跑通”路径。

### 策略 B：写入自定义 shellcode（更难）

- 你要保证机器码正确  
- 调用约定正确  
- 地址与重定位正确  
- 崩溃时极难排错  

**入门不建议先走 B。**

---

## 3. CreateRemoteThread 参数怎么读

简化理解：

| 参数 | 含义 |
|------|------|
| `hProcess` | 目标进程句柄 |
| `lpStartAddress` | 远程线程入口（目标地址空间中的地址） |
| `lpParameter` | 传给线程函数的参数（一个指针宽度） |
| 返回值 | 新线程句柄 |

线程函数在概念上类似：

```c
DWORD WINAPI ThreadProc(LPVOID lpParameter);
```

因此：

- 如果你把入口设为 `LoadLibraryW`  
- 那么 `lpParameter` 就应是**目标进程中**那条 DLL 路径字符串的地址  

---

## 4. 权限清单（经典组合）

通常需要：

```text
PROCESS_CREATE_THREAD
PROCESS_QUERY_INFORMATION
PROCESS_VM_OPERATION
PROCESS_VM_WRITE
PROCESS_VM_READ
```

少任何一个，都可能在不同步骤失败。

---

## 5. 学习型实现路径（推荐顺序）

### Step 1：只创建远程线程，入口用目标已有安全函数？

不现实地乱指定入口会崩。所以：

### Step 1'（推荐）：直接做 DLL 注入

见 [05-dll-injection.md](./05-dll-injection.md)。

它完整覆盖了 `CreateRemoteThread` 的关键点，但结果可观察：

- 模块列表多了一个 DLL  
- DLL 可写文件 / 弹窗 / 打日志  

### Step 2：回头总结

写一篇笔记回答：

1. 远程线程入口地址从哪里来？  
2. 参数为什么必须在目标地址空间？  
3. 线程返回值你如何获取？  
4. 为什么 RWX 匿名内存里的入口比模块内入口更可疑？  

---

## 6. 伪代码级骨架（帮助理解，不等于直接抄作业）

```text
hProcess = OpenProcess(...)
remoteStr = VirtualAllocEx(...)
WriteProcessMemory(hProcess, remoteStr, dllPathUTF16)

hKernel32 = GetModuleHandle("kernel32.dll")
pLoadLibraryW = GetProcAddress(hKernel32, "LoadLibraryW")
// 说明：kernel32 在多数进程中基址一致的假设，是经典教学简化；
// 真实世界 ASLR/差异场景需要更严谨处理。

hThread = CreateRemoteThread(
    hProcess,
    start = pLoadLibraryW,
    param = remoteStr,
)

WaitForSingleObject(hThread, timeout)
CloseHandle(hThread)
CloseHandle(hProcess)
```

---

## 7. 架构与地址问题（高频坑）

1. **注入器与目标位数必须一致**  
2. `GetProcAddress` 在注入器进程拿到的地址，只有在“目标中同模块映射基址相同”时才能直接复用  
   - 对 `kernel32.dll` 的经典实验常常能工作  
   - 这不代表所有 DLL 都能这么做  
3. WOW64（64 位系统上的 32 位进程）情况更复杂，入门先全用 64 位  

---

## 8. 观察方法

注入成功后检查：

- Process Hacker → Threads：是否出现新线程  
- Modules：是否出现 payload DLL  
- target 控制台 / 日志文件是否有副作用  

失败时检查：

- `CreateRemoteThread` 错误码  
- 入口地址是否为 0  
- 参数地址是否为远程地址  
- DLL 路径是否存在、是否可被目标访问  

---

## 9. 本阶段完成标准

- [ ] 能手绘 CreateRemoteThread 注入时序图  
- [ ] 能解释为何“只写入 shellcode 而不创建线程”不会自动执行  
- [ ] 能说明 DLL 注入为什么是 CreateRemoteThread 的最佳入门实例  
- [ ] 知道至少 3 个失败原因与对应排查手段  

---

## 10. 完成后去哪里

→ [05-dll-injection.md](./05-dll-injection.md)  
→ 完成后读 [06-debug-and-defense.md](./06-debug-and-defense.md)
