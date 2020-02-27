package giu

import (
	"image/color"

	"github.com/AllenDang/giu/imgui"
)

func PushFont(font imgui.Font) {
	imgui.PushFont(font)
}

func PopFont() {
	imgui.PopFont()
}

func PushColorText(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorText, ToVec4Color(col))
}

func PushColorTextDisabled(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorTextDisabled, ToVec4Color(col))
}

func PushColorWindowBg(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorWindowBg, ToVec4Color(col))
}

func PushColorFrameBg(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorFrameBg, ToVec4Color(col))
}

func PushColorButton(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorButton, ToVec4Color(col))
}

func PushColorButtonHovered(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorButtonHovered, ToVec4Color(col))
}

func PushColorButtonActive(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorButtonActive, ToVec4Color(col))
}

func PushWindowPadding(width, height float32) {
	width *= Context.GetPlatform().GetContentScale()
	height *= Context.GetPlatform().GetContentScale()
	imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{X: width, Y: height})
}

func PushFramePadding(width, height float32) {
	width *= Context.GetPlatform().GetContentScale()
	height *= Context.GetPlatform().GetContentScale()
	imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: width, Y: height})
}

func PushItemSpacing(width, height float32) {
	width *= Context.GetPlatform().GetContentScale()
	height *= Context.GetPlatform().GetContentScale()
	imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: width, Y: height})
}

// Alignment for button text. Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func PushButtonTextAlign(width, height float32) {
	width *= Context.GetPlatform().GetContentScale()
	height *= Context.GetPlatform().GetContentScale()
	imgui.PushStyleVarVec2(imgui.StyleVarButtonTextAlign, imgui.Vec2{X: width, Y: height})
}

// Alignment for selectable text. Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func PushSelectableTextAlign(width, height float32) {
	width *= Context.GetPlatform().GetContentScale()
	height *= Context.GetPlatform().GetContentScale()
	imgui.PushStyleVarVec2(imgui.StyleVarSelectableTextAlign, imgui.Vec2{X: width, Y: height})
}

func PopStyle() {
	imgui.PopStyleVar()
}

func PopStyleV(count int) {
	imgui.PopStyleVarV(count)
}

func PopStyleColor() {
	imgui.PopStyleColor()
}

func PopStyleColorV(count int) {
	imgui.PopStyleColorV(count)
}

// AlignTextToFramePadding vertically aligns upcoming text baseline to
// FramePadding.y so that it will align properly to regularly framed
// items. Call if you have text on a line before a framed item.
func AlignTextToFramePadding() {
	imgui.AlignTextToFramePadding()
}

func PushItemWidth(width float32) {
	width *= Context.GetPlatform().GetContentScale()
	imgui.PushItemWidth(width)
}

func PopItemWidth() {
	imgui.PopItemWidth()
}

func PushTextWrapPos() {
	imgui.PushTextWrapPos()
}

func PopTextWrapPos() {
	imgui.PopTextWrapPos()
}

type MouseCursorType int

const (
	// MouseCursorNone no mouse cursor
	MouseCursorNone MouseCursorType = -1
	// MouseCursorArrow standard arrow mouse cursor
	MouseCursorArrow MouseCursorType = 0
	// MouseCursorTextInput when hovering over InputText, etc.
	MouseCursorTextInput MouseCursorType = 1
	// MouseCursorResizeAll (Unused by imgui functions)
	MouseCursorResizeAll MouseCursorType = 2
	// MouseCursorResizeNS when hovering over an horizontal border
	MouseCursorResizeNS MouseCursorType = 3
	// MouseCursorResizeEW when hovering over a vertical border or a column
	MouseCursorResizeEW MouseCursorType = 4
	// MouseCursorResizeNESW when hovering over the bottom-left corner of a window
	MouseCursorResizeNESW MouseCursorType = 5
	// MouseCursorResizeNWSE when hovering over the bottom-right corner of a window
	MouseCursorResizeNWSE MouseCursorType = 6
	// MouseCursorHand (Unused by imgui functions. Use for e.g. hyperlinks)
	MouseCursorHand  MouseCursorType = 7
	MouseCursorCount MouseCursorType = 8
)

func SetMouseCursor(cursor MouseCursorType) {
	imgui.SetMouseCursor(int(cursor))
}

func GetWindowPadding() (float32, float32) {
	vec2 := imgui.CurrentStyle().WindowPadding()
	return vec2.X, vec2.Y
}

func GetItemInnerSpacing() (float32, float32) {
	vec2 := imgui.CurrentStyle().ItemInnerSpacing()
	return vec2.X, vec2.Y
}

func GetFramePadding() (float32, float32) {
	vec2 := imgui.CurrentStyle().FramePadding()
	return vec2.X, vec2.Y
}
