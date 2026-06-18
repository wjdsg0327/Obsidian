//go:build windows

package kernel32Dll

import "unsafe"

var (
	procGetTickCount64            = kernel32.NewProc("GetTickCount64")
	procQueryPerformanceCounter   = kernel32.NewProc("QueryPerformanceCounter")
	procQueryPerformanceFrequency = kernel32.NewProc("QueryPerformanceFrequency")
	procGetSystemTime             = kernel32.NewProc("GetSystemTime")
	procGetLocalTime              = kernel32.NewProc("GetLocalTime")
	procSystemTimeToFileTime      = kernel32.NewProc("SystemTimeToFileTime")
	procFileTimeToSystemTime      = kernel32.NewProc("FileTimeToSystemTime")
	procGetSystemInfo             = kernel32.NewProc("GetSystemInfo")
	procGetNativeSystemInfo       = kernel32.NewProc("GetNativeSystemInfo")
	procGetVersion                = kernel32.NewProc("GetVersion")
	procGetVersionExW             = kernel32.NewProc("GetVersionExW")
	procGetSystemPowerStatus      = kernel32.NewProc("GetSystemPowerStatus")
)

// SYSTEM_INFO 表示系统处理器和内存页等基础信息。
type SYSTEM_INFO struct {
	WProcessorArchitecture      uint16
	WReserved                   uint16
	DwPageSize                  uint32
	LpMinimumApplicationAddress uintptr
	LpMaximumApplicationAddress uintptr
	DwActiveProcessorMask       uintptr
	DwNumberOfProcessors        uint32
	DwProcessorType             uint32
	DwAllocationGranularity     uint32
	WProcessorLevel             uint16
	WProcessorRevision          uint16
}

// OSVERSIONINFOEX 表示 Windows 版本信息。
type OSVERSIONINFOEX struct {
	DwOSVersionInfoSize uint32
	DwMajorVersion      uint32
	DwMinorVersion      uint32
	DwBuildNumber       uint32
	DwPlatformId        uint32
	SzCSDVersion        [128]uint16
	WServicePackMajor   uint16
	WServicePackMinor   uint16
	WSuiteMask          uint16
	WProductType        byte
	WReserved           byte
}

// SYSTEM_POWER_STATUS 表示系统电源和电池状态。
type SYSTEM_POWER_STATUS struct {
	ACLineStatus        byte
	BatteryFlag         byte
	BatteryLifePercent  byte
	SystemStatusFlag    byte
	BatteryLifeTime     uint32
	BatteryFullLifeTime uint32
}

// GetTickCount64 获取系统启动后的毫秒数。
// 返回值：返回系统启动后经过的毫秒数。
func GetTickCount64() uint64 {
	ret, _, _ := procGetTickCount64.Call()
	return uint64(ret)
}

// QueryPerformanceCounter 获取高精度性能计数器当前值。
// 返回值：返回计数器值和是否获取成功；true 表示 counter 有效。
func QueryPerformanceCounter() (counter int64, ok bool) {
	ret, _, _ := procQueryPerformanceCounter.Call(uintptr(unsafe.Pointer(&counter)))
	return counter, ret != 0
}

// QueryPerformanceFrequency 获取高精度性能计数器频率。
// 返回值：返回每秒计数次数和是否获取成功；true 表示 frequency 有效。
func QueryPerformanceFrequency() (frequency int64, ok bool) {
	ret, _, _ := procQueryPerformanceFrequency.Call(uintptr(unsafe.Pointer(&frequency)))
	return frequency, ret != 0
}

// GetSystemTime 获取 UTC 系统时间。
// 返回值：返回当前 UTC 系统时间。
func GetSystemTime() SYSTEMTIME {
	var st SYSTEMTIME
	procGetSystemTime.Call(uintptr(unsafe.Pointer(&st)))
	return st
}

// GetLocalTime 获取本地系统时间。
// 返回值：返回当前本地系统时间。
func GetLocalTime() SYSTEMTIME {
	var st SYSTEMTIME
	procGetLocalTime.Call(uintptr(unsafe.Pointer(&st)))
	return st
}

// SystemTimeToFileTime 将 SYSTEMTIME 转换为 FILETIME。
// 参数systemTime：系统时间结构指针。
// 返回值：返回文件时间和是否转换成功；true 表示 fileTime 有效。
func SystemTimeToFileTime(systemTime *SYSTEMTIME) (fileTime FILETIME, ok bool) {
	ret, _, _ := procSystemTimeToFileTime.Call(uintptr(unsafe.Pointer(systemTime)), uintptr(unsafe.Pointer(&fileTime)))
	return fileTime, ret != 0
}

// FileTimeToSystemTime 将 FILETIME 转换为 SYSTEMTIME。
// 参数fileTime：文件时间结构指针。
// 返回值：返回系统时间和是否转换成功；true 表示 systemTime 有效。
func FileTimeToSystemTime(fileTime *FILETIME) (systemTime SYSTEMTIME, ok bool) {
	ret, _, _ := procFileTimeToSystemTime.Call(uintptr(unsafe.Pointer(fileTime)), uintptr(unsafe.Pointer(&systemTime)))
	return systemTime, ret != 0
}

// GetSystemInfo 获取当前系统信息。
// 返回值：返回系统处理器、页大小、地址范围等信息。
func GetSystemInfo() SYSTEM_INFO {
	var info SYSTEM_INFO
	procGetSystemInfo.Call(uintptr(unsafe.Pointer(&info)))
	return info
}

// GetNativeSystemInfo 获取原生系统信息。
// 返回值：返回原生系统处理器、页大小、地址范围等信息；在 WOW64 下区别于 GetSystemInfo。
func GetNativeSystemInfo() SYSTEM_INFO {
	var info SYSTEM_INFO
	procGetNativeSystemInfo.Call(uintptr(unsafe.Pointer(&info)))
	return info
}

// GetVersion 获取旧式 Windows 版本值。
// 返回值：返回旧式版本编码值；新系统可能受兼容性清单影响。
func GetVersion() uint32 {
	ret, _, _ := procGetVersion.Call()
	return uint32(ret)
}

// GetVersionEx 获取 Windows 版本信息。
// 返回值：返回版本信息和是否获取成功；true 表示结构体内容有效。
func GetVersionEx() (OSVERSIONINFOEX, bool) {
	info := OSVERSIONINFOEX{DwOSVersionInfoSize: uint32(unsafe.Sizeof(OSVERSIONINFOEX{}))}
	ret, _, _ := procGetVersionExW.Call(uintptr(unsafe.Pointer(&info)))
	return info, ret != 0
}

// GetSystemPowerStatus 获取系统电源状态。
// 返回值：返回电源状态和是否获取成功；true 表示结构体内容有效。
func GetSystemPowerStatus() (SYSTEM_POWER_STATUS, bool) {
	var status SYSTEM_POWER_STATUS
	ret, _, _ := procGetSystemPowerStatus.Call(uintptr(unsafe.Pointer(&status)))
	return status, ret != 0
}
