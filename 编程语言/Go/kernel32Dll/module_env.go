//go:build windows

package kernel32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procLoadLibraryW              = kernel32.NewProc("LoadLibraryW")
	procFreeLibrary               = kernel32.NewProc("FreeLibrary")
	procGetProcAddress            = kernel32.NewProc("GetProcAddress")
	procGetModuleHandleW          = kernel32.NewProc("GetModuleHandleW")
	procGetModuleFileNameW        = kernel32.NewProc("GetModuleFileNameW")
	procGetCurrentDirectoryW      = kernel32.NewProc("GetCurrentDirectoryW")
	procSetCurrentDirectoryW      = kernel32.NewProc("SetCurrentDirectoryW")
	procGetWindowsDirectoryW      = kernel32.NewProc("GetWindowsDirectoryW")
	procGetSystemDirectoryW       = kernel32.NewProc("GetSystemDirectoryW")
	procGetComputerNameW          = kernel32.NewProc("GetComputerNameW")
	procGetEnvironmentVariableW   = kernel32.NewProc("GetEnvironmentVariableW")
	procSetEnvironmentVariableW   = kernel32.NewProc("SetEnvironmentVariableW")
	procExpandEnvironmentStringsW = kernel32.NewProc("ExpandEnvironmentStringsW")
)

// LoadLibrary 加载 DLL 模块。
// 参数path：DLL 文件路径或模块名。
// 返回值：返回模块句柄；返回 0 表示加载失败。
func LoadLibrary(path string) uintptr {
	ret, _, _ := procLoadLibraryW.Call(utf16Ptr(path))
	return ret
}

// FreeLibrary 释放 DLL 模块引用。
// 参数module：模块句柄。
// 返回值：true 表示模块引用释放成功，false 表示释放失败。
func FreeLibrary(module uintptr) bool {
	ret, _, _ := procFreeLibrary.Call(module)
	return ret != 0
}

// GetProcAddress 获取模块导出函数地址。
// 参数module：模块句柄。
// 参数name：导出函数名称。
// 返回值：返回函数地址；返回 0 表示未找到或获取失败。
func GetProcAddress(module uintptr, name string) uintptr {
	nameBytes := append([]byte(name), 0)
	ret, _, _ := procGetProcAddress.Call(module, uintptr(unsafe.Pointer(&nameBytes[0])))
	return ret
}

// GetModuleHandle 获取已加载模块句柄。
// 参数moduleName：模块名；为空表示当前可执行模块。
// 返回值：返回模块句柄；返回 0 表示未找到或获取失败。
func GetModuleHandle(moduleName string) uintptr {
	ret, _, _ := procGetModuleHandleW.Call(utf16PtrOrNil(moduleName))
	return ret
}

// GetModuleFileName 获取模块文件路径。
// 参数module：模块句柄；传 0 表示当前可执行文件。
// 返回值：返回模块文件完整路径；获取失败时返回空字符串。
func GetModuleFileName(module uintptr) string {
	buf := make([]uint16, 1024)
	ret, _, _ := procGetModuleFileNameW.Call(module, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetCurrentDirectory 获取当前工作目录。
// 返回值：返回当前工作目录；获取失败时返回空字符串。
func GetCurrentDirectory() string {
	buf := make([]uint16, 1024)
	ret, _, _ := procGetCurrentDirectoryW.Call(uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 || int(ret) > len(buf) {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// SetCurrentDirectory 设置当前工作目录。
// 参数path：新的工作目录路径。
// 返回值：true 表示工作目录设置成功，false 表示设置失败。
func SetCurrentDirectory(path string) bool {
	ret, _, _ := procSetCurrentDirectoryW.Call(utf16Ptr(path))
	return ret != 0
}

// GetWindowsDirectory 获取 Windows 目录。
// 返回值：返回 Windows 目录路径；获取失败时返回空字符串。
func GetWindowsDirectory() string {
	buf := make([]uint16, 260)
	ret, _, _ := procGetWindowsDirectoryW.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 || int(ret) > len(buf) {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetSystemDirectory 获取 System32 系统目录。
// 返回值：返回系统目录路径；获取失败时返回空字符串。
func GetSystemDirectory() string {
	buf := make([]uint16, 260)
	ret, _, _ := procGetSystemDirectoryW.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 || int(ret) > len(buf) {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetComputerName 获取本机计算机名。
// 返回值：返回计算机名和是否获取成功；true 表示名称有效，false 表示获取失败。
func GetComputerName() (string, bool) {
	buf := make([]uint16, 256)
	size := uint32(len(buf))
	ret, _, _ := procGetComputerNameW.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&size)))
	return syscall.UTF16ToString(buf[:size]), ret != 0
}

// GetEnvironmentVariable 获取环境变量值。
// 参数name：环境变量名称。
// 返回值：返回环境变量值和是否存在；true 表示变量存在且 value 有效。
func GetEnvironmentVariable(name string) (value string, ok bool) {
	buf := make([]uint16, 32767)
	ret, _, _ := procGetEnvironmentVariableW.Call(utf16Ptr(name), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 || int(ret) > len(buf) {
		return "", false
	}
	return syscall.UTF16ToString(buf[:ret]), true
}

// SetEnvironmentVariable 设置或删除环境变量。
// 参数name：环境变量名称。
// 参数value：环境变量值；空字符串表示设置为空值。
// 返回值：true 表示环境变量设置成功，false 表示设置失败。
func SetEnvironmentVariable(name, value string) bool {
	ret, _, _ := procSetEnvironmentVariableW.Call(utf16Ptr(name), utf16Ptr(value))
	return ret != 0
}

// DeleteEnvironmentVariable 删除环境变量。
// 参数name：环境变量名称。
// 返回值：true 表示环境变量删除成功，false 表示删除失败。
func DeleteEnvironmentVariable(name string) bool {
	ret, _, _ := procSetEnvironmentVariableW.Call(utf16Ptr(name), 0)
	return ret != 0
}

// ExpandEnvironmentStrings 展开字符串中的环境变量引用。
// 参数text：包含环境变量引用的字符串，例如 %TEMP%。
// 返回值：返回展开后的字符串；展开失败时返回空字符串。
func ExpandEnvironmentStrings(text string) string {
	buf := make([]uint16, 32767)
	ret, _, _ := procExpandEnvironmentStringsW.Call(utf16Ptr(text), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 || int(ret) > len(buf) {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret-1])
}
