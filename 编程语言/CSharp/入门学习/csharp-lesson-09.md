# 第九课：EF Core 入门——DbContext、Entity、Migration、CRUD、LINQ 数据库查询

> 目标：掌握 C# / ASP.NET Core 后端开发最常见的数据访问方式，建立 EF Core 的完整基础认知

---

## 1. 本课目标

本课学完后，你应该能掌握：

- 什么是 EF Core
- Entity、DbContext 的作用
- 如何配置数据库连接
- Migration 是什么
- 如何做常见 CRUD
- 如何使用 LINQ 查询数据库
- 如何把 EF Core 接入 ASP.NET Core 项目

---

## 2. 什么是 EF Core

EF Core 全称是 **Entity Framework Core**。

你可以把它理解为：

> .NET 生态最主流的 ORM 框架之一。

ORM = Object Relational Mapping，对象关系映射。

它的核心思想是：

- 用 C# 类表示数据库表
- 用对象操作表示数据库增删改查
- 由框架帮你生成 SQL

### 对比

- **Java**：类似 JPA / Hibernate
- **Go**：类似 GORM
- **Python**：类似 SQLAlchemy ORM / Django ORM

### 速记

- EF Core = .NET 主流 ORM
- 用对象和 LINQ 操作数据库

---

## 3. 为什么要学 EF Core

因为在 ASP.NET Core 后端项目里，你会高频做：

- 定义实体类
- 建表
- 查数据
- 插入数据
- 更新数据
- 删除数据
- 用 LINQ 写查询

EF Core 基本就是 C# Web 后端的数据访问基础设施之一。

### 速记

- 学会 EF Core，才能真正写完整 CRUD 后端

---

## 4. Entity 是什么

Entity 可以理解为：

> 映射数据库表的一种 C# 类。

### 示例

```csharp
public class User
{
    public int Id { get; set; }
    public string Name { get; set; } = "";
    public int Age { get; set; }
}
```

### 说明

通常：

- 类名对应表名
- 属性对应字段
- `Id` 常作为主键

### 速记

- Entity = 数据库表对应的 C# 类
- 属性 = 表字段

---

## 5. DbContext 是什么

DbContext 可以理解为：

> EF Core 操作数据库的核心入口。

它负责：

- 管理实体集合
- 跟踪实体状态
- 执行数据库查询和保存

### 示例

```csharp
using Microsoft.EntityFrameworkCore;

public class AppDbContext : DbContext
{
    public AppDbContext(DbContextOptions<AppDbContext> options) : base(options)
    {
    }

    public DbSet<User> Users { get; set; }
}
```

### 说明

- `AppDbContext` 继承 `DbContext`
- `DbSet<User>` 可以理解为 `users` 这张表的操作入口

### 速记

- `DbContext` 是数据库上下文
- `DbSet<T>` 表示某张表

---

## 6. DbSet<T> 是什么

### 示例

```csharp
public DbSet<User> Users { get; set; }
```

### 说明

你可以把它理解为：

> 数据库表在代码里的集合入口。

通过它你可以：

- 查询
- 添加
- 删除
- 更新

### 常见操作对象

```csharp
_context.Users
```

### 速记

- `DbSet<User>` ≈ `User` 表的代码入口

---

## 7. 配置数据库连接

在 ASP.NET Core 项目里，通常先在配置文件写连接串。

### appsettings.json 示例

```json
{
  "ConnectionStrings": {
    "DefaultConnection": "Server=.;Database=DemoDb;Trusted_Connection=True;TrustServerCertificate=True;"
  }
}
```

### Program.cs 注册 DbContext

```csharp
using Microsoft.EntityFrameworkCore;

builder.Services.AddDbContext<AppDbContext>(options =>
    options.UseSqlServer(builder.Configuration.GetConnectionString("DefaultConnection")));
```

### 说明

- `GetConnectionString("DefaultConnection")` 从配置中读取连接串
- `UseSqlServer(...)` 表示使用 SQL Server

### 速记

- 连接串通常放 `appsettings.json`
- `AddDbContext` 把数据库上下文注册进 DI 容器

---

## 8. 常见数据库提供程序

EF Core 支持多种数据库：

- SQL Server
- MySQL
- PostgreSQL
- SQLite

### 示例

```csharp
options.UseSqlServer(...)
options.UseSqlite(...)
options.UseNpgsql(...)
```

### 速记

- EF Core 是抽象层
- 底层数据库通过 provider 决定

---

## 9. Migration 是什么

Migration 可以理解为：

> 用代码记录数据库结构变更的一种机制。

比如你新增一个字段：

- 改了 Entity 类
- 然后生成 Migration
- 再执行 Migration 更新数据库表结构

### 速记

- Migration = 数据库结构演进记录
- 类似数据库版本管理

---

## 10. 创建 Migration

### 常用命令

```bash
dotnet ef migrations add InitialCreate
```

### 说明

这会根据当前模型生成一次迁移。

### 速记

- 新建迁移：`dotnet ef migrations add 名称`

---

## 11. 应用 Migration 到数据库

### 常用命令

```bash
dotnet ef database update
```

### 说明

这会把尚未应用的迁移更新到数据库。

### 速记

- 更新数据库：`dotnet ef database update`

---

## 12. 安装 EF Core 常见包

### 例如 SQL Server

```bash
dotnet add package Microsoft.EntityFrameworkCore.SqlServer
dotnet add package Microsoft.EntityFrameworkCore.Design
dotnet add package Microsoft.EntityFrameworkCore.Tools
```

### 说明

- `SqlServer`：数据库提供程序
- `Design` / `Tools`：迁移等开发工具支持

### 速记

- 做 EF Core 通常要装 provider + design/tools 包

---

## 13. 新增数据 Create

### 示例

```csharp
var user = new User
{
    Name = "Tom",
    Age = 20
};

_context.Users.Add(user);
await _context.SaveChangesAsync();
```

### 说明

- `Add()`：把实体加入上下文
- `SaveChangesAsync()`：真正提交到数据库

### 速记

- 改动不会自动落库
- 落库要调用 `SaveChanges()` / `SaveChangesAsync()`

---

## 14. 查询数据 Read

### 查询全部

```csharp
List<User> users = await _context.Users.ToListAsync();
```

### 按主键查找

```csharp
User? user = await _context.Users.FindAsync(id);
```

### 条件查询

```csharp
User? user = await _context.Users.FirstOrDefaultAsync(x => x.Id == id);
```

### 说明

- `ToListAsync()`：查全部并转成列表
- `FindAsync(id)`：按主键查询
- `FirstOrDefaultAsync(...)`：按条件查第一条或默认值

### 速记

- 查全部：`ToListAsync`
- 按主键：`FindAsync`
- 条件查一条：`FirstOrDefaultAsync`

---

## 15. 更新数据 Update

### 示例

```csharp
User? user = await _context.Users.FindAsync(id);

if (user == null)
{
    return;
}

user.Name = "NewName";
user.Age = 30;

await _context.SaveChangesAsync();
```

### 说明

如果实体是从当前上下文查出来的，EF Core 会自动跟踪它。
你修改属性后再调用 `SaveChangesAsync()`，就会生成更新 SQL。

### 速记

- 查出来 -> 改属性 -> `SaveChangesAsync()`
- 被上下文跟踪的实体会自动识别修改

---

## 16. 删除数据 Delete

### 示例

```csharp
User? user = await _context.Users.FindAsync(id);

if (user == null)
{
    return;
}

_context.Users.Remove(user);
await _context.SaveChangesAsync();
```

### 速记

- 删除：`Remove`
- 最后别忘了 `SaveChangesAsync()`

---

## 17. 为什么优先异步数据库操作

在 Web 项目中，数据库操作通常建议优先用异步：

- `ToListAsync()`
- `FirstOrDefaultAsync()`
- `FindAsync()`
- `SaveChangesAsync()`

### 原因

数据库 IO 是等待型操作。
异步能更高效利用线程，提高 Web 服务吞吐。

### 速记

- Web 项目里数据库操作优先 async 版本

---

## 18. LINQ 查询数据库

EF Core 支持用 LINQ 写数据库查询。

### 示例

```csharp
var users = await _context.Users
    .Where(x => x.Age >= 18)
    .OrderBy(x => x.Name)
    .ToListAsync();
```

### 说明

这段代码表达的是：

- 查询年龄 >= 18 的用户
- 按名字升序排序
- 转成列表

框架会把它翻译成 SQL。

### 速记

- 你写 LINQ，EF Core 帮你翻成 SQL

---

## 19. 常见 LINQ 数据库查询方法

### Where

```csharp
.Where(x => x.Age > 18)
```

过滤条件。

### Select

```csharp
.Select(x => new UserDto
{
    Id = x.Id,
    Name = x.Name
})
```

投影成需要的结果。

### OrderBy / OrderByDescending

```csharp
.OrderBy(x => x.Name)
.OrderByDescending(x => x.Id)
```

排序。

### AnyAsync

```csharp
bool exists = await _context.Users.AnyAsync(x => x.Name == "Tom");
```

判断是否存在。

### CountAsync

```csharp
int count = await _context.Users.CountAsync();
```

统计数量。

### 速记

- 过滤：`Where`
- 投影：`Select`
- 排序：`OrderBy`
- 存在性判断：`AnyAsync`
- 计数：`CountAsync`

---

## 20. Select 投影很重要

在实际项目里，你常常不应该直接返回 Entity。
而是应该投影成 DTO。

### 示例

```csharp
var users = await _context.Users
    .Select(x => new UserDto
    {
        Id = x.Id,
        Name = x.Name
    })
    .ToListAsync();
```

### 说明

这样做的好处：

- 只查需要字段
- 不暴露内部字段
- 接口层更稳定

### 速记

- 查询接口数据时，经常 `Select -> DTO`

---

## 21. FirstOrDefaultAsync vs SingleOrDefaultAsync

### FirstOrDefaultAsync

```csharp
var user = await _context.Users.FirstOrDefaultAsync(x => x.Name == "Tom");
```

- 查第一条，没有则返回 `null`
- 即使多条也不会报错

### SingleOrDefaultAsync

```csharp
var user = await _context.Users.SingleOrDefaultAsync(x => x.Name == "Tom");
```

- 期望最多只有一条
- 如果多于一条会抛异常

### 速记

- 不确定唯一性时优先 `FirstOrDefaultAsync`
- 明确业务上必须唯一时才考虑 `SingleOrDefaultAsync`

---

## 22. Controller 中使用 DbContext

### 示例

```csharp
[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase
{
    private readonly AppDbContext _context;

    public UsersController(AppDbContext context)
    {
        _context = context;
    }

    [HttpGet]
    public async Task<IActionResult> GetAll()
    {
        var users = await _context.Users.ToListAsync();
        return Ok(users);
    }
}
```

### 说明

入门阶段这样写没问题。
但项目稍大后，更推荐：

- Controller -> Service -> DbContext

### 速记

- 小例子里可直接注入 `DbContext`
- 正式项目通常通过 Service 包一层

---

## 23. Service 层使用 DbContext

### 示例

```csharp
public interface IUserService
{
    Task<List<UserDto>> GetAllAsync();
}

public class UserService : IUserService
{
    private readonly AppDbContext _context;

    public UserService(AppDbContext context)
    {
        _context = context;
    }

    public async Task<List<UserDto>> GetAllAsync()
    {
        return await _context.Users
            .Select(x => new UserDto
            {
                Id = x.Id,
                Name = x.Name
            })
            .ToListAsync();
    }
}
```

### 说明

这种结构更清晰：

- Controller 负责 HTTP
- Service 负责业务
- DbContext 负责数据访问

### 速记

- 中大型项目优先分层，不把所有逻辑堆在 Controller

---

## 24. 导航属性基础认知

两个实体之间常有关系，比如：

- 一个用户有多个订单

### 示例

```csharp
public class User
{
    public int Id { get; set; }
    public string Name { get; set; } = "";

    public List<Order> Orders { get; set; } = new();
}

public class Order
{
    public int Id { get; set; }
    public string ProductName { get; set; } = "";

    public int UserId { get; set; }
    public User? User { get; set; }
}
```

### 说明

- `User.Orders`：导航到订单集合
- `Order.User`：导航到所属用户
- `UserId`：外键

### 速记

- 导航属性用于表达表关系
- 外键字段和导航属性常一起出现

---

## 25. Include 基础认知

如果你想把关联数据一起查出来，可以用 `Include`。

### 示例

```csharp
var users = await _context.Users
    .Include(x => x.Orders)
    .ToListAsync();
```

### 说明

这表示查询用户时，把关联订单也加载出来。

### 速记

- 查关联数据常用 `Include`
- 入门先理解概念，不必一开始就深挖复杂关系映射

---

## 26. AsNoTracking 基础认知

### 示例

```csharp
var users = await _context.Users
    .AsNoTracking()
    .ToListAsync();
```

### 说明

如果你只是读数据，不准备修改并保存，`AsNoTracking()` 常能减少跟踪开销。

### 速记

- 只读查询可考虑 `AsNoTracking()`

---

## 27. 一个最小完整 CRUD 示例

### Entity

```csharp
public class User
{
    public int Id { get; set; }
    public string Name { get; set; } = "";
    public int Age { get; set; }
}
```

### DbContext

```csharp
public class AppDbContext : DbContext
{
    public AppDbContext(DbContextOptions<AppDbContext> options) : base(options)
    {
    }

    public DbSet<User> Users { get; set; }
}
```

### Controller

```csharp
[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase
{
    private readonly AppDbContext _context;

    public UsersController(AppDbContext context)
    {
        _context = context;
    }

    [HttpGet]
    public async Task<IActionResult> GetAll()
    {
        var users = await _context.Users.ToListAsync();
        return Ok(users);
    }

    [HttpGet("{id}")]
    public async Task<IActionResult> GetById(int id)
    {
        var user = await _context.Users.FindAsync(id);
        if (user == null)
        {
            return NotFound();
        }

        return Ok(user);
    }

    [HttpPost]
    public async Task<IActionResult> Create(User user)
    {
        _context.Users.Add(user);
        await _context.SaveChangesAsync();
        return Ok(user);
    }

    [HttpPut("{id}")]
    public async Task<IActionResult> Update(int id, User input)
    {
        var user = await _context.Users.FindAsync(id);
        if (user == null)
        {
            return NotFound();
        }

        user.Name = input.Name;
        user.Age = input.Age;
        await _context.SaveChangesAsync();
        return Ok(user);
    }

    [HttpDelete("{id}")]
    public async Task<IActionResult> Delete(int id)
    {
        var user = await _context.Users.FindAsync(id);
        if (user == null)
        {
            return NotFound();
        }

        _context.Users.Remove(user);
        await _context.SaveChangesAsync();
        return NoContent();
    }
}
```

---

## 28. 常见误区

### 误区 1：改了实体类，数据库会自动跟着变

不会。
通常你还要生成 Migration 并执行更新。

### 误区 2：`Add()` 之后数据已经进数据库了

不对。
真正写入数据库要靠 `SaveChanges()` / `SaveChangesAsync()`。

### 误区 3：EF Core 就是不写 SQL

不准确。
你是少写 SQL，但本质还是在生成 SQL。
所以要理解查询行为。

### 误区 4：API 直接返回 Entity 永远没问题

不推荐长期这样做。
正式项目更推荐 DTO。

### 误区 5：数据库查询都可以同步写

Web 项目里通常优先异步，尤其是高并发场景。

---

## 29. 本课重点总结

### 你必须先掌握的 10 个核心点

1. EF Core 是 .NET 主流 ORM 之一
2. Entity 对应数据库表
3. `DbContext` 是数据库操作核心入口
4. `DbSet<T>` 表示表的操作集合
5. 数据库连接通常在 `Program.cs` 注册
6. Migration 用来管理数据库结构变更
7. `Add` / `FindAsync` / `Remove` / `SaveChangesAsync` 是 CRUD 基础
8. EF Core 可以把 LINQ 翻译为 SQL
9. 查询接口数据时常用 `Select -> DTO`
10. Web 项目里数据库操作优先异步方法

---

## 30. Java / Go / Python 对比总结

| 知识点 | C# / EF Core | Java / JPA | Go / GORM | Python / SQLAlchemy |
|---|---|---|---|---|
| ORM 核心入口 | `DbContext` | `EntityManager` / Repository | `DB` 对象 | Session |
| 实体类 | Entity | Entity | Model/Struct | Model |
| 迁移 | Migration | Flyway/Liquibase/JPA 生态 | GORM migrate | Alembic |
| 查询方式 | LINQ | JPQL / Criteria / Repository | 链式 API | ORM Query |
| 异步支持 | Web 中高频 async | 传统同步更多 | 同步常见 | 框架不同 |

---

## 31. 面试 / 实战高频提醒

- `DbContext` 和 `DbSet<T>` 必须牢固掌握
- `SaveChangesAsync()` 是数据落库关键
- 迁移命令要会：`migrations add`、`database update`
- 查询时优先异步 API
- 列表接口常用 `Where + OrderBy + Select + ToListAsync`
- 正式项目尽量别让 Controller 直接堆满 EF Core 细节

---

## 32. 本课小练习

### 练习 1：定义 Entity 和 DbContext

要求：

- 定义 `Product` 实体，包含 `Id`、`Name`、`Price`
- 定义 `AppDbContext`
- 加入 `DbSet<Product>`

### 练习 2：配置数据库连接

要求：

- 在 `appsettings.json` 写连接串
- 在 `Program.cs` 注册 `AppDbContext`

### 练习 3：新增和查询

要求：

- 向数据库新增一个 `Product`
- 查询全部产品并输出

### 练习 4：更新和删除

要求：

- 根据 `id` 查询产品
- 修改价格并保存
- 删除该产品并保存

### 练习 5：LINQ 查询

要求：

- 查询价格大于 100 的产品
- 按价格倒序排序
- 投影成只包含 `Name` 和 `Price` 的 DTO

### 练习 6：迁移命令

要求：

- 执行一次 `dotnet ef migrations add InitialCreate`
- 执行一次 `dotnet ef database update`
- 理解这两条命令分别做什么

---

## 33. 一页速背版

### Entity

```csharp
public class User
{
    public int Id { get; set; }
    public string Name { get; set; } = "";
}
```

### DbContext

```csharp
public class AppDbContext : DbContext
{
    public AppDbContext(DbContextOptions<AppDbContext> options) : base(options) { }
    public DbSet<User> Users { get; set; }
}
```

### 注册

```csharp
builder.Services.AddDbContext<AppDbContext>(options =>
    options.UseSqlServer(builder.Configuration.GetConnectionString("DefaultConnection")));
```

### CRUD

```csharp
_context.Users.Add(user);
await _context.SaveChangesAsync();

var user = await _context.Users.FindAsync(id);

_context.Users.Remove(user);
await _context.SaveChangesAsync();
```

### LINQ

```csharp
var users = await _context.Users
    .Where(x => x.Age >= 18)
    .Select(x => new UserDto { Id = x.Id, Name = x.Name })
    .ToListAsync();
```

### 迁移命令

```bash
dotnet ef migrations add InitialCreate
dotnet ef database update
```

### 核心结论

- Entity 对应表
- `DbContext` 管数据库
- `DbSet<T>` 管表
- 查询写 LINQ
- 结构变更靠 Migration
- 数据落库靠 `SaveChangesAsync()`

---

## 34. 下一课预告

下一课学习：

- C# 项目实战建议
- 分层架构
- Controller / Service / Repository 设计
- 日志、配置、异常处理中间件
- 常见开发规范
- 如何从 Java / Go 背景平滑转成 C# 工程思维

这课会帮你把前面的知识真正串成“能干活”的体系。
