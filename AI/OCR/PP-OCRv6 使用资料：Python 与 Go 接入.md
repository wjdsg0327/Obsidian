# PP-OCRv6 使用资料：Python 与 Go 接入

> 资料整理时间：2026-06-13  
> 主要来源：PaddleOCR 3.x 官方文档、PaddleOCR GitHub 文档。

## 1. PP-OCRv6 是什么

PP-OCRv6 是 PaddleOCR 3.x 通用 OCR Pipeline 支持的新一代 PP-OCR 模型系列。官方文档中说明：通用 OCR Pipeline 支持 PP-OCRv3、PP-OCRv4、PP-OCRv5、PP-OCRv6，并且 PaddleOCR 3.7 发布后默认模型为 **PP-OCRv6_medium**。

PP-OCRv6 的特点：

- 基于新的 **PPLCNetV4** 统一骨干网络。
- 分为 **tiny / small / medium** 三个档位。
- medium 档偏服务端精度；small 偏精度与效率平衡；tiny 偏边缘设备、IoT、轻量化场景。
- 单个识别模型支持 50 种语言，包括中文、英文、日文和多种拉丁语系语言。
- 通用 OCR Pipeline 由 5 个模块组成：
  - 文档图像方向分类，可选
  - 文本图像矫正，可选
  - 文本行方向分类，可选
  - 文本检测
  - 文本识别

## 2. 什么时候用 PP-OCRv6

适合：

- 图片文字识别。
- 截图、票据、扫描图、普通文档图片 OCR。
- PDF / 多页 TIFF 的 OCR 服务化处理。
- 需要中英文、多语言混合识别的场景。
- 希望直接调用成熟 OCR Pipeline，而不是自己训练检测/识别模型。

不太适合：

- 复杂版面解析、表格、公式、印章、版面结构还原：更适合看 **PP-StructureV3**。
- 文档问答、关键信息抽取：更适合 **PP-ChatOCRv4**。
- 大模型视觉文档理解：更适合 **PaddleOCR-VL**。

## 3. 安装方式

### 3.1 基础安装

```bash
# 只安装通用 OCR 和文档图像预处理能力
pip install paddleocr

# 安装全部能力：文档解析、信息抽取、文档翻译等
pip install "paddleocr[all]"
```

Python 版本：

- `paddleocr` 基础包支持 Python 3.8+。
- 一些可选能力如 `doc-parser`、`ie`、`trans`、`all` 等通常要求 Python 3.9+。

### 3.2 推理引擎

PaddleOCR 3.x 支持多种推理引擎：

| 引擎 | 说明 |
|---|---|
| `paddle` / `paddle_static` / `paddle_dynamic` | PaddlePaddle 框架推理，默认主路径 |
| `onnxruntime` | ONNX Runtime 推理 |
| `transformers` | Hugging Face Transformers 生态推理，主要用于相关模型 |

官方建议：一个环境里尽量只装一个推理引擎，避免依赖冲突。

如果使用默认 PaddlePaddle 推理，需要按机器环境安装 PaddlePaddle：

```bash
# 示例：CPU 环境，具体版本以 PaddlePaddle 官方安装页为准
pip install paddlepaddle
```

## 4. Python 怎么使用 PP-OCRv6

Python 是 PaddleOCR 官方最直接、最推荐的集成方式。

### 4.1 最小可用示例

```python
from paddleocr import PaddleOCR

ocr = PaddleOCR(
    ocr_version="PP-OCRv6",
    use_doc_orientation_classify=False,
    use_doc_unwarping=False,
    use_textline_orientation=False,
)

result = ocr.predict("./demo.jpg")

for res in result:
    res.print()
    res.save_to_img("output")
    res.save_to_json("output")
```

说明：

- `ocr_version="PP-OCRv6"`：指定使用 PP-OCRv6 系列模型。
- `use_doc_orientation_classify=False`：不启用文档方向分类，速度更快；如果图片可能旋转，可改为 `True`。
- `use_doc_unwarping=False`：不启用图像矫正；拍照歪斜、弯曲文档可考虑开启。
- `use_textline_orientation=False`：不启用文本行方向分类；如果存在倒置文本可开启。
- `save_to_img("output")`：保存可视化识别结果。
- `save_to_json("output")`：保存结构化 JSON 结果。

### 4.2 使用 GPU

```python
from paddleocr import PaddleOCR

ocr = PaddleOCR(
    ocr_version="PP-OCRv6",
    device="gpu:0",
    use_doc_orientation_classify=False,
    use_doc_unwarping=False,
    use_textline_orientation=False,
)

result = ocr.predict("./demo.jpg")
for res in result:
    res.print()
```

如果没有 GPU，可以用：

```python
ocr = PaddleOCR(ocr_version="PP-OCRv6", device="cpu")
```

### 4.3 指定模型档位

PP-OCRv6 有 detection 和 recognition 两部分模型。可以手动指定：

```python
from paddleocr import PaddleOCR

ocr = PaddleOCR(
    text_detection_model_name="PP-OCRv6_medium_det",
    text_recognition_model_name="PP-OCRv6_medium_rec",
    use_doc_orientation_classify=False,
    use_doc_unwarping=False,
    use_textline_orientation=False,
)

result = ocr.predict("./demo.jpg")
for res in result:
    res.print()
```

可选模型：

| 档位 | 检测模型 | 识别模型 | 适合场景 |
|---|---|---|---|
| medium | `PP-OCRv6_medium_det` | `PP-OCRv6_medium_rec` | 服务端、高精度 |
| small | `PP-OCRv6_small_det` | `PP-OCRv6_small_rec` | 速度与精度平衡 |
| tiny | `PP-OCRv6_tiny_det` | `PP-OCRv6_tiny_rec` | 边缘设备、轻量化 |

### 4.4 命令行快速测试

```bash
paddleocr ocr -i ./demo.jpg \
  --ocr_version PP-OCRv6 \
  --use_doc_orientation_classify False \
  --use_doc_unwarping False \
  --use_textline_orientation False \
  --save_path ./output \
  --device cpu
```

GPU：

```bash
paddleocr ocr -i ./demo.jpg \
  --ocr_version PP-OCRv6 \
  --save_path ./output \
  --device gpu:0
```

## 5. Go 怎么使用 PP-OCRv6

重点结论：**PaddleOCR 官方没有把 Go 当成本地 SDK 主路径；Go 推荐通过 HTTP 服务调用 PaddleOCR / PaddleX 部署出来的 OCR 服务。**

也就是说：

- Python 负责部署 PP-OCRv6 服务。
- Go 负责上传图片 / PDF，调用 HTTP 接口拿结果。

这是生产环境最稳的方式：Go 项目不用直接处理 PaddlePaddle、CUDA、模型依赖，只调用服务。

## 6. 部署 OCR HTTP 服务

PaddleOCR 3.x 官方推荐通过 PaddleX Serving 做服务化。

### 6.1 安装 Serving 插件

```bash
paddlex --install serving
```

### 6.2 启动通用 OCR 服务

```bash
paddlex --serve --pipeline OCR
```

默认服务地址：

```text
http://localhost:8080
```

OCR 接口：

```text
POST /ocr
```

请求体是 JSON。

核心字段：

| 字段 | 类型 | 说明 |
|---|---|---|
| `file` | string | 图片/PDF 的 Base64 内容，或服务端可访问的 URL |
| `fileType` | int/null | `0` 表示 PDF，`1` 表示图片；不填时可从 URL 推断 |
| `useDocOrientationClassify` | bool/null | 是否启用文档方向分类 |
| `useDocUnwarping` | bool/null | 是否启用文档矫正 |
| `useTextlineOrientation` | bool/null | 是否启用文本行方向识别 |
| `visualize` | bool/null | 是否返回可视化图片 |

成功响应中：

- `result.ocrResults[].prunedResult`：精简后的 OCR 结果。
- `result.ocrResults[].ocrImage`：OCR 可视化图片，默认 Base64。
- `result.dataInfo`：输入文件信息。

## 7. Go 调用示例

### 7.1 最小 HTTP 调用

```go
package main

import (
    "bytes"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
)

type OCRResponse struct {
    ErrorCode int    `json:"errorCode"`
    ErrorMsg  string `json:"errorMsg"`
    Result    struct {
        OCRResults []struct {
            PrunedResult map[string]any `json:"prunedResult"`
            OCRImage     *string        `json:"ocrImage"`
        } `json:"ocrResults"`
        DataInfo any `json:"dataInfo"`
    } `json:"result"`
}

func main() {
    apiURL := "http://localhost:8080/ocr"
    imagePath := "./demo.jpg"

    fileBytes, err := os.ReadFile(imagePath)
    if err != nil {
        panic(err)
    }

    payload := map[string]any{
        "file":     base64.StdEncoding.EncodeToString(fileBytes),
        "fileType": 1,
        "useDocOrientationClassify": false,
        "useDocUnwarping": false,
        "useTextlineOrientation": false,
        "visualize": false,
    }

    body, err := json.Marshal(payload)
    if err != nil {
        panic(err)
    }

    client := &http.Client{Timeout: 60 * time.Second}
    req, err := http.NewRequest("POST", apiURL, bytes.NewReader(body))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    if resp.StatusCode != http.StatusOK {
        panic(fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(respBody)))
    }

    var data OCRResponse
    if err := json.Unmarshal(respBody, &data); err != nil {
        panic(err)
    }

    if data.ErrorCode != 0 {
        panic(fmt.Sprintf("OCR error: %s", data.ErrorMsg))
    }

    for i, item := range data.Result.OCRResults {
        fmt.Printf("page/image %d:\n%+v\n", i, item.PrunedResult)
    }
}
```

### 7.2 保存 OCR 可视化图片

如果请求中 `visualize` 为 `true` 或服务端默认返回可视化图，可以保存 `ocrImage`：

```go
if item.OCRImage != nil {
    imgBytes, err := base64.StdEncoding.DecodeString(*item.OCRImage)
    if err == nil {
        _ = os.WriteFile(fmt.Sprintf("ocr_%d.jpg", i), imgBytes, 0644)
    }
}
```

## 8. Python 调 HTTP 服务示例

如果 Python 项目也想走服务化接口，而不是本地加载模型：

```python
import base64
import requests

api_url = "http://localhost:8080/ocr"
file_path = "./demo.jpg"

with open(file_path, "rb") as f:
    file_data = base64.b64encode(f.read()).decode("ascii")

payload = {
    "file": file_data,
    "fileType": 1,
    "useDocOrientationClassify": False,
    "useDocUnwarping": False,
    "useTextlineOrientation": False,
    "visualize": False,
}

resp = requests.post(api_url, json=payload, timeout=60)
resp.raise_for_status()

data = resp.json()
for item in data["result"]["ocrResults"]:
    print(item["prunedResult"])
```

## 9. 推荐架构

### 9.1 单机脚本 / 离线批处理

推荐：Python 直接调用 PaddleOCR。

```text
图片 / PDF -> Python PaddleOCR -> JSON / 图片结果
```

优点：开发快，代码少。  
缺点：Python 环境、模型、GPU 依赖跟业务代码耦合。

### 9.2 Go 后端项目

推荐：Go 调 OCR 服务。

```text
Go 业务服务 -> HTTP -> PaddleOCR Serving -> OCR JSON 结果
```

优点：

- Go 服务保持干净。
- OCR 服务可独立扩容。
- Python/PaddlePaddle/CUDA 环境独立维护。
- 多语言客户端都能复用同一个 OCR 服务。

### 9.3 高并发生产环境

建议：

```text
Go API 网关 / 业务服务
    -> 队列 / 任务系统
    -> OCR Worker / PaddleOCR Serving
    -> 对象存储 / 数据库
```

注意点：

- 大图片、PDF 不建议长期用 Base64 直接塞请求，体积会膨胀。
- 可以用文件 URL 或对象存储 URL。
- 多页 PDF 默认可能只处理前 10 页，可通过 Serving 配置调整。
- 返回可视化图片会让响应很大，生产中建议 `visualize=false`，只返回结构化结果。

## 10. 常见问题

### 10.1 Go 能不能直接本地加载 PP-OCRv6 模型？

理论上可以折腾：比如导出 ONNX，用 Go 调 ONNX Runtime；或用 cgo 调 C++/Paddle Inference。但这不是官方最顺的路线，工程复杂度高。

推荐优先级：

1. **Go 调 HTTP OCR 服务**：最推荐。
2. Go 调命令行 `paddleocr ocr ...`：能用，但进程开销大，不适合高并发。
3. Go + ONNX Runtime：适合强本地化、边缘部署，但需要处理前处理、后处理、字典、检测框还原等细节。
4. cgo / C++ 推理：复杂度最高。

### 10.2 PP-OCRv6 默认就是最新吗？

在 PaddleOCR 3.7 的通用 OCR Pipeline 文档中，默认模型是 `PP-OCRv6_medium`。为了避免未来默认值变化，项目里建议显式写：

```python
PaddleOCR(ocr_version="PP-OCRv6")
```

或者显式指定：

```python
PaddleOCR(
    text_detection_model_name="PP-OCRv6_medium_det",
    text_recognition_model_name="PP-OCRv6_medium_rec",
)
```

### 10.3 medium / small / tiny 怎么选？

- 服务器、有 GPU、追求精度：`medium`
- 普通 CPU / 轻量服务：`small`
- 移动端、边缘设备、IoT：`tiny`

### 10.4 中文识别需要 `lang="ch"` 吗？

PP-OCRv6 的识别模型本身是多语言方向，支持中文、英文、日文等多种语言。实际项目里可以先使用默认 PP-OCRv6；如果你希望明确语言策略，可以结合 `lang` 参数测试效果。

## 11. 参考链接

- PaddleOCR GitHub：<https://github.com/PaddlePaddle/PaddleOCR>
- PaddleOCR 3.x 安装文档：<https://www.paddleocr.ai/latest/en/version3.x/installation.html>
- 通用 OCR Pipeline 文档：<https://www.paddleocr.ai/latest/en/version3.x/pipeline_usage/OCR.html>
- Serving 部署文档：<https://www.paddleocr.ai/latest/en/version3.x/inference_deployment/serving/serving.html>
- 推理引擎配置文档：<https://www.paddleocr.ai/latest/en/version3.x/inference_deployment/local_inference/inference_engine.html>
