---
title: Grok Build 运行与环境整理
date: 2026-07-17
tags:
  - AI
  - AI编程
  - Grok
  - 开源工具
  - 本地部署
---

# Grok Build 运行与环境整理

## 1. Grok Build 是什么

Grok Build 是 xAI 开源的一个 **终端 AI 编程助手 / TUI 工具**。

它更像下面这些工具：

- Claude Code
- OpenCode
- Codex CLI
- Cursor Agent 的终端版

它可以做：

- 读取代码库上下文
- 修改文件
- 执行 shell 命令
- 搜索网页
- 跑长任务
- 交互式 TUI 操作
- Headless 模式，用于脚本或 CI
- 支持 MCP、skills、plugins、hooks、sandbox 等扩展

> 注意：开源的是 **Grok Build 工具框架**，不是最新版 Grok 大模型权重。  
> 工具可以本地运行，但 AI 推理通常还需要接云端 API 或本地模型。

---

## 2. 官方资料

- GitHub：<https://github.com/xai-org/grok-build>
- 官方 CLI 页面：<https://x.ai/cli>
- 文档入口：<https://docs.x.ai/build/overview>
- 开源说明：<https://x.ai/news/grok-build-open-source>

---

## 3. 安装方式

### 3.1 macOS / Linux / Git Bash

```bash
curl -fsSL https://x.ai/cli/install.sh | bash
grok --version
```

安装后一般直接运行：

```bash
grok
```

第一次启动会打开浏览器进行认证。

### 3.2 Windows PowerShell

```powershell
irm https://x.ai/cli/install.ps1 | iex
grok --version
```

---

## 4. 从源码编译

源码仓库：

```bash
git clone https://github.com/xai-org/grok-build.git
cd grok-build
```

### 4.1 环境要求

需要：

1. **Rust**
   - 仓库里有 `rust-toolchain.toml`
   - 使用 `rustup` 时会自动安装指定版本工具链

2. **protoc**
   - 用于 protobuf 代码生成
   - 仓库会优先找 `bin/protoc`
   - 或使用系统 PATH / `$PROTOC` 中的 protoc

3. **系统支持**
   - macOS：支持
   - Linux：支持
   - Windows：有预编译版；源码树构建属于 best-effort，官方未重点测试

### 4.2 源码运行

```bash
cargo run -p xai-grok-pager-bin
```

### 4.3 构建 Release

```bash
cargo build -p xai-grok-pager-bin --release
```

生成的二进制文件：

```bash
target/release/xai-grok-pager
```

官方安装版会把它包装成：

```bash
grok
```

### 4.4 快速检查

```bash
cargo check -p xai-grok-pager-bin
```

---

## 5. 运行需要什么环境

### 5.1 只用官方安装包

最低需要：

- Windows / Linux / macOS
- 终端环境
- 网络
- xAI/Grok 账号或相关认证
- Git，方便管理代码项目

不需要自己安装 Rust，也不需要编译。

### 5.2 从源码编译

需要额外安装：

- Rust / rustup
- protoc / protobuf-compiler
- Git
- C/C++ 构建工具链

Ubuntu / Debian 可参考：

```bash
sudo apt update
sudo apt install -y curl git build-essential protobuf-compiler
```

---

## 6. 能不能完全本地离线跑

要分两层看：

### 6.1 工具本体本地跑

可以。Grok Build 的 CLI/TUI 工具本体可以在本机运行。

### 6.2 AI 模型完全本地跑

不一定。

Grok Build 只是工具框架，真正负责写代码和推理的是模型。可以接：

- xAI / Grok API
- OpenAI API
- Claude API
- 其他兼容 OpenAI API 的服务
- 自己本地部署的模型服务

本地模型服务可以考虑：

- Ollama
- llama.cpp server
- vLLM
- LM Studio
- 其他 OpenAI-compatible 本地 API

适合的本地代码模型：

- Qwen Coder
- DeepSeek Coder
- GLM Coder
- CodeLlama
- StarCoder 系列

---

## 7. MS-02 Ultra 上的建议

如果机器是 **铭凡 MS-02 Ultra 285HX**：

- 跑 Grok Build 工具本体：没问题
- 接云端 Grok / OpenAI / Claude API：很适合
- 跑本地小模型：可以
- 跑大模型：主要受内存和显卡限制

### 7.1 32G 内存

适合：

- Grok Build 工具本体
- 云端 API 编程助手
- 7B / 14B 量化本地模型

不太适合：

- 多虚拟机 + 本地大模型同时跑
- 32B / 70B 级模型

### 7.2 后期升级建议

- **64G**：比较舒服
- **128G**：适合 PVE + Docker + 多服务 + 本地模型
- **256G**：适合重度虚拟化和更大模型折腾

---

## 8. 推荐部署方案

### 方案 A：Windows 11 直接装

适合先快速试用。

步骤：

1. 安装 Git
2. 安装 Grok Build
3. 登录认证
4. 进入项目目录运行：

```bash
grok
```

优点：简单。  
缺点：服务器化不如 Linux / PVE。

---

### 方案 B：Ubuntu Server / Debian

适合稳定跑开发环境。

```bash
sudo apt update
sudo apt install -y curl git build-essential protobuf-compiler

curl -fsSL https://x.ai/cli/install.sh | bash

grok --version
```

进入项目目录：

```bash
cd /path/to/project
grok
```

优点：干净、稳定、适合 Docker。  
缺点：桌面使用不如 Windows。

---

### 方案 C：PVE + Ubuntu VM（推荐）

底层安装 Proxmox VE，然后开一个 Ubuntu Server VM 专门跑：

- Grok Build
- Docker
- 开发环境
- AI 工具
- 本地模型服务

优点：

- 系统隔离
- 方便快照
- 折腾坏了能回滚
- 后续可继续加 NAS、Windows VM、OpenWrt VM

推荐结论：

> MS-02 Ultra 这种机器更适合 **PVE + Ubuntu VM + 云端模型 / 本地轻量模型** 的玩法。

---

## 9. 一句话总结

Grok Build 是一个开源 AI 编程工具框架，适合作为本地开发终端助手使用。  
它本体可以本地运行，但不等于免费本地运行满血 Grok 模型。  
如果使用 MS-02 Ultra，推荐底层装 PVE，在 Ubuntu VM 中运行 Grok Build，先接云端模型，后续再根据内存和显卡条件折腾本地模型。
