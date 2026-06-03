# 第六课：C# 可空引用类型、值类型与引用类型、record、模式匹配、switch expression、var/object/dynamic

> 目标：掌握现代 C# 最容易“看起来简单、实际很关键”的几组特性，建立更清晰的类型系统认知

---

## 1. 本课目标

本课学完后，你应该能掌握：

- 什么是可空引用类型
- 值类型和引用类型的核心区别
- `record` 是什么，适合解决什么问题
- 模式匹配的基本思路
- `switch expression` 的写法
- `var`、`object`、`dynamic` 的区别
- 为什么这些特性会直接影响代码安全性和可读性

---

## 2. 为什么这一课重要

这一课不是“多几个语法糖”这么简单。
它决定你写出来的 C#：

- 容不容易空指针
- 是否表达清晰的数据语义
- 能不能写出更现代、更易读的分支逻辑
- 是否正确理解类型系统

很多 Java 开发者刚学 C# 时，最容易在这里混淆：

- `null` 到底怎么处理
- `class` / `struct` / `record` 到底怎么选
- `var` 和 `dynamic` 到底是不是一回事

---

## 3. 可空引用类型是什么

在旧习惯里，引用类型默认都可能是 `null`。

例如：

```csharp
string name = null;
```

这会让代码里到处潜伏空引用问题。

现代 C# 引入了 **可空引用类型**，核心思想是：

- `string`：默认不该为 `null`
- `string?`：明确表示它可能为 `null`

### 示例

```csharp
string name = "Tom";
string? nickName = null;
```

### 速记

- `string`：不应为 `null`
- `string?`：可能为 `null`
- `?` 表示“可空”

---

## 4. 为什么可空引用类型重要

它的价值是：

- 提前发现空引用风险
- 让编译器帮你检查代码
- 让接口语义更清晰

比如：

```csharp
string? name = Console.ReadLine();
Console.WriteLine(name.Length);
```

这里编译器会提醒你：

- `name` 可能是 `null`
- 直接访问 `Length` 不安全

### 速记

- 可空引用类型的核心不是语法，而是“编译期空安全”

---

## 5. 可空值检查

### 方式 1：if 判断

```csharp
string? name = Console.ReadLine();

if (name != null)
{
    Console.WriteLine(name.Length);
}
```

### 方式 2：空条件运算符

```csharp
Console.WriteLine(name?.Length);
```

### 方式 3：空合并运算符

```csharp
string result = name ?? "default";
```

### 说明

- `?.`：如果对象不为 `null` 才继续访问
- `??`：如果左边是 `null`，就用右边默认值

### 速记

- 判空：`if (x != null)`
- 空安全访问：`?.`
- 提供默认值：`??`

---

## 6. null-forgiving 运算符 `!`

### 示例

```csharp
string? name = Console.ReadLine();
Console.WriteLine(name!.Length);
```

### 说明

`!` 的意思是：

> 我确信这里不是 null，请编译器别警告我。

但这只是“压警告”，不是运行时保护。
如果你判断错了，运行时还是可能报错。

### 速记

- `!` 是“我保证不为空”
- 能不用就少用
- 优先真正做好判空

---

## 7. 值类型 vs 引用类型

这是 C# 类型系统的核心概念。

### 值类型常见例子

- `int`
- `double`
- `bool`
- `DateTime`
- `struct`
- `enum`

### 引用类型常见例子

- `string`
- `class`
- `object`
- `array`
- `delegate`
- `record class`

### 核心区别

- 值类型：赋值时复制“值本身”
- 引用类型：赋值时复制“引用地址”

---

## 8. 值类型复制示例

```csharp
int a = 10;
int b = a;
b = 20;

Console.WriteLine(a); // 10
Console.WriteLine(b); // 20
```

### 说明

- `b = a` 是值拷贝
- 改 `b` 不影响 `a`

### 速记

- 值类型赋值 = 拷贝数据
- 两边互不影响

---

## 9. 引用类型复制示例

```csharp
class Person
{
    public string Name { get; set; } = "";
}

Person p1 = new Person { Name = "Tom" };
Person p2 = p1;
p2.Name = "Jack";

Console.WriteLine(p1.Name); // Jack
```

### 说明

- `p2 = p1` 拷贝的是引用
- `p1` 和 `p2` 指向同一个对象

### 速记

- 引用类型赋值 = 拷贝引用
- 改一边，另一边可能受影响

---

## 10. Nullable 值类型

值类型本来不能为 `null`，但可以通过 `?` 变成可空值类型。

### 示例

```csharp
int? age = null;
DateTime? time = null;
```

### 说明

- `int` 不能是 `null`
- `int?` 可以是 `null`

### 常见用法

```csharp
if (age.HasValue)
{
    Console.WriteLine(age.Value);
}
```

或：

```csharp
int result = age ?? 0;
```

### 速记

- 值类型可空写法：`int?`、`DateTime?`
- 常配合 `??` 使用

---

## 11. boxing / unboxing 基础认知

### 示例

```csharp
int x = 10;
object obj = x;      // boxing
int y = (int)obj;    // unboxing
```

### 说明

- boxing：把值类型装箱成 `object`
- unboxing：再拆回来

### 注意

- 会带来额外开销
- 泛型很多时候就是为了减少这类问题

### 速记

- 值类型转 `object`：boxing
- `object` 转回值类型：unboxing

---

## 12. record 是什么

你可以把 `record` 理解为：

> 用来表达“数据对象”的一种更现代写法。

### 示例

```csharp
public record User(string Name, int Age);
```

使用：

```csharp
User user = new User("Tom", 20);
Console.WriteLine(user.Name);
```

### 和 class 的区别直觉

`class` 更偏“对象行为”
`record` 更偏“数据表达”

### 速记

- `record` 适合表示 DTO、配置、返回结果、值对象
- 强调“数据本身”而不是复杂行为

---

## 13. record 的值相等

### class 默认比较

```csharp
class User
{
    public string Name { get; set; } = "";
    public int Age { get; set; }
}

var a = new User { Name = "Tom", Age = 20 };
var b = new User { Name = "Tom", Age = 20 };

Console.WriteLine(a == b); // False
```

### record 默认比较

```csharp
public record User(string Name, int Age);

var a = new User("Tom", 20);
var b = new User("Tom", 20);

Console.WriteLine(a == b); // True
```

### 说明

- `class` 默认更偏引用比较
- `record` 默认更偏值比较

### 速记

- `record` 天然适合“两个对象内容一样就算一样”的场景

---

## 14. record 的 with 表达式

### 示例

```csharp
public record User(string Name, int Age);

var u1 = new User("Tom", 20);
var u2 = u1 with { Age = 21 };
```

### 说明

- `with` 会基于旧对象创建一个新对象
- 很适合不可变数据修改

### 对比

- Java 老风格没这么方便
- 很像函数式风格的数据拷贝更新

### 速记

- `with` = 基于原对象复制并修改部分值

---

## 15. record class 和 record struct

### 示例

```csharp
public record class User(string Name, int Age);
public record struct Point(int X, int Y);
```

### 说明

- `record class`：引用类型
- `record struct`：值类型

入门阶段你先重点掌握普通 `record`，大多数情况下它等价于 `record class` 直觉使用。

### 速记

- `record` 默认更多当作数据类来理解
- `record struct` 是值类型数据对象

---

## 16. 模式匹配是什么

模式匹配可以理解为：

> 不只是判断“是不是等于某个值”，还可以判断“是不是某种类型、某种结构、某种条件”。

这是现代 C# 很强的一部分。

### 基础例子：类型判断

```csharp
object obj = "hello";

if (obj is string s)
{
    Console.WriteLine(s.Length);
}
```

### 说明

- `obj is string s` 表示：
  - 如果 `obj` 是 `string`
  - 那就把它赋给变量 `s`

### 速记

- `is` 不只做判断，还能顺便完成类型转换绑定

---

## 17. 常量模式

### 示例

```csharp
int score = 100;

if (score is 100)
{
    Console.WriteLine("满分");
}
```

### 速记

- `is 100` 这种写法可以直接匹配常量

---

## 18. null 模式

### 示例

```csharp
if (name is null)
{
    Console.WriteLine("为空");
}
```

### 说明

有时你会看到这种写法，它比 `== null` 更符合现代 C# 风格语义。

### 速记

- 判空也能用模式匹配：`is null`

---

## 19. 属性模式

### 示例

```csharp
if (user is { Age: >= 18 })
{
    Console.WriteLine("成年人");
}
```

### 说明

这表示：

- `user` 不为 `null`
- 并且 `user.Age >= 18`

### 速记

- 可以直接在对象属性上做匹配判断
- 写复杂条件时可读性很好

---

## 20. switch expression 是什么

这是 `switch` 的更现代写法。

### 传统 switch

```csharp
string level;

switch (score)
{
    case >= 90:
        level = "A";
        break;
    case >= 80:
        level = "B";
        break;
    default:
        level = "C";
        break;
}
```

### switch expression

```csharp
string level = score switch
{
    >= 90 => "A",
    >= 80 => "B",
    _ => "C"
};
```

### 说明

- 更简洁
- 更适合“根据条件产出一个值”的场景

### 速记

- `switch expression` 常用于赋值
- `_` 表示默认分支

---

## 21. 类型模式配合 switch

### 示例

```csharp
object obj = 123;

string result = obj switch
{
    int n => $"int: {n}",
    string s => $"string length: {s.Length}",
    null => "null",
    _ => "unknown"
};
```

### 速记

- `switch` 不只能判断值，还能判断类型
- 这是现代 C# 很常见的写法

---

## 22. when 子句

### 示例

```csharp
string result = score switch
{
    int n when n >= 90 => "A",
    int n when n >= 80 => "B",
    _ => "C"
};
```

### 说明

- `when` 用来加额外条件
- 适合更复杂的匹配逻辑

### 速记

- `when` = 模式匹配里的附加条件

---

## 23. var 是什么

`var` 是 **编译期类型推断**。

### 示例

```csharp
var name = "Tom";
var age = 20;
```

实际类型分别是：

- `name` -> `string`
- `age` -> `int`

### 关键点

- 变量类型在编译期就确定了
- 不是动态类型

### 速记

- `var` = 编译器帮你写类型
- `var` ≠ `dynamic`

---

## 24. object 是什么

`object` 是 C# 所有类型的基类。

### 示例

```csharp
object a = 123;
object b = "hello";
```

### 说明

- 什么都能装进 `object`
- 但取出来时，通常需要转换或判断类型

### 示例

```csharp
object obj = "hello";
if (obj is string s)
{
    Console.WriteLine(s.Length);
}
```

### 速记

- `object` 很通用
- 但用多了会损失类型信息和可读性

---

## 25. dynamic 是什么

`dynamic` 表示：

> 把成员解析推迟到运行时。

### 示例

```csharp
dynamic x = "hello";
Console.WriteLine(x.Length);
```

这个编译器通常不会严格检查 `Length` 是否存在。
真正到运行时才决定。

### 再看一个例子

```csharp
dynamic x = 123;
Console.WriteLine(x.Length);
```

这段代码编译可能过，但运行时会报错。

### 说明

- `dynamic` 更像“关闭一部分编译期类型检查”
- 会降低安全性

### 对比

- 有点像 Python 风格的动态行为
- 但 C# 默认并不是这种模式

### 速记

- `dynamic` = 运行时绑定
- 功能强，但风险高
- 能不用尽量少用

---

## 26. var / object / dynamic 对比

| 关键字 | 类型确定时机 | 是否保留静态类型信息 | 编译期检查 | 典型场景 |
|---|---|---|---|---|
| `var` | 编译期 | 是 | 强 | 简化局部变量声明 |
| `object` | 编译期 | 弱，需要转换 | 有限 | 通用容器、顶层基类 |
| `dynamic` | 运行时 | 弱 | 弱 | 动态调用、特殊互操作 |

### 一句话理解

- `var`：我懒得写类型，但编译器知道
- `object`：我故意装成最通用的父类型
- `dynamic`：我把检查推迟到运行时

---

## 27. 什么时候用 record，什么时候用 class

### 更适合用 record 的场景

- DTO
- API 返回模型
- 配置对象
- 值对象
- 不强调复杂行为的数据结构

### 更适合用 class 的场景

- 有较强业务行为
- 对象生命周期复杂
- 有可变状态管理
- 面向对象建模更明显

### 速记

- 数据为主：优先考虑 `record`
- 行为为主：优先考虑 `class`

---

## 28. 常见误区

### 误区 1：`string?` 只是语法糖

不止。
它本质是让编译器参与空安全检查。

### 误区 2：`var` 是动态类型

不是。
`var` 仍然是静态类型。

### 误区 3：`object` 和 `dynamic` 一样

不一样。
`object` 仍然受编译器类型系统约束，`dynamic` 更多在运行时决定。

### 误区 4：`record` 只是 `class` 的简写

不完全对。
`record` 还带来了值相等、`with` 等语义优势。

### 误区 5：模式匹配只是 switch 换皮

不对。
模式匹配可以匹配类型、属性、条件、null 等，比传统 switch 强很多。

---

## 29. 本课重点总结

### 你必须先掌握的 10 个核心点

1. `string` 和 `string?` 语义不同
2. `?.` 和 `??` 是高频空安全工具
3. 值类型复制值本身，引用类型复制引用
4. `int?` 是可空值类型
5. `record` 适合表示数据对象
6. `record` 默认更强调值相等
7. 模式匹配能做类型、null、属性等判断
8. `switch expression` 比传统 `switch` 更简洁
9. `var` 是静态类型推断
10. `dynamic` 是运行时绑定，谨慎使用

---

## 30. Java / Go / Python 对比总结

| 知识点 | C# | Java | Go | Python |
|---|---|---|---|---|
| null 安全 | 可空引用类型 | 较弱，靠规范/工具 | 指针 nil 语义不同 | 运行时报错为主 |
| 数据类 | `record` | `record`（现代 Java） | struct | dataclass |
| 模式匹配 | 很强 | 近年增强但生态习惯不同 | 较弱 | `match`（3.10+） |
| 类型推断 | `var` | `var` | `:=` | 动态 |
| 顶层父类 | `object` | `Object` | `any/interface{}` | 一切皆对象风格 |
| 动态能力 | `dynamic` | 很弱 | 反射/接口 | 默认动态 |

---

## 31. 面试 / 实战高频提醒

- 看到 `?` 要先想“是否可能为 null”
- 可空安全优先靠类型和判空，不要靠运气
- 业务 DTO 很适合用 `record`
- 需要生成结果值时优先考虑 `switch expression`
- `var` 提升简洁性，但别滥用到降低可读性
- `dynamic` 只在确有必要时使用

---

## 32. 本课小练习

### 练习 1：可空引用类型

要求：

- 定义一个 `string? name`
- 如果为 `null`，输出 `unknown`
- 如果不为 `null`，输出长度

### 练习 2：值类型与引用类型

要求：

- 用 `int` 写一个赋值拷贝示例
- 再用 `class Person` 写一个引用拷贝示例
- 观察输出差异

### 练习 3：record

要求：

- 定义一个 `record User(string Name, int Age)`
- 创建两个内容相同的对象
- 比较它们是否相等

### 练习 4：with 表达式

要求：

- 基于已有 `User` 对象创建一个年龄不同的新对象
- 输出两个对象内容

### 练习 5：switch expression

要求：

- 根据分数返回等级：90+ A，80+ B，其他 C
- 用 `switch expression` 实现

### 练习 6：var / object / dynamic

要求：

- 分别定义这三种变量
- 打印并观察代码差异
- 思考哪种最安全、哪种最灵活

---

## 33. 一页速背版

### 可空引用类型

```csharp
string name = "Tom";
string? nickName = null;
```

### 空安全

```csharp
name?.Length;
string result = name ?? "default";
```

### 可空值类型

```csharp
int? age = null;
int value = age ?? 0;
```

### record

```csharp
public record User(string Name, int Age);
var u2 = u1 with { Age = 21 };
```

### 模式匹配

```csharp
if (obj is string s)
{
    Console.WriteLine(s.Length);
}
```

### switch expression

```csharp
string level = score switch
{
    >= 90 => "A",
    >= 80 => "B",
    _ => "C"
};
```

### var / object / dynamic

```csharp
var a = 123;
object b = 123;
dynamic c = 123;
```

### 核心结论

- `string?` 表示可能为空
- `?.` / `??` 高频
- `record` 适合数据对象
- 模式匹配和 `switch expression` 很现代、很好用
- `var` 是静态类型推断
- `dynamic` 谨慎使用

---

## 34. 下一课预告

下一课学习：

- 面向接口编程
- 继承与多态
- 抽象类
- 接口 `interface`
- 依赖注入基础
- C# 常见设计习惯

这一课会把你从“会写语法”推进到“会理解 C# 工程设计”。
