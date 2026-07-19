// 第 13 课：生命周期
//
// 引用的生命周期描述“这份借用必须持续多久”。生命周期标注不会改变引用的存活时间，
// 而是把不同引用之间的关系告诉编译器，让编译器检查返回引用是否仍然有效。

// 返回值与两个输入引用中的较短者具有相同生命周期。
fn longest<'a>(left: &'a str, right: &'a str) -> &'a str {
    if left.len() >= right.len() {
        left
    } else {
        right
    }
}

// 只有一个输入引用时，Rust 可以根据生命周期省略规则自动推导下面的标注。
fn first_character(text: &str) -> Option<char> {
    text.chars().next()
}

#[derive(Debug)]
struct ImportantExcerpt<'a> {
    part: &'a str,
}

impl<'a> ImportantExcerpt<'a> {
    fn level(&self) -> i32 {
        3
    }

    fn announce(&self, announcement: &str) -> String {
        format!("{announcement}: {}", self.part)
    }
}

// &'static str 表示引用整个程序期间都有效的数据，例如字符串字面量。
fn default_title() -> &'static str {
    "Rust 学习"
}

fn main() {
    let first = String::from("较长的第一段文字");
    let second = String::from("第二段");
    let result = longest(&first, &second);
    println!("较长的文本：{result}");

    // result 只能在 first 和 second 都有效的范围内使用。
    println!("first_character(first) = {:?}", first_character(&first));

    let novel = String::from("很久很久以前。有一位 Rust 学习者。");
    let first_sentence = novel.split('.').next().unwrap_or("");
    let excerpt = ImportantExcerpt {
        part: first_sentence,
    };
    println!("excerpt = {excerpt:?}");
    println!("excerpt level = {}", excerpt.level());
    println!("{}", excerpt.announce("摘录"));

    println!("static 字符串：{}", default_title());

    // 下面这些写法会被借用检查器拒绝，因为返回引用可能指向已经释放的数据：
    // fn invalid() -> &str {
    //     let text = String::from("临时数据");
    //     &text
    // }
    // 生命周期规则确保这种悬垂引用无法进入可运行程序。
}
