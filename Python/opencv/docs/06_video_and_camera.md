# 第 6 章：摄像头、视频与实时处理

本章目标：掌握视频流读取、逐帧处理和视频保存的基本方法，理解 OpenCV 在实时场景中的工作方式。

## 知识点 1：读取摄像头或模拟视频流

### 说明

`cv2.VideoCapture()` 可以读取摄像头，也可以读取视频文件。学习时最常见的写法是：

```python
cap = cv2.VideoCapture(0)
```

这里的 `0` 通常表示默认摄像头。如果没有摄像头，或者环境不支持摄像头访问，你也可以用程序生成“模拟视频帧”来练习逐帧处理逻辑。

### 案例

下面的代码会优先尝试打开摄像头；如果失败，就自动切换成一个带运动小球的模拟视频流。

### 完整代码

```python
import cv2
import numpy as np


def generate_frame(frame_id, width=640, height=360):
    frame = np.full((height, width, 3), 30, dtype=np.uint8)
    x = 60 + (frame_id * 8) % (width - 120)
    y = height // 2 + int(70 * np.sin(frame_id / 8))

    cv2.circle(frame, (x, y), 35, (0, 220, 255), -1)
    cv2.putText(frame, "Synthetic Stream", (20, 35), cv2.FONT_HERSHEY_SIMPLEX, 1.0, (255, 255, 255), 2)
    cv2.putText(frame, f"frame: {frame_id}", (20, 75), cv2.FONT_HERSHEY_SIMPLEX, 0.9, (0, 255, 0), 2)
    return frame


def main():
    cap = cv2.VideoCapture(0)
    use_camera = cap.isOpened()
    frame_id = 0

    while True:
        if use_camera:
            ok, frame = cap.read()
            if not ok:
                break
        else:
            frame = generate_frame(frame_id)

        frame_id += 1
        cv2.imshow("video stream", frame)

        key = cv2.waitKey(30) & 0xFF
        if key == ord("q"):
            break

    if use_camera:
        cap.release()
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 2：实时灰度化、边缘检测与 FPS 显示

### 说明

视频处理的本质是“对每一帧重复图像处理流程”。因此你可以把已经学过的灰度化、滤波、边缘提取等操作直接放到循环里。为了观察程序性能，我们还常常把 FPS 显示出来。

### 案例

下面的代码会持续读取视频帧，把它转为灰度图并做 Canny 边缘检测，同时在画面左上角显示 FPS。

### 完整代码

```python
import time
import cv2
import numpy as np


def generate_frame(frame_id, width=640, height=360):
    frame = np.full((height, width, 3), 25, dtype=np.uint8)
    x = 60 + (frame_id * 10) % (width - 120)
    y = height // 2 + int(80 * np.sin(frame_id / 10))
    cv2.circle(frame, (x, y), 30, (255, 120, 0), -1)
    cv2.rectangle(frame, (420, 110), (560, 250), (0, 200, 120), 3)
    return frame


def main():
    cap = cv2.VideoCapture(0)
    use_camera = cap.isOpened()
    frame_id = 0
    prev_time = time.time()

    while True:
        if use_camera:
            ok, frame = cap.read()
            if not ok:
                break
        else:
            frame = generate_frame(frame_id)

        frame_id += 1

        gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
        blur = cv2.GaussianBlur(gray, (5, 5), 0)
        edges = cv2.Canny(blur, 60, 150)
        edges_bgr = cv2.cvtColor(edges, cv2.COLOR_GRAY2BGR)

        current_time = time.time()
        fps = 1.0 / max(current_time - prev_time, 1e-6)
        prev_time = current_time

        cv2.putText(frame, f"FPS: {fps:.1f}", (20, 35), cv2.FONT_HERSHEY_SIMPLEX, 1.0, (0, 255, 255), 2)

        panel = cv2.hconcat([frame, edges_bgr])
        cv2.imshow("real time processing", panel)

        key = cv2.waitKey(1) & 0xFF
        if key == ord("q"):
            break

    if use_camera:
        cap.release()
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
```

## 知识点 3：把处理结果保存为视频文件

### 说明

实时处理完的结果往往还需要存盘，这时就要用 `cv2.VideoWriter()`。它的关键参数有：

- 输出文件名
- 编码器，例如 `mp4v`
- 帧率
- 输出分辨率

如果写出的视频无法播放，通常是编码器和系统环境不匹配，可以尝试换成 `XVID` 或 `MJPG`。

### 案例

下面的代码会生成一段 150 帧的模拟视频，并把结果保存为 `demo_output.mp4`。

### 完整代码

```python
import cv2
import numpy as np


def generate_frame(frame_id, width=640, height=360):
    frame = np.full((height, width, 3), 20, dtype=np.uint8)
    x = 50 + (frame_id * 7) % (width - 100)
    y = 80 + (frame_id * 3) % (height - 160)
    cv2.circle(frame, (x, y), 28, (0, 220, 255), -1)
    cv2.rectangle(frame, (width - 180, 80), (width - 60, 220), (255, 100, 0), -1)
    cv2.putText(frame, f"frame {frame_id:03d}", (20, 35), cv2.FONT_HERSHEY_SIMPLEX, 1.0, (255, 255, 255), 2)
    return frame


def main():
    width, height = 640, 360
    writer = cv2.VideoWriter(
        "demo_output.mp4",
        cv2.VideoWriter_fourcc(*"mp4v"),
        25,
        (width, height),
    )

    for frame_id in range(150):
        frame = generate_frame(frame_id, width, height)
        gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
        edges = cv2.Canny(gray, 60, 150)
        edges_bgr = cv2.cvtColor(edges, cv2.COLOR_GRAY2BGR)
        output = cv2.addWeighted(frame, 0.75, edges_bgr, 0.25, 0)
        writer.write(output)

    writer.release()
    print("saved: demo_output.mp4")


if __name__ == "__main__":
    main()
```
