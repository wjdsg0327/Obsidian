# 学习资料与工具清单

> 以官方文档和体系化教材为主。  
> 网络文章质量参差不齐：优先用来对照，不优先用来“抄现成攻击器”。

---

## 1. Microsoft 官方文档（必读）

### 进程与线程

- [OpenProcess](https://learn.microsoft.com/windows/win32/api/processthreadsapi/nf-processthreadsapi-openprocess)
- [CreateRemoteThread](https://learn.microsoft.com/windows/win32/api/processthreadsapi/nf-processthreadsapi-createremotethread)
- [GetExitCodeThread](https://learn.microsoft.com/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodethread)
- [WaitForSingleObject](https://learn.microsoft.com/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject)

### 内存

- [VirtualAlloc](https://learn.microsoft.com/windows/win32/api/memoryapi/nf-memoryapi-virtualalloc)
- [VirtualAllocEx](https://learn.microsoft.com/windows/win32/api/memoryapi/nf-memoryapi-virtualallocex)
- [VirtualProtect / VirtualProtectEx](https://learn.microsoft.com/windows/win32/api/memoryapi/nf-memoryapi-virtualprotect)
- [WriteProcessMemory](https://learn.microsoft.com/windows/win32/api/memoryapi/nf-memoryapi-writeprocessmemory)
- [ReadProcessMemory](https://learn.microsoft.com/windows/win32/api/memoryapi/nf-memoryapi-readprocessmemory)
- [Memory Protection Constants](https://learn.microsoft.com/windows/win32/memory/memory-protection-constants)
- [Memory Management 总览](https://learn.microsoft.com/windows/win32/memory/memory-management)

### 模块加载

- [LoadLibraryW](https://learn.microsoft.com/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibraryw)
- [GetModuleHandle](https://learn.microsoft.com/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew)
- [GetProcAddress](https://learn.microsoft.com/windows/win32/api/libloaderapi/nf-libloaderapi-getprocaddress)
- [DllMain](https://learn.microsoft.com/windows/win32/dlls/dllmain)

### 权限与安全

- [Process Security and Access Rights](https://learn.microsoft.com/windows/win32/procthread/process-security-and-access-rights)
- [Access Rights for Process Objects 相关说明](https://learn.microsoft.com/windows/win32/procthread/process-security-and-access-rights)

---

## 2. 书籍（体系化）

| 书名 | 为什么看 | 怎么用 |
|------|----------|--------|
| *Windows Internals* | 进程/线程/内存管理权威 | 精读相关章节，不求一次读完 |
| *Practical Malware Analysis* | 动态分析与常见恶意手法 | 用实验理解注入在样本中的样子 |
| *Windows via C/C++*（或同类系统编程书） | API 使用直觉 | 对照 API 调法 |
| *The Rootkit Arsenal* / 现代替代分析资料 | 扩展视野 | 进阶选读，注意年代 |

中文资料可作辅助，但涉及 API 细节请回英文官方文档核对。

---

## 3. Go 相关

- [Go 官方安装](https://go.dev/dl/)
- [golang.org/x/sys/windows](https://pkg.go.dev/golang.org/x/sys/windows)
- [Go unsafe 包说明](https://pkg.go.dev/unsafe)
- 搜索关键词（自学时）：
  - `go VirtualAllocEx`
  - `go CreateRemoteThread`
  - `golang windows process injection lab`

注意：GitHub 上大量“injection”仓库包含攻击向完整代码与免杀讨论。  
**学习阶段请只吸收 API 用法与错误处理，不要把目标设为复制攻击链。**

---

## 4. 工具

### 必装级

| 工具 | 用途 |
|------|------|
| Process Hacker / System Informer | 看进程、内存、模块、线程 |
| x64dbg | 用户态调试 |
| Go toolchain | 编译实验程序 |

### 很有用

| 工具 | 用途 |
|------|------|
| Process Explorer | 轻量进程观察 |
| VMMap | 虚拟内存布局 |
| API Monitor | API 调用轨迹 |
| PE-bear | PE 结构查看 |
| Sysmon | 系统事件采集（防御练习） |
| DebugView | 内核/用户调试输出（若你用 OutputDebugString） |

### 编译 DLL

- Visual Studio / MSVC Build Tools  
- 或 MinGW-w64  

---

## 5. 课程与专题（可选）

选择标准：

- 明确教学/防御/认证方向  
- 有实验环境说明  
- 不主打“免杀变现”

可关注方向：

- Windows 恶意软件分析课  
- 蓝队 / 威胁狩猎中的 Process Injection 专题  
- OSCP/类似认证中与 Windows 相关的基础（注意范围与授权）  

MITRE ATT&CK：

- [Process Injection 技术条目](https://attack.mitre.org/techniques/T1055/)

用它建立“分类地图”，而不是当攻击手册。

---

## 6. 建议阅读顺序（两周版）

### 第 1～3 天

1. 本教程 `01-foundations.md`  
2. MSDN：`VirtualAlloc` / 内存保护常量  
3. 做 phase1 本进程实验  

### 第 4～7 天

1. MSDN：`OpenProcess` / `WriteProcessMemory`  
2. 做 phase2 远程读写  
3. Process Hacker 内存观察练习  

### 第 8～12 天

1. MSDN：`CreateRemoteThread` / `LoadLibraryW`  
2. 做 DLL 注入  
3. x64dbg 附加观察 `LoadLibraryW`  

### 第 13～14 天

1. MITRE T1055 分类速览  
2. 写防御笔记  
3. 复盘所有错误记录  

---

## 7. 记笔记的方式

推荐三本笔记：

1. **API 卡片**：每个 API 一页（参数、权限、常见错误）  
2. **实验日志**：每次命令、PID、地址、结果  
3. **对照表**：注入技术 vs 检测点  

模板：

```markdown
### API: WriteProcessMemory
- 需要权限：
- 输入：
- 成功返回：
- 我踩过的坑：
- 相关实验：
```

---

## 8. 资料质量自检

看到一篇“Go 注入教程”时先问：

1. 是否只鼓励攻击未授权目标？  
2. 是否跳过权限与架构解释直接给 shellcode？  
3. 是否完全不谈如何验证与排错？  
4. 是否把免杀当核心卖点？  

若是，降权处理：只挑其中 API 名称做索引，回到官方文档重学。
