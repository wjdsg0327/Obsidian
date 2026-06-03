# TypeScript 教程

> TypeScript 是 JavaScript 的超集，添加了静态类型系统和其他特性。

## 目录结构

```
TypeScript/
├── README.md                    ← 本文件
├── 01-基础入门.md               ← 基本类型、变量声明
├── 02-类型系统.md               ← 类型推断、类型断言、类型守卫
├── 03-接口与类型别名.md         ← interface、type、交叉类型
├── 04-函数.md                   ← 函数类型、重载、this
├── 05-类与面向对象.md           ← 类、继承、抽象类、装饰器
├── 06-泛型.md                   ← 泛型函数、泛型约束、工具类型
├── 07-高级类型.md               ← 联合类型、交叉类型、条件类型
├── 08-模块与命名空间.md         ← import/export、命名空间
├── 09-配置与工程化.md           ← tsconfig.json、构建工具
├── 10-React+TypeScript.md       ← React 组件类型
├── 11-Vue3+TypeScript.md        ← Vue3 组合式 API 类型
├── 12-常见问题与技巧.md         ← 实用技巧、避坑指南
```

## 快速开始

### 安装 TypeScript
```bash
# 全局安装
npm install -g typescript

# 项目安装
npm install -D typescript

# 初始化配置
tsc --init
```

### 编译运行
```bash
# 编译单个文件
tsc index.ts

# 监听模式
tsc --watch

# 使用 ts-node 直接运行
npx ts-node index.ts
```

## 为什么学 TypeScript？

| 优势 | 说明 |
|------|------|
| **静态类型** | 编译时发现错误，减少运行时 bug |
| **智能提示** | IDE 自动补全、类型检查 |
| **代码可读性** | 类型即文档，更容易理解代码 |
| **重构安全** | 修改代码时有类型保护 |
| **生态完善** | Vue3、React、Node.js 都原生支持 |

## 学习路线

```
基础类型 → 接口 → 函数 → 类 → 泛型 → 高级类型 → 工程化
    ↓
框架实战（Vue3/React）
    ↓
项目实践
```

## 推荐资源

- [TypeScript 官方文档](https://www.typescriptlang.org/docs/)
- [TypeScript 中文手册](https://www.typescriptlang.org/zh/docs/)
- [TypeScript Deep Dive](https://basarat.gitbook.io/typescript/)
- [Type Challenges](https://github.com/type-challenges/type-challenges)

---

*创建时间：2026-06-03*
