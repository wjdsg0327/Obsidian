"""SIFT/SURF 学习包的公共函数。"""

from __future__ import annotations

import json
from dataclasses import dataclass
from pathlib import Path
from time import perf_counter

import cv2 as cv
import numpy as np


PACKAGE_ROOT = Path(__file__).resolve().parents[1]
SAMPLE_DIR = PACKAGE_ROOT / "09_样例与数据集" / "opencv_samples"
DEFAULT_IMAGE_A = SAMPLE_DIR / "synthetic_object.png"
DEFAULT_IMAGE_B = SAMPLE_DIR / "synthetic_scene.png"
SURF_GUIDE = PACKAGE_ROOT / "03_SURF" / "SURF_Windows独立编译指南.md"


class DetectorUnavailable(RuntimeError):
    """当前 OpenCV 构建无法提供请求的检测器。"""


@dataclass
class MatchResult:
    keypoints_a: list
    keypoints_b: list
    descriptors_a: np.ndarray
    descriptors_b: np.ndarray
    good_matches: list
    homography: np.ndarray | None
    inlier_mask: np.ndarray | None
    detect_ms: float
    match_ms: float

    @property
    def inliers(self) -> int:
        return int(self.inlier_mask.sum()) if self.inlier_mask is not None else 0

    @property
    def inlier_ratio(self) -> float:
        return self.inliers / len(self.good_matches) if self.good_matches else 0.0


def read_image(path: str | Path, flags: int = cv.IMREAD_COLOR) -> np.ndarray:
    """兼容中文 Windows 路径的图像读取。"""
    path = Path(path)
    if not path.exists():
        raise FileNotFoundError(f"图像不存在：{path}")
    data = np.fromfile(path, dtype=np.uint8)
    image = cv.imdecode(data, flags)
    if image is None:
        raise ValueError(f"OpenCV 无法解码图像：{path}")
    return image


def write_image(path: str | Path, image: np.ndarray) -> Path:
    """兼容中文 Windows 路径的图像保存。"""
    path = Path(path)
    path.parent.mkdir(parents=True, exist_ok=True)
    suffix = path.suffix or ".png"
    ok, encoded = cv.imencode(suffix, image)
    if not ok:
        raise ValueError(f"OpenCV 无法编码输出：{path}")
    encoded.tofile(path)
    return path


def create_detector(name: str):
    name = name.lower()
    if name == "sift":
        if not hasattr(cv, "SIFT_create"):
            raise DetectorUnavailable("当前 OpenCV 不含 SIFT_create。")
        return cv.SIFT_create(), cv.NORM_L2
    if name == "surf":
        factory = getattr(getattr(cv, "xfeatures2d", None), "SURF_create", None)
        if factory is None:
            raise DetectorUnavailable(
                "当前 OpenCV 未启用 SURF（通常是 Non-free algorithms: NO）。\n"
                f"请阅读独立编译指南：{SURF_GUIDE}"
            )
        try:
            return factory(400), cv.NORM_L2
        except cv.error as exc:
            raise DetectorUnavailable(
                "SURF 接口存在但当前构建未启用 nonfree 算法。\n"
                f"请阅读独立编译指南：{SURF_GUIDE}\n原始错误：{exc}"
            ) from exc
    if name == "orb":
        return cv.ORB_create(nfeatures=2500), cv.NORM_HAMMING
    if name == "akaze":
        return cv.AKAZE_create(), cv.NORM_HAMMING
    raise ValueError(f"未知算法：{name}；可选 sift/surf/orb/akaze")


def detect_and_match(
    gray_a: np.ndarray,
    gray_b: np.ndarray,
    algorithm: str = "sift",
    ratio: float = 0.75,
    ransac_threshold: float = 5.0,
) -> MatchResult:
    if not 0.0 < ratio < 1.0:
        raise ValueError("ratio 必须位于 (0, 1) 区间。")

    detector, norm = create_detector(algorithm)
    start = perf_counter()
    keypoints_a, descriptors_a = detector.detectAndCompute(gray_a, None)
    keypoints_b, descriptors_b = detector.detectAndCompute(gray_b, None)
    detect_ms = (perf_counter() - start) * 1000.0

    if descriptors_a is None or descriptors_b is None:
        raise RuntimeError("没有得到描述子；图片可能太小、过于平坦或阈值过高。")

    matcher = cv.BFMatcher(normType=norm, crossCheck=False)
    start = perf_counter()
    pairs = matcher.knnMatch(descriptors_a, descriptors_b, k=2)
    good = [pair[0] for pair in pairs if len(pair) == 2 and pair[0].distance < ratio * pair[1].distance]
    match_ms = (perf_counter() - start) * 1000.0

    homography = None
    inlier_mask = None
    if len(good) >= 4:
        src = np.float32([keypoints_a[m.queryIdx].pt for m in good]).reshape(-1, 1, 2)
        dst = np.float32([keypoints_b[m.trainIdx].pt for m in good]).reshape(-1, 1, 2)
        homography, mask = cv.findHomography(src, dst, cv.RANSAC, ransac_threshold)
        if mask is not None:
            inlier_mask = mask.ravel().astype(np.uint8)

    return MatchResult(
        keypoints_a=keypoints_a,
        keypoints_b=keypoints_b,
        descriptors_a=descriptors_a,
        descriptors_b=descriptors_b,
        good_matches=good,
        homography=homography,
        inlier_mask=inlier_mask,
        detect_ms=detect_ms,
        match_ms=match_ms,
    )


def metrics_dict(result: MatchResult, algorithm: str, ratio: float) -> dict:
    return {
        "algorithm": algorithm,
        "ratio": ratio,
        "keypoints_a": len(result.keypoints_a),
        "keypoints_b": len(result.keypoints_b),
        "descriptor_shape_a": list(result.descriptors_a.shape),
        "descriptor_shape_b": list(result.descriptors_b.shape),
        "good_matches": len(result.good_matches),
        "inliers": result.inliers,
        "inlier_ratio": round(result.inlier_ratio, 6),
        "detect_and_describe_ms": round(result.detect_ms, 3),
        "match_ms": round(result.match_ms, 3),
        "homography_found": result.homography is not None,
    }


def draw_keypoints(image: np.ndarray, keypoints: list) -> np.ndarray:
    return cv.drawKeypoints(
        image,
        keypoints,
        None,
        flags=cv.DRAW_MATCHES_FLAGS_DRAW_RICH_KEYPOINTS,
    )


def draw_matches(image_a: np.ndarray, image_b: np.ndarray, result: MatchResult, inliers_only: bool = False) -> np.ndarray:
    matches = result.good_matches
    mask = None
    if inliers_only and result.inlier_mask is not None:
        mask = result.inlier_mask.tolist()
    return cv.drawMatches(
        image_a,
        result.keypoints_a,
        image_b,
        result.keypoints_b,
        matches,
        None,
        matchesMask=mask,
        flags=cv.DrawMatchesFlags_NOT_DRAW_SINGLE_POINTS,
    )


def save_json(path: str | Path, payload: dict | list) -> Path:
    path = Path(path)
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, ensure_ascii=False, indent=2), encoding="utf-8")
    return path
