from __future__ import annotations

import json
import ssl
import urllib.error
import urllib.request
from concurrent.futures import ThreadPoolExecutor, as_completed
from datetime import datetime
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
MANIFEST = ROOT / "resources_manifest.json"
OUTPUT = ROOT / "outputs" / "失效链接报告.md"


def check(url: str) -> tuple[str, str]:
    request = urllib.request.Request(url, method="HEAD", headers={"User-Agent": "Mozilla/5.0 SIFT-SURF-study-pack"})
    try:
        with urllib.request.urlopen(request, timeout=8, context=ssl.create_default_context()) as response:
            return str(response.status), response.geturl()
    except urllib.error.HTTPError as exc:
        if exc.code in (403, 405):
            get_request = urllib.request.Request(url, headers={"User-Agent": "Mozilla/5.0 SIFT-SURF-study-pack", "Range": "bytes=0-0"})
            try:
                with urllib.request.urlopen(get_request, timeout=8) as response:
                    return str(response.status), response.geturl()
            except Exception as inner:  # noqa: BLE001
                return f"ERROR {type(inner).__name__}", str(inner)
        return str(exc.code), str(exc)
    except Exception as exc:  # noqa: BLE001
        return f"ERROR {type(exc).__name__}", str(exc)


def main() -> int:
    resources = json.loads(MANIFEST.read_text(encoding="utf-8"))
    lines = ["# 链接检查报告", "", f"检查时间：{datetime.now().isoformat(timespec='seconds')}", "", "| 状态 | 资源 | URL | 说明 |", "|---|---|---|---|"]
    results = {}
    online_items = [item for item in resources if item["source_url"]]
    with ThreadPoolExecutor(max_workers=8) as pool:
        futures = {pool.submit(check, item["source_url"]): item for item in online_items}
        for future in as_completed(futures):
            item = futures[future]
            results[item["source_url"]] = future.result()

    http_failures = 0
    unverified = 0
    for item in resources:
        if not item["source_url"]:
            lines.append(f"| 本地生成 | {item['title']} |  | 无外部链接 |")
            continue
        status, detail = results[item["source_url"]]
        ok = status.startswith("2") or status.startswith("3")
        if status.startswith("ERROR"):
            unverified += 1
        elif not ok:
            http_failures += 1
        label = "正常" if ok else ("网络环境无法验证" if status.startswith("ERROR") else status)
        lines.append(f"| {label} | {item['title']} | {item['source_url']} | {detail.replace('|', '/')} |")
    lines.extend([
        "",
        f"明确 HTTP 异常：{http_failures}",
        f"因当前网络环境无法验证：{unverified}",
        "",
        "说明：连接失败表示当前运行环境无法联网，不等同于链接失效；403 也可能表示站点拒绝自动检查。",
    ])
    OUTPUT.parent.mkdir(parents=True, exist_ok=True)
    OUTPUT.write_text("\n".join(lines), encoding="utf-8")
    print(f"报告：{OUTPUT}，HTTP 异常：{http_failures}，网络环境无法验证：{unverified}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
