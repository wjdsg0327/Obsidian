# MySQL Docker Compose

## 用途

使用 Docker Compose 启动 MySQL 8.0。

## 目录建议

```text
mysql/
├── docker-compose.yml
├── data/
├── conf/
└── log/
```

## docker-compose.yml

```yaml
services:
  mysql:
    image: mysql:8.0
    container_name: mysql_db
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword # 请替换为自己的密码
    ports:
      - "3306:3306"
    volumes:
      - ./data:/var/lib/mysql
      - ./conf:/etc/mysql/conf.d
      - ./log:/var/log/mysql
```

## 启动

```bash
docker compose up -d
```

## 注意

- `MYSQL_ROOT_PASSWORD` 不要在生产环境中使用弱密码。
- `data` 目录保存数据库数据，删除前必须确认。
