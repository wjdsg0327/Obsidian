### 场景 1：把各种“乱七八糟”的类型转成数字

在处理 JSON 或配置项时，你永远不知道对方传过来的是 `"100"`（字符串）还是 `100`（数字）。

```go
package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	// 无论输入是 字符串、浮点数 还是 布尔值
	val1 := cast.ToInt("88")      // 88
	val2 := cast.ToInt(88.88)     // 88 (自动取整)
	val3 := cast.ToInt(true)      // 1  (true 转 1, false 转 0)
	val4 := cast.ToInt(nil)       // 0  (nil 安全转 0)

	fmt.Println(val1, val2, val3, val4)
	// Output: 88 88 1 0
}
```

### 场景 2：超宽容的布尔值转换

原生 `strconv.ParseBool` 非常严格，而 `cast` 认为 "yes", "on", "1" 全都是 `true`。

```go
package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	// 模拟从环境变量或配置文件读取开关
	p1 := cast.ToBool("yes") // true
	p2 := cast.ToBool("on")  // true
	p3 := cast.ToBool("1")   // true
	p4 := cast.ToBool(0)     // false

	fmt.Println(p1, p2, p3, p4)
}
```

### 场景 3：快速转换字符串切片

当你拿到一个 `[]interface{}`（比如从数据库或 API 返回的列表），想直接当 `[]string` 用时。

```go
package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	// 这是一个混合类型的切片
	mixed := []interface{}{"apple", 100, 3.14, true}

	// 一行代码全部变字符串
	fruits := cast.ToStringSlice(mixed)

	fmt.Printf("%T: %v\\n", fruits, fruits)
	// Output: []string: [apple 100 3.14 true]
}
```

### 场景 4：处理 Map 的键值对转换

这是处理配置文件（如 YAML/JSON）最强大的地方，把 `map[interface{}]interface{}` 转为标准的 `map[string]interface{}`。

```go
package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	// 很多库解析出来的数据长这样，Key 是 interface{} 很不方便
	rawConfig := map[interface{}]interface{}{
		"port":    8080,
		"enabled": "true",
	}

	// 转换为 string 类型的 key
	cleanConfig := cast.ToStringMap(rawConfig)

	port := cast.ToInt(cleanConfig["port"])
	enabled := cast.ToBool(cleanConfig["enabled"])

	fmt.Printf("Port: %d, Enabled: %v\\n", port, enabled)
	// Output: Port: 8080, Enabled: true
}
```

### 场景 5：带 Error 检查的转换 (ToE 系列)

如果你不希望静默失败（即转换失败返回默认值），而是想捕获错误。

```go
package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	val, err := cast.ToIntE("这是个苹果")
	
	if err != nil {
		fmt.Printf("转换失败: %v\\n", err)
		// Output: 转换失败: unable to cast type string to int
	} else {
		fmt.Println(val)
	}
}
```