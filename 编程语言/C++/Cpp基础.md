# C++ 基础语法速查

> 目标：先能读懂和写出常见 C++ 代码，再逐步进入 STL、现代 C++ 和框架项目。

## 1. 第一个 C++ 程序

```cpp
#include <iostream>

int main() {
    std::cout << "Hello World!" << std::endl;
    return 0;
}
```

编译运行：

```bash
g++ main.cpp -std=c++17 -O2 -Wall -o main
./main
```

说明：

- `#include <iostream>`：引入输入输出库
- `int main()`：程序入口
- `std::cout`：标准输出
- `return 0`：程序正常结束

小项目里可以写 `using namespace std;`，但正式项目更推荐写 `std::cout`、`std::vector`，避免名字冲突。

---

## 2. 变量与基础类型

```cpp
#include <iostream>
#include <string>

int main() {
    int age = 18;
    double score = 95.5;
    bool passed = true;
    char level = 'A';
    std::string name = "Lao Wang";

    std::cout << name << " " << age << " " << score << std::endl;
}
```

常见类型：

| 类型 | 用途 |
|---|---|
| `int` | 整数 |
| `long long` | 大整数 |
| `float` / `double` | 小数，常用 `double` |
| `bool` | true / false |
| `char` | 单个字符 |
| `std::string` | 字符串 |

---

## 3. 输入输出

```cpp
#include <iostream>
#include <string>

int main() {
    std::string name;
    int age;

    std::cout << "name: ";
    std::cin >> name;

    std::cout << "age: ";
    std::cin >> age;

    std::cout << name << " is " << age << " years old\n";
}
```

读取一整行：

```cpp
std::string line;
std::getline(std::cin, line);
```

---

## 4. 条件与循环

```cpp
if (score >= 90) {
    std::cout << "A\n";
} else if (score >= 60) {
    std::cout << "Pass\n";
} else {
    std::cout << "Fail\n";
}
```

```cpp
for (int i = 0; i < 5; ++i) {
    std::cout << i << "\n";
}

int n = 3;
while (n > 0) {
    std::cout << n << "\n";
    --n;
}
```

---

## 5. 函数

```cpp
#include <iostream>

int add(int a, int b) {
    return a + b;
}

void printHello() {
    std::cout << "hello\n";
}

int main() {
    std::cout << add(1, 2) << std::endl;
    printHello();
}
```

参数传递建议：

```cpp
void readOnly(const std::string& text);  // 只读大对象：const 引用
void modify(std::string& text);          // 需要修改：引用
void consume(std::string text);          // 需要拷贝/接管：值传递
```

---

## 6. 数组、字符串、vector

优先用 `std::vector` 和 `std::string`，少用裸数组。

```cpp
#include <iostream>
#include <vector>
#include <string>

int main() {
    std::vector<int> nums = {1, 2, 3};
    nums.push_back(4);

    for (int x : nums) {
        std::cout << x << "\n";
    }

    std::string s = "cpp";
    std::cout << s.size() << "\n";
}
```

常用操作：

```cpp
nums.size();      // 长度
nums.empty();     // 是否为空
nums.push_back(5);
nums.pop_back();
nums[0];          // 不检查越界
nums.at(0);       // 检查越界
```

---

## 7. 指针与引用

引用像“别名”，指针像“地址变量”。

```cpp
int x = 10;
int& ref = x;
ref = 20;         // x 也变成 20

int* ptr = &x;
*ptr = 30;        // x 变成 30
```

优先级建议：

1. 普通对象
2. 引用：`T&` / `const T&`
3. 智能指针：`std::unique_ptr` / `std::shared_ptr`
4. 裸指针：只在不得不用时使用

---

## 8. struct 与 class

`struct` 默认成员是 public，适合简单数据；`class` 默认成员是 private，适合封装行为。

```cpp
#include <iostream>
#include <string>

class User {
private:
    std::string name;
    int age;

public:
    User(std::string name, int age) : name(name), age(age) {}

    void sayHello() const {
        std::cout << "Hi, I am " << name << "\n";
    }
};

int main() {
    User user("Lao Wang", 18);
    user.sayHello();
}
```

重点：

- 构造函数负责初始化对象
- 析构函数负责释放资源
- `const` 成员函数表示不会修改对象
- 初始化列表 `: name(name), age(age)` 是 C++ 常见写法

---

## 9. RAII 与智能指针

RAII：资源获取即初始化。对象创建时拿资源，对象销毁时自动释放资源。

```cpp
#include <memory>

struct Image {
    int width;
    int height;
};

int main() {
    auto img = std::make_unique<Image>(Image{1920, 1080});
}
```

常用智能指针：

| 类型 | 含义 |
|---|---|
| `std::unique_ptr<T>` | 独占所有权，推荐默认使用 |
| `std::shared_ptr<T>` | 共享所有权，有引用计数 |
| `std::weak_ptr<T>` | 弱引用，常用于打破循环引用 |

---

## 10. STL 容器

```cpp
#include <vector>
#include <map>
#include <unordered_map>
#include <set>

std::vector<int> list = {1, 2, 3};
std::map<std::string, int> ordered;
std::unordered_map<std::string, int> dict;
std::set<int> uniqueNumbers;
```

怎么选：

| 场景 | 容器 |
|---|---|
| 顺序列表 | `std::vector` |
| 队列 | `std::queue` |
| 栈 | `std::stack` |
| 键值映射，要求有序 | `std::map` |
| 键值映射，追求查询速度 | `std::unordered_map` |
| 去重集合 | `std::set` / `std::unordered_set` |

---

## 11. STL 算法与 Lambda

```cpp
#include <algorithm>
#include <iostream>
#include <vector>

int main() {
    std::vector<int> nums = {3, 1, 4, 2};

    std::sort(nums.begin(), nums.end());

    auto it = std::find(nums.begin(), nums.end(), 4);
    if (it != nums.end()) {
        std::cout << "found\n";
    }

    std::sort(nums.begin(), nums.end(), [](int a, int b) {
        return a > b;
    });
}
```

现代 C++ 写法要点：

- 能用标准算法就少手写循环
- Lambda 适合临时排序规则、过滤规则、回调函数
- `auto` 适合类型很长但含义清楚的地方

---

## 12. 文件读写

```cpp
#include <fstream>
#include <iostream>
#include <string>

int main() {
    std::ofstream out("hello.txt");
    out << "hello cpp\n";

    std::ifstream in("hello.txt");
    std::string line;
    while (std::getline(in, line)) {
        std::cout << line << "\n";
    }
}
```

---

## 13. 最小 CMake 项目

目录：

```text
hello-cpp/
├── CMakeLists.txt
└── main.cpp
```

`CMakeLists.txt`：

```cmake
cmake_minimum_required(VERSION 3.16)
project(hello_cpp)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_executable(hello main.cpp)
```

构建：

```bash
cmake -S . -B build
cmake --build build
```

---

## 14. 必须掌握的 C++ 关键词

| 关键词 | 重点 |
|---|---|
| `const` | 不可修改，常用于只读参数和成员函数 |
| `static` | 静态变量、静态成员、内部链接 |
| `virtual` | 多态 |
| `override` | 明确重写父类虚函数 |
| `template` | 泛型编程 |
| `auto` | 自动类型推导 |
| `nullptr` | 空指针，替代 `NULL` |
| `using` | 类型别名、命名空间引入 |

---

## 15. 小练习

1. 写一个命令行计算器，支持加减乘除
2. 用 `vector` 保存学生成绩，计算最高分、最低分、平均分
3. 写一个 `User` 类，包含姓名、年龄、打印信息方法
4. 用 `unordered_map` 统计一段文本里每个单词出现次数
5. 用文件读写保存和读取学生列表

下一步看：[C++快速学习路线](C++快速学习路线.md)。
