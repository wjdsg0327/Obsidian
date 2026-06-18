//go:build windows

package user32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procFindWindowW                 = user32.NewProc("FindWindowW")
	procFindWindowExW               = user32.NewProc("FindWindowExW")
	procGetForegroundWindow         = user32.NewProc("GetForegroundWindow")
	procGetWindowTextW              = user32.NewProc("GetWindowTextW")
	procGetWindowTextLengthW        = user32.NewProc("GetWindowTextLengthW")
	procSetWindowTextW              = user32.NewProc("SetWindowTextW")
	procGetClassNameW               = user32.NewProc("GetClassNameW")
	procGetWindowThreadProcessId    = user32.NewProc("GetWindowThreadProcessId")
	procIsWindow                    = user32.NewProc("IsWindow")
	procIsWindowVisible             = user32.NewProc("IsWindowVisible")
	procIsWindowEnabled             = user32.NewProc("IsWindowEnabled")
	procIsIconic                    = user32.NewProc("IsIconic")
	procIsZoomed                    = user32.NewProc("IsZoomed")
	procShowWindow                  = user32.NewProc("ShowWindow")
	procShowWindowAsync             = user32.NewProc("ShowWindowAsync")
	procSetForegroundWindow         = user32.NewProc("SetForegroundWindow")
	procBringWindowToTop            = user32.NewProc("BringWindowToTop")
	procSetWindowPos                = user32.NewProc("SetWindowPos")
	procMoveWindow                  = user32.NewProc("MoveWindow")
	procGetWindowRect               = user32.NewProc("GetWindowRect")
	procGetClientRect               = user32.NewProc("GetClientRect")
	procClientToScreen              = user32.NewProc("ClientToScreen")
	procScreenToClient              = user32.NewProc("ScreenToClient")
	procWindowFromPoint             = user32.NewProc("WindowFromPoint")
	procChildWindowFromPoint        = user32.NewProc("ChildWindowFromPoint")
	procEnumWindows                 = user32.NewProc("EnumWindows")
	procEnumChildWindows            = user32.NewProc("EnumChildWindows")
	procGetParent                   = user32.NewProc("GetParent")
	procGetAncestor                 = user32.NewProc("GetAncestor")
	procGetWindow                   = user32.NewProc("GetWindow")
	procGetDesktopWindow            = user32.NewProc("GetDesktopWindow")
	procGetShellWindow              = user32.NewProc("GetShellWindow")
	procSetFocus                    = user32.NewProc("SetFocus")
	procGetFocus                    = user32.NewProc("GetFocus")
	procSetActiveWindow             = user32.NewProc("SetActiveWindow")
	procGetActiveWindow             = user32.NewProc("GetActiveWindow")
	procEnableWindow                = user32.NewProc("EnableWindow")
	procDestroyWindow               = user32.NewProc("DestroyWindow")
	procGetWindowLongPtrW           = user32.NewProc("GetWindowLongPtrW")
	procSetWindowLongPtrW           = user32.NewProc("SetWindowLongPtrW")
	procSetLayeredWindowAttributes  = user32.NewProc("SetLayeredWindowAttributes")
	procRegisterClassExW            = user32.NewProc("RegisterClassExW")
	procCreateWindowExW             = user32.NewProc("CreateWindowExW")
	procDefWindowProcW              = user32.NewProc("DefWindowProcW")
	procUpdateWindow                = user32.NewProc("UpdateWindow")
	procInvalidateRect              = user32.NewProc("InvalidateRect")
	procRedrawWindow                = user32.NewProc("RedrawWindow")
	procBeginPaint                  = user32.NewProc("BeginPaint")
	procEndPaint                    = user32.NewProc("EndPaint")
	procGetDC                       = user32.NewProc("GetDC")
	procReleaseDC                   = user32.NewProc("ReleaseDC")
	procCallWindowProcW             = user32.NewProc("CallWindowProcW")
	procChangeWindowMessageFilterEx = user32.NewProc("ChangeWindowMessageFilterEx")
	procAttachThreadInput           = user32.NewProc("AttachThreadInput")
	procAllowSetForegroundWindow    = user32.NewProc("AllowSetForegroundWindow")
)

const (
	SW_HIDE          = 0
	SW_SHOWNORMAL    = 1
	SW_SHOWMINIMIZED = 2
	SW_SHOWMAXIMIZED = 3
	SW_SHOW          = 5
	SW_MINIMIZE      = 6
	SW_RESTORE       = 9

	HWND_TOP       = 0
	HWND_TOPMOST   = ^uintptr(0)
	HWND_NOTOPMOST = ^uintptr(1)

	SWP_NOSIZE     = 0x0001
	SWP_NOMOVE     = 0x0002
	SWP_NOZORDER   = 0x0004
	SWP_SHOWWINDOW = 0x0040

	GA_ROOT       = 2
	GW_HWNDNEXT   = 2
	GW_HWNDPREV   = 3
	GWL_STYLE     = ^uintptr(15)
	GWL_EXSTYLE   = ^uintptr(19)
	WS_EX_LAYERED = 0x00080000
	LWA_ALPHA     = 0x00000002

	WM_DESTROY          = 0x0002
	WS_OVERLAPPEDWINDOW = 0x00CF0000
	CW_USEDEFAULT       = 0x80000000
	MSGFLT_ALLOW        = 1
)

// WNDCLASSEX 表示 RegisterClassExW 使用的窗口类结构。
type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     uintptr
	HIcon         uintptr
	HCursor       uintptr
	HbrBackground uintptr
	LpszMenuName  uintptr
	LpszClassName uintptr
	HIconSm       uintptr
}

// PAINTSTRUCT 表示 BeginPaint 和 EndPaint 使用的绘制状态。
type PAINTSTRUCT struct {
	Hdc         uintptr
	FErase      int32
	RcPaint     RECT
	FRestore    int32
	FIncUpdate  int32
	RgbReserved [32]byte
}

var wndProcPtr uintptr

// FindWindowByTitle 按窗口标题查找顶层窗口句柄。
// 参数title：窗口标题或提示框标题。
// 返回值：返回匹配标题的窗口句柄；返回 0 表示未找到。
func FindWindowByTitle(title string) uintptr {
	hwnd, _, _ := procFindWindowW.Call(0, utf16Ptr(title))
	return hwnd
}

// FindWindow 按类名和标题查找顶层窗口句柄，空字符串表示不限制。
// 参数className：窗口类名，空字符串表示不按类名过滤。
// 参数title：窗口标题或提示框标题。
// 返回值：返回匹配类名和标题的窗口句柄；返回 0 表示未找到。
func FindWindow(className, title string) uintptr {
	hwnd, _, _ := procFindWindowW.Call(utf16PtrOrNil(className), utf16PtrOrNil(title))
	return hwnd
}

// FindChildWindow 在父窗口下查找指定类名或标题的子窗口。
// 参数parent：父窗口句柄。
// 参数className：窗口类名，空字符串表示不按类名过滤。
// 参数title：窗口标题或提示框标题。
// 返回值：返回匹配条件的子窗口句柄；返回 0 表示未找到。
func FindChildWindow(parent uintptr, className, title string) uintptr {
	hwnd, _, _ := procFindWindowExW.Call(parent, 0, utf16PtrOrNil(className), utf16PtrOrNil(title))
	return hwnd
}

// GetForeground 获取当前前台窗口句柄。
// 返回值：返回当前前台窗口句柄；返回 0 表示获取失败或当前没有前台窗口。
func GetForeground() uintptr {
	hwnd, _, _ := procGetForegroundWindow.Call()
	return hwnd
}

// GetWindowTitle 读取窗口标题。
// 参数hwnd：窗口句柄。
// 返回值：返回窗口标题文本；窗口无标题或读取失败时返回空字符串。
func GetWindowTitle(hwnd uintptr) string {
	buf := make([]uint16, 512)
	procGetWindowTextW.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return syscall.UTF16ToString(buf)
}

// GetWindowTitleExact 按窗口标题实际长度读取标题。
// 参数hwnd：窗口句柄。
// 返回值：返回窗口标题文本；窗口无标题或读取失败时返回空字符串。
func GetWindowTitleExact(hwnd uintptr) string {
	n, _, _ := procGetWindowTextLengthW.Call(hwnd)
	if n == 0 {
		return ""
	}
	buf := make([]uint16, n+1)
	procGetWindowTextW.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return syscall.UTF16ToString(buf)
}

// SetWindowTitle 修改窗口标题。
// 参数hwnd：窗口句柄。
// 参数title：窗口标题或提示框标题。
// 返回值：true 表示窗口标题设置成功，false 表示设置失败。
func SetWindowTitle(hwnd uintptr, title string) bool {
	ret, _, _ := procSetWindowTextW.Call(hwnd, utf16Ptr(title))
	return ret != 0
}

// GetClassName 获取窗口类名。
// 参数hwnd：窗口句柄。
// 返回值：返回窗口类名；窗口无效或读取失败时返回空字符串。
func GetClassName(hwnd uintptr) string {
	buf := make([]uint16, 256)
	procGetClassNameW.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return syscall.UTF16ToString(buf)
}

// GetWindowPID 获取窗口所属线程 ID 和进程 ID。
// 参数hwnd：窗口句柄。
// 参数processID：类型为 uint32 的调用参数。
// 返回值：返回窗口所属线程 ID 和进程 ID；窗口无效时进程 ID 可能为 0。
func GetWindowPID(hwnd uintptr) (threadID, processID uint32) {
	var pid uint32
	tid, _, _ := procGetWindowThreadProcessId.Call(hwnd, uintptr(unsafe.Pointer(&pid)))
	return uint32(tid), pid
}

// IsValidWindow 判断窗口句柄是否有效。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口句柄当前有效，false 表示无效。
func IsValidWindow(hwnd uintptr) bool {
	ret, _, _ := procIsWindow.Call(hwnd)
	return ret != 0
}

// IsVisible 判断窗口是否可见。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口可见，false 表示不可见或句柄无效。
func IsVisible(hwnd uintptr) bool {
	ret, _, _ := procIsWindowVisible.Call(hwnd)
	return ret != 0
}

// IsEnabled 判断窗口是否可操作。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口可接收输入，false 表示被禁用或句柄无效。
func IsEnabled(hwnd uintptr) bool {
	ret, _, _ := procIsWindowEnabled.Call(hwnd)
	return ret != 0
}

// IsMinimized 判断窗口是否最小化。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口最小化，false 表示未最小化或句柄无效。
func IsMinimized(hwnd uintptr) bool {
	ret, _, _ := procIsIconic.Call(hwnd)
	return ret != 0
}

// IsMaximized 判断窗口是否最大化。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口最大化，false 表示未最大化或句柄无效。
func IsMaximized(hwnd uintptr) bool {
	ret, _, _ := procIsZoomed.Call(hwnd)
	return ret != 0
}

// RestoreWindow 还原窗口显示状态。
// 参数hwnd：窗口句柄。
// 返回值：true 表示调用前窗口是可见的，false 表示调用前窗口是隐藏的；这不是简单的成功失败标志。
func RestoreWindow(hwnd uintptr) bool {
	ret, _, _ := procShowWindow.Call(hwnd, SW_RESTORE)
	return ret != 0
}

// HideWindow 隐藏窗口。
// 参数hwnd：窗口句柄。
// 返回值：true 表示调用前窗口是可见的，false 表示调用前窗口是隐藏的；这不是简单的成功失败标志。
func HideWindow(hwnd uintptr) bool {
	ret, _, _ := procShowWindow.Call(hwnd, SW_HIDE)
	return ret != 0
}

// ShowWindowByCmd 按指定 ShowWindow 命令显示或隐藏窗口。
// 参数hwnd：窗口句柄。
// 参数cmd：ShowWindow 或 GetWindow 命令。
// 返回值：true 表示调用前窗口是可见的，false 表示调用前窗口是隐藏的；这不是简单的成功失败标志。
func ShowWindowByCmd(hwnd uintptr, cmd int) bool {
	ret, _, _ := procShowWindow.Call(hwnd, uintptr(cmd))
	return ret != 0
}

// RestoreWindowAsync 异步还原窗口显示状态。
// 参数hwnd：窗口句柄。
// 返回值：true 表示异步显示命令投递成功，false 表示投递失败。
func RestoreWindowAsync(hwnd uintptr) bool {
	ret, _, _ := procShowWindowAsync.Call(hwnd, SW_RESTORE)
	return ret != 0
}

// ActivateWindow 还原并尝试激活窗口到前台。
// 参数hwnd：窗口句柄。
// 返回值：true 表示系统接受前台激活请求，false 表示请求被限制或失败。
func ActivateWindow(hwnd uintptr) bool {
	RestoreWindow(hwnd)
	ret, _, _ := procSetForegroundWindow.Call(hwnd)
	return ret != 0
}

// BringTop 将窗口提升到 Z 顺序顶部。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口提升到 Z 顺序顶部成功，false 表示失败。
func BringTop(hwnd uintptr) bool {
	ret, _, _ := procBringWindowToTop.Call(hwnd)
	return ret != 0
}

// MoveResizeWindow 移动窗口并调整窗口大小。
// 参数hwnd：窗口句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 参数w：宽度。
// 参数h：高度。
// 返回值：true 表示窗口位置和大小设置成功，false 表示设置失败。
func MoveResizeWindow(hwnd uintptr, x, y, w, h int32) bool {
	ret, _, _ := procSetWindowPos.Call(hwnd, HWND_TOP, uintptr(x), uintptr(y), uintptr(w), uintptr(h), SWP_SHOWWINDOW)
	return ret != 0
}

// SetTopMost 设置或取消窗口置顶。
// 参数hwnd：窗口句柄。
// 参数yes：是否启用该状态。
// 返回值：true 表示置顶状态设置成功，false 表示设置失败。
func SetTopMost(hwnd uintptr, yes bool) bool {
	after := HWND_NOTOPMOST
	if yes {
		after = HWND_TOPMOST
	}
	ret, _, _ := procSetWindowPos.Call(hwnd, after, 0, 0, 0, 0, SWP_NOMOVE|SWP_NOSIZE)
	return ret != 0
}

// MoveWindowTo 移动窗口并调整尺寸，同时触发重绘。
// 参数hwnd：窗口句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 参数w：宽度。
// 参数h：高度。
// 返回值：true 表示窗口移动和缩放成功，false 表示操作失败。
func MoveWindowTo(hwnd uintptr, x, y, w, h int32) bool {
	ret, _, _ := procMoveWindow.Call(hwnd, uintptr(x), uintptr(y), uintptr(w), uintptr(h), 1)
	return ret != 0
}

// GetWindowRect 获取窗口外框矩形。
// 参数hwnd：窗口句柄。
// 返回值：返回窗口外框矩形和是否获取成功；true 表示矩形有效，false 表示获取失败。
func GetWindowRect(hwnd uintptr) (RECT, bool) {
	var r RECT
	ret, _, _ := procGetWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&r)))
	return r, ret != 0
}

// GetClientRect 获取窗口客户区矩形。
// 参数hwnd：窗口句柄。
// 返回值：返回客户区矩形和是否获取成功；true 表示矩形有效，false 表示获取失败。
func GetClientRect(hwnd uintptr) (RECT, bool) {
	var r RECT
	ret, _, _ := procGetClientRect.Call(hwnd, uintptr(unsafe.Pointer(&r)))
	return r, ret != 0
}

// ClientToScreen 将客户区坐标转换为屏幕坐标。
// 参数hwnd：窗口句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 返回值：返回转换后的屏幕坐标和是否转换成功；true 表示坐标有效，false 表示转换失败。
func ClientToScreen(hwnd uintptr, x, y int32) (POINT, bool) {
	pt := POINT{x, y}
	ret, _, _ := procClientToScreen.Call(hwnd, uintptr(unsafe.Pointer(&pt)))
	return pt, ret != 0
}

// ScreenToClient 将屏幕坐标转换为客户区坐标。
// 参数hwnd：窗口句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 返回值：返回转换后的客户区坐标和是否转换成功；true 表示坐标有效，false 表示转换失败。
func ScreenToClient(hwnd uintptr, x, y int32) (POINT, bool) {
	pt := POINT{x, y}
	ret, _, _ := procScreenToClient.Call(hwnd, uintptr(unsafe.Pointer(&pt)))
	return pt, ret != 0
}

// WindowAt 根据屏幕坐标获取窗口句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 返回值：返回指定屏幕坐标处的窗口句柄；返回 0 表示未找到。
func WindowAt(x, y int32) uintptr {
	hwnd, _, _ := procWindowFromPoint.Call(packPoint(x, y))
	return hwnd
}

// ChildAt 根据父窗口客户区坐标获取子窗口句柄。
// 参数parent：父窗口句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 返回值：返回父窗口客户区坐标处的子窗口句柄；返回 0 表示未找到。
func ChildAt(parent uintptr, x, y int32) uintptr {
	hwnd, _, _ := procChildWindowFromPoint.Call(parent, packPoint(x, y))
	return hwnd
}

// EnumTopWindows 枚举所有顶层窗口句柄。
// 返回值：返回枚举到的顶层窗口句柄列表；没有结果时返回空切片。
func EnumTopWindows() []uintptr {
	var windows []uintptr
	cb := syscall.NewCallback(func(hwnd uintptr, lparam uintptr) uintptr {
		windows = append(windows, hwnd)
		return 1
	})
	procEnumWindows.Call(cb, 0)
	return windows
}

// EnumChildren 枚举指定父窗口下的所有子窗口句柄。
// 参数parent：父窗口句柄。
// 返回值：返回枚举到的子窗口句柄列表；没有结果时返回空切片。
func EnumChildren(parent uintptr) []uintptr {
	var children []uintptr
	cb := syscall.NewCallback(func(hwnd uintptr, lparam uintptr) uintptr {
		children = append(children, hwnd)
		return 1
	})
	procEnumChildWindows.Call(parent, cb, 0)
	return children
}

// GetParentWindow 获取窗口的父窗口句柄。
// 参数hwnd：窗口句柄。
// 返回值：返回父窗口句柄；返回 0 表示没有父窗口或获取失败。
func GetParentWindow(hwnd uintptr) uintptr {
	parent, _, _ := procGetParent.Call(hwnd)
	return parent
}

// GetRootWindow 获取窗口的根祖先窗口句柄。
// 参数hwnd：窗口句柄。
// 返回值：返回根祖先窗口句柄；返回 0 表示获取失败。
func GetRootWindow(hwnd uintptr) uintptr {
	root, _, _ := procGetAncestor.Call(hwnd, GA_ROOT)
	return root
}

// GetRelatedWindow 按 GetWindow 命令获取相关窗口。
// 参数hwnd：窗口句柄。
// 参数cmd：ShowWindow 或 GetWindow 命令。
// 返回值：返回与指定窗口相关的窗口句柄；返回 0 表示没有对应窗口或获取失败。
func GetRelatedWindow(hwnd uintptr, cmd uint32) uintptr {
	w, _, _ := procGetWindow.Call(hwnd, uintptr(cmd))
	return w
}

// GetNextWindow 获取 Z 顺序中的下一个窗口。
// 参数hwnd：窗口句柄。
// 返回值：返回 Z 顺序中的下一个窗口句柄；返回 0 表示没有下一个窗口。
func GetNextWindow(hwnd uintptr) uintptr {
	return GetRelatedWindow(hwnd, GW_HWNDNEXT)
}

// DesktopWindow 获取桌面窗口句柄。
// 返回值：返回桌面窗口句柄。
func DesktopWindow() uintptr {
	hwnd, _, _ := procGetDesktopWindow.Call()
	return hwnd
}

// ShellWindow 获取 Shell 桌面窗口句柄。
// 返回值：返回 Shell 桌面窗口句柄；返回 0 表示当前没有 Shell 窗口。
func ShellWindow() uintptr {
	hwnd, _, _ := procGetShellWindow.Call()
	return hwnd
}

// FocusWindow 设置当前线程键盘焦点窗口。
// 参数hwnd：窗口句柄。
// 返回值：返回之前拥有键盘焦点的窗口句柄；返回 0 表示之前没有焦点窗口或调用失败。
func FocusWindow(hwnd uintptr) uintptr {
	old, _, _ := procSetFocus.Call(hwnd)
	return old
}

// CurrentThreadFocus 获取当前线程焦点窗口。
// 返回值：返回当前线程焦点窗口句柄；返回 0 表示当前线程没有焦点窗口。
func CurrentThreadFocus() uintptr {
	hwnd, _, _ := procGetFocus.Call()
	return hwnd
}

// SetActive 设置当前线程活动窗口。
// 参数hwnd：窗口句柄。
// 返回值：返回之前的活动窗口句柄；返回 0 表示之前没有活动窗口或调用失败。
func SetActive(hwnd uintptr) uintptr {
	old, _, _ := procSetActiveWindow.Call(hwnd)
	return old
}

// GetActive 获取当前线程活动窗口。
// 返回值：返回当前线程活动窗口句柄；返回 0 表示当前线程没有活动窗口。
func GetActive() uintptr {
	hwnd, _, _ := procGetActiveWindow.Call()
	return hwnd
}

// Enable 启用或禁用窗口输入。
// 参数hwnd：窗口句柄。
// 参数enabled：是否启用。
// 返回值：true 表示调用前窗口处于启用状态，false 表示调用前窗口处于禁用状态；Win32 返回的是旧状态。
func Enable(hwnd uintptr, enabled bool) bool {
	ret, _, _ := procEnableWindow.Call(hwnd, boolArg(enabled))
	return ret != 0
}

// Destroy 销毁窗口。
// 参数hwnd：窗口句柄。
// 返回值：true 表示窗口销毁成功，false 表示销毁失败。
func Destroy(hwnd uintptr) bool {
	ret, _, _ := procDestroyWindow.Call(hwnd)
	return ret != 0
}

// GetWindowStyle 获取窗口样式。
// 参数hwnd：窗口句柄。
// 返回值：返回窗口样式位掩码；返回 0 可能表示无样式或获取失败。
func GetWindowStyle(hwnd uintptr) uintptr {
	style, _, _ := procGetWindowLongPtrW.Call(hwnd, GWL_STYLE)
	return style
}

// SetWindowStyle 设置窗口样式并返回旧样式。
// 参数hwnd：窗口句柄。
// 参数style：窗口样式。
// 返回值：返回修改前的窗口样式位掩码；返回 0 可能表示旧样式为 0 或调用失败。
func SetWindowStyle(hwnd uintptr, style uintptr) uintptr {
	old, _, _ := procSetWindowLongPtrW.Call(hwnd, GWL_STYLE, style)
	return old
}

// SetWindowAlpha 设置分层窗口透明度。
// 参数hwnd：窗口句柄。
// 参数alpha：透明度，0 表示全透明，255 表示不透明。
// 返回值：true 表示窗口透明度设置成功，false 表示设置失败。
func SetWindowAlpha(hwnd uintptr, alpha byte) bool {
	exStyle, _, _ := procGetWindowLongPtrW.Call(hwnd, GWL_EXSTYLE)
	procSetWindowLongPtrW.Call(hwnd, GWL_EXSTYLE, exStyle|WS_EX_LAYERED)
	ret, _, _ := procSetLayeredWindowAttributes.Call(hwnd, 0, uintptr(alpha), LWA_ALPHA)
	return ret != 0
}

// wndProc 处理 CreateSimpleWindow 创建窗口的默认窗口过程。
func wndProc(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case WM_DESTROY:
		QuitApp(0)
		return 0
	}
	ret, _, _ := procDefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
	return ret
}

// CreateSimpleWindow 注册窗口类并创建一个简单 Win32 原生窗口。
// 返回值：返回新创建的窗口句柄；返回 0 表示创建失败。
func CreateSimpleWindow() uintptr {
	wndProcPtr = syscall.NewCallback(wndProc)
	className := "GoUser32Window"
	wc := WNDCLASSEX{
		CbSize:        uint32(unsafe.Sizeof(WNDCLASSEX{})),
		LpfnWndProc:   wndProcPtr,
		LpszClassName: utf16Ptr(className),
	}
	procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))
	hwnd, _, _ := procCreateWindowExW.Call(0, utf16Ptr(className), utf16Ptr("Go 创建的 user32 窗口"), WS_OVERLAPPEDWINDOW, CW_USEDEFAULT, CW_USEDEFAULT, 800, 600, 0, 0, 0, 0)
	return hwnd
}

// ForceRepaint 标记窗口无效并立即请求重绘。
// 参数hwnd：窗口句柄。
// 返回值：无返回值；函数会请求窗口重绘。
func ForceRepaint(hwnd uintptr) {
	procInvalidateRect.Call(hwnd, 0, 1)
	procUpdateWindow.Call(hwnd)
}

// RedrawWindowNow 立即按指定标志重绘窗口区域。
// 参数hwnd：窗口句柄。
// 参数rect：矩形区域指针，nil 表示不指定。
// 参数flags：调用标志位。
// 返回值：true 表示重绘请求成功，false 表示请求失败。
func RedrawWindowNow(hwnd uintptr, rect *RECT, flags uint32) bool {
	var rectPtr uintptr
	if rect != nil {
		rectPtr = uintptr(unsafe.Pointer(rect))
	}
	ret, _, _ := procRedrawWindow.Call(hwnd, rectPtr, 0, uintptr(flags))
	return ret != 0
}

// OnPaint 执行一次 BeginPaint 和 EndPaint 绘制流程并返回设备上下文。
// 参数hwnd：窗口句柄。
// 返回值：返回 BeginPaint 得到的设备上下文句柄；返回 0 表示开始绘制失败。
func OnPaint(hwnd uintptr) uintptr {
	var ps PAINTSTRUCT
	hdc, _, _ := procBeginPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))
	procEndPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))
	return hdc
}

// WithWindowDC 获取窗口设备上下文并在回调后释放。
// 参数hwnd：窗口句柄。
// 参数fn：持有资源期间执行的回调函数。
// 返回值：true 表示成功获取设备上下文并执行回调，false 表示获取设备上下文失败。
func WithWindowDC(hwnd uintptr, fn func(hdc uintptr)) bool {
	hdc, _, _ := procGetDC.Call(hwnd)
	if hdc == 0 {
		return false
	}
	defer procReleaseDC.Call(hwnd, hdc)
	fn(hdc)
	return true
}

// CallOldWndProc 调用子类化前保存的原窗口过程。
// 参数oldProc：原窗口过程指针。
// 参数hwnd：窗口句柄。
// 参数msg：窗口消息编号。
// 参数wParam：消息的 wParam 参数。
// 参数lParam：消息的 lParam 参数。
// 返回值：返回原窗口过程处理消息后的结果值，具体含义由消息类型决定。
func CallOldWndProc(oldProc uintptr, hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procCallWindowProcW.Call(oldProc, hwnd, uintptr(msg), wParam, lParam)
	return ret
}

// AllowMessage 允许指定消息通过窗口消息过滤。
// 参数hwnd：窗口句柄。
// 参数msg：窗口消息编号。
// 返回值：true 表示指定消息已允许通过过滤，false 表示设置失败。
func AllowMessage(hwnd uintptr, msg uint32) bool {
	ret, _, _ := procChangeWindowMessageFilterEx.Call(hwnd, uintptr(msg), MSGFLT_ALLOW, 0)
	return ret != 0
}

// AttachInput 连接或断开两个线程的输入队列。
// 参数srcThread：源线程 ID。
// 参数dstThread：目标线程 ID。
// 参数attach：true 表示连接输入队列，false 表示断开。
// 返回值：true 表示线程输入队列连接或断开成功，false 表示操作失败。
func AttachInput(srcThread, dstThread uint32, attach bool) bool {
	ret, _, _ := procAttachThreadInput.Call(uintptr(srcThread), uintptr(dstThread), boolArg(attach))
	return ret != 0
}

// AllowProcessSetForeground 允许指定进程设置前台窗口。
// 参数pid：进程 ID。
// 返回值：true 表示已允许指定进程设置前台窗口，false 表示设置失败。
func AllowProcessSetForeground(pid uint32) bool {
	ret, _, _ := procAllowSetForegroundWindow.Call(uintptr(pid))
	return ret != 0
}
