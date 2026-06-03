# C# 快速学习笔记总目录

> 面向有 Java / Go / Python 基础的开发者，目标是用最短路径掌握 C# 与 ASP.NET Core。

## 学习说明

这套笔记按 **语言基础 -> 集合与异步 -> 工程化 -> Web API -> EF Core -> 项目实战** 的顺序组织。

建议学习方式：

1. 每课先通读一遍
2. 把示例代码自己手敲一遍
3. 至少完成该课 2~3 个小练习
4. 学完第 8~10 课后，马上做一个 Web API 小项目巩固

---

## 课程目录

### 语言基础

1. [第一课：C# 基础语法快速入门](csharp-lesson-01.md)
   - 程序结构、变量、基本类型、`var`、字符串、流程控制、数组、方法

2. [第二课：C# 面向对象入门](csharp-lesson-02.md)
   - 类、对象、字段、属性、构造函数、`this`、`static`、`class` 和 `struct`

3. [第三课：C# 集合、泛型、Lambda、LINQ](csharp-lesson-03.md)
   - `List<T>`、`Dictionary<TKey, TValue>`、泛型、Lambda、LINQ 常见操作

4. [第四课：异常处理、委托、Action/Func、事件、async/await](csharp-lesson-04.md)
   - 异常、委托、事件、异步编程、`Task`

5. [第五课：文件操作、JSON、时间处理、命名空间、项目结构、NuGet、.NET CLI](csharp-lesson-05.md)
   - 文件读写、JSON 序列化、`DateTime`、命名空间、CLI 常用命令

6. [第六课：可空引用类型、record、模式匹配、switch expression、var/object/dynamic](csharp-lesson-06.md)
   - 可空引用类型、值/引用类型、`record`、模式匹配、现代 C# 语法

7. [第七课：继承、多态、抽象类、接口、依赖注入、面向接口编程](csharp-lesson-07.md)
   - OOP 设计、接口、抽象类、依赖注入、面向接口编程

### Web 与数据库

8. [第八课：ASP.NET Core 入门——Web API 基本结构、路由、Controller、请求响应、DTO、依赖注入](csharp-lesson-08.md)
   - Web API 项目结构、路由、Controller、DTO、DI

9. [第九课：EF Core 入门——DbContext、Entity、Migration、CRUD、LINQ 数据库查询](csharp-lesson-09.md)
   - EF Core、实体、数据库上下文、迁移、CRUD、数据库查询

10. [第十课：C# 项目实战与工程化思维](csharp-lesson-10.md)
    - 分层架构、日志、配置、异常处理中间件、工程实践

---

## 推荐学习顺序

### 路线 A：最快上手路线

适合你这种已有开发经验的人：

- 第 1 课
- 第 2 课
- 第 3 课
- 第 4 课
- 第 7 课
- 第 8 课
- 第 9 课
- 第 10 课

如果目的是尽快写后端，这条路线最快。

### 路线 B：完整系统路线

- 第 1 课 → 第 2 课 → 第 3 课 → 第 4 课 → 第 5 课 → 第 6 课 → 第 7 课 → 第 8 课 → 第 9 课 → 第 10 课

如果你想把 C# 的语言体系一起补完整，走这条。

---

## 学完后的能力对照

学完 1~3 课：

- 能写基础 C# 语法
- 能写集合处理与 LINQ

学完 4~7 课：

- 能理解异步、委托、接口、依赖注入
- 能看懂大部分常见 C# 业务代码

学完 8~10 课：

- 能搭建一个 ASP.NET Core Web API 项目
- 能接入 EF Core 做 CRUD
- 能写基础分层架构
- 能做日志、配置、异常处理中间件

---

## 建议你下一步怎么学

### 第一阶段：过一遍课程

- 每天 1~2 课
- 每课自己敲代码
- 每课至少做 2 个练习

### 第二阶段：做一个小项目

推荐项目：

- 用户管理系统 API
- 任务管理 API
- 博客后台 API

建议功能：

- 用户 CRUD
- 分页查询
- 登录 / JWT
- 全局异常处理
- 日志
- Swagger
- EF Core + SQLite / SQL Server

### 第三阶段：进阶专题

后续可继续学：

- JWT 鉴权
- 单元测试
- Redis
- Docker
- 部署
- 性能优化

---

## 配套实战路线图

已经为你额外生成：

- [C# Web API 实战路线图（7 天 / 14 天）](csharp-webapi-roadmap.md)

---

## 一句话建议

如果你的目标是快速转 C# 后端，不要只看语法，**学到第 8 课开始就要立刻做 Web API 项目**。
