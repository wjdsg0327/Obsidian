// 第 4 课：流程控制
//
// Rust 提供 if、loop、while 和 for。if 是表达式，loop 可以用 break 返回值，
// 标签可以让 break/continue 作用于指定的嵌套循环。

fn main() {
    let number = 6;
    let description = if number % 2 == 0 { "偶数" } else { "奇数" };
    println!("{number} 是{description}");

    // loop 会一直执行，直到 break；break 后面的值可以作为 loop 表达式的结果。
    let mut attempts = 0;
    let result = loop {
        attempts += 1;
        if attempts == 3 {
            break attempts * 10;
        }
    };
    println!("loop 在第 {attempts} 次尝试后返回 {result}");

    // 标签 'rows 让 break 跳出外层循环，而不只是内层循环。
    'rows: for row in 1..=3 {
        for column in 1..=3 {
            if row == 2 && column == 2 {
                println!("在 row={row}, column={column} 处跳出外层循环");
                break 'rows;
            }
            println!("访问单元格 ({row}, {column})");
        }
    }


    // while 适合“条件为真就继续”的场景。
    let mut countdown = 3;
    while countdown > 0 {
        println!("倒计时：{countdown}");
        countdown -= 1;
    }

    // for 直接遍历集合的引用，不取得集合所有权。
    let values = [1, 2, 3, 4, 5];
    let mut odd_sum = 0;
    for value in values {
        if value % 2 == 0 {
            continue; // 跳过偶数，继续下一次循环。
        }
        odd_sum += value;
    }
    println!("数组中的奇数之和 = {odd_sum}");

    for index in 0..values.len() {
        println!("values[{index}] = {}", values[index]);
    }
}
