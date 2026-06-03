# 第七课：C# 继承、多态、抽象类、接口、依赖注入、面向接口编程

> 目标：掌握 C# 面向对象设计里最重要的一组能力，理解为什么现代 .NET 项目更偏向“面向接口编程”而不是“到处 new 具体类”

---

## 1. 本课目标

本课学完后，你应该能掌握：

- 什么是继承
- 什么是多态
- `virtual`、`override`、`base` 的用法
- 抽象类 `abstract class` 的作用
- 接口 `interface` 的作用
- 抽象类和接口怎么选
- 什么是依赖注入（DI）
- 为什么现代 C# 项目强调面向接口编程

---

## 2. 为什么这一课重要

如果前面几课让你“会写 C#”，
这一课会让你开始“理解 C# 项目为什么这么组织”。

你后面学习 ASP.NET Core、EF Core、测试、架构设计时，会反复遇到：

- `interface`
- `override`
- `abstract`
- `services.AddScoped<...>()`
- 构造函数注入

所以这一课是工程化的分水岭。

---

## 3. 继承是什么

继承可以理解为：

> 子类复用父类已有的属性和方法，并在此基础上扩展自己的能力。

### 示例

```csharp
class Animal
{
    public string Name { get; set; } = "";

    public void Eat()
    {
        Console.WriteLine($"{Name} is eating");
    }
}

class Dog : Animal
{
    public void Bark()
    {
        Console.WriteLine($"{Name} is barking");
    }
}
```

使用：

```csharp
Dog dog = new Dog();
dog.Name = "Buddy";
dog.Eat();
dog.Bark();
```

### 说明

- `Dog : Animal` 表示 `Dog` 继承 `Animal`
- 子类可以直接使用父类的公开成员

### 对比

- **Java**：语法和理解几乎一样
- **Go**：没有传统 class 继承，更偏组合
- **Python**：也支持继承，但类型系统约束更弱

### 速记

- `:` 表示继承
- 子类能复用父类成员
- 继承表示一种 “is-a” 关系

---

## 4. 继承的意义

继承的核心价值：

- 复用公共逻辑
- 建立层次关系
- 支持多态

比如：

- `Animal` -> `Dog` / `Cat`
- `Shape` -> `Circle` / `Rectangle`
- `Employee` -> `Manager` / `Developer`

### 速记

- 有明确父子层次关系时才考虑继承
- 不是所有“代码复用”都该用继承

---

## 5. virtual 和 override

如果你希望子类可以重写父类方法，需要用：

- 父类：`virtual`
- 子类：`override`

### 示例

```csharp
class Animal
{
    public virtual void Speak()
    {
        Console.WriteLine("Animal speaks");
    }
}

class Dog : Animal
{
    public override void Speak()
    {
        Console.WriteLine("Dog barks");
    }
}
```

使用：

```csharp
Animal animal = new Dog();
animal.Speak();
```

输出：

```csharp
Dog barks
```

### 说明

- `virtual` 表示“允许子类重写”
- `override` 表示“这里重写父类实现”

### 速记

- 想让子类改行为：父类方法加 `virtual`
- 子类改写时用 `override`

---

## 6. 多态是什么

多态可以理解为：

> 同一个父类引用，指向不同子类对象时，调用相同方法会表现出不同行为。

### 示例

```csharp
Animal a1 = new Dog();
Animal a2 = new Cat();

a1.Speak();
a2.Speak();
```

如果 `Dog` 和 `Cat` 都重写了 `Speak()`，那么结果会不同。

### 说明

这就是多态的核心价值：

- 写调用方代码时依赖抽象
- 运行时再决定具体行为

### 对比

- 和 Java 多态几乎一致
- Go 更偏接口多态

### 速记

- 多态 = 同一接口，不同实现
- 面向父类/接口编程，运行时表现不同

---

## 7. base 关键字

`base` 用于访问父类成员。

### 示例 1：调用父类方法

```csharp
class Animal
{
    public virtual void Speak()
    {
        Console.WriteLine("Animal speaks");
    }
}

class Dog : Animal
{
    public override void Speak()
    {
        base.Speak();
        Console.WriteLine("Dog barks");
    }
}
```

### 示例 2：调用父类构造函数

```csharp
class Animal
{
    public string Name { get; set; }

    public Animal(string name)
    {
        Name = name;
    }
}

class Dog : Animal
{
    public Dog(string name) : base(name)
    {
    }
}
```

### 速记

- `base.xxx`：访问父类成员
- `: base(...)`：调用父类构造函数

---

## 8. sealed

如果你不希望一个类被继承，或不希望某个重写继续被重写，可以用 `sealed`。

### 示例

```csharp
sealed class FinalClass
{
}
```

或：

```csharp
class Dog : Animal
{
    public sealed override void Speak()
    {
        Console.WriteLine("Dog barks");
    }
}
```

### 速记

- `sealed class`：禁止继续继承
- `sealed override`：禁止继续重写

---

## 9. 抽象类 abstract class

抽象类可以理解为：

> 不能直接实例化的“半成品父类”，用于定义通用结构和部分默认实现。

### 示例

```csharp
abstract class Animal
{
    public string Name { get; set; } = "";

    public abstract void Speak();

    public void Eat()
    {
        Console.WriteLine($"{Name} is eating");
    }
}
```

子类：

```csharp
class Dog : Animal
{
    public override void Speak()
    {
        Console.WriteLine("Dog barks");
    }
}
```

### 说明

- `abstract` 类不能 `new`
- 抽象方法没有方法体
- 子类必须实现抽象方法
- 抽象类也可以包含普通方法

### 速记

- 抽象类不能直接实例化
- 抽象方法必须由子类实现
- 抽象类适合“有共性实现 + 有待子类定制”的场景

---

## 10. 接口 interface

接口可以理解为：

> 一组能力约定，只描述“能做什么”，不强调“怎么做”。

### 示例

```csharp
interface ILogger
{
    void Log(string message);
}
```

实现类：

```csharp
class ConsoleLogger : ILogger
{
    public void Log(string message)
    {
        Console.WriteLine($"[LOG] {message}");
    }
}
```

使用：

```csharp
ILogger logger = new ConsoleLogger();
logger.Log("hello");
```

### 说明

- 接口定义规范
- 实现类负责具体实现
- 调用方依赖接口，不依赖具体类

### 对比

- 和 Java interface 很像
- Go 的接口概念更轻量，但思想接近

### 速记

- 接口定义能力，不写具体业务细节
- 变量类型优先写接口，不优先写具体类

---

## 11. 接口的意义

接口的核心价值：

- 解耦
- 更容易替换实现
- 更容易测试
- 更适合扩展

例如：

- `ILogger`
  - `ConsoleLogger`
  - `FileLogger`
  - `DbLogger`

调用方不用关心具体是哪一种 logger。

### 速记

- 接口是解耦的核心工具
- 面向接口编程能降低模块耦合度

---

## 12. 一个类可以实现多个接口

### 示例

```csharp
interface ILogger
{
    void Log(string message);
}

interface ISavable
{
    void Save();
}

class Document : ILogger, ISavable
{
    public void Log(string message)
    {
        Console.WriteLine(message);
    }

    public void Save()
    {
        Console.WriteLine("saved");
    }
}
```

### 说明

- C# 类只支持单继承
- 但支持实现多个接口

### 速记

- 类：单继承
- 接口：可多实现

---

## 13. 抽象类 vs 接口怎么选

### 更适合抽象类的场景

- 多个子类确实有明显共同父类
- 需要共享部分实现
- 需要共享字段或公共状态

### 更适合接口的场景

- 更关注“能力约定”
- 需要解耦
- 需要多实现切换
- 更适合依赖注入

### 一个直觉理解

- 抽象类：一种“是什么”
- 接口：一种“能做什么”

### 速记

- 有共享实现：优先考虑抽象类
- 要解耦、可替换：优先考虑接口

---

## 14. 面向接口编程是什么

你可以把它理解为：

> 调用方依赖抽象接口，而不是直接依赖具体实现类。

### 不推荐写法

```csharp
class UserService
{
    private readonly ConsoleLogger _logger = new ConsoleLogger();
}
```

问题：

- 强耦合
- 不方便替换
- 不方便测试

### 更推荐写法

```csharp
class UserService
{
    private readonly ILogger _logger;

    public UserService(ILogger logger)
    {
        _logger = logger;
    }
}
```

### 说明

这样 `UserService` 不关心传进来的是：

- `ConsoleLogger`
- `FileLogger`
- `MockLogger`

### 速记

- 依赖接口，不依赖具体类
- 这样代码更容易替换、扩展、测试

---

## 15. 依赖注入 DI 是什么

依赖注入可以理解为：

> 一个对象需要什么依赖，不自己 new，而是由外部传进来。

### 示例

```csharp
class UserService
{
    private readonly ILogger _logger;

    public UserService(ILogger logger)
    {
        _logger = logger;
    }

    public void CreateUser(string name)
    {
        _logger.Log($"create user: {name}");
    }
}
```

调用：

```csharp
ILogger logger = new ConsoleLogger();
UserService service = new UserService(logger);
service.CreateUser("Tom");
```

### 说明

- `UserService` 并没有自己去 `new ConsoleLogger()`
- 这就是最基础的 DI 思想

### 速记

- 不自己创建依赖
- 由外部注入依赖
- 最常见方式是构造函数注入

---

## 16. 构造函数注入

这是最常用、最推荐的依赖注入方式。

### 示例

```csharp
class OrderService
{
    private readonly ILogger _logger;
    private readonly IOrderRepository _repository;

    public OrderService(ILogger logger, IOrderRepository repository)
    {
        _logger = logger;
        _repository = repository;
    }
}
```

### 说明

优点：

- 依赖关系清晰
- 创建对象时必须提供完整依赖
- 更方便测试

### 速记

- DI 优先构造函数注入
- 比属性注入、方法注入更清晰

---

## 17. 为什么 DI 对测试友好

因为你可以很容易替换成假的实现。

### 示例

```csharp
class FakeLogger : ILogger
{
    public void Log(string message)
    {
        Console.WriteLine($"fake: {message}");
    }
}
```

测试时：

```csharp
ILogger logger = new FakeLogger();
UserService service = new UserService(logger);
```

### 说明

如果你在类里写死：

```csharp
new ConsoleLogger()
```

那测试替换就麻烦很多。

### 速记

- 接口 + DI = 更易测试
- 这也是企业项目非常重视它们的原因

---

## 18. ASP.NET Core 里的 DI 直觉理解

你以后会经常看到：

```csharp
builder.Services.AddScoped<IUserService, UserService>();
```

这表示：

- 当程序需要 `IUserService` 时
- 容器就给它一个 `UserService`

你现在不用死记容器细节，先理解思想：

- 注册接口与实现的映射
- 用的时候自动注入

### 速记

- ASP.NET Core 默认深度使用 DI
- 你现在先掌握“接口 + 构造函数注入”就够了

---

## 19. override 和 new 的区别

### 示例

```csharp
class Parent
{
    public virtual void Show()
    {
        Console.WriteLine("Parent");
    }
}

class Child : Parent
{
    public override void Show()
    {
        Console.WriteLine("Child");
    }
}
```

这是重写。

而：

```csharp
class Parent
{
    public void Show()
    {
        Console.WriteLine("Parent");
    }
}

class Child : Parent
{
    public new void Show()
    {
        Console.WriteLine("Child");
    }
}
```

这是隐藏，不是真正的多态重写。

### 速记

- `override`：真正多态
- `new`：隐藏父类成员
- 日常开发里优先理解和使用 `override`

---

## 20. is / as 与多态判断

### 示例

```csharp
Animal animal = new Dog();

if (animal is Dog dog)
{
    dog.Bark();
}
```

### 说明

- `is` 可以做类型判断并转换
- `as` 也可以做引用类型安全转换，但现代代码更常见 `is` 模式匹配

### 速记

- 类型判断优先 `is Xxx x`
- 可读性比老式强转更好

---

## 21. 组合优于继承

这是你后面做工程设计时非常重要的一条经验。

### 说明

不是所有复用都要靠继承。
很多时候：

- 继承层次太深会变复杂
- 组合会更灵活

### 示例思路

与其让：

- `SuperDog : Dog : Animal : ...`

不如让某个对象“拥有一个能力组件”。

### 速记

- 有明确 is-a 关系才考虑继承
- 否则优先考虑接口 + 组合

---

## 22. 常见误区

### 误区 1：为了复用代码就应该继承

不一定。
继承是强耦合关系，不是万能复用手段。

### 误区 2：接口只是为了“规范好看”

不对。
接口直接影响解耦、测试、扩展能力。

### 误区 3：抽象类和接口完全一样

不一样。
抽象类可以带实现和状态，接口更强调能力约定。

### 误区 4：DI 很玄学，只是框架强行搞复杂

不是。
它本质上只是“依赖别自己 new，让外部传进来”。

### 误区 5：`new` 和 `override` 差不多

不对。
`override` 是多态行为，`new` 是成员隐藏。

---

## 23. 本课重点总结

### 你必须先掌握的 10 个核心点

1. 继承表示 is-a 关系
2. 多态让同一抽象有不同实现表现
3. 父类可重写方法用 `virtual`
4. 子类重写用 `override`
5. `base` 用于访问父类成员或构造函数
6. 抽象类适合“有共性实现”的父类建模
7. 接口适合“能力约定”和解耦
8. 类单继承，但可实现多个接口
9. 依赖注入的核心是不自己 new 依赖
10. 面向接口编程是现代 .NET 工程开发核心习惯

---

## 24. Java / Go / Python 对比总结

| 知识点 | C# | Java | Go | Python |
|---|---|---|---|---|
| 继承 | 单继承 | 单继承 | 无传统 class 继承 | 支持 |
| 接口 | 很常用 | 很常用 | 更核心、更轻量 | 协议/鸭子类型更常见 |
| 抽象类 | `abstract class` | `abstract class` | 无直接对应 | `abc` 模块 |
| 多态 | `virtual/override` + 接口 | 类似 | 主要靠接口 | 动态多态 |
| 依赖注入 | 框架默认支持强 | 常靠 Spring 等框架 | 手动注入更多 | 框架各异 |

---

## 25. 面试 / 实战高频提醒

- 父类方法不加 `virtual`，子类不能正常 `override`
- 接口是 DI 和解耦的基础
- 构造函数注入是最推荐的 DI 方式
- 业务代码里尽量避免直接 `new` 底层依赖
- 继承层次不要过深
- 能力扩展很多时候优先考虑接口 + 组合

---

## 26. 本课小练习

### 练习 1：继承与重写

要求：

- 定义 `Animal` 类，包含 `Speak()`
- 定义 `Dog`、`Cat` 继承它
- 分别重写 `Speak()`

### 练习 2：抽象类

要求：

- 定义抽象类 `Shape`
- 写抽象方法 `GetArea()`
- 定义 `Circle` 和 `Rectangle` 实现它

### 练习 3：接口

要求：

- 定义接口 `ILogger`
- 写 `ConsoleLogger` 实现它
- 调用 `Log()` 输出日志

### 练习 4：面向接口编程

要求：

- 定义 `UserService`
- 通过构造函数注入 `ILogger`
- 在 `CreateUser()` 中输出日志

### 练习 5：多态

要求：

- 建一个 `List<Animal>`
- 放入 `Dog` 和 `Cat`
- 遍历调用 `Speak()`
- 观察不同输出

---

## 27. 一页速背版

### 继承

```csharp
class Dog : Animal
{
}
```

### 重写

```csharp
class Animal
{
    public virtual void Speak() { }
}

class Dog : Animal
{
    public override void Speak()
    {
        Console.WriteLine("Dog barks");
    }
}
```

### 抽象类

```csharp
abstract class Shape
{
    public abstract double GetArea();
}
```

### 接口

```csharp
interface ILogger
{
    void Log(string message);
}
```

### 构造函数注入

```csharp
class UserService
{
    private readonly ILogger _logger;

    public UserService(ILogger logger)
    {
        _logger = logger;
    }
}
```

### 核心结论

- 继承：父子关系
- 多态：同一抽象，不同行为
- 抽象类：共性 + 部分实现
- 接口：约定 + 解耦
- DI：依赖外部传入，不自己 new

---

## 28. 下一课预告

下一课学习：

- ASP.NET Core 入门
- Web API 基本结构
- 路由
- Controller
- 请求与响应
- DTO
- 依赖注入在 Web 项目里的使用

这课学完，你就开始真正进入 C# Web 开发了。
