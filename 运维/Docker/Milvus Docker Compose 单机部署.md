# Milvus Docker Compose 单机部署

## 用途

这个项目用于启动 Milvus 单机版环境，适合本地测试、学习、轻量开发环境。

## 服务组件

- `etcd`：Milvus 元数据存储
- `minio`：对象存储，保存向量索引、日志等数据
- `standalone`：Milvus 单机服务

## 端口

- Milvus：`19530`
- Milvus 健康检查：`9091`
- MinIO API：`9000`
- MinIO 控制台：`9001`

## 数据目录

默认会在当前目录下创建：

```text
volumes/etcd
volumes/minio
volumes/milvus
```

也可以通过环境变量 `DOCKER_VOLUME_DIRECTORY` 指定数据目录。

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

## 访问

- Milvus SDK / 客户端连接：`localhost:19530`
- MinIO 控制台：`http://localhost:9001`

## 注意事项

- `MINIO_ROOT_PASSWORD` 已从原始默认值修正为 `change_me_minio_password`，使用前请改成自己的强密码。
- 当前镜像使用 `docker.1panel.live` 镜像源。
- `standalone` 服务使用 `seccomp:unconfined`，如果生产环境使用，需要结合安全策略重新评估。
- `depends_on` 已核对为 `etcd` 和 `minio`。
- 生产环境建议使用 Milvus 集群部署，而不是 standalone 单机模式。

## Compose 文件

正式可运行配置见：

```text
./docker-compose.yml
```
