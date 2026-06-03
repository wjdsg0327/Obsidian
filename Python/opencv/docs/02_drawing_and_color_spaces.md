# 第 2 章：绘图、文字与颜色空间

本章目标：学会在图像上绘制图形和文字，理解 OpenCV 中的 BGR、灰度、HSV 等颜色表示方式。

## 知识点 1：绘制线条、矩形、圆和多边形

### 说明

OpenCV 提供了丰富的绘图函数，例如：

- `cv2.line()`：画线
- `cv2.rectangle()`：画矩形
- `cv2.circle()`：画圆
- `cv2.ellipse()`：画椭圆
- `cv2.polylines()`：画多边形

这些函数在调试、标注检测结果、制作测试图时都非常有用。

### 案例

下面的代码会在同一张画布上画出多种图形，让你熟悉参数的基本用法。

### 完整代码

```python
import cv2
import numpy as np


def main():
    canvas = np.full((420, 520, 3), 240, dtype=np.uint8)

    cv2.line(canvas, (30, 40), (220, 180), (255, 0, 0), 3)
    cv2.rectangle(canvas, (260, 40), (470, 170), (0, 255, 0), 3)
    cv2.circle(canvas, (120, 300), 70, (0, 0, 255), -1)
    cv2.ellipse(canvas, (340, 300), (100, 50), 30, 0, 300, (255, 180, 0), 4)

    points = np.array([[260, 220], [440, 250], [390, 380], [240, 340]], dtype=np.int32)
    cv2.polylines(canvas, [points], True, (120, 0, 200), 4)

    cv2.imshow("drawing primitives", canvas)
    cv2.imwrite("drawing_primitives.png", canvas)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 2：添加文字、水印和透明叠加

### 说明

OpenCV 的 `cv2.putText()` 可以把文字写到图像上，而 `cv2.addWeighted()` 可以用来做半透明叠加。二者结合后，就能做标题、水印、区域高亮、结果标注等常见效果。

### 案例

下面的代码先生成一张背景图，再叠加一块半透明信息面板，并写上标题和说明文字。

### 完整代码

```python
import cv2
import numpy as np


def main():
    image = np.zeros((320, 520, 3), dtype=np.uint8)

    for y in range(image.shape[0]):
        color = 50 + y * 180 // image.shape[0]
        image[y, :] = (color // 2, color, 200)

    overlay = image.copy()
    cv2.rectangle(overlay, (30, 30), (490, 130), (20, 20, 20), -1)
    result = cv2.addWeighted(overlay, 0.6, image, 0.4, 0)

    cv2.putText(result, "OpenCV Overlay Demo", (50, 75), cv2.FONT_HERSHEY_SIMPLEX, 1.0, (0, 255, 255), 2)
    cv2.putText(result, "Transparent panel + text", (50, 110), cv2.FONT_HERSHEY_SIMPLEX, 0.7, (255, 255, 255), 2)

    cv2.imshow("text and overlay", result)
    cv2.imwrite("text_overlay_demo.png", result)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 3：BGR、灰度、HSV 与通道分离

### 说明

OpenCV 默认使用 BGR，不是 RGB。除此之外，灰度图适合做阈值、边缘等计算，HSV 更适合按颜色范围做分割。通道分离后，你能看到某种颜色在图像中“有多强”。

### 案例

下面的代码构造一张彩色渐变图，然后转换成灰度和 HSV，并分别显示 B、G、R、H、S、V 的效果。

### 完整代码

```python
import cv2
import numpy as np


def gray_to_bgr(channel):
    return cv2.cvtColor(channel, cv2.COLOR_GRAY2BGR)


def main():
    height, width = 240, 360
    y, x = np.indices((height, width))

    image = np.zeros((height, width, 3), dtype=np.uint8)
    image[..., 0] = x * 255 // width
    image[..., 1] = y * 255 // height
    image[..., 2] = (x + y) * 255 // (height + width)

    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    hsv = cv2.cvtColor(image, cv2.COLOR_BGR2HSV)

    b, g, r = cv2.split(image)
    h, s, v = cv2.split(hsv)

    row1 = cv2.hconcat([image, gray_to_bgr(gray), gray_to_bgr(b), gray_to_bgr(g)])
    row2 = cv2.hconcat([gray_to_bgr(r), gray_to_bgr(h), gray_to_bgr(s), gray_to_bgr(v)])
    panel = cv2.vconcat([row1, row2])

    cv2.imshow("color spaces and channels", panel)
    cv2.imwrite("color_spaces_panel.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```
