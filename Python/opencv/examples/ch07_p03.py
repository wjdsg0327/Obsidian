"""
特征提取、匹配、霍夫变换与模板匹配
知识点 3：霍夫直线检测

来源：docs/07_features_matching_and_hough.md
"""
import cv2
import numpy as np


def gray_to_bgr(image):
    return cv2.cvtColor(image, cv2.COLOR_GRAY2BGR)


def create_line_scene():
    image = np.zeros((320, 480, 3), dtype=np.uint8)
    cv2.line(image, (40, 280), (220, 40), (255, 255, 255), 3)
    cv2.line(image, (120, 300), (400, 80), (255, 255, 255), 3)
    cv2.line(image, (40, 60), (430, 60), (255, 255, 255), 3)
    cv2.line(image, (320, 40), (320, 290), (255, 255, 255), 3)
    return image


def main():
    image = create_line_scene()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    edges = cv2.Canny(gray, 50, 150)

    result = image.copy()
    lines = cv2.HoughLinesP(
        edges,
        rho=1,
        theta=np.pi / 180,
        threshold=60,
        minLineLength=60,
        maxLineGap=10,
    )

    if lines is not None:
        for line in lines:
            x1, y1, x2, y2 = line[0]
            cv2.line(result, (x1, y1), (x2, y2), (0, 0, 255), 2)

    panel = cv2.hconcat([image, gray_to_bgr(edges), result])
    cv2.imshow("hough lines", panel)
    cv2.imwrite("hough_lines.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
