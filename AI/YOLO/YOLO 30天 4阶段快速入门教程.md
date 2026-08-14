---
title: YOLO 30天 / 4阶段快速入门教程
date: 2026-08-14
tags: [AI, 计算机视觉, YOLO, 目标检测, 深度学习]
---

# YOLO 30天 / 4阶段快速入门教程

这是一份从零开始的 YOLO 学习文档，目标是在 30 天内走完一条完整路径：配置环境、跑通 Demo、理解目标检测核心原理、训练自己的数据集、完成一个小项目，并初步掌握导出、加速和部署。

主线工具选择 **Ultralytics YOLO**。它封装成熟、命令统一、社区资料多，适合从入门到项目实战。本文示例默认使用 `yolov8n.pt` 作为入门模型；后续如果 Ultralytics 发布新版本，学习方法仍然适用，只需要替换模型名和少量命令。

## 学习目标

30 天结束后，你应该能做到：

- 独立安装 YOLO 开发环境。
- 使用预训练模型检测图片、视频、摄像头画面。
- 理解边界框、置信度、IoU、NMS、mAP 等核心概念。
- 制作自己的 YOLO 格式数据集。
- 使用迁移学习训练自定义目标检测模型。
- 评估模型效果，并根据误检、漏检做第一轮优化。
- 将模型导出为 ONNX，并理解 TensorRT、OpenVINO、量化等部署优化方向。

## 适合人群

- Python 有基础，能运行脚本和安装依赖。
- 想学习目标检测，但不想一开始被论文和公式卡住。
- 想做自己的检测项目，例如安全帽检测、口罩检测、车辆检测、桌面物品检测、工件缺陷检测等。
- 想后续把模型接入摄像头、Web 服务、桌面工具或边缘设备。

## 环境建议

硬件：

- 有 NVIDIA 显卡最好，显存 6GB 以上更舒服。
- 没有显卡也可以学习，先用 CPU 跑 Demo，训练时用小模型和小数据集。
- 磁盘预留 20GB 以上，用于环境、数据、权重和训练结果。

软件：

- Windows 10/11 或 Linux。
- Anaconda 或 Miniconda。
- Python 3.10 或 3.11。
- PyTorch。
- Ultralytics。
- VS Code。
- 标注工具：LabelImg、CVAT、Roboflow、Label Studio 任意一种。

## 总体路线

| 阶段 | 时间 | 目标 | 产出 |
|---|---:|---|---|
| 第 1 阶段 | 第 1-7 天 | 配置环境、跑通 Demo、理解基本概念 | 图片/摄像头检测 Demo |
| 第 2 阶段 | 第 8-14 天 | 掌握 YOLO 核心原理、理解版本演进 | 原理笔记和思维导图 |
| 第 3 阶段 | 第 15-24 天 | 训练自己的数据集、完成小项目 | 自定义模型 `best.pt` |
| 第 4 阶段 | 第 25-30 天 | 模型导出、优化加速、简单部署 | ONNX 模型和部署计划 |

---

# 阶段 1：基础搭建与快速上手（第 1-7 天）

## 阶段目标

扫清环境障碍，快速跑通第一个 YOLO 模型，建立直观感受。这个阶段不要纠结太多公式，先看到模型工作起来。

阶段验收：

- 能创建并激活独立 Python 环境。
- 能安装 PyTorch 和 Ultralytics。
- 能检测图片、视频、摄像头。
- 能解释边界框、类别、置信度、阈值的基本含义。

## 第 1 天：环境地基：安装 Anaconda 并创建虚拟环境

### 今日目标

安装 Anaconda 或 Miniconda，创建独立环境 `yolo-env`。

### 为什么要虚拟环境

深度学习项目依赖复杂，PyTorch、OpenCV、NumPy、CUDA 等版本容易互相影响。虚拟环境可以把 YOLO 项目单独隔离，避免污染系统 Python 或其他项目。

### 操作步骤

安装 Anaconda 或 Miniconda 后，打开 Anaconda Prompt 或终端：

```bash
conda create -n yolo-env python=3.10 -y
conda activate yolo-env
python --version
```

创建学习目录：

```bash
mkdir yolo-30days
cd yolo-30days
```

推荐目录：

```text
yolo-30days/
  demos/
  datasets/
  notes/
  weights/
```

### 验收标准

终端前缀出现：

```text
(yolo-env)
```

并且 `python --version` 输出 Python 3.10.x 或 3.11.x。

### 常见坑

- 忘记执行 `conda activate yolo-env`，导致依赖装到 base 环境。
- Python 版本过新，部分包暂时不兼容。
- 项目路径包含过多中文、空格或特殊符号，后续脚本容易出路径问题。

## 第 2 天：安装 PyTorch

### 今日目标

安装 PyTorch，并确认 CPU 或 GPU 是否可用。

### 核心概念

Ultralytics YOLO 底层使用 PyTorch。PyTorch 负责张量计算、模型加载、训练和推理。

### 操作步骤

建议打开 PyTorch 官网，根据系统、显卡和 CUDA 版本生成安装命令：

```text
https://pytorch.org/get-started/locally/
```

如果有 NVIDIA 显卡，先检查驱动：

```bash
nvidia-smi
```

安装完成后验证：

```bash
python - <<'PY'
import torch
print('torch:', torch.__version__)
print('cuda available:', torch.cuda.is_available())
if torch.cuda.is_available():
    print('gpu:', torch.cuda.get_device_name(0))
PY
```

### 验收标准

- `import torch` 不报错。
- 有 NVIDIA 显卡时，`torch.cuda.is_available()` 输出 `True`。

### 常见坑

- 明明有显卡，但安装成了 CPU 版 PyTorch。
- 显卡驱动太旧，CUDA 不可用。
- Windows 上多个 Python 环境混用，实际运行的不是 `yolo-env`。

## 第 3 天：安装 Ultralytics YOLO

### 今日目标

安装 Ultralytics 工具包，确认 `yolo` 命令可用。

### 安装命令

```bash
pip install -U ultralytics
```

验证：

```bash
yolo version
```

或：

```bash
python - <<'PY'
import ultralytics
print(ultralytics.__version__)
PY
```

### Ultralytics 常用能力

- `yolo predict`：推理预测。
- `yolo train`：训练模型。
- `yolo val`：验证模型。
- `yolo export`：导出模型。
- Python API：在自己的程序里调用 YOLO。

### 验收标准

能看到 Ultralytics 版本号。

### 常见坑

- pip 下载慢，可以换国内镜像源。
- Linux 服务器没有图形界面，OpenCV 可能缺少 GUI 依赖。
- 安装完成后命令不可用，通常是环境未激活。

## 第 4 天：运行官方 Demo：图片检测

### 今日目标

使用预训练模型检测一张图片，并保存结果。

### 命令行方式

准备一张图片，例如 `test.jpg`：

```bash
yolo predict model=yolov8n.pt source=test.jpg conf=0.25 save=True
```

第一次运行会自动下载 `yolov8n.pt` 权重。

结果通常保存在：

```text
runs/detect/predict/
```

### Python 方式

创建 `demos/demo_image.py`：

```python
from ultralytics import YOLO

model = YOLO('yolov8n.pt')
results = model.predict('test.jpg', conf=0.25, save=True)

for result in results:
    print(result.boxes)
```

运行：

```bash
python demos/demo_image.py
```

### 参数解释

- `model=yolov8n.pt`：使用 YOLOv8 nano 模型，速度快，适合入门。
- `source=test.jpg`：输入图片。
- `conf=0.25`：置信度阈值，低于该值的检测框会被过滤。
- `save=True`：保存可视化结果。

### 验收标准

能在 `runs/detect/predict/` 中看到带检测框的结果图。

### 常见坑

- 图片路径写错。
- 首次下载模型失败，换网络后重试。
- 没有检测结果，不一定是程序错，可能模型不认识图片里的物体，或置信度阈值太高。

## 第 5 天：核心概念初识

### 今日目标

理解 YOLO 输出结果中的边界框、类别和置信度。

### 边界框 Bounding Box

边界框是模型预测的目标位置，通常是一个矩形。

常见表示方式：

- `x1, y1, x2, y2`：左上角和右下角坐标。
- `x_center, y_center, width, height`：中心点和宽高。

YOLO 标签格式通常是：

```text
class_id x_center y_center width height
```

这些坐标一般是归一化值，范围 0-1。

例如：

```text
0 0.512 0.430 0.230 0.180
```

意思是第 0 类目标，中心点在图片宽度 51.2%、高度 43.0% 的位置，框宽占图片 23.0%，框高占图片 18.0%。

### 类别 Class

类别是模型识别出的物体类型，例如 `person`、`car`、`dog`、`bottle`。

预训练 YOLO 通常基于 COCO 数据集，能识别常见的 80 类物体。

### 置信度 Confidence

置信度表示模型对某个检测结果的把握程度。它不是绝对真实概率，而是模型评分。

简单理解：

- 0.90：模型很确定。
- 0.50：模型比较确定。
- 0.20：模型不太确定，可能误检。

### 今日练习

用同一张图片分别运行：

```bash
yolo predict model=yolov8n.pt source=test.jpg conf=0.1 save=True
yolo predict model=yolov8n.pt source=test.jpg conf=0.5 save=True
yolo predict model=yolov8n.pt source=test.jpg conf=0.8 save=True
```

观察检测框数量变化。

### 验收标准

能用自己的话说明：模型检测到了什么、框在哪里、置信度代表什么。

## 第 6 天：调用摄像头检测

### 今日目标

把输入从图片改成摄像头，实现实时目标检测。

### 命令行方式

```bash
yolo predict model=yolov8n.pt source=0 show=True conf=0.5
```

`source=0` 表示默认摄像头。如果有多个摄像头，可尝试 `source=1`。

### Python 方式

创建 `demos/demo_camera.py`：

```python
from ultralytics import YOLO

model = YOLO('yolov8n.pt')
model.predict(source=0, show=True, conf=0.5)
```

### 今日练习

- 对准手机、水杯、键盘、椅子等常见物体。
- 调整 `conf=0.3`、`conf=0.5`、`conf=0.8`。
- 尝试 `yolov8n.pt` 和 `yolov8s.pt` 对比速度。

### 常见坑

- 摄像头被浏览器、微信、会议软件占用。
- WSL 中访问 Windows 摄像头不方便，建议用 Windows 原生 Python 环境运行。
- `show=True` 没窗口，可能是在无图形界面的服务器环境。

### 验收标准

能看到实时检测窗口，并能通过参数调整检测效果。

## 第 7 天：阶段复习与参数实验

### 今日目标

整理第一周成果，把 Demo 变成可重复运行的小项目。

### 推荐项目结构

```text
yolo-30days/
  demos/
    demo_image.py
    demo_camera.py
  images/
    test1.jpg
    test2.jpg
  outputs/
  notes/
    week1.md
```

### 复习清单

- conda 环境如何创建和激活。
- PyTorch 如何验证 GPU 可用。
- Ultralytics 如何安装。
- 图片检测、摄像头检测如何运行。
- `model`、`source`、`conf`、`imgsz`、`save` 分别是什么意思。

### 参数实验

```bash
yolo predict model=yolov8n.pt source=test.jpg conf=0.25 imgsz=640 save=True
yolo predict model=yolov8n.pt source=test.jpg conf=0.50 imgsz=640 save=True
yolo predict model=yolov8n.pt source=test.jpg conf=0.25 imgsz=1280 save=True
```

观察：

- `conf` 越高，框通常越少。
- `imgsz` 越大，小目标可能更容易检出，但速度更慢。
- 小模型速度快，大模型通常更准但更慢。

### 阶段产出

写一份 `notes/week1.md`：

```markdown
# YOLO 第一周笔记

## 我跑通了什么
- 图片检测：
- 摄像头检测：

## 我理解的概念
- 边界框：
- 置信度：
- 类别：

## 遇到的问题
- 问题：
- 解决：

## 下一阶段问题
- YOLO 为什么这么快？
- NMS 是什么？
- mAP 怎么看？
```

---

# 阶段 2：核心原理与版本演进（第 8-14 天）

## 阶段目标

理解 YOLO 为什么能检测物体，理清目标检测基础、YOLO 核心思想、IoU、NMS、mAP 和版本演进。

阶段验收：

- 能说清分类、检测、分割的区别。
- 能说清两阶段检测器和单阶段检测器的区别。
- 能理解 YOLO 的一次性预测思路。
- 能解释 IoU、NMS、Precision、Recall、mAP。

## 第 8 天：目标检测基础

### 今日目标

理解目标检测任务是什么，以及它和图像分类、实例分割的区别。

### 三类视觉任务

图像分类：回答“这张图是什么”。

```text
输入：一张猫的图片
输出：cat
```

目标检测：回答“图里有哪些物体，它们在哪里”。

```text
输出：
- person: [x1, y1, x2, y2]
- bottle: [x1, y1, x2, y2]
```

实例分割：回答“物体的精确轮廓在哪里”。

```text
输出：
- person mask
- bottle mask
```

### 两阶段检测器 vs 单阶段检测器

两阶段检测器，例如 Faster R-CNN：

1. 先找可能有物体的区域。
2. 再对这些区域分类和修正边界框。

特点：精度强，但流程复杂，速度通常较慢。

单阶段检测器，例如 YOLO、SSD、RetinaNet：

1. 直接在一次网络前向传播中预测类别和框。

特点：速度快，适合实时检测。

### 类比

两阶段检测像先圈出可疑区域，再逐个检查。YOLO 像扫一眼全图，同时说出“哪里有什么”。

### 今日练习

判断下面任务适合分类、检测还是分割：

- 判断图片中有没有火焰。
- 统计仓库里有多少箱货。
- 抠出人物轮廓换背景。

### 常见误区

误区：所有视觉任务都应该用目标检测。

正确理解：如果只需要判断有没有某类东西，分类可能更简单。如果需要精确轮廓，检测框又不够，需要分割。

## 第 9 天：YOLO 核心思想：把检测视为回归问题

### 今日目标

理解 YOLO 的核心：一次性预测边界框、类别和置信度。

### 核心思想

YOLO 全称 You Only Look Once，意思是“只看一次”。它把目标检测看成一个回归问题：输入图像，网络直接输出所有目标的位置和类别。

早期 YOLO 的经典解释：

1. 把图像划分成网格。
2. 每个网格负责预测中心落在该网格内的目标。
3. 每个位置预测多个候选框。
4. 最后筛选出最可信的框。

现代 YOLO 结构更复杂，但“一次性密集预测”的思想仍然是核心。

### 简化流程

```text
图像 -> 特征提取网络 -> 多尺度检测头 -> 候选框和类别分数 -> NMS -> 最终检测结果
```

### 多尺度检测

为了同时检测大物体和小物体，YOLO 会在不同尺度的特征图上预测：

- 高分辨率特征图：更适合小目标。
- 低分辨率特征图：更适合大目标。

### 今日练习

找一张人、车、小物体同时出现的图片，运行检测后观察：

- 小目标是否更容易漏检。
- 遮挡物体是否更容易误检。
- 大目标框是否更稳定。

### 常见误区

误区：YOLO 只是把图片切成格子，所以很粗糙。

正确理解：网格思想适合理解入门，但现代 YOLO 使用多尺度特征、改进检测头、anchor-free 等机制，实际能力远强于最早版本。

## 第 10 天：关键机制上：锚框与 IoU

### 今日目标

理解锚框和 IoU，它们是目标检测的基础概念。

### 锚框 Anchor Box

锚框是预设参考框。早期 YOLO 会在每个位置放几个不同形状的参考框，然后让模型在这些框基础上调整位置和尺寸。

例子：

- 行人：高瘦框。
- 车辆：宽扁框。
- 水杯：竖向矩形框。

YOLOv8 默认是 anchor-free，不再强依赖手工锚框。但理解锚框有助于理解 YOLOv3、YOLOv5 等版本。

### IoU：交并比

IoU 全称 Intersection over Union，用于衡量预测框和真实框的重叠程度。

```text
IoU = 预测框和真实框的交集面积 / 预测框和真实框的并集面积
```

IoU 越接近 1，预测框越准。

### 简单例子

```text
预测框 A：基本贴合目标，IoU = 0.85
预测框 B：只框住目标一半，IoU = 0.40
预测框 C：框到旁边背景，IoU = 0.05
```

### 今日练习

打开一张检测结果图，肉眼判断哪些框贴合度高，哪些框贴合度低。

### 常见误区

误区：框只要碰到目标就算检测正确。

正确理解：框的位置也要足够准，IoU 太低通常不算正确检测。

## 第 11 天：关键机制下：非极大值抑制 NMS

### 今日目标

理解为什么模型会产生很多重叠框，以及 NMS 如何筛选最终结果。

### 为什么需要 NMS

模型预测时会产生大量候选框。同一个物体附近可能出现多个相似框。NMS 的作用是只保留最可信的框，删除重复框。

### NMS 基本流程

1. 按置信度从高到低排序。
2. 选出最高分框作为保留框。
3. 删除和它 IoU 太高的其他框。
4. 对剩余框重复以上步骤。

### 关键参数

- `conf`：置信度阈值，低分框先过滤。
- `iou`：NMS 的 IoU 阈值，决定重叠到什么程度算重复。

实验命令：

```bash
yolo predict model=yolov8n.pt source=test.jpg conf=0.25 iou=0.7 save=True
yolo predict model=yolov8n.pt source=test.jpg conf=0.25 iou=0.3 save=True
```

### 今日练习

找一张人群、车流、货架等密集图片，调整 `iou` 参数，观察重复框变化。

### 常见误区

误区：NMS 阈值越低越好。

正确理解：阈值太低，密集目标可能被误删；阈值太高，同一目标可能保留多个重复框。

## 第 12 天：YOLO 版本简史

### 今日目标

了解 YOLO 从 v1 到现代版本的大致演进方向。

### 版本脉络

| 版本 | 关键词 | 重点理解 |
|---|---|---|
| YOLOv1 | 单阶段检测开端 | 把检测视为回归问题 |
| YOLOv2 | 锚框、BN、YOLO9000 | 训练和泛化增强 |
| YOLOv3 | Darknet-53、多尺度预测 | 经典工业版本之一 |
| YOLOv4 | CSPDarknet、Mosaic | 大量工程技巧整合 |
| YOLOv5 | PyTorch、工程化 | 易用、部署资料多 |
| YOLOv6/v7 | 工业优化、结构改进 | 速度和精度继续提升 |
| YOLOv8 | Ultralytics 新工具链 | 检测、分割、姿态等统一接口 |
| YOLOv9/v10/后续 | 新结构、端到端、部署友好 | 关注趋势，不必一开始深追 |

### 版本选择建议

- 入门和项目实战：优先 Ultralytics YOLOv8 或当前稳定版。
- 公司项目已有 YOLOv5：按项目版本学习。
- 做论文复现：再深入具体版本结构。

### 今日练习

整理自己的版本对比表，写清楚“我现在应该重点掌握哪个版本，为什么”。

## 第 13 天：评价指标：Precision、Recall、mAP

### 今日目标

学会看模型效果，不只凭肉眼感觉判断。

### 四种情况

- TP：正确检测到目标。
- FP：误检，把不存在的目标检测出来。
- FN：漏检，真实存在但没检测到。
- TN：目标检测里通常不重点讨论。

### Precision 精确率

```text
Precision = TP / (TP + FP)
```

它回答：模型检出来的东西里，有多少是真的。

精确率低，说明误检多。

### Recall 召回率

```text
Recall = TP / (TP + FN)
```

它回答：真实存在的目标里，有多少被模型找到了。

召回率低，说明漏检多。

### mAP 平均精度均值

mAP 是目标检测最常用综合指标。

常见指标：

- `mAP50`：IoU 阈值为 0.5 时的 mAP。
- `mAP50-95`：IoU 从 0.5 到 0.95 多个阈值平均，更严格。

### 解读方式

```text
Precision 高，Recall 低：模型保守，检出来多半对，但漏掉不少。
Precision 低，Recall 高：模型激进，找得多，但误检多。
mAP50 高，mAP50-95 低：大概能找到目标，但框不够精细。
```

### 常见误区

误区：mAP 越高，业务一定越好。

正确理解：业务还要看速度、模型大小、误检成本、漏检成本和部署平台。

## 第 14 天：知识梳理：画思维导图

### 今日目标

把阶段 2 的知识串起来。

### 推荐结构

```text
YOLO 目标检测
├── 输入输出
│   ├── 输入：图片/视频/摄像头
│   └── 输出：类别、边界框、置信度
├── 核心思想
│   ├── 单阶段检测
│   ├── 一次性预测
│   └── 多尺度特征
├── 关键机制
│   ├── 锚框/无锚框
│   ├── IoU
│   └── NMS
├── 评价指标
│   ├── Precision
│   ├── Recall
│   └── mAP
└── 工程流程
    ├── 数据
    ├── 标注
    ├── 训练
    ├── 验证
    └── 部署
```

### 阶段复盘问题

- YOLO 为什么适合实时检测？
- NMS 解决了什么问题？
- 为什么 `mAP50-95` 比 `mAP50` 更严格？
- 为什么小目标更难检测？
- 为什么同一个模型在不同场景效果差很多？

### 阶段产出

写 `notes/week2.md`，用自己的话解释目标检测、YOLO、IoU、NMS 和 mAP。

---

# 阶段 3：项目实战：训练自己的模型（第 15-24 天）

## 阶段目标

从 Demo 进入真实项目：收集图片、标注数据、组织数据集、写 YAML 配置、训练模型、评估结果、调参优化。

阶段验收：

- 完成一个自定义检测项目。
- 至少标注 100-200 张图片。
- 能用自己的数据训练 YOLO 模型。
- 能查看训练结果和验证指标。
- 能根据结果做一轮调优。

## 第 15 天：数据准备：确定第一个小项目

### 今日目标

选择一个简单、明确、容易收集图片的检测任务。

### 选题原则

第一个项目不要太难：

- 类别少，建议 1-3 类。
- 目标清楚，边界容易画。
- 图片容易收集。
- 场景相对固定。

推荐选题：

- 检测口罩：`mask` / `no-mask`。
- 检测饮料瓶：`bottle`。
- 检测硬币：`coin`。
- 检测安全帽：`helmet` / `no_helmet`。
- 检测桌面物品：`phone` / `cup` / `mouse`。

### 数据量建议

入门最低：

- 每类 50-100 张。
- 总计 100-200 张。

更稳定：

- 每类 300-500 张。
- 覆盖不同角度、光照、背景、距离。

### 数据采集注意

尽量覆盖：

- 正面、侧面、远景、近景。
- 明亮、昏暗、反光、遮挡。
- 单目标、多目标、无目标负样本。
- 与真实使用场景接近的图片。

### 今日产出

建立目录：

```text
datasets/my_first_yolo/raw/images/
datasets/my_first_yolo/raw/labels/
```

写项目定义：

```markdown
# 项目定义

任务：检测桌面上的饮料瓶
类别：bottle
使用场景：摄像头实时检测桌面物体
第一版数据量：150 张
```

## 第 16 天：数据标注：使用 LabelImg

### 今日目标

学会标注 YOLO 格式数据。

### 安装 LabelImg

```bash
pip install labelImg
labelImg
```

也可以使用 CVAT、Roboflow 或 Label Studio。

### YOLO 标签格式

每张图片对应一个同名 `.txt` 文件：

```text
images/raw/img001.jpg
labels/raw/img001.txt
```

标签内容：

```text
0 0.512 0.430 0.230 0.180
```

含义：

```text
class_id x_center y_center width height
```

坐标是 0-1 的归一化值。

### 标注规则

- 框尽量贴合目标边缘。
- 遮挡物体的标注规则要统一。
- 太模糊看不清的目标可以不标，但标准要一致。
- 类别名称不要频繁修改。

### 今日产出

至少标注 30-50 张图片，熟悉流程。

### 常见坑

- 保存成 VOC/XML，而不是 YOLO/txt。
- 图片名和标签名不一致。
- 类别 id 从 1 开始，YOLO 应该从 0 开始。
- 漏标图片里的目标，导致模型学到错误背景。

## 第 17 天：数据集划分

### 今日目标

把数据划分为训练集、验证集和测试集。

### 推荐比例

```text
train: 70%
val: 20%
test: 10%
```

小数据集也可以先用：

```text
train: 80%
val: 20%
```

### 目标目录

```text
datasets/my_first_yolo/
  images/
    train/
    val/
    test/
  labels/
    train/
    val/
    test/
```

### 简单划分脚本

创建 `split_dataset.py`：

```python
from pathlib import Path
import random
import shutil

random.seed(42)

root = Path('datasets/my_first_yolo')
raw_images = root / 'raw' / 'images'
raw_labels = root / 'raw' / 'labels'

images = sorted([p for p in raw_images.iterdir() if p.suffix.lower() in ['.jpg', '.jpeg', '.png']])
random.shuffle(images)

n = len(images)
train_end = int(n * 0.7)
val_end = int(n * 0.9)

splits = {
    'train': images[:train_end],
    'val': images[train_end:val_end],
    'test': images[val_end:],
}

for split, files in splits.items():
    (root / 'images' / split).mkdir(parents=True, exist_ok=True)
    (root / 'labels' / split).mkdir(parents=True, exist_ok=True)
    for img in files:
        label = raw_labels / f'{img.stem}.txt'
        shutil.copy2(img, root / 'images' / split / img.name)
        if label.exists():
            shutil.copy2(label, root / 'labels' / split / label.name)

print({k: len(v) for k, v in splits.items()})
```

### 常见坑

- 同一视频连续帧随机分到 train 和 val，验证结果虚高。
- 标签文件缺失。
- train 和 val 路径写反。

## 第 18 天：组织 YOLO 数据集结构

### 今日目标

按 Ultralytics 要求整理数据集。

### 标准结构

```text
datasets/my_first_yolo/
  images/
    train/
      001.jpg
    val/
      101.jpg
    test/
      151.jpg
  labels/
    train/
      001.txt
    val/
      101.txt
    test/
      151.txt
  data.yaml
```

图片和标签必须同名：

```text
images/train/001.jpg
labels/train/001.txt
```

### 检查脚本

```python
from pathlib import Path

root = Path('datasets/my_first_yolo')

for split in ['train', 'val', 'test']:
    img_dir = root / 'images' / split
    label_dir = root / 'labels' / split
    if not img_dir.exists():
        continue
    images = [p for p in img_dir.iterdir() if p.suffix.lower() in ['.jpg', '.jpeg', '.png']]
    missing = []
    for img in images:
        label = label_dir / f'{img.stem}.txt'
        if not label.exists():
            missing.append(img.name)
    print(split, 'images:', len(images), 'missing labels:', len(missing))
    if missing[:10]:
        print('examples:', missing[:10])
```

### 常见坑

- 类别 id 超出类别数量。
- 标注坐标不是 0-1 归一化值。
- 空标签文件不一定是错，无目标图片可以作为负样本。

## 第 19 天：编写 data.yaml

### 今日目标

创建 YOLO 训练所需的数据集配置文件。

### 单类别示例

`datasets/my_first_yolo/data.yaml`：

```yaml
path: datasets/my_first_yolo
train: images/train
val: images/val
test: images/test

names:
  0: bottle
```

### 多类别示例

```yaml
path: datasets/safety_helmet
train: images/train
val: images/val
test: images/test

names:
  0: helmet
  1: no_helmet
  2: person
```

### 配置解释

- `path`：数据集根目录。
- `train`：训练图片目录，相对 `path`。
- `val`：验证图片目录。
- `test`：测试图片目录，可选。
- `names`：类别 id 到类别名的映射。

### 格式验证

先跑 1 个 epoch，确认路径和标签没问题：

```bash
yolo detect train model=yolov8n.pt data=datasets/my_first_yolo/data.yaml epochs=1 imgsz=640 batch=4
```

### 常见坑

- YAML 缩进错误。
- Windows 路径用反斜杠导致转义问题，建议用 `/`。
- 类别顺序和标注工具里的顺序不一致。

## 第 20 天：首次训练：迁移学习

### 今日目标

使用预训练权重训练自己的检测模型。

### 为什么用迁移学习

从零训练需要大量数据和算力。预训练模型已经学会通用视觉特征，迁移学习是在这个基础上适配你的新类别。

### 训练命令

```bash
yolo detect train \
  model=yolov8n.pt \
  data=datasets/my_first_yolo/data.yaml \
  epochs=50 \
  imgsz=640 \
  batch=8 \
  project=runs/my_first_yolo \
  name=exp01
```

显存不足时：

```bash
yolo detect train model=yolov8n.pt data=datasets/my_first_yolo/data.yaml epochs=50 imgsz=640 batch=4
```

还不行就降低输入尺寸：

```bash
yolo detect train model=yolov8n.pt data=datasets/my_first_yolo/data.yaml epochs=50 imgsz=512 batch=2
```

### 模型大小选择

- `yolov8n.pt`：nano，最快，适合入门和低算力。
- `yolov8s.pt`：small，速度和精度平衡。
- `yolov8m.pt`：medium，更准但更慢。
- `yolov8l.pt` / `yolov8x.pt`：更大，训练成本更高。

第一版建议用 `yolov8n.pt` 或 `yolov8s.pt`。

### 训练结果

训练完成后通常生成：

```text
runs/my_first_yolo/exp01/
  weights/
    best.pt
    last.pt
  results.png
  confusion_matrix.png
  args.yaml
```

重点关注 `weights/best.pt`。

## 第 21 天：监控训练过程

### 今日目标

学会看训练日志和 loss 曲线。

### 常见 loss

- box loss：框位置损失。
- cls loss：类别分类损失。
- dfl loss：边界框分布相关损失。

### 正常训练表现

- train loss 逐步下降。
- val loss 大体下降，但可能波动。
- mAP 上升后趋于稳定。

### 异常表现

- loss 不下降：数据、标签、学习率或环境可能有问题。
- train 很好 val 很差：过拟合。
- 指标突然为 0：标签格式、类别 id 或路径可能出错。

### 训练记录模板

```markdown
# exp01 训练记录

模型：yolov8n.pt
数据量：train 120, val 30, test 20
epochs：50
batch：8
imgsz：640

结果：
- mAP50：
- mAP50-95：
- Precision：
- Recall：

问题：
-

下一轮改动：
-
```

## 第 22 天：模型评估与测试集验证

### 今日目标

用验证集或测试集评估模型，并用未见过图片检查真实效果。

### 验证命令

```bash
yolo detect val \
  model=runs/my_first_yolo/exp01/weights/best.pt \
  data=datasets/my_first_yolo/data.yaml \
  split=val
```

如果有 test 集：

```bash
yolo detect val \
  model=runs/my_first_yolo/exp01/weights/best.pt \
  data=datasets/my_first_yolo/data.yaml \
  split=test
```

### 预测测试图片

```bash
yolo detect predict \
  model=runs/my_first_yolo/exp01/weights/best.pt \
  source=test_images/ \
  conf=0.25 \
  save=True
```

### 人工检查清单

- 是否有明显漏检？
- 是否有明显误检？
- 框是否贴合目标？
- 暗光、遮挡、远距离、小目标效果如何？
- 真实摄像头画面是否比训练图片复杂？

### 结果解读

- 训练集效果好，真实图片差：数据分布不一致。
- 大目标准，小目标漏：需要更多小目标样本，或提高输入分辨率。
- 某个类别总错：类别定义混乱或样本太少。
- 背景误检多：需要加入负样本。

## 第 23 天：第一次调优

### 今日目标

根据第一轮结果做小范围调参。

### 调优原则

一次只改 1-2 个变量，否则不知道效果变化来自哪里。

### 可尝试方向

#### 1. 增加训练轮数

```bash
yolo detect train model=yolov8n.pt data=datasets/my_first_yolo/data.yaml epochs=100 imgsz=640 batch=8 name=exp02_epochs100
```

适合 loss 还在下降、mAP 还没稳定的情况。

#### 2. 换更大模型

```bash
yolo detect train model=yolov8s.pt data=datasets/my_first_yolo/data.yaml epochs=50 imgsz=640 batch=8 name=exp03_yolov8s
```

适合数据质量不错、小模型欠拟合的情况。

#### 3. 提高输入尺寸

```bash
yolo detect train model=yolov8n.pt data=datasets/my_first_yolo/data.yaml epochs=50 imgsz=960 batch=4 name=exp04_img960
```

适合小目标多的情况。

#### 4. 补充数据

多数时候，补数据比调参更有效。优先补：

- 漏检场景。
- 误检背景。
- 光照变化。
- 遮挡样本。
- 真实部署场景图片。

### 实验对比表

```markdown
| 实验 | 模型 | epochs | imgsz | batch | mAP50 | mAP50-95 | 速度 | 结论 |
|---|---|---:|---:|---:|---:|---:|---|---|
| exp01 | yolov8n | 50 | 640 | 8 |  |  |  | baseline |
| exp02 | yolov8n | 100 | 640 | 8 |  |  |  |  |
| exp03 | yolov8s | 50 | 640 | 8 |  |  |  |  |
```

### 常见坑

- 盲目调学习率，不先检查数据。
- 指标好看但业务不可用。
- 用测试集反复调参，导致测试集不再客观。

## 第 24 天：项目小结

### 今日目标

把第一个自定义 YOLO 项目整理成可复现工程。

### 推荐项目结构

```text
my_first_yolo_project/
  README.md
  data.yaml
  scripts/
    split_dataset.py
    predict_image.py
    predict_camera.py
  datasets/
    my_first_yolo/
  runs/
    exp01/
  weights/
    best.pt
  docs/
    training_log.md
```

### README 模板

````markdown
# 我的第一个 YOLO 目标检测项目

## 任务
检测：

## 类别
- 0:

## 环境
- Python:
- PyTorch:
- Ultralytics:
- GPU:

## 数据集
- train:
- val:
- test:

## 训练命令
```bash
yolo detect train model=yolov8n.pt data=data.yaml epochs=50 imgsz=640 batch=8
```

## 结果
- mAP50:
- mAP50-95:
- Precision:
- Recall:

## 问题和改进
-
````

### 阶段产出

你现在应该完成了一个真实的 YOLO 小项目。哪怕精度不高，也已经走通了目标检测项目最重要的闭环。

---

# 阶段 4：优化部署与扩展（第 25-30 天）

## 阶段目标

学习模型导出、推理加速、目标跟踪、实例分割和未来方向。重点建立部署意识：训练出来只是第一步，真正使用还要考虑速度、格式、平台和稳定性。

阶段验收：

- 能把 `.pt` 模型导出为 ONNX。
- 知道 TensorRT、OpenVINO、NCNN 等加速方向。
- 能跑简单的视频跟踪 Demo。
- 知道实例分割和检测的区别。
- 能规划下一个更贴近实际场景的项目。

## 第 25 天：模型导出 ONNX

### 今日目标

把 PyTorch `.pt` 权重导出为 ONNX 格式，为部署做准备。

### 为什么导出

`.pt` 适合 PyTorch 环境，但部署时可能希望在不同平台运行：

- Python 服务端。
- C++ 程序。
- Windows 桌面应用。
- 边缘设备。
- 推理引擎，例如 ONNX Runtime、TensorRT、OpenVINO。

ONNX 是比较通用的中间格式。

### 导出命令

```bash
yolo export \
  model=runs/my_first_yolo/exp01/weights/best.pt \
  format=onnx \
  imgsz=640
```

导出后通常得到：

```text
best.onnx
```

### Python 导出

```python
from ultralytics import YOLO

model = YOLO('runs/my_first_yolo/exp01/weights/best.pt')
model.export(format='onnx', imgsz=640)
```

### 验证 ONNX

```bash
yolo predict model=best.onnx source=test.jpg
```

### 常见坑

- 导出成功不代表部署性能一定好，还要测试推理速度。
- 动态输入尺寸、NMS 是否内置、后处理方式会影响部署代码。
- 部署环境和训练环境不同，结果可能有细微差异。

## 第 26 天：速度优化初探

### 今日目标

了解常见推理加速方案。

### 影响速度的因素

- 模型大小：n/s/m/l/x 越大通常越慢。
- 输入尺寸：`imgsz` 越大越慢。
- 硬件：CPU、NVIDIA GPU、Intel 核显、ARM 芯片差异很大。
- 推理引擎：PyTorch、ONNX Runtime、TensorRT 等速度不同。
- 后处理：NMS 也会消耗时间。

### 常见加速方向

#### 使用小模型

如果业务允许，`yolov8n` 或 `yolov8s` 更适合实时检测。

#### 降低输入尺寸

```bash
yolo predict model=best.pt source=video.mp4 imgsz=512
```

代价是小目标可能更容易漏检。

#### ONNX Runtime

适合跨平台部署，比直接 PyTorch 更通用。

#### TensorRT

适合 NVIDIA GPU 部署，推理速度通常很强。

```bash
yolo export model=best.pt format=engine imgsz=640
```

注意：TensorRT 对 CUDA、驱动、TensorRT 版本要求较严格。

#### 量化

把 FP32 模型转换为 FP16 或 INT8，降低计算和显存占用。

- FP16：速度更快，精度损失通常较小。
- INT8：速度和体积更好，但需要校准，精度可能下降。

### 今日练习

用同一张图片测试不同模型和尺寸：

```bash
yolo predict model=yolov8n.pt source=test.jpg imgsz=640
yolo predict model=yolov8s.pt source=test.jpg imgsz=640
yolo predict model=yolov8n.pt source=test.jpg imgsz=960
```

记录速度和效果。

## 第 27 天：扩展任务：目标跟踪

### 今日目标

学习 YOLO 检测和目标跟踪的区别，并跑通视频跟踪。

### 检测 vs 跟踪

检测只回答每一帧里有什么。

跟踪还要回答同一个物体在不同帧中是不是同一个。例如视频里有 3 个人，检测每一帧都能框出来；跟踪会给每个人分配 ID。

### 跟踪命令

```bash
yolo track model=yolov8n.pt source=video.mp4 tracker=botsort.yaml show=True
```

摄像头：

```bash
yolo track model=yolov8n.pt source=0 tracker=botsort.yaml show=True
```

### 应用场景

- 人流计数。
- 车辆计数。
- 工位动作监控。
- 物流包裹跟踪。
- 运动目标轨迹分析。

### 简单计数思路

1. 检测目标。
2. 跟踪目标 ID。
3. 设置一条计数线。
4. 当目标轨迹穿过计数线时，计数 +1。

### 常见坑

- 遮挡后 ID 切换。
- 多目标密集时跟踪混乱。
- 摄像头抖动影响轨迹稳定。

## 第 28 天：扩展任务：实例分割

### 今日目标

了解 YOLO 检测和 YOLO 分割的区别。

### 检测输出

```text
类别 + 边界框 + 置信度
```

适合：只需要知道物体位置大概范围。

### 分割输出

```text
类别 + 边界框 + mask + 置信度
```

适合：需要更精细的物体区域，例如抠图、面积计算、缺陷区域定位。

### 运行分割 Demo

```bash
yolo segment predict model=yolov8n-seg.pt source=test.jpg save=True
```

Python：

```python
from ultralytics import YOLO

model = YOLO('yolov8n-seg.pt')
model.predict('test.jpg', save=True)
```

### 检测和分割怎么选

选检测：

- 统计数量。
- 判断位置。
- 实时性要求高。
- 标注成本要低。

选分割：

- 需要精确轮廓。
- 需要面积、形状、边缘。
- 检测框不够精细。

### 常见坑

分割标注比检测标注更费时间。如果业务只需要框，不要一开始就做分割。

## 第 29 天：探索前沿：新模型和新能力

### 今日目标

了解 YOLO 之外和 YOLO 后续的发展方向，避免只会一个命令行工具。

### 关注方向

#### 更强的实时检测器

关注速度、精度和部署成本之间的平衡：

- 端到端检测，减少 NMS 依赖。
- 更轻量的 backbone。
- 更好的多尺度特征融合。
- 部署友好的结构设计。

#### 开放词汇检测 Open-Vocabulary Detection

传统 YOLO 只能检测训练过的类别。开放词汇检测希望通过文本提示检测新类别。

相关方向：

- Grounding DINO。
- YOLO-World。
- CLIP 相关视觉语言模型。

#### 零样本检测 Zero-Shot Detection

不专门训练某个类别，也能根据文本描述找到目标。

优点是灵活，缺点是速度、稳定性、部署成本通常更复杂。

#### 多模态视觉模型

大型视觉语言模型可以理解图像内容、回答问题、做 OCR、描述场景，但实时检测和精确框定位仍然有独立价值。

### 今日练习

阅读 2-3 篇最新文章或官方文档，记录：

```markdown
# 前沿模型阅读笔记

模型名称：
解决什么问题：
比普通 YOLO 强在哪里：
部署是否方便：
适合我的场景吗：
```

### 常见误区

误区：新模型一定比旧模型适合项目。

正确理解：项目落地更看稳定性、速度、数据闭环、维护成本。新模型适合研究和验证，不一定适合马上替换生产方案。

## 第 30 天：规划未来：从入门项目走向真实应用

### 今日目标

复盘 30 天学习成果，规划下一步深入方向。

### 复盘清单

你应该已经掌握：

- YOLO 环境搭建。
- 预训练模型推理。
- 图片、视频、摄像头检测。
- 目标检测核心概念。
- 数据标注和 YOLO 数据格式。
- 自定义数据集训练。
- 训练日志和指标分析。
- ONNX 导出。
- 跟踪、分割和部署方向。

### 下一步方向 1：做更真实的数据集

把数据从 100-200 张扩展到 1000+ 张，覆盖真实场景。精度提升往往来自数据，而不是神秘参数。

重点：

- 真实场景采集。
- 负样本补充。
- 错误样本回流。
- 标注规范统一。

### 下一步方向 2：做完整应用

例如：

- 摄像头实时检测 + Web 页面展示。
- 视频文件批量分析。
- 检测结果写入数据库。
- 异常检测后推送通知。
- 工厂/仓库/桌面场景自动计数。

推荐架构：

```text
摄像头/视频 -> YOLO 推理 -> 后处理/业务规则 -> 存储/告警/API -> 前端展示
```

### 下一步方向 3：学习部署优化

根据平台选择：

- NVIDIA GPU：TensorRT。
- Intel CPU/GPU：OpenVINO。
- 跨平台服务端：ONNX Runtime。
- 移动端/嵌入式：NCNN、TFLite、CoreML。
- Web 前端：ONNX Runtime Web、WebGPU。

### 下一步方向 4：读源码和论文

推荐顺序：

1. 先看 Ultralytics 文档和训练配置。
2. 再看 YOLOv1 论文，理解起点。
3. 再看 YOLOv3、YOLOv5/YOLOv8 结构讲解。
4. 最后关注最新检测器和开放词汇检测。

---

# 常用命令速查

## 环境

```bash
conda create -n yolo-env python=3.10 -y
conda activate yolo-env
pip install -U ultralytics
yolo version
```

## 验证 PyTorch

```bash
python - <<'PY'
import torch
print(torch.__version__)
print(torch.cuda.is_available())
if torch.cuda.is_available():
    print(torch.cuda.get_device_name(0))
PY
```

## 图片检测

```bash
yolo predict model=yolov8n.pt source=test.jpg conf=0.25 save=True
```

## 摄像头检测

```bash
yolo predict model=yolov8n.pt source=0 show=True conf=0.5
```

## 视频检测

```bash
yolo predict model=yolov8n.pt source=video.mp4 save=True
```

## 训练

```bash
yolo detect train model=yolov8n.pt data=datasets/my_first_yolo/data.yaml epochs=50 imgsz=640 batch=8
```

## 验证

```bash
yolo detect val model=runs/detect/train/weights/best.pt data=datasets/my_first_yolo/data.yaml
```

## 自训练模型预测

```bash
yolo detect predict model=runs/detect/train/weights/best.pt source=test_images/ conf=0.25 save=True
```

## 导出 ONNX

```bash
yolo export model=runs/detect/train/weights/best.pt format=onnx imgsz=640
```

## 跟踪

```bash
yolo track model=yolov8n.pt source=video.mp4 tracker=botsort.yaml show=True
```

## 分割

```bash
yolo segment predict model=yolov8n-seg.pt source=test.jpg save=True
```

---

# 常见问题排错表

| 问题 | 可能原因 | 解决方法 |
|---|---|---|
| `yolo` 命令不存在 | 环境没激活或安装失败 | `conda activate yolo-env` 后重装 `pip install -U ultralytics` |
| `torch.cuda.is_available()` 是 False | PyTorch 装成 CPU 版、驱动/CUDA 不匹配 | 去 PyTorch 官网按 CUDA 版本重新安装 |
| 训练爆显存 | batch 太大、imgsz 太大、模型太大 | 减小 `batch`，降低 `imgsz`，换 `yolov8n.pt` |
| mAP 为 0 | 标签格式错、类别 id 错、路径错 | 检查 data.yaml、标签 class_id、图片和标签同名 |
| 检测框很多误检 | conf 太低、负样本少、数据不够 | 提高 `conf`，补充负样本，清理错误标注 |
| 漏检严重 | 数据少、小目标难、模型太小 | 补数据，提高 `imgsz`，尝试 `yolov8s.pt` |
| 训练集好验证集差 | 过拟合或数据分布不同 | 补真实场景数据、减少训练轮数、检查数据划分 |
| LabelImg 保存格式不对 | 没切换 YOLO 格式 | 在工具里选择 YOLO 格式，确认生成 `.txt` |
| Windows 路径报错 | 反斜杠或空格问题 | YAML 中使用 `/`，路径尽量简单 |
| 摄像头打不开 | 摄像头被占用或编号不对 | 关闭占用程序，尝试 `source=1`、`source=2` |

---

# 后续深入路线

## 路线 A：工程应用方向

适合目标：把 YOLO 用到实际项目里。

学习内容：

- FastAPI + YOLO 推理服务。
- 摄像头 RTSP 拉流。
- 视频抽帧和批量检测。
- 结果写入 SQLite/PostgreSQL。
- Web 前端展示检测框。
- Docker 部署。

## 路线 B：算法理解方向

适合目标：深入理解目标检测算法。

学习内容：

- CNN、Backbone、Neck、Head。
- FPN/PAN 多尺度特征融合。
- Anchor-based 和 Anchor-free。
- Loss 设计。
- 数据增强策略。
- NMS 变体。

## 路线 C：模型优化方向

适合目标：让模型更快、更小、更适合部署。

学习内容：

- ONNX Runtime。
- TensorRT。
- FP16/INT8 量化。
- 剪枝和蒸馏。
- OpenVINO、NCNN、TFLite。
- 边缘设备性能测试。

## 路线 D：视觉大模型方向

适合目标：结合最新多模态能力。

学习内容：

- Grounding DINO。
- SAM / SAM2。
- YOLO-World。
- CLIP。
- 多模态模型和传统检测模型结合。

---

# 推荐资料

- Ultralytics 官方文档：`https://docs.ultralytics.com/`
- PyTorch 安装指南：`https://pytorch.org/get-started/locally/`
- COCO 数据集：`https://cocodataset.org/`
- LabelImg：`https://github.com/HumanSignal/labelImg`
- CVAT：`https://www.cvat.ai/`
- Roboflow：`https://roboflow.com/`

---

# 30 天学习打卡表

| 天数 | 主题 | 是否完成 | 输出物 |
|---:|---|---|---|
| 1 | 创建 Anaconda 虚拟环境 |  | `yolo-env` |
| 2 | 安装 PyTorch |  | GPU/CPU 验证结果 |
| 3 | 安装 Ultralytics |  | 版本号 |
| 4 | 图片检测 Demo |  | 检测结果图 |
| 5 | 核心概念初识 |  | 概念笔记 |
| 6 | 摄像头检测 |  | 实时检测运行成功 |
| 7 | 阶段复习 |  | `week1.md` |
| 8 | 目标检测基础 |  | 任务对比笔记 |
| 9 | YOLO 核心思想 |  | 流程图/笔记 |
| 10 | 锚框与 IoU |  | IoU 理解笔记 |
| 11 | NMS |  | 参数实验记录 |
| 12 | YOLO 版本简史 |  | 版本对比表 |
| 13 | 评价指标 |  | 指标解释笔记 |
| 14 | 知识梳理 |  | 思维导图 |
| 15 | 确定项目和收集数据 |  | raw 图片目录 |
| 16 | 数据标注 |  | YOLO txt 标签 |
| 17 | 数据集划分 |  | train/val/test |
| 18 | 组织数据结构 |  | images/labels 目录 |
| 19 | 编写 data.yaml |  | `data.yaml` |
| 20 | 首次训练 |  | `best.pt` |
| 21 | 监控训练 |  | `training_log.md` |
| 22 | 模型评估 |  | 验证指标 |
| 23 | 第一次调优 |  | 实验对比表 |
| 24 | 项目小结 |  | `README.md` |
| 25 | 导出 ONNX |  | `best.onnx` |
| 26 | 速度优化 |  | 速度对比表 |
| 27 | 目标跟踪 |  | track demo |
| 28 | 实例分割 |  | segmentation demo |
| 29 | 探索前沿 |  | 阅读笔记 |
| 30 | 规划未来 |  | 下一项目计划 |
