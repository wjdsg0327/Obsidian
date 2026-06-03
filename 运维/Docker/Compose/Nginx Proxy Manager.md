# Nginx Proxy Manager Docker Compose

## 用途

使用 Docker Compose 启动 Nginx Proxy Manager，用于可视化管理反向代理和 HTTPS 证书。

## docker-compose.yml

```yaml
version: '3.8'
services:
  app:
    image: 'jc21/nginx-proxy-manager:latest'
    restart: unless-stopped
    ports:
      - '80:80'
      - '443:443'
      - '81:81'
    volumes:
      - ./data:/data
      - ./letsencrypt:/etc/letsencrypt
```

## 端口

- `80`：HTTP
- `443`：HTTPS
- `81`：管理后台

## 启动

```bash
docker compose up -d
```

## 注意

- 需要确保宿主机 80、443、81 端口未被占用。
- 公网环境下建议立刻修改默认管理员账号密码。
