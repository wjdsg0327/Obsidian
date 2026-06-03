"""
图像分割与高级处理
知识点 3：GrabCut 前景提取

来源：docs/08_segmentation_and_advanced_processing.md
"""
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
