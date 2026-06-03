"""
摄像头、视频与实时处理
知识点 3：把处理结果保存为视频文件

来源：docs/06_video_and_camera.md
"""
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
