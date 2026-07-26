$ErrorActionPreference = 'Stop'
$root = Split-Path -Parent $PSScriptRoot
$paperDir = Join-Path $PSScriptRoot '经典论文'
$docsDir = Join-Path $PSScriptRoot 'OpenCV官方文档'
$sampleDir = Join-Path $root '09_样例与数据集\opencv_samples'

$downloads = @(
    @('https://www.cs.ubc.ca/~lowe/papers/ijcv04.pdf', (Join-Path $paperDir 'Lowe_2004_SIFT.pdf')),
    @('https://people.ee.ethz.ch/~surf/eccv06.pdf', (Join-Path $paperDir 'Bay_2006_SURF.pdf')),
    @('https://arxiv.org/pdf/1712.07629', (Join-Path $paperDir 'DeTone_2018_SuperPoint.pdf')),
    @('https://docs.opencv.org/4.x/da/df5/tutorial_py_sift_intro.html', (Join-Path $docsDir 'SIFT教程.html')),
    @('https://docs.opencv.org/4.x/df/dd2/tutorial_py_surf_intro.html', (Join-Path $docsDir 'SURF教程.html')),
    @('https://docs.opencv.org/4.x/dc/dc3/tutorial_py_matcher.html', (Join-Path $docsDir '特征匹配.html')),
    @('https://docs.opencv.org/4.x/d1/de0/tutorial_py_feature_homography.html', (Join-Path $docsDir '单应性目标定位.html')),
    @('https://raw.githubusercontent.com/opencv/opencv/4.x/samples/data/box.png', (Join-Path $sampleDir 'box.png')),
    @('https://raw.githubusercontent.com/opencv/opencv/4.x/samples/data/box_in_scene.png', (Join-Path $sampleDir 'box_in_scene.png')),
    @('https://raw.githubusercontent.com/opencv/opencv/4.x/samples/data/graf1.png', (Join-Path $sampleDir 'graf1.png')),
    @('https://raw.githubusercontent.com/opencv/opencv/4.x/samples/data/graf3.png', (Join-Path $sampleDir 'graf3.png'))
)

foreach ($item in $downloads) {
    New-Item -ItemType Directory -Force -Path (Split-Path -Parent $item[1]) | Out-Null
    Write-Host "下载 $($item[0])"
    curl.exe -L --fail --retry 2 --output $item[1] $item[0]
}

Write-Host '开放资料下载完成。请重新运行 scripts\validate_package.py 和资源索引生成程序以刷新校验值。'

