# 01 Rust 是什么，适合做什么

Rust 是一门系统级编程语言，目标是同时做到：

- 运行速度接近 C/C++
- 内存安全，不需要垃圾回收 GC
- 并发安全，减少多线程数据竞争
- 包管理和工程体验现代化

一句话理解：

> Rust 想解决 C/C++ 很快但容易写出内存 bug 的问题，同时又不想像 Java/C# 那样依赖 GC。

## Rust 适合做什么

Rust 很适合：

- 命令行工具 CLI
- 后端服务
- 网络服务
- 高性能组件
- 嵌入式
- WebAssembly
- 数据处理工具
- 替代部分 C/C++ 的底层模块

典型项目：

- `ripgrep`：非常快的文本搜索工具
- `fd`：更好用的 find
- `deno`：JavaScript/TypeScript 运行时
- `ruff`：Python linter/formatter
- `tauri`：桌面应用框架

## Rust 不太适合什么

Rust 不是所有场景都最舒服：

- 快速写一次性脚本：Python 更快
- 传统企业 CRUD：Java/C# 生态更成熟
- 新手只想轻松入门编程：Python 更适合
- 极度依赖成熟 GUI 生态：Rust 还在发展

## Rust 和其它语言对比

| 语言 | 优点 | 代价 |
|---|---|---|
| C/C++ | 快、底层能力强 | 内存 bug 多，工程复杂 |
| Java/C# | 生态成熟，开发效率高 | 有 GC，底层控制弱一些 |
| Go | 简洁，部署方便，并发好用 | 表达力和底层控制不如 Rust |
| Python | 入门快，生态广 | 性能较弱，类型和工程约束弱 |
| Rust | 快、安全、现代工程 | 学习曲线陡 |

## Rust 的核心关键词

学 Rust，要围绕这些词：

- `ownership`：所有权
- `borrow`：借用
- `lifetime`：生命周期
- `trait`：类似接口，但更强
- `enum`：非常强大的枚举
- `match`：模式匹配
- `Result` / `Option`：显式错误处理和空值处理
- `Cargo`：包管理和构建工具

## 第一个 Rust 程序长什么样

```rust
fn main() {
    println!("Hello, Rust!");
}
```

注意：

- `fn` 定义函数
- `main` 是程序入口
- `println!` 后面有 `!`，说明它是宏 macro，不是普通函数
- 每条语句通常用 `;` 结尾

## 本节小结

Rust 是一门偏硬核但很值得学的语言。不要一上来就被生命周期吓到，先按顺序学：环境 → 基础语法 → 所有权 → 常用类型 → 项目实战。
