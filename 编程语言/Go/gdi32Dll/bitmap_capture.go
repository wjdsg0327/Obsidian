//go:build windows

package main

import "unsafe"

var (
    procCreateCompatibleDC     = gdi32.NewProc("CreateCompatibleDC")
    procCreateCompatibleBitmap = gdi32.NewProc("CreateCompatibleBitmap")
    procBitBlt                 = gdi32.NewProc("BitBlt")
    procStretchBlt             = gdi32.NewProc("StretchBlt")
    procDeleteDC               = gdi32.NewProc("DeleteDC")
    procCreateDIBSection       = gdi32.NewProc("CreateDIBSection")
    procGetDIBits              = gdi32.NewProc("GetDIBits")
)

const (
    SRCCOPY        = 0x00CC0020
    BI_RGB         = 0
    DIB_RGB_COLORS = 0
)

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

func CaptureToBitmap(x, y, w, h int32) uintptr {
    screen, _, _ := procGetDC.Call(0)
    defer procReleaseDC.Call(0, screen)

    memdc, _, _ := procCreateCompatibleDC.Call(screen)
    defer procDeleteDC.Call(memdc)

    bmp, _, _ := procCreateCompatibleBitmap.Call(screen, uintptr(w), uintptr(h))
    old, _, _ := procSelectObject.Call(memdc, bmp)
    procBitBlt.Call(memdc, 0, 0, uintptr(w), uintptr(h), screen, uintptr(x), uintptr(y), SRCCOPY)
    procSelectObject.Call(memdc, old)
    return bmp
}

func NewDIBSection(hdc uintptr, w, h int32) (bmp uintptr, bits uintptr) {
    bmi := BITMAPINFO{}
    bmi.Header.BiSize = uint32(unsafe.Sizeof(bmi.Header))
    bmi.Header.BiWidth = w
    bmi.Header.BiHeight = -h
    bmi.Header.BiPlanes = 1
    bmi.Header.BiBitCount = 32
    bmi.Header.BiCompression = BI_RGB

    var ppvBits uintptr
    bmp, _, _ = procCreateDIBSection.Call(hdc, uintptr(unsafe.Pointer(&bmi)), DIB_RGB_COLORS, uintptr(unsafe.Pointer(&ppvBits)), 0, 0)
    return bmp, ppvBits
}

func ReadBitmapPixels(hdc, bmp uintptr, w, h int32) []byte {
    bmi := BITMAPINFO{}
    bmi.Header.BiSize = uint32(unsafe.Sizeof(bmi.Header))
    bmi.Header.BiWidth = w
    bmi.Header.BiHeight = -h
    bmi.Header.BiPlanes = 1
    bmi.Header.BiBitCount = 32
    bmi.Header.BiCompression = BI_RGB

    buf := make([]byte, int(w*h*4))
    procGetDIBits.Call(hdc, bmp, 0, uintptr(h), uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&bmi)), DIB_RGB_COLORS)
    return buf
}
