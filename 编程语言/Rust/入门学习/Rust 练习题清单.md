# Rust 练习题清单

这份练习按难度排列。建议每个题先自己写，再看资料。

## 基础语法

### 练习 1：奇偶判断

写函数：

```rust
fn is_even(n: i32) -> bool
```

输入整数，返回是否偶数。

### 练习 2：阶乘

写函数：

```rust
fn factorial(n: u32) -> u32
```

例如：

```text
factorial(5) = 120
```

### 练习 3：数组求和

给定：

```rust
let nums = vec![1, 2, 3, 4, 5];
```

求总和。

## 所有权与借用

### 练习 4：打印字符串长度

写函数：

```rust
fn print_len(s: &str)
```

要求调用后原字符串还能继续使用。

### 练习 5：修改字符串

写函数：

```rust
fn add_suffix(s: &mut String)
```

给字符串末尾追加 `"_done"`。

## 集合

### 练习 6：找最大值

写函数：

```rust
fn max_num(nums: &[i32]) -> Option<i32>
```

空数组返回 `None`。

### 练习 7：单词计数

统计字符串中每个单词出现次数。

输入：

```text
rust is fast rust is safe
```

输出类似：

```text
rust: 2
is: 2
fast: 1
safe: 1
```

## 结构体与枚举

### 练习 8：用户结构体

定义：

```rust
struct User {
    name: String,
    age: u32,
}
```

实现方法：

```rust
fn is_adult(&self) -> bool
```

### 练习 9：任务状态

定义枚举：

```rust
enum Status {
    Todo,
    Doing,
    Done,
}
```

写函数把状态转成中文字符串。

## 错误处理

### 练习 10：安全除法

写函数：

```rust
fn divide(a: f64, b: f64) -> Result<f64, String>
```

当 `b == 0.0` 返回错误。

### 练习 11：读取文件行数

写函数读取文件内容，返回行数：

```rust
fn count_lines(path: &str) -> Result<usize, std::io::Error>
```

## 小项目

### 练习 12：猜数字游戏

功能：

- 随机生成 1-100 数字
- 用户输入猜测
- 提示大了/小了/猜对了

用到：

- `rand`
- 标准输入
- `match`
- 循环

### 练习 13：JSON 通讯录

功能：

- 添加联系人
- 列出联系人
- 删除联系人
- 保存到 JSON 文件

用到：

- `serde`
- `serde_json`
- `Vec`
- 文件读写

### 练习 14：命令行 Todo

参考 [[10 实战：写一个命令行待办工具]]。

## 建议提交方式

每个练习一个小项目：

```bash
cargo new rust_ex_01
```

写完后至少跑：

```bash
cargo fmt
cargo check
cargo run
```

## 学习原则

1. 不要只复制代码
2. 每个例子至少改一个地方
3. 报错先读英文，不懂再查
4. 能编译只是第一步，能解释给别人听才算懂
