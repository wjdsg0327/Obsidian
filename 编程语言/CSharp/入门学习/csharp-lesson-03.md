# 第三课：C# 集合、泛型、Lambda、LINQ

> 目标：掌握 C# 日常开发最高频的一组能力：集合处理、泛型抽象、Lambda 写法和 LINQ 查询

---

## 1. 本课目标

本课学完后，你应该能掌握：

- 数组和 `List<T>` 的区别
- 常见集合类型的使用方式
- 泛型的基本思想
- `Dictionary<TKey, TValue>` 的使用
- Lambda 表达式的写法
- LINQ 的常见操作：筛选、投影、排序、聚合
- C# 集合处理和 Java / Go / Python 的对比

---

## 2. 为什么这一课很重要

如果说前两课是“会写 C#”，那这一课就是“开始像 C# 开发者一样写代码”。

在 C# 业务开发里，你会非常频繁地写：

- 列表遍历
- 条件过滤
- 数据转换
- 字典查找
- 按条件排序
- 集合统计

这些能力大多围绕：

- `List<T>`
- `Dictionary<TKey, TValue>`
- Lambda
- LINQ

---

## 3. 数组 vs List< T>

## 数组

```csharp
int[] nums = { 1, 2, 3 };
```

特点：

- 长度固定
- 适合已知数量的数据

### List< T>

```csharp
List<int> nums = new List<int> { 1, 2, 3 };
nums.Add(4);
```

特点：

- 长度可动态变化
- 实战中比数组更常用

### 对比

- **Java**：类似 `ArrayList<T>`
- **Python**：类似 `list`
- **Go**：类似 `slice`

### 速记

- 固定长度用数组
- 日常业务开发优先 `List<T>`
- `List<T>` 是 C# 最常用集合之一

---

## 4. List< T> 基本操作

### 创建列表

```csharp
List<string> names = new List<string>();
```

或：

```csharp
List<string> names = new List<string> { "Tom", "Jack", "Lucy" };
```

### 添加元素

```csharp
names.Add("Bob");
```

### 访问元素

```csharp
Console.WriteLine(names[0]);
```

### 删除元素

```csharp
names.Remove("Jack");
names.RemoveAt(0);
```

### 获取数量

```csharp
Console.WriteLine(names.Count);
```

### 遍历

```csharp
foreach (string name in names)
{
    Console.WriteLine(name);
}
```

### 速记

- 添加：`Add`
- 删除指定值：`Remove`
- 删除指定下标：`RemoveAt`
- 长度：`Count`
- 遍历：`foreach`

---

## 5. Dictionary<TKey, TValue>

### 示例

```csharp
Dictionary<string, int> scores = new Dictionary<string, int>();
scores["Tom"] = 95;
scores["Jack"] = 88;
```

### 读取值

```csharp
Console.WriteLine(scores["Tom"]);
```

### 判断 key 是否存在

```csharp
if (scores.ContainsKey("Tom"))
{
    Console.WriteLine("exists");
}
```

### 更安全的读取方式

```csharp
if (scores.TryGetValue("Tom", out int score))
{
    Console.WriteLine(score);
}
```

### 遍历字典

```csharp
foreach (KeyValuePair<string, int> item in scores)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}
```

### 对比

- **Java**：类似 `HashMap<K, V>`
- **Python**：类似 `dict`
- **Go**：类似 `map[K]V`

### 速记

- 键值对集合用 `Dictionary<TKey, TValue>`
- 判断是否存在优先 `ContainsKey`
- 取值更稳妥用 `TryGetValue`

---

## 6. 泛型是什么

### 示例

```csharp
List<int> nums = new List<int>();
List<string> names = new List<string>();
```

这里的 `<int>`、`<string>` 就是泛型参数。

### 说明

泛型的核心价值：

- 代码复用
- 类型安全
- 避免强制类型转换

### 类比

- **Java**：和 Java 泛型非常接近
- **Go**：新版也支持泛型，但生态和使用方式还不如 C# / Java 历史久
- **Python**：更多依赖鸭子类型和 typing 注解

### 速记

- 泛型 = 参数化类型
- 常见写法：`List<T>`、`Dictionary<TKey, TValue>`
- 好处：复用 + 安全 + 清晰

---

## 7. 泛型方法

### 示例

```csharp
static T Echo<T>(T value)
{
    return value;
}
```

调用：

```csharp
int a = Echo(123);
string b = Echo("hello");
```

### 说明

- `T` 是类型占位符
- 调用时编译器会自动推断类型

### 速记

- 泛型方法格式：`方法名<T>(T value)`
- `T` 只是一个约定俗成的名字
- 不局限于 `T`，也可以写 `TKey`、`TValue`

---

## 8. IEnumerable 基础理解

很多 LINQ 操作都基于可枚举对象。

常见可枚举集合包括：

- 数组
- `List<T>`
- `Dictionary<TKey, TValue>`
- 很多查询结果

只要一个对象“可以被 foreach 遍历”，通常你就可以把它理解为适合做 LINQ 操作。

### 速记

- 能 `foreach` 的对象，通常就能配合 LINQ 使用
- LINQ 常处理的是一批数据，而不是单个对象

---

## 9. Lambda 表达式

### 最简单例子

```csharp
x => x * 2
```

### 含义

- `=>` 左边是参数
- 右边是表达式或代码块

### 示例：传给方法

```csharp
List<int> nums = new List<int> { 1, 2, 3, 4, 5 };
List<int> result = nums.Where(x => x > 3).ToList();
```

### 多参数写法

```csharp
(a, b) => a + b
```

### 代码块写法

```csharp
x =>
{
    Console.WriteLine(x);
    return x * 2;
}
```

### 对比

- **Java**：类似 lambda
- **Python**：类似 `lambda x: x * 2`
- **Go**：更像匿名函数，但语法不同

### 速记

- `=>` 是 Lambda 的标志
- 简单逻辑可直接写表达式
- 复杂逻辑可写代码块

---

## 10. LINQ 是什么

LINQ = Language Integrated Query。

你可以理解为：

> 用统一、简洁、链式的方式操作集合数据。

最常见用途：

- 筛选数据
- 映射数据
- 排序
- 求和
- 分组
- 取第一条/单条

### 速记

- LINQ 是 C# 最重要的生产力工具之一
- 写集合处理时优先想到 LINQ

---

## 11. Where：筛选

### 示例

```csharp
List<int> nums = new List<int> { 1, 2, 3, 4, 5, 6 };
List<int> evenNums = nums.Where(x => x % 2 == 0).ToList();
```

### 说明

- `Where` 用来过滤数据
- 满足条件的保留，不满足的丢掉

### 对比

- **Java**：类似 Stream 的 `filter`
- **Python**：类似列表推导式过滤
- **Go**：通常手写循环

### 速记

- `Where` = 过滤
- 最后常接 `.ToList()`

---

## 12. Select：投影 / 转换

### 示例

```csharp
List<string> names = new List<string> { "tom", "jack" };
List<string> upperNames = names.Select(x => x.ToUpper()).ToList();
```

### 说明

- `Select` 用来把一种数据转换成另一种形式
- 一对一映射

### 对比

- **Java**：类似 Stream 的 `map`
- **Python**：类似 `map` 或列表推导式

### 速记

- `Select` = 转换 / 映射
- `Where` 是过滤，`Select` 是变形

---

## 13. OrderBy / OrderByDescending：排序

### 示例

```csharp
List<int> nums = new List<int> { 5, 2, 8, 1 };
List<int> sorted = nums.OrderBy(x => x).ToList();
List<int> desc = nums.OrderByDescending(x => x).ToList();
```

### 对对象排序

```csharp
List<Student> students = new List<Student>
{
    new Student { Name = "Tom", Score = 90 },
    new Student { Name = "Jack", Score = 85 },
    new Student { Name = "Lucy", Score = 95 }
};

List<Student> sortedStudents = students.OrderBy(x => x.Score).ToList();
```

### 速记

- 升序：`OrderBy`
- 降序：`OrderByDescending`
- 参数通常写排序依据

---

## 14. First / FirstOrDefault

### 示例

```csharp
int first = nums.First();
int firstEven = nums.First(x => x % 2 == 0);
```

### 更安全写法

```csharp
int firstEven = nums.FirstOrDefault(x => x % 2 == 0);
```

### 说明

- `First()`：取第一个元素，没有会抛异常
- `FirstOrDefault()`：没有时返回默认值

### 注意

对引用类型，默认值通常是 `null`
对 `int`，默认值通常是 `0`

### 速记

- 确定有数据时用 `First`
- 不确定时优先 `FirstOrDefault`

---

## 15. Any / All

### 示例

```csharp
bool hasEven = nums.Any(x => x % 2 == 0);
bool allPositive = nums.All(x => x > 0);
```

### 说明

- `Any`：是否存在至少一个满足条件
- `All`：是否全部满足条件

### 速记

- `Any` = 至少一个
- `All` = 全部

---

## 16. Count / Sum / Max / Min / Average

### 示例

```csharp
int count = nums.Count();
int sum = nums.Sum();
int max = nums.Max();
int min = nums.Min();
double avg = nums.Average();
```

### 对对象集合操作

```csharp
int totalScore = students.Sum(x => x.Score);
int maxScore = students.Max(x => x.Score);
```

### 速记

- `Count` 统计个数
- `Sum` 求和
- `Max` / `Min` 求最大最小
- `Average` 求平均值

---

## 17. 匿名对象

### 示例

```csharp
var result = students.Select(x => new
{
    x.Name,
    Level = x.Score >= 90 ? "优秀" : "普通"
}).ToList();
```

### 说明

- `new { ... }` 可以快速创建临时对象
- 常用于查询结果转换

### 对比

- Java 没有这么直接的匿名对象写法
- Python 字典能起到部分类似作用，但语义不同

### 速记

- 临时结构可用匿名对象
- 经常和 `Select` 搭配使用

---

## 18. 组合使用示例

### 示例

```csharp
List<Student> students = new List<Student>
{
    new Student { Name = "Tom", Score = 90 },
    new Student { Name = "Jack", Score = 85 },
    new Student { Name = "Lucy", Score = 95 },
    new Student { Name = "Bob", Score = 70 }
};

var topStudents = students
    .Where(x => x.Score >= 85)
    .OrderByDescending(x => x.Score)
    .Select(x => new
    {
        x.Name,
        x.Score
    })
    .ToList();
```

### 说明

这段代码完成了：

1. 过滤出分数 >= 85 的学生
2. 按分数倒序排列
3. 只保留姓名和分数
4. 转为列表

### 速记

- LINQ 最常见写法就是链式调用
- 典型顺序：`Where -> OrderBy -> Select -> ToList`

---

## 19. 延迟执行的基础认知

LINQ 很多方法是延迟执行的。

例如：

```csharp
var query = nums.Where(x => x > 2);
```

这时通常还没有真正开始遍历结果。
只有在你：

- `foreach`
- `ToList()`
- `Count()`
- `First()`

这种真正取结果时，查询才会执行。

### 速记

- 很多 LINQ 不是立刻执行
- 想固定结果时，常用 `ToList()`

---

## 20. 常见误区

### 误区 1：`Where` 会直接修改原集合

不会。
它返回的是新的查询结果。

### 误区 2：`Select` 是过滤

不是。
`Select` 是转换，过滤要用 `Where`。

### 误区 3：字典取值直接用 `dict[key]` 永远安全

不安全。
key 不存在时可能报错。
优先考虑：

```csharp
TryGetValue
```

### 误区 4：数组和 `List<T>` 一样

不一样。
数组固定长度，`List<T>` 可动态扩容。

---

## 21. 本课重点总结

### 你必须先掌握的 8 个核心点

1. 日常开发中，`List<T>` 比数组更常用
2. 键值对集合用 `Dictionary<TKey, TValue>`
3. 泛型让代码更安全、更通用
4. Lambda 是给方法传逻辑的一种简洁写法
5. `Where` 用于过滤
6. `Select` 用于转换
7. `OrderBy` / `OrderByDescending` 用于排序
8. LINQ 是 C# 集合处理核心能力

---

## 22. Java / Go / Python 对比总结

| 知识点 | C# | Java | Go | Python |
|---|---|---|---|---|
| 动态数组 | `List<T>` | `ArrayList<T>` | `slice` | `list` |
| 字典 | `Dictionary<K,V>` | `HashMap<K,V>` | `map` | `dict` |
| 泛型 | 很成熟 | 很成熟 | 新增但生态较新 | typing 为主 |
| Lambda | `x => ...` | `x -> ...` | 匿名函数 | `lambda` |
| 过滤 | `Where` | `filter` / Stream | 手写循环 | 列表推导式 |
| 映射 | `Select` | `map` / Stream | 手写循环 | `map` / 推导式 |

---

## 23. 面试 / 实战高频提醒

- `List<T>` 是高频核心集合
- `Dictionary<TKey, TValue>` 取值优先 `TryGetValue`
- `Where` 是过滤，`Select` 是映射
- `FirstOrDefault` 通常比 `First` 更稳妥
- 很多 LINQ 操作是延迟执行的
- 链式查询是 C# 非常常见的业务写法

---

## 24. 本课小练习

### 练习 1：List 基础操作

要求：

- 创建一个 `List<int>`
- 加入 5 个数字
- 删除一个数字
- 遍历输出

### 练习 2：Dictionary 基础操作

要求：

- 创建一个 `Dictionary<string, int>`
- 保存 3 个学生成绩
- 用 `TryGetValue` 查询某个学生的成绩

### 练习 3：LINQ 过滤偶数

要求：

- 有一个整数列表
- 用 `Where` 过滤出偶数
- 输出结果

### 练习 4：LINQ 转大写

要求：

- 有一个字符串列表
- 用 `Select` 转成大写
- 输出结果

### 练习 5：学生成绩排序

要求：

- 定义 `Student` 类，包含 `Name` 和 `Score`
- 创建学生列表
- 按分数降序排序
- 输出姓名和分数

---

## 25. 一页速背版

### List

```csharp
List<int> nums = new List<int> { 1, 2, 3 };
nums.Add(4);
nums.Remove(2);
Console.WriteLine(nums.Count);
```

### Dictionary

```csharp
Dictionary<string, int> scores = new Dictionary<string, int>();
scores["Tom"] = 95;

if (scores.TryGetValue("Tom", out int score))
{
    Console.WriteLine(score);
}
```

### Lambda

```csharp
x => x * 2
(a, b) => a + b
```

### LINQ 常用

```csharp
nums.Where(x => x > 2).ToList();
nums.Select(x => x * 2).ToList();
nums.OrderBy(x => x).ToList();
nums.FirstOrDefault();
nums.Any(x => x > 10);
nums.Sum();
```

### 核心结论

- `List<T>` 高频
- `Dictionary<TKey, TValue>` 高频
- `Where` 过滤
- `Select` 转换
- `OrderBy` 排序
- LINQ 是生产力核心

---

## 26. 下一课预告

下一课学习：

- 异常处理
- 委托 delegate
- `Action` / `Func`
- 事件 event
- 异步编程 `async/await`

这部分会让你真正理解 C# 为什么非常适合工程开发。
