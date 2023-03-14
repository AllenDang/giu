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

	imgui "github.com/AllenDang/cimgui-go"
)

func fatal(widgetName, method, message string, args ...any) {
	if widgetName != "" {
		widgetName = fmt.Sprintf("(*%s)", widgetName)
	}

	log.Panicf("giu: %s.%s: %s", widgetName, method, fmt.Sprintf(message, args...))
}

// Assert checks if cond. If not cond, it alls golang panic.
func Assert(cond bool, t, method, msg string, args ...any) {
	if !cond {
		fatal(t, method, msg, args...)
	}
}

// GetAvailableRegion returns region available for rendering.
// it is always WindowSize-WindowPadding*2.
func GetAvailableRegion() (width, height float32) {
	region := imgui.ContentRegionAvail()
	return region.X, region.Y
}

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

func ToU32(col color.Color) uint32 {
	return imgui.ColorU32Vec4(ToVec4Color(col))
}

// ToVec2 converts image.Point to imgui.Vec2.
func ToVec2(pt image.Point) imgui.Vec2 {
	return imgui.Vec2{
		X: float32(pt.X),
		Y: float32(pt.Y),
	}
}

// Vec4ToRGBA converts imgui's Vec4 to golang rgba color.
func Vec4ToRGBA(vec4 *imgui.Vec4) color.RGBA {
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
		imgui.Refresh()
	}
}

// GetDrawCursorScreenPos returns imgui drawing cursor on the screen.
func GetDrawCursorScreenPos() image.Point {
	pos := imgui.CursorScreenPos()
	return image.Pt(int(pos.X), int(pos.Y))
}

// SetDrawCursorScreenPos sets imgui drawing cursor on the screen.
func SetDrawCursorScreenPos(pos image.Point) {
	imgui.SetCursorScreenPos(imgui.Vec2{X: float32(pos.X), Y: float32(pos.Y)})
}

// GetDrawCursorPos gets imgui drawing cursor inside of current window.
func GetDrawCursorPos() image.Point {
	pos := imgui.CursorPos()
	return image.Pt(int(pos.X), int(pos.Y))
}

// SetDrawCursorPos sets imgui drawing cursor inside of current window.
func SetDrawCursorPos(pos image.Point) {
	imgui.SetCursorPos(imgui.Vec2{X: float32(pos.X), Y: float32(pos.Y)})
}

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

func GetMousePos() image.Point {
	pt := imgui.MousePos()
	return image.Pt(int(pt.X), int(pt.Y))
}

func GetCursorPos() image.Point {
	pt := imgui.CursorPos()
	return image.Pt(int(pt.X), int(pt.Y))
}

func SetCursorPos(pos image.Point) {
	imgui.SetCursorPosX(float32(pos.X))
	imgui.SetCursorPosY(float32(pos.Y))
}
