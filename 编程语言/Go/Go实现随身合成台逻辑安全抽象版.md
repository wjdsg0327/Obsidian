---
title: Go 实现随身合成台逻辑（安全抽象版）
date: 2026-06-10
tags:
  - Go
  - 游戏工具
  - 逆向学习
  - DLL注入
  - 安全抽象
  - 芙芙启动器
---

# Go 实现随身合成台逻辑（安全抽象版）

> 说明：本笔记是把“随身合成台”这种功能抽象成 Go 代码来学习其**业务逻辑**：打开合成界面、检查材料、扣除材料、产出物品。  
> 这里不包含游戏进程注入、Hook、内存修改、绕过反作弊或调用第三方游戏内部函数的代码。

## 背景

芙芙启动器这类第三方启动器可能包含“随身合成台”功能。真实实现如果作用于游戏进程，通常可能涉及：

- 启动器进程控制
- DLL 注入
- Hook 游戏函数
- 修改或调用游戏内部 UI / 交互逻辑
- 模拟打开合成台界面

这些内容如果用于第三方商业游戏，可能违反服务条款或触发反作弊风险。因此这里仅整理一个**安全学习版**：用 Go 模拟合成台的状态机和配方系统。

## 功能目标

这个 Go 示例实现：

1. 注册合成配方
2. 管理玩家背包材料
3. 判断材料是否足够
4. 执行合成
5. 扣除材料
6. 增加产物
7. 输出背包状态

## Go 示例代码

```go
package main

import (
	"errors"
	"fmt"
	"sort"
)

// ItemID 表示物品 ID。真实游戏里通常会对应内部物品表。
type ItemID string

// ItemStack 表示一种物品及数量。
type ItemStack struct {
	ID    ItemID
	Name  string
	Count int
}

// Recipe 表示一条合成配方。
type Recipe struct {
	ID          string
	Name        string
	Ingredients []ItemStack
	Output      ItemStack
}

// Inventory 表示背包。这里用 map 模拟：ItemID -> 数量。
type Inventory struct {
	items map[ItemID]int
	names map[ItemID]string
}

func NewInventory() *Inventory {
	return &Inventory{
		items: make(map[ItemID]int),
		names: make(map[ItemID]string),
	}
}

func (inv *Inventory) Add(item ItemStack) {
	if item.Count <= 0 {
		return
	}
	inv.items[item.ID] += item.Count
	if item.Name != "" {
		inv.names[item.ID] = item.Name
	}
}

func (inv *Inventory) Has(item ItemStack) bool {
	return inv.items[item.ID] >= item.Count
}

func (inv *Inventory) Remove(item ItemStack) error {
	if item.Count <= 0 {
		return nil
	}
	if !inv.Has(item) {
		return fmt.Errorf("材料不足: %s 需要 %d，当前 %d", item.Name, item.Count, inv.items[item.ID])
	}
	inv.items[item.ID] -= item.Count
	if inv.items[item.ID] == 0 {
		delete(inv.items, item.ID)
	}
	return nil
}

func (inv *Inventory) Print() {
	fmt.Println("\n当前背包：")
	ids := make([]string, 0, len(inv.items))
	for id := range inv.items {
		ids = append(ids, string(id))
	}
	sort.Strings(ids)

	for _, rawID := range ids {
		id := ItemID(rawID)
		name := inv.names[id]
		if name == "" {
			name = rawID
		}
		fmt.Printf(" - %-10s x%d\n", name, inv.items[id])
	}
}

// CraftingTable 表示合成台系统。
type CraftingTable struct {
	recipes map[string]Recipe
}

func NewCraftingTable() *CraftingTable {
	return &CraftingTable{recipes: make(map[string]Recipe)}
}

func (ct *CraftingTable) Register(recipe Recipe) error {
	if recipe.ID == "" {
		return errors.New("配方 ID 不能为空")
	}
	if len(recipe.Ingredients) == 0 {
		return errors.New("配方材料不能为空")
	}
	if recipe.Output.ID == "" || recipe.Output.Count <= 0 {
		return errors.New("配方产物无效")
	}
	ct.recipes[recipe.ID] = recipe
	return nil
}

func (ct *CraftingTable) CanCraft(inv *Inventory, recipeID string) error {
	recipe, ok := ct.recipes[recipeID]
	if !ok {
		return fmt.Errorf("未知配方: %s", recipeID)
	}
	for _, ingredient := range recipe.Ingredients {
		if !inv.Has(ingredient) {
			return fmt.Errorf(
				"无法合成 %s：缺少 %s，需要 %d，当前 %d",
				recipe.Name,
				ingredient.Name,
				ingredient.Count,
				inv.items[ingredient.ID],
			)
		}
	}
	return nil
}

func (ct *CraftingTable) Craft(inv *Inventory, recipeID string) error {
	recipe, ok := ct.recipes[recipeID]
	if !ok {
		return fmt.Errorf("未知配方: %s", recipeID)
	}

	// 先检查，避免扣了一半才发现材料不够。
	if err := ct.CanCraft(inv, recipeID); err != nil {
		return err
	}

	// 扣除材料。
	for _, ingredient := range recipe.Ingredients {
		if err := inv.Remove(ingredient); err != nil {
			return err
		}
	}

	// 增加产物。
	inv.Add(recipe.Output)
	fmt.Printf("合成成功：%s -> %s x%d\n", recipe.Name, recipe.Output.Name, recipe.Output.Count)
	return nil
}

func main() {
	// 初始化背包。
	inv := NewInventory()
	inv.Add(ItemStack{ID: "mint", Name: "薄荷", Count: 5})
	inv.Add(ItemStack{ID: "sweet_flower", Name: "甜甜花", Count: 3})
	inv.Add(ItemStack{ID: "crystal_core", Name: "晶核", Count: 1})

	// 初始化合成台。
	ct := NewCraftingTable()

	// 示例配方：这里是演示用，不对应真实游戏数据。
	_ = ct.Register(Recipe{
		ID:   "portable_potion",
		Name: "随身合成：简易药剂",
		Ingredients: []ItemStack{
			{ID: "mint", Name: "薄荷", Count: 2},
			{ID: "sweet_flower", Name: "甜甜花", Count: 1},
		},
		Output: ItemStack{ID: "potion", Name: "简易药剂", Count: 1},
	})

	_ = ct.Register(Recipe{
		ID:   "portable_resin_tool",
		Name: "随身合成：树脂小工具",
		Ingredients: []ItemStack{
			{ID: "crystal_core", Name: "晶核", Count: 1},
			{ID: "mint", Name: "薄荷", Count: 3},
		},
		Output: ItemStack{ID: "resin_tool", Name: "树脂小工具", Count: 1},
	})

	inv.Print()

	// 执行合成。
	if err := ct.Craft(inv, "portable_potion"); err != nil {
		fmt.Println("合成失败:", err)
	}

	if err := ct.Craft(inv, "portable_resin_tool"); err != nil {
		fmt.Println("合成失败:", err)
	}

	inv.Print()
}
```

## 运行方式

保存为 `main.go` 后运行：

```bash
go run main.go
```

示例输出类似：

```text
当前背包：
 - 晶核         x1
 - 薄荷         x5
 - 甜甜花       x3
合成成功：随身合成：简易药剂 -> 简易药剂 x1
合成成功：随身合成：树脂小工具 -> 树脂小工具 x1

当前背包：
 - 简易药剂     x1
 - 树脂小工具   x1
 - 甜甜花       x2
```

## 和真实“随身合成台”的关系

真实工具如果要在游戏里实现随身合成台，通常会多出这些层：

```text
启动器 UI
  ↓
配置/功能开关
  ↓
注入模块或插件模块
  ↓
游戏进程内 Hook / 调用游戏内部函数
  ↓
打开或模拟合成台 UI
  ↓
调用游戏原本的合成逻辑
```

而本 Go 示例只覆盖最上层的**业务抽象**：

```text
配方 + 背包 + 材料检查 + 扣材料 + 给产物
```

它适合用于学习：

- 游戏道具系统建模
- 合成配方建模
- Go 结构体设计
- Go map 管理库存
- 事务式操作：先检查再扣除
- 将游戏功能拆成安全的纯逻辑模块

## 如果以后继续扩展

可以继续加：

- 从 JSON 加载配方
- 从 JSON 保存背包
- 支持批量合成
- 支持材料别名
- 支持配方分类
- 支持合成耗时
- 支持概率产物
- 支持命令行参数
- 支持 Web UI
- 支持单元测试

## JSON 配方扩展示例

后续可以把配方从 Go 代码中移出来：

```json
[
  {
    "id": "portable_potion",
    "name": "随身合成：简易药剂",
    "ingredients": [
      {"id": "mint", "name": "薄荷", "count": 2},
      {"id": "sweet_flower", "name": "甜甜花", "count": 1}
    ],
    "output": {"id": "potion", "name": "简易药剂", "count": 1}
  }
]
```

这样就可以把“程序逻辑”和“游戏数据”分开。

## 安全边界

这份笔记不涉及：

- DLL 注入代码
- 远程线程创建
- Hook 游戏函数
- 修改游戏内存
- 绕过反作弊
- 绕过限制或自动化第三方游戏行为

如需继续学习 DLL 注入，应只在自己写的测试进程和测试 DLL 中做实验。

## 相关笔记

- [[Go 读取 Windows DLL 导入表示例]]
- [[DLL逆向分析]]
- [[PE文件格式]]
- [[Windows静态分析]]
