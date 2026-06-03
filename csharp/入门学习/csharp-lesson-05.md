# 第五课：C# 文件操作、JSON、时间处理、命名空间、项目结构、NuGet、.NET CLI

> 目标：从“会写语法”进入“会做项目”，掌握 C# / .NET 日常工程开发最常见的一组基础能力

---

## 1. 本课目标

本课学完后，你应该能掌握：

- 常见文件和目录操作
- 文本读写
- JSON 序列化和反序列化
- `DateTime` 的基本使用
- 命名空间的作用
- .NET 项目的基本结构
- NuGet 包管理基础
- 常用 `dotnet` CLI 命令

---

## 2. 为什么这一课很重要

前几课你学的是“语言能力”。
这一课学的是“工程能力”。

实际工作中，你会经常做：

- 读取配置文件
- 写日志文件
- 解析 JSON
- 调接口并处理时间字段
- 管理项目依赖
- 使用命令行创建、构建、运行项目

这些都是 .NET 开发的日常。

---

## 3. 文件操作基础：File

`File` 是最常用的文件工具类之一。

### 写入文件

```csharp
File.WriteAllText("test.txt", "hello world");
```

### 读取文件

```csharp
string content = File.ReadAllText("test.txt");
Console.WriteLine(content);
```

### 追加写入

```csharp
File.AppendAllText("test.txt", "\nnew line");
```

### 判断文件是否存在

```csharp
if (File.Exists("test.txt"))
{
    Console.WriteLine("文件存在");
}
```

### 速记

- 读文本：`ReadAllText`
- 写文本：`WriteAllText`
- 追加：`AppendAllText`
- 判断存在：`Exists`

---

## 4. 按行读写文件

### 一次性读取所有行

```csharp
string[] lines = File.ReadAllLines("test.txt");
foreach (string line in lines)
{
    Console.WriteLine(line);
}
```

### 一次性写入多行

```csharp
string[] lines = { "line1", "line2", "line3" };
File.WriteAllLines("test.txt", lines);
```

### 速记

- 按行读：`ReadAllLines`
- 按行写：`WriteAllLines`
- 小文件处理很方便

---

## 5. Directory：目录操作

### 创建目录

```csharp
Directory.CreateDirectory("logs");
```

### 判断目录是否存在

```csharp
if (Directory.Exists("logs"))
{
    Console.WriteLine("目录存在");
}
```

### 获取目录下文件

```csharp
string[] files = Directory.GetFiles("logs");
foreach (string file in files)
{
    Console.WriteLine(file);
}
```

### 速记

- 创建目录：`CreateDirectory`
- 判断目录：`Exists`
- 列出文件：`GetFiles`

---

## 6. Path：路径处理

跨平台开发时，尽量别手写拼路径。
优先使用 `Path`。

### 示例

```csharp
string path = Path.Combine("logs", "app.txt");
Console.WriteLine(path);
```

### 获取扩展名

```csharp
string ext = Path.GetExtension("demo.json");
```

### 获取文件名

```csharp
string fileName = Path.GetFileName("logs/app.txt");
```

### 说明

- `Path.Combine` 会自动处理分隔符
- 避免你自己写 `/` 或 `\`

### 速记

- 拼路径优先 `Path.Combine`
- 文件名：`GetFileName`
- 扩展名：`GetExtension`

---

## 7. using 与资源释放

某些对象使用完后需要及时释放，比如文件流。

### 示例

```csharp
using StreamWriter writer = new StreamWriter("test.txt");
writer.WriteLine("hello");
```

### 说明

- `using` 可以确保对象用完后自动释放资源
- 常用于文件、网络流、数据库连接等

### 对比

- **Java**：类似 try-with-resources
- **Python**：类似 `with open(...) as f:`

### 速记

- 资源型对象优先考虑 `using`
- 用完自动释放，避免资源泄漏

---

## 8. JSON 处理基础

C# 里常用 `System.Text.Json` 处理 JSON。

### 示例类

```csharp
class User
{
    public string Name { get; set; } = "";
    public int Age { get; set; }
}
```

---

## 9. JSON 序列化

把对象转成 JSON 字符串。

### 示例

```csharp
using System.Text.Json;

User user = new User
{
    Name = "Tom",
    Age = 20
};

string json = JsonSerializer.Serialize(user);
Console.WriteLine(json);
```

输出可能是：

```json
{"Name":"Tom","Age":20}
```

### 速记

- 对象转 JSON：`JsonSerializer.Serialize(obj)`

---

## 10. JSON 反序列化

把 JSON 字符串转成对象。

### 示例

```csharp
string json = "{\"Name\":\"Tom\",\"Age\":20}";
User? user = JsonSerializer.Deserialize<User>(json);

if (user != null)
{
    Console.WriteLine(user.Name);
}
```

### 说明

- 返回值可能为 `null`
- 所以通常写成 `User?`

### 速记

- JSON 转对象：`Deserialize<T>(json)`
- 反序列化结果要考虑 `null`

---

## 11. JSON 格式化输出

### 示例

```csharp
var options = new JsonSerializerOptions
{
    WriteIndented = true
};

string json = JsonSerializer.Serialize(user, options);
Console.WriteLine(json);
```

### 说明

- `WriteIndented = true` 会让 JSON 更易读
- 常用于调试、配置文件生成

### 速记

- 美化 JSON：`WriteIndented = true`

---

## 12. JSON 属性名映射

有时 JSON 字段名和 C# 属性名不一致。

### 示例

```csharp
using System.Text.Json.Serialization;

class User
{
    [JsonPropertyName("user_name")]
    public string Name { get; set; } = "";
}
```

### 说明

这样 JSON 中的 `user_name` 就能映射到 `Name`。

### 速记

- 字段映射常用 `[JsonPropertyName("xxx")]`

---

## 13. DateTime 基础

### 获取当前时间

```csharp
DateTime now = DateTime.Now;
Console.WriteLine(now);
```

### 获取 UTC 时间

```csharp
DateTime utcNow = DateTime.UtcNow;
```

### 创建指定时间

```csharp
DateTime dt = new DateTime(2026, 4, 2, 10, 30, 0);
```

### 速记

- 本地时间：`DateTime.Now`
- UTC 时间：`DateTime.UtcNow`
- 指定时间：`new DateTime(...)`

---

## 14. 时间格式化

### 示例

```csharp
DateTime now = DateTime.Now;
Console.WriteLine(now.ToString("yyyy-MM-dd HH:mm:ss"));
```

### 常用格式

- `yyyy-MM-dd`
- `yyyy-MM-dd HH:mm:ss`
- `HH:mm:ss`

### 对比

- 和 Java 的时间格式化思路接近
- Python 里类似 `strftime`

### 速记

- 时间转字符串常用 `ToString("格式")`

---

## 15. 时间计算

### 示例

```csharp
DateTime now = DateTime.Now;
DateTime tomorrow = now.AddDays(1);
DateTime nextHour = now.AddHours(1);
```

### 时间差

```csharp
TimeSpan diff = tomorrow - now;
Console.WriteLine(diff.TotalHours);
```

### 说明

- `AddDays` / `AddHours` / `AddMinutes` 很常用
- 两个时间相减会得到 `TimeSpan`

### 速记

- 时间加减：`AddDays`、`AddHours`
- 时间差：`TimeSpan`

---

## 16. DateTime 解析

### 示例

```csharp
DateTime dt = DateTime.Parse("2026-04-02 10:30:00");
```

更稳妥写法：

```csharp
if (DateTime.TryParse("2026-04-02 10:30:00", out DateTime dt))
{
    Console.WriteLine(dt);
}
```

### 说明

- `Parse` 失败会抛异常
- `TryParse` 更稳妥

### 速记

- 不确定输入是否合法时优先 `TryParse`

---

## 17. 命名空间 namespace

命名空间可以理解为：

> 用来组织代码，避免命名冲突。

### 示例

```csharp
namespace MyApp.Services;

class UserService
{
}
```

或者传统写法：

```csharp
namespace MyApp.Services
{
    class UserService
    {
    }
}
```

### 说明

- 类似 Java 的 `package`
- 用于划分模块和代码层次

### 速记

- `namespace` 类似 Java 的 `package`
- 用来分模块、避冲突

---

## 18. using 的含义

### 示例

```csharp
using System;
using System.Collections.Generic;
```

### 说明

- `using` 用于导入命名空间
- 类似 Java 的 `import`

注意区分：

- `using System;`：导入命名空间
- `using StreamWriter writer = ...`：资源释放语法

### 速记

- `using` 有两种常见用途：导入命名空间、自动释放资源

---

## 19. .NET 项目基本结构

一个典型项目里常会看到：

- `Program.cs`：程序入口
- `*.csproj`：项目文件
- `bin/`：编译输出
- `obj/`：中间构建文件

### `.csproj` 是什么

它是项目配置文件，用来描述：

- 目标框架
- 依赖包
- 编译选项

### 示例

```xml
<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>net8.0</TargetFramework>
    <ImplicitUsings>enable</ImplicitUsings>
    <Nullable>enable</Nullable>
  </PropertyGroup>
</Project>
```

### 速记

- `Program.cs`：入口
- `.csproj`：项目配置核心
- `bin/`：输出目录
- `obj/`：中间文件

---

## 20. TargetFramework 是什么

### 示例

```xml
<TargetFramework>net8.0</TargetFramework>
```

### 说明

表示你的项目目标运行在什么 .NET 版本上。

常见：

- `net6.0`
- `net7.0`
- `net8.0`

### 速记

- `TargetFramework` 决定项目运行时目标环境

---

## 21. NuGet 是什么

NuGet 是 .NET 的包管理器。

你可以把它理解为：

- Java 的 Maven Central / Gradle 依赖体系
- Python 的 pip
- Go 的 module 依赖管理生态的一部分

通过 NuGet，你可以安装：

- JSON 库
- ORM
- HTTP 客户端扩展
- 日志库
- 测试框架

### 速记

- NuGet = .NET 的依赖包管理系统

---

## 22. 安装 NuGet 包

### CLI 示例

```bash
dotnet add package Newtonsoft.Json
```

### 说明

执行后会修改 `.csproj`，增加包依赖。

### 速记

- 加包：`dotnet add package 包名`

---

## 23. 常用 dotnet CLI 命令

### 创建控制台项目

```bash
dotnet new console -n DemoApp
```

### 进入目录后运行

```bash
dotnet run
```

### 构建项目

```bash
dotnet build
```

### 发布项目

```bash
dotnet publish
```

### 还原依赖

```bash
dotnet restore
```

### 查看 SDK 信息

```bash
dotnet --info
```

### 速记

- 新建项目：`dotnet new`
- 运行：`dotnet run`
- 构建：`dotnet build`
- 发布：`dotnet publish`
- 还原依赖：`dotnet restore`

---

## 24. 最小项目运行流程

### 第一步：创建项目

```bash
dotnet new console -n HelloCSharp
```

### 第二步：进入项目目录

```bash
cd HelloCSharp
```

### 第三步：运行

```bash
dotnet run
```

### 速记

- 学习 C# 最快方式就是用 `dotnet new console` 开一个小项目练习

---

## 25. 配置文件和 JSON 文件实战思路

实际开发里很常见的做法：

1. 从 JSON 文件读取配置
2. 反序列化成 C# 对象
3. 在程序中使用这个配置对象

### 示例思路

```csharp
string json = File.ReadAllText("appsettings.json");
AppConfig? config = JsonSerializer.Deserialize<AppConfig>(json);
```

这类写法在控制台程序、Web 项目、工具项目里都很常见。

---

## 26. 常见误区

### 误区 1：拼路径直接写字符串就行

不推荐。
跨平台时可能出问题，优先 `Path.Combine`。

### 误区 2：JSON 反序列化一定成功

不一定。
JSON 不合法、字段不匹配、空值问题都可能导致失败。

### 误区 3：时间都用 `DateTime.Now` 就够了

不总是。
跨时区、接口传输、数据库存储时，经常要考虑 UTC。

### 误区 4：用了 `File.ReadAllText` 就不用考虑异常

不对。
文件不存在、权限问题、路径错误都可能抛异常。

### 误区 5：NuGet 只是下载库

不止。
它还是 .NET 工程生态的一部分，影响依赖管理、版本管理、项目构建。

---

## 27. 本课重点总结

### 你必须先掌握的 10 个核心点

1. 文本文件读写常用 `File`
2. 目录操作常用 `Directory`
3. 路径拼接优先 `Path.Combine`
4. 资源对象优先考虑 `using`
5. JSON 常用 `System.Text.Json`
6. 对象转 JSON 用 `Serialize`
7. JSON 转对象用 `Deserialize<T>`
8. 时间处理核心类型是 `DateTime` 和 `TimeSpan`
9. `namespace` 用于组织代码
10. `.NET` 工程操作要熟悉 `dotnet` CLI 和 NuGet

---

## 28. Java / Go / Python 对比总结

| 知识点 | C# | Java | Go | Python |
|---|---|---|---|---|
| 文件读写 | `File` | `Files` / IO | `os` / `io` | `open()` |
| 路径处理 | `Path.Combine` | `Paths.get` | `path/filepath` | `os.path.join` |
| JSON | `System.Text.Json` | Jackson/Gson | `encoding/json` | `json` |
| 时间 | `DateTime` | `LocalDateTime` 等 | `time.Time` | `datetime` |
| 包管理 | NuGet | Maven/Gradle | go mod | pip |
| CLI | `dotnet` | `mvn` / `gradle` | `go` | `python` / `pip` |

---

## 29. 面试 / 实战高频提醒

- 路径拼接优先 `Path.Combine`
- 文件/目录操作常配合异常处理
- JSON 反序列化结果要考虑 `null`
- `DateTime.Now` 和 `DateTime.UtcNow` 语义不同
- `using` 既能导包，也能自动释放资源
- `.csproj` 是 .NET 项目配置核心文件
- `dotnet run/build/publish` 是必会命令

---

## 30. 本课小练习

### 练习 1：写文本文件

要求：

- 创建一个 `notes.txt`
- 写入三行文字
- 再读取并输出内容

### 练习 2：路径拼接

要求：

- 用 `Path.Combine` 拼接一个日志文件路径
- 输出结果

### 练习 3：JSON 序列化

要求：

- 定义一个 `User` 类
- 创建对象
- 序列化成 JSON
- 打印 JSON 字符串

### 练习 4：JSON 反序列化

要求：

- 准备一个 JSON 字符串
- 反序列化成对象
- 输出对象属性

### 练习 5：时间处理

要求：

- 输出当前时间
- 输出明天同一时刻
- 输出两者间隔小时数

### 练习 6：dotnet CLI

要求：

- 新建一个控制台项目
- 运行它
- 修改 `Program.cs` 输出一句自定义内容

---

## 31. 一页速背版

### 文件读写

```csharp
File.WriteAllText("a.txt", "hello");
string s = File.ReadAllText("a.txt");
File.AppendAllText("a.txt", "\nmore");
```

### 目录与路径

```csharp
Directory.CreateDirectory("logs");
string path = Path.Combine("logs", "app.txt");
```

### JSON

```csharp
string json = JsonSerializer.Serialize(obj);
User? user = JsonSerializer.Deserialize<User>(json);
```

### 时间

```csharp
DateTime now = DateTime.Now;
DateTime utc = DateTime.UtcNow;
DateTime tomorrow = now.AddDays(1);
```

### dotnet CLI

```bash
dotnet new console -n DemoApp
dotnet run
dotnet build
dotnet publish
dotnet add package Newtonsoft.Json
```

### 核心结论

- 文件用 `File`
- 目录用 `Directory`
- 路径用 `Path.Combine`
- JSON 用 `System.Text.Json`
- 时间核心是 `DateTime`
- 工程管理靠 `.csproj`、NuGet、`dotnet` CLI

---

## 32. 下一课预告

下一课学习：

- C# 可空引用类型
- 值类型 vs 引用类型再深入
- `record`
- 模式匹配
- `switch expression`
- `var`、`object`、`dynamic` 的区别

这部分会让你真正理解 C# 的现代语法体系。
