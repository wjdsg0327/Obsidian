```go
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"os"
	"sort"

	ort "github.com/yalue/onnxruntime_go"
	"golang.org/x/image/draw"
)

// 定义边界框结构体
type BoundingBox struct {
	Xmin, Ymin, Xmax, Ymax float32
	Score                  float32
}

// ---------------------------------------------------------------------------
// 1. 图像预处理 (修改版：返回缩放后的图片对象，方便后续画框)
// ---------------------------------------------------------------------------
func loadAndPreprocessImage(imagePath string) ([]float32, *image.RGBA, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, nil, err
	}

	bounds := image.Rect(0, 0, 640, 640)
	rgba := image.NewRGBA(bounds)
	draw.BiLinear.Scale(rgba, rgba.Bounds(), img, img.Bounds(), draw.Src, nil)

	channelSize := 640 * 640
	inputData := make([]float32, 3*channelSize)

	for y := 0; y < 640; y++ {
		for x := 0; x < 640; x++ {
			c := rgba.RGBAAt(x, y)
			idx := y*640 + x
			inputData[0*channelSize+idx] = float32(c.R) / 255.0
			inputData[1*channelSize+idx] = float32(c.G) / 255.0
			inputData[2*channelSize+idx] = float32(c.B) / 255.0
		}
	}

	return inputData, rgba, nil
}

// ---------------------------------------------------------------------------
// 2. 计算两个框的交并比 (IoU)，用于 NMS
// ---------------------------------------------------------------------------
func calculateIoU(b1, b2 BoundingBox) float32 {
	interXmin := max(b1.Xmin, b2.Xmin)
	interYmin := max(b1.Ymin, b2.Ymin)
	interXmax := min(b1.Xmax, b2.Xmax)
	interYmax := min(b1.Ymax, b2.Ymax)

	interW := max(0, interXmax-interXmin)
	interH := max(0, interYmax-interYmin)
	interArea := interW * interH

	b1Area := (b1.Xmax - b1.Xmin) * (b1.Ymax - b1.Ymin)
	b2Area := (b2.Xmax - b2.Xmin) * (b2.Ymax - b2.Ymin)

	return interArea / (b1Area + b2Area - interArea)
}

// ---------------------------------------------------------------------------
// 3. 在图片上画一个粗框
// ---------------------------------------------------------------------------
func drawRect(img *image.RGBA, x1, y1, x2, y2 int, col color.Color) {
	thickness := 3 // 框的粗细(像素)
	for t := 0; t < thickness; t++ {
		for x := x1 - t; x <= x2+t; x++ {
			if x >= 0 && x < img.Bounds().Dx() {
				img.Set(x, max(0, y1-t), col)
				img.Set(x, min(img.Bounds().Dy()-1, y2+t), col)
			}
		}
		for y := y1 - t; y <= y2+t; y++ {
			if y >= 0 && y < img.Bounds().Dy() {
				img.Set(max(0, x1-t), y, col)
				img.Set(min(img.Bounds().Dx()-1, x2+t), y, col)
			}
		}
	}
}

func main() {
	// 初始化环境 (请确认 DLL 路径)
	ort.SetSharedLibraryPath("D:\\subject\\GoWorks\\src\\aaa\\onnxruntime-win-x64-gpu_cuda13-1.25.0\\onnxruntime-win-x64-gpu-1.25.0\\lib\\onnxruntime.dll")
	ort.InitializeEnvironment()
	defer ort.DestroyEnvironment()

	// 1. 读取并预处理图片
	imagePath := "test.jpg" // 替换为你的真实图片
	fmt.Println("1. 正在读取并处理图片...")
	inputData, rgbaImage, err := loadAndPreprocessImage(imagePath)
	if err != nil {
		panic(err)
	}

	// 2. 准备推理 Tensor
	inputTensor, _ := ort.NewTensor(ort.NewShape(1, 3, 640, 640), inputData)
	defer inputTensor.Destroy()
	outputTensor, _ := ort.NewEmptyTensor[float32](ort.NewShape(1, 5, 8400))
	defer outputTensor.Destroy()

	// 3. 运行推理
	session, _ := ort.NewAdvancedSession(
		"bgi_mine.onnx",
		[]string{"images"}, []string{"output0"},
		[]ort.ArbitraryTensor{inputTensor}, []ort.ArbitraryTensor{outputTensor}, nil,
	)
	defer session.Destroy()

	fmt.Println("2. 正在执行 AI 推理...")
	session.Run()
	resultData := outputTensor.GetData()

	// ---------------------------------------------------------------------------
	// 核心逻辑：数据解析与 NMS
	// ---------------------------------------------------------------------------
	fmt.Println("3. 正在解析结果并执行 NMS (非极大值抑制)...")
	var candidateBoxes []BoundingBox

	// 解析 8400 个框 (按照 cx, cy, w, h, score 排布)
	for i := 0; i < 8400; i++ {
		score := resultData[33600+i] // 得分起始索引是 4 * 8400 = 33600

		// 置信度过滤：只保留得分大于 0.4 的框
		if score > 0.4 {
			cx := resultData[0+i]       // 中心 X
			cy := resultData[8400+i]    // 中心 Y
			w := resultData[16800+i]    // 宽度
			h := resultData[25200+i]    // 高度

			// 将中心坐标和宽高，转换为左上角和右下角的坐标
			candidateBoxes = append(candidateBoxes, BoundingBox{
				Xmin:  cx - w/2.0,
				Ymin:  cy - h/2.0,
				Xmax:  cx + w/2.0,
				Ymax:  cy + h/2.0,
				Score: score,
			})
		}
	}

	// 执行 NMS (非极大值抑制)
	// 1) 按照得分从高到低排序
	sort.Slice(candidateBoxes, func(i, j int) bool {
		return candidateBoxes[i].Score > candidateBoxes[j].Score
	})

	// 2) 过滤重叠框 (IoU 阈值设为 0.45)
	var finalBoxes []BoundingBox
	nmsThreshold := float32(0.45)

	for _, box := range candidateBoxes {
		keep := true
		for _, fBox := range finalBoxes {
			if calculateIoU(box, fBox) > nmsThreshold {
				keep = false // 和已保留的高分框重叠度太高，丢弃
				break
			}
		}
		if keep {
			finalBoxes = append(finalBoxes, box)
		}
	}

	fmt.Printf("过滤完成！最终保留了 %d 个目标。\n", len(finalBoxes))

	// ---------------------------------------------------------------------------
	// 绘图并保存
	// ---------------------------------------------------------------------------
	fmt.Println("4. 正在将框绘制到图片上...")
	redColor := color.RGBA{R: 255, G: 0, B: 0, A: 255} // 画笔颜色

	for i, box := range finalBoxes {
		fmt.Printf("  目标 %d: 坐标(%.1f, %.1f) 到 (%.1f, %.1f), 置信度: %.2f\n",
			i+1, box.Xmin, box.Ymin, box.Xmax, box.Ymax, box.Score)

		drawRect(rgbaImage, int(box.Xmin), int(box.Ymin), int(box.Xmax), int(box.Ymax), redColor)
	}

	// 保存图片
	outFile, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, rgbaImage, &jpeg.Options{Quality: 95})

	fmt.Println("\n🎉 大功告成！请查看根目录下的 output.jpg。")
}
```