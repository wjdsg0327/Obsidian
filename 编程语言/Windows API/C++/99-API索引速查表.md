# API 索引速查表

## 高频 API 按场景

| 场景 | C++ API | Go 建议 |
|---|---|---|
| 打开文件 | `CreateFileW` | 标准库 `os`；底层用 `windows.CreateFile` |
| 读写文件 | `ReadFile` / `WriteFile` | `os.File`；底层用 `windows.ReadFile/WriteFile` |
| 关闭句柄 | `CloseHandle` | `windows.CloseHandle` |
| 创建进程 | `CreateProcessW` | `os/exec`；底层用 `windows.CreateProcess` |
| 等待对象 | `WaitForSingleObject` | `windows.WaitForSingleObject` |
| 注册表 | `RegOpenKeyExW` / `RegQueryValueExW` | `golang.org/x/sys/windows/registry` |
| 服务 | `OpenSCManager` / `CreateService` | `x/sys/windows/svc` 或 `windows.*Service*` |
| 权限 | `OpenProcessToken` / `AdjustTokenPrivileges` | `windows.OpenProcessToken` 等 |
| 网络 Socket | Winsock2 | Go 标准库 `net` 优先 |
| HTTP | WinHTTP / WinINet | Go `net/http` 优先；必要时 DLL |
| COM | `CoInitializeEx` / `CoCreateInstance` | 第三方 `go-ole` 或手写 syscall |
| 错误文本 | `GetLastError` / `FormatMessage` | `syscall.Errno` / `windows.FormatMessage` |
| 枚举进程 | PSAPI / ToolHelp | `windows.EnumProcesses` 或 ToolHelp 封装 |
| 设备控制 | `DeviceIoControl` | `windows.DeviceIoControl` |

## 头文件/库速查

| 模块 | 常见头文件 | 常见库 |
|---|---|---|
| 基础/进程/文件 | `windows.h` | `Kernel32.lib` |
| 窗口/消息 | `windows.h` | `User32.lib` |
| GDI | `wingdi.h` | `Gdi32.lib` |
| 服务/注册表/安全 | `windows.h`、`aclapi.h` | `Advapi32.lib` |
| Winsock | `winsock2.h`、`ws2tcpip.h` | `Ws2_32.lib` |
| WinHTTP | `winhttp.h` | `Winhttp.lib` |
| Shell | `shlobj.h`、`shellapi.h` | `Shell32.lib` |
| COM | `objbase.h` | `Ole32.lib` |
| SetupAPI | `setupapi.h` | `Setupapi.lib` |
| ETW | `evntrace.h` | `Advapi32.lib` |
