// 第 5 课：所有权与借用
//
// 所有权是 Rust 管理内存的核心规则：每个值有一个所有者，所有者离开作用域时，
// 值会被自动清理。借用允许函数使用数据而不取得所有权。

fn length(text: &str) -> usize {
    text.len()
}

fn add_exclamation(text: &mut String) {
    text.push('!');
}

fn first_word(text: &str) -> &str {
    text.split_whitespace().next().unwrap_or("")
}

fn main() {
    // String 存在堆上。赋值会移动所有权，移动后不能再使用原绑定。
    let original = String::from("所有权会移动");
    let moved = original;
    println!("移动后的字符串：{moved}");
    // println!("{original}"); // 编译错误：original 的值已经被移动。

    // 整数实现了 Copy，赋值会复制值，不会让第一个绑定失效。
    let first_number = 10;
    let second_number = first_number;
    println!("Copy 类型：{first_number} 和 {second_number}");

    // clone 会显式复制堆数据，代价比 Copy 更明显。
    let cloned = moved.clone();
    println!("clone 的副本：{cloned}");

    // 不可变借用可以同时存在多个，函数只读取数据。
    let text = String::from("借用不转移所有权");
    let text_length = length(&text);
    let word = first_word(&text);
    println!("text = {text}, 长度 = {text_length}, 首词 = {word}");

    // 可变借用允许修改数据，但同一时间不能与其他借用重叠。
    let mut editable = String::from("可变借用");
    add_exclamation(&mut editable);
    println!("修改后的 editable = {editable}");

    // 借用规则的要点：可以有多个不可变借用，或一个可变借用，不能同时两者都有。
    let slice = &editable[..];
    println!("字符串切片也是借用：{slice}");
}
