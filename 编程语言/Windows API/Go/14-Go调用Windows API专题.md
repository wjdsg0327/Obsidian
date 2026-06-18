# Windows API Go 总览与学习路线

Go 调 Windows API 常用三层方式：

1. 优先用标准库：`os`、`net`、`syscall` 已经封装了很多平台差异。
2. 需要 Windows 专有能力时，用 `golang.org/x/sys/windows`。
3. 仍未封装时，用 `windows.NewLazySystemDLL` / `NewProc` 或自己写 syscall 封装。

## Go 调用规则

- Windows API 字符串大多是 UTF-16：用 `windows.UTF16PtrFromString`。
- `HANDLE` 对应 `windows.Handle`，多数要 `windows.CloseHandle`。
- 返回值与错误需要按 API 文档判断：有些返回 0 表示失败，有些返回非 0 表示失败。
- 涉及结构体、指针、回调时需要 `unsafe`，必须保证内存布局与 Windows SDK 一致。
- 能用 `x/sys/windows` 已封装函数就不要手写 DLL 调用。

## Go 最小 API 调用案例：MessageBoxW

```go
//go:build windows
package main

import (
    "syscall"
    "unsafe"
)

var (
    user32     = syscall.NewLazyDLL("user32.dll")
    messageBox = user32.NewProc("MessageBoxW")
)

func utf16Ptr(s string) uintptr {
    p, _ := syscall.UTF16PtrFromString(s)
    return uintptr(unsafe.Pointer(p))
}

func main() {
    messageBox.Call(
        0,
        utf16Ptr("你好，Windows API"),
        utf16Ptr("Go 调用 Win32"),
        0,
    )
}
```

## Go 使用 x/sys/windows 案例：读取计算机名

```go
//go:build windows
package main

import (
    "fmt"
    "golang.org/x/sys/windows"
)

func main() {
    buf := make([]uint16, windows.MAX_COMPUTERNAME_LENGTH+1)
    n := uint32(len(buf))
    if err := windows.GetComputerName(&buf[0], &n); err != nil {
        panic(err)
    }
    fmt.Println(windows.UTF16ToString(buf[:n]))
}
```


## 三种封装层级

### 1. 标准库优先

- 文件：`os`
- 网络：`net`、`net/http`
- 进程：`os/exec`
- 时间：`time`

### 2. x/sys/windows

适合：句柄、注册表、服务、Token、Winsock、证书、文件底层能力。

### 3. LazyDLL 手写调用

适合：API 尚未被 x/sys/windows 封装，或者临时验证。

```go
var (
    kernel32 = windows.NewLazySystemDLL("kernel32.dll")
    procGetTickCount64 = kernel32.NewProc("GetTickCount64")
)

func uptimeMs() uint64 {
    r1, _, _ := procGetTickCount64.Call()
    return uint64(r1)
}
```

## 常见坑

- Go GC 与 Windows 回调：回调函数、buffer、结构体必须保证调用期间仍然存活。
- `uintptr` 不是普通整数，跨调用保存指针要小心。
- `ERROR_INSUFFICIENT_BUFFER` 很常见：先调一次取长度，再分配 buffer。
- Windows API 往往需要管理员权限或特定 privilege，不是代码错。
