# 第 9 章：综合实战与性能优化

本章目标：把前面学过的知识串成完整流程，并理解如何把 OpenCV 代码写得更接近真实项目。

## 知识点 1：实战项目一：文档扫描与透视矫正

### 说明

文档扫描是 OpenCV 里的经典小项目。它会综合使用：

- 灰度化
- 模糊与边缘检测
- 最大轮廓查找
- 多边形逼近
- 透视变换

只要你能把“纸张四个角”找出来，就可以把倾斜拍摄的文档拉正。

### 案例

下面的代码会先生成一张“倾斜拍摄的文档图”，再自动找出文档边界，并把它透视矫正成俯视效果。

### 完整代码

```python
import cv2
import numpy as np


def order_points(points):
    points = points.astype(np.float32)
    rect = np.zeros((4, 2), dtype=np.float32)
    s = points.sum(axis=1)
    diff = np.diff(points, axis=1)

    rect[0] = points[np.argmin(s)]
    rect[2] = points[np.argmax(s)]
    rect[1] = points[np.argmin(diff)]
    rect[3] = points[np.argmax(diff)]
    return rect


def four_point_transform(image, points):
    rect = order_points(points)
    tl, tr, br, bl = rect

    width_top = np.linalg.norm(tr - tl)
    width_bottom = np.linalg.norm(br - bl)
    max_width = int(max(width_top, width_bottom))

    height_left = np.linalg.norm(bl - tl)
    height_right = np.linalg.norm(br - tr)
    max_height = int(max(height_left, height_right))

    dst = np.array(
        [[0, 0], [max_width - 1, 0], [max_width - 1, max_height - 1], [0, max_height - 1]],
        dtype=np.float32,
    )

    matrix = cv2.getPerspectiveTransform(rect, dst)
    warped = cv2.warpPerspective(image, matrix, (max_width, max_height))
    return warped


def create_document_scene():
    document = np.full((300, 420, 3), 250, dtype=np.uint8)
    cv2.putText(document, "Invoice", (120, 60), cv2.FONT_HERSHEY_SIMPLEX, 1.4, (30, 30, 30), 3)
    for y in range(100, 250, 28):
        cv2.line(document, (35, y), (380, y), (80, 80, 80), 2)
    cv2.rectangle(document, (20, 20), (399, 279), (40, 40, 40), 3)

    scene = np.full((420, 620, 3), 70, dtype=np.uint8)
    src = np.float32([[0, 0], [419, 0], [419, 299], [0, 299]])
    dst = np.float32([[120, 60], [520, 90], [470, 350], [80, 320]])

    matrix = cv2.getPerspectiveTransform(src, dst)
    warped_doc = cv2.warpPerspective(document, matrix, (620, 420))
    warped_mask = cv2.warpPerspective(np.full((300, 420), 255, dtype=np.uint8), matrix, (620, 420))

    scene[warped_mask > 0] = warped_doc[warped_mask > 0]
    scene = cv2.GaussianBlur(scene, (3, 3), 0)
    return scene


def main():
    image = create_document_scene()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    blur = cv2.GaussianBlur(gray, (5, 5), 0)
    edges = cv2.Canny(blur, 50, 150)

    contours, _ = cv2.findContours(edges, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
    contours = sorted(contours, key=cv2.contourArea, reverse=True)

    screen_contour = None
    for contour in contours:
        peri = cv2.arcLength(contour, True)
        approx = cv2.approxPolyDP(contour, 0.02 * peri, True)
        if len(approx) == 4:
            screen_contour = approx.reshape(4, 2)
            break

    if screen_contour is None:
        raise RuntimeError("未找到四边形文档轮廓")

    warped = four_point_transform(image, screen_contour)

    outlined = image.copy()
    cv2.polylines(outlined, [screen_contour.astype(np.int32)], True, (0, 0, 255), 3)

    warped_small = cv2.resize(warped, (outlined.shape[1], outlined.shape[0]))
    panel = cv2.hconcat([outlined, warped_small])

    cv2.imshow("document scanner", panel)
    cv2.imwrite("document_scanner_result.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 2：实战项目二：圆形目标计数

### 说明

目标计数是图像处理中非常常见的任务。一个典型流程是：

1. 准备图像
2. 转灰度并二值化
3. 去噪或形态学处理
4. 查找轮廓
5. 按面积过滤并统计数量

### 案例

下面的代码会生成多枚不重叠的圆形目标，自动完成计数，并把结果写在图像上。

### 完整代码

```python
import cv2
import numpy as np


def create_coin_scene():
    image = np.full((340, 520, 3), 35, dtype=np.uint8)
    centers = [(80, 90), (210, 80), (360, 105), (120, 230), (270, 220), (420, 240)]
    radii = [35, 40, 32, 38, 42, 36]

    for center, radius in zip(centers, radii):
        cv2.circle(image, center, radius, (220, 220, 220), -1)
        cv2.circle(image, center, radius, (180, 180, 180), 3)

    return image


def main():
    image = create_coin_scene()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    _, binary = cv2.threshold(gray, 100, 255, cv2.THRESH_BINARY)

    kernel = np.ones((3, 3), np.uint8)
    binary = cv2.morphologyEx(binary, cv2.MORPH_OPEN, kernel)

    contours, _ = cv2.findContours(binary, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)

    output = image.copy()
    count = 0
    for contour in contours:
        area = cv2.contourArea(contour)
        if area < 1000:
            continue

        count += 1
        x, y, w, h = cv2.boundingRect(contour)
        cv2.drawContours(output, [contour], -1, (0, 255, 255), 2)
        cv2.rectangle(output, (x, y), (x + w, y + h), (0, 0, 255), 2)
        cv2.putText(output, f"#{count}", (x + 10, y + 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (255, 255, 255), 2)

    cv2.putText(output, f"count = {count}", (20, 35), cv2.FONT_HERSHEY_SIMPLEX, 1.0, (0, 255, 0), 2)

    panel = cv2.hconcat([image, cv2.cvtColor(binary, cv2.COLOR_GRAY2BGR), output])
    cv2.imshow("object counting", panel)
    cv2.imwrite("object_counting_result.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 3：性能分析与批处理模板

### 说明

当代码开始处理大量图像或实时视频时，“能跑”已经不够了，还要关心“跑得快不快”。性能优化的基本思路包括：

- 先缩小图像，再做复杂计算
- 尽量复用统一的处理管线
- 用 `cv2.TickMeter()` 或时间统计找到瓶颈
- 批量处理时建立稳定的输入输出目录

### 案例

下面的代码提供了一个简单的批处理模板。如果 `input_images` 目录里没有图片，它会自动生成测试图像；随后会执行统一处理流程，并打印总耗时与平均耗时。

### 完整代码

```python
from pathlib import Path
import cv2
import numpy as np


class BatchProcessor:
    def __init__(self, input_dir="input_images", output_dir="batch_outputs", target_width=640):
        self.input_dir = Path(input_dir)
        self.output_dir = Path(output_dir)
        self.target_width = target_width
        self.output_dir.mkdir(exist_ok=True)

    def generate_samples(self, count=5):
        samples = []
        for i in range(count):
            image = np.full((480, 720, 3), 30, dtype=np.uint8)
            cv2.circle(image, (120 + i * 80, 140 + i * 20), 45, (0, 220, 255), -1)
            cv2.rectangle(image, (320, 100 + i * 20), (560, 260 + i * 20), (255, 120, 0), 3)
            cv2.putText(image, f"sample {i}", (40, 420), cv2.FONT_HERSHEY_SIMPLEX, 1.4, (255, 255, 255), 3)
            samples.append((f"sample_{i}.png", image))
        return samples

    def load_images(self):
        patterns = ["*.png", "*.jpg", "*.jpeg", "*.bmp"]
        files = []
        for pattern in patterns:
            files.extend(sorted(self.input_dir.glob(pattern)))

        if not files:
            return self.generate_samples()

        images = []
        for path in files:
            image = cv2.imread(str(path))
            if image is not None:
                images.append((path.name, image))
        return images

    def process(self, image):
        h, w = image.shape[:2]
        if w > self.target_width:
            scale = self.target_width / w
            image = cv2.resize(image, None, fx=scale, fy=scale, interpolation=cv2.INTER_AREA)

        gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
        blur = cv2.GaussianBlur(gray, (5, 5), 0)
        edges = cv2.Canny(blur, 60, 150)
        result = cv2.cvtColor(edges, cv2.COLOR_GRAY2BGR)
        return result

    def run(self):
        items = self.load_images()
        meter = cv2.TickMeter()

        meter.start()
        for name, image in items:
            result = self.process(image)
            output_path = self.output_dir / f"processed_{name}"
            cv2.imwrite(str(output_path), result)
        meter.stop()

        total_ms = meter.getTimeMilli()
        avg_ms = total_ms / max(len(items), 1)

        print(f"processed images: {len(items)}")
        print(f"total time: {total_ms:.2f} ms")
        print(f"average per image: {avg_ms:.2f} ms")
        print(f"outputs saved to: {self.output_dir.resolve()}")


def main():
    processor = BatchProcessor()
    processor.run()


if __name__ == "__main__":
    main()
```
