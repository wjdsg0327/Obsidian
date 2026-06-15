# C++ 指针专题教程

> 目标：把 C++ 指针从“看着吓人”学到“能安全使用”。重点不是背符号，而是理解地址、所有权、生命周期。

## 1. 指针是什么

变量有两个东西：

- 值：变量里存的数据
- 地址：变量在内存里的位置

指针就是“保存地址的变量”。

```cpp
#include <iostream>

int main() {
    int x = 10;
    int* p = &x;

    std::cout << "x 的值: " << x << "\n";
    std::cout << "x 的地址: " << &x << "\n";
    std::cout << "p 保存的地址: " << p << "\n";
    std::cout << "p 指向的值: " << *p << "\n";
}
```

记住三个符号：

| 符号 | 含义 |
|---|---|
| `int* p` | 定义一个指向 `int` 的指针 |
| `&x` | 取变量 `x` 的地址 |
| `*p` | 解引用，访问 `p` 指向的值 |

一句话：`p` 是地址，`*p` 是地址里的值。

---

## 2. 修改指针指向的值

```cpp
#include <iostream>

int main() {
    int x = 10;
    int* p = &x;

    *p = 99;

    std::cout << x << "\n";  // 99
}
```

`*p = 99` 的意思是：把 `p` 指向的那块内存改成 `99`。

所以指针强大的地方是：它可以通过地址间接修改变量。

---

## 3. 指针变量本身也可以改

```cpp
#include <iostream>

int main() {
    int a = 10;
    int b = 20;

    int* p = &a;
    std::cout << *p << "\n";  // 10

    p = &b;
    std::cout << *p << "\n";  // 20
}
```

这里变化的是 `p` 保存的地址。

- `p = &a`：指向 `a`
- `p = &b`：改为指向 `b`

---

## 4. 空指针 nullptr

空指针表示“现在不指向任何有效对象”。

```cpp
int* p = nullptr;
```

使用前要判断：

```cpp
if (p != nullptr) {
    std::cout << *p << "\n";
}
```

不要这样：

```cpp
int* p = nullptr;
std::cout << *p << "\n";  // 错误：解引用空指针
```

解引用空指针是未定义行为，程序可能崩溃，也可能表现得很诡异。

---

## 5. 指针和引用的区别

引用是变量的别名，指针是保存地址的变量。

```cpp
int x = 10;

int& r = x;   // 引用
int* p = &x;  // 指针
```

对比：

| 对比项 | 引用 `T&` | 指针 `T*` |
|---|---|---|
| 是否可以为空 | 不应该为空 | 可以是 `nullptr` |
| 是否可以换目标 | 不可以 | 可以 |
| 使用方式 | `r` | `*p` |
| 常见用途 | 函数参数、别名 | 可选对象、数组、底层接口 |

优先级建议：

1. 能用普通对象就用普通对象
2. 传参优先用引用：`T&` / `const T&`
3. 需要表示“可能为空”时用指针
4. 需要表达所有权时用智能指针

---

## 6. 指针作为函数参数

### 6.1 用指针修改外部变量

```cpp
#include <iostream>

void addOne(int* p) {
    if (p != nullptr) {
        *p += 1;
    }
}

int main() {
    int x = 10;
    addOne(&x);
    std::cout << x << "\n";  // 11
}
```

### 6.2 更常见的写法：引用

```cpp
void addOne(int& x) {
    x += 1;
}
```

如果参数必须存在，用引用更清楚；如果参数可以不传，用指针更合适。

```cpp
void update(User& user);       // user 必须存在
void update(User* userOrNull); // user 可以为空
```

---

## 7. const 和指针

这块容易绕，分开看。

### 7.1 指向常量的指针

不能通过指针修改值。

```cpp
int x = 10;
const int* p = &x;

// *p = 20;  // 错误
p = nullptr; // 可以，指针本身能改
```

读法：`p` 指向的 `int` 是 const。

### 7.2 常量指针

指针本身不能换目标，但可以改指向的值。

```cpp
int x = 10;
int y = 20;

int* const p = &x;
*p = 99;   // 可以
// p = &y; // 错误
```

读法：`p` 这个指针是 const。

### 7.3 指向常量的常量指针

指针不能换目标，也不能通过它改值。

```cpp
int x = 10;
const int* const p = &x;
```

快速判断法：看 `const` 修饰谁。

```cpp
const int* p;        // 不能改 *p，可以改 p
int* const p = &x;   // 可以改 *p，不能改 p
const int* const p;  // 两个都不能改
```

---

## 8. 指针和数组

数组名在很多场景下会退化成指向首元素的指针。

```cpp
#include <iostream>

int main() {
    int arr[3] = {10, 20, 30};
    int* p = arr;

    std::cout << *p << "\n";       // 10
    std::cout << *(p + 1) << "\n"; // 20
    std::cout << *(p + 2) << "\n"; // 30
}
```

`p + 1` 不是地址加 1 个字节，而是移动到下一个 `int`。

更推荐现代写法：

```cpp
#include <array>
#include <vector>

std::array<int, 3> fixed = {10, 20, 30};
std::vector<int> nums = {10, 20, 30};
```

日常开发里，优先用 `std::vector`、`std::array`，少直接操作裸数组指针。

---

## 9. 指针和字符串

C 风格字符串本质是字符数组，以 `'\0'` 结尾。

```cpp
const char* text = "hello";
```

更推荐：

```cpp
#include <string>

std::string text = "hello";
```

规则很简单：

- 业务代码优先用 `std::string`
- 调用 C 接口时可能需要 `const char*`
- 从 `std::string` 获取 C 字符串：`text.c_str()`

```cpp
std::string name = "cpp";
const char* raw = name.c_str();
```

注意：`raw` 指向的是 `name` 内部数据。`name` 被销毁或大幅修改后，`raw` 可能失效。

---

## 10. 动态内存 new/delete

传统写法：

```cpp
int* p = new int(10);
std::cout << *p << "\n";
delete p;
p = nullptr;
```

数组：

```cpp
int* arr = new int[3]{1, 2, 3};
delete[] arr;
arr = nullptr;
```

两个危险点：

- `new` 后忘记 `delete`：内存泄漏
- `delete` 后继续使用：悬空指针

现代 C++ 里，业务代码尽量不要手写 `new/delete`，优先用智能指针和标准容器。

---

## 11. 悬空指针

悬空指针：指针还保存着地址，但那块内存已经无效。

```cpp
int* badPointer() {
    int x = 10;
    return &x;  // 错误：返回局部变量地址
}
```

`x` 是局部变量，函数结束后就被销毁了，返回它的地址没有意义。

正确写法：

```cpp
int getValue() {
    int x = 10;
    return x;
}
```

或者返回对象：

```cpp
#include <string>

std::string makeName() {
    return "cpp";
}
```

现代 C++ 返回对象通常很高效，不要为了“省拷贝”乱返回指针。

---

## 12. 指针的所有权

这是学 C++ 指针最重要的地方。

指针本身只是地址，它不一定表示“我负责释放这块内存”。

常见语义：

```cpp
void read(const User* user);       // 借用，只读，可以为空
void update(User* user);           // 借用，可修改，可以为空
void update(User& user);           // 借用，可修改，必须存在
std::unique_ptr<User> createUser(); // 创建并转移所有权
```

判断一个指针危险不危险，要问两件事：

1. 它指向的对象现在还活着吗？
2. 谁负责释放它？

如果这两个问题说不清，代码就容易出事。

---

## 13. 智能指针

### 13.1 unique_ptr

独占所有权。默认优先使用。

```cpp
#include <memory>
#include <iostream>

struct User {
    std::string name;
};

int main() {
    auto user = std::make_unique<User>(User{"Lao Wang"});
    std::cout << user->name << "\n";
}
```

`user` 离开作用域后，对象自动释放。

转移所有权：

```cpp
std::unique_ptr<User> a = std::make_unique<User>();
std::unique_ptr<User> b = std::move(a);

// 此时 a 为空，b 拥有对象
```

### 13.2 shared_ptr

共享所有权。最后一个 `shared_ptr` 销毁时，对象释放。

```cpp
#include <memory>

auto a = std::make_shared<User>();
auto b = a;
```

不要滥用 `shared_ptr`。它让所有权变模糊，代码看起来方便，维护时容易绕。

### 13.3 weak_ptr

弱引用，不增加引用计数，常用于避免循环引用。

```cpp
std::weak_ptr<User> weak = a;

if (auto locked = weak.lock()) {
    // 对象还活着
}
```

---

## 14. 指针访问成员：->

如果有对象：

```cpp
User user;
user.name = "cpp";
```

如果有指针：

```cpp
User* p = &user;
p->name = "cpp";
```

`p->name` 等价于 `(*p).name`。

智能指针也支持 `->`：

```cpp
auto user = std::make_unique<User>();
user->name = "cpp";
```

---

## 15. 二级指针

二级指针就是“指向指针的指针”。

```cpp
int x = 10;
int* p = &x;
int** pp = &p;

std::cout << **pp << "\n";  // 10
```

初学阶段只需要知道它存在。常见于：

- C 语言接口
- 需要在函数里修改一个指针变量本身
- 一些底层数据结构

C++ 业务代码里，二级指针出现太多通常说明设计可以换成引用、容器或智能指针。

---

## 16. 常见错误清单

### 16.1 未初始化指针

```cpp
int* p;
// *p = 10; // 错误
```

正确：

```cpp
int* p = nullptr;
```

### 16.2 空指针解引用

```cpp
int* p = nullptr;
std::cout << *p; // 错误
```

### 16.3 数组越界

```cpp
int arr[3] = {1, 2, 3};
std::cout << arr[3]; // 错误
```

### 16.4 delete 后继续使用

```cpp
int* p = new int(10);
delete p;
std::cout << *p; // 错误
```

### 16.5 new 和 delete 不匹配

```cpp
int* arr = new int[10];
delete arr;   // 错误，应该 delete[]
```

正确：

```cpp
delete[] arr;
```

---

## 17. 指针使用原则

建议按这个顺序选择：

| 场景 | 推荐写法 |
|---|---|
| 单个普通对象 | 直接定义对象 |
| 函数只读大对象 | `const T&` |
| 函数修改对象且对象必须存在 | `T&` |
| 参数可以为空 | `T*` |
| 独占所有权 | `std::unique_ptr<T>` |
| 共享所有权 | `std::shared_ptr<T>` |
| 动态数组 | `std::vector<T>` |
| 固定长度数组 | `std::array<T, N>` |
| 字符串 | `std::string` |

一句话：**裸指针适合表达“借用”和“可为空”，智能指针适合表达“所有权”。**

---

## 18. 练习

### 练习 1：基础地址

定义一个 `int x = 10`，打印：

- `x`
- `&x`
- 指针 `p`
- `*p`

### 练习 2：交换变量

分别用指针和引用写两个交换函数：

```cpp
void swapByPointer(int* a, int* b);
void swapByReference(int& a, int& b);
```

### 练习 3：数组求和

用指针遍历数组求和：

```cpp
int sum(const int* arr, int size);
```

### 练习 4：可选参数

写一个函数：

```cpp
void printUser(const User* user);
```

当 `user == nullptr` 时打印 `"no user"`。

### 练习 5：智能指针

用 `std::unique_ptr` 创建一个 `User` 对象，并打印它的字段。

---

## 19. 学习路线建议

学完这篇后，按顺序补：

1. 引用和 `const` 参数
2. 类的构造函数和析构函数
3. RAII
4. `std::unique_ptr`
5. `std::vector` 和 `std::string`

不要一开始死磕二级指针、函数指针、复杂内存模型。先把“对象是否还活着”和“谁负责释放”这两个问题养成习惯，C++ 指针就顺很多了。
