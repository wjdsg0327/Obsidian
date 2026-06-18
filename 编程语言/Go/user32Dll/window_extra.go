//go:build windows

package user32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procCreateWindowW       = user32.NewProc("CreateWindowW")
	procUnregisterClassW    = user32.NewProc("UnregisterClassW")
	procGetClassInfoExW     = user32.NewProc("GetClassInfoExW")
	procGetClassLongPtrW    = user32.NewProc("GetClassLongPtrW")
	procSetClassLongPtrW    = user32.NewProc("SetClassLongPtrW")
	procGetWindowPlacement  = user32.NewProc("GetWindowPlacement")
	procSetWindowPlacement  = user32.NewProc("SetWindowPlacement")
	procGetWindowInfo       = user32.NewProc("GetWindowInfo")
	procGetTitleBarInfo     = user32.NewProc("GetTitleBarInfo")
	procFlashWindowEx       = user32.NewProc("FlashWindowEx")
	procIsWindowArranged    = user32.NewProc("IsWindowArranged")
	procGetLastActivePopup  = user32.NewProc("GetLastActivePopup")
	procGetWindowModuleFile = user32.NewProc("GetWindowModuleFileNameW")
)

const (
	FLASHW_STOP      = 0x00000000
	FLASHW_CAPTION   = 0x00000001
	FLASHW_TRAY      = 0x00000002
	FLASHW_ALL       = FLASHW_CAPTION | FLASHW_TRAY
	FLASHW_TIMER     = 0x00000004
	FLASHW_TIMERNOFG = 0x0000000C
)

// WINDOWPLACEMENT 表示窗口显示状态、最小化位置、最大化位置和还原矩形。
type WINDOWPLACEMENT struct {
	Length           uint32
	Flags            uint32
	ShowCmd          uint32
	PtMinPosition    POINT
	PtMaxPosition    POINT
	RcNormalPosition RECT
}

// WINDOWINFO 表示窗口边框、客户区、样式和状态信息。
type WINDOWINFO struct {
	CbSize          uint32
	RcWindow        RECT
	RcClient        RECT
	DwStyle         uint32
	DwExStyle       uint32
	DwWindowStatus  uint32
	CxWindowBorders uint32
	CyWindowBorders uint32
	AtomWindowType  uint16
	WCreatorVersion uint16
}

// TITLEBARINFO 表示标题栏矩形和标题栏元素状态。
type TITLEBARINFO struct {
	CbSize     uint32
	RcTitleBar RECT
	Rgstate    [6]uint32
}

// FLASHWINFO 表示 FlashWindowEx 的闪烁参数。
type FLASHWINFO struct {
	CbSize    uint32
	Hwnd      uintptr
	DwFlags   uint32
	UCount    uint32
	DwTimeout uint32
}

// CreateWindow 创建一个标准 Win32 窗口。
// 参数className：窗口类名，空字符串表示不按类名过滤。
// 参数title：窗口标题或提示框标题。
// 参数style：窗口样式。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 参数w：宽度。
// 参数h：高度。
// 参数parent：父窗口句柄。
// 参数menu：菜单句柄。
// 参数instance：模块实例句柄。
// 参数param：传给窗口创建流程的自定义参数指针。
// 返回值：返回新创建的窗口句柄；返回 0 表示创建失败。
func CreateWindow(className, title string, style uint32, x, y, w, h int32, parent, menu, instance, param uintptr) uintptr {
	hwnd, _, _ := procCreateWindowW.Call(utf16Ptr(className), utf16Ptr(title), uintptr(style), uintptr(x), uintptr(y), uintptr(w), uintptr(h), parent, menu, instance, param)
	return hwnd
}

// UnregisterWindowClass 注销已注册的窗口类。
// 参数className：窗口类名，空字符串表示不按类名过滤。
// 参数hInstance：注册窗口类时使用的模块实例句柄。
// 返回值：true 表示窗口类注销成功，false 表示注销失败。
func UnregisterWindowClass(className string, hInstance uintptr) bool {
	ret, _, _ := procUnregisterClassW.Call(utf16Ptr(className), hInstance)
	return ret != 0
}

// GetWindowClassInfo 获取窗口类注册信息。
// 参数className：窗口类名，空字符串表示不按类名过滤。
// 参数hInstance：窗口类所属模块实例句柄。
// 返回值：返回窗口类信息和是否获取成功；true 表示结构体内容有效，false 表示获取失败。
func GetWindowClassInfo(className string, hInstance uintptr) (WNDCLASSEX, bool) {
	wc := WNDCLASSEX{CbSize: uint32(unsafe.Sizeof(WNDCLASSEX{}))}
	ret, _, _ := procGetClassInfoExW.Call(hInstance, utf16Ptr(className), uintptr(unsafe.Pointer(&wc)))
	return wc, ret != 0
}

// GetClassLongPtr 读取窗口类长指针属性。
// 参数hwnd：窗口句柄。
// 参数index：要读取的类属性索引。
// 返回值：返回指定类属性值；返回 0 可能表示值为 0 或读取失败。
func GetClassLongPtr(hwnd uintptr, index int32) uintptr {
	ret, _, _ := procGetClassLongPtrW.Call(hwnd, uintptr(index))
	return ret
}

// SetClassLongPtr 设置窗口类长指针属性。
// 参数hwnd：窗口句柄。
// 参数index：要设置的类属性索引。
// 参数value：新的类属性值。
// 返回值：返回修改前的类属性值；返回 0 可能表示旧值为 0 或设置失败。
func SetClassLongPtr(hwnd uintptr, index int32, value uintptr) uintptr {
	ret, _, _ := procSetClassLongPtrW.Call(hwnd, uintptr(index), value)
	return ret
}

// GetWindowPlacementInfo 获取窗口显示位置和状态。
// 参数hwnd：窗口句柄。
// 返回值：返回窗口位置状态和是否获取成功；true 表示结构体内容有效，false 表示获取失败。
func GetWindowPlacementInfo(hwnd uintptr) (WINDOWPLACEMENT, bool) {
	wp := WINDOWPLACEMENT{Length: uint32(unsafe.Sizeof(WINDOWPLACEMENT{}))}
	ret, _, _ := procGetWindowPlacement.Call(hwnd, uintptr(unsafe.Pointer(&wp)))
	return wp, ret != 0
}

// SetWindowPlacementInfo 设置窗口显示位置和状态。
// 参数hwnd：窗口句柄。
// 参数placement：窗口位置状态结构指针。
// 返回值：true 表示窗口位置状态设置成功，false 表示设置失败。
func SetWindowPlacementInfo(hwnd uintptr, placement *WINDOWPLACEMENT) bool {
	if placement != nil && placement.Length == 0 {
		placement.Length = uint32(unsafe.Sizeof(WINDOWPLACEMENT{}))
	}
	ret, _, _ := procSetWindowPlacement.Call(hwnd, uintptr(unsafe.Pointer(placement)))
	return ret != 0
}

// GetWindowInfo 获取窗口边框、客户区和样式信息。
// 参数hwnd：窗口句柄。
// 返回值：返回窗口信息和是否获取成功；true 表示结构体内容有效，false 表示获取失败。
func GetWindowInfo(hwnd uintptr) (WINDOWINFO, bool) {
	info := WINDOWINFO{CbSize: uint32(unsafe.Sizeof(WINDOWINFO{}))}
	ret, _, _ := procGetWindowInfo.Call(hwnd, uintptr(unsafe.Pointer(&info)))
	return info, ret != 0
}

// GetTitleBarInfo 获取窗口标题栏信息。
// 参数hwnd：窗口句柄。
// 返回值：返回标题栏信息和是否获取成功；true 表示结构体内容有效，false 表示获取失败。
func GetTitleBarInfo(hwnd uintptr) (TITLEBARINFO, bool) {
	info := TITLEBARINFO{CbSize: uint32(unsafe.Sizeof(TITLEBARINFO{}))}
	ret, _, _ := procGetTitleBarInfo.Call(hwnd, uintptr(unsafe.Pointer(&info)))
	return info, ret != 0
}

// FlashWindow 闪烁窗口标题栏或任务栏按钮。
// 参数hwnd：窗口句柄。
// 参数flags：闪烁方式标志，例如 FLASHW_ALL。
// 参数count：闪烁次数。
// 参数timeout：每次闪烁间隔，0 表示使用系统默认值。
// 返回值：true 表示调用前窗口处于活动状态，false 表示调用前窗口不处于活动状态；这不是简单的成功失败标志。
func FlashWindow(hwnd uintptr, flags, count, timeout uint32) bool {
	info := FLASHWINFO{CbSize: uint32(unsafe.Sizeof(FLASHWINFO{})), Hwnd: hwnd, DwFlags: flags, UCount: count, DwTimeout: timeout}
	ret, _, _ := procFlashWindowEx.Call(uintptr(unsafe.Pointer(&info)))
	return ret != 0
}

// IsWindowArranged 判断窗口是否由系统排列布局管理。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口处于系统排列状态，false 表示不处于排列状态或调用失败。
func IsWindowArranged(hwnd uintptr) bool {
	ret, _, _ := procIsWindowArranged.Call(hwnd)
	return ret != 0
}

// GetLastActivePopup 获取指定窗口拥有的最后一个活动弹出窗口。
// 参数hwnd：窗口句柄。
// 返回值：返回最后一个活动弹出窗口句柄；没有弹出窗口时通常返回原窗口句柄。
func GetLastActivePopup(hwnd uintptr) uintptr {
	ret, _, _ := procGetLastActivePopup.Call(hwnd)
	return ret
}

// GetWindowModuleFileName 获取创建窗口的模块文件名。
// 参数hwnd：窗口句柄。
// 返回值：返回模块文件名；读取失败时返回空字符串。
func GetWindowModuleFileName(hwnd uintptr) string {
	buf := make([]uint16, 1024)
	ret, _, _ := procGetWindowModuleFile.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}
