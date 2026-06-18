# Cocos Creator 从零开始快速掌握教程

> 版本基准：Cocos Creator 3.8 LTS，脚本语言默认 TypeScript。  
> 学习目标：用最短路径理解编辑器、节点组件、脚本、UI、动画、物理、资源、发布流程，并能独立做出一个 2D 小游戏原型。

## 0. 你将学到什么

学完这份文档后，你应该能做到：

- 会安装和管理 Cocos Creator 版本。
- 会创建项目、认识项目目录、使用编辑器主要面板。
- 理解“场景、节点、组件、资源、预制体”的关系。
- 会写 TypeScript 组件，理解生命周期函数。
- 会做玩家移动、键盘/触摸输入、碰撞检测、UI 分数显示。
- 会使用 Prefab 批量生成对象。
- 会做基础动画、音效、资源加载。
- 会构建发布 Web、Android 或小游戏平台前的基本检查。
- 会用一个练习项目把知识串起来。

官方资料入口：

- Cocos Creator 3.8 中文手册：https://docs.cocos.com/creator/3.8/manual/zh/
- 新手入门：https://docs.cocos.com/creator/3.8/manual/zh/getting-started/
- 下载安装：https://www.cocos.com/creator-download

## 1. 学习路线总览

建议按 10 天节奏学习。如果每天时间少，也可以拆成 2 到 3 周。

| 阶段 | 内容 | 目标 |
| --- | --- | --- |
| 第 1 天 | 安装、创建项目、认识编辑器 | 能打开项目并运行预览 |
| 第 2 天 | 节点、组件、坐标、层级 | 能手动搭一个简单场景 |
| 第 3 天 | TypeScript 脚本与生命周期 | 能写自己的组件 |
| 第 4 天 | 输入与移动 | 能控制角色移动 |
| 第 5 天 | UI、按钮、Label | 能显示分数和状态 |
| 第 6 天 | Prefab 与对象生成 | 能生成金币、障碍物、敌人 |
| 第 7 天 | 碰撞与物理 | 能处理拾取、受伤、失败 |
| 第 8 天 | 动画、音效、粒子 | 能让游戏更像成品 |
| 第 9 天 | 资源加载、场景切换、数据保存 | 能组织中型项目 |
| 第 10 天 | 构建发布与调试 | 能导出可运行版本 |

## 2. 安装与创建第一个项目

### 2.1 安装

1. 打开 Cocos Creator 下载页。
2. 安装 Cocos Dashboard。
3. 在 Dashboard 中安装 Cocos Creator 3.8.x。
4. 新建项目时选择 2D 模板，方便先上手。

建议：

- 初学先固定一个版本，不要频繁换版本。
- 项目路径不要带太复杂的特殊字符。
- 一个项目对应一个 Creator 版本，升级前先备份。

### 2.2 创建项目

在 Dashboard 中：

1. 点击“新建项目”。
2. 选择 2D 模板。
3. 项目名可以叫 `CatchCoin`。
4. 选择一个容易找到的位置。
5. 点击创建并打开。

打开后先点顶部的预览按钮，确认能正常运行。

## 3. 编辑器核心面板

Cocos Creator 的编辑器可以理解为“可视化搭场景 + 写组件脚本”的组合。

常用面板：

| 面板 | 作用 |
| --- | --- |
| 层级管理器 | 查看当前场景里的所有节点 |
| 场景编辑器 | 拖拽、摆放、缩放、旋转节点 |
| 资源管理器 | 管理图片、脚本、音频、预制体、场景 |
| 属性检查器 | 修改节点和组件参数 |
| 控制台 | 查看日志和报错 |
| 动画编辑器 | 制作关键帧动画 |

初学时你要形成一个习惯：

1. 在资源管理器里准备资源。
2. 在层级管理器里创建节点。
3. 在属性检查器里给节点加组件。
4. 用脚本控制组件行为。

## 4. 项目目录结构

一个典型项目会看到：

```text
assets/       游戏资源与脚本，最常用
settings/     项目设置
extensions/   编辑器扩展，可选
temp/         临时文件，不要手动维护
library/      导入后的缓存，不要手动维护
build/        构建输出
```

你主要关心 `assets/`。建议初学阶段这样组织：

```text
assets/
  scenes/       场景
  scripts/      TypeScript 脚本
  prefabs/      预制体
  textures/     图片
  audio/        音效和音乐
  animations/   动画资源
```

命名建议：

- 场景：`Main.scene`、`Menu.scene`
- 脚本：`PlayerController.ts`、`GameManager.ts`
- 预制体：`Coin.prefab`、`Enemy.prefab`
- 图片：`player.png`、`coin.png`

## 5. 核心概念：场景、节点、组件、资源、预制体

### 5.1 场景 Scene

场景是一张“游戏关卡/页面”。比如：

- `Menu`：主菜单
- `Main`：游戏主场景
- `Result`：结算页面

一个场景里面有很多节点。

### 5.2 节点 Node

节点是场景中的一个对象。角色、金币、按钮、背景、摄像机都可以是节点。

节点本身主要提供：

- 位置
- 旋转
- 缩放
- 父子层级
- 激活/隐藏状态

节点本身不决定“它是什么”，组件决定它能做什么。

### 5.3 组件 Component

组件挂在节点上，用来赋予节点能力。

常见组件：

- `Sprite`：显示图片
- `Label`：显示文字
- `Button`：按钮
- `AudioSource`：播放音频
- `Animation`：播放动画
- 自己写的 TypeScript 组件：控制逻辑

一句话理解：节点是身体，组件是器官和能力。

### 5.4 资源 Asset

图片、音频、脚本、字体、动画、预制体、场景，都是资源。

### 5.5 预制体 Prefab

Prefab 是“可重复创建的节点模板”。

适合做：

- 子弹
- 金币
- 敌人
- 特效
- 弹窗

当一个对象需要反复生成时，优先考虑做成 Prefab。

## 6. 第一个脚本组件

在 `assets/scripts/` 下创建 `HelloCocos.ts`：

```ts
import { _decorator, Component } from 'cc';

const { ccclass } = _decorator;

@ccclass('HelloCocos')
export class HelloCocos extends Component {
    start() {
        console.log('Hello Cocos Creator');
    }
}
```

使用方法：

1. 在层级管理器选中一个节点。
2. 属性检查器点击“添加组件”。
3. 搜索 `HelloCocos`。
4. 运行预览。
5. 打开控制台看日志。

### 6.1 常用生命周期

```ts
import { _decorator, Component } from 'cc';

const { ccclass } = _decorator;

@ccclass('LifeCycleDemo')
export class LifeCycleDemo extends Component {
    onLoad() {
        console.log('节点加载时调用，适合初始化引用');
    }

    start() {
        console.log('第一次 update 前调用，适合初始化游戏状态');
    }

    update(deltaTime: number) {
        console.log('每帧调用，deltaTime 是上一帧到这一帧的秒数');
    }

    onDestroy() {
        console.log('节点销毁时调用，适合取消事件监听');
    }
}
```

常用规则：

- 初始化节点引用：`onLoad`
- 初始化游戏数据：`start`
- 每帧移动或检测：`update`
- 注册事件：`onEnable`
- 取消事件：`onDisable`
- 清理资源：`onDestroy`

## 7. 属性暴露：让脚本参数显示到编辑器

很多参数不应该写死在代码里，而应该暴露到属性检查器中。

```ts
import { _decorator, Component } from 'cc';

const { ccclass, property } = _decorator;

@ccclass('PlayerConfig')
export class PlayerConfig extends Component {
    @property
    speed = 300;

    @property
    hp = 3;
}
```

挂到节点后，你会在属性检查器里看到 `speed` 和 `hp`。

这非常重要。以后你做游戏调参时，不要每改一次数值就重新写代码。

## 8. 玩家移动

目标：用键盘控制一个角色左右上下移动。

### 8.1 准备场景

1. 新建一个 `Main.scene`。
2. 创建一个 `Sprite` 节点，命名为 `Player`。
3. 给 `Player` 放一张图片。
4. 创建脚本 `PlayerController.ts`。
5. 把脚本挂到 `Player` 上。

### 8.2 代码

```ts
import { _decorator, Component, input, Input, EventKeyboard, KeyCode, Vec3 } from 'cc';

const { ccclass, property } = _decorator;

@ccclass('PlayerController')
export class PlayerController extends Component {
    @property
    speed = 320;

    private moveX = 0;
    private moveY = 0;
    private tempPosition = new Vec3();

    onEnable() {
        input.on(Input.EventType.KEY_DOWN, this.onKeyDown, this);
        input.on(Input.EventType.KEY_UP, this.onKeyUp, this);
    }

    onDisable() {
        input.off(Input.EventType.KEY_DOWN, this.onKeyDown, this);
        input.off(Input.EventType.KEY_UP, this.onKeyUp, this);
    }

    update(deltaTime: number) {
        this.node.getPosition(this.tempPosition);
        this.tempPosition.x += this.moveX * this.speed * deltaTime;
        this.tempPosition.y += this.moveY * this.speed * deltaTime;
        this.node.setPosition(this.tempPosition);
    }

    private onKeyDown(event: EventKeyboard) {
        switch (event.keyCode) {
            case KeyCode.KEY_A:
            case KeyCode.ARROW_LEFT:
                this.moveX = -1;
                break;
            case KeyCode.KEY_D:
            case KeyCode.ARROW_RIGHT:
                this.moveX = 1;
                break;
            case KeyCode.KEY_W:
            case KeyCode.ARROW_UP:
                this.moveY = 1;
                break;
            case KeyCode.KEY_S:
            case KeyCode.ARROW_DOWN:
                this.moveY = -1;
                break;
        }
    }

    private onKeyUp(event: EventKeyboard) {
        switch (event.keyCode) {
            case KeyCode.KEY_A:
            case KeyCode.ARROW_LEFT:
            case KeyCode.KEY_D:
            case KeyCode.ARROW_RIGHT:
                this.moveX = 0;
                break;
            case KeyCode.KEY_W:
            case KeyCode.ARROW_UP:
            case KeyCode.KEY_S:
            case KeyCode.ARROW_DOWN:
                this.moveY = 0;
                break;
        }
    }
}
```

练习：

- 把 `speed` 改成 100、500，感受移动速度。
- 加入边界限制，别让玩家跑出屏幕。
- 让角色移动时改变朝向。

## 9. UI：分数、按钮、游戏状态

### 9.1 创建分数文本

1. 在场景中新建 `Canvas`。
2. 在 `Canvas` 下创建 `Label`，命名为 `ScoreLabel`。
3. 创建空节点 `GameManager`。
4. 创建脚本 `GameManager.ts` 并挂到 `GameManager`。

```ts
import { _decorator, Component, Label } from 'cc';

const { ccclass, property } = _decorator;

@ccclass('GameManager')
export class GameManager extends Component {
    @property(Label)
    scoreLabel: Label | null = null;

    private score = 0;

    start() {
        this.refreshScore();
    }

    addScore(value: number) {
        this.score += value;
        this.refreshScore();
    }

    private refreshScore() {
        if (this.scoreLabel) {
            this.scoreLabel.string = `Score: ${this.score}`;
        }
    }
}
```

挂好脚本后，把场景里的 `ScoreLabel` 拖到脚本属性的 `scoreLabel` 槽位。

### 9.2 按钮事件

创建按钮后，在按钮组件的 Click Events 里：

1. 拖入带脚本的节点。
2. 选择组件。
3. 选择公开方法。

示例：

```ts
import { _decorator, Component, director } from 'cc';

const { ccclass } = _decorator;

@ccclass('MenuController')
export class MenuController extends Component {
    startGame() {
        director.loadScene('Main');
    }
}
```

把 `startGame` 绑定到按钮点击事件，就可以点击按钮切换到 `Main` 场景。

## 10. Prefab：批量生成金币

目标：让金币随机出现，玩家碰到后加分。

### 10.1 制作 Coin Prefab

1. 创建一个 `Sprite` 节点，命名为 `Coin`。
2. 设置金币图片。
3. 给它添加碰撞组件，例如 `CircleCollider2D`。
4. 把节点从层级管理器拖到 `assets/prefabs/`。
5. 得到 `Coin.prefab`。
6. 场景里的原始 `Coin` 可以删除。

### 10.2 生成器脚本

```ts
import { _decorator, Component, Prefab, instantiate, Node, Vec3, randomRange } from 'cc';

const { ccclass, property } = _decorator;

@ccclass('CoinSpawner')
export class CoinSpawner extends Component {
    @property(Prefab)
    coinPrefab: Prefab | null = null;

    @property(Node)
    coinRoot: Node | null = null;

    @property
    spawnInterval = 1.2;

    private timer = 0;

    update(deltaTime: number) {
        this.timer += deltaTime;
        if (this.timer >= this.spawnInterval) {
            this.timer = 0;
            this.spawnCoin();
        }
    }

    private spawnCoin() {
        if (!this.coinPrefab || !this.coinRoot) {
            return;
        }

        const coin = instantiate(this.coinPrefab);
        const x = randomRange(-300, 300);
        const y = randomRange(-180, 180);
        coin.setPosition(new Vec3(x, y, 0));
        this.coinRoot.addChild(coin);
    }
}
```

场景准备：

- 创建空节点 `CoinRoot`。
- 创建空节点 `CoinSpawner`。
- 把脚本挂到 `CoinSpawner`。
- 把 `Coin.prefab` 拖到 `coinPrefab`。
- 把 `CoinRoot` 拖到 `coinRoot`。

## 11. 碰撞检测与物理

### 11.1 开启 2D 物理

在项目设置中确认 2D 物理可用。常见流程：

1. 给玩家添加 `RigidBody2D`。
2. 给玩家添加 `Collider2D`，例如 `CircleCollider2D` 或 `BoxCollider2D`。
3. 给金币也添加 `Collider2D`。
4. 确认碰撞分组和碰撞矩阵设置正确。

### 11.2 金币碰撞脚本

给金币挂 `Coin.ts`：

```ts
import { _decorator, Component, Collider2D, Contact2DType, IPhysics2DContact, Node } from 'cc';
import { GameManager } from './GameManager';

const { ccclass, property } = _decorator;

@ccclass('Coin')
export class Coin extends Component {
    @property(Node)
    gameManagerNode: Node | null = null;

    private gameManager: GameManager | null = null;

    onLoad() {
        if (this.gameManagerNode) {
            this.gameManager = this.gameManagerNode.getComponent(GameManager);
        }
    }

    onEnable() {
        const collider = this.getComponent(Collider2D);
        if (collider) {
            collider.on(Contact2DType.BEGIN_CONTACT, this.onBeginContact, this);
        }
    }

    onDisable() {
        const collider = this.getComponent(Collider2D);
        if (collider) {
            collider.off(Contact2DType.BEGIN_CONTACT, this.onBeginContact, this);
        }
    }

    private onBeginContact(self: Collider2D, other: Collider2D, contact: IPhysics2DContact | null) {
        this.gameManager?.addScore(1);
        this.node.destroy();
    }
}
```

注意：Prefab 里直接拖场景节点有时不方便。更常见的做法是由生成器创建金币后，把 `GameManager` 引用赋给金币脚本。初学可以先用简单拖拽理解流程，后面再优化。

## 12. 动画与 Tween

Cocos Creator 有两条常用动画路线：

- 动画编辑器：适合可视化关键帧动画。
- Tween：适合代码控制的简单动画。

### 12.1 Tween 示例：金币上下浮动

```ts
import { _decorator, Component, tween, Vec3 } from 'cc';

const { ccclass } = _decorator;

@ccclass('CoinFloat')
export class CoinFloat extends Component {
    start() {
        const startPos = this.node.position.clone();
        const upPos = startPos.clone().add(new Vec3(0, 20, 0));

        tween(this.node)
            .to(0.5, { position: upPos })
            .to(0.5, { position: startPos })
            .union()
            .repeatForever()
            .start();
    }
}
```

适合用 Tween 的情况：

- 淡入淡出
- 弹窗缩放
- 金币跳动
- 按钮点击反馈
- 简单位移

适合用动画编辑器的情况：

- 角色帧动画
- 复杂 UI 动画
- 多属性关键帧

## 13. 音效与音乐

给节点添加 `AudioSource` 组件，然后通过脚本播放。

```ts
import { _decorator, Component, AudioSource } from 'cc';

const { ccclass, property } = _decorator;

@ccclass('AudioDemo')
export class AudioDemo extends Component {
    @property(AudioSource)
    coinAudio: AudioSource | null = null;

    playCoinSound() {
        this.coinAudio?.playOneShot(this.coinAudio.clip!);
    }
}
```

建议：

- 背景音乐用一个常驻节点管理。
- 短音效用 `playOneShot`。
- 音量设置不要写死，后期做设置界面。

## 14. 资源加载

初学阶段尽量使用编辑器拖拽引用。等项目变大后，再学习动态加载。

### 14.1 编辑器引用

```ts
@property(SpriteFrame)
icon: SpriteFrame | null = null;
```

优点：

- 简单直观。
- 不容易路径写错。
- 适合角色、UI、常驻资源。

### 14.2 resources 动态加载

如果要通过代码加载，需要把资源放进 `assets/resources/`。

```ts
import { _decorator, Component, resources, SpriteFrame } from 'cc';

const { ccclass } = _decorator;

@ccclass('LoadResourceDemo')
export class LoadResourceDemo extends Component {
    start() {
        resources.load('textures/coin/spriteFrame', SpriteFrame, (err, spriteFrame) => {
            if (err) {
                console.error(err);
                return;
            }

            console.log('加载成功', spriteFrame);
        });
    }
}
```

提醒：

- `resources` 很方便，但不要把所有资源都塞进去。
- 大项目要学习 Asset Bundle。
- 动态加载资源后，也要考虑释放。

## 15. 场景切换与全局状态

### 15.1 切换场景

```ts
import { _decorator, Component, director } from 'cc';

const { ccclass } = _decorator;

@ccclass('SceneLoader')
export class SceneLoader extends Component {
    loadMenu() {
        director.loadScene('Menu');
    }

    loadMain() {
        director.loadScene('Main');
    }
}
```

### 15.2 常驻节点

如果你想让音乐管理器、全局数据管理器跨场景存在：

```ts
import { _decorator, Component, director } from 'cc';

const { ccclass } = _decorator;

@ccclass('PersistentRoot')
export class PersistentRoot extends Component {
    onLoad() {
        director.addPersistRootNode(this.node);
    }
}
```

不要滥用常驻节点。只有真正全局的系统才适合：

- 音乐管理
- 玩家存档
- 网络管理
- 全局配置

## 16. 本地数据保存

简单数据可以用 `sys.localStorage`。

```ts
import { _decorator, Component, sys } from 'cc';

const { ccclass } = _decorator;

@ccclass('SaveDemo')
export class SaveDemo extends Component {
    saveHighScore(score: number) {
        sys.localStorage.setItem('highScore', String(score));
    }

    loadHighScore() {
        const value = sys.localStorage.getItem('highScore');
        return value ? Number(value) : 0;
    }
}
```

适合保存：

- 最高分
- 设置项
- 新手引导状态
- 简单关卡进度

不适合保存：

- 大量关卡数据
- 复杂背包系统
- 需要防作弊的核心数据

## 17. 实战项目：接金币小游戏

这是你快速掌握 Cocos Creator 的第一个完整练习。

### 17.1 游戏规则

- 玩家控制角色移动。
- 场景中随机生成金币。
- 玩家碰到金币加 1 分。
- 倒计时 60 秒。
- 时间结束后显示结算界面。

### 17.2 需要的节点

```text
Main
  Canvas
    Background
    Player
    CoinRoot
    UI
      ScoreLabel
      TimeLabel
      ResultPanel
        FinalScoreLabel
        RestartButton
  GameManager
  CoinSpawner
```

### 17.3 脚本拆分

| 脚本 | 挂载节点 | 职责 |
| --- | --- | --- |
| `PlayerController.ts` | Player | 玩家移动 |
| `GameManager.ts` | GameManager | 分数、时间、游戏状态 |
| `CoinSpawner.ts` | CoinSpawner | 随机生成金币 |
| `Coin.ts` | Coin Prefab | 碰撞后加分并销毁 |
| `ResultPanel.ts` | ResultPanel | 显示结算、重新开始 |

### 17.4 GameManager 完整版示例

```ts
import { _decorator, Component, Label, Node, director } from 'cc';

const { ccclass, property } = _decorator;

@ccclass('GameManager')
export class GameManager extends Component {
    @property(Label)
    scoreLabel: Label | null = null;

    @property(Label)
    timeLabel: Label | null = null;

    @property(Node)
    resultPanel: Node | null = null;

    @property(Label)
    finalScoreLabel: Label | null = null;

    @property
    totalTime = 60;

    private score = 0;
    private leftTime = 0;
    private playing = true;

    start() {
        this.leftTime = this.totalTime;
        if (this.resultPanel) {
            this.resultPanel.active = false;
        }
        this.refreshUI();
    }

    update(deltaTime: number) {
        if (!this.playing) {
            return;
        }

        this.leftTime -= deltaTime;
        if (this.leftTime <= 0) {
            this.leftTime = 0;
            this.gameOver();
        }

        this.refreshUI();
    }

    addScore(value: number) {
        if (!this.playing) {
            return;
        }

        this.score += value;
        this.refreshUI();
    }

    restart() {
        director.loadScene('Main');
    }

    private gameOver() {
        this.playing = false;

        if (this.resultPanel) {
            this.resultPanel.active = true;
        }

        if (this.finalScoreLabel) {
            this.finalScoreLabel.string = `Final Score: ${this.score}`;
        }
    }

    private refreshUI() {
        if (this.scoreLabel) {
            this.scoreLabel.string = `Score: ${this.score}`;
        }

        if (this.timeLabel) {
            this.timeLabel.string = `Time: ${Math.ceil(this.leftTime)}`;
        }
    }
}
```

做完这个小游戏，你已经理解了 Cocos Creator 2D 游戏开发的主流程。

## 18. 常见错误与排查

### 18.1 脚本没有出现在添加组件列表

检查：

- 类名、文件名是否清晰一致。
- 是否写了 `@ccclass('类名')`。
- 是否有 TypeScript 编译错误。
- 控制台是否有红色报错。

### 18.2 拖拽属性时拖不进去

检查：

- `@property(Label)` 的类型是不是和拖入对象匹配。
- 你拖的是组件还是节点。
- 例如 `@property(Label)` 要拖带有 Label 组件的节点。

### 18.3 碰撞没有触发

检查：

- 双方是否都有 Collider2D。
- 至少一方是否有 RigidBody2D。
- 碰撞分组矩阵是否允许碰撞。
- Collider 的 sensor、group、mask 设置是否合理。
- 脚本是否注册了 `BEGIN_CONTACT`。

### 18.4 预制体改了但场景没变化

检查：

- 你改的是 Prefab 资源，还是场景中的实例。
- 实例是否覆盖了 Prefab 属性。
- 是否需要应用到 Prefab。

### 18.5 资源路径加载失败

检查：

- 是否放在 `assets/resources/` 下。
- 路径是否不带扩展名。
- SpriteFrame 是否需要写成 `xxx/spriteFrame`。
- 控制台里的真实报错是什么。

## 19. 性能与工程习惯

初学阶段先完成游戏，不要过早优化。但这些习惯要早点养成：

- 不要在 `update` 里频繁 `find` 节点。
- 常用节点引用用 `@property` 拖拽。
- 频繁生成销毁的对象后面要学对象池。
- 图片尺寸尽量合理，不要拿大图当小图用。
- UI、玩法、数据管理尽量拆脚本。
- 一个脚本只做一类事。
- 不要把所有逻辑都塞进 `GameManager`。
- 定期提交 Git，尤其是升级引擎版本前。

## 20. 构建发布

常见发布目标：

- Web Mobile
- Web Desktop
- Android
- iOS
- 微信小游戏等小游戏平台

构建前检查：

1. 场景是否加入构建列表。
2. 首场景是否正确。
3. 资源是否都能加载。
4. 控制台是否无红色报错。
5. 分辨率适配是否正常。
6. 真机或目标浏览器是否测试过。

Web 版本最适合初学测试。移动端发布需要额外配置 SDK、JDK、Android Studio 或平台工具链。

## 21. 继续进阶的主题

当你能独立做完“接金币小游戏”后，再按这个顺序进阶：

1. 对象池：优化子弹、金币、敌人。
2. 状态机：管理角色 idle、run、attack、die。
3. 摄像机跟随：做横版或俯视角地图。
4. TileMap：制作 2D 地图。
5. Spine/DragonBones：接入骨骼动画。
6. Asset Bundle：管理大型项目资源。
7. 存档系统：封装本地数据。
8. 音频管理器：统一控制音乐和音效。
9. UI 框架：弹窗、页面栈、红点、列表。
10. 网络通信：排行榜、登录、联机。
11. 3D 基础：模型、材质、灯光、相机、物理。

## 22. 推荐练习项目

按难度递增：

| 项目 | 训练重点 |
| --- | --- |
| 接金币 | 移动、碰撞、分数、Prefab |
| 打砖块 | 物理、反弹、关卡 |
| 飞机大战 | 子弹、敌人、对象池、碰撞 |
| 横版跳跃 | 重力、地面检测、动画状态 |
| 塔防原型 | 寻路、敌人波次、炮塔攻击 |
| 背包 UI | ScrollView、数据绑定、拖拽 |
| 消消乐 | 网格、交换、匹配检测、掉落 |

## 23. 最小知识地图

你可以用这张知识图检查自己是否入门：

```text
Cocos Creator
  编辑器
    场景
    层级
    属性
    资源
  游戏对象
    Node
    Component
    Prefab
  脚本
    TypeScript
    生命周期
    property 暴露属性
    输入事件
  玩法
    移动
    碰撞
    生成
    分数
  表现
    UI
    动画
    音效
    粒子
  工程
    资源组织
    场景切换
    数据保存
    构建发布
```

## 24. 每日学习打卡表

| 天数 | 任务 | 完成 |
| --- | --- | --- |
| Day 1 | 安装 Creator，创建并运行空项目 | [ ] |
| Day 2 | 创建场景、Player、Background、Canvas | [ ] |
| Day 3 | 写第一个脚本，理解生命周期 | [ ] |
| Day 4 | 实现玩家移动 | [ ] |
| Day 5 | 做分数 UI 和倒计时 UI | [ ] |
| Day 6 | 制作 Coin Prefab 并随机生成 | [ ] |
| Day 7 | 实现碰撞加分 | [ ] |
| Day 8 | 添加金币动画和音效 | [ ] |
| Day 9 | 添加开始、结束、重开流程 | [ ] |
| Day 10 | 构建 Web 版本并测试 | [ ] |

## 25. 入门后你应该能回答的问题

如果下面问题都能回答，说明你已经真正入门：

- Node 和 Component 有什么区别？
- 为什么要用 Prefab？
- `onLoad`、`start`、`update` 分别什么时候用？
- `@property` 有什么用？
- 如何让按钮调用脚本方法？
- 如何随机生成一个金币？
- 碰撞不触发时从哪几步排查？
- 什么资源适合拖拽引用，什么资源适合动态加载？
- Web 构建前要检查哪些东西？
- 为什么不应该把所有代码都写在一个脚本里？

## 26. 你的下一步

打开 Cocos Creator，新建 `CatchCoin` 项目，然后按第 17 章做完“接金币小游戏”。  
不要等所有概念都懂了再动手，Cocos Creator 最快的学习方式就是：拖一个节点，挂一个组件，跑一次预览，看一次报错，改一次代码。

