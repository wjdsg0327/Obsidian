# Superpowers & gstack 使用指南

> 两个针对 **AI 编程 Agent** 的工作流框架，让你的 coding agent 从一个"会写代码的机器人"变成一个"完整的开发团队"。

---

## 一、Superpowers

**作者：** Jesse Vincent (obra)  
**仓库：** <https://github.com/obra/superpowers>  
**定位：** 完整的软件开发方法论 + 可组合技能集，专注 **TDD + 规范驱动** 的开发流程

### 1.1 核心理念

Superpowers 让 coding agent 不再是"上来就写代码"，而是：

1. **先问清楚需求** — 用苏格拉底式提问帮你理清真正想要什么
2. **写设计文档** — 分块输出，让你能真正读完
3. **制定实现计划** — 分解成小任务，每项 2-5 分钟
4. **子 Agent 驱动开发** — 自动派发子 agent 逐项实现 + 审查
5. **始终坚持 TDD** — RED-GREEN-REFACTOR，先写测试再写代码
6. **代码审查** — 每完成一项任务自动审查，阻断关键问题

### 1.2 支持的平台

| 平台 | 安装方式 |
|------|----------|
| **Claude Code** | `/plugin install superpowers@claude-plugins-official` 或插件市场 |
| **Codex CLI / Codex App** | `/plugins` → 搜索 `superpowers` |
| **Cursor** | `/add-plugin superpowers` |
| **Gemini CLI** | `gemini extensions install https://github.com/obra/superpowers` |
| **Factory Droid** | `droid plugin install superpowers@superpowers` |
| **OpenCode** | 执行安装脚本 |
| **GitHub Copilot CLI** | `copilot plugin install superpowers@superpowers-marketplace` |

### 1.3 技能列表

Superpowers 的技能是**自动触发的** — 你什么也不用做，agent 会根据任务自动选择合适的技能。

#### 🧠 核心开发流程技能（自动按顺序触发）

| 技能 | 触发时机 | 功能 |
|------|----------|------|
| **brainstorming** | 写代码前 | 苏格拉底式需求探询，输出分块设计文档 |
| **using-git-worktrees** | 设计通过后 | 创建隔离工作区和新分支，跑通基线测试 |
| **writing-plans** | 设计确认后 | 拆解为 2-5 分钟的小任务，每项含文件路径、完整代码、验证步骤 |
| **subagent-driven-development** | 开始实现 | 派发子 agent 逐项实现，两步审查（合规性→代码质量） |
| **test-driven-development** | 实现中 | 强制 RED-GREEN-REFACTOR 循环 |
| **requesting-code-review** | 任务间 | 按计划审查，按严重级别报告问题 |
| **finishing-a-development-branch** | 全部完成 | 验证测试，提供选项（合并/PR/保留/丢弃） |

#### 🔧 其他实用技能

| 分类 | 技能 | 说明 |
|------|------|------|
| **调试** | systematic-debugging | 4 阶段根因分析 |
| **调试** | verification-before-completion | 确保问题真正修复 |
| **协作** | executing-plans | 分批执行 + 人工检查点 |
| **协作** | dispatching-parallel-agents | 并发子 agent 工作流 |
| **协作** | requesting-code-review | 审查前检查清单 |
| **协作** | receiving-code-review | 回应反馈 |
| **元能力** | writing-skills | 创建新的技能 |
| **元能力** | using-superpowers | 技能系统入门 |

### 1.4 使用示例

安装后**什么都不用做**，正常对话即可：

```
你：帮我做一个待办事项管理应用
Agent：[brainstorming 自动触发]
      → 提问你的真实需求、目标用户、关键功能
      → 输出设计文档，分块让你确认
你：确认方案
Agent：[writing-plans 自动触发]
      → 生成详细实现计划，拆解为多个小任务
你：开干
Agent：[subagent-driven-development + TDD 自动触发]
      → 逐个任务派发子 agent 实现
      → 先写测试 → 写代码 → 审查 → 提交
      → 自动运行数小时，不需要你干预
```

### 1.5 核心原则

- **TDD 优先** — 永远先写测试
- **系统化 > 拍脑袋** — 过程驱动而非猜测
- **复杂度最低化** — 简洁是第一目标
- **用证据说话** — 完成前先验证

---

## 二、gstack

**作者：** Garry Tan（Y Combinator CEO）  
**仓库：** <https://github.com/garrytan/gstack>  
**定位：** 把 Claude Code 变成一支虚拟工程团队（CEO + 架构师 + 设计师 + QA + 安全 + 运维）

### 2.1 核心理念

gstack 不是"一个工具"，而是**一套完整的软件工厂流程**：

> **Think → Plan → Build → Review → Test → Ship → Reflect**

每个 `/命令` 都是一个角色。命令之间天然衔接——上一个的输出自动成为下一个的输入。

### 2.2 安装

#### 🚀 基础安装（Claude Code）

```bash
git clone --single-branch --depth 1 \
  https://github.com/garrytan/gstack.git ~/.claude/skills/gstack \
  && cd ~/.claude/skills/gstack \
  && ./setup
```

然后在 `CLAUDE.md` 中添加 gstack 节，声明可用技能。

**前置依赖：** Claude Code、Git、Bun v1.0+

#### 👥 团队模式（推送到仓库）

```bash
(cd ~/.claude/skills/gstack && ./setup --team) \
  && ~/.claude/skills/gstack/bin/gstack-team-init required \
  && git add .claude/ CLAUDE.md \
  && git commit -m "require gstack for AI-assisted work"
```

#### 🔌 其他平台

```bash
git clone --single-branch --depth 1 https://github.com/garrytan/gstack.git ~/gstack
cd ~/gstack && ./setup --host <name>
```

| 平台 | flag |
|------|------|
| Codex CLI | `--host codex` |
| OpenCode | `--host opencode` |
| Cursor | `--host cursor` |
| Factory Droid | `--host factory` |
| Slate / Kiro / Hermes | 对应 `--host` 参数 |

#### 🤖 OpenClaw 集成

在 `AGENTS.md` 中添加 Coding Tasks 节：

```
安全审计 → "Load gstack. Run /cso"
代码审查 → "Load gstack. Run /review"
QA 测试 → "Load gstack. Run /qa https://..."
完整功能 → "Load gstack. Run /autoplan, implement, then /ship"
```

OpenClaw 原生技能（ClawHub 安装）：

```bash
clawhub install gstack-openclaw-office-hours \
  gstack-openclaw-ceo-review \
  gstack-openclaw-investigate \
  gstack-openclaw-retro
```

### 2.3 技能详解（23+ 个斜杠命令）

#### 🅰️ 规划阶段（Think → Plan）

| 命令 | 角色 | 功能 |
|------|------|------|
| **`/office-hours`** | 🏢 YC Office Hours | 6 个强力问题重新框定产品。挑战假设，生成实现方案，输出设计文档。 |
| **`/plan-ceo-review`** | 👑 CEO | 重新思考问题。4 种模式：Expansion / Selective Expansion / Hold Scope / Reduction。 |
| **`/plan-eng-review`** | 🏗️ 工程经理 | 锁定架构、ASCII 图表、数据流、边界情况、失败模式。 |
| **`/plan-design-review`** | 🎨 高级设计师 | 维度打分（0-10），AI Slop 检测，交互式确认。 |
| **`/plan-devex-review`** | 👩‍💻 DX 主管 | 开发者体验审查，竞品 TTHW 基准测试，摩擦点追踪。 |
| **`/design-consultation`** | 💡 设计合伙人 | 从零构建完整设计系统。 |
| **`/autoplan`** | 🔄 审查管道 | 一键全流程规划：CEO → 设计 → 工程。 |

#### 🅱️ 实现与审查阶段（Build → Review）

| 命令 | 角色 | 功能 |
|------|------|------|
| **`/review`** | 🔍 高级工程师 | 找能过 CI 但上线会炸的 bug。自动修复 + 标记完整性问题。 |
| **`/investigate`** | 🐞 调试专家 | 系统化根因分析。铁律：不调查不修复。 |
| **`/design-review`** | 🎨 会写代码的设计师 | 审查 + 直接修复 + before/after 截图。 |
| **`/design-shotgun`** | 🎯 设计探索器 | 生成 4-6 个原型变体，浏览器对比，品味记忆学习。 |
| **`/design-html`** | 🖥️ 设计工程师 | 原型图 → 可上线的 HTML（30KB 零依赖）。 |
| **`/cso`** | 🔒 首席安全官 | OWASP Top 10 + STRIDE 威胁建模，独立验证。 |
| **`/codex`** | 🔄 第二意见 | Codex 独立代码审查，三模式。 |
| **`/pair-agent`** | 🤝 多 Agent 协调 | 共享浏览器给任意 AI agent。 |

#### 🅲 测试与发布阶段（Test → Ship）

| 命令 | 角色 | 功能 |
|------|------|------|
| **`/qa`** | ✅ QA 主管 | 真实浏览器测试 → 修复 → 回归测试。 |
| **`/qa-only`** | 📋 QA 报告员 | 同上但只出报告不改代码。 |
| **`/ship`** | 🚢 发布工程师 | Sync main → 跑测试 → 审计覆盖率 → 开 PR。 |
| **`/land-and-deploy`** | 🚀 发布工程师 | 合并 PR → CI → 部署 → 验证线上健康。 |
| **`/canary`** | 📡 SRE | 部署后监控：console 错误、性能回归、页面故障。 |
| **`/benchmark`** | ⚡ 性能工程师 | Core Web Vitals 基准 + PR 前后对比。 |
| **`/browse`** | 👁️ QA 工程师 | 真实 Chromium 浏览器操作与截图。 |
| **`/spec`** | 📝 规格撰写者 | 模糊想法 → 精确可执行规格文档。五阶段流程。 |

#### 🅳 反思与维护阶段（Reflect）

| 命令 | 角色 | 功能 |
|------|------|------|
| **`/retro`** | 📊 工程经理 | 周度回顾，跨项目汇总。 |
| **`/document-release`** | ✍️ 技术文档作者 | 自动更新文档、发现过期 README。 |
| **`/document-generate`** | 📖 文档作者 | 从零生成缺失文档。 |
| **`/learn`** | 🧠 记忆管理 | 管理跨 session 项目知识。 |

#### 🛡️ 安全防护工具

| 命令 | 功能 |
|------|------|
| **`/careful`** | 危险操作前警告（rm -rf, DROP TABLE 等） |
| **`/freeze`** | 限制文件编辑到单目录 |
| **`/guard`** | `/careful` + `/freeze` 合体 |
| **`/unfreeze`** | 解除限制 |

---

## 三、结合使用案例

> 💡 **黄金搭档：Superpowers 管 TDD 流水线，gstack 管角色化质量门。**
> 前者自动运行，后者按需指挥。两者互补，互不冲突。

---

### 案例 1：从零构建「待办事项 + 习惯追踪」Web 应用

**一句话 → 上线 staging，含测试 + PR + 文档**

#### 阶段 A：需求探索 + 产品策略

```text
👤 打开 Claude Code，说：
"我想做一个待办事项加习惯追踪的应用"

━━━ 【Superpowers → brainstorming 自动触发】 ━━━

Claude 不会直接写代码。它会说：
"我们先聊聊。你说的'习惯追踪'具体指什么？
 - 你现在用什么管任务？
 - 想养成什么习惯？
 - 用户只有你还是开放注册？"

经过苏格拉底式对话，输出分块设计文档：
→ 功能清单（CRUD、习惯打卡、热力图、导出）
→ 技术栈建议（React + SQLite/Supabase）
→ MVP vs 完整版 分阶段路线
→ 你在聊天里逐块确认 ✅
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

👤 确认后说："我们可以开始规划了"

【Superpowers → writing-plans】
→ 拆解为 8-12 个小任务，每项 2-5 分钟
→ 每个任务含：文件路径 + 完整代码 + 验证步骤

━━━ 【gstack 介入 — 产品/架构层面审视】 ━━━

👤 输入：/plan-ceo-review

→ 读 Superpowers 的设计文档
→ 挑战范围："热力图有必要做在第一版吗？"
→ 输出项目精简建议 + 风险清单

👤 同意缩减方案后输入：/plan-eng-review

→ ASCII 架构图（组件树、数据流、DB schema）
→ 状态机图（任务流转）
→ 边界情况清单（空状态、并发、首次加载）
→ 测试策略

═════════════════════════════════════════
阶段小结：
Superpowers → 需求摸清 + 设计文档 + 实现计划
gstack      → CEO 砍范围 + 工程经理锁定架构
═════════════════════════════════════════
```

#### 阶段 B：实现（Superpowers TDD 流水线）

```text
👤 审查后说："开干"

【Superpowers → subagent-driven-development + TDD 自动触发】

对每个小任务：
  1. 创建 git worktree（隔离分支）
  2. 派发子 agent 执行：
     - 写测试 → 运行 → 看它失败（RED）
     - 写最简实现 → 运行 → 看它通过（GREEN）
     - 重构 → 跑测试 → 确认不破坏（REFACTOR）
  3. 发起代码审查 → 通过后进入下个任务
  4. 每个任务 atomic commit

☕ 这阶段你可以去喝杯咖啡，Claude 自动运行 1-2 小时
```

#### 阶段 C：审查 + QA + 发布（gstack 接管）

```text
👤 实现完成：
你：/review

【gstack → Staff Engineer】
→ [AUTO-FIXED] 未使用的 import、命名不一致
→ [ASK] "习惯打卡数据量大，建议加复合索引"
→ 你确认 → 自动创建索引

👤 部署 staging：
你：/qa https://staging.myapp.com

【gstack → QA Lead — 真实浏览器测试】
→ 打开 Chromium，截图每个页面
→ 新增待办 → 勾选完成 → 打卡 → 查看热力图
→ 发现 bug：热力图日期偏移一天
→ 自动原子修复 + 添加回归测试

👤 确认后：
你：/ship

【gstack → Release Engineer】
→ 同步 main → 全套测试（42→56 新增14个）→ 覆盖率审计
→ 创建 PR + 自动生成描述（变更列表、截图、测试报告）

👤 事后：
你：/retro
→ 自动生成回顾报告（速度统计、问题分类、改进建议）

💡 核心洞察：
   Superpowers 确保每行代码都有测试覆盖 🧪
   gstack 确保代码经过多角色审视后才上线 🔍
   → 两者交集 = 质量的铁壁
```

---

### 案例 2：给现有项目加「通知推送」功能

**场景：** 已有项目（DB + API + 前端），新增 Slack 通知 + 站内通知

```text
👤 打开 Claude Code：
"给项目加通知功能：用户完成动作时发 Slack 消息
  + 站内顶部铃铛通知"

【Superpowers → brainstorming 自动触发】
→ 读现有代码了解架构
→ 追问：哪些动作触发？频率要求？
→ 输出设计文档：
   - 通知模板系统
   - 异步队列（防阻塞）
   - DB 迁移（notifications 表）

👤 确认后：
你：/spec

【gstack → 规格撰写】
→ 五阶段：why → scope → 技术调研 → draft → file
→ 自动读现有代码确保架构一致
→ Codex 质量门控（< 7/10 不通过）

👤 确认规格：
你：/plan-eng-review

【gstack → 工程经理】
→ 架构影响分析：改动波及哪些模块？
→ 异步队列方案（RabbitMQ / Redis Stream）
→ DB 迁移计划 + 回滚方案

👤 确认后说："开始实现"

【Superpowers → TDD 流水线】
→ writing-plans 拆解任务
→ subagent-driven-development:
   Task 1: notifications 表 + 迁移（先测迁移）
   Task 2: 通知模型 + API（先测 ORM）
   Task 3: 异步队列（先测队列）
   Task 4: Slack 集成（先测发消息）
   Task 5: 前端铃铛组件（先测组件）

👤 实现完成：
你：/review
你：/cso

【gstack → 安全官】
→ OWASP Top 10 + STRIDE
→ 发现："通知 API 缺频率限制，可能被刷"
→ 自动添加速率限制

你：/qa https://staging.myapp.com
→ 真实浏览器：完成动作 → 检查通知 → 验证 Slack 发出

你：/ship
→ 测试 91 → 106
→ PR: github.com/you/app/pull/14
```

---

### 案例 3：修复棘手的生产 Bug

**场景：** 订单结算页偶尔白屏，本地无法复现

```text
👤 你说：
"线上有个结算白屏 bug，偶尔出现，本地测不出来"

【Superpowers → systematic-debugging 自动触发】
→ 4 阶段根因分析：收集信息 → 假设 → 测试 → 确认
→ 追问：什么浏览器？频率？最近上线了什么？

👤 提供上下文后：
你：/investigate

【gstack → 调试专家】
→ 铁律：不调查清楚决不修复
→ 追踪数据流：请求→路由→控制器→服务→DB→模板
→ 发现：第三方 API 偶尔超时，代码没 catch
→ 输出根因报告

👤 确认根因后：

【Superpowers → TDD 修复】
→ 写模拟超时测试 → RED
→ 加错误处理 + 优雅降级 → GREEN
→ 重构 → REFACTOR

你：/review
→ gstack 检查遗漏的边界情况

你：/qa https://staging.myapp.com/settle
→ 浏览器验证结算流程正常

你：/ship
→ 自动部署
```

---

### 案例 4：重构历史遗留模块

**场景：** 遗留的用户模块，500 行的巨型方法，没测试

```text
👤 你说：
"重构 user 模块，那个 500 行的 updateUser 方法
  太可怕了，需要拆开而且加上测试"

【Superpowers → brainstorming 自动触发】
→ 先读完整代码，理解现有行为
→ 输出重构设计文档：
   - 现有问题清单（耦合、副作用、错误处理）
   - 拆分方案（验证/业务/持久化三层）
   - 迁移策略（Strangler Fig 模式）

👤 确认后：
你：/plan-eng-review

【gstack → 工程经理】
→ 检查重构对下游的影响
→ 输出灰度切换方案 + 回滚计划

👤 确认后说："开始重构"

【Superpowers → TDD 流水线】
→ 第一步：先给现有代码写特性测试（Characterization Tests）
   锁定当前行为，防止重构跑偏
→ 第二步：逐块提取为独立函数/类，每个提取都：
   写测试 → RED → 实现 → GREEN → REFACTOR
→ 第三步：全部重构完后跑完整测试套件

你：/review
→ gstack 检查重构是否引入新问题

你：/benchmark
→ 重构前 vs 重构后性能对比

你：/ship
```

---

### 案例 5：一个完整 Sprint 的双工具配合时间线

```text
Day 1 — 产品规划
  你：说想法 → Superpowers brainstorming → 设计文档
  你：/office-hours → gstack 重新框定产品思路
  你：/plan-ceo-review → 砍掉 30% 不必要的功能

Day 2 — 技术规划
  你：/plan-eng-review → 架构定型
  你：/autoplan → 一键全流程审查
  你：确认 → Superpowers writing-plans → 拆解为 15 个任务

Day 3-4 — 开发
  Superpowers 自动派发子 agent 逐个实现
  每个任务：TDD 严格 RED-GREEN-REFACTOR
  你只需要每天 review 两次进度

Day 5 — 审查 + QA
  你：/review → 代码审查 + 自动修复
  你：/cso → 安全审计
  你：/qa staging_url → 浏览器端到端测试

Day 5 pm — 发布
  你：/ship → PR 自动创建
  你：review PR → merge → /land-and-deploy

Day 6 — 回顾
  你：/retro → 回顾报告
  gstack + Superpowers 的速度统计
  → 对比之前不带工具的开发速度
```

---

## 四、对比总结

| 维度 | Superpowers | gstack |
|------|-------------|--------|
| **作者** | Jesse Vincent (Prime Radiant) | Garry Tan (Y Combinator CEO) |
| **哲学** | TDD + 方法论驱动，自动触发 | 角色扮演 + 斜杠命令，手动调用 |
| **上手方式** | 安装后自动运行 | 安装后用 `/命令` 触发 |
| **核心特色** | 严格 TDD、子 agent 驱动、技能链自动串联 | 多角色虚拟工程团队、真实浏览器 QA、一键发布 |
| **最适合** | 需要严谨 TDD 和规范的项目 | 需要团队式工作流、快速迭代的项目 |
| **冲突吗？** | **不冲突！** 可以同时安装使用。Superpowers 管 TDD 流程，gstack 管角色分工和发布。 |

### 分工协作图

```
           用户说想法
               │
               ▼
     ┌─────────────────┐
     │  Superpowers     │  ← 自动触发
     │  brainstorming   │     输出设计文档
     └────────┬────────┘
               │
               ▼
     ┌─────────────────┐
     │  gstack          │  ← /office-hours /plan-ceo-review
     │  产品/架构审视   │     砍范围、锁架构
     └────────┬────────┘
               │
               ▼
     ┌─────────────────┐
     │  Superpowers     │  ← 自动触发
     │  TDD 流水线      │     writing-plans → subagent + TDD
     └────────┬────────┘
               │
               ▼
     ┌─────────────────┐
     │  gstack          │  ← /review /cso /qa /ship
     │  质量门 + 发布   │     审查 → 安全 → 测试 → 上线
     └─────────────────┘
```

## 五、快速安装命令汇总

### Superpowers

```bash
# Claude Code 插件市场
/plugin install superpowers@claude-plugins-official

# 或先添加市场
/plugin marketplace add obra/superpowers-marketplace
/plugin install superpowers@superpowers-marketplace
```

### gstack

```bash
# Claude Code
git clone --single-branch --depth 1 \
  https://github.com/garrytan/gstack.git ~/.claude/skills/gstack \
  && cd ~/.claude/skills/gstack && ./setup

# OpenClaw 技能
clawhub install gstack-openclaw-office-hours gstack-openclaw-ceo-review \
  gstack-openclaw-investigate gstack-openclaw-retro
```

---

> **一句话总结：**
> - **Superpowers** = 一个严格遵循 TDD 的开发方法论，装上就不用管了
> - **gstack** = 一个虚拟工程团队，用斜杠命令随时调动 CEO/架构师/QA/安全/运维
> - **两者可以一起用** — Superpowers 管开发流水线，gstack 管质量门和团队角色
