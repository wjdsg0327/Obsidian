from __future__ import annotations

import argparse
import json
import platform
import sys

import cv2 as cv
import numpy as np

from common import SURF_GUIDE, create_detector, DetectorUnavailable


def nonfree_status() -> str:
    for line in cv.getBuildInformation().splitlines():
        if "Non-free algorithms:" in line:
            return line.split(":", 1)[1].strip()
    return "UNKNOWN"


def main() -> int:
    parser = argparse.ArgumentParser(description="检查 SIFT/SURF 学习环境")
    parser.add_argument("--json", action="store_true", help="输出 JSON")
    args = parser.parse_args()

    detectors = {}
    for name in ("sift", "surf", "orb", "akaze"):
        try:
            create_detector(name)
            detectors[name] = {"available": True, "message": "可用"}
        except DetectorUnavailable as exc:
            detectors[name] = {"available": False, "message": str(exc)}

    report = {
        "python": sys.version.split()[0],
        "platform": platform.platform(),
        "opencv": cv.__version__,
        "opencv_path": cv.__file__,
        "numpy": np.__version__,
        "nonfree_algorithms": nonfree_status(),
        "detectors": detectors,
        "surf_guide": str(SURF_GUIDE),
    }

    if args.json:
        print(json.dumps(report, ensure_ascii=False, indent=2))
    else:
        print("SIFT/SURF 学习环境检查")
        print(f"Python: {report['python']}")
        print(f"OpenCV: {report['opencv']} ({report['opencv_path']})")
        print(f"NumPy: {report['numpy']}")
        print(f"Non-free algorithms: {report['nonfree_algorithms']}")
        for name, info in detectors.items():
            marker = "OK" if info["available"] else "--"
            first_line = info["message"].splitlines()[0]
            print(f"[{marker}] {name.upper()}: {first_line}")
        if not detectors["surf"]["available"]:
            print(f"SURF 指南：{SURF_GUIDE}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

