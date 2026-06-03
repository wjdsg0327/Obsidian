"""
综合实战与性能优化
知识点 3：性能分析与批处理模板

来源：docs/09_projects_and_optimization.md
"""
from pathlib import Path
import cv2
import numpy as np


class BatchProcessor:
    def __init__(self, input_dir="input_images", output_dir="batch_outputs", target_width=640):
        self.input_dir = Path(input_dir)
        self.output_dir = Path(output_dir)
        self.target_width = target_width
        self.output_dir.mkdir(exist_ok=True)

    def generate_samples(self, count=5):
        samples = []
        for i in range(count):
            image = np.full((480, 720, 3), 30, dtype=np.uint8)
            cv2.circle(image, (120 + i * 80, 140 + i * 20), 45, (0, 220, 255), -1)
            cv2.rectangle(image, (320, 100 + i * 20), (560, 260 + i * 20), (255, 120, 0), 3)
            cv2.putText(image, f"sample {i}", (40, 420), cv2.FONT_HERSHEY_SIMPLEX, 1.4, (255, 255, 255), 3)
            samples.append((f"sample_{i}.png", image))
        return samples

    def load_images(self):
        patterns = ["*.png", "*.jpg", "*.jpeg", "*.bmp"]
        files = []
        for pattern in patterns:
            files.extend(sorted(self.input_dir.glob(pattern)))

        if not files:
            return self.generate_samples()

        images = []
        for path in files:
            image = cv2.imread(str(path))
            if image is not None:
                images.append((path.name, image))
        return images

    def process(self, image):
        h, w = image.shape[:2]
        if w > self.target_width:
            scale = self.target_width / w
            image = cv2.resize(image, None, fx=scale, fy=scale, interpolation=cv2.INTER_AREA)

        gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
        blur = cv2.GaussianBlur(gray, (5, 5), 0)
        edges = cv2.Canny(blur, 60, 150)
        result = cv2.cvtColor(edges, cv2.COLOR_GRAY2BGR)
        return result

    def run(self):
        items = self.load_images()
        meter = cv2.TickMeter()

        meter.start()
        for name, image in items:
            result = self.process(image)
            output_path = self.output_dir / f"processed_{name}"
            cv2.imwrite(str(output_path), result)
        meter.stop()

        total_ms = meter.getTimeMilli()
        avg_ms = total_ms / max(len(items), 1)

        print(f"processed images: {len(items)}")
        print(f"total time: {total_ms:.2f} ms")
        print(f"average per image: {avg_ms:.2f} ms")
        print(f"outputs saved to: {self.output_dir.resolve()}")


def main():
    processor = BatchProcessor()
    processor.run()


if __name__ == "__main__":
    main()
