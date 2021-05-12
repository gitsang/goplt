package main

import (
	"fmt"
	"go.uber.org/zap"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/gitsang/golog"
	"github.com/kbinani/screenshot"
)

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Error("save failed", zap.Error(err))
	}
	defer func() { _ = file.Close() }()

	err = png.Encode(file, img)
	if err != nil {
		log.Error("png encode failed", zap.Error(err))
	}

	log.Info("save image success", zap.String("path", filePath))
}

func genImgName(prefix string) string {
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()
	return fmt.Sprintf("screenshot-%s-%04d%02d%02d-%02d%02d%02d.png", prefix, year, month, day, hour, min, sec)
}

func screenshotCustomize(x, y, w, h int) {
	img, err := screenshot.Capture(x, y, w, h)
	if err != nil {
		panic(err)
	}
	save(img, genImgName("customize"))
}

func screenshotAll() {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		log.Error("no active displays")
	}

	for i := 0; i < n; i++ {
		img, err := screenshot.CaptureDisplay(i)
		if err != nil {
			log.Error("screenshot failed", zap.Error(err))
			continue
		}
		save(img, genImgName(fmt.Sprintf("screen%d", i)))
	}
}

func screenshotRectangle(x0, y0, x1, y1 int)  {
	rect := image.Rect(x0, y0, x1, y1)
	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		panic(err)
	}
	save(img, genImgName("rectangle"))
}

func screenshotBounds(idx int) {
	bounds := screenshot.GetDisplayBounds(idx)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Error("screenshotBounds failed", zap.Error(err))
	}
	save(img, genImgName("bounds"))
}

func screenshotUnion() {
	all := image.Rect(0, 0, 0, 0)
	bounds := screenshot.GetDisplayBounds(0)
	all = bounds.Union(all)
	fmt.Println(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
}

func main() {
	log.InitLogger(log.WithLogFile("screenshot.log"))

	for {
		select {
		case <-time.Tick(60 * time.Second):
		}

		screenshotAll()
	}
}
