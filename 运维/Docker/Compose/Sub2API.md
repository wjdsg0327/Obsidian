# Sub2API Docker Compose

## 用途

使用 Docker Compose 启动 Sub2API，并配套 PostgreSQL 与 Redis。

## 服务组件

- `sub2api`：主程序
- `postgres`：数据库
- `redis`：缓存

## docker-compose.yml

```yaml
services:
  sub2api:
    image: weishaw/sub2api:latest
    container_name: sub2api
    restart: unless-stopped
    ports:
      - "6780:8080"
    volumes:
      - ./data/sub2api:/app/data
    environment:
      - AUTO_SETUP=true
      - TZ=Asia/Shanghai
      - SERVER_MODE=release
      - DATABASE_HOST=postgres
      - DATABASE_PORT=5432
      - DATABASE_USER=sub2api
      - DATABASE_PASSWORD=sub2api_password
      - DATABASE_DBNAME=sub2api
      - REDIS_HOST=redis
      - REDIS_PORT=6379
      - REDIS_PASSWORD=redis_password
      - ADMIN_EMAIL=admin@sub2api.local
      - ADMIN_PASSWORD=admin123456
      - JWT_SECRET=
      - TOTP_ENCRYPTION_KEY=
      - SECURITY_URL_ALLOWLIST_ENABLED=false
      - SECURITY_URL_ALLOWLIST_ALLOW_INSECURE_HTTP=true
    depends_on:
      postgres:
        condition: service_healthy
      redis:
        condition: service_healthy
    networks:
      - sub2api-network

  postgres:
    image: postgres:18-alpine
    container_name: sub2api-postgres
    restart: unless-stopped
    volumes:
      - ./data/postgres:/var/lib/postgresql/data
    environment:
      - POSTGRES_USER=sub2api
      - POSTGRES_PASSWORD=sub2api_password
      - POSTGRES_DB=sub2api
      - PGDATA=/var/lib/postgresql/data
      - TZ=Asia/Shanghai
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U sub2api -d sub2api"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - sub2api-network

  redis:
    image: redis:8-alpine
    container_name: sub2api-redis
    restart: unless-stopped
    volumes:
      - ./data/redis:/data
    command: ["redis-server", "--requirepass", "redis_password", "--appendonly", "yes"]
    environment:
      - REDISCLI_AUTH=redis_password
      - TZ=Asia/Shanghai
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - sub2api-network

networks:
  sub2api-network:
    driver: bridge
```

## 启动

```bash
docker compose up -d
```

## 注意

- 生产环境必须修改默认密码：数据库、Redis、管理员密码、JWT 密钥、2FA 密钥。
- 原始笔记里 PostgreSQL 用户名与健康检查用户不一致，这里已整理为 `sub2api`。
- `SECURITY_URL_ALLOWLIST_ALLOW_INSECURE_HTTP=true` 允许 HTTP URL，公网环境谨慎使用。
