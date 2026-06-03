https://github.com/carlos7ags/folio

基于你提供的 `carlos7ags/folio` 这个 Go 库，下面补充它的**使用方法**和**案例**。

## 📦 安装

```bash
go get github.com/carlos7ags/folio
```

## 🚀 基本用法

`folio` 是一个轻量级的 PDF 生成库，主要特点是用**链式调用**构建 PDF 内容。

### 1. 创建简单 PDF

```go
package main

import (
    "os"
    "github.com/carlos7ags/folio"
)

func main() {
    // 创建文档
    doc := folio.New("My Document")
    
    // 添加页面
    page := doc.AddPage()
    
    // 添加文字段落（自动换行）
    page.AddText("Hello, World! This is a sample PDF generated with Folio.")
    
    // 添加标题
    page.AddTitle("Main Title", 1)
    
    // 保存文件
    doc.SaveToFile("output.pdf")
}
```

### 2. 添加图片

```go
// 从文件加载图片
img, _ := folio.LoadImage("image.png")
page.AddImage(img, nil) // nil表示使用原尺寸

// 指定尺寸添加图片（宽500，高自动）
page.AddImage(img, &folio.ImageOption{Width: 500})
```

### 3. 表格生成

```go
// 创建表格
table := folio.NewTable(3) // 3列
table.AddRow([]string{"Name", "Age", "Country"})
table.AddRow([]string{"Alice", "30", "USA"})
table.AddRow([]string{"Bob", "25", "Canada"})

page.AddTable(table)
```

### 4. 自定义样式

```go
// 设置字体大小和样式
page.SetFontSize(14)
page.SetBold(true)

// 添加行间距
page.SetLineSpacing(1.5)

// 设置边距（单位：磅）
page.SetMargins(72, 72, 72, 72) // 上下左右各1英寸
```

## 💡 完整示例

```go
package main

import (
    "github.com/carlos7ags/folio"
)

func main() {
    doc := folio.New("Report " + time.Now().Format("2006-01-02"))
    page := doc.AddPage()
    
    // 标题区域
    page.SetBold(true)
    page.SetFontSize(24)
    page.AddText("Monthly Sales Report")
    page.SetBold(false)
    
    // 分隔线
    page.AddHorizontalLine(2)
    page.AddSpace(20)
    
    // 普通文本
    page.SetFontSize(12)
    page.AddText("This report shows the sales performance for January 2024.")
    
    // 表格数据
    salesTable := folio.NewTable(3)
    salesTable.SetHeaderStyle(folio.Style{Bold: true, BackgroundColor: "#DDDDDD"})
    salesTable.AddRow([]string{"Product", "Units Sold", "Revenue ($)"})
    salesTable.AddRow([]string{"Laptop", "150", "225,000"})
    salesTable.AddRow([]string{"Mouse", "500", "12,500"})
    salesTable.AddRow([]string{"Keyboard", "300", "45,000"})
    page.AddTable(salesTable)
    
    // 保存
    doc.SaveToFile("sales_report.pdf")
}
```

## 📌 注意事项

- 该库支持 **UTF-8** 中文（需系统有中文字体支持）
- 默认单位是**磅**（1英寸 = 72磅）
- 自动分页：内容超出页面高度会自动换页
- 需要**手动导入**图片库：`image/png`, `image/jpeg` 等

## 🔗 延伸

如果需要更复杂的 PDF 操作（如表单、水印、加密），可以考虑配合 `gofpdf` 或 `unidoc` 使用，但 `folio` 适合**快速生成结构化文档**的场景。
