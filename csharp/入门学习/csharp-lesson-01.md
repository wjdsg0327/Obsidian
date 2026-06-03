# 第一课：C# 基础语法快速入门

> 适合有 Java / Go / Python 基础的开发者快速入门 C#

---

## 1. 本课目标

本课学完后，你应该能掌握：

- C# 程序的基本结构
- 变量与常见基本类型
- `var` 的含义
- 字符串与字符串插值
- 条件判断与循环
- 数组与方法定义
- C# 与 Java / Go / Python 的基础差异

---

## 2. C# 程序最小结构

### 示例

```csharp
using System;

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Hello, World!");
    }
}
```

### 说明

- `using System;`
  - 导入命名空间
  - 类似 Java 的 `import`
- `class Program`
  - C# 代码通常写在类中
- `static void Main(string[] args)`
  - 程序入口
  - 类似 Java 的 `public static void main(String[] args)`
- `Console.WriteLine(...)`
  - 控制台输出
  - 类似 Java 的 `System.out.println(...)`

### 速记

- C# 基础结构和 Java 很像
- 程序入口通常是 `Main`
- 输出常用 `Console.WriteLine`

---

## 3. 变量与基本类型

### 示例

```csharp
int age = 18;
double price = 19.99;
char grade = 'A';
string name = "Tom";
bool isAdmin = true;
decimal salary = 9999.99m;
```

### 常见类型

| 类型 | 含义 | 示例 |
|---|---|---|
| `int` | 整数 | `int age = 18;` |
| `long` | 长整数 | `long count = 100000L;` |
| `float` | 单精度浮点 | `float x = 1.2f;` |
| `double` | 双精度浮点 | `double pi = 3.14;` |
| `decimal` | 高精度小数，常用于金额 | `decimal money = 99.9m;` |
| `char` | 单个字符 | `char c = 'A';` |
| `string` | 字符串 | `string s = "hello";` |
| `bool` | 布尔值 | `bool ok = true;` |

### 注意点

- `string` 在 C# 中非常常用
- 金额推荐使用 `decimal`
- `decimal` 字面量要加 `m` 后缀

### 对比

- **Java**：整体非常接近
- **Go**：Go 没有内建 `decimal`
- **Python**：Python 不强制声明类型，C# 是强类型语言

### 速记

- C# 是强类型语言
- 金额优先 `decimal`
- `string`、`int`、`bool` 是最高频类型

---

## 4. var：类型推断

### 示例

```csharp
var name = "Alice";
var count = 10;
var price = 19.9;
```

### 说明

`var` 表示让编译器自动推断类型，不是动态类型。

上面的实际类型分别是：

- `name` -> `string`
- `count` -> `int`
- `price` -> `double`

### 对比

- 类似 Java 的 `var`
- 类似 Go 的 `:=`
- **不像 Python**
  - Python 是运行时动态类型
  - C# 的 `var` 仍然是静态类型

### 速记

- `var` = 自动推断类型
- `var` ≠ 动态类型
- 写法更简洁，但本质还是强类型

---

## 5. 常量

### 示例

```csharp
const double PI = 3.14159;
```

### 说明

- `const` 表示编译期常量
- 一旦定义，不能修改

### 补充

后续还会接触：

```csharp
readonly int id;
```

- `readonly`：运行时只读，一般在构造函数中赋值
- `const` 和 `readonly` 的区别，后面会详细讲

### 速记

- `const`：编译期常量
- `readonly`：运行时只读
- 现阶段先记住常量用 `const`

---

## 6. 输入输出

### 输出

```csharp
Console.WriteLine("Hello");
Console.Write("Hi");
```

### 输入

```csharp
string? name = Console.ReadLine();
Console.WriteLine($"你好, {name}");
```

### 说明

- `Console.WriteLine()`：输出并换行
- `Console.Write()`：输出但不换行
- `Console.ReadLine()`：读取用户输入

### 注意点

- 新版本 C# 中，`Console.ReadLine()` 常按可空字符串处理：`string?`
- 以后会学到“可空引用类型”

### 速记

- 输出：`Console.WriteLine`
- 输入：`Console.ReadLine`
- 输入结果常看作 `string?`

---

## 7. 字符串与字符串插值

### 普通拼接

```csharp
string name = "Tom";
int age = 20;
Console.WriteLine("My name is " + name + ", age is " + age);
```

### 推荐：字符串插值

```csharp
string name = "Tom";
int age = 20;
Console.WriteLine($"My name is {name}, age is {age}");
```

### 说明

- `$"..."` 是 C# 很常用的语法
- 可读性比 `+` 拼接更好

### 对比

- **Java**：更像增强版字符串格式化
- **Python**：非常像 f-string
- **Go**：类似 `fmt.Printf`

### 速记

- 推荐用 `$""`
- 插值语法：`{变量名}`
- 日常开发中非常高频

---

## 8. 条件判断

### if / else

```csharp
int age = 20;

if (age >= 18)
{
    Console.WriteLine("成年人");
}
else
{
    Console.WriteLine("未成年");
}
```

### switch

```csharp
int day = 2;

switch (day)
{
    case 1:
        Console.WriteLine("Mon");
        break;
    case 2:
        Console.WriteLine("Tue");
        break;
    default:
        Console.WriteLine("Unknown");
        break;
}
```

### 说明

- `if / else` 和 Java 几乎一致
- `switch` 也是常规多分支判断
- 后面还会学更强的 `switch expression`

### 速记

- 条件判断和 Java 非常像
- `switch` 每个 `case` 常配 `break`
- 默认分支用 `default`

---

## 9. 循环

### for

```csharp
for (int i = 0; i < 5; i++)
{
    Console.WriteLine(i);
}
```

### while

```csharp
int i = 0;
while (i < 5)
{
    Console.WriteLine(i);
    i++;
}
```

### foreach

```csharp
string[] names = { "Tom", "Jack", "Lucy" };

foreach (string name in names)
{
    Console.WriteLine(name);
}
```

### 说明

- `for`：已知次数时常用
- `while`：条件循环
- `foreach`：遍历集合最方便

### 对比

- **Java**：类似增强 for
- **Go**：类似 `for range`
- **Python**：类似 `for x in list`

### 速记

- 遍历集合优先考虑 `foreach`
- 计数循环用 `for`
- 条件循环用 `while`

---

## 10. 数组

### 示例 1：直接初始化

```csharp
int[] nums = { 1, 2, 3 };
Console.WriteLine(nums[0]);
```

### 示例 2：指定长度

```csharp
int[] nums = new int[3];
nums[0] = 10;
nums[1] = 20;
nums[2] = 30;
```

### 说明

- 数组长度固定
- 下标从 0 开始
- 越界会抛异常

### 速记

- C# 数组写法：`int[] nums`
- 数组长度固定
- 遍历数组常配合 `foreach`

---

## 11. 方法

### 示例

```csharp
static int Add(int a, int b)
{
    return a + b;
}
```

调用：

```csharp
int result = Add(3, 5);
Console.WriteLine(result);
```

### 说明

- 方法必须声明返回类型
- 参数类型必须明确
- `return` 返回结果

### 对比

- 和 Java 很像
- 比 Python 更严格
- 比 Go 更偏面向对象风格

### 速记

- 方法格式：`返回类型 方法名(参数)`
- 有返回值就写具体类型
- 没返回值用 `void`

---

## 12. 可选参数

### 示例

```csharp
static void Greet(string name = "Guest")
{
    Console.WriteLine($"Hello, {name}");
}
```

调用：

```csharp
Greet();
Greet("Tom");
```

### 说明

- 参数可以有默认值
- 调用时可不传

### 对比

- Java 原生不如 C# 直接
- Python 里也常见默认参数

### 速记

- 默认参数很实用
- 语法：`类型 参数名 = 默认值`

---

## 13. 命名规范

### 常见规范

- 类名、方法名、属性名：`PascalCase`
  - `StudentInfo`
  - `GetUserName`
- 变量名、参数名：`camelCase`
  - `userName`
  - `totalCount`

### 速记

- 大部分场景和 Java 命名习惯接近
- 类/方法/属性首字母大写
- 局部变量首字母小写

---

## 14. 本课重点总结

### 你必须先掌握的 5 个核心点

1. C# 基础结构和 Java 非常像
2. `var` 是类型推断，不是动态类型
3. 字符串推荐使用 `$""` 插值
4. 遍历集合优先使用 `foreach`
5. 方法、数组、条件判断都不难，重点是熟悉语法写法

---

## 15. Java / Go / Python 对比总结

| 知识点 | C# | Java | Go | Python |
|---|---|---|---|---|
| 程序入口 | `Main` | `main` | `main` | 不强制 |
| 类型系统 | 强类型 | 强类型 | 强类型 | 动态类型 |
| 类型推断 | `var` | `var` | `:=` | 天然动态 |
| 字符串插值 | `$""` | 较弱 | `fmt.Printf` | f-string |
| 遍历集合 | `foreach` | 增强 for | `range` | `for in` |
| OOP 风格 | 强 | 强 | 弱一些 | 灵活 |

---

## 16. 面试/实战高频提醒

- `var` 不是动态类型
- `decimal` 常用于金额
- `Console.ReadLine()` 可能是 `null`
- 字符串推荐用插值，而不是反复 `+`
- 数组长度固定，后面要学更常用的 `List<T>`

---

## 17. 本课小练习

### 练习 1：定义变量并打印

要求：

- 定义姓名、年龄、工资
- 使用字符串插值输出

参考目标：

```csharp
string name = "Alice";
int age = 25;
decimal salary = 12000.50m;

Console.WriteLine($"name={name}, age={age}, salary={salary}");
```

### 练习 2：写一个加法方法

要求：

- 写 `Add(int a, int b)`
- 返回两数之和

### 练习 3：遍历数组

要求：

- 定义一个字符串数组
- 用 `foreach` 输出每个元素

---

## 18. 一页速背版

### 语法骨架

```csharp
using System;

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Hello, World!");
    }
}
```

### 常用类型

```csharp
int
double
decimal
bool
char
string
```

### 常用写法

```csharp
var name = "Tom";
const double PI = 3.14;
Console.WriteLine($"name={name}");
```

### 常用流程控制

```csharp
if / else
switch
for
while
foreach
```

### 常用结论

- `var` 是类型推断
- `$""` 是字符串插值
- 金额优先 `decimal`
- 遍历集合优先 `foreach`

---

## 19. 下一课预告

下一课学习：

- 类和对象
- 字段和属性
- `get; set;`
- 构造函数
- `this`
- `class` 和 `struct`
- `static`

这是 C# 真正和 Java 拉开细节差异的地方。