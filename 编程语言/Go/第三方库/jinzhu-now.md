### 场景 1：获取“当前”语义下的整点时刻

不用再手写 `time.Date(t.Year(), t.Month(), ...)`，一行代码搞定。

Go

```
package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)

func main() {
	// 1. 获取今天的开始时刻 (00:00:00)
	fmt.Println(now.BeginningOfDay()) 

	// 2. 获取本季度的末尾时刻
	fmt.Println(now.EndOfQuarter())

	// 3. 获取本周一的时刻
	// 注意：默认周日是每周第一天，可以配置
	fmt.Println(now.BeginningOfWeek()) 
}
```

---

### 场景 2：基于特定时间进行偏移

如果你想知道“2023年5月20日”那周的周六是什么时候：

Go

```
func main() {
	t := time.Date(2023, 5, 20, 0, 0, 0, 0, time.Local)
	
	// 基于 t 这个时间点进行计算
	monday := now.New(t).BeginningOfWeek()
	fmt.Println(monday) // 2023-05-14 ... (默认周日开始)
}
```

---

### 场景 3：自定义“周一”为每周第一天

很多业务逻辑要求周一才是新的一周，`now` 提供了全局和实例级的配置。

Go

```
func main() {
	// 设置全局配置：周一作为一周的第一天
	now.WeekStartDay = time.Monday
	
	fmt.Println(now.BeginningOfWeek()) // 现在会返回本周一的 00:00
}
```

---

### 场景 4：解析“模糊”的时间字符串

这是 `now` 最惊艳的功能之一，它内置了大量的 Layout，能自动识别各种格式。

Go

```
func main() {
	// 自动识别 "2024-12-25" 或 "2024/12/25 13:00"
	t, err := now.Parse("2024-12-25 15:30")
	if err == nil {
		fmt.Println(t.Unix())
	}

	// 甚至支持一些简单的相对语义（取决于版本支持情况）
	t2, _ := now.Parse("15:30") // 今天的 15:30
	fmt.Println(t2)
}
```

---

### 场景 5：获取常用的时间区间 (Must-Have)

写 SQL 查询时，经常需要 `BETWEEN ? AND ?`，用 `now` 取区间非常丝滑。

Go

```
func main() {
	// 快速获取上个月的起止时间
	lastMonth := now.With(time.Now().AddDate(0, -1, 0))
	start := lastMonth.BeginningOfMonth()
	end := lastMonth.EndOfMonth()

	fmt.Printf("上月区间: %v 至 %v\n", start, end)
}
```