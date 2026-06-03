"""
绘图、文字与颜色空间
知识点 3：BGR、灰度、HSV 与通道分离

来源：docs/02_drawing_and_color_spaces.md
"""
import cv2
import numpy as np


def gray_to_bgr(channel):
    return cv2.cvtColor(channel, cv2.COLOR_GRAY2BGR)


def main():
    height, width = 240, 360
    y, x = np.indices((height, width))

    image = np.zeros((height, width, 3), dtype=np.uint8)
    image[..., 0] = x * 255 // width
    image[..., 1] = y * 255 // height
    image[..., 2] = (x + y) * 255 // (height + width)

    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    hsv = cv2.cvtColor(image, cv2.COLOR_BGR2HSV)

    b, g, r = cv2.split(image)
    h, s, v = cv2.split(hsv)

    row1 = cv2.hconcat([image, gray_to_bgr(gray), gray_to_bgr(b), gray_to_bgr(g)])
    row2 = cv2.hconcat([gray_to_bgr(r), gray_to_bgr(h), gray_to_bgr(s), gray_to_bgr(v)])
    panel = cv2.vconcat([row1, row2])

    cv2.imshow("color spaces and channels", panel)
    cv2.imwrite("color_spaces_panel.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
