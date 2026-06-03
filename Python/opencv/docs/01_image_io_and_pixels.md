# 第 1 章：图像读写、像素访问与 ROI

本章目标：掌握最基础也最常用的图像 I/O 操作，理解如何直接修改像素，并学会用 ROI 处理图像局部区域。

## 知识点 1：读取、显示、保存图像

### 说明

OpenCV 中最常用的三个入口函数是：

- `cv2.imread()`：读取图像
- `cv2.imshow()`：显示图像
- `cv2.imwrite()`：保存图像

初学时要注意两点：第一，`imread()` 读取失败会返回 `None`；第二，OpenCV 读彩色图时默认使用 BGR 顺序。

### 案例

下面的代码先生成一张测试图，再把它保存到磁盘，随后重新读回，显示彩色图和灰度图，并把灰度结果另存为文件。

### 完整代码

```python
import cv2
import numpy as np


def create_sample_image():
    image = np.zeros((320, 480, 3), dtype=np.uint8)
    image[:] = (30, 30, 30)
    cv2.rectangle(image, (40, 40), (200, 220), (255, 120, 0), -1)
    cv2.circle(image, (340, 160), 70, (0, 220, 255), -1)
    cv2.putText(image, "OpenCV", (120, 290), cv2.FONT_HERSHEY_SIMPLEX, 1.2, (255, 255, 255), 2)
    return image


def main():
    source = create_sample_image()
    cv2.imwrite("sample_input.png", source)

    color = cv2.imread("sample_input.png")
    if color is None:
        raise FileNotFoundError("sample_input.png 读取失败")

    gray = cv2.imread("sample_input.png", cv2.IMREAD_GRAYSCALE)

    cv2.imshow("color", color)
    cv2.imshow("gray", gray)
    cv2.imwrite("sample_gray.png", gray)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 2：像素访问与局部修改

### 说明

当你写 `image[y, x]` 时，拿到的是图像中第 `y` 行、第 `x` 列像素的 BGR 值。直接修改像素适合做简单的遮罩、颜色替换、区域标记，但大规模逐像素循环通常不如 NumPy 切片高效。

### 案例

下面的代码会生成一张深色背景图，在中间写入多个彩色块，并打印指定像素的 BGR 值。你会看到切片赋值比逐个点赋值更直观。

### 完整代码

```python
import cv2
import numpy as np


def main():
    image = np.zeros((260, 360, 3), dtype=np.uint8)
    image[:] = (20, 20, 20)

    image[30:110, 40:120] = (255, 0, 0)
    image[30:110, 140:220] = (0, 255, 0)
    image[30:110, 240:320] = (0, 0, 255)

    for i in range(0, 360, 20):
        image[160:220, i:i + 10] = (0, 255, 255)

    print("pixel at (50, 60):", image[50, 60])
    print("pixel at (180, 200):", image[180, 200])

    image[120:150, 100:260] = (255, 255, 255)

    cv2.imshow("pixel editing", image)
    cv2.imwrite("pixel_editing.png", image)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 3：ROI 裁剪、复制与拼接

### 说明

ROI 是 Region of Interest，也就是“感兴趣区域”。在 OpenCV 里，ROI 最常见的写法就是 NumPy 切片，例如 `roi = image[y1:y2, x1:x2]`。这是后面做人脸区域、目标区域、局部增强时的核心技巧。

### 案例

下面的代码会创建一张分区图像，裁剪出其中一个区域，把它贴到其他位置，并额外生成一个拼接结果，帮助你理解 ROI 的作用。

### 完整代码

```python
import cv2
import numpy as np


def create_layout():
    image = np.zeros((300, 420, 3), dtype=np.uint8)
    image[:] = (40, 40, 40)
    cv2.rectangle(image, (20, 30), (180, 140), (255, 120, 0), -1)
    cv2.rectangle(image, (220, 30), (390, 140), (0, 200, 120), -1)
    cv2.rectangle(image, (20, 170), (180, 270), (120, 80, 255), -1)
    cv2.putText(image, "ROI", (70, 100), cv2.FONT_HERSHEY_SIMPLEX, 1.2, (255, 255, 255), 2)
    return image


def main():
    image = create_layout()

    roi = image[30:140, 20:180].copy()
    image[170:280, 220:380] = roi

    left = image[:, :210]
    right = image[:, 210:]
    stitched = cv2.hconcat([left, right])

    cv2.imshow("roi copy", image)
    cv2.imshow("stitched", stitched)
    cv2.imwrite("roi_copy.png", image)
    cv2.imwrite("roi_stitched.png", stitched)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```
