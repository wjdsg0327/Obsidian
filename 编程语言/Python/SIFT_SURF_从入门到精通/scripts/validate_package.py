from __future__ import annotations

import json
import re
from pathlib import Path

import cv2 as cv
import numpy as np


ROOT = Path(__file__).resolve().parents[1]
OUTPUT = ROOT / "outputs" / "验收报告.md"


REQUIRED = [
    "README.md",
    "requirements.txt",
    "environment.yml",
    "资源索引.csv",
    "资源索引.xlsx",
    "07_论文与扩展资料/经典论文/经典论文下载入口.url",
    "07_论文与扩展资料/经典论文/论文精读指南.pdf",
    "09_样例与数据集/opencv_samples/synthetic_object.png",
    "09_样例与数据集/opencv_samples/synthetic_scene.png",
    "outputs/sift_matches.png",
    "outputs/planar_localization.png",
    "outputs/simple_panorama.png",
    "outputs/benchmark_chart.png",
    "outputs/notebook_01_sift_keypoints.png",
    "outputs/notebook_02_inlier_matches.png",
]


def is_image_ok(path: Path) -> bool:
    data = np.fromfile(path, dtype=np.uint8)
    return cv.imdecode(data, cv.IMREAD_UNCHANGED) is not None


def local_markdown_links() -> list[tuple[Path, str]]:
    broken = []
    pattern = re.compile(r"\[[^\]]+\]\(([^)]+)\)")
    for md in ROOT.rglob("*.md"):
        for raw in pattern.findall(md.read_text(encoding="utf-8")):
            target = raw.split("#", 1)[0]
            if not target or "://" in target or target.startswith("mailto:"):
                continue
            if not (md.parent / target).resolve().exists():
                broken.append((md.relative_to(ROOT), raw))
    return broken


def main() -> int:
    checks = []
    for relative in REQUIRED:
        path = ROOT / relative
        checks.append((f"必需文件：{relative}", path.exists()))

    for pdf in (ROOT / "07_论文与扩展资料" / "经典论文").glob("*.pdf"):
        checks.append((f"PDF 文件头：{pdf.name}", pdf.read_bytes()[:5] == b"%PDF-"))
    for image in (ROOT / "09_样例与数据集" / "opencv_samples").glob("*"):
        if image.is_file():
            checks.append((f"图像解码：{image.name}", is_image_ok(image)))

    broken_links = local_markdown_links()
    checks.append(("Markdown 本地链接", not broken_links))
    manifest = json.loads((ROOT / "resources_manifest.json").read_text(encoding="utf-8"))
    checks.append(("资源清单不少于 12 项", len(manifest) >= 12))
    checks.append(("资源索引 CSV 为 UTF-8 BOM", (ROOT / "资源索引.csv").read_bytes().startswith(b"\xef\xbb\xbf")))
    checks.append(("资源索引 XLSX 文件头", (ROOT / "资源索引.xlsx").read_bytes().startswith(b"PK")))
    for notebook in (ROOT / "notebooks").glob("*.ipynb"):
        payload = json.loads(notebook.read_text(encoding="utf-8"))
        code_cells = [cell for cell in payload["cells"] if cell["cell_type"] == "code"]
        checks.append((f"Notebook 已执行：{notebook.name}", bool(code_cells) and all(cell.get("execution_count") is not None for cell in code_cells)))

    passed = sum(ok for _, ok in checks)
    lines = ["# 学习包验收报告", "", f"通过：{passed}/{len(checks)}", "", "| 检查项 | 结果 |", "|---|---|"]
    lines.extend(f"| {name} | {'通过' if ok else '失败'} |" for name, ok in checks)
    if broken_links:
        lines.extend(["", "## 损坏的本地链接", ""])
        lines.extend(f"- `{path}` -> `{target}`" for path, target in broken_links)
    OUTPUT.parent.mkdir(parents=True, exist_ok=True)
    OUTPUT.write_text("\n".join(lines), encoding="utf-8")
    print(f"验收：{passed}/{len(checks)}；报告：{OUTPUT}")
    return 0 if passed == len(checks) else 1


if __name__ == "__main__":
    raise SystemExit(main())
