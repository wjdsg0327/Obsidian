"""
特征提取、匹配、霍夫变换与模板匹配
知识点 2：使用 ORB + BFMatcher 做特征匹配

来源：docs/07_features_matching_and_hough.md
"""
import cv2
import numpy as np


def create_feature_scene():
    image = np.full((360, 520, 3), 245, dtype=np.uint8)
    cv2.rectangle(image, (30, 50), (180, 180), (255, 120, 0), -1)
    cv2.circle(image, (310, 120), 70, (0, 200, 255), -1)
    cv2.putText(image, "MATCH", (100, 300), cv2.FONT_HERSHEY_SIMPLEX, 1.8, (30, 30, 30), 4)
    for i in range(8):
        cv2.line(image, (260 + i * 25, 210), (240 + i * 25, 340), (0, 0, 180), 2)
    return image


def main():
    image1 = create_feature_scene()
    h, w = image1.shape[:2]

    rotate_matrix = cv2.getRotationMatrix2D((w // 2, h // 2), 18, 0.95)
    rotate_matrix[:, 2] += np.array([15, 10])
    image2 = cv2.warpAffine(image1, rotate_matrix, (w, h))

    orb = cv2.ORB_create(nfeatures=400)
    kp1, des1 = orb.detectAndCompute(cv2.cvtColor(image1, cv2.COLOR_BGR2GRAY), None)
    kp2, des2 = orb.detectAndCompute(cv2.cvtColor(image2, cv2.COLOR_BGR2GRAY), None)

    if des1 is None or des2 is None:
        raise RuntimeError("特征提取失败，无法进行匹配")

    matcher = cv2.BFMatcher(cv2.NORM_HAMMING, crossCheck=True)
    matches = matcher.match(des1, des2)
    matches = sorted(matches, key=lambda m: m.distance)

    matched = cv2.drawMatches(
        image1,
        kp1,
        image2,
        kp2,
        matches[:40],
        None,
        flags=cv2.DrawMatchesFlags_NOT_DRAW_SINGLE_POINTS,
    )

    print("total matches:", len(matches))
    cv2.imshow("orb matching", matched)
    cv2.imwrite("orb_matching.png", matched)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
