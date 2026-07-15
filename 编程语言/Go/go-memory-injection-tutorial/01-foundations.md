# 第 0 阶段：基础概念

> 目标：不写注入代码，也能讲清“内存注入”在操作系统里意味着什么。  
> 建议用时：2～4 小时阅读 + 自己画 1 张图。

---

## 1. 进程（Process）是什么

一个进程大致包含：

- **私有虚拟地址空间**（每个进程觉得自己独占一套地址）
- **一个或多个线程**
- **打开的句柄**（文件、进程、线程、同步对象等）
- **加载的模块**（exe 本体 + DLL）
- **安全令牌**（你是谁、有什么权限）

关键点：

> 进程 A 默认**不能**随便读写进程 B 的内存。  
> 跨进程操作必须通过操作系统提供的 API，并持有足够权限的句柄。

这也是注入的“门槛”：你要先合法地拿到目标进程句柄。

---

## 2. 线程（Thread）是什么

线程是 CPU 调度的基本单位。线程有：

- 指令指针（x64 上是 `RIP`）
- 寄存器上下文
- 栈
- 线程本地存储等

**注入与线程的关系**：

很多注入手法的本质是：

> 让目标进程里的某个线程，去执行你准备好的代码入口。

常见方式：

1. **新建远程线程**（`CreateRemoteThread`）→ 新线程从你给的地址开始跑  
2. **借用已有线程**（APC、线程劫持）→ 让老线程拐去执行你的代码  

---

## 3. 虚拟内存与页保护

### 3.1 虚拟地址空间

程序里的指针通常是**虚拟地址**。CPU + 操作系统把它翻译到物理页。

对注入来说，你关心的是：

- 在目标进程地址空间里**分配**一块区域  
- 这块区域对目标进程是否**可读 / 可写 / 可执行**

### 3.2 常见保护常量

| 常量 | 含义 | 注入场景 |
|------|------|----------|
| `PAGE_READWRITE` | 可读可写，不可执行 | 先写入 payload 时常用 |
| `PAGE_EXECUTE_READ` | 可读可执行，不可写 | 写完后改成执行更“规矩” |
| `PAGE_EXECUTE_READWRITE` | 可读可写可执行（RWX） | 省事但非常显眼 |
| `PAGE_READONLY` | 只读 | 数据区 |

教学建议：

```text
先 VirtualAllocEx(..., PAGE_READWRITE)
写完 payload
再 VirtualProtectEx(..., PAGE_EXECUTE_READ)
```

比一上来 RWX 更接近“正常加载器”思路，也便于理解 DEP（数据执行保护）。

### 3.3 分配类型

- `MEM_RESERVE`：预约地址范围  
- `MEM_COMMIT`：真正提交物理/页文件支持  
- 教学里常见：`MEM_RESERVE | MEM_COMMIT`

---

## 4. 句柄与权限

### 4.1 句柄（Handle）

`OpenProcess` 成功后返回的是一个**句柄**，不是 PID 本身。  
后续 `VirtualAllocEx`、`WriteProcessMemory` 都基于这个句柄。

用完要 `CloseHandle`。

### 4.2 进程访问权限（节选）

| 权限 | 用途 |
|------|------|
| `PROCESS_CREATE_THREAD` | 创建远程线程 |
| `PROCESS_VM_OPERATION` | 分配/改保护等内存操作 |
| `PROCESS_VM_WRITE` | 写内存 |
| `PROCESS_VM_READ` | 读内存 |
| `PROCESS_QUERY_INFORMATION` | 查询信息 |
| `PROCESS_ALL_ACCESS` | 全开（教学省事，真实环境权限过大） |

最小权限原则：教学阶段可以先用较宽权限保证跑通，但你要知道每项权限对应哪一步。

### 4.3 为什么有时 OpenProcess 会失败

- PID 不存在或已退出  
- 目标是受保护进程（PPL 等）  
- 完整性级别不够（例如低完整性去开高完整性）  
- 权限掩码给错  
- 安全软件拦截  

---

## 5. PE 文件与 DLL 加载（入门级）

Windows 可执行文件（EXE/DLL）是 **PE（Portable Executable）** 格式。

你暂时只需知道：

1. 文件被映射进进程地址空间  
2. 导入表描述“我依赖哪些 DLL 的哪些函数”  
3. `LoadLibrary` 负责把 DLL 加载进来，并完成必要的重定位、导入解析  
4. DLL 的入口 `DllMain` 可能在加载时被调用  

因此 **DLL 注入** 的经典思路是：

> 不自己实现加载器，而是让目标进程调用已有的 `LoadLibraryW("C:\\path\\payload.dll")`。

这比手写 shellcode 容易验证得多。

---

## 6. “内存注入”到底在做什么

把花活剥掉，核心通常是四步：

```text
1) Attach  : 打开目标进程（拿到句柄）
2) Allocate: 在目标地址空间准备内存
3) Write   : 把代码或数据写进去
4) Execute : 让目标以某种方式执行到该入口
```

对应最经典的实现：

```text
OpenProcess
VirtualAllocEx
WriteProcessMemory
CreateRemoteThread
```

你后面所有变体（APC、劫持、手动映射）都是在改 **第 4 步如何执行**，或 **第 2/3 步如何更隐蔽地落内存**。

---

## 7. Go 在这条链路中的位置

Go 不提供“注入标准库”，它做的是：

- 用 `golang.org/x/sys/windows` 声明并调用 Win32 API  
- 管理字符串、字节切片、错误  
- 组织你的实验程序结构  

注意：

- Go 运行时有 GC；传给系统 API 的内存要在调用期间保持有效  
- Go 程序默认可能是动态栈；教学代码里尽量用明确的 `[]byte` / UTF-16 buffer  
- 先把 Windows 语义学对，再纠结 Go 语法糖  

---

## 8. 一张总览图（请自己重画一遍）

```text
[Injector 进程]
    |  OpenProcess(pid, access)
    v
[Target 进程句柄]
    |  VirtualAllocEx
    v
[Target 虚拟内存空洞] ---- WriteProcessMemory ----> [payload 字节]
    |
    |  CreateRemoteThread(start = payload 或 LoadLibrary)
    v
[Target 新线程开始执行]
    |
    +--> 直接跑 shellcode
    +--> 或 LoadLibrary 加载 DLL → DllMain
```

---

## 9. 本阶段自测题

不看资料回答：

1. 为什么注入通常需要 `PROCESS_VM_WRITE`？  
2. 为什么只写内存而不创建执行路径，目标不会自动跑 payload？  
3. `PAGE_READWRITE` 的内存能直接当代码执行吗？默认策略下通常怎样？  
4. DLL 注入为什么要把**路径字符串**写到目标进程，而不是写在注入器自己的栈上？  
5. 64 位注入器去注入 32 位进程，最可能出什么问题？  

参考答案要点：

1. 没有写权限，`WriteProcessMemory` 会失败。  
2. 内存里有代码 ≠ 有线程在执行它；必须新建线程、劫持线程或排队 APC 等。  
3. 有 DEP 时，不可执行页取指会出问题；应改为可执行保护。  
4. `LoadLibrary` 在目标进程上下文中执行，它读的是目标进程地址空间里的字符串。  
5. 指针宽度、子系统、模块地址、系统 DLL 布局都不兼容，极易失败。  

---

## 10. 完成后去哪里

基础过关后，进入：

→ [02-local-memory-lab.md](./02-local-memory-lab.md) 本进程内存实验
