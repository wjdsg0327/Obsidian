# Go：设备、驱动、打印、HID

设备管理、SetupAPI、DeviceIoControl、串口、HID、打印、WIA、传感器。

## Go 调用方式

| API | 用途 | 库/模块 |
|---|---|---|
| `SetupDiGetClassDevs` | 枚举设备 | golang.org/x/sys/windows 或 syscall/DLL |
| `SetupDiEnumDeviceInfo` | 设备信息 | golang.org/x/sys/windows 或 syscall/DLL |
| `CreateFile("\\.\Device")` | 打开设备/串口 | golang.org/x/sys/windows 或 syscall/DLL |
| `DeviceIoControl` | 发送 IOCTL | go-ole 或 syscall + unsafe |
| `GetCommState/SetCommState` | 串口配置 | go-ole 或 syscall + unsafe |
| `HidD_GetAttributes` | HID 属性 | golang.org/x/sys/windows 或 syscall/DLL |
| `OpenPrinter/StartDocPrinter` | 打印 | golang.org/x/sys/windows 或 syscall/DLL |
| `EnumPrinters` | 枚举打印机 | golang.org/x/sys/windows 或 syscall/DLL |
| `CoCreateInstance(WIA)` | 图像采集 | go-ole 或 syscall + unsafe |
| `RegisterDeviceNotification` | 设备插拔通知 | golang.org/x/sys/windows 或 syscall/DLL |

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
