//go:build windows

package user32Dll

import "unsafe"

var (
	procInsertMenuW        = user32.NewProc("InsertMenuW")
	procInsertMenuItemW    = user32.NewProc("InsertMenuItemW")
	procModifyMenuW        = user32.NewProc("ModifyMenuW")
	procDeleteMenu         = user32.NewProc("DeleteMenu")
	procRemoveMenu         = user32.NewProc("RemoveMenu")
	procGetMenuItemInfoW   = user32.NewProc("GetMenuItemInfoW")
	procSetMenuItemInfoW   = user32.NewProc("SetMenuItemInfoW")
	procGetSubMenu         = user32.NewProc("GetSubMenu")
	procGetMenuState       = user32.NewProc("GetMenuState")
	procGetMenuItemCount   = user32.NewProc("GetMenuItemCount")
	procGetMenuItemID      = user32.NewProc("GetMenuItemID")
	procSendDlgItemMessage = user32.NewProc("SendDlgItemMessageW")
	procMapDialogRect      = user32.NewProc("MapDialogRect")
	procEndDialog          = user32.NewProc("EndDialog")
	procCreateDialogParamW = user32.NewProc("CreateDialogParamW")
	procDialogBoxParamW    = user32.NewProc("DialogBoxParamW")
)

const (
	MIIM_STATE    = 0x00000001
	MIIM_ID       = 0x00000002
	MIIM_SUBMENU  = 0x00000004
	MIIM_TYPE     = 0x00000010
	MIIM_DATA     = 0x00000020
	MIIM_STRING   = 0x00000040
	MIIM_BITMAP   = 0x00000080
	MIIM_FTYPE    = 0x00000100
	MF_BYPOSITION = 0x00000400
)

// MENUITEMINFO 表示菜单项信息结构。
type MENUITEMINFO struct {
	CbSize        uint32
	FMask         uint32
	FType         uint32
	FState        uint32
	WID           uint32
	HSubMenu      uintptr
	HbmpChecked   uintptr
	HbmpUnchecked uintptr
	DwItemData    uintptr
	DwTypeData    uintptr
	Cch           uint32
	HbmpItem      uintptr
}

// InsertMenuString 插入字符串菜单项。
// 参数menu：菜单句柄。
// 参数position：插入位置或命令 ID。
// 参数flags：菜单插入标志。
// 参数id：新菜单项 ID。
// 参数text：要传入的文本内容。
// 返回值：true 表示菜单项插入成功，false 表示插入失败。
func InsertMenuString(menu uintptr, position uint32, flags uint32, id uint32, text string) bool {
	ret, _, _ := procInsertMenuW.Call(menu, uintptr(position), uintptr(flags|MF_STRING), uintptr(id), utf16Ptr(text))
	return ret != 0
}

// ModifyMenuString 修改字符串菜单项。
// 参数menu：菜单句柄。
// 参数position：要修改的位置或命令 ID。
// 参数flags：菜单定位和类型标志。
// 参数id：新的菜单项 ID。
// 参数text：要传入的文本内容。
// 返回值：true 表示菜单项修改成功，false 表示修改失败。
func ModifyMenuString(menu uintptr, position uint32, flags uint32, id uint32, text string) bool {
	ret, _, _ := procModifyMenuW.Call(menu, uintptr(position), uintptr(flags|MF_STRING), uintptr(id), utf16Ptr(text))
	return ret != 0
}

// DeleteMenu 删除菜单项并销毁其子菜单。
// 参数menu：菜单句柄。
// 参数position：要删除的位置或命令 ID。
// 参数flags：菜单定位标志。
// 返回值：true 表示菜单项删除成功，false 表示删除失败。
func DeleteMenu(menu uintptr, position uint32, flags uint32) bool {
	ret, _, _ := procDeleteMenu.Call(menu, uintptr(position), uintptr(flags))
	return ret != 0
}

// RemoveMenu 移除菜单项但不销毁其子菜单。
// 参数menu：菜单句柄。
// 参数position：要移除的位置或命令 ID。
// 参数flags：菜单定位标志。
// 返回值：true 表示菜单项移除成功，false 表示移除失败。
func RemoveMenu(menu uintptr, position uint32, flags uint32) bool {
	ret, _, _ := procRemoveMenu.Call(menu, uintptr(position), uintptr(flags))
	return ret != 0
}

// InsertMenuItem 插入菜单项信息。
// 参数menu：菜单句柄。
// 参数item：插入位置或命令 ID。
// 参数byPosition：true 表示按位置插入，false 表示按命令 ID 插入。
// 参数info：菜单项信息结构指针。
// 返回值：true 表示菜单项插入成功，false 表示插入失败。
func InsertMenuItem(menu uintptr, item uint32, byPosition bool, info *MENUITEMINFO) bool {
	if info != nil && info.CbSize == 0 {
		info.CbSize = uint32(unsafe.Sizeof(MENUITEMINFO{}))
	}
	ret, _, _ := procInsertMenuItemW.Call(menu, uintptr(item), boolArg(byPosition), uintptr(unsafe.Pointer(info)))
	return ret != 0
}

// GetMenuItemInfo 获取菜单项信息。
// 参数menu：菜单句柄。
// 参数item：菜单位置或命令 ID。
// 参数byPosition：true 表示按位置查找，false 表示按命令 ID 查找。
// 参数info：菜单项信息结构指针。
// 返回值：true 表示菜单项信息获取成功，false 表示获取失败。
func GetMenuItemInfo(menu uintptr, item uint32, byPosition bool, info *MENUITEMINFO) bool {
	if info != nil && info.CbSize == 0 {
		info.CbSize = uint32(unsafe.Sizeof(MENUITEMINFO{}))
	}
	ret, _, _ := procGetMenuItemInfoW.Call(menu, uintptr(item), boolArg(byPosition), uintptr(unsafe.Pointer(info)))
	return ret != 0
}

// SetMenuItemInfo 设置菜单项信息。
// 参数menu：菜单句柄。
// 参数item：菜单位置或命令 ID。
// 参数byPosition：true 表示按位置查找，false 表示按命令 ID 查找。
// 参数info：菜单项信息结构指针。
// 返回值：true 表示菜单项信息设置成功，false 表示设置失败。
func SetMenuItemInfo(menu uintptr, item uint32, byPosition bool, info *MENUITEMINFO) bool {
	if info != nil && info.CbSize == 0 {
		info.CbSize = uint32(unsafe.Sizeof(MENUITEMINFO{}))
	}
	ret, _, _ := procSetMenuItemInfoW.Call(menu, uintptr(item), boolArg(byPosition), uintptr(unsafe.Pointer(info)))
	return ret != 0
}

// GetSubMenu 获取指定位置的子菜单。
// 参数menu：菜单句柄。
// 参数position：菜单项位置。
// 返回值：返回子菜单句柄；返回 0 表示没有子菜单或获取失败。
func GetSubMenu(menu uintptr, position int32) uintptr {
	ret, _, _ := procGetSubMenu.Call(menu, uintptr(position))
	return ret
}

// GetMenuState 获取菜单项状态。
// 参数menu：菜单句柄。
// 参数id：资源、命令或计时器 ID。
// 参数flags：菜单定位标志。
// 返回值：返回菜单项状态标志；返回 0xFFFFFFFF 表示获取失败。
func GetMenuState(menu uintptr, id uint32, flags uint32) uint32 {
	ret, _, _ := procGetMenuState.Call(menu, uintptr(id), uintptr(flags))
	return uint32(ret)
}

// GetMenuItemCount 获取菜单项数量。
// 参数menu：菜单句柄。
// 返回值：返回菜单项数量；返回 -1 表示获取失败。
func GetMenuItemCount(menu uintptr) int32 {
	ret, _, _ := procGetMenuItemCount.Call(menu)
	return int32(ret)
}

// GetMenuItemID 获取指定位置菜单项的命令 ID。
// 参数menu：菜单句柄。
// 参数position：菜单项位置。
// 返回值：返回菜单项命令 ID；返回 0xFFFFFFFF 表示该项是子菜单或获取失败。
func GetMenuItemID(menu uintptr, position int32) uint32 {
	ret, _, _ := procGetMenuItemID.Call(menu, uintptr(position))
	return uint32(ret)
}

// SendDialogItemMessage 向对话框控件发送消息。
// 参数hwndDlg：对话框窗口句柄。
// 参数id：资源、命令或计时器 ID。
// 参数msg：窗口消息编号。
// 参数wParam：消息的 wParam 参数。
// 参数lParam：消息的 lParam 参数。
// 返回值：返回控件处理消息后的结果值，具体含义由消息类型决定。
func SendDialogItemMessage(hwndDlg uintptr, id int, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procSendDlgItemMessage.Call(hwndDlg, uintptr(id), uintptr(msg), wParam, lParam)
	return ret
}

// MapDialogRect 将对话框单位转换为像素矩形。
// 参数hwndDlg：对话框窗口句柄。
// 参数rect：矩形区域指针，nil 表示不指定。
// 返回值：true 表示转换成功，false 表示转换失败。
func MapDialogRect(hwndDlg uintptr, rect *RECT) bool {
	ret, _, _ := procMapDialogRect.Call(hwndDlg, uintptr(unsafe.Pointer(rect)))
	return ret != 0
}

// EndDialog 结束模态对话框。
// 参数hwndDlg：对话框窗口句柄。
// 参数result：对话框返回给调用者的结果值。
// 返回值：true 表示对话框结束成功，false 表示结束失败。
func EndDialog(hwndDlg uintptr, result uintptr) bool {
	ret, _, _ := procEndDialog.Call(hwndDlg, result)
	return ret != 0
}

// CreateDialogParam 创建非模态对话框。
// 参数instance：模块实例句柄。
// 参数templateName：对话框模板名称。
// 参数parent：父窗口句柄。
// 参数dialogProc：对话框过程回调指针。
// 参数param：传给对话框过程的自定义参数。
// 返回值：返回对话框窗口句柄；返回 0 表示创建失败。
func CreateDialogParam(instance uintptr, templateName string, parent uintptr, dialogProc uintptr, param uintptr) uintptr {
	ret, _, _ := procCreateDialogParamW.Call(instance, utf16Ptr(templateName), parent, dialogProc, param)
	return ret
}

// DialogBoxParam 创建并运行模态对话框。
// 参数instance：模块实例句柄。
// 参数templateName：对话框模板名称。
// 参数parent：父窗口句柄。
// 参数dialogProc：对话框过程回调指针。
// 参数param：传给对话框过程的自定义参数。
// 返回值：返回 EndDialog 传出的结果值；返回 -1 表示创建对话框失败。
func DialogBoxParam(instance uintptr, templateName string, parent uintptr, dialogProc uintptr, param uintptr) int32 {
	ret, _, _ := procDialogBoxParamW.Call(instance, utf16Ptr(templateName), parent, dialogProc, param)
	return int32(ret)
}
