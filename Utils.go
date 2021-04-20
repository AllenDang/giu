package giu

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/AllenDang/imgui-go"
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

func Vec4ToRGBA(vec4 imgui.Vec4) color.RGBA {
	return color.RGBA{
		R: uint8(vec4.X * 255),
		G: uint8(vec4.Y * 255),
		B: uint8(vec4.Z * 255),
		A: uint8(vec4.W * 255),
	}
}

func Update() {
	if Context.isAlive {
		Context.platform.Update()
		Context.IO().SetFrameCountSinceLastInput(0)
	}
}

func GetCursorScreenPos() image.Point {
	pos := imgui.CursorScreenPos()
	return image.Pt(int(pos.X), int(pos.Y))
}

func GetCursorPos() image.Point {
	pos := imgui.CursorPos()
	return image.Pt(int(pos.X), int(pos.Y))
}

func GetMousePos() image.Point {
	pos := imgui.MousePos()
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

func SetNextWindowSize(width, height float32) {
	imgui.SetNextWindowSize(imgui.Vec2{X: width * Context.platform.GetContentScale(), Y: height * Context.platform.GetContentScale()})
}

type ExecCondition imgui.Condition

const (
	ConditionAlways       ExecCondition = ExecCondition(imgui.ConditionAlways)
	ConditionOnce         ExecCondition = ExecCondition(imgui.ConditionOnce)
	ConditionFirstUseEver ExecCondition = ExecCondition(imgui.ConditionFirstUseEver)
	ConditionAppearing    ExecCondition = ExecCondition(imgui.ConditionAppearing)
)

func SetNextWindowPos(x, y float32) {
	imgui.SetNextWindowPos(imgui.Vec2{X: x, Y: y})
}

func SetNextWindowSizeV(width, height float32, condition ExecCondition) {
	imgui.SetNextWindowSizeV(imgui.Vec2{X: width * Context.platform.GetContentScale(), Y: height * Context.platform.GetContentScale()}, imgui.Condition(condition))
}

func SetItemDefaultFocus() {
	imgui.SetItemDefaultFocus()
}

func SetKeyboardFocusHere() {
	SetKeyboardFocusHereV(0)
}

func SetKeyboardFocusHereV(i int) {
	imgui.SetKeyboardFocusHereV(i)
}

func PushClipRect(clipRectMin, clipRectMax image.Point, intersectWithClipRect bool) {
	imgui.PushClipRect(ToVec2(clipRectMin), ToVec2(clipRectMax), intersectWithClipRect)
}

func PopClipRect() {
	imgui.PopClipRect()
}
