from __future__ import annotations

import cv2 as cv
import numpy as np

from common import SAMPLE_DIR, write_image


def textured_canvas(width: int, height: int, seed: int) -> np.ndarray:
    rng = np.random.default_rng(seed)
    image = np.full((height, width, 3), (235, 238, 242), dtype=np.uint8)
    for x in range(0, width, 40):
        cv.line(image, (x, 0), (x, height - 1), (205, 210, 218), 1)
    for y in range(0, height, 40):
        cv.line(image, (0, y), (width - 1, y), (205, 210, 218), 1)
    for _ in range(max(80, width * height // 4000)):
        center = (int(rng.integers(8, width - 8)), int(rng.integers(8, height - 8)))
        radius = int(rng.integers(2, 8))
        color = tuple(int(v) for v in rng.integers(25, 215, size=3))
        cv.circle(image, center, radius, color, -1, cv.LINE_AA)
    return image


def make_object() -> np.ndarray:
    obj = textured_canvas(360, 260, 20260726)
    cv.rectangle(obj, (6, 6), (353, 253), (15, 23, 42), 7)
    cv.putText(obj, "SIFT", (35, 92), cv.FONT_HERSHEY_DUPLEX, 2.2, (15, 23, 42), 4, cv.LINE_AA)
    cv.putText(obj, "SURF", (122, 170), cv.FONT_HERSHEY_DUPLEX, 1.6, (180, 35, 35), 3, cv.LINE_AA)
    cv.circle(obj, (65, 200), 29, (28, 135, 84), 6, cv.LINE_AA)
    cv.line(obj, (245, 195), (325, 225), (10, 80, 180), 8, cv.LINE_AA)
    cv.line(obj, (325, 195), (245, 225), (10, 80, 180), 8, cv.LINE_AA)
    return obj


def make_scene(obj: np.ndarray) -> np.ndarray:
    scene = textured_canvas(820, 600, 20260727)
    h, w = obj.shape[:2]
    source = np.float32([[0, 0], [w - 1, 0], [w - 1, h - 1], [0, h - 1]])
    target = np.float32([[180, 145], [635, 95], [680, 455], [135, 485]])
    transform = cv.getPerspectiveTransform(source, target)
    warped = cv.warpPerspective(obj, transform, (scene.shape[1], scene.shape[0]))
    mask = cv.warpPerspective(np.full((h, w), 255, np.uint8), transform, (scene.shape[1], scene.shape[0]))
    scene[mask > 0] = warped[mask > 0]
    cv.putText(scene, "PLANAR SCENE", (20, 45), cv.FONT_HERSHEY_SIMPLEX, 1.1, (35, 40, 55), 2, cv.LINE_AA)
    return scene


def make_panorama_pair() -> tuple[np.ndarray, np.ndarray]:
    base = textured_canvas(1100, 520, 20260728)
    for i, label in enumerate(["A", "B", "C", "D", "E"]):
        x = 80 + i * 210
        cv.rectangle(base, (x, 120), (x + 120, 330), (30 + i * 30, 70, 170 - i * 20), 5)
        cv.putText(base, label, (x + 30, 255), cv.FONT_HERSHEY_DUPLEX, 2.4, (20, 25, 35), 4, cv.LINE_AA)
    cv.putText(base, "LOCAL FEATURES PANORAMA", (250, 465), cv.FONT_HERSHEY_DUPLEX, 1.2, (30, 35, 45), 3, cv.LINE_AA)
    return base[:, :720].copy(), base[:, 380:].copy()


def main() -> int:
    SAMPLE_DIR.mkdir(parents=True, exist_ok=True)
    obj = make_object()
    scene = make_scene(obj)
    pano_a, pano_b = make_panorama_pair()
    write_image(SAMPLE_DIR / "synthetic_object.png", obj)
    write_image(SAMPLE_DIR / "synthetic_scene.png", scene)
    write_image(SAMPLE_DIR / "synthetic_panorama_a.png", pano_a)
    write_image(SAMPLE_DIR / "synthetic_panorama_b.png", pano_b)
    print(f"已生成样例：{SAMPLE_DIR}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

