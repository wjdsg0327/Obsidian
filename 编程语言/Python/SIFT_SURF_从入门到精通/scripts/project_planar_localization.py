from __future__ import annotations

import argparse
from pathlib import Path

import cv2 as cv
import numpy as np

from common import (
    DEFAULT_IMAGE_A,
    DEFAULT_IMAGE_B,
    detect_and_match,
    draw_matches,
    metrics_dict,
    read_image,
    save_json,
    write_image,
)


def main() -> int:
    parser = argparse.ArgumentParser(description="SIFT + RANSAC 平面目标定位")
    parser.add_argument("--object", type=Path, default=DEFAULT_IMAGE_A)
    parser.add_argument("--scene", type=Path, default=DEFAULT_IMAGE_B)
    parser.add_argument("--ratio", type=float, default=0.75)
    parser.add_argument("--output-dir", type=Path, default=Path(__file__).resolve().parents[1] / "outputs")
    args = parser.parse_args()

    object_image = read_image(args.object)
    scene_image = read_image(args.scene)
    result = detect_and_match(
        cv.cvtColor(object_image, cv.COLOR_BGR2GRAY),
        cv.cvtColor(scene_image, cv.COLOR_BGR2GRAY),
        "sift",
        args.ratio,
        5.0,
    )
    if result.homography is None or result.inliers < 4:
        raise RuntimeError(f"无法得到可靠单应性：good={len(result.good_matches)}, inliers={result.inliers}")

    h, w = object_image.shape[:2]
    corners = np.float32([[0, 0], [w - 1, 0], [w - 1, h - 1], [0, h - 1]]).reshape(-1, 1, 2)
    projected = cv.perspectiveTransform(corners, result.homography)
    localized = scene_image.copy()
    cv.polylines(localized, [np.int32(projected)], True, (0, 255, 0), 3, cv.LINE_AA)

    args.output_dir.mkdir(parents=True, exist_ok=True)
    write_image(args.output_dir / "planar_localization.png", localized)
    write_image(args.output_dir / "planar_inlier_matches.png", draw_matches(object_image, scene_image, result, inliers_only=True))
    metrics = metrics_dict(result, "sift", args.ratio)
    metrics["projected_corners"] = projected.reshape(-1, 2).round(3).tolist()
    save_json(args.output_dir / "planar_localization_metrics.json", metrics)
    print(metrics)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

