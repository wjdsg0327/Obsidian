// 第 7 课：枚举与模式匹配
//
// 枚举用一个类型表示多个可能的变体。match 必须覆盖所有可能性，
// 因此它适合把数据和分支逻辑放在一起表达。

#[derive(Debug)]
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

fn describe(message: &Message) -> String {
    match message {
        Message::Quit => String::from("退出消息"),
        Message::Move { x, y } => format!("移动到 ({x}, {y})"),
        Message::Write(text) => format!("文本消息：{text}"),
        Message::ChangeColor(red, green, blue) => {
            format!("颜色改为 RGB({red}, {green}, {blue})")
        }
    }
}

fn plus_one(value: Option<i32>) -> Option<i32> {
    match value {
        Some(number) => Some(number + 1),
        None => None,
    }
}

fn main() {
    let messages = [
        Message::Quit,
        Message::Move { x: 3, y: 4 },
        Message::Write(String::from("你好，枚举")),
        Message::ChangeColor(255, 128, 0),
    ];

    for message in &messages {
        println!("{message:?} -> {}", describe(message));
    }

    // Option<T> 表示“有一个 T”或“没有值”，比空指针更安全。
    let some_number = Some(41);
    let no_number: Option<i32> = None;
    println!("plus_one(Some(41)) = {:?}", plus_one(some_number));
    println!("plus_one(None) = {:?}", plus_one(no_number));

    let temperature = 28;
    // 模式可以绑定值，并且可以使用守卫增加条件。
    match temperature {
        value if value < 0 => println!("{value} 度：结冰"),
        value if value > 30 => println!("{value} 度：炎热"),
        value => println!("{value} 度：温和"),
    }

    // if let 适合只关心一个模式的情况。
    if let Some(number) = plus_one(Some(9)) {
        println!("if let 取到数字：{number}");
    }

    // while let 会在模式匹配成功时重复执行。
    let mut stack = vec![1, 2, 3];
    while let Some(top) = stack.pop() {
        println!("从栈顶取出 {top}");
    }
}
