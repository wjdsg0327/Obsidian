//go:build windows

package main

import "unsafe"

var (
    procGetPixel      = gdi32.NewProc("GetPixel")
    procSetPixel      = gdi32.NewProc("SetPixel")
    procCreateRectRgn = gdi32.NewProc("CreateRectRgn")
    procSelectClipRgn = gdi32.NewProc("SelectClipRgn")
    procCombineRgn    = gdi32.NewProc("CombineRgn")
    procBeginPath     = gdi32.NewProc("BeginPath")
    procEndPath       = gdi32.NewProc("EndPath")
    procStrokePath    = gdi32.NewProc("StrokePath")
    procPolyline      = gdi32.NewProc("Polyline")
    procPolygon       = gdi32.NewProc("Polygon")
)

const RGN_OR = 2

func PixelDemo(hdc uintptr) uintptr {
    old, _, _ := procGetPixel.Call(hdc, 10, 10)
    procSetPixel.Call(hdc, 10, 10, RGB(255, 0, 0))
    return old
}

func DrawWithClip(hdc uintptr) {
    rgn, _, _ := procCreateRectRgn.Call(50, 50, 250, 150)
    procSelectClipRgn.Call(hdc, rgn)
    DrawTextOut(hdc, 60, 80, "只在裁剪区域内显示")
    procSelectClipRgn.Call(hdc, 0)
    procDeleteObject.Call(rgn)
}

func CombineRegionDemo() uintptr {
    r1, _, _ := procCreateRectRgn.Call(0, 0, 100, 100)
    r2, _, _ := procCreateRectRgn.Call(50, 50, 150, 150)
    dst, _, _ := procCreateRectRgn.Call(0, 0, 0, 0)
    procCombineRgn.Call(dst, r1, r2, RGN_OR)
    procDeleteObject.Call(r1)
    procDeleteObject.Call(r2)
    return dst
}

func PathDemo(hdc uintptr) {
    procBeginPath.Call(hdc)
    procMoveToEx.Call(hdc, 50, 50, 0)
    procLineTo.Call(hdc, 150, 50)
    procLineTo.Call(hdc, 150, 150)
    procLineTo.Call(hdc, 50, 150)
    procLineTo.Call(hdc, 50, 50)
    procEndPath.Call(hdc)
    procStrokePath.Call(hdc)
}

func DrawPoly(hdc uintptr) {
    pts := []POINT{{50, 50}, {120, 80}, {180, 30}, {260, 140}}
    procPolyline.Call(hdc, uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
    tri := []POINT{{80, 220}, {180, 220}, {130, 150}}
    procPolygon.Call(hdc, uintptr(unsafe.Pointer(&tri[0])), uintptr(len(tri)))
}
