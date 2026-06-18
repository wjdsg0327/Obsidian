//go:build windows

package user32Dll

import (
	"syscall"
	"unsafe"
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
)

// POINT 表示 Win32 屏幕或客户区坐标点。
type POINT struct {
	X int32
	Y int32
}

// RECT 表示 Win32 矩形区域。
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// utf16Ptr 将 Go 字符串转换为 Win32 UTF-16 指针。
func utf16Ptr(s string) uintptr {
	p, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		panic(err)
	}
	return uintptr(unsafe.Pointer(p))
}

// utf16PtrOrNil 将空字符串转换为空指针，非空字符串转换为 UTF-16 指针。
func utf16PtrOrNil(s string) uintptr {
	if s == "" {
		return 0
	}
	return utf16Ptr(s)
}

// boolArg 将 Go 布尔值转换为 Win32 BOOL 参数。
func boolArg(v bool) uintptr {
	if v {
		return 1
	}
	return 0
}

// packPoint 将两个 int32 坐标打包为 POINT 参数。
func packPoint(x, y int32) uintptr {
	return uintptr(uint32(x)) | uintptr(uint64(uint32(y))<<32)
}
