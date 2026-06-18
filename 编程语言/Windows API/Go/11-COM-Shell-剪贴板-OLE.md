# Go：COM、Shell、剪贴板、OLE

COM 初始化、接口、Shell、属性系统、快捷方式、剪贴板、拖放、OLE。

## Go 调用方式

| API | 用途 | 库/模块 |
|---|---|---|
| `CoInitializeEx/CoUninitialize` | COM 初始化/反初始化 | go-ole 或 syscall + unsafe |
| `CoCreateInstance` | 创建 COM 对象 | go-ole 或 syscall + unsafe |
| `QueryInterface/AddRef/Release` | COM 生命周期 | golang.org/x/sys/windows 或 syscall/DLL |
| `CoTaskMemFree` | 释放 COM 分配内存 | go-ole 或 syscall + unsafe |
| `SHGetKnownFolderPath` | 已知文件夹 | golang.org/x/sys/windows 或 syscall/DLL |
| `ShellExecuteEx` | Shell 执行 | go-ole 或 syscall + unsafe |
| `IShellLink` | 快捷方式 | go-ole 或 syscall + unsafe |
| `OpenClipboard/GetClipboardData` | 剪贴板读取 | golang.org/x/sys/windows 或 syscall/DLL |
| `SetClipboardData/CloseClipboard` | 剪贴板写入 | golang.org/x/sys/windows 或 syscall/DLL |
| `DoDragDrop` | OLE 拖放 | golang.org/x/sys/windows 或 syscall/DLL |

## 使用要点

- 优先使用 Go 标准库；标准库不够时再用 `golang.org/x/sys/windows`。
- Windows 字符串通常是 UTF-16：`windows.UTF16PtrFromString`、`windows.UTF16ToString`。
- 结构体传参要注意字段类型、对齐、`unsafe.Sizeof`。
- `HANDLE`、证书上下文、COM 内存、注册表 Key 等都要使用对应释放函数。
- `syscall.NewLazyDLL` 适合补充未封装 API，但要按官方签名核对参数宽度。

## 案例

```go
//go:build windows
// COM/Shell 在 Go 中通常要借助 go-ole 或手写 syscall。
// 基本原则：CoInitializeEx -> CoCreateInstance -> 调方法 -> Release -> CoUninitialize。
```

## 官方入口

- Go x/sys/windows：https://pkg.go.dev/golang.org/x/sys/windows
- Windows API index：https://learn.microsoft.com/windows/win32/apiindex/windows-api-list
