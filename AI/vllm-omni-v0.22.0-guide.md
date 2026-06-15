# vLLM-Omni v0.22.0 资料、部署方法与使用案例

> 说明：你提到的 `vLLM-Omnivo0.22.0`，经检索官方项目后，准确项目名应为 **vLLM-Omni**，版本为 **v0.22.0**。  
> 官方仓库：https://github.com/vllm-project/vllm-omni  
> Release：https://github.com/vllm-project/vllm-omni/releases/tag/v0.22.0

---

## 1. 项目简介

**vLLM-Omni** 是 vLLM 官方项目下的全模态推理与服务框架，目标是为文本、图像、视频、音频等多模态/全模态模型提供统一、高性能、低成本的部署能力。

它是在 vLLM 原有文本自回归推理能力基础上扩展的全模态框架，核心能力包括：

- **Omni-modality / 全模态处理**：支持文本、图像、视频、音频等输入输出。
- **非自回归架构支持**：扩展到 Diffusion Transformer、DiT、并行生成模型等。
- **异构输出**：支持从文本生成到图像、视频、音频/TTS 等多种输出形式。
- **多阶段 pipeline 抽象**：适合 Qwen3-Omni、Qwen3-TTS 等 Thinker/Talker/Code2Wav 或多阶段模型。
- **OpenAI-compatible API server**：可通过 OpenAI 风格 API 服务化。

官方定位语：

> Easy, fast, and cheap omni-modality model serving for everyone.

官方 README 中列出的典型支持模型包括：

- Qwen3-Omni / Qwen2.5-Omni
- Qwen3-TTS
- Qwen-Image / Qwen-Image-Edit / Qwen-Image-Layered
- Cosmos3
- GLM-TTS、Fish Speech、VoxCPM、MOSS-TTS、Hunyuan、Wan、Bagel 等模型族

---

## 2. v0.22.0 Release 亮点

vLLM-Omni `v0.22.0` 是一个与 **vLLM 0.22 release line 对齐** 的全模态世界模型版本。

官方 Release 摘要：

- 339 commits
- 124 contributors
- 52 位新贡献者
- 发布时间：2026-06-06

### 2.1 World Model 支持增强

v0.22.0 提供了 **NVIDIA Cosmos3** world model 的 Day-0 支持，覆盖：

- text
- image
- audio
- video
- action

相关能力包括：

- Cosmos3 base model
- sound generation
- action modality
- DreamZero integration
- OpenPI robot serving

### 2.2 量化与硬件覆盖扩大

新增或增强：

- Blackwell diffusion attention backends
- W4A16 / Intel AutoRound
- FP8 / INT8
- MXFP4 / MXFP8
- ModelOpt mixed FP8 / NVFP4
- batched ModelOpt FP8 serving
- ROCm AITER
- Intel XPU
- Ascend NPU

### 2.3 音频与 TTS 生产化增强

增强模型和能力：

- Qwen3-TTS
- Qwen3-Omni
- VoxCPM2
- Fish Speech S2 Pro
- OmniVoice
- async audio input
- custom voices
- ref-context cache
- high-concurrency serving

### 2.4 RL / veRL-Omni 集成

支持与 veRL-Omni 相关的强化学习集成，覆盖：

- Qwen-Image
- Bagel
- SD 3.5
- WAN 2.2

### 2.5 与 vLLM 0.22 对齐

包括：

- vLLM 0.21 / 0.22 rebase
- 依赖兼容性更新
- release image builds
- PyPI upload 支持

---

## 3. 安装与环境要求

官方 Quickstart 要求：

- OS：Linux
- Python：3.12
- vLLM 与 vLLM-Omni 的 major/minor 版本应保持一致

官方特别提示：

> vLLM-Omni `0.22.0` 应搭配 vLLM `0.22.0`。如果版本不一致，可能出现 `vllm` 命令无法正确处理 `--omni` flag 等问题。

---

## 4. Python / CUDA GPU 安装

推荐使用 `uv` 创建 Python 3.12 环境：

```bash
uv venv --python 3.12 --seed
source .venv/bin/activate
```

安装 vLLM：

```bash
uv pip install vllm==0.22.0 --torch-backend=auto
```

安装 vLLM-Omni：

```bash
uv pip install vllm-omni
```

如果需要运行 Gradio demo：

```bash
uv pip install 'vllm-omni[demo]'
```

---

## 5. 从源码安装

```bash
uv venv --python 3.12 --seed
source .venv/bin/activate

uv pip install vllm==0.22.0 --torch-backend=auto

git clone https://github.com/vllm-project/vllm-omni.git
cd vllm-omni
git checkout v0.22.0
uv pip install -e .
```

如需 demo 依赖：

```bash
uv pip install -e '.[demo]'
```

---

## 6. ROCm 安装

官方 Quickstart 给出的 ROCm 安装方式：

```bash
uv pip install vllm==0.22.0+rocm721 \
  --extra-index-url https://wheels.vllm.ai/rocm/0.22.0/rocm721
```

随后安装 vLLM-Omni：

```bash
git clone https://github.com/vllm-project/vllm-omni.git
cd vllm-omni
git checkout v0.22.0
uv pip install -e .
```

---

## 7. Docker 部署

官方 CUDA 安装文档说明 vLLM-Omni 提供官方 Docker 镜像：

```text
vllm/vllm-omni:v0.22.0
```

官方文档示例是在 **2 x H100** 上验证的部署命令：

```bash
docker run --runtime nvidia --gpus 2 \
    -v ~/.cache/huggingface:/root/.cache/huggingface \
    --env "HF_TOKEN=***" \
    -p 8091:8091 \
    --ipc=host \
    vllm/vllm-omni:v0.22.0 \
    vllm serve Qwen/Qwen3-Omni-30B-A3B-Instruct --omni --port 8091
```

说明：

- CUDA image 没有默认 entrypoint，因此镜像名后要加 `vllm serve ... --omni`。
- 如果模型需要 Hugging Face 授权，需传入 `HF_TOKEN`。
- 建议挂载 Hugging Face cache，避免重复下载模型。
- 端口示例使用 `8091`。

更完整的推荐模板：

```bash
docker run -d --name vllm-omni \
  --restart unless-stopped \
  --runtime nvidia \
  --gpus all \
  --ipc=host \
  -p 8091:8091 \
  -e HF_TOKEN="你的HF_TOKEN" \
  -v ~/.cache/huggingface:/root/.cache/huggingface \
  vllm/vllm-omni:v0.22.0 \
  vllm serve Qwen/Qwen3-Omni-30B-A3B-Instruct --omni --port 8091
```

验证服务：

```bash
curl http://127.0.0.1:8091/v1/models
```

---

## 8. OpenAI-compatible API Server

vLLM-Omni 使用下面形式启动服务：

```bash
vllm serve <MODEL_ID> --omni --port 8091
```

示例：

```bash
vllm serve Qwen/Qwen3-Omni-30B-A3B-Instruct --omni --port 8091
```

支持的常见 API：

| API | 用途 |
|---|---|
| `/v1/chat/completions` | 多模态对话、图像生成、图像编辑、Qwen3-Omni 文本/音频输出 |
| `/v1/audio/speech` | TTS、streaming PCM、voice cloning、custom voice |
| `/v1/audio/voices` | voice list / upload / delete |
| `/v1/images/generations` | OpenAI-style 文生图 |
| `/v1/images/edits` | OpenAI-style 图像编辑 |
| `/v1/videos` | 异步视频生成 |
| `/v1/videos/sync` | 同步视频生成 |

### 8.1 Diffusion 参数传递方式

OpenAI Chat API schema 本身没有 `height`、`width`、`num_inference_steps` 等 diffusion 字段。

因此：

- 用 curl / requests 时，把参数放在 JSON 的 `extra_body` 字段。
- 用 OpenAI Python SDK 时，把参数作为 `extra_body` keyword argument 传入。

curl 示例：

```bash
curl -s http://localhost:8091/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "messages": [
      {"role": "user", "content": "A beautiful landscape painting"}
    ],
    "extra_body": {
      "num_inference_steps": 50,
      "seed": 42
    }
  }'
```

OpenAI Python SDK 示例：

```python
from openai import OpenAI

client = OpenAI(
    base_url="http://localhost:8091/v1",
    api_key="EMPTY",
)

response = client.chat.completions.create(
    model="Qwen/Qwen-Image",
    messages=[{"role": "user", "content": "A beautiful landscape painting"}],
    extra_body={
        "num_inference_steps": 50,
        "seed": 42,
    },
)

print(response)
```

---

## 9. 使用案例一：Qwen3-Omni

### 9.1 启动服务

```bash
vllm serve Qwen/Qwen3-Omni-30B-A3B-Instruct --omni --port 8091
```

官方说明：

- 默认部署配置位于：

```text
vllm_omni/deploy/qwen3_omni_moe.yaml
```

- 标准部署场景下会通过模型注册表自动解析加载，通常不需要显式传 `--deploy-config`。
- 默认启用 asynchronous chunk streaming。
- 默认设置 `VLLM_USE_FLASHINFER_MOE_FP16=0`，避免 FlashInfer CUTLASS unquantized MoE 路径上的性能回退。

如需显式指定自定义部署 YAML：

```bash
vllm serve Qwen/Qwen3-Omni-30B-A3B-Instruct --omni --port 8091 \
    --deploy-config /path/to/deploy_config_file
```

### 9.2 多阶段 CLI 部署

Stage 0：Thinker + API server

```bash
CUDA_VISIBLE_DEVICES=0 vllm serve Qwen/Qwen3-Omni-30B-A3B-Instruct --omni \
    --port 8091 \
    --stage-id 0 \
    --omni-master-address 127.0.0.1 \
    --omni-master-port 26000
```

Stage 1：Talker

```bash
CUDA_VISIBLE_DEVICES=1 vllm serve Qwen/Qwen3-Omni-30B-A3B-Instruct --omni \
    --stage-id 1 \
    --headless \
    --omni-master-address 127.0.0.1 \
    --omni-master-port 26000
```

Stage 2：Code2Wav

```bash
CUDA_VISIBLE_DEVICES=1 vllm serve Qwen/Qwen3-Omni-30B-A3B-Instruct --omni \
    --stage-id 2 \
    --headless \
    --omni-master-address 127.0.0.1 \
    --omni-master-port 26000
```

### 9.3 模态控制

Qwen3-Omni 支持通过 `modalities` 控制输出，例如：

- text only
- text + audio

具体请求格式见官方 Qwen3-Omni online serving 文档。

---

## 10. 使用案例二：Qwen3-TTS

### 10.1 支持模型

| Task Type | 模型 | 说明 |
|---|---|---|
| CustomVoice | `Qwen/Qwen3-TTS-12Hz-1.7B-CustomVoice` | 预置 speaker voices，可带 style/emotion control |
| VoiceDesign | `Qwen/Qwen3-TTS-12Hz-1.7B-VoiceDesign` | 根据自然语言描述生成语音风格 |
| Base | `Qwen/Qwen3-TTS-12Hz-1.7B-Base` | 基于参考音频与 transcript 的 voice cloning |
| CustomVoice 小模型 | `Qwen/Qwen3-TTS-12Hz-0.6B-CustomVoice` | 更小/更快 |
| Base 小模型 | `Qwen/Qwen3-TTS-12Hz-0.6B-Base` | 更小/更快的 voice cloning |

### 10.2 启动服务

```bash
vllm serve Qwen/Qwen3-TTS-12Hz-1.7B-CustomVoice --omni --port 8091
```

带 deploy config：

```bash
vllm serve Qwen/Qwen3-TTS-12Hz-1.7B-CustomVoice \
    --deploy-config vllm_omni/deploy/qwen3_tts.yaml \
    --omni --port 8091
```

VoiceDesign：

```bash
vllm serve Qwen/Qwen3-TTS-12Hz-1.7B-VoiceDesign \
    --deploy-config vllm_omni/deploy/qwen3_tts.yaml \
    --omni --port 8091
```

Base / voice cloning：

```bash
vllm serve Qwen/Qwen3-TTS-12Hz-1.7B-Base \
    --deploy-config vllm_omni/deploy/qwen3_tts.yaml \
    --omni --port 8091
```

### 10.3 生成语音

```bash
curl -X POST http://localhost:8091/v1/audio/speech \
    -H "Content-Type: application/json" \
    -d '{
        "input": "Hello, how are you?",
        "voice": "vivian",
        "language": "English"
    }' --output output.wav
```

带 emotion instruction：

```bash
curl -X POST http://localhost:8091/v1/audio/speech \
    -H "Content-Type: application/json" \
    -d '{
        "input": "I am so excited!",
        "voice": "vivian",
        "instructions": "Speak with great enthusiasm"
    }' --output excited.wav
```

列出 voices：

```bash
curl http://localhost:8091/v1/audio/voices
```

Streaming PCM：

```bash
curl -X POST http://localhost:8091/v1/audio/speech \
    -H "Content-Type: application/json" \
    -d '{
        "input": "Hello, how are you?",
        "voice": "vivian",
        "language": "English",
        "stream": true,
        "response_format": "pcm"
    }' --no-buffer | play -t raw -r 24000 -e signed -b 16 -c 1 -
```

说明：Qwen3-TTS 输出为 24 kHz；其他 TTS 模型采样率可能不同。

---

## 11. 使用案例三：Qwen-Image 文生图

### 11.1 启动服务

```bash
vllm serve Qwen/Qwen-Image --omni --port 8091
```

显存不足时可以尝试：

```text
--vae-use-slicing --vae-use-tiling
```

并行加速示例：

```bash
# Tensor Parallelism，需要 >= 2 GPUs
vllm serve Qwen/Qwen-Image --omni --port 8091 --tensor-parallel-size 2

# Tensor Parallelism + VAE Patch Parallelism，需要 >= 2 GPUs
vllm serve Qwen/Qwen-Image --omni --port 8091 --tensor-parallel-size 2 --vae-patch-parallel-size 2 --vae-use-tiling

# Ulysses sequence parallelism，需要 >= 2 GPUs
vllm serve Qwen/Qwen-Image --omni --port 8091 --usp 2

# Ring-Attention，需要 >= 2 GPUs
vllm serve Qwen/Qwen-Image --omni --port 8091 --ring 2

# Ulysses + Ring，需要 >= 4 GPUs
vllm serve Qwen/Qwen-Image --omni --port 8091 --usp 2 --ring 2
```

### 11.2 `/v1/chat/completions` 文生图请求

```bash
curl -s http://localhost:8091/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "messages": [
      {"role": "user", "content": "A beautiful landscape painting"}
    ],
    "extra_body": {
      "height": 1024,
      "width": 1024,
      "num_inference_steps": 50,
      "true_cfg_scale": 4.0,
      "seed": 42
    }
  }' | jq -r '.choices[0].message.content[0].image_url.url' | cut -d',' -f2- | base64 -d > output.png
```

### 11.3 `/v1/images/generations` 请求

```bash
curl http://localhost:8091/v1/images/generations \
  -H "Content-Type: application/json" \
  -d '{
    "model": "Qwen/Qwen-Image",
    "prompt": "A ceramic teapot on a wooden table",
    "size": "1024x1024",
    "num_inference_steps": 20,
    "seed": 42
  }'
```

---

## 12. 使用案例四：Qwen-Image-Edit 图像编辑

启动服务：

```bash
vllm serve Qwen/Qwen-Image-Edit --omni --port 8092
```

多图编辑模型：

```bash
vllm serve Qwen/Qwen-Image-Edit-2509 --omni --port 8092
```

图像编辑请求：

```bash
IMG_B64=$(base64 -w0 input.png)

cat <<EOF > request.json
{
  "messages": [{
    "role": "user",
    "content": [
      {"type": "text", "text": "Convert this image to watercolor style"},
      {"type": "image_url", "image_url": {"url": "data:image/png;base64,$IMG_B64"}}
    ]
  }],
  "extra_body": {
    "height": 1024,
    "width": 1024,
    "num_inference_steps": 50,
    "guidance_scale": 1,
    "seed": 42
  }
}
EOF

curl -s http://localhost:8092/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d @request.json \
  | jq -r '.choices[0].message.content[0].image_url.url' \
  | cut -d',' -f2 | base64 -d > output.png
```

官方 recipe 中记录的 `Qwen/Qwen-Image-Edit` 验证环境：

- 1x H200 141GB
- 2x H200 TP=2
- workload：1024×1024、50 steps、batch size 1、`guidance_scale=1.0`、seed 42
- 1x H200 warm latency 约 17.78s
- 2x H200 TP=2 warm latency 约 11.09s

注意：该性能数据来自官方 recipe，不代表所有环境。

---

## 13. 使用案例五：Cosmos3-Nano / Cosmos3-Super

v0.22.0 的一个重要新增能力是 Cosmos3 world model。

### 13.1 Cosmos3-Nano

模型：

```text
nvidia/Cosmos3-Nano
```

支持：

- Text-to-image
- Text-to-video
- Image-to-video
- Video + synchronized audio
- Action policy / Physical-AI tasks

启动服务：

```bash
vllm serve nvidia/Cosmos3-Nano \
  --omni \
  --host 0.0.0.0 --port 8000 \
  --init-timeout 1800
```

说明：

- Guardrails 默认开启。
- 默认会加载 gated `nvidia/Cosmos-1.0-Guardrail`。
- 如果保持 guardrails，需要：
  - 安装 `cosmos-guardrail`
  - 接受 Hugging Face license
  - 设置有访问权限的 `HF_TOKEN`

关闭 guardrails：

```bash
--no-guardrails
```

Text-to-image 请求：

```bash
curl -sS -X POST http://localhost:8000/v1/images/generations \
  -H "Content-Type: application/json" \
  -d '{
    "model": "nvidia/Cosmos3-Nano",
    "prompt": "A photorealistic red sports car on a city street at golden hour, cinematic lighting.",
    "negative_prompt": "blurry, distorted, low quality",
    "size": "1024x1024",
    "n": 1,
    "response_format": "b64_json",
    "num_inference_steps": 50,
    "guidance_scale": 7.0,
    "seed": 42
  }'
```

Text-to-video 请求：

```bash
curl -sS -X POST http://localhost:8000/v1/videos/sync \
  -H "Accept: video/mp4" \
  -F "model=nvidia/Cosmos3-Nano" \
  -F "prompt=A robot arm is cleaning a plate in the kitchen" \
  -F "negative_prompt=blurry, distorted, low quality, jittery, deformed" \
  -F "size=1280x720" \
  -F "num_frames=189" \
  -F "fps=24" \
  -F "num_inference_steps=35" \
  -F "guidance_scale=6.0" \
  -F "max_sequence_length=4096" \
  -F "flow_shift=10.0" \
  -F 'extra_params={"use_resolution_template":false,"use_duration_template":false,"guardrails":true}' \
  -F "seed=123" \
  -o cosmos3_t2v.mp4
```

### 13.2 Cosmos3-Super

模型：

```text
nvidia/Cosmos3-Super
```

说明：

- 64B world model
- 支持 T2I、T2V、I2V、optional synchronized audio
- 官方 recipe 推荐 8x H200/H100/A100
- 2x H200 / B300 是官方 recipe 里列出的 minimum 配置

推荐 8 卡启动：

```bash
vllm serve nvidia/Cosmos3-Super \
  --omni \
  --host 0.0.0.0 --port 8000 \
  --cfg-parallel-size 2 \
  --ulysses-degree 4 \
  --use-hsdp --hsdp-shard-size 8 \
  --init-timeout 1800
```

2 卡启动：

```bash
vllm serve nvidia/Cosmos3-Super \
  --omni \
  --host 0.0.0.0 --port 8000 \
  --cfg-parallel-size 2 \
  --use-hsdp --hsdp-shard-size 2 \
  --init-timeout 1800
```

T2V + sound 示例：

```bash
curl -sS -X POST http://localhost:8000/v1/videos/sync -H "Accept: video/mp4" \
  -F "model=nvidia/Cosmos3-Super" \
  -F "prompt=A robot arm is cleaning a plate in the kitchen" \
  -F "size=1280x720" \
  -F "num_frames=189" \
  -F "fps=24" \
  -F "num_inference_steps=35" \
  -F "guidance_scale=6.0" \
  -F "max_sequence_length=4096" \
  -F "flow_shift=10.0" \
  -F "generate_sound=true" \
  -F "sound_duration=7.875" \
  -F 'extra_params={"use_resolution_template":false,"use_duration_template":false,"guardrails":true}' \
  -F "seed=17" \
  -o cosmos3_super_t2vs.mp4
```

---

## 14. Offline Inference 示例

官方 Quickstart 的 text-to-image offline 示例：

```python
from vllm_omni.entrypoints.omni import Omni

if __name__ == "__main__":
    omni = Omni(model="Tongyi-MAI/Z-Image-Turbo")
    prompt = "a cup of coffee on the table"
    outputs = omni.generate(prompt)
    images = outputs[0].request_output.images
    images[0].save("coffee.png")
```

批量 prompts 示例官方也有，但官方明确提示：

> 当前不推荐依赖 batch inference，因为并非所有模型支持 batch，而且 batch 请求通常不一定带来明显性能提升。该接口主要用于与 vLLM 兼容及未来扩展。

---

## 15. 生产部署建议

### 15.1 版本固定

建议显式固定：

```text
vllm==0.22.0
vllm-omni==0.22.0 或源码 tag v0.22.0
Docker 镜像：vllm/vllm-omni:v0.22.0
```

不要在生产里直接用 latest。

### 15.2 Hugging Face 缓存

建议挂载：

```text
~/.cache/huggingface:/root/.cache/huggingface
```

避免容器重建后重复下载大模型。

### 15.3 端口与健康检查

启动后先验证：

```bash
curl http://127.0.0.1:8091/v1/models
```

### 15.4 GPU 规划

不同模型资源差异很大：

- Qwen3-Omni 30B：建议多 GPU。
- Qwen3-TTS 0.6B/1.7B：资源要求相对低。
- Qwen-Image：建议高显存 GPU，显存不足时开启 VAE slicing / tiling。
- Cosmos3-Super：官方 recipe 建议 8x H200/H100/A100，2x H200/B300 为 minimum。

### 15.5 反向代理

如果对外提供服务，建议在前面加 Nginx/Caddy：

- HTTPS
- API key 认证
- IP 白名单
- 请求体大小限制
- 超时配置

---

## 16. 常见问题

### 16.1 `--omni` 不生效

通常是版本不匹配。官方强调：

```text
vLLM-Omni 0.22.0 应搭配 vLLM 0.22.0
```

### 16.2 模型下载失败

检查：

- 是否设置 `HF_TOKEN`
- 是否接受模型 license / gated model 访问许可
- 是否挂载 Hugging Face cache
- 网络是否可访问 Hugging Face

### 16.3 图像/视频生成 OOM

可尝试：

- 降低分辨率
- 降低 batch size
- 使用 `--vae-use-slicing`
- 使用 `--vae-use-tiling`
- 使用 tensor parallel / USP / Ring attention
- 换更大显存 GPU

### 16.4 Cosmos3 Guardrail 报错

如果默认 guardrails 开启，需要访问 gated `nvidia/Cosmos-1.0-Guardrail`。否则可尝试：

```bash
--no-guardrails
```

---

## 17. 官方链接汇总

### GitHub / Release

- vLLM-Omni GitHub：  
  https://github.com/vllm-project/vllm-omni

- v0.22.0 Release：  
  https://github.com/vllm-project/vllm-omni/releases/tag/v0.22.0

- v0.22.0 tag 源码：  
  https://github.com/vllm-project/vllm-omni/tree/v0.22.0

### 官方文档

- ReadTheDocs 首页：  
  https://vllm-omni.readthedocs.io/en/latest/

- Installation：  
  https://vllm-omni.readthedocs.io/en/latest/getting_started/installation/

- Quickstart：  
  https://vllm-omni.readthedocs.io/en/latest/getting_started/quickstart/

- Supported Models：  
  https://vllm-omni.readthedocs.io/en/latest/models/supported_models/

- CLI serve：  
  https://vllm-omni.readthedocs.io/en/latest/cli/serve/

### API 文档

- Chat Completions API：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/docs/serving/chat_completions_api.md

- Speech API：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/docs/serving/speech_api.md

- Image Generation API：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/docs/serving/image_generation_api.md

- Image Edit API：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/docs/serving/image_edit_api.md

- Videos API：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/docs/serving/videos_api.md

### 使用案例文档

- Qwen3-Omni recipe：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/recipes/Qwen/Qwen3-Omni.md

- Qwen3-TTS recipe：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/recipes/Qwen/Qwen3-TTS.md

- Qwen-Image recipe：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/recipes/Qwen/Qwen-Image.md

- Qwen-Image-Edit recipe：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/recipes/Qwen/Qwen-Image-Edit.md

- Cosmos3-Nano recipe：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/recipes/cosmos3/Cosmos3-Nano.md

- Cosmos3-Super recipe：  
  https://github.com/vllm-project/vllm-omni/blob/v0.22.0/recipes/cosmos3/Cosmos3-Super.md

### Docker

- Docker Hub vLLM-Omni：  
  https://hub.docker.com/r/vllm/vllm-omni/tags

### 论文 / 社区

- vLLM-Omni Paper：  
  https://arxiv.org/abs/2602.02204

- vLLM forum：  
  https://discuss.vllm.ai

- vLLM Slack：  
  https://slack.vllm.ai
