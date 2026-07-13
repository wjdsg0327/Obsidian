# Rust 新手常见报错速查

这篇专门用来查新手常见错误。Rust 报错虽然长，但通常很有价值。

## borrow of moved value

典型代码：

```rust
let s1 = String::from("hello");
let s2 = s1;
println!("{}", s1);
```

原因：

`s1` 的所有权已经移动给 `s2`。

解决：

```rust
let s2 = s1.clone();
```

或者用引用：

```rust
let s2 = &s1;
```

## cannot borrow as mutable

典型代码：

```rust
let s = String::from("hello");
s.push_str(" world");
```

原因：

变量默认不可变。

解决：

```rust
let mut s = String::from("hello");
s.push_str(" world");
```

## cannot borrow as mutable because it is also borrowed as immutable

典型代码：

```rust
let mut s = String::from("hello");
let r1 = &s;
let r2 = &mut s;
println!("{}", r1);
```

原因：

不可变引用和可变引用同时存在。

解决：

```rust
let mut s = String::from("hello");
let r1 = &s;
println!("{}", r1);

let r2 = &mut s;
r2.push_str(" world");
```

## mismatched types

典型代码：

```rust
let x: i32 = "hello";
```

原因：

类型不匹配。

解决：

看清楚函数参数、返回值、变量声明的类型。

## expected String, found &str

典型代码：

```rust
fn say(s: String) {}
say("hello");
```

原因：

函数要 `String`，你传了 `&str`。

解决一：转换：

```rust
say("hello".to_string());
```

解决二：函数参数改成更通用的 `&str`：

```rust
fn say(s: &str) {}
say("hello");
```

## cannot return reference to local variable

典型代码：

```rust
fn get_name() -> &String {
    let name = String::from("老王");
    &name
}
```

原因：

函数结束后 `name` 被释放，返回引用会悬垂。

解决：

```rust
fn get_name() -> String {
    String::from("老王")
}
```

## index out of bounds

典型代码：

```rust
let nums = vec![1, 2, 3];
println!("{}", nums[10]);
```

原因：

数组越界，程序 panic。

解决：

```rust
match nums.get(10) {
    Some(n) => println!("{}", n),
    None => println!("没有这个元素"),
}
```

## use of undeclared crate or module

典型情况：

```rust
use serde::Serialize;
```

但 `Cargo.toml` 没加依赖。

解决：

```toml
[dependencies]
serde = { version = "1", features = ["derive"] }
```

然后：

```bash
cargo check
```

## method not found

可能原因：

1. 类型不对
2. trait 没导入
3. 依赖 feature 没打开

例子：

```rust
use rand::Rng;
```

有些方法来自 trait，不导入 trait 就找不到。

## temporary value dropped while borrowed

典型原因：

你借用了一个临时值，但临时值很快被释放。

常见解决：

把临时值绑定到变量：

```rust
let text = format!("hello {}", "rust");
let s = text.as_str();
println!("{}", s);
```

## 读报错建议

Rust 报错建议按这个顺序看：

1. 第一行错误类型
2. 报错代码位置
3. note 解释
4. help 建议
5. 如果看不懂，把最小代码复制出来单独试

## 最实用命令

```bash
cargo check
cargo clippy
cargo fmt
```

- `cargo check`：快速看能不能编译
- `cargo clippy`：看更地道写法
- `cargo fmt`：自动格式化
