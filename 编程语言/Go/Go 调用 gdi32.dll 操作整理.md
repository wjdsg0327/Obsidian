---
title: Go 调用 gdi32.dll 操作整理
date: 2026-06-27
tags:
  - Go
  - Windows
  - Win32
  - gdi32.dll
  - Windows-GDI
  - 图形编程
---

# Go 调用 gdi32.dll 操作整理

> 目标：整理 Go 语言中通过 Win32 API 调用 `gdi32.dll` 能做的主要操作、常用函数、调用方式、注意事项与示例代码。  
> 范围：`gdi32.dll` 主要负责 Windows GDI（Graphics Device Interface，图形设备接口），包括设备上下文 DC、画笔/画刷/字体/位图、文本绘制、图形绘制、屏幕截图、打印、区域、路径、坐标变换、调色板等。  
> 说明：`gdi32.dll` 的 API 很多，完整清单以 Microsoft Learn 的 `wingdi.h` / Win32 GDI 文档为准。本文按 Go 调用时最常用的功能域整理。

---

## 1. gdi32.dll 是什么

`gdi32.dll` 是 Windows 图形设备接口 GDI 的核心 DLL。

它的职责是让应用程序通过统一 API 在不同设备上绘制图形和文字，例如：

- 屏幕窗口
- 内存位图
- 打印机
- 图元文件 EMF
- 兼容设备上下文

Windows 程序通常不会直接访问显卡或打印机硬件，而是通过 GDI 调用设备驱动完成绘制。

### 1.1 和 user32.dll 的关系

可以这样理解：

| DLL | 主要负责 | 典型能力 |
|---|---|---|
| `user32.dll` | 窗口、消息、输入、菜单、剪贴板、光标 | 找窗口、发消息、模拟输入、管理窗口 |
| `gdi32.dll` | 图形绘制、文本输出、位图、字体、打印、DC | 画线、画矩形、写字、截图、创建字体、位图拷贝 |

常见组合：

1. 用 `user32.dll` 的 `GetDC` / `ReleaseDC` 获取窗口或屏幕 DC。
2. 用 `gdi32.dll` 在这个 DC 上绘制。
3. 或用 `gdi32.dll` 创建内存 DC / 位图，再用 `BitBlt` 截图或拷贝图像。

---

## 2. Go 调用 gdi32.dll 的基本方式

Go 调用 Windows DLL 常见方式：

1. 标准库 `syscall`
2. `golang.org/x/sys/windows`
3. 第三方封装，例如 `github.com/lxn/win`

简单学习和实验时，可以直接使用：

```go
syscall.NewLazyDLL("gdi32.dll")
```

### 2.1 基础模板：加载 gdi32.dll

```go
package main

import (
    "syscall"
    "unsafe"
)

var (
    gdi32       = syscall.NewLazyDLL("gdi32.dll")
    user32      = syscall.NewLazyDLL("user32.dll")
    procTextOut = gdi32.NewProc("TextOutW")
)

func utf16Ptr(s string) *uint16 {
    p, _ := syscall.UTF16PtrFromString(s)
    return p
}

func main() {
    hwnd := uintptr(0) // 0 表示整个屏幕 DC
    getDC := user32.NewProc("GetDC")
    releaseDC := user32.NewProc("ReleaseDC")

    hdc, _, _ := getDC.Call(hwnd)
    defer releaseDC.Call(hwnd, hdc)

    text := "你好，gdi32.dll"
    p := utf16Ptr(text)
    procTextOut.Call(
        hdc,
        100,
        100,
        uintptr(unsafe.Pointer(p)),
        uintptr(len([]rune(text))),
    )
}
```

> 注意：直接在屏幕 DC 上画东西只是临时显示，窗口刷新或桌面重绘后可能消失。正式 GUI 程序一般在 `WM_PAINT` 中绘制。

### 2.2 A / W 函数区别

GDI 里很多文本相关 API 有两个版本：

- `xxxA`：ANSI 版本
- `xxxW`：Unicode / UTF-16 版本

Go 字符串是 UTF-8，调用 Windows API 通常转成 UTF-16，并优先调用 `xxxW`。

常见：

```text
TextOutW
DrawTextW
CreateFontW
GetTextExtentPoint32W
AddFontResourceW
```

### 2.3 常见 Win32 类型在 Go 里的映射

| Win32 类型 | Go 常见表示 |
|---|---|
| `HDC` | `uintptr` / `windows.Handle` |
| `HBITMAP` | `uintptr` / `windows.Handle` |
| `HBRUSH` | `uintptr` / `windows.Handle` |
| `HPEN` | `uintptr` / `windows.Handle` |
| `HFONT` | `uintptr` / `windows.Handle` |
| `HGDIOBJ` | `uintptr` / `windows.Handle` |
| `COLORREF` | `uint32` / `uintptr` |
| `BOOL` | 返回值 `0=false`，非 `0=true` |
| `RECT` | Go struct，字段通常是 `int32` |
| `POINT` | Go struct，字段通常是 `int32` |
| `SIZE` | Go struct，字段通常是 `int32` |
| `LPCWSTR` | `*uint16` / `uintptr(unsafe.Pointer(ptr))` |

---

## 3. gdi32.dll 能做什么：功能总览

`gdi32.dll` 覆盖的操作大致包括：

1. 设备上下文 DC 创建、获取、释放、保存与恢复
2. 画笔 Pen、画刷 Brush、字体 Font、位图 Bitmap 等 GDI 对象管理
3. 线条、矩形、圆角矩形、椭圆、弧线、多边形绘制
4. 文本绘制、字体选择、文本尺寸计算
5. 位图创建、屏幕截图、图像拷贝、缩放拷贝、透明/Alpha 混合
6. 填充、反色、区域裁剪、复杂区域合并
7. 坐标系、映射模式、视口/窗口原点与缩放
8. 路径 Path 绘制、填充、描边、裁剪
9. 打印机 DC、打印任务、页面输出
10. 图元文件 EMF 创建、回放和枚举
11. 调色板、颜色管理、ICC / WCS 相关操作
12. 字体资源加载、枚举字体、字形轮廓
13. OpenGL 上下文像素格式辅助函数

---

## 4. 设备上下文 DC

设备上下文（Device Context，DC）是 GDI 绘图的核心。几乎所有 GDI 绘制 API 都需要 `HDC`。

### 常用函数

| 函数 | 作用 |
|---|---|
| `CreateCompatibleDC` | 创建与指定 DC 兼容的内存 DC |
| `DeleteDC` | 删除由 `CreateCompatibleDC` / `CreateDC` 创建的 DC |
| `SaveDC` | 保存 DC 当前状态 |
| `RestoreDC` | 恢复 DC 状态 |
| `GetDeviceCaps` | 获取设备能力，例如分辨率、DPI、颜色深度 |
| `CreateDCW` | 创建指定设备 DC，例如打印机 |
| `ResetDCW` | 重置打印机或设备 DC |

> `GetDC` / `ReleaseDC` 在 `user32.dll` 里，不在 `gdi32.dll` 里；但它们经常和 GDI 配合使用。

### Go 示例：获取屏幕 DPI

```go
var (
    user32GetDC      = user32.NewProc("GetDC")
    user32ReleaseDC = user32.NewProc("ReleaseDC")
    getDeviceCaps   = gdi32.NewProc("GetDeviceCaps")
)

const (
    LOGPIXELSX = 88
    LOGPIXELSY = 90
)

func ScreenDPI() (int, int) {
    hdc, _, _ := user32GetDC.Call(0)
    defer user32ReleaseDC.Call(0, hdc)

    x, _, _ := getDeviceCaps.Call(hdc, LOGPIXELSX)
    y, _, _ := getDeviceCaps.Call(hdc, LOGPIXELSY)
    return int(x), int(y)
}
```

---

## 5. GDI 对象：画笔、画刷、字体、位图

GDI 通过对象控制绘制效果。

| 对象 | 用途 | 常用函数 |
|---|---|---|
| `HPEN` | 线条颜色、宽度、样式 | `CreatePen`, `ExtCreatePen` |
| `HBRUSH` | 填充颜色、图案 | `CreateSolidBrush`, `CreateHatchBrush`, `GetStockObject` |
| `HFONT` | 文本字体 | `CreateFontW`, `CreateFontIndirectW` |
| `HBITMAP` | 位图图像 | `CreateCompatibleBitmap`, `CreateDIBSection` |
| `HRGN` | 区域 | `CreateRectRgn`, `CreateEllipticRgn` |
| `HPALETTE` | 调色板 | `CreatePalette`, `SelectPalette` |

### 重要规则：SelectObject 与 DeleteObject

GDI 对象使用流程通常是：

1. 创建对象，例如 `CreatePen`
2. 用 `SelectObject(hdc, obj)` 选入 DC
3. 绘制
4. 把旧对象选回去
5. 用 `DeleteObject(obj)` 删除自己创建的对象

```go
old, _, _ := selectObject.Call(hdc, pen)
// draw...
selectObject.Call(hdc, old)
deleteObject.Call(pen)
```

> 不要删除系统库存对象（`GetStockObject` 返回的对象），也不要删除仍被 DC 选中的对象。

---

## 6. 绘制基础图形

### 常用函数

| 函数 | 作用 |
|---|---|
| `MoveToEx` | 移动当前绘图点 |
| `LineTo` | 从当前位置画线 |
| `Rectangle` | 画矩形 |
| `RoundRect` | 画圆角矩形 |
| `Ellipse` | 画椭圆 |
| `Arc` | 画椭圆弧 |
| `Pie` | 画扇形 |
| `Chord` | 画弦形 |
| `Polyline` | 画折线 |
| `Polygon` | 画多边形 |
| `PolyBezier` | 画贝塞尔曲线 |
| `PatBlt` | 用当前画刷填充或进行光栅操作 |

### 示例：画线和矩形

```go
var (
    createPen    = gdi32.NewProc("CreatePen")
    selectObject = gdi32.NewProc("SelectObject")
    deleteObject = gdi32.NewProc("DeleteObject")
    moveToEx     = gdi32.NewProc("MoveToEx")
    lineTo       = gdi32.NewProc("LineTo")
    rectangle    = gdi32.NewProc("Rectangle")
)

const PS_SOLID = 0

func RGB(r, g, b byte) uintptr {
    return uintptr(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}

func DrawDemo(hdc uintptr) {
    pen, _, _ := createPen.Call(PS_SOLID, 3, RGB(255, 0, 0))
    old, _, _ := selectObject.Call(hdc, pen)

    moveToEx.Call(hdc, 50, 50, 0)
    lineTo.Call(hdc, 250, 50)
    rectangle.Call(hdc, 50, 80, 250, 180)

    selectObject.Call(hdc, old)
    deleteObject.Call(pen)
}
```

---

## 7. 文本绘制与字体

### 常用函数

| 函数 | 作用 |
|---|---|
| `TextOutW` | 在指定位置输出文本 |
| `ExtTextOutW` | 更复杂的文本输出，可指定裁剪矩形、字符间距等 |
| `DrawTextW` | 在矩形区域中格式化绘制文本 |
| `SetTextColor` | 设置文本颜色 |
| `SetBkColor` | 设置文本背景颜色 |
| `SetBkMode` | 设置背景模式，透明或不透明 |
| `CreateFontW` | 创建字体 |
| `CreateFontIndirectW` | 根据 `LOGFONT` 创建字体 |
| `GetTextExtentPoint32W` | 计算文本像素尺寸 |
| `EnumFontFamiliesExW` | 枚举字体 |
| `GetGlyphOutlineW` | 获取字形轮廓 |

### 示例：透明背景文字

```go
const TRANSPARENT = 1

var (
    setTextColor = gdi32.NewProc("SetTextColor")
    setBkMode    = gdi32.NewProc("SetBkMode")
    textOutW     = gdi32.NewProc("TextOutW")
)

func DrawTextDemo(hdc uintptr, text string) {
    p, _ := syscall.UTF16PtrFromString(text)
    setTextColor.Call(hdc, RGB(0, 120, 215))
    setBkMode.Call(hdc, TRANSPARENT)
    textOutW.Call(hdc, 40, 40, uintptr(unsafe.Pointer(p)), uintptr(len([]rune(text))))
}
```

---

## 8. 位图、截图与图像拷贝

这是 Go 调用 `gdi32.dll` 最常见的实用场景之一。

### 常用函数

| 函数 | 作用 |
|---|---|
| `CreateCompatibleDC` | 创建内存 DC |
| `CreateCompatibleBitmap` | 创建兼容位图 |
| `CreateDIBSection` | 创建可直接访问像素内存的 DIB 位图 |
| `SelectObject` | 把位图选入内存 DC |
| `BitBlt` | 位块传输，常用于截图 |
| `StretchBlt` | 缩放拷贝 |
| `GetDIBits` | 从位图取像素数据 |
| `SetDIBits` | 把像素数据写入位图 |
| `GetPixel` | 获取某点颜色 |
| `SetPixel` | 设置某点颜色 |
| `AlphaBlend` | 半透明混合 |
| `TransparentBlt` | 透明色拷贝 |
| `PlgBlt` | 平行四边形位图拷贝 |

### 示例：屏幕截图基本流程

1. `user32.GetDC(0)` 获取屏幕 DC
2. `CreateCompatibleDC(screenDC)` 创建内存 DC
3. `CreateCompatibleBitmap(screenDC, width, height)` 创建位图
4. `SelectObject(memoryDC, bitmap)`
5. `BitBlt(memoryDC, 0, 0, width, height, screenDC, x, y, SRCCOPY)`
6. `GetDIBits` 读取像素，或后续保存 BMP/PNG
7. 选回旧对象，释放 DC / 删除对象

---

## 9. 区域、裁剪与填充

### 常用函数

| 函数 | 作用 |
|---|---|
| `CreateRectRgn` | 创建矩形区域 |
| `CreateEllipticRgn` | 创建椭圆区域 |
| `CreatePolygonRgn` | 创建多边形区域 |
| `CombineRgn` | 合并区域 |
| `SelectClipRgn` | 设置 DC 裁剪区域 |
| `ExcludeClipRect` | 从裁剪区域排除矩形 |
| `IntersectClipRect` | 和矩形相交作为裁剪区 |
| `FillRgn` | 填充区域 |
| `FrameRgn` | 绘制区域边框 |
| `InvertRgn` | 反转区域颜色 |
| `PtInRegion` | 判断点是否在区域内 |
| `RectInRegion` | 判断矩形是否与区域相交 |

---

## 10. 坐标系与映射模式

GDI 支持逻辑坐标到设备坐标的映射。

### 常用函数

| 函数 | 作用 |
|---|---|
| `SetMapMode` | 设置映射模式 |
| `GetMapMode` | 获取映射模式 |
| `SetViewportOrgEx` | 设置视口原点 |
| `SetWindowOrgEx` | 设置窗口原点 |
| `SetViewportExtEx` | 设置视口范围 |
| `SetWindowExtEx` | 设置窗口范围 |
| `LPtoDP` | 逻辑坐标转设备坐标 |
| `DPtoLP` | 设备坐标转逻辑坐标 |
| `SetWorldTransform` | 设置世界变换 |
| `ModifyWorldTransform` | 修改世界变换 |

常见映射模式：

| 模式 | 含义 |
|---|---|
| `MM_TEXT` | 默认模式，单位是像素，x 向右，y 向下 |
| `MM_LOMETRIC` | 0.1 mm，x 向右，y 向上 |
| `MM_HIMETRIC` | 0.01 mm，x 向右，y 向上 |
| `MM_TWIPS` | 1/1440 inch，常用于打印 |
| `MM_ISOTROPIC` | 自定义比例，x/y 等比例 |
| `MM_ANISOTROPIC` | 自定义比例，x/y 可不等比例 |

---

## 11. 路径 Path

Path 可以记录一组线条、曲线和图形，然后统一填充、描边或作为裁剪区域。

### 常用函数

| 函数 | 作用 |
|---|---|
| `BeginPath` | 开始记录路径 |
| `EndPath` | 结束记录路径 |
| `AbortPath` | 放弃当前路径 |
| `CloseFigure` | 闭合当前图形 |
| `StrokePath` | 描边路径 |
| `FillPath` | 填充路径 |
| `StrokeAndFillPath` | 描边并填充路径 |
| `SelectClipPath` | 把路径设为裁剪区域 |
| `FlattenPath` | 把曲线近似为线段 |
| `WidenPath` | 按当前画笔加宽路径 |

---

## 12. 打印相关

GDI 仍然是 Windows 打印体系中的重要基础。

### 常用函数

| 函数 | 作用 |
|---|---|
| `CreateDCW` | 创建打印机 DC |
| `StartDocW` | 开始打印任务 |
| `StartPage` | 开始一页 |
| `EndPage` | 结束一页 |
| `EndDoc` | 结束打印任务 |
| `AbortDoc` | 中止打印任务 |
| `SetAbortProc` | 设置取消打印回调 |

基本流程：

```text
CreateDCW -> StartDocW -> StartPage -> 绘制 -> EndPage -> EndDoc -> DeleteDC
```

---

## 13. 图元文件 EMF

EMF（Enhanced Metafile）可以把 GDI 绘制命令记录成可回放文件。

### 常用函数

| 函数 | 作用 |
|---|---|
| `CreateEnhMetaFileW` | 创建增强图元文件 DC |
| `CloseEnhMetaFile` | 关闭并得到 EMF 句柄 |
| `PlayEnhMetaFile` | 回放 EMF |
| `DeleteEnhMetaFile` | 删除 EMF 句柄 |
| `EnumEnhMetaFile` | 枚举 EMF 记录 |
| `GetEnhMetaFileW` | 从文件加载 EMF |

---

## 14. 字体资源与字体枚举

### 常用函数

| 函数 | 作用 |
|---|---|
| `AddFontResourceW` | 添加字体资源 |
| `AddFontResourceExW` | 添加字体资源，可设置私有字体 |
| `RemoveFontResourceW` | 移除字体资源 |
| `CreateFontW` | 创建字体对象 |
| `EnumFontFamiliesExW` | 枚举字体 |
| `GetOutlineTextMetricsW` | 获取轮廓字体信息 |
| `GetGlyphOutlineW` | 获取字形位图或轮廓 |

---

## 15. 颜色、调色板与颜色管理

### 常用函数

| 函数 | 作用 |
|---|---|
| `RGB` 宏 | 组合 `COLORREF`，Go 中需要自己实现 |
| `GetRValue` / `GetGValue` / `GetBValue` | C 宏，Go 中自己拆位 |
| `CreatePalette` | 创建逻辑调色板 |
| `SelectPalette` | 选择调色板 |
| `RealizePalette` | 映射逻辑调色板到系统调色板 |
| `GetNearestColor` | 获取设备可显示的最近颜色 |
| `SetICMMode` | 设置图像颜色管理模式 |
| `GetICMProfileW` | 获取颜色配置文件 |
| `SetICMProfileW` | 设置颜色配置文件 |

Go 中常用 `COLORREF`：

```go
func RGB(r, g, b byte) uintptr {
    return uintptr(uint32(r) | uint32(g)<<8 | uint32(b)<<16)
}
```

---

## 16. 常见坑与注意事项

### 16.1 GDI 对象泄漏

每次 `CreatePen`、`CreateSolidBrush`、`CreateFontW`、`CreateCompatibleBitmap`、`CreateDIBSection` 后，都要考虑是否需要 `DeleteObject`。

每次 `CreateCompatibleDC`、`CreateDCW` 后，都要考虑 `DeleteDC`。

每次 `user32.GetDC` 后，都要对应 `user32.ReleaseDC`。

### 16.2 一定要选回旧对象

`SelectObject` 返回旧对象。结束后应该把旧对象选回 DC，再删除自己创建的新对象。

错误示例：

```go
selectObject.Call(hdc, pen)
deleteObject.Call(pen) // pen 还在 DC 里，风险高
```

正确示例：

```go
old, _, _ := selectObject.Call(hdc, pen)
// draw...
selectObject.Call(hdc, old)
deleteObject.Call(pen)
```

### 16.3 注意 DC 来源

| DC 来源 | 释放方式 |
|---|---|
| `user32.GetDC` | `user32.ReleaseDC` |
| `user32.BeginPaint` | `user32.EndPaint` |
| `gdi32.CreateCompatibleDC` | `gdi32.DeleteDC` |
| `gdi32.CreateDCW` | `gdi32.DeleteDC` |

不要混用释放函数。

### 16.4 32 位 / 64 位

Go 的 `uintptr` 会随平台变化，适合存放句柄和指针。结构体字段需要按 Win32 定义对齐，常见坐标字段是 `LONG`，Go 中用 `int32`。

### 16.5 屏幕缩放和 DPI

截图、窗口绘制、坐标计算时要注意 DPI 缩放。可以配合：

- `GetDeviceCaps(LOGPIXELSX / LOGPIXELSY)`
- `user32.SetProcessDPIAware`
- `user32.SetProcessDpiAwarenessContext`

### 16.6 GDI 与 Direct2D / DirectWrite

GDI 是经典 Win32 图形 API，兼容性强，适合学习和简单工具。现代高质量图形、抗锯齿文本、高性能渲染通常会优先考虑：

- Direct2D
- DirectWrite
- Direct3D
- Windows Composition

---

## 17. 常用 API 速查

| 功能 | 常用 API |
|---|---|
| 获取设备能力 | `GetDeviceCaps` |
| 内存 DC | `CreateCompatibleDC`, `DeleteDC` |
| 兼容位图 | `CreateCompatibleBitmap` |
| DIB 位图 | `CreateDIBSection`, `GetDIBits`, `SetDIBits` |
| 位图拷贝 | `BitBlt`, `StretchBlt`, `PatBlt` |
| 透明混合 | `AlphaBlend`, `TransparentBlt` |
| 像素 | `GetPixel`, `SetPixel`, `SetPixelV` |
| 画线 | `MoveToEx`, `LineTo`, `Polyline` |
| 形状 | `Rectangle`, `RoundRect`, `Ellipse`, `Arc`, `Pie`, `Chord`, `Polygon` |
| 文本 | `TextOutW`, `ExtTextOutW`, `DrawTextW`, `GetTextExtentPoint32W` |
| 字体 | `CreateFontW`, `CreateFontIndirectW`, `EnumFontFamiliesExW` |
| 画笔/画刷 | `CreatePen`, `ExtCreatePen`, `CreateSolidBrush`, `GetStockObject` |
| 对象管理 | `SelectObject`, `DeleteObject`, `GetObjectW` |
| 裁剪区域 | `CreateRectRgn`, `CombineRgn`, `SelectClipRgn` |
| 坐标转换 | `SetMapMode`, `LPtoDP`, `DPtoLP` |
| 路径 | `BeginPath`, `EndPath`, `StrokePath`, `FillPath` |
| 打印 | `StartDocW`, `StartPage`, `EndPage`, `EndDoc`, `AbortDoc` |
| EMF | `CreateEnhMetaFileW`, `PlayEnhMetaFile`, `CloseEnhMetaFile` |

---

## 18. 学习路线建议

建议按这个顺序学：

1. `HDC` 是什么，如何获取和释放
2. `SelectObject` / `DeleteObject` 的生命周期
3. `TextOutW`、`LineTo`、`Rectangle` 等基础绘制
4. `CreateFontW`、`SetTextColor`、`SetBkMode` 文本绘制
5. `CreateCompatibleDC` + `CreateCompatibleBitmap` + `BitBlt` 截图
6. `GetDIBits` / `CreateDIBSection` 处理像素
7. 区域、裁剪、路径、打印等进阶内容

---

## 19. 参考资料

- Microsoft Learn: Windows GDI
- Microsoft Learn: `wingdi.h` header
- Microsoft Learn: `BitBlt`, `CreateCompatibleDC`, `CreateCompatibleBitmap`, `TextOutW`, `CreateFontW`, `GetDIBits`
- Windows SDK: `wingdi.h`
