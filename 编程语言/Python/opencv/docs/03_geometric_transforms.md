# 第 3 章：缩放、旋转与几何变换

本章目标：理解图像的空间变换方式，掌握缩放、旋转、平移、仿射变换和透视变换。

## 知识点 1：缩放与插值方法

### 说明

图像缩放最常用的是 `cv2.resize()`。当图像变大时，常见插值方式有：

- `INTER_NEAREST`：最近邻，速度快，但边缘会明显锯齿化
- `INTER_LINEAR`：线性插值，常用默认值
- `INTER_CUBIC`：三次插值，放大后通常更平滑

### 案例

下面的程序会生成一张棋盘图，再用三种插值方式把它放大，对比效果差异。

### 完整代码

```python
import cv2
import numpy as np


def create_checkerboard(rows=8, cols=10, cell=20):
    board = np.zeros((rows * cell, cols * cell), dtype=np.uint8)
    for r in range(rows):
        for c in range(cols):
            if (r + c) % 2 == 0:
                board[r * cell:(r + 1) * cell, c * cell:(c + 1) * cell] = 255
    return board


def add_title(image, title):
    image = image.copy()
    cv2.putText(image, title, (10, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return image


def main():
    board = create_checkerboard()
    board_bgr = cv2.cvtColor(board, cv2.COLOR_GRAY2BGR)

    nearest = cv2.resize(board_bgr, None, fx=4, fy=4, interpolation=cv2.INTER_NEAREST)
    linear = cv2.resize(board_bgr, None, fx=4, fy=4, interpolation=cv2.INTER_LINEAR)
    cubic = cv2.resize(board_bgr, None, fx=4, fy=4, interpolation=cv2.INTER_CUBIC)

    panel = cv2.hconcat([
        add_title(nearest, "NEAREST"),
        add_title(linear, "LINEAR"),
        add_title(cubic, "CUBIC"),
    ])

    cv2.imshow("resize interpolation", panel)
    cv2.imwrite("resize_interpolation.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 2：旋转、翻转与平移

### 说明

平移和旋转都属于二维仿射变换。OpenCV 中常用 `cv2.warpAffine()` 完成这类操作。翻转则可以直接使用 `cv2.flip()`。

### 案例

下面的代码会先生成一张带有坐标感的图像，然后分别做旋转、水平翻转和右下平移。

### 完整代码

```python
import cv2
import numpy as np


def create_demo_image():
    image = np.full((320, 320, 3), 245, dtype=np.uint8)
    cv2.rectangle(image, (40, 40), (280, 280), (0, 180, 255), 3)
    cv2.line(image, (0, 160), (319, 160), (120, 120, 120), 1)
    cv2.line(image, (160, 0), (160, 319), (120, 120, 120), 1)
    cv2.putText(image, "A", (130, 180), cv2.FONT_HERSHEY_SIMPLEX, 2.5, (255, 0, 0), 5)
    return image


def add_title(image, title):
    output = image.copy()
    cv2.putText(output, title, (12, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return output


def main():
    image = create_demo_image()
    h, w = image.shape[:2]

    rotate_matrix = cv2.getRotationMatrix2D((w // 2, h // 2), 30, 1.0)
    rotated = cv2.warpAffine(image, rotate_matrix, (w, h))

    flipped = cv2.flip(image, 1)

    move_matrix = np.float32([[1, 0, 35], [0, 1, 45]])
    translated = cv2.warpAffine(image, move_matrix, (w, h))

    panel = cv2.hconcat([
        add_title(image, "ORIGINAL"),
        add_title(rotated, "ROTATED"),
        add_title(flipped, "FLIPPED"),
        add_title(translated, "TRANSLATED"),
    ])

    cv2.imshow("basic transforms", panel)
    cv2.imwrite("basic_transforms.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 3：仿射变换与透视变换

### 说明

仿射变换可以同时完成平移、旋转、缩放、倾斜，但它保持平行线平行。透视变换更进一步，能模拟“近大远小”的视觉效果，常用于文档矫正、俯视图拉正、车牌校正等场景。

### 案例

下面的代码先做一个三点仿射变换，再做一个四点透视变换，让你直观看到两类几何变换的差异。

### 完整代码

```python
import cv2
import numpy as np


def create_source_image():
    image = np.full((320, 420, 3), 255, dtype=np.uint8)
    cv2.rectangle(image, (60, 60), (360, 250), (0, 200, 255), -1)
    cv2.putText(image, "DOC", (150, 170), cv2.FONT_HERSHEY_SIMPLEX, 2.0, (40, 40, 40), 4)
    return image


def add_title(image, title):
    result = image.copy()
    cv2.putText(result, title, (10, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return result


def main():
    image = create_source_image()

    src_affine = np.float32([[60, 60], [360, 60], [60, 250]])
    dst_affine = np.float32([[40, 100], [380, 40], [90, 280]])
    affine_matrix = cv2.getAffineTransform(src_affine, dst_affine)
    affine_result = cv2.warpAffine(image, affine_matrix, (420, 320))

    src_perspective = np.float32([[60, 60], [360, 60], [360, 250], [60, 250]])
    dst_perspective = np.float32([[100, 40], [330, 80], [370, 260], [80, 240]])
    perspective_matrix = cv2.getPerspectiveTransform(src_perspective, dst_perspective)
    warped = cv2.warpPerspective(image, perspective_matrix, (420, 320))

    rectified_matrix = cv2.getPerspectiveTransform(dst_perspective, src_perspective)
    rectified = cv2.warpPerspective(warped, rectified_matrix, (420, 320))

    panel = cv2.hconcat([
        add_title(image, "SOURCE"),
        add_title(affine_result, "AFFINE"),
        add_title(warped, "PERSPECTIVE"),
        add_title(rectified, "RECTIFIED"),
    ])

    cv2.imshow("affine and perspective", panel)
    cv2.imwrite("affine_perspective_panel.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```
