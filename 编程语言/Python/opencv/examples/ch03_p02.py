"""
缩放、旋转与几何变换
知识点 2：旋转、翻转与平移

来源：docs/03_geometric_transforms.md
"""
import cv2
import numpy as np


def create_demo_image():
    image = np.full((320, 320, 3), 245, dtype=np.uint8)
    cv2.rectangle(image, (40, 40), (280, 280), (0, 180, 255), 3)
    cv2.line(image, (0, 160), (319, 160), (120, 120, 120), 1)
    cv2.line(image, (160, 0), (160, 319), (120, 120, 120), 1)
    cv2.putText(image, "A", (130, 180), cv2.FONT_HERSHEY_SIMPLEX, 2.5, (255, 0, 0), 5)
    return image


def add_title(image, title):
    output = image.copy()
    cv2.putText(output, title, (12, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return output


def main():
    image = create_demo_image()
    h, w = image.shape[:2]

    rotate_matrix = cv2.getRotationMatrix2D((w // 2, h // 2), 30, 1.0)
    rotated = cv2.warpAffine(image, rotate_matrix, (w, h))

    flipped = cv2.flip(image, 1)

    move_matrix = np.float32([[1, 0, 35], [0, 1, 45]])
    translated = cv2.warpAffine(image, move_matrix, (w, h))

    panel = cv2.hconcat([
        add_title(image, "ORIGINAL"),
        add_title(rotated, "ROTATED"),
        add_title(flipped, "FLIPPED"),
        add_title(translated, "TRANSLATED"),
    ])

    cv2.imshow("basic transforms", panel)
    cv2.imwrite("basic_transforms.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
