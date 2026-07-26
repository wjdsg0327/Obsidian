# SURF 完整讲义

SURF（Speeded-Up Robust Features）由 Herbert Bay、Tinne Tuytelaars 和 Luc Van Gool 提出，目标是在保留尺度、旋转鲁棒性的同时提高计算速度。

## 1. Fast-Hessian 检测器

SURF 使用 Hessian 行列式定位斑点型结构：

```text
det(H) = Lxx * Lyy - Lxy^2
```

实际计算用盒式滤波器近似二阶高斯导数，并用权重修正不同滤波响应。由于盒式滤波器可通过积分图快速计算，滤波器扩大时计算量不会按面积同步增长。

## 2. 尺度空间策略

SIFT 主要通过逐级缩小图像构建 octave；SURF 更强调保持图像大小并增大盒式滤波器尺寸。两者都在位置与尺度中寻找稳定极值，但实现取舍不同。

## 3. 方向分配

SURF 在关键点邻域计算 Haar 小波响应 `dx`、`dy`，用滑动方向窗口寻找响应向量和最大的方向。若启用 upright 模式（U-SURF），跳过方向估计，速度更快，但失去旋转不变性。

## 4. 描述子

标准 SURF 将对齐后的邻域分成 `4 x 4` 个子区域。每个区域统计：

```text
sum(dx), sum(dy), sum(abs(dx)), sum(abs(dy))
```

因此标准描述子为 `4 x 4 x 4 = 64` 维。扩展模式进一步区分响应符号，得到 128 维描述子。

## 5. OpenCV 接口

启用相应模块后，典型接口是：

```python
surf = cv.xfeatures2d.SURF_create(
    hessianThreshold=400,
    nOctaves=4,
    nOctaveLayers=3,
    extended=False,
    upright=False,
)
keypoints, descriptors = surf.detectAndCompute(gray, None)
```

## 参数理解

- `hessianThreshold`：越大通常关键点越少、响应越强。
- `nOctaves`：覆盖的尺度范围。
- `nOctaveLayers`：每个 octave 的中间层数量。
- `extended=True`：64 维变 128 维，信息和成本都增加。
- `upright=True`：跳过方向分配，适合相机姿态较稳定的场景。

## SIFT 与 SURF 对比

| 维度 | SIFT | SURF |
|---|---|---|
| 检测核心 | DoG 极值 | Hessian 行列式近似 |
| 加速手段 | 金字塔与局部统计 | 积分图与盒式滤波 |
| 方向 | 梯度方向直方图 | Haar 小波响应 |
| 描述子 | 默认 128 维 | 默认 64 维，可扩展 128 维 |
| OpenCV 可用性 | 当前主模块可用 | 常需 contrib + nonfree 自编译 |
| 实际选型 | 质量稳定、资料丰富 | 需考虑构建与许可成本 |

## 学习提醒

“Speeded-Up”来自 SURF 提出时的软硬件背景。现代 OpenCV、CPU SIMD、GPU 和替代算法会改变实际速度排名，应在目标设备上测量，不要仅凭算法名称判断。

