---
title: Crow crow.h C++ Web 框架资料整理
date: 2026-06-16
tags:
  - Cpp
  - Web
  - HTTP
  - REST
  - Crow
aliases:
  - crow.h
  - CrowCpp
---

# Crow（crow.h）C++ Web 框架资料整理

Crow 是一个 C++ 轻量级 Web 微框架，常被叫做 `crow.h`。它的写法接近 Python Flask：用宏注册路由，用 lambda 处理请求，适合快速写 HTTP API、REST 服务、小型后台、WebSocket 服务。

官方仓库：<https://github.com/CrowCpp/Crow>  
官方文档：<https://crowcpp.org>

## 一句话理解

> Crow = C++ 里的“Flask 风格”Web 框架。

特点：

- 路由简单：`CROW_ROUTE(app, "/path")([] { ... });`
- 支持 URL 参数类型检查，如 `<int>`、`<string>`。
- 内置 JSON：`crow::json::wvalue`、`crow::json::load`。
- 支持 middleware。
- 支持 WebSocket。
- 可以作为单头文件使用，也可通过 CMake / vcpkg / Conan 集成。
- 基于现代 C++，常见项目里通常用 C++11/14 起步。

## 适合用在什么场景

适合：

- 写 C++ 小型 HTTP API。
- 给已有 C++ 程序暴露控制接口。
- 写本地工具的 Web 管理端。
- 做高性能但结构不太复杂的 REST 服务。
- 做 WebSocket 通信 Demo 或轻量服务。

不太适合：

- 大型企业后端全家桶项目。
- 需要 ORM、权限、迁移、任务队列等完整生态的场景。
- 非常复杂的异步业务编排。
- 希望框架替你处理所有工程结构和约定的场景。

如果要做更大型 C++ Web 服务，可以顺便比较：

- `Drogon`：更完整，工程化能力更强。
- `oatpp`：偏 REST API 生态。
- `cpp-httplib`：更轻，适合简单 HTTP server/client。

## 最小 Hello World

```cpp
#include "crow.h"

int main()
{
    crow::SimpleApp app;

    CROW_ROUTE(app, "/")([](){
        return "Hello world";
    });

    app.port(18080).multithreaded().run();
}
```

访问：

```bash
curl http://127.0.0.1:18080/
```

说明：

- `crow::SimpleApp app;` 创建应用。
- `CROW_ROUTE(app, "/")` 注册路径。
- `([](){ ... })` 是处理函数。
- `port(18080)` 设置端口。
- `multithreaded()` 开启多线程处理。
- `run()` 启动服务。

## 安装与引入方式

### 方式 1：单头文件

适合学习、Demo、小工具。

```cpp
#include "crow.h"
```

把 `crow.h` 放到项目 include 路径里即可。

注意：旧版本 Crow v0.3 需要在且仅在一个源文件顶部写：

```cpp
#define CROW_MAIN
#include "crow.h"
```

新版本一般不需要这样写，按当前使用版本文档为准。

### 方式 2：CMake 集成

常见写法示意：

```cmake
cmake_minimum_required(VERSION 3.15)
project(crow_demo)

set(CMAKE_CXX_STANDARD 17)

add_executable(crow_demo main.cpp)
target_include_directories(crow_demo PRIVATE path/to/crow/include)
```

如果用包管理器，优先查当前环境支持：

- `vcpkg install crow`
- `conan` 对应 Crow 包
- 系统包管理器或源码安装

## 路由基础

### GET 路由

```cpp
CROW_ROUTE(app, "/hello")([](){
    return "hello crow";
});
```

### URL 参数

```cpp
CROW_ROUTE(app, "/hello/<string>")
([](const std::string& name){
    return "hello " + name;
});
```

数字参数：

```cpp
CROW_ROUTE(app, "/square/<int>")
([](int n){
    return std::to_string(n * n);
});
```

Crow 会对 handler 参数做编译期类型检查。比如路由里只有一个 `<int>`，lambda 却写两个参数，会编译报错。

### 指定 HTTP 方法

```cpp
CROW_ROUTE(app, "/users")
.methods("POST"_method)
([](const crow::request& req){
    return crow::response(201, "created");
});
```

常用方法：

```cpp
.methods("GET"_method)
.methods("POST"_method)
.methods("PUT"_method)
.methods("DELETE"_method)
```

也可以使用 Crow 的方法枚举或宏写法，具体以版本 API 为准。

## 请求与响应

### 获取请求体

```cpp
CROW_ROUTE(app, "/echo")
.methods("POST"_method)
([](const crow::request& req){
    return crow::response(req.body);
});
```

### 设置状态码

```cpp
return crow::response(400, "bad request");
```

常见状态码：

- `200`：成功。
- `201`：创建成功。
- `204`：成功但无内容。
- `400`：请求参数错误。
- `401`：未认证。
- `403`：无权限。
- `404`：不存在。
- `500`：服务端错误。

### 设置响应头

```cpp
crow::response res;
res.code = 200;
res.set_header("Content-Type", "text/plain; charset=utf-8");
res.body = "ok";
return res;
```

## JSON 使用

Crow 内置 JSON 支持。

### 返回 JSON

```cpp
CROW_ROUTE(app, "/json")
([]{
    crow::json::wvalue data;
    data["message"] = "Hello, World!";
    data["ok"] = true;
    data["count"] = 123;
    return data;
});
```

### 解析 JSON 请求

```cpp
CROW_ROUTE(app, "/add")
.methods("POST"_method)
([](const crow::request& req){
    auto body = crow::json::load(req.body);
    if (!body) {
        return crow::response(400, "invalid json");
    }

    int a = body["a"].i();
    int b = body["b"].i();

    crow::json::wvalue result;
    result["sum"] = a + b;
    return crow::response(result.dump());
});
```

更稳妥的实际写法要检查字段是否存在、类型是否正确，避免客户端传错数据导致异常或未定义行为。

## 一个完整 REST 示例

```cpp
#include "crow.h"
#include <unordered_map>
#include <string>

int main()
{
    crow::SimpleApp app;
    std::unordered_map<int, std::string> users;

    CROW_ROUTE(app, "/users/<int>")
    .methods("GET"_method)
    ([&users](int id){
        auto it = users.find(id);
        if (it == users.end()) {
            return crow::response(404, "user not found");
        }

        crow::json::wvalue data;
        data["id"] = id;
        data["name"] = it->second;
        return crow::response(data.dump());
    });

    CROW_ROUTE(app, "/users/<int>")
    .methods("POST"_method)
    ([&users](const crow::request& req, int id){
        auto body = crow::json::load(req.body);
        if (!body || !body.has("name")) {
            return crow::response(400, "missing name");
        }

        users[id] = body["name"].s();
        return crow::response(201, "created");
    });

    app.port(18080).multithreaded().run();
}
```

测试：

```bash
curl -X POST http://127.0.0.1:18080/users/1 \
  -H 'Content-Type: application/json' \
  -d '{"name":"laowang"}'

curl http://127.0.0.1:18080/users/1
```

注意：上面用 `unordered_map` 只是 Demo。开启 `multithreaded()` 后，共享数据需要考虑线程安全，实际项目应加锁或使用线程安全的数据层。

## Query 参数

Crow 的请求对象可以读取 URL 参数。

```cpp
CROW_ROUTE(app, "/search")
([](const crow::request& req){
    const char* q = req.url_params.get("q");
    if (!q) {
        return crow::response(400, "missing q");
    }
    return crow::response(std::string("query = ") + q);
});
```

访问：

```bash
curl 'http://127.0.0.1:18080/search?q=cpp'
```

## 静态文件

Crow 支持静态路由，官方 API 里有：

- `route_static(url)`
- `static_file(url, internalPath)`

不同版本用法略有差异，写静态资源服务前应查当前版本文档。

如果要对外提供大量静态资源，更推荐用 Nginx / Caddy 处理静态文件，Crow 专注 API。

## Middleware 中间件

中间件用于在请求前后执行逻辑，例如日志、鉴权、CORS、统计耗时。

示意结构：

```cpp
struct ExampleMiddleware
{
    struct context
    {
    };

    void before_handle(crow::request& req, crow::response& res, context& ctx)
    {
        // 请求进入路由前执行
    }

    void after_handle(crow::request& req, crow::response& res, context& ctx)
    {
        // 路由处理完成后执行
    }
};

int main()
{
    crow::App<ExampleMiddleware> app;

    CROW_ROUTE(app, "/")([]{
        return "ok";
    });

    app.port(18080).run();
}
```

实际常见用途：

- 请求日志。
- Token 校验。
- CORS 响应头。
- 限流。
- 请求耗时统计。

## WebSocket

Crow 支持 WebSocket，可以注册连接、消息、关闭事件。

示意写法：

```cpp
CROW_ROUTE(app, "/ws")
.websocket()
.onopen([](crow::websocket::connection& conn){
    CROW_LOG_INFO << "websocket open";
})
.onmessage([](crow::websocket::connection& conn, const std::string& data, bool is_binary){
    conn.send_text("echo: " + data);
})
.onclose([](crow::websocket::connection& conn, const std::string& reason, uint16_t code){
    CROW_LOG_INFO << "websocket close";
});
```

官方 API 里还可以配置默认最大 payload：

```cpp
app.websocket_max_payload(1024 * 1024);
```

## 常用服务配置

### 端口

```cpp
app.port(18080).run();
```

### 绑定地址

```cpp
app.bindaddr("127.0.0.1").port(18080).run();
```

默认绑定地址通常是 `0.0.0.0`，会监听所有网卡。开发环境可绑定 `127.0.0.1` 更安全。

### 多线程

```cpp
app.port(18080).multithreaded().run();
```

多线程下要注意：

- lambda 捕获的共享变量要线程安全。
- 数据库连接最好用连接池或每线程连接。
- 日志、缓存、全局对象都要考虑并发访问。

### 超时

官方 API 支持：

```cpp
app.timeout(5);
```

默认连接超时时间一般是 5 秒，具体以版本为准。

### Server Header

官方 API 支持设置响应里的 `Server` header：

```cpp
app.server_name("my-crow-service");
```

如果设置为空字符串，部分版本会省略该 header。

## 常见坑

### 1. 版本差异

网上很多教程来自旧版 `ipkn/crow` 或 Crow v0.3。现在更常用的是 `CrowCpp/Crow`。如果遇到编译错误，先确认：

- include 的 `crow.h` 来自哪里。
- 当前 Crow 版本。
- 教程是否写给旧版。
- 是否需要 `#define CROW_MAIN`。

### 2. JSON 字段不校验

不要默认客户端一定传了字段：

```cpp
int a = body["a"].i();
```

实际项目要做：

- JSON 是否解析成功。
- 字段是否存在。
- 类型是否符合预期。
- 错误时返回清晰的 `400`。

### 3. 多线程共享数据

`multithreaded()` 很方便，但也容易踩并发问题。

危险示例：

```cpp
std::unordered_map<int, std::string> users;
```

如果多个请求同时读写，需要加锁：

```cpp
std::mutex users_mutex;
```

或者把状态放到数据库里。

### 4. 对外暴露服务

如果只是本机开发：

```cpp
app.bindaddr("127.0.0.1")
```

如果对外提供服务，建议前面放 Caddy / Nginx：

- 处理 HTTPS。
- 处理静态文件。
- 做反向代理。
- 做访问日志。
- 做限流和基本安全策略。

### 5. CORS

浏览器前端调用 Crow API 时可能遇到 CORS。可以用中间件统一加响应头，也可以在反向代理层处理。

常见 header：

```text
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Content-Type, Authorization
```

生产环境不要无脑 `*`，应限制为可信域名。

## 推荐工程结构

小项目：

```text
crow-demo/
├── CMakeLists.txt
├── include/
│   └── crow.h
└── src/
    └── main.cpp
```

稍复杂一点：

```text
crow-service/
├── CMakeLists.txt
├── include/
│   ├── routes/
│   ├── services/
│   └── middleware/
├── src/
│   ├── main.cpp
│   ├── routes/
│   ├── services/
│   └── middleware/
└── third_party/
    └── crow/
```

建议：

- `main.cpp` 只负责启动、注册路由、加载配置。
- 路由只负责 HTTP 输入输出。
- 业务逻辑放到 service 层。
- 数据库访问单独封装。
- 中间件单独放目录。

## 最小 CMake 示例

```cmake
cmake_minimum_required(VERSION 3.15)
project(crow_demo)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_executable(crow_demo src/main.cpp)
target_include_directories(crow_demo PRIVATE include)
```

构建：

```bash
cmake -S . -B build
cmake --build build
./build/crow_demo
```

## 实战模板：API 服务骨架

```cpp
#include "crow.h"

int main()
{
    crow::SimpleApp app;

    CROW_ROUTE(app, "/health")
    ([] {
        crow::json::wvalue res;
        res["status"] = "ok";
        return res;
    });

    CROW_ROUTE(app, "/api/echo")
    .methods("POST"_method)
    ([](const crow::request& req) {
        auto body = crow::json::load(req.body);
        if (!body) {
            return crow::response(400, "invalid json");
        }

        crow::json::wvalue res;
        res["received"] = body;
        return crow::response(res.dump());
    });

    app.bindaddr("127.0.0.1")
       .port(18080)
       .multithreaded()
       .run();
}
```

## 学习路线

1. 跑通 Hello World。
2. 学会 `CROW_ROUTE` 和 URL 参数。
3. 学会 GET / POST / JSON。
4. 写一个 CRUD Demo。
5. 加日志 middleware。
6. 加 CORS / Token 鉴权。
7. 接数据库。
8. 用 Caddy / Nginx 反代成 HTTPS 服务。

## 参考链接

- Crow GitHub：<https://github.com/CrowCpp/Crow>
- Crow 文档：<https://crowcpp.org>
- Crow API Reference：<https://crowcpp.org/master/reference/classcrow_1_1_crow.html>
- Crow examples：<https://github.com/CrowCpp/Crow/tree/master/examples>

## 速查

```cpp
crow::SimpleApp app;

CROW_ROUTE(app, "/")([]{
    return "ok";
});

CROW_ROUTE(app, "/hello/<string>")
([](const std::string& name){
    return "hello " + name;
});

CROW_ROUTE(app, "/json")
([]{
    crow::json::wvalue x;
    x["ok"] = true;
    return x;
});

CROW_ROUTE(app, "/post")
.methods("POST"_method)
([](const crow::request& req){
    auto body = crow::json::load(req.body);
    if (!body) return crow::response(400);
    return crow::response(200);
});

app.bindaddr("127.0.0.1")
   .port(18080)
   .multithreaded()
   .run();
```
