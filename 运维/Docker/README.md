# Docker 运维知识库

本目录统一收纳 Docker / Docker Compose / 容器化部署相关内容。

## 目录结构

```text
Docker/
├── README.md
├── 安装/
│   ├── CentOS 7.9 安装 Docker.md
│   └── Ubuntu 24.04 安装 Docker.md
├── Compose/
│   ├── MySQL.md
│   ├── Nginx Proxy Manager.md
│   └── Sub2API.md
├── Qdrant/
│   ├── README.md
│   └── docker-compose.yml
├── MinIO/
│   ├── README.md
│   └── docker-compose.yml
├── Milvus项目/
│   ├── README.md
│   └── docker-compose.yml
└── _原始归档/
```

## 安装指南

- [[安装/CentOS 7.9 安装 Docker]]
- [[安装/Ubuntu 24.04 安装 Docker]]

## Compose 模板

- [[Compose/MySQL]]
- [[Compose/Nginx Proxy Manager]]
- [[Compose/Sub2API]]

## 项目部署

- [[运维/Docker/Compose/Qdrant/README|Qdrant 向量数据库部署]]
- [[运维/Docker/Compose/MinIO/README|MinIO 对象存储部署]]
- [[运维/Docker/Compose/Milvus项目/README|Milvus + Attu 单机部署]]

## 常用访问地址

- Qdrant Dashboard：`http://192.168.1.139:6333/dashboard`
- Qdrant API：`http://192.168.1.139:6333`
- MinIO 控制台：`http://192.168.1.139:19001`
- MinIO S3 API：`http://192.168.1.139:19000`

## 维护规则

- 新增 Docker 内容统一放到 `运维/Docker/`。
- 安装教程放 `安装/`。
- 通用服务模板放 `Compose/`。
- 独立项目部署放项目目录。
- Docker 持久化数据优先放到 `D:\work\docker`，WSL 路径是 `/mnt/d/work/docker`。
- 真实账号、密码、Token 不写入知识库；统一使用占位符。
- 可运行配置和解释文档分开保存：
  - `docker-compose.yml` 保存可运行配置
  - `README.md` 保存说明
