# PP-OCRv6 对比 PP-OCRv5 变化整理

> 资料来源：PaddleOCR 官方 `version3.x` 文档、OCR Pipeline、Text Detection、Text Recognition 模块说明。  
> 结论时间：2026-06-13。  
> 注意：官方文档说明 PP-OCRv6 与 PP-OCRv5 的部分指标来自不同评测集，不能简单按表格数值直接横向比较；更适合看官方给出的相对提升、模型结构和能力变化。

## 一句话结论

PP-OCRv6 相比 PP-OCRv5，核心变化是：**换了新 backbone 和模型结构，默认模型升级为 PP-OCRv6_medium，提供 tiny / small / medium 三档，在识别、检测、多语言覆盖、边缘端部署上都做了重新设计**。

如果是新项目，优先用 **PP-OCRv6_medium**；如果是移动端或边缘端，可按算力选择 **PP-OCRv6_small** 或 **PP-OCRv6_tiny**。

---

## 1. 默认模型变化

PaddleOCR 3.7 的通用 OCR Pipeline 默认模型已经变为：

```text
PP-OCRv6_medium
```

官方 OCR Pipeline 文档说明：

- 通用 OCR Pipeline 支持 `PP-OCRv3`、`PP-OCRv4`、`PP-OCRv5`、`PP-OCRv6`。
- 默认模型是随 PaddleOCR 3.7 发布的 `PP-OCRv6_medium`。
- PP-OCRv6 基于新设计的 `PPLCNetV4` 统一 backbone。
- v6 提供 `tiny`、`small`、`medium` 三个档位。

这意味着：**以后直接用 PaddleOCR 3.7+ 的默认通用 OCR，默认就会走 PP-OCRv6_medium，而不是 PP-OCRv5。**

---

## 2. 模型档位变化

PP-OCRv5 的主要命名方式是：

| 场景 | PP-OCRv5 模型 |
|---|---|
| 服务端 | `PP-OCRv5_server_det`、`PP-OCRv5_server_rec` |
| 移动端 | `PP-OCRv5_mobile_det`、`PP-OCRv5_mobile_rec` |

PP-OCRv6 改成三档：

| 档位 | 适合场景 | 特点 |
|---|---|---|
| `PP-OCRv6_medium` | 服务端、高精度场景 | 精度最高，适合服务器部署 |
| `PP-OCRv6_small` | 移动端、普通边缘设备 | 精度和速度折中 |
| `PP-OCRv6_tiny` | 极轻量边缘设备、IoT | 模型最小，适合资源受限环境 |

这个变化比 v5 的 `server/mobile` 更细，可以按资源和精度需求选模型。

---

## 3. 检测模型变化

### 3.1 PP-OCRv6 检测模型

| 模型 | 检测 Hmean | 模型大小 | 官方描述 |
|---|---:|---:|---|
| `PP-OCRv6_medium_det` | 86.2* | 59.4 MB | 基于 `PPLCNetV4 + RepLKFPN`，精度最高，适合服务端 |
| `PP-OCRv6_small_det` | 84.1* | 9.6 MB | 平衡精度和效率，适合移动端 |
| `PP-OCRv6_tiny_det` | 80.6* | 1.9 MB | 超轻量检测模型，约 0.43M 参数，适合边缘/IoT |

### 3.2 PP-OCRv5 检测模型

| 模型 | 检测 Hmean | GPU 推理时间 | CPU 推理时间 | 模型大小 |
|---|---:|---:|---:|---:|
| `PP-OCRv5_server_det` | 83.8 | 89.55 / 70.19 ms | 383.15 / 383.15 ms | 84.3 MB |
| `PP-OCRv5_mobile_det` | 79.0 | 10.67 / 6.36 ms | 57.77 / 28.15 ms | 4.7 MB |

### 3.3 检测侧变化总结

PP-OCRv6 检测模型主要变化：

1. **结构升级**：`PP-OCRv6_medium_det` 使用 `PPLCNetV4 + RepLKFPN`。
2. **服务端模型更小**：v6 medium det 为 59.4 MB，v5 server det 为 84.3 MB。
3. **多了 tiny 档**：v6 tiny det 只有 1.9 MB，适合更低资源设备。
4. **官方称检测提升明显**：OCR Pipeline 文档写到 v6 medium 相比 v5 server 检测提升约 `+4.6%`，同时推理更快。

注意：官方表格中 v6 与 v5 的检测指标评测集不同，所以不要直接只看 `86.2` vs `83.8` 做绝对结论，应以官方相对说明和实际业务测试为准。

---

## 4. 识别模型变化

### 4.1 PP-OCRv6 识别模型

| 模型 | 平均识别准确率 | 模型大小 | 官方描述 |
|---|---:|---:|---|
| `PP-OCRv6_medium_rec` | 83.2* | 73.3 MB | 基于 `PPLCNetV4 + LightSVTR + CTC/NRTR 多头解码器` |
| `PP-OCRv6_small_rec` | 81.3* | 20.4 MB | 平衡体积和效果 |
| `PP-OCRv6_tiny_rec` | 73.5* | 4.4 MB | 极轻量识别模型 |

PP-OCRv6 识别模型的重点结构：

```text
PPLCNetV4 + LightSVTR + CTC/NRTR multi-head decoder
```

### 4.2 PP-OCRv5 识别模型

| 模型 | 平均识别准确率 | GPU 推理时间 | CPU 推理时间 | 模型大小 |
|---|---:|---:|---:|---:|
| `PP-OCRv5_server_rec` | 86.38 | 8.46 / 2.36 ms | 31.21 / 31.21 ms | 81 MB |
| `PP-OCRv5_mobile_rec` | 81.29 | 5.43 / 1.46 ms | 21.20 / 5.32 ms | 16 MB |

### 4.3 识别侧变化总结

PP-OCRv6 识别模型主要变化：

1. **结构换代**：v6 使用 `PPLCNetV4 + LightSVTR + CTC/NRTR 多头解码器`。
2. **语言覆盖扩大**：v6 单模型支持 50 种语言，tiny 支持 49 种语言。
3. **官方称 medium 识别提升**：官方文档写到 `PP-OCRv6_medium_rec` 相比 `PP-OCRv5_server_rec` 提升约 `+5.1%`。
4. **模型档位更完整**：从 v5 的 server/mobile，扩展为 medium/small/tiny。
5. **服务端模型略小**：v6 medium rec 为 73.3 MB，v5 server rec 为 81 MB。

注意：识别表格里的 v6 指标带 `*`，官方说明 v6 与 v5/v4 使用的评测集不同，不能直接把表格数值硬比。

---

## 5. 多语言能力变化

PP-OCRv5 的重点是：

- 简体中文
- 繁体中文
- 英文
- 日文
- 手写、竖排、拼音、生僻字等复杂场景

PP-OCRv6 的重点升级是：

- 单模型支持 **50 种语言**
- 包括中文、英文、日文，以及 46 种拉丁语系语言
- tiny 档支持 49 种语言

所以如果业务中有多语种 OCR，尤其是中英日 + 拉丁语系混合场景，PP-OCRv6 更合适。

---

## 6. 部署和使用变化

### Python 使用

PaddleOCR 3.7+ 中，通用 OCR Pipeline 默认使用 `PP-OCRv6_medium`。Python 中通常直接用：

```python
from paddleocr import PaddleOCR

ocr = PaddleOCR()
result = ocr.predict("test.jpg")
```

如果需要显式指定模型，可以按官方参数传入对应的检测、识别模型名或模型目录。

### Go 使用

PaddleOCR 官方主路径仍然是 Python / PaddleX / Serving。Go 项目不建议直接把 PaddleOCR 深度绑定进 Go 进程里，推荐：

```text
Go 业务服务  ->  HTTP 调用 PaddleOCR / PaddleX Serving  ->  返回 OCR JSON
```

这样 Go 项目只负责业务逻辑，OCR 模型、GPU、CUDA、PaddlePaddle 环境放在独立服务中维护。

### 推理引擎

PaddleOCR 3.x 支持统一推理引擎配置机制，可选择：

- `paddle_static`：多数本地推理场景默认推荐
- `onnxruntime`
- `transformers`
- `paddle_dynamic`
- 高性能推理模式 HPI
- TensorRT / OpenVINO 等后端组合

---

## 7. 什么时候升级到 PP-OCRv6

建议升级的情况：

1. 新项目，没有历史兼容负担。
2. 想使用 PaddleOCR 3.7+ 默认 OCR Pipeline。
3. 需要更强多语言识别能力。
4. 需要服务端更高精度，同时希望模型体积更合理。
5. 需要 tiny/small/medium 三档，方便服务端、移动端、边缘端统一选型。
6. 当前 v5 在复杂图片、手写、旋转、模糊、多语种场景下效果不够好。

可以暂缓升级的情况：

1. 现有 PP-OCRv5 已在生产稳定运行。
2. 已针对 v5 模型做过大量参数调优、阈值调优或后处理规则。
3. 业务只处理固定模板，v5 已经够用。
4. 当前部署环境严格锁定旧版 PaddleOCR / PaddlePaddle。

---

## 8. 迁移建议

从 PP-OCRv5 迁移到 PP-OCRv6，建议不要直接全量替换，按下面流程做：

1. **保留 v5 线上版本**：不要一上来替换生产服务。
2. **准备业务样本集**：至少覆盖清晰图、模糊图、倾斜图、表格截图、票据、手写、多语言等真实样本。
3. **并行跑 v5/v6**：同一批图片分别跑 v5 和 v6。
4. **统计业务指标**：不要只看官方指标，要看自己的字段提取准确率、漏检率、误检率、耗时、GPU/CPU 占用。
5. **优先测试 medium**：服务端先测 `PP-OCRv6_medium`。
6. **边缘端再测 small/tiny**：移动端、IoT 或低配置机器再选 `small` / `tiny`。
7. **灰度上线**：先小流量切换，再逐步替换。

---

## 9. v5 / v6 对比总表

| 维度 | PP-OCRv5 | PP-OCRv6 |
|---|---|---|
| 默认地位 | 旧一代通用 OCR 模型 | PaddleOCR 3.7+ 默认通用 OCR 模型 |
| 模型档位 | server / mobile | medium / small / tiny |
| 检测结构 | v5 检测模型 | `PPLCNetV4 + RepLKFPN` |
| 识别结构 | v5 识别模型 | `PPLCNetV4 + LightSVTR + CTC/NRTR 多头解码器` |
| 语言覆盖 | 中、繁中、英、日等重点语言 | 单模型 50 种语言，tiny 49 种 |
| 服务端检测模型大小 | server det 84.3 MB | medium det 59.4 MB |
| 服务端识别模型大小 | server rec 81 MB | medium rec 73.3 MB |
| 边缘端选择 | mobile | small / tiny 两档更细 |
| 官方相对提升 | - | medium 相比 v5 server：识别约 +5.1%，检测约 +4.6%，且推理更快 |
| 适合场景 | 已稳定生产、旧项目兼容 | 新项目、多语言、高精度、统一多端选型 |

---

## 10. 老王项目里的推荐结论

如果后续要在 Go / Python 项目里做 OCR：

### 推荐方案

```text
Go / Python 业务层
        ↓ HTTP / 内网 API
PaddleOCR / PaddleX Serving 独立 OCR 服务
        ↓
PP-OCRv6_medium 默认模型
```

### 推荐选择

- **服务端默认**：`PP-OCRv6_medium`
- **普通边缘端**：`PP-OCRv6_small`
- **极低资源设备**：`PP-OCRv6_tiny`
- **已有稳定 v5 生产环境**：先并行测试，不要直接替换

### 我的判断

PP-OCRv6 更像是 PaddleOCR 3.7 后的新默认基线，适合新项目直接采用。PP-OCRv5 仍然可用，但更适合作为已有项目的稳定版本保留。真正迁移时，应以老王自己的图片样本跑一轮 A/B 测试，而不是只看官方表格。

---

## 参考资料

- PaddleOCR 官方 OCR Pipeline 文档：`docs/version3.x/pipeline_usage/OCR.en.md`
- PaddleOCR 官方文本检测模块文档：`docs/version3.x/module_usage/text_detection.en.md`
- PaddleOCR 官方文本识别模块文档：`docs/version3.x/module_usage/text_recognition.en.md`
- PaddleOCR GitHub：`https://github.com/PaddlePaddle/PaddleOCR`
