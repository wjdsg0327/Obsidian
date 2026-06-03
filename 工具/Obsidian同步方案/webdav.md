# WebDAV 自建方案

## 简介
自建 WebDAV 服务器，完全控制数据，配合 Remotely Save 插件或直接挂载使用。

---

## 优点
- ✅ **完全控制**：数据在自己服务器上
- ✅ **隐私安全**：不经过第三方
- ✅ **免费**：只需要有服务器
- ✅ **稳定**：自己维护，不受第三方影响
- ✅ **灵活**：可自定义配置

## 缺点
- ❌ **需要服务器**：需要有 VPS 或 NAS
- ❌ **技术门槛**：需要会部署 Docker
- ❌ **维护成本**：需要自己维护服务器
- ❌ **公网访问**：需要公网 IP 或内网穿透

---

## Docker 方案对比

| 方案 | 镜像 | 优点 | 缺点 | 推荐度 |
|------|------|------|------|--------|
| bytemark/webdav | bytemark/webdav | 简单 | 功能少，已停更 | ⭐⭐ |
| hacdias/webdav | hacdias/webdav | 轻量，Go 实现 | 功能较少 | ⭐⭐⭐⭐ |
| Nginx WebDAV | nginx:alpine | 稳定，功能全 | 配置复杂 | ⭐⭐⭐⭐⭐ |
| Rclone | rclone/rclone | 支持多种存储 | 配置复杂 | ⭐⭐⭐ |
| Caddy WebDAV | kencx/webdav | 自动 HTTPS | 配置稍复杂 | ⭐⭐⭐⭐ |

---

## 方案一：hacdias/webdav（推荐）

### docker-compose.yml
```yaml
version: '3'
services:
  webdav:
    image: hacdias/webdav:latest
    container_name: webdav
    restart: always
    ports:
      - "1122:80"
    volumes:
      - ./data:/data
      - ./config.yml:/config.yml
    command: --config /config.yml
```

### config.yml
```yaml
address: 0.0.0.0
port: 80
scope: /data
modify: true
auth: true
users:
  - username: admin
    password: your_secure_password
```

### 启动命令
```bash
mkdir -p data
docker-compose up -d
```

### 测试
```bash
curl -u admin:your_secure_password http://localhost:1122/
```

---

## 方案二：Nginx WebDAV（最稳定）

### docker-compose.yml
```yaml
version: '3'
services:
  webdav:
    image: nginx:alpine
    container_name: webdav
    restart: always
    ports:
      - "1122:80"
    volumes:
      - ./data:/usr/share/nginx/html
      - ./nginx.conf:/etc/nginx/nginx.conf
```

### nginx.conf
```nginx
events {
    worker_connections 1024;
}

http {
    server {
        listen 80;
        server_name localhost;

        root /usr/share/nginx/html;
        
        # WebDAV 方法
        dav_methods PUT DELETE MKCOL COPY MOVE;
        dav_ext_methods PROPFIND OPTIONS;
        dav_access user:rw group:rw all:r;

        # 认证
        auth_basic "WebDAV";
        auth_basic_user_file /etc/nginx/.htpasswd;

        # 允许的文件大小
        client_max_body_size 0;

        # 日志
        access_log /var/log/nginx/webdav.access.log;
        error_log /var/log/nginx/webdav.error.log;
    }
}
```

### 创建认证文件
```bash
# 安装 htpasswd 工具
sudo apt install apache2-utils

# 创建密码文件
htpasswd -c .htpasswd admin

# 复制到容器
docker cp .htpasswd webdav:/etc/nginx/.htpasswd
```

### 启动
```bash
mkdir -p data
docker-compose up -d
```

---

## 方案三：Rclone WebDAV（支持多种存储）

### rclone.conf
```ini
[local]
type = local
nounc = true

[webdav]
type = webdav
url = http://localhost:1122
vendor = other
user = admin
pass = your_encrypted_password
```

### docker-compose.yml
```yaml
version: '3'
services:
  rclone:
    image: rclone/rclone:latest
    container_name: rclone-webdav
    restart: always
    ports:
      - "1122:80"
    volumes:
      - ./rclone.conf:/config/rclone/rclone.conf
      - ./data:/data
    command: serve webdav /data --addr :80 --user admin --pass your_password
```

---

## 方案四：Caddy WebDAV（自动 HTTPS）

### Caddyfile
```
:80 {
    webdav {
        root /data
        prefix /dav
    }
    
    basicauth {
        admin $2a$14$HASHED_PASSWORD
    }
}
```

### docker-compose.yml
```yaml
version: '3'
services:
  webdav:
    image: kencx/webdav:latest
    container_name: webdav
    restart: always
    ports:
      - "1122:80"
    volumes:
      - ./data:/data
      - ./Caddyfile:/etc/caddy/Caddyfile
```

---

## 配置 Obsidian 使用 WebDAV

### 方法一：Remotely Save 插件
1. 安装 Remotely Save 插件
2. 选择 WebDAV
3. 填入信息：
   - **Address**：`http://your-server:1122/`
   - **Username**：admin
   - **Password**：your_password
4. 测试连接并保存

### 方法二：系统 WebDAV 挂载

#### Windows
1. 打开文件资源管理器
2. 右键"此电脑" → "映射网络驱动器"
3. 输入：`http://your-server:1122/`
4. 输入用户名密码

#### macOS
1. Finder → 前往 → 连接服务器
2. 输入：`http://your-server:1122/`
3. 输入用户名密码

#### Linux
```bash
# 安装 davfs2
sudo apt install davfs2

# 创建挂载点
sudo mkdir /mnt/webdav

# 挂载
sudo mount -t davfs http://your-server:1122/ /mnt/webdav
```

---

## 安全建议

### 1. 使用 HTTPS
- 配置 SSL 证书
- 使用 Let's Encrypt 免费证书
- 或使用 Caddy 自动 HTTPS

### 2. 强密码
- 使用强密码
- 定期更换密码

### 3. 限制访问
- 配置防火墙
- 只允许特定 IP 访问
- 使用 VPN 访问

### 4. 定期备份
- 备份 WebDAV 数据
- 备份配置文件

---

## 常见问题

### Q：无法连接怎么办？
A：检查防火墙、端口、服务是否启动。

### Q：上传大文件失败？
A：检查 nginx 的 client_max_body_size 设置。

### Q：速度慢怎么办？
A：检查服务器带宽，考虑使用 CDN。

### Q：多人使用会冲突吗？
A：WebDAV 本身不处理冲突，需要客户端处理。

---

## 适用场景
- 有自己的服务器/VPS
- 注重数据隐私
- 有一定技术基础
- 想要完全控制

---

*相关链接：*
- hacdias/webdav：https://github.com/hacdias/webdav
- Nginx WebDAV：https://nginx.org/en/docs/http/ngx_http_dav_module.html
- Rclone：https://rclone.org/
- Caddy：https://caddyserver.com/
