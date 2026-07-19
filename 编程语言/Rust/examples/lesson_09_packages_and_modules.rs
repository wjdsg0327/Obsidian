// 第 9 课：包、crate 和模块
//
// package 是 Cargo 管理的项目；一个 package 可以包含多个 binary crate 和一个
// library crate。每个 crate 是一个编译单元，module 用来组织 crate 内的代码。
// 本示例把多个模块写在一个文件中；实际项目通常会把模块拆到 src/*.rs 中。

mod front_of_house {
    pub mod hosting {
        pub fn add_to_waitlist() {
            println!("已加入等待列表");
        }
    }

    pub mod serving {
        pub fn take_order() {
            println!("服务员已接单");
        }

        pub fn complete_order() {
            // super 从当前模块回到父模块，再通过 crate 访问根模块。
            super::super::back_of_house::fix_incorrect_order();
        }
    }
}

mod back_of_house {
    pub(crate) fn fix_incorrect_order() {
        println!("厨房正在修正订单");
    }

    pub fn cook_order() {
        println!("厨房正在制作订单");
    }
}

mod restaurant {
    // use 创建一个较短的路径别名。
    use crate::front_of_house::hosting;

    pub fn eat_at_restaurant() {
        hosting::add_to_waitlist();
        crate::front_of_house::serving::take_order();
        crate::back_of_house::cook_order();
        crate::back_of_house::fix_incorrect_order();
    }
}

fn main() {
    restaurant::eat_at_restaurant();
    front_of_house::serving::complete_order();

    // `pub` 控制跨模块可见性；没有 pub 的项默认只在自己的模块及其子模块可见。
    println!("模块路径可以使用 crate、super 或从当前作用域开始书写");
}
