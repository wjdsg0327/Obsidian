1、下载插件
```
https://github.com/jgm/pandoc/releases/tag/3.9.0.2
```
```
https://wkhtmltopdf.org/downloads.html
```
2、代码
```go
package main  
  
import (  
    "fmt"  
    "os"    "os/exec"    "path/filepath"    "strings")  
  
func main() {  
    inputDir := "./input"  
    outputDir := "./output"  
    cssPath := "./github.css" // 确保该目录下有这个 CSS 文件  
  
    // 自动创建输出目录  
    os.MkdirAll(outputDir, os.ModePerm)  
  
    files, _ := filepath.Glob(filepath.Join(inputDir, "*.md"))  
  
    for _, file := range files {  
       fileName := filepath.Base(file)  
       outputFile := filepath.Join(outputDir, strings.TrimSuffix(fileName, ".md")+".pdf")  
  
       // --- 核心配置 ---       args := []string{  
          file,  
          "-o", outputFile,  
          "--pdf-engine=wkhtmltopdf", // 更换为更稳定的引擎  
          "--css", cssPath,           // 加载 Typora 主题  
          "--self-contained",               // 将图片等资源嵌入  
          "-V", "mainfont=Microsoft YaHei", // 强制中文字体防止乱码  
       }  
  
       cmd := exec.Command("pandoc", args...)  
  
       // 重点：捕获标准错误输出，这样你就能看到详细的报错信息  
       stderr, err := cmd.CombinedOutput()  
  
       if err != nil {  
          fmt.Printf("❌ 转换 %s 失败\n", fileName)  
          fmt.Printf("错误信息: %s\n", string(stderr))  
       } else {  
          fmt.Printf("✅ 成功生成: %s\n", outputFile)  
       }  
    }  
}
```