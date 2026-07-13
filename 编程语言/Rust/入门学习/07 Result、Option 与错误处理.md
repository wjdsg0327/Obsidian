# 07 Result、Option 与错误处理

Rust 不鼓励到处抛异常，而是把错误作为返回值显式处理。

最重要的两个类型：

- `Option<T>`：可能有值，也可能没有
- `Result<T, E>`：可能成功，也可能失败

## Option：处理空值

```rust
fn first_char(s: &str) -> Option<char> {
    s.chars().next()
}

fn main() {
    match first_char("Rust") {
        Some(c) => println!("第一个字符: {}", c),
        None => println!("空字符串"),
    }
}
```

常用方法：

```rust
let x = Some(10);
println!("{}", x.unwrap_or(0));

let y: Option<i32> = None;
println!("{}", y.unwrap_or(0));
```

## 不要滥用 unwrap

```rust
let x: Option<i32> = None;
let value = x.unwrap(); // panic
```

`unwrap()` 的意思是：我确信这里一定有值，否则程序崩溃。

新手可以在临时代码里用，但正式代码尽量用 `match`、`if let` 或 `?`。

## Result：处理错误

`Result` 定义大概是：

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

读取文件：

```rust
use std::fs;

fn main() {
    let result = fs::read_to_string("hello.txt");

    match result {
        Ok(content) => println!("文件内容: {}", content),
        Err(err) => println!("读取失败: {}", err),
    }
}
```

## ? 操作符

如果函数返回 `Result`，可以用 `?` 把错误往外传：

```rust
use std::fs;
use std::io;

fn read_file(path: &str) -> Result<String, io::Error> {
    let content = fs::read_to_string(path)?;
    Ok(content)
}

fn main() {
    match read_file("hello.txt") {
        Ok(content) => println!("{}", content),
        Err(err) => println!("错误: {}", err),
    }
}
```

`?` 的含义：

- 如果是 `Ok(value)`，取出 value 继续执行
- 如果是 `Err(e)`，立刻返回这个错误

## main 也可以返回 Result

```rust
use std::error::Error;
use std::fs;

fn main() -> Result<(), Box<dyn Error>> {
    let content = fs::read_to_string("hello.txt")?;
    println!("{}", content);
    Ok(())
}
```

这在小工具里很好用。

## panic 是什么

`panic!` 表示程序遇到不可恢复错误，直接崩溃：

```rust
panic!("出大问题了");
```

常见 panic：

- 数组越界
- `unwrap()` 一个 `None`
- `unwrap()` 一个 `Err`

## 什么时候用 Result，什么时候 panic

用 `Result`：

- 文件不存在
- 网络失败
- 用户输入错误
- 配置解析失败
- 数据库连接失败

这些都是可预期错误。

用 `panic`：

- 程序内部逻辑不可能出错却出错
- 测试里快速失败
- 原型阶段临时代码

## 自定义错误简化版

新手可以先用字符串错误：

```rust
fn divide(a: i32, b: i32) -> Result<i32, String> {
    if b == 0 {
        Err(String::from("除数不能为 0"))
    } else {
        Ok(a / b)
    }
}
```

使用：

```rust
match divide(10, 0) {
    Ok(v) => println!("结果: {}", v),
    Err(e) => println!("错误: {}", e),
}
```

## anyhow：应用程序常用错误库

写 CLI 或后端应用时，常用 `anyhow` 简化错误处理。

添加依赖：

```toml
[dependencies]
anyhow = "1"
```

示例：

```rust
use anyhow::Result;
use std::fs;

fn main() -> Result<()> {
    let content = fs::read_to_string("hello.txt")?;
    println!("{}", content);
    Ok(())
}
```

## thiserror：库常用错误库

如果你写的是给别人用的库，常用 `thiserror` 定义清晰错误类型。

这个后面再学，不是入门第一优先级。

## 本节练习

写一个除法函数：

```rust
fn safe_divide(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        Err(String::from("除数不能为 0"))
    } else {
        Ok(a / b)
    }
}

fn main() {
    match safe_divide(10.0, 2.0) {
        Ok(v) => println!("结果: {}", v),
        Err(e) => println!("错误: {}", e),
    }
}
```

## 本节小结

- `Option<T>` 表示可能没有值
- `Result<T,E>` 表示可能失败
- `match` 可以显式处理
- `?` 可以快速传播错误
- `unwrap()` 新手能用，但别滥用
