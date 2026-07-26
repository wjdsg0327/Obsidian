# C++ 附录

该示例演示 SIFT、Lowe 比率测试和单应性内点统计。它默认使用系统或 `OpenCV_DIR` 指定的 OpenCV，不会修改 Python 环境。

```powershell
cmake -S cpp -B cpp/build -D OpenCV_DIR=E:\opencv-surf-build\install\x64\vc17\lib
cmake --build cpp/build --config Release
cpp\build\Release\sift_match.exe 09_样例与数据集\opencv_samples\box.png 09_样例与数据集\opencv_samples\box_in_scene.png
```

若只使用系统已安装的 OpenCV，可省略 `OpenCV_DIR`。需要在 C++ 中运行 SURF 时，还要包含 `opencv2/xfeatures2d.hpp`，并使用启用了 nonfree 的独立构建。

