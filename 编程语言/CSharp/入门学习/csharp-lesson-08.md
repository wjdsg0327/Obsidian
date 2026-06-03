# 第八课：ASP.NET Core 入门——Web API 基本结构、路由、Controller、请求响应、DTO、依赖注入

> 目标：快速建立 ASP.NET Core Web API 的完整认知，知道一个 C# Web 项目是怎么跑起来的、请求是怎么流转的、代码应该怎么组织

---

## 1. 本课目标

本课学完后，你应该能掌握：

- 什么是 ASP.NET Core
- Web API 项目的基本结构
- 请求从进入程序到返回响应的基本流程
- 路由怎么定义
- Controller 怎么写
- 请求参数怎么接收
- 响应怎么返回
- DTO 是什么，为什么重要
- ASP.NET Core 里的依赖注入怎么用

---

## 2. 什么是 ASP.NET Core

你可以把 ASP.NET Core 理解为：

> C# / .NET 生态里做 Web 开发的主流框架。

它可以用来开发：

- Web API
- MVC 网站
- 实时应用
- 微服务
- 后台管理系统

如果你有 Java 背景，可以把它类比成：

- **Spring Boot + MVC + 内置依赖注入体系**

如果你有 Go 背景，可以把它理解为：

- 比 Gin / Echo 更“框架化”和工程化的一套 Web 体系

### 速记

- C# 做后端开发，主流就是 ASP.NET Core
- Web API 是最常见入门方向

---

## 3. 创建一个 Web API 项目

### CLI 命令

```bash
dotnet new webapi -n DemoApi
```

### 运行项目

```bash
cd DemoApi
dotnet run
```

### 说明

创建后你会得到一个可运行的 API 项目。

### 速记

- 新建 API 项目：`dotnet new webapi -n 项目名`
- 启动：`dotnet run`

---

## 4. Program.cs 是启动入口

现代 ASP.NET Core 项目通常从 `Program.cs` 开始。

### 典型示例

```csharp
var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.UseAuthorization();
app.MapControllers();

app.Run();
```

### 说明

这里大致做了几件事：

1. 创建应用构建器
2. 注册服务
3. 构建应用
4. 配置中间件
5. 映射控制器路由
6. 启动程序

### 速记

- `builder.Services`：注册服务
- `app.UseXxx()`：配置中间件
- `app.MapControllers()`：启用控制器路由
- `app.Run()`：启动 Web 应用

---

## 5. 请求处理流程的整体直觉

一个 HTTP 请求进来后，大致会经历：

1. 进入 ASP.NET Core 应用
2. 经过中间件管道
3. 根据路由匹配到某个 Controller / Action
4. 执行业务逻辑
5. 返回响应

你可以把它理解为：

> 请求先过“管道”，再进“控制器”，最后产出“响应”。

### 速记

- 中间件负责通用处理
- Controller 负责具体接口逻辑

---

## 6. Controller 是什么

Controller 可以理解为：

> 用来接收 HTTP 请求并返回结果的类。

### 示例

```csharp
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase
{
    [HttpGet]
    public IActionResult GetAll()
    {
        return Ok(new[] { "Tom", "Jack" });
    }
}
```

### 说明

- `UsersController` 是一个控制器
- `ControllerBase` 是 Web API 控制器常用基类
- `[ApiController]` 提供很多 API 友好特性
- `[Route(...)]` 定义路由规则

### 速记

- 控制器类通常以 `Controller` 结尾
- Web API 常继承 `ControllerBase`

---

## 7. 路由 Route 是什么

路由可以理解为：

> 把某个 URL 请求映射到某个控制器方法。

### 示例

```csharp
[Route("api/[controller]")]
```

如果控制器名是 `UsersController`，那么：

- `[controller]` 会替换成 `Users`
- 最终路由类似：`api/users`

### 速记

- 路由决定 URL 访问路径
- `[controller]` 会自动取控制器名去掉 `Controller`

---

## 8. HTTP 动词与 Action

常见 HTTP 动词：

- `GET`：查询
- `POST`：创建
- `PUT`：整体更新
- `PATCH`：部分更新
- `DELETE`：删除

在 ASP.NET Core 中，通常通过特性标记：

### 示例

```csharp
[HttpGet]
public IActionResult GetAll() { ... }

[HttpGet("{id}")]
public IActionResult GetById(int id) { ... }

[HttpPost]
public IActionResult Create(CreateUserDto dto) { ... }

[HttpPut("{id}")]
public IActionResult Update(int id, UpdateUserDto dto) { ... }

[HttpDelete("{id}")]
public IActionResult Delete(int id) { ... }
```

### 速记

- 查询：`[HttpGet]`
- 创建：`[HttpPost]`
- 更新：`[HttpPut]`
- 删除：`[HttpDelete]`

---

## 9. 路由参数

### 示例

```csharp
[HttpGet("{id}")]
public IActionResult GetById(int id)
{
    return Ok(new { Id = id, Name = "Tom" });
}
```

请求可能是：

```http
GET /api/users/1
```

### 说明

- `{id}` 表示路由占位符
- 方法参数 `int id` 会自动接收它

### 速记

- URL 路径中的参数常用 `{id}`
- 方法参数名通常要对应上

---

## 10. 查询参数 Query String

### 示例

```csharp
[HttpGet("search")]
public IActionResult Search(string keyword)
{
    return Ok($"search: {keyword}");
}
```

请求可能是：

```http
GET /api/users/search?keyword=tom
```

### 说明

- `keyword` 来自 URL 查询参数
- 默认会自动绑定

### 速记

- `?a=1&b=2` 这种就是查询参数
- 简单场景可直接写到方法参数里

---

## 11. 从请求体接收 JSON

### DTO 示例

```csharp
public class CreateUserDto
{
    public string Name { get; set; } = "";
    public int Age { get; set; }
}
```

### Controller 示例

```csharp
[HttpPost]
public IActionResult Create(CreateUserDto dto)
{
    return Ok(new { Message = $"Created user: {dto.Name}" });
}
```

### 请求体示例

```json
{
  "name": "Tom",
  "age": 20
}
```

### 说明

ASP.NET Core 会自动把 JSON 反序列化成 `CreateUserDto`。

### 速记

- POST/PUT 常从请求体接收 JSON
- JSON 会自动绑定到 DTO 对象

---

## 12. DTO 是什么

DTO = Data Transfer Object。

你可以把它理解为：

> 专门用于接口请求和响应的数据对象。

### 为什么要 DTO

因为你通常不应该：

- 直接把数据库实体裸露给前端
- 直接让所有内部字段暴露出去

DTO 的价值：

- 控制输入输出结构
- 避免泄露内部字段
- 让接口语义更稳定
- 方便后续演进

### 速记

- DTO 是接口层的数据模型
- 请求 DTO 和响应 DTO 最好分开考虑

---

## 13. 返回响应：IActionResult

### 示例

```csharp
[HttpGet]
public IActionResult Get()
{
    return Ok(new { Message = "success" });
}
```

### 常见返回方法

- `Ok(...)`：200
- `Created(...)` / `CreatedAtAction(...)`：201
- `NoContent()`：204
- `BadRequest(...)`：400
- `NotFound()`：404
- `Unauthorized()`：401

### 速记

- `IActionResult` 是最常见控制器返回类型
- 成功返回常用 `Ok()`
- 查不到常用 `NotFound()`
- 参数错误常用 `BadRequest()`

---

## 14. 返回强类型结果

除了 `IActionResult`，也可以直接返回对象：

### 示例

```csharp
[HttpGet("{id}")]
public UserDto GetById(int id)
{
    return new UserDto { Id = id, Name = "Tom" };
}
```

### 说明

框架会自动序列化成 JSON。

但入门阶段你先优先掌握：

- `IActionResult`
- `ActionResult<T>`

因为更灵活。

### 速记

- 简单场景能直接返回对象
- 更通用的控制器返回方式还是 `IActionResult`

---

## 15. ActionResult<T>

### 示例

```csharp
[HttpGet("{id}")]
public ActionResult<UserDto> GetById(int id)
{
    if (id <= 0)
    {
        return BadRequest();
    }

    return new UserDto { Id = id, Name = "Tom" };
}
```

### 说明

`ActionResult<T>` 的好处：

- 既能返回 `T`
- 也能返回 `BadRequest()`、`NotFound()` 这类结果

### 速记

- `ActionResult<T>` = 强类型结果 + 状态码结果都能兼顾

---

## 16. [ApiController] 的作用

### 示例

```csharp
[ApiController]
public class UsersController : ControllerBase
{
}
```

### 它带来的常见好处

- 更友好的参数绑定
- 自动模型验证错误响应
- API 场景默认行为更合理

### 速记

- Web API 控制器一般都加 `[ApiController]`

---

## 17. 模型验证基础

可以在 DTO 上加验证特性。

### 示例

```csharp
using System.ComponentModel.DataAnnotations;

public class CreateUserDto
{
    [Required]
    public string Name { get; set; } = "";

    [Range(1, 150)]
    public int Age { get; set; }
}
```

### 说明

- `[Required]`：必填
- `[Range(1, 150)]`：范围限制

如果用了 `[ApiController]`，很多情况下校验失败会自动返回 400。

### 速记

- DTO 上可加数据校验特性
- `[ApiController]` 配合模型验证很常见

---

## 18. 依赖注入在 Web 项目里怎么用

先定义接口：

```csharp
public interface IUserService
{
    string GetName(int id);
}
```

实现类：

```csharp
public class UserService : IUserService
{
    public string GetName(int id)
    {
        return $"user-{id}";
    }
}
```

注册服务：

```csharp
builder.Services.AddScoped<IUserService, UserService>();
```

控制器注入：

```csharp
[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase
{
    private readonly IUserService _userService;

    public UsersController(IUserService userService)
    {
        _userService = userService;
    }

    [HttpGet("{id}")]
    public IActionResult GetById(int id)
    {
        return Ok(new { Name = _userService.GetName(id) });
    }
}
```

### 说明

这就是 ASP.NET Core 最典型的开发模式之一：

- 接口
- 实现类
- DI 注册
- Controller 构造函数注入

### 速记

- 服务注册到 `builder.Services`
- 控制器通过构造函数拿依赖

---

## 19. 生命周期：Transient / Scoped / Singleton

ASP.NET Core 注册服务时常见三种生命周期：

### Transient

```csharp
builder.Services.AddTransient<IUserService, UserService>();
```

- 每次请求依赖时都创建新实例

### Scoped

```csharp
builder.Services.AddScoped<IUserService, UserService>();
```

- 同一个请求内共用一个实例
- Web 项目里非常常见

### Singleton

```csharp
builder.Services.AddSingleton<IUserService, UserService>();
```

- 整个应用生命周期只一个实例

### 入门建议

先记住：

- 普通业务服务常用 `Scoped`
- 无状态轻量服务有时可用 `Transient`
- 全局共享配置/缓存类才考虑 `Singleton`

### 速记

- `Transient`：每次都新建
- `Scoped`：每个请求一个
- `Singleton`：全局一个

---

## 20. 项目常见分层

一个常见 Web API 项目通常会有：

- `Controllers/`：控制器
- `Services/`：业务逻辑
- `Dtos/`：请求响应模型
- `Entities/` 或 `Models/`：领域/数据库实体
- `Repositories/`：数据访问层

### 说明

请求大致流转：

- Controller 接收请求
- Service 处理业务
- Repository 访问数据库

### 速记

- Controller 不要塞太多业务逻辑
- 业务逻辑尽量放 Service

---

## 21. Swagger 是什么

你创建 Web API 项目时通常会自带 Swagger。

### 作用

- 自动生成接口文档
- 可以在浏览器里调试 API
- 对前后端联调很方便

### Program.cs 常见配置

```csharp
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
```

```csharp
app.UseSwagger();
app.UseSwaggerUI();
```

### 速记

- Swagger = API 文档 + 调试页面
- 学 Web API 时非常好用

---

## 22. 最小完整示例

### DTO

```csharp
public class CreateUserDto
{
    public string Name { get; set; } = "";
    public int Age { get; set; }
}
```

### Service

```csharp
public interface IUserService
{
    string CreateUser(CreateUserDto dto);
}

public class UserService : IUserService
{
    public string CreateUser(CreateUserDto dto)
    {
        return $"Created: {dto.Name}, Age: {dto.Age}";
    }
}
```

### Program.cs

```csharp
builder.Services.AddScoped<IUserService, UserService>();
```

### Controller

```csharp
[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase
{
    private readonly IUserService _userService;

    public UsersController(IUserService userService)
    {
        _userService = userService;
    }

    [HttpPost]
    public IActionResult Create(CreateUserDto dto)
    {
        string result = _userService.CreateUser(dto);
        return Ok(new { Message = result });
    }
}
```

### 这段代码体现了

- Controller
- DTO
- Service
- DI
- 请求体 JSON 绑定
- JSON 响应返回

---

## 23. 常见误区

### 误区 1：Controller 里把所有业务都写完

不推荐。
Controller 应该更偏“接收请求 + 调用服务 + 返回响应”。

### 误区 2：DTO 和数据库实体混着用

不推荐。
接口模型和内部数据模型尽量分开。

### 误区 3：服务类直接在 Controller 里 `new`

不推荐。
优先交给 DI 容器管理。

### 误区 4：所有返回都用 `string`

不推荐。
API 应尽量返回结构化 JSON 和正确状态码。

### 误区 5：只知道写 `[HttpGet]`，不知道 URL 怎么映射

要理解：

- Controller 路由
- Action 路由
- HTTP 方法

三者共同决定一个接口。

---

## 24. 本课重点总结

### 你必须先掌握的 10 个核心点

1. ASP.NET Core 是 C# Web 开发主流框架
2. `Program.cs` 是现代 Web 项目启动入口
3. 请求会经过中间件，再路由到 Controller
4. Controller 用来处理 HTTP 请求
5. 路由决定 URL 到方法的映射关系
6. POST/PUT 常从请求体接收 JSON DTO
7. 响应通常返回 `IActionResult` 或 `ActionResult<T>`
8. DTO 用于控制接口输入输出模型
9. DI 是 ASP.NET Core 的默认核心机制
10. Controller、Service、DTO 分层是高频项目结构

---

## 25. Java / Go / Python 对比总结

| 知识点 | C# / ASP.NET Core | Java / Spring Boot | Go | Python |
|---|---|---|---|---|
| Controller | Controller | RestController | Handler | Flask/FastAPI 路由函数 |
| 路由 | 特性路由 | 注解路由 | 手动注册 | 装饰器路由 |
| DI | 内置强支持 | Spring 强支持 | 多手动注入 | 框架各异 |
| DTO | 很常见 | 很常见 | 常手写 struct | Pydantic / schema |
| Swagger | 支持方便 | 支持方便 | 依赖库 | 框架支持各异 |

---

## 26. 面试 / 实战高频提醒

- 控制器尽量薄，业务逻辑放 Service
- DTO 不要直接等于数据库实体
- `builder.Services` 是依赖注入注册入口
- `AddScoped` 是最常见业务服务注册方式
- `IActionResult` / `ActionResult<T>` 必须熟悉
- `[ApiController]` 和路由特性是基础中的基础

---

## 27. 本课小练习

### 练习 1：创建 API 项目

要求：

- 用 `dotnet new webapi -n DemoApi` 创建项目
- 运行项目
- 在浏览器打开 Swagger

### 练习 2：写一个 UsersController

要求：

- 路由为 `api/users`
- 写一个 `[HttpGet]` 方法
- 返回一个字符串数组

### 练习 3：路由参数

要求：

- 写 `[HttpGet("{id}")]`
- 接收 `id`
- 返回包含 `id` 的 JSON

### 练习 4：POST + DTO

要求：

- 定义 `CreateUserDto`
- 写 `[HttpPost]`
- 从请求体接收 JSON
- 返回创建结果

### 练习 5：服务注入

要求：

- 定义 `IUserService` 和 `UserService`
- 注册到 DI 容器
- 在 Controller 中注入并调用

---

## 28. 一页速背版

### 创建项目

```bash
dotnet new webapi -n DemoApi
dotnet run
```

### Program.cs 关键点

```csharp
builder.Services.AddControllers();
app.MapControllers();
app.Run();
```

### Controller

```csharp
[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase
{
    [HttpGet]
    public IActionResult GetAll()
    {
        return Ok(new[] { "Tom", "Jack" });
    }
}
```

### POST + DTO

```csharp
[HttpPost]
public IActionResult Create(CreateUserDto dto)
{
    return Ok(dto);
}
```

### DI 注册

```csharp
builder.Services.AddScoped<IUserService, UserService>();
```

### 构造函数注入

```csharp
public UsersController(IUserService userService)
{
    _userService = userService;
}
```

### 核心结论

- `Program.cs` 配启动和服务
- Controller 接请求
- DTO 做输入输出
- Service 放业务逻辑
- DI 做解耦

---

## 29. 下一课预告

下一课学习：

- EF Core 入门
- DbContext
- 实体类 Entity
- 数据库迁移 Migration
- 常见 CRUD
- LINQ 查询在数据库中的使用

这课学完，你就具备做完整 C# CRUD 后端的基础了。
