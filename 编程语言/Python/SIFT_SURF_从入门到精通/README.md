# SIFT 与 SURF：从入门到精通

这是一套以 **Python + OpenCV** 为主线、以 C++ 为附录的局部特征学习包。你将从“关键点是什么”开始，逐步完成特征检测、描述、匹配、RANSAC、单应性估计、目标定位、图像拼接和算法评测。

> 推荐入口：先读 [8 周学习路线](00_开始这里/学习路线_8周.md)，然后运行 `python scripts/check_env.py`。

## 你最终应该会什么

- 能从尺度空间、DoG、Hessian、梯度直方图解释 SIFT 与 SURF。
- 能正确选择 BFMatcher/FLANN、距离度量、Lowe 比率阈值和 RANSAC 参数。
- 能区分“匹配数量多”和“几何上可信”的差别。
- 能完成平面目标定位、图像配准/拼接和 SIFT/ORB/AKAZE 性能对比。
- 能诊断重复纹理、弱纹理、模糊、光照、视角和尺度变化造成的失败。
- 知道何时继续使用经典特征，何时改用 ORB、AKAZE 或学习型特征。

## 五分钟开始

在本目录打开 PowerShell：

```powershell
python scripts/check_env.py
python scripts/run_demo.py --algorithm sift
python scripts/project_planar_localization.py
python scripts/benchmark.py --algorithms sift,orb,akaze
```

结果写入 `outputs/`。如果想在独立环境中学习：

```powershell
conda env create -f environment.yml
conda activate sift-surf-study
```

## 推荐学习顺序

1. `01_基础知识`：建立尺度、梯度、关键点、描述子和距离的直觉。
2. `02_SIFT`：逐步掌握 SIFT 的四个阶段与参数。
3. `04_特征匹配与几何`：把局部描述子变成可信的空间对应关系。
4. `05_实战项目`：完成目标定位、拼接与基准测试。
5. `03_SURF`：理解它如何用积分图和 Hessian 近似换取速度。
6. `06_进阶与工程化`：建立评测、选型和失败分析能力。
7. `08_练习与答案`：先独立完成，再核对答案。

## 目录导航

| 目录 | 内容 | 阶段验收 |
|---|---|---|
| `00_开始这里` | 路线、每日任务、环境与进度 | 能运行环境检查 |
| `01_基础知识` | 尺度空间、DoG、Hessian、描述子 | 能解释不变性和距离 |
| `02_SIFT` | 原理、OpenCV 用法、参数调优 | 能调参并解释关键点 |
| `03_SURF` | 原理、限制、独立编译指南 | 能比较 SIFT/SURF |
| `04_特征匹配与几何` | BF/FLANN、比率测试、RANSAC、单应性 | 能过滤错误匹配 |
| `05_实战项目` | 目标定位、拼接、基准测试 | 能交付可复现实验 |
| `06_进阶与工程化` | 评估、失败模式、现代替代方案 | 能做算法选型 |
| `07_论文与扩展资料` | 论文、官方文档、资源索引 | 能查阅一手资料 |
| `08_练习与答案` | 分级练习与答案 | 能独立完成综合题 |
| `09_样例与数据集` | OpenCV 示例图、数据集说明 | 能更换自己的图片 |

## 当前机器的重要说明

环境检查显示本机 OpenCV 4.12.0 的 SIFT 可用，但构建信息为 `Non-free algorithms: NO`，当前 Python 包不含可调用的 SURF。SURF 示例会友好提示并正常退出。需要真实运行 SURF 时，请看 [Windows 独立编译指南](03_SURF/SURF_Windows独立编译指南.md)，不要覆盖当前 Anaconda 环境。

## 常见错误速查

- **图片读不到**：确认路径存在；Windows 路径可加引号，尽量使用本包自带样例。
- **`descriptors is None`**：图像可能过于平坦、太小或阈值太高。
- **匹配很多但定位错误**：增加 Lowe 比率约束，检查 RANSAC 内点比例，不要只看匹配总数。
- **FLANN 报类型错误**：SIFT/SURF 使用 `float32` 描述子；ORB/AKAZE 二进制描述子通常使用 Hamming/BFMatcher。
- **SURF 不存在**：这通常是 OpenCV 构建选项，不是 `pip install opencv-contrib-python` 就一定能解决。
- **拼接画布过大或黑边多**：透视差太大、非平面场景或匹配退化；换图并检查单应性内点。

## 资料与版权

经典论文、OpenCV 官方页面和样例图均保留原始来源。视频、付费课程和版权受限内容只提供导读与链接。完整来源、访问日期、授权状态和 SHA-256 校验值见 `资源索引.csv` / `资源索引.xlsx`。

