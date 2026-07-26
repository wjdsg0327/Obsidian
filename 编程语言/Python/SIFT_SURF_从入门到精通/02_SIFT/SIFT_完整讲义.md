# SIFT 完整讲义

SIFT（Scale-Invariant Feature Transform）由 David G. Lowe 提出。经典流程可分为四步：尺度空间极值检测、关键点精定位、方向分配、描述子生成。

## 1. 尺度空间极值检测

构建多个 octave，每个 octave 内用若干高斯尺度表示图像，再计算 DoG。候选点必须大于或小于位置与相邻尺度的 26 个邻居。

- octave 解决大范围尺度变化。
- octave 内尺度层解决更细的尺度定位。
- DoG 以较低代价近似尺度归一化 LoG。

## 2. 关键点精定位

候选极值只是离散网格点。SIFT 用局部泰勒展开估计亚像素、亚尺度位置，并进行两类过滤：

- **低对比度过滤**：不稳定且易受噪声影响的响应被删除。
- **边缘响应过滤**：通过 Hessian 主曲率比排除沿边缘定位不稳定的点。

OpenCV 的 `contrastThreshold` 和 `edgeThreshold` 对应这两类行为。注意：`edgeThreshold` 越大，保留的边缘型候选通常越多，并不是“阈值越大越严格”。

## 3. 方向分配

在关键点尺度对应的邻域中统计梯度方向直方图，主峰成为关键点方向。接近主峰的次峰也可能生成一个同位置、同尺度但不同方向的关键点。因此关键点数量可能在方向阶段增加。

方向归一化让后续描述子具有旋转不变性。

## 4. 128 维描述子

以关键点方向对齐邻域，通常划分为 `4 x 4` 个子区域，每个子区域统计 8 个梯度方向桶：

```text
4 x 4 x 8 = 128
```

贡献按空间距离和方向进行插值，最后归一化、截断较大分量、再次归一化，以提升对光照和局部非线性变化的稳定性。

## OpenCV 最小用法

```python
import cv2 as cv

gray = cv.imread("image.png", cv.IMREAD_GRAYSCALE)
sift = cv.SIFT_create()
keypoints, descriptors = sift.detectAndCompute(gray, None)
print(len(keypoints), descriptors.shape)
```

## 重要输出

- `KeyPoint.pt`：浮点坐标。
- `KeyPoint.size`：关键点邻域直径，不是半径。
- `KeyPoint.angle`：方向，单位为度。
- `KeyPoint.response`：检测响应。
- `KeyPoint.octave`：编码后的 octave/层信息。
- `descriptors`：通常形状为 `(N, 128)`、类型 `float32`。

## SIFT 的优势

- 对尺度和旋转变化稳定。
- 对中等光照、视角和噪声变化有较好鲁棒性。
- 描述子区分性强，经典资料和实现成熟。

## SIFT 的局限

- 比 ORB 等二进制特征更耗时、占内存。
- 弱纹理、严重模糊、强反光、重复纹理和大视差仍会失败。
- 单应性只适合平面目标或纯旋转相机等条件；SIFT 本身不能消除模型假设错误。

## 论文阅读重点

第一次阅读经典论文时，重点看算法总览、尺度空间极值、方向分配、描述子构造和实验结论。第二遍再推导泰勒展开与边缘响应过滤。

