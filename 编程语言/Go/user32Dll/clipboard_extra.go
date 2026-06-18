//go:build windows

package user32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procCountClipboardFormats      = user32.NewProc("CountClipboardFormats")
	procEnumClipboardFormats       = user32.NewProc("EnumClipboardFormats")
	procRegisterClipboardFormatW   = user32.NewProc("RegisterClipboardFormatW")
	procGetClipboardFormatNameW    = user32.NewProc("GetClipboardFormatNameW")
	procGetClipboardOwner          = user32.NewProc("GetClipboardOwner")
	procGetOpenClipboardWindow     = user32.NewProc("GetOpenClipboardWindow")
	procGetClipboardSequenceNumber = user32.NewProc("GetClipboardSequenceNumber")
)

// CountClipboardFormats 获取剪贴板当前可用格式数量。
// 返回值：返回剪贴板格式数量；返回 0 表示没有格式或调用失败。
func CountClipboardFormats() int {
	ret, _, _ := procCountClipboardFormats.Call()
	return int(ret)
}

// EnumClipboardFormats 枚举剪贴板格式。
// 参数format：上一次返回的格式 ID，首次调用传 0。
// 返回值：返回下一个剪贴板格式 ID；返回 0 表示枚举结束或失败。
func EnumClipboardFormats(format uint32) uint32 {
	ret, _, _ := procEnumClipboardFormats.Call(uintptr(format))
	return uint32(ret)
}

// RegisterClipboardFormat 注册自定义剪贴板格式名称。
// 参数name：剪贴板格式名称。
// 返回值：返回注册得到的格式 ID；返回 0 表示注册失败。
func RegisterClipboardFormat(name string) uint32 {
	ret, _, _ := procRegisterClipboardFormatW.Call(utf16Ptr(name))
	return uint32(ret)
}

// GetClipboardFormatName 获取剪贴板格式名称。
// 参数format：剪贴板格式 ID。
// 返回值：返回格式名称；标准格式、未知格式或读取失败时返回空字符串。
func GetClipboardFormatName(format uint32) string {
	buf := make([]uint16, 256)
	ret, _, _ := procGetClipboardFormatNameW.Call(uintptr(format), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetClipboardOwner 获取当前剪贴板所有者窗口。
// 返回值：返回剪贴板所有者窗口句柄；返回 0 表示没有所有者或获取失败。
func GetClipboardOwner() uintptr {
	ret, _, _ := procGetClipboardOwner.Call()
	return ret
}

// GetOpenClipboardWindow 获取当前打开剪贴板的窗口。
// 返回值：返回打开剪贴板的窗口句柄；返回 0 表示剪贴板未打开或获取失败。
func GetOpenClipboardWindow() uintptr {
	ret, _, _ := procGetOpenClipboardWindow.Call()
	return ret
}

// GetClipboardSequenceNumber 获取剪贴板内容变化序号。
// 返回值：返回剪贴板序号；序号变大表示剪贴板内容发生过变化。
func GetClipboardSequenceNumber() uint32 {
	ret, _, _ := procGetClipboardSequenceNumber.Call()
	return uint32(ret)
}
