---
title: CosyVoice 部署与听小说项目教程
date: 2026-06-05
tags:
  - AI
  - TTS
  - CosyVoice
  - Docker
  - 听小说
  - 语音合成
source:
  - https://github.com/FunAudioLLM/CosyVoice
  - https://github.com/FunAudioLLM/CosyVoice/blob/main/docker/Dockerfile
  - https://github.com/FunAudioLLM/CosyVoice/tree/main/runtime/python
  - https://github.com/FunAudioLLM/CosyVoice/tree/main/runtime/triton_trtllm
  - https://github.com/jianchang512/cosyvoice-api
---

# CosyVoice 部署与听小说项目教程

> 目标：为“听小说 / 有声小说生成”项目选择并部署开源 TTS 服务，重点关注 CosyVoice 的 Docker/API 部署、OpenAI 兼容接口、长文本切分、章节音频生成和工程落地方案。

## 结论先行

如果项目目标是 **中文小说朗读 / 听书 App / 章节批量转音频**，推荐优先使用：

```text
正式章节生成：CosyVoice2 / Fun-CosyVoice3
快速试听或低成本预览：Kokoro / Piper，可作为辅助
高质量角色音色克隆：后续可引入 GPT-SoVITS
```

对于当前项目，建议路线：

```text
MVP：CosyVoice2-0.5B + FastAPI + 章节离线生成
进阶：Fun-CosyVoice3 + 分角色音色 + 队列任务
高性能：Triton TensorRT-LLM / vLLM / GPU 多实例
```

## CosyVoice 是什么

CosyVoice 是 FunAudioLLM 团队开源的多语言大模型 TTS 项目，提供文本转语音、零样本音色克隆、跨语言合成、流式合成、训练与部署能力。

官方仓库：

```text
https://github.com/FunAudioLLM/CosyVoice
```

当前官方资料中主要有三代：

| 版本 | 模型 | 说明 |
| --- | --- | --- |
| CosyVoice 1.0 | CosyVoice-300M / SFT / Instruct | 资料多，部署简单，适合入门 |
| CosyVoice 2.0 | CosyVoice2-0.5B | 更适合流式 TTS，质量和能力更好 |
| Fun-CosyVoice 3.0 | Fun-CosyVoice3-0.5B-2512 | 官方推荐的新一代，内容一致性、音色相似度、韵律自然度更强 |

官方 README 中提到 Fun-CosyVoice3 支持：

- 9 种常见语言：中文、英文、日语、韩语、德语、西语、法语、意大利语、俄语
- 18+ 中文方言/口音
- 多语言 / 跨语言 zero-shot voice cloning
- 拼音与英文音素级 pronunciation inpainting
- 数字、符号、复杂文本格式归一化
- 双流式：text-in streaming + audio-out streaming
- 最低约 150ms 延迟的流式输出
- 指令控制：语言、方言、情绪、语速、音量等

## 为什么适合听小说项目

听小说项目对 TTS 的要求通常是：

1. **中文自然度要够好**
2. **长文本稳定**
3. **支持批量生成章节音频**
4. **可以做角色音色区分**
5. **可以 API 化部署**
6. **最好支持流式或准实时试听**
7. **部署成本可控**

CosyVoice 的优势：

- 中文效果比传统轻量 TTS 更自然。
- 支持 zero-shot 音色克隆，可做旁白、男声、女声、角色声。
- 官方提供 FastAPI / gRPC 部署示例。
- 支持 Docker 部署。
- 后续可用 TensorRT-LLM / Triton 做加速。
- 社区资料多，适合工程落地。

不足：

- 比 Piper / Kokoro 更重，最好有 NVIDIA GPU。
- 首次部署依赖较多：CUDA、PyTorch、Conda、模型文件、submodule。
- 小说长文本需要自己做分段、任务队列、音频拼接和缓存。
- 真正生产级 OpenAI 兼容接口可能需要自己封装或采用社区项目。

## 推荐部署架构

### MVP 架构

适合先跑通项目：

```text
前端 / 管理后台
  ↓
小说文本导入
  ↓
章节切分
  ↓
段落切分
  ↓
TTS API 任务队列
  ↓
CosyVoice FastAPI 服务
  ↓
生成 wav
  ↓
转 mp3 / m4a
  ↓
对象存储 / 本地文件
  ↓
播放器按章节播放
```

### 推荐 Docker 数据目录

按老王当前 Docker 数据卷约定，建议放在：

```text
Windows: D:\work\docker\cosyvoice
WSL:     /mnt/d/work/docker/cosyvoice
```

建议目录：

```text
/mnt/d/work/docker/cosyvoice/
├── models/              # 模型文件
├── input/               # 原始小说文本
├── output/              # 生成后的音频
├── cache/               # 分段缓存
├── voices/              # 参考音频 / 角色音色
├── logs/                # 日志
└── compose/             # docker-compose 文件
```

## 方案选择

### 方案 A：官方 FastAPI 部署

适合：想尽量贴近官方，少依赖第三方封装。

官方 README 给出的部署路径：

```bash
cd runtime/python
docker build -t cosyvoice:v1.0 .
```

FastAPI 启动示例：

```bash
docker run -d --runtime=nvidia \
  -p 50000:50000 \
  cosyvoice:v1.0 \
  /bin/bash -c "cd /opt/CosyVoice/CosyVoice/runtime/python/fastapi && python3 server.py --port 50000 --model_dir iic/CosyVoice-300M && sleep infinity"
```

客户端测试：

```bash
cd runtime/python/fastapi
python3 client.py --port 50000 --mode <sft|zero_shot|cross_lingual|instruct>
```

可选模式：

```text
sft
zero_shot
cross_lingual
instruct
```

优点：

- 官方支持。
- 与项目更新最一致。
- 同时有 gRPC / FastAPI 示例。

缺点：

- 接口不一定是 OpenAI 兼容格式。
- 对业务系统接入不如 `/v1/audio/speech` 方便。

### 方案 B：社区 OpenAI 兼容 API 封装

参考项目：

```text
https://github.com/jianchang512/cosyvoice-api
```

它提供的接口包括：

| 接口 | 用途 |
| --- | --- |
| `/tts` | 使用内置角色直接合成文字 |
| `/clone_eq` | 同语言音色克隆 |
| `/cone` | 跨语言音色克隆，项目 README 中接口名疑似为 cone |
| `/v1/audio/speech` | OpenAI-compatible TTS 接口 |

OpenAI 兼容接口示例：

```python
from openai import OpenAI

client = OpenAI(api_key="dummy", base_url="http://127.0.0.1:9933/v1")

with client.audio.speech.with_streaming_response.create(
    model="tts-1",
    voice="中文女",
    input="你好啊，亲爱的朋友们。",
    speed=1.0,
) as response:
    with open("test.wav", "wb") as f:
        for chunk in response.iter_bytes():
            f.write(chunk)
```

内置角色示例：

```text
中文女
中文男
日语男
粤语女
英文女
英文男
韩语女
```

优点：

- 很适合接入现有 OpenAI TTS 客户端。
- 听小说项目后端可以统一用 OpenAI SDK 调用。
- API 简洁。

缺点：

- 第三方封装，需自行检查代码质量和维护状态。
- 生产环境建议自己 fork 后固定版本。
- 需要确认是否支持 CosyVoice2 / Fun-CosyVoice3 的最新能力。

### 方案 C：Triton TensorRT-LLM 高性能部署

官方 README 提到：使用 TensorRT-LLM 加速 CosyVoice2 LLM，相比 HuggingFace Transformers 实现可有约 4x 加速。

启动方式：

```bash
cd runtime/triton_trtllm
docker compose up -d
```

适合：

- 有 NVIDIA GPU。
- 需要批量生成大量章节。
- 需要更高并发。
- 后期做生产化部署。

不适合：

- MVP 初期。
- 没有 GPU 的环境。
- 不熟悉 Triton / TensorRT-LLM 的个人项目。

## 本地 Conda 安装流程

适合调试和开发，不一定适合生产部署。

### 1. 克隆仓库

```bash
git clone --recursive https://github.com/FunAudioLLM/CosyVoice.git
cd CosyVoice
```

如果 submodule 拉取失败：

```bash
git submodule update --init --recursive
```

### 2. 创建环境

```bash
conda create -n cosyvoice -y python=3.10
conda activate cosyvoice
```

### 3. 安装依赖

国内环境可使用阿里云 PyPI 镜像：

```bash
pip install -r requirements.txt \
  -i https://mirrors.aliyun.com/pypi/simple/ \
  --trusted-host=mirrors.aliyun.com
```

如果遇到 sox 问题：

Ubuntu：

```bash
sudo apt-get install sox libsox-dev
```

CentOS：

```bash
sudo yum install sox sox-devel
```

## 模型下载

官方推荐下载：

- `FunAudioLLM/Fun-CosyVoice3-0.5B-2512`
- `iic/CosyVoice2-0.5B`
- `iic/CosyVoice-300M`
- `iic/CosyVoice-300M-SFT`
- `iic/CosyVoice-300M-Instruct`
- `iic/CosyVoice-ttsfrd`

### ModelScope 下载

国内优先：

```python
from modelscope import snapshot_download

snapshot_download('FunAudioLLM/Fun-CosyVoice3-0.5B-2512', local_dir='pretrained_models/Fun-CosyVoice3-0.5B')
snapshot_download('iic/CosyVoice2-0.5B', local_dir='pretrained_models/CosyVoice2-0.5B')
snapshot_download('iic/CosyVoice-300M', local_dir='pretrained_models/CosyVoice-300M')
snapshot_download('iic/CosyVoice-300M-SFT', local_dir='pretrained_models/CosyVoice-300M-SFT')
snapshot_download('iic/CosyVoice-300M-Instruct', local_dir='pretrained_models/CosyVoice-300M-Instruct')
snapshot_download('iic/CosyVoice-ttsfrd', local_dir='pretrained_models/CosyVoice-ttsfrd')
```

### HuggingFace 下载

海外环境：

```python
from huggingface_hub import snapshot_download

snapshot_download('FunAudioLLM/Fun-CosyVoice3-0.5B-2512', local_dir='pretrained_models/Fun-CosyVoice3-0.5B')
snapshot_download('FunAudioLLM/CosyVoice2-0.5B', local_dir='pretrained_models/CosyVoice2-0.5B')
snapshot_download('FunAudioLLM/CosyVoice-300M', local_dir='pretrained_models/CosyVoice-300M')
snapshot_download('FunAudioLLM/CosyVoice-300M-SFT', local_dir='pretrained_models/CosyVoice-300M-SFT')
snapshot_download('FunAudioLLM/CosyVoice-300M-Instruct', local_dir='pretrained_models/CosyVoice-300M-Instruct')
snapshot_download('FunAudioLLM/CosyVoice-ttsfrd', local_dir='pretrained_models/CosyVoice-ttsfrd')
```

## ttsfrd 文本归一化

官方说明：`ttsfrd` 可选，不安装时默认使用 WeTextProcessing。

如果需要更好的文本归一化能力，可安装：

```bash
cd pretrained_models/CosyVoice-ttsfrd/
unzip resource.zip -d .
pip install ttsfrd_dependency-0.1-py3-none-any.whl
pip install ttsfrd-0.4.2-cp310-cp310-linux_x86_64.whl
```

听小说项目建议：

- MVP：可以先不装，降低部署复杂度。
- 正式生产：建议测试安装，特别是小说里有大量数字、时间、金额、符号时。

## WebUI 体验

先用 WebUI 体验模型效果：

```bash
python3 webui.py --port 50000 --model_dir pretrained_models/CosyVoice-300M
```

如果要测试其他模型，修改 `--model_dir`：

```bash
python3 webui.py --port 50000 --model_dir pretrained_models/CosyVoice2-0.5B
python3 webui.py --port 50000 --model_dir pretrained_models/Fun-CosyVoice3-0.5B
```

## Docker 部署建议

官方 Dockerfile 基于：

```text
nvidia/cuda:12.4.1-cudnn-devel-ubuntu22.04
```

包含：

- git
- build-essential
- curl / wget
- ffmpeg
- unzip
- git-lfs
- sox / libsox-dev
- Miniforge / Conda
- Python 3.10
- CosyVoice repo
- requirements.txt

### 注意

官方 Dockerfile 会在镜像构建阶段直接 clone CosyVoice 仓库。生产环境更建议：

1. fork 或固定 commit。
2. 将模型目录通过 volume 挂载。
3. 不要每次 build 都重新下载大模型。
4. 将输入、输出、缓存目录挂载到 `/mnt/d/work/docker/cosyvoice`。

### Docker Compose 示例：官方 FastAPI 风格

> 这是整理后的参考配置，需要根据实际 server.py 参数和模型路径调整。

```yaml
services:
  cosyvoice-api:
    image: cosyvoice:v1.0
    container_name: cosyvoice-api
    restart: unless-stopped
    ports:
      - "50000:50000"
    volumes:
      - /mnt/d/work/docker/cosyvoice/models:/opt/CosyVoice/CosyVoice/pretrained_models
      - /mnt/d/work/docker/cosyvoice/output:/data/output
      - /mnt/d/work/docker/cosyvoice/voices:/data/voices
      - /mnt/d/work/docker/cosyvoice/logs:/data/logs
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: all
              capabilities: [gpu]
    command: >
      /bin/bash -lc "cd /opt/CosyVoice/CosyVoice/runtime/python/fastapi &&
      python3 server.py --port 50000 --model_dir /opt/CosyVoice/CosyVoice/pretrained_models/CosyVoice2-0.5B"
```

如果 Docker 版本不支持 `deploy.resources`，可改用：

```yaml
    runtime: nvidia
```

或运行时使用：

```bash
docker run --gpus all ...
```

## 小说长文本切分策略

TTS 不能直接把整章几十万字丢进去。建议流程：

```text
原始章节
→ 清洗文本
→ 按段落切分
→ 过长段落再按标点切分
→ 每段生成音频
→ 音频拼接
→ 生成章节音频
→ 保存元数据
```

### 推荐切分规则

中文小说建议：

- 单段控制在 80～200 个汉字左右。
- 对话句尽量不要拆断。
- 优先按自然段切。
- 自然段过长时按标点切：
  - `。`
  - `！`
  - `？`
  - `；`
  - `……`
- 避免在引号中间切断。
- 每段之间插入 300～800ms 静音。
- 章节标题后插入 800～1200ms 静音。

### 文本清洗

需要处理：

- 多余空白
- HTML 标签
- 章节广告
- 作者的话
- 重复标题
- 特殊符号
- Markdown 标记
- 乱码
- 数字读法

建议保留：

- 中文引号
- 破折号
- 省略号
- 问号 / 感叹号

这些对语气有帮助。

## 角色音色设计

听小说项目可以先做两种模式。

### 简单模式

```text
旁白：中文女 / 中文男
所有角色：同一个音色
```

优点：稳定、简单、成本低。

### 进阶模式

```text
旁白：中性旁白音
男主：男声
女主：女声
老人：低沉慢速
小孩：高音色
反派：低沉或冷感
```

需要先做：

1. 角色识别
2. 对话归属判断
3. 角色到 voice/profile 的映射
4. 分段调用 TTS
5. 拼接音频

建议 MVP 不要一开始做自动角色识别，容易复杂化。可以先支持手动配置：

```json
{
  "narrator": "中文女",
  "default_male": "中文男",
  "default_female": "中文女",
  "角色A": "voice_001",
  "角色B": "voice_002"
}
```

## 后端任务表设计

### book

```text
id
name
author
source_path
status
created_at
updated_at
```

### chapter

```text
id
book_id
chapter_index
title
text_path
audio_path
status
word_count
duration_ms
created_at
updated_at
```

### tts_segment

```text
id
chapter_id
segment_index
speaker
text
audio_path
status
retry_count
error
created_at
updated_at
```

### voice_profile

```text
id
name
type
reference_audio_path
reference_text
provider
params_json
created_at
updated_at
```

## 章节生成流程

```text
1. 上传 / 导入小说 txt
2. 解析章节
3. 保存 chapter text
4. 每章切成 segments
5. segment 入队
6. worker 调用 CosyVoice API
7. 保存 segment wav
8. ffmpeg 拼接 segment wav
9. 转 mp3 / m4a
10. 写入 chapter.audio_path
11. 前端播放器读取章节音频
```

## 音频格式建议

生成阶段：

```text
wav
```

存储 / 分发阶段：

```text
mp3 或 m4a/aac
```

原因：

- wav 便于拼接和调试。
- mp3/m4a 文件更小，适合 Web/App 播放。

ffmpeg 转码示例：

```bash
ffmpeg -i chapter.wav -codec:a libmp3lame -b:a 128k chapter.mp3
```

拼接示例：

```bash
ffmpeg -f concat -safe 0 -i segments.txt -c copy chapter.wav
```

`segments.txt` 示例：

```text
file '/data/output/chapter_001/0001.wav'
file '/data/output/chapter_001/silence_500ms.wav'
file '/data/output/chapter_001/0002.wav'
```

## API 调用设计

推荐后端封装统一接口：

```http
POST /api/tts
Content-Type: application/json

{
  "text": "要合成的文本",
  "voice": "中文女",
  "speed": 1.0,
  "format": "wav",
  "speaker": "narrator"
}
```

内部再转发给：

- 官方 CosyVoice FastAPI
- 或 OpenAI-compatible `/v1/audio/speech`

这样以后替换 Kokoro、Piper、GPT-SoVITS、商业 TTS 都不影响业务系统。

## OpenAI 兼容接口建议

如果采用 `/v1/audio/speech`，业务侧可以统一：

```python
from openai import OpenAI

client = OpenAI(
    api_key="dummy",
    base_url="http://cosyvoice-api:9933/v1",
)

response = client.audio.speech.create(
    model="tts-1",
    voice="中文女",
    input="这一章，从一个雨夜开始。",
    speed=1.0,
)

response.stream_to_file("segment.wav")
```

优点：

- 以后换 OpenAI TTS / Azure / 自建 Kokoro 兼容服务更容易。
- 客户端 SDK 成熟。
- 服务边界清晰。

## 性能建议

### MVP

- 单 GPU 单 worker。
- 每次处理一个 segment。
- 先离线生成整章，不追求实时。
- 前端只播放已生成好的 mp3。

### 中期

- 增加任务队列：Redis + Celery / RQ / Dramatiq。
- segment 级缓存：相同文本 + voice + speed 不重复生成。
- 多 worker 并发。
- 每章生成进度可视化。

### 后期

- TensorRT-LLM / Triton。
- 多 GPU。
- 热门章节预生成。
- 新章节增量生成。
- 流式试听。
- 语音质量打分和重试机制。

## 缓存 Key 设计

```text
sha256(model + voice + speed + text + normalize_version)
```

缓存路径：

```text
/mnt/d/work/docker/cosyvoice/cache/{hash}.wav
```

如果文本、音色、语速、模型没变，就直接复用。

## 常见坑

### 1. CUDA / PyTorch / 驱动不匹配

表现：容器能启动，但推理报 CUDA 错误。

建议：

- 确认宿主机 NVIDIA 驱动正常。
- 确认 Docker 支持 GPU：

```bash
docker run --rm --gpus all nvidia/cuda:12.4.1-base-ubuntu22.04 nvidia-smi
```

### 2. 模型下载慢或失败

国内优先 ModelScope。

建议提前下载到：

```text
/mnt/d/work/docker/cosyvoice/models
```

再挂载进容器。

### 3. submodule 没拉全

如果源码方式部署，必须执行：

```bash
git submodule update --init --recursive
```

### 4. 长文本一次性请求失败

不要整章直接请求。必须切段。

### 5. 音频拼接有爆音或停顿怪

建议：

- 每段前后加短静音。
- 统一采样率、声道、编码。
- 拼接前先转成统一 wav。

### 6. 多角色朗读效果不稳定

先不要自动识别角色。MVP 可以只做旁白单音色。

### 7. 数字、金额、符号读法不自然

可考虑安装 `ttsfrd`，或者在进入 TTS 前自行做文本归一化。

## 推荐 MVP 实施清单

第一阶段：跑通服务

- [ ] 准备 NVIDIA GPU 环境
- [ ] 拉取 CosyVoice 官方仓库
- [ ] 下载 CosyVoice2-0.5B 或 Fun-CosyVoice3-0.5B
- [ ] 启动 WebUI 测试音质
- [ ] 启动 FastAPI 服务
- [ ] 用短文本生成 wav

第二阶段：小说生成

- [ ] 导入 txt 小说
- [ ] 解析章节
- [ ] 按段落/标点切分
- [ ] 调 TTS 生成 segment wav
- [ ] ffmpeg 拼接章节 wav
- [ ] 转 mp3
- [ ] 前端播放

第三阶段：产品化

- [ ] 任务队列
- [ ] 生成进度
- [ ] 缓存
- [ ] 失败重试
- [ ] 音色配置
- [ ] 批量章节生成
- [ ] 后台管理界面

第四阶段：精品化

- [ ] 多角色音色
- [ ] 参考音频管理
- [ ] 情绪 / 语速控制
- [ ] 章节试听与重生成
- [ ] 热门书籍预生成

## 推荐技术栈

后端：

```text
Python FastAPI / Node.js NestJS / Go Gin 均可
```

如果要快速做：

```text
FastAPI + Redis + Celery/RQ + PostgreSQL + MinIO/本地文件
```

前端：

```text
Vue / React + audio player
```

音频处理：

```text
ffmpeg
pydub
```

存储：

```text
本地文件：/mnt/d/work/docker/cosyvoice/output
后期可迁移：MinIO / S3
```

## 一个实际推荐架构

```text
novel-tts-api
├── FastAPI 后端
├── Redis 队列
├── Worker
├── PostgreSQL
├── CosyVoice API client
├── ffmpeg audio pipeline
└── local/S3 storage

cosyvoice-api
└── GPU TTS 服务

web-admin
└── 上传小说、查看章节、生成进度、播放音频
```

## 参考链接

- CosyVoice 官方仓库：https://github.com/FunAudioLLM/CosyVoice
- 官方 Dockerfile：https://github.com/FunAudioLLM/CosyVoice/blob/main/docker/Dockerfile
- 官方 Python runtime：https://github.com/FunAudioLLM/CosyVoice/tree/main/runtime/python
- Triton TensorRT-LLM 部署：https://github.com/FunAudioLLM/CosyVoice/tree/main/runtime/triton_trtllm
- 社区 OpenAI 兼容 API 封装：https://github.com/jianchang512/cosyvoice-api
- CosyVoice2 ModelScope：https://www.modelscope.cn/models/iic/CosyVoice2-0.5B
- CosyVoice2 HuggingFace：https://huggingface.co/FunAudioLLM/CosyVoice2-0.5B
- Fun-CosyVoice3 ModelScope：https://www.modelscope.cn/models/FunAudioLLM/Fun-CosyVoice3-0.5B-2512
- Fun-CosyVoice3 HuggingFace：https://huggingface.co/FunAudioLLM/Fun-CosyVoice3-0.5B-2512
