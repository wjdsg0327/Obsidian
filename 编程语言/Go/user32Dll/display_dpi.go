//go:build windows

package user32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procGetSystemMetrics              = user32.NewProc("GetSystemMetrics")
	procEnumDisplayMonitors           = user32.NewProc("EnumDisplayMonitors")
	procGetMonitorInfoW               = user32.NewProc("GetMonitorInfoW")
	procMonitorFromWindow             = user32.NewProc("MonitorFromWindow")
	procMonitorFromPoint              = user32.NewProc("MonitorFromPoint")
	procMonitorFromRect               = user32.NewProc("MonitorFromRect")
	procGetDpiForWindow               = user32.NewProc("GetDpiForWindow")
	procGetDpiForSystem               = user32.NewProc("GetDpiForSystem")
	procSetProcessDPIAware            = user32.NewProc("SetProcessDPIAware")
	procSetProcessDpiAwarenessContext = user32.NewProc("SetProcessDpiAwarenessContext")
	procAdjustWindowRectEx            = user32.NewProc("AdjustWindowRectEx")
	procAdjustWindowRectExForDpi      = user32.NewProc("AdjustWindowRectExForDpi")
)

const (
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1

	MONITOR_DEFAULTTONEAREST = 2
)

var DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2 = ^uintptr(3)

// MONITORINFO 表示显示器矩形和工作区信息。
type MONITORINFO struct {
	CbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	DwFlags   uint32
}

// ScreenSize 获取主屏幕宽度和高度。
// 参数) (w：类型为 int 的调用参数。
// 参数h：高度。
// 返回值：返回主屏幕宽度 w 和高度 h，单位为像素。
func ScreenSize() (w, h int) {
	rw, _, _ := procGetSystemMetrics.Call(SM_CXSCREEN)
	rh, _, _ := procGetSystemMetrics.Call(SM_CYSCREEN)
	return int(rw), int(rh)
}

// GetSystemMetric 获取指定系统尺寸或状态指标。
// 参数index：类型为 int 的调用参数。
// 返回值：返回指定系统指标的整数值；具体含义由 index 决定。
func GetSystemMetric(index int) int {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(index))
	return int(ret)
}

// EnumMonitors 枚举所有显示器矩形。
// 返回值：返回所有显示器的屏幕矩形列表；枚举失败或没有结果时返回空切片。
func EnumMonitors() []RECT {
	var rects []RECT
	cb := syscall.NewCallback(func(hMonitor, hdc uintptr, lprcMonitor uintptr, dwData uintptr) uintptr {
		var mi MONITORINFO
		mi.CbSize = uint32(unsafe.Sizeof(mi))
		procGetMonitorInfoW.Call(hMonitor, uintptr(unsafe.Pointer(&mi)))
		rects = append(rects, mi.RcMonitor)
		return 1
	})
	procEnumDisplayMonitors.Call(0, 0, cb, 0)
	return rects
}

// GetMonitorInfo 获取指定显示器信息。
// 参数hMonitor：显示器句柄。
// 返回值：返回显示器信息和是否获取成功；true 表示信息有效，false 表示获取失败。
func GetMonitorInfo(hMonitor uintptr) (MONITORINFO, bool) {
	mi := MONITORINFO{CbSize: uint32(unsafe.Sizeof(MONITORINFO{}))}
	ret, _, _ := procGetMonitorInfoW.Call(hMonitor, uintptr(unsafe.Pointer(&mi)))
	return mi, ret != 0
}

// MonitorOfWindow 获取窗口所在显示器句柄。
// 参数hwnd：窗口句柄。
// 返回值：返回窗口所在显示器句柄；按 MONITOR_DEFAULTTONEAREST 规则通常会返回最近的显示器。
func MonitorOfWindow(hwnd uintptr) uintptr {
	h, _, _ := procMonitorFromWindow.Call(hwnd, MONITOR_DEFAULTTONEAREST)
	return h
}

// MonitorOfPoint 获取指定屏幕坐标所在显示器句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 返回值：返回坐标所在显示器句柄；坐标不在任何显示器内时返回最近的显示器。
func MonitorOfPoint(x, y int32) uintptr {
	h, _, _ := procMonitorFromPoint.Call(packPoint(x, y), MONITOR_DEFAULTTONEAREST)
	return h
}

// MonitorOfRect 获取指定矩形所在显示器句柄。
// 参数r：矩形区域。
// 返回值：返回矩形所在显示器句柄；矩形不在任何显示器内时返回最近的显示器。
func MonitorOfRect(r RECT) uintptr {
	h, _, _ := procMonitorFromRect.Call(uintptr(unsafe.Pointer(&r)), MONITOR_DEFAULTTONEAREST)
	return h
}

// DpiForWindow 获取指定窗口 DPI。
// 参数hwnd：窗口句柄。
// 返回值：返回窗口当前 DPI 值；窗口无效或系统不支持时可能返回 0。
func DpiForWindow(hwnd uintptr) uint32 {
	dpi, _, _ := procGetDpiForWindow.Call(hwnd)
	return uint32(dpi)
}

// SystemDPI 获取系统 DPI。
// 返回值：返回系统 DPI 值。
func SystemDPI() uint32 {
	dpi, _, _ := procGetDpiForSystem.Call()
	return uint32(dpi)
}

// EnableDPIAwareOld 启用旧版进程 DPI 感知。
// 返回值：true 表示已成功启用旧版进程 DPI 感知，false 表示设置失败。
func EnableDPIAwareOld() bool {
	ret, _, _ := procSetProcessDPIAware.Call()
	return ret != 0
}

// EnablePerMonitorDPIAwareV2 启用 Per-Monitor DPI Aware V2。
// 返回值：true 表示已成功启用 Per-Monitor DPI Aware V2，false 表示设置失败。
func EnablePerMonitorDPIAwareV2() bool {
	ret, _, _ := procSetProcessDpiAwarenessContext.Call(DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2)
	return ret != 0
}

// CalcWindowRectForClient 根据客户区尺寸计算窗口外框尺寸。
// 参数w：宽度。
// 参数h：高度。
// 参数style：窗口样式。
// 参数exStyle：窗口扩展样式。
// 返回值：返回为了容纳指定客户区所需的窗口外框矩形。
func CalcWindowRectForClient(w, h int32, style, exStyle uint32) RECT {
	r := RECT{0, 0, w, h}
	procAdjustWindowRectEx.Call(uintptr(unsafe.Pointer(&r)), uintptr(style), 0, uintptr(exStyle))
	return r
}

// CalcWindowRectForClientDPI 根据客户区尺寸和 DPI 计算窗口外框尺寸。
// 参数w：宽度。
// 参数h：高度。
// 参数style：窗口样式。
// 参数exStyle：窗口扩展样式。
// 参数dpi：DPI 数值。
// 返回值：返回在指定 DPI 下为了容纳客户区所需的窗口外框矩形。
func CalcWindowRectForClientDPI(w, h int32, style, exStyle, dpi uint32) RECT {
	r := RECT{0, 0, w, h}
	procAdjustWindowRectExForDpi.Call(uintptr(unsafe.Pointer(&r)), uintptr(style), 0, uintptr(exStyle), uintptr(dpi))
	return r
}
