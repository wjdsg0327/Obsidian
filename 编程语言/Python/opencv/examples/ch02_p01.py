"""
绘图、文字与颜色空间
知识点 1：绘制线条、矩形、圆和多边形

来源：docs/02_drawing_and_color_spaces.md
"""
import cv2
import numpy as np


def main():
    canvas = np.full((420, 520, 3), 240, dtype=np.uint8)

    cv2.line(canvas, (30, 40), (220, 180), (255, 0, 0), 3)
    cv2.rectangle(canvas, (260, 40), (470, 170), (0, 255, 0), 3)
    cv2.circle(canvas, (120, 300), 70, (0, 0, 255), -1)
    cv2.ellipse(canvas, (340, 300), (100, 50), 30, 0, 300, (255, 180, 0), 4)

    points = np.array([[260, 220], [440, 250], [390, 380], [240, 340]], dtype=np.int32)
    cv2.polylines(canvas, [points], True, (120, 0, 200), 4)

    cv2.imshow("drawing primitives", canvas)
    cv2.imwrite("drawing_primitives.png", canvas)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
