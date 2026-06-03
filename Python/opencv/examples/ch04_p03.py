"""
滤波、阈值与图像增强
知识点 3：直方图均衡化与 CLAHE

来源：docs/04_filtering_threshold_and_enhancement.md
"""
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
