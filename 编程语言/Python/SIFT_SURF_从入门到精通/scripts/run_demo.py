from __future__ import annotations

import argparse
from pathlib import Path

import cv2 as cv

from common import (
    DEFAULT_IMAGE_A,
    DEFAULT_IMAGE_B,
    DetectorUnavailable,
    detect_and_match,
    draw_keypoints,
    draw_matches,
    metrics_dict,
    read_image,
    save_json,
    write_image,
)


def main() -> int:
    parser = argparse.ArgumentParser(description="局部特征检测、匹配和几何验证演示")
    parser.add_argument("--algorithm", default="sift", choices=["sift", "surf", "orb", "akaze"])
    parser.add_argument("--image-a", type=Path, default=DEFAULT_IMAGE_A)
    parser.add_argument("--image-b", type=Path, default=DEFAULT_IMAGE_B)
    parser.add_argument("--ratio", type=float, default=0.75)
    parser.add_argument("--ransac-threshold", type=float, default=5.0)
    parser.add_argument("--output-dir", type=Path, default=Path(__file__).resolve().parents[1] / "outputs")
    args = parser.parse_args()

    image_a = read_image(args.image_a)
    image_b = read_image(args.image_b)
    gray_a = cv.cvtColor(image_a, cv.COLOR_BGR2GRAY)
    gray_b = cv.cvtColor(image_b, cv.COLOR_BGR2GRAY)

    try:
        result = detect_and_match(gray_a, gray_b, args.algorithm, args.ratio, args.ransac_threshold)
    except DetectorUnavailable as exc:
        print(f"[不可用] {exc}")
        return 0

    prefix = args.algorithm.lower()
    args.output_dir.mkdir(parents=True, exist_ok=True)
    write_image(args.output_dir / f"{prefix}_keypoints_a.png", draw_keypoints(image_a, result.keypoints_a))
    write_image(args.output_dir / f"{prefix}_keypoints_b.png", draw_keypoints(image_b, result.keypoints_b))
    write_image(args.output_dir / f"{prefix}_matches.png", draw_matches(image_a, image_b, result))
    if result.inlier_mask is not None:
        write_image(args.output_dir / f"{prefix}_inlier_matches.png", draw_matches(image_a, image_b, result, inliers_only=True))
    metrics = metrics_dict(result, args.algorithm, args.ratio)
    metrics["image_a"] = str(args.image_a)
    metrics["image_b"] = str(args.image_b)
    save_json(args.output_dir / f"{prefix}_metrics.json", metrics)
    print(metrics)
    print(f"输出目录：{args.output_dir.resolve()}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

