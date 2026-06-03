### 场景 1：基础 GET 请求（带自动 JSON 解析）

不需要手动 `ioutil.ReadAll` 和 `json.Unmarshal`，直接绑定到结构体。

```go
package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	client := resty.New()
	user := &User{}

	// Result 会自动将响应体 Unmarshal 到 user 变量中
	resp, err := client.R().
		SetResult(user).
		Get("<https://api.example.com/users/1>")

	if err == nil {
		fmt.Printf("用户名: %s, 状态码: %d\\n", user.Name, resp.StatusCode())
	}
}
```

### 场景 2：POST 请求（发送 JSON 内容）

直接把结构体或 Map 丢进去，Resty 会自动设置 `Content-Type: application/json`。

```go
func main() {
	client := resty.New()

	// 发送 Body 数据
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"title": "foo", "body": "bar", "userId": 1}).
		Post("<https://jsonplaceholder.typicode.com/posts>")

	if err == nil {
		fmt.Println("响应结果:", resp.String())
	}
}
```

### 场景 3：公共配置（BaseURL、Auth、Timeout）

如果你有一堆请求要发往同一个域名，没必要每次都写全路径和 Token。

```go
func main() {
	client := resty.New()

	// 全局公共配置
	client.SetBaseURL("<https://api.github.com>").
		SetAuthToken("YOUR_TOKEN_HERE").
		SetTimeout(5 * time.Second)

	// 这里的路径会自动拼接到 BaseURL 后面
	resp, _ := client.R().Get("/user/repos")
	fmt.Println(resp.StatusCode())
}
```

### 场景 4：自动失败重试机制

这是 Resty 的杀手锏。网络波动时，它能自动按策略重试。

```go
func main() {
	client := resty.New()

	// 设置重试次数为 3，重试间隔 100 毫秒，最大间隔 2 秒
	client.SetRetryCount(3).
		SetRetryWaitTime(100 * time.Millisecond).
		SetRetryMaxWaitTime(2 * time.Second).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				// 只有当状态码为 500 或有错误时才重试
				return r.StatusCode() >= 500
			},
		)

	client.R().Get("<https://unstable-api.com/data>")
}
```

### 场景 5：调试模式（Debug）

排查问题时，一行代码开启“详细日志”，打印所有的 Request 和 Response 详情。

```go
func main() {
	client := resty.New()
	
	// 开启调试，你会看到所有的 Header、Body、URL 打印在控制台
	client.SetDebug(true)

	client.R().
		SetQueryParams(map[string]string{
			"page": "1",
			"limit": "10",
		}).
		Get("<https://httpbin.org/get>")
}
```

### 场景 6：下载文件

Resty 处理文件流非常简单，直接指定保存路径。

```go
func main() {
	client := resty.New()

	// 将返回的内容直接写入文件
	_, err := client.R().
		SetOutput("./logo.png").
		Get("<https://example.com/image.png>")

	if err == nil {
		fmt.Println("文件下载成功！")
	}
}
```