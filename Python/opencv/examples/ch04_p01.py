"""
滤波、阈值与图像增强
知识点 1：均值滤波、高斯滤波与中值滤波

来源：docs/04_filtering_threshold_and_enhancement.md
"""
import cv2
import numpy as np


def create_clean_image():
    image = np.full((260, 420, 3), 235, dtype=np.uint8)
    cv2.rectangle(image, (40, 40), (180, 200), (0, 180, 255), -1)
    cv2.circle(image, (310, 130), 65, (255, 120, 0), -1)
    cv2.putText(image, "NOISE", (110, 245), cv2.FONT_HERSHEY_SIMPLEX, 1.2, (40, 40, 40), 2)
    return image


def add_noise(image):
    rng = np.random.default_rng(42)
    noisy = image.astype(np.int16)
    gaussian_noise = rng.normal(0, 22, image.shape).astype(np.int16)
    noisy = np.clip(noisy + gaussian_noise, 0, 255).astype(np.uint8)

    for _ in range(1800):
        y = rng.integers(0, image.shape[0])
        x = rng.integers(0, image.shape[1])
        noisy[y, x] = (255, 255, 255) if rng.random() > 0.5 else (0, 0, 0)

    return noisy


def label(image, title):
    output = image.copy()
    cv2.putText(output, title, (10, 28), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2)
    return output


def main():
    clean = create_clean_image()
    noisy = add_noise(clean)

    mean_blur = cv2.blur(noisy, (5, 5))
    gaussian_blur = cv2.GaussianBlur(noisy, (5, 5), 0)
    median_blur = cv2.medianBlur(noisy, 5)

    blank = np.full_like(clean, 255)

    row1 = cv2.hconcat([
        label(clean, "CLEAN"),
        label(noisy, "NOISY"),
        label(blank, "BLANK"),
    ])
    row2 = cv2.hconcat([
        label(mean_blur, "MEAN"),
        label(gaussian_blur, "GAUSSIAN"),
        label(median_blur, "MEDIAN"),
    ])
    panel = cv2.vconcat([row1, row2])

    cv2.imshow("filter comparison", panel)
    cv2.imwrite("filter_comparison.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
