# 第 8 章：图像分割与高级处理

本章目标：学习如何把图像中的目标从背景中分离出来，并理解连通域、分水岭和 GrabCut 这类更高级的处理思路。

## 知识点 1：连通域分析

### 说明

当一张二值图里有多个彼此分离的白色目标时，我们可以通过连通域分析为它们编号，并统计面积、边界框、质心等信息。这个思路非常适合做目标计数、字符分割、简单颗粒分析。

### 案例

下面的代码会生成多个彼此分离的目标，然后使用 `connectedComponentsWithStats()` 给每个连通区域着色并标注编号与面积。

### 完整代码

```python
import cv2
import numpy as np


def create_binary_scene():
    image = np.zeros((300, 420), dtype=np.uint8)
    cv2.circle(image, (80, 90), 35, 255, -1)
    cv2.rectangle(image, (180, 40), (290, 150), 255, -1)
    cv2.ellipse(image, (330, 220), (50, 30), 20, 0, 360, 255, -1)
    cv2.circle(image, (120, 230), 28, 255, -1)
    return image


def color_for_label(label):
    return (
        int((label * 70) % 255),
        int((label * 130) % 255),
        int((label * 200) % 255),
    )


def main():
    binary = create_binary_scene()
    count, labels, stats, centroids = cv2.connectedComponentsWithStats(binary, connectivity=8)

    color_map = np.zeros((binary.shape[0], binary.shape[1], 3), dtype=np.uint8)

    for label in range(1, count):
        color_map[labels == label] = color_for_label(label)

        x, y, w, h, area = stats[label]
        cx, cy = centroids[label]
        cv2.rectangle(color_map, (x, y), (x + w, y + h), (255, 255, 255), 2)
        cv2.putText(
            color_map,
            f"#{label} A={area}",
            (x, max(20, y - 8)),
            cv2.FONT_HERSHEY_SIMPLEX,
            0.55,
            (255, 255, 255),
            2,
        )
        cv2.circle(color_map, (int(cx), int(cy)), 4, (0, 0, 255), -1)

    panel = cv2.hconcat([cv2.cvtColor(binary, cv2.COLOR_GRAY2BGR), color_map])
    cv2.imshow("connected components", panel)
    cv2.imwrite("connected_components.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 2：分水岭算法分离粘连目标

### 说明

当目标彼此接触时，仅靠轮廓可能很难把它们分开。分水岭算法会把图像看成地形图，再从“确定的前景”和“确定的背景”往中间推进，最后把粘连区域拆开。

这是硬币分割、细胞分割、颗粒分析中的常见方法。

### 案例

下面的代码会生成几枚互相接触的“硬币”，然后使用阈值、距离变换和分水岭把它们分离开。

### 完整代码

```python
import cv2
import numpy as np


def create_touching_objects():
    image = np.full((320, 420, 3), 40, dtype=np.uint8)
    cv2.circle(image, (120, 150), 70, (220, 220, 220), -1)
    cv2.circle(image, (200, 170), 70, (220, 220, 220), -1)
    cv2.circle(image, (290, 150), 70, (220, 220, 220), -1)
    cv2.circle(image, (220, 240), 65, (220, 220, 220), -1)
    return image


def main():
    image = create_touching_objects()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

    _, binary = cv2.threshold(gray, 0, 255, cv2.THRESH_BINARY + cv2.THRESH_OTSU)
    kernel = np.ones((3, 3), np.uint8)
    opening = cv2.morphologyEx(binary, cv2.MORPH_OPEN, kernel, iterations=2)

    sure_bg = cv2.dilate(opening, kernel, iterations=3)
    dist = cv2.distanceTransform(opening, cv2.DIST_L2, 5)
    _, sure_fg = cv2.threshold(dist, 0.45 * dist.max(), 255, 0)
    sure_fg = sure_fg.astype(np.uint8)
    unknown = cv2.subtract(sure_bg, sure_fg)

    count, markers = cv2.connectedComponents(sure_fg)
    markers = markers + 1
    markers[unknown == 255] = 0

    markers = cv2.watershed(image.copy(), markers)
    result = image.copy()
    result[markers == -1] = (0, 0, 255)

    display_dist = cv2.normalize(dist, None, 0, 255, cv2.NORM_MINMAX).astype(np.uint8)
    panel = cv2.hconcat([
        image,
        cv2.cvtColor(binary, cv2.COLOR_GRAY2BGR),
        cv2.cvtColor(display_dist, cv2.COLOR_GRAY2BGR),
        result,
    ])

    print("connected foreground components:", count - 1)
    cv2.imshow("watershed demo", panel)
    cv2.imwrite("watershed_demo.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 3：GrabCut 前景提取

### 说明

GrabCut 适合在“我大概知道目标在哪”的情况下，把前景从背景中抠出来。它通常需要一个初始矩形，算法会在矩形内外分别估计前景和背景模型，然后反复优化。

### 案例

下面的代码会生成一张带背景和主体的示意图，用一个大致包住主体的矩形初始化 GrabCut，最后得到前景提取结果。

### 完整代码

```python
import cv2
import numpy as np


def create_scene():
    height, width = 320, 460
    y, x = np.indices((height, width))

    image = np.zeros((height, width, 3), dtype=np.uint8)
    image[..., 0] = 40 + x * 60 // width
    image[..., 1] = 120 + y * 80 // height
    image[..., 2] = 40 + (x + y) * 30 // (width + height)

    cv2.circle(image, (230, 160), 85, (20, 60, 220), -1)
    cv2.rectangle(image, (205, 90), (255, 235), (0, 220, 255), -1)
    cv2.putText(image, "FG", (205, 175), cv2.FONT_HERSHEY_SIMPLEX, 1.4, (255, 255, 255), 3)
    return image


def main():
    image = create_scene()
    mask = np.zeros(image.shape[:2], np.uint8)
    bgd_model = np.zeros((1, 65), np.float64)
    fgd_model = np.zeros((1, 65), np.float64)

    rect = (120, 60, 220, 210)
    cv2.grabCut(image, mask, rect, bgd_model, fgd_model, 5, cv2.GC_INIT_WITH_RECT)

    foreground_mask = np.where(
        (mask == cv2.GC_FGD) | (mask == cv2.GC_PR_FGD),
        255,
        0,
    ).astype("uint8")

    result = cv2.bitwise_and(image, image, mask=foreground_mask)

    display = image.copy()
    x, y, w, h = rect
    cv2.rectangle(display, (x, y), (x + w, y + h), (255, 255, 255), 2)

    panel = cv2.hconcat([
        display,
        cv2.cvtColor(foreground_mask, cv2.COLOR_GRAY2BGR),
        result,
    ])

    cv2.imshow("grabcut demo", panel)
    cv2.imwrite("grabcut_demo.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```
