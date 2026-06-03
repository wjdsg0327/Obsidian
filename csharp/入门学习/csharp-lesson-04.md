# 第四课：C# 异常处理、委托、Action/Func、事件、async/await

> 目标：掌握 C# 工程开发里最常见的一组高级基础能力，理解“把方法当值传递”和“异步编程”的核心思想

---

## 1. 本课目标

本课学完后，你应该能掌握：

- C# 异常处理的基本写法
- `try / catch / finally` 的使用场景
- 什么是委托 `delegate`
- `Action` 和 `Func` 的区别
- 什么是事件 `event`
- `async/await` 的基本写法
- `Task` 的基本概念
- 这几个知识点和 Java / Go / Python 的差异

---

## 2. 为什么这一课很重要

前面三课让你能写出“能跑”的代码。
这一课会让你开始理解：

- 如何优雅处理错误
- 如何把逻辑当参数传递
- 如何做异步编程
- 为什么 C# 很适合大型工程开发

C# 的很多现代写法，本质都和下面几件事强相关：

- 委托
- Lambda
- 事件
- Task
- async/await

---

## 3. 异常处理基础

### 示例

```csharp
try
{
    int x = 10;
    int y = 0;
    int result = x / y;
    Console.WriteLine(result);
}
catch (DivideByZeroException ex)
{
    Console.WriteLine($"发生异常: {ex.Message}");
}
finally
{
    Console.WriteLine("finally 总会执行");
}
```

### 说明

- `try`：放可能出错的代码
- `catch`：捕获异常
- `finally`：无论是否异常都会执行，常用于资源释放

### 对比

- **Java**：几乎一样
- **Python**：类似 `try / except / finally`
- **Go**：Go 更常用 `error` 返回值，而不是异常机制

### 速记

- 可能出错的代码放 `try`
- 异常处理写 `catch`
- 收尾逻辑放 `finally`

---

## 4. 捕获多个异常

### 示例

```csharp
try
{
    string? s = null;
    Console.WriteLine(s.Length);
}
catch (NullReferenceException ex)
{
    Console.WriteLine($"空引用异常: {ex.Message}");
}
catch (Exception ex)
{
    Console.WriteLine($"其他异常: {ex.Message}");
}
```

### 说明

- 可以写多个 `catch`
- 一般先写更具体的异常，再写通用的 `Exception`

### 速记

- `catch` 可以有多个
- 先具体，后通用
- `Exception` 相当于兜底异常

---

## 5. throw：主动抛异常

### 示例

```csharp
static void CheckAge(int age)
{
    if (age < 0)
    {
        throw new ArgumentException("age 不能小于 0");
    }
}
```

### 说明

- `throw` 用来主动抛出异常
- 常用于参数校验、业务规则校验

### 常见异常类型

- `ArgumentException`
- `ArgumentNullException`
- `InvalidOperationException`
- `NotImplementedException`

### 速记

- 参数非法常想到 `ArgumentException`
- `throw` 用于主动中断流程并报告错误

---

## 6. 什么时候该用异常

### 适合用异常的场景

- 真正的异常情况
- 非法输入
- 外部资源调用失败
- 无法继续执行的错误

### 不适合滥用异常的场景

- 正常业务分支判断
- 高频控制流程

比如：

- 查字典 key 是否存在，不要靠异常，优先 `TryGetValue`
- 字符串转数字，不要靠异常，优先 `TryParse`

### 速记

- 异常处理“异常情况”
- 不要拿异常代替普通 if 判断

---

## 7. 委托 delegate 是什么

你可以把委托理解为：

> “可以指向方法的类型”。

也就是：

- 变量不只可以存数据
- 还可以“存一个方法”

### 示例

```csharp
delegate int Calc(int a, int b);
```

这表示定义了一个委托类型 `Calc`，它可以指向：

- 接收两个 `int`
- 返回一个 `int`

这样的方法。

### 示例方法

```csharp
static int Add(int a, int b)
{
    return a + b;
}
```

绑定委托：

```csharp
Calc calc = Add;
int result = calc(3, 5);
Console.WriteLine(result);
```

### 对比

- **Java**：可类比函数式接口
- **Python**：函数本身就是一等公民，天然可以传递
- **Go**：函数也可以作为值传递

### 速记

- 委托 = 方法签名的类型
- 可以把方法赋值给委托变量
- 可以像调用方法一样调用委托

---

## 8. 委托的意义

委托的价值在于：

- 让方法可以作为参数传递
- 让逻辑更灵活
- 是 Lambda、事件、很多框架机制的基础

### 示例

```csharp
static int Compute(int a, int b, Calc calc)
{
    return calc(a, b);
}
```

调用：

```csharp
int result = Compute(3, 5, Add);
```

### 速记

- 委托让“传方法”成为可能
- 很多高级特性底层都依赖委托

---

## 9. 多播委托

一个委托可以绑定多个方法。

### 示例

```csharp
delegate void Notify();

static void A()
{
    Console.WriteLine("A");
}

static void B()
{
    Console.WriteLine("B");
}

Notify notify = A;
notify += B;
notify();
```

输出：

```csharp
A
B
```

### 说明

- `+=` 添加方法
- `-=` 移除方法

### 速记

- 委托可以绑定多个方法
- 常见于事件通知场景

---

## 10. Action 和 Func

在实际开发里，手写 `delegate` 不一定是最高频。
更常见的是：

- `Action`
- `Func`

### Action

表示：

- 有参数或无参数
- **没有返回值**

#### 示例

```csharp
Action sayHello = () => Console.WriteLine("Hello");
sayHello();
```

```csharp
Action<string> greet = name => Console.WriteLine($"Hello, {name}");
greet("Tom");
```

### Func

表示：

- 有参数或无参数
- **有返回值**

#### 示例

```csharp
Func<int, int, int> add = (a, b) => a + b;
Console.WriteLine(add(3, 5));
```

这里：

- 前两个 `int` 是参数类型
- 最后一个 `int` 是返回值类型

### 速记

- `Action`：无返回值
- `Func`：有返回值
- 最后一个泛型参数是 `Func` 的返回类型

---

## 11. Predicate< T>

还有一个常见类型：

```csharp
Predicate<int> isEven = x => x % 2 == 0;
```

表示：

- 接收一个 `T`
- 返回 `bool`

虽然现在很多时候直接用 `Func< T, bool>` 也可以，但你需要知道它的存在。

### 速记

- `Predicate<T>` 本质上是“判断函数”
- 输入一个值，返回 true/false

---

## 12. 事件 event 是什么

事件可以理解为：

> “一种受限制的委托，用于发布-订阅通知”。

它常用于：

- 按钮点击
- 状态变化通知
- 消息广播
- 系统回调

### 示例

```csharp
class Alarm
{
    public event Action? OnRing;

    public void Ring()
    {
        Console.WriteLine("Alarm ring...");
        OnRing?.Invoke();
    }
}
```

订阅事件：

```csharp
Alarm alarm = new Alarm();
alarm.OnRing += () => Console.WriteLine("Wake up!");
alarm.Ring();
```

### 说明

- `event` 限制了外部只能订阅/取消订阅
- 外部不能随便直接触发事件
- 真正触发通常在类内部完成

### 速记

- 事件基于委托
- 常用于通知机制
- 外部订阅：`+=`
- 外部取消：`-=`
- 内部触发：`Invoke()`

---

## 13. 空安全调用 `?.Invoke()`

### 示例

```csharp
OnRing?.Invoke();
```

### 说明

意思是：

- 如果 `OnRing` 不为 `null`，就调用
- 如果为 `null`，就不调用

因为可能根本没人订阅这个事件。

### 速记

- `?.` 表示空安全调用
- 事件触发常见写法：`SomeEvent?.Invoke()`

---

## 14. async/await 是什么

你可以把 `async/await` 理解为：

> 用同步代码的写法写异步逻辑。

这是 C# 非常核心的现代能力。

最常见场景：

- 网络请求
- 数据库调用
- 文件 IO
- 调第三方接口

### 为什么要异步

因为这些操作通常要等待：

- 等网络返回
- 等磁盘读取
- 等数据库响应

异步可以避免线程一直傻等，提高程序吞吐和响应能力。

### 速记

- `async/await` 主要解决 IO 等待问题
- 不是为了“让代码更快执行”，而是为了更高效利用线程

---

## 15. Task 是什么

在 C# 里，异步方法通常返回：

- `Task`
- `Task<T>`

### 示例

```csharp
static async Task DoWorkAsync()
{
    await Task.Delay(1000);
    Console.WriteLine("done");
}
```

```csharp
static async Task<int> GetNumberAsync()
{
    await Task.Delay(1000);
    return 100;
}
```

### 说明

- `Task`：表示异步操作，没有返回值
- `Task<T>`：表示异步操作，返回 `T`

### 对比

- **Java**：类似 `CompletableFuture`
- **Python**：类似 `async def` + coroutine
- **Go**：Go 没有 `await` 机制，更多依赖 goroutine + channel

### 速记

- 异步无返回值：`Task`
- 异步有返回值：`Task<T>`

---

## 16. async 方法基本写法

### 示例

```csharp
static async Task SayHelloAsync()
{
    await Task.Delay(1000);
    Console.WriteLine("Hello after 1 second");
}
```

### 说明

- 方法前加 `async`
- 方法内部可以写 `await`
- `await` 后面通常接一个 `Task`

### 速记

- `async` 修饰方法
- `await` 等待异步结果
- `await` 只能出现在 `async` 方法里

---

## 17. async 返回值规则

### 常见写法

```csharp
async Task FooAsync()
async Task<int> FooAsync()
async ValueTask FooAsync()
```

入门阶段先重点掌握：

- `Task`
- `Task<T>`

### 不推荐初学者写法

```csharp
async void FooAsync()
```

### 为什么

因为 `async void`：

- 不好等待
- 不好捕获异常
- 一般只用于事件处理器

### 速记

- 普通异步方法优先 `Task` / `Task<T>`
- `async void` 通常只留给事件处理

---

## 18. await 的效果

### 示例

```csharp
static async Task<int> GetDataAsync()
{
    await Task.Delay(1000);
    return 42;
}

static async Task RunAsync()
{
    int value = await GetDataAsync();
    Console.WriteLine(value);
}
```

### 说明

`await` 会：

- 暂停当前方法后续执行
- 不阻塞整个线程做无意义等待
- 异步操作完成后再继续执行后面的代码

### 速记

- `await` 不是“卡死线程等待”
- 它是异步地等待结果回来

---

## 19. Task.Delay 示例

### 示例

```csharp
await Task.Delay(1000);
```

### 说明

- 表示异步等待 1 秒
- 常用于演示异步，不会阻塞线程

### 注意对比

```csharp
Thread.Sleep(1000);
```

- `Thread.Sleep` 会阻塞线程
- `Task.Delay` 是异步等待

### 速记

- 异步等待：`Task.Delay`
- 阻塞线程：`Thread.Sleep`

---

## 20. 同步 vs 异步直觉理解

### 同步

```csharp
var data = GetData();
Console.WriteLine(data);
```

调用方会一直等到结果回来。

### 异步

```csharp
var data = await GetDataAsync();
Console.WriteLine(data);
```

写法看起来像同步，但底层是异步等待。

### 速记

- 异步不是把代码写乱
- `await` 的价值就是让异步逻辑保持可读性

---

## 21. 异常和 async 的关系

### 示例

```csharp
static async Task RunAsync()
{
    try
    {
        await Task.Delay(1000);
        throw new Exception("boom");
    }
    catch (Exception ex)
    {
        Console.WriteLine(ex.Message);
    }
}
```

### 说明

- 异步方法内部同样可以 `try/catch`
- `await` 抛出的异常也可以被捕获

### 速记

- `async/await` 不影响基本异常处理模式
- 异步代码照样 `try/catch`

---

## 22. 一个完整组合示例

### 示例

```csharp
class Downloader
{
    public event Action<string>? OnCompleted;

    public async Task DownloadAsync(string fileName)
    {
        try
        {
            Console.WriteLine($"开始下载: {fileName}");
            await Task.Delay(1000);
            OnCompleted?.Invoke(fileName);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"下载失败: {ex.Message}");
        }
    }
}
```

调用：

```csharp
Downloader downloader = new Downloader();
downloader.OnCompleted += fileName => Console.WriteLine($"下载完成: {fileName}");
await downloader.DownloadAsync("demo.zip");
```

### 这段代码包含了

- 事件
- Lambda
- `Action`
- 异步方法
- `await`
- 异常处理

---

## 23. 常见误区

### 误区 1：`async` 就等于多线程

不完全对。
`async/await` 主要解决异步等待问题，不等于你一定新开线程。

### 误区 2：所有方法都该写成 async

不对。
只有真正有异步 IO 场景的方法，才有必要异步。

### 误区 3：`async void` 很方便

不推荐。
除非是事件处理器。

### 误区 4：异常都用 `catch (Exception)` 就行

不够好。
具体异常优先，`Exception` 兜底。

### 误区 5：事件和委托完全一样

不一样。
事件是“受限制的委托”，更适合发布订阅。

---

## 24. 本课重点总结

### 你必须先掌握的 10 个核心点

1. `try / catch / finally` 是异常处理基础
2. 异常适用于异常情况，不是普通流程控制
3. 委托本质是“方法类型”
4. `Action` 表示无返回值委托
5. `Func` 表示有返回值委托
6. 事件基于委托，适合通知场景
7. `event` 用于发布-订阅
8. 异步方法通常返回 `Task` 或 `Task<T>`
9. `await` 用来等待异步操作完成
10. `async/await` 是 C# 工程开发核心能力

---

## 25. Java / Go / Python 对比总结

| 知识点 | C# | Java | Go | Python |
|---|---|---|---|---|
| 异常处理 | `try/catch/finally` | `try/catch/finally` | 主要 `error` | `try/except/finally` |
| 方法当参数 | 委托 / `Func` / `Action` | 函数式接口 | 函数值 | 函数一等公民 |
| 事件 | `event` | 常靠监听器接口 | 手写回调 | 手写回调/框架机制 |
| 异步返回 | `Task` / `Task<T>` | `CompletableFuture` | goroutine | coroutine |
| 异步关键字 | `async/await` | `CompletableFuture` 链式较多 | 无 await | `async/await` |

---

## 26. 面试 / 实战高频提醒

- `Action` 无返回值，`Func` 有返回值
- `Func<T, bool>` 经常用于条件判断
- `event` 常用于通知和回调机制
- 异步方法尽量返回 `Task`，不要随便用 `async void`
- `Task.Delay` 和 `Thread.Sleep` 的语义不同
- `await` 常用于 IO 异步操作

---

## 27. 本课小练习

### 练习 1：异常处理

要求：

- 写一个除法方法
- 当除数为 0 时抛异常
- 调用时用 `try/catch` 捕获

### 练习 2：自定义委托

要求：

- 定义一个委托 `Calc`
- 写 `Add` 和 `Sub` 两个方法
- 用委托调用它们

### 练习 3：Action / Func

要求：

- 写一个 `Action<string>` 打印问候语
- 写一个 `Func<int, int, int>` 计算乘积

### 练习 4：事件

要求：

- 定义一个 `Door` 类
- 写一个 `OnOpened` 事件
- 打开门时触发事件

### 练习 5：async/await

要求：

- 写一个 `Task<int> GetNumberAsync()`
- 内部 `await Task.Delay(1000)`
- 返回一个整数
- 在调用方使用 `await` 获取结果

---

## 28. 一页速背版

### 异常处理

```csharp
try
{
    // 可能出错
}
catch (Exception ex)
{
    Console.WriteLine(ex.Message);
}
finally
{
    // 收尾
}
```

### 委托

```csharp
delegate int Calc(int a, int b);
Calc calc = Add;
int result = calc(1, 2);
```

### Action / Func

```csharp
Action<string> greet = name => Console.WriteLine(name);
Func<int, int, int> add = (a, b) => a + b;
```

### 事件

```csharp
public event Action? OnRing;
OnRing?.Invoke();
```

### async/await

```csharp
async Task<int> GetNumberAsync()
{
    await Task.Delay(1000);
    return 42;
}
```

### 核心结论

- 异常用于异常情况
- 委托是方法类型
- `Action` 无返回值
- `Func` 有返回值
- 事件适合通知
- 异步返回 `Task` / `Task<T>`
- `await` 是 C# 核心能力

---

## 29. 下一课预告

下一课学习：

- 文件操作
- JSON 序列化 / 反序列化
- 时间处理
- 命名空间
- 项目结构
- NuGet 包管理
- .NET CLI 常用命令

这部分会让你从“会语法”进入“会做项目”。
