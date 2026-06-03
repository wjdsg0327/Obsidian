"""
特征提取、匹配、霍夫变换与模板匹配
知识点 4：模板匹配

来源：docs/07_features_matching_and_hough.md
"""
import cv2
import numpy as np


def create_template():
    icon = np.full((70, 70, 3), 255, dtype=np.uint8)
    cv2.rectangle(icon, (5, 5), (64, 64), (0, 0, 0), 2)
    cv2.circle(icon, (35, 35), 18, (0, 120, 255), -1)
    cv2.line(icon, (20, 20), (50, 50), (255, 255, 255), 3)
    cv2.line(icon, (50, 20), (20, 50), (255, 255, 255), 3)
    return icon


def create_scene(template):
    scene = np.full((260, 420, 3), 235, dtype=np.uint8)
    cv2.rectangle(scene, (20, 20), (160, 100), (255, 180, 0), -1)
    cv2.circle(scene, (330, 70), 40, (0, 200, 120), -1)
    cv2.putText(scene, "Template Matching", (40, 220), cv2.FONT_HERSHEY_SIMPLEX, 0.9, (30, 30, 30), 2)

    y, x = 110, 260
    scene[y:y + template.shape[0], x:x + template.shape[1]] = template
    return scene


def main():
    template = create_template()
    scene = create_scene(template)

    result = cv2.matchTemplate(scene, template, cv2.TM_CCOEFF_NORMED)
    _, max_val, _, max_loc = cv2.minMaxLoc(result)

    h, w = template.shape[:2]
    top_left = max_loc
    bottom_right = (top_left[0] + w, top_left[1] + h)

    output = scene.copy()
    cv2.rectangle(output, top_left, bottom_right, (0, 0, 255), 3)
    cv2.putText(output, f"score={max_val:.2f}", (top_left[0] - 10, top_left[1] - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.7, (0, 0, 255), 2)

    display_size = (template.shape[1] * 3, template.shape[0] * 3)
    template_large = cv2.resize(template, display_size, interpolation=cv2.INTER_NEAREST)
    scene_large = cv2.resize(scene, display_size)
    output_large = cv2.resize(output, display_size)
    panel = cv2.hconcat([template_large, scene_large, output_large])

    cv2.imshow("template matching", panel)
    cv2.imwrite("template_matching.png", panel)
    cv2.waitKey(0)
    cv2.destroyAllWindows()


if __name__ == "__main__":
    main()
