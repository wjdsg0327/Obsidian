from __future__ import annotations

import re
from pathlib import Path


DOCS_DIR = Path("docs")
EXAMPLES_DIR = Path("examples")
EXERCISES_DIR = Path("exercises")

CHAPTER_META = {
    "00_setup_and_first_program.md": ("第 0 章：环境准备与第一段 OpenCV 代码", "环境准备与第一段 OpenCV 代码"),
    "01_image_io_and_pixels.md": ("第 1 章：图像读写、像素访问与 ROI", "图像读写、像素访问与 ROI"),
    "02_drawing_and_color_spaces.md": ("第 2 章：绘图、文字与颜色空间", "绘图、文字与颜色空间"),
    "03_geometric_transforms.md": ("第 3 章：缩放、旋转与几何变换", "缩放、旋转与几何变换"),
    "04_filtering_threshold_and_enhancement.md": ("第 4 章：滤波、阈值与图像增强", "滤波、阈值与图像增强"),
    "05_edges_morphology_and_contours.md": ("第 5 章：边缘、形态学与轮廓分析", "边缘、形态学与轮廓分析"),
    "06_video_and_camera.md": ("第 6 章：摄像头、视频与实时处理", "摄像头、视频与实时处理"),
    "07_features_matching_and_hough.md": ("第 7 章：特征提取、匹配、霍夫变换与模板匹配", "特征提取、匹配、霍夫变换与模板匹配"),
    "08_segmentation_and_advanced_processing.md": ("第 8 章：图像分割与高级处理", "图像分割与高级处理"),
    "09_projects_and_optimization.md": ("第 9 章：综合实战与性能优化", "综合实战与性能优化"),
}

EXERCISE_CONTENT = {
    "00": {
        "title": "环境准备与第一段 OpenCV 代码",
        "goals": [
            "确认 `opencv-python` 和 `numpy` 已正确安装。",
            "理解图像是 `numpy.ndarray` 这一基本概念。",
            "学会用 `cv2.imwrite()` 保存第一张测试图像。",
        ],
        "tasks": [
            ("基础练习 1", "编写一个脚本，打印 OpenCV 版本号、NumPy 版本号，并生成一张 300x300 的纯黑图。"),
            ("基础练习 2", "在纯黑图中心写上你的名字拼音，并保存成 `my_first_opencv.png`。"),
            ("进阶练习", "把背景色改成深蓝色，文字颜色改成亮黄色，再额外画一条对角线。"),
            ("思考题", "为什么 OpenCV 图像对象适合用 NumPy 来表示？请从数据结构角度简单总结。"),
        ],
    },
    "01": {
        "title": "图像读写、像素访问与 ROI",
        "goals": [
            "掌握 `imread()`、`imshow()`、`imwrite()` 的基本用法。",
            "理解像素访问和 NumPy 切片的关系。",
            "学会裁剪和复制局部区域。",
        ],
        "tasks": [
            ("基础练习 1", "自己创建一张 400x300 的测试图，并把左上角 80x80 区域改成红色。"),
            ("基础练习 2", "读取一张图片后，打印左上角、中心点、右下角三个像素的 BGR 值。"),
            ("进阶练习", "从图像中裁剪一个 ROI，把它复制到图像右下角，再保存结果。"),
            ("思考题", "为什么在处理大面积区域时，切片赋值通常比双层 for 循环更推荐？"),
        ],
    },
    "02": {
        "title": "绘图、文字与颜色空间",
        "goals": [
            "学会绘制常用几何图形和文字。",
            "理解 BGR、灰度、HSV 三类常见表示。",
            "学会做简单的图像标注和通道观察。",
        ],
        "tasks": [
            ("基础练习 1", "创建一张 500x400 画布，至少画出线、矩形、圆、文字四种元素。"),
            ("基础练习 2", "把一张彩色测试图转换成灰度图和 HSV 图，并分别保存。"),
            ("进阶练习", "分离一张彩色图的 B、G、R 三个通道，并以三张灰度图形式输出。"),
            ("思考题", "什么时候更适合用 HSV 而不是 BGR 做颜色分割？"),
        ],
    },
    "03": {
        "title": "缩放、旋转与几何变换",
        "goals": [
            "掌握缩放、翻转、旋转、平移。",
            "理解仿射变换与透视变换的差异。",
            "学会用几何变换完成简单矫正任务。",
        ],
        "tasks": [
            ("基础练习 1", "把同一张图分别缩小到 50% 和放大到 200%，比较效果差异。"),
            ("基础练习 2", "将一张图顺时针旋转 30 度，再水平翻转。"),
            ("进阶练习", "自己构造四个点，完成一次透视变换，把倾斜矩形拉正。"),
            ("思考题", "为什么透视变换适合做文档扫描，而普通旋转和平移不够？"),
        ],
    },
    "04": {
        "title": "滤波、阈值与图像增强",
        "goals": [
            "理解均值、高斯、中值滤波的作用差异。",
            "掌握全局阈值、自适应阈值和 Otsu。",
            "学会使用直方图均衡化提升对比度。",
        ],
        "tasks": [
            ("基础练习 1", "生成一张带噪声图像，分别用均值滤波和高斯滤波处理。"),
            ("基础练习 2", "把一张灰度图分别做普通阈值和 Otsu 阈值，比较结果。"),
            ("进阶练习", "准备一张偏暗图像，对比普通均衡化和 CLAHE 的增强效果。"),
            ("思考题", "为什么中值滤波对椒盐噪声比均值滤波更有效？"),
        ],
    },
    "05": {
        "title": "边缘、形态学与轮廓分析",
        "goals": [
            "理解边缘检测的用途。",
            "掌握腐蚀、膨胀、开闭运算。",
            "学会通过轮廓提取目标几何信息。",
        ],
        "tasks": [
            ("基础练习 1", "使用 Canny 检测一张图中的边缘，并调节两组阈值观察区别。"),
            ("基础练习 2", "对一张有噪点的二值图做开运算和闭运算。"),
            ("进阶练习", "查找图像中的所有外轮廓，并标出面积最大的目标。"),
            ("思考题", "为什么轮廓分析前通常需要先二值化或提边？"),
        ],
    },
    "06": {
        "title": "摄像头、视频与实时处理",
        "goals": [
            "掌握逐帧读取视频流的基本循环。",
            "学会把图像处理流程迁移到实时场景。",
            "理解视频写出和 FPS 统计的基本方法。",
        ],
        "tasks": [
            ("基础练习 1", "打开默认摄像头或模拟视频流，显示实时画面。"),
            ("基础练习 2", "把实时视频改成灰度显示，并按 `q` 退出。"),
            ("进阶练习", "把实时边缘检测结果写入一段 MP4 视频。"),
            ("思考题", "为什么视频处理通常要更关注图像尺寸和算法复杂度？"),
        ],
    },
    "07": {
        "title": "特征提取、匹配、霍夫变换与模板匹配",
        "goals": [
            "理解关键点和描述子的基本含义。",
            "掌握 ORB 特征和暴力匹配器的基本用法。",
            "学会用霍夫变换和模板匹配解决定位问题。",
        ],
        "tasks": [
            ("基础练习 1", "对一张纹理丰富的图像提取 ORB 关键点，并绘制结果。"),
            ("基础练习 2", "用 ORB 匹配两张相似图像，显示前 30 个匹配结果。"),
            ("进阶练习", "自己生成带多条直线的图像，使用 `HoughLinesP` 检出这些线段。"),
            ("思考题", "模板匹配和特征匹配分别适合什么场景？它们的局限在哪里？"),
        ],
    },
    "08": {
        "title": "图像分割与高级处理",
        "goals": [
            "理解连通域分析的输出信息。",
            "学会用分水岭分离粘连目标。",
            "认识 GrabCut 的前景提取思路。",
        ],
        "tasks": [
            ("基础练习 1", "准备一张含多个白色目标的二值图，统计连通域数量和面积。"),
            ("基础练习 2", "制作两三个相互接触的圆形目标，尝试用分水岭分开它们。"),
            ("进阶练习", "用 GrabCut 从一张图中抠出主要前景，并保存掩膜结果。"),
            ("思考题", "为什么分水岭需要‘确定前景’和‘确定背景’这两个步骤？"),
        ],
    },
    "09": {
        "title": "综合实战与性能优化",
        "goals": [
            "把多个基础算子串成完整处理流程。",
            "理解文档扫描和目标计数这类小项目的实现思路。",
            "建立基础的性能分析和批处理意识。",
        ],
        "tasks": [
            ("基础练习 1", "参考文档扫描案例，自己画一张倾斜矩形并尝试透视拉正。"),
            ("基础练习 2", "生成多个圆形目标，完成面积过滤与自动计数。"),
            ("进阶练习", "写一个批处理脚本，对一个目录中的图片统一执行灰度化、模糊和边缘检测。"),
            ("思考题", "如果一段 OpenCV 程序运行很慢，你会按什么顺序排查问题？"),
        ],
    },
}


def parse_points(markdown: str) -> list[dict[str, str]]:
    pattern = re.compile(
        r"## 知识点\s*(\d+)：(.+?)\n(.*?)(?=\n## 知识点\s*\d+：|\Z)",
        re.S,
    )
    code_pattern = re.compile(r"### 完整代码\s+```python\n(.*?)```", re.S)
    points = []
    for match in pattern.finditer(markdown):
        point_no = match.group(1).strip()
        title = match.group(2).strip()
        body = match.group(3)
        code_match = code_pattern.search(body)
        if not code_match:
            raise ValueError(f"未在知识点 {point_no} 中找到 Python 代码块")
        points.append(
            {
                "point_no": point_no,
                "title": title,
                "code": code_match.group(1).strip() + "\n",
            }
        )
    return points


def build_examples() -> None:
    EXAMPLES_DIR.mkdir(exist_ok=True)
    readme_lines = [
        "# OpenCV 配套示例脚本",
        "",
        "这个目录收集了从 `docs/` 教程中拆分出来的可直接运行示例脚本。",
        "",
        "如果运行脚本时出现 `ModuleNotFoundError: No module named 'cv2'`，先安装依赖：",
        "",
        "```bash",
        "pip install -r requirements.txt",
        "```",
        "",
        "## 使用方式",
        "",
        "```bash",
        "python examples/ch00_p01.py",
        "```",
        "",
        "大部分脚本会弹出 `cv2.imshow()` 窗口，按任意键关闭；视频类脚本一般按 `q` 退出。",
        "",
        "## 文件索引",
        "",
    ]

    for doc_name, (chapter_heading, chapter_short) in CHAPTER_META.items():
        path = DOCS_DIR / doc_name
        markdown = path.read_text(encoding="utf-8")
        points = parse_points(markdown)

        readme_lines.append(f"### {chapter_heading}")
        readme_lines.append("")

        chapter_prefix = doc_name.split("_", 1)[0]
        for point in points:
            script_name = f"ch{chapter_prefix}_p{int(point['point_no']):02d}.py"
            script_path = EXAMPLES_DIR / script_name
            header = [
                '"""',
                f"{chapter_short}",
                f"知识点 {point['point_no']}：{point['title']}",
                "",
                f"来源：docs/{doc_name}",
                '"""',
                "",
            ]
            script_path.write_text("\n".join(header) + point["code"], encoding="utf-8")
            readme_lines.append(f"- `{script_name}`: 知识点 {point['point_no']} - {point['title']}")

        readme_lines.append("")

    (EXAMPLES_DIR / "README.md").write_text("\n".join(readme_lines).rstrip() + "\n", encoding="utf-8")


def build_exercises() -> None:
    EXERCISES_DIR.mkdir(exist_ok=True)
    readme_lines = [
        "# OpenCV 配套练习",
        "",
        "这个目录提供与教程章节一一对应的练习题，建议在学完一章后立刻完成当章练习。",
        "",
        "做练习前请先安装依赖：",
        "",
        "```bash",
        "pip install -r requirements.txt",
        "```",
        "",
        "## 建议顺序",
        "",
        "1. 先完成基础练习，熟悉 API。",
        "2. 再做进阶练习，把多个知识点串起来。",
        "3. 最后回答思考题，把“会写”升级成“会解释”。",
        "",
        "## 文件索引",
        "",
    ]

    for chapter_no, payload in EXERCISE_CONTENT.items():
        file_name = f"ch{chapter_no}_exercises.md"
        readme_lines.append(f"- `{file_name}`: {payload['title']}")

        lines = [
            f"# 第 {int(chapter_no)} 章练习：{payload['title']}",
            "",
            "## 学习目标",
            "",
        ]
        for goal in payload["goals"]:
            lines.append(f"- {goal}")

        lines.extend(["", "## 练习题", ""])
        for title, content in payload["tasks"]:
            lines.append(f"### {title}")
            lines.append("")
            lines.append(content)
            lines.append("")

        lines.extend(
            [
                "## 提交建议",
                "",
                "- 每道练习尽量保存运行结果图片或视频。",
                "- 修改参数时记录你的观察，例如阈值变大后发生了什么。",
                "- 如果你做不出来，先回看同章节的完整代码，再自己独立重写一遍。",
                "",
            ]
        )

        (EXERCISES_DIR / file_name).write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")

    (EXERCISES_DIR / "README.md").write_text("\n".join(readme_lines).rstrip() + "\n", encoding="utf-8")


def main() -> None:
    build_examples()
    build_exercises()
    print("Study pack generated.")
    print(f"Examples: {EXAMPLES_DIR.resolve()}")
    print(f"Exercises: {EXERCISES_DIR.resolve()}")


if __name__ == "__main__":
    main()
