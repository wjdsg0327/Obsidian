# C++：进程、线程、DLL、同步

进程线程创建、句柄、作业对象、线程池、DLL、互斥量、事件、信号量。

## 常用 API 速查

| API | 用途 | 库/模块 |
|---|---|---|
| `CreateProcess` | 创建进程 | Kernel32.lib |
| `OpenProcess/TerminateProcess` | 打开/终止进程 | Kernel32.lib |
| `GetCurrentProcessId/GetCurrentThreadId` | 当前 ID | Kernel32.lib |
| `CreateThread` | 创建线程 | Kernel32.lib |
| `WaitForSingleObject/WaitForMultipleObjects` | 等待对象 | Kernel32.lib |
| `CreateMutex/CreateEvent/CreateSemaphore` | 同步对象 | Kernel32.lib |
| `InitializeCriticalSection` | 临界区 | Kernel32.lib |
| `LoadLibrary/GetProcAddress/FreeLibrary` | 动态加载 DLL | Kernel32.lib |
| `CreateJobObject/AssignProcessToJobObject` | 作业对象 | Kernel32.lib |
| `CreateThreadpoolWork` | 线程池 | Kernel32.lib |

## 使用要点

- 查看每个 API 文档中的 **Header / Library / DLL / Minimum supported client**。
- C++ 默认建议走 Unicode 版本，避免 ANSI 代码页问题。
- 明确资源释放函数：`CloseHandle`、`LocalFree`、`CoTaskMemFree`、`DeleteObject`、`Release` 等不可混用。
- 失败处理：Win32 多数用 `GetLastError()`，COM 多数用 `HRESULT`。

## 案例

```cpp
#include <windows.h>
#include <iostream>

int wmain() {
    STARTUPINFOW si{ sizeof(si) };
    PROCESS_INFORMATION pi{};
    wchar_t cmd[] = L"notepad.exe";
    if (!CreateProcessW(nullptr, cmd, nullptr, nullptr, FALSE, 0, nullptr, nullptr, &si, &pi)) {
        std::wcerr << L"CreateProcess failed: " << GetLastError() << L"
";
        return 1;
    }
    CloseHandle(pi.hThread);
    CloseHandle(pi.hProcess);
}
```

## 官方入口

- Windows API index：https://learn.microsoft.com/windows/win32/apiindex/windows-api-list
- Win32 API reference：https://learn.microsoft.com/windows/win32/api/
