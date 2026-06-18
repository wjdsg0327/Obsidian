# Go：诊断、调试、ETW、错误处理

GetLastError、异常处理、Debug API、WER、ETW、性能日志、PSAPI。

## Go 调用方式

| API | 用途 | 库/模块 |
|---|---|---|
| `GetLastError` | 错误码 | golang.org/x/sys/windows 或 syscall/DLL |
| `FormatMessage` | 错误文本 | golang.org/x/sys/windows 或 syscall/DLL |
| `SetUnhandledExceptionFilter` | 未处理异常 | golang.org/x/sys/windows 或 syscall/DLL |
| `RaiseException` | 抛异常 | golang.org/x/sys/windows 或 syscall/DLL |
| `OutputDebugString` | 调试输出 | golang.org/x/sys/windows 或 syscall/DLL |
| `DebugActiveProcess` | 调试进程 | golang.org/x/sys/windows 或 syscall/DLL |
| `MiniDumpWriteDump` | 生成 dump | golang.org/x/sys/windows 或 syscall/DLL |
| `EnumProcesses` | PSAPI 枚举进程 | golang.org/x/sys/windows 或 syscall/DLL |
| `StartTrace/EnableTraceEx2` | ETW 控制 | golang.org/x/sys/windows 或 syscall/DLL |
| `EventRegister/EventWrite` | ETW Provider | golang.org/x/sys/windows 或 syscall/DLL |

## 使用要点

- 优先使用 Go 标准库；标准库不够时再用 `golang.org/x/sys/windows`。
- Windows 字符串通常是 UTF-16：`windows.UTF16PtrFromString`、`windows.UTF16ToString`。
- 结构体传参要注意字段类型、对齐、`unsafe.Sizeof`。
- `HANDLE`、证书上下文、COM 内存、注册表 Key 等都要使用对应释放函数。
- `syscall.NewLazyDLL` 适合补充未封装 API，但要按官方签名核对参数宽度。

## 案例

```go
// Go 调 Windows API 通用骨架：
// 1. 优先查 golang.org/x/sys/windows 是否已经封装
// 2. 字符串转 UTF-16：windows.UTF16PtrFromString
// 3. 调用 API，检查 err / 返回值
// 4. defer windows.CloseHandle 或调用对应释放函数
```

## 官方入口

- Go x/sys/windows：https://pkg.go.dev/golang.org/x/sys/windows
- Windows API index：https://learn.microsoft.com/windows/win32/apiindex/windows-api-list
