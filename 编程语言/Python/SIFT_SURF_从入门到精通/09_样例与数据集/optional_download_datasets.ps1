param(
    [ValidateSet('oxford-info','hpatches-info')]
    [string]$Dataset = 'oxford-info'
)

$links = @{
    'oxford-info' = 'https://www.robots.ox.ac.uk/~vgg/research/affine/'
    'hpatches-info' = 'https://github.com/hpatches/hpatches-dataset'
}

Write-Host "请先阅读许可、数据规模和用途说明："
Write-Host $links[$Dataset]
Write-Host "本脚本不会自动下载大型数据集，避免未经确认占用空间或违反条款。"

