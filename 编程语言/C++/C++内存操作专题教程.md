# C++ 内存操作专题教程

> 目标：理解 C++ 程序里的内存从哪里来、什么时候释放、怎么避免泄漏和越界。重点是“生命周期”和“所有权”，不是鼓励到处手动操作内存。

## 1. 为什么 C++ 要学内存

C++ 和 Python、Java、C# 最大的不同之一是：C++ 允许你非常接近内存。

这带来两个结果：

- 好处：性能高、控制力强，适合游戏、图形、嵌入式、底层系统、AI 推理框架
- 风险：越界、泄漏、悬空引用、重复释放都可能让程序崩溃

学内存不是为了到处写 `new/delete`，而是为了知道：

1. 对象什么时候创建
2. 对象什么时候销毁
3. 谁拥有这块资源
4. 这块内存还能不能访问

---

## 2. C++ 常见内存区域

可以先粗略理解为这几块：

| 区域 | 放什么 | 生命周期 |
|---|---|---|
| 栈 stack | 局部变量、函数参数 | 作用域结束自动释放 |
| 堆 heap | 动态创建的对象 | 手动释放或智能指针释放 |
| 全局/静态区 | 全局变量、`static` 变量 | 程序结束释放 |
| 常量区 | 字符串字面量等 | 程序运行期间存在 |

示例：

```cpp
#include <iostream>
#include <string>

int globalValue = 100;          // 全局/静态区

int main() {
    int localValue = 10;        // 栈
    static int staticValue = 5; // 全局/静态区
    std::string text = "cpp";   // text 对象在栈上，内部字符可能在堆上

    int* heapValue = new int(20); // 堆
    delete heapValue;
}
```

---

## 3. 栈内存

栈内存由系统自动管理，函数调用时创建，函数返回时释放。

```cpp
#include <iostream>

void run() {
    int x = 10;
    std::cout << x << "\n";
} // x 在这里自动销毁

int main() {
    run();
}
```

优点：

- 快
- 自动释放
- 不容易泄漏

限制：

- 空间有限
- 不能返回局部变量的地址或引用

错误示例：

```cpp
int* makeBadPointer() {
    int x = 10;
    return &x; // 错误：x 离开函数就销毁了
}
```

正确做法：

```cpp
int makeValue() {
    return 10;
}
```

---

## 4. 堆内存

堆内存适合需要动态生命周期的对象。

传统写法：

```cpp
int* p = new int(10);
delete p;
p = nullptr;
```

数组：

```cpp
int* arr = new int[3]{1, 2, 3};
delete[] arr;
arr = nullptr;
```

问题是：你必须保证每一次 `new` 都有正确的 `delete`。

如果忘记：

```cpp
void leak() {
    int* p = new int(10);
} // 内存泄漏
```

所以现代 C++ 里，业务代码优先用：

```cpp
#include <memory>
#include <vector>

auto value = std::make_unique<int>(10);
std::vector<int> nums = {1, 2, 3};
```

---

## 5. 对象生命周期

对象生命周期就是：对象从构造开始，到析构结束。

```cpp
#include <iostream>

class User {
public:
    User() {
        std::cout << "构造 User\n";
    }

    ~User() {
        std::cout << "析构 User\n";
    }
};

int main() {
    User user;
    std::cout << "main running\n";
}
```

输出顺序大致是：

```text
构造 User
main running
析构 User
```

理解生命周期后，你就能判断：

- 这个对象现在是否还存在
- 指针/引用是否还有效
- 析构函数什么时候会自动执行

---

## 6. RAII：C++ 管资源的核心方式

RAII：Resource Acquisition Is Initialization，资源获取即初始化。

简单说：把资源交给对象管理，对象活着资源就活着，对象销毁资源就释放。

```cpp
#include <fstream>

void writeFile() {
    std::ofstream file("log.txt");
    file << "hello\n";
} // file 析构，文件自动关闭
```

RAII 可以管理：

- 内存
- 文件
- 锁
- 网络连接
- 图像资源
- GPU 资源

自己写一个简化版 RAII：

```cpp
#include <iostream>

class Buffer {
private:
    int* data;

public:
    Buffer(int size) : data(new int[size]) {
        std::cout << "分配内存\n";
    }

    ~Buffer() {
        delete[] data;
        std::cout << "释放内存\n";
    }
};

int main() {
    Buffer buffer(10);
} // 自动调用析构函数
```

这个例子用于理解。实际项目里更推荐 `std::vector<int>`。

---

## 7. 用 vector 替代动态数组

不推荐：

```cpp
int* arr = new int[100];
delete[] arr;
```

推荐：

```cpp
#include <vector>

std::vector<int> arr(100);
arr[0] = 10;
```

`vector` 的好处：

- 自动管理内存
- 自动扩容
- 知道自己的长度
- 可以用 `.at()` 做越界检查

```cpp
std::vector<int> nums = {1, 2, 3};
nums.push_back(4);

std::cout << nums.size() << "\n";
std::cout << nums.at(0) << "\n";
```

初学阶段记住：**需要动态数组时，优先 `std::vector`。**

---

## 8. 用 string 替代字符数组

不推荐：

```cpp
char name[20] = "cpp";
```

推荐：

```cpp
#include <string>

std::string name = "cpp";
```

需要传给 C 接口时：

```cpp
const char* raw = name.c_str();
```

注意：`raw` 只是借用了 `name` 内部内存，不要在 `name` 销毁后继续用。

---

## 9. 智能指针管理单个对象

### 9.1 unique_ptr

独占所有权，默认优先使用。

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

离开作用域后，`User` 自动释放。

转移所有权：

```cpp
auto a = std::make_unique<User>();
auto b = std::move(a);

// a 变空，b 拥有对象
```

### 9.2 shared_ptr

共享所有权。

```cpp
#include <memory>

auto a = std::make_shared<User>();
auto b = a;
```

最后一个 `shared_ptr` 销毁时，对象才释放。

注意：`shared_ptr` 很方便，但容易让“谁负责释放”变模糊，不要一上来全用它。

---

## 10. 拷贝和移动

内存操作绕不开拷贝和移动。

### 10.1 拷贝

拷贝会产生一份新数据。

```cpp
#include <string>

std::string a = "hello";
std::string b = a; // 拷贝
```

### 10.2 移动

移动表示资源所有权转移。

```cpp
#include <memory>

auto a = std::make_unique<int>(10);
auto b = std::move(a);
```

移动后：

- `b` 拥有那块内存
- `a` 不再拥有对象

不要继续假设 `a` 还能正常访问对象。

---

## 11. 浅拷贝和深拷贝

手写裸指针类时最容易出问题。

错误示例：

```cpp
class BadBuffer {
public:
    int* data;

    BadBuffer(int size) {
        data = new int[size];
    }

    ~BadBuffer() {
        delete[] data;
    }
};

int main() {
    BadBuffer a(10);
    BadBuffer b = a; // 默认浅拷贝，两个对象指向同一块内存
} // 析构时重复 delete
```

解决思路：

- 最好直接用 `std::vector<int>`
- 或者禁用拷贝
- 或者自己实现深拷贝

禁用拷贝：

```cpp
class Buffer {
public:
    Buffer(const Buffer&) = delete;
    Buffer& operator=(const Buffer&) = delete;
};
```

实际项目里，优先让标准库替你管理资源。

---

## 12. 内存泄漏

内存泄漏：申请的内存没有释放。

```cpp
void leak() {
    int* p = new int(10);
} // p 丢了，但堆上的 int 没释放
```

修复：

```cpp
void noLeak() {
    auto p = std::make_unique<int>(10);
}
```

常见泄漏来源：

- `new` 后忘记 `delete`
- 异常提前返回，跳过释放逻辑
- 容器里保存裸指针但没有清理
- `shared_ptr` 循环引用

---

## 13. 重复释放

同一块内存释放两次非常危险。

```cpp
int* p = new int(10);
delete p;
delete p; // 错误
```

如果不得不用裸指针，释放后置空：

```cpp
delete p;
p = nullptr;
```

但更推荐直接避免裸 `new/delete`。

---

## 14. 越界访问

```cpp
int arr[3] = {1, 2, 3};
arr[3] = 99; // 错误
```

`arr[3]` 已经越界，因为合法下标是 `0, 1, 2`。

`vector` 也可能越界：

```cpp
std::vector<int> nums = {1, 2, 3};
nums[3] = 99; // 错误，operator[] 不检查越界
```

需要检查时用 `.at()`：

```cpp
nums.at(3) = 99; // 抛出异常，便于发现问题
```

---

## 15. use-after-free

释放后继续使用，就是 use-after-free。

```cpp
int* p = new int(10);
delete p;
std::cout << *p << "\n"; // 错误
```

更隐蔽的情况：

```cpp
std::vector<int> nums = {1, 2, 3};
int* p = &nums[0];

nums.push_back(4); // 可能触发扩容，旧地址失效
std::cout << *p << "\n"; // 可能错误
```

`vector` 扩容后，原来的元素地址可能变化。不要长期保存容器内部元素的指针。

---

## 16. 内存初始化

未初始化变量的值是不确定的。

```cpp
int x;
std::cout << x << "\n"; // 不要这样
```

推荐初始化：

```cpp
int x = 0;
int* p = nullptr;
std::vector<int> nums;
```

动态数组初始化：

```cpp
int* arr = new int[3]{}; // 全部初始化为 0
delete[] arr;
```

更推荐：

```cpp
std::vector<int> arr(3); // 全部初始化为 0
```

---

## 17. memcpy、memset 要谨慎

C 里常见：

```cpp
#include <cstring>

int arr[3];
std::memset(arr, 0, sizeof(arr));
```

这对简单整数数组通常可以。

但不要随便对 C++ 对象用 `memset`：

```cpp
std::string s = "hello";
// std::memset(&s, 0, sizeof(s)); // 极其危险
```

因为 `std::string` 内部有自己的资源管理，直接清内存会破坏对象状态。

规则：

- POD/简单字节数据可以谨慎用 `memcpy/memset`
- `std::string`、`std::vector`、自定义类不要乱用
- 优先用构造函数、赋值、标准算法

---

## 18. 内存对齐和 sizeof

`sizeof` 返回类型或对象占用的字节数。

```cpp
#include <iostream>

struct A {
    char c;
    int x;
};

int main() {
    std::cout << sizeof(char) << "\n";
    std::cout << sizeof(int) << "\n";
    std::cout << sizeof(A) << "\n";
}
```

`sizeof(A)` 可能不是 `1 + 4 = 5`，因为编译器会为了访问效率做内存对齐。

初学阶段知道这件事即可。做网络协议、文件格式、嵌入式结构体时再深入。

---

## 19. 调试内存问题

### 19.1 编译器警告

先打开警告：

```bash
g++ main.cpp -std=c++17 -Wall -Wextra -g -o main
```

### 19.2 AddressSanitizer

Linux / MinGW / Clang 常用：

```bash
g++ main.cpp -std=c++17 -fsanitize=address -g -o main
./main
```

它能发现：

- 越界访问
- use-after-free
- 部分内存泄漏

### 19.3 Visual Studio

Windows + Visual Studio 可以用：

- 调试器断点
- 内存窗口
- AddressSanitizer
- CRT Debug Heap

不需要一开始全会。先学会断点、单步、看变量，已经很有用了。

---

## 20. 现代 C++ 内存使用原则

建议按这个顺序选：

| 需求 | 推荐 |
|---|---|
| 普通局部对象 | 直接定义变量 |
| 动态数组 | `std::vector<T>` |
| 固定数组 | `std::array<T, N>` |
| 字符串 | `std::string` |
| 独占动态对象 | `std::unique_ptr<T>` |
| 共享动态对象 | `std::shared_ptr<T>` |
| 只借用对象 | `T&` / `const T&` / `T*` |
| 原始字节缓冲区 | `std::vector<std::byte>` 或 `std::array<std::byte, N>` |

尽量做到：

- 少写裸 `new/delete`
- 少保存裸指针
- 不返回局部变量地址
- 不越界
- 让对象自动释放资源
- 所有权关系写清楚

---

## 21. 小练习

### 练习 1：观察生命周期

写一个类，在构造函数和析构函数里打印日志。分别创建：

- 局部对象
- `std::unique_ptr` 对象
- `std::vector` 中的对象

观察它们什么时候构造、什么时候析构。

### 练习 2：vector 替代 new 数组

把下面代码改成 `std::vector<int>`：

```cpp
int* nums = new int[100];
delete[] nums;
```

### 练习 3：修复内存泄漏

把下面代码改成不会泄漏：

```cpp
void run() {
    int* value = new int(10);
    if (*value > 0) {
        return;
    }
    delete value;
}
```

### 练习 4：找出悬空指针

解释这段代码为什么错：

```cpp
int* makeValue() {
    int x = 10;
    return &x;
}
```

### 练习 5：用 AddressSanitizer 抓越界

写一段数组越界代码，用 `-fsanitize=address` 编译运行，看报错信息。

---

## 22. 和指针专题的关系

[C++指针专题教程](C++指针专题教程.md) 主要讲“地址怎么用”。

这篇主要讲“内存怎么活、怎么释放、怎么别炸”。

学习顺序建议：

1. 先看指针专题，理解 `&`、`*`、`nullptr`
2. 再看本篇，理解栈、堆、生命周期、所有权
3. 然后看类、构造函数、析构函数和 RAII
4. 最后进入智能指针、STL 容器和框架项目
