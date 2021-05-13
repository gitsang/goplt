package main

import (
	"fmt"
	"image"

	log "github.com/gitsang/golog"
	"github.com/kbinani/screenshot"
	"go.uber.org/zap"
)

func screenshotCustomize(x, y, w, h int) {
	img, err := screenshot.Capture(x, y, w, h)
	if err != nil {
		panic(err)
	}
	save(img, genImgPath("customize"))
}

func screenshotRectangle(x0, y0, x1, y1 int) {
	rect := image.Rect(x0, y0, x1, y1)
	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		panic(err)
	}
	save(img, genImgPath("rectangle"))
}

func screenshotBounds(idx int) {
	bounds := screenshot.GetDisplayBounds(idx)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Error("screenshotBounds failed", zap.Error(err))
	}
	save(img, genImgPath("bounds"))
}

func screenshotUnion() {
	all := image.Rect(0, 0, 0, 0)
	bounds := screenshot.GetDisplayBounds(0)
	all = bounds.Union(all)
	fmt.Println(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
}
