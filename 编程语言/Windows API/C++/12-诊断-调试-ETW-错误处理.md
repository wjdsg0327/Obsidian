# C++：诊断、调试、ETW、错误处理

GetLastError、异常处理、Debug API、WER、ETW、性能日志、PSAPI。

## 常用 API 速查

| API | 用途 | 库/模块 |
|---|---|---|
| `GetLastError` | 错误码 | Kernel32.lib |
| `FormatMessage` | 错误文本 | Kernel32.lib |
| `SetUnhandledExceptionFilter` | 未处理异常 | Kernel32.lib |
| `RaiseException` | 抛异常 | Kernel32.lib |
| `OutputDebugString` | 调试输出 | Kernel32.lib |
| `DebugActiveProcess` | 调试进程 | Kernel32.lib |
| `MiniDumpWriteDump` | 生成 dump | Dbghelp.lib |
| `EnumProcesses` | PSAPI 枚举进程 | Psapi.lib |
| `StartTrace/EnableTraceEx2` | ETW 控制 | Advapi32.lib |
| `EventRegister/EventWrite` | ETW Provider | Advapi32.lib |

## 使用要点

- 查看每个 API 文档中的 **Header / Library / DLL / Minimum supported client**。
- C++ 默认建议走 Unicode 版本，避免 ANSI 代码页问题。
- 明确资源释放函数：`CloseHandle`、`LocalFree`、`CoTaskMemFree`、`DeleteObject`、`Release` 等不可混用。
- 失败处理：Win32 多数用 `GetLastError()`，COM 多数用 `HRESULT`。

## 案例

```cpp
// 伪代码骨架：查看具体 API 文档确认头文件、库、返回值和错误规则。
// 1. 准备输入结构体
// 2. 调用 Windows API
// 3. 判断返回值
// 4. 失败时 GetLastError / HRESULT
// 5. 释放句柄、内存、COM 对象或 GDI 对象
```

## 官方入口

- Windows API index：https://learn.microsoft.com/windows/win32/apiindex/windows-api-list
- Win32 API reference：https://learn.microsoft.com/windows/win32/api/
