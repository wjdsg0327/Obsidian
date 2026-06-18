# Windows API 知识库索引


## 主要参考来源

- Microsoft Learn：Windows API index — https://learn.microsoft.com/windows/win32/apiindex/windows-api-list
- Microsoft Learn：Programming reference for the Win32 API — https://learn.microsoft.com/windows/win32/api/
- Microsoft Learn：Get Started with Win32 and C++ — https://learn.microsoft.com/windows/win32/learnwin32/learn-to-program-for-windows
- Go：golang.org/x/sys/windows — https://pkg.go.dev/golang.org/x/sys/windows

> 说明：Windows API 体量极大，微软官方按技术域、头文件、函数三种维度组织。这里整理的是“学习与工程实践用”的全局索引 + 常用 API + 可运行案例骨架；需要函数级全集时，建议后续按本目录逐类扩展。

## 目录

- [[C++/01-总览与学习路线|C++：总览与学习路线]]
- [[Go/01-总览与学习路线|Go：总览与学习路线]]
- [[C++/02-窗口-消息-控件|C++：窗口、消息、控件]]
- [[Go/02-窗口-消息-控件|Go：窗口、消息、控件]]
- [[C++/03-文件-目录-磁盘-IO|C++：文件、目录、磁盘、I/O]]
- [[Go/03-文件-目录-磁盘-IO|Go：文件、目录、磁盘、I/O]]
- [[C++/04-进程-线程-DLL-同步|C++：进程、线程、DLL、同步]]
- [[Go/04-进程-线程-DLL-同步|Go：进程、线程、DLL、同步]]
- [[C++/05-内存-系统信息-注册表|C++：内存、系统信息、注册表]]
- [[Go/05-内存-系统信息-注册表|Go：内存、系统信息、注册表]]
- [[C++/05A-内存管理专题|C++：内存管理专题]]
- [[Go/05A-内存管理专题|Go：内存管理专题]]
- [[C++/06-安全-身份-权限-加密|C++：安全、身份、权限、加密]]
- [[Go/06-安全-身份-权限-加密|Go：安全、身份、权限、加密]]
- [[C++/07-网络-WinHTTP-WinINet-Winsock|C++：网络、HTTP、Winsock]]
- [[Go/07-网络-WinHTTP-WinINet-Winsock|Go：网络、HTTP、Winsock]]
- [[C++/08-服务-计划任务-系统管理|C++：服务、计划任务、系统管理]]
- [[Go/08-服务-计划任务-系统管理|Go：服务、计划任务、系统管理]]
- [[C++/09-图形-音视频-多媒体|C++：图形、音视频、多媒体]]
- [[Go/09-图形-音视频-多媒体|Go：图形、音视频、多媒体]]
- [[C++/10-设备-驱动-打印-HID|C++：设备、驱动、打印、HID]]
- [[Go/10-设备-驱动-打印-HID|Go：设备、驱动、打印、HID]]
- [[C++/11-COM-Shell-剪贴板-OLE|C++：COM、Shell、剪贴板、OLE]]
- [[Go/11-COM-Shell-剪贴板-OLE|Go：COM、Shell、剪贴板、OLE]]
- [[C++/12-诊断-调试-ETW-错误处理|C++：诊断、调试、ETW、错误处理]]
- [[Go/12-诊断-调试-ETW-错误处理|Go：诊断、调试、ETW、错误处理]]
- [[C++/13-安装-部署-包管理|C++：安装、部署、包管理]]
- [[Go/13-安装-部署-包管理|Go：安装、部署、包管理]]
- [[C++/14-C++工程模板与编译|C++：工程模板与编译]]
- [[Go/14-Go调用Windows API专题|Go：Go 调用 Windows API 专题]]
- [[C++/99-API索引速查表|C++：API 索引速查表]]
- [[Go/99-API索引速查表|Go：API 索引速查表]]


## 两套目录的区别

- **C++**：贴近微软官方文档，包含头文件、链接库、生命周期和原生案例。
- **Go**：强调 `golang.org/x/sys/windows`、`syscall`、UTF-16、`unsafe`、DLL 动态调用和 Go 风格封装。

## 后续扩展建议

如果要继续补到“函数级全集”，建议按以下顺序扩：

1. 文件/进程/注册表/服务/网络：最常用，收益最高。
2. COM/Shell/安全：工程中常见，但坑多。
3. 图形、多媒体、设备驱动：按具体项目需要扩。
