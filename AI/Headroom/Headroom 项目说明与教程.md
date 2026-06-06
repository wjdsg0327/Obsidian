# Headroom 项目说明与教程

> 整理日期：2026-06-06  
> 项目仓库：<https://github.com/chopratejas/headroom>  
> 官方文档：<https://headroom-docs.vercel.app/docs>  
> 许可证：Apache-2.0  
> 定位：AI Agent / LLM 应用的本地优先上下文压缩层。

---

## 1. 一句话说明

**Headroom 是一个给 AI Agent 减少上下文 token 消耗的压缩层**：它会在日志、工具输出、代码搜索结果、RAG 片段、文件内容、对话历史等内容发送给大模型之前，先进行压缩和结构化处理，从而在尽量不影响答案质量的前提下减少 token、降低成本、提高长任务可持续性。

可以把它理解成：

> Agent 和 LLM 之间的“上下文瘦身器 + 可逆缓存层 + 代理服务器”。

---

## 2. 它解决什么问题

AI 编程助手和 Agent 在真实任务里，经常会把大量原始文本塞进上下文：

- `grep` / `ripgrep` 搜索结果；
- 构建、测试、安装日志；
- 报错堆栈；
- 大文件片段；
- GitHub issue / PR 讨论；
- RAG 检索出来的大段资料；
- 多轮对话历史；
- 多 Agent 之间传递的中间结果。

这些内容常见问题是：

1. **token 消耗大**：同样一次请求成本更高。
2. **上下文窗口被挤满**：真正重要的问题、约束、代码反而被淹没。
3. **重复信息多**：日志、搜索结果里经常有大量重复行。
4. **Agent 长任务容易劣化**：上下文越来越长，模型越容易遗漏重点。
5. **多工具、多 Agent 难共享压缩后的上下文**。

Headroom 的思路是：

> 在内容进入 LLM 之前，先按内容类型选择合适的压缩方式；必要时保存原文，让模型可以按需取回。

---

## 3. 核心能力

### 3.1 多种接入方式

Headroom 支持几类用法：

| 用法 | 适合场景 | 说明 |
|---|---|---|
| Library / SDK | 自己写 Python / TypeScript 应用 | 在代码里调用 `compress(messages)` |
| Proxy | 想尽量零改代码 | 启动本地代理，让 OpenAI/Anthropic 兼容请求先经过 Headroom |
| Agent wrap | 直接包装 AI 编程工具 | 如 `headroom wrap claude`、`headroom wrap codex` |
| MCP server | MCP 客户端 | 提供 `headroom_compress`、`headroom_retrieve`、`headroom_stats` 等工具 |
| Memory / SharedContext | 多 Agent / 长会话 | 跨会话、跨 Agent 共享压缩上下文和记忆 |
| headroom learn | 失败经验沉淀 | 从失败会话中提取规则，写入 `CLAUDE.md` / `AGENTS.md` / `GEMINI.md` 等 |

### 3.2 支持的内容类型

官方 README 表示它面向这些内容做压缩：

- tool outputs：工具输出；
- logs：日志；
- files：文件内容；
- RAG chunks：检索增强生成片段；
- code search results：代码搜索结果；
- conversation history：对话历史；
- JSON / 结构化数据；
- 代码 AST；
- 普通文本。

### 3.3 本地优先与可逆压缩

Headroom 的重要特点：

- **本地运行**：核心压缩流程在本机，隐私和可控性更好。
- **可逆压缩 CCR**：原文不会直接丢掉，而是保存在本地；模型需要细节时，可以通过检索工具取回原文。
- **跨 Agent 记忆**：Claude、Codex、Gemini 等不同 Agent 可以共享记忆/上下文。
- **缓存优化**：通过 CacheAligner 稳定提示词前缀，提高供应商 KV cache 命中率。

---

## 4. 核心机制简图

```text
AI Agent / 应用
  │
  │  prompts / tool outputs / logs / RAG / files / history
  ▼
Headroom 本地压缩层
  │
  ├─ CacheAligner：稳定前缀，改善缓存命中
  ├─ ContentRouter：识别内容类型，选择压缩器
  ├─ SmartCrusher：压缩 JSON / 结构化数据
  ├─ CodeCompressor：基于 AST 压缩代码
  ├─ Kompress-base：压缩普通文本
  ├─ IntelligentContext：按重要性裁剪上下文
  └─ CCR：保存原文，可按需 retrieve
  │
  ▼
LLM Provider
  Anthropic / OpenAI / Bedrock / Azure / Vertex / OpenRouter ...
```

---

## 5. 什么时候适合用 Headroom

### 适合

- 经常使用 Claude Code、Codex、Cursor、Aider、Copilot CLI、OpenClaw 等 AI 编程工具；
- 经常跑长任务，日志和搜索结果很多；
- 想降低 token 成本；
- 想减少上下文爆炸导致的模型遗忘；
- 想让多个 Agent 共享记忆和上下文；
- 想把 RAG 检索结果先压缩再交给模型；
- 想以代理方式给现有 OpenAI/Anthropic 兼容应用加一层压缩。

### 不适合 / 暂时没必要

- 只是偶尔聊天，token 成本不敏感；
- 只用单一模型自带的上下文压缩，不需要跨 Agent 记忆；
- 当前环境不允许启动本地服务；
- 对压缩引入的任何语义偏差都不能接受的高风险场景；
- 不想让工具修改已有 Agent 配置文件时，不建议直接用 wrap，应先 dry-run 或手动检查。

---

## 6. 安装教程

### 6.1 Python 安装

Headroom 要求 **Python 3.10+**。

最小安装：

```bash
pip install headroom-ai
```

安装全部功能：

```bash
pip install "headroom-ai[all]"
```

按需安装 extras：

```bash
pip install "headroom-ai[proxy]"
pip install "headroom-ai[mcp]"
pip install "headroom-ai[proxy,langchain,ml]"
```

常见 extras：

| extra | 说明 |
|---|---|
| `proxy` | 本地代理服务 |
| `mcp` | MCP 工具服务 |
| `ml` | 机器学习压缩相关能力，如 Kompress-base |
| `code` | 代码压缩相关能力 |
| `memory` | 记忆功能 |
| `langchain` | LangChain 集成 |
| `agno` | Agno 集成 |
| `evals` | 评测相关工具 |

验证安装：

```bash
python -c "import headroom; print(headroom.__version__)"
```

### 6.2 pipx 安装

如果用 `pipx`，官方建议显式选择支持的 Python 版本，例如：

```bash
pipx install --python python3.13 "headroom-ai[all]"
```

### 6.3 Node / TypeScript 安装

```bash
npm install headroom-ai
```

之后可在 TypeScript / Node 项目里直接调用 SDK。

### 6.4 Docker 安装

拉取镜像：

```bash
docker pull ghcr.io/chopratejas/headroom:latest
```

启动代理服务：

```bash
docker run -p 8787:8787 ghcr.io/chopratejas/headroom:latest
```

访问健康检查：

```bash
curl http://localhost:8787/health
```

---

## 7. 快速开始：Library 模式

Library 模式适合自己写应用时直接调用压缩函数。

### 7.1 Python 示例

```python
from headroom import compress

messages = [
    {
        "role": "user",
        "content": "这里放很长的日志、搜索结果、RAG 文档或代码片段……"
    }
]

compressed = compress(messages)
print(compressed)
```

之后把 `compressed` 发送给 OpenAI / Anthropic / 其他 LLM 即可。

### 7.2 TypeScript 示例

```ts
import { compress } from 'headroom-ai';

const messages = [
  {
    role: 'user',
    content: '这里放很长的日志、搜索结果、RAG 文档或代码片段……'
  }
];

const compressed = await compress(messages, { model: 'gpt-4o-mini' });
console.log(compressed);
```

### 7.3 适用场景

- 自己开发 AI 应用；
- 自己控制 LLM 请求流程；
- 想在发送请求前做一次显式压缩；
- RAG 系统里想压缩 chunks；
- 批处理日志摘要、issue 摘要、代码搜索结果。

---

## 8. 快速开始：Proxy 模式

Proxy 模式适合“尽量不改代码”：启动本地代理，让请求先经过 Headroom，再转发给真实 LLM 服务。

### 8.1 启动代理

```bash
headroom proxy --port 8787
```

默认监听本地端口 `8787`。

指定 host 和 port：

```bash
headroom proxy --host 0.0.0.0 --port 8080
```

### 8.2 健康检查

```bash
curl http://localhost:8787/health
```

### 8.3 查看统计

```bash
curl http://localhost:8787/stats
```

可能看到类似统计：

```json
{
  "requests_total": 42,
  "tokens_saved_total": 125000
}
```

### 8.4 代理支持的常见接口

官方文档列出代理支持这些接口：

- `GET /health`
- `GET /stats`
- `GET /stats-history`
- `GET /metrics`
- `POST /v1/messages`
- `POST /v1/chat/completions`
- `POST /v1/responses`
- `POST /v1internal:streamGenerateContent`
- `POST /v1/compress`

### 8.5 代理模式常用参数

```bash
# 最大化 token 压缩
headroom proxy --mode token

# 优先保持 provider prefix cache 稳定
headroom proxy --mode cache

# 禁用优化，变成透传模式
headroom proxy --no-optimize

# 禁用语义缓存
headroom proxy --no-cache

# 启用记忆
headroom proxy --memory

# 启用失败学习
headroom proxy --learn --min-evidence 3
```

---

## 9. 快速开始：Agent wrap 模式

wrap 模式适合直接包装现有 AI 编程工具。

官方 README 中列出的兼容情况包括：

| Agent | wrap 支持 | 说明 |
|---|---:|---|
| Claude Code | ✅ | 支持 memory、code graph 等选项 |
| Codex | ✅ | 可与 Claude 共享 memory |
| Cursor | ✅ | 通常会打印配置，按提示粘贴一次 |
| Aider | ✅ | 启动 proxy 并启动 Aider |
| Copilot CLI | ✅ | 启动 proxy 并启动 Copilot CLI |
| OpenClaw | ✅ | 可作为 ContextEngine plugin 安装 |

### 9.1 Claude Code

```bash
headroom wrap claude
```

可选：

```bash
headroom wrap claude --memory
```

### 9.2 Codex

```bash
headroom wrap codex
```

只准备配置、不立即启动：

```bash
headroom wrap codex --prepare-only
```

### 9.3 Cursor

```bash
headroom wrap cursor
```

Cursor 场景下通常需要按输出提示手动设置 base URL / API 配置。

### 9.4 Aider

```bash
headroom wrap aider
```

### 9.5 Copilot CLI

普通模式：

```bash
headroom wrap copilot
```

订阅模式示例：

```bash
headroom wrap copilot --subscription -- --model gpt-4o
```

### 9.6 使用 wrap 前的注意事项

wrap 类命令可能会：

- 启动本地 proxy；
- 修改或生成目标工具的配置；
- 注入环境变量；
- 安装 MCP / 插件；
- 写入 `~/.headroom` 或目标 Agent 的配置目录。

建议第一次使用时：

1. 在测试项目里试；
2. 先备份相关配置文件；
3. 仔细看命令输出；
4. 如果支持 `--prepare-only` / dry-run，优先使用；
5. 不要在生产环境里直接首次试验。

---

## 10. MCP 模式教程

MCP 模式适合让支持 MCP 的客户端调用 Headroom 工具。

### 10.1 安装 MCP 功能

```bash
pip install "headroom-ai[mcp]"
```

或者同时安装代理：

```bash
pip install "headroom-ai[proxy]"
```

### 10.2 注册到 Claude Code

```bash
headroom mcp install
```

之后启动 Claude Code，它应能看到 Headroom 的 MCP 工具。

### 10.3 使用自定义代理地址

```bash
headroom mcp install --proxy-url http://host:9000
```

### 10.4 常用 MCP 命令

```bash
# 安装 / 注册
headroom mcp install

# 覆盖已有配置
headroom mcp install --force

# 查看状态
headroom mcp status

# 卸载
headroom mcp uninstall

# 调试模式启动 MCP server
headroom mcp serve --debug
```

### 10.5 MCP 工具

| 工具 | 作用 |
|---|---|
| `headroom_compress` | 压缩输入内容 |
| `headroom_retrieve` | 从 CCR 存储里取回原文 |
| `headroom_stats` | 查看压缩统计 |

---

## 11. headroom learn：失败经验学习

`headroom learn` 用来分析失败会话，把经验写入项目规则文件，帮助 Agent 下次少犯同类错误。

### 11.1 快速使用

只查看建议，不修改文件：

```bash
headroom learn
```

应用建议，写入规则文件：

```bash
headroom learn --apply
```

分析指定项目：

```bash
headroom learn --project ~/my-project --apply
```

分析所有项目：

```bash
headroom learn --all --apply
```

### 11.2 它会学什么

官方文档示例包括：

- 环境事实：例如项目使用 pnpm 而不是 npm；
- 文件路径纠正：例如真实配置文件位置；
- 搜索范围：哪些目录应该排除；
- 命令模式：测试、构建、格式化的正确命令；
- 大文件提醒：避免 Agent 反复读取超大文件。

### 11.3 写到哪里

可能写入：

- `CLAUDE.md`
- `AGENTS.md`
- `GEMINI.md`
- `MEMORY.md`

一般会通过 marker 区块追加/更新，例如：

```markdown
## Headroom Learned Patterns
...
```

**注意**：这类命令会修改项目规则文件，建议先用 dry-run 看建议，再决定是否 `--apply`。

---

## 12. Memory / Shared Context

Headroom 支持记忆和共享上下文，用于跨会话、跨 Agent 保存重要事实。

### 12.1 Python 记忆示例

```python
from headroom import with_memory

# 用 with_memory 包装 LLM client 后，用户偏好、项目事实等可被抽取和复用
client = with_memory(your_llm_client, user_id="laowang")
```

### 12.2 记忆适合保存什么

- 用户偏好；
- 项目技术栈；
- 反复出现的路径；
- 常用命令；
- 代码库结构；
- 已验证过的排障结论。

### 12.3 后端选择

官方文档提到可使用：

- ONNX embeddings：推荐，本地、快速、免费；
- OpenAI embeddings：质量高但需要付费和外部 API；
- Ollama embeddings：本地模型服务；
- 存储后端可按配置选择。

---

## 13. 配置要点

### 13.1 CLI Context Tool

默认直接运行：

```bash
headroom wrap claude
```

只准备 Codex 配置：

```bash
headroom wrap codex --prepare-only
```

也可以通过环境变量选择 CLI context tool，例如官方 README 提到可选 `lean-ctx`：

```bash
export HEADROOM_CONTEXT_TOOL=lean-ctx
headroom wrap claude
```

### 13.2 Proxy 配置

```bash
headroom proxy \
  --host 0.0.0.0 \
  --port 8787 \
  --mode token
```

常用开关：

```bash
headroom proxy --no-optimize
headroom proxy --no-cache
headroom proxy --mode cache
headroom proxy --memory
headroom proxy --learn --min-evidence 3
```

### 13.3 常见 provider 环境变量

根据不同后端，通常需要这些密钥之一：

```bash
export OPENAI_API_KEY="..."
export ANTHROPIC_API_KEY="..."
export AWS_REGION="us-east-1"
export AZURE_OPENAI_API_KEY="..."
```

具体以所用 provider 和官方配置文档为准。

### 13.4 配置优先级

一般建议遵循：

1. 命令行参数优先；
2. 环境变量其次；
3. 项目配置文件再次；
4. 默认配置最后。

---

## 14. 与其他方案对比

官方 README 中的对比大意：

| 方案 | 覆盖范围 | 部署方式 | 本地 | 可逆 |
|---|---|---|---:|---:|
| Headroom | tools / RAG / logs / files / history | Proxy / library / middleware / MCP | ✅ | ✅ |
| RTK | CLI 命令输出 | CLI wrapper | ✅ | ❌ |
| lean-ctx | CLI commands / MCP tools / editor rules | CLI wrapper / MCP | ✅ | ❌ |
| Compresr / Token Co. | 发送给 API 的文本 | Hosted API | ❌ | ❌ |
| OpenAI Compaction | 对话历史 | Provider 原生 | ❌ | ❌ |

Headroom 也说明它会集成/使用 RTK 这类工具来优化 shell 输出，但 Headroom 自身目标更广：工具输出、RAG、日志、文件、历史等都覆盖。

---

## 15. 实操路线：老王可以怎么试

### 路线 A：最小验证，不动现有 Agent 配置

适合第一次尝试。

```bash
python3 --version
pip install "headroom-ai[all]"
python -c "import headroom; print(headroom.__version__)"
headroom perf
```

如果安装成功，再测试 library 压缩或 proxy。

### 路线 B：本地代理验证

```bash
pip install "headroom-ai[proxy]"
headroom proxy --port 8787
```

另开终端：

```bash
curl http://localhost:8787/health
curl http://localhost:8787/stats
```

确认代理服务正常后，再把某个测试客户端的 base URL 指向 `http://localhost:8787`。

### 路线 C：包装 Claude Code / Codex

先备份配置，再试：

```bash
headroom wrap claude
```

或：

```bash
headroom wrap codex --prepare-only
```

建议在临时项目目录里先跑，不要直接在重要项目中首次启用。

### 路线 D：接入自研 RAG / Agent

如果自己写 Python / TypeScript 应用，推荐优先 library 模式：

1. RAG 检索得到 chunks；
2. 用 Headroom 压缩 chunks；
3. 把压缩后的内容发送给 LLM；
4. 记录压缩前后 token；
5. 对比回答质量。

---

## 16. 常见排障

### 16.1 `headroom` 命令不存在

检查是否安装在当前 Python 环境：

```bash
python -m pip show headroom-ai
python -m pip install "headroom-ai[all]"
```

如果使用虚拟环境，确认已激活：

```bash
source .venv/bin/activate
```

### 16.2 Python 版本不支持

Headroom 要求 Python 3.10+：

```bash
python --version
```

如果系统默认 Python 太旧，用 pyenv、conda 或 pipx 指定版本。

### 16.3 代理端口冲突

换端口：

```bash
headroom proxy --port 8877
```

检查端口占用：

```bash
lsof -i :8787
```

### 16.4 LLM 请求失败

检查：

- API key 是否设置；
- base URL 是否指向 Headroom proxy；
- proxy 是否在运行；
- provider 是否兼容 OpenAI / Anthropic 接口；
- 是否需要设置 `OPENAI_BASE_URL` / `ANTHROPIC_BASE_URL` 或客户端对应配置。

### 16.5 wrap 后 Agent 表现异常

建议：

1. 停止当前 Agent；
2. 查看 Headroom 输出的配置改动；
3. 暂时用 `--no-optimize` 验证是否是压缩导致；
4. 回退 Agent 配置；
5. 改用 proxy / library 模式做更可控的测试。

### 16.6 压缩后答案质量下降

可尝试：

```bash
headroom proxy --mode cache
```

或降低压缩激进程度，必要时对关键内容不压缩。

---

## 17. 风险与注意事项

1. **项目迭代很快**：命令参数可能变化，实操前应以官方 README / docs 为准。
2. **wrap 可能改配置**：首次使用前备份 Claude/Codex/Cursor 等配置。
3. **压缩不是无损理解**：虽然 CCR 可取回原文，但模型未必总能主动取回关键细节。
4. **本地服务也要管权限**：如果 `--host 0.0.0.0` 对外监听，要注意防火墙和访问控制。
5. **密钥安全**：不要把 API key 写入仓库；尽量用环境变量或安全密钥管理。
6. **桌面/GUI 版本需谨慎**：如果安装 GUI 或第三方封装，先确认来源、权限、会修改哪些文件。

---

## 18. 推荐学习顺序

1. 看 README，理解 Headroom 是什么；
2. 用 `pip install "headroom-ai[all]"` 安装；
3. 跑 `headroom perf` 看压缩统计；
4. 用 library 模式压缩一段长日志；
5. 启动 proxy，看 `/health` 和 `/stats`；
6. 在临时项目里测试 `headroom wrap claude` 或 `headroom wrap codex --prepare-only`；
7. 研究 MCP 模式；
8. 最后再考虑 Memory / headroom learn。

---

## 19. 适合老王的实用场景

### 19.1 AI 编程助手省 token

如果老王经常让 Agent 读代码库、跑测试、看日志，Headroom 可以减少大量无效上下文。

### 19.2 整理大日志

例如 Docker、系统服务、Python 报错、CI 输出，可以先压缩再交给模型分析。

### 19.3 RAG 知识库压缩

检索出 20 段资料后，先压缩再喂给模型，节省上下文窗口。

### 19.4 多 Agent 协作

如果同时用 OpenClaw、Claude Code、Codex 等工具，可以研究它的 shared memory / shared context 能力。

### 19.5 项目规则沉淀

`headroom learn` 可以把失败经验写入 `AGENTS.md` 这类规则文件，适合长期维护项目。

---

## 20. 资料来源

- GitHub 仓库：<https://github.com/chopratejas/headroom>
- README：<https://raw.githubusercontent.com/chopratejas/headroom/main/README.md>
- 官方文档首页：<https://headroom-docs.vercel.app/docs>
- Quickstart：<https://headroom-docs.vercel.app/docs/quickstart>
- Installation：<https://headroom-docs.vercel.app/docs/installation>
- Proxy：<https://headroom-docs.vercel.app/docs/proxy>
- MCP：<https://headroom-docs.vercel.app/docs/mcp>
- Configuration：<https://headroom-docs.vercel.app/docs/configuration>
- Architecture：<https://headroom-docs.vercel.app/docs/architecture>
- How compression works：<https://headroom-docs.vercel.app/docs/how-compression-works>
- CCR reversible compression：<https://headroom-docs.vercel.app/docs/ccr>
- Memory：<https://headroom-docs.vercel.app/docs/memory>
- Failure learning：<https://headroom-docs.vercel.app/docs/failure-learning>
- Docker install：<https://headroom-docs.vercel.app/docs/docker-install>
- Troubleshooting：<https://headroom-docs.vercel.app/docs/troubleshooting>

---

## 21. 快速命令备忘

```bash
# 安装
pip install "headroom-ai[all]"
npm install headroom-ai

# 验证
python -c "import headroom; print(headroom.__version__)"
headroom perf

# Proxy
headroom proxy --port 8787
curl http://localhost:8787/health
curl http://localhost:8787/stats

# Agent wrap
headroom wrap claude
headroom wrap codex
headroom wrap cursor
headroom wrap aider
headroom wrap copilot

# MCP
headroom mcp install
headroom mcp status
headroom mcp uninstall

# Learn
headroom learn
headroom learn --apply
```
