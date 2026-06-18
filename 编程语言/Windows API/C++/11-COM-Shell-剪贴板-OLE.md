# C++：COM、Shell、剪贴板、OLE

COM 初始化、接口、Shell、属性系统、快捷方式、剪贴板、拖放、OLE。

## 常用 API 速查

| API | 用途 | 库/模块 |
|---|---|---|
| `CoInitializeEx/CoUninitialize` | COM 初始化/反初始化 | Ole32.lib |
| `CoCreateInstance` | 创建 COM 对象 | Ole32.lib |
| `QueryInterface/AddRef/Release` | COM 生命周期 | Ole32.lib |
| `CoTaskMemFree` | 释放 COM 分配内存 | Ole32.lib |
| `SHGetKnownFolderPath` | 已知文件夹 | Shell32.lib |
| `ShellExecuteEx` | Shell 执行 | Shell32.lib |
| `IShellLink` | 快捷方式 | Shell32/Ole32 |
| `OpenClipboard/GetClipboardData` | 剪贴板读取 | User32.lib |
| `SetClipboardData/CloseClipboard` | 剪贴板写入 | User32.lib |
| `DoDragDrop` | OLE 拖放 | Ole32.lib |

## 使用要点

- 查看每个 API 文档中的 **Header / Library / DLL / Minimum supported client**。
- C++ 默认建议走 Unicode 版本，避免 ANSI 代码页问题。
- 明确资源释放函数：`CloseHandle`、`LocalFree`、`CoTaskMemFree`、`DeleteObject`、`Release` 等不可混用。
- 失败处理：Win32 多数用 `GetLastError()`，COM 多数用 `HRESULT`。

## 案例

```cpp
#include <windows.h>
#include <shlobj.h>
#include <iostream>
#pragma comment(lib, "Ole32.lib")
#pragma comment(lib, "Shell32.lib")

int wmain() {
    CoInitializeEx(nullptr, COINIT_APARTMENTTHREADED);
    PWSTR path = nullptr;
    if (SUCCEEDED(SHGetKnownFolderPath(FOLDERID_Documents, 0, nullptr, &path))) {
        std::wcout << path << L"
";
        CoTaskMemFree(path);
    }
    CoUninitialize();
}
```

## 官方入口

- Windows API index：https://learn.microsoft.com/windows/win32/apiindex/windows-api-list
- Win32 API reference：https://learn.microsoft.com/windows/win32/api/
