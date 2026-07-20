---
title: chromem-go：纯 Go 嵌入式向量数据库与 RAG
date: 2026-07-20
source: https://mp.weixin.qq.com/s/1Kh0JarDF22_6zGfCPMGag
author: Go语言中文网
tags:
  - Go
  - RAG
  - 向量数据库
  - AI
  - 知识库
---

# chromem-go：纯 Go 嵌入式向量数据库与 RAG

> 来源文章：[chromem-go：纯 Go 嵌入式向量数据库，像 SQLite 一样给应用加上 RAG](https://mp.weixin.qq.com/s/1Kh0JarDF22_6zGfCPMGag)  
> 来源：Go语言中文网  
> 整理时间：2026-07-20

## 一句话结论

**chromem-go 是一个纯 Go、零依赖、可嵌入进应用进程的向量数据库。**  
它的定位类似“向量数据库领域的 SQLite”：适合给 Go 应用快速加上 RAG、语义搜索、知识库问答能力，而不用额外部署 Qdrant、Milvus、PostgreSQL/pgvector 等服务。

## 它解决什么问题

传统向量数据库通常是客户端-服务器架构：

```text
Go 应用 → HTTP/gRPC → 独立向量数据库服务
```

这会带来：

- 额外部署和运维成本
- 网络调用延迟
- 容器 / 数据库 / 进程管理复杂度
- 小项目或本地工具场景下显得过重

chromem-go 的方式是：

```go
db := chromem.NewDB()
collection, _ := db.CreateCollection("docs", nil, nil)
```

也就是直接在 Go 进程内创建和查询向量库。

## 核心特点

### 1. 纯 Go、零依赖

- 不依赖第三方 Go 包
- 无 CGO
- 不需要 C++/Python 绑定
- 不需要外部数据库服务
- 适合交叉编译、单二进制部署、轻量 Docker 镜像

### 2. 嵌入模型支持丰富

内置支持多种 Embedding Provider：

- OpenAI
- Azure OpenAI
- GCP Vertex AI
- Cohere
- Mistral
- Jina
- Ollama
- LocalAI
- 自定义 `chromem.EmbeddingFunc`

本地离线方案示例：

```go
embeddingFunc := chromem.NewEmbeddingFuncOllama(
    "nomic-embed-text",
    "http://localhost:11434",
)
collection, _ := db.CreateCollection("local-docs", nil, embeddingFunc)
```

### 3. 性能适合中小规模知识库

文章给出的基准大致是：

| 文档数量 | 查询耗时 |
|---:|---:|
| 100 | 0.09 ms |
| 1,000 | 0.52 ms |
| 5,000 | 2.1 ms |
| 25,000 | 9.9 ms |
| 100,000 | 39.6 ms |

结论：**10 万篇以内的知识库、FAQ、文档搜索、个人 AI 助手记忆系统基本够用。**

### 4. 支持持久化、导入导出

```go
// 持久化到磁盘
db, _ := chromem.NewDB("./data/vector.db")

// 导出备份
file, _ := os.Create("backup.gob.gz")
db.Export(file, encryptionKey)

// 恢复
file, _ := os.Open("backup.gob.gz")
db2 := chromem.NewDB()
db2.Import(file, encryptionKey)
```

### 5. 支持并发添加和过滤查询

```go
collection.AddDocuments(ctx, documents, runtime.NumCPU())
```

支持：

- 元数据过滤
- 内容包含过滤
- 批量添加
- 并发处理

## 最小使用流程

```bash
go get github.com/philippgille/chromem-go@latest
```

典型流程：

1. 创建数据库
2. 创建 collection
3. 添加文档
4. 自动生成 embedding
5. 根据问题做语义搜索
6. 把检索结果拼进 prompt
7. 发给 LLM 回答

示意：

```go
results, _ := collection.Query(
    ctx,
    "Go 语言有什么性能优化新特性？",
    3,
    nil,
    nil,
)
```

## 适用场景

适合：

- Go CLI 工具里的语义搜索
- 桌面应用内置知识库
- 小型 RAG 系统
- FAQ / 文档搜索
- 个人 AI 助手记忆系统
- 代码搜索原型
- 不想部署额外数据库的小项目

不太适合：

- 百万级以上文档检索
- 多服务共享同一向量数据库
- 高并发、多租户、生产级集中式向量服务
- 需要复杂索引、实时更新、集群能力的场景

## 与主流方案对比

| 方案 | 架构 | 优点 | 代价 |
|---|---|---|---|
| chromem-go | 嵌入式 | 纯 Go、零依赖、部署极简 | 更适合中小规模 |
| Chroma | 客户端-服务器 / Python 生态 | 上手简单，AI 生态常见 | Python 依赖 |
| Qdrant | 独立服务 | Rust 实现，性能强，生产友好 | 需要部署服务 |
| Milvus | 分布式向量数据库 | 大规模能力强 | 运维复杂 |
| pgvector | PostgreSQL 扩展 | 结合业务数据库方便 | 需要 PostgreSQL 运维 |

## 我的判断

如果目标是：

> 给一个 Go 项目快速加上 RAG / 本地知识库 / 语义搜索能力

那 chromem-go 很值得试。它的优势不是“最强”，而是 **轻、简单、Go 原生、无需额外服务**。

但如果后面数据规模涨到百万级、需要多服务共享、需要成熟索引和高可用，还是应该换 Qdrant / Milvus / pgvector 这类方案。

## 可关注的路线图

文章提到后续值得关注：

- SIMD 点积计算：提升相似度计算速度
- HNSW 索引：支持更大规模近似最近邻搜索
- IVFFlat 索引
- WAL 写前日志
- 图片 / 视频等多模态向量搜索

其中 **HNSW** 最关键；如果实现成熟，chromem-go 的适用范围会明显扩大。

## 相关链接

- GitHub：`github.com/philippgille/chromem-go`
- GoDoc：搜索 `chromem-go godoc`
- 示例方向：chromem-go + Ollama + nomic-embed-text 构建本地 RAG

## 关键词

`Go`、`chromem-go`、`RAG`、`Embedding`、`向量数据库`、`语义搜索`、`Ollama`、`本地知识库`
