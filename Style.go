package giu

import (
	"image/color"

	"github.com/AllenDang/cimgui-go"
)

// You may want to use styles in order to make your app looking more beautiful.
// You have two ways to apply style to a widget:
// 1. Use the StyleSetter e.g.:
//    ```golang
//   	giu.Style().
//  		SetStyle(giu.StyleVarWindowPadding, cimgui.Vec2{10, 10})
//  		SetStyleFloat(giu.StyleVarGrabRounding, 5)
//  		SetColor(giu.StyleColorButton, colornames.Red).
// 			To(/*your widgets here*/),
//   ```
// NOTE/TODO: style variables could be Vec2 or float32 for details see comments
// 2. use PushStyle/PushStyleColor in giu.Custom widget
//    NOTE: remember about calling PopStyle/PopStyleColor at the end of styled section!
//    example:
//    ```golang
// 	  	giu.Custom(func() {
// 		  	cimgui.PushStyleVarFlot(giu.StyleVarFrameRounding, 2)
//    	}),
// 		/*your widgets here*/
//   	giu.Custom(func() {
//   		cimgui.PopStyleVar()
//   	}),
//    ```
// below, you can find a few giu wrappers like PushItemSpacing PushColorFrameBG that
// can be used in a similar way as shown above but without specifying style ID.
//
// See also:
// - examples/setstyle for code example
// - StyleIDs.go for list of all style/color IDs
// - StyleSetter.go for user-friendly giu api for styles

// PushFont sets font to "font"
// NOTE: PopFont has to be called
// NOTE: Don't use PushFont. use StyleSetter instead.
func PushFont(font *FontInfo) bool {
	if font == nil {
		return false
	}

	if f, ok := Context.FontAtlas.extraFontMap[font.String()]; ok {
		cimgui.PushFont(f)
		return true
	}

	return false
}

// PopFont pops the font (should be called after PushFont).
func PopFont() {
	cimgui.PopFont()
}

// PushStyleColorVec4 wraps cimgui.PushStyleColor
// NOTE: don't forget to call PopStyleColor()!
func PushStyleColorVec4(id StyleColorID, col color.Color) {
	cimgui.PushStyleColor_Vec4(cimgui.ImGuiCol(id), ToVec4Color(col))
}

func PushStyleVarVec2(id StyleVarID, width, height float32) {
	cimgui.PushStyleVar_Vec2(cimgui.ImGuiStyleVar(id), cimgui.ImVec2{X: width, Y: height})
}

func PushStyleVarFloat(id StyleVarID, value float32) {
	cimgui.PushStyleVar_Float(cimgui.ImGuiStyleVar(id), value)
}

// PushColorText calls PushStyleColorVec4(StyleColorText,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorText(col color.Color) {
	PushStyleColorVec4(cimgui.ImGuiCol_Text, col)
}

// PushColorTextDisabled calls PushStyleColorVec4(StyleColorTextDisabled,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorTextDisabled(col color.Color) {
	PushStyleColorVec4(cimgui.ImGuiCol_TextDisabled, col)
}

// PushColorWindowBg calls PushStyleColorVec4(StyleColorWindowBg,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorWindowBg(col color.Color) {
	PushStyleColorVec4(cimgui.ImGuiCol_WindowBg, col)
}

// PushColorFrameBg calls PushStyleColorVec4(StyleColorFrameBg,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorFrameBg(col color.Color) {
	PushStyleColorVec4(cimgui.ImGuiCol_FrameBg, col)
}

// PushColorButton calls PushStyleColorVec4(StyleColorButton,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButton(col color.Color) {
	PushStyleColorVec4(cimgui.ImGuiCol_Button, col)
}

// PushColorButtonHovered calls PushStyleColorVec4(StyleColorButtonHovered,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButtonHovered(col color.Color) {
	PushStyleColorVec4(cimgui.ImGuiCol_ButtonHovered, col)
}

// PushColorButtonActive calls PushStyleColorVec4(StyleColorButtonActive,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButtonActive(col color.Color) {
	PushStyleColorVec4(cimgui.ImGuiCol_ButtonActive, col)
}

// PushWindowPadding calls PushStyleVar(StyleWindowPadding,...)
func PushWindowPadding(width, height float32) {
	PushStyleVarVec2(StyleVarWindowPadding, width, height)
}

// PushFramePadding calls PushStyleVar(StyleFramePadding,...)
func PushFramePadding(width, height float32) {
	PushStyleVarVec2(StyleVarFramePadding, width, height)
}

// PushItemSpacing calls PushStyleVar(StyleVarItemSpacing,...)
func PushItemSpacing(width, height float32) {
	PushStyleVarVec2(StyleVarItemSpacing, width, height)
}

// PushButtonTextAlign sets alignment for button text. Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func PushButtonTextAlign(width, height float32) {
	PushStyleVarVec2(StyleVarButtonTextAlign, width, height)
}

// PushSelectableTextAlign sets alignment for selectable text. Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func PushSelectableTextAlign(width, height float32) {
	PushStyleVarVec2(StyleVarSelectableTextAlign, width, height)
}

// PopStyle should be called to stop applying style.
// It should be called as many times, as you called PushStyle...
// NOTE: If you don't call PopStyle cimgui will panic.
func PopStyle() {
	cimgui.PopStyleVar()
}

// PopStyleV does similarly to PopStyle, but allows to specify number
// of styles you're going to pop.
func PopStyleV(count int) {
	cimgui.PopStyleVarV(int32(count))
}

// PopStyleColor is used to stop applying colors styles.
// It should be called after each PushStyleColor... (for each push)
// If PopStyleColor wasn't called after PushColor... or was called
// improperly, cimgui will panic.
func PopStyleColor() {
	cimgui.PopStyleColor()
}

// PopStyleColorV does similar to PopStyleColor, but allows to specify
// how much style colors would you like to pop.
func PopStyleColorV(count int) {
	cimgui.PopStyleColorV(int32(count))
}

// AlignTextToFramePadding vertically aligns upcoming text baseline to
// FramePadding.y so that it will align properly to regularly framed
// items. Call if you have text on a line before a framed item.
func AlignTextToFramePadding() {
	cimgui.AlignTextToFramePadding()
}

// PushItemWidth sets following item's widths
// NOTE: don't forget to call PopItemWidth! If you don't do so, cimgui
// will panic.
func PushItemWidth(width float32) {
	cimgui.PushItemWidth(width)
}

// PopItemWidth should be called to stop applying PushItemWidth effect
// If it isn't called cimgui will panic.
func PopItemWidth() {
	cimgui.PopItemWidth()
}

// PushTextWrapPos adds the position, where the text should be wrapped.
// use PushTextWrapPos, render text. If text reaches frame end,
// rendering will be continued at the start pos in line below.
// NOTE: Don't forget to call PopWrapTextPos
// NOTE: it is done automatically in LabelWidget (see (*LabelWidget).Wrapped()).
func PushTextWrapPos() {
	cimgui.PushTextWrapPos()
}

// PopTextWrapPos should be called as many times as PushTextWrapPos
// on each frame.
func PopTextWrapPos() {
	cimgui.PopTextWrapPos()
}

// MouseCursorType represents a type (layout) of mouse cursor.
type MouseCursorType cimgui.ImGuiMouseCursor

// cursor types.
const (
	// MouseCursorNone no mouse cursor.
	MouseCursorNone MouseCursorType = cimgui.ImGuiMouseCursor_None
	// MouseCursorArrow standard arrow mouse cursor.
	MouseCursorArrow MouseCursorType = cimgui.ImGuiMouseCursor_Arrow
	// MouseCursorTextInput when hovering over InputText, etc.
	MouseCursorTextInput MouseCursorType = cimgui.ImGuiMouseCursor_TextInput
	// MouseCursorResizeAll (Unused by cimgui functions).
	MouseCursorResizeAll MouseCursorType = cimgui.ImGuiMouseCursor_ResizeAll
	// MouseCursorResizeNS when hovering over an horizontal border.
	MouseCursorResizeNS MouseCursorType = cimgui.ImGuiMouseCursor_ResizeNS
	// MouseCursorResizeEW when hovering over a vertical border or a column.
	MouseCursorResizeEW MouseCursorType = cimgui.ImGuiMouseCursor_ResizeEW
	// MouseCursorResizeNESW when hovering over the bottom-left corner of a window.
	MouseCursorResizeNESW MouseCursorType = cimgui.ImGuiMouseCursor_ResizeNESW
	// MouseCursorResizeNWSE when hovering over the bottom-right corner of a window.
	MouseCursorResizeNWSE MouseCursorType = cimgui.ImGuiMouseCursor_ResizeNWSE
	// MouseCursorHand (Unused by cimgui functions. Use for e.g. hyperlinks).
	MouseCursorHand       MouseCursorType = cimgui.ImGuiMouseCursor_Hand
	MouseCursorNotAllowed MouseCursorType = cimgui.ImGuiMouseCursor_NotAllowed
	MouseCursorCount      MouseCursorType = cimgui.ImGuiMouseCursor_COUNT
)

// SetMouseCursor sets mouse cursor layout.
func SetMouseCursor(cursor MouseCursorType) {
	cimgui.SetMouseCursor(cimgui.ImGuiMouseCursor(cursor))
}

// GetWindowPadding returns window padding.
func GetWindowPadding() (x, y float32) {
	vec2 := cimgui.GetStyle().GetWindowPadding()
	return vec2.X, vec2.Y
}

// GetItemSpacing returns current item spacing.
func GetItemSpacing() (w, h float32) {
	vec2 := cimgui.GetStyle().GetItemSpacing()
	return vec2.X, vec2.Y
}

// GetItemInnerSpacing returns current item inner spacing.
func GetItemInnerSpacing() (w, h float32) {
	vec2 := cimgui.GetStyle().GetItemInnerSpacing()
	return vec2.X, vec2.Y
}

// GetFramePadding returns current frame padding.
func GetFramePadding() (x, y float32) {
	vec2 := cimgui.GetStyle().GetFramePadding()
	return vec2.X, vec2.Y
}
