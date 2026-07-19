// 第 6 课：结构体
//
// 结构体把相关数据组合成一个有意义的类型。impl 块可以为结构体添加方法
// 和关联函数；方法的第一个参数通常是 self、&self 或 &mut self。

#[derive(Debug, Clone)]
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

fn build_user(username: String, email: String) -> User {
    User {
        username,
        email,
        sign_in_count: 1,
        active: true,
    }
}

#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }

    fn square(size: u32) -> Rectangle {
        Rectangle {
            width: size,
            height: size,
        }
    }

    fn resize(&mut self, width: u32, height: u32) {
        self.width = width;
        self.height = height;
    }
}

#[derive(Debug)]
struct Color(u8, u8, u8); // 元组结构体没有字段名。

#[derive(Debug)]
struct Marker; // 类单元结构体没有字段，适合表达一个标记类型。

fn main() {
    let user1 = build_user(String::from("alice"), String::from("alice@example.com"));
    println!("user1 = {user1:?}");

    // 结构体更新语法会复用 user1 中未显式指定的字段。
    let user2 = User {
        email: String::from("alice@rust.example"),
        ..user1.clone()
    };
    println!("user2 = {user2:?}");
    println!(
        "user2 的用户名：{}，是否启用：{}",
        user2.username, user2.active
    );
    println!("登录次数：{}，邮箱：{}", user2.sign_in_count, user2.email);

    let mut rectangle = Rectangle {
        width: 30,
        height: 50,
    };
    println!("rectangle = {rectangle:?}, 面积 = {}", rectangle.area());
    rectangle.resize(40, 60);
    let square = Rectangle::square(20);
    println!("调整后面积 = {}，square = {square:?}", rectangle.area());
    println!(
        "rectangle 能容纳 square 吗？{}",
        rectangle.can_hold(&square)
    );

    let color = Color(255, 128, 0);
    let marker = Marker;
    println!(
        "元组结构体 color = {color:?}，RGB = ({}, {}, {})，单元结构体 marker = {marker:?}",
        color.0, color.1, color.2
    );
}
