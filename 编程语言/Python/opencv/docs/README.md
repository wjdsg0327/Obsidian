# Python OpenCV 学习路线

这是一套按章节拆分的 Python OpenCV 教学文档，目标是带你从 0 开始，逐步走到进阶与高级应用。整套内容坚持统一结构：每个知识点都包含 **说明**、**案例**、**完整代码**，方便你边学边练。

## 学习建议

1. 按章节顺序阅读，先建立图像处理直觉，再进入特征、分割和综合实战。
2. 每读完一个知识点，亲手运行一次完整代码，并尝试修改参数。
3. 对 `cv2.imshow()` 类示例，运行后按任意键关闭窗口。
4. 如果你的环境没有 GUI，可以把 `imshow()` 改成 `imwrite()`，先观察输出图片。

## 环境依赖

```bash
pip install opencv-python numpy
```

如果你之后要学习更丰富的扩展模块，也可以安装：

```bash
pip install opencv-contrib-python numpy
```

## 章节目录

| 章节 | 文件 | 内容重点 | 难度 |
|------|------|---------|------|
| 0 | `00_setup_and_first_program.md` | 安装、验证环境、第一段代码 | 入门 |
| 1 | `01_image_io_and_pixels.md` | 图像读写、像素访问、ROI | 入门 |
| 2 | `02_drawing_and_color_spaces.md` | 绘图、文字、颜色空间、通道 | 入门 |
| 3 | `03_geometric_transforms.md` | 缩放、旋转、仿射、透视变换 | 入门到进阶 |
| 4 | `04_filtering_threshold_and_enhancement.md` | 滤波、阈值、直方图增强 | 进阶 |
| 5 | `05_edges_morphology_and_contours.md` | 边缘、形态学、轮廓分析 | 进阶 |
| 6 | `06_video_and_camera.md` | 摄像头、视频读写、实时处理 | 进阶 |
| 7 | `07_features_matching_and_hough.md` | ORB 特征、匹配、霍夫变换、模板匹配 | 进阶到高级 |
| 8 | `08_segmentation_and_advanced_processing.md` | 连通域、分水岭、GrabCut | 高级 |
| 9 | `09_projects_and_optimization.md` | 综合实战、性能优化、工程化思路 | 高级 |

## 推荐学习路径

- 第 0 到 2 章：先理解 OpenCV 中“图像就是 NumPy 数组”这件事。
- 第 3 到 5 章：建立图像变换、滤波、二值化、轮廓分析的基础能力。
- 第 6 到 8 章：进入实时处理、特征提取、分割等更接近实际项目的主题。
- 第 9 章：把前面知识串起来，形成“能做项目”的能力。

## 你会学到什么

- 如何读写、显示、修改图像和视频
- 如何做缩放、旋转、透视校正等几何变换
- 如何做模糊、去噪、阈值分割、边缘提取
- 如何分析轮廓、提取特征、做模板匹配
- 如何完成文档扫描、目标计数等小型实战项目
- 如何优化性能，写出更适合复用的 OpenCV 代码

## 使用方式

- 建议为每个章节新建一个单独的实验脚本，把文档中的“完整代码”复制进去运行。
- 如果某段代码里用到了输出文件，例如 `output.png`、`output.mp4`，运行后可以直接在当前目录查看结果。
- 高级章节中的算法不要求一次全部掌握，重要的是先理解输入、输出和处理流程。

## 配套材料

- 示例脚本索引：`../examples/README.md`
- 练习题索引：`../exercises/README.md`
- 基础依赖文件：`../requirements.txt`
- 扩展模块依赖：`../requirements-contrib.txt`
