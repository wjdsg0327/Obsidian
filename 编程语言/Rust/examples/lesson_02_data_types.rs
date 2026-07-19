// 第 2 课：数据类型
//
// Rust 是静态类型语言。编译器通常可以根据上下文推导类型，
// 也可以使用类型标注明确指定类型。

fn main() {
    // 标量类型：整数、浮点数、布尔值和字符。
    let unsigned: u32 = 42;
    let signed: i32 = -12;
    let decimal: f64 = 3.14;
    let is_ready: bool = true;
    let ferris: char = '🦀';

    println!(
        "标量：unsigned = {unsigned}, signed = {signed}, decimal = {decimal}, \
         is_ready = {is_ready}, ferris = {ferris}"
    );

    // 数值类型不会自动混合运算；需要显式转换。
    let converted = unsigned as i64 + signed as i64;
    println!("显式转换后的加法结果 = {converted}");

    // 元组可以保存不同类型的值，长度固定。
    let profile: (&str, u8, bool) = ("Alice", 30, true);
    let (name, age, active) = profile;
    println!("元组解构：name = {name}, age = {age}, active = {active}");

    // 数组的元素类型相同、长度固定；[value; length] 可以快速创建数组。
    let months = ["Jan", "Feb", "Mar", "Apr"];
    let repeated = [0; 5];
    println!(
        "数组第一个月 = {}, repeated 长度 = {}",
        months[0],
        repeated.len()
    );

    // 切片是对连续集合的一段借用，不拥有数据。
    let first_three = &months[0..3];
    println!("切片 first_three = {first_three:?}");

    // 字符串按 UTF-8 保存；char 表示一个 Unicode 标量值。
    let word = "你好 Rust";
    let characters = word.chars().count();
    let bytes = word.len();
    println!("{word} 有 {characters} 个字符、{bytes} 个 UTF-8 字节");
}
