# Go 内存注入入门教程

> **学习定位**：安全研究 / 逆向分析 / 恶意软件防御理解 / 授权测试环境  
> **禁止用途**：未授权注入他人进程、制作木马、免杀对抗、破坏性攻击  
> **练习原则**：只对自己编译、自己启动的目标进程动手；建议全程在虚拟机中完成

本文是一份**从零到能动手实验**的路线图。目标不是做一个武器化注入器，而是让你真正理解：

1. Windows 进程与虚拟内存如何工作  
2. “注入”在系统 API 层面到底发生了什么  
3. 如何用 Go 调用这些 API，并完成可控、可调试的实验  
4. 防御方如何识别这类行为

---

## 目录

1. [学习前准备](#1-学习前准备)
2. [整体学习路线（建议 2～4 周）](#2-整体学习路线建议-24-周)
3. [第 0 阶段：先补这些基础](#3-第-0-阶段先补这些基础)
4. [第 1 阶段：本进程内存实验（最安全）](#4-第-1-阶段本进程内存实验最安全)
5. [第 2 阶段：远程进程打开与内存读写](#5-第-2-阶段远程进程打开与内存读写)
6. [第 3 阶段：经典 CreateRemoteThread 注入](#6-第-3-阶段经典-createremotethread-注入)
7. [第 4 阶段：DLL 注入（LoadLibrary 路径）](#7-第-4-阶段dll-注入loadlibrary-路径)
8. [第 5 阶段：其他注入技术概览](#8-第-5-阶段其他注入技术概览)
9. [第 6 阶段：调试、排错与观察](#9-第-6-阶段调试排错与观察)
10. [第 7 阶段：防御视角](#10-第-7-阶段防御视角)
11. [推荐资料清单](#11-推荐资料清单)
12. [实验项目结构建议](#12-实验项目结构建议)
13. [检查清单与常见坑](#13-检查清单与常见坑)
14. [下一步你可以做什么](#14-下一步你可以做什么)

配套章节文档：

| 文档 | 内容 |
|------|------|
| [01-foundations.md](./01-foundations.md) | 进程 / 线程 / 虚拟内存 / PE 基础 |
| [02-local-memory-lab.md](./02-local-memory-lab.md) | 本进程 VirtualAlloc 实验 |
| [03-remote-process-lab.md](./03-remote-process-lab.md) | OpenProcess / 读写远程内存 |
| [04-classic-injection.md](./04-classic-injection.md) | CreateRemoteThread 经典注入 |
| [05-dll-injection.md](./05-dll-injection.md) | LoadLibrary DLL 注入 |
| [06-debug-and-defense.md](./06-debug-and-defense.md) | 调试、排错、防御识别 |
| [resources.md](./resources.md) | 书籍 / 文档 / 工具 / 课程清单 |
| [lab-checklist.md](./lab-checklist.md) | 每日练习清单 |

---

## 1. 学习前准备

### 1.1 环境要求

| 项目 | 建议 |
|------|------|
| 操作系统 | Windows 10/11 x64（优先虚拟机） |
| 语言 | Go 1.21+ |
| 架构 | 注入器与目标进程位数一致（都用 amd64） |
| 权限 | 普通用户即可起步；部分实验可能需要管理员 |
| 安全 | 虚拟机快照；关闭“对真实主机的测试”念头 |

### 1.2 建议安装的工具

- **Go**：https://go.dev/dl/
- **Visual Studio Build Tools** 或 **MinGW-w64**（若以后要编原生 DLL）
- **Process Hacker** 或 **Process Explorer**：看进程、内存、句柄、线程
- **x64dbg** 或 **WinDbg**：动态调试
- **PE-bear / CFF Explorer**：看 PE 结构（可选）
- **API Monitor**（可选）：观察 API 调用

### 1.3 学习心态

先问“系统为什么允许这样做”，再问“代码怎么写”。

很多失败不是 Go 语法问题，而是：

- 权限不够（`OpenProcess` 失败）
- 架构不匹配（64 打 32）
- 内存保护属性不对
- 字符串编码错误（ANSI / UTF-16）
- 目标进程已退出或受保护

---

## 2. 整体学习路线（建议 2～4 周）

```text
第 0 周：基础概念
   进程 / 线程 / 虚拟地址空间 / 页保护 / PE 入门
        ↓
第 1 周：本进程实验
   VirtualAlloc → 写内存 → 改保护 → 本地调用
        ↓
第 2 周：跨进程读写
   自写 target.exe → OpenProcess → VirtualAllocEx → Write/Read
        ↓
第 3 周：经典注入
   CreateRemoteThread 执行远程地址
   LoadLibrary 路径 DLL 注入
        ↓
第 4 周：观察与防御
   调试器验证 / 事件日志思路 / 检测点理解
```

**不要跳级**：跳过本进程实验直接写远程注入，排错成本会高很多。

---

## 3. 第 0 阶段：先补这些基础

详细内容见 [01-foundations.md](./01-foundations.md)。

你至少要能回答：

1. 进程和线程有什么区别？
2. 虚拟地址和物理地址差在哪？
3. `PAGE_READWRITE` 和 `PAGE_EXECUTE_READ` 有什么区别？
4. 为什么“可写且可执行（RWX）”很显眼？
5. PE 文件大概由哪些部分组成？`LoadLibrary` 大致做了什么？

### 必学 Windows API（名字先混个脸熟）

| API | 作用 |
|-----|------|
| `OpenProcess` | 打开目标进程，拿到句柄 |
| `VirtualAlloc` / `VirtualAllocEx` | 分配内存（本进程 / 远程） |
| `VirtualProtect` / `VirtualProtectEx` | 修改页保护属性 |
| `WriteProcessMemory` | 向目标进程写内存 |
| `ReadProcessMemory` | 从目标进程读内存 |
| `CreateRemoteThread` | 在目标进程创建线程并指定入口 |
| `LoadLibraryA/W` | 加载 DLL |
| `GetModuleHandle` / `GetProcAddress` | 取模块与导出函数地址 |
| `CloseHandle` | 关闭句柄 |

官方文档入口：  
https://learn.microsoft.com/windows/win32/api/

---

## 4. 第 1 阶段：本进程内存实验（最安全）

详细实验见 [02-local-memory-lab.md](./02-local-memory-lab.md)。

### 目标

在**当前 Go 进程**里：

1. 分配一块内存  
2. 写入数据或机器码 stub  
3. 修改保护属性  
4. 以函数指针方式调用（或仅验证读写）

### 你将学会

- Go 如何通过 `golang.org/x/sys/windows` 调 Win32 API  
- 内存页保护的意义  
- 为什么 payload 缓冲区要稳定（避免 GC / 指针问题）

### 完成标准

- [ ] 能成功 `VirtualAlloc` 并打印返回地址  
- [ ] 能写入自定义字节并读回校验  
- [ ] 能解释 `MEM_COMMIT | MEM_RESERVE` 的含义  
- [ ] 能把保护从 `PAGE_READWRITE` 改成 `PAGE_EXECUTE_READ`（若做执行实验）

---

## 5. 第 2 阶段：远程进程打开与内存读写

详细实验见 [03-remote-process-lab.md](./03-remote-process-lab.md)。

### 目标

1. 自己写一个长期运行的 `target.exe`（例如循环打印心跳）  
2. 用另一个 Go 程序按 PID 打开它  
3. 在目标中分配内存、写入字符串、再读回来

### 关键 API 链

```text
OpenProcess
  → VirtualAllocEx
    → WriteProcessMemory
      → ReadProcessMemory
        → VirtualFreeEx（可选清理）
          → CloseHandle
```

### 完成标准

- [ ] `OpenProcess` 成功，失败时能打印 `GetLastError` / Windows 错误含义  
- [ ] 远程分配成功，能在 Process Hacker 中看到新内存区  
- [ ] 读写内容一致  
- [ ] 目标崩溃时能定位是权限、地址还是编码问题

---

## 6. 第 3 阶段：经典 CreateRemoteThread 注入

详细实验见 [04-classic-injection.md](./04-classic-injection.md)。

### 目标

理解最经典的跨进程执行模型：

```text
在目标分配可执行内存
  → 写入一段“无害、可验证”的代码或调用桩
    → CreateRemoteThread(入口=远程地址)
      → 观察目标行为变化 / 线程创建
```

### 学习重点

- 远程线程入口地址必须对目标进程有意义  
- 64 位进程的调用约定（RCX/RDX/R8/R9…）  
- 为什么教学上更推荐下一步的 **DLL 注入**（更容易验证）

### 完成标准

- [ ] 能说清 `CreateRemoteThread` 每个参数含义  
- [ ] 知道需要哪些进程访问权限  
- [ ] 能在调试器或工具中确认新线程启动  

---

## 7. 第 4 阶段：DLL 注入（LoadLibrary 路径）

详细实验见 [05-dll-injection.md](./05-dll-injection.md)。

### 目标

实现教学版流程：

```text
把 DLL 完整路径写入目标进程
  → 远程线程入口设为 LoadLibraryW
    → 参数为远程字符串地址
      → DLL 的 DllMain 被调用（例如弹消息框 / 写日志文件）
```

### 为什么这一步很适合入门

- 不需要手写复杂 shellcode  
- 结果可观察（日志、MessageBox、文件）  
- 完整覆盖“分配 → 写 → 远程执行 → 模块加载”链路  

### 完成标准

- [ ] 能编译一个简单 DLL（可用 C/C++ 或合适工具链）  
- [ ] Go 注入器成功让 target 加载该 DLL  
- [ ] 在 Process Hacker 模块列表中看到该 DLL  
- [ ] 能解释为何路径必须是目标进程可访问的绝对路径  

---

## 8. 第 5 阶段：其他注入技术概览

先建立分类地图，**不必立刻全部实现**。

| 技术 | 核心思想 | 难度 | 学习建议 |
|------|----------|------|----------|
| CreateRemoteThread | 远程新建线程执行 | ★★☆ | 必做 |
| LoadLibrary DLL 注入 | 远程调用加载器 | ★★☆ | 必做 |
| APC 注入 | 队列 APC 到可告警线程 | ★★★ | 了解原理后选做 |
| 线程劫持 | 挂起线程改上下文再恢复 | ★★★ | 了解原理 |
| 映射/手动映射 | 自己实现加载逻辑 | ★★★★ | 进阶 |
| 模块 stomping 等 | 复用已有模块空间 | ★★★★ | 进阶 / 分析向 |

进阶阅读建议从**恶意软件分析资料**的 Process Injection 分类章节入手，用防御视角学会识别。

---

## 9. 第 6 阶段：调试、排错与观察

详细内容见 [06-debug-and-defense.md](./06-debug-and-defense.md)。

### 每次实验都要做的 4 件事

1. **打印错误码**：失败不要只看 `err != nil`，要落到具体 Windows 错误  
2. **用工具验证**：Process Hacker 看内存区、线程、模块  
3. **最小化变量**：同一架构、同一用户、自己的 target  
4. **保留日志**：PID、地址、权限、路径、返回值全部记下来

### 高频失败原因

| 现象 | 常见原因 |
|------|----------|
| `OpenProcess` 失败 | 权限不足 / PID 错 / 受保护进程 |
| 写入失败 | 没有 `PROCESS_VM_WRITE` / 地址无效 |
| 远程线程失败 | 入口不可执行 / 架构不匹配 |
| DLL 没加载 | 路径错 / 依赖缺失 / 位数不匹配 |
| 目标闪退 | DLL 初始化崩溃 / 调用约定错误 |

---

## 10. 第 7 阶段：防御视角

学注入如果只学“怎么打”，视野是残缺的。至少要知道检测点：

- 异常的跨进程 `VirtualAllocEx` + `WriteProcessMemory` + `CreateRemoteThread` 组合  
- 新线程起始地址落在匿名可执行内存，而不是正常模块  
- 突然加载来路不明的 DLL  
- 父进程与注入行为不匹配的告警链  

建议练习：

1. 对自己的注入实验开 Sysmon（若方便）观察事件  
2. 在 Process Hacker 中对比“正常程序”和“被注入后”的差异  
3. 写一页笔记：若你是 EDR，会盯哪些 API 序列  

---

## 11. 推荐资料清单

完整清单与链接见 [resources.md](./resources.md)。下面是精简版。

### 官方文档（必收藏）

- [OpenProcess](https://learn.microsoft.com/windows/win32/api/processthreadsapi/nf-processthreadsapi-openprocess)
- [VirtualAllocEx](https://learn.microsoft.com/windows/win32/api/memoryapi/nf-memoryapi-virtualallocex)
- [WriteProcessMemory](https://learn.microsoft.com/windows/win32/api/memoryapi/nf-memoryapi-writeprocessmemory)
- [CreateRemoteThread](https://learn.microsoft.com/windows/win32/api/processthreadsapi/nf-processthreadsapi-createremotethread)
- [Memory Protection Constants](https://learn.microsoft.com/windows/win32/memory/memory-protection-constants)

### 体系化书籍

- *Windows Internals*（进程、线程、内存管理章节）  
- *Practical Malware Analysis*（动态分析与注入相关实验）  
- *Windows Kernel Programming* / 相关 Windows 系统编程书（按兴趣）

### Go 相关

- [`golang.org/x/sys/windows`](https://pkg.go.dev/golang.org/x/sys/windows)  
- Go 官方文档：`unsafe`、CGO（进阶才需要）

### 工具文档

- Process Hacker / System Informer 帮助  
- x64dbg 官方文档  
- Microsoft Sysinternals 套件说明  

---

## 12. 实验项目结构建议

建议你本地建这样一个练习仓库（本教程可与之对应）：

```text
go-memory-injection-lab/
├── README.md                 # 你的实验记录
├── target/
│   └── main.go               # 被注入的目标程序（心跳循环）
├── phase1_local/
│   └── main.go               # 本进程分配与读写
├── phase2_remote_rw/
│   └── main.go               # 远程读写
├── phase3_remote_thread/
│   └── main.go               # CreateRemoteThread 教学版
├── phase4_dll_inject/
│   ├── injector/main.go
│   └── payload_dll/          # 简单 DLL 源码
└── notes/
    ├── errors.md             # 失败与解决
    └── observations.md       # 工具截图与笔记
```

### 推荐依赖

```bash
go get golang.org/x/sys/windows
```

---

## 13. 检查清单与常见坑

每日练习表见 [lab-checklist.md](./lab-checklist.md)。

### Go 特有注意点

1. **架构一致**：`GOARCH=amd64` 编注入器，目标也必须是 64 位  
2. **字符串**：`LoadLibraryW` 需要 UTF-16，可用 `windows.UTF16PtrFromString`  
3. **切片稳定性**：传给系统 API 的 buffer 确保在调用期间有效  
4. **错误处理**：封装 `windows.GetLastError()` 或直接使用 `x/sys/windows` 返回的 `syscall.Errno`  
5. **不要一上来追求 shellcode**：先 DLL 路径注入，成功率与可观测性都更高  

### 安全红线

- 不注入浏览器、杀软、系统关键进程做“炫技”  
- 不传播注入器给他人用于未授权场景  
- 不把重点放在免杀、持久化、对抗上（那是另一条且高风险路径）  
- 课程作业 / 研究需要时，固定在隔离虚拟机  

---

## 14. 下一步你可以做什么

按你的目标选一条：

### 路线 A：扎实入门（推荐）

1. 读完 [01-foundations.md](./01-foundations.md)  
2. 做 [02-local-memory-lab.md](./02-local-memory-lab.md)  
3. 做 [03-remote-process-lab.md](./03-remote-process-lab.md)  
4. 再进入 DLL 注入  

### 路线 B：我直接帮你写实验代码

告诉我你现在的基础，例如：

- 是否会 Go  
- 是否装了虚拟机  
- 是否会一点 C/C++（编 DLL）  
- 想从 **本进程实验** 还是 **DLL 注入** 开始  

我可以按章节给你生成可运行的 `target` + 实验程序骨架。

### 路线 C：只补理论

先精读 foundations + resources，不写代码，把 API 调用链和权限模型画成一张图。

---

## 版本与维护

| 项 | 内容 |
|----|------|
| 文档版本 | v1.0 |
| 适用平台 | Windows x64 + Go |
| 更新原则 | 以官方 API 文档为准；工具名可能随社区更名（如 Process Hacker → System Informer） |

---

**开始建议**：打开 [01-foundations.md](./01-foundations.md)，用 1～2 小时把概念读完，再进入第 1 个动手实验。
