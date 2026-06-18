//go:build windows

package kernel32Dll

import (
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32.dll")

const (
	INFINITE = 0xFFFFFFFF

	WAIT_OBJECT_0 = 0x00000000
	WAIT_TIMEOUT  = 0x00000102
	WAIT_FAILED   = 0xFFFFFFFF

	TRUE  = 1
	FALSE = 0
)

// SECURITY_ATTRIBUTES 表示 Win32 安全属性结构。
type SECURITY_ATTRIBUTES struct {
	NLength              uint32
	LpSecurityDescriptor uintptr
	BInheritHandle       int32
}

// FILETIME 表示 Win32 文件时间。
type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

// SYSTEMTIME 表示 Win32 系统时间。
type SYSTEMTIME struct {
	WYear         uint16
	WMonth        uint16
	WDayOfWeek    uint16
	WDay          uint16
	WHour         uint16
	WMinute       uint16
	WSecond       uint16
	WMilliseconds uint16
}

// OVERLAPPED 表示 Win32 异步 I/O 重叠结构。
type OVERLAPPED struct {
	Internal     uintptr
	InternalHigh uintptr
	Offset       uint32
	OffsetHigh   uint32
	HEvent       uintptr
}

// utf16Ptr 将 Go 字符串转换为 Win32 UTF-16 指针。
// 参数s：要转换的 Go 字符串。
// 返回值：返回 UTF-16 字符串指针；字符串包含非法空字符时会 panic。
func utf16Ptr(s string) uintptr {
	p, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		panic(err)
	}
	return uintptr(unsafe.Pointer(p))
}

// utf16PtrOrNil 将空字符串转换为空指针，非空字符串转换为 UTF-16 指针。
// 参数s：要转换的 Go 字符串。
// 返回值：返回 UTF-16 字符串指针；s 为空时返回 0。
func utf16PtrOrNil(s string) uintptr {
	if s == "" {
		return 0
	}
	return utf16Ptr(s)
}

// boolArg 将 Go 布尔值转换为 Win32 BOOL 参数。
// 参数v：Go 布尔值。
// 返回值：true 转换为 1，false 转换为 0。
func boolArg(v bool) uintptr {
	if v {
		return 1
	}
	return 0
}

// saPtr 将安全属性指针转换为 syscall 调用参数。
// 参数sa：安全属性结构指针，nil 表示默认安全属性。
// 返回值：返回可传给 Win32 API 的指针值；sa 为 nil 时返回 0。
func saPtr(sa *SECURITY_ATTRIBUTES) uintptr {
	if sa == nil {
		return 0
	}
	return uintptr(unsafe.Pointer(sa))
}
