"""
综合实战与性能优化
知识点 2：实战项目二：圆形目标计数

来源：docs/09_projects_and_optimization.md
"""
import cv2
import numpy as np


def create_coin_scene():
    image = np.full((340, 520, 3), 35, dtype=np.uint8)
    centers = [(80, 90), (210, 80), (360, 105), (120, 230), (270, 220), (420, 240)]
    radii = [35, 40, 32, 38, 42, 36]

    for center, radius in zip(centers, radii):
        cv2.circle(image, center, radius, (220, 220, 220), -1)
        cv2.circle(image, center, radius, (180, 180, 180), 3)

    return image


def main():
    image = create_coin_scene()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    _, binary = cv2.threshold(gray, 100, 255, cv2.THRESH_BINARY)

    kernel = np.ones((3, 3), np.uint8)
    binary = cv2.morphologyEx(binary, cv2.MORPH_OPEN, kernel)

    contours, _ = cv2.findContours(binary, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)

    output = image.copy()
    count = 0
    for contour in contours:
        area = cv2.contourArea(contour)
        if area < 1000:
            continue

        count += 1
        x, y, w, h = cv2.boundingRect(contour)
        cv2.drawContours(output, [contour], -1, (0, 255, 255), 2)
        cv2.rectangle(output, (x, y), (x + w, y + h), (0, 0, 255), 2)
        cv2.putText(output, f"#{count}", (x + 10, y + 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (255, 255, 255), 2)

    cv2.putText(output, f"count = {count}", (20, 35), cv2.FONT_HERSHEY_SIMPLEX, 1.0, (0, 255, 0), 2)

    panel = cv2.hconcat([image, cv2.cvtColor(binary, cv2.COLOR_GRAY2BGR), output])
    cv2.imshow("object counting", panel)
    cv2.imwrite("object_counting_result.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
