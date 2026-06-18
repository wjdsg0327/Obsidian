# C++：设备、驱动、打印、HID

设备管理、SetupAPI、DeviceIoControl、串口、HID、打印、WIA、传感器。

## 常用 API 速查

| API | 用途 | 库/模块 |
|---|---|---|
| `SetupDiGetClassDevs` | 枚举设备 | Setupapi.lib |
| `SetupDiEnumDeviceInfo` | 设备信息 | Setupapi.lib |
| `CreateFile("\\.\Device")` | 打开设备/串口 | Kernel32.lib |
| `DeviceIoControl` | 发送 IOCTL | Kernel32.lib |
| `GetCommState/SetCommState` | 串口配置 | Kernel32.lib |
| `HidD_GetAttributes` | HID 属性 | Hid.lib |
| `OpenPrinter/StartDocPrinter` | 打印 | Winspool.lib |
| `EnumPrinters` | 枚举打印机 | Winspool.lib |
| `CoCreateInstance(WIA)` | 图像采集 | Wiaaut/Wiaservc |
| `RegisterDeviceNotification` | 设备插拔通知 | User32.lib |

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
