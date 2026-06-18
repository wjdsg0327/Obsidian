# Go：网络、HTTP、Winsock

Winsock2、WinHTTP、WinINet、IP Helper、DNS、RPC、HTTP Server API。

## Go 调用方式

| API | 用途 | 库/模块 |
|---|---|---|
| `WSAStartup/WSACleanup` | 初始化 Winsock | golang.org/x/sys/windows 或 syscall/DLL |
| `socket/bind/listen/accept/connect` | TCP/UDP socket | golang.org/x/sys/windows 或 syscall/DLL |
| `send/recv/closesocket` | 收发与关闭 | golang.org/x/sys/windows 或 syscall/DLL |
| `getaddrinfo/freeaddrinfo` | DNS/地址解析 | golang.org/x/sys/windows 或 syscall/DLL |
| `WinHttpOpen/WinHttpConnect` | WinHTTP 会话 | go-ole 或 syscall + unsafe |
| `WinHttpOpenRequest/WinHttpSendRequest` | HTTP 请求 | golang.org/x/sys/windows 或 syscall/DLL |
| `InternetOpen/InternetOpenUrl` | WinINet 简易访问 | 标准库 net/http 优先；必要时 DLL |
| `GetAdaptersAddresses` | 网卡地址 | golang.org/x/sys/windows 或 syscall/DLL |
| `GetIfTable2` | 接口表 | golang.org/x/sys/windows 或 syscall/DLL |
| `DnsQuery` | DNS 查询 | golang.org/x/sys/windows 或 syscall/DLL |

## 使用要点

- 优先使用 Go 标准库；标准库不够时再用 `golang.org/x/sys/windows`。
- Windows 字符串通常是 UTF-16：`windows.UTF16PtrFromString`、`windows.UTF16ToString`。
- 结构体传参要注意字段类型、对齐、`unsafe.Sizeof`。
- `HANDLE`、证书上下文、COM 内存、注册表 Key 等都要使用对应释放函数。
- `syscall.NewLazyDLL` 适合补充未封装 API，但要按官方签名核对参数宽度。

## 案例

```go
// Go 网络优先使用标准库 net/http、net。
// 只有需要 WinHTTP/WinINet 代理、系统证书、IE 配置等 Windows 特性时再调 WinHTTP API。
package main

import (
    "fmt"
    "net"
)

func main() {
    addrs, _ := net.LookupHost("example.com")
    fmt.Println(addrs)
}
```

## 官方入口

- Go x/sys/windows：https://pkg.go.dev/golang.org/x/sys/windows
- Windows API index：https://learn.microsoft.com/windows/win32/apiindex/windows-api-list
