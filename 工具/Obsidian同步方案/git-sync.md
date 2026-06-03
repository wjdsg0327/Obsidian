# Git 同步方案

## 简介
使用 Git 管理 Obsidian vault，推送到 GitHub/Gitee 私有仓库，配合 Obsidian Git 插件实现自动同步。

---

## 优点
- ✅ **版本历史**：完整记录每次修改，可回滚任意版本
- ✅ **免费**：GitHub/Gitee 私有仓库免费
- ✅ **冲突处理**：Git 有成熟的冲突解决机制
- ✅ **开发者友好**：熟悉 Git 的用户上手快
- ✅ **跨平台**：支持所有平台
- ✅ **可离线**：本地 commit，有网再 push

## 缺点
- ❌ **学习成本**：需要了解 Git 基础操作
- ❌ **手动同步**：需要手动或定时 pull/push
- ❌ **大文件限制**：GitHub 单文件限制 100MB
- ❌ **冲突需手动处理**：合并冲突需要一定经验
- ❌ **手机支持弱**：移动端 Git 工具较少

---

## 安装方法

### 1. 安装 Git
```bash
# Windows
winget install Git.Git

# Linux
sudo apt install git

# macOS
brew install git
```

### 2. 安装 Obsidian Git 插件
1. 打开 Obsidian
2. 设置 → 第三方插件 → 浏览
3. 搜索 "Obsidian Git"
4. 安装并启用

---

## 配置步骤

### 1. 初始化仓库
```bash
cd /path/to/your/vault
git init
git remote add origin https://github.com/username/repo.git
```

### 2. 创建 .gitignore
```gitignore
# Obsidian 工作区状态
.obsidian/workspace.json
.obsidian/workspace-mobile.json

# 系统文件
.DS_Store
Thumbs.db

# 大文件（可选）
*.mp4
*.zip
```

### 3. 首次提交
```bash
git add .
git commit -m "Initial commit"
git push -u origin main
```

### 4. 配置 Obsidian Git 插件
在 Obsidian 设置 → Obsidian Git：
- **Auto backup interval**：10（分钟）
- **Auto pull interval**：10（分钟）
- **Commit message**：`vault backup: {{date}}`
- **Pull updates on startup**：✅

---

## 日常使用

### 自动模式（推荐）
插件会自动：
- 每 10 分钟 commit 一次
- 每 10 分钟 pull 一次远程更新
- 启动时自动 pull

### 手动模式
- `Ctrl+P` → 输入 "Git"
- 选择 "Git: Pull" 拉取更新
- 选择 "Git: Commit all changes" 提交
- 选择 "Git: Push" 推送到远程

---

## 多设备同步流程

```
设备A 编辑 → commit → push → 远程仓库
                              ↓
设备B 编辑 → commit → pull ← 远程仓库
                              ↓
                         merge 冲突（如有）
```

---

## 冲突处理

### 自动合并
Git 会自动合并不冲突的修改

### 手动冲突
如果两边修改了同一文件的同一位置：
1. Git 会标记冲突区域
2. 手动选择保留哪个版本
3. 删除冲突标记
4. commit 解决后的文件

### 使用工具
```bash
# 使用 VS Code 解决冲突
code .

# 或使用专用工具
git mergetool
```

---

## 常见问题

### Q：GitHub 仓库大小限制？
A：建议 < 1GB，单文件 < 100MB。大文件用 Git LFS。

### Q：手机怎么同步？
A：Android 可用 MGit，iOS 可用 Working Copy（收费）。

### Q：忘了 push 怎么办？
A：本地 commit 不会丢，有网后 push 即可。

### Q：怎么回滚到某个版本？
A：在 GitHub 查看历史，或用 `git revert`。

---

## 适用场景
- 开发者，熟悉 Git
- 需要版本历史
- 多人协作（共享 vault）
- 想要备份保障

---

*相关链接：*
- Obsidian Git 插件：https://github.com/denolehov/obsidian-git
- GitHub：https://github.com
- Gitee（国内）：https://gitee.com
