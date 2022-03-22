package giu

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/AllenDang/imgui-go"
	"github.com/pkg/browser"
)

// LoadImage loads image from file and returns *image.RGBA.
func LoadImage(imgPath string) (*image.RGBA, error) {
	imgFile, err := os.Open(filepath.Clean(imgPath))
	if err != nil {
		return nil, fmt.Errorf("LoadImage: error opening image file %s: %w", imgPath, err)
	}

	defer func() {
		// nolint:govet // we want to reuse this err variable here
		if err := imgFile.Close(); err != nil {
			panic(fmt.Sprintf("error closing image file: %s", imgPath))
		}
	}()

	img, err := png.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("LoadImage: error decoding png image: %w", err)
	}

	return ImageToRgba(img), nil
}

// ImageToRgba converts image.Image to *image.RGBA.
func ImageToRgba(img image.Image) *image.RGBA {
	switch trueImg := img.(type) {
	case *image.RGBA:
		return trueImg
	default:
		rgba := image.NewRGBA(trueImg.Bounds())
		draw.Draw(rgba, trueImg.Bounds(), trueImg, image.Pt(0, 0), draw.Src)
		return rgba
	}
}

// ToVec4Color converts rgba color to imgui.Vec4.
func ToVec4Color(col color.Color) imgui.Vec4 {
	const mask = 0xffff

	r, g, b, a := col.RGBA()
	return imgui.Vec4{
		X: float32(r) / mask,
		Y: float32(g) / mask,
		Z: float32(b) / mask,
		W: float32(a) / mask,
	}
}

// ToVec2 converts image.Point to imgui.Vec2.
func ToVec2(pt image.Point) imgui.Vec2 {
	return imgui.Vec2{
		X: float32(pt.X),
		Y: float32(pt.Y),
	}
}

// Vec4ToRGBA converts imgui's Vec4 to golang rgba color.
func Vec4ToRGBA(vec4 imgui.Vec4) color.RGBA {
	return color.RGBA{
		R: uint8(vec4.X * 255),
		G: uint8(vec4.Y * 255),
		B: uint8(vec4.Z * 255),
		A: uint8(vec4.W * 255),
	}
}

// Update updates giu app
// it is done by default after each frame.
// However because frames stops rendering, when no user
// action is done, it may be necessary to
// Update ui manually at some point.
func Update() {
	if Context.isAlive {
		Context.platform.Update()
	}
}

// GetCursorScreenPos returns imgui drawing cursor on the screen.
func GetCursorScreenPos() image.Point {
	pos := imgui.CursorScreenPos()
	return image.Pt(int(pos.X), int(pos.Y))
}

// SetCursorScreenPos sets imgui drawing cursor on the screen.
func SetCursorScreenPos(pos image.Point) {
	imgui.SetCursorScreenPos(imgui.Vec2{X: float32(pos.X), Y: float32(pos.Y)})
}

// GetCursorPos gets imgui drawing cursor inside of current window.
func GetCursorPos() image.Point {
	pos := imgui.CursorPos()
	return image.Pt(int(pos.X), int(pos.Y))
}

// SetCursorPos sets imgui drawing cursor inside of current window.
func SetCursorPos(pos image.Point) {
	imgui.SetCursorPos(imgui.Vec2{X: float32(pos.X), Y: float32(pos.Y)})
}

// GetMousePos returns mouse position.
func GetMousePos() image.Point {
	pos := imgui.MousePos()
	return image.Pt(int(pos.X), int(pos.Y))
}

// GetAvailableRegion returns region available for rendering.
// it is always WindowSize-WindowPadding*2.
func GetAvailableRegion() (width, height float32) {
	region := imgui.ContentRegionAvail()
	return region.X, region.Y
}

// CalcTextSize calls CalcTextSizeV(text, false, -1).
func CalcTextSize(text string) (width, height float32) {
	return CalcTextSizeV(text, false, -1)
}

// CalcTextSizeV calculates text dimensions.
func CalcTextSizeV(text string, hideAfterDoubleHash bool, wrapWidth float32) (w, h float32) {
	size := imgui.CalcTextSize(text, hideAfterDoubleHash, wrapWidth)
	return size.X, size.Y
}

// SetNextWindowSize sets size of the next window.
func SetNextWindowSize(width, height float32) {
	imgui.SetNextWindowSize(imgui.Vec2{X: width, Y: height})
}

// ExecCondition represents imgui.Condition.
type ExecCondition imgui.Condition

// imgui conditions.
const (
	ConditionAlways       ExecCondition = ExecCondition(imgui.ConditionAlways)
	ConditionOnce         ExecCondition = ExecCondition(imgui.ConditionOnce)
	ConditionFirstUseEver ExecCondition = ExecCondition(imgui.ConditionFirstUseEver)
	ConditionAppearing    ExecCondition = ExecCondition(imgui.ConditionAppearing)
)

// SetNextWindowPos sets position of next window.
func SetNextWindowPos(x, y float32) {
	imgui.SetNextWindowPos(imgui.Vec2{X: x, Y: y})
}

// SetNextWindowSizeV does similar to SetNextWIndowSize but allows to specify imgui.Condition.
func SetNextWindowSizeV(width, height float32, condition ExecCondition) {
	imgui.SetNextWindowSizeV(
		imgui.Vec2{
			X: width,
			Y: height,
		},
		imgui.Condition(condition),
	)
}

// SetItemDefaultFocus set the item focused by default.
func SetItemDefaultFocus() {
	imgui.SetItemDefaultFocus()
}

// SetKeyboardFocusHere sets keyboard focus at *NEXT* widget.
func SetKeyboardFocusHere() {
	SetKeyboardFocusHereV(0)
}

// SetKeyboardFocusHereV sets keyboard on the next widget. Use positive 'offset' to access sub components of a multiple component widget. Use -1 to access previous widget.
func SetKeyboardFocusHereV(i int) {
	imgui.SetKeyboardFocusHereV(i)
}

// PushClipRect pushes a clipping rectangle for both ImGui logic (hit-testing etc.) and low-level ImDrawList rendering.
func PushClipRect(clipRectMin, clipRectMax image.Point, intersectWithClipRect bool) {
	imgui.PushClipRect(ToVec2(clipRectMin), ToVec2(clipRectMax), intersectWithClipRect)
}

// PopClipRect should be called to end PushClipRect.
func PopClipRect() {
	imgui.PopClipRect()
}

// Assert checks if cond. If not cond, it alls golang panic.
func Assert(cond bool, t, method, msg string, args ...any) {
	if !cond {
		fatal(t, method, msg, args...)
	}
}

func fatal(widgetName, method, message string, args ...any) {
	if widgetName != "" {
		widgetName = fmt.Sprintf("(*%s)", widgetName)
	}

	log.Panicf("giu: %s.%s: %s", widgetName, method, fmt.Sprintf(message, args...))
}

// OpenURL opens `url` in default browser.
func OpenURL(url string) {
	if err := browser.OpenURL(url); err != nil {
		log.Printf("Error opening %s: %v", url, err)
	}
}
