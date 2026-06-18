//go:build windows

package user32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procSetCapture         = user32.NewProc("SetCapture")
	procReleaseCapture     = user32.NewProc("ReleaseCapture")
	procGetCapture         = user32.NewProc("GetCapture")
	procClipCursor         = user32.NewProc("ClipCursor")
	procGetClipCursor      = user32.NewProc("GetClipCursor")
	procCreateCaret        = user32.NewProc("CreateCaret")
	procShowCaret          = user32.NewProc("ShowCaret")
	procHideCaret          = user32.NewProc("HideCaret")
	procDestroyCaret       = user32.NewProc("DestroyCaret")
	procSetCaretPos        = user32.NewProc("SetCaretPos")
	procGetCaretPos        = user32.NewProc("GetCaretPos")
	procGetDoubleClickTime = user32.NewProc("GetDoubleClickTime")
	procSetDoubleClickTime = user32.NewProc("SetDoubleClickTime")
	procSwapMouseButton    = user32.NewProc("SwapMouseButton")
	procToUnicodeEx        = user32.NewProc("ToUnicodeEx")
	procMapVirtualKeyExW   = user32.NewProc("MapVirtualKeyExW")
	procGetKeyNameTextW    = user32.NewProc("GetKeyNameTextW")
	procOemKeyScan         = user32.NewProc("OemKeyScan")
)

// SetCapture 捕获鼠标输入到指定窗口。
// 参数hwnd：窗口句柄。
// 返回值：返回之前捕获鼠标输入的窗口句柄；返回 0 表示之前没有捕获窗口。
func SetCapture(hwnd uintptr) uintptr {
	ret, _, _ := procSetCapture.Call(hwnd)
	return ret
}

// ReleaseCapture 释放当前线程的鼠标捕获。
// 返回值：true 表示鼠标捕获释放成功，false 表示释放失败。
func ReleaseCapture() bool {
	ret, _, _ := procReleaseCapture.Call()
	return ret != 0
}

// GetCapture 获取当前捕获鼠标输入的窗口。
// 返回值：返回当前捕获鼠标的窗口句柄；返回 0 表示当前没有捕获窗口。
func GetCapture() uintptr {
	ret, _, _ := procGetCapture.Call()
	return ret
}

// ClipCursor 限制鼠标光标活动区域。
// 参数rect：限制区域指针，nil 表示解除限制。
// 返回值：true 表示光标限制设置成功，false 表示设置失败。
func ClipCursor(rect *RECT) bool {
	var ptr uintptr
	if rect != nil {
		ptr = uintptr(unsafe.Pointer(rect))
	}
	ret, _, _ := procClipCursor.Call(ptr)
	return ret != 0
}

// GetClipCursor 获取当前鼠标光标限制区域。
// 返回值：返回光标限制矩形和是否获取成功；true 表示矩形有效，false 表示获取失败。
func GetClipCursor() (RECT, bool) {
	var r RECT
	ret, _, _ := procGetClipCursor.Call(uintptr(unsafe.Pointer(&r)))
	return r, ret != 0
}

// CreateCaret 为窗口创建插入符。
// 参数hwnd：窗口句柄。
// 参数bitmap：插入符位图句柄，0 表示使用实心插入符。
// 参数width：插入符宽度。
// 参数height：插入符高度。
// 返回值：true 表示插入符创建成功，false 表示创建失败。
func CreateCaret(hwnd uintptr, bitmap uintptr, width, height int32) bool {
	ret, _, _ := procCreateCaret.Call(hwnd, bitmap, uintptr(width), uintptr(height))
	return ret != 0
}

// ShowCaret 显示指定窗口的插入符。
// 参数hwnd：窗口句柄。
// 返回值：true 表示插入符显示成功，false 表示显示失败。
func ShowCaret(hwnd uintptr) bool {
	ret, _, _ := procShowCaret.Call(hwnd)
	return ret != 0
}

// HideCaret 隐藏指定窗口的插入符。
// 参数hwnd：窗口句柄。
// 返回值：true 表示插入符隐藏成功，false 表示隐藏失败。
func HideCaret(hwnd uintptr) bool {
	ret, _, _ := procHideCaret.Call(hwnd)
	return ret != 0
}

// DestroyCaret 销毁当前插入符。
// 返回值：true 表示插入符销毁成功，false 表示销毁失败。
func DestroyCaret() bool {
	ret, _, _ := procDestroyCaret.Call()
	return ret != 0
}

// SetCaretPos 设置插入符位置。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 返回值：true 表示插入符位置设置成功，false 表示设置失败。
func SetCaretPos(x, y int32) bool {
	ret, _, _ := procSetCaretPos.Call(uintptr(x), uintptr(y))
	return ret != 0
}

// GetCaretPos 获取当前插入符位置。
// 返回值：返回插入符坐标和是否获取成功；true 表示坐标有效，false 表示获取失败。
func GetCaretPos() (POINT, bool) {
	var pt POINT
	ret, _, _ := procGetCaretPos.Call(uintptr(unsafe.Pointer(&pt)))
	return pt, ret != 0
}

// GetDoubleClickTime 获取系统双击时间。
// 返回值：返回双击判定时间，单位毫秒。
func GetDoubleClickTime() uint32 {
	ret, _, _ := procGetDoubleClickTime.Call()
	return uint32(ret)
}

// SetDoubleClickTime 设置系统双击时间。
// 参数milliseconds：双击判定时间，单位毫秒。
// 返回值：true 表示双击时间设置成功，false 表示设置失败。
func SetDoubleClickTime(milliseconds uint32) bool {
	ret, _, _ := procSetDoubleClickTime.Call(uintptr(milliseconds))
	return ret != 0
}

// SwapMouseButton 交换或恢复鼠标左右键。
// 参数swap：true 表示交换左右键，false 表示恢复默认。
// 返回值：true 表示调用前左右键已经处于交换状态，false 表示调用前未交换。
func SwapMouseButton(swap bool) bool {
	ret, _, _ := procSwapMouseButton.Call(boolArg(swap))
	return ret != 0
}

// ToUnicodeEx 使用指定键盘布局将按键转换为 Unicode 字符。
// 参数vk：虚拟键码。
// 参数scanCode：键盘扫描码。
// 参数state：键盘状态数组指针。
// 参数hkl：键盘布局句柄。
// 返回值：返回转换得到的 Unicode 字符串；无法转换、死键或失败时返回空字符串。
func ToUnicodeEx(vk, scanCode uint32, state *[256]byte, hkl uintptr) string {
	buf := make([]uint16, 8)
	ret, _, _ := procToUnicodeEx.Call(uintptr(vk), uintptr(scanCode), uintptr(unsafe.Pointer(&state[0])), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)), 0, hkl)
	if int32(ret) <= 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// MapVirtualKeyEx 使用指定键盘布局转换虚拟键码或扫描码。
// 参数code：要转换的键码或扫描码。
// 参数mapType：转换类型。
// 参数hkl：键盘布局句柄。
// 返回值：返回转换后的键码或扫描码；返回 0 通常表示无法转换。
func MapVirtualKeyEx(code, mapType uint32, hkl uintptr) uint32 {
	ret, _, _ := procMapVirtualKeyExW.Call(uintptr(code), uintptr(mapType), hkl)
	return uint32(ret)
}

// GetKeyNameText 获取按键名称文本。
// 参数lParam：包含扫描码和扩展键标志的参数。
// 返回值：返回按键名称；读取失败时返回空字符串。
func GetKeyNameText(lParam uintptr) string {
	buf := make([]uint16, 128)
	ret, _, _ := procGetKeyNameTextW.Call(lParam, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// OemKeyScan 将 OEM 字符转换为扫描码和 Shift 状态。
// 参数ch：要转换的字符。
// 返回值：返回扫描码和 Shift 状态组合；返回 0xFFFFFFFF 表示转换失败。
func OemKeyScan(ch uint16) uint32 {
	ret, _, _ := procOemKeyScan.Call(uintptr(ch))
	return uint32(ret)
}
