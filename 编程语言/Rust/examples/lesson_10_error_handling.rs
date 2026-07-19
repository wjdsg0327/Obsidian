// 第 10 课：错误处理
//
// panic! 表示无法继续的不可恢复错误；Result<T, E> 表示调用者可以处理的错误。
// 生产代码通常让错误通过 Result 向上传播，由更高层决定如何展示或记录。

use std::num::ParseIntError;

fn parse_port(text: &str) -> Result<u16, ParseIntError> {
    text.parse::<u16>()
}

fn read_port(text: &str) -> Result<u16, String> {
    let port = parse_port(text).map_err(|error| format!("端口解析失败：{error}"))?;
    if port == 0 {
        return Err(String::from("端口不能是 0"));
    }
    Ok(port)
}

fn divide(dividend: f64, divisor: f64) -> Result<f64, String> {
    if divisor == 0.0 {
        Err(String::from("除数不能为 0"))
    } else {
        Ok(dividend / divisor)
    }
}

fn main() {
    // panic! 会立即终止当前线程。这里仅展示语法，不执行它，避免终止示例。
    // panic!("遇到了无法恢复的状态");

    let inputs = ["8080", "0", "not-a-port"];
    for input in inputs {
        match read_port(input) {
            Ok(port) => println!("端口 {input} 有效：{port}"),
            Err(error) => println!("端口 {input} 无效：{error}"),
        }
    }

    // unwrap 在你确定一定成功时简洁有用；expect 可以提供更明确的失败说明。
    let parsed: i32 = "42".parse().expect("示例中的数字应该可以解析");
    println!("expect 得到 parsed = {parsed}");

    match divide(10.0, 2.0) {
        Ok(value) => println!("10 / 2 = {value}"),
        Err(error) => println!("计算失败：{error}"),
    }

    match divide(10.0, 0.0) {
        Ok(value) => println!("结果 = {value}"),
        Err(error) => println!("计算失败：{error}"),
    }

    // Option 也可以用 ? 处理“有值/无值”的流程。
    let maybe_name = Some("Rust");
    if let Some(name) = maybe_name {
        println!("Option 中的名字：{name}");
    }
}
