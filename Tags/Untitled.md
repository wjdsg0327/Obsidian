
```go
package main  
  
import (  
    "fmt"  
    "io"    "log"    "net/http"    "net/url"    "os"    "path/filepath"    "strings"    "time"  
    "github.com/playwright-community/playwright-go"    "github.com/xuri/excelize/v2"    "gopkg.in/yaml.v3")  
  
// Config 对应 YAML 结构  
type Config struct {  
    URL              string `yaml:"url"`  
    TargetCount      int    `yaml:"target_count"`  
    DownloadInterval int    `yaml:"download_interval"`  
    Proxy            struct {  
       Enabled  bool   `yaml:"enabled"`  
       Server   string `yaml:"server"`  
       Username string `yaml:"username"`  
       Password string `yaml:"password"`  
    } `yaml:"proxy"`  
}  
  
type Item struct {  
    ID        string  
    Title     string  
    Price     string  
    RemoteURL string  
    LocalPath string  
}  
  
func main() {  
    err := playwright.Install()  
    if err != nil {  
       log.Fatalf("安装 Playwright 依赖失败: %v", err)  
    }  
  
    // 1. 加载配置  
    cfg := loadConfig("config.yml")  
  
    // 2. 初始化 Playwright    pw, _ := playwright.Run()  
    defer pw.Stop()  
  
    // 配置浏览器代理  
    launchOptions := playwright.BrowserTypeLaunchOptions{  
       Headless: playwright.Bool(false),  
    }  
    if cfg.Proxy.Enabled {  
       launchOptions.Proxy = &playwright.Proxy{  
          Server:   cfg.Proxy.Server,  
          Username: &cfg.Proxy.Username,  
          Password: &cfg.Proxy.Password,  
       }  
    }  
  
    browser, _ := pw.Chromium.Launch(launchOptions)  
    defer browser.Close()  
    page, _ := browser.NewPage()  
  
    imgDir := "./images"  
    os.MkdirAll(imgDir, os.ModePerm)  
  
    // 3. 访问页面  
    log.Printf("正在访问: %s", cfg.URL)  
    if _, err := page.Goto(cfg.URL, playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateNetworkidle}); err != nil {  
       log.Fatal(err)  
    }  
  
    // 关闭弹窗  
    closeBtn := "div.closeIconBg--cubvOqVh > img"  
    if _, err := page.WaitForSelector(closeBtn, playwright.PageWaitForSelectorOptions{Timeout: playwright.Float(3000)}); err == nil {  
       page.Click(closeBtn)  
       time.Sleep(1 * time.Second)  
    }  
  
    allItems := make(map[string]Item)  
  
    // 4. 抓取主循环  
    for len(allItems) < cfg.TargetCount {  
       entries, _ := page.QuerySelectorAll(".cardWarp--dZodM57A")  
  
       for _, entry := range entries {  
          linkEl, _ := entry.QuerySelector("a.feeds-item-wrap--rGdH_KoF")  
          if linkEl == nil {  
             continue  
          }  
          href, _ := linkEl.GetAttribute("href")  
          id := extractID(href)  
  
          if _, exists := allItems[id]; !exists {  
             entry.ScrollIntoViewIfNeeded()  
             time.Sleep(500 * time.Millisecond)  
  
             titleEl, _ := entry.QuerySelector(".main-title--sMrtWSJa")  
             title, _ := titleEl.InnerText()  
             priceEl, _ := entry.QuerySelector(".price-wrap--YzmU5cUl")  
             price, _ := priceEl.InnerText()  
  
             var cleanImgURL string  
             imgEl, _ := entry.QuerySelector(".feeds-image--TDRC4fV1")  
             if imgEl != nil {  
                for i := 0; i < 5; i++ {  
                   rawSrc, _ := imgEl.GetAttribute("src")  
                   if rawSrc != "" && !strings.Contains(rawSrc, "data:image") {  
                      cleanImgURL = formatImgURL(rawSrc)  
                      break  
                   }  
                   time.Sleep(300 * time.Millisecond)  
                }  
             }  
  
             if cleanImgURL == "" {  
                continue  
             }  
  
             localPath := filepath.Join(imgDir, id+".jpg")  
  
             log.Printf("[进度 %d/%d] 正在抓取: %s", len(allItems)+1, cfg.TargetCount, title)  
  
             // 使用配置的代理和间隔下载图片  
             if err := downloadFile(cleanImgURL, localPath, cfg); err == nil {  
                allItems[id] = Item{  
                   ID: id, Title: title, Price: price, RemoteURL: cleanImgURL, LocalPath: localPath,  
                }  
                // 使用 YAML 设定的间隔  
                time.Sleep(time.Duration(cfg.DownloadInterval) * time.Second)  
             }  
          }  
          if len(allItems) >= cfg.TargetCount {  
             break  
          }  
       }  
  
       page.Evaluate("window.scrollBy(0, window.innerHeight * 1.5)")  
       time.Sleep(2 * time.Second)  
    }  
  
    saveToExcel(allItems, "goofish_data.xlsx")  
    log.Println("任务圆满完成！")  
}  
  
// --- 辅助功能 ---  
func loadConfig(path string) *Config {  
    content, err := os.ReadFile(path)  
    if err != nil {  
       log.Fatalf("读取配置文件失败: %v", err)  
    }  
    var cfg Config  
    if err := yaml.Unmarshal(content, &cfg); err != nil {  
       log.Fatalf("解析配置文件失败: %v", err)  
    }  
    return &cfg  
}  
  
func downloadFile(targetURL string, filePath string, cfg *Config) error {  
    transport := &http.Transport{}  
    if cfg.Proxy.Enabled {  
       proxyAddr, _ := url.Parse(cfg.Proxy.Server)  
       transport.Proxy = http.ProxyURL(proxyAddr)  
       // 如果代理有账号密码，通常格式为 http://user:pass@host:port       // 这里简单处理，如果 server 没带，可通过 Request Header 注入  
    }  
  
    client := &http.Client{Transport: transport, Timeout: 15 * time.Second}  
    req, _ := http.NewRequest("GET", targetURL, nil)  
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")  
  
    resp, err := client.Do(req)  
    if err != nil {  
       return err  
    }  
    defer resp.Body.Close()  
  
    out, _ := os.Create(filePath)  
    defer out.Close()  
    _, err = io.Copy(out, resp.Body)  
    return err  
}  
  
// (extractID, formatImgURL, saveToExcel 函数保持之前逻辑不变)  
func extractID(u string) string {  
    if strings.Contains(u, "id=") {  
       parts := strings.Split(u, "id=")  
       return strings.Split(parts[1], "&")[0]  
    }  
    return fmt.Sprintf("%d", time.Now().UnixNano())  
}  
  
func formatImgURL(u string) string {  
    if u == "" {  
       return ""  
    }  
    if strings.HasPrefix(u, "//") {  
       u = "https:" + u  
    }  
    for _, ext := range []string{".jpg_", ".png_", ".jpeg_"} {  
       if idx := strings.Index(u, ext); idx != -1 {  
          return u[:idx+len(ext)-1]  
       }  
    }  
    return u  
}  
  
func saveToExcel(items map[string]Item, fileName string) {  
    f := excelize.NewFile()  
    sheet := "Sheet1"  
    linkStyle, _ := f.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "0000FF", Underline: "single"}})  
    headers := []string{"商品ID", "标题", "价格", "原图链接", "本地超链接"}  
    for i, h := range headers {  
       cell, _ := excelize.CoordinatesToCellName(i+1, 1)  
       f.SetCellValue(sheet, cell, h)  
    }  
    row := 2  
    for _, item := range items {  
       f.SetCellValue(sheet, fmt.Sprintf("A%d", row), item.ID)  
       f.SetCellValue(sheet, fmt.Sprintf("B%d", row), item.Title)  
       f.SetCellValue(sheet, fmt.Sprintf("C%d", row), item.Price)  
       f.SetCellValue(sheet, fmt.Sprintf("D%d", row), item.RemoteURL)  
       pathCell := fmt.Sprintf("E%d", row)  
       f.SetCellValue(sheet, pathCell, item.LocalPath)  
       absPath, _ := filepath.Abs(item.LocalPath)  
       f.SetCellHyperLink(sheet, pathCell, absPath, "External")  
       f.SetCellStyle(sheet, pathCell, pathCell, linkStyle)  
       row++  
    }  
    f.SetColWidth(sheet, "B", "B", 40)  
    f.SetColWidth(sheet, "E", "E", 35)  
    f.SaveAs(fileName)  
}
```