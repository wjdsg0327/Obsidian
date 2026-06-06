# Taste Skill 技能说明与教程

> 整理日期：2026-06-06  
> 项目仓库：<https://github.com/Leonxlnx/taste-skill>  
> 官网：<https://tasteskill.dev>  
> 许可证：MIT  
> 类型：AI Agent Skills / 前端设计审美技能包

---

## 1. 一句话说明

**Taste Skill 是一组给 AI 编程 Agent 使用的“前端审美 / 反模板味 UI”技能。**

它的目标是让 Claude Code、Codex、Cursor 等 Agent 在生成界面时，不再产出常见的 AI 模板味页面，而是更注重：

- 布局层次；
- 字体选择；
- 间距节奏；
- 动效质感；
- 视觉密度；
- 设计语言；
- 响应式细节；
- 反 AI 味的 UI 限制。

简单说：

> Taste Skill 不是一个 UI 组件库，而是一套让 AI 写前端时“更有审美”的规则包。

---

## 2. 它解决什么问题

AI 写前端很容易出现这些问题：

- 一眼 AI 味：紫蓝渐变、霓虹光、模板 SaaS hero；
- 默认字体和默认卡片太多；
- 三列等宽卡片反复出现；
- 间距、层级、留白都很平均，没有设计重点；
- 动效要么没有，要么很廉价；
- 页面看起来“能用”，但不像真实设计师做的；
- 生成代码时经常半成品、placeholder、注释代替真实实现；
- 根据截图还原页面时，缺少图像分析和设计系统抽取步骤。

Taste Skill 把这些审美判断写成可安装的 `SKILL.md`，让 Agent 在合适任务中自动加载。

---

## 3. 项目定位

官方 README 的定位是：

> Portable Agent Skills that upgrade AI-built interfaces: stronger layout, typography, motion, and spacing instead of boilerplate-looking UIs.

中文理解：

> 一组可移植的 Agent 技能，用来升级 AI 生成界面的布局、字体、动效和间距，避免样板化 UI。

它还包含面向图像生成的技能，用于生成：

- 网站参考图；
- 移动端界面参考图；
- 品牌套件 / brand kit；
- 图片到代码的流程。

---

## 4. 适合哪些工具

README 明确提到支持或适合：

- Codex
- Cursor
- Claude / Claude Code
- ChatGPT Images
- 其他可以读取/安装 Agent Skill 的工具

它是 framework agnostic，即不绑定 React、Vue、Svelte 等某一个框架。规则主要约束设计意图和输出质量，而不是框架 API。

---

## 5. 安装方式

### 5.1 安装全部技能

使用 Vercel Labs 的 `skills` CLI：

```bash
npx skills add https://github.com/Leonxlnx/taste-skill
```

这个命令会扫描仓库里的 `skills/` 目录，并安装其中可发现的技能。

### 5.2 只安装默认 Taste Skill

默认主技能安装名是：

```text
design-taste-frontend
```

安装命令：

```bash
npx skills add https://github.com/Leonxlnx/taste-skill --skill "design-taste-frontend"
```

### 5.3 安装 v1 旧版

当前默认 `taste-skill` 是 v2 experimental。如果依赖旧行为，可安装 v1：

```bash
npx skills add https://github.com/Leonxlnx/taste-skill --skill "design-taste-frontend-v1"
```

### 5.4 手动安装

也可以复制某个技能目录下的 `SKILL.md` 到目标 Agent 支持的技能目录，例如：

```text
.cursor/skills/<skill-name>/SKILL.md
.claude/skills/<skill-name>/SKILL.md
.agents/skills/<skill-name>/SKILL.md
```

不同工具路径可能不同，按所用 Agent 文档为准。

---

## 6. 技能列表

### 6.1 实现类技能

| 技能目录 | 安装名 | 用途 |
|---|---|---|
| `taste-skill` | `design-taste-frontend` | 默认主技能，v2 experimental。读取 brief，推断设计语言，调节 VARIANCE / MOTION / DENSITY 三个旋钮。 |
| `taste-skill-v1` | `design-taste-frontend-v1` | 原始 v1，适合需要旧行为的项目。 |
| `gpt-tasteskill` | `gpt-taste` | 面向 GPT / Codex 更严格的版本，更强布局变化、GSAP 动效方向、反 slop 约束。 |
| `image-to-code-skill` | `image-to-code` | 图片优先流程：先生成/分析参考图，再实现前端。 |
| `redesign-skill` | `redesign-existing-projects` | 用于已有项目重设计：先审计 UI，再改布局、间距、层级和样式。 |
| `soft-skill` | `high-end-visual-design` | 柔和、高端、留白充分、低对比但精致的视觉方向。 |
| `output-skill` | `full-output-enforcement` | 防止模型输出半成品、placeholder、截断代码。 |
| `minimalist-skill` | `minimalist-ui` | 极简、编辑感、Notion/Linear 风格，克制配色和扁平结构。 |
| `brutalist-skill` | `industrial-brutalist-ui` | 工业粗野主义，强对比、瑞士字体、机械感和实验布局。 |
| `stitch-skill` | `stitch-design-taste` | 面向 Google Stitch，可生成 `DESIGN.md` 设计系统规则。 |

### 6.2 图像生成类技能

| 技能目录 | 安装名 | 用途 |
|---|---|---|
| `imagegen-frontend-web` | `imagegen-frontend-web` | 生成网站/landing page/hero/多区块参考图。 |
| `imagegen-frontend-mobile` | `imagegen-frontend-mobile` | 生成移动端屏幕、流程、mockup。 |
| `brandkit` | `brandkit` | 生成品牌套件图：logo 方向、配色、字体、品牌应用等。 |

---

## 7. 默认主技能 taste-skill 的三个旋钮

`taste-skill` 里有三个 1-10 的调节项：

### 7.1 DESIGN_VARIANCE

布局实验程度。

- 低：居中、干净、保守；
- 高：非对称、现代、更有设计感。

### 7.2 MOTION_INTENSITY

动效强度。

- 低：hover、简单 transition；
- 高：滚动动画、磁吸、分层动效、GSAP 方向。

### 7.3 VISUAL_DENSITY

每屏信息密度。

- 低：宽松、留白多；
- 高：数据面板、dashboard、信息密集。

---

## 8. 推荐怎么选技能

### 8.1 不知道用哪个

先用默认：

```bash
npx skills add https://github.com/Leonxlnx/taste-skill --skill "design-taste-frontend"
```

适合大部分新页面、landing page、前端页面优化。

### 8.2 用 Codex / GPT 写前端

可以试：

```bash
npx skills add https://github.com/Leonxlnx/taste-skill --skill "gpt-taste"
```

它更偏严格执行，适合 Codex 这类编码 Agent。

### 8.3 已有项目要美化

用：

```bash
npx skills add https://github.com/Leonxlnx/taste-skill --skill "redesign-existing-projects"
```

提示词可以写：

```text
使用 redesign-existing-projects 技能，先审计当前页面 UI，再分步骤改进布局、间距、字体层级和视觉风格。
```

### 8.4 想先生成设计图再写代码

用：

```bash
npx skills add https://github.com/Leonxlnx/taste-skill --skill "image-to-code"
```

提示词示例：

```text
follow the skill: generate images, then analyze, then code.
我要做一个高端 SaaS 首页，先生成视觉参考图，再分析设计系统，最后实现 React + Tailwind 页面。
```

### 8.5 想要品牌设计图

用：

```bash
npx skills add https://github.com/Leonxlnx/taste-skill --skill "brandkit"
```

---

## 9. 使用教程

### 9.1 基础流程

1. 安装技能；
2. 重启或刷新目标 Agent；
3. 在提示词中明确让 Agent 使用对应 skill；
4. 给足业务上下文、目标用户、风格方向；
5. 要求它先做设计判断，再写代码；
6. 浏览器检查；
7. 继续让 Agent 根据截图/反馈迭代。

### 9.2 示例：生成高端 landing page

```text
使用 design-taste-frontend 技能。
帮我做一个 AI 知识库产品的 landing page，要求：
- 高端、克制，不要模板 SaaS 味
- 不要紫蓝霓虹渐变
- 首页 hero 要有非对称布局
- 字体层级清楚
- 移动端不能横向滚动
- 用 React + Tailwind 实现
先给设计方向，再写代码。
```

### 9.3 示例：已有页面重设计

```text
使用 redesign-existing-projects 技能。
请先审计当前项目的 UI，指出模板感、间距、字体、卡片、动效问题。
然后按最小改动原则重设计，不要破坏业务逻辑。
```

### 9.4 示例：图片到代码

```text
使用 image-to-code 技能。
流程必须是：先生成参考图，再分析参考图里的布局/字体/颜色/间距/组件，最后实现页面。
不要直接开始写代码。
```

### 9.5 示例：极简 UI

```text
使用 minimalist-ui 技能。
做一个极简、编辑感、类似 Linear/Notion 气质的管理后台首页。
要求暖灰底、克制卡片、清晰文字层级，不要重阴影、不要渐变。
```

---

## 10. Taste Skill 的反 AI 味规则

各技能里反复出现的约束包括：

- 不要默认 `Inter` 字体；
- 不要大面积紫蓝霓虹渐变；
- 不要廉价 glow；
- 不要三列等宽卡片套娃；
- 不要无意义 emoji；
- 不要假数据大词，如 `99.99%`、`Next-Gen`、`Seamless`、`Elevate`；
- 不要纯黑 `#000000`；
- 不要过度圆角和 pill；
- 不要移动端横向滚动；
- hero 不要永远居中；
- 动效优先 `transform` 和 `opacity`，避免动画 `top/left/width/height`。

---

## 11. 安全与注意事项

1. **技能是提示词/规则，不是魔法组件库。** 仍需要人审美把关。
2. **安装第三方 skill 前要看源码。** `SKILL.md` 可能包含让 Agent 执行命令、访问外链、修改文件的指令。
3. **不要盲目全装。** 技能太多会增加上下文负担，也可能互相冲突。
4. **v2 是 experimental。** 如果输出不稳定，可以回退 `design-taste-frontend-v1`。
5. **设计规则可能偏主观。** 它偏“高端、反模板”，不一定适合所有企业后台或政企系统。
6. **重设计已有项目时，先备份/开分支。** 防止大范围 UI 改动不好回退。

---

## 12. 常见排障

### 12.1 `npx skills add` 不可用

检查 Node.js / npm：

```bash
node -v
npm -v
```

如果没有 Node.js，需要先安装 Node.js。

### 12.2 安装后 Agent 没反应

可能原因：

- 技能装到了错误目录；
- Agent 没重启；
- 当前工具不支持自动发现 skills；
- 提示词没有触发技能描述；
- 技能名写错。

解决：

```text
请明确使用 design-taste-frontend 技能来完成这个前端页面。
```

### 12.3 输出仍然很模板

加更明确的约束：

```text
不要三列等宽卡片，不要紫蓝渐变，不要模板 SaaS hero，不要默认 Inter 字体。
先列出设计语言和反模式，再写代码。
```

### 12.4 输出代码不完整

安装或使用：

```bash
npx skills add https://github.com/Leonxlnx/taste-skill --skill "full-output-enforcement"
```

提示：

```text
使用 full-output-enforcement 技能。不要省略代码，不要 placeholder，不要“其余同上”。
```

---

## 13. 适合老王的使用场景

- 让 AI 写网页时提升 UI 质感；
- 给现有项目做前端美化；
- 生成网站参考图后交给 Codex / Cursor / Claude Code 实现；
- 做产品首页、后台、移动端页面、品牌视觉稿；
- 把“审美要求”固化成可复用技能，而不是每次重新写一大段提示词。

---

## 14. 推荐落地 SOP

1. 新项目先安装默认 `design-taste-frontend`；
2. 明确业务和风格，不要只说“做得好看”；
3. 要求 Agent 先输出设计方向；
4. 再写代码；
5. 用浏览器截图检查；
6. 根据截图继续让 Agent 修；
7. 如果是旧项目，优先用 `redesign-existing-projects`；
8. 如果 Agent 输出半截，用 `full-output-enforcement`。

---

## 15. 资料来源

- GitHub 仓库：<https://github.com/Leonxlnx/taste-skill>
- 官网：<https://tasteskill.dev>
- README：<https://github.com/Leonxlnx/taste-skill/blob/main/README.md>
- SkillsLLM 项目页：<https://skillsllm.com/skill/taste-skill>
- Skills Over MCP 项目页：<https://skillsovermcp.com/connect/Leonxlnx/taste-skill>

---

## 16. 快速命令备忘

```bash
# 安装全部
npx skills add https://github.com/Leonxlnx/taste-skill

# 默认主技能
npx skills add https://github.com/Leonxlnx/taste-skill --skill "design-taste-frontend"

# v1 旧版
npx skills add https://github.com/Leonxlnx/taste-skill --skill "design-taste-frontend-v1"

# GPT/Codex 严格版
npx skills add https://github.com/Leonxlnx/taste-skill --skill "gpt-taste"

# 图片到代码
npx skills add https://github.com/Leonxlnx/taste-skill --skill "image-to-code"

# 已有项目重设计
npx skills add https://github.com/Leonxlnx/taste-skill --skill "redesign-existing-projects"

# 防止半成品输出
npx skills add https://github.com/Leonxlnx/taste-skill --skill "full-output-enforcement"
```
