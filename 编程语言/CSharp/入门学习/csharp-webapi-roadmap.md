# C# Web API 实战路线图（7 天 / 14 天）

> 目标：用一个完整的小项目，把你已经学过的 C#、ASP.NET Core、EF Core、分层架构真正串起来。

---

## 1. 实战项目建议

推荐项目名：**用户管理 Web API**

### 最小功能范围

- 用户列表
- 用户详情
- 创建用户
- 更新用户
- 删除用户
- 分页查询
- 关键字搜索
- 全局异常处理中间件
- Swagger
- 日志
- EF Core + SQLite / SQL Server

### 进阶可选功能

- 登录接口
- JWT 鉴权
- 角色字段
- 审计字段（创建时间、更新时间）
- Docker 打包
- 单元测试

---

## 2. 建议技术栈

### 必选

- .NET 8
- ASP.NET Core Web API
- EF Core
- SQLite 或 SQL Server
- Swagger

### 建议选择

如果你想快速跑通：

- **SQLite**：更轻便，适合本地练手

如果你想更接近企业项目：

- **SQL Server**：更贴近常见 .NET 企业场景

---

## 3. 建议项目结构

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
│   ├── UpdateUserDto.cs
│   ├── UserDto.cs
│   └── UserQueryDto.cs
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

---

## 4. User 实体建议字段

```csharp
public class User
{
    public int Id { get; set; }
    public string Name { get; set; } = "";
    public string Email { get; set; } = "";
    public int Age { get; set; }
    public DateTime CreatedAt { get; set; }
    public DateTime UpdatedAt { get; set; }
}
```

### 为什么这么设计

- `Id`：主键
- `Name` / `Email`：基础信息
- `Age`：演示校验和查询条件
- `CreatedAt` / `UpdatedAt`：练时间字段和审计字段

---

## 5. DTO 建议

### CreateUserDto

```csharp
public class CreateUserDto
{
    public string Name { get; set; } = "";
    public string Email { get; set; } = "";
    public int Age { get; set; }
}
```

### UpdateUserDto

```csharp
public class UpdateUserDto
{
    public string Name { get; set; } = "";
    public string Email { get; set; } = "";
    public int Age { get; set; }
}
```

### UserDto

```csharp
public class UserDto
{
    public int Id { get; set; }
    public string Name { get; set; } = "";
    public string Email { get; set; } = "";
    public int Age { get; set; }
}
```

### UserQueryDto

```csharp
public class UserQueryDto
{
    public string? Keyword { get; set; }
    public int PageIndex { get; set; } = 1;
    public int PageSize { get; set; } = 10;
}
```

---

## 6. 7 天实战路线

适合你这种已有开发经验、想快速跑通一套 C# Web API 的情况。

### Day 1：搭项目骨架

目标：把项目跑起来。

任务：

- 创建 Web API 项目
- 熟悉 `Program.cs`
- 建立目录结构
- 创建第一个 `UsersController`
- 跑通 Swagger

验收标准：

- 能启动项目
- Swagger 可访问
- `/api/users` 能返回测试数据

建议命令：

```bash
dotnet new webapi -n DemoApi
dotnet run
```

---

### Day 2：补 DTO、Service、Repository 分层

目标：不要让 Controller 直接堆逻辑。

任务：

- 新建 `Dtos/`
- 新建 `Services/`
- 新建 `Repositories/`
- 实现 `IUserService` / `UserService`
- Controller 改为调用 Service

验收标准：

- Controller 只保留接参、调 Service、回响应
- 至少完成一个查询接口和一个创建接口

---

### Day 3：接入 EF Core

目标：打通数据库。

任务：

- 安装 EF Core 包
- 创建 `User` 实体
- 创建 `AppDbContext`
- 配置连接串
- 注册 `DbContext`
- 创建第一次 Migration
- 更新数据库

验收标准：

- 数据库成功建表
- 能查到 `Users` 表
- 项目可以正常启动

建议命令：

```bash
dotnet add package Microsoft.EntityFrameworkCore.Sqlite
dotnet add package Microsoft.EntityFrameworkCore.Design
dotnet ef migrations add InitialCreate
dotnet ef database update
```

---

### Day 4：完成用户 CRUD

目标：把基础接口做完整。

任务：

- 查询列表
- 查询详情
- 创建用户
- 更新用户
- 删除用户
- 用 DTO 组织请求和响应

验收标准：

- 5 个基础 CRUD 接口可在 Swagger 中调通
- 能正确返回 200 / 404 / 400 等状态码

---

### Day 5：做分页、搜索、排序

目标：让接口更像真实业务接口。

任务：

- 增加 `UserQueryDto`
- 支持关键字查询
- 支持分页 `PageIndex/PageSize`
- 支持按创建时间或 Id 排序
- 使用 `Where + OrderBy + Skip + Take`

验收标准：

- 列表接口支持分页
- 支持关键词搜索
- 代码放在 Service 层实现

---

### Day 6：补日志、配置、异常处理中间件

目标：引入工程化能力。

任务：

- 在 Service 注入 `ILogger<T>`
- 新增全局异常处理中间件
- 把连接串和应用配置放入 `appsettings.json`
- 统一错误返回结构

验收标准：

- 报错时能统一返回 JSON
- 控制台能看到结构化日志
- 不再每个接口都手写重复 try/catch

---

### Day 7：收尾 + 可选 JWT

目标：把项目整理成可展示作品。

任务：

- 整理目录结构
- 优化 DTO 命名
- 检查接口返回是否统一
- 补一个 README 说明项目结构和接口
- 可选：增加简单登录接口和 JWT 鉴权

验收标准：

- 项目结构清晰
- Swagger 可完整演示
- 你能自己讲清楚每一层职责

---

## 7. 14 天实战路线

适合你想学得更稳一点，边做边消化每个工程点。

### 第 1 天：创建项目、理解 Program.cs
- 创建项目
- 运行 Swagger
- 了解请求处理流程

### 第 2 天：写第一个 Controller
- `UsersController`
- `GET /api/users`
- `GET /api/users/{id}`

### 第 3 天：抽 DTO
- 建 `CreateUserDto`
- 建 `UserDto`
- 区分请求模型和响应模型

### 第 4 天：抽 Service
- 写 `IUserService`
- Controller 改为依赖接口
- 熟悉 DI

### 第 5 天：定义 Entity 和 DbContext
- 建 `User` 实体
- 建 `AppDbContext`
- 注册数据库上下文

### 第 6 天：EF Core Migration
- 安装包
- 生成 Migration
- 更新数据库
- 看懂建表结果

### 第 7 天：实现创建和查询
- `POST /api/users`
- `GET /api/users`
- `GET /api/users/{id}`

### 第 8 天：实现更新和删除
- `PUT /api/users/{id}`
- `DELETE /api/users/{id}`
- 统一 404 处理

### 第 9 天：做分页和搜索
- `UserQueryDto`
- `Keyword`
- `PageIndex`
- `PageSize`

### 第 10 天：加日志
- 注入 `ILogger<T>`
- 记录创建、更新、删除日志
- 学结构化日志

### 第 11 天：加异常处理中间件
- 捕获未处理异常
- 统一错误返回结构
- 不再到处 try/catch

### 第 12 天：整理配置
- 连接串放 `appsettings.json`
- 可选：绑定 Options 配置类

### 第 13 天：加认证（可选但推荐）
- 增加登录接口
- 返回 JWT
- 给查询接口加鉴权

### 第 14 天：做项目复盘
- 清理目录结构
- 完善 README
- 自己讲一遍项目架构
- 列出还能优化的点

---

## 8. 每天都要做的 4 件事

1. **手敲代码，不只看**
2. **每做完一个接口，就用 Swagger 调一次**
3. **每新增一个分层，就问自己：职责清楚吗**
4. **每天写一小段复盘**

建议复盘模板：

```md
## 今日复盘
- 今天完成了什么
- 卡在哪里
- 新学会了什么
- 明天要做什么
```

---

## 9. 每个阶段的验收清单

### 基础完成

- [ ] 项目可以启动
- [ ] Swagger 正常
- [ ] 有 `UsersController`
- [ ] 有 DTO / Service / Entity / DbContext

### CRUD 完成

- [ ] 查询列表
- [ ] 查询详情
- [ ] 创建用户
- [ ] 更新用户
- [ ] 删除用户

### 工程化完成

- [ ] 日志注入完成
- [ ] 全局异常处理中间件完成
- [ ] 配置从 `appsettings.json` 读取
- [ ] 接口返回风格统一

### 进阶完成

- [ ] 分页查询完成
- [ ] 搜索完成
- [ ] JWT 登录完成（可选）
- [ ] 项目文档完成

---

## 10. 你做项目时最容易踩的坑

### 坑 1：Controller 写太厚

解决：

- Controller 只处理 HTTP
- 业务逻辑放 Service

### 坑 2：直接返回 Entity

解决：

- 用 DTO 映射输出

### 坑 3：忘记 `SaveChangesAsync()`

解决：

- 牢记 EF Core 改动不会自动落库

### 坑 4：改了实体没做 Migration

解决：

- 改模型后及时 `migrations add` + `database update`

### 坑 5：每个接口都 try/catch

解决：

- 用统一异常处理中间件

### 坑 6：日志全用字符串拼接

解决：

- 用结构化日志参数占位写法

---

## 11. 如果你只做一个项目，我建议你做到这一步

最小可展示成果：

- 一个能运行的 ASP.NET Core Web API
- 用户 CRUD
- EF Core + SQLite/SQL Server
- Swagger
- DTO / Service / Repository 分层
- 全局异常处理中间件
- 结构化日志
- 分页查询

如果你能把这套完整做出来，并且自己讲清楚：

- 为什么要 DTO
- 为什么要 Service
- 为什么要 DI
- 为什么要 Migration
- 为什么要异常处理中间件

那你已经具备很扎实的 C# Web 入门能力了。

---

## 12. 下一步进阶建议

做完这条路线后，再按顺序继续：

1. JWT 鉴权
2. 单元测试
3. Redis
4. Docker
5. 部署
6. 性能优化
7. 消息队列 / 后台任务

---

## 13. 一句话执行建议

如果你想最快把 C# 学成能干活的程度：

> 用 **7 天路线** 先快速跑通一个项目，再用 **14 天路线** 补齐工程细节和理解深度。
