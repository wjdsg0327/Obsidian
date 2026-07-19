// 第 12 课：Trait
//
// Trait 定义一组共享行为。类型实现 Trait 后，就可以在泛型约束或函数参数中
// 使用这种行为。Trait 不等于继承，更接近“满足某个能力接口”。

use std::fmt::Display;

trait Summary {
    fn summarize_author(&self) -> String;

    // 默认方法可以复用 Trait 中的其他方法。
    fn summarize(&self) -> String {
        format!("（作者：{}）", self.summarize_author())
    }
}

struct Article {
    headline: String,
    author: String,
}

struct Tweet {
    username: String,
    content: String,
}

impl Summary for Article {
    fn summarize_author(&self) -> String {
        self.author.clone()
    }

    fn summarize(&self) -> String {
        format!("{} - {}", self.headline, self.author)
    }
}

impl Summary for Tweet {
    fn summarize_author(&self) -> String {
        format!("@{}", self.username)
    }
}

fn notify(item: &impl Summary) {
    println!("通知：{}", item.summarize());
}

fn notify_with_bounds<T: Summary + Display>(item: &T) {
    println!("{} | {}", item, item.summarize());
}

fn returns_summary() -> impl Summary {
    Tweet {
        username: String::from("rustacean"),
        content: String::from("Trait 很实用"),
    }
}

impl Display for Tweet {
    fn fmt(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(formatter, "@{}: {}", self.username, self.content)
    }
}

fn main() {
    let article = Article {
        headline: String::from("Rust 学习笔记"),
        author: String::from("Alice"),
    };
    let tweet = Tweet {
        username: String::from("bob"),
        content: String::from("我正在学习 Trait"),
    };

    notify(&article);
    notify(&tweet);
    notify_with_bounds(&tweet);

    // impl Trait 允许隐藏具体返回类型，但返回值必须属于同一个具体类型。
    let summary = returns_summary();
    println!("返回的摘要：{}", summary.summarize());
}
