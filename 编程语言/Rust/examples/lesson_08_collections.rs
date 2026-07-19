// 第 8 课：常见集合
//
// Vec、String、HashMap 和 HashSet 都在堆上保存可变大小的数据。
// 集合通常拥有其中的元素，因此要遵守所有权和借用规则。

use std::collections::{HashMap, HashSet};

fn main() {
    // Vec<T> 是同类型元素的动态数组。
    let mut numbers = vec![10, 20, 30];
    numbers.push(40);
    numbers[0] = 11;
    println!("numbers = {numbers:?}, 第一个元素 = {}", numbers[0]);

    match numbers.get(10) {
        Some(value) => println!("索引 10 的值是 {value}"),
        None => println!("索引 10 越界，get 返回 None"),
    }

    let sum: i32 = numbers.iter().sum();
    let doubled: Vec<i32> = numbers.iter().map(|number| number * 2).collect();
    println!("sum = {sum}, doubled = {doubled:?}");

    // String 是拥有所有权的 UTF-8 字符串，不能用整数直接索引字符。
    let mut message = String::from("Rust");
    message.push(' ');
    message.push_str("学习");
    println!("message = {message}");
    for character in message.chars() {
        print!("[{character}]");
    }
    println!();

    let first = String::from("hello");
    let second = String::from("Rust");
    let joined = first + " " + &second; // + 会取得左侧 String 的所有权。
    println!("拼接结果：{joined}");
    println!("右侧字符串仍可使用：{second}");

    // HashMap 保存键值对；entry API 适合“存在则更新，不存在则插入”。
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);
    scores
        .entry(String::from("Blue"))
        .and_modify(|score| *score += 5);
    scores.entry(String::from("Red")).or_insert(25);
    println!("scores = {scores:?}");
    if let Some(score) = scores.get("Blue") {
        println!("Blue 的分数是 {score}");
    }

    // HashSet 只保存唯一值，可以进行交集、并集和差集运算。
    let a: HashSet<_> = [1, 2, 3, 4].into_iter().collect();
    let b: HashSet<_> = [3, 4, 5, 6].into_iter().collect();
    let intersection: Vec<_> = a.intersection(&b).copied().collect();
    let union: Vec<_> = a.union(&b).copied().collect();
    let difference: Vec<_> = a.difference(&b).copied().collect();
    println!("交集 = {intersection:?}");
    println!("并集 = {union:?}");
    println!("a 相对 b 的差集 = {difference:?}");
}
