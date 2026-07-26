from __future__ import annotations

import argparse
from pathlib import Path

import cv2 as cv
import numpy as np

from common import PACKAGE_ROOT, detect_and_match, read_image, write_image


def simple_panorama(image_a: np.ndarray, image_b: np.ndarray, homography: np.ndarray) -> np.ndarray:
    h_a, w_a = image_a.shape[:2]
    h_b, w_b = image_b.shape[:2]
    corners_a = np.float32([[0, 0], [w_a, 0], [w_a, h_a], [0, h_a]]).reshape(-1, 1, 2)
    corners_b = np.float32([[0, 0], [w_b, 0], [w_b, h_b], [0, h_b]]).reshape(-1, 1, 2)
    warped_a = cv.perspectiveTransform(corners_a, homography)
    all_corners = np.concatenate([warped_a, corners_b], axis=0).reshape(-1, 2)
    min_xy = np.floor(all_corners.min(axis=0)).astype(int)
    max_xy = np.ceil(all_corners.max(axis=0)).astype(int)
    translate = np.array([[1, 0, -min_xy[0]], [0, 1, -min_xy[1]], [0, 0, 1]], dtype=np.float64)
    width, height = (max_xy - min_xy).tolist()
    if width <= 0 or height <= 0 or width * height > 80_000_000:
        raise RuntimeError(f"异常拼接画布：{width}x{height}；请检查匹配与单应性。")

    canvas = cv.warpPerspective(image_a, translate @ homography, (width, height))
    x0, y0 = -min_xy
    roi = canvas[y0 : y0 + h_b, x0 : x0 + w_b]
    if roi.shape[:2] != image_b.shape[:2]:
        raise RuntimeError("拼接画布边界计算失败。")
    mask_a = np.any(roi != 0, axis=2)
    blended = image_b.copy()
    overlap = mask_a
    blended[overlap] = ((roi[overlap].astype(np.uint16) + image_b[overlap].astype(np.uint16)) // 2).astype(np.uint8)
    blended[~overlap] = image_b[~overlap]
    canvas[y0 : y0 + h_b, x0 : x0 + w_b] = np.where(mask_a[..., None], blended, image_b)
    return canvas


def main() -> int:
    samples = PACKAGE_ROOT / "09_样例与数据集" / "opencv_samples"
    parser = argparse.ArgumentParser(description="用 SIFT 和单应性完成简单图像拼接")
    parser.add_argument("--image-a", type=Path, default=samples / "synthetic_panorama_a.png")
    parser.add_argument("--image-b", type=Path, default=samples / "synthetic_panorama_b.png")
    parser.add_argument("--output", type=Path, default=PACKAGE_ROOT / "outputs" / "simple_panorama.png")
    args = parser.parse_args()

    image_a = read_image(args.image_a)
    image_b = read_image(args.image_b)
    result = detect_and_match(
        cv.cvtColor(image_a, cv.COLOR_BGR2GRAY),
        cv.cvtColor(image_b, cv.COLOR_BGR2GRAY),
        "sift",
        0.75,
        5.0,
    )
    if result.homography is None or result.inliers < 8:
        raise RuntimeError(f"拼接几何不足：good={len(result.good_matches)}, inliers={result.inliers}")
    panorama = simple_panorama(image_a, image_b, result.homography)
    write_image(args.output, panorama)
    print(f"拼接完成：{args.output.resolve()} | good={len(result.good_matches)} inliers={result.inliers}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
