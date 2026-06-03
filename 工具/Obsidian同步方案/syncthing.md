# Syncthing 同步方案

## 简介
Syncthing 是一款开源免费的点对点文件同步工具，不经过第三方服务器，隐私性最好。

---

## 优点
- ✅ **免费开源**：完全免费，代码开源可审计
- ✅ **隐私安全**：点对点加密传输，不经过第三方
- ✅ **跨平台**：支持 Windows、Mac、Linux、Android
- ✅ **实时同步**：文件变化后自动同步
- ✅ **无需服务器**：设备间直接连接
- ✅ **支持中继**：不在同一网络也能同步

## 缺点
- ❌ **需要安装**：每台设备都要安装客户端
- ❌ **设备需在线**：同步时双方设备需要在线（有中继可缓解）
- ❌ **iOS 支持差**：iOS 只有第三方收费 App（Möbius Sync）
- ❌ **配置稍复杂**：需要互相添加设备和文件夹

---

## 安装方法

### Windows
```powershell
# 方法1：winget（推荐）
winget install Syncthing.Syncthing

# 方法2：Scoop
scoop install syncthing

# 方法3：官网下载
# https://syncthing.net/downloads/
```

### Linux
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install syncthing

# 或下载二进制
curl -sL https://github.com/syncthing/syncthing/releases/latest/download/syncthing-linux-amd64.tar.gz | tar xz
```

### macOS
```bash
# Homebrew
brew install syncthing
```

### Android
- Google Play 搜索 "Syncthing"
- F-Droid 搜索 "Syncthing"

---

## 配置步骤

### 1. 启动 Syncthing
```bash
# Linux/Mac
syncthing

# Windows：安装后自动启动，或在开始菜单找到
```

启动后自动打开浏览器，访问 `http://127.0.0.1:8384`

### 2. 添加设备
1. 在 Web 界面点击 **"Add Remote Device"**
2. 输入另一台设备的 **Device ID**
3. 给设备起名（如 "家庭电脑"、"手机"）

### 3. 添加同步文件夹
1. 点击 **"Add Folder"**
2. **Folder Path**：选择 Obsidian vault 路径
3. **Folder Label**：起名（如 "Obsidian Vault"）
4. 在 **"Sharing"** 标签页勾选要同步的设备

### 4. 其他设备重复操作
- 添加相同文件夹，选择同步设备
- 两边自动开始同步

---

## 关键设置

| 设置项 | 推荐值 | 说明 |
|--------|--------|------|
| Folder Type | Send & Receive | 双向同步 |
| Watch for Changes | ✅ 开启 | 实时监测文件变化 |
| File Versioning | Simple | 保留旧版本，防止误删 |

---

## Obsidian 专用忽略规则

在同步文件夹创建 `.stignore` 文件：
```
// 忽略工作区状态
.obsidian/workspace.json
.obsidian/workspace-mobile.json

// 忽略缓存
.trash/
.DS_Store
```

---

## 常见问题

### Q：设备不在同一局域网能同步吗？
A：可以，Syncthing 会通过中继服务器连接，速度可能慢点但能用。

### Q：会冲突吗？
A：两边同时编辑同一文件会生成冲突副本（`.sync-conflict-*`），手动合并即可。

### Q：手机能同步吗？
A：Android 有官方 App，iOS 有 Möbius Sync（第三方，收费）。

---

## 适用场景
- 注重隐私安全
- 多设备（电脑+手机）同步
- 不想依赖第三方服务
- 有一定技术基础

---

*相关链接：*
- 官网：https://syncthing.net/
- 文档：https://docs.syncthing.net/
- GitHub：https://github.com/syncthing/syncthing
