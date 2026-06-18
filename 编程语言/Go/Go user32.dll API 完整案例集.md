---
title: Go user32.dll API 完整案例集
date: 2026-06-18
tags:
  - Go
  - Windows
  - Win32
  - user32.dll
  - API案例
---

# Go user32.dll API 完整案例集

> 这是《Go 调用 user32.dll 操作整理》的案例版。  
> 目标：不是只列 API 名，而是让每个常用 `user32.dll` API 都有一个可理解、可改造的 Go 使用案例。  
> 说明：`user32.dll` / `winuser.h` 的 API 非常多，微软文档完整列表超过数百项。本文先覆盖桌面自动化、窗口管理、输入、剪贴板、Hook、菜单、DPI、显示器、系统参数等开发中最常用的一批 API。  
> 如果要把 Microsoft Learn `winuser.h` 页面里的全部 API 逐个展开，会形成一本很大的手册，建议后续拆成多篇：窗口篇、输入篇、剪贴板篇、Hook 篇、DPI 篇、菜单篇。

---

## 0. 通用准备代码

后面的案例默认使用下面这套基础代码。每个 API 案例为了避免重复，不一定都重新写 `package main`，但都能嵌入这个模板中使用。

```go
package main

import (
    "fmt"
    "syscall"
    "unsafe"
)

var (
    user32   = syscall.NewLazyDLL("user32.dll")
    kernel32 = syscall.NewLazyDLL("kernel32.dll")
)

func utf16Ptr(s string) uintptr {
    p, err := syscall.UTF16PtrFromString(s)
    if err != nil {
        panic(err)
    }
    return uintptr(unsafe.Pointer(p))
}

func boolRet(r uintptr) bool {
    return r != 0
}

type POINT struct {
    X int32
    Y int32
}

type RECT struct {
    Left   int32
    Top    int32
    Right  int32
    Bottom int32
}
```

---

## 1. MessageBoxW：显示消息框

### 场景

程序启动后弹出一个 Windows 原生提示框。

### API

```text
int MessageBoxW(HWND hWnd, LPCWSTR lpText, LPCWSTR lpCaption, UINT uType);
```

### Go 案例

```go
var procMessageBoxW = user32.NewProc("MessageBoxW")

const (
    MB_OK              = 0x00000000
    MB_ICONINFORMATION = 0x00000040
)

func ShowInfoBox(title, text string) int {
    ret, _, _ := procMessageBoxW.Call(
        0,
        utf16Ptr(text),
        utf16Ptr(title),
        MB_OK|MB_ICONINFORMATION,
    )
    return int(ret)
}

func main() {
    ShowInfoBox("Go + user32.dll", "这是 MessageBoxW 的完整调用案例")
}
```

---

## 2. MessageBeep：播放系统提示音

### 场景

程序完成任务后播放一个系统提示音。

### Go 案例

```go
var procMessageBeep = user32.NewProc("MessageBeep")

const MB_ICONASTERISK = 0x00000040

func PlayNotifySound() bool {
    ret, _, _ := procMessageBeep.Call(MB_ICONASTERISK)
    return ret != 0
}

func main() {
    PlayNotifySound()
}
```

---

## 3. FindWindowW：按标题或类名查找窗口

### 场景

查找标题为“计算器”的窗口句柄。

### Go 案例

```go
var procFindWindowW = user32.NewProc("FindWindowW")

func FindWindowByTitle(title string) uintptr {
    hwnd, _, _ := procFindWindowW.Call(
        0,              // lpClassName，0 表示不按类名查找
        utf16Ptr(title), // lpWindowName
    )
    return hwnd
}

func main() {
    hwnd := FindWindowByTitle("计算器")
    if hwnd == 0 {
        fmt.Println("没找到窗口")
        return
    }
    fmt.Printf("窗口句柄: 0x%X\n", hwnd)
}
```

---

## 4. FindWindowExW：查找子窗口

### 场景

在某个父窗口下面查找指定类名或标题的子控件。

### Go 案例

```go
var procFindWindowExW = user32.NewProc("FindWindowExW")

func FindChildWindow(parent uintptr, className, title string) uintptr {
    var classPtr uintptr
    var titlePtr uintptr

    if className != "" {
        classPtr = utf16Ptr(className)
    }
    if title != "" {
        titlePtr = utf16Ptr(title)
    }

    hwnd, _, _ := procFindWindowExW.Call(
        parent,
        0,        // 从第一个子窗口开始查找
        classPtr,
        titlePtr,
    )
    return hwnd
}
```

---

## 5. GetForegroundWindow：获取当前前台窗口

### 场景

获取用户当前正在操作的窗口。

### Go 案例

```go
var procGetForegroundWindow = user32.NewProc("GetForegroundWindow")

func GetForeground() uintptr {
    hwnd, _, _ := procGetForegroundWindow.Call()
    return hwnd
}

func main() {
    hwnd := GetForeground()
    fmt.Printf("当前前台窗口: 0x%X\n", hwnd)
}
```

---

## 6. GetWindowTextW：读取窗口标题

### 场景

读取当前前台窗口标题。

### Go 案例

```go
var procGetWindowTextW = user32.NewProc("GetWindowTextW")

func GetWindowTitle(hwnd uintptr) string {
    buf := make([]uint16, 512)
    procGetWindowTextW.Call(
        hwnd,
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(len(buf)),
    )
    return syscall.UTF16ToString(buf)
}

func main() {
    hwnd := GetForeground()
    fmt.Println(GetWindowTitle(hwnd))
}
```

---

## 7. GetWindowTextLengthW：获取窗口标题长度

### 场景

先获取标题长度，再分配刚好够用的缓冲区。

### Go 案例

```go
var procGetWindowTextLengthW = user32.NewProc("GetWindowTextLengthW")

func GetWindowTitleExact(hwnd uintptr) string {
    n, _, _ := procGetWindowTextLengthW.Call(hwnd)
    if n == 0 {
        return ""
    }

    buf := make([]uint16, n+1)
    procGetWindowTextW.Call(
        hwnd,
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(len(buf)),
    )
    return syscall.UTF16ToString(buf)
}
```

---

## 8. SetWindowTextW：修改窗口标题

### 场景

把自己创建的窗口标题改成新文字。

### Go 案例

```go
var procSetWindowTextW = user32.NewProc("SetWindowTextW")

func SetWindowTitle(hwnd uintptr, title string) bool {
    ret, _, _ := procSetWindowTextW.Call(hwnd, utf16Ptr(title))
    return ret != 0
}
```

> 注意：跨进程修改别的软件窗口标题不一定可靠，推荐只对自己创建的窗口使用。

---

## 9. GetClassNameW：获取窗口类名

### 场景

判断一个窗口是不是某类控件，例如 `Button`、`Edit`。

### Go 案例

```go
var procGetClassNameW = user32.NewProc("GetClassNameW")

func GetClassName(hwnd uintptr) string {
    buf := make([]uint16, 256)
    procGetClassNameW.Call(
        hwnd,
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(len(buf)),
    )
    return syscall.UTF16ToString(buf)
}
```

---

## 10. GetWindowThreadProcessId：获取窗口所属进程 ID

### 场景

知道一个窗口属于哪个进程。

### Go 案例

```go
var procGetWindowThreadProcessId = user32.NewProc("GetWindowThreadProcessId")

func GetWindowPID(hwnd uintptr) (threadID, processID uint32) {
    var pid uint32
    tid, _, _ := procGetWindowThreadProcessId.Call(
        hwnd,
        uintptr(unsafe.Pointer(&pid)),
    )
    return uint32(tid), pid
}

func main() {
    hwnd := GetForeground()
    tid, pid := GetWindowPID(hwnd)
    fmt.Println("线程ID:", tid, "进程ID:", pid)
}
```

---

## 11. IsWindow：判断窗口句柄是否有效

### 场景

保存了一个窗口句柄，后续使用前判断它是否仍然有效。

### Go 案例

```go
var procIsWindow = user32.NewProc("IsWindow")

func IsValidWindow(hwnd uintptr) bool {
    ret, _, _ := procIsWindow.Call(hwnd)
    return ret != 0
}
```

---

## 12. IsWindowVisible：判断窗口是否可见

```go
var procIsWindowVisible = user32.NewProc("IsWindowVisible")

func IsVisible(hwnd uintptr) bool {
    ret, _, _ := procIsWindowVisible.Call(hwnd)
    return ret != 0
}
```

使用场景：枚举窗口时过滤不可见窗口。

---

## 13. IsWindowEnabled：判断窗口是否可操作

```go
var procIsWindowEnabled = user32.NewProc("IsWindowEnabled")

func IsEnabled(hwnd uintptr) bool {
    ret, _, _ := procIsWindowEnabled.Call(hwnd)
    return ret != 0
}
```

使用场景：自动化点击前判断按钮是否启用。

---

## 14. IsIconic：判断窗口是否最小化

```go
var procIsIconic = user32.NewProc("IsIconic")

func IsMinimized(hwnd uintptr) bool {
    ret, _, _ := procIsIconic.Call(hwnd)
    return ret != 0
}
```

---

## 15. IsZoomed：判断窗口是否最大化

```go
var procIsZoomed = user32.NewProc("IsZoomed")

func IsMaximized(hwnd uintptr) bool {
    ret, _, _ := procIsZoomed.Call(hwnd)
    return ret != 0
}
```

---

## 16. ShowWindow：显示、隐藏、最小化、还原窗口

### 场景

找到窗口后将它恢复显示。

### Go 案例

```go
var procShowWindow = user32.NewProc("ShowWindow")

const (
    SW_HIDE            = 0
    SW_SHOWNORMAL      = 1
    SW_SHOWMINIMIZED   = 2
    SW_SHOWMAXIMIZED   = 3
    SW_SHOW            = 5
    SW_MINIMIZE        = 6
    SW_RESTORE         = 9
)

func RestoreWindow(hwnd uintptr) bool {
    ret, _, _ := procShowWindow.Call(hwnd, SW_RESTORE)
    return ret != 0
}

func HideWindow(hwnd uintptr) bool {
    ret, _, _ := procShowWindow.Call(hwnd, SW_HIDE)
    return ret != 0
}
```

---

## 17. ShowWindowAsync：异步显示窗口

### 场景

跨线程或跨进程显示窗口时，避免同步阻塞。

```go
var procShowWindowAsync = user32.NewProc("ShowWindowAsync")

func RestoreWindowAsync(hwnd uintptr) bool {
    ret, _, _ := procShowWindowAsync.Call(hwnd, SW_RESTORE)
    return ret != 0
}
```

---

## 18. SetForegroundWindow：激活前台窗口

```go
var procSetForegroundWindow = user32.NewProc("SetForegroundWindow")

func ActivateWindow(hwnd uintptr) bool {
    RestoreWindow(hwnd)
    ret, _, _ := procSetForegroundWindow.Call(hwnd)
    return ret != 0
}
```

> Windows 有前台窗口限制，不是任何时候都能把别的窗口抢到前台。

---

## 19. BringWindowToTop：窗口放到 Z 顺序顶部

```go
var procBringWindowToTop = user32.NewProc("BringWindowToTop")

func BringTop(hwnd uintptr) bool {
    ret, _, _ := procBringWindowToTop.Call(hwnd)
    return ret != 0
}
```

---

## 20. SetWindowPos：移动、调整大小、置顶

### 场景

把窗口移动到左上角并设置大小为 800x600。

```go
var procSetWindowPos = user32.NewProc("SetWindowPos")

const (
    HWND_TOP       = 0
    HWND_TOPMOST   = ^uintptr(0) // -1
    HWND_NOTOPMOST = ^uintptr(1) // -2

    SWP_NOSIZE     = 0x0001
    SWP_NOMOVE     = 0x0002
    SWP_NOZORDER   = 0x0004
    SWP_SHOWWINDOW = 0x0040
)

func MoveResizeWindow(hwnd uintptr, x, y, w, h int32) bool {
    ret, _, _ := procSetWindowPos.Call(
        hwnd,
        HWND_TOP,
        uintptr(x), uintptr(y), uintptr(w), uintptr(h),
        SWP_SHOWWINDOW,
    )
    return ret != 0
}

func SetTopMost(hwnd uintptr, yes bool) bool {
    after := HWND_NOTOPMOST
    if yes {
        after = HWND_TOPMOST
    }
    ret, _, _ := procSetWindowPos.Call(
        hwnd, after,
        0, 0, 0, 0,
        SWP_NOMOVE|SWP_NOSIZE,
    )
    return ret != 0
}
```

---

## 21. MoveWindow：移动窗口并调整尺寸

```go
var procMoveWindow = user32.NewProc("MoveWindow")

func MoveWindowTo(hwnd uintptr, x, y, w, h int32) bool {
    ret, _, _ := procMoveWindow.Call(
        hwnd,
        uintptr(x), uintptr(y), uintptr(w), uintptr(h),
        1, // repaint
    )
    return ret != 0
}
```

---

## 22. GetWindowRect：获取窗口外框矩形

```go
var procGetWindowRect = user32.NewProc("GetWindowRect")

func GetWindowRect(hwnd uintptr) (RECT, bool) {
    var r RECT
    ret, _, _ := procGetWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&r)))
    return r, ret != 0
}
```

使用：

```go
r, ok := GetWindowRect(hwnd)
if ok {
    fmt.Println(r.Left, r.Top, r.Right, r.Bottom)
}
```

---

## 23. GetClientRect：获取客户区大小

```go
var procGetClientRect = user32.NewProc("GetClientRect")

func GetClientRect(hwnd uintptr) (RECT, bool) {
    var r RECT
    ret, _, _ := procGetClientRect.Call(hwnd, uintptr(unsafe.Pointer(&r)))
    return r, ret != 0
}
```

---

## 24. ClientToScreen：客户区坐标转屏幕坐标

```go
var procClientToScreen = user32.NewProc("ClientToScreen")

func ClientToScreen(hwnd uintptr, x, y int32) (POINT, bool) {
    pt := POINT{x, y}
    ret, _, _ := procClientToScreen.Call(hwnd, uintptr(unsafe.Pointer(&pt)))
    return pt, ret != 0
}
```

---

## 25. ScreenToClient：屏幕坐标转客户区坐标

```go
var procScreenToClient = user32.NewProc("ScreenToClient")

func ScreenToClient(hwnd uintptr, x, y int32) (POINT, bool) {
    pt := POINT{x, y}
    ret, _, _ := procScreenToClient.Call(hwnd, uintptr(unsafe.Pointer(&pt)))
    return pt, ret != 0
}
```

---

## 26. WindowFromPoint：根据屏幕坐标获取窗口

```go
var procWindowFromPoint = user32.NewProc("WindowFromPoint")

func WindowAt(x, y int32) uintptr {
    // POINT 作为两个 int32，在 64 位下可打包为一个 uintptr
    packed := uintptr(uint32(x)) | uintptr(uint64(uint32(y))<<32)
    hwnd, _, _ := procWindowFromPoint.Call(packed)
    return hwnd
}
```

更稳妥的写法可以使用 `x/sys/windows` 定义结构体调用。

---

## 27. ChildWindowFromPoint：根据父窗口客户区坐标找子窗口

```go
var procChildWindowFromPoint = user32.NewProc("ChildWindowFromPoint")

func ChildAt(parent uintptr, x, y int32) uintptr {
    packed := uintptr(uint32(x)) | uintptr(uint64(uint32(y))<<32)
    hwnd, _, _ := procChildWindowFromPoint.Call(parent, packed)
    return hwnd
}
```

---

## 28. SendMessageW：同步发送窗口消息

### 场景

向窗口发送关闭消息，并等待窗口处理。

```go
var procSendMessageW = user32.NewProc("SendMessageW")

const WM_CLOSE = 0x0010

func CloseBySendMessage(hwnd uintptr) uintptr {
    ret, _, _ := procSendMessageW.Call(hwnd, WM_CLOSE, 0, 0)
    return ret
}
```

> `SendMessageW` 会阻塞，目标窗口卡住时调用方也会卡住。

---

## 29. SendMessageTimeoutW：带超时发送消息

### 场景

跨进程发送消息，避免目标窗口无响应导致自己卡死。

```go
var procSendMessageTimeoutW = user32.NewProc("SendMessageTimeoutW")

const (
    SMTO_ABORTIFHUNG = 0x0002
)

func SendCloseWithTimeout(hwnd uintptr, timeoutMs uint32) bool {
    var result uintptr
    ret, _, _ := procSendMessageTimeoutW.Call(
        hwnd,
        WM_CLOSE,
        0,
        0,
        SMTO_ABORTIFHUNG,
        uintptr(timeoutMs),
        uintptr(unsafe.Pointer(&result)),
    )
    return ret != 0
}
```

---

## 30. PostMessageW：异步投递消息

```go
var procPostMessageW = user32.NewProc("PostMessageW")

func CloseByPostMessage(hwnd uintptr) bool {
    ret, _, _ := procPostMessageW.Call(hwnd, WM_CLOSE, 0, 0)
    return ret != 0
}
```

---

## 31. PostThreadMessageW：向线程消息队列投递消息

```go
var procPostThreadMessageW = user32.NewProc("PostThreadMessageW")

const WM_USER = 0x0400

func NotifyThread(threadID uint32) bool {
    ret, _, _ := procPostThreadMessageW.Call(
        uintptr(threadID),
        WM_USER+1,
        123,
        456,
    )
    return ret != 0
}
```

---

## 32. RegisterWindowMessageW：注册唯一窗口消息

### 场景

多个程序用同一个字符串注册消息，得到相同 message id，用于进程间通信。

```go
var procRegisterWindowMessageW = user32.NewProc("RegisterWindowMessageW")

func RegisterMyMessage(name string) uint32 {
    msg, _, _ := procRegisterWindowMessageW.Call(utf16Ptr(name))
    return uint32(msg)
}
```

---

## 33. GetMessageW / TranslateMessage / DispatchMessageW：消息循环

### 场景

写 Win32 GUI、热键、Hook、剪贴板监听时维持消息循环。

```go
var (
    procGetMessageW      = user32.NewProc("GetMessageW")
    procTranslateMessage = user32.NewProc("TranslateMessage")
    procDispatchMessageW = user32.NewProc("DispatchMessageW")
)

type MSG struct {
    Hwnd    uintptr
    Message uint32
    WParam  uintptr
    LParam  uintptr
    Time    uint32
    Pt      POINT
}

func MessageLoop() int {
    var msg MSG
    for {
        ret, _, _ := procGetMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
        if int32(ret) == -1 {
            return -1
        }
        if ret == 0 {
            return int(msg.WParam)
        }
        procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
        procDispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
    }
}
```

---

## 34. PeekMessageW：非阻塞读取消息

```go
var procPeekMessageW = user32.NewProc("PeekMessageW")

const PM_REMOVE = 0x0001

func PumpOnce() bool {
    var msg MSG
    ret, _, _ := procPeekMessageW.Call(
        uintptr(unsafe.Pointer(&msg)),
        0,
        0,
        0,
        PM_REMOVE,
    )
    if ret == 0 {
        return false
    }
    procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
    procDispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
    return true
}
```

---

## 35. PostQuitMessage：退出消息循环

```go
var procPostQuitMessage = user32.NewProc("PostQuitMessage")

func QuitApp(exitCode int) {
    procPostQuitMessage.Call(uintptr(exitCode))
}
```

通常在窗口过程收到 `WM_DESTROY` 时调用。

---

## 36. GetAsyncKeyState：检测按键是否按下

```go
var procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")

const VK_CONTROL = 0x11

func IsKeyDown(vk int) bool {
    ret, _, _ := procGetAsyncKeyState.Call(uintptr(vk))
    return ret&0x8000 != 0
}

func main() {
    if IsKeyDown(VK_CONTROL) {
        fmt.Println("Ctrl 正在按下")
    }
}
```

---

## 37. GetKeyState：获取按键状态

```go
var procGetKeyState = user32.NewProc("GetKeyState")

const VK_CAPITAL = 0x14

func IsCapsLockOn() bool {
    ret, _, _ := procGetKeyState.Call(VK_CAPITAL)
    return ret&1 != 0
}
```

---

## 38. GetKeyboardState：获取 256 个虚拟键状态

```go
var procGetKeyboardState = user32.NewProc("GetKeyboardState")

func GetKeyboardState() ([256]byte, bool) {
    var state [256]byte
    ret, _, _ := procGetKeyboardState.Call(uintptr(unsafe.Pointer(&state[0])))
    return state, ret != 0
}
```

---

## 39. MapVirtualKeyW：键码与扫描码转换

```go
var procMapVirtualKeyW = user32.NewProc("MapVirtualKeyW")

const MAPVK_VK_TO_VSC = 0

func VirtualKeyToScanCode(vk uint32) uint32 {
    ret, _, _ := procMapVirtualKeyW.Call(uintptr(vk), MAPVK_VK_TO_VSC)
    return uint32(ret)
}
```

---

## 40. ToUnicode：键盘状态转换为字符

```go
var procToUnicode = user32.NewProc("ToUnicode")

func KeyToUnicode(vk, scanCode uint32, state *[256]byte) string {
    buf := make([]uint16, 8)
    ret, _, _ := procToUnicode.Call(
        uintptr(vk),
        uintptr(scanCode),
        uintptr(unsafe.Pointer(&state[0])),
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(len(buf)),
        0,
    )
    if int32(ret) <= 0 {
        return ""
    }
    return syscall.UTF16ToString(buf[:ret])
}
```

---

## 41. GetCursorPos：获取鼠标坐标

```go
var procGetCursorPos = user32.NewProc("GetCursorPos")

func GetMousePos() (POINT, bool) {
    var pt POINT
    ret, _, _ := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
    return pt, ret != 0
}
```

---

## 42. SetCursorPos：设置鼠标坐标

```go
var procSetCursorPos = user32.NewProc("SetCursorPos")

func SetMousePos(x, y int32) bool {
    ret, _, _ := procSetCursorPos.Call(uintptr(x), uintptr(y))
    return ret != 0
}
```

---

## 43. SendInput：模拟键盘 / 鼠标输入

### 场景

模拟按下并松开 A 键。

```go
var procSendInput = user32.NewProc("SendInput")

const (
    INPUT_MOUSE    = 0
    INPUT_KEYBOARD = 1
    KEYEVENTF_KEYUP = 0x0002
    VK_A = 0x41
)

type KEYBDINPUT struct {
    WVk         uint16
    WScan       uint16
    DwFlags     uint32
    Time        uint32
    DwExtraInfo uintptr
}

type INPUT struct {
    Type uint32
    Ki   KEYBDINPUT
    Padding [16]byte // 简化对齐；生产代码建议严格按平台定义结构
}

func PressA() uint32 {
    inputs := []INPUT{
        {Type: INPUT_KEYBOARD, Ki: KEYBDINPUT{WVk: VK_A}},
        {Type: INPUT_KEYBOARD, Ki: KEYBDINPUT{WVk: VK_A, DwFlags: KEYEVENTF_KEYUP}},
    }
    ret, _, _ := procSendInput.Call(
        uintptr(len(inputs)),
        uintptr(unsafe.Pointer(&inputs[0])),
        unsafe.Sizeof(inputs[0]),
    )
    return uint32(ret)
}
```

> `SendInput` 的结构体对齐很关键，实际项目建议直接使用成熟封装或仔细对照 Windows 结构体。

---

## 44. keybd_event：老式键盘模拟

```go
var procKeybdEvent = user32.NewProc("keybd_event")

const KEYEVENTF_KEYUP_OLD = 0x0002

func PressEnterOld() {
    const VK_RETURN = 0x0D
    procKeybdEvent.Call(VK_RETURN, 0, 0, 0)
    procKeybdEvent.Call(VK_RETURN, 0, KEYEVENTF_KEYUP_OLD, 0)
}
```

推荐新代码使用 `SendInput`。

---

## 45. mouse_event：老式鼠标模拟

```go
var procMouseEvent = user32.NewProc("mouse_event")

const (
    MOUSEEVENTF_LEFTDOWN = 0x0002
    MOUSEEVENTF_LEFTUP   = 0x0004
)

func LeftClickOld() {
    procMouseEvent.Call(MOUSEEVENTF_LEFTDOWN, 0, 0, 0, 0)
    procMouseEvent.Call(MOUSEEVENTF_LEFTUP, 0, 0, 0, 0)
}
```

---

## 46. RegisterHotKey / UnregisterHotKey：注册全局快捷键

### 场景

注册 Ctrl+Alt+K，按下后收到 `WM_HOTKEY`。

```go
var (
    procRegisterHotKey   = user32.NewProc("RegisterHotKey")
    procUnregisterHotKey = user32.NewProc("UnregisterHotKey")
)

const (
    MOD_ALT     = 0x0001
    MOD_CONTROL = 0x0002
    WM_HOTKEY   = 0x0312
)

func RegisterCtrlAltK() bool {
    id := uintptr(1)
    ret, _, _ := procRegisterHotKey.Call(
        0,
        id,
        MOD_CONTROL|MOD_ALT,
        'K',
    )
    return ret != 0
}

func UnregisterCtrlAltK() {
    procUnregisterHotKey.Call(0, 1)
}
```

消息循环里判断：

```go
if msg.Message == WM_HOTKEY {
    fmt.Println("Ctrl+Alt+K 被按下")
}
```

---

## 47. OpenClipboard / CloseClipboard：打开和关闭剪贴板

```go
var (
    procOpenClipboard  = user32.NewProc("OpenClipboard")
    procCloseClipboard = user32.NewProc("CloseClipboard")
)

func WithClipboard(fn func()) bool {
    ret, _, _ := procOpenClipboard.Call(0)
    if ret == 0 {
        return false
    }
    defer procCloseClipboard.Call()
    fn()
    return true
}
```

---

## 48. IsClipboardFormatAvailable：判断剪贴板是否有 Unicode 文本

```go
var procIsClipboardFormatAvailable = user32.NewProc("IsClipboardFormatAvailable")

const CF_UNICODETEXT = 13

func HasUnicodeText() bool {
    ret, _, _ := procIsClipboardFormatAvailable.Call(CF_UNICODETEXT)
    return ret != 0
}
```

---

## 49. GetClipboardData：读取剪贴板文本

需要配合 `kernel32.GlobalLock` / `GlobalUnlock`。

```go
var (
    procGetClipboardData = user32.NewProc("GetClipboardData")
    procGlobalLock       = kernel32.NewProc("GlobalLock")
    procGlobalUnlock     = kernel32.NewProc("GlobalUnlock")
)

func GetClipboardText() string {
    var text string
    WithClipboard(func() {
        if !HasUnicodeText() {
            return
        }
        h, _, _ := procGetClipboardData.Call(CF_UNICODETEXT)
        if h == 0 {
            return
        }
        p, _, _ := procGlobalLock.Call(h)
        if p == 0 {
            return
        }
        defer procGlobalUnlock.Call(h)

        // 找 UTF-16 结尾
        ptr := (*[1 << 20]uint16)(unsafe.Pointer(p))
        n := 0
        for ptr[n] != 0 {
            n++
        }
        text = syscall.UTF16ToString(ptr[:n])
    })
    return text
}
```

---

## 50. EmptyClipboard / SetClipboardData：写入剪贴板文本

需要配合 `GlobalAlloc`，并且调用 `SetClipboardData` 成功后内存所有权交给系统。

```go
var (
    procEmptyClipboard = user32.NewProc("EmptyClipboard")
    procSetClipboardData = user32.NewProc("SetClipboardData")
    procGlobalAlloc = kernel32.NewProc("GlobalAlloc")
)

const GMEM_MOVEABLE = 0x0002

func SetClipboardText(text string) bool {
    utf16, _ := syscall.UTF16FromString(text)
    size := uintptr(len(utf16) * 2)

    ok := false
    WithClipboard(func() {
        procEmptyClipboard.Call()

        h, _, _ := procGlobalAlloc.Call(GMEM_MOVEABLE, size)
        if h == 0 {
            return
        }

        p, _, _ := procGlobalLock.Call(h)
        if p == 0 {
            return
        }
        copy((*[1 << 20]uint16)(unsafe.Pointer(p))[:len(utf16)], utf16)
        procGlobalUnlock.Call(h)

        ret, _, _ := procSetClipboardData.Call(CF_UNICODETEXT, h)
        ok = ret != 0
    })
    return ok
}
```

---

## 51. AddClipboardFormatListener / RemoveClipboardFormatListener：监听剪贴板变化

### 场景

创建一个隐藏窗口，调用 `AddClipboardFormatListener(hwnd)`，当剪贴板变化时窗口收到 `WM_CLIPBOARDUPDATE`。

```go
var (
    procAddClipboardFormatListener = user32.NewProc("AddClipboardFormatListener")
    procRemoveClipboardFormatListener = user32.NewProc("RemoveClipboardFormatListener")
)

const WM_CLIPBOARDUPDATE = 0x031D

func AddClipboardListener(hwnd uintptr) bool {
    ret, _, _ := procAddClipboardFormatListener.Call(hwnd)
    return ret != 0
}

func RemoveClipboardListener(hwnd uintptr) bool {
    ret, _, _ := procRemoveClipboardFormatListener.Call(hwnd)
    return ret != 0
}
```

窗口过程里：

```go
if msg == WM_CLIPBOARDUPDATE {
    fmt.Println("剪贴板变了：", GetClipboardText())
}
```

---

## 52. EnumWindows：枚举所有顶层窗口

```go
var procEnumWindows = user32.NewProc("EnumWindows")

func EnumTopWindows() []uintptr {
    var windows []uintptr

    cb := syscall.NewCallback(func(hwnd uintptr, lparam uintptr) uintptr {
        if IsVisible(hwnd) {
            windows = append(windows, hwnd)
        }
        return 1 // 继续枚举
    })

    procEnumWindows.Call(cb, 0)
    return windows
}

func main() {
    for _, hwnd := range EnumTopWindows() {
        fmt.Printf("0x%X %s\n", hwnd, GetWindowTitle(hwnd))
    }
}
```

---

## 53. EnumChildWindows：枚举子窗口

```go
var procEnumChildWindows = user32.NewProc("EnumChildWindows")

func EnumChildren(parent uintptr) []uintptr {
    var children []uintptr
    cb := syscall.NewCallback(func(hwnd uintptr, lparam uintptr) uintptr {
        children = append(children, hwnd)
        return 1
    })
    procEnumChildWindows.Call(parent, cb, 0)
    return children
}
```

---

## 54. GetParent：获取父窗口

```go
var procGetParent = user32.NewProc("GetParent")

func GetParentWindow(hwnd uintptr) uintptr {
    parent, _, _ := procGetParent.Call(hwnd)
    return parent
}
```

---

## 55. GetAncestor：获取祖先窗口

```go
var procGetAncestor = user32.NewProc("GetAncestor")

const GA_ROOT = 2

func GetRootWindow(hwnd uintptr) uintptr {
    root, _, _ := procGetAncestor.Call(hwnd, GA_ROOT)
    return root
}
```

---

## 56. GetWindow：获取相关窗口

```go
var procGetWindow = user32.NewProc("GetWindow")

const (
    GW_HWNDFIRST = 0
    GW_HWNDLAST  = 1
    GW_HWNDNEXT  = 2
    GW_HWNDPREV  = 3
    GW_OWNER     = 4
    GW_CHILD     = 5
)

func GetNextWindow(hwnd uintptr) uintptr {
    next, _, _ := procGetWindow.Call(hwnd, GW_HWNDNEXT)
    return next
}
```

---

## 57. GetDesktopWindow：获取桌面窗口句柄

```go
var procGetDesktopWindow = user32.NewProc("GetDesktopWindow")

func DesktopWindow() uintptr {
    hwnd, _, _ := procGetDesktopWindow.Call()
    return hwnd
}
```

---

## 58. GetShellWindow：获取 Shell 桌面窗口

```go
var procGetShellWindow = user32.NewProc("GetShellWindow")

func ShellWindow() uintptr {
    hwnd, _, _ := procGetShellWindow.Call()
    return hwnd
}
```

---

## 59. SetFocus：设置键盘焦点

```go
var procSetFocus = user32.NewProc("SetFocus")

func FocusWindow(hwnd uintptr) uintptr {
    old, _, _ := procSetFocus.Call(hwnd)
    return old
}
```

> 通常只对同线程窗口可靠。跨进程焦点控制常受限制。

---

## 60. GetFocus：获取当前线程焦点窗口

```go
var procGetFocus = user32.NewProc("GetFocus")

func CurrentThreadFocus() uintptr {
    hwnd, _, _ := procGetFocus.Call()
    return hwnd
}
```

---

## 61. SetActiveWindow / GetActiveWindow：活动窗口

```go
var (
    procSetActiveWindow = user32.NewProc("SetActiveWindow")
    procGetActiveWindow = user32.NewProc("GetActiveWindow")
)

func SetActive(hwnd uintptr) uintptr {
    old, _, _ := procSetActiveWindow.Call(hwnd)
    return old
}

func GetActive() uintptr {
    hwnd, _, _ := procGetActiveWindow.Call()
    return hwnd
}
```

---

## 62. EnableWindow：启用或禁用窗口

```go
var procEnableWindow = user32.NewProc("EnableWindow")

func Enable(hwnd uintptr, enabled bool) bool {
    v := uintptr(0)
    if enabled {
        v = 1
    }
    old, _, _ := procEnableWindow.Call(hwnd, v)
    return old != 0
}
```

---

## 63. DestroyWindow：销毁窗口

```go
var procDestroyWindow = user32.NewProc("DestroyWindow")

func Destroy(hwnd uintptr) bool {
    ret, _, _ := procDestroyWindow.Call(hwnd)
    return ret != 0
}
```

> 只能安全销毁自己线程 / 自己程序创建的窗口。不要乱销毁其他程序窗口。

---

## 64. GetWindowLongPtrW / SetWindowLongPtrW：读写窗口属性

### 场景

读取窗口样式，然后添加 `WS_EX_TOPMOST` 等扩展样式通常配合 `SetWindowPos`。

```go
var (
    procGetWindowLongPtrW = user32.NewProc("GetWindowLongPtrW")
    procSetWindowLongPtrW = user32.NewProc("SetWindowLongPtrW")
)

const (
    GWL_STYLE   = ^uintptr(16 - 1) // -16 的 uintptr 表示较绕，实际建议用 uintptr(^uint(15))
    GWL_EXSTYLE = ^uintptr(20 - 1) // -20
)

func GetWindowStyle(hwnd uintptr) uintptr {
    style, _, _ := procGetWindowLongPtrW.Call(hwnd, uintptr(int32(-16)))
    return style
}

func SetWindowStyle(hwnd uintptr, style uintptr) uintptr {
    old, _, _ := procSetWindowLongPtrW.Call(hwnd, uintptr(int32(-16)), style)
    return old
}
```

> 注意：负数常量转换到 `uintptr` 容易写错，生产代码建议封装成专门函数或用 `x/sys/windows`。

---

## 65. SetLayeredWindowAttributes：设置窗口透明度

```go
var procSetLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes")

const (
    WS_EX_LAYERED = 0x00080000
    LWA_ALPHA     = 0x00000002
)

func SetWindowAlpha(hwnd uintptr, alpha byte) bool {
    exStyle, _, _ := procGetWindowLongPtrW.Call(hwnd, uintptr(int32(-20)))
    procSetWindowLongPtrW.Call(hwnd, uintptr(int32(-20)), exStyle|WS_EX_LAYERED)

    ret, _, _ := procSetLayeredWindowAttributes.Call(hwnd, 0, uintptr(alpha), LWA_ALPHA)
    return ret != 0
}
```

---

## 66. RegisterClassExW / CreateWindowExW / DefWindowProcW：创建原生窗口

这是完整 Win32 GUI 的核心组合。

```go
var (
    procRegisterClassExW = user32.NewProc("RegisterClassExW")
    procCreateWindowExW  = user32.NewProc("CreateWindowExW")
    procDefWindowProcW   = user32.NewProc("DefWindowProcW")
)

const (
    WM_DESTROY = 0x0002
    WS_OVERLAPPEDWINDOW = 0x00CF0000
    CW_USEDEFAULT = 0x80000000
)

type WNDCLASSEX struct {
    CbSize        uint32
    Style         uint32
    LpfnWndProc   uintptr
    CbClsExtra    int32
    CbWndExtra    int32
    HInstance     uintptr
    HIcon         uintptr
    HCursor       uintptr
    HbrBackground uintptr
    LpszMenuName  uintptr
    LpszClassName uintptr
    HIconSm       uintptr
}

var wndProcPtr uintptr

func wndProc(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
    switch msg {
    case WM_DESTROY:
        QuitApp(0)
        return 0
    }
    ret, _, _ := procDefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
    return ret
}

func CreateSimpleWindow() uintptr {
    wndProcPtr = syscall.NewCallback(wndProc)
    className := "GoUser32Window"

    wc := WNDCLASSEX{
        CbSize:        uint32(unsafe.Sizeof(WNDCLASSEX{})),
        LpfnWndProc:   wndProcPtr,
        LpszClassName: utf16Ptr(className),
    }

    procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))

    hwnd, _, _ := procCreateWindowExW.Call(
        0,
        utf16Ptr(className),
        utf16Ptr("Go 创建的 user32 窗口"),
        WS_OVERLAPPEDWINDOW,
        CW_USEDEFAULT, CW_USEDEFAULT,
        800, 600,
        0, 0, 0, 0,
    )
    return hwnd
}
```

---

## 67. UpdateWindow / InvalidateRect / RedrawWindow：触发重绘

```go
var (
    procUpdateWindow = user32.NewProc("UpdateWindow")
    procInvalidateRect = user32.NewProc("InvalidateRect")
    procRedrawWindow = user32.NewProc("RedrawWindow")
)

func ForceRepaint(hwnd uintptr) {
    procInvalidateRect.Call(hwnd, 0, 1)
    procUpdateWindow.Call(hwnd)
}
```

---

## 68. BeginPaint / EndPaint：处理 WM_PAINT

```go
var (
    procBeginPaint = user32.NewProc("BeginPaint")
    procEndPaint   = user32.NewProc("EndPaint")
)

type PAINTSTRUCT struct {
    Hdc         uintptr
    FErase      int32
    RcPaint     RECT
    FRestore    int32
    FIncUpdate  int32
    RgbReserved [32]byte
}

func OnPaint(hwnd uintptr) {
    var ps PAINTSTRUCT
    hdc, _, _ := procBeginPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))
    _ = hdc
    // 这里可以调用 DrawTextW 或 gdi32 绘图
    procEndPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))
}
```

---

## 69. DrawTextW：绘制文本

```go
var procDrawTextW = user32.NewProc("DrawTextW")

const DT_CENTER = 0x00000001
const DT_VCENTER = 0x00000004
const DT_SINGLELINE = 0x00000020

func DrawCenteredText(hdc uintptr, text string, rect RECT) {
    procDrawTextW.Call(
        hdc,
        utf16Ptr(text),
        uintptr(len([]rune(text))),
        uintptr(unsafe.Pointer(&rect)),
        DT_CENTER|DT_VCENTER|DT_SINGLELINE,
    )
}
```

---

## 70. GetDC / ReleaseDC：获取并释放设备上下文

```go
var (
    procGetDC = user32.NewProc("GetDC")
    procReleaseDC = user32.NewProc("ReleaseDC")
)

func WithWindowDC(hwnd uintptr, fn func(hdc uintptr)) bool {
    hdc, _, _ := procGetDC.Call(hwnd)
    if hdc == 0 {
        return false
    }
    defer procReleaseDC.Call(hwnd, hdc)
    fn(hdc)
    return true
}
```

---

## 71. CreateMenu / AppendMenuW / SetMenu / DrawMenuBar

### 场景

给自己创建的窗口添加一个菜单。

```go
var (
    procCreateMenu = user32.NewProc("CreateMenu")
    procAppendMenuW = user32.NewProc("AppendMenuW")
    procSetMenu = user32.NewProc("SetMenu")
    procDrawMenuBar = user32.NewProc("DrawMenuBar")
)

const MF_STRING = 0x00000000

func AddSimpleMenu(hwnd uintptr) bool {
    menu, _, _ := procCreateMenu.Call()
    if menu == 0 {
        return false
    }
    procAppendMenuW.Call(menu, MF_STRING, 1001, utf16Ptr("打开"))
    procAppendMenuW.Call(menu, MF_STRING, 1002, utf16Ptr("退出"))
    ret, _, _ := procSetMenu.Call(hwnd, menu)
    procDrawMenuBar.Call(hwnd)
    return ret != 0
}
```

---

## 72. CreatePopupMenu / TrackPopupMenu：右键菜单

```go
var (
    procCreatePopupMenu = user32.NewProc("CreatePopupMenu")
    procTrackPopupMenu = user32.NewProc("TrackPopupMenu")
)

const TPM_RIGHTBUTTON = 0x0002

func ShowContextMenu(hwnd uintptr, x, y int32) {
    menu, _, _ := procCreatePopupMenu.Call()
    procAppendMenuW.Call(menu, MF_STRING, 2001, utf16Ptr("复制"))
    procAppendMenuW.Call(menu, MF_STRING, 2002, utf16Ptr("粘贴"))
    procTrackPopupMenu.Call(menu, TPM_RIGHTBUTTON, uintptr(x), uintptr(y), 0, hwnd, 0)
}
```

---

## 73. DestroyMenu：销毁菜单

```go
var procDestroyMenu = user32.NewProc("DestroyMenu")

func DestroyMenu(menu uintptr) bool {
    ret, _, _ := procDestroyMenu.Call(menu)
    return ret != 0
}
```

---

## 74. EnableMenuItem / CheckMenuItem：启用和勾选菜单项

```go
var (
    procEnableMenuItem = user32.NewProc("EnableMenuItem")
    procCheckMenuItem = user32.NewProc("CheckMenuItem")
)

const (
    MF_BYCOMMAND = 0x00000000
    MF_GRAYED    = 0x00000001
    MF_ENABLED   = 0x00000000
    MF_CHECKED   = 0x00000008
    MF_UNCHECKED = 0x00000000
)

func DisableMenuCommand(menu uintptr, id uint32) {
    procEnableMenuItem.Call(menu, uintptr(id), MF_BYCOMMAND|MF_GRAYED)
}

func CheckMenuCommand(menu uintptr, id uint32, checked bool) {
    flag := MF_UNCHECKED
    if checked {
        flag = MF_CHECKED
    }
    procCheckMenuItem.Call(menu, uintptr(id), MF_BYCOMMAND|uintptr(flag))
}
```

---

## 75. LoadCursorW / SetCursor / ShowCursor

```go
var (
    procLoadCursorW = user32.NewProc("LoadCursorW")
    procSetCursor   = user32.NewProc("SetCursor")
    procShowCursor  = user32.NewProc("ShowCursor")
)

var IDC_HAND = uintptr(32649)

func SetHandCursor() {
    cursor, _, _ := procLoadCursorW.Call(0, IDC_HAND)
    procSetCursor.Call(cursor)
}

func HideCursor() {
    procShowCursor.Call(0)
}

func ShowCursorAgain() {
    procShowCursor.Call(1)
}
```

---

## 76. LoadIconW / DrawIcon / DestroyIcon

```go
var (
    procLoadIconW = user32.NewProc("LoadIconW")
    procDrawIcon = user32.NewProc("DrawIcon")
    procDestroyIcon = user32.NewProc("DestroyIcon")
)

var IDI_INFORMATION = uintptr(32516)

func DrawInfoIcon(hdc uintptr, x, y int32) {
    icon, _, _ := procLoadIconW.Call(0, IDI_INFORMATION)
    procDrawIcon.Call(hdc, uintptr(x), uintptr(y), icon)
    // 系统共享图标通常不需要 DestroyIcon；自己创建 / 复制的图标才销毁
}
```

---

## 77. SetTimer / KillTimer：窗口计时器

```go
var (
    procSetTimer = user32.NewProc("SetTimer")
    procKillTimer = user32.NewProc("KillTimer")
)

const WM_TIMER = 0x0113

func StartTimer(hwnd uintptr) uintptr {
    id, _, _ := procSetTimer.Call(hwnd, 1, 1000, 0) // 每 1000ms 收到 WM_TIMER
    return id
}

func StopTimer(hwnd uintptr, id uintptr) bool {
    ret, _, _ := procKillTimer.Call(hwnd, id)
    return ret != 0
}
```

窗口过程里：

```go
if msg == WM_TIMER {
    fmt.Println("timer tick")
}
```

---

## 78. GetSystemMetrics：获取系统尺寸和状态

```go
var procGetSystemMetrics = user32.NewProc("GetSystemMetrics")

const (
    SM_CXSCREEN = 0
    SM_CYSCREEN = 1
)

func ScreenSize() (w, h int) {
    rw, _, _ := procGetSystemMetrics.Call(SM_CXSCREEN)
    rh, _, _ := procGetSystemMetrics.Call(SM_CYSCREEN)
    return int(rw), int(rh)
}
```

---

## 79. SystemParametersInfoW：读取 / 修改系统参数

### 场景：读取鼠标速度

```go
var procSystemParametersInfoW = user32.NewProc("SystemParametersInfoW")

const SPI_GETMOUSESPEED = 0x0070

func GetMouseSpeed() int {
    var speed uint32
    procSystemParametersInfoW.Call(
        SPI_GETMOUSESPEED,
        0,
        uintptr(unsafe.Pointer(&speed)),
        0,
    )
    return int(speed)
}
```

> 修改系统参数会影响全局，写入类操作要谨慎。

---

## 80. LockWorkStation：锁定 Windows

```go
var procLockWorkStation = user32.NewProc("LockWorkStation")

func LockWindows() bool {
    ret, _, _ := procLockWorkStation.Call()
    return ret != 0
}
```

---

## 81. ExitWindowsEx：注销 / 关机 / 重启

```go
var procExitWindowsEx = user32.NewProc("ExitWindowsEx")

const EWX_LOGOFF = 0x00000000

func LogoffWindows() bool {
    ret, _, _ := procExitWindowsEx.Call(EWX_LOGOFF, 0)
    return ret != 0
}
```

> 关机 / 重启通常还需要权限调整，不建议随便调用。

---

## 82. EnumDisplayMonitors / GetMonitorInfoW：枚举显示器

```go
var (
    procEnumDisplayMonitors = user32.NewProc("EnumDisplayMonitors")
    procGetMonitorInfoW = user32.NewProc("GetMonitorInfoW")
)

type MONITORINFO struct {
    CbSize    uint32
    RcMonitor RECT
    RcWork    RECT
    DwFlags   uint32
}

func EnumMonitors() []RECT {
    var rects []RECT
    cb := syscall.NewCallback(func(hMonitor, hdc uintptr, lprcMonitor uintptr, dwData uintptr) uintptr {
        var mi MONITORINFO
        mi.CbSize = uint32(unsafe.Sizeof(mi))
        procGetMonitorInfoW.Call(hMonitor, uintptr(unsafe.Pointer(&mi)))
        rects = append(rects, mi.RcMonitor)
        return 1
    })
    procEnumDisplayMonitors.Call(0, 0, cb, 0)
    return rects
}
```

---

## 83. MonitorFromWindow / MonitorFromPoint / MonitorFromRect

```go
var (
    procMonitorFromWindow = user32.NewProc("MonitorFromWindow")
    procMonitorFromPoint  = user32.NewProc("MonitorFromPoint")
    procMonitorFromRect   = user32.NewProc("MonitorFromRect")
)

const MONITOR_DEFAULTTONEAREST = 2

func MonitorOfWindow(hwnd uintptr) uintptr {
    h, _, _ := procMonitorFromWindow.Call(hwnd, MONITOR_DEFAULTTONEAREST)
    return h
}

func MonitorOfPoint(x, y int32) uintptr {
    packed := uintptr(uint32(x)) | uintptr(uint64(uint32(y))<<32)
    h, _, _ := procMonitorFromPoint.Call(packed, MONITOR_DEFAULTTONEAREST)
    return h
}
```

---

## 84. GetDpiForWindow / GetDpiForSystem

```go
var (
    procGetDpiForWindow = user32.NewProc("GetDpiForWindow")
    procGetDpiForSystem = user32.NewProc("GetDpiForSystem")
)

func DpiForWindow(hwnd uintptr) uint32 {
    dpi, _, _ := procGetDpiForWindow.Call(hwnd)
    return uint32(dpi)
}

func SystemDPI() uint32 {
    dpi, _, _ := procGetDpiForSystem.Call()
    return uint32(dpi)
}
```

---

## 85. SetProcessDPIAware：设置进程 DPI 感知

```go
var procSetProcessDPIAware = user32.NewProc("SetProcessDPIAware")

func EnableDPIAwareOld() bool {
    ret, _, _ := procSetProcessDPIAware.Call()
    return ret != 0
}
```

> 新程序更建议使用 Manifest 或 `SetProcessDpiAwarenessContext`。

---

## 86. SetProcessDpiAwarenessContext：设置 DPI 感知上下文

```go
var procSetProcessDpiAwarenessContext = user32.NewProc("SetProcessDpiAwarenessContext")

var DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2 = ^uintptr(3) // -4

func EnablePerMonitorDPIAwareV2() bool {
    ret, _, _ := procSetProcessDpiAwarenessContext.Call(DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2)
    return ret != 0
}
```

---

## 87. AdjustWindowRectEx / AdjustWindowRectExForDpi：根据客户区计算窗口大小

```go
var (
    procAdjustWindowRectEx = user32.NewProc("AdjustWindowRectEx")
    procAdjustWindowRectExForDpi = user32.NewProc("AdjustWindowRectExForDpi")
)

func CalcWindowRectForClient(w, h int32, style, exStyle uint32) RECT {
    r := RECT{0, 0, w, h}
    procAdjustWindowRectEx.Call(uintptr(unsafe.Pointer(&r)), uintptr(style), 0, uintptr(exStyle))
    return r
}

func CalcWindowRectForClientDPI(w, h int32, style, exStyle, dpi uint32) RECT {
    r := RECT{0, 0, w, h}
    procAdjustWindowRectExForDpi.Call(uintptr(unsafe.Pointer(&r)), uintptr(style), 0, uintptr(exStyle), uintptr(dpi))
    return r
}
```

---

## 88. SetWindowsHookExW / CallNextHookEx / UnhookWindowsHookEx：低级键盘 Hook

### 场景

监听全局键盘事件。

```go
var (
    procSetWindowsHookExW = user32.NewProc("SetWindowsHookExW")
    procCallNextHookEx = user32.NewProc("CallNextHookEx")
    procUnhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
)

const WH_KEYBOARD_LL = 13
const WM_KEYDOWN = 0x0100

type KBDLLHOOKSTRUCT struct {
    VkCode      uint32
    ScanCode    uint32
    Flags       uint32
    Time        uint32
    DwExtraInfo uintptr
}

var keyboardHook uintptr
var keyboardHookProc uintptr

func keyboardProc(nCode int, wParam uintptr, lParam uintptr) uintptr {
    if nCode >= 0 && wParam == WM_KEYDOWN {
        info := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
        fmt.Println("按键:", info.VkCode)
    }
    ret, _, _ := procCallNextHookEx.Call(keyboardHook, uintptr(nCode), wParam, lParam)
    return ret
}

func InstallKeyboardHook() bool {
    keyboardHookProc = syscall.NewCallback(keyboardProc)
    h, _, _ := procSetWindowsHookExW.Call(
        WH_KEYBOARD_LL,
        keyboardHookProc,
        0,
        0,
    )
    keyboardHook = h
    return h != 0
}

func UninstallKeyboardHook() {
    if keyboardHook != 0 {
        procUnhookWindowsHookEx.Call(keyboardHook)
        keyboardHook = 0
    }
}
```

完整运行时要调用：

```go
func main() {
    InstallKeyboardHook()
    defer UninstallKeyboardHook()
    MessageLoop()
}
```

---

## 89. SetWindowsHookExW：低级鼠标 Hook

```go
const WH_MOUSE_LL = 14
const WM_LBUTTONDOWN = 0x0201

type MSLLHOOKSTRUCT struct {
    Pt          POINT
    MouseData   uint32
    Flags       uint32
    Time        uint32
    DwExtraInfo uintptr
}

var mouseHook uintptr
var mouseHookProc uintptr

func mouseProc(nCode int, wParam uintptr, lParam uintptr) uintptr {
    if nCode >= 0 && wParam == WM_LBUTTONDOWN {
        info := (*MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))
        fmt.Println("鼠标左键:", info.Pt.X, info.Pt.Y)
    }
    ret, _, _ := procCallNextHookEx.Call(mouseHook, uintptr(nCode), wParam, lParam)
    return ret
}

func InstallMouseHook() bool {
    mouseHookProc = syscall.NewCallback(mouseProc)
    h, _, _ := procSetWindowsHookExW.Call(WH_MOUSE_LL, mouseHookProc, 0, 0)
    mouseHook = h
    return h != 0
}
```

---

## 90. SetWinEventHook / UnhookWinEvent：监听前台窗口变化

```go
var (
    procSetWinEventHook = user32.NewProc("SetWinEventHook")
    procUnhookWinEvent  = user32.NewProc("UnhookWinEvent")
)

const (
    EVENT_SYSTEM_FOREGROUND = 0x0003
    WINEVENT_OUTOFCONTEXT   = 0x0000
)

var winEventHook uintptr
var winEventProc uintptr

func onWinEvent(hWinEventHook uintptr, event uint32, hwnd uintptr, idObject int32, idChild int32, dwEventThread uint32, dwmsEventTime uint32) uintptr {
    fmt.Println("前台窗口变化:", hwnd, GetWindowTitle(hwnd))
    return 0
}

func InstallForegroundEventHook() bool {
    winEventProc = syscall.NewCallback(onWinEvent)
    h, _, _ := procSetWinEventHook.Call(
        EVENT_SYSTEM_FOREGROUND,
        EVENT_SYSTEM_FOREGROUND,
        0,
        winEventProc,
        0,
        0,
        WINEVENT_OUTOFCONTEXT,
    )
    winEventHook = h
    return h != 0
}

func UninstallWinEventHook() {
    if winEventHook != 0 {
        procUnhookWinEvent.Call(winEventHook)
    }
}
```

---

## 91. RegisterRawInputDevices / GetRawInputData：原始输入

### 场景

让窗口接收 `WM_INPUT`，再读取原始输入数据。

```go
var (
    procRegisterRawInputDevices = user32.NewProc("RegisterRawInputDevices")
    procGetRawInputData = user32.NewProc("GetRawInputData")
)

const (
    WM_INPUT = 0x00FF
    RIDEV_INPUTSINK = 0x00000100
    RID_INPUT = 0x10000003
)

type RAWINPUTDEVICE struct {
    UsUsagePage uint16
    UsUsage     uint16
    DwFlags     uint32
    HwndTarget  uintptr
}

func RegisterRawKeyboard(hwnd uintptr) bool {
    rid := RAWINPUTDEVICE{
        UsUsagePage: 0x01,
        UsUsage:     0x06, // keyboard
        DwFlags:     RIDEV_INPUTSINK,
        HwndTarget:  hwnd,
    }
    ret, _, _ := procRegisterRawInputDevices.Call(
        uintptr(unsafe.Pointer(&rid)),
        1,
        unsafe.Sizeof(rid),
    )
    return ret != 0
}
```

窗口过程收到 `WM_INPUT` 后，再调用 `GetRawInputData` 获取详细数据。Raw Input 结构较复杂，建议单独成篇。

---

## 92. RegisterTouchWindow / GetTouchInputInfo / CloseTouchInputHandle

```go
var (
    procRegisterTouchWindow = user32.NewProc("RegisterTouchWindow")
    procGetTouchInputInfo = user32.NewProc("GetTouchInputInfo")
    procCloseTouchInputHandle = user32.NewProc("CloseTouchInputHandle")
)

const WM_TOUCH = 0x0240

func RegisterTouch(hwnd uintptr) bool {
    ret, _, _ := procRegisterTouchWindow.Call(hwnd, 0)
    return ret != 0
}
```

收到 `WM_TOUCH` 后：

```go
// wParam 低位是触摸点数量，lParam 是 HTOUCHINPUT
// 调用 GetTouchInputInfo 读取 TOUCHINPUT 数组，最后 CloseTouchInputHandle(lParam)
```

---

## 93. RegisterPointerInputTarget / GetPointerInfo

```go
var (
    procRegisterPointerInputTarget = user32.NewProc("RegisterPointerInputTarget")
    procGetPointerInfo = user32.NewProc("GetPointerInfo")
)

const PT_TOUCH = 0x00000002

func RegisterPointerTouch(hwnd uintptr) bool {
    ret, _, _ := procRegisterPointerInputTarget.Call(hwnd, PT_TOUCH)
    return ret != 0
}
```

Pointer API 结构体较大，适合单独拆一篇。

---

## 94. GetKeyboardLayout / ActivateKeyboardLayout

```go
var (
    procGetKeyboardLayout = user32.NewProc("GetKeyboardLayout")
    procActivateKeyboardLayout = user32.NewProc("ActivateKeyboardLayout")
)

func CurrentKeyboardLayout() uintptr {
    hkl, _, _ := procGetKeyboardLayout.Call(0)
    return hkl
}

func ActivateLayout(hkl uintptr) uintptr {
    old, _, _ := procActivateKeyboardLayout.Call(hkl, 0)
    return old
}
```

---

## 95. LoadKeyboardLayoutW / UnloadKeyboardLayout

```go
var (
    procLoadKeyboardLayoutW = user32.NewProc("LoadKeyboardLayoutW")
    procUnloadKeyboardLayout = user32.NewProc("UnloadKeyboardLayout")
)

const KLF_ACTIVATE = 0x00000001

func LoadUSKeyboard() uintptr {
    // 00000409 = English US
    hkl, _, _ := procLoadKeyboardLayoutW.Call(utf16Ptr("00000409"), KLF_ACTIVATE)
    return hkl
}

func UnloadLayout(hkl uintptr) bool {
    ret, _, _ := procUnloadKeyboardLayout.Call(hkl)
    return ret != 0
}
```

---

## 96. VkKeyScanW：字符转虚拟键

```go
var procVkKeyScanW = user32.NewProc("VkKeyScanW")

func CharToVK(ch rune) int16 {
    ret, _, _ := procVkKeyScanW.Call(uintptr(ch))
    return int16(ret)
}
```

---

## 97. DragDetect：检测拖动意图

```go
var procDragDetect = user32.NewProc("DragDetect")

func IsDrag(hwnd uintptr, x, y int32) bool {
    packed := uintptr(uint32(x)) | uintptr(uint64(uint32(y))<<32)
    ret, _, _ := procDragDetect.Call(hwnd, packed)
    return ret != 0
}
```

---

## 98. DrawFocusRect：绘制焦点矩形

```go
var procDrawFocusRect = user32.NewProc("DrawFocusRect")

func DrawFocus(hdc uintptr, r RECT) bool {
    ret, _, _ := procDrawFocusRect.Call(hdc, uintptr(unsafe.Pointer(&r)))
    return ret != 0
}
```

---

## 99. DrawFrameControl：绘制标准按钮/控件外观

```go
var procDrawFrameControl = user32.NewProc("DrawFrameControl")

const (
    DFC_BUTTON = 4
    DFCS_BUTTONPUSH = 0x0010
)

func DrawPushButtonFrame(hdc uintptr, r RECT) bool {
    ret, _, _ := procDrawFrameControl.Call(
        hdc,
        uintptr(unsafe.Pointer(&r)),
        DFC_BUTTON,
        DFCS_BUTTONPUSH,
    )
    return ret != 0
}
```

---

## 100. 滚动条：GetScrollInfo / SetScrollInfo / ShowScrollBar

```go
var (
    procGetScrollInfo = user32.NewProc("GetScrollInfo")
    procSetScrollInfo = user32.NewProc("SetScrollInfo")
    procShowScrollBar = user32.NewProc("ShowScrollBar")
)

type SCROLLINFO struct {
    CbSize    uint32
    FMask     uint32
    NMin      int32
    NMax      int32
    NPage     uint32
    NPos      int32
    NTrackPos int32
}

const (
    SB_VERT = 1
    SIF_ALL = 0x17
)

func ShowVerticalScrollBar(hwnd uintptr, show bool) {
    v := uintptr(0)
    if show { v = 1 }
    procShowScrollBar.Call(hwnd, SB_VERT, v)
}
```

---

## 101. Dialog：GetDlgItem / SetDlgItemTextW / GetDlgItemTextW

```go
var (
    procGetDlgItem = user32.NewProc("GetDlgItem")
    procSetDlgItemTextW = user32.NewProc("SetDlgItemTextW")
    procGetDlgItemTextW = user32.NewProc("GetDlgItemTextW")
)

func SetDialogItemText(hwndDlg uintptr, id int, text string) bool {
    ret, _, _ := procSetDlgItemTextW.Call(hwndDlg, uintptr(id), utf16Ptr(text))
    return ret != 0
}

func GetDialogItemText(hwndDlg uintptr, id int) string {
    buf := make([]uint16, 256)
    procGetDlgItemTextW.Call(hwndDlg, uintptr(id), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
    return syscall.UTF16ToString(buf)
}
```

---

## 102. CheckDlgButton / IsDlgButtonChecked

```go
var (
    procCheckDlgButton = user32.NewProc("CheckDlgButton")
    procIsDlgButtonChecked = user32.NewProc("IsDlgButtonChecked")
)

const BST_CHECKED = 1
const BST_UNCHECKED = 0

func SetCheckbox(hwndDlg uintptr, id int, checked bool) {
    v := BST_UNCHECKED
    if checked { v = BST_CHECKED }
    procCheckDlgButton.Call(hwndDlg, uintptr(id), uintptr(v))
}

func IsCheckboxChecked(hwndDlg uintptr, id int) bool {
    ret, _, _ := procIsDlgButtonChecked.Call(hwndDlg, uintptr(id))
    return ret == BST_CHECKED
}
```

---

## 103. CallWindowProcW：调用原窗口过程

### 场景

子类化窗口后，在自定义窗口过程里继续调用原窗口过程。

```go
var procCallWindowProcW = user32.NewProc("CallWindowProcW")

func CallOldWndProc(oldProc uintptr, hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
    ret, _, _ := procCallWindowProcW.Call(oldProc, hwnd, uintptr(msg), wParam, lParam)
    return ret
}
```

---

## 104. CallMsgFilterW：调用消息过滤 Hook

```go
var procCallMsgFilterW = user32.NewProc("CallMsgFilterW")

func CallMessageFilter(msg *MSG, code int) bool {
    ret, _, _ := procCallMsgFilterW.Call(uintptr(unsafe.Pointer(msg)), uintptr(code))
    return ret != 0
}
```

---

## 105. BroadcastSystemMessageW：广播系统消息

```go
var procBroadcastSystemMessageW = user32.NewProc("BroadcastSystemMessageW")

const BSF_POSTMESSAGE = 0x00000010
const BSM_APPLICATIONS = 0x00000008

func BroadcastToApplications(message uint32, wParam, lParam uintptr) int32 {
    recipients := uint32(BSM_APPLICATIONS)
    ret, _, _ := procBroadcastSystemMessageW.Call(
        BSF_POSTMESSAGE,
        uintptr(unsafe.Pointer(&recipients)),
        uintptr(message),
        wParam,
        lParam,
    )
    return int32(ret)
}
```

---

## 106. ChangeWindowMessageFilterEx：修改窗口消息过滤

用于 UIPI 场景，让低权限进程的某些消息能发到高权限窗口。此类 API 安全敏感。

```go
var procChangeWindowMessageFilterEx = user32.NewProc("ChangeWindowMessageFilterEx")

const MSGFLT_ALLOW = 1

func AllowMessage(hwnd uintptr, msg uint32) bool {
    ret, _, _ := procChangeWindowMessageFilterEx.Call(hwnd, uintptr(msg), MSGFLT_ALLOW, 0)
    return ret != 0
}
```

---

## 107. AttachThreadInput：连接两个线程输入队列

### 场景

有时为了设置焦点，需要临时把当前线程和目标窗口线程输入队列关联。

```go
var procAttachThreadInput = user32.NewProc("AttachThreadInput")

func AttachInput(srcThread, dstThread uint32, attach bool) bool {
    v := uintptr(0)
    if attach { v = 1 }
    ret, _, _ := procAttachThreadInput.Call(uintptr(srcThread), uintptr(dstThread), v)
    return ret != 0
}
```

使用后一定要 detach。

---

## 108. AllowSetForegroundWindow：允许指定进程设置前台窗口

```go
var procAllowSetForegroundWindow = user32.NewProc("AllowSetForegroundWindow")

func AllowProcessSetForeground(pid uint32) bool {
    ret, _, _ := procAllowSetForegroundWindow.Call(uintptr(pid))
    return ret != 0
}
```

---

## 109. BlockInput：阻止键鼠输入

```go
var procBlockInput = user32.NewProc("BlockInput")

func BlockUserInput(block bool) bool {
    v := uintptr(0)
    if block { v = 1 }
    ret, _, _ := procBlockInput.Call(v)
    return ret != 0
}
```

> 高风险：可能让用户暂时无法操作电脑。一定要确保能恢复。

---

## 110. 总结：不同任务该用哪些 API

### 查找并控制已有窗口

```text
FindWindowW
EnumWindows
GetWindowTextW
GetClassNameW
GetWindowThreadProcessId
ShowWindow
SetForegroundWindow
SetWindowPos
PostMessageW
```

### 做桌面自动化

```text
GetForegroundWindow
GetWindowRect
WindowFromPoint
SendInput
GetAsyncKeyState
RegisterHotKey
SetWinEventHook
```

### 做剪贴板工具

```text
OpenClipboard
CloseClipboard
IsClipboardFormatAvailable
GetClipboardData
EmptyClipboard
SetClipboardData
AddClipboardFormatListener
```

### 做 GUI 程序

```text
RegisterClassExW
CreateWindowExW
DefWindowProcW
ShowWindow
GetMessageW
TranslateMessage
DispatchMessageW
BeginPaint
EndPaint
CreateMenu
SetTimer
```

### 做全局监听

```text
SetWindowsHookExW
CallNextHookEx
UnhookWindowsHookEx
SetWinEventHook
UnhookWinEvent
RegisterRawInputDevices
```

---

## 111. 后续建议拆分

如果要做到真正“Microsoft Learn 上 winuser.h 每一个 API 都有完整案例”，建议拆成这些文件：

1. `Go user32.dll 窗口 API 完整案例.md`
2. `Go user32.dll 消息 API 完整案例.md`
3. `Go user32.dll 输入 API 完整案例.md`
4. `Go user32.dll 剪贴板 API 完整案例.md`
5. `Go user32.dll Hook API 完整案例.md`
6. `Go user32.dll 菜单和对话框 API 完整案例.md`
7. `Go user32.dll DPI 和显示器 API 完整案例.md`
8. `Go user32.dll 系统参数 API 完整案例.md`

这样更容易维护，也不会单个 md 文件过大。
