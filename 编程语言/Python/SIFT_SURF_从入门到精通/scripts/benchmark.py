from __future__ import annotations

import argparse
import csv
from pathlib import Path
from statistics import median
from time import perf_counter

import cv2 as cv
import numpy as np

from common import (
    DEFAULT_IMAGE_A,
    DEFAULT_IMAGE_B,
    DetectorUnavailable,
    create_detector,
    detect_and_match,
    read_image,
    save_json,
    write_image,
)


def draw_benchmark_chart(rows: list[dict]) -> np.ndarray:
    canvas = np.full((500, 1100, 3), 248, dtype=np.uint8)
    cv.putText(canvas, "Local feature benchmark", (35, 48), cv.FONT_HERSHEY_DUPLEX, 1.2, (25, 35, 50), 2, cv.LINE_AA)
    panels = [
        (40, 90, 510, 450, "Median total time (ms)", "median_total_ms", (210, 90, 35)),
        (590, 90, 1060, 450, "RANSAC inlier ratio", "inlier_ratio", (60, 155, 60)),
    ]
    for left, top, right, bottom, title, field, color in panels:
        cv.rectangle(canvas, (left, top), (right, bottom), (215, 220, 228), 1)
        cv.putText(canvas, title, (left + 18, top + 35), cv.FONT_HERSHEY_SIMPLEX, 0.68, (35, 40, 50), 2, cv.LINE_AA)
        values = [float(row[field]) for row in rows]
        maximum = 1.0 if field == "inlier_ratio" else max(values + [1.0]) * 1.12
        chart_top, chart_bottom = top + 65, bottom - 55
        usable = chart_bottom - chart_top
        slot = (right - left - 50) / max(len(rows), 1)
        for index, (row, value) in enumerate(zip(rows, values)):
            x1 = int(left + 30 + index * slot + slot * 0.18)
            x2 = int(left + 30 + (index + 1) * slot - slot * 0.18)
            y1 = int(chart_bottom - usable * value / maximum)
            cv.rectangle(canvas, (x1, y1), (x2, chart_bottom), color, -1)
            label = row["algorithm"].upper()
            cv.putText(canvas, label, (x1, chart_bottom + 28), cv.FONT_HERSHEY_SIMPLEX, 0.55, (35, 40, 50), 1, cv.LINE_AA)
            value_text = f"{value:.3f}" if field == "inlier_ratio" else f"{value:.1f}"
            cv.putText(canvas, value_text, (x1, max(chart_top + 15, y1 - 8)), cv.FONT_HERSHEY_SIMPLEX, 0.5, (25, 35, 50), 1, cv.LINE_AA)
    return canvas


def main() -> int:
    parser = argparse.ArgumentParser(description="SIFT/SURF/ORB/AKAZE 小型基准测试")
    parser.add_argument("--algorithms", default="sift,orb,akaze")
    parser.add_argument("--image-a", type=Path, default=DEFAULT_IMAGE_A)
    parser.add_argument("--image-b", type=Path, default=DEFAULT_IMAGE_B)
    parser.add_argument("--ratio", type=float, default=0.75)
    parser.add_argument("--repeats", type=int, default=5)
    parser.add_argument("--output-dir", type=Path, default=Path(__file__).resolve().parents[1] / "outputs")
    args = parser.parse_args()
    if args.repeats < 1:
        raise ValueError("repeats 必须大于 0")

    gray_a = cv.cvtColor(read_image(args.image_a), cv.COLOR_BGR2GRAY)
    gray_b = cv.cvtColor(read_image(args.image_b), cv.COLOR_BGR2GRAY)
    rows = []

    for algorithm in [x.strip().lower() for x in args.algorithms.split(",") if x.strip()]:
        try:
            create_detector(algorithm)
            detect_and_match(gray_a, gray_b, algorithm, args.ratio)
        except DetectorUnavailable as exc:
            rows.append({"algorithm": algorithm, "status": "unavailable", "message": str(exc).splitlines()[0]})
            continue

        times = []
        last = None
        for _ in range(args.repeats):
            start = perf_counter()
            last = detect_and_match(gray_a, gray_b, algorithm, args.ratio)
            times.append((perf_counter() - start) * 1000.0)
        assert last is not None
        rows.append(
            {
                "algorithm": algorithm,
                "status": "ok",
                "keypoints_a": len(last.keypoints_a),
                "keypoints_b": len(last.keypoints_b),
                "good_matches": len(last.good_matches),
                "inliers": last.inliers,
                "inlier_ratio": round(last.inlier_ratio, 6),
                "median_total_ms": round(median(times), 3),
                "repeats": args.repeats,
                "message": "",
            }
        )

    args.output_dir.mkdir(parents=True, exist_ok=True)
    fields = ["algorithm", "status", "keypoints_a", "keypoints_b", "good_matches", "inliers", "inlier_ratio", "median_total_ms", "repeats", "message"]
    with (args.output_dir / "benchmark_results.csv").open("w", encoding="utf-8-sig", newline="") as f:
        writer = csv.DictWriter(f, fieldnames=fields)
        writer.writeheader()
        writer.writerows(rows)
    save_json(args.output_dir / "benchmark_results.json", rows)

    valid = [r for r in rows if r["status"] == "ok"]
    if valid:
        write_image(args.output_dir / "benchmark_chart.png", draw_benchmark_chart(valid))

    for row in rows:
        print(row)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
