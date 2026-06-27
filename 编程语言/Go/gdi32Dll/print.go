//go:build windows

package main

import "unsafe"

var (
    procCreateDCW = gdi32.NewProc("CreateDCW")
    procStartDocW = gdi32.NewProc("StartDocW")
    procStartPage = gdi32.NewProc("StartPage")
    procEndPage   = gdi32.NewProc("EndPage")
    procEndDoc    = gdi32.NewProc("EndDoc")
)

type DOCINFO struct {
    CbSize       int32
    LpszDocName  uintptr
    LpszOutput   uintptr
    LpszDatatype uintptr
    FwType       uint32
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
