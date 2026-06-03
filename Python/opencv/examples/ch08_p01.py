"""
图像分割与高级处理
知识点 1：连通域分析

来源：docs/08_segmentation_and_advanced_processing.md
"""
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
