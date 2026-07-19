// 第 11 课：泛型
//
// 泛型让一段代码适用于多种类型。编译器会在编译期为实际使用的类型生成专门版本，
// 这个过程叫单态化，因此泛型通常不会带来运行时类型判断成本。

#[derive(Debug)]
struct Point<T> {
    x: T,
    y: T,
}

#[derive(Debug)]
struct MixedPoint<T, U> {
    x: T,
    y: U,
}

enum Maybe<T> {
    Value(T),
    Empty,
}

fn largest<T: PartialOrd + Copy>(items: &[T]) -> T {
    let mut largest = items[0];
    for &item in &items[1..] {
        if item > largest {
            largest = item;
        }
    }
    largest
}

impl<T> Point<T> {
    fn x(&self) -> &T {
        &self.x
    }
}

impl Point<f64> {
    fn distance_from_origin(&self) -> f64 {
        (self.x * self.x + self.y * self.y).sqrt()
    }
}

impl<T, U> MixedPoint<T, U> {
    fn mix<V, W>(self, other: MixedPoint<V, W>) -> MixedPoint<T, W> {
        MixedPoint {
            x: self.x,
            y: other.y,
        }
    }
}

fn main() {
    let integer_point = Point { x: 5, y: 10 };
    let float_point = Point { x: 1.5, y: 2.5 };
    let mixed_point = MixedPoint { x: 5, y: 4.0 };
    println!(
        "integer_point = {integer_point:?}, x = {}",
        integer_point.x()
    );
    println!(
        "float_point = {float_point:?}, 到原点距离 = {}",
        float_point.distance_from_origin()
    );
    println!("mixed_point = {mixed_point:?}");

    let mixed_result = mixed_point.mix(MixedPoint { x: 'x', y: "new y" });
    println!("mix 后得到 mixed_result = {mixed_result:?}");

    let numbers = [3, 9, 2, 7];
    let words = ["apple", "pear", "banana"];
    println!("最大数字 = {}", largest(&numbers));
    println!("最大单词 = {}", largest(&words));

    let present = Maybe::Value(String::from("泛型枚举"));
    let absent: Maybe<i32> = Maybe::Empty;
    match present {
        Maybe::Value(value) => println!("Maybe 中有值：{value}"),
        Maybe::Empty => println!("Maybe 为空"),
    }
    match absent {
        Maybe::Value(value) => println!("Maybe 中有数字：{value}"),
        Maybe::Empty => println!("Maybe 为空"),
    }
}
