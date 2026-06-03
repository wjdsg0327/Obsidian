### 场景 1：实时监听 Binlog 变更 (Canal)

这是该库最著名的用途：实时捕获数据库的增删改操作。常用于 **缓存同步（Redis）** 或 **异构数据库同步（ES/ClickHouse）**。

```go
package main

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
)

// 1. 定义处理器：嵌入 DummyEventHandler，只需要重写你感兴趣的方法
type MyHandler struct {
	canal.DummyEventHandler
}

// 2. 实现 OnRow 方法：当数据库发生 Insert/Update/Delete 时触发
func (h *MyHandler) OnRow(e *canal.RowsEvent) error {
	// e.Action: 动作类型 (insert, update, delete)
	// e.Table:  数据库名 e.Table.Schema, 表名 e.Table.Name
	fmt.Printf("[%s] 在表 %s.%s 上干了坏事\\n", e.Action, e.Table.Schema, e.Table.Name)

	// e.Rows: 具体的行数据
	for _, row := range e.Rows {
		fmt.Printf("📦 行数据: %v\\n", row)
	}
	return nil
}

func main() {
	// 3. 基础配置
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "gossip_boy"
	cfg.Password = "secret123"
	cfg.Flavor = "mysql"
	
	// ServerID 非常重要：同一集群内不能重复，否则会互踢掉线
	cfg.ServerID = 1234 

	// 4. 初始化并运行
	c, _ := canal.NewCanal(cfg)
	c.SetEventHandler(&MyHandler{})
	c.Run()
}
```