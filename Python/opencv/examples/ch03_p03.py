"""
缩放、旋转与几何变换
知识点 3：仿射变换与透视变换

来源：docs/03_geometric_transforms.md
"""
import cv2
import numpy as np


def create_source_image():
    image = np.full((320, 420, 3), 255, dtype=np.uint8)
    cv2.rectangle(image, (60, 60), (360, 250), (0, 200, 255), -1)
    cv2.putText(image, "DOC", (150, 170), cv2.FONT_HERSHEY_SIMPLEX, 2.0, (40, 40, 40), 4)
    return image


def add_title(image, title):
    result = image.copy()
    cv2.putText(result, title, (10, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return result


def main():
    image = create_source_image()

    src_affine = np.float32([[60, 60], [360, 60], [60, 250]])
    dst_affine = np.float32([[40, 100], [380, 40], [90, 280]])
    affine_matrix = cv2.getAffineTransform(src_affine, dst_affine)
    affine_result = cv2.warpAffine(image, affine_matrix, (420, 320))

    src_perspective = np.float32([[60, 60], [360, 60], [360, 250], [60, 250]])
    dst_perspective = np.float32([[100, 40], [330, 80], [370, 260], [80, 240]])
    perspective_matrix = cv2.getPerspectiveTransform(src_perspective, dst_perspective)
    warped = cv2.warpPerspective(image, perspective_matrix, (420, 320))

    rectified_matrix = cv2.getPerspectiveTransform(dst_perspective, src_perspective)
    rectified = cv2.warpPerspective(warped, rectified_matrix, (420, 320))

    panel = cv2.hconcat([
        add_title(image, "SOURCE"),
        add_title(affine_result, "AFFINE"),
        add_title(warped, "PERSPECTIVE"),
        add_title(rectified, "RECTIFIED"),
    ])

    cv2.imshow("affine and perspective", panel)
    cv2.imwrite("affine_perspective_panel.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
