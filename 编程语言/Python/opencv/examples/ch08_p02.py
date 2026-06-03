"""
图像分割与高级处理
知识点 2：分水岭算法分离粘连目标

来源：docs/08_segmentation_and_advanced_processing.md
"""
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
