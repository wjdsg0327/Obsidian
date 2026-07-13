# 05 String、Vec、HashMap 常用集合

写实际程序时，最常用的三个集合类型是：

- `String`：可增长字符串
- `Vec<T>`：动态数组
- `HashMap<K, V>`：键值对字典

## String 和 &str 的区别

### &str

字符串切片，通常是借用来的：

```rust
let s: &str = "hello";
```

它不能随便增长。

### String

拥有数据、可以增长：

```rust
let mut s = String::from("hello");
s.push_str(" world");
println!("{}", s);
```

## 创建 String

```rust
let s1 = String::new();
let s2 = String::from("hello");
let s3 = "hello".to_string();
```

## 修改 String

```rust
let mut s = String::from("hello");
s.push('!');
s.push_str(" rust");
println!("{}", s);
```

## 拼接 String

```rust
let s1 = String::from("hello");
let s2 = String::from("world");
let s3 = s1 + " " + &s2;

// println!("{}", s1); // 报错，s1 被移动
println!("{}", s2);
println!("{}", s3);
```

更推荐：

```rust
let name = "老王";
let age = 18;
let text = format!("{} 今年 {} 岁", name, age);
```

## 遍历字符串

Rust 字符串是 UTF-8，不能直接按下标取字符：

```rust
let s = "你好";
// s[0] // 不允许
```

按字符遍历：

```rust
for c in s.chars() {
    println!("{}", c);
}
```

按字节遍历：

```rust
for b in s.bytes() {
    println!("{}", b);
}
```

## Vec 动态数组

创建：

```rust
let mut nums: Vec<i32> = Vec::new();
nums.push(1);
nums.push(2);
nums.push(3);
```

更常用：

```rust
let nums = vec![1, 2, 3];
```

## 读取 Vec

```rust
let nums = vec![10, 20, 30];

let a = nums[0];
println!("{}", a);

match nums.get(10) {
    Some(value) => println!("值: {}", value),
    None => println!("没有这个位置"),
}
```

区别：

- `nums[10]` 越界会 panic
- `nums.get(10)` 返回 `Option`，更安全

## 遍历 Vec

只读：

```rust
let nums = vec![1, 2, 3];
for n in &nums {
    println!("{}", n);
}
println!("{:?}", nums); // 还能用
```

可变修改：

```rust
let mut nums = vec![1, 2, 3];
for n in &mut nums {
    *n *= 2;
}
println!("{:?}", nums);
```

注意 `*n` 是解引用。

## HashMap

使用前导入：

```rust
use std::collections::HashMap;
```

创建和插入：

```rust
use std::collections::HashMap;

fn main() {
    let mut scores = HashMap::new();
    scores.insert(String::from("老王"), 90);
    scores.insert(String::from("小张"), 80);

    println!("{:?}", scores);
}
```

## 读取 HashMap

```rust
let name = String::from("老王");
match scores.get(&name) {
    Some(score) => println!("分数: {}", score),
    None => println!("没找到"),
}
```

## 遍历 HashMap

```rust
for (name, score) in &scores {
    println!("{}: {}", name, score);
}
```

HashMap 默认无序，不保证输出顺序。

## 更新 HashMap

覆盖：

```rust
scores.insert(String::from("老王"), 95);
```

不存在才插入：

```rust
scores.entry(String::from("小李")).or_insert(70);
```

计数例子：

```rust
use std::collections::HashMap;

fn main() {
    let text = "hello rust hello world";
    let mut map = HashMap::new();

    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;
    }

    println!("{:?}", map);
}
```

## 本节练习

写一个程序统计一句话里每个单词出现次数：

```rust
use std::collections::HashMap;

fn main() {
    let text = "rust is fast and rust is safe";
    let mut counts = HashMap::new();

    for word in text.split_whitespace() {
        let count = counts.entry(word).or_insert(0);
        *count += 1;
    }

    println!("{:?}", counts);
}
```

## 本节小结

- `String` 用来拥有和修改字符串
- `&str` 更适合函数参数
- `Vec<T>` 是动态数组
- `HashMap<K,V>` 是字典
- 读取集合时，优先考虑 `get()`，避免越界 panic
