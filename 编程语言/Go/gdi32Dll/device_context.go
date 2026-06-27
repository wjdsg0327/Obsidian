//go:build windows

package main

import "fmt"

var (
    procGetDC          = user32.NewProc("GetDC")
    procReleaseDC     = user32.NewProc("ReleaseDC")
    procGetDeviceCaps = gdi32.NewProc("GetDeviceCaps")
    procSaveDC        = gdi32.NewProc("SaveDC")
    procRestoreDC     = gdi32.NewProc("RestoreDC")
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

func WithSavedDC(hdc uintptr, fn func()) {
    state, _, _ := procSaveDC.Call(hdc)
    defer procRestoreDC.Call(hdc, state)
    fn()
}
