# 第 0 章：环境准备与第一段 OpenCV 代码

本章目标：完成 OpenCV 环境准备，理解 OpenCV 图像对象的基本形式，并跑通第一段可视化代码。

## 知识点 1：安装并验证 OpenCV 环境

### 说明

学习 OpenCV 的第一步不是直接记函数，而是先确认环境正确。只要 `cv2` 能成功导入，并且我们能创建一张图像保存到磁盘，就说明基础环境已经可用。

### 案例

下面的程序会打印 OpenCV 和 NumPy 版本，生成一张简单图片，并保存为 `hello_opencv.png`。如果程序能正常执行，说明你的环境已经准备好了。

### 完整代码

```python
import cv2
import numpy as np


def main():
    print("OpenCV version:", cv2.__version__)
    print("NumPy version:", np.__version__)

    canvas = np.zeros((220, 420, 3), dtype=np.uint8)
    canvas[:] = (35, 35, 35)

    cv2.putText(
        canvas,
        "Hello OpenCV",
        (35, 120),
        cv2.FONT_HERSHEY_SIMPLEX,
        1.2,
        (0, 255, 0),
        2,
        cv2.LINE_AA,
    )

    cv2.imwrite("hello_opencv.png", canvas)
    print("image shape:", canvas.shape)
    print("image dtype:", canvas.dtype)
    print("saved file: hello_opencv.png")


if __name__ == "__main__":
    main()
```

## 知识点 2：生成你的第一张测试图像

### 说明

OpenCV 中的图像本质上是一个 `numpy.ndarray`。对于彩色图像，通常形状是 `(高, 宽, 3)`，最后一个维度表示 BGR 三个颜色通道。只要会创建数组，就会创建图像。

### 案例

下面我们手工生成一张由蓝、绿、红三个色块和一条白线组成的测试图，并用窗口显示它。

### 完整代码

```python
import cv2
import numpy as np


def main():
    image = np.zeros((300, 450, 3), dtype=np.uint8)

    image[:, 0:150] = (255, 0, 0)
    image[:, 150:300] = (0, 255, 0)
    image[:, 300:450] = (0, 0, 255)

    cv2.line(image, (0, 150), (449, 150), (255, 255, 255), 4)

    cv2.imshow("first image", image)
    cv2.imwrite("first_blocks.png", image)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 3：理解图像的形状、类型和像素值

### 说明

理解 `shape`、`dtype`、像素访问方式，是后面做滤波、二值化、轮廓分析的基础。OpenCV 图像的数值通常是 `uint8`，范围是 `0~255`。对于彩色图像，一个像素会有三个值，顺序是 B、G、R。

### 案例

下面的代码会生成一张渐变图，打印图像属性，读取中心像素值，再手动修改一块区域的颜色。

### 完整代码

```python
import cv2
import numpy as np


def main():
    height, width = 240, 320
    y, x = np.indices((height, width))

    image = np.zeros((height, width, 3), dtype=np.uint8)
    image[..., 0] = x * 255 // width
    image[..., 1] = y * 255 // height
    image[..., 2] = (x + y) * 255 // (height + width)

    print("shape:", image.shape)
    print("dtype:", image.dtype)
    print("ndim:", image.ndim)
    print("size:", image.size)
    print("itemsize:", image.itemsize)

    center_pixel = image[height // 2, width // 2]
    print("center pixel (BGR):", center_pixel)

    image[80:160, 100:220] = (0, 255, 255)

    cv2.imshow("image info demo", image)
    cv2.imwrite("image_info_demo.png", image)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```
