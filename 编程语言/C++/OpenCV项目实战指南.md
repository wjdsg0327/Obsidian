# OpenCV C++ 项目实战指南

## 1. 人脸检测 (Face Detection)

### 1.1 Haar Cascade 方法 (传统方法)

**原理：** 使用预训练的 Haar 特征分类器，通过滑动窗口检测人脸。

**核心代码：**
```cpp
#include <opencv2/opencv.hpp>
#include <opencv2/objdetect.hpp>

int main() {
    // 加载 Haar Cascade 分类器
    cv::CascadeClassifier faceCascade;
    if (!faceCascade.load("haarcascade_frontalface_default.xml")) {
        std::cerr << "Error loading cascade file!" << std::endl;
        return -1;
    }

    // 打开摄像头
    cv::VideoCapture cap(0);
    if (!cap.isOpened()) {
        std::cerr << "Error opening camera!" << std::endl;
        return -1;
    }

    cv::Mat frame, gray;
    std::vector<cv::Rect> faces;

    while (true) {
        cap >> frame;
        if (frame.empty()) break;

        // 转换为灰度图
        cv::cvtColor(frame, gray, cv::COLOR_BGR2GRAY);

        // 直方图均衡化（增强对比度）
        cv::equalizeHist(gray, gray);

        // 检测人脸
        faceCascade.detectMultiScale(
            gray,           // 输入图像
            faces,          // 输出人脸矩形
            1.1,            // scaleFactor - 图像缩放比例
            3,              // minNeighbors - 最小邻居数
            0,              // flags
            cv::Size(30, 30) // minSize - 最小人脸尺寸
        );

        // 绘制人脸矩形
        for (const auto& face : faces) {
            cv::rectangle(frame, face, cv::Scalar(0, 255, 0), 2);
            // 可选：在人脸上方显示文字
            cv::putText(frame, "Face", 
                       cv::Point(face.x, face.y - 10),
                       cv::FONT_HERSHEY_SIMPLEX, 0.9, 
                       cv::Scalar(0, 255, 0), 2);
        }

        cv::imshow("Face Detection", frame);
        if (cv::waitKey(30) >= 0) break;
    }

    return 0;
}
```

**关键参数说明：**
- `scaleFactor`: 图像缩放比例，越小检测越细致（但更慢）
- `minNeighbors`: 值越高误检越少，但可能漏检
- `minSize`: 忽略小于此尺寸的人脸

**Haar Cascade XML 文件位置：**
- OpenCV 自带：`data/haarcascades/`
- 常用文件：
  - `haarcascade_frontalface_default.xml` - 正面人脸
  - `haarcascade_frontalface_alt.xml` - 正面人脸（备选）
  - `haarcascade_eye.xml` - 眼睛检测

**官方教程：** https://docs.opencv.org/db/d28/tutorial_cascade_classifier.html

---

### 1.2 DNN 方法 (深度学习方法)

**原理：** 使用预训练的深度神经网络（如 SSD、ResNet），精度更高。

**核心代码：**
```cpp
#include <opencv2/opencv.hpp>
#include <opencv2/dnn.hpp>

int main() {
    // 加载 DNN 模型
    cv::dnn::Net net = cv::dnn::readNetFromCaffe(
        "deploy.prototxt",                    // 模型定义文件
        "res10_300x300_ssd_iter_140000.caffemodel"  // 预训练权重
    );

    // 设置首选后端和目标（可选，使用 GPU 加速）
    // net.setPreferableBackend(cv::dnn::DNN_BACKEND_CUDA);
    // net.setPreferableTarget(cv::dnn::DNN_TARGET_CUDA);

    cv::VideoCapture cap(0);
    cv::Mat frame;

    while (true) {
        cap >> frame;
        if (frame.empty()) break;

        // 创建 blob（预处理）
        cv::Mat blob = cv::dnn::blobFromImage(
            frame,          // 输入图像
            1.0,            // 缩放因子
            cv::Size(300, 300),  // 目标尺寸
            cv::Scalar(104, 177, 123),  // 均值减法（BGR）
            false,          // swapRB
            false           // crop
        );

        // 设置输入并前向传播
        net.setInput(blob);
        cv::Mat detections = net.forward();

        // 解析检测结果
        cv::Mat detectionMat(detections.size[2], detections.size[3], 
                            CV_32F, detections.ptr<float>());

        float confidenceThreshold = 0.5;  // 置信度阈值

        for (int i = 0; i < detectionMat.rows; i++) {
            float confidence = detectionMat.at<float>(i, 2);

            if (confidence > confidenceThreshold) {
                // 获取边界框坐标
                int x1 = static_cast<int>(detectionMat.at<float>(i, 3) * frame.cols);
                int y1 = static_cast<int>(detectionMat.at<float>(i, 4) * frame.rows);
                int x2 = static_cast<int>(detectionMat.at<float>(i, 5) * frame.cols);
                int y2 = static_cast<int>(detectionMat.at<float>(i, 6) * frame.rows);

                // 绘制边界框
                cv::rectangle(frame, cv::Point(x1, y1), cv::Point(x2, y2), 
                            cv::Scalar(0, 255, 0), 2);

                // 显示置信度
                std::string label = cv::format("Face: %.2f", confidence);
                cv::putText(frame, label, cv::Point(x1, y1 - 10),
                           cv::FONT_HERSHEY_SIMPLEX, 0.7, 
                           cv::Scalar(0, 255, 0), 2);
            }
        }

        cv::imshow("DNN Face Detection", frame);
        if (cv::waitKey(30) >= 0) break;
    }

    return 0;
}
```

**模型文件下载：**
- deploy.prototxt: https://github.com/opencv/opencv/blob/master/samples/dnn/face_detector/deploy.prototxt
- caffemodel: https://github.com/opencv/opencv_3rdparty/raw/dnn_samples_face_detector_20170830/res10_300x300_ssd_iter_140000.caffemodel

**Haar vs DNN 对比：**

| 特性 | Haar Cascade | DNN (SSD) |
|------|-------------|-----------|
| 速度 | 快 | 较慢 |
| 精度 | 中等 | 高 |
| 误检 | 较多 | 较少 |
| 角度变化 | 差 | 好 |
| 遮挡处理 | 差 | 好 |
| GPU 加速 | 不支持 | 支持 |

---

### 1.3 GitHub 仓库

- **opencv/opencv** - 官方示例：`samples/cpp/tutorial_code/objectDetection/`
- **opencv/opencv_contrib** - 包含额外的人脸识别模块（face module）

---

## 2. 车道检测 (Lane Detection)

### 2.1 传统方法（Canny + Hough Transform）

**处理流程：**
1. 灰度转换
2. 高斯模糊（降噪）
3. Canny 边缘检测
4. ROI 区域掩码
5. Hough 直线检测
6. 直线平均/外推
7. 叠加显示

**核心代码：**
```cpp
#include <opencv2/opencv.hpp>
#include <vector>
#include <cmath>

// 计算直线斜率
double slope(cv::Vec4i line) {
    return static_cast<double>(line[3] - line[1]) / 
           static_cast<double>(line[2] - line[0]);
}

// 绘制车道线
void drawLaneLines(cv::Mat& frame, 
                   const std::vector<cv::Vec4i>& leftLines,
                   const std::vector<cv::Vec4i>& rightLines) {
    
    // 计算左车道线平均
    double leftSlopeSum = 0, leftInterceptSum = 0;
    for (const auto& line : leftLines) {
        double s = slope(line);
        leftSlopeSum += s;
        leftInterceptSum += line[1] - s * line[0];  // y = mx + b -> b = y - mx
    }
    double avgLeftSlope = leftSlopeSum / leftLines.size();
    double avgLeftIntercept = leftInterceptSum / leftLines.size();
    
    // 计算右车道线平均
    double rightSlopeSum = 0, rightInterceptSum = 0;
    for (const auto& line : rightLines) {
        double s = slope(line);
        rightSlopeSum += s;
        rightInterceptSum += line[1] - s * line[0];
    }
    double avgRightSlope = rightSlopeSum / rightLines.size();
    double avgRightIntercept = rightInterceptSum / rightLines.size();
    
    // 计算线段端点（从底部到图像中部）
    int yBottom = frame.rows;
    int yMiddle = static_cast<int>(frame.rows * 0.6);
    
    // 左车道线
    int xLeftBottom = static_cast<int>((yBottom - avgLeftIntercept) / avgLeftSlope);
    int xLeftMiddle = static_cast<int>((yMiddle - avgLeftIntercept) / avgLeftSlope);
    cv::line(frame, cv::Point(xLeftBottom, yBottom), 
             cv::Point(xLeftMiddle, yMiddle), 
             cv::Scalar(0, 255, 0), 5);
    
    // 右车道线
    int xRightBottom = static_cast<int>((yBottom - avgRightIntercept) / avgRightSlope);
    int xRightMiddle = static_cast<int>((yMiddle - avgRightIntercept) / avgRightSlope);
    cv::line(frame, cv::Point(xRightBottom, yBottom), 
             cv::Point(xRightMiddle, yMiddle), 
             cv::Scalar(0, 255, 0), 5);
}

int main() {
    cv::VideoCapture cap("road_video.mp4");
    cv::Mat frame;

    while (true) {
        cap >> frame;
        if (frame.empty()) break;

        // 1. 灰度转换
        cv::Mat gray;
        cv::cvtColor(frame, gray, cv::COLOR_BGR2GRAY);

        // 2. 高斯模糊
        cv::Mat blurred;
        cv::GaussianBlur(gray, blurred, cv::Size(5, 5), 0);

        // 3. Canny 边缘检测
        cv::Mat edges;
        cv::Canny(blurred, edges, 50, 150);

        // 4. ROI 区域掩码（梯形区域）
        cv::Mat mask = cv::Mat::zeros(edges.size(), edges.type());
        std::vector<cv::Point> vertices = {
            cv::Point(0, frame.rows),                          // 左下
            cv::Point(frame.cols * 0.45, frame.rows * 0.6),   // 左上
            cv::Point(frame.cols * 0.55, frame.rows * 0.6),   // 右上
            cv::Point(frame.cols, frame.rows)                   // 右下
        };
        cv::fillPoly(mask, std::vector<std::vector<cv::Point>>{vertices}, 
                     cv::Scalar(255));
        
        cv::Mat maskedEdges;
        cv::bitwise_and(edges, mask, maskedEdges);

        // 5. Hough 直线检测
        std::vector<cv::Vec4i> lines;
        cv::HoughLinesP(maskedEdges, lines, 2, CV_PI/180, 
                        100,    // 阈值
                        40,     // 最小线段长度
                        5);     // 最大线段间隙

        // 6. 分离左右车道线
        std::vector<cv::Vec4i> leftLines, rightLines;
        for (const auto& line : lines) {
            double s = slope(line);
            if (std::abs(s) < 0.5) continue;  // 过滤水平线
            
            if (s < 0) {
                leftLines.push_back(line);   // 负斜率 = 左车道
            } else {
                rightLines.push_back(line);  // 正斜率 = 右车道
            }
        }

        // 7. 绘制车道线
        cv::Mat laneFrame = frame.clone();
        if (!leftLines.empty() && !rightLines.empty()) {
            drawLaneLines(laneFrame, leftLines, rightLines);
        }

        // 8. 叠加显示
        cv::Mat result;
        cv::addWeighted(frame, 0.8, laneFrame, 0.3, 0, result);

        cv::imshow("Lane Detection", result);
        if (cv::waitKey(30) >= 0) break;
    }

    return 0;
}
```

**关键概念：**

**Canny 边缘检测：**
- 使用梯度计算边缘
- 两个阈值：低阈值（弱边缘）和高阈值（强边缘）
- 非极大值抑制 + 滞后阈值

**Hough Transform：**
- 将图像空间的点映射到参数空间 (ρ, θ)
- `HoughLinesP` 返回线段的起止点（更实用）
- 参数：阈值、最小线长、最大间隙

**ROI 掩码：**
- 只关注图像底部的梯形区域
- 避免检测到树木、建筑物等干扰

**GitHub 仓库：**
- UjjwalSaxena/Automated-Lane-Detection
- cardboardcode/opencv-lane-detection
- 搜索 "lane detection opencv c++" 可找到 50+ 仓库

**官方教程：** https://docs.opencv.org/da/d22/tutorial_py_canny.html

---

## 3. 对象跟踪 (Object Tracking)

### 3.1 颜色跟踪（HSV 颜色空间）

**原理：** 在 HSV 颜色空间中，颜色更容易分离，对光照变化更鲁棒。

**核心代码：**
```cpp
#include <opencv2/opencv.hpp>

int main() {
    cv::VideoCapture cap(0);
    cv::Mat frame, hsv, mask;

    // 定义 HSV 颜色范围（以蓝色为例）
    cv::Scalar lowerBlue(100, 50, 50);
    cv::Scalar upperBlue(130, 255, 255);

    while (true) {
        cap >> frame;
        if (frame.empty()) break;

        // 1. 转换为 HSV
        cv::cvtColor(frame, hsv, cv::COLOR_BGR2HSV);

        // 2. 颜色阈值化
        cv::inRange(hsv, lowerBlue, upperBlue, mask);

        // 3. 形态学操作（去除噪声）
        cv::Mat kernel = cv::getStructuringElement(cv::MORPH_ELLIPSE, cv::Size(5, 5));
        cv::morphologyEx(mask, mask, cv::MORPH_OPEN, kernel);   // 开运算
        cv::morphologyEx(mask, mask, cv::MORPH_CLOSE, kernel);  // 闭运算

        // 4. 查找轮廓
        std::vector<std::vector<cv::Point>> contours;
        cv::findContours(mask, contours, cv::RETR_EXTERNAL, 
                        cv::CHAIN_APPROX_SIMPLE);

        // 5. 找到最大轮廓
        if (!contours.empty()) {
            auto maxContour = std::max_element(contours.begin(), contours.end(),
                [](const auto& a, const auto& b) {
                    return cv::contourArea(a) < cv::contourArea(b);
                });

            // 6. 计算最小外接圆
            cv::Point2f center;
            float radius;
            cv::minEnclosingCircle(*maxContour, center, radius);

            // 7. 绘制跟踪结果
            if (radius > 10) {  // 过滤噪声
                cv::circle(frame, center, static_cast<int>(radius), 
                          cv::Scalar(0, 255, 0), 2);
                cv::circle(frame, center, 5, cv::Scalar(0, 0, 255), -1);
                
                // 显示坐标
                std::string coord = cv::format("(%d, %d)", 
                                              static_cast<int>(center.x), 
                                              static_cast<int>(center.y));
                cv::putText(frame, coord, cv::Point(center.x + 10, center.y),
                           cv::FONT_HERSHEY_SIMPLEX, 0.5, 
                           cv::Scalar(0, 255, 0), 2);
            }
        }

        cv::imshow("Color Tracking", frame);
        cv::imshow("Mask", mask);
        
        if (cv::waitKey(30) >= 0) break;
    }

    return 0;
}
```

**常见颜色 HSV 范围：**

| 颜色 | H 范围 | S 范围 | V 范围 |
|------|--------|--------|--------|
| 红色 | 0-10, 170-180 | 50-255 | 50-255 |
| 橙色 | 10-25 | 50-255 | 50-255 |
| 黄色 | 25-35 | 50-255 | 50-255 |
| 绿色 | 35-85 | 50-255 | 50-255 |
| 蓝色 | 100-130 | 50-255 | 50-255 |
| 紫色 | 130-170 | 50-255 | 50-255 |

---

### 3.2 运动检测（帧差法）

**原理：** 比较连续帧的差异，检测运动物体。

**核心代码：**
```cpp
#include <opencv2/opencv.hpp>

int main() {
    cv::VideoCapture cap(0);
    cv::Mat frame, prevFrame, gray, prevGray, diff, thresh;

    cap >> frame;
    cv::cvtColor(frame, prevGray, cv::COLOR_BGR2GRAY);

    while (true) {
        cap >> frame;
        if (frame.empty()) break;

        // 1. 灰度转换
        cv::cvtColor(frame, gray, cv::COLOR_BGR2GRAY);

        // 2. 帧差法
        cv::absdiff(prevGray, gray, diff);

        // 3. 二值化
        cv::threshold(diff, thresh, 30, 255, cv::THRESH_BINARY);

        // 4. 形态学操作（去除噪声）
        cv::Mat kernel = cv::getStructuringElement(cv::MORPH_ELLIPSE, cv::Size(5, 5));
        cv::morphologyEx(thresh, thresh, cv::MORPH_OPEN, kernel);
        cv::morphologyEx(thresh, thresh, cv::MORPH_CLOSE, kernel);

        // 5. 查找轮廓
        std::vector<std::vector<cv::Point>> contours;
        cv::findContours(thresh, contours, cv::RETR_EXTERNAL, 
                        cv::CHAIN_APPROX_SIMPLE);

        // 6. 绘制边界框
        for (const auto& contour : contours) {
            double area = cv::contourArea(contour);
            if (area > 500) {  // 过滤小区域
                cv::Rect bbox = cv::boundingRect(contour);
                cv::rectangle(frame, bbox, cv::Scalar(0, 255, 0), 2);
            }
        }

        // 更新前一帧
        prevGray = gray.clone();

        cv::imshow("Motion Detection", frame);
        cv::imshow("Difference", thresh);
        
        if (cv::waitKey(30) >= 0) break;
    }

    return 0;
}
```

---

### 3.3 光流法跟踪

**原理：** 计算像素在连续帧间的运动矢量。

**核心代码（稀疏光流 - Lucas-Kanade）：**
```cpp
#include <opencv2/opencv.hpp>
#include <vector>

int main() {
    cv::VideoCapture cap(0);
    cv::Mat frame, gray, prevGray;

    // 特征点参数
    std::vector<cv::Point2f> points, prevPoints;
    std::vector<uchar> status;
    std::vector<float> err;

    // Shi-Tomasi 角点检测参数
    cv::Size winSize(15, 15);
    cv::TermCriteria criteria(cv::TermCriteria::COUNT + cv::TermCriteria::EPS, 
                             30, 0.01);

    bool initialized = false;

    while (true) {
        cap >> frame;
        if (frame.empty()) break;

        cv::cvtColor(frame, gray, cv::COLOR_BGR2GRAY);

        if (!initialized) {
            // 初始化：检测特征点
            cv::goodFeaturesToTrack(gray, points, 100, 0.3, 7);
            prevGray = gray.clone();
            prevPoints = points;
            initialized = true;
            continue;
        }

        // 计算光流
        cv::calcOpticalFlowPyrLK(prevGray, gray, prevPoints, points, 
                                 status, err, winSize, 3, criteria);

        // 绘制跟踪结果
        std::vector<cv::Point2f> goodPoints;
        for (size_t i = 0; i < points.size(); i++) {
            if (status[i]) {  // 跟踪成功
                // 绘制运动矢量
                cv::arrowedLine(frame, prevPoints[i], points[i], 
                               cv::Scalar(0, 255, 0), 2);
                // 绘制当前点
                cv::circle(frame, points[i], 5, cv::Scalar(0, 0, 255), -1);
                goodPoints.push_back(points[i]);
            }
        }

        // 更新
        prevGray = gray.clone();
        prevPoints = goodPoints;

        // 定期重新检测特征点
        if (prevPoints.size() < 20) {
            cv::goodFeaturesToTrack(gray, prevPoints, 100, 0.3, 7);
        }

        cv::imshow("Optical Flow", frame);
        if (cv::waitKey(30) >= 0) break;
    }

    return 0;
}
```

**GitHub 仓库：**
- 搜索 "object tracking opencv c++"
- LearnOpenCV.com 有详细教程

---

## 4. 图像处理基础

### 4.1 边缘检测

```cpp
// Canny 边缘检测
cv::Mat edges;
cv::Canny(image, edges, 50, 150);

// Sobel 算子
cv::Mat gradX, gradY;
cv::Sobel(image, gradX, CV_16S, 1, 0);
cv::Sobel(image, gradY, CV_16S, 0, 1);
cv::convertScaleAbs(gradX, gradX);
cv::convertScaleAbs(gradY, gradY);
cv::addWeighted(gradX, 0.5, gradY, 0.5, 0, edges);

// Laplacian 算子
cv::Mat laplacian;
cv::Laplacian(image, laplacian, CV_16S);
cv::convertScaleAbs(laplacian, laplacian);
```

### 4.2 轮廓检测

```cpp
// 查找轮廓
std::vector<std::vector<cv::Point>> contours;
std::vector<cv::Vec4i> hierarchy;
cv::findContours(binary, contours, hierarchy, 
                 cv::RETR_TREE, cv::CHAIN_APPROX_SIMPLE);

// 绘制轮廓
cv::Mat contourImage = cv::Mat::zeros(image.size(), CV_8UC3);
for (size_t i = 0; i < contours.size(); i++) {
    cv::drawContours(contourImage, contours, static_cast<int>(i), 
                     cv::Scalar(0, 255, 0), 2);
}

// 计算轮廓属性
for (const auto& contour : contours) {
    double area = cv::contourArea(contour);
    double perimeter = cv::arcLength(contour, true);
    cv::Moments mu = cv::moments(contour);
    cv::Point center(mu.m10 / mu.m00, mu.m01 / mu.m00);
    
    // 最小外接矩形
    cv::Rect bbox = cv::boundingRect(contour);
    cv::RotatedRect rotRect = cv::minAreaRect(contour);
}
```

### 4.3 形态学操作

```cpp
cv::Mat kernel = cv::getStructuringElement(cv::MORPH_RECT, cv::Size(5, 5));

// 腐蚀（去除小的白色噪声）
cv::erode(binary, binary, kernel);

// 膨胀（填补小的黑色空洞）
cv::dilate(binary, binary, kernel);

// 开运算（先腐蚀后膨胀 - 去除噪声）
cv::morphologyEx(binary, binary, cv::MORPH_OPEN, kernel);

// 闭运算（先膨胀后腐蚀 - 填补空洞）
cv::morphologyEx(binary, binary, cv::MORPH_CLOSE, kernel);

// 形态学梯度（膨胀 - 腐蚀 = 边缘）
cv::morphologyEx(binary, binary, cv::MORPH_GRADIENT, kernel);
```

---

## 5. 构建配置

### CMakeLists.txt

```cmake
cmake_minimum_required(VERSION 3.10)
project(OpenCVProjects)

set(CMAKE_CXX_STANDARD 17)

# 查找 OpenCV
find_package(OpenCV REQUIRED)
include_directories(${OpenCV_INCLUDE_DIRS})

# 人脸检测
add_executable(face_detection src/face_detection.cpp)
target_link_libraries(face_detection ${OpenCV_LIBS})

# 车道检测
add_executable(lane_detection src/lane_detection.cpp)
target_link_libraries(lane_detection ${OpenCV_LIBS})

# 颜色跟踪
add_executable(color_tracking src/color_tracking.cpp)
target_link_libraries(color_tracking ${OpenCV_LIBS})

# 运动检测
add_executable(motion_detection src/motion_detection.cpp)
target_link_libraries(motion_detection ${OpenCV_LIBS})

# 光流跟踪
add_executable(optical_flow src/optical_flow.cpp)
target_link_libraries(optical_flow ${OpenCV_LIBS})
```

### 安装 OpenCV

**Ubuntu/Debian：**
```bash
sudo apt update
sudo apt install libopencv-dev
```

**从源码编译（获取最新版本）：**
```bash
# 安装依赖
sudo apt install build-essential cmake git libgtk2.0-dev pkg-config \
    libavcodec-dev libavformat-dev libswscale-dev libtbb2 libtbb-dev \
    libjpeg-dev libpng-dev libtiff-dev libdc1394-dev

# 克隆 OpenCV 和 contrib
git clone https://github.com/opencv/opencv.git
git clone https://github.com/opencv/opencv_contrib.git

# 编译
cd opencv
mkdir build && cd build
cmake -D CMAKE_BUILD_TYPE=Release \
      -D CMAKE_INSTALL_PREFIX=/usr/local \
      -D OPENCV_EXTRA_MODULES_PATH=../../opencv_contrib/modules \
      -D BUILD_EXAMPLES=ON ..
make -j$(nproc)
sudo make install
sudo ldconfig
```

**验证安装：**
```bash
pkg-config --modversion opencv4
```

---

## 6. 学习资源

### 官方文档
- **OpenCV Tutorials**: https://docs.opencv.org/4.x/d9/df8/tutorial_root.html
- **C++ Samples**: https://github.com/opencv/opencv/tree/master/samples/cpp

### 教程网站
- **LearnOpenCV**: https://learnopencv.com/
- **GeeksforGeeks OpenCV**: https://www.geeksforgeeks.org/opencv-c-tutorial/

### GitHub 项目
- **opencv/opencv**: 官方仓库，包含示例代码
- **搜索 "opencv-cpp"**: https://github.com/topics/opencv-cpp

### 书籍
- **"Learning OpenCV 4"** - Packt Publishing
- **"OpenCV 4 Computer Vision Application Programming"** - Packt Publishing

---

## 相关笔记
- [[C++学习资源指南]]
- [[Python/opencv/docs/README|Python OpenCV 教程]]

## 参考
- 整理时间：2026-06-03
- 来源：OpenCV 官方文档、LearnOpenCV、GitHub 社区
