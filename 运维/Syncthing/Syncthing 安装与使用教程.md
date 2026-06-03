# Syncthing 安装与使用教程

整理时间：2026-06-03

## 1. Syncthing 是什么

Syncthing 是一个开源的点对点文件同步工具，可以在多台设备之间实时同步文件夹。

它的特点：

- 免费开源。
- 不依赖中心服务器。
- 设备之间点对点同步。
- 数据传输加密。
- 支持 Windows、Linux、macOS、Android、NAS、Docker。
- 很适合同步 Obsidian 知识库、配置文件、项目文档。

简单理解：

```text
A 电脑的某个文件夹 <——> B 电脑的某个文件夹 <——> 手机上的某个文件夹
```

只要设备在线，Syncthing 会自动同步变化。

---

## 2. Syncthing 适合做什么

适合：

- 多设备同步 Obsidian 知识库。
- Windows 和 Linux/WSL/服务器之间同步文件。
- 手机和电脑同步笔记、文档、照片。
- 局域网高速同步大文件。
- 替代部分网盘同步功能。

不太适合：

- 作为完整备份方案。
- 多人同时编辑同一个文件。
- 需要历史版本管理的场景。
- 文件冲突非常频繁的团队协作。

如果同步 Obsidian，建议：

```text
Syncthing 负责设备间同步文件
GitHub 负责版本历史和远程备份
```

两者可以配合，但要注意冲突。

---

## 3. 核心概念

### 3.1 设备 Device

每台安装 Syncthing 的机器都是一个设备。

每台设备都有一个唯一的 Device ID。

### 3.2 文件夹 Folder

你要同步的是某个文件夹，例如：

```text
D:\老王
/home/wang/notes
/storage/emulated/0/Documents/Obsidian
```

### 3.3 共享 Sharing

你需要把某个文件夹共享给另一台设备，对方接受后才会同步。

### 3.4 同步方向

Syncthing 支持几种文件夹类型：

| 类型 | 含义 |
|---|---|
| Send & Receive | 双向同步，默认模式 |
| Send Only | 只发送，不接收修改 |
| Receive Only | 只接收，不主动发送修改 |

一般个人多设备同步用：

```text
Send & Receive
```

如果某台机器只是备份机，可以用：

```text
Receive Only
```

---

## 4. Windows 安装

### 4.1 下载

官网：

```text
https://syncthing.net/downloads/
```

Windows 推荐两种方式：

1. **Syncthing 原版**：命令行 + Web 管理界面。
2. **SyncTrayzor**：Windows 托盘版，更适合桌面使用。

推荐普通用户使用 SyncTrayzor，体验更像普通软件。

SyncTrayzor 项目地址：

```text
https://github.com/canton7/SyncTrayzor
```

### 4.2 使用 SyncTrayzor

步骤：

1. 下载 SyncTrayzor 安装包。
2. 安装并启动。
3. 首次启动会自动启动 Syncthing。
4. 浏览器或软件窗口中打开管理界面。
5. 设置 GUI 用户名和密码。
6. 添加需要同步的文件夹。

默认管理地址通常是：

```text
http://127.0.0.1:8384
```

---

## 5. Linux 安装

### 5.1 Ubuntu / Debian 安装

可以直接用官方包源，或者简单使用 apt 安装。

简单安装：

```bash
sudo apt update
sudo apt install syncthing
```

查看版本：

```bash
syncthing --version
```

### 5.2 启动 Syncthing

当前用户启动：

```bash
syncthing
```

启动后访问：

```text
http://127.0.0.1:8384
```

### 5.3 systemd 用户服务

如果希望开机自动启动：

```bash
systemctl --user enable syncthing.service
systemctl --user start syncthing.service
```

查看状态：

```bash
systemctl --user status syncthing.service
```

如果用户服务开机不启动，可以开启 linger：

```bash
sudo loginctl enable-linger $USER
```

---

## 6. Docker 安装

适合服务器、NAS、长期运行环境。

### 6.1 Docker Compose 示例

按老王的 Docker 数据卷习惯，推荐放到：

```text
D:\work\docker\syncthing
WSL 路径：/mnt/d/work/docker/syncthing
```

`docker-compose.yml`：

```yaml
services:
  syncthing:
    image: syncthing/syncthing:latest
    container_name: syncthing
    restart: unless-stopped
    hostname: syncthing-server
    ports:
      - "8384:8384"       # Web 管理界面
      - "22000:22000/tcp" # TCP 同步端口
      - "22000:22000/udp" # QUIC 同步端口
      - "21027:21027/udp" # 局域网发现
    volumes:
      - /mnt/d/work/docker/syncthing/config:/var/syncthing/config
      - /mnt/d/work/docker/syncthing/data:/data
```

启动：

```bash
cd /mnt/d/work/docker/syncthing
docker compose up -d
```

查看日志：

```bash
docker logs -f syncthing
```

访问管理界面：

```text
http://服务器IP:8384
```

### 6.2 Docker 注意事项

如果管理界面暴露到局域网或公网，一定要设置：

- GUI 用户名。
- GUI 密码。
- 防火墙限制访问来源。

不要把 `8384` 直接暴露到公网。

---

## 7. Android 安装

Android 可以安装：

- Syncthing-Fork，常用且适合 Android。
- 或从 F-Droid 安装相关版本。

一般推荐：

```text
Syncthing-Fork
```

安装后：

1. 打开 App。
2. 查看本机 Device ID。
3. 添加电脑设备的 Device ID。
4. 选择手机上的同步目录。
5. 电脑端接受手机设备。
6. 双方确认共享文件夹。

Android 需要注意电池优化：

- 给 Syncthing 关闭电池优化。
- 允许后台运行。
- 如果系统杀后台严重，可能无法实时同步。

---

## 8. 两台设备同步教程

假设：

- 设备 A：Windows 电脑。
- 设备 B：Linux 服务器。
- 要同步目录：`D:\老王`。

### 8.1 在两台设备安装 Syncthing

Windows 安装 SyncTrayzor 或 Syncthing。

Linux 安装：

```bash
sudo apt install syncthing
systemctl --user enable --now syncthing.service
```

### 8.2 打开 Web 管理界面

Windows：

```text
http://127.0.0.1:8384
```

Linux 如果是远程服务器，默认只监听本地。

可以 SSH 端口转发：

```bash
ssh -L 8384:127.0.0.1:8384 user@server-ip
```

然后本机浏览器打开：

```text
http://127.0.0.1:8384
```

### 8.3 互相添加设备

在 A 设备：

1. 右上角 Actions。
2. Show ID。
3. 复制 Device ID。

在 B 设备：

1. Add Remote Device。
2. 粘贴 A 的 Device ID。
3. 保存。

另一边会弹出提示，选择接受。

### 8.4 添加同步文件夹

在 A 设备添加文件夹：

```text
Folder Label: 老王知识库
Folder Path: D:\老王
Folder Type: Send & Receive
```

共享给 B 设备。

B 设备收到共享请求后，选择本地路径：

```text
/home/wang/老王
```

保存后开始同步。

---

## 9. 同步 Obsidian 知识库建议

### 9.1 推荐同步范围

可以同步整个 vault，例如：

```text
D:\老王
```

但建议排除一些不必要目录。

### 9.2 Syncthing 忽略规则

在同步文件夹根目录创建：

```text
.stignore
```

推荐内容：

```gitignore
// Git 内部目录，通常不建议用 Syncthing 同步 Git 仓库内部状态
.git

// Obsidian 临时状态
.obsidian/workspace*.json
.obsidian/cache

// 回收站和临时文件
.trash
.DS_Store
Thumbs.db
*.tmp
*.log

// Python 缓存
__pycache__
*.pyc

// Docker / 数据库运行数据
**/volumes
**/data
**/logs
**/wal
```

说明：

- 如果同时用 GitHub，同步 `.git` 容易出问题，建议让每台设备自己 `git pull/push`。
- Syncthing 同步正文 Markdown 文件即可。
- Obsidian 插件是否同步看个人习惯；如果多设备系统不同，插件目录可能产生冲突。

### 9.3 和 GitHub 配合

推荐模式：

```text
Syncthing：同步设备间当前文件
GitHub：保存版本历史和远程备份
```

不要依赖 Syncthing 当版本管理。误删、冲突、覆盖时，Git 才方便恢复。

---

## 10. 常见配置

### 10.1 设置管理界面密码

进入 Web UI：

```text
Actions -> Settings -> GUI
```

设置：

- GUI Authentication User
- GUI Authentication Password

### 10.2 修改设备名称

```text
Actions -> Settings -> General -> Device Name
```

### 10.3 设置文件版本控制

Syncthing 支持简单版本保留：

```text
Folder -> Edit -> File Versioning
```

常用选择：

- Simple File Versioning。
- Staggered File Versioning。

建议重要笔记开启简单版本控制。

### 10.4 限速

```text
Actions -> Settings -> Connections
```

可以设置上传/下载速度限制。

---

## 11. 远程服务器访问 GUI

Syncthing 默认只监听本地地址 `127.0.0.1:8384`。

如果服务器没有桌面环境，推荐 SSH 转发：

```bash
ssh -L 8384:127.0.0.1:8384 user@server-ip
```

然后本地浏览器打开：

```text
http://127.0.0.1:8384
```

不推荐直接把 GUI 暴露公网。

如果必须监听局域网，可以修改配置：

```text
Actions -> Settings -> GUI -> GUI Listen Address
```

改成：

```text
0.0.0.0:8384
```

但必须设置密码，并用防火墙限制访问。

---

## 12. 防火墙端口

Syncthing 常用端口：

| 端口 | 协议 | 作用 |
|---|---|---|
| 8384 | TCP | Web 管理界面 |
| 22000 | TCP/UDP | 文件同步 |
| 21027 | UDP | 局域网发现 |

服务器防火墙示例：

```bash
sudo ufw allow 22000/tcp
sudo ufw allow 22000/udp
sudo ufw allow 21027/udp
```

如果不需要远程访问 Web UI，不要开放 `8384`。

---

## 13. 常见问题

### 13.1 设备发现不了

检查：

- 两边是否都添加了对方 Device ID。
- 防火墙是否放行 22000。
- 是否能访问互联网中继。
- 局域网发现端口 21027 是否被阻止。

### 13.2 一直显示 Out of Sync

可能原因：

- 文件被占用。
- 权限不足。
- 路径不存在。
- 文件名在不同系统不兼容。
- 某些文件被忽略规则排除了。

### 13.3 出现冲突文件

Syncthing 会生成类似：

```text
xxx.sync-conflict-日期-设备名.md
```

处理方式：

1. 打开原文件和 conflict 文件。
2. 对比内容。
3. 合并需要保留的部分。
4. 删除 conflict 文件。
5. 如果使用 Git，提交一次合并结果。

### 13.4 手机不能实时同步

Android 常见原因：

- 系统杀后台。
- 电池优化限制。
- 没给存储权限。
- App 没设置后台运行。

解决：

- 关闭电池优化。
- 允许自启动。
- 允许后台活动。
- 打开 Syncthing-Fork 的运行通知。

---

## 14. 推荐方案：老王知识库同步

老王当前知识库路径：

```text
D:\老王
WSL 路径：/mnt/d/老王
```

推荐方案：

```text
Windows 主力机：D:\老王
其他电脑/服务器：各自本地目录
手机：Documents/Obsidian/老王
GitHub：远程版本备份
Syncthing：设备间实时同步
```

建议策略：

1. 用 Syncthing 同步 Markdown 正文和必要附件。
2. 不用 Syncthing 同步 `.git`。
3. 每台电脑单独配置 Git 远程仓库。
4. 重要修改后执行 Git commit/push。
5. 冲突文件手动合并，不要直接删除。

---

## 15. 最小上手流程

如果只想最快跑起来：

1. Windows 安装 SyncTrayzor。
2. 手机安装 Syncthing-Fork。
3. 两边互相添加 Device ID。
4. Windows 添加文件夹 `D:\老王`。
5. 共享给手机。
6. 手机选择 Obsidian vault 路径。
7. 等同步完成。
8. 设置 GUI 密码。
9. 添加 `.stignore` 排除临时文件。
10. 重要笔记继续用 GitHub 做版本备份。

---

## 16. 一句话总结

Syncthing 适合做多设备实时同步，尤其适合 Obsidian 知识库；但它不是 Git，也不是网盘。推荐用：

```text
Syncthing 同步文件 + GitHub 保存版本
```

这样既方便多设备使用，又能保留历史版本，出问题也容易恢复。
