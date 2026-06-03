"""
边缘、形态学与轮廓分析
知识点 1：Canny 边缘检测

来源：docs/05_edges_morphology_and_contours.md
"""
import cv2
import numpy as np


def gray_to_bgr(image):
    return cv2.cvtColor(image, cv2.COLOR_GRAY2BGR)


def create_source():
    image = np.full((280, 420, 3), 230, dtype=np.uint8)
    cv2.rectangle(image, (40, 50), (170, 220), (255, 120, 0), -1)
    cv2.circle(image, (290, 130), 65, (0, 180, 255), -1)
    cv2.putText(image, "EDGE", (120, 260), cv2.FONT_HERSHEY_SIMPLEX, 1.4, (30, 30, 30), 3)
    return image


def main():
    image = create_source()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    blur = cv2.GaussianBlur(gray, (5, 5), 0)
    edges = cv2.Canny(blur, 60, 150)

    panel = cv2.hconcat([image, gray_to_bgr(blur), gray_to_bgr(edges)])

    cv2.imshow("canny edge", panel)
    cv2.imwrite("canny_edge_demo.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
