cmake_minimum_required(VERSION 3.20)
project(sift_match_demo LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(OpenCV 4.12 REQUIRED COMPONENTS core imgcodecs imgproc features2d calib3d)
add_executable(sift_match sift_match.cpp)
target_link_libraries(sift_match PRIVATE ${OpenCV_LIBS})
target_include_directories(sift_match PRIVATE ${OpenCV_INCLUDE_DIRS})

