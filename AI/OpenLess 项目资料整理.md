---
title: OpenLess 项目资料整理
date: 2026-06-05
tags:
  - AI
  - 语音输入
  - 开源项目
  - Tauri
  - Rust
  - TypeScript
source:
  - https://github.com/Open-Less/openless
  - https://openless.top
  - https://raw.githubusercontent.com/Open-Less/openless/beta/README.zh.md
  - https://raw.githubusercontent.com/Open-Less/openless/beta/USAGE.md
---

# OpenLess 项目资料整理

> OpenLess 是一款面向 macOS 与 Windows 的开源语音输入工具：按住快捷键说话，经过 ASR 转写与 AI 润色后，文本会插入到当前光标位置；如果插入失败，则回退复制到剪贴板。

## 一句话定位

**OpenLess = 开源、本地优先、可自带模型/凭据的 AI 语音输入工具，重点不是“逐字听写”，而是把口语整理成可直接使用的书面文本，尤其适合生成 AI Prompt。**

## 基本信息

- **项目名：** OpenLess
- **GitHub：** https://github.com/Open-Less/openless
- **官网：** https://openless.top
- **许可证：** MIT License
- **默认分支：** `beta`
- **支持平台：**
  - macOS 12+
  - Windows 10+
- **主要技术栈：**
  - Tauri 2
  - Rust
  - React
  - TypeScript
- **项目状态：** README 中标注当前版本为 v1.3.6
- **社区：** QQ 群 1078960553，Discord 社区

## 核心价值

OpenLess 想解决的问题是：很多时候不是没有想法，而是打字太慢、口语太乱、手动整理 Prompt 太耗时。

它的核心价值包括：

1. **用说话代替打字**
   - 在任意输入框中触发语音输入。
   - 支持 ChatGPT、Claude、Cursor、Notion、邮件、聊天框等场景。

2. **把口语整理成可用文本**
   - 去掉语气词。
   - 修正标点和明显错字。
   - 整理段落与结构。
   - 按指定模式润色。

3. **特别适合生成 AI Prompt**
   - 用户可以随口说需求。
   - OpenLess 将其整理为带上下文、约束、目标的结构化 Prompt。
   - 可直接粘贴给 ChatGPT、Claude、Cursor 等工具。

4. **开源与本地优先**
   - 对标 Typeless、Wispr Flow、Lazy、Superwhisper 等商业工具。
   - 代码公开。
   - 数据、历史、词典等尽量留在本机。
   - 凭据存放在系统凭据保管库，而不是明文配置文件。

## 工作流程

典型流程：

```text
按下/触发全局快捷键
→ 录音
→ ASR 转写
→ LLM 润色/结构化
→ 插入到当前光标位置
→ 插入失败时复制到剪贴板
→ 保存历史记录
```

更底层的流水线：

```text
hotkey edge
→ Recorder.start + ASR.openSession
→ audio frames
→ hotkey edge
→ Recorder.stop + ASR.sendLastFrame
→ Polish
→ Insert
→ History.save
```

## 主要功能

### 1. 全局语音输入

- 在任意文本输入框中使用。
- 默认快捷键：
  - macOS：右 Option
  - Windows：右 Control
- 支持按住说话（push-to-talk）和切换式录音。
- 支持 MediaPlayPause 触发，方便耳机线控开始/停止录音。
- `Esc` 可取消当前录音、润色或插入过程。

### 2. 输出模式

OpenLess 支持多种输出模式：

| 模式 | 说明 |
| --- | --- |
| 原文 | 直接输出转写文本，不做修改 |
| 轻度润色 | 修正语气词、标点、明显错字，保留原意 |
| 清晰结构 / AI Prompt 模式 | 把口语整理成有结构、有上下文、有约束的 Prompt |
| 正式表达 | 将口语转换为正式书面语 |

此外还有一个翻译快捷键，可将语音直接转换为配置的目标语言。

### 3. 风格包与市场

OpenLess 不只提供固定的润色风格，而是支持 **Style Pack / 风格包**：

- 每个风格包有自己的系统提示词。
- 可以创建自己的风格，例如：
  - 简洁工程 commit message
  - 温暖客服回复
  - 小红书文案
  - 正式报告
  - 团队统一语气
- 可用快捷键切换当前风格。
- 可从 Marketplace 安装社区风格包。
- 可用 GitHub 身份发布自己的风格包。
- 市场内容会经过审核后再公开。

### 4. 流式插入

- 润色后的文本可逐字符写入当前光标位置。
- 这样感知延迟更低，体验接近“边想边说边落字”。
- 如果目标应用不支持流式按键，会自动回退为一次性粘贴。

### 5. 词典

词典用于提升专有名词识别与润色准确率：

- 支持手动添加正确拼写、分类、备注。
- 词条作为 Volcengine ASR 的 `context.hotwords` 注入，提高识别准确率。
- 词条也会作为语义提示参与润色，例如把上下文中的错误识别修正为正确产品名。
- 应用会从历史中学习候选纠正并推荐。

适合加入词典的内容：

- 产品名
- 人名
- 公司名
- 技术名词
- 项目名
- 特殊缩写

### 6. 历史记录

主窗口的「历史」页可以查看录音记录，包括：

- 原始转写
- 润色结果
- 可能的词典命中信息

### 7. 选区问答面板

README 中提到：OpenLess 支持一个独立快捷键打开浮动面板，对任意应用中高亮选中的文本进行语音问答。

### 8. 本地模型管理

OpenLess 支持在设置中管理本地 ASR 模型在磁盘上的存储。

### 9. 多语言界面

支持语言：

- 简体中文
- 繁體中文
- English
- 日本語
- 한국어

首次启动会自动检测语言，也可在设置中切换。

### 10. 自动更新与 Beta 频道

- Stable 用户通过 Tauri updater 获取正式更新。
- Beta 频道需要用户手动加入。
- Beta 版本不会自动推送给 Stable 用户。

## ASR 与 LLM 提供方

### ASR 语音识别

OpenLess 支持：

- **云端 ASR**
  - Volcengine 流式 ASR
  - OpenAI Whisper 兼容的批量 ASR
  - Apple Speech（macOS）

- **本地 ASR**
  - 内置 Qwen3-ASR 0.6B / 1.7B，通过 vendored `Open-Less/qwen-asr`
  - Windows 上的 Foundry Local Whisper 变体

### 润色模型

支持 OpenAI-compatible Chat Completions 形式的模型提供方，包括：

- Ark
- DeepSeek
- OpenAI
- Doubao
- Anthropic 兼容接口
- 任意用户自带的 OpenAI 兼容端点

README 中 Ark 默认端点：

```text
https://ark.cn-beijing.volces.com/api/v3/chat/completions
```

## 凭据与隐私

OpenLess 的凭据存放在操作系统凭据保管库中：

```text
service = com.openless.app
```

对应平台：

- macOS：Keychain
- Windows：Credential Manager
- Linux：keyring

旧版明文凭据文件仅作为迁移来源读取，迁移成功后删除：

```text
macOS / Linux: ~/.openless/credentials.json
Windows:       %APPDATA%\OpenLess\credentials.json
```

需要准备的凭据：

- Volcengine 流式 ASR
  - App ID
  - Access Token
  - Resource ID
- Ark 润色
  - API Key
  - Model ID
  - Endpoint

## 安装与首次配置

### macOS

1. 从 GitHub Releases 下载：
   - Apple Silicon：`OpenLess_<version>_aarch64.dmg`
   - Intel：`OpenLess_<version>_x64.dmg`
2. 打开 DMG，将 OpenLess.app 拖入 `/Applications`。
3. 首次启动后授予权限：
   - 麦克风权限
   - 辅助功能权限
4. 授权辅助功能后需要完全退出并重新打开应用。
5. 在设置中填写 Volcengine ASR 与 Ark 凭据。

如果遇到 macOS Gatekeeper 的“已损坏”提示，README 中给出的处理方式是：

```bash
xattr -cr /Applications/OpenLess.app
```

Homebrew 安装：

```bash
brew tap appergb/openless https://github.com/appergb/openless
brew install --cask openless
xattr -cr /Applications/OpenLess.app
```

升级：

```bash
brew update && brew upgrade openless
```

### Windows

1. 从 GitHub Releases 下载：
   - `OpenLess_<version>_x64-setup.exe`
2. 运行安装程序。
3. 首次启动时授予麦克风权限。
4. 打开「设置 → 权限」，确认全局快捷键监听器已启动。
5. 在设置中填写 Volcengine ASR 与 Ark 凭据。

## 常见问题

### 快捷键没反应

- macOS：确认已授予辅助功能权限，并且授权后重启过 OpenLess。
- Windows：在「设置 → 权限」中检查全局快捷键监听器状态。

### 识别结果为空或异常

优先检查火山引擎 ASR 凭据是否填写正确。

### 文字没有插入，只复制到了剪贴板

某些应用不支持辅助功能写入或模拟输入，OpenLess 会自动回退到剪贴板。此时手动粘贴即可。

### Windows 全屏游戏中不可用

README 中说明这主要是 Windows 系统层限制：

- 独占全屏应用不会显示 OpenLess capsule。
- 若目标程序以管理员运行而 OpenLess 不是，Windows UIPI 会阻止按键注入。
- 在 Minecraft 等游戏中，需要先打开聊天框，文本才能输入进去。

建议：

- 使用无边框窗口化全屏。
- 保持 OpenLess 与目标应用权限一致。
- 确认目标输入框已经获得焦点。

## 开发者资料

### 代码结构

活跃代码库位于：

```text
openless-all/app/
```

这是一个 Tauri 2 + Rust + React/TypeScript 应用。

macOS 构建会链接 vendored C ASR 引擎：

```text
Open-Less/qwen-asr
```

该引擎 fork 自：

```text
antirez/qwen-asr
```

作为 git 子模块位于：

```text
openless-all/app/src-tauri/vendor/qwen-asr/
```

首次克隆后需要初始化子模块：

```bash
git submodule update --init --recursive
```

### 后端模块

Tauri 后端由 Rust 实现，README 中列出的模块职责：

```text
types.rs         纯值类型：DictationSession、PolishMode、HotkeyBinding、errors
hotkey.rs        全局快捷键：macOS CGEventTap、Windows WH_KEYBOARD_LL、Linux rdev
recorder.rs      麦克风 → 16 kHz mono Int16 PCM，RMS callback
asr/             Volcengine streaming ASR + Whisper HTTP
polish.rs        OpenAI-compatible chat completions：Ark / DeepSeek / etc.
insertion.rs     AX focused-element → clipboard + Cmd+V → copy-only fallback
persistence.rs   History / preferences / vocab JSON + platform credential vault
permissions.rs   TCC checks：Accessibility / Microphone
coordinator.rs   状态机：Idle → Starting → Listening → Processing
commands.rs      Tauri IPC surface
```

### 前端结构

- React 前端位于 `src/`。
- 状态通过 Recoil atoms 管理：`pages/_atoms.tsx`。
- 快捷键能力与绑定通过 `HotkeySettingsContext`。
- 后端调用统一经由 `lib/ipc.ts`。

### 开发运行

```bash
cd openless-all/app
npm ci
npm run tauri dev
```

### macOS 构建

推荐使用项目脚本，而不是直接调用 `tauri build`：

```bash
cd openless-all/app
INSTALL=0 ./scripts/build-mac.sh
```

仅构建不安装：

```bash
INSTALL=0 ./scripts/build-mac.sh
```

构建并本地安装：

```bash
./scripts/build-mac.sh
```

### Rust 类型检查

```bash
cargo check --manifest-path src-tauri/Cargo.toml
```

### 前端检查

```bash
npm run build
```

### Windows 构建

Windows 构建前建议先跑预检：

```powershell
cd openless-all/app
powershell -ExecutionPolicy Bypass -File .\scripts\windows-preflight.ps1
```

#### MSVC 路线

适合安装了 Visual Studio Build Tools 与 Windows SDK 的环境：

```powershell
cd openless-all/app
npm ci
npm run tauri -- build
```

需要组件：

- `Microsoft.VisualStudio.Workload.VCTools`
- MSVC v143 x64/x86 build tools
- Windows 10/11 SDK，需包含 `kernel32.lib`

#### GNU / MinGW 路线

```powershell
cd openless-all/app
scoop install rustup mingw
rustup toolchain install stable-x86_64-pc-windows-gnu
rustup target add x86_64-pc-windows-gnu
powershell -ExecutionPolicy Bypass -File .\scripts\windows-preflight.ps1 -Toolchain gnu
powershell -ExecutionPolicy Bypass -File .\scripts\windows-build-gnu.ps1
```

### 日志路径

```text
macOS:   ~/Library/Logs/OpenLess/openless.log
Windows: %LOCALAPPDATA%\OpenLess\Logs\openless.log
```

## 发布流程

OpenLess 采用双频道分支模型：

```text
beta → main → release
```

### 分支

- `beta`
  - 默认分支。
  - Beta 频道。
  - 所有进行中的开发进入这里。
  - Beta 构建不会推送给普通用户，只提供给主动加入 Beta 的用户。

- `main`
  - Stable 频道。
  - 始终可发布。
  - 普通用户默认获得的正式构建。

### 贡献流程

```text
your fork / topic branch
→ PR to beta
→ AI review（一次，仅供参考）
→ maintainer 轻量检查
→ merge into beta
→ 双平台冒烟测试
→ merge beta into main
→ tag v<version>-tauri
→ release CI
→ Stable users
```

规则：

- PR 应提交到 `beta`，不要提交到 `main`。
- 提 PR 前，需要在目标平台验证改动。
- AI review 只跑一次，仅供参考，不建议围绕它反复循环。
- Beta 工作不得泄漏到 Stable。
- Stable 发布从 `main` 打 `v<version>-tauri` 标签。

### 发布标签

- Beta：

```text
v<version>-beta-tauri
```

- Stable：

```text
v<version>-tauri
```

### 版本号同步

发布前需要在以下 5 个位置提升版本号：

- `package.json`
- `package-lock.json`
  - 根级版本
  - `packages.""` 下的嵌套条目
- `src-tauri/tauri.conf.json`
- `src-tauri/Cargo.toml`
- `Cargo.lock` 中 `name = "openless"` 块

### 发布前验证

- 运行 macOS 构建脚本：

```bash
INSTALL=0 ./scripts/build-mac.sh
```

- 确认 `.app` 可启动。
- 在干净机器做冒烟测试：
  - 权限流程
  - 快捷键
  - 录音
  - ASR
  - 润色
  - 插入
  - 剪贴板回退
- 确认发布签名密钥已配置。

### GitHub Actions / 签名

Tagged releases `v*-tauri` 需要 Developer ID 签名与公证，让用户无需手动移除 quarantine 即可打开 macOS 应用。

需要 GitHub secrets：

- `APPLE_CERTIFICATE`
- `APPLE_CERTIFICATE_PASSWORD`
- `APPLE_ID`
- `APPLE_PASSWORD`
- `APPLE_TEAM_ID`

可选：

- `APPLE_PROVIDER_SHORT_NAME`
- `KEYCHAIN_PASSWORD`

## 与竞品对比

| 工具 | 形态 | OpenLess 的不同之处 |
| --- | --- | --- |
| Typeless | 闭源，macOS / Windows / iOS，订阅制 | 开源；显式 AI Prompt 模式；可自带 ASR + LLM；数据和词典留本机 |
| Wispr Flow | 闭源，macOS / Windows，订阅制 | 开源；自带 ASR + LLM；文本处理规则透明 |
| Lazy | 闭源笔记 / 速记工具 | 不是笔记容器，直接插入到任意输入框 |
| Superwhisper | 闭源 macOS，订阅制 | 开源；当前云端 ASR，本地 ASR 在路线图中 |

## 项目边界

OpenLess 明确只做一件事：

> 把语音变成可用的书面文字，尤其是 AI Prompt，并落在当前光标处。

它不做：

- 不回答用户问题。
- 不执行任务。
- 不分析项目。
- 不积累跨会话上下文。
- 每次听写都是独立的文本清理请求。

这一点对产品定位很重要：OpenLess 不是 AI Agent，也不是聊天机器人，而是语音输入与 Prompt 整理工具。

## 路线图 / 已规划能力

README 中提到的已规划或增强方向包括：

- 听写翻译模式：用一种语言说，插入为目标语言。
- 跨会话风格记忆：润色随时间学习用户语气。
- Snippets 片段功能。
- 历史增强：复制按钮、搜索、重新润色、重新插入。
- “粘贴上次结果”快捷键。

## 适合使用 OpenLess 的场景

- 给 ChatGPT / Claude / Cursor 写 Prompt。
- 快速写长消息、邮件、说明文档。
- 写代码注释、commit message、PR 描述。
- 把脑子里的零散想法快速整理成结构化文本。
- 需要跨应用输入，但不想维护一个额外笔记容器的场景。

## 需要注意的点

1. **凭据配置是首次使用门槛**
   - 尤其是 Volcengine ASR 与 Ark API。

2. **macOS 权限配置很关键**
   - 麦克风 + 辅助功能权限缺一不可。
   - 辅助功能授权后需要重启应用。

3. **插入能力受目标应用限制**
   - 某些应用无法直接写入，只能复制到剪贴板。

4. **它不是对话式 AI**
   - 用户说问题时，它只会整理成问题文本，而不会回答问题。

5. **Stable / Beta 分支与发布标签要严格区分**
   - PR 到 `beta`。
   - Stable 从 `main` 打 `v<version>-tauri`。
   - Beta 打 `v<version>-beta-tauri`。

## 快速摘要

OpenLess 是一个开源语音输入项目，面向 macOS 和 Windows。它通过全局快捷键录音，使用 ASR 转写，再通过 LLM 将口语润色或结构化为可用文本，并插入到当前光标位置。它最大的特点是面向 AI Prompt 场景：用户可以随口说需求，OpenLess 负责整理成清晰、带约束的 Prompt。技术上采用 Tauri 2 + Rust + React/TypeScript，支持云端和本地 ASR、多种 LLM 提供方、词典、历史、风格包市场、流式插入、多语言界面和双频道发布流程。它对标 Typeless、Wispr Flow、Lazy、Superwhisper 等商业工具，但强调开源、本地优先和可自带凭据。

## 参考链接

- GitHub 仓库：https://github.com/Open-Less/openless
- 官网：https://openless.top
- 中文 README：https://github.com/Open-Less/openless/blob/beta/README.zh.md
- 使用指南：https://github.com/Open-Less/openless/blob/beta/USAGE.md
- All-platform README：https://github.com/Open-Less/openless/blob/beta/openless-all/README.md
