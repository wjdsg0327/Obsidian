"""
图像读写、像素访问与 ROI
知识点 2：像素访问与局部修改

来源：docs/01_image_io_and_pixels.md
"""
import cv2
import numpy as np


def main():
    image = np.zeros((260, 360, 3), dtype=np.uint8)
    image[:] = (20, 20, 20)

    image[30:110, 40:120] = (255, 0, 0)
    image[30:110, 140:220] = (0, 255, 0)
    image[30:110, 240:320] = (0, 0, 255)

    for i in range(0, 360, 20):
        image[160:220, i:i + 10] = (0, 255, 255)

    print("pixel at (50, 60):", image[50, 60])
    print("pixel at (180, 200):", image[180, 200])

    image[120:150, 100:260] = (255, 255, 255)

    cv2.imshow("pixel editing", image)
    cv2.imwrite("pixel_editing.png", image)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
