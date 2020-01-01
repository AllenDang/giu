package giu

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	"time"

	"github.com/AllenDang/giu/imgui"
)

func LoadImage(imgPath string) (*image.RGBA, error) {
	imgFile, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	switch trueImg := img.(type) {
	case *image.RGBA:
		return trueImg, nil
	default:
		rgba := image.NewRGBA(trueImg.Bounds())
		draw.Draw(rgba, trueImg.Bounds(), trueImg, image.Pt(0, 0), draw.Src)
		return rgba, nil
	}
}

func Update(t time.Duration) {
	f := float32(t) / float32(time.Second)
	imgui.SetMaxWaitBeforeNextFrame(f)
}
