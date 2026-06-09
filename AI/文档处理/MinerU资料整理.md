# MinerU 资料整理：说明、教程、使用方法与 Python 案例

> 适用项目：AI国际学校智能作业系统  
> 整理时间：2026-06-07  
> 资料来源：MinerU 官方 GitHub、官方文档、MinerU-Ecosystem 等公开资料。

---

## 1. MinerU 是什么

MinerU 是 OpenDataLab 开源的文档解析工具，主要用于把复杂文档转换成适合大模型、RAG、知识库、题库导入等场景使用的结构化内容。

它可以把以下类型的输入转换为 Markdown、JSON、图片、表格、公式等中间结果：

- PDF
- 图片
- DOCX
- PPTX
- XLSX

在实际项目里，可以把 MinerU 理解为一个 **文档解析 Provider**：

```text
PDF / 扫描件 / 教材 / 真题 / 讲义
        ↓
MinerU 文档解析
        ↓
Markdown / JSON / 图片资源 / 表格 HTML / 公式 LaTeX
        ↓
题库导入、知识库、RAG、AI 批改、内容审核
```

---

## 2. 核心能力

### 2.1 文档结构解析

MinerU 能识别并尽量保留文档中的结构信息，例如：

- 标题
- 段落
- 列表
- 多栏排版
- 页眉页脚
- 页码
- 脚注
- 阅读顺序

它的目标不是简单 OCR，而是把文档转成更适合机器处理的结构化内容。

### 2.2 公式识别

MinerU 支持识别文档中的数学公式，并转换为 LaTeX。

适合场景：

- 数学试卷
- 物理试卷
- 科学教材
- 论文
- 含公式的讲义

### 2.3 表格识别

MinerU 支持识别表格，并可输出为 HTML 等形式。

适合场景：

- 成绩表
- 实验数据表
- 对照表
- 财务/统计类文档
- 教材中的表格题

### 2.4 图片与图文混排处理

MinerU 可以提取文档中的图片，并在输出结果中保留图片引用关系。

适合场景：

- 带图题目
- 几何图形
- 图表题
- 地理/生物/物理配图
- 教材插图

### 2.5 OCR 能力

MinerU 能自动检测扫描版 PDF、乱码 PDF，并启用 OCR。

官方资料中提到 OCR 支持 100+ 种语言识别。对于学校系统，主要关注：

- 中文
- 英文
- 中英混排
- 数学公式
- 表格
- 图片题

---

## 3. 适合放在本项目里的定位

对于“AI国际学校智能作业系统”，MinerU 不建议直接替代所有 OCR，而是作为 **PDF / 文档解析 Provider 候选方案**。

### 3.1 推荐使用场景

适合用 MinerU：

- PDF 真题导入
- 教材 / 教辅资料解析
- 扫描版试卷初步解析
- 含公式、表格、图片的复杂文档
- 需要转成 Markdown / JSON 后进入题库或知识库的材料
- RAG 知识库资料预处理

### 3.2 不建议直接使用的场景

不建议 MinerU 直接承担：

- 学生手写答案识别
- 实时拍照批改
- 手机端低延迟 OCR
- 对单张答题卡进行秒级识别

这些场景更适合独立 OCR Provider，例如：

- PaddleOCR
- 云厂商 OCR
- 专门的手写识别模型
- 多模态大模型 OCR

### 3.3 项目内推荐架构

```text
文件上传
  ↓
判断文件类型
  ├─ PDF / DOCX / PPTX / XLSX → Document Parse Provider
  │                              └─ MinerU 候选
  │
  ├─ 图片 / 拍照作业 → OCR Provider
  │                     └─ PaddleOCR / 云 OCR / 多模态模型
  │
  └─ 手写答案 → Handwriting OCR / VLM Provider

解析结果
  ↓
人工审核
  ↓
题库 / 作业 / 知识库
```

---

## 4. 安装教程

### 4.1 Python 版本要求

官方快速入门资料中提到支持 Python 3.10 - 3.13。

注意：

- Windows 因部分依赖限制，建议 Python 3.10 - 3.12。
- Linux / macOS 环境相对更适合部署。
- 生产环境建议优先用 Docker，减少依赖冲突。

---

### 4.2 使用 pip / uv 安装

官方推荐方式之一：

```bash
pip install --upgrade pip
pip install uv
uv pip install -U "mineru[all]"
```

如果从源码安装：

```bash
git clone https://github.com/opendatalab/MinerU.git
cd MinerU
uv pip install -e .[all]
```

安装后检查版本：

```bash
mineru --version
```

查看帮助：

```bash
mineru --help
```

---

### 4.3 模型源配置

MinerU 默认模型源可能使用 HuggingFace。

如果国内访问 HuggingFace 不稳定，可以切换到 ModelScope：

```bash
export MINERU_MODEL_SOURCE=modelscope
```

如果已经提前下载好模型，也可以配置本地模型路径，并使用：

```bash
export MINERU_MODEL_SOURCE=local
```

生产环境建议：

- 首次部署时提前下载模型。
- 不要依赖服务运行时临时下载。
- 将模型目录持久化。
- Docker 部署时挂载模型目录。

---

## 5. 命令行使用

### 5.1 最简单的解析命令

```bash
mineru -p input.pdf -o ./output
```

含义：

- `-p` / `--path`：输入文件或目录
- `-o` / `--output`：输出目录

MinerU 支持本地文件，也支持目录批量解析。

---

### 5.2 指定解析后端

```bash
mineru -p input.pdf -o ./output -b pipeline
```

常见后端：

```text
pipeline
hybrid-auto-engine
hybrid-http-client
vlm-auto-engine
vlm-http-client
```

简单理解：

| 后端 | 说明 | 适合场景 |
|---|---|---|
| pipeline | 传统解析管线，兼容性较好 | CPU / 普通 GPU / 稳定解析 |
| hybrid | 混合模式 | 兼顾兼容性和效果 |
| vlm | 视觉语言模型模式 | 复杂版面、高精度需求，但硬件要求更高 |
| http-client | 调用外部 OpenAI-compatible 服务 | 已部署远程模型服务时 |

---

### 5.3 指定解析方法

```bash
mineru -p input.pdf -o ./output -m auto
```

可选：

```text
auto
txt
ocr
```

说明：

- `auto`：自动判断，默认推荐。
- `txt`：优先使用 PDF 内文本层。
- `ocr`：强制走 OCR，适合扫描件。

---

### 5.4 指定语言

```bash
mineru -p input.pdf -o ./output --lang ch
```

常见语言参数：

```text
ch        中文
ch_server 中文高精度服务模型相关配置
en        英文
korean    韩文
japan     日文
```

中英混排试卷可以先用：

```bash
mineru -p paper.pdf -o ./output --lang ch
```

如果英文文档识别效果不好，可以试：

```bash
mineru -p paper.pdf -o ./output --lang en
```

---

### 5.5 只解析指定页码

```bash
mineru -p input.pdf -o ./output --start 0 --end 4
```

说明：

- 页码从 0 开始。
- `--start 0 --end 4` 通常表示解析第 1 页到第 5 页。

适合先做 PoC 测试，不必解析整本教材。

---

### 5.6 开启 / 关闭公式解析

```bash
mineru -p input.pdf -o ./output --formula true
```

关闭公式解析：

```bash
mineru -p input.pdf -o ./output --formula false
```

如果文档没有公式，关闭公式解析可能减少耗时。

---

### 5.7 开启 / 关闭表格解析

```bash
mineru -p input.pdf -o ./output --table true
```

关闭表格解析：

```bash
mineru -p input.pdf -o ./output --table false
```

---

### 5.8 GPU 指定

如果服务器有多张 GPU，可以用：

```bash
CUDA_VISIBLE_DEVICES=0 mineru -p input.pdf -o ./output
```

使用第 2 张 GPU：

```bash
CUDA_VISIBLE_DEVICES=1 mineru -p input.pdf -o ./output
```

禁用 GPU，只用 CPU：

```bash
CUDA_VISIBLE_DEVICES="" mineru -p input.pdf -o ./output -b pipeline
```

---

## 6. FastAPI / WebUI 使用

### 6.1 启动 FastAPI 服务

```bash
mineru-api --host 0.0.0.0 --port 8000
```

启动后可访问：

```text
http://localhost:8000/docs
```

在接口文档里可以看到可调用接口。

常见文件解析接口：

```text
POST /file_parse
```

注意：

- `/file_parse` 是 MinerU FastAPI 服务的接口。
- 如果使用 Docker 或远程服务，要确认启动的是 `mineru-api`，而不是其他推理服务。

---

### 6.2 启动 Gradio WebUI

```bash
mineru-gradio --host 0.0.0.0 --port 7860
```

然后浏览器访问：

```text
http://localhost:7860
```

适合：

- 人工测试解析效果
- 对比不同参数
- 给业务人员演示
- PoC 阶段快速验证

---

## 7. Docker 使用建议

MinerU 官方提供 Docker 部署方式，适合生产环境或 PoC 环境。

推荐原因：

- 减少 Python 依赖冲突
- 方便固定版本
- 方便挂载模型目录
- 方便暴露 FastAPI 服务
- 更适合服务器部署

典型思路：

```text
Docker 容器
  ├─ MinerU Runtime
  ├─ 模型目录挂载
  ├─ 输入文件目录挂载
  ├─ 输出文件目录挂载
  └─ FastAPI 端口 8000
```

本项目若做 PoC，建议：

```text
AI国际学校系统后端
  ↓ HTTP 上传 PDF
MinerU FastAPI 服务
  ↓ 返回 Markdown / JSON / 图片资源
后端保存解析结果
  ↓
人工审核后入题库
```

---

## 8. 输出结果说明

MinerU 通常会输出多类结果，具体文件名可能随版本和参数变化，但核心类型包括：

### 8.1 Markdown

用于直接阅读、RAG 入库、知识库沉淀。

示例：

```markdown
# 第一章 函数

## 1.1 函数的概念

设 A、B 是两个非空集合……

$$
f(x)=x^2+1
$$
```

### 8.2 JSON / 中间结构

用于程序处理，例如：

- 页面信息
- 文本块
- 表格块
- 图片块
- 公式块
- 坐标信息
- 阅读顺序

项目中可以把 JSON 作为题库导入的中间格式。

### 8.3 图片资源

文档中的图片、题图、表格截图等可能会被提取出来，供 Markdown 或 JSON 引用。

### 8.4 可视化结果

MinerU 支持一些布局可视化、span 可视化结果，用于检查解析质量。

PoC 阶段建议保留这些结果，方便判断：

- 版面识别是否正确
- 阅读顺序是否正确
- 表格边界是否正确
- 题图是否丢失
- 公式是否识别正确

---

## 9. Python 案例 1：用 subprocess 调用命令行

这是最稳妥的 Python 调用方式之一。后端系统不用依赖 MinerU 内部 API，只要调用命令行即可。

```python
from pathlib import Path
import subprocess


def parse_pdf_with_mineru(input_pdf: str, output_dir: str):
    input_pdf = Path(input_pdf)
    output_dir = Path(output_dir)
    output_dir.mkdir(parents=True, exist_ok=True)

    cmd = [
        "mineru",
        "-p", str(input_pdf),
        "-o", str(output_dir),
        "-b", "pipeline",
        "-m", "auto",
        "--lang", "ch",
        "--formula", "true",
        "--table", "true",
    ]

    result = subprocess.run(
        cmd,
        text=True,
        capture_output=True,
        check=False,
    )

    if result.returncode != 0:
        raise RuntimeError(
            f"MinerU parse failed\nSTDOUT:\n{result.stdout}\nSTDERR:\n{result.stderr}"
        )

    return {
        "input": str(input_pdf),
        "output_dir": str(output_dir),
        "stdout": result.stdout,
    }


if __name__ == "__main__":
    info = parse_pdf_with_mineru(
        input_pdf="./data/test.pdf",
        output_dir="./output/test",
    )
    print(info)
```

优点：

- 简单稳定
- 不绑定 MinerU 内部 Python API
- 容易做进程隔离
- 适合后端异步任务

缺点：

- 每次调用会启动进程
- 大批量任务要做好队列和并发控制

---

## 10. Python 案例 2：调用本地 MinerU FastAPI `/file_parse`

先启动服务：

```bash
mineru-api --host 0.0.0.0 --port 8000
```

Python 调用：

```python
from pathlib import Path
import requests


def parse_files_by_mineru_api(pdf_paths, api_base="http://127.0.0.1:8000"):
    url = f"{api_base.rstrip('/')}/file_parse"

    files = []
    opened_files = []

    try:
        for pdf_path in pdf_paths:
            path = Path(pdf_path)
            f = path.open("rb")
            opened_files.append(f)
            files.append(("files", (path.name, f, "application/pdf")))

        data = {
            # 多个文件时，可按接口要求传逗号分隔或统一语言配置
            "lang_list": "ch",
            "backend": "pipeline",
            "parse_method": "auto",
            "formula_enable": "true",
            "table_enable": "true",
        }

        response = requests.post(
            url,
            files=files,
            data=data,
            timeout=300,
        )
        response.raise_for_status()
        return response.json()

    finally:
        for f in opened_files:
            f.close()


if __name__ == "__main__":
    result = parse_files_by_mineru_api([
        "./data/math_paper.pdf",
    ])
    print(result)
```

适合后端项目集成：

```text
Spring Boot / Python / Node.js 后端
        ↓ HTTP multipart/form-data
MinerU FastAPI /file_parse
        ↓ JSON 结果
业务系统保存解析结果
```

注意：

- 要确认 MinerU API 服务已启动。
- 请求的是 MinerU FastAPI 服务，不是 VLLM / SGLang 推理服务。
- 大文件要增加超时时间。
- 生产环境建议用异步任务队列，不要让用户请求一直阻塞。

---

## 11. Python 案例 3：解析结果入题库的中间结构

MinerU 的输出不建议直接进入正式题库，应该先转成中间结构，再人工审核。

示例结构：

```python
from dataclasses import dataclass, field
from typing import List, Optional


@dataclass
class ParsedQuestionCandidate:
    source_file: str
    page_no: int
    raw_text: str
    markdown: str
    images: List[str] = field(default_factory=list)
    formulas: List[str] = field(default_factory=list)
    tables: List[str] = field(default_factory=list)
    confidence: Optional[float] = None
    provider: str = "mineru"
    status: str = "pending_review"


candidate = ParsedQuestionCandidate(
    source_file="math_paper_2025.pdf",
    page_no=3,
    raw_text="已知函数 f(x)=x^2+1，求……",
    markdown="已知函数 $f(x)=x^2+1$，求……",
    formulas=["f(x)=x^2+1"],
    images=["images/page_3_fig_1.png"],
)

print(candidate)
```

推荐进入题库前的流程：

```text
MinerU 解析
  ↓
题目候选切分
  ↓
AI 辅助识别题干 / 选项 / 答案 / 解析
  ↓
人工审核
  ↓
正式题库
```

---

## 12. 在 AI 国际学校系统中的接入建议

### 12.1 后端服务设计

建议新增一个文档解析服务接口：

```http
POST /api/document-parse/tasks
```

请求：

```json
{
  "fileId": "file_123",
  "provider": "mineru",
  "parseMethod": "auto",
  "language": "ch",
  "formulaEnable": true,
  "tableEnable": true
}
```

返回：

```json
{
  "taskId": "parse_task_001",
  "status": "pending"
}
```

查询任务：

```http
GET /api/document-parse/tasks/{taskId}
```

返回：

```json
{
  "taskId": "parse_task_001",
  "status": "success",
  "provider": "mineru",
  "outputs": {
    "markdownFileId": "file_md_001",
    "jsonFileId": "file_json_001",
    "imageFileIds": ["img_001", "img_002"]
  },
  "reviewRequired": true
}
```

---

### 12.2 数据库字段建议

可以新增文档解析任务表：

```sql
CREATE TABLE document_parse_task (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    file_id BIGINT NOT NULL,
    provider VARCHAR(50) NOT NULL,
    parse_method VARCHAR(50),
    language VARCHAR(50),
    formula_enable TINYINT DEFAULT 1,
    table_enable TINYINT DEFAULT 1,
    status VARCHAR(50) NOT NULL,
    output_markdown_file_id BIGINT NULL,
    output_json_file_id BIGINT NULL,
    error_message TEXT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);
```

---

### 12.3 MVP PoC 范围

建议先选 5-10 份真实材料测试：

- 2 份中文数学 PDF 真题
- 2 份英文 PDF 讲义
- 1 份扫描版试卷
- 1 份含表格材料
- 1 份含大量图片材料
- 1 份 DOCX / PPTX，如实际项目需要

评估指标：

| 指标 | 说明 |
|---|---|
| 文本准确率 | 是否漏字、错字、乱码 |
| 公式准确率 | LaTeX 是否正确 |
| 表格准确率 | 表格结构是否保留 |
| 图片保留 | 题图是否丢失 |
| 阅读顺序 | 多栏排版是否顺序正确 |
| 解析耗时 | 单页 / 单文件耗时 |
| 人工修正成本 | 老师审核是否省时间 |

---

## 13. 常见问题与注意事项

### 13.1 第一次运行很慢

可能原因：

- 首次下载模型
- 首次加载模型
- GPU / CPU 环境初始化

建议：

- 部署阶段提前下载模型。
- 服务启动后做一次预热解析。

---

### 13.2 国内网络无法下载模型

解决方向：

```bash
export MINERU_MODEL_SOURCE=modelscope
```

或者提前下载模型后使用本地模型源：

```bash
export MINERU_MODEL_SOURCE=local
```

---

### 13.3 解析效果不稳定

建议尝试：

- 换 `parse_method`：`auto` / `txt` / `ocr`
- 换语言参数：`ch` / `en`
- 对扫描件强制 OCR
- 对复杂版面尝试 VLM 后端
- 关闭不需要的公式 / 表格解析以减少干扰

---

### 13.4 不要直接入正式题库

MinerU 解析结果应作为候选内容。

正式流程必须包括：

```text
解析 → AI 辅助切题 → 人工审核 → 正式入库
```

原因：

- OCR 可能识别错字。
- 公式可能识别错误。
- 表格可能结构错位。
- 题目边界可能切分错误。
- 图片题可能丢失上下文。

---

### 13.5 Python 内部 API 不建议强绑定

MinerU 版本升级较快，内部 Python API 可能变化。

项目集成优先级建议：

1. **FastAPI 服务调用**：最适合系统集成。
2. **命令行 subprocess 调用**：适合 PoC / 后台任务。
3. **内部 Python API 直接调用**：除非明确锁定版本，否则不推荐生产强依赖。

---

## 14. 推荐落地方案

### 阶段 1：PoC

目标：验证 MinerU 对项目材料是否有效。

动作：

- 部署 MinerU 本地环境或 Docker。
- 选 5-10 份真实 PDF / 扫描材料。
- 使用 CLI 或 WebUI 解析。
- 记录准确率、耗时、人工修正量。

---

### 阶段 2：后端集成

目标：让系统可以上传文件并调用 MinerU。

动作：

- 启动 MinerU FastAPI 服务。
- 后端新增文档解析任务接口。
- 文件上传后异步调用 `/file_parse`。
- 保存 Markdown / JSON / 图片。
- 前端提供审核页面。

---

### 阶段 3：题库导入

目标：从解析结果中生成题目候选。

动作：

- Markdown / JSON 切题。
- AI 辅助识别题干、选项、答案、解析、知识点。
- 老师审核。
- 正式入题库。

---

## 15. 和本项目相关的结论

MinerU 很适合作为“AI国际学校智能作业系统”的 **PDF / 文档解析 Provider 候选**。

推荐定位：

```text
不是直接替代 OCR，
而是负责教材、教辅、真题、扫描 PDF 等复杂文档的结构化解析。
```

最推荐先做：

```text
MinerU + PDF 真题导入 PoC
```

验证通过后，再接入正式系统。

---

## 16. 参考链接

- MinerU GitHub：<https://github.com/opendatalab/MinerU>
- MinerU 官方文档：<https://opendatalab.github.io/MinerU/>
- MinerU 快速入门：<https://opendatalab.github.io/MinerU/quick_start>
- MinerU 命令行工具：<https://opendatalab.github.io/MinerU/usage/cli_tools>
- MinerU-Ecosystem：<https://github.com/opendatalab/MinerU-Ecosystem>

---

## 补充：物理公式需要二次清洗（2026-06-08）

在物理试卷、讲义、答案解析等场景中，MinerU / marker 的输出不应直接作为最终稿。推荐流程是：

1. 先用 MinerU / marker 抽取文本、图片、表格和基础 Markdown。
2. 再用脚本规则或人工规则把常见物理公式重写成 LaTeX。
3. 最后按题号整理 Markdown，统一题干、选项、答案、解析、分值结构。

核心原则：**PDF 转换工具只负责初稿抽取，物理公式必须二次清洗，不能完全指望自动识别。**

详见：

- [[PDF物理题转Markdown二次清洗流程]]
- [[物理公式LaTeX二次清洗规则]]
