# Go：进程、线程、DLL、同步

进程线程创建、句柄、作业对象、线程池、DLL、互斥量、事件、信号量。

## Go 调用方式

| API | 用途 | 库/模块 |
|---|---|---|
| `CreateProcess` | 创建进程 | golang.org/x/sys/windows 或 syscall/DLL |
| `OpenProcess/TerminateProcess` | 打开/终止进程 | golang.org/x/sys/windows 或 syscall/DLL |
| `GetCurrentProcessId/GetCurrentThreadId` | 当前 ID | golang.org/x/sys/windows 或 syscall/DLL |
| `CreateThread` | 创建线程 | golang.org/x/sys/windows 或 syscall/DLL |
| `WaitForSingleObject/WaitForMultipleObjects` | 等待对象 | golang.org/x/sys/windows 或 syscall/DLL |
| `CreateMutex/CreateEvent/CreateSemaphore` | 同步对象 | golang.org/x/sys/windows 或 syscall/DLL |
| `InitializeCriticalSection` | 临界区 | golang.org/x/sys/windows 或 syscall/DLL |
| `LoadLibrary/GetProcAddress/FreeLibrary` | 动态加载 DLL | golang.org/x/sys/windows 或 syscall/DLL |
| `CreateJobObject/AssignProcessToJobObject` | 作业对象 | golang.org/x/sys/windows 或 syscall/DLL |
| `CreateThreadpoolWork` | 线程池 | golang.org/x/sys/windows 或 syscall/DLL |

## 使用要点

- 优先使用 Go 标准库；标准库不够时再用 `golang.org/x/sys/windows`。
- Windows 字符串通常是 UTF-16：`windows.UTF16PtrFromString`、`windows.UTF16ToString`。
- 结构体传参要注意字段类型、对齐、`unsafe.Sizeof`。
- `HANDLE`、证书上下文、COM 内存、注册表 Key 等都要使用对应释放函数。
- `syscall.NewLazyDLL` 适合补充未封装 API，但要按官方签名核对参数宽度。

## 案例

```go
//go:build windows
package main

import (
    "unsafe"
    "golang.org/x/sys/windows"
)

func main() {
    app, _ := windows.UTF16PtrFromString("C:\Windows\System32\notepad.exe")
    var si windows.StartupInfo
    var pi windows.ProcessInformation
    si.Cb = uint32(unsafe.Sizeof(si))
    err := windows.CreateProcess(app, nil, nil, nil, false, 0, nil, nil, &si, &pi)
    if err != nil { panic(err) }
    windows.CloseHandle(pi.Thread)
    windows.CloseHandle(pi.Process)
}
```

## 官方入口

- Go x/sys/windows：https://pkg.go.dev/golang.org/x/sys/windows
- Windows API index：https://learn.microsoft.com/windows/win32/apiindex/windows-api-list
