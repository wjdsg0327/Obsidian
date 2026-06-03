"""
环境准备与第一段 OpenCV 代码
知识点 1：安装并验证 OpenCV 环境

来源：docs/00_setup_and_first_program.md
"""
import cv2
import numpy as np


def main():
    print("OpenCV version:", cv2.__version__)
    print("NumPy version:", np.__version__)

    canvas = np.zeros((220, 420, 3), dtype=np.uint8)
    canvas[:] = (35, 35, 35)

    cv2.putText(
        canvas,
        "Hello OpenCV",
        (35, 120),
        cv2.FONT_HERSHEY_SIMPLEX,
        1.2,
        (0, 255, 0),
        2,
        cv2.LINE_AA,
    )

    cv2.imwrite("hello_opencv.png", canvas)
    print("image shape:", canvas.shape)
    print("image dtype:", canvas.dtype)
    print("saved file: hello_opencv.png")


if __name__ == "__main__":
    main()
