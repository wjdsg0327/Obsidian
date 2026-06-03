# Milvus + Attu 单机部署

这个目录只保留 Milvus 相关的核心内容：

- `README.md`：部署说明、访问方式、注意事项
- `docker-compose.yml`：可直接运行的 Compose 配置

## 服务组件

| 服务 | 作用 | 端口 |
|---|---|---|
| `etcd` | Milvus 元数据存储 | 容器内使用 |
| `minio` | 对象存储 | `9000` API / `9001` 控制台 |
| `standalone` | Milvus 单机服务 | `19530` / `9091` 健康检查 |
| `attu` | Milvus 可视化管理界面 | `8000` |

## 启动

在当前目录执行：

```bash
docker compose up -d
```

只启动 Attu：

```bash
docker compose up -d attu
```

查看状态：

```bash
docker compose ps
```

停止：

```bash
docker compose down
```

## 访问

### Attu 可视化界面

```text
http://服务器IP:8000
```

Attu 连接 Milvus 地址：

```text
milvus-standalone:19530
```

因为 Attu 和 Milvus 在同一个 Docker 网络里，所以用容器名连接即可。

### Milvus

```text
服务器IP:19530
```

### MinIO

```text
http://服务器IP:9001
```

默认用户名：

```text
minioadmin
```

默认密码在 compose 中使用占位值：

```text
change_me_minio_password
```

生产环境必须修改。

## 数据目录

默认数据保存在当前目录：

```text
volumes/etcd
volumes/minio
volumes/milvus
```

也可以通过环境变量指定：

```bash
DOCKER_VOLUME_DIRECTORY=/data/milvus docker compose up -d
```

## 当前镜像

- etcd：`quay.io/coreos/etcd:v3.5.25`
- MinIO：`docker.1panel.live/minio/minio:latest`
- Milvus：`docker.1panel.live/milvusdb/milvus:v2.6.17`
- Attu：`docker.1panel.live/zilliz/attu:latest`

## 当前关键配置

- 数据存储使用 Docker 命名卷：`milvus_etcd`、`milvus_minio`、`milvus_data`
- 单机 MQ 使用：`MQ_TYPE=rocksmq`

## 注意事项

- 这是单机版 Milvus，适合学习、测试、小型项目。
- 生产环境建议评估 Milvus 集群部署。
- `MINIO_ROOT_PASSWORD` 不要使用默认值。
- `standalone` 使用了 `seccomp:unconfined`，生产环境需要重新评估安全策略。
- 如果 Attu 镜像拉取失败，优先检查 Docker Desktop 代理或镜像源。
- 如果 Milvus 健康检查长期返回 `Not all components are healthy`，优先检查 MQ 配置和数据卷位置；WSL/Windows 盘挂载目录不适合直接承载 Milvus/etcd 的 WAL 数据。

## 文件维护规则

Milvus 相关知识不再拆多个说明文件，统一维护在本 README 中。

需要修改部署时，只改：

```text
README.md
docker-compose.yml
```
