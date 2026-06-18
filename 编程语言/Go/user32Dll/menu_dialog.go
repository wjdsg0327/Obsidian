//go:build windows

package user32Dll

import (
	"syscall"
	"unsafe"
)

var (
	procCreateMenu         = user32.NewProc("CreateMenu")
	procAppendMenuW        = user32.NewProc("AppendMenuW")
	procSetMenu            = user32.NewProc("SetMenu")
	procDrawMenuBar        = user32.NewProc("DrawMenuBar")
	procCreatePopupMenu    = user32.NewProc("CreatePopupMenu")
	procTrackPopupMenu     = user32.NewProc("TrackPopupMenu")
	procDestroyMenu        = user32.NewProc("DestroyMenu")
	procEnableMenuItem     = user32.NewProc("EnableMenuItem")
	procCheckMenuItem      = user32.NewProc("CheckMenuItem")
	procGetScrollInfo      = user32.NewProc("GetScrollInfo")
	procSetScrollInfo      = user32.NewProc("SetScrollInfo")
	procShowScrollBar      = user32.NewProc("ShowScrollBar")
	procGetDlgItem         = user32.NewProc("GetDlgItem")
	procSetDlgItemTextW    = user32.NewProc("SetDlgItemTextW")
	procGetDlgItemTextW    = user32.NewProc("GetDlgItemTextW")
	procCheckDlgButton     = user32.NewProc("CheckDlgButton")
	procIsDlgButtonChecked = user32.NewProc("IsDlgButtonChecked")
)

const (
	MF_STRING    = 0x00000000
	MF_BYCOMMAND = 0x00000000
	MF_GRAYED    = 0x00000001
	MF_ENABLED   = 0x00000000
	MF_UNCHECKED = 0x00000000
	MF_CHECKED   = 0x00000008

	TPM_RIGHTBUTTON = 0x0002

	SB_VERT = 1
	SIF_ALL = 0x17

	BST_CHECKED   = 1
	BST_UNCHECKED = 0
)

// SCROLLINFO 表示滚动条信息结构。
type SCROLLINFO struct {
	CbSize    uint32
	FMask     uint32
	NMin      int32
	NMax      int32
	NPage     uint32
	NPos      int32
	NTrackPos int32
}

// AddSimpleMenu 创建简单菜单并挂到窗口上。
// 参数hwnd：窗口句柄。
// 返回值：true 表示菜单已成功创建并设置到窗口，false 表示创建或设置失败。
func AddSimpleMenu(hwnd uintptr) bool {
	menu, _, _ := procCreateMenu.Call()
	if menu == 0 {
		return false
	}
	procAppendMenuW.Call(menu, MF_STRING, 1001, utf16Ptr("打开"))
	procAppendMenuW.Call(menu, MF_STRING, 1002, utf16Ptr("退出"))
	ret, _, _ := procSetMenu.Call(hwnd, menu)
	procDrawMenuBar.Call(hwnd)
	return ret != 0
}

// CreateMenuHandle 创建一个空菜单句柄。
// 返回值：返回新建菜单句柄；返回 0 表示创建失败。
func CreateMenuHandle() uintptr {
	menu, _, _ := procCreateMenu.Call()
	return menu
}

// AppendMenuString 向菜单追加字符串菜单项。
// 参数menu：菜单句柄。
// 参数id：资源、命令或计时器 ID。
// 参数text：要传入的文本内容。
// 返回值：true 表示菜单项追加成功，false 表示追加失败。
func AppendMenuString(menu uintptr, id uint32, text string) bool {
	ret, _, _ := procAppendMenuW.Call(menu, MF_STRING, uintptr(id), utf16Ptr(text))
	return ret != 0
}

// SetWindowMenu 将菜单设置到窗口上。
// 参数hwnd：窗口句柄。
// 参数menu：菜单句柄。
// 返回值：true 表示菜单已成功设置到窗口，false 表示设置失败。
func SetWindowMenu(hwnd, menu uintptr) bool {
	ret, _, _ := procSetMenu.Call(hwnd, menu)
	procDrawMenuBar.Call(hwnd)
	return ret != 0
}

// DrawMenuBar 重绘窗口菜单栏。
// 参数hwnd：窗口句柄。
// 返回值：true 表示菜单栏重绘成功，false 表示重绘失败。
func DrawMenuBar(hwnd uintptr) bool {
	ret, _, _ := procDrawMenuBar.Call(hwnd)
	return ret != 0
}

// CreatePopupMenuHandle 创建一个空弹出菜单句柄。
// 返回值：返回新建弹出菜单句柄；返回 0 表示创建失败。
func CreatePopupMenuHandle() uintptr {
	menu, _, _ := procCreatePopupMenu.Call()
	return menu
}

// TrackPopupMenu 在指定位置显示弹出菜单。
// 参数menu：菜单句柄。
// 参数flags：调用标志位。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 参数hwnd：窗口句柄。
// 返回值：true 表示弹出菜单显示并成功跟踪，false 表示显示或跟踪失败。
func TrackPopupMenu(menu uintptr, flags uint32, x, y int32, hwnd uintptr) bool {
	ret, _, _ := procTrackPopupMenu.Call(menu, uintptr(flags), uintptr(x), uintptr(y), 0, hwnd, 0)
	return ret != 0
}

// ShowContextMenu 创建并显示一个简单右键菜单。
// 参数hwnd：窗口句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 返回值：无返回值；函数会创建并显示一个包含复制和粘贴项的右键菜单。
func ShowContextMenu(hwnd uintptr, x, y int32) {
	menu, _, _ := procCreatePopupMenu.Call()
	if menu == 0 {
		return
	}
	defer procDestroyMenu.Call(menu)
	procAppendMenuW.Call(menu, MF_STRING, 2001, utf16Ptr("复制"))
	procAppendMenuW.Call(menu, MF_STRING, 2002, utf16Ptr("粘贴"))
	procTrackPopupMenu.Call(menu, TPM_RIGHTBUTTON, uintptr(x), uintptr(y), 0, hwnd, 0)
}

// DestroyMenu 销毁菜单句柄。
// 参数menu：菜单句柄。
// 返回值：true 表示菜单句柄销毁成功，false 表示销毁失败。
func DestroyMenu(menu uintptr) bool {
	ret, _, _ := procDestroyMenu.Call(menu)
	return ret != 0
}

// DisableMenuCommand 禁用指定菜单命令。
// 参数menu：菜单句柄。
// 参数id：资源、命令或计时器 ID。
// 返回值：无返回值；函数会把指定菜单命令置为灰色禁用状态。
func DisableMenuCommand(menu uintptr, id uint32) {
	procEnableMenuItem.Call(menu, uintptr(id), MF_BYCOMMAND|MF_GRAYED)
}

// EnableMenuCommand 启用或禁用指定菜单命令。
// 参数menu：菜单句柄。
// 参数id：资源、命令或计时器 ID。
// 参数enabled：是否启用。
// 返回值：无返回值；函数会根据 enabled 启用或禁用指定菜单命令。
func EnableMenuCommand(menu uintptr, id uint32, enabled bool) {
	flag := uintptr(MF_GRAYED)
	if enabled {
		flag = MF_ENABLED
	}
	procEnableMenuItem.Call(menu, uintptr(id), MF_BYCOMMAND|flag)
}

// CheckMenuCommand 设置指定菜单命令的勾选状态。
// 参数menu：菜单句柄。
// 参数id：资源、命令或计时器 ID。
// 参数checked：是否勾选。
// 返回值：无返回值；函数会根据 checked 设置菜单命令的勾选状态。
func CheckMenuCommand(menu uintptr, id uint32, checked bool) {
	flag := uintptr(MF_UNCHECKED)
	if checked {
		flag = MF_CHECKED
	}
	procCheckMenuItem.Call(menu, uintptr(id), MF_BYCOMMAND|flag)
}

// ShowVerticalScrollBar 显示或隐藏垂直滚动条。
// 参数hwnd：窗口句柄。
// 参数show：是否显示。
// 返回值：无返回值；函数会显示或隐藏窗口的垂直滚动条。
func ShowVerticalScrollBar(hwnd uintptr, show bool) {
	procShowScrollBar.Call(hwnd, SB_VERT, boolArg(show))
}

// ShowScrollBar 显示或隐藏指定滚动条。
// 参数hwnd：窗口句柄。
// 参数bar：滚动条类型。
// 参数show：是否显示。
// 返回值：true 表示滚动条显示状态设置成功，false 表示设置失败。
func ShowScrollBar(hwnd uintptr, bar int, show bool) bool {
	ret, _, _ := procShowScrollBar.Call(hwnd, uintptr(bar), boolArg(show))
	return ret != 0
}

// GetVerticalScrollInfo 获取垂直滚动条信息。
// 参数hwnd：窗口句柄。
// 返回值：返回垂直滚动条信息和是否获取成功；true 表示信息有效，false 表示获取失败。
func GetVerticalScrollInfo(hwnd uintptr) (SCROLLINFO, bool) {
	info := SCROLLINFO{CbSize: uint32(unsafe.Sizeof(SCROLLINFO{})), FMask: SIF_ALL}
	ret, _, _ := procGetScrollInfo.Call(hwnd, SB_VERT, uintptr(unsafe.Pointer(&info)))
	return info, ret != 0
}

// GetScrollInfo 获取指定滚动条信息。
// 参数hwnd：窗口句柄。
// 参数bar：滚动条类型。
// 参数info：滚动条信息结构指针。
// 返回值：true 表示滚动条信息获取成功，false 表示获取失败。
func GetScrollInfo(hwnd uintptr, bar int, info *SCROLLINFO) bool {
	if info != nil && info.CbSize == 0 {
		info.CbSize = uint32(unsafe.Sizeof(SCROLLINFO{}))
	}
	ret, _, _ := procGetScrollInfo.Call(hwnd, uintptr(bar), uintptr(unsafe.Pointer(info)))
	return ret != 0
}

// SetVerticalScrollInfo 设置垂直滚动条信息。
// 参数hwnd：窗口句柄。
// 参数info：滚动条信息结构指针。
// 参数redraw：是否立即重绘。
// 返回值：返回设置后的滚动条位置；返回值含义与 SetScrollInfo 一致。
func SetVerticalScrollInfo(hwnd uintptr, info *SCROLLINFO, redraw bool) int32 {
	if info != nil && info.CbSize == 0 {
		info.CbSize = uint32(unsafe.Sizeof(SCROLLINFO{}))
	}
	ret, _, _ := procSetScrollInfo.Call(hwnd, SB_VERT, uintptr(unsafe.Pointer(info)), boolArg(redraw))
	return int32(ret)
}

// SetScrollInfo 设置指定滚动条信息。
// 参数hwnd：窗口句柄。
// 参数bar：滚动条类型。
// 参数info：滚动条信息结构指针。
// 参数redraw：是否立即重绘。
// 返回值：返回设置后的滚动条位置；返回 0 可能表示失败或位置为 0。
func SetScrollInfo(hwnd uintptr, bar int, info *SCROLLINFO, redraw bool) int32 {
	if info != nil && info.CbSize == 0 {
		info.CbSize = uint32(unsafe.Sizeof(SCROLLINFO{}))
	}
	ret, _, _ := procSetScrollInfo.Call(hwnd, uintptr(bar), uintptr(unsafe.Pointer(info)), boolArg(redraw))
	return int32(ret)
}

// GetDialogItem 获取对话框中的控件句柄。
// 参数hwndDlg：对话框窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 返回值：返回对话框控件窗口句柄；返回 0 表示未找到或获取失败。
func GetDialogItem(hwndDlg uintptr, id int) uintptr {
	hwnd, _, _ := procGetDlgItem.Call(hwndDlg, uintptr(id))
	return hwnd
}

// SetDialogItemText 设置对话框控件文本。
// 参数hwndDlg：对话框窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 参数text：要传入的文本内容。
// 返回值：true 表示控件文本设置成功，false 表示设置失败。
func SetDialogItemText(hwndDlg uintptr, id int, text string) bool {
	ret, _, _ := procSetDlgItemTextW.Call(hwndDlg, uintptr(id), utf16Ptr(text))
	return ret != 0
}

// GetDialogItemText 获取对话框控件文本。
// 参数hwndDlg：对话框窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 返回值：返回控件文本；控件无文本、未找到或读取失败时返回空字符串。
func GetDialogItemText(hwndDlg uintptr, id int) string {
	buf := make([]uint16, 256)
	procGetDlgItemTextW.Call(hwndDlg, uintptr(id), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return syscall.UTF16ToString(buf)
}

// SetCheckbox 设置对话框复选框状态。
// 参数hwndDlg：对话框窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 参数checked：是否勾选。
// 返回值：无返回值；函数会设置复选框为勾选或未勾选状态。
func SetCheckbox(hwndDlg uintptr, id int, checked bool) {
	v := uintptr(BST_UNCHECKED)
	if checked {
		v = BST_CHECKED
	}
	procCheckDlgButton.Call(hwndDlg, uintptr(id), v)
}

// CheckDlgButton 设置对话框按钮的勾选状态。
// 参数hwndDlg：对话框窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 参数check：按钮勾选状态值。
// 返回值：true 表示按钮勾选状态设置成功，false 表示设置失败。
func CheckDlgButton(hwndDlg uintptr, id int, check uint32) bool {
	ret, _, _ := procCheckDlgButton.Call(hwndDlg, uintptr(id), uintptr(check))
	return ret != 0
}

// IsCheckboxChecked 判断对话框复选框是否已勾选。
// 参数hwndDlg：对话框窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 返回值：true 表示复选框当前为勾选状态，false 表示未勾选或读取失败。
func IsCheckboxChecked(hwndDlg uintptr, id int) bool {
	ret, _, _ := procIsDlgButtonChecked.Call(hwndDlg, uintptr(id))
	return ret == BST_CHECKED
}
