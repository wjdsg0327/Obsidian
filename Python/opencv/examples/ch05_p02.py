"""
边缘、形态学与轮廓分析
知识点 2：腐蚀、膨胀、开运算与闭运算

来源：docs/05_edges_morphology_and_contours.md
"""
import cv2
import numpy as np


def gray_to_bgr(image):
    return cv2.cvtColor(image, cv2.COLOR_GRAY2BGR)


def create_binary_image():
    image = np.zeros((260, 420), dtype=np.uint8)
    cv2.rectangle(image, (40, 60), (170, 200), 255, -1)
    cv2.rectangle(image, (210, 70), (360, 180), 255, -1)
    cv2.circle(image, (285, 125), 22, 0, -1)
    cv2.line(image, (100, 60), (100, 200), 0, 8)

    rng = np.random.default_rng(7)
    for _ in range(500):
        y = rng.integers(0, image.shape[0])
        x = rng.integers(0, image.shape[1])
        image[y, x] = 255

    return image


def add_title(image, title):
    output = image.copy()
    cv2.putText(output, title, (10, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return output


def main():
    binary = create_binary_image()
    kernel = np.ones((5, 5), dtype=np.uint8)

    eroded = cv2.erode(binary, kernel, iterations=1)
    dilated = cv2.dilate(binary, kernel, iterations=1)
    opened = cv2.morphologyEx(binary, cv2.MORPH_OPEN, kernel)
    closed = cv2.morphologyEx(binary, cv2.MORPH_CLOSE, kernel)

    row1 = cv2.hconcat([
        add_title(gray_to_bgr(binary), "SOURCE"),
        add_title(gray_to_bgr(eroded), "ERODE"),
        add_title(gray_to_bgr(dilated), "DILATE"),
    ])
    row2 = cv2.hconcat([
        add_title(gray_to_bgr(opened), "OPEN"),
        add_title(gray_to_bgr(closed), "CLOSE"),
        np.full_like(row1[:, :binary.shape[1]], 255),
    ])

    panel = cv2.vconcat([row1, row2])

    cv2.imshow("morphology demo", panel)
    cv2.imwrite("morphology_demo.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
