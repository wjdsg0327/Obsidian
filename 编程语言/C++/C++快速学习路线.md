# C++ 快速学习路线：从语法到框架

> 目标：14 天建立 C++ 基础能力，随后选择一个框架方向做项目。

## 总路线

```text
语法基础
  ↓
指针 / 引用 / 内存 / 生命周期
  ↓
类 / 对象 / RAII / 智能指针
  ↓
STL 容器 / 算法 / Lambda
  ↓
现代 C++：C++11/14/17/20 常用特性
  ↓
CMake / 调试 / 测试 / 第三方库
  ↓
框架实战：OpenCV / Qt / raylib / SFML / Crow / Drogon
```

---

## 学习前准备

建议标准：先学 **C++17**。它够现代、生态支持好，也不会一上来被 C++20/23 的新语法分散注意力。

### Windows 环境

推荐：

- Visual Studio 2022：安装“使用 C++ 的桌面开发”
- CMake：管理项目
- vcpkg：安装第三方 C++ 库
- VS Code：轻量编辑器，可选

验证：

```bash
cl
cmake --version
```

或者使用 MinGW：

```bash
g++ --version
cmake --version
```

### Linux / WSL 环境

```bash
sudo apt update
sudo apt install build-essential cmake gdb
g++ --version
cmake --version
```

---

## 14 天快速路线

### 第 1 天：编译运行与基础语法

学习：

- `main` 函数
- `#include`
- `std::cout` / `std::cin`
- 变量、基础类型、字符串
- `if` / `for` / `while`

练习：

- Hello World
- 输入姓名和年龄并输出
- 判断成绩等级
- 打印 1 到 100 的偶数

参考：[Cpp基础](Cpp基础.md)

---

### 第 2 天：函数、数组、字符串、vector

学习：

- 函数定义与返回值
- 参数传递
- `std::string`
- `std::vector`
- 范围 `for`

练习：

- 写 `add`、`max`、`isPrime` 函数
- 用 `vector` 保存成绩并计算平均分
- 反转字符串

---

### 第 3 天：指针、引用与内存

学习：

- 栈和堆
- 指针 `T*`
- 引用 `T&`
- `const T&`
- `nullptr`

专题笔记：[C++指针专题教程](C++指针专题教程.md)

延伸专题：[C++内存操作专题教程](C++内存操作专题教程.md)

重点理解：

- 引用通常用于“传参不拷贝”
- 裸指针不代表所有权
- 能不用 `new/delete` 就不用

练习：

- 写一个交换两个整数的函数：`void swap(int& a, int& b)`
- 用指针访问数组
- 对比值传递和引用传递

---

### 第 4 天：类、对象、构造函数、析构函数

学习：

- `struct` 与 `class`
- `public` / `private`
- 构造函数
- 析构函数
- 成员函数
- `const` 成员函数

练习：

- 写 `User` 类
- 写 `Student` 类，保存姓名和成绩
- 写 `Timer` 类，在析构时打印耗时

---

### 第 5 天：面向对象与 RAII

学习：

- 封装
- 继承
- 虚函数与多态
- `override`
- RAII

重点：

- C++ 的核心不是“到处继承”，而是“对象生命周期和资源管理”
- 文件、锁、内存、网络连接都应该用对象自动管理

练习：

- 写 `Shape` 基类，派生 `Circle` 和 `Rectangle`
- 用虚函数计算面积
- 写一个自动打开/关闭文件的类

---

### 第 6 天：STL 容器

学习：

- `std::vector`
- `std::array`
- `std::map`
- `std::unordered_map`
- `std::set`
- `std::queue`
- `std::stack`

练习：

- 用 `unordered_map` 统计单词出现次数
- 用 `set` 给数字去重
- 用 `queue` 模拟任务队列

---

### 第 7 天：STL 算法与 Lambda

学习：

- `std::sort`
- `std::find`
- `std::count_if`
- `std::transform`
- Lambda 表达式

练习：

- 对学生按分数排序
- 过滤出所有及格学生
- 把一组字符串统一转成小写

---

### 第 8 天：现代 C++ 常用特性

学习：

- `auto`
- 范围 `for`
- `nullptr`
- `enum class`
- 结构化绑定
- `std::optional`
- `std::variant`

练习：

- 用 `optional` 表示“查找用户可能失败”
- 用结构化绑定遍历 `map`

---

### 第 9 天：智能指针与移动语义

学习：

- `std::unique_ptr`
- `std::shared_ptr`
- `std::weak_ptr`
- 拷贝与移动
- `std::move`

重点：

- 默认用 `unique_ptr`
- 只有确实需要共享所有权时才用 `shared_ptr`
- 不要把 `std::move` 理解成“移动内存”，它更像是允许资源转移

练习：

- 用 `unique_ptr` 管理一个对象
- 写一个不能拷贝但可以移动的类

---

### 第 10 天：文件、异常、日志、JSON

学习：

- `fstream`
- 异常处理 `try/catch`
- 错误返回与异常的取舍
- 第三方库：`nlohmann/json`、`fmt`

练习：

- 读取文本文件并统计行数
- 保存学生列表到 JSON
- 读取 JSON 并打印表格

---

### 第 11 天：CMake 与项目结构

学习：

- `CMakeLists.txt`
- `add_executable`
- `target_link_libraries`
- `include` / `src` / `tests` 目录结构

推荐结构：

```text
my_project/
├── CMakeLists.txt
├── include/
├── src/
├── tests/
└── README.md
```

练习：

- 把前面的学生管理程序整理成 CMake 项目
- 拆分 `.h` 和 `.cpp`

---

### 第 12 天：调试、测试与代码质量

学习：

- gdb / Visual Studio 调试
- 断点、单步、变量查看
- GoogleTest
- `clang-format`
- `clang-tidy`

练习：

- 给 `isPrime`、成绩统计函数写单元测试
- 用调试器找一个数组越界问题

---

### 第 13 天：框架选择

按目标选框架：

| 方向 | 推荐 | 为什么 |
|---|---|---|
| 图像处理 / AI 视觉 | OpenCV | 资料多，项目感强，适合快速出成果 |
| 桌面应用 | Qt | 成熟跨平台 GUI 框架 |
| 小游戏 / 可视化 | raylib | 上手快，API 简单 |
| 多媒体 / 游戏 | SFML | 图形、音频、窗口、网络都比较清晰 |
| 工具 UI / 编辑器 | Dear ImGui | 很适合调试面板和内部工具 |
| Web API | Crow / Drogon | C++ Web 服务方向 |
| 网络编程 | Boost.Asio | 异步网络基础能力强 |

建议第一阶段先选 **OpenCV** 或 **raylib**。它们反馈快，不容易陷进复杂工程配置。

---

### 第 14 天：做第一个项目

推荐项目：

| 项目 | 用到知识 |
|---|---|
| 学生管理系统 CLI | 基础语法、类、STL、文件 |
| 图片批量处理工具 | CMake、文件、OpenCV |
| 摄像头人脸检测 | OpenCV、视频流、图像处理 |
| 2D 贪吃蛇 | raylib/SFML、类、游戏循环 |
| 简单 Todo API | Crow/Drogon、JSON、路由 |

如果你想跟知识库现有资料衔接，优先做：

1. OpenCV 图片读取与显示
2. 摄像头颜色跟踪
3. 人脸检测
4. 车道线检测

参考：[OpenCV项目实战指南](OpenCV项目实战指南.md)

---

## 4 周完整路线

如果不赶时间，建议这样安排：

| 周数 | 目标 | 产出 |
|---|---|---|
| 第 1 周 | 基础语法 + 指针引用 + 类 | 学生管理系统 CLI |
| 第 2 周 | STL + 现代 C++ + 文件 JSON | 可保存数据的小工具 |
| 第 3 周 | CMake + 调试 + 测试 + 第三方库 | 结构清晰的 CMake 项目 |
| 第 4 周 | 框架实战 | OpenCV / Qt / raylib 小项目 |

---

## 学完后你应该能做到

基础能力：

- 能写 C++17 命令行程序
- 能看懂常见类、函数、STL 代码
- 能理解引用、指针、智能指针的区别
- 能用 `vector`、`map`、`unordered_map`、`sort` 等标准库

工程能力：

- 能用 CMake 创建项目
- 能拆分头文件和源文件
- 能接入第三方库
- 能做简单调试和单元测试

框架能力：

- 能选择合适的 C++ 框架
- 能跑通 OpenCV / Qt / raylib / Crow 这类项目的最小 Demo
- 能基于示例扩展成自己的项目

---

## 不建议一开始深挖的内容

这些内容重要，但不适合第一轮学习就钻太深：

- 模板元编程
- SFINAE
- C++20 Concepts
- 协程
- 多线程底层细节
- 操作系统级内存模型
- 复杂继承体系

第一轮目标是“能写项目”，不是“读懂标准委员会论文”。

---

## 学习资料入口

- [Cpp基础](Cpp基础.md)
- [C++指针专题教程](C++指针专题教程.md)
- [C++内存操作专题教程](C++内存操作专题教程.md)
- [C++学习资源指南](C++学习资源指南.md)
- [OpenCV项目实战指南](OpenCV项目实战指南.md)

一句话路线：**先 C++17 基础，再 STL 和 RAII，然后 CMake，最后用 OpenCV 或 raylib 做项目。**
