# Qdrant 向量数据库部署

Qdrant 是一个轻量、稳定、易部署的向量数据库，适合个人知识库、RAG、语义搜索和中小型 AI 项目。

## 文件

- `docker-compose.yml`：Qdrant 可运行配置
- `README.md`：部署说明

## 数据目录

按老王的 Docker 数据卷约定，持久化数据放在：

```text
D:\work\docker\qdrant\storage
```

WSL 路径：

```text
/mnt/d/work/docker/qdrant/storage
```

## 端口

| 端口 | 用途 |
|---|---|
| `6333` | HTTP API / Dashboard |
| `6334` | gRPC API |

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

## 访问 Dashboard

本机访问：

```text
http://localhost:6333/dashboard
```

局域网访问：

```text
http://192.168.1.139:6333/dashboard
```

## API 地址

```text
http://192.168.1.139:6333
```

## 账号密码

默认没有账号密码。

如果要暴露到公网，必须配置 API Key、反向代理认证或防火墙限制。

## Python 示例

```bash
pip install qdrant-client
```

```python
from qdrant_client import QdrantClient

client = QdrantClient(url="http://192.168.1.139:6333")
print(client.get_collections())
```

## 注意事项

- Qdrant 默认没有鉴权，只建议局域网使用。
- 生产环境需要启用 API Key 或放到内网。
- 数据目录不要随便删除，否则 collection 和向量数据会丢失。
