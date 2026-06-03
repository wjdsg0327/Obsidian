# 第 4 章：滤波、阈值与图像增强

本章目标：掌握去噪、平滑、二值化和对比度增强的常用方法，为后续边缘、轮廓、分割打下基础。

## 知识点 1：均值滤波、高斯滤波与中值滤波

### 说明

滤波的核心目标通常是“去噪”或“平滑”。不同滤波器适合不同噪声：

- `cv2.blur()`：均值滤波，简单平均，速度快
- `cv2.GaussianBlur()`：高斯滤波，边缘保留通常比均值更自然
- `cv2.medianBlur()`：中值滤波，对椒盐噪声很有效

### 案例

下面的代码先生成一张干净图像，再添加高斯噪声和椒盐噪声，最后比较三种滤波结果。

### 完整代码

```python
import cv2
import numpy as np


def create_clean_image():
    image = np.full((260, 420, 3), 235, dtype=np.uint8)
    cv2.rectangle(image, (40, 40), (180, 200), (0, 180, 255), -1)
    cv2.circle(image, (310, 130), 65, (255, 120, 0), -1)
    cv2.putText(image, "NOISE", (110, 245), cv2.FONT_HERSHEY_SIMPLEX, 1.2, (40, 40, 40), 2)
    return image


def add_noise(image):
    rng = np.random.default_rng(42)
    noisy = image.astype(np.int16)
    gaussian_noise = rng.normal(0, 22, image.shape).astype(np.int16)
    noisy = np.clip(noisy + gaussian_noise, 0, 255).astype(np.uint8)

    for _ in range(1800):
        y = rng.integers(0, image.shape[0])
        x = rng.integers(0, image.shape[1])
        noisy[y, x] = (255, 255, 255) if rng.random() > 0.5 else (0, 0, 0)

    return noisy


def label(image, title):
    output = image.copy()
    cv2.putText(output, title, (10, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return output


def main():
    clean = create_clean_image()
    noisy = add_noise(clean)

    mean_blur = cv2.blur(noisy, (5, 5))
    gaussian_blur = cv2.GaussianBlur(noisy, (5, 5), 0)
    median_blur = cv2.medianBlur(noisy, 5)

    blank = np.full_like(clean, 255)

    row1 = cv2.hconcat([
        label(clean, "CLEAN"),
        label(noisy, "NOISY"),
        label(blank, "BLANK"),
    ])
    row2 = cv2.hconcat([
        label(mean_blur, "MEAN"),
        label(gaussian_blur, "GAUSSIAN"),
        label(median_blur, "MEDIAN"),
    ])
    panel = cv2.vconcat([row1, row2])

    cv2.imshow("filter comparison", panel)
    cv2.imwrite("filter_comparison.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 2：全局阈值、自适应阈值与 Otsu

### 说明

阈值化是把灰度图变成黑白图的常用手段。典型场景包括文字提取、目标分割、轮廓准备等。

- 全局阈值：整张图使用同一个阈值
- 自适应阈值：局部区域各自计算阈值，适合光照不均匀
- Otsu：自动寻找更合适的阈值

### 案例

下面的代码会生成一张“光照不均匀”的灰度图，然后分别使用三种阈值方法进行二值化。

### 完整代码

```python
import cv2
import numpy as np


def gray_to_bgr(image):
    return cv2.cvtColor(image, cv2.COLOR_GRAY2BGR)


def create_uneven_gray():
    height, width = 280, 420
    y, x = np.indices((height, width))
    background = 60 + x * 120 // width + y * 40 // height
    image = background.astype(np.uint8)

    cv2.putText(image, "OPENCV", (40, 110), cv2.FONT_HERSHEY_SIMPLEX, 1.5, 40, 4)
    cv2.putText(image, "Threshold", (90, 200), cv2.FONT_HERSHEY_SIMPLEX, 1.0, 25, 2)
    cv2.circle(image, (340, 80), 35, 30, -1)
    cv2.rectangle(image, (280, 150), (390, 240), 50, -1)
    return image


def add_title(image, title):
    output = image.copy()
    cv2.putText(output, title, (10, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return output


def main():
    gray = create_uneven_gray()

    _, global_binary = cv2.threshold(gray, 110, 255, cv2.THRESH_BINARY)
    adaptive = cv2.adaptiveThreshold(
        gray,
        255,
        cv2.ADAPTIVE_THRESH_GAUSSIAN_C,
        cv2.THRESH_BINARY,
        31,
        10,
    )
    _, otsu = cv2.threshold(gray, 0, 255, cv2.THRESH_BINARY + cv2.THRESH_OTSU)

    panel = cv2.hconcat([
        add_title(gray_to_bgr(gray), "SOURCE"),
        add_title(gray_to_bgr(global_binary), "GLOBAL"),
        add_title(gray_to_bgr(adaptive), "ADAPTIVE"),
        add_title(gray_to_bgr(otsu), "OTSU"),
    ])

    cv2.imshow("threshold methods", panel)
    cv2.imwrite("threshold_methods.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 3：直方图均衡化与 CLAHE

### 说明

当图像整体偏暗、对比度不足时，细节往往不明显。直方图均衡化会重新分配灰度，让亮暗差异更明显。CLAHE 则是在局部范围内做增强，通常比普通均衡化更稳定，不容易让整幅图看起来过于生硬。

### 案例

下面的代码会创建一张偏暗的灰度图，然后比较原图、普通均衡化、CLAHE 三种结果。

### 完整代码

```python
import cv2
import numpy as np


def gray_to_bgr(image):
    return cv2.cvtColor(image, cv2.COLOR_GRAY2BGR)


def create_dark_gray():
    height, width = 280, 420
    y, x = np.indices((height, width))
    base = 20 + x * 50 // width + y * 40 // height
    image = base.astype(np.uint8)

    cv2.circle(image, (110, 120), 60, 100, -1)
    cv2.rectangle(image, (230, 70), (360, 210), 85, -1)
    cv2.putText(image, "LOW LIGHT", (60, 250), cv2.FONT_HERSHEY_SIMPLEX, 1.0, 120, 2)
    return image


def add_title(image, title):
    output = image.copy()
    cv2.putText(output, title, (10, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return output


def main():
    gray = create_dark_gray()

    equalized = cv2.equalizeHist(gray)
    clahe = cv2.createCLAHE(clipLimit=2.0, tileGridSize=(8, 8)).apply(gray)

    panel = cv2.hconcat([
        add_title(gray_to_bgr(gray), "ORIGINAL"),
        add_title(gray_to_bgr(equalized), "EQUALIZE"),
        add_title(gray_to_bgr(clahe), "CLAHE"),
    ])

    cv2.imshow("contrast enhancement", panel)
    cv2.imwrite("contrast_enhancement.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```
