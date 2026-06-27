---
title: Go gdi32.dll API 完整案例集
date: 2026-06-27
tags:
  - Go
  - Windows
  - Win32
  - gdi32.dll
  - API案例
  - GDI
---

# Go gdi32.dll API 完整案例集

> 这是《Go 调用 gdi32.dll 操作整理》的案例版。  
> 目标：不是只列 API 名，而是让常用 `gdi32.dll` API 都有一个可理解、可改造的 Go 使用案例。  
> 说明：GDI 真实 API 数量很大，本文先覆盖 DC、GDI 对象、基础绘图、文本、位图截图、区域、路径、坐标、打印等常用部分。

---

## 0. 通用准备代码

后面的案例默认使用下面这套基础代码。每个 API 案例为了避免重复，不一定都重新写 `package main`，但都能嵌入这个模板中使用。

```go
package main

import (
    "fmt"
    "syscall"
    "unsafe"
)

var (
    user32 = syscall.NewLazyDLL("user32.dll")
    gdi32  = syscall.NewLazyDLL("gdi32.dll")
)

func utf16Ptr(s string) uintptr {
    p, err := syscall.UTF16PtrFromString(s)
    if err != nil {
        panic(err)
    }
    return uintptr(unsafe.Pointer(p))
}

func boolRet(r uintptr) bool { return r != 0 }

func RGB(r, g, b byte) uintptr {
    return uintptr(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

type POINT struct { X, Y int32 }
type SIZE struct { CX, CY int32 }
type RECT struct { Left, Top, Right, Bottom int32 }
```

---

## 1. GetDeviceCaps：获取屏幕 DPI / 分辨率信息

### 场景

获取当前屏幕 DC 的 DPI、颜色深度、屏幕宽高等信息。

### API

```text
int GetDeviceCaps(HDC hdc, int index);
```

### Go 案例

```go
var (
    procGetDC         = user32.NewProc("GetDC")
    procReleaseDC    = user32.NewProc("ReleaseDC")
    procGetDeviceCaps = gdi32.NewProc("GetDeviceCaps")
)

const (
    HORZRES    = 8
    VERTRES    = 10
    BITSPIXEL  = 12
    PLANES     = 14
    LOGPIXELSX = 88
    LOGPIXELSY = 90
)

func PrintScreenCaps() {
    hdc, _, _ := procGetDC.Call(0)
    defer procReleaseDC.Call(0, hdc)

    w, _, _ := procGetDeviceCaps.Call(hdc, HORZRES)
    h, _, _ := procGetDeviceCaps.Call(hdc, VERTRES)
    dpiX, _, _ := procGetDeviceCaps.Call(hdc, LOGPIXELSX)
    dpiY, _, _ := procGetDeviceCaps.Call(hdc, LOGPIXELSY)
    bits, _, _ := procGetDeviceCaps.Call(hdc, BITSPIXEL)
    planes, _, _ := procGetDeviceCaps.Call(hdc, PLANES)

    fmt.Println("screen:", w, h, "dpi:", dpiX, dpiY, "color bits:", bits*planes)
}
```

---

## 2. TextOutW：在屏幕上输出文字

### 场景

在指定 DC 的指定坐标输出一段 Unicode 文本。

### API

```text
BOOL TextOutW(HDC hdc, int x, int y, LPCWSTR lpString, int c);
```

### Go 案例

```go
var procTextOutW = gdi32.NewProc("TextOutW")

func DrawTextOut(hdc uintptr, x, y int32, text string) bool {
    p, _ := syscall.UTF16PtrFromString(text)
    ret, _, _ := procTextOutW.Call(
        hdc,
        uintptr(x), uintptr(y),
        uintptr(unsafe.Pointer(p)),
        uintptr(len([]rune(text))),
    )
    return ret != 0
}
```

---

## 3. SetTextColor / SetBkMode：设置文字颜色和透明背景

### 场景

绘制蓝色文字，并让文字背景透明。

### Go 案例

```go
var (
    procSetTextColor = gdi32.NewProc("SetTextColor")
    procSetBkMode    = gdi32.NewProc("SetBkMode")
)

const TRANSPARENT = 1

func DrawBlueTransparentText(hdc uintptr, text string) {
    procSetTextColor.Call(hdc, RGB(0, 120, 215))
    procSetBkMode.Call(hdc, TRANSPARENT)
    DrawTextOut(hdc, 80, 80, text)
}
```

---

## 4. CreatePen / MoveToEx / LineTo：画线

### 场景

创建红色画笔，在 DC 上画一条线。

### API

```text
HPEN CreatePen(int iStyle, int cWidth, COLORREF color);
BOOL MoveToEx(HDC hdc, int x, int y, LPPOINT lppt);
BOOL LineTo(HDC hdc, int x, int y);
```

### Go 案例

```go
var (
    procCreatePen    = gdi32.NewProc("CreatePen")
    procSelectObject = gdi32.NewProc("SelectObject")
    procDeleteObject = gdi32.NewProc("DeleteObject")
    procMoveToEx     = gdi32.NewProc("MoveToEx")
    procLineTo       = gdi32.NewProc("LineTo")
)

const PS_SOLID = 0

func DrawRedLine(hdc uintptr) {
    pen, _, _ := procCreatePen.Call(PS_SOLID, 3, RGB(255, 0, 0))
    old, _, _ := procSelectObject.Call(hdc, pen)

    procMoveToEx.Call(hdc, 50, 50, 0)
    procLineTo.Call(hdc, 300, 50)

    procSelectObject.Call(hdc, old)
    procDeleteObject.Call(pen)
}
```

---

## 5. CreateSolidBrush / Rectangle：画填充矩形

### 场景

用蓝色画刷填充矩形。

### Go 案例

```go
var (
    procCreateSolidBrush = gdi32.NewProc("CreateSolidBrush")
    procRectangle        = gdi32.NewProc("Rectangle")
)

func DrawFilledRectangle(hdc uintptr) {
    brush, _, _ := procCreateSolidBrush.Call(RGB(200, 230, 255))
    oldBrush, _, _ := procSelectObject.Call(hdc, brush)

    procRectangle.Call(hdc, 50, 80, 300, 200)

    procSelectObject.Call(hdc, oldBrush)
    procDeleteObject.Call(brush)
}
```

---

## 6. Ellipse / RoundRect：画椭圆和圆角矩形

```go
var (
    procEllipse   = gdi32.NewProc("Ellipse")
    procRoundRect = gdi32.NewProc("RoundRect")
)

func DrawShapes(hdc uintptr) {
    procEllipse.Call(hdc, 50, 50, 200, 160)
    procRoundRect.Call(hdc, 220, 50, 420, 160, 24, 24)
}
```

---

## 7. Polyline / Polygon：画折线和多边形

```go
var (
    procPolyline = gdi32.NewProc("Polyline")
    procPolygon  = gdi32.NewProc("Polygon")
)

func DrawPoly(hdc uintptr) {
    pts := []POINT{{50, 50}, {120, 80}, {180, 30}, {260, 140}}
    procPolyline.Call(hdc, uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))

    tri := []POINT{{80, 220}, {180, 220}, {130, 150}}
    procPolygon.Call(hdc, uintptr(unsafe.Pointer(&tri[0])), uintptr(len(tri)))
}
```

---

## 8. CreateFontW：创建字体并绘制文字

### 场景

使用微软雅黑 32 像素高度绘制文本。

```go
var procCreateFontW = gdi32.NewProc("CreateFontW")

func DrawWithFont(hdc uintptr, text string) {
    font, _, _ := procCreateFontW.Call(
        32, 0, 0, 0,
        400, // FW_NORMAL
        0, 0, 0,
        1,  // DEFAULT_CHARSET
        0, 0, 0, 0,
        utf16Ptr("Microsoft YaHei"),
    )
    old, _, _ := procSelectObject.Call(hdc, font)

    DrawTextOut(hdc, 50, 50, text)

    procSelectObject.Call(hdc, old)
    procDeleteObject.Call(font)
}
```

---

## 9. GetTextExtentPoint32W：计算文本尺寸

```go
var procGetTextExtentPoint32W = gdi32.NewProc("GetTextExtentPoint32W")

func TextSize(hdc uintptr, text string) SIZE {
    p, _ := syscall.UTF16PtrFromString(text)
    var size SIZE
    procGetTextExtentPoint32W.Call(
        hdc,
        uintptr(unsafe.Pointer(p)),
        uintptr(len([]rune(text))),
        uintptr(unsafe.Pointer(&size)),
    )
    return size
}
```

---

## 10. CreateCompatibleDC / CreateCompatibleBitmap / BitBlt：屏幕截图核心

### 场景

把屏幕左上角一块区域拷贝到内存位图。

```go
var (
    procCreateCompatibleDC     = gdi32.NewProc("CreateCompatibleDC")
    procCreateCompatibleBitmap = gdi32.NewProc("CreateCompatibleBitmap")
    procBitBlt                 = gdi32.NewProc("BitBlt")
    procDeleteDC               = gdi32.NewProc("DeleteDC")
)

const SRCCOPY = 0x00CC0020

func CaptureToBitmap(x, y, w, h int32) uintptr {
    screen, _, _ := procGetDC.Call(0)
    defer procReleaseDC.Call(0, screen)

    memdc, _, _ := procCreateCompatibleDC.Call(screen)
    defer procDeleteDC.Call(memdc)

    bmp, _, _ := procCreateCompatibleBitmap.Call(screen, uintptr(w), uintptr(h))
    old, _, _ := procSelectObject.Call(memdc, bmp)

    procBitBlt.Call(
        memdc, 0, 0, uintptr(w), uintptr(h),
        screen, uintptr(x), uintptr(y), SRCCOPY,
    )

    procSelectObject.Call(memdc, old)
    return bmp // 调用者负责 DeleteObject(bmp)
}
```

---

## 11. StretchBlt：缩放拷贝图像

```go
var procStretchBlt = gdi32.NewProc("StretchBlt")

func StretchCopy(dst, src uintptr, dstW, dstH, srcW, srcH int32) bool {
    ret, _, _ := procStretchBlt.Call(
        dst, 0, 0, uintptr(dstW), uintptr(dstH),
        src, 0, 0, uintptr(srcW), uintptr(srcH),
        SRCCOPY,
    )
    return ret != 0
}
```

---

## 12. GetPixel / SetPixel：读写单个像素

```go
var (
    procGetPixel = gdi32.NewProc("GetPixel")
    procSetPixel = gdi32.NewProc("SetPixel")
)

func PixelDemo(hdc uintptr) {
    color, _, _ := procGetPixel.Call(hdc, 10, 10)
    fmt.Printf("color = 0x%06X\n", color)

    procSetPixel.Call(hdc, 10, 10, RGB(255, 0, 0))
}
```

> 单像素 API 很直观，但大量像素处理时性能差。建议用 `CreateDIBSection` 直接处理内存。

---

## 13. CreateDIBSection：创建可直接访问像素内存的位图

```go
var procCreateDIBSection = gdi32.NewProc("CreateDIBSection")

type BITMAPINFOHEADER struct {
    BiSize          uint32
    BiWidth         int32
    BiHeight        int32
    BiPlanes        uint16
    BiBitCount      uint16
    BiCompression   uint32
    BiSizeImage     uint32
    BiXPelsPerMeter int32
    BiYPelsPerMeter int32
    BiClrUsed       uint32
    BiClrImportant  uint32
}

type RGBQUAD struct { Blue, Green, Red, Reserved byte }

type BITMAPINFO struct {
    Header BITMAPINFOHEADER
    Colors [1]RGBQUAD
}

const BI_RGB = 0
const DIB_RGB_COLORS = 0

func NewDIBSection(hdc uintptr, w, h int32) (bmp uintptr, bits uintptr) {
    bmi := BITMAPINFO{}
    bmi.Header.BiSize = uint32(unsafe.Sizeof(bmi.Header))
    bmi.Header.BiWidth = w
    bmi.Header.BiHeight = -h // 负数表示 top-down DIB
    bmi.Header.BiPlanes = 1
    bmi.Header.BiBitCount = 32
    bmi.Header.BiCompression = BI_RGB

    var ppvBits uintptr
    bmp, _, _ = procCreateDIBSection.Call(
        hdc,
        uintptr(unsafe.Pointer(&bmi)),
        DIB_RGB_COLORS,
        uintptr(unsafe.Pointer(&ppvBits)),
        0, 0,
    )
    return bmp, ppvBits
}
```

---

## 14. GetDIBits：从 HBITMAP 读取像素数据

```go
var procGetDIBits = gdi32.NewProc("GetDIBits")

func ReadBitmapPixels(hdc, bmp uintptr, w, h int32) []byte {
    bmi := BITMAPINFO{}
    bmi.Header.BiSize = uint32(unsafe.Sizeof(bmi.Header))
    bmi.Header.BiWidth = w
    bmi.Header.BiHeight = -h
    bmi.Header.BiPlanes = 1
    bmi.Header.BiBitCount = 32
    bmi.Header.BiCompression = BI_RGB

    buf := make([]byte, int(w*h*4))
    procGetDIBits.Call(
        hdc, bmp,
        0, uintptr(h),
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(unsafe.Pointer(&bmi)),
        DIB_RGB_COLORS,
    )
    return buf
}
```

---

## 15. PatBlt：快速填充 / 反色

```go
var procPatBlt = gdi32.NewProc("PatBlt")

const (
    PATCOPY = 0x00F00021
    DSTINVERT = 0x00550009
)

func InvertRect(hdc uintptr, x, y, w, h int32) {
    procPatBlt.Call(hdc, uintptr(x), uintptr(y), uintptr(w), uintptr(h), DSTINVERT)
}
```

---

## 16. CreateRectRgn / SelectClipRgn：裁剪绘制区域

```go
var (
    procCreateRectRgn = gdi32.NewProc("CreateRectRgn")
    procSelectClipRgn = gdi32.NewProc("SelectClipRgn")
)

func DrawOnlyInRect(hdc uintptr) {
    rgn, _, _ := procCreateRectRgn.Call(50, 50, 250, 150)
    procSelectClipRgn.Call(hdc, rgn)

    DrawBlueTransparentText(hdc, "只有裁剪区域里能看到")

    procSelectClipRgn.Call(hdc, 0)
    procDeleteObject.Call(rgn)
}
```

---

## 17. CombineRgn：组合区域

```go
var procCombineRgn = gdi32.NewProc("CombineRgn")

const RGN_OR = 2

func CombineRegionDemo() uintptr {
    r1, _, _ := procCreateRectRgn.Call(0, 0, 100, 100)
    r2, _, _ := procCreateRectRgn.Call(50, 50, 150, 150)
    dst, _, _ := procCreateRectRgn.Call(0, 0, 0, 0)

    procCombineRgn.Call(dst, r1, r2, RGN_OR)

    procDeleteObject.Call(r1)
    procDeleteObject.Call(r2)
    return dst // 调用者负责 DeleteObject(dst)
}
```

---

## 18. SetMapMode / LPtoDP / DPtoLP：逻辑坐标和设备坐标转换

```go
var (
    procSetMapMode = gdi32.NewProc("SetMapMode")
    procLPtoDP     = gdi32.NewProc("LPtoDP")
    procDPtoLP     = gdi32.NewProc("DPtoLP")
)

const MM_TEXT = 1

func ConvertPointDemo(hdc uintptr) {
    procSetMapMode.Call(hdc, MM_TEXT)
    pts := []POINT{{100, 100}}
    procLPtoDP.Call(hdc, uintptr(unsafe.Pointer(&pts[0])), 1)
    fmt.Println("device point:", pts[0])
}
```

---

## 19. BeginPath / StrokePath：路径绘制

```go
var (
    procBeginPath  = gdi32.NewProc("BeginPath")
    procEndPath    = gdi32.NewProc("EndPath")
    procStrokePath = gdi32.NewProc("StrokePath")
)

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
```

---

## 20. SaveDC / RestoreDC：保存和恢复 DC 状态

```go
var (
    procSaveDC    = gdi32.NewProc("SaveDC")
    procRestoreDC = gdi32.NewProc("RestoreDC")
)

func WithSavedDC(hdc uintptr) {
    state, _, _ := procSaveDC.Call(hdc)
    defer procRestoreDC.Call(hdc, state)

    procSetTextColor.Call(hdc, RGB(255, 0, 0))
    DrawTextOut(hdc, 20, 20, "临时修改 DC 状态")
}
```

---

## 21. StartDocW / StartPage / EndPage / EndDoc：打印基本流程

```go
var (
    procCreateDCW  = gdi32.NewProc("CreateDCW")
    procStartDocW  = gdi32.NewProc("StartDocW")
    procStartPage  = gdi32.NewProc("StartPage")
    procEndPage    = gdi32.NewProc("EndPage")
    procEndDoc     = gdi32.NewProc("EndDoc")
)

type DOCINFO struct {
    CbSize      int32
    LpszDocName uintptr
    LpszOutput  uintptr
    LpszDatatype uintptr
    FwType      uint32
}

func PrintText(printerName string) {
    hdc, _, _ := procCreateDCW.Call(utf16Ptr("WINSPOOL"), utf16Ptr(printerName), 0, 0)
    if hdc == 0 { return }
    defer procDeleteDC.Call(hdc)

    di := DOCINFO{CbSize: int32(unsafe.Sizeof(DOCINFO{})), LpszDocName: utf16Ptr("Go GDI Print")}
    procStartDocW.Call(hdc, uintptr(unsafe.Pointer(&di)))
    procStartPage.Call(hdc)
    DrawTextOut(hdc, 200, 200, "Hello Printer")
    procEndPage.Call(hdc)
    procEndDoc.Call(hdc)
}
```

---

## 22. DeleteObject / DeleteDC：资源释放总结

```go
func CleanupExample(hdc uintptr) {
    pen, _, _ := procCreatePen.Call(PS_SOLID, 1, RGB(0, 0, 0))
    old, _, _ := procSelectObject.Call(hdc, pen)

    // 使用 pen 绘制

    procSelectObject.Call(hdc, old)
    procDeleteObject.Call(pen)
}
```

释放原则：

| 创建 / 获取 | 释放 |
|---|---|
| `GetDC` | `ReleaseDC` |
| `CreateCompatibleDC` | `DeleteDC` |
| `CreateDCW` | `DeleteDC` |
| `CreatePen` | `DeleteObject` |
| `CreateSolidBrush` | `DeleteObject` |
| `CreateFontW` | `DeleteObject` |
| `CreateCompatibleBitmap` | `DeleteObject` |
| `CreateDIBSection` | `DeleteObject` |
| `CreateRectRgn` | `DeleteObject` |

---

## 23. 一个完整小程序：在屏幕上画文字和图形

```go
package main

import (
    "syscall"
    "time"
    "unsafe"
)

var (
    user32 = syscall.NewLazyDLL("user32.dll")
    gdi32  = syscall.NewLazyDLL("gdi32.dll")

    getDC         = user32.NewProc("GetDC")
    releaseDC    = user32.NewProc("ReleaseDC")
    textOutW      = gdi32.NewProc("TextOutW")
    createPen     = gdi32.NewProc("CreatePen")
    selectObject  = gdi32.NewProc("SelectObject")
    deleteObject  = gdi32.NewProc("DeleteObject")
    moveToEx      = gdi32.NewProc("MoveToEx")
    lineTo        = gdi32.NewProc("LineTo")
    rectangle     = gdi32.NewProc("Rectangle")
)

func rgb(r, g, b byte) uintptr {
    return uintptr(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

func utf16Ptr(s string) uintptr {
    p, _ := syscall.UTF16PtrFromString(s)
    return uintptr(unsafe.Pointer(p))
}

func main() {
    hdc, _, _ := getDC.Call(0)
    defer releaseDC.Call(0, hdc)

    pen, _, _ := createPen.Call(0, 4, rgb(255, 0, 0))
    old, _, _ := selectObject.Call(hdc, pen)

    moveToEx.Call(hdc, 100, 100, 0)
    lineTo.Call(hdc, 500, 100)
    rectangle.Call(hdc, 100, 140, 500, 300)

    msg := "Go 调用 gdi32.dll 绘制"
    textOutW.Call(hdc, 120, 180, utf16Ptr(msg), uintptr(len([]rune(msg))))

    selectObject.Call(hdc, old)
    deleteObject.Call(pen)

    time.Sleep(3 * time.Second)
}
```

> 这个程序只是演示。正式绘图不要长期直接画桌面，应该在窗口的 `WM_PAINT` 中绘制。
