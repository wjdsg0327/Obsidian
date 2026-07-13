# 02 安装 Rust 与 Cargo 基本用法

Rust 官方推荐用 `rustup` 安装。`rustup` 负责管理 Rust 编译器、标准库、工具链版本。

## 安装 Rust

Linux / macOS / WSL：

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

Windows：

- 访问 https://www.rust-lang.org/tools/install
- 下载 `rustup-init.exe`
- 按提示安装

安装后检查：

```bash
rustc --version
cargo --version
rustup --version
```

## 三个重要工具

### rustc

Rust 编译器。

```bash
rustc main.rs
```

可以直接编译单个文件，但日常开发很少直接用它。

### cargo

Rust 的项目管理工具，最常用。

它负责：

- 创建项目
- 构建项目
- 运行项目
- 下载依赖
- 运行测试
- 发布 crate

### rustup

管理 Rust 工具链。

常用命令：

```bash
rustup update       # 更新 Rust
rustup show         # 查看当前工具链
rustup component add rustfmt clippy
```

## 创建第一个项目

```bash
cargo new hello_rust
cd hello_rust
cargo run
```

目录结构：

```text
hello_rust/
├── Cargo.toml
└── src/
    └── main.rs
```

`Cargo.toml` 是项目配置文件：

```toml
[package]
name = "hello_rust"
version = "0.1.0"
edition = "2021"

[dependencies]
```

`src/main.rs` 是入口文件：

```rust
fn main() {
    println!("Hello, world!");
}
```

## Cargo 常用命令

```bash
cargo new 项目名        # 创建新项目
cargo run               # 编译并运行
cargo build             # 编译 debug 版本
cargo build --release   # 编译 release 优化版本
cargo check             # 只检查能不能编译，不生成程序，速度快
cargo test              # 运行测试
cargo fmt               # 自动格式化代码
cargo clippy            # 代码检查，给出改进建议
cargo clean             # 清理构建产物
```

## debug 和 release 的区别

```bash
cargo build
```

生成 debug 版本：

- 编译快
- 运行慢一点
- 方便调试

```bash
cargo build --release
```

生成 release 版本：

- 编译慢
- 运行快
- 发布时使用

生成文件位置：

```text
target/debug/项目名
target/release/项目名
```

## 添加第三方依赖

例如添加随机数库 `rand`：

```toml
[dependencies]
rand = "0.8"
```

然后代码里使用：

```rust
use rand::Rng;

fn main() {
    let n = rand::thread_rng().gen_range(1..=10);
    println!("随机数: {}", n);
}
```

也可以用命令添加：

```bash
cargo add rand
```

如果没有 `cargo add`，安装：

```bash
cargo install cargo-edit
```

## 本节练习

1. 安装 Rust
2. 创建 `hello_rust`
3. 修改 `main.rs` 输出自己的名字
4. 执行：

```bash
cargo fmt
cargo check
cargo run
```

## 本节小结

Rust 项目基本都围绕 Cargo 展开。新手阶段最常用三个命令：

```bash
cargo new
cargo check
cargo run
```
