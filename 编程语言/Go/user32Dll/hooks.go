//go:build windows

package user32Dll

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	procSetWindowsHookExW   = user32.NewProc("SetWindowsHookExW")
	procCallNextHookEx      = user32.NewProc("CallNextHookEx")
	procUnhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	procSetWinEventHook     = user32.NewProc("SetWinEventHook")
	procUnhookWinEvent      = user32.NewProc("UnhookWinEvent")
)

const (
	WH_KEYBOARD_LL = 13
	WH_MOUSE_LL    = 14

	WM_KEYDOWN     = 0x0100
	WM_LBUTTONDOWN = 0x0201

	EVENT_SYSTEM_FOREGROUND = 0x0003
	WINEVENT_OUTOFCONTEXT   = 0x0000
)

// KBDLLHOOKSTRUCT 表示低级键盘 Hook 回调中的键盘事件。
type KBDLLHOOKSTRUCT struct {
	VkCode      uint32
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DwExtraInfo uintptr
}

// MSLLHOOKSTRUCT 表示低级鼠标 Hook 回调中的鼠标事件。
type MSLLHOOKSTRUCT struct {
	Pt          POINT
	MouseData   uint32
	Flags       uint32
	Time        uint32
	DwExtraInfo uintptr
}

var keyboardHook uintptr
var keyboardHookProc uintptr
var mouseHook uintptr
var mouseHookProc uintptr
var winEventHook uintptr
var winEventProc uintptr

// keyboardProc 处理低级键盘 Hook 回调并输出按键码。
func keyboardProc(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 && wParam == WM_KEYDOWN {
		info := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		fmt.Println("按键:", info.VkCode)
	}
	ret, _, _ := procCallNextHookEx.Call(keyboardHook, uintptr(nCode), wParam, lParam)
	return ret
}

// InstallKeyboardHook 安装全局低级键盘 Hook。
// 返回值：true 表示低级键盘 Hook 安装成功，false 表示安装失败。
func InstallKeyboardHook() bool {
	keyboardHookProc = syscall.NewCallback(keyboardProc)
	h, _, _ := procSetWindowsHookExW.Call(WH_KEYBOARD_LL, keyboardHookProc, 0, 0)
	keyboardHook = h
	return h != 0
}

// UninstallKeyboardHook 卸载全局低级键盘 Hook。
// 返回值：无返回值；函数会在已安装时卸载键盘 Hook。
func UninstallKeyboardHook() {
	if keyboardHook != 0 {
		procUnhookWindowsHookEx.Call(keyboardHook)
		keyboardHook = 0
	}
}

// mouseProc 处理低级鼠标 Hook 回调并输出左键坐标。
func mouseProc(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 && wParam == WM_LBUTTONDOWN {
		info := (*MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		fmt.Println("鼠标左键:", info.Pt.X, info.Pt.Y)
	}
	ret, _, _ := procCallNextHookEx.Call(mouseHook, uintptr(nCode), wParam, lParam)
	return ret
}

// InstallMouseHook 安装全局低级鼠标 Hook。
// 返回值：true 表示低级鼠标 Hook 安装成功，false 表示安装失败。
func InstallMouseHook() bool {
	mouseHookProc = syscall.NewCallback(mouseProc)
	h, _, _ := procSetWindowsHookExW.Call(WH_MOUSE_LL, mouseHookProc, 0, 0)
	mouseHook = h
	return h != 0
}

// UninstallMouseHook 卸载全局低级鼠标 Hook。
// 返回值：无返回值；函数会在已安装时卸载鼠标 Hook。
func UninstallMouseHook() {
	if mouseHook != 0 {
		procUnhookWindowsHookEx.Call(mouseHook)
		mouseHook = 0
	}
}

// onWinEvent 处理前台窗口变化事件。
func onWinEvent(hWinEventHook uintptr, event uint32, hwnd uintptr, idObject int32, idChild int32, dwEventThread uint32, dwmsEventTime uint32) uintptr {
	fmt.Println("前台窗口变化:", hwnd, GetWindowTitle(hwnd))
	return 0
}

// InstallForegroundEventHook 安装前台窗口变化事件 Hook。
// 返回值：true 表示前台窗口事件 Hook 安装成功，false 表示安装失败。
func InstallForegroundEventHook() bool {
	winEventProc = syscall.NewCallback(onWinEvent)
	h, _, _ := procSetWinEventHook.Call(EVENT_SYSTEM_FOREGROUND, EVENT_SYSTEM_FOREGROUND, 0, winEventProc, 0, 0, WINEVENT_OUTOFCONTEXT)
	winEventHook = h
	return h != 0
}

// UninstallWinEventHook 卸载窗口事件 Hook。
// 返回值：无返回值；函数会在已安装时卸载窗口事件 Hook。
func UninstallWinEventHook() {
	if winEventHook != 0 {
		procUnhookWinEvent.Call(winEventHook)
		winEventHook = 0
	}
}
