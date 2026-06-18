//go:build windows

package user32Dll

import "unsafe"

var (
	procFillRect          = user32.NewProc("FillRect")
	procFrameRect         = user32.NewProc("FrameRect")
	procInvertRect        = user32.NewProc("InvertRect")
	procDrawEdge          = user32.NewProc("DrawEdge")
	procDrawCaption       = user32.NewProc("DrawCaption")
	procDrawIconEx        = user32.NewProc("DrawIconEx")
	procScrollWindow      = user32.NewProc("ScrollWindow")
	procScrollWindowEx    = user32.NewProc("ScrollWindowEx")
	procDrawAnimatedRects = user32.NewProc("DrawAnimatedRects")
)

const (
	DI_MASK        = 0x0001
	DI_IMAGE       = 0x0002
	DI_NORMAL      = DI_MASK | DI_IMAGE
	DI_COMPAT      = 0x0004
	DI_DEFAULTSIZE = 0x0008

	BDR_RAISEDOUTER = 0x0001
	BDR_SUNKENOUTER = 0x0002
	EDGE_RAISED     = BDR_RAISEDOUTER | BDR_RAISEDOUTER
	EDGE_SUNKEN     = BDR_SUNKENOUTER | BDR_SUNKENOUTER
)

// FillRect 填充矩形区域。
// 参数hdc：设备上下文句柄。
// 参数rect：矩形区域指针，nil 表示不指定。
// 参数brush：画刷句柄。
// 返回值：true 表示填充成功，false 表示填充失败。
func FillRect(hdc uintptr, rect *RECT, brush uintptr) bool {
	ret, _, _ := procFillRect.Call(hdc, uintptr(unsafe.Pointer(rect)), brush)
	return ret != 0
}

// FrameRect 绘制矩形边框。
// 参数hdc：设备上下文句柄。
// 参数rect：矩形区域指针，nil 表示不指定。
// 参数brush：画刷句柄。
// 返回值：true 表示边框绘制成功，false 表示绘制失败。
func FrameRect(hdc uintptr, rect *RECT, brush uintptr) bool {
	ret, _, _ := procFrameRect.Call(hdc, uintptr(unsafe.Pointer(rect)), brush)
	return ret != 0
}

// InvertRect 反转矩形区域颜色。
// 参数hdc：设备上下文句柄。
// 参数rect：矩形区域指针，nil 表示不指定。
// 返回值：true 表示颜色反转成功，false 表示反转失败。
func InvertRect(hdc uintptr, rect *RECT) bool {
	ret, _, _ := procInvertRect.Call(hdc, uintptr(unsafe.Pointer(rect)))
	return ret != 0
}

// DrawEdge 绘制边框边缘。
// 参数hdc：设备上下文句柄。
// 参数rect：矩形区域指针，nil 表示不指定。
// 参数edge：边缘样式标志。
// 参数flags：绘制标志位。
// 返回值：true 表示边缘绘制成功，false 表示绘制失败。
func DrawEdge(hdc uintptr, rect *RECT, edge uint32, flags uint32) bool {
	ret, _, _ := procDrawEdge.Call(hdc, uintptr(unsafe.Pointer(rect)), uintptr(edge), uintptr(flags))
	return ret != 0
}

// DrawCaption 绘制窗口标题栏。
// 参数hwnd：窗口句柄。
// 参数hdc：设备上下文句柄。
// 参数rect：矩形区域指针，nil 表示不指定。
// 参数flags：绘制标志位。
// 返回值：true 表示标题栏绘制成功，false 表示绘制失败。
func DrawCaption(hwnd uintptr, hdc uintptr, rect *RECT, flags uint32) bool {
	ret, _, _ := procDrawCaption.Call(hwnd, hdc, uintptr(unsafe.Pointer(rect)), uintptr(flags))
	return ret != 0
}

// DrawIconEx 绘制图标并可指定尺寸和绘制标志。
// 参数hdc：设备上下文句柄。
// 参数x：横向坐标。
// 参数y：纵向坐标。
// 参数icon：图标句柄。
// 参数width：绘制宽度。
// 参数height：绘制高度。
// 参数stepIfAniCur：动画光标帧索引。
// 参数brush：闪烁背景画刷句柄。
// 参数flags：绘制标志位。
// 返回值：true 表示图标绘制成功，false 表示绘制失败。
func DrawIconEx(hdc uintptr, x, y int32, icon uintptr, width, height int32, stepIfAniCur uint32, brush uintptr, flags uint32) bool {
	ret, _, _ := procDrawIconEx.Call(hdc, uintptr(x), uintptr(y), icon, uintptr(width), uintptr(height), uintptr(stepIfAniCur), brush, uintptr(flags))
	return ret != 0
}

// ScrollWindow 滚动窗口客户区内容。
// 参数hwnd：窗口句柄。
// 参数dx：水平滚动像素，正数向右滚动。
// 参数dy：垂直滚动像素，正数向下滚动。
// 参数rect：滚动区域指针，nil 表示整个客户区。
// 参数clipRect：裁剪区域指针，nil 表示不指定。
// 返回值：true 表示滚动请求成功，false 表示滚动失败。
func ScrollWindow(hwnd uintptr, dx, dy int32, rect *RECT, clipRect *RECT) bool {
	ret, _, _ := procScrollWindow.Call(hwnd, uintptr(dx), uintptr(dy), uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(clipRect)))
	return ret != 0
}

// ScrollWindowEx 扩展滚动窗口客户区内容。
// 参数hwnd：窗口句柄。
// 参数dx：水平滚动像素，正数向右滚动。
// 参数dy：垂直滚动像素，正数向下滚动。
// 参数rect：滚动区域指针，nil 表示整个客户区。
// 参数clipRect：裁剪区域指针，nil 表示不指定。
// 参数updateRegion：接收更新区域的区域句柄。
// 参数updateRect：接收更新矩形的指针。
// 参数flags：滚动标志位。
// 返回值：返回滚动后无效区域的复杂度代码；具体含义与 ScrollWindowEx 返回值一致。
func ScrollWindowEx(hwnd uintptr, dx, dy int32, rect *RECT, clipRect *RECT, updateRegion uintptr, updateRect *RECT, flags uint32) int32 {
	ret, _, _ := procScrollWindowEx.Call(hwnd, uintptr(dx), uintptr(dy), uintptr(unsafe.Pointer(rect)), uintptr(unsafe.Pointer(clipRect)), updateRegion, uintptr(unsafe.Pointer(updateRect)), uintptr(flags))
	return int32(ret)
}

// DrawAnimatedRects 绘制窗口最小化或还原动画矩形。
// 参数hwnd：窗口句柄。
// 参数animation：动画类型。
// 参数fromRect：动画起始矩形。
// 参数toRect：动画结束矩形。
// 返回值：true 表示动画矩形绘制成功，false 表示绘制失败。
func DrawAnimatedRects(hwnd uintptr, animation int32, fromRect *RECT, toRect *RECT) bool {
	ret, _, _ := procDrawAnimatedRects.Call(hwnd, uintptr(animation), uintptr(unsafe.Pointer(fromRect)), uintptr(unsafe.Pointer(toRect)))
	return ret != 0
}
