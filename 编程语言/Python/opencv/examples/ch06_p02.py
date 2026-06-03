"""
摄像头、视频与实时处理
知识点 2：实时灰度化、边缘检测与 FPS 显示

来源：docs/06_video_and_camera.md
"""
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
