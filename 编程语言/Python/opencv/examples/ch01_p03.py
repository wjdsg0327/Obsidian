"""
图像读写、像素访问与 ROI
知识点 3：ROI 裁剪、复制与拼接

来源：docs/01_image_io_and_pixels.md
"""
import cv2
import numpy as np


def create_layout():
    image = np.zeros((300, 420, 3), dtype=np.uint8)
    image[:] = (40, 40, 40)
    cv2.rectangle(image, (20, 30), (180, 140), (255, 120, 0), -1)
    cv2.rectangle(image, (220, 30), (390, 140), (0, 200, 120), -1)
    cv2.rectangle(image, (20, 170), (180, 270), (120, 80, 255), -1)
    cv2.putText(image, "ROI", (70, 100), cv2.FONT_HERSHEY_SIMPLEX, 1.2, (255, 255, 255), 2)
    return image


def main():
    image = create_layout()

    roi = image[30:140, 20:180].copy()
    image[170:280, 220:380] = roi

    left = image[:, :210]
    right = image[:, 210:]
    stitched = cv2.hconcat([left, right])

    cv2.imshow("roi copy", image)
    cv2.imshow("stitched", stitched)
    cv2.imwrite("roi_copy.png", image)
    cv2.imwrite("roi_stitched.png", stitched)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
