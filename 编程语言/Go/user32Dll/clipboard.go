//go:build windows

package user32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procOpenClipboard                 = user32.NewProc("OpenClipboard")
	procCloseClipboard                = user32.NewProc("CloseClipboard")
	procIsClipboardFormatAvailable    = user32.NewProc("IsClipboardFormatAvailable")
	procGetClipboardData              = user32.NewProc("GetClipboardData")
	procEmptyClipboard                = user32.NewProc("EmptyClipboard")
	procSetClipboardData              = user32.NewProc("SetClipboardData")
	procAddClipboardFormatListener    = user32.NewProc("AddClipboardFormatListener")
	procRemoveClipboardFormatListener = user32.NewProc("RemoveClipboardFormatListener")
	procGlobalLock                    = kernel32.NewProc("GlobalLock")
	procGlobalUnlock                  = kernel32.NewProc("GlobalUnlock")
	procGlobalAlloc                   = kernel32.NewProc("GlobalAlloc")
)

const (
	CF_UNICODETEXT     = 13
	GMEM_MOVEABLE      = 0x0002
	WM_CLIPBOARDUPDATE = 0x031D
)

// WithClipboard 打开剪贴板并在回调结束后自动关闭。
// 参数fn：持有资源期间执行的回调函数。
// 返回值：true 表示成功打开剪贴板并执行回调，false 表示剪贴板打开失败。
func WithClipboard(fn func()) bool {
	ret, _, _ := procOpenClipboard.Call(0)
	if ret == 0 {
		return false
	}
	defer procCloseClipboard.Call()
	fn()
	return true
}

// HasUnicodeText 判断剪贴板是否包含 Unicode 文本。
// 返回值：true 表示剪贴板当前包含 Unicode 文本，false 表示没有该格式或检查失败。
func HasUnicodeText() bool {
	ret, _, _ := procIsClipboardFormatAvailable.Call(CF_UNICODETEXT)
	return ret != 0
}

// GetClipboardText 读取剪贴板中的 Unicode 文本。
// 返回值：返回剪贴板中的 Unicode 文本；没有文本或读取失败时返回空字符串。
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
		ptr := (*[1 << 20]uint16)(unsafe.Pointer(p))
		n := 0
		for ptr[n] != 0 {
			n++
		}
		text = syscall.UTF16ToString(ptr[:n])
	})
	return text
}

// SetClipboardText 写入 Unicode 文本到剪贴板。
// 参数text：要传入的文本内容。
// 返回值：true 表示文本已写入剪贴板，false 表示写入失败。
func SetClipboardText(text string) bool {
	utf16, err := syscall.UTF16FromString(text)
	if err != nil {
		return false
	}
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

// AddClipboardListener 注册窗口监听剪贴板变化。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口已成功注册剪贴板变化监听，false 表示注册失败。
func AddClipboardListener(hwnd uintptr) bool {
	ret, _, _ := procAddClipboardFormatListener.Call(hwnd)
	return ret != 0
}

// RemoveClipboardListener 移除窗口剪贴板变化监听。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口已成功移除剪贴板监听，false 表示移除失败。
func RemoveClipboardListener(hwnd uintptr) bool {
	ret, _, _ := procRemoveClipboardFormatListener.Call(hwnd)
	return ret != 0
}
