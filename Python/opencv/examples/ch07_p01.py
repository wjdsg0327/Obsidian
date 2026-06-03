"""
特征提取、匹配、霍夫变换与模板匹配
知识点 1：使用 ORB 检测关键点

来源：docs/07_features_matching_and_hough.md
"""
import cv2
import numpy as np


def create_feature_scene():
    image = np.full((360, 520, 3), 240, dtype=np.uint8)
    cv2.rectangle(image, (30, 40), (170, 170), (255, 120, 0), -1)
    cv2.circle(image, (290, 120), 65, (0, 210, 255), -1)
    cv2.line(image, (390, 40), (490, 170), (0, 0, 255), 5)
    cv2.line(image, (490, 40), (390, 170), (0, 0, 255), 5)
    cv2.putText(image, "ORB", (70, 280), cv2.FONT_HERSHEY_SIMPLEX, 2.0, (30, 30, 30), 4)

    for x in range(250, 500, 20):
        cv2.line(image, (x, 220), (x, 330), (80, 80, 80), 1)
    for y in range(220, 331, 20):
        cv2.line(image, (250, y), (500, y), (80, 80, 80), 1)

    return image


def main():
    image = create_feature_scene()
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

    orb = cv2.ORB_create(nfeatures=300)
    keypoints, descriptors = orb.detectAndCompute(gray, None)

    if descriptors is None:
        raise RuntimeError("未检测到可用特征，请检查输入图像")

    output = cv2.drawKeypoints(
        image,
        keypoints,
        None,
        color=(0, 0, 255),
        flags=cv2.DRAW_MATCHES_FLAGS_DRAW_RICH_KEYPOINTS,
    )

    print("keypoints:", len(keypoints))
    cv2.imshow("orb keypoints", output)
    cv2.imwrite("orb_keypoints.png", output)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
