# C++：网络、HTTP、Winsock

Winsock2、WinHTTP、WinINet、IP Helper、DNS、RPC、HTTP Server API。

## 常用 API 速查

| API | 用途 | 库/模块 |
|---|---|---|
| `WSAStartup/WSACleanup` | 初始化 Winsock | Ws2_32.lib |
| `socket/bind/listen/accept/connect` | TCP/UDP socket | Ws2_32.lib |
| `send/recv/closesocket` | 收发与关闭 | Ws2_32.lib |
| `getaddrinfo/freeaddrinfo` | DNS/地址解析 | Ws2_32.lib |
| `WinHttpOpen/WinHttpConnect` | WinHTTP 会话 | Winhttp.lib |
| `WinHttpOpenRequest/WinHttpSendRequest` | HTTP 请求 | Winhttp.lib |
| `InternetOpen/InternetOpenUrl` | WinINet 简易访问 | Wininet.lib |
| `GetAdaptersAddresses` | 网卡地址 | Iphlpapi.lib |
| `GetIfTable2` | 接口表 | Iphlpapi.lib |
| `DnsQuery` | DNS 查询 | Dnsapi.lib |

## 使用要点

- 查看每个 API 文档中的 **Header / Library / DLL / Minimum supported client**。
- C++ 默认建议走 Unicode 版本，避免 ANSI 代码页问题。
- 明确资源释放函数：`CloseHandle`、`LocalFree`、`CoTaskMemFree`、`DeleteObject`、`Release` 等不可混用。
- 失败处理：Win32 多数用 `GetLastError()`，COM 多数用 `HRESULT`。

## 案例

```cpp
#include <winsock2.h>
#include <ws2tcpip.h>
#include <iostream>
#pragma comment(lib, "Ws2_32.lib")

int main() {
    WSADATA wsa{};
    if (WSAStartup(MAKEWORD(2,2), &wsa) != 0) return 1;
    addrinfo hints{}; hints.ai_family = AF_UNSPEC; hints.ai_socktype = SOCK_STREAM;
    addrinfo* result = nullptr;
    if (getaddrinfo("example.com", "80", &hints, &result) == 0) {
        std::cout << "resolved
";
        freeaddrinfo(result);
    }
    WSACleanup();
}
```

## 官方入口

- Windows API index：https://learn.microsoft.com/windows/win32/apiindex/windows-api-list
- Win32 API reference：https://learn.microsoft.com/windows/win32/api/
