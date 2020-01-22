package giu

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

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

func ToVec4Color(col color.RGBA) imgui.Vec4 {
	return imgui.Vec4{
		X: float32(col.R) / 255,
		Y: float32(col.G) / 255,
		Z: float32(col.B) / 255,
		W: float32(col.A) / 255,
	}
}

func ToVec2(pt image.Point) imgui.Vec2 {
	return imgui.Vec2{
		X: float32(pt.X),
		Y: float32(pt.Y),
	}
}

func Update() {
	Context.platform.Update()
}

func GetCursorScreenPos() image.Point {
	pos := imgui.CursorScreenPos()
	return image.Pt(int(pos.X), int(pos.Y))
}

func GetCursorPos() image.Point {
	pos := imgui.CursorPos()
	return image.Pt(int(pos.X), int(pos.Y))
}

func GetAvaiableRegion() (width, height float32) {
	region := imgui.ContentRegionAvail()
	return region.X, region.Y
}

func CalcTextSize(text string) (width, height float32) {
	size := imgui.CalcTextSize(text, false, 0)
	return size.X, size.Y
}
