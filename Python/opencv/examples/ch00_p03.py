"""
环境准备与第一段 OpenCV 代码
知识点 3：理解图像的形状、类型和像素值

来源：docs/00_setup_and_first_program.md
"""
import cv2
import numpy as np


def main():
    height, width = 240, 320
    y, x = np.indices((height, width))

    image = np.zeros((height, width, 3), dtype=np.uint8)
    image[..., 0] = x * 255 // width
    image[..., 1] = y * 255 // height
    image[..., 2] = (x + y) * 255 // (height + width)

    print("shape:", image.shape)
    print("dtype:", image.dtype)
    print("ndim:", image.ndim)
    print("size:", image.size)
    print("itemsize:", image.itemsize)

    center_pixel = image[height // 2, width // 2]
    print("center pixel (BGR):", center_pixel)

    image[80:160, 100:220] = (0, 255, 255)

    cv2.imshow("image info demo", image)
    cv2.imwrite("image_info_demo.png", image)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
