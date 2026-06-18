# 03-TypeScript 脚本基础

## 为什么 Cocos Creator 用 TypeScript

TypeScript 可以理解成“带类型的 JavaScript”。

好处：

- 编辑器提示更好
- 错误更早暴露
- 大项目更容易维护

你不用先成为 TS 高手。入门 Cocos 只要掌握一小部分。

## 变量

```ts
let score = 0;
let playerName = 'hero';
let isGameOver = false;
```

带类型写法：

```ts
let score: number = 0;
let playerName: string = 'hero';
let isGameOver: boolean = false;
```

## 函数

```ts
function addScore(value: number) {
  score += value;
}
```

类里的方法：

```ts
class GameManager {
  addScore(value: number) {
    console.log(value);
  }
}
```

## 类

Cocos 脚本通常是一个类：

```ts
@ccclass('GameManager')
export class GameManager extends Component {
  score = 0;

  addScore(value: number) {
    this.score += value;
  }
}
```

`this` 表示当前这个组件实例。

## 导入 Cocos API

你要用什么，就从 `cc` 里导入什么：

```ts
import { _decorator, Component, Node, Vec3 } from 'cc';
```

常见导入：

- Component：组件基类
- Node：节点类型
- Vec3：三维向量，2D 里也常用于位置
- Label：文字组件
- Prefab：预制体
- instantiate：实例化预制体

## property 装饰器

想在编辑器里拖节点或填参数，就用 `@property`。

```ts
@property
speed: number = 300;

@property(Node)
player: Node | null = null;
```

这样不用把所有东西写死在代码里。

## null 是什么意思

```ts
player: Node | null = null;
```

意思是：

- player 可以是 Node
- 也可以暂时是 null

使用前最好判断：

```ts
if (!this.player) return;
this.player.setPosition(0, 0, 0);
```

## update 和 deltaTime

`update(deltaTime)` 每一帧执行。

```ts
update(deltaTime: number) {
  this.node.translate(new Vec3(100 * deltaTime, 0, 0));
}
```

为什么要乘 `deltaTime`？

因为不同机器帧率不同。乘了以后，速度按“每秒”计算，而不是按“每帧”计算。

## 常见错误

### 忘记导入

用了 `Vec3`，但没 import，会报错。

### 脚本类名和文件理解混乱

文件叫 `PlayerController.ts`，类最好也叫 `PlayerController`。

### 改了脚本但编辑器没刷新

保存脚本后等编辑器编译完成，Console 不报错再运行。

## 本章练习

写一个 `ScoreTest.ts`：

- 有 `score = 0`
- `start()` 里加 10 分
- 打印当前分数

```ts
console.log('score:', this.score);
```
