# 第二课：C# 面向对象入门——类、对象、属性、构造函数

> 目标：快速掌握 C# 面向对象核心语法，并理解它和 Java / Go / Python 的关键差异

---

## 1. 本课目标

本课学完后，你应该能掌握：

- 什么是类和对象
- 字段、属性的区别
- `get; set;` 的用法
- 构造函数怎么写
- `this` 的作用
- `static` 成员的含义
- `class` 和 `struct` 的区别
- 访问修饰符的基础用法

---

## 2. 类和对象

### 示例

```csharp
class Person
{
    public string Name;
    public int Age;

    public void SayHello()
    {
        Console.WriteLine($"Hello, I'm {Name}, age {Age}");
    }
}
```

创建对象：

```csharp
Person p = new Person();
p.Name = "Tom";
p.Age = 20;
p.SayHello();
```

### 说明

- `class` 用来定义类
- `new` 用来创建对象
- 对象通过 `.` 访问字段和方法

### 对比

- **Java**：几乎一样
- **Go**：Go 也有结构体和方法，但不是典型 class 风格 OOP
- **Python**：类更灵活，但不如 C# 严格

### 速记

- `class` 定义模板
- `new` 创建实例
- 对象用 `.` 调属性和方法

---

## 3. 字段 Field

### 示例

```csharp
class Person
{
    public string Name;
    public int Age;
}
```

### 说明

- `Name`、`Age` 是字段
- 字段就是类里定义的变量
- 可以直接读写

### 问题

字段虽然简单，但直接暴露通常不够规范，因为：

- 不方便控制赋值逻辑
- 不方便校验数据
- 不利于后续扩展

所以 C# 更常用 **属性 Property**。

### 速记

- 字段 = 类中的变量
- 能直接存数据
- 实战中更推荐属性而不是公开字段

---

## 4. 属性 Property

### 示例

```csharp
class Person
{
    public string Name { get; set; }
    public int Age { get; set; }
}
```

使用：

```csharp
Person p = new Person();
p.Name = "Alice";
p.Age = 25;
Console.WriteLine(p.Name);
```

### 说明

- 属性是 C# 非常核心的语法
- 表面上看像字段，实际上底层是方法形式的封装
- `get` 表示读取
- `set` 表示写入

### 为什么属性重要

因为属性兼顾：

- 写起来像字段
- 本质上可以加入逻辑控制
- 比公开字段更符合封装思想

### 对比

- **Java**：通常写 `private field + getter/setter`
- **C#**：常直接写 `public string Name { get; set; }`
- **Python**：有 `@property`，但使用习惯不同

### 速记

- C# 中优先用属性，不优先用公开字段
- `get; set;` 是最常见写法
- 属性看起来像字段，用起来更安全、更规范

---

## 5. 自动属性与完整属性

### 自动属性

```csharp
public string Name { get; set; }
```

适合：只是存值，没有额外逻辑。

### 完整属性

```csharp
private int age;

public int Age
{
    get { return age; }
    set
    {
        if (value >= 0)
        {
            age = value;
        }
    }
}
```

### 说明

完整属性适合在赋值时做校验。

### 对比

这很像 Java 手写 getter/setter，只是 C# 语法更紧凑。

### 速记

- 简单存值：自动属性
- 要做校验：完整属性
- `value` 表示 setter 传进来的值

---

## 6. 只读属性

### 示例

```csharp
class Person
{
    public string Name { get; }

    public Person(string name)
    {
        Name = name;
    }
}
```

### 说明

- 只有 `get`，没有 `set`
- 一般在构造函数中赋值
- 适合不希望对象创建后被修改的数据

### 速记

- `get;` 表示外部可读
- 没有 `set;` 就不能随便改
- 常用于不可变对象设计

---

## 7. 构造函数

### 示例

```csharp
class Person
{
    public string Name { get; set; }
    public int Age { get; set; }

    public Person(string name, int age)
    {
        Name = name;
        Age = age;
    }
}
```

使用：

```csharp
Person p = new Person("Tom", 20);
```

### 说明

- 构造函数名称和类名相同
- 没有返回值类型
- 创建对象时自动调用

### 对比

- **Java**：完全类似
- **Python**：类似 `__init__`

### 速记

- 构造函数用于初始化对象
- 名字必须和类名一致
- 没有返回值类型

---

## 8. this 关键字

### 示例

```csharp
class Person
{
    public string Name { get; set; }
    public int Age { get; set; }

    public Person(string name, int age)
    {
        this.Name = name;
        this.Age = age;
    }
}
```

### 说明

- `this` 代表当前对象
- 常用于区分“成员变量/属性”和“参数”同名的情况

### 对比

- 和 Java 里的 `this` 基本一致
- Python 里对应 `self`

### 速记

- `this` = 当前对象
- 参数和成员同名时很常用
- 和 Java 理解方式一样

---

## 9. 方法重载

### 示例

```csharp
class Calculator
{
    public int Add(int a, int b)
    {
        return a + b;
    }

    public double Add(double a, double b)
    {
        return a + b;
    }
}
```

### 说明

- 方法名相同
- 参数列表不同
- 返回类型不同 **不能单独构成重载**

### 速记

- 重载看参数，不看返回值
- 方法名一样没问题
- 参数类型/个数不同才行

---

## 10. static 静态成员

### 示例

```csharp
class MathHelper
{
    public static int Add(int a, int b)
    {
        return a + b;
    }
}
```

调用：

```csharp
int result = MathHelper.Add(3, 5);
```

### 说明

- `static` 成员属于类，不属于某个对象
- 不需要 `new` 就能调用

### 静态字段示例

```csharp
class Counter
{
    public static int Total = 0;
}
```

### 对比

- 和 Java 的 `static` 很像
- Python 没有完全等价的语言级写法

### 速记

- `static` 属于类
- 调用方式：`类名.成员名`
- 工具方法经常写成静态方法

---

## 11. 访问修饰符

### 常见修饰符

| 修饰符 | 含义 |
|---|---|
| `public` | 任何地方可访问 |
| `private` | 只能在当前类内部访问 |
| `protected` | 当前类和子类可访问 |
| `internal` | 当前程序集内可访问 |

### 示例

```csharp
class Person
{
    private int age;
    public string Name { get; set; }
}
```

### 说明

- 默认推荐：字段 `private`
- 对外暴露：属性或方法 `public`

### 速记

- 封装的常见组合：`private field + public property`
- 实战中 `public` 和 `private` 最常见

---

## 12. 封装：推荐写法

### 示例

```csharp
class Person
{
    private int age;

    public string Name { get; set; }

    public int Age
    {
        get { return age; }
        set
        {
            if (value >= 0)
            {
                age = value;
            }
        }
    }
}
```

### 说明

这就是封装：

- 隐藏内部数据
- 对外通过属性控制访问
- 保证数据合法性

### 对比

- Java 也强调封装
- C# 用属性实现会更自然、更简洁

### 速记

- 内部字段尽量别裸露
- 对外暴露属性
- 有规则的数据要加校验

---

## 13. class 和 struct 的区别

### class

```csharp
class Person
{
    public string Name { get; set; }
}
```

### struct

```csharp
struct Point
{
    public int X { get; set; }
    public int Y { get; set; }
}
```

### 核心区别

- `class` 是 **引用类型**
- `struct` 是 **值类型**

### 示例理解

```csharp
class Person
{
    public string Name { get; set; }
}

Person p1 = new Person { Name = "Tom" };
Person p2 = p1;
p2.Name = "Jack";

Console.WriteLine(p1.Name); // Jack
```

因为 `class` 复制的是引用。

```csharp
struct Point
{
    public int X { get; set; }
}

Point p1 = new Point { X = 10 };
Point p2 = p1;
p2.X = 20;

Console.WriteLine(p1.X); // 10
```

因为 `struct` 复制的是值。

### 什么时候用 struct

适合：

- 数据小
- 表示值语义
- 不需要继承
- 比如坐标、日期片段、小型数据对象

### 速记

- `class`：引用类型
- `struct`：值类型
- 大多数业务对象优先 `class`
- 小型值对象可以考虑 `struct`

---

## 14. 对象初始化器

### 示例

```csharp
Person p = new Person
{
    Name = "Tom",
    Age = 20
};
```

### 说明

这在 C# 里很常见，前提是属性可写。

### 对比

- 比 Java 传统写法更简洁
- 类似一些框架里的 builder 直观赋值风格

### 速记

- `new 对象 { 属性 = 值 }`
- 代码更简洁
- 配合自动属性很好用

---

## 15. null 基础认知

### 示例

```csharp
Person? p = null;
```

### 说明

- 引用类型可以为 `null`
- 新版 C# 引入了可空引用类型概念
- `Person` 和 `Person?` 在语义上不同

这一块后面会单独展开，现在先知道：

- 访问对象前要考虑是否为空

### 速记

- 引用类型可能为 `null`
- 新版 C# 更强调可空安全
- 看到 `?` 要想到“可能为空”

---

## 16. 本课重点总结

### 你必须先掌握的 7 个核心点

1. `class` 用来定义对象模板
2. `new` 用来创建实例
3. C# 更推荐属性，不推荐公开字段
4. `get; set;` 是最常见属性写法
5. 构造函数用于初始化对象
6. `static` 属于类，不属于实例
7. `class` 是引用类型，`struct` 是值类型

---

## 17. Java / Go / Python 对比总结

| 知识点 | C# | Java | Go | Python |
|---|---|---|---|---|
| 类 | `class` | `class` | 无传统 class OOP | `class` |
| 属性 | `get; set;` | getter/setter 方法 | 无属性语法糖 | `@property` |
| 构造函数 | 类名同名 | 类名同名 | 通常工厂函数 | `__init__` |
| this/self | `this` | `this` | 无固定 this 用法 | `self` |
| 静态成员 | `static` | `static` | 包级函数更多 | 类变量/静态方法 |
| 值类型 | `struct` | 基本类型为值，类为引用 | struct 常见 | 一般对象语义 |

---

## 18. 面试 / 实战高频提醒

- C# 中 **属性** 比公开字段更常用
- `get; set;` 是必须熟悉的高频语法
- `class` 是引用类型，赋值时复制引用
- `struct` 是值类型，赋值时复制值
- `static` 方法不需要实例化对象
- 参数和成员同名时常用 `this`

---

## 19. 本课小练习

### 练习 1：定义一个 Student 类

要求：

- 有 `Name` 和 `Age` 两个属性
- 有一个 `SayHello()` 方法
- 输出自己的名字和年龄

### 练习 2：写构造函数

要求：

- 在 `Student` 类中写构造函数
- 创建对象时传入姓名和年龄

### 练习 3：给 Age 加校验

要求：

- 年龄不能小于 0
- 用完整属性或字段 + 属性实现

### 练习 4：写一个静态工具类

要求：

- 创建 `MathHelper`
- 写一个静态方法 `Add(int a, int b)`

---

## 20. 一页速背版

### 类和对象

```csharp
class Person
{
    public string Name { get; set; }
    public int Age { get; set; }
}

Person p = new Person();
```

### 构造函数

```csharp
public Person(string name, int age)
{
    this.Name = name;
    this.Age = age;
}
```

### 静态方法

```csharp
class MathHelper
{
    public static int Add(int a, int b)
    {
        return a + b;
    }
}
```

### 完整属性

```csharp
private int age;

public int Age
{
    get { return age; }
    set
    {
        if (value >= 0)
        {
            age = value;
        }
    }
}
```

### 核心结论

- 属性比公开字段更常用
- `get; set;` 是高频语法
- `static` 属于类
- `class` 是引用类型
- `struct` 是值类型

---

## 21. 下一课预告

下一课学习：

- `List<T>`、`Dictionary<TKey, TValue>`
- 泛型
- 集合遍历
- Lambda 表达式
- LINQ 基础

这是 C# 真正提升开发效率的关键部分。
