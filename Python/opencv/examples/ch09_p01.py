"""
综合实战与性能优化
知识点 1：实战项目一：文档扫描与透视矫正

来源：docs/09_projects_and_optimization.md
"""
import cv2
import numpy as np


def order_points(points):
    points = points.astype(np.float32)
    rect = np.zeros((4, 2), dtype=np.float32)
    s = points.sum(axis=1)
    diff = np.diff(points, axis=1)

    rect[0] = points[np.argmin(s)]
    rect[2] = points[np.argmax(s)]
    rect[1] = points[np.argmin(diff)]
    rect[3] = points[np.argmax(diff)]
    return rect


def four_point_transform(image, points):
    rect = order_points(points)
    tl, tr, br, bl = rect

    width_top = np.linalg.norm(tr - tl)
    width_bottom = np.linalg.norm(br - bl)
    max_width = int(max(width_top, width_bottom))

    height_left = np.linalg.norm(bl - tl)
    height_right = np.linalg.norm(br - tr)
    max_height = int(max(height_left, height_right))

    dst = np.array(
        [[0, 0], [max_width - 1, 0], [max_width - 1, max_height - 1], [0, max_height - 1]],
        dtype=np.float32,
    )

    matrix = cv2.getPerspectiveTransform(rect, dst)
    warped = cv2.warpPerspective(image, matrix, (max_width, max_height))
    return warped


def create_document_scene():
    document = np.full((300, 420, 3), 250, dtype=np.uint8)
    cv2.putText(document, "Invoice", (120, 60), cv2.FONT_HERSHEY_SIMPLEX, 1.4, (30, 30, 30), 3)
    for y in range(100, 250, 28):
        cv2.line(document, (35, y), (380, y), (80, 80, 80), 2)
    cv2.rectangle(document, (20, 20), (399, 279), (40, 40, 40), 3)

    scene = np.full((420, 620, 3), 70, dtype=np.uint8)
    src = np.float32([[0, 0], [419, 0], [419, 299], [0, 299]])
    dst = np.float32([[120, 60], [520, 90], [470, 350], [80, 320]])

    matrix = cv2.getPerspectiveTransform(src, dst)
    warped_doc = cv2.warpPerspective(document, matrix, (620, 420))
    warped_mask = cv2.warpPerspective(np.full((300, 420), 255, dtype=np.uint8), matrix, (620, 420))

    scene[warped_mask > 0] = warped_doc[warped_mask > 0]
    scene = cv2.GaussianBlur(scene, (3, 3), 0)
    return scene


def main():
    image = create_document_scene()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    blur = cv2.GaussianBlur(gray, (5, 5), 0)
    edges = cv2.Canny(blur, 50, 150)

    contours, _ = cv2.findContours(edges, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
    contours = sorted(contours, key=cv2.contourArea, reverse=True)

    screen_contour = None
    for contour in contours:
        peri = cv2.arcLength(contour, True)
        approx = cv2.approxPolyDP(contour, 0.02 * peri, True)
        if len(approx) == 4:
            screen_contour = approx.reshape(4, 2)
            break

    if screen_contour is None:
        raise RuntimeError("未找到四边形文档轮廓")

    warped = four_point_transform(image, screen_contour)

    outlined = image.copy()
    cv2.polylines(outlined, [screen_contour.astype(np.int32)], True, (0, 0, 255), 3)

    warped_small = cv2.resize(warped, (outlined.shape[1], outlined.shape[0]))
    panel = cv2.hconcat([outlined, warped_small])

    cv2.imshow("document scanner", panel)
    cv2.imwrite("document_scanner_result.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
