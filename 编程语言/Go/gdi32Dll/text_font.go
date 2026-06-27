//go:build windows

package main

import (
    "syscall"
    "unsafe"
)

var (
    procTextOutW              = gdi32.NewProc("TextOutW")
    procSetTextColor          = gdi32.NewProc("SetTextColor")
    procSetBkMode             = gdi32.NewProc("SetBkMode")
    procCreateFontW           = gdi32.NewProc("CreateFontW")
    procGetTextExtentPoint32W = gdi32.NewProc("GetTextExtentPoint32W")
)

const TRANSPARENT = 1

func DrawTextOut(hdc uintptr, x, y int32, text string) bool {
    p, _ := syscall.UTF16PtrFromString(text)
    ret, _, _ := procTextOutW.Call(hdc, uintptr(x), uintptr(y), uintptr(unsafe.Pointer(p)), uintptr(len([]rune(text))))
    return ret != 0
}

func DrawTextWithFont(hdc uintptr, text string) {
    font, _, _ := procCreateFontW.Call(32, 0, 0, 0, 400, 0, 0, 0, 1, 0, 0, 0, 0, utf16Ptr("Microsoft YaHei"))
    old, _, _ := procSelectObject.Call(hdc, font)

    procSetTextColor.Call(hdc, RGB(0, 120, 215))
    procSetBkMode.Call(hdc, TRANSPARENT)
    DrawTextOut(hdc, 60, 60, text)

    procSelectObject.Call(hdc, old)
    procDeleteObject.Call(font)
}

func TextSize(hdc uintptr, text string) SIZE {
    p, _ := syscall.UTF16PtrFromString(text)
    var size SIZE
    procGetTextExtentPoint32W.Call(hdc, uintptr(unsafe.Pointer(p)), uintptr(len([]rune(text))), uintptr(unsafe.Pointer(&size)))
    return size
}
