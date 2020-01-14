package giu

import (
	"image/color"

	"github.com/AllenDang/giu/imgui"
)

func PushFont(font imgui.Font) {
	imgui.PushFont(font)
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
	imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{X: width, Y: height})
}

func PushFramePadding(width, height float32) {
	imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: width, Y: height})
}

func PushItemSpacing(width, height float32) {
	imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: width, Y: height})
}

// Alignment for button text. Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func PushButtonTextAlign(width, height float32) {
	imgui.PushStyleVarVec2(imgui.StyleVarButtonTextAlign, imgui.Vec2{X: width, Y: height})
}

// Alignment for selectable text. Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func PushSelectableTextAlign(width, height float32) {
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
	imgui.PushItemWidth(width)
}
