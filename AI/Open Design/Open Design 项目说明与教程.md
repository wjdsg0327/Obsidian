# Open Design 项目说明与教程

> 整理时间：2026-06-06  
> 项目：Open Design / open-design / Open Design AI  
> 官方站点：https://open-design.ai/  
> GitHub：https://github.com/nexu-io/open-design  
> License：Apache-2.0

## 1. 一句话说明

Open Design 是一个**本地优先、开源的 Claude Design 替代品**：它把 Claude Code、OpenClaw、Codex、Cursor、OpenCode、Qwen、Copilot 等本地/命令行 Agent 变成“设计引擎”，通过可组合的 **Skill、Design System、Plugin** 来生成 Web 原型、移动端原型、仪表盘、演示文稿、图片、视频和动态图形等设计产物。

它的核心不是“再做一个 AI 聊天工具”，而是把 AI 生成结果落成真实文件：HTML、PDF、PPTX、MP4、Markdown、ZIP 等。

## 2. 核心定位

### 2.1 Claude Design 的开源替代

Open Design 延续了 Claude Design 的“artifact-first / 产物优先”思路：

1. 发现需求：先理解用户要做什么；
2. 锁定方向：选择视觉方向、设计系统、输出类型；
3. 流式生成：Agent 实时写出 artifact；
4. 预览评审：在沙箱 iframe 中即时预览；
5. 交付导出：保存到磁盘或导出为 HTML/PDF/PPTX/MP4 等。

不同点在于：Open Design 开源、本地优先、可自托管，可接入自己的模型、密钥和 Agent，不被单一云服务或模型供应商锁死。

### 2.2 Agent 时代的设计工作台

Open Design 更像“Agent 时代的 Figma 替代品”：不是在画布上手动拖像素，而是用自然语言 + 设计系统 + Skill，直接生成真实 CSS、真实字体、真实组件构成的单页产物。

适合开发者、产品、设计师、独立开发者快速产出：

- Web / SaaS 原型
- 管理后台 / 数据仪表盘
- 移动端页面原型
- 产品 Landing Page
- 文档页 / 博客页
- Pitch Deck / 周报 / 杂志式 PPT
- 图片视觉素材
- 视频 / HyperFrames 动态图形

## 3. 关键能力

### 3.1 Skills：技能决定“做什么”

Skill 是 Open Design 的工作流说明文件，通常由 `SKILL.md` 描述。它告诉 Agent 当前任务的输出规范、流程、约束和质量要求。

常见内置 Skill：

- `web-prototype`：通用 Web 原型；
- `saas-landing`：SaaS 落地页；
- `dashboard`：后台 / 分析仪表盘；
- `pricing-page`：定价页；
- `docs-page`：三栏文档页；
- `blog-post`：编辑风长文；
- `mobile-app`：移动端单屏原型；
- `simple-deck`：简单横向翻页 Deck；
- `magazine-web-ppt` / `guizang-ppt`：杂志式 PPT / Deck。

### 3.2 Design System：设计系统决定“长什么样”

Design System 通常由 `DESIGN.md` 描述，是品牌视觉契约，包括色板、字体、间距、布局、动效、语气和反模式等。

Open Design 官方资料提到内置大量设计系统，例如 Neutral Modern、Warm Editorial，以及 Linear、Stripe、Vercel、Airbnb、Apple、Tesla、Notion、Anthropic、Cursor、Supabase、Figma 等风格系统。

实际使用时，可以把 Skill 和 Design System 组合：

> `dashboard` + `Linear 风格` → Linear 风的管理后台  
> `simple-deck` + `Warm Editorial` → 暖色编辑风 PPT  
> `mobile-app` + `Apple 风格` → Apple 风移动端原型

### 3.3 Plugins：插件扩展工作流

Plugin 用于扩展生成流程和能力，例如接入外部系统、模板、媒体生成或自动化流程。官方 README 提到 Open Design 已包含大量开箱即用插件。

### 3.4 支持多种 Agent / 模型

Open Design 可以运行在多种本地 Agent 或编码 CLI 上，例如：

- Claude Code
- OpenClaw
- Codex CLI
- Cursor
- VS Code + GitHub Copilot / Copilot CLI
- Gemini CLI
- OpenCode
- Qwen
- Kimi CLI
- Hermes Agent
- Antigravity
- Cline / Trae 等

如果没有本地 CLI，也可以使用 BYOK API 模式，接入 OpenAI、Anthropic、Azure OpenAI、Google Gemini、Ollama、LM Studio、vLLM 或其他 OpenAI-compatible endpoint。

## 4. 输出物与使用场景

| 输出类型 | 典型用途 | 可导出形式 |
|---|---|---|
| Web 原型 | SaaS 页面、官网、活动页、管理页 | HTML / ZIP |
| 移动端原型 | App onboarding、单页流程、交互草图 | HTML / ZIP |
| Dashboard | KPI 大屏、运营看板、GitHub 风格报表 | HTML |
| Deck / PPT | 路演、周报、产品方案、杂志式展示 | HTML / PDF / PPTX / Markdown |
| 图片 | 品牌图、封面图、海报视觉 | 图片文件 |
| 视频 / HyperFrames | 产品宣传片、动态图形、短视频素材 | MP4 |

## 5. 本地开发模式教程

### 5.1 环境要求

官方 Quickstart 给出的主要要求：

- Node.js：`~24`，即 Node 24.x；
- pnpm：`10.33.x`，仓库通过 `packageManager` 固定版本；
- 操作系统：macOS、Linux、WSL2 是主要路径；Windows 原生也支持，但 WSL2 通常更稳定；
- 可选：安装一个本地 Agent CLI，例如 Claude Code、Codex、Gemini、OpenCode、Cursor Agent、Qwen、Copilot CLI 等。

### 5.2 安装 Node 24（可选）

如果用 nvm：

```bash
nvm install 24
nvm use 24
```

如果用 fnm：

```bash
fnm install 24
fnm use 24
```

启用 Corepack，并确认 pnpm 版本：

```bash
corepack enable
corepack pnpm --version
```

期望输出接近：

```text
10.33.2
```

### 5.3 克隆与启动

```bash
git clone https://github.com/nexu-io/open-design.git
cd open-design
corepack enable
pnpm install
pnpm tools-dev run web
```

启动后打开终端输出的 Web URL。

首次加载时，Open Design 会扫描本机已安装的 Agent CLI，并自动选择可用项；默认 Skill 一般是 `web-prototype`，默认 Design System 一般是 `Neutral Modern`。

### 5.4 第一次生成

1. 打开 Open Design Web 页面；
2. 在顶部选择 Skill，例如 `web-prototype` / `dashboard` / `mobile-app`；
3. 选择 Design System，例如 `Neutral Modern`；
4. 在输入框描述需求，例如：

```text
为一个 AI 知识库产品生成一个 SaaS Landing Page，包含 Hero、功能区、价格区、FAQ 和 CTA，风格偏 Linear。
```

5. 点击 Send；
6. 左侧会显示 Agent 流式输出；
7. 右侧会解析 `<artifact>` 并实时渲染 HTML；
8. 完成后点击 Save to disk，产物会保存到：

```text
./.od/artifacts/<timestamp>-<slug>/index.html
```

## 6. Docker 部署教程

Docker 适合不想在本机安装 Node.js / pnpm 的场景。

### 6.1 要求

- Docker Desktop
- Docker Compose v2

验证：

```bash
docker compose version
```

### 6.2 启动步骤

从仓库根目录开始：

```bash
cd deploy
cp .env.example .env
```

生成安全 token：

```bash
openssl rand -hex 32
```

编辑 `.env`，把生成的 token 填入：

```env
OD_API_TOKEN=你的随机token
```

启动：

```bash
docker compose up -d
```

浏览器打开：

```text
http://localhost:7456
```

### 6.3 常用 Docker 命令

查看日志：

```bash
docker compose logs -f
```

重启：

```bash
docker compose restart
```

停止：

```bash
docker compose down
```

拉取最新镜像并重启：

```bash
docker compose pull
docker compose up -d
```

删除所有本地应用数据：

```bash
docker compose down -v
```

注意：最后一条会删除 Docker 卷中的本地数据，谨慎执行。

## 7. 目录结构速览

官方 Quickstart 中的项目结构大意：

```text
open-design/
├── apps/
│   ├── daemon/      # Node/Express 后端，负责启动本地 Agent、提供 API
│   ├── web/         # Next.js + React 前端
│   └── desktop/     # Electron 桌面端
├── packages/        # 共享协议、平台、sidecar 等包
├── tools/dev/       # pnpm tools-dev 生命周期管理
├── e2e/             # Playwright / Vitest 测试
├── skills/          # 各类 SKILL.md
├── design-systems/  # 各类 DESIGN.md
├── docs/            # 架构、协议、模式、路线图等文档
├── .od/             # 运行时数据，自动创建，通常 gitignored
│   ├── app.sqlite
│   ├── artifacts/
│   └── projects/<id>/
└── package.json
```

## 8. 两种执行模式

### 8.1 Local CLI 模式

当 daemon 检测到本地 Agent CLI 时，默认走 Local CLI：

```text
前端 → daemon /api/chat → spawn(<agent>) → stdout → SSE → artifact parser → iframe 预览
```

适合已经安装 Claude Code、OpenClaw、Codex、Cursor Agent 等 CLI 的用户。

### 8.2 API 模式

没有本地 CLI，或想直接使用模型 API 时，可使用 API 模式：

```text
前端 → daemon /api/proxy/{provider}/stream → provider SSE → artifact parser → iframe 预览
```

支持 OpenAI / Anthropic / Azure OpenAI / Google Gemini / Ollama / vLLM 等。

## 9. Prompt 组合机制

每次发送需求时，Open Design 会组合三层提示：

```text
BASE_SYSTEM_PROMPT
+ 当前 Design System 正文（DESIGN.md）
+ 当前 Skill 正文（SKILL.md）
```

这意味着：

- Skill 决定流程和输出类型；
- Design System 决定品牌视觉和风格约束；
- 用户 prompt 决定具体业务需求。

所以写 prompt 时不必重复所有风格细节，优先说明业务目标、受众、页面内容、关键模块和限制条件。

## 10. 常见排障

### 10.1 no agents found on PATH

含义：Open Design 没找到本地 Agent CLI。

处理：

- 安装 `claude`、`codex`、`gemini`、`opencode`、`cursor-agent`、`qwen`、`copilot` 等其中之一；
- 或进入 Settings，切换到 API mode，填入 provider key；
- 如果已经安装，确认 CLI 所在目录在启动 daemon 的 PATH 中。

### 10.2 artifact 始终不渲染

可能原因：模型输出了普通文本，但没有按要求用 `<artifact>` 包裹。

处理：

- 换更强模型；
- 换更严格的 Skill；
- 检查 daemon 日志，确认 system prompt 正常传递。

### 10.3 daemon /api/chat 500

通常是本地 CLI 拒绝参数或认证失败。

处理：

- 查看 daemon 终端 stderr；
- 确认 CLI 可直接运行；
- 若是 Claude Code，先检查 `claude --version` 与认证状态；
- 不同 CLI 参数格式不同，必要时查看 `apps/daemon/src/agents.ts`。

### 10.4 媒体生成报 OD_BIN 缺失或 daemon URL 为 :0

处理：

```bash
pnpm --filter @open-design/daemon build
pnpm tools-dev restart --daemon-port 7457 --web-port 5175
ls -la apps/daemon/dist/cli.js
curl -s http://127.0.0.1:7457/api/health
```

然后从 Open Design 应用里重新打开项目，不要复用旧 terminal Agent 会话。

### 10.5 nginx 反代 SSE 断流

如果前面加 nginx，SSE 路由要关闭 buffering 和 gzip，否则可能出现 `net::ERR_INCOMPLETE_CHUNKED_ENCODING 200 (OK)`。

参考配置：

```nginx
location /api/ {
    proxy_pass http://127.0.0.1:7456;
    proxy_buffering off;
    gzip off;
    proxy_read_timeout 86400s;
    proxy_send_timeout 86400s;
    proxy_http_version 1.1;
    proxy_set_header Connection "";
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
}
```

## 11. 推荐学习路线

### 入门路线

1. 先用桌面版或 Docker 跑起来；
2. 用默认 `web-prototype + Neutral Modern` 生成一个简单页面；
3. 尝试切换 Design System，看同一个 prompt 的视觉差异；
4. 尝试 `dashboard`、`mobile-app`、`simple-deck` 等不同 Skill；
5. 保存产物到磁盘，并用 Cursor / OpenClaw / Claude Code 继续改。

### 进阶路线

1. 阅读 `skills/` 下的 `SKILL.md`，理解 Skill 如何约束 Agent；
2. 阅读 `design-systems/` 下的 `DESIGN.md`，学习设计系统 schema；
3. 自己创建一个品牌 `DESIGN.md`；
4. 自己创建一个专用 Skill，例如“知识库教程页生成器”；
5. 尝试通过 MCP 或 `od mcp install <agent>` 接入常用编码 Agent。

### 工程路线

1. 阅读 `docs/architecture.md`；
2. 阅读 `docs/skills-protocol.md`；
3. 阅读 `docs/agent-adapters.md`；
4. 查看 `apps/daemon/src/agents.ts`，理解 Agent 适配；
5. 查看 `apps/web/src/artifacts/`，理解 `<artifact>` 流式解析和预览。

## 12. 适合老王的实用场景

结合日常“处理文件、查资料、整理教程”的工作，Open Design 可重点用于：

- 把技术笔记快速变成可视化教程页；
- 把项目说明生成 Landing Page / README 风格展示；
- 把学习资料生成 PPT / Deck；
- 给 Docker、AI、Python、安卓等主题生成知识库封面图或教程页面；
- 生成用于 Obsidian 知识库的可视化 HTML 附件；
- 让 OpenClaw / Codex / Cursor 接着把设计稿转成真实前端项目。

## 13. 快速参考命令

```bash
# 克隆
git clone https://github.com/nexu-io/open-design.git
cd open-design

# 本地开发启动
corepack enable
pnpm install
pnpm tools-dev run web

# 后台启动 daemon + web + desktop
pnpm tools-dev

# 状态 / 日志 / 检查
pnpm tools-dev status
pnpm tools-dev logs
pnpm tools-dev check

# 停止
pnpm tools-dev stop

# Docker 启动
cd deploy
cp .env.example .env
openssl rand -hex 32
# 填写 OD_API_TOKEN 后：
docker compose up -d

# Docker 访问
# http://localhost:7456
```

## 14. 资料来源

- 官方网站：https://open-design.ai/
- 官方确认页：https://open-design.ai/official/
- GitHub 仓库：https://github.com/nexu-io/open-design
- README：https://github.com/nexu-io/open-design/blob/main/README.md
- Quickstart：https://github.com/nexu-io/open-design/blob/main/QUICKSTART.md
- 简体中文 README：https://github.com/nexu-io/open-design/blob/main/docs/i18n/README.zh-CN.md
- 简体中文 Quickstart：https://github.com/nexu-io/open-design/blob/main/docs/i18n/QUICKSTART.zh-CN.md
