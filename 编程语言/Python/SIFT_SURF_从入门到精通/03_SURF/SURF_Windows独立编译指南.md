# SURF：Windows 独立编译指南

目标：在**独立目录/独立环境**编译 OpenCV 4.12.0 与 opencv_contrib 4.12.0，并启用 `OPENCV_ENABLE_NONFREE=ON`。不覆盖现有 Anaconda 或系统 OpenCV。

## 前置条件

- Visual Studio 2022 Build Tools，安装“使用 C++ 的桌面开发”。
- CMake 3.24 或更新版本。
- Git。
- 足够的磁盘空间（建议预留 15–25 GB）。
- 已确认相关算法在你的司法辖区和用途下可使用。

## 推荐目录

```text
E:\opencv-surf-build\
  opencv\
  opencv_contrib\
  build\
  install\
```

## 获取同版本源码

```powershell
git clone --branch 4.12.0 --depth 1 https://github.com/opencv/opencv.git E:\opencv-surf-build\opencv
git clone --branch 4.12.0 --depth 1 https://github.com/opencv/opencv_contrib.git E:\opencv-surf-build\opencv_contrib
```

## CMake 配置

在“x64 Native Tools Command Prompt for VS 2022”中执行：

```powershell
cmake -S E:\opencv-surf-build\opencv -B E:\opencv-surf-build\build `
  -G "Visual Studio 17 2022" -A x64 `
  -D CMAKE_INSTALL_PREFIX=E:\opencv-surf-build\install `
  -D OPENCV_EXTRA_MODULES_PATH=E:\opencv-surf-build\opencv_contrib\modules `
  -D OPENCV_ENABLE_NONFREE=ON `
  -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D BUILD_EXAMPLES=OFF
```

检查 CMake 输出必须包含：

```text
Non-free algorithms: YES
```

## 编译和安装

```powershell
cmake --build E:\opencv-surf-build\build --config Release --target INSTALL --parallel
```

## 验证 C++

使用本包 `cpp` 示例时，把 `OpenCV_DIR` 指向：

```text
E:\opencv-surf-build\install\x64\vc17\lib
```

## Python 说明

要得到带 SURF 的 Python 模块，还需要在 CMake 中让 `PYTHON3_EXECUTABLE`、`PYTHON3_INCLUDE_DIR`、`PYTHON3_LIBRARY` 和 `PYTHON3_PACKAGES_PATH` 指向一个**单独创建的 Python 环境**。不同 Python/NumPy 版本的 ABI 配置容易出错，因此本学习包不自动覆盖当前 `cv2`。

## 常见失败

- `SURF_create` 不存在：可能只装了官方 wheel，或没有编译 `opencv_contrib`。
- 抛出 patented/nonfree 错误：配置时没有真正启用 `OPENCV_ENABLE_NONFREE=ON`。
- `opencv2/xfeatures2d.hpp` 找不到：`OPENCV_EXTRA_MODULES_PATH` 指错，应指向 `opencv_contrib\modules`。
- Debug/Release DLL 不匹配：应用和 OpenCV 使用相同配置。
- Python 导入了旧 cv2：打印 `cv2.__file__`，确认没有加载 Anaconda 原包。

