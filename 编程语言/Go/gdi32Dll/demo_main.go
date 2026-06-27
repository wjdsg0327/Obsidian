//go:build windows

package main

import "time"

func main() {
    hdc, _, _ := procGetDC.Call(0)
    defer procReleaseDC.Call(0, hdc)

    DrawBasicShapes(hdc)
    DrawTextWithFont(hdc, "Go 调用 gdi32.dll 绘制")
    DrawPoly(hdc)

    time.Sleep(3 * time.Second)
}
