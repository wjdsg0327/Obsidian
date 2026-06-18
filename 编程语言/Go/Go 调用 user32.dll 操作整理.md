---
title: Go 调用 user32.dll 操作整理
date: 2026-06-18
tags:
  - Go
  - Windows
  - Win32
  - user32.dll
  - Windows-API
---

# Go 调用 user32.dll 操作整理

> 目标：整理 Go 语言中通过 Win32 API 调用 `user32.dll` 能做的主要操作、常用函数、调用方式、注意事项与示例代码。  
> 范围：`user32.dll` 主要负责 Windows 图形界面、窗口、消息、输入、剪贴板、菜单、光标、显示、DPI、Hook 等用户交互相关能力。  
> 说明：`user32.dll` 的 API 非常多，本文按功能域整理“可做什么”和常用函数。完整函数清单以 Microsoft Learn 的 `winuser.h` / Win32 API 文档为准。

---

## 1. Go 调用 user32.dll 的基本方式

Go 在 Windows 下调用 DLL 常见有三种方式：

1. 标准库 `syscall`
2. `golang.org/x/sys/windows`
3. 第三方封装，例如 `github.com/lxn/win`、`github.com/AllenDang/w32` 等

推荐新代码优先使用：

```go
import "golang.org/x/sys/windows"
```

不过很多简单示例仍用 `syscall.NewLazyDLL`，写起来直观。

### 1.1 基础模板：加载 user32.dll

```go
package main

import (
    "syscall"
    "unsafe"
)

var (
    user32              = syscall.NewLazyDLL("user32.dll")
    procMessageBoxW     = user32.NewProc("MessageBoxW")
)

func utf16Ptr(s string) *uint16 {
    p, _ := syscall.UTF16PtrFromString(s)
    return p
}

func main() {
    procMessageBoxW.Call(
        0,
        uintptr(unsafe.Pointer(utf16Ptr("你好，user32.dll"))),
        uintptr(unsafe.Pointer(utf16Ptr("Go 调用 Win32"))),
        0,
    )
}
```

### 1.2 A / W 函数区别

Windows API 中经常有：

- `xxxA`：ANSI 版本
- `xxxW`：Unicode / UTF-16 版本

Go 字符串是 UTF-8，调用 Windows API 时通常转成 UTF-16，并优先调用 `xxxW`。

```go
MessageBoxW
FindWindowW
SendMessageW
CreateWindowExW
```

### 1.3 常见 Win32 类型在 Go 里的映射

| Win32 类型 | Go 常见表示 |
|---|---|
| `HWND` | `uintptr` / `windows.Handle` |
| `HINSTANCE` | `uintptr` |
| `HMENU` | `uintptr` |
| `HICON` | `uintptr` |
| `HCURSOR` | `uintptr` |
| `HHOOK` | `uintptr` |
| `LPWSTR` | `*uint16` / `uintptr(unsafe.Pointer(ptr))` |
| `LPCWSTR` | `*uint16` / `uintptr(unsafe.Pointer(ptr))` |
| `BOOL` | 返回值 `0=false`，非 `0=true` |
| `UINT` | `uint32` / `uintptr` |
| `WPARAM` | `uintptr` |
| `LPARAM` | `uintptr` / `unsafe.Pointer` |
| `LRESULT` | `uintptr` |

---

## 2. user32.dll 能做什么：功能总览

`user32.dll` 覆盖的操作大致包括：

1. 弹窗与对话框
2. 窗口查找、枚举、创建、显示、隐藏、移动、置顶
3. 窗口标题、类名、样式、属性读取与修改
4. 窗口消息发送、投递、消息循环
5. 鼠标、键盘输入读取与模拟
6. 全局快捷键注册
7. 剪贴板读写与监听
8. 光标、图标、鼠标位置控制
9. 菜单创建与修改
10. 系统参数读取与修改
11. 屏幕、显示器、分辨率相关操作
12. DPI 感知与缩放相关操作
13. Hook：键盘、鼠标、消息、窗口事件监听
14. 窗口站、桌面、前台窗口、焦点控制
15. 原始输入 Raw Input
16. 触摸、指针、手势相关输入
17. Accessibility / WinEvent 事件监听
18. 计时器 Timer
19. 多显示器枚举
20. IME / 键盘布局相关操作

---

## 3. 弹窗与简单 UI

### 常用函数

| 函数 | 作用 |
|---|---|
| `MessageBoxW` | 显示消息框 |
| `MessageBoxExW` | 指定语言环境的消息框 |
| `MessageBeep` | 播放系统提示音 |

### 示例：MessageBoxW

```go
package main

import (
    "syscall"
    "unsafe"
)

var (
    user32          = syscall.NewLazyDLL("user32.dll")
    messageBoxW     = user32.NewProc("MessageBoxW")
)

func strPtr(s string) uintptr {
    p, _ := syscall.UTF16PtrFromString(s)
    return uintptr(unsafe.Pointer(p))
}

func main() {
    messageBoxW.Call(
        0,
        strPtr("这是内容"),
        strPtr("这是标题"),
        0x00000040, // MB_ICONINFORMATION
    )
}
```

---

## 4. 窗口查找、枚举与信息读取

### 常用函数

| 函数 | 作用 |
|---|---|
| `FindWindowW` | 按窗口类名或标题查找顶层窗口 |
| `FindWindowExW` | 查找子窗口 |
| `EnumWindows` | 枚举所有顶层窗口 |
| `EnumChildWindows` | 枚举某窗口的子窗口 |
| `GetForegroundWindow` | 获取当前前台窗口 |
| `GetActiveWindow` | 获取当前线程活动窗口 |
| `GetDesktopWindow` | 获取桌面窗口句柄 |
| `GetShellWindow` | 获取 Shell 桌面窗口 |
| `GetParent` | 获取父窗口 |
| `GetAncestor` | 获取祖先窗口 |
| `GetWindow` | 获取相关窗口，如 owner、next、prev |
| `WindowFromPoint` | 通过屏幕坐标获取窗口 |
| `ChildWindowFromPoint` | 通过父窗口内坐标获取子窗口 |
| `RealChildWindowFromPoint` | 更精确地通过坐标获取子窗口 |
| `GetWindowTextW` | 获取窗口标题 |
| `GetWindowTextLengthW` | 获取窗口标题长度 |
| `GetClassNameW` | 获取窗口类名 |
| `GetWindowThreadProcessId` | 获取窗口所属线程与进程 ID |
| `IsWindow` | 判断句柄是否为有效窗口 |
| `IsWindowVisible` | 判断窗口是否可见 |
| `IsWindowEnabled` | 判断窗口是否可用 |
| `IsIconic` | 判断窗口是否最小化 |
| `IsZoomed` | 判断窗口是否最大化 |

### 示例：获取当前前台窗口标题

```go
package main

import (
    "fmt"
    "syscall"
    "unsafe"
)

var (
    user32                   = syscall.NewLazyDLL("user32.dll")
    getForegroundWindow       = user32.NewProc("GetForegroundWindow")
    getWindowTextW            = user32.NewProc("GetWindowTextW")
)

func main() {
    hwnd, _, _ := getForegroundWindow.Call()
    if hwnd == 0 {
        fmt.Println("没有前台窗口")
        return
    }

    buf := make([]uint16, 256)
    getWindowTextW.Call(
        hwnd,
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(len(buf)),
    )

    fmt.Println("前台窗口标题：", syscall.UTF16ToString(buf))
}
```

### 示例：按标题查找窗口

```go
var findWindowW = user32.NewProc("FindWindowW")

func FindWindowByTitle(title string) uintptr {
    p, _ := syscall.UTF16PtrFromString(title)
    hwnd, _, _ := findWindowW.Call(
        0,
        uintptr(unsafe.Pointer(p)),
    )
    return hwnd
}
```

---

## 5. 窗口显示、隐藏、移动与层级控制

### 常用函数

| 函数 | 作用 |
|---|---|
| `ShowWindow` | 显示、隐藏、最小化、最大化窗口 |
| `ShowWindowAsync` | 异步显示窗口 |
| `SetWindowPos` | 改变位置、大小、Z 顺序、置顶等 |
| `MoveWindow` | 移动并调整窗口大小 |
| `BringWindowToTop` | 将窗口放到 Z 顺序顶部 |
| `SetForegroundWindow` | 设置前台窗口 |
| `AllowSetForegroundWindow` | 允许指定进程设置前台窗口 |
| `SetActiveWindow` | 激活窗口 |
| `SetFocus` | 设置键盘焦点 |
| `CloseWindow` | 最小化窗口 |
| `OpenIcon` | 恢复最小化窗口 |
| `UpdateWindow` | 立即触发绘制更新 |
| `RedrawWindow` | 重绘窗口 |
| `InvalidateRect` | 使窗口区域无效，等待重绘 |
| `ValidateRect` | 验证窗口区域 |

### ShowWindow 常量

| 常量 | 值 | 作用 |
|---|---:|---|
| `SW_HIDE` | 0 | 隐藏 |
| `SW_SHOWNORMAL` | 1 | 正常显示 |
| `SW_SHOWMINIMIZED` | 2 | 最小化显示 |
| `SW_SHOWMAXIMIZED` | 3 | 最大化显示 |
| `SW_SHOWNOACTIVATE` | 4 | 显示但不激活 |
| `SW_SHOW` | 5 | 显示 |
| `SW_MINIMIZE` | 6 | 最小化 |
| `SW_SHOWMINNOACTIVE` | 7 | 最小化但不激活 |
| `SW_SHOWNA` | 8 | 显示但不激活 |
| `SW_RESTORE` | 9 | 还原 |

### 示例：隐藏 / 显示窗口

```go
var showWindow = user32.NewProc("ShowWindow")

const (
    SW_HIDE    = 0
    SW_SHOW    = 5
    SW_RESTORE = 9
)

func HideWindow(hwnd uintptr) {
    showWindow.Call(hwnd, SW_HIDE)
}

func ShowWindowNormal(hwnd uintptr) {
    showWindow.Call(hwnd, SW_RESTORE)
}
```

### 示例：移动窗口

```go
var moveWindow = user32.NewProc("MoveWindow")

func Move(hwnd uintptr, x, y, w, h int32) {
    moveWindow.Call(
        hwnd,
        uintptr(x),
        uintptr(y),
        uintptr(w),
        uintptr(h),
        1, // repaint=true
    )
}
```

### 示例：置顶窗口

```go
var setWindowPos = user32.NewProc("SetWindowPos")

const (
    HWND_TOPMOST   = ^uintptr(0) // -1
    HWND_NOTOPMOST = ^uintptr(1) // -2
    SWP_NOMOVE     = 0x0002
    SWP_NOSIZE     = 0x0001
)

func SetTopMost(hwnd uintptr, topmost bool) {
    insertAfter := HWND_NOTOPMOST
    if topmost {
        insertAfter = HWND_TOPMOST
    }
    setWindowPos.Call(hwnd, insertAfter, 0, 0, 0, 0, SWP_NOMOVE|SWP_NOSIZE)
}
```

---

## 6. 窗口创建与窗口过程

Go 也可以用 `user32.dll` 直接创建原生 Win32 窗口，但代码比较繁琐。

### 常用函数

| 函数 | 作用 |
|---|---|
| `RegisterClassW` / `RegisterClassExW` | 注册窗口类 |
| `UnregisterClassW` | 注销窗口类 |
| `CreateWindowExW` | 创建窗口 |
| `DestroyWindow` | 销毁窗口 |
| `DefWindowProcW` | 默认窗口过程 |
| `CallWindowProcW` | 调用窗口过程 |
| `SetWindowLongPtrW` | 修改窗口过程或窗口属性 |
| `GetWindowLongPtrW` | 读取窗口属性 |

### 基本流程

1. 定义窗口过程 `WndProc`
2. 注册窗口类 `RegisterClassExW`
3. 创建窗口 `CreateWindowExW`
4. 显示窗口 `ShowWindow`
5. 进入消息循环 `GetMessageW` → `TranslateMessage` → `DispatchMessageW`
6. 收到 `WM_DESTROY` 后 `PostQuitMessage`

### 重要消息

| 消息 | 作用 |
|---|---|
| `WM_CREATE` | 窗口创建 |
| `WM_DESTROY` | 窗口销毁 |
| `WM_CLOSE` | 请求关闭 |
| `WM_PAINT` | 绘制 |
| `WM_SIZE` | 大小改变 |
| `WM_MOVE` | 位置改变 |
| `WM_COMMAND` | 菜单、按钮等命令 |
| `WM_KEYDOWN` / `WM_KEYUP` | 键盘按下 / 抬起 |
| `WM_LBUTTONDOWN` / `WM_LBUTTONUP` | 鼠标左键 |
| `WM_MOUSEMOVE` | 鼠标移动 |
| `WM_TIMER` | 计时器 |
| `WM_HOTKEY` | 全局热键 |
| `WM_CLIPBOARDUPDATE` | 剪贴板变更 |

---

## 7. 消息发送、投递与消息循环

### 常用函数

| 函数 | 作用 |
|---|---|
| `SendMessageW` | 同步发送窗口消息 |
| `PostMessageW` | 异步投递窗口消息 |
| `PostThreadMessageW` | 向线程消息队列投递消息 |
| `SendNotifyMessageW` | 发送通知消息 |
| `SendMessageTimeoutW` | 带超时发送消息 |
| `BroadcastSystemMessageW` | 广播系统消息 |
| `RegisterWindowMessageW` | 注册唯一字符串消息 |
| `GetMessageW` | 获取消息，阻塞等待 |
| `PeekMessageW` | 非阻塞查看消息 |
| `TranslateMessage` | 翻译键盘消息 |
| `DispatchMessageW` | 分发消息到窗口过程 |
| `PostQuitMessage` | 请求退出消息循环 |
| `ReplyMessage` | 回复消息 |
| `WaitMessage` | 等待新消息 |

### SendMessage 和 PostMessage 区别

| 项目 | `SendMessageW` | `PostMessageW` |
|---|---|---|
| 调用方式 | 同步 | 异步 |
| 是否等待处理完成 | 等待 | 不等待 |
| 返回值 | 窗口过程返回值 | 是否投递成功 |
| 风险 | 目标窗口卡死会阻塞 | 消息可能晚点处理 |

### 示例：给窗口发送关闭消息

```go
var postMessageW = user32.NewProc("PostMessageW")

const WM_CLOSE = 0x0010

func CloseByMessage(hwnd uintptr) {
    postMessageW.Call(hwnd, WM_CLOSE, 0, 0)
}
```

---

## 8. 鼠标与键盘输入读取

### 常用函数

| 函数 | 作用 |
|---|---|
| `GetAsyncKeyState` | 获取按键当前状态 |
| `GetKeyState` | 获取按键状态 |
| `GetKeyboardState` | 获取键盘状态数组 |
| `SetKeyboardState` | 设置键盘状态 |
| `GetKeyboardLayout` | 获取键盘布局 |
| `ActivateKeyboardLayout` | 激活键盘布局 |
| `MapVirtualKeyW` | 虚拟键码与扫描码转换 |
| `ToUnicode` | 按键转换成 Unicode 字符 |
| `GetCursorPos` | 获取鼠标位置 |
| `SetCursorPos` | 设置鼠标位置 |
| `ScreenToClient` | 屏幕坐标转客户区坐标 |
| `ClientToScreen` | 客户区坐标转屏幕坐标 |
| `GetDoubleClickTime` | 获取双击时间 |
| `SetDoubleClickTime` | 设置双击时间 |
| `SwapMouseButton` | 左右键交换 |

### 示例：检测某个按键是否按下

```go
var getAsyncKeyState = user32.NewProc("GetAsyncKeyState")

const VK_SHIFT = 0x10

func IsKeyDown(vk int) bool {
    ret, _, _ := getAsyncKeyState.Call(uintptr(vk))
    return ret&0x8000 != 0
}

func main() {
    if IsKeyDown(VK_SHIFT) {
        println("Shift 正在按下")
    }
}
```

---

## 9. 模拟鼠标与键盘输入

### 常用函数

| 函数 | 作用 |
|---|---|
| `SendInput` | 推荐的键盘/鼠标输入模拟 API |
| `keybd_event` | 老 API，模拟键盘事件 |
| `mouse_event` | 老 API，模拟鼠标事件 |
| `BlockInput` | 阻止键鼠输入到应用程序 |

### 注意

- 推荐使用 `SendInput`，`keybd_event` 和 `mouse_event` 已较旧。
- 对管理员窗口、UAC、不同完整性级别进程可能无效。
- 自动化别人的软件时，要考虑焦点、权限、前台窗口限制。

### 示例：移动鼠标

```go
var setCursorPos = user32.NewProc("SetCursorPos")

func SetMousePos(x, y int) {
    setCursorPos.Call(uintptr(x), uintptr(y))
}
```

---

## 10. 全局快捷键

### 常用函数

| 函数 | 作用 |
|---|---|
| `RegisterHotKey` | 注册系统范围热键 |
| `UnregisterHotKey` | 注销热键 |

### 基本流程

1. 调用 `RegisterHotKey`
2. 启动消息循环
3. 监听 `WM_HOTKEY`
4. 退出时调用 `UnregisterHotKey`

### 常量

| 常量 | 值 | 作用 |
|---|---:|---|
| `MOD_ALT` | `0x0001` | Alt |
| `MOD_CONTROL` | `0x0002` | Ctrl |
| `MOD_SHIFT` | `0x0004` | Shift |
| `MOD_WIN` | `0x0008` | Windows 键 |
| `WM_HOTKEY` | `0x0312` | 热键消息 |

---

## 11. 剪贴板操作

### 常用函数

| 函数 | 作用 |
|---|---|
| `OpenClipboard` | 打开剪贴板 |
| `CloseClipboard` | 关闭剪贴板 |
| `EmptyClipboard` | 清空剪贴板 |
| `GetClipboardData` | 获取剪贴板数据 |
| `SetClipboardData` | 设置剪贴板数据 |
| `IsClipboardFormatAvailable` | 判断格式是否可用 |
| `EnumClipboardFormats` | 枚举剪贴板格式 |
| `RegisterClipboardFormatW` | 注册自定义剪贴板格式 |
| `GetClipboardFormatNameW` | 获取格式名称 |
| `CountClipboardFormats` | 获取格式数量 |
| `AddClipboardFormatListener` | 添加剪贴板监听窗口 |
| `RemoveClipboardFormatListener` | 移除剪贴板监听窗口 |
| `SetClipboardViewer` | 老式剪贴板查看器链 |
| `ChangeClipboardChain` | 从剪贴板查看器链移除 |

### 常见格式

| 格式 | 值 | 说明 |
|---|---:|---|
| `CF_TEXT` | 1 | ANSI 文本 |
| `CF_BITMAP` | 2 | 位图 |
| `CF_UNICODETEXT` | 13 | Unicode 文本 |
| `CF_HDROP` | 15 | 文件拖放列表 |

### 注意

`SetClipboardData` 通常需要配合 `kernel32.dll` 的 `GlobalAlloc`、`GlobalLock`、`GlobalUnlock` 使用。也就是说剪贴板操作不只涉及 `user32.dll`。

---

## 12. 光标、图标与资源

### 常用函数

| 函数 | 作用 |
|---|---|
| `LoadCursorW` | 加载系统光标或资源光标 |
| `SetCursor` | 设置当前光标 |
| `GetCursor` | 获取当前光标 |
| `ShowCursor` | 显示或隐藏光标计数 |
| `SetSystemCursor` | 设置系统光标 |
| `CopyCursor` | 复制光标 |
| `DestroyCursor` | 销毁光标 |
| `LoadIconW` | 加载图标 |
| `PrivateExtractIconsW` | 从文件提取图标 |
| `DrawIcon` | 绘制图标 |
| `DrawIconEx` | 绘制图标扩展版本 |
| `DestroyIcon` | 销毁图标 |

---

## 13. 菜单操作

### 常用函数

| 函数 | 作用 |
|---|---|
| `CreateMenu` | 创建菜单 |
| `CreatePopupMenu` | 创建弹出菜单 |
| `DestroyMenu` | 销毁菜单 |
| `AppendMenuW` | 追加菜单项 |
| `InsertMenuW` | 插入菜单项 |
| `ModifyMenuW` | 修改菜单项 |
| `RemoveMenu` | 移除菜单项 |
| `DeleteMenu` | 删除菜单项 |
| `GetMenu` | 获取窗口菜单 |
| `SetMenu` | 设置窗口菜单 |
| `GetSystemMenu` | 获取系统菜单 |
| `TrackPopupMenu` | 显示弹出菜单 |
| `DrawMenuBar` | 重绘菜单栏 |
| `CheckMenuItem` | 勾选菜单项 |
| `EnableMenuItem` | 启用 / 禁用菜单项 |
| `GetMenuItemInfoW` | 获取菜单项信息 |
| `SetMenuItemInfoW` | 设置菜单项信息 |

---

## 14. 窗口样式、属性与扩展属性

### 常用函数

| 函数 | 作用 |
|---|---|
| `GetWindowLongW` / `GetWindowLongPtrW` | 获取窗口长整型属性 |
| `SetWindowLongW` / `SetWindowLongPtrW` | 设置窗口长整型属性 |
| `GetClassLongW` / `GetClassLongPtrW` | 获取窗口类属性 |
| `SetClassLongW` / `SetClassLongPtrW` | 设置窗口类属性 |
| `GetPropW` | 获取窗口属性 |
| `SetPropW` | 设置窗口属性 |
| `RemovePropW` | 移除窗口属性 |
| `EnumPropsW` | 枚举窗口属性 |
| `GetLayeredWindowAttributes` | 获取分层窗口属性 |
| `SetLayeredWindowAttributes` | 设置窗口透明度 / 颜色键 |
| `UpdateLayeredWindow` | 更新分层窗口 |

### 常用样式

| 样式 | 说明 |
|---|---|
| `WS_VISIBLE` | 可见窗口 |
| `WS_CHILD` | 子窗口 |
| `WS_POPUP` | 弹出窗口 |
| `WS_OVERLAPPEDWINDOW` | 普通顶层窗口 |
| `WS_DISABLED` | 禁用窗口 |
| `WS_EX_TOPMOST` | 置顶扩展样式 |
| `WS_EX_LAYERED` | 分层窗口，支持透明 |
| `WS_EX_TOOLWINDOW` | 工具窗口 |

---

## 15. 窗口矩形、坐标与布局

### 常用函数

| 函数 | 作用 |
|---|---|
| `GetWindowRect` | 获取窗口外框矩形 |
| `GetClientRect` | 获取客户区矩形 |
| `AdjustWindowRect` | 根据客户区计算窗口外框 |
| `AdjustWindowRectEx` | 扩展版本 |
| `AdjustWindowRectExForDpi` | DPI 感知版本 |
| `MapWindowPoints` | 坐标在窗口间转换 |
| `ScreenToClient` | 屏幕坐标转客户区 |
| `ClientToScreen` | 客户区坐标转屏幕 |
| `GetSystemMetrics` | 获取屏幕、边框、滚动条等系统尺寸 |

### RECT 结构

```go
type RECT struct {
    Left   int32
    Top    int32
    Right  int32
    Bottom int32
}
```

### 示例：获取窗口矩形

```go
var getWindowRect = user32.NewProc("GetWindowRect")

type RECT struct {
    Left, Top, Right, Bottom int32
}

func GetRect(hwnd uintptr) RECT {
    var r RECT
    getWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&r)))
    return r
}
```

---

## 16. 屏幕、显示器与分辨率

### 常用函数

| 函数 | 作用 |
|---|---|
| `EnumDisplayMonitors` | 枚举显示器 |
| `GetMonitorInfoW` | 获取显示器信息 |
| `MonitorFromWindow` | 根据窗口获取显示器 |
| `MonitorFromPoint` | 根据点获取显示器 |
| `MonitorFromRect` | 根据矩形获取显示器 |
| `EnumDisplayDevicesW` | 枚举显示设备 |
| `EnumDisplaySettingsW` | 枚举显示模式 |
| `ChangeDisplaySettingsW` | 修改默认显示设备设置 |
| `ChangeDisplaySettingsExW` | 修改指定显示设备设置 |
| `GetDisplayConfigBufferSizes` | 显示配置相关，实际在 user32 |
| `QueryDisplayConfig` | 查询显示配置 |
| `SetDisplayConfig` | 设置显示配置 |

### 注意

修改分辨率、显示器配置会影响系统环境，属于高风险操作，建议只做读取，修改前需要用户确认。

---

## 17. DPI 与缩放

### 常用函数

| 函数 | 作用 |
|---|---|
| `SetProcessDPIAware` | 设置进程 DPI 感知，旧 API |
| `SetProcessDpiAwarenessContext` | 设置进程 DPI 感知上下文 |
| `GetProcessDpiAwarenessContext` | 获取进程 DPI 感知上下文 |
| `GetThreadDpiAwarenessContext` | 获取线程 DPI 感知上下文 |
| `SetThreadDpiAwarenessContext` | 设置线程 DPI 感知上下文 |
| `GetDpiForWindow` | 获取窗口 DPI |
| `GetDpiForSystem` | 获取系统 DPI |
| `GetSystemMetricsForDpi` | 按 DPI 获取系统尺寸 |
| `AdjustWindowRectExForDpi` | 按 DPI 调整窗口矩形 |
| `AreDpiAwarenessContextsEqual` | 比较 DPI 上下文 |
| `GetAwarenessFromDpiAwarenessContext` | 获取 DPI awareness 枚举 |

### 建议

如果写 GUI 程序，需要尽早设置 DPI 感知，否则窗口尺寸、鼠标坐标、截图坐标可能出现缩放偏差。

---

## 18. Hook：键盘、鼠标、消息与窗口事件

### 常用函数

| 函数 | 作用 |
|---|---|
| `SetWindowsHookExW` | 安装 Hook |
| `CallNextHookEx` | 调用下一个 Hook |
| `UnhookWindowsHookEx` | 卸载 Hook |
| `SetWinEventHook` | 安装 WinEvent 事件 Hook |
| `UnhookWinEvent` | 卸载 WinEvent Hook |
| `NotifyWinEvent` | 通知系统事件 |

### 常见 Hook 类型

| Hook | 说明 |
|---|---|
| `WH_KEYBOARD_LL` | 低级键盘 Hook，全局键盘监听常用 |
| `WH_MOUSE_LL` | 低级鼠标 Hook，全局鼠标监听常用 |
| `WH_KEYBOARD` | 线程键盘 Hook |
| `WH_MOUSE` | 线程鼠标 Hook |
| `WH_GETMESSAGE` | 监听消息队列取出的消息 |
| `WH_CALLWNDPROC` | 消息发送到窗口过程前监听 |
| `WH_CALLWNDPROCRET` | 窗口过程处理消息后监听 |
| `WH_CBT` | 窗口创建、激活、移动等 CBT 事件 |
| `WH_SHELL` | Shell 事件 |
| `WH_FOREGROUNDIDLE` | 前台线程空闲 |

### Go 中使用 Hook 的注意点

1. Hook 回调函数不能被 GC 回收，要保持引用。
2. 回调里不要做耗时操作，否则会卡系统输入。
3. 低级键鼠 Hook 要保持消息循环运行。
4. 退出时必须调用 `UnhookWindowsHookEx`。
5. 监听全局键盘/鼠标涉及隐私与安全，必须谨慎使用。

### Hook 相关结构

```go
type KBDLLHOOKSTRUCT struct {
    VkCode      uint32
    ScanCode    uint32
    Flags       uint32
    Time        uint32
    DwExtraInfo uintptr
}

type MSLLHOOKSTRUCT struct {
    Pt          POINT
    MouseData   uint32
    Flags       uint32
    Time        uint32
    DwExtraInfo uintptr
}

type POINT struct {
    X int32
    Y int32
}
```

---

## 19. WinEvent / 无障碍事件监听

`SetWinEventHook` 可以监听窗口标题变化、前台窗口变化、对象创建销毁等事件，常用于辅助工具、自动化、窗口监控。

### 常用事件

| 事件 | 说明 |
|---|---|
| `EVENT_SYSTEM_FOREGROUND` | 前台窗口变化 |
| `EVENT_OBJECT_CREATE` | 对象创建 |
| `EVENT_OBJECT_DESTROY` | 对象销毁 |
| `EVENT_OBJECT_SHOW` | 对象显示 |
| `EVENT_OBJECT_HIDE` | 对象隐藏 |
| `EVENT_OBJECT_NAMECHANGE` | 名称变化，如标题变化 |
| `EVENT_SYSTEM_MINIMIZESTART` | 开始最小化 |
| `EVENT_SYSTEM_MINIMIZEEND` | 最小化结束 |

---

## 20. Raw Input 原始输入

Raw Input 可以在不依赖窗口焦点的情况下接收底层输入设备数据。

### 常用函数

| 函数 | 作用 |
|---|---|
| `RegisterRawInputDevices` | 注册原始输入设备 |
| `GetRawInputData` | 获取原始输入数据 |
| `GetRawInputDeviceInfoW` | 获取设备信息 |
| `GetRawInputDeviceList` | 获取原始输入设备列表 |
| `DefRawInputProc` | 默认 Raw Input 处理 |

### 常见用途

- 游戏输入
- 多键盘 / 多鼠标区分
- 低延迟输入采集
- 后台输入监听

---

## 21. 触摸、指针与手势

### 常用函数

| 函数 | 作用 |
|---|---|
| `RegisterTouchWindow` | 注册触摸窗口 |
| `UnregisterTouchWindow` | 取消注册触摸窗口 |
| `GetTouchInputInfo` | 获取触摸输入信息 |
| `CloseTouchInputHandle` | 关闭触摸输入句柄 |
| `GetPointerInfo` | 获取指针信息 |
| `GetPointerTouchInfo` | 获取触摸指针信息 |
| `GetPointerPenInfo` | 获取笔输入信息 |
| `GetPointerFrameInfo` | 获取一帧指针信息 |
| `EnableMouseInPointer` | 启用鼠标作为指针输入 |
| `RegisterPointerInputTarget` | 注册指针输入目标 |
| `GetGestureInfo` | 获取手势信息 |
| `CloseGestureInfoHandle` | 关闭手势句柄 |
| `SetGestureConfig` | 设置手势配置 |
| `GetGestureConfig` | 获取手势配置 |

---

## 22. 计时器 Timer

### 常用函数

| 函数 | 作用 |
|---|---|
| `SetTimer` | 创建窗口计时器 |
| `KillTimer` | 删除计时器 |

计时器触发后，窗口会收到 `WM_TIMER` 消息。

---

## 23. 系统参数与环境设置

### 常用函数

| 函数 | 作用 |
|---|---|
| `SystemParametersInfoW` | 读取或设置大量系统参数 |
| `GetSystemMetrics` | 获取系统尺寸、配置、状态 |
| `GetSysColor` | 获取系统颜色 |
| `SetSysColors` | 设置系统颜色 |
| `LockWorkStation` | 锁定工作站 |
| `ExitWindowsEx` | 注销、关机、重启等 |
| `ShutdownBlockReasonCreate` | 创建阻止关机理由 |
| `ShutdownBlockReasonDestroy` | 移除阻止关机理由 |
| `ShutdownBlockReasonQuery` | 查询阻止关机理由 |

### SystemParametersInfoW 可做的事

- 获取 / 设置鼠标速度
- 获取 / 设置键盘延迟
- 获取 / 设置桌面壁纸
- 获取 / 设置屏保参数
- 获取 / 设置工作区大小
- 获取 / 设置动画、菜单淡入淡出等 UI 参数

### 注意

`SystemParametersInfoW` 和 `ExitWindowsEx` 会改变系统状态，应谨慎调用。

---

## 24. 桌面、窗口站与安全边界

### 常用函数

| 函数 | 作用 |
|---|---|
| `GetProcessWindowStation` | 获取当前进程窗口站 |
| `SetProcessWindowStation` | 设置进程窗口站 |
| `OpenWindowStationW` | 打开窗口站 |
| `CloseWindowStation` | 关闭窗口站 |
| `EnumWindowStationsW` | 枚举窗口站 |
| `GetThreadDesktop` | 获取线程桌面 |
| `SetThreadDesktop` | 设置线程桌面 |
| `OpenDesktopW` | 打开桌面 |
| `CreateDesktopW` | 创建桌面 |
| `CloseDesktop` | 关闭桌面 |
| `EnumDesktopsW` | 枚举桌面 |
| `SwitchDesktop` | 切换桌面 |

### 说明

这些 API 更偏系统级，常用于服务、远程桌面、隔离桌面、安全桌面等场景。普通应用很少直接使用。

---

## 25. 输入法与键盘布局

### 常用函数

| 函数 | 作用 |
|---|---|
| `GetKeyboardLayout` | 获取键盘布局 |
| `GetKeyboardLayoutList` | 获取键盘布局列表 |
| `ActivateKeyboardLayout` | 激活键盘布局 |
| `LoadKeyboardLayoutW` | 加载键盘布局 |
| `UnloadKeyboardLayout` | 卸载键盘布局 |
| `GetKeyboardLayoutNameW` | 获取键盘布局名称 |
| `VkKeyScanW` | 字符转虚拟键 |
| `MapVirtualKeyW` | 键码映射 |
| `ToUnicode` | 键盘状态转字符 |

---

## 26. 拖放、文件、Shell 交互相关

部分拖放能力在 `user32.dll`，部分在 `shell32.dll`、`ole32.dll`。

### user32.dll 相关函数

| 函数 | 作用 |
|---|---|
| `DragDetect` | 检测拖动开始 |
| `DrawFocusRect` | 绘制焦点矩形 |
| `DrawFrameControl` | 绘制标准控件外观 |
| `DrawTextW` | 绘制文本 |
| `TabbedTextOutW` | 绘制带 Tab 的文本 |

---

## 27. 绘制与窗口刷新相关

`user32.dll` 提供部分高层绘制函数，但真正的 GDI 绘图主要在 `gdi32.dll`。

### 常用函数

| 函数 | 作用 |
|---|---|
| `BeginPaint` | 开始绘制 |
| `EndPaint` | 结束绘制 |
| `GetDC` | 获取设备上下文 DC |
| `ReleaseDC` | 释放 DC |
| `GetWindowDC` | 获取整个窗口 DC |
| `DrawTextW` | 绘制文本 |
| `DrawEdge` | 绘制边框 |
| `DrawFrameControl` | 绘制按钮、滚动条等标准控件 |
| `FillRect` | 填充矩形，常与 GDI 画刷配合 |
| `FrameRect` | 绘制矩形边框 |
| `InvertRect` | 反色矩形区域 |

---

## 28. 滚动条与控件辅助

### 常用函数

| 函数 | 作用 |
|---|---|
| `ShowScrollBar` | 显示 / 隐藏滚动条 |
| `EnableScrollBar` | 启用 / 禁用滚动条 |
| `GetScrollInfo` | 获取滚动条信息 |
| `SetScrollInfo` | 设置滚动条信息 |
| `GetScrollPos` | 获取滚动位置 |
| `SetScrollPos` | 设置滚动位置 |
| `GetScrollRange` | 获取滚动范围 |
| `SetScrollRange` | 设置滚动范围 |
| `ScrollWindow` | 滚动窗口内容 |
| `ScrollWindowEx` | 扩展滚动窗口内容 |

---

## 29. 对话框相关

### 常用函数

| 函数 | 作用 |
|---|---|
| `DialogBoxParamW` | 创建模态对话框 |
| `CreateDialogParamW` | 创建非模态对话框 |
| `EndDialog` | 结束模态对话框 |
| `GetDlgItem` | 获取对话框控件 |
| `SetDlgItemTextW` | 设置控件文本 |
| `GetDlgItemTextW` | 获取控件文本 |
| `CheckDlgButton` | 设置复选框状态 |
| `IsDlgButtonChecked` | 获取复选框状态 |
| `SendDlgItemMessageW` | 给对话框控件发消息 |
| `MapDialogRect` | 对话框单位转像素 |

---

## 30. 加速键 Accelerator

### 常用函数

| 函数 | 作用 |
|---|---|
| `LoadAcceleratorsW` | 加载加速键表 |
| `CreateAcceleratorTableW` | 创建加速键表 |
| `DestroyAcceleratorTable` | 销毁加速键表 |
| `TranslateAcceleratorW` | 翻译加速键消息 |
| `CopyAcceleratorTableW` | 复制加速键表 |

---

## 31. Go 调用 user32.dll 时的坑

### 31.1 字符串编码

- 优先调用 `W` 版本函数。
- 使用 `syscall.UTF16PtrFromString` 转换。
- 注意指针生命周期，不要让临时对象过早释放。

### 31.2 32 位 / 64 位差异

- `HWND`、`WPARAM`、`LPARAM`、指针都应使用 `uintptr`。
- 不要把句柄硬转成 `int32`。
- `GetWindowLongPtrW` 在 64 位下尤其重要。

### 31.3 回调函数

涉及这些 API 时要特别小心：

- `EnumWindows`
- `EnumChildWindows`
- `SetWindowsHookExW`
- `SetWinEventHook`
- `RegisterClassExW` 的窗口过程

Go 中通常使用：

```go
syscall.NewCallback(fn)
```

回调函数必须匹配 Win32 调用约定和参数数量。

### 31.4 消息循环

使用窗口、热键、Hook、剪贴板监听时，经常需要消息循环，否则收不到消息。

基本形态：

```go
for {
    ret, _, _ := getMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
    if int32(ret) <= 0 {
        break
    }
    translateMessage.Call(uintptr(unsafe.Pointer(&msg)))
    dispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
}
```

### 31.5 权限与完整性级别

Windows 有 UIPI / 完整性级别限制：

- 普通权限进程难以操作管理员权限窗口。
- 给高权限窗口发消息可能失败。
- 模拟输入到安全桌面、UAC 提示通常不可行。

### 31.6 阻塞风险

`SendMessageW` 是同步的，如果目标窗口卡死，调用方也会卡住。跨进程自动化建议优先考虑：

- `PostMessageW`
- `SendMessageTimeoutW`

### 31.7 资源释放

需要配对释放的常见函数：

| 获取 / 创建 | 释放 |
|---|---|
| `SetWindowsHookExW` | `UnhookWindowsHookEx` |
| `SetWinEventHook` | `UnhookWinEvent` |
| `CreateMenu` | `DestroyMenu` |
| `CreateWindowExW` | `DestroyWindow` |
| `SetTimer` | `KillTimer` |
| `OpenClipboard` | `CloseClipboard` |
| `BeginPaint` | `EndPaint` |
| `GetDC` | `ReleaseDC` |
| `RegisterHotKey` | `UnregisterHotKey` |

---

## 32. 常用 user32.dll API 按场景速查

### 32.1 自动化已有窗口

| 目标 | 函数 |
|---|---|
| 找窗口 | `FindWindowW`, `EnumWindows` |
| 获取标题 | `GetWindowTextW` |
| 获取进程 ID | `GetWindowThreadProcessId` |
| 激活窗口 | `SetForegroundWindow`, `ShowWindow` |
| 移动窗口 | `MoveWindow`, `SetWindowPos` |
| 发送点击 / 按键消息 | `SendMessageW`, `PostMessageW` |
| 模拟真实输入 | `SendInput` |
| 关闭窗口 | `PostMessageW(hwnd, WM_CLOSE, 0, 0)` |

### 32.2 写 GUI 程序

| 目标 | 函数 |
|---|---|
| 注册窗口类 | `RegisterClassExW` |
| 创建窗口 | `CreateWindowExW` |
| 默认窗口过程 | `DefWindowProcW` |
| 消息循环 | `GetMessageW`, `TranslateMessage`, `DispatchMessageW` |
| 绘制 | `BeginPaint`, `EndPaint`, `DrawTextW` |
| 菜单 | `CreateMenu`, `AppendMenuW`, `SetMenu` |
| 定时器 | `SetTimer`, `KillTimer` |

### 32.3 做键鼠工具

| 目标 | 函数 |
|---|---|
| 读取键盘状态 | `GetAsyncKeyState`, `GetKeyState` |
| 读取鼠标位置 | `GetCursorPos` |
| 设置鼠标位置 | `SetCursorPos` |
| 模拟输入 | `SendInput` |
| 全局快捷键 | `RegisterHotKey`, `UnregisterHotKey` |
| 全局键盘 Hook | `SetWindowsHookExW(WH_KEYBOARD_LL)` |
| 全局鼠标 Hook | `SetWindowsHookExW(WH_MOUSE_LL)` |

### 32.4 做剪贴板工具

| 目标 | 函数 |
|---|---|
| 打开 / 关闭 | `OpenClipboard`, `CloseClipboard` |
| 读取 | `GetClipboardData` |
| 写入 | `SetClipboardData` |
| 清空 | `EmptyClipboard` |
| 监听变化 | `AddClipboardFormatListener`, `WM_CLIPBOARDUPDATE` |

### 32.5 做窗口监控工具

| 目标 | 函数 |
|---|---|
| 当前前台窗口 | `GetForegroundWindow` |
| 监听前台变化 | `SetWinEventHook(EVENT_SYSTEM_FOREGROUND)` |
| 监听标题变化 | `SetWinEventHook(EVENT_OBJECT_NAMECHANGE)` |
| 获取窗口进程 | `GetWindowThreadProcessId` |
| 获取窗口位置 | `GetWindowRect` |
| 判断窗口状态 | `IsWindowVisible`, `IsIconic`, `IsZoomed` |

---

## 33. 一个实用封装示例：窗口工具类

```go
package winuser

import (
    "syscall"
    "unsafe"
)

var (
    user32              = syscall.NewLazyDLL("user32.dll")
    procFindWindowW     = user32.NewProc("FindWindowW")
    procGetWindowTextW  = user32.NewProc("GetWindowTextW")
    procShowWindow      = user32.NewProc("ShowWindow")
    procSetForeground   = user32.NewProc("SetForegroundWindow")
    procPostMessageW    = user32.NewProc("PostMessageW")
)

const (
    SW_RESTORE = 9
    WM_CLOSE   = 0x0010
)

func utf16Ptr(s string) uintptr {
    p, _ := syscall.UTF16PtrFromString(s)
    return uintptr(unsafe.Pointer(p))
}

func FindWindow(title string) uintptr {
    hwnd, _, _ := procFindWindowW.Call(0, utf16Ptr(title))
    return hwnd
}

func GetWindowText(hwnd uintptr) string {
    buf := make([]uint16, 512)
    procGetWindowTextW.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
    return syscall.UTF16ToString(buf)
}

func ActivateWindow(hwnd uintptr) {
    procShowWindow.Call(hwnd, SW_RESTORE)
    procSetForeground.Call(hwnd)
}

func CloseWindow(hwnd uintptr) {
    procPostMessageW.Call(hwnd, WM_CLOSE, 0, 0)
}
```

---

## 34. 不建议滥用的操作

这些 API 虽然能用，但容易影响用户体验或触发安全问题：

| 操作 | 风险 |
|---|---|
| 全局键盘 Hook | 涉及隐私，容易被安全软件拦截 |
| `BlockInput` | 会阻断用户键鼠输入 |
| `SystemParametersInfoW` 修改系统参数 | 改变系统设置，影响全局 |
| `ChangeDisplaySettingsExW` | 改分辨率 / 显示器配置，可能黑屏 |
| `ExitWindowsEx` | 注销 / 关机 / 重启 |
| 操作管理员窗口 | 受权限限制，可能失败 |
| 给其它进程乱发消息 | 可能导致目标程序异常 |
| `SetWindowsHookExW` 线程 / 全局 Hook | 需要正确卸载，否则影响系统稳定性 |

---

## 35. user32.dll 操作分类清单

下面按功能域列出 user32.dll 常见操作，便于查漏：

### 窗口生命周期

- 注册窗口类：`RegisterClassW`, `RegisterClassExW`
- 注销窗口类：`UnregisterClassW`
- 创建窗口：`CreateWindowExW`
- 销毁窗口：`DestroyWindow`
- 默认窗口过程：`DefWindowProcW`
- 替换窗口过程：`SetWindowLongPtrW(GWLP_WNDPROC)`

### 窗口查找与枚举

- 查找窗口：`FindWindowW`, `FindWindowExW`
- 枚举窗口：`EnumWindows`, `EnumChildWindows`, `EnumThreadWindows`
- 获取前台窗口：`GetForegroundWindow`
- 获取桌面 / Shell：`GetDesktopWindow`, `GetShellWindow`
- 通过坐标找窗口：`WindowFromPoint`, `ChildWindowFromPoint`, `RealChildWindowFromPoint`

### 窗口信息

- 标题：`GetWindowTextW`, `SetWindowTextW`
- 类名：`GetClassNameW`
- 位置：`GetWindowRect`, `GetClientRect`
- 状态：`IsWindow`, `IsWindowVisible`, `IsWindowEnabled`, `IsIconic`, `IsZoomed`
- 进程线程：`GetWindowThreadProcessId`

### 窗口控制

- 显示隐藏：`ShowWindow`, `ShowWindowAsync`
- 移动调整：`MoveWindow`, `SetWindowPos`
- 激活焦点：`SetForegroundWindow`, `SetActiveWindow`, `SetFocus`
- 重绘：`InvalidateRect`, `ValidateRect`, `UpdateWindow`, `RedrawWindow`
- 关闭：发送 `WM_CLOSE` 或调用 `DestroyWindow`

### 消息机制

- 发送消息：`SendMessageW`, `SendMessageTimeoutW`
- 投递消息：`PostMessageW`, `PostThreadMessageW`
- 消息循环：`GetMessageW`, `PeekMessageW`, `TranslateMessage`, `DispatchMessageW`
- 退出消息循环：`PostQuitMessage`
- 注册消息：`RegisterWindowMessageW`

### 输入

- 键盘状态：`GetAsyncKeyState`, `GetKeyState`, `GetKeyboardState`
- 鼠标位置：`GetCursorPos`, `SetCursorPos`
- 输入模拟：`SendInput`, `keybd_event`, `mouse_event`
- 全局热键：`RegisterHotKey`, `UnregisterHotKey`
- 键盘布局：`GetKeyboardLayout`, `ActivateKeyboardLayout`, `LoadKeyboardLayoutW`

### Hook 与事件

- Windows Hook：`SetWindowsHookExW`, `CallNextHookEx`, `UnhookWindowsHookEx`
- WinEvent Hook：`SetWinEventHook`, `UnhookWinEvent`
- 常用 Hook：`WH_KEYBOARD_LL`, `WH_MOUSE_LL`, `WH_GETMESSAGE`, `WH_CALLWNDPROC`

### 剪贴板

- 打开关闭：`OpenClipboard`, `CloseClipboard`
- 读写：`GetClipboardData`, `SetClipboardData`
- 清空：`EmptyClipboard`
- 格式：`IsClipboardFormatAvailable`, `EnumClipboardFormats`, `RegisterClipboardFormatW`
- 监听：`AddClipboardFormatListener`, `RemoveClipboardFormatListener`

### 菜单

- 创建销毁：`CreateMenu`, `CreatePopupMenu`, `DestroyMenu`
- 修改：`AppendMenuW`, `InsertMenuW`, `ModifyMenuW`, `RemoveMenu`, `DeleteMenu`
- 绑定窗口：`GetMenu`, `SetMenu`, `DrawMenuBar`
- 弹出菜单：`TrackPopupMenu`

### 光标和图标

- 光标：`LoadCursorW`, `SetCursor`, `GetCursor`, `ShowCursor`, `SetSystemCursor`
- 图标：`LoadIconW`, `DrawIcon`, `DrawIconEx`, `DestroyIcon`

### 显示和 DPI

- 系统指标：`GetSystemMetrics`, `GetSystemMetricsForDpi`
- 显示器：`EnumDisplayMonitors`, `GetMonitorInfoW`, `MonitorFromWindow`
- 分辨率：`EnumDisplaySettingsW`, `ChangeDisplaySettingsExW`
- DPI：`SetProcessDpiAwarenessContext`, `GetDpiForWindow`, `GetDpiForSystem`

### 桌面与窗口站

- 窗口站：`OpenWindowStationW`, `GetProcessWindowStation`, `SetProcessWindowStation`
- 桌面：`OpenDesktopW`, `CreateDesktopW`, `GetThreadDesktop`, `SetThreadDesktop`, `SwitchDesktop`

### 系统参数

- 系统参数：`SystemParametersInfoW`
- 系统颜色：`GetSysColor`, `SetSysColors`
- 锁定 / 关机：`LockWorkStation`, `ExitWindowsEx`

---

## 36. 推荐学习路线

如果只是想用 Go 操作 Windows 窗口，建议按这个顺序学：

1. `MessageBoxW`：理解 DLL 调用、UTF-16 字符串、`uintptr`
2. `FindWindowW` + `GetWindowTextW`：理解窗口句柄
3. `ShowWindow` + `SetWindowPos`：控制窗口
4. `SendMessageW` / `PostMessageW`：理解 Win32 消息
5. `GetMessageW` 消息循环：理解 GUI 程序核心
6. `RegisterHotKey`：做全局快捷键
7. `OpenClipboard`：做剪贴板工具
8. `SetWindowsHookExW`：做键鼠监听，但要谨慎
9. `CreateWindowExW`：写完整 Win32 原生窗口

---

## 37. 参考资料

- Microsoft Learn：Winuser.h header / Win32 apps  
  https://learn.microsoft.com/windows/win32/api/winuser
- Microsoft Learn：Windows and Messages  
  https://learn.microsoft.com/windows/win32/winmsg/windows-and-messages
- Microsoft Learn：Keyboard and Mouse Input  
  https://learn.microsoft.com/windows/win32/inputdev/user-input
- Microsoft Learn：Clipboard  
  https://learn.microsoft.com/windows/win32/dataxchg/clipboard
- Microsoft Learn：Hooks Overview  
  https://learn.microsoft.com/windows/win32/winmsg/about-hooks
- Go `syscall` package  
  https://pkg.go.dev/syscall
- Go `x/sys/windows` package  
  https://pkg.go.dev/golang.org/x/sys/windows

---

## 38. 总结

Go 调用 `user32.dll` 的核心是：

1. 加载 DLL：`syscall.NewLazyDLL("user32.dll")`
2. 获取函数：`NewProc("FunctionNameW")`
3. 转换参数：字符串转 UTF-16，句柄 / 指针转 `uintptr`
4. 调用 API：`proc.Call(...)`
5. 正确处理返回值、错误、资源释放

`user32.dll` 基本覆盖了 Windows 用户界面层的大部分能力：窗口、消息、输入、剪贴板、Hook、菜单、显示器、DPI、系统参数等。  
如果只是自动化窗口或做桌面小工具，最常用的是：

```text
FindWindowW
EnumWindows
GetWindowTextW
GetForegroundWindow
ShowWindow
SetWindowPos
SendMessageW
PostMessageW
GetAsyncKeyState
SendInput
RegisterHotKey
OpenClipboard
SetWindowsHookExW
SetWinEventHook
```
