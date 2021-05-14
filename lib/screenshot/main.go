package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"image"
	"image/png"
	"os"
	"path"
	"time"

	"github.com/gitsang/golog"
	"github.com/kbinani/screenshot"
)

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	dir := path.Dir(filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Error("mkdir failed", zap.Error(err))
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Error("create failed", zap.Error(err))
	}
	defer func() { _ = file.Close() }()

	encoder := png.Encoder{
		CompressionLevel: png.BestCompression,
		BufferPool:       nil,
	}
	err = encoder.Encode(file, img)
	if err != nil {
		log.Error("png encode failed", zap.Error(err))
	}

	log.Info("save image success", zap.String("path", filePath))
}

func genImgPath(savePath, prefix string) string {
	now := time.Now()
	year, month, day := now.Date()
	hour, min, sec := now.Clock()
	dir := fmt.Sprintf("screenshot-%04d%02d%02d", year, month, day)
	file := fmt.Sprintf("%s-%04d%02d%02d-%02d%02d%02d.png", prefix, year, month, day, hour, min, sec)
	return path.Join(savePath, dir, file)
}

func screenshotAll(savePath string) {
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
		prefix := fmt.Sprintf("screen%d", i)
		save(img, genImgPath(savePath, prefix))
	}
}

func main() {
	savePathPtr := flag.String("s", "", "save path")
	logPathPtr := flag.String("l", "screenshot.log", "log path")
	intervalPtr := flag.Int64("i", 10, "interval(minute)")
	flag.Parse()

	savePath := *savePathPtr
	logPath := *logPathPtr
	interval := time.Duration(*intervalPtr) * time.Minute
	log.InitLogger(log.WithLogFile(logPath))

	log.Info("config", zap.String("savePath", savePath),
		zap.String("logPath", logPath), zap.Any("interval", interval))
	for {
		screenshotAll(savePath)

		select {
		case <-time.Tick(interval):
		}
	}
}
