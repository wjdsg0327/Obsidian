// 第 3 课：函数
//
// 函数签名必须声明参数类型和返回类型。Rust 用表达式组成函数体，
// 最后一条没有分号的表达式会成为返回值。

fn greet(name: &str) {
    println!("你好，{name}！");
}

fn add(left: i32, right: i32) -> i32 {
    left + right
}

fn classify(number: i32) -> &'static str {
    if number > 0 {
        "正数"
    } else if number < 0 {
        "负数"
    } else {
        "零"
    }
}

fn double_if_positive(number: i32) -> i32 {
    if number <= 0 {
        return 0; // return 可以提前结束函数。
    }

    number * 2
}

fn main() {
    greet("Rustacean");

    let total = add(20, 22);
    println!("add(20, 22) = {total}");
    println!("-3 是{}", classify(-3));
    println!("double_if_positive(9) = {}", double_if_positive(9));

    // 语句执行动作但不产生值；表达式会产生值。
    let block_value = {
        let base = 5;
        base + 1 // 没有分号，所以这个值会从代码块返回。
    };
    println!("代码块表达式的值 = {block_value}");

    // 闭包是可以保存在变量中的匿名函数，参数和返回类型常可推导。
    let square = |number: i32| number * number;
    println!("闭包 square(6) = {}", square(6));
}
