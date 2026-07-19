// 第 1 课：变量
//
// Rust 中的变量默认不可变。使用 `mut` 才能重新赋值；使用变量遮蔽（shadowing）
// 可以用同一个名字绑定一个新值，甚至改变它的类型。

const MAX_POINTS: u32 = 100;

fn main() {
    // 不可变变量：绑定创建后，不能再次赋值。
    let language = "Rust";
    println!("不可变变量 language = {language}");

    // 可变变量：必须在 let 后显式写出 mut。
    let mut score: u32 = 80;
    score += 10;
    println!("可变变量 score = {score}");

    // 类型标注可以帮助读者和编译器理解数据的类型。
    let temperature: i32 = -7;
    let ratio = 0.75_f64; // 后缀也可以直接指定数值类型。
    println!("temperature = {temperature}, ratio = {ratio}");

    // 常量没有固定的运行时绑定，不属于某个函数作用域。
    println!("MAX_POINTS = {MAX_POINTS}");

    // 遮蔽会创建一个新的绑定；下面的 spaces 从字符串变成了 usize。
    let spaces = "   ";
    let spaces = spaces.len();
    println!("遮蔽后的 spaces 长度 = {spaces}");

    // 元组可以一次解构成多个变量。
    let (x, y, z) = (10, 20, 30);
    println!("解构结果：x = {x}, y = {y}, z = {z}");

    // Rust 的变量绑定默认要求初始化后才能使用。
    let initialized_later;
    initialized_later = "现在已经初始化";
    println!("{initialized_later}");
}
