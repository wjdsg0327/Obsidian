# 第十课：C# 项目实战与工程化思维——分层架构、日志、配置、异常处理中间件、开发规范

> 目标：把前面 1~9 课的语言和框架知识串起来，形成一套能真正落地做项目的 C# / ASP.NET Core 工程思维

---

## 1. 本课目标

本课学完后，你应该能掌握：

- 一个典型 C# Web 项目应该怎么分层
- Controller / Service / Repository 各自负责什么
- 配置、日志、异常处理应该放在哪里
- 为什么中间件在 Web 项目里很重要
- 日常开发中哪些写法更工程化
- 从 Java / Go / Python 迁移到 C# 时要建立什么思维方式

---

## 2. 为什么这一课重要

如果前面 1~9 课是：

- 会写语法
- 会写 Web API
- 会用 EF Core

那这一课就是：

> 怎么把这些东西组织成“像团队项目”的代码。

很多人学完语法以后，最大的问题不是不会写代码，而是：

- 不知道代码该放哪
- 不知道 Controller 该写多少
- 不知道异常、日志、配置怎么统一处理
- 不知道哪些写法后面会把项目搞乱

这课就是解决这些问题。

---

## 3. 一个典型 Web API 项目的常见分层

你可以先记这个最常见版本：

- `Controllers/`：接收请求、返回响应
- `Services/`：业务逻辑
- `Repositories/`：数据访问
- `Entities/`：数据库实体
- `Dtos/`：请求响应模型
- `Data/`：`DbContext`、迁移等
- `Middlewares/`：中间件
- `Extensions/`：扩展方法
- `Common/` 或 `Shared/`：公共工具、常量等

### 速记

- Controller 处理 HTTP
- Service 处理业务
- Repository 处理数据库访问
- DTO 负责接口入参与出参

---

## 4. 为什么要分层

分层的核心价值：

- 职责清晰
- 更容易维护
- 更容易测试
- 更容易扩展
- 更容易多人协作

如果不分层，常见问题是：

- Controller 里全是业务逻辑
- SQL / EF Core 查询散落到处都是
- 输入输出模型和数据库实体混在一起
- 后面一改就牵一大片

### 速记

- 分层不是为了“好看”，而是为了降低混乱

---

## 5. Controller 应该负责什么

Controller 的职责应该尽量薄。

它通常负责：

- 接收请求
- 参数绑定
- 调用 Service
- 返回 HTTP 响应

### 推荐示例

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
    public async Task<ActionResult<UserDto>> GetById(int id)
    {
        UserDto? user = await _userService.GetByIdAsync(id);
        if (user == null)
        {
            return NotFound();
        }

        return user;
    }
}
```

### 不推荐现象

- Controller 里堆几十行查询逻辑
- Controller 里写复杂业务判断
- Controller 里直接拼大量 DTO 映射、事务、权限、日志细节

### 速记

- Controller 要“薄”
- 主要做请求入口和响应出口

---

## 6. Service 应该负责什么

Service 层负责：

- 业务规则
- 业务流程编排
- 调用 Repository / DbContext
- 聚合多个依赖

### 示例

```csharp
public interface IUserService
{
    Task<UserDto?> GetByIdAsync(int id);
}

public class UserService : IUserService
{
    private readonly IUserRepository _userRepository;

    public UserService(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    public async Task<UserDto?> GetByIdAsync(int id)
    {
        var user = await _userRepository.GetByIdAsync(id);
        if (user == null)
        {
            return null;
        }

        return new UserDto
        {
            Id = user.Id,
            Name = user.Name
        };
    }
}
```

### 速记

- Service 层是业务逻辑核心
- Controller 不该替代 Service

---

## 7. Repository 应该负责什么

Repository 层负责：

- 封装数据访问
- 统一数据库查询逻辑
- 隔离具体 ORM 细节

### 示例

```csharp
public interface IUserRepository
{
    Task<User?> GetByIdAsync(int id);
}

public class UserRepository : IUserRepository
{
    private readonly AppDbContext _context;

    public UserRepository(AppDbContext context)
    {
        _context = context;
    }

    public async Task<User?> GetByIdAsync(int id)
    {
        return await _context.Users.FirstOrDefaultAsync(x => x.Id == id);
    }
}
```

### 说明

并不是所有项目都必须显式 Repository。

小项目常见两种写法：

- Controller -> DbContext
- Controller -> Service -> DbContext

中大项目更常见：

- Controller -> Service -> Repository -> DbContext

### 速记

- Repository 不是绝对必须
- 但在复杂项目里很常见

---

## 8. DTO 和 Entity 一定要分开吗

入门时小项目不一定强制，但正式项目里通常建议分开。

### Entity

- 对应数据库结构
- 往往包含导航属性、数据库字段约束等

### DTO

- 对应接口输入输出
- 更关注前后端交互契约

### 为什么分开

- 避免暴露内部字段
- 避免接口直接绑定数据库结构
- 更方便接口演进

### 速记

- 小 demo 可以简单些
- 正式项目建议 Entity / DTO 分离

---

## 9. 配置管理：appsettings.json

ASP.NET Core 项目里，配置通常放在：

- `appsettings.json`
- `appsettings.Development.json`
- 环境变量

### 示例

```json
{
  "ConnectionStrings": {
    "DefaultConnection": "Server=.;Database=DemoDb;Trusted_Connection=True;TrustServerCertificate=True;"
  },
  "Jwt": {
    "Key": "your-secret-key",
    "Issuer": "DemoApi"
  }
}
```

### 说明

常见配置内容：

- 数据库连接串
- Redis
- JWT
- 第三方接口地址
- 功能开关

### 速记

- 配置不要硬编码在业务代码里
- 优先放配置文件或环境变量

---

## 10. 读取配置

### 示例

```csharp
string? conn = builder.Configuration.GetConnectionString("DefaultConnection");
```

或：

```csharp
string? issuer = builder.Configuration["Jwt:Issuer"];
```

### 更推荐：绑定配置类

```csharp
public class JwtOptions
{
    public string Key { get; set; } = "";
    public string Issuer { get; set; } = "";
}
```

注册：

```csharp
builder.Services.Configure<JwtOptions>(builder.Configuration.GetSection("Jwt"));
```

### 速记

- 简单值可直接读配置
- 成组配置更推荐绑定配置类

---

## 11. 日志 Logging 是什么

日志是后端项目非常关键的基础设施。

你需要它来：

- 记录程序运行过程
- 排查问题
- 分析错误
- 观察业务行为

ASP.NET Core 默认就有日志体系。

### 注入 ILogger

```csharp
public class UserService : IUserService
{
    private readonly ILogger<UserService> _logger;

    public UserService(ILogger<UserService> logger)
    {
        _logger = logger;
    }
}
```

### 记录日志

```csharp
_logger.LogInformation("Creating user: {Name}", name);
_logger.LogWarning("User not found: {Id}", id);
_logger.LogError(ex, "Create user failed");
```

### 速记

- ASP.NET Core 默认支持日志
- 优先通过 DI 注入 `ILogger<T>`

---

## 12. 常见日志级别

- `Trace`
- `Debug`
- `Information`
- `Warning`
- `Error`
- `Critical`

### 入门直觉

- 正常业务过程：`Information`
- 可疑但未必出错：`Warning`
- 失败异常：`Error`

### 速记

- 不是所有东西都该打 Error
- 日志级别要有层次感

---

## 13. 日志的推荐写法

### 推荐

```csharp
_logger.LogInformation("User {UserId} login success", userId);
```

### 不推荐

```csharp
_logger.LogInformation($"User {userId} login success");
```

### 为什么

结构化日志写法更适合后续日志检索、聚合、分析。

### 速记

- 优先参数化日志，不优先字符串插值日志

---

## 14. 异常处理应该放哪里

很多初学者喜欢在每个 Controller 方法里都写：

```csharp
try { ... } catch { ... }
```

这会导致：

- 重复代码多
- 逻辑分散
- 返回风格不统一

更工程化的做法通常是：

- 业务异常在必要处抛出
- 统一异常处理中间件负责兜底

### 速记

- 不是所有异常都要在 Controller 里本地 try/catch
- 很多项目会统一处理中间件兜底

---

## 15. 中间件 Middleware 是什么

中间件可以理解为：

> 处理 HTTP 请求和响应的一段管道逻辑。

它常用于：

- 异常处理
- 日志记录
- 鉴权
- 跨域
- 请求耗时统计

### Program.cs 中常见写法

```csharp
app.UseHttpsRedirection();
app.UseAuthorization();
```

这些就是中间件。

### 速记

- 中间件处理“通用请求逻辑”
- Controller 处理“具体业务逻辑”

---

## 16. 统一异常处理中间件

### 示例

```csharp
public class ExceptionMiddleware
{
    private readonly RequestDelegate _next;
    private readonly ILogger<ExceptionMiddleware> _logger;

    public ExceptionMiddleware(RequestDelegate next, ILogger<ExceptionMiddleware> logger)
    {
        _next = next;
        _logger = logger;
    }

    public async Task InvokeAsync(HttpContext context)
    {
        try
        {
            await _next(context);
        }
        catch (Exception ex)
        {
            _logger.LogError(ex, "Unhandled exception");
            context.Response.StatusCode = 500;
            context.Response.ContentType = "application/json";
            await context.Response.WriteAsJsonAsync(new
            {
                Message = "Internal Server Error"
            });
        }
    }
}
```

注册：

```csharp
app.UseMiddleware<ExceptionMiddleware>();
```

### 说明

这样未处理异常就能统一记录和返回，不需要每个 Controller 都重复写。

### 速记

- 异常处理中间件 = 全局兜底异常处理

---

## 17. 依赖注册集中管理

当项目变大后，`Program.cs` 很容易堆很多注册代码。

### 常见做法

把注册逻辑拆到扩展方法里：

```csharp
public static class ServiceCollectionExtensions
{
    public static IServiceCollection AddApplicationServices(this IServiceCollection services)
    {
        services.AddScoped<IUserService, UserService>();
        services.AddScoped<IUserRepository, UserRepository>();
        return services;
    }
}
```

然后在 `Program.cs` 里写：

```csharp
builder.Services.AddApplicationServices();
```

### 速记

- 项目一大，依赖注册最好模块化

---

## 18. API 返回风格要统一

项目里最好统一：

- 成功返回什么结构
- 失败返回什么结构
- 分页返回什么格式
- 错误码怎么设计

### 示例思路

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "name": "Tom"
  }
}
```

或直接使用更 RESTful 的状态码 + JSON 结果。

### 说明

关键不是哪种唯一正确，而是：

- 风格一致
- 团队能长期维护

### 速记

- 接口返回要统一，不要每个接口一种风格

---

## 19. 分页、过滤、排序要有固定套路

列表接口经常需要：

- 分页
- 过滤
- 排序

### 示例思路

```csharp
public class UserQueryDto
{
    public string? Keyword { get; set; }
    public int PageIndex { get; set; } = 1;
    public int PageSize { get; set; } = 20;
}
```

### 说明

然后在 Service 里统一处理：

- `Where`
- `Skip`
- `Take`
- `OrderBy`

### 速记

- 列表接口尽量标准化：查询 DTO + LINQ 分页套路

---

## 20. 不要过早搞复杂架构

工程化不等于一上来就：

- 过多抽象
- 到处接口
- 到处泛型基类
- 到处“万能”工具类

### 正确理解

小项目：

- 简单清晰最重要

中项目：

- 再逐步分层、抽象

大项目：

- 再考虑更强的模块边界和规范

### 速记

- 工程化不是复杂化
- 先清晰，再抽象

---

## 21. 从 Java / Go / Python 转到 C# 的思维要点

### 如果你来自 Java

你会觉得很多东西很熟：

- OOP
- interface
- DI
- Web Controller
- ORM

你要重点适应的是：

- 属性 `get; set;`
- LINQ
- 委托 / 事件 / Lambda 风格
- `async/await`
- C# 语法更现代、更简洁

### 如果你来自 Go

你要重点适应的是：

- 更强的面向对象建模
- 更重的框架体系
- DI / 中间件 / ORM 的框架化风格
- C# 更依赖约定和生态整合

### 如果你来自 Python

你要重点适应的是：

- 强类型
- 编译期检查
- 泛型、接口、可空引用类型
- 工程结构更显式

### 速记

- Java 背景：重点适应现代 C# 特性
- Go / Python 背景：重点适应 .NET 工程化体系

---

## 22. 一个建议的项目学习路线

如果你想真正把 C# 用起来，建议你接下来做一个小项目，功能包括：

- 用户 CRUD
- 登录接口
- JWT 鉴权
- EF Core + SQL Server / SQLite
- 分页查询
- 全局异常处理中间件
- Swagger
- 日志

### 为什么

因为这个小项目几乎会把你前面 1~10 课的知识全串起来。

### 速记

- 学语言最快的方式不是只看语法，而是做一个完整小项目

---

## 23. 一个建议的最小实战目录结构

```text
DemoApi/
├── Controllers/
│   └── UsersController.cs
├── Services/
│   ├── IUserService.cs
│   └── UserService.cs
├── Repositories/
│   ├── IUserRepository.cs
│   └── UserRepository.cs
├── Dtos/
│   ├── CreateUserDto.cs
│   └── UserDto.cs
├── Entities/
│   └── User.cs
├── Data/
│   └── AppDbContext.cs
├── Middlewares/
│   └── ExceptionMiddleware.cs
├── Program.cs
├── appsettings.json
└── DemoApi.csproj
```

### 速记

- 这个结构不是唯一标准，但非常适合入门和中小项目

---

## 24. 常见误区

### 误区 1：工程化就是多建目录、多写接口

不对。
工程化核心是职责清晰、可维护，不是表面复杂。

### 误区 2：Controller 里查数据库也没问题，反正能跑

小 demo 可以，正式项目容易失控。

### 误区 3：日志就是随便打印字符串

不够。
要考虑级别、结构化、上下文。

### 误区 4：异常都本地 try/catch 最安全

不一定。
很多异常应该统一处理中间件兜底。

### 误区 5：所有项目都必须上 Repository

不绝对。
要看项目规模和复杂度。

### 误区 6：抽象越多越专业

不对。
无价值抽象只会增加复杂度。

---

## 25. 本课重点总结

### 你必须先掌握的 12 个核心点

1. 项目代码要按职责分层
2. Controller 应该尽量薄
3. Service 是业务逻辑核心层
4. Repository 负责数据访问封装，但不是所有项目都强制需要
5. DTO 和 Entity 正式项目中尽量分离
6. 配置不要硬编码，优先放配置系统
7. 日志通过 `ILogger<T>` 注入
8. 结构化日志优于字符串拼接日志
9. 全局异常处理中间件能统一处理未捕获异常
10. 中间件处理通用请求逻辑
11. 工程化不是复杂化，重点是清晰和可维护
12. 最好的学习方式是做一个完整小项目把知识串起来

---

## 26. Java / Go / Python 对比总结

| 工程点 | C# / ASP.NET Core | Java / Spring Boot | Go | Python |
|---|---|---|---|---|
| 分层 | Controller/Service/Repository 常见 | 非常类似 | 常更轻量 | 框架差异大 |
| DI | 内置强支持 | Spring 强支持 | 多手动 | 框架各异 |
| 中间件 | 非常核心 | Filter / Interceptor / Advice 等 | Middleware | Middleware |
| 配置 | appsettings + options | yml/properties | env/config | env/config |
| 日志 | `ILogger<T>` | slf4j/logback | zap/logrus | logging |

---

## 27. 面试 / 实战高频提醒

- Controller 薄、Service 厚是常见项目习惯
- DTO 不等于 Entity
- `ILogger<T>` 是必须熟悉的基础设施
- 全局异常处理中间件非常高频
- `appsettings.json` 和环境变量都很重要
- 代码清晰比过度抽象更重要
- 会搭一个最小完整 CRUD API 项目，比背很多语法更有价值

---

## 28. 本课小练习

### 练习 1：分层改造

要求：

- 把一个简单用户查询接口拆成 Controller + Service
- 再进一步拆成 Controller + Service + Repository
- 对比代码可读性

### 练习 2：日志注入

要求：

- 在 `UserService` 中注入 `ILogger<UserService>`
- 在创建用户时记录一条信息日志
- 在用户不存在时记录一条警告日志

### 练习 3：配置读取

要求：

- 在 `appsettings.json` 中配置一个 `AppName`
- 在程序中读取并输出它

### 练习 4：异常处理中间件

要求：

- 自定义一个全局异常处理中间件
- 捕获异常并返回统一 JSON 错误响应

### 练习 5：目录结构整理

要求：

- 按本课建议创建目录结构
- 把 DTO、Entity、Service、Repository 分开放置

### 练习 6：做一个小项目草图

要求：

- 设计一个最小用户管理 API
- 包括：查询列表、查询详情、创建、更新、删除
- 明确每一层各放什么代码

---

## 29. 一页速背版

### 常见分层

```text
Controllers/
Services/
Repositories/
Dtos/
Entities/
Data/
Middlewares/
```

### Controller 原则

- 接请求
- 调 Service
- 回响应

### Service 原则

- 放业务逻辑
- 编排流程

### Repository 原则

- 封装数据访问

### 日志

```csharp
private readonly ILogger<UserService> _logger;
_logger.LogInformation("Creating user: {Name}", name);
```

### 配置

```csharp
builder.Configuration.GetConnectionString("DefaultConnection");
builder.Configuration["Jwt:Issuer"];
```

### 中间件

```csharp
app.UseMiddleware<ExceptionMiddleware>();
```

### 核心结论

- 分层要清晰
- Controller 要薄
- 日志、配置、异常处理要工程化
- 中间件处理通用逻辑
- 先做成，再做优雅，不要过度设计

---

## 30. 结语：你现在该怎么继续学

如果你已经跟着学到这里，你已经完成了 C# 快速入门里最有价值的主体部分。

你现在建议进入的阶段是：

### 第一阶段：巩固

- 把 1~10 课每课代码自己敲一遍
- 每课小练习至少做 3 个

### 第二阶段：做项目

做一个最小 Web API 项目，功能包含：

- 用户 CRUD
- DTO
- EF Core
- SQL Server / SQLite
- 日志
- 全局异常处理中间件
- Swagger

### 第三阶段：进阶

后面再深入：

- JWT 鉴权
- 过滤器
- 中间件进阶
- 缓存
- Redis
- 单元测试
- Docker 部署
- 性能优化
- 微服务/消息队列

### 一句话总结

> 学到这里，你已经不是“C# 零基础”，而是已经具备做一个标准 ASP.NET Core CRUD 项目的基础能力了。
