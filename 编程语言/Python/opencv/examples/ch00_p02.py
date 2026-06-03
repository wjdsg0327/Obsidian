"""
环境准备与第一段 OpenCV 代码
知识点 2：生成你的第一张测试图像

来源：docs/00_setup_and_first_program.md
"""
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
