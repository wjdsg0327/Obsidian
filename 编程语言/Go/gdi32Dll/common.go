//go:build windows

package main

import (
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

func RGB(r, g, b byte) uintptr {
    return uintptr(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

type POINT struct { X, Y int32 }
type SIZE struct { CX, CY int32 }
type RECT struct { Left, Top, Right, Bottom int32 }
