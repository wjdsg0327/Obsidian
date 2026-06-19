---
title: Win32 API 官方入口速览
date: 2026-06-19
tags: [Windows, Win32, API, Cpp]
source: https://learn.microsoft.com/en-us/windows/win32/api/
---

# Win32 API 官方入口速览

> 来源：Microsoft Learn - Programming reference for the Win32 API  
> 链接：https://learn.microsoft.com/en-us/windows/win32/api/

## 说明

Win32 API 是 Windows 原生桌面开发的底层接口集合，主要面向 C / C++ 程序员，也常被 C#、Python、Rust、Go 等语言通过 FFI / PInvoke / ctypes 等方式间接调用。

这个页面不是单个函数文档，而是 **Win32 API 按技术领域划分的总入口**。下面按官网页面列出的每个 API/技术分类做简要说明。

---

## 入门

### Build desktop Windows apps using the Win32 API
Windows 桌面应用开发总览，介绍如何使用 Win32 API 构建传统 Windows 桌面程序，包括窗口、消息循环、资源、控件等基础概念。

### Development tools
Windows 开发工具下载入口，通常包括 Windows SDK、Visual Studio、调试工具等。

### Windows code samples
微软官方 Windows 示例代码集合，可用于学习具体 API 的调用方式。

---

# User Interface and desktop：用户界面与桌面

### Windows controls
Windows 标准控件 API，例如按钮、编辑框、列表框、树形控件、进度条等，用于构建传统桌面界面。

### Windows and messages
窗口与消息机制 API，是 Win32 GUI 编程核心，包括创建窗口、窗口过程、消息循环、键盘鼠标消息等。

### Menus and other resources
菜单、图标、光标、字符串表、对话框资源等相关 API，用于管理应用程序资源。

### Windows Shell
Windows Shell 相关 API，包括文件资源管理器交互、快捷方式、任务栏、通知区域、文件关联、Shell 扩展等。

### Accessibility features
辅助功能 API，用于屏幕阅读器、无障碍访问、UI 自动化、键盘导航等场景。

### Internationalization
国际化与本地化 API，用于处理语言、区域设置、字符编码、日期时间格式、排序规则等。

---

# Graphics and gaming：图形与游戏

### Direct2D
2D 硬件加速图形 API，用于绘制文本、几何图形、位图、矢量图形等。

### Direct3D 11 Graphics
Direct3D 11 图形 API，用于 3D 渲染、游戏、图形引擎和 GPU 加速应用。

### Direct3D 12 Graphics
Direct3D 12 图形 API，提供更底层、更高性能的 GPU 控制能力，适合游戏引擎和高性能渲染。

### DirectML
Direct Machine Learning API，用于在 Windows 上通过 GPU 加速机器学习推理。

### DXGI
DirectX Graphics Infrastructure，负责图形适配器、显示输出、交换链、帧缓冲等底层图形基础设施。

### Windows GDI
传统 2D 图形接口，用于绘制文本、线条、矩形、位图等，是老式 Win32 程序常用绘图系统。

### GDI+
GDI 的增强版，支持更现代的 2D 图形能力，如抗锯齿、渐变、图像格式处理等。

### Windows Imaging Component
WIC 图像处理 API，用于读取、写入、解码、编码 PNG、JPEG、TIFF、GIF 等图片格式。

---

# Audio and video：音频与视频

### Microsoft Media Foundation
现代 Windows 多媒体框架，用于音视频采集、播放、编码、解码、转码和流媒体处理。

### Windows Multimedia
较早期的多媒体 API，包括 wave、midi、计时器、简单音频播放等传统接口。

### XAudio2 APIs
低延迟音频引擎 API，常用于游戏音效、实时音频混音和音频处理。

### Core Audio APIs
Windows 核心音频 API，用于控制音频设备、音量、音频会话、音频流等。

### Audio Devices DDI Reference
音频设备驱动接口文档，主要面向驱动开发人员。

---

# Data access and storage：数据访问与存储

### Data access and storage
文件系统、文件 I/O、目录、卷、磁盘、路径、文件属性等基础存储 API。

### Backup
备份相关 API，用于文件备份、恢复、卷影复制等场景。

### Background Intelligent Transfer Service
BITS 后台智能传输服务 API，用于可靠的后台下载/上传，Windows Update 等服务也会使用类似机制。

### Cloud Filter API
云文件占位与同步 API，常用于 OneDrive 一类的云盘集成，实现“按需下载”的文件体验。

### Data Exchange
数据交换 API，包括剪贴板、拖放、动态数据交换等，用于应用之间传递数据。

### Structured storage
结构化存储 API，可以把多个数据流组织在一个复合文件中，类似文件中的小型文件系统。

### Virtual Storage
虚拟存储相关 API，用于虚拟磁盘、存储抽象或兼容性存储场景。

---

# Devices：设备与驱动

### Kernel-mode driver reference
内核模式驱动开发参考，面向 Windows 驱动开发，包括内核对象、IRP、驱动入口等。

### Device and driver installation reference
设备与驱动安装 API，用于设备枚举、驱动安装、设备属性管理等。

### Storage driver DDI reference
存储驱动接口文档，面向磁盘、卷、存储控制器等驱动开发。

### USB driver reference
USB 驱动开发参考，包括 USB 设备、端点、传输、描述符等接口。

### Display devices reference
显示设备相关接口，面向显卡、显示器、显示驱动、显示配置等。

### Human Interface Devices reference
HID 人机接口设备 API，包括键盘、鼠标、手柄、触摸板、自定义 HID 设备等。

### Print DDI reference
打印驱动接口文档，面向打印机驱动、打印管线和打印设备开发。

---

# Networking and internet：网络与互联网

### Bluetooth
蓝牙 API，用于发现蓝牙设备、连接、服务枚举、数据通信等。

### HTTP Server API
Windows 内置 HTTP 服务器 API，可让程序直接处理 HTTP 请求，底层常见组件是 HTTP.sys。

### IP Helper
IP Helper API，用于查询和配置网络适配器、IP 地址、路由表、ARP 表、TCP/UDP 连接等。

### Network Management
网络管理 API，用于用户、共享、服务器、工作站、网络资源等传统 Windows 网络管理功能。

### Remote Procedure Call (RPC)
远程过程调用 API，用于进程间或跨机器调用服务接口，是很多 Windows 系统服务的基础机制。

### Windows HTTP Services (WinHTTP)
WinHTTP API，适合服务端程序或后台程序进行 HTTP/HTTPS 请求。

### Windows Internet
WinINet API，适合客户端应用进行 HTTP、FTP 等网络访问，和用户登录态、代理设置关系更紧密。

### Windows Sockets 2
Winsock 2 网络套接字 API，是 Windows 上 TCP/UDP 网络编程的基础接口。

---

# Security and identity：安全与身份

### Security and identity
Windows 安全模型相关 API，包括访问令牌、ACL、SID、权限、身份验证、加密等。

### Network Access Protection
网络访问保护相关 API，属于较老的企业网络健康检查与准入控制技术。

### Network Policy Server
网络策略服务器 API，主要与 RADIUS、企业网络认证和访问策略有关。

### Parental controls
家长控制 API，用于限制用户访问内容、应用、时间等，偏系统策略管理。

### Windows Biometric Framework
Windows 生物识别框架 API，用于指纹、人脸等生物识别设备与认证流程。

### TPM Base Services
TPM 基础服务 API，用于访问可信平台模块，常见于密钥保护、安全启动、设备身份等场景。

---

# Diagnostics：诊断与监控

### Event tracing
ETW 事件跟踪 API，用于高性能系统日志、性能分析、内核/应用事件采集。

### Network Diagnostics Framework
网络诊断框架 API，用于检测和修复网络连接问题。

### Performance counters
性能计数器 API，用于读取 CPU、内存、磁盘、网络、应用程序等性能指标。

### Application Recovery and Restart
应用恢复与重启 API，用于程序崩溃或系统重启后自动恢复状态。

### TraceLogging
TraceLogging 是 ETW 的一种现代封装方式，用于更方便地写入结构化跟踪日志。

### Windows Event Collector
Windows 事件收集器 API，用于集中收集多台机器的事件日志。

### Windows Error Reporting
Windows 错误报告 API，用于收集崩溃、挂起、错误诊断信息并上报或本地分析。

---

# Application installation：应用安装与部署

### Application installation and servicing
应用安装、维护、修复、卸载相关 API，涉及安装包、组件、补丁等。

### Packaging and deployment Windows 10 apps
Windows 10 应用打包与部署 API，主要与 MSIX / APPX 包、应用安装和部署相关。

### Developer licensing
开发者许可相关 API，用于开发、测试或部署许可相关场景。

### Restart Manager
重启管理器 API，用于安装/更新时检测占用文件的程序，并协调关闭与重启。

---

# System services：系统服务与基础组件

### Component Object Model (COM)
COM 组件对象模型 API，是 Windows 重要的二进制组件技术，用于跨语言、跨进程组件调用。

### COM+ (Component Services)
COM+ 组件服务，提供事务、对象池、安全、分布式组件等企业级能力。

### Microsoft Interface Definition Language (MIDL)
MIDL 接口定义语言，用于定义 COM/RPC 接口并生成代理、存根等代码。

### Compression API
Windows 压缩 API，用于数据压缩和解压缩。

### Activity Coordinator
活动协调器 API，用于协调后台活动与系统资源策略，减少对前台体验的影响。

### Hardware Requirement Evaluator (HWREQCHK)
硬件需求评估相关 API，用于检查设备是否满足特定硬件要求。

---

## 怎么查具体函数

如果要查某一个具体 Win32 函数，可以直接在 Microsoft Learn 搜索函数名，例如：

- `CreateWindowExW`：创建窗口
- `GetMessageW`：从线程消息队列取消息
- `DispatchMessageW`：分发消息到窗口过程
- `CreateFileW`：打开或创建文件、设备、管道等对象
- `ReadFile` / `WriteFile`：读取和写入文件或设备
- `CreateProcessW`：创建新进程
- `RegOpenKeyExW`：打开注册表键
- `ShellExecuteW`：通过 Shell 打开文件、URL 或启动程序
- `CoInitializeEx`：初始化 COM 运行环境

---

## 学习建议

如果只是想入门 Win32 编程，建议按这个顺序看：

1. Windows and messages：窗口和消息机制
2. Windows controls：标准控件
3. Menus and other resources：菜单和资源
4. Data access and storage：文件读写
5. Windows GDI / Direct2D：绘图
6. Windows Sockets 2 / WinHTTP：网络
7. COM：组件与系统接口
8. Security and identity：权限与安全

