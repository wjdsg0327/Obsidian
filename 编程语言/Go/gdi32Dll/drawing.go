//go:build windows

package main

var (
    procCreatePen        = gdi32.NewProc("CreatePen")
    procCreateSolidBrush = gdi32.NewProc("CreateSolidBrush")
    procSelectObject     = gdi32.NewProc("SelectObject")
    procDeleteObject     = gdi32.NewProc("DeleteObject")
    procMoveToEx         = gdi32.NewProc("MoveToEx")
    procLineTo           = gdi32.NewProc("LineTo")
    procRectangle        = gdi32.NewProc("Rectangle")
    procEllipse          = gdi32.NewProc("Ellipse")
    procRoundRect        = gdi32.NewProc("RoundRect")
)

const PS_SOLID = 0

func DrawBasicShapes(hdc uintptr) {
    pen, _, _ := procCreatePen.Call(PS_SOLID, 3, RGB(255, 0, 0))
    oldPen, _, _ := procSelectObject.Call(hdc, pen)

    brush, _, _ := procCreateSolidBrush.Call(RGB(220, 240, 255))
    oldBrush, _, _ := procSelectObject.Call(hdc, brush)

    procMoveToEx.Call(hdc, 50, 50, 0)
    procLineTo.Call(hdc, 300, 50)
    procRectangle.Call(hdc, 50, 80, 300, 200)
    procEllipse.Call(hdc, 330, 80, 520, 200)
    procRoundRect.Call(hdc, 50, 230, 300, 340, 24, 24)

    procSelectObject.Call(hdc, oldBrush)
    procSelectObject.Call(hdc, oldPen)
    procDeleteObject.Call(brush)
    procDeleteObject.Call(pen)
}
