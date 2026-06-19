---
title: Win32 内存管理 API 速查
date: 2026-06-19
tags: [Windows, Win32, API, 内存管理, Cpp]
source: https://learn.microsoft.com/en-us/windows/win32/memory/memory-management-functions
---

# Win32 内存管理 API 速查

> 来源：Microsoft Learn - Memory Management Functions  
> 链接：https://learn.microsoft.com/en-us/windows/win32/memory/memory-management-functions

## 总览

Windows 的内存管理器负责：

- 虚拟内存
- 堆内存
- 内存映射文件
- 写时复制 Copy-on-write
- 大页内存 Large Pages
- 文件缓存
- DEP 数据执行保护
- AWE 物理页映射
- Enclave 隔离内存区域

实际开发里最常用的是这几组：

| 类别 | 常见用途 |
|---|---|
| Virtual memory | 直接申请/释放虚拟内存页，控制读写/执行权限 |
| Heap | 普通动态内存分配，类似 malloc/free 的底层能力 |
| File mapping | 内存映射文件、共享内存、大文件访问 |
| General memory | 拷贝、清零、查询系统内存状态 |
| DEP | 查询/设置数据执行保护 |
| Global/Local | 老式兼容 API，剪贴板/OLE 场景还会见到 |

---

# 一、General memory functions：通用内存函数

| API | 简要说明 |
|---|---|
| AddSecureMemoryCacheCallback | 注册回调：安全内存区域被释放或保护属性变化时触发。 |
| CopyDeviceMemory | 复制设备内存，避免编译器优化干扰，适合特殊设备内存访问。 |
| CopyMemory | 复制一块内存，类似 `memcpy`。 |
| CopyVolatileMemory | 复制 volatile 内存块，避免某些优化导致访问被省略。 |
| CreateMemoryResourceNotification | 创建内存资源通知对象，用于监听系统内存资源状态。 |
| FillDeviceMemory | 填充设备内存，避免编译器优化干扰。 |
| FillMemory | 用指定值填充一块内存，类似 `memset`。 |
| FillVolatileMemory | 填充 volatile 内存。 |
| GetLargePageMinimum | 获取大页内存的最小页大小。 |
| GetPhysicallyInstalledSystemMemory | 获取机器实际安装的物理内存容量。 |
| GetSystemFileCacheSize | 获取系统文件缓存工作集大小限制。 |
| GetWriteWatch | 获取某段虚拟内存中被写入过的页地址。 |
| GlobalMemoryStatusEx | 查询当前物理内存和虚拟内存使用情况。 |
| MoveMemory | 移动内存块，支持源和目标区域重叠，类似 `memmove`。 |
| MoveVolatileMemory | 移动 volatile 内存块，支持重叠区域。 |
| QueryMemoryResourceNotification | 查询内存资源通知对象当前状态。 |
| RemoveSecureMemoryCacheCallback | 取消注册安全内存回调。 |
| ResetWriteWatch | 重置虚拟内存区域的写入跟踪状态。 |
| SecureMemoryCacheCallback | 应用自定义回调函数，用于安全内存区域变化通知。 |
| SecureZeroMemory | 安全清零内存，常用于清除密码、密钥，避免被编译器优化掉。 |
| SecureZeroMemory2 | 更严格保证的安全清零函数。 |
| SetSystemFileCacheSize | 设置系统文件缓存工作集大小限制。 |
| ZeroDeviceMemory | 把设备内存清零，避免编译器优化干扰。 |
| ZeroMemory | 清零内存，类似 `memset(ptr, 0, size)`。 |
| ZeroVolatileMemory | 清零 volatile 内存。 |

## 常用重点

- `GlobalMemoryStatusEx`：看系统内存状态。
- `CopyMemory` / `MoveMemory` / `ZeroMemory`：传统内存操作宏/函数。
- `SecureZeroMemory`：清理密码、Token、密钥时优先用。

---

# 二、DEP functions：数据执行保护函数

DEP：Data Execution Prevention，数据执行保护。它的核心思想是：**普通数据内存不应该被当作代码执行**，用于降低缓冲区溢出类攻击风险。

| API | 简要说明 |
|---|---|
| GetProcessDEPPolicy | 获取某个进程的 DEP 设置。 |
| GetSystemDEPPolicy | 获取系统级 DEP 策略。 |
| SetProcessDEPPolicy | 修改当前进程的 DEP 策略。 |

## 常用重点

普通程序一般不用手动设置 DEP；做安全、壳、JIT、逆向、漏洞分析时会遇到。

---

# 三、File mapping functions：内存映射文件函数

内存映射文件可以把文件内容映射到进程地址空间，像访问内存一样访问文件。也可用于进程间共享内存。

| API | 简要说明 |
|---|---|
| CreateFileMappingA | 创建或打开 ANSI 版本文件映射对象。 |
| CreateFileMappingW | 创建或打开 Unicode 版本文件映射对象。常用。 |
| CreateFileMapping2 | 创建文件映射对象，支持扩展参数，例如 NUMA 节点。 |
| CreateFileMappingFromApp | Windows Store/UWP 应用可用的文件映射创建函数。 |
| CreateFileMappingNuma | 创建文件映射对象，并指定 NUMA 节点。 |
| FlushViewOfFile | 把映射视图中的修改刷新到磁盘。 |
| GetMappedFileName | 查询某个地址是否属于内存映射文件，并返回对应文件名。 |
| MapViewOfFile | 把文件映射对象映射到当前进程地址空间。常用。 |
| MapViewOfFile2 | 把文件或页文件支持的 section 映射到指定进程。 |
| MapViewOfFile3 | 更现代的映射函数，支持更多参数控制。 |
| MapViewOfFile3FromApp | Store/UWP 版本的 MapViewOfFile3。 |
| MapViewOfFileEx | 映射文件视图，可建议映射到指定地址。 |
| MapViewOfFileExNuma | 映射文件视图，并指定 NUMA 节点。 |
| MapViewOfFileFromApp | Store/UWP 版本映射函数。 |
| MapViewOfFileNuma2 | 支持 NUMA 的现代映射函数。 |
| OpenFileMapping | 打开已有的命名文件映射对象。 |
| OpenFileMappingFromApp | Store/UWP 版本打开文件映射对象。 |
| UnmapViewOfFile | 取消映射当前进程中的文件视图。常用。 |
| UnmapViewOfFile2 | 取消映射文件或 pagefile-backed section 视图。 |
| UnmapViewOfFileEx | 扩展版取消映射函数。 |

## 典型流程

```cpp
CreateFileW        // 打开文件
CreateFileMappingW // 创建文件映射对象
MapViewOfFile      // 映射到内存
// 像访问内存一样读写文件内容
FlushViewOfFile    // 需要时刷新
UnmapViewOfFile    // 取消映射
CloseHandle        // 关闭句柄
```

## 常用重点

- 大文件读取：`CreateFileMappingW` + `MapViewOfFile`
- 共享内存：`CreateFileMappingW(INVALID_HANDLE_VALUE, ...)`
- 释放必须配对：`MapViewOfFile` 对应 `UnmapViewOfFile`

---

# 四、AWE functions：地址窗口扩展函数

AWE：Address Windowing Extensions。用于把物理内存页映射/取消映射到进程虚拟地址空间。现在普通应用较少用，更多是数据库、特殊高性能服务或旧 32 位大内存场景。

| API | 简要说明 |
|---|---|
| AllocateUserPhysicalPages | 分配可映射到 AWE 区域的物理内存页。 |
| AllocateUserPhysicalPagesNuma | 分配物理页，并指定 NUMA 节点。 |
| FreeUserPhysicalPages | 释放之前分配的物理页。 |
| MapUserPhysicalPages | 把已分配物理页映射到 AWE 虚拟地址区域。 |
| MapUserPhysicalPagesScatter | 分散方式映射物理页。 |

## 常用重点

一般业务程序很少直接用 AWE；如果只是普通内存申请，用 `HeapAlloc` 或 `VirtualAlloc`。

---

# 五、Heap functions：堆内存函数

Heap API 是 Windows 进程堆管理接口。C 运行库的 `malloc/free` 底层通常也会间接使用系统堆能力。

| API | 简要说明 |
|---|---|
| GetProcessHeap | 获取当前进程默认堆句柄。常用。 |
| GetProcessHeaps | 获取当前进程所有堆句柄。 |
| HeapAlloc | 从指定堆分配内存。常用。 |
| HeapCompact | 合并堆中的相邻空闲块，尝试整理碎片。 |
| HeapCreate | 创建一个私有堆。 |
| HeapDestroy | 销毁私有堆。 |
| HeapFree | 释放由 HeapAlloc/HeapReAlloc 分配的内存。常用。 |
| HeapLock | 锁定指定堆。 |
| HeapQueryInformation | 查询堆信息。 |
| HeapReAlloc | 重新分配堆内存大小。 |
| HeapSetInformation | 设置堆信息，例如启用某些堆特性。 |
| HeapSize | 获取某块堆内存大小。 |
| HeapUnlock | 解锁指定堆。 |
| HeapValidate | 校验堆或堆块是否有效。 |
| HeapWalk | 枚举堆中的内存块。 |

## 典型流程

```cpp
HANDLE heap = GetProcessHeap();
void* p = HeapAlloc(heap, 0, size);
// 使用 p
HeapFree(heap, 0, p);
```

## 常用重点

- 普通 C/C++ 推荐优先用 `new/delete`、`malloc/free` 或 STL 容器。
- 需要直接调用 Win32 堆时，用 `GetProcessHeap` + `HeapAlloc` + `HeapFree`。
- 自建堆场景：大量同生命周期对象、插件隔离、特殊内存管理。

---

# 六、Virtual memory functions：虚拟内存函数

这是 Win32 内存管理里非常核心的一组。它直接操作虚拟地址空间，以“页”为单位申请、提交、释放和改权限。

| API | 简要说明 |
|---|---|
| DiscardVirtualMemory | 丢弃一段内存页内容，但不取消提交；之后内容未定义，需要重新写入。 |
| OfferVirtualMemory | 告诉系统这段内存暂时不需要，内存紧张时可以回收。 |
| PrefetchVirtualMemory | 预取虚拟地址范围到物理内存。 |
| QueryVirtualMemoryInformation | 查询指定进程虚拟地址空间中页或页集合的信息。 |
| ReclaimVirtualMemory | 重新取回之前 Offer 给系统的内存。 |
| SetProcessValidCallTargets | 为 CFG 控制流保护设置合法间接调用目标。 |
| VirtualAlloc | 在当前进程虚拟地址空间保留或提交内存页。最常用。 |
| VirtualAlloc2 | 更现代的虚拟内存分配函数，支持扩展参数。 |
| VirtualAlloc2FromApp | Store/UWP 可用版本。 |
| VirtualAllocEx | 在指定进程地址空间分配内存。常用于调试、注入、跨进程操作。 |
| VirtualAllocExNuma | 在指定进程中分配内存，并指定 NUMA 节点。 |
| VirtualAllocFromApp | Store/UWP 可用版本。 |
| VirtualFree | 释放或取消提交当前进程虚拟内存页。常用。 |
| VirtualFreeEx | 释放指定进程中的虚拟内存。 |
| VirtualLock | 将指定虚拟内存页锁进物理内存，尽量避免换出到页面文件。 |
| VirtualProtect | 修改当前进程已提交内存页的访问权限。常用。 |
| VirtualProtectEx | 修改指定进程内存页访问权限。 |
| VirtualProtectFromApp | Store/UWP 可用版本。 |
| VirtualQuery | 查询当前进程某段虚拟地址的信息。常用。 |
| VirtualQueryEx | 查询指定进程某段虚拟地址的信息。 |
| VirtualUnlock | 解锁之前 VirtualLock 的内存页。 |

## 典型流程

```cpp
void* p = VirtualAlloc(
    nullptr,
    size,
    MEM_RESERVE | MEM_COMMIT,
    PAGE_READWRITE
);

// 使用 p

VirtualFree(p, 0, MEM_RELEASE);
```

## 常用重点

- `VirtualAlloc`：申请大块内存、页级内存、可执行内存、保护页。
- `VirtualProtect`：修改内存页权限，例如只读、可写、可执行。
- `VirtualQuery`：扫描进程地址空间，逆向/调试常用。
- `VirtualAllocEx` + `WriteProcessMemory` + `CreateRemoteThread`：经典远程注入链条的一部分。

---

# 七、Global and local functions：全局/本地内存函数

这是早期 Windows 遗留 API。现代 Windows 中 Global/Local 内存基本都来自堆，但某些老接口仍要求使用它们，例如剪贴板、OLE、DDE 等。

| API | 简要说明 |
|---|---|
| GlobalAlloc / LocalAlloc | 分配全局/本地内存块。 |
| GlobalDiscard / LocalDiscard | 丢弃指定内存块。 |
| GlobalFlags / LocalFlags | 获取内存对象标志信息。 |
| GlobalFree / LocalFree | 释放全局/本地内存对象。 |
| GlobalHandle / LocalHandle | 从指针获取对应内存句柄；主要用于 OLE/剪贴板兼容。 |
| GlobalLock / LocalLock | 锁定内存对象并返回可访问指针。 |
| GlobalReAlloc / LocalReAlloc | 调整内存对象大小或属性。 |
| GlobalSize / LocalSize | 获取内存对象当前大小。 |
| GlobalUnlock / LocalUnlock | 解锁内存对象。 |

## 常用重点

- 现代代码一般不用 `GlobalAlloc/LocalAlloc`。
- 剪贴板 API 常见组合：`GlobalAlloc` → `GlobalLock` → 写入数据 → `GlobalUnlock` → `SetClipboardData`。

---

# 八、Bad memory functions：坏内存通知函数

用于检测和处理硬件层面的坏内存页通知，普通应用很少用。

| API | 简要说明 |
|---|---|
| BadMemoryCallbackRoutine | 应用自定义回调：检测到坏内存页时调用。 |
| GetMemoryErrorHandlingCapabilities | 获取系统内存错误处理能力。 |
| RegisterBadMemoryNotification | 注册坏内存页通知。 |
| UnregisterBadMemoryNotification | 关闭坏内存通知句柄。 |

---

# 九、Enclave functions：隔离区函数

Enclave 是进程地址空间中的隔离代码和数据区域，只有 enclave 内部代码能访问 enclave 内部数据。用于安全计算、保护敏感数据等场景。

| API | 简要说明 |
|---|---|
| CreateEnclave | 创建未初始化的 enclave。 |
| InitializeEnclave | 初始化已创建并加载数据的 enclave。 |
| IsEnclaveTypeSupported | 查询系统是否支持指定类型的 enclave。 |
| LoadEnclaveData | 向未初始化 enclave 加载数据。 |

---

# 十、ATL thunk functions：ATL thunk 函数

ATL thunk 用于 ATL 框架内部生成小段跳转/适配代码，普通 Win32 开发很少直接调用。

| API | 简要说明 |
|---|---|
| AtlThunk_AllocateData | 为 ATL thunk 分配内存。 |
| AtlThunk_DataToCode | 根据 thunk 数据返回可执行函数地址。 |
| AtlThunk_FreeData | 释放 thunk 相关内存。 |
| AtlThunk_InitData | 初始化 ATL thunk 数据。 |

---

# 十一、按用途速查

## 1. 想申请普通内存

优先：

- C++：`new/delete`、`std::vector`、`std::string`
- C：`malloc/free`

需要 Win32：

- `HeapAlloc`
- `HeapFree`
- `GetProcessHeap`

## 2. 想申请页级内存

用：

- `VirtualAlloc`
- `VirtualFree`
- `VirtualProtect`
- `VirtualQuery`

适合：

- 大块内存
- 自定义内存池
- JIT
- 保护页
- 逆向/调试/注入分析

## 3. 想做内存映射文件

用：

- `CreateFileMappingW`
- `MapViewOfFile`
- `FlushViewOfFile`
- `UnmapViewOfFile`

适合：

- 大文件读取
- 进程间共享内存
- 高性能文件访问

## 4. 想查系统内存状态

用：

- `GlobalMemoryStatusEx`
- `GetPhysicallyInstalledSystemMemory`

## 5. 想清除敏感数据

用：

- `SecureZeroMemory`
- `SecureZeroMemory2`

不要只用普通 `memset`，因为编译器可能认为“后面不用了”而优化掉。

## 6. 想修改内存权限

用：

- `VirtualProtect`
- `VirtualProtectEx`

常见权限：

| 权限 | 含义 |
|---|---|
| PAGE_READONLY | 只读 |
| PAGE_READWRITE | 可读写 |
| PAGE_EXECUTE | 只执行 |
| PAGE_EXECUTE_READ | 可执行 + 可读 |
| PAGE_EXECUTE_READWRITE | 可执行 + 可读写 |
| PAGE_NOACCESS | 不可访问 |
| PAGE_GUARD | 保护页，访问时触发异常 |

---

# 十二、学习顺序建议

如果是为了理解 Windows 内存机制，建议按这个顺序学：

1. `GlobalMemoryStatusEx`：先学会看系统内存状态
2. `HeapAlloc / HeapFree`：理解普通堆内存
3. `VirtualAlloc / VirtualFree`：理解虚拟内存页
4. `VirtualProtect / VirtualQuery`：理解内存权限与地址空间
5. `CreateFileMapping / MapViewOfFile`：理解内存映射文件和共享内存
6. `SecureZeroMemory`：理解敏感内存清理
7. `DEP / CFG / Enclave`：再看安全相关机制

---

# 十三、几个核心概念

## Reserve 与 Commit

`VirtualAlloc` 里常见两个动作：

- `MEM_RESERVE`：保留一段虚拟地址空间，但不一定占用实际物理内存。
- `MEM_COMMIT`：提交内存页，使其可以实际读写，由系统分配物理页或页面文件支持。

简单理解：

- Reserve：先占坑。
- Commit：真正能用。

## Page Protection

Windows 虚拟内存以页为单位设置访问权限，例如只读、可读写、可执行、不可访问。

这就是为什么很多安全机制会关注：

- 数据页是否可执行
- 代码页是否可写
- 是否存在 `PAGE_EXECUTE_READWRITE`

## Working Set

工作集是进程当前驻留在物理内存中的页面集合。进程虚拟地址空间很大，不代表所有页面都真的在物理内存里。

## Memory-mapped file

内存映射文件把文件内容映射进内存地址空间。读写内存，就像读写文件的一部分。

常用于：

- 大文件处理
- 共享内存 IPC
- 数据库/搜索引擎/缓存系统

