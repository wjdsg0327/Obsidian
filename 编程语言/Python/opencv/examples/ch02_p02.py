"""
绘图、文字与颜色空间
知识点 2：添加文字、水印和透明叠加

来源：docs/02_drawing_and_color_spaces.md
"""
import cv2
import numpy as np


def main():
    image = np.zeros((320, 520, 3), dtype=np.uint8)

    for y in range(image.shape[0]):
        color = 50 + y * 180 // image.shape[0]
        image[y, :] = (color // 2, color, 200)

    overlay = image.copy()
    cv2.rectangle(overlay, (30, 30), (490, 130), (20, 20, 20), -1)
    result = cv2.addWeighted(overlay, 0.6, image, 0.4, 0)

    cv2.putText(result, "OpenCV Overlay Demo", (50, 75), cv2.FONT_HERSHEY_SIMPLEX, 1.0, (0, 255, 255), 2)
    cv2.putText(result, "Transparent panel + text", (50, 110), cv2.FONT_HERSHEY_SIMPLEX, 0.7, (255, 255, 255), 2)

    cv2.imshow("text and overlay", result)
    cv2.imwrite("text_overlay_demo.png", result)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
