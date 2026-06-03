# Syncthing Docker 服务器部署教程

整理时间：2026-06-03

## 1. 适用场景

如果你有一台服务器，Syncthing 很适合用 Docker 部署成一个常驻同步节点。

典型用途：

```text
Windows 电脑 <——> 服务器 Syncthing <——> 手机 / 其他电脑
```

这样即使电脑和手机没有同时在线，只要服务器在线，就可以作为中间节点同步文件。

适合同步：

- Obsidian 知识库
- 文档目录
- 配置文件
- 项目资料
- 手机和电脑之间的文件

---

## 2. 推荐目录规划

按老王的 Docker 数据目录约定，服务器上推荐这样放：

```text
/mnt/d/work/docker/syncthing
├── config/          # Syncthing 配置
├── data/            # 需要同步的数据根目录
└── docker-compose.yml
```

如果是普通 Linux 服务器，没有 `/mnt/d`，可以改成：

```text
/data/docker/syncthing
├── config/
├── data/
└── docker-compose.yml
```

---

## 3. Docker Compose 部署

创建目录：

```bash
mkdir -p /data/docker/syncthing/config
mkdir -p /data/docker/syncthing/data
cd /data/docker/syncthing
```

创建 `docker-compose.yml`：

```yaml
services:
  syncthing:
    image: syncthing/syncthing:latest
    container_name: syncthing
    hostname: syncthing-server
    restart: unless-stopped
    ports:
      - "8384:8384"       # Web 管理界面，不建议公网裸露
      - "22000:22000/tcp" # TCP 同步端口
      - "22000:22000/udp" # QUIC 同步端口
      - "21027:21027/udp" # 局域网发现
    volumes:
      - ./config:/var/syncthing/config
      - ./data:/data
```

启动：

```bash
docker compose up -d
```

查看日志：

```bash
docker logs -f syncthing
```

查看容器：

```bash
docker ps | grep syncthing
```

---

## 4. 访问管理界面

浏览器访问：

```text
http://服务器IP:8384
```

第一次打开后，建议立刻设置管理账号密码：

```text
Actions -> Settings -> GUI
```

设置：

- GUI Authentication User
- GUI Authentication Password

不要让 `8384` 裸奔公网。

---

## 5. 防火墙端口

Syncthing 常用端口：

| 端口 | 协议 | 说明 |
|---|---|---|
| 8384 | TCP | Web 管理界面 |
| 22000 | TCP | 文件同步 |
| 22000 | UDP | QUIC 同步 |
| 21027 | UDP | 局域网发现 |

服务器建议开放同步端口：

```bash
sudo ufw allow 22000/tcp
sudo ufw allow 22000/udp
```

如果服务器在公网，`21027/udp` 局域网发现意义不大，可以不开。

管理界面 `8384` 推荐只允许自己 IP 访问，或只通过 SSH 隧道访问。

不建议：

```bash
sudo ufw allow 8384/tcp
```

更安全做法是 SSH 隧道：

```bash
ssh -L 8384:127.0.0.1:8384 root@服务器IP
```

然后本地浏览器打开：

```text
http://127.0.0.1:8384
```

---

## 6. 如果想让 GUI 只本地访问

如果服务器只通过 SSH 隧道管理，可以不要把 `8384` 暴露到公网。

Compose 可以改成：

```yaml
ports:
  - "127.0.0.1:8384:8384"
  - "22000:22000/tcp"
  - "22000:22000/udp"
```

这样外部无法直接访问：

```text
http://服务器IP:8384
```

只能通过 SSH 隧道访问，更安全。

---

## 7. 添加同步目录

服务器容器里同步数据统一放在：

```text
/data
```

例如要同步 Obsidian 知识库，可以在服务器 Syncthing 里添加文件夹：

```text
Folder Label: 老王知识库
Folder Path: /data/老王
Folder Type: Send & Receive
```

宿主机实际路径是：

```text
/data/docker/syncthing/data/老王
```

如果是 Windows / WSL 目录约定：

```text
/mnt/d/work/docker/syncthing/data/老王
```

---

## 8. 和 Windows 电脑同步

Windows 端推荐安装 SyncTrayzor。

流程：

1. Windows 打开 SyncTrayzor。
2. 复制 Windows 设备 ID。
3. 服务器 Syncthing 添加 Windows 设备 ID。
4. Windows 接受服务器设备。
5. Windows 添加文件夹，例如：

```text
D:\老王
```

6. 共享给服务器。
7. 服务器接受共享，路径填：

```text
/data/老王
```

---

## 9. 和手机同步

Android 推荐 Syncthing-Fork。

流程：

1. 手机安装 Syncthing-Fork。
2. 手机复制 Device ID。
3. 服务器添加手机设备。
4. 手机接受服务器设备。
5. 服务器把 `/data/老王` 文件夹共享给手机。
6. 手机选择本地 Obsidian vault 路径。

注意：Android 要关闭电池优化，否则后台同步可能不稳定。

---

## 10. 推荐 `.stignore`

如果同步 Obsidian 知识库，建议在知识库根目录放 `.stignore`：

```gitignore
// 不同步 Git 内部数据，避免多设备 Git 状态互相污染
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

- Syncthing 同步正文文件。
- GitHub 负责版本历史。
- 不建议用 Syncthing 同步 `.git` 目录。

---

## 11. 推荐部署方案

### 方案 A：公网服务器常驻节点

```text
电脑 <——> 公网服务器 <——> 手机
```

优点：

- 服务器 24 小时在线。
- 手机和电脑不需要同时在线。
- 跨网络也能同步。

注意：

- 开放 `22000/tcp` 和 `22000/udp`。
- 管理界面 `8384` 不要裸露公网。

### 方案 B：内网服务器 / NAS

```text
电脑 <——> NAS/家用服务器 <——> 手机
```

优点：

- 局域网速度快。
- 数据不出家里。

注意：

- 外网同步需要内网穿透、VPN 或公网 IP。

---

## 12. 更新 Syncthing

进入目录：

```bash
cd /data/docker/syncthing
```

拉取新镜像：

```bash
docker compose pull
```

重启：

```bash
docker compose up -d
```

清理旧镜像：

```bash
docker image prune -f
```

---

## 13. 备份

至少备份这两个目录：

```text
/data/docker/syncthing/config
/data/docker/syncthing/data
```

其中：

- `config`：设备 ID、配置、共享关系。
- `data`：同步文件。

如果只备份数据，不备份 config，重新部署后设备 ID 会变化，需要重新配对。

---

## 14. 一句话建议

老王有服务器的话，推荐直接用 Docker 部署 Syncthing，当作 24 小时在线同步节点：

```text
服务器 Docker Syncthing + Windows SyncTrayzor + 手机 Syncthing-Fork + GitHub 版本备份
```

这是同步 Obsidian 知识库比较稳的一套组合。
