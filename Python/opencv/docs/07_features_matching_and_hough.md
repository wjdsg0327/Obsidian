# 第 7 章：特征提取、匹配、霍夫变换与模板匹配

本章目标：理解如何从图像中提取“更高级的结构信息”，并学会把这些信息用于匹配、定位和检测。

## 知识点 1：使用 ORB 检测关键点

### 说明

关键点可以理解为图像中“值得关注”的局部位置，比如角点、纹理丰富区域、边缘交叉点等。ORB 是 OpenCV 中非常常用的一种局部特征算法，速度快、无需额外安装模型，适合教学和项目入门。

### 案例

下面的代码会生成一张特征较丰富的图像，并使用 ORB 检测关键点，然后把检测结果画出来。

### 完整代码

```python
import cv2
import numpy as np


def create_feature_scene():
    image = np.full((360, 520, 3), 240, dtype=np.uint8)
    cv2.rectangle(image, (30, 40), (170, 170), (255, 120, 0), -1)
    cv2.circle(image, (290, 120), 65, (0, 210, 255), -1)
    cv2.line(image, (390, 40), (490, 170), (0, 0, 255), 5)
    cv2.line(image, (490, 40), (390, 170), (0, 0, 255), 5)
    cv2.putText(image, "ORB", (70, 280), cv2.FONT_HERSHEY_SIMPLEX, 2.0, (30, 30, 30), 4)

    for x in range(250, 500, 20):
        cv2.line(image, (x, 220), (x, 330), (80, 80, 80), 1)
    for y in range(220, 331, 20):
        cv2.line(image, (250, y), (500, y), (80, 80, 80), 1)

    return image


def main():
    image = create_feature_scene()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

    orb = cv2.ORB_create(nfeatures=300)
    keypoints, descriptors = orb.detectAndCompute(gray, None)

    if descriptors is None:
        raise RuntimeError("未检测到可用特征，请检查输入图像")

    output = cv2.drawKeypoints(
        image,
        keypoints,
        None,
        color=(0, 0, 255),
        flags=cv2.DRAW_MATCHES_FLAGS_DRAW_RICH_KEYPOINTS,
    )

    print("keypoints:", len(keypoints))
    cv2.imshow("orb keypoints", output)
    cv2.imwrite("orb_keypoints.png", output)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 2：使用 ORB + BFMatcher 做特征匹配

### 说明

特征匹配的基本流程是：

1. 在两张图上提取关键点和描述子
2. 使用匹配器比较描述子相似度
3. 按距离排序，保留较好的匹配

ORB 通常配合 `BFMatcher` 和 `NORM_HAMMING` 使用，因为它的描述子是二进制形式。

### 案例

下面的代码会生成一张原图，再对它做旋转和平移，最后使用 ORB 匹配两张图的相似特征点。

### 完整代码

```python
import cv2
import numpy as np


def create_feature_scene():
    image = np.full((360, 520, 3), 245, dtype=np.uint8)
    cv2.rectangle(image, (30, 50), (180, 180), (255, 120, 0), -1)
    cv2.circle(image, (310, 120), 70, (0, 200, 255), -1)
    cv2.putText(image, "MATCH", (100, 300), cv2.FONT_HERSHEY_SIMPLEX, 1.8, (30, 30, 30), 4)
    for i in range(8):
        cv2.line(image, (260 + i * 25, 210), (240 + i * 25, 340), (0, 0, 180), 2)
    return image


def main():
    image1 = create_feature_scene()
    h, w = image1.shape[:2]

    rotate_matrix = cv2.getRotationMatrix2D((w // 2, h // 2), 18, 0.95)
    rotate_matrix[:, 2] += np.array([15, 10])
    image2 = cv2.warpAffine(image1, rotate_matrix, (w, h))

    orb = cv2.ORB_create(nfeatures=400)
    kp1, des1 = orb.detectAndCompute(cv2.cvtColor(image1, cv2.COLOR_BGR2GRAY), None)
    kp2, des2 = orb.detectAndCompute(cv2.cvtColor(image2, cv2.COLOR_BGR2GRAY), None)

    if des1 is None or des2 is None:
        raise RuntimeError("特征提取失败，无法进行匹配")

    matcher = cv2.BFMatcher(cv2.NORM_HAMMING, crossCheck=True)
    matches = matcher.match(des1, des2)
    matches = sorted(matches, key=lambda m: m.distance)

    matched = cv2.drawMatches(
        image1,
        kp1,
        image2,
        kp2,
        matches[:40],
        None,
        flags=cv2.DrawMatchesFlags_NOT_DRAW_SINGLE_POINTS,
    )

    print("total matches:", len(matches))
    cv2.imshow("orb matching", matched)
    cv2.imwrite("orb_matching.png", matched)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 3：霍夫直线检测

### 说明

霍夫变换适合从边缘图中找出规则几何结构，比如直线、圆。它的一个经典应用是车道线检测、表格线检测、票据边框检测等。

在 OpenCV 里，线检测常用：

- `cv2.HoughLines()`：标准霍夫变换
- `cv2.HoughLinesP()`：概率霍夫变换，直接返回线段

### 案例

下面的代码会生成多条不同角度的白线，先用 Canny 提边，再用 `HoughLinesP` 找出这些线段。

### 完整代码

```python
import cv2
import numpy as np


def gray_to_bgr(image):
    return cv2.cvtColor(image, cv2.COLOR_GRAY2BGR)


def create_line_scene():
    image = np.zeros((320, 480, 3), dtype=np.uint8)
    cv2.line(image, (40, 280), (220, 40), (255, 255, 255), 3)
    cv2.line(image, (120, 300), (400, 80), (255, 255, 255), 3)
    cv2.line(image, (40, 60), (430, 60), (255, 255, 255), 3)
    cv2.line(image, (320, 40), (320, 290), (255, 255, 255), 3)
    return image


def main():
    image = create_line_scene()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    edges = cv2.Canny(gray, 50, 150)

    result = image.copy()
    lines = cv2.HoughLinesP(
        edges,
        rho=1,
        theta=np.pi / 180,
        threshold=60,
        minLineLength=60,
        maxLineGap=10,
    )

    if lines is not None:
        for line in lines:
            x1, y1, x2, y2 = line[0]
            cv2.line(result, (x1, y1), (x2, y2), (0, 0, 255), 2)

    panel = cv2.hconcat([image, gray_to_bgr(edges), result])
    cv2.imshow("hough lines", panel)
    cv2.imwrite("hough_lines.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 4：模板匹配

### 说明

模板匹配适合做“已知小图案在大图中的定位”。它的前提是目标外观变化不大，比如尺寸接近、旋转不明显、角度变化有限。OpenCV 通过 `cv2.matchTemplate()` 完成这类搜索。

### 案例

下面的代码会先生成一个小模板，再把它嵌入更大的场景图，最后使用模板匹配自动定位它的位置。

### 完整代码

```python
import cv2
import numpy as np


def create_template():
    icon = np.full((70, 70, 3), 255, dtype=np.uint8)
    cv2.rectangle(icon, (5, 5), (64, 64), (0, 0, 0), 2)
    cv2.circle(icon, (35, 35), 18, (0, 120, 255), -1)
    cv2.line(icon, (20, 20), (50, 50), (255, 255, 255), 3)
    cv2.line(icon, (50, 20), (20, 50), (255, 255, 255), 3)
    return icon


def create_scene(template):
    scene = np.full((260, 420, 3), 235, dtype=np.uint8)
    cv2.rectangle(scene, (20, 20), (160, 100), (255, 180, 0), -1)
    cv2.circle(scene, (330, 70), 40, (0, 200, 120), -1)
    cv2.putText(scene, "Template Matching", (40, 220), cv2.FONT_HERSHEY_SIMPLEX, 0.9, (30, 30, 30), 2)

    y, x = 110, 260
    scene[y:y + template.shape[0], x:x + template.shape[1]] = template
    return scene


def main():
    template = create_template()
    scene = create_scene(template)

    result = cv2.matchTemplate(scene, template, cv2.TM_CCOEFF_NORMED)
    _, max_val, _, max_loc = cv2.minMaxLoc(result)

    h, w = template.shape[:2]
    top_left = max_loc
    bottom_right = (top_left[0] + w, top_left[1] + h)

    output = scene.copy()
    cv2.rectangle(output, top_left, bottom_right, (0, 0, 255), 3)
    cv2.putText(output, f"score={max_val:.2f}", (top_left[0] - 10, top_left[1] - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.7, (0, 0, 255), 2)

    display_size = (template.shape[1] * 3, template.shape[0] * 3)
    template_large = cv2.resize(template, display_size, interpolation=cv2.INTER_NEAREST)
    scene_large = cv2.resize(scene, display_size)
    output_large = cv2.resize(output, display_size)
    panel = cv2.hconcat([template_large, scene_large, output_large])

    cv2.imshow("template matching", panel)
    cv2.imwrite("template_matching.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```
