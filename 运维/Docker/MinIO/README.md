# MinIO Docker 部署

MinIO 是 S3 兼容对象存储，可以用来存图片、文件、模型数据、向量数据库对象数据等。

## 数据目录

按老王的 Docker 数据卷约定，持久化数据放在：

```text
D:\work\docker\minio\data
```

WSL 路径：

```text
/mnt/d/work/docker/minio/data
```

## 端口

为了避免和 Milvus 里的 MinIO 冲突，独立 MinIO 使用外部端口：

| 外部端口 | 容器端口 | 用途 |
|---|---|---|
| `19000` | `9000` | S3 API |
| `19001` | `9001` | Web 控制台 |

## 访问

Web 控制台：

```text
http://192.168.1.139:19001
```

S3 API：

```text
http://192.168.1.139:19000
```

## 默认账号

```text
minioadmin
```

## 默认密码

```text
change_me_minio_password_123
```

首次部署后建议尽快修改密码，生产环境不要使用默认密码。

## 启动

```bash
docker compose up -d
```

## 查看状态

```bash
docker compose ps
```

## 停止

```bash
docker compose down
```
