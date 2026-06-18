# kernel32.dll 封装进度

当前 `kernel32Dll` 包覆盖了项目里最常用的一批 kernel32.dll API：

1. 句柄和错误码：`CloseHandle`、`DuplicateHandle`、`GetLastError` 等。
2. 文件和目录：`CreateFileW`、`ReadFile`、`WriteFile`、`FindFirstFileW`、`CopyFileW`、`MoveFileExW` 等。
3. 进程和线程：`CreateProcessW`、`OpenProcess`、`CreateThread`、`WaitForSingleObject` 等。
4. 内存：`VirtualAlloc`、`VirtualProtect`、`HeapAlloc`、`GlobalAlloc` 等。
5. 模块和环境：`LoadLibraryW`、`GetProcAddress`、`GetModuleFileNameW`、环境变量读写等。
6. 同步对象：Mutex、Event、Semaphore、WaitableTimer。
7. 时间和系统信息：`QueryPerformanceCounter`、`GetSystemTime`、`GetSystemInfo`、`GetSystemPowerStatus` 等。

仍建议后续按需补充的部分：

1. Console 控制台 API：`AllocConsole`、`GetStdHandle`、`ReadConsoleW`、`WriteConsoleW`、`SetConsoleTextAttribute` 等。
2. Pipe / Named Pipe：`CreatePipe`、`CreateNamedPipeW`、`ConnectNamedPipe`、`PeekNamedPipe` 等。
3. File Mapping：`CreateFileMappingW`、`MapViewOfFile`、`UnmapViewOfFile`、`OpenFileMappingW`。
4. ToolHelp 快照：`CreateToolhelp32Snapshot`、`Process32FirstW`、`Thread32First`、`Module32FirstW`。
5. IOCP：`CreateIoCompletionPort`、`GetQueuedCompletionStatus`、`PostQueuedCompletionStatus`。
6. TLS / FLS：`TlsAlloc`、`TlsGetValue`、`FlsAlloc` 等。
7. Fiber：`CreateFiber`、`SwitchToFiber`、`DeleteFiber`。
8. ActCtx / SxS：`CreateActCtxW`、`ActivateActCtx`、`DeactivateActCtx`。
9. Locale / NLS：`GetLocaleInfoEx` 等很多函数已迁移到 KernelBase，但项目里仍常见。
10. Debug API：`OutputDebugStringW`、`IsDebuggerPresent`、`DebugBreak`、`WaitForDebugEvent`。

说明：`kernel32.dll` 与 `KernelBase.dll`、`ntdll.dll`、`advapi32.dll`、`psapi.dll` 等边界有不少历史兼容关系。这个包定位为“项目常用 kernel32.dll Go 封装”，不是逐导出符号全集。
