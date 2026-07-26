#include <iostream>
#include <string>
#include <vector>

#include <opencv2/calib3d.hpp>
#include <opencv2/features2d.hpp>
#include <opencv2/imgcodecs.hpp>

int main(int argc, char** argv) {
    if (argc != 3) {
        std::cerr << "Usage: sift_match <image-a> <image-b>\n";
        return 2;
    }

    cv::Mat imageA = cv::imread(argv[1], cv::IMREAD_GRAYSCALE);
    cv::Mat imageB = cv::imread(argv[2], cv::IMREAD_GRAYSCALE);
    if (imageA.empty() || imageB.empty()) {
        std::cerr << "Failed to read input images.\n";
        return 3;
    }

    auto sift = cv::SIFT::create();
    std::vector<cv::KeyPoint> keypointsA, keypointsB;
    cv::Mat descriptorsA, descriptorsB;
    sift->detectAndCompute(imageA, cv::noArray(), keypointsA, descriptorsA);
    sift->detectAndCompute(imageB, cv::noArray(), keypointsB, descriptorsB);

    cv::BFMatcher matcher(cv::NORM_L2);
    std::vector<std::vector<cv::DMatch>> pairs;
    matcher.knnMatch(descriptorsA, descriptorsB, pairs, 2);

    std::vector<cv::DMatch> good;
    for (const auto& pair : pairs) {
        if (pair.size() == 2 && pair[0].distance < 0.75F * pair[1].distance) {
            good.push_back(pair[0]);
        }
    }

    int inliers = 0;
    if (good.size() >= 4) {
        std::vector<cv::Point2f> source, destination;
        for (const auto& match : good) {
            source.push_back(keypointsA[match.queryIdx].pt);
            destination.push_back(keypointsB[match.trainIdx].pt);
        }
        cv::Mat mask;
        cv::findHomography(source, destination, cv::RANSAC, 5.0, mask);
        inliers = cv::countNonZero(mask);
    }

    std::cout << "keypoints_a=" << keypointsA.size()
              << " keypoints_b=" << keypointsB.size()
              << " good_matches=" << good.size()
              << " inliers=" << inliers << '\n';
    return 0;
}

