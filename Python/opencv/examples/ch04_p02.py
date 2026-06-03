"""
滤波、阈值与图像增强
知识点 2：全局阈值、自适应阈值与 Otsu

来源：docs/04_filtering_threshold_and_enhancement.md
"""
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
