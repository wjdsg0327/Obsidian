"""
缩放、旋转与几何变换
知识点 1：缩放与插值方法

来源：docs/03_geometric_transforms.md
"""
import cv2
import numpy as np


def create_checkerboard(rows=8, cols=10, cell=20):
    board = np.zeros((rows * cell, cols * cell), dtype=np.uint8)
    for r in range(rows):
        for c in range(cols):
            if (r + c) % 2 == 0:
                board[r * cell:(r + 1) * cell, c * cell:(c + 1) * cell] = 255
    return board


def add_title(image, title):
    image = image.copy()
    cv2.putText(image, title, (10, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return image


def main():
    board = create_checkerboard()
    board_bgr = cv2.cvtColor(board, cv2.COLOR_GRAY2BGR)

    nearest = cv2.resize(board_bgr, None, fx=4, fy=4, interpolation=cv2.INTER_NEAREST)
    linear = cv2.resize(board_bgr, None, fx=4, fy=4, interpolation=cv2.INTER_LINEAR)
    cubic = cv2.resize(board_bgr, None, fx=4, fy=4, interpolation=cv2.INTER_CUBIC)

    panel = cv2.hconcat([
        add_title(nearest, "NEAREST"),
        add_title(linear, "LINEAR"),
        add_title(cubic, "CUBIC"),
    ])

    cv2.imshow("resize interpolation", panel)
    cv2.imwrite("resize_interpolation.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
