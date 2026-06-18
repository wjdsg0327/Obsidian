//go:build windows

package user32Dll

import "unsafe"

var (
	procDrawTextW        = user32.NewProc("DrawTextW")
	procLoadCursorW      = user32.NewProc("LoadCursorW")
	procSetCursor        = user32.NewProc("SetCursor")
	procShowCursor       = user32.NewProc("ShowCursor")
	procLoadIconW        = user32.NewProc("LoadIconW")
	procDrawIcon         = user32.NewProc("DrawIcon")
	procDestroyIcon      = user32.NewProc("DestroyIcon")
	procDrawFocusRect    = user32.NewProc("DrawFocusRect")
	procDrawFrameControl = user32.NewProc("DrawFrameControl")
)

const (
	DT_CENTER     = 0x00000001
	DT_VCENTER    = 0x00000004
	DT_SINGLELINE = 0x00000020

	IDC_HAND        = uintptr(32649)
	IDI_INFORMATION = uintptr(32516)

	DFC_BUTTON      = 4
	DFCS_BUTTONPUSH = 0x0010
)

// DrawCenteredText 在指定矩形中绘制居中文本。
// 参数hdc：设备上下文句柄。
// 参数text：要传入的文本内容。
// 参数rect：矩形区域指针，nil 表示不指定。
// 返回值：无返回值；函数会向指定设备上下文绘制文本。
func DrawCenteredText(hdc uintptr, text string, rect RECT) {
	procDrawTextW.Call(hdc, utf16Ptr(text), ^uintptr(0), uintptr(unsafe.Pointer(&rect)), DT_CENTER|DT_VCENTER|DT_SINGLELINE)
}

// SetHandCursor 加载并设置系统手型光标。
// 返回值：无返回值；函数会把当前光标设置为系统手型光标。
func SetHandCursor() {
	cursor, _, _ := procLoadCursorW.Call(0, IDC_HAND)
	procSetCursor.Call(cursor)
}

// HideCursor 隐藏鼠标光标。
// 返回值：无返回值；函数会减少系统光标显示计数，计数小于 0 时光标隐藏。
func HideCursor() {
	procShowCursor.Call(0)
}

// ShowCursorAgain 显示鼠标光标。
// 返回值：无返回值；函数会增加系统光标显示计数，计数大于等于 0 时光标显示。
func ShowCursorAgain() {
	procShowCursor.Call(1)
}

// LoadSystemCursor 加载指定系统光标。
// 参数cursorID：系统光标资源 ID。
// 返回值：返回系统光标句柄；返回 0 表示加载失败。
func LoadSystemCursor(cursorID uintptr) uintptr {
	cursor, _, _ := procLoadCursorW.Call(0, cursorID)
	return cursor
}

// SetCursorHandle 设置当前光标句柄。
// 参数cursor：光标句柄。
// 返回值：返回之前的光标句柄；返回 0 表示之前没有光标或设置失败。
func SetCursorHandle(cursor uintptr) uintptr {
	old, _, _ := procSetCursor.Call(cursor)
	return old
}

// DrawInfoIcon 绘制系统信息图标。
// 参数hdc：设备上下文句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 返回值：无返回值；函数会在指定位置绘制系统信息图标。
func DrawInfoIcon(hdc uintptr, x, y int32) {
	icon, _, _ := procLoadIconW.Call(0, IDI_INFORMATION)
	procDrawIcon.Call(hdc, uintptr(x), uintptr(y), icon)
}

// LoadSystemIcon 加载指定系统图标。
// 参数iconID：系统图标资源 ID。
// 返回值：返回系统图标句柄；返回 0 表示加载失败。
func LoadSystemIcon(iconID uintptr) uintptr {
	icon, _, _ := procLoadIconW.Call(0, iconID)
	return icon
}

// DrawIconAt 在指定设备上下文位置绘制图标。
// 参数hdc：设备上下文句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 参数icon：图标句柄。
// 返回值：true 表示图标绘制成功，false 表示绘制失败。
func DrawIconAt(hdc uintptr, x, y int32, icon uintptr) bool {
	ret, _, _ := procDrawIcon.Call(hdc, uintptr(x), uintptr(y), icon)
	return ret != 0
}

// DestroyIcon 销毁图标句柄。
// 参数icon：图标句柄。
// 返回值：true 表示图标句柄已销毁，false 表示销毁失败。
func DestroyIcon(icon uintptr) bool {
	ret, _, _ := procDestroyIcon.Call(icon)
	return ret != 0
}

// DrawFocus 绘制焦点矩形。
// 参数hdc：设备上下文句柄。
// 参数r：矩形区域。
// 返回值：true 表示焦点矩形绘制成功，false 表示绘制失败。
func DrawFocus(hdc uintptr, r RECT) bool {
	ret, _, _ := procDrawFocusRect.Call(hdc, uintptr(unsafe.Pointer(&r)))
	return ret != 0
}

// DrawPushButtonFrame 绘制标准按钮外框。
// 参数hdc：设备上下文句柄。
// 参数r：矩形区域。
// 返回值：true 表示标准按钮边框绘制成功，false 表示绘制失败。
func DrawPushButtonFrame(hdc uintptr, r RECT) bool {
	ret, _, _ := procDrawFrameControl.Call(hdc, uintptr(unsafe.Pointer(&r)), DFC_BUTTON, DFCS_BUTTONPUSH)
	return ret != 0
}
