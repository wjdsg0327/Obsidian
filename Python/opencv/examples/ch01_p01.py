"""
图像读写、像素访问与 ROI
知识点 1：读取、显示、保存图像

来源：docs/01_image_io_and_pixels.md
"""
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
