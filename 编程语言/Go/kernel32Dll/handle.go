//go:build windows

package kernel32Dll

import "unsafe"

var (
	procCloseHandle     = kernel32.NewProc("CloseHandle")
	procDuplicateHandle = kernel32.NewProc("DuplicateHandle")
	procGetHandleInfo   = kernel32.NewProc("GetHandleInformation")
	procSetHandleInfo   = kernel32.NewProc("SetHandleInformation")
	procGetLastError    = kernel32.NewProc("GetLastError")
	procSetLastError    = kernel32.NewProc("SetLastError")
	procFormatMessageW  = kernel32.NewProc("FormatMessageW")
	procLocalFree       = kernel32.NewProc("LocalFree")
)

const (
	DUPLICATE_CLOSE_SOURCE = 0x00000001
	DUPLICATE_SAME_ACCESS  = 0x00000002

	HANDLE_FLAG_INHERIT            = 0x00000001
	HANDLE_FLAG_PROTECT_FROM_CLOSE = 0x00000002

	FORMAT_MESSAGE_ALLOCATE_BUFFER = 0x00000100
	FORMAT_MESSAGE_FROM_SYSTEM     = 0x00001000
	FORMAT_MESSAGE_IGNORE_INSERTS  = 0x00000200
)

// CloseHandle 关闭内核对象句柄。
// 参数handle：需要关闭的内核对象句柄。
// 返回值：true 表示句柄关闭成功，false 表示关闭失败。
func CloseHandle(handle uintptr) bool {
	ret, _, _ := procCloseHandle.Call(handle)
	return ret != 0
}

// DuplicateHandleForProcess 复制内核对象句柄到目标进程。
// 参数sourceProcess：源进程句柄。
// 参数sourceHandle：源对象句柄。
// 参数targetProcess：目标进程句柄。
// 参数desiredAccess：目标句柄访问权限。
// 参数inheritHandle：目标句柄是否可继承。
// 参数options：复制选项，例如 DUPLICATE_SAME_ACCESS。
// 返回值：返回新句柄和是否复制成功；true 表示句柄有效，false 表示复制失败。
func DuplicateHandleForProcess(sourceProcess, sourceHandle, targetProcess uintptr, desiredAccess uint32, inheritHandle bool, options uint32) (uintptr, bool) {
	var targetHandle uintptr
	ret, _, _ := procDuplicateHandle.Call(sourceProcess, sourceHandle, targetProcess, uintptr(unsafe.Pointer(&targetHandle)), uintptr(desiredAccess), boolArg(inheritHandle), uintptr(options))
	return targetHandle, ret != 0
}

// GetHandleInformation 获取句柄标志。
// 参数handle：内核对象句柄。
// 返回值：返回句柄标志和是否获取成功；true 表示 flags 有效，false 表示获取失败。
func GetHandleInformation(handle uintptr) (flags uint32, ok bool) {
	ret, _, _ := procGetHandleInfo.Call(handle, uintptr(unsafe.Pointer(&flags)))
	return flags, ret != 0
}

// SetHandleInformation 设置句柄标志。
// 参数handle：内核对象句柄。
// 参数mask：要修改的标志掩码。
// 参数flags：新的标志值。
// 返回值：true 表示句柄标志设置成功，false 表示设置失败。
func SetHandleInformation(handle uintptr, mask, flags uint32) bool {
	ret, _, _ := procSetHandleInfo.Call(handle, uintptr(mask), uintptr(flags))
	return ret != 0
}

// GetLastErrorCode 获取当前线程最后一个 Win32 错误码。
// 返回值：返回当前线程最后一个错误码；0 通常表示没有错误。
func GetLastErrorCode() uint32 {
	ret, _, _ := procGetLastError.Call()
	return uint32(ret)
}

// SetLastErrorCode 设置当前线程最后一个 Win32 错误码。
// 参数code：要设置的错误码。
// 返回值：无；函数只更新当前线程错误码。
func SetLastErrorCode(code uint32) {
	procSetLastError.Call(uintptr(code))
}
