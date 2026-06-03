"""
边缘、形态学与轮廓分析
知识点 3：查找轮廓并计算面积、周长、外接矩形

来源：docs/05_edges_morphology_and_contours.md
"""
import cv2
import numpy as np


def create_shapes():
    binary = np.zeros((320, 460), dtype=np.uint8)
    cv2.rectangle(binary, (30, 50), (140, 210), 255, -1)
    cv2.circle(binary, (250, 130), 60, 255, -1)
    pts = np.array([[330, 240], [420, 180], [430, 290], [350, 300]], dtype=np.int32)
    cv2.fillPoly(binary, [pts], 255)
    return binary


def main():
    binary = create_shapes()
    contours, _ = cv2.findContours(binary, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)

    canvas = cv2.cvtColor(binary, cv2.COLOR_GRAY2BGR)

    for index, contour in enumerate(contours, start=1):
        area = cv2.contourArea(contour)
        perimeter = cv2.arcLength(contour, True)
        x, y, w, h = cv2.boundingRect(contour)

        cv2.drawContours(canvas, [contour], -1, (0, 255, 0), 2)
        cv2.rectangle(canvas, (x, y), (x + w, y + h), (0, 0, 255), 2)
        cv2.putText(
            canvas,
            f"#{index} A={int(area)} P={int(perimeter)}",
            (x, max(20, y - 10)),
            cv2.FONT_HERSHEY_SIMPLEX,
            0.55,
            (255, 0, 0),
            2,
        )

    cv2.imshow("contour analysis", canvas)
    cv2.imwrite("contour_analysis.png", canvas)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
