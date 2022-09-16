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

	"github.com/AllenDang/cimgui-go"
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
	var region imgui.ImVec2
	imgui.GetContentRegionAvail(&region)
	return region.X, region.Y
}

func ToVec4Color(col color.Color) imgui.ImVec4 {
	const mask = 0xffff

	r, g, b, a := col.RGBA()
	return imgui.ImVec4{
		X: float32(r) / mask,
		Y: float32(g) / mask,
		Z: float32(b) / mask,
		W: float32(a) / mask,
	}
}

func ToU32(col color.Color) uint32 {
	return imgui.GetColorU32_Vec4(ToVec4Color(col))
}

// ToVec2 converts image.Point to imgui.Vec2.
func ToVec2(pt image.Point) imgui.ImVec2 {
	return imgui.ImVec2{
		X: float32(pt.X),
		Y: float32(pt.Y),
	}
}

// Vec4ToRGBA converts imgui's Vec4 to golang rgba color.
func Vec4ToRGBA(vec4 imgui.ImVec4) color.RGBA {
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
		cimgui.Refresh()
	}
}

// GetDrawCursorScreenPos returns imgui drawing cursor on the screen.
func GetDrawCursorScreenPos() image.Point {
	var pos imgui.ImVec2
	imgui.GetDrawCursorScreenPos(&pos)
	return image.Pt(int(pos.X), int(pos.Y))
}

// SetDrawCursorScreenPos sets imgui drawing cursor on the screen.
func SetDrawCursorScreenPos(pos image.Point) {
	imgui.SetDrawCursorScreenPos(imgui.ImVec2{X: float32(pos.X), Y: float32(pos.Y)})
}

// GetDrawCursorPos gets imgui drawing cursor inside of current window.
func GetDrawCursorPos() image.Point {
	var pos imgui.ImVec2
	imgui.GetDrawCursorPos(&pos)
	return image.Pt(int(pos.X), int(pos.Y))
}

// SetDrawCursorPos sets imgui drawing cursor inside of current window.
func SetDrawCursorPos(pos image.Point) {
	imgui.SetDrawCursorPos(imgui.ImVec2{X: float32(pos.X), Y: float32(pos.Y)})
}

// GetItemSpacing returns current item spacing.
func GetItemSpacing() (w, h float32) {
	vec2 := imgui.GetStyle().GetItemSpacing()
	return vec2.X, vec2.Y
}

// GetItemInnerSpacing returns current item inner spacing.
func GetItemInnerSpacing() (w, h float32) {
	vec2 := imgui.GetStyle().GetItemInnerSpacing()
	return vec2.X, vec2.Y
}

// GetFramePadding returns current frame padding.
func GetFramePadding() (x, y float32) {
	vec2 := imgui.GetStyle().GetFramePadding()
	return vec2.X, vec2.Y
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
	var pt imgui.ImVec2
	imgui.GetMousePos(&pt)
	return image.Pt(int(pt.X), int(pt.Y))
}
