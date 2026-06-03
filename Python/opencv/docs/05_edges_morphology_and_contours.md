# 第 5 章：边缘、形态学与轮廓分析

本章目标：学会从图像中提取结构信息，理解边缘、二值形态学和轮廓分析之间的关系。

## 知识点 1：Canny 边缘检测

### 说明

边缘通常对应图像灰度变化剧烈的位置。`cv2.Canny()` 是最经典的边缘检测方法之一，常用在文档检测、目标轮廓提取、车道线识别等任务前。

使用 Canny 时，通常会先做高斯模糊，以减小噪声对边缘结果的干扰。

### 案例

下面的代码创建一张带形状和文字的图像，先模糊，再做 Canny 边缘检测。

### 完整代码

```python
import cv2
import numpy as np


def gray_to_bgr(image):
    return cv2.cvtColor(image, cv2.COLOR_GRAY2BGR)


def create_source():
    image = np.full((280, 420, 3), 230, dtype=np.uint8)
    cv2.rectangle(image, (40, 50), (170, 220), (255, 120, 0), -1)
    cv2.circle(image, (290, 130), 65, (0, 180, 255), -1)
    cv2.putText(image, "EDGE", (120, 260), cv2.FONT_HERSHEY_SIMPLEX, 1.4, (30, 30, 30), 3)
    return image


def main():
    image = create_source()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    blur = cv2.GaussianBlur(gray, (5, 5), 0)
    edges = cv2.Canny(blur, 60, 150)

    panel = cv2.hconcat([image, gray_to_bgr(blur), gray_to_bgr(edges)])

    cv2.imshow("canny edge", panel)
    cv2.imwrite("canny_edge_demo.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 2：腐蚀、膨胀、开运算与闭运算

### 说明

形态学操作主要针对二值图像：

- 腐蚀：让白色区域变小
- 膨胀：让白色区域变大
- 开运算：先腐蚀后膨胀，常用于去小噪声
- 闭运算：先膨胀后腐蚀，常用于补小孔洞

### 案例

下面的代码会构造一张二值图，里面既有噪点，也有带缺口的目标，然后观察不同形态学操作的效果。

### 完整代码

```python
import cv2
import numpy as np


def gray_to_bgr(image):
    return cv2.cvtColor(image, cv2.COLOR_GRAY2BGR)


def create_binary_image():
    image = np.zeros((260, 420), dtype=np.uint8)
    cv2.rectangle(image, (40, 60), (170, 200), 255, -1)
    cv2.rectangle(image, (210, 70), (360, 180), 255, -1)
    cv2.circle(image, (285, 125), 22, 0, -1)
    cv2.line(image, (100, 60), (100, 200), 0, 8)

    rng = np.random.default_rng(7)
    for _ in range(500):
        y = rng.integers(0, image.shape[0])
        x = rng.integers(0, image.shape[1])
        image[y, x] = 255

    return image


def add_title(image, title):
    output = image.copy()
    cv2.putText(output, title, (10, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return output


def main():
    binary = create_binary_image()
    kernel = np.ones((5, 5), dtype=np.uint8)

    eroded = cv2.erode(binary, kernel, iterations=1)
    dilated = cv2.dilate(binary, kernel, iterations=1)
    opened = cv2.morphologyEx(binary, cv2.MORPH_OPEN, kernel)
    closed = cv2.morphologyEx(binary, cv2.MORPH_CLOSE, kernel)

    row1 = cv2.hconcat([
        add_title(gray_to_bgr(binary), "SOURCE"),
        add_title(gray_to_bgr(eroded), "ERODE"),
        add_title(gray_to_bgr(dilated), "DILATE"),
    ])
    row2 = cv2.hconcat([
        add_title(gray_to_bgr(opened), "OPEN"),
        add_title(gray_to_bgr(closed), "CLOSE"),
        np.full_like(row1[:, :binary.shape[1]], 255),
    ])

    panel = cv2.vconcat([row1, row2])

    cv2.imshow("morphology demo", panel)
    cv2.imwrite("morphology_demo.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 3：查找轮廓并计算面积、周长、外接矩形

### 说明

轮廓可以理解为目标边界的点集。`cv2.findContours()` 往往接在二值化或边缘提取之后，用于：

- 统计目标数量
- 计算面积和周长
- 获取外接矩形
- 找出最大轮廓

### 案例

下面的代码会生成多个几何图形，通过轮廓分析把每个目标圈出来，并在图像上标出面积和周长。

### 完整代码

```python
import cv2
import numpy as np


def create_shapes():
    binary = np.zeros((320, 460), dtype=np.uint8)
    cv2.rectangle(binary, (30, 50), (140, 210), 255, -1)
    cv2.circle(binary, (250, 130), 60, 255, -1)
    pts = np.array([[330, 240], [420, 180], [430, 290], [350, 300]], dtype=np.int32)
    cv2.fillPoly(binary, [pts], 255)
    return binary


def main():
    binary = create_shapes()
    contours, _ = cv2.findContours(binary, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)

    canvas = cv2.cvtColor(binary, cv2.COLOR_GRAY2BGR)

    for index, contour in enumerate(contours, start=1):
        area = cv2.contourArea(contour)
        perimeter = cv2.arcLength(contour, True)
        x, y, w, h = cv2.boundingRect(contour)

        cv2.drawContours(canvas, [contour], -1, (0, 255, 0), 2)
        cv2.rectangle(canvas, (x, y), (x + w, y + h), (0, 0, 255), 2)
        cv2.putText(
            canvas,
            f"#{index} A={int(area)} P={int(perimeter)}",
            (x, max(20, y - 10)),
            cv2.FONT_HERSHEY_SIMPLEX,
            0.55,
            (255, 0, 0),
            2,
        )

    cv2.imshow("contour analysis", canvas)
    cv2.imwrite("contour_analysis.png", canvas)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```
