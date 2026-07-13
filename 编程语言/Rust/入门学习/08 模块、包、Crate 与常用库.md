# 08 模块、包、Crate 与常用库

学会写单个文件之后，就要理解 Rust 项目怎么组织。

## 几个概念

### Package

一个 Cargo 项目就是一个 package，里面有 `Cargo.toml`。

### Crate

crate 是 Rust 的编译单元。

常见两种：

- binary crate：可执行程序，有 `main.rs`
- library crate：库，有 `lib.rs`

### Module

module 是代码组织单元，用 `mod` 声明。

## 基本项目结构

```text
my_app/
├── Cargo.toml
└── src/
    ├── main.rs
    ├── lib.rs
    └── utils.rs
```

## 在 main.rs 中使用模块

`src/main.rs`：

```rust
mod utils;

fn main() {
    utils::hello();
}
```

`src/utils.rs`：

```rust
pub fn hello() {
    println!("hello from utils");
}
```

注意：函数要加 `pub` 才能被模块外访问。

## 子模块目录写法

```text
src/
├── main.rs
└── user/
    ├── mod.rs
    └── service.rs
```

`main.rs`：

```rust
mod user;

fn main() {
    user::service::create_user();
}
```

`user/mod.rs`：

```rust
pub mod service;
```

`user/service.rs`：

```rust
pub fn create_user() {
    println!("create user");
}
```

## use 引入路径

```rust
use std::collections::HashMap;

fn main() {
    let mut map = HashMap::new();
    map.insert("a", 1);
}
```

也可以给别名：

```rust
use std::collections::HashMap as Map;
```

## pub 控制可见性

默认私有：

```rust
fn private_fn() {}
```

公开：

```rust
pub fn public_fn() {}
```

结构体字段也默认私有：

```rust
pub struct User {
    pub name: String,
    age: u32,
}
```

这里 `User` 和 `name` 公开，但 `age` 私有。

## 添加依赖

在 `Cargo.toml`：

```toml
[dependencies]
serde = { version = "1", features = ["derive"] }
serde_json = "1"
```

## 常用库推荐

### serde / serde_json

序列化和反序列化，处理 JSON 必备。

```rust
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
struct User {
    name: String,
    age: u32,
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let user = User { name: "老王".to_string(), age: 18 };
    let json = serde_json::to_string(&user)?;
    println!("{}", json);
    Ok(())
}
```

### clap

命令行参数解析。

```toml
clap = { version = "4", features = ["derive"] }
```

```rust
use clap::Parser;

#[derive(Parser)]
struct Args {
    #[arg(short, long)]
    name: String,
}

fn main() {
    let args = Args::parse();
    println!("hello {}", args.name);
}
```

运行：

```bash
cargo run -- --name 老王
```

### tokio

异步运行时，写网络服务常用。

```toml
tokio = { version = "1", features = ["full"] }
```

```rust
#[tokio::main]
async fn main() {
    println!("async hello");
}
```

### reqwest

HTTP 客户端。

```toml
reqwest = { version = "0.12", features = ["json"] }
tokio = { version = "1", features = ["full"] }
```

```rust
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let text = reqwest::get("https://www.rust-lang.org").await?.text().await?;
    println!("{}", &text[..100]);
    Ok(())
}
```

### anyhow

应用程序错误处理。

```toml
anyhow = "1"
```

### tracing

日志系统，比 `println!` 更适合正式项目。

```toml
tracing = "0.1"
tracing-subscriber = "0.3"
```

## crates.io

Rust 官方包仓库：

https://crates.io

查库文档：

https://docs.rs

一般流程：

1. 在 crates.io 搜库
2. 看下载量、更新时间、README
3. 在 docs.rs 看 API 文档
4. 加到 `Cargo.toml`
5. `cargo check`

## 本节练习

创建一个项目：

```bash
cargo new json_demo
cd json_demo
```

添加依赖：

```toml
serde = { version = "1", features = ["derive"] }
serde_json = "1"
```

把结构体转 JSON 打印出来。

## 本节小结

- `Cargo.toml` 管依赖
- `mod` 划分模块
- `pub` 控制公开
- `use` 引入路径
- `crates.io` 找库，`docs.rs` 看文档
