"""
摄像头、视频与实时处理
知识点 1：读取摄像头或模拟视频流

来源：docs/06_video_and_camera.md
"""
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
