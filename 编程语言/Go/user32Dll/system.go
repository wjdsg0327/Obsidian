//go:build windows

package user32Dll

import "unsafe"

var (
	procSystemParametersInfoW = user32.NewProc("SystemParametersInfoW")
	procLockWorkStation       = user32.NewProc("LockWorkStation")
	procExitWindowsEx         = user32.NewProc("ExitWindowsEx")
)

const (
	SPI_GETMOUSESPEED = 0x0070
	EWX_LOGOFF        = 0x00000000
)

// GetMouseSpeed 读取系统鼠标速度。
// 返回值：返回系统鼠标速度，范围通常为 1 到 20；读取失败时可能返回 0。
func GetMouseSpeed() int {
	var speed uint32
	procSystemParametersInfoW.Call(SPI_GETMOUSESPEED, 0, uintptr(unsafe.Pointer(&speed)), 0)
	return int(speed)
}

// SystemParametersInfo 调用 SystemParametersInfoW 读取或修改系统参数。
// 参数action：SystemParametersInfoW 的动作编号。
// 参数uiParam：动作相关的整数参数。
// 参数pvParam：动作相关的数据指针。
// 参数winIni：系统参数更新标志。
// 返回值：true 表示系统参数读取或修改成功，false 表示调用失败。
func SystemParametersInfo(action, uiParam uint32, pvParam uintptr, winIni uint32) bool {
	ret, _, _ := procSystemParametersInfoW.Call(uintptr(action), uintptr(uiParam), pvParam, uintptr(winIni))
	return ret != 0
}

// LockWindows 锁定 Windows 工作站。
// 返回值：true 表示锁定工作站请求成功，false 表示请求失败。
func LockWindows() bool {
	ret, _, _ := procLockWorkStation.Call()
	return ret != 0
}

// LogoffWindows 注销当前 Windows 用户。
// 返回值：true 表示注销请求已提交成功，false 表示提交失败。
func LogoffWindows() bool {
	ret, _, _ := procExitWindowsEx.Call(EWX_LOGOFF, 0)
	return ret != 0
}

// ExitWindows 调用 ExitWindowsEx 执行注销、关机或重启等操作。
// 参数flags：调用标志位。
// 参数reason：关机或注销原因代码。
// 返回值：true 表示注销、关机或重启请求已提交成功，false 表示提交失败。
func ExitWindows(flags, reason uint32) bool {
	ret, _, _ := procExitWindowsEx.Call(uintptr(flags), uintptr(reason))
	return ret != 0
}
