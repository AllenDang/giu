package giu

import (
	"image/color"

	"github.com/AllenDang/imgui-go"
)

// PushFont sets font to "font"
// NOTE: PopFont has to be called
// NOTE: Don't use PushFont. use StyleSetter instead
func PushFont(font *FontInfo) bool {
	if f, ok := extraFontMap[font.String()]; ok {
		imgui.PushFont(*f)
		return true
	}
	return false
}

// PopFont pops the font (should be called after PushFont)
func PopFont() {
	imgui.PopFont()
}

// PushStyleColor wrapps imgui.PushStyleColor
// NOTE: don't forget to call PopStyleColor()!
func PushStyleColor(id StyleColorID, col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorID(id), ToVec4Color(col))
}

// PushColorText calls PushStyleColor(StyleColorText,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorText(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorText, ToVec4Color(col))
}

// PushColorTextDisabled calls PushStyleColor(StyleColorTextDisabled,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorTextDisabled(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorTextDisabled, ToVec4Color(col))
}

// PushColorWindowBg calls PushStyleColor(StyleColorWindowBg,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorWindowBg(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorWindowBg, ToVec4Color(col))
}

// PushColorFrameBg calls PushStyleColor(StyleColorFrameBg,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorFrameBg(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorFrameBg, ToVec4Color(col))
}

// PushColorButton calls PushStyleColor(StyleColorButton,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButton(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorButton, ToVec4Color(col))
}

// PushColorButtonHovered calls PushStyleColor(StyleColorButtonHovered,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButtonHovered(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorButtonHovered, ToVec4Color(col))
}

// PushColorButtonActive calls PushStyleColor(StyleColorButtonActive,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButtonActive(col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorButtonActive, ToVec4Color(col))
}

// PushWindowPadding calls PushStyleVar(StyleWindowPadding,...)
func PushWindowPadding(width, height float32) {
	imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{X: width, Y: height})
}

// PushFramePadding calls PushStyleVar(StyleFramePadding,...)
func PushFramePadding(width, height float32) {
	imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: width, Y: height})
}

// PushItemSpacing calls PushStyleVar(StyleVarItemSpacing,...)
func PushItemSpacing(width, height float32) {
	imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: width, Y: height})
}

// PushButtonTextAlign sets alignment for button text. Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func PushButtonTextAlign(width, height float32) {
	imgui.PushStyleVarVec2(imgui.StyleVarButtonTextAlign, imgui.Vec2{X: width, Y: height})
}

// PushSelectableTextAlign sets alignment for selectable text. Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func PushSelectableTextAlign(width, height float32) {
	imgui.PushStyleVarVec2(imgui.StyleVarSelectableTextAlign, imgui.Vec2{X: width, Y: height})
}

// PopStyle should be called to stop applying style.
// It should be called as much times, as you Called PushStyle...
// NOTE: If you don't call PopStyle imgui will panic
func PopStyle() {
	imgui.PopStyleVar()
}

// PopStyleV does similarly to PopStyle, but allows to specify number
// of styles you're going to pop
func PopStyleV(count int) {
	imgui.PopStyleVarV(count)
}

// PopStyleColor is used to stop applying colors styles.
// It should be called after each PushStyleColor... (for each push)
// If PopStyleColor wasn't called after PushColor... or was called
// inproperly, imgui will panic
func PopStyleColor() {
	imgui.PopStyleColor()
}

// PopStyleColorV does similar to PopStyleColor, but allows to specify
// how much style colors would you like to pop
func PopStyleColorV(count int) {
	imgui.PopStyleColorV(count)
}

// AlignTextToFramePadding vertically aligns upcoming text baseline to
// FramePadding.y so that it will align properly to regularly framed
// items. Call if you have text on a line before a framed item.
func AlignTextToFramePadding() {
	imgui.AlignTextToFramePadding()
}

// PushItemWidth sets following item's widths
// NOTE: don't forget to call PopItemWidth! If you don't do so, imgui
// will panic
func PushItemWidth(width float32) {
	imgui.PushItemWidth(width)
}

// PopItemWidth should be called to stop applying PushItemWidth effect
// If it isn't called imgui will panic
func PopItemWidth() {
	imgui.PopItemWidth()
}

func PushTextWrapPos() {
	imgui.PushTextWrapPos()
}

func PopTextWrapPos() {
	imgui.PopTextWrapPos()
}

// MouseCursorType represents a type (layout) of mouse cursor
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

// SetMouseCursor sets mouse cursor layout
func SetMouseCursor(cursor MouseCursorType) {
	imgui.SetMouseCursor(int(cursor))
}

// GetWindowPadding returns window padding
func GetWindowPadding() (x, y float32) {
	vec2 := imgui.CurrentStyle().WindowPadding()
	return vec2.X, vec2.Y
}

// GetItemSpacing returns current item spacing
func GetItemSpacing() (w, h float32) {
	vec2 := imgui.CurrentStyle().ItemSpacing()
	return vec2.X, vec2.Y
}

// GetItemInnerSpacing returns current item inner spacing
func GetItemInnerSpacing() (w, h float32) {
	vec2 := imgui.CurrentStyle().ItemInnerSpacing()
	return vec2.X, vec2.Y
}

// GetFramePadding returns current frame padding
func GetFramePadding() (x, y float32) {
	vec2 := imgui.CurrentStyle().FramePadding()
	return vec2.X, vec2.Y
}

// StyleColorID identifies a color in the UI style.
type StyleColorID int

// StyleColor identifier
const (
	StyleColorText                  StyleColorID = 0
	StyleColorTextDisabled          StyleColorID = 1
	StyleColorWindowBg              StyleColorID = 2
	StyleColorChildBg               StyleColorID = 3
	StyleColorPopupBg               StyleColorID = 4
	StyleColorBorder                StyleColorID = 5
	StyleColorBorderShadow          StyleColorID = 6
	StyleColorFrameBg               StyleColorID = 7
	StyleColorFrameBgHovered        StyleColorID = 8
	StyleColorFrameBgActive         StyleColorID = 9
	StyleColorTitleBg               StyleColorID = 10
	StyleColorTitleBgActive         StyleColorID = 11
	StyleColorTitleBgCollapsed      StyleColorID = 12
	StyleColorMenuBarBg             StyleColorID = 13
	StyleColorScrollbarBg           StyleColorID = 14
	StyleColorScrollbarGrab         StyleColorID = 15
	StyleColorScrollbarGrabHovered  StyleColorID = 16
	StyleColorScrollbarGrabActive   StyleColorID = 17
	StyleColorCheckMark             StyleColorID = 18
	StyleColorSliderGrab            StyleColorID = 19
	StyleColorSliderGrabActive      StyleColorID = 20
	StyleColorButton                StyleColorID = 21
	StyleColorButtonHovered         StyleColorID = 22
	StyleColorButtonActive          StyleColorID = 23
	StyleColorHeader                StyleColorID = 24
	StyleColorHeaderHovered         StyleColorID = 25
	StyleColorHeaderActive          StyleColorID = 26
	StyleColorSeparator             StyleColorID = 27
	StyleColorSeparatorHovered      StyleColorID = 28
	StyleColorSeparatorActive       StyleColorID = 29
	StyleColorResizeGrip            StyleColorID = 30
	StyleColorResizeGripHovered     StyleColorID = 31
	StyleColorResizeGripActive      StyleColorID = 32
	StyleColorTab                   StyleColorID = 33
	StyleColorTabHovered            StyleColorID = 34
	StyleColorTabActive             StyleColorID = 35
	StyleColorTabUnfocused          StyleColorID = 36
	StyleColorTabUnfocusedActive    StyleColorID = 37
	StyleColorPlotLines             StyleColorID = 38
	StyleColorPlotLinesHovered      StyleColorID = 39
	StyleColorProgressBarActive     StyleColorID = 40
	StyleColorPlotHistogram         StyleColorID = 40
	StyleColorPlotHistogramHovered  StyleColorID = 41
	StyleColorTableHeaderBg         StyleColorID = 42
	StyleColorTableBorderStrong     StyleColorID = 43
	StyleColorTableBorderLight      StyleColorID = 44
	StyleColorTableRowBg            StyleColorID = 45
	StyleColorTableRowBgAlt         StyleColorID = 46
	StyleColorTextSelectedBg        StyleColorID = 47
	StyleColorDragDropTarget        StyleColorID = 48
	StyleColorNavHighlight          StyleColorID = 49
	StyleColorNavWindowingHighlight StyleColorID = 50
	StyleColorNavWindowingDimBg     StyleColorID = 51
	StyleColorModalWindowDimBg      StyleColorID = 52
)

// StyleVarID identifies a style variable in the UI style.
type StyleVarID int

// Style IDs
const (
	// StyleVarAlpha is a float
	StyleVarAlpha StyleVarID = iota
	// float     DisabledAlpha
	StyleVarDisabledAlpha
	// StyleVarWindowPadding is a Vec2
	StyleVarWindowPadding
	// StyleVarWindowRounding is a float
	StyleVarWindowRounding
	// StyleVarWindowBorderSize is a float
	StyleVarWindowBorderSize
	// StyleVarWindowMinSize is a Vec2
	StyleVarWindowMinSize
	// StyleVarWindowTitleAlign is a Vec2
	StyleVarWindowTitleAlign
	// StyleVarChildRounding is a float
	StyleVarChildRounding
	// StyleVarChildBorderSize is a float
	StyleVarChildBorderSize
	// StyleVarPopupRounding is a float
	StyleVarPopupRounding
	// StyleVarPopupBorderSize is a float
	StyleVarPopupBorderSize
	// StyleVarFramePadding is a Vec2
	StyleVarFramePadding
	// StyleVarFrameRounding is a float
	StyleVarFrameRounding
	// StyleVarFrameBorderSize is a float
	StyleVarFrameBorderSize
	// StyleVarItemSpacing is a Vec2
	StyleVarItemSpacing
	// StyleVarItemInnerSpacing is a Vec2
	StyleVarItemInnerSpacing
	// StyleVarIndentSpacing is a float
	StyleVarIndentSpacing
	// StyleVarScrollbarSize is a float
	StyleVarScrollbarSize
	// StyleVarScrollbarRounding is a float
	StyleVarScrollbarRounding
	// StyleVarGrabMinSize is a float
	StyleVarGrabMinSize
	// StyleVarGrabRounding is a float
	StyleVarGrabRounding
	// StyleVarTabRounding is a float
	StyleVarTabRounding
	// StyleVarButtonTextAlign is a Vec2
	StyleVarButtonTextAlign
	// StyleVarSelectableTextAlign is a Vec2
	StyleVarSelectableTextAlign
)

var _ Widget = &StyleSetter{}

// StyleSetter is a user-friendly way to manage imgui styles
type StyleSetter struct {
	colors   map[StyleColorID]color.RGBA
	styles   map[StyleVarID]imgui.Vec2
	font     *FontInfo
	disabled bool
	layout   Layout
}

// Style initializes a style setter (see examples/setstyle)
func Style() *StyleSetter {
	var ss StyleSetter
	ss.colors = make(map[StyleColorID]color.RGBA)
	ss.styles = make(map[StyleVarID]imgui.Vec2)

	return &ss
}

// SetColor sets colorID's color
func (ss *StyleSetter) SetColor(colorID StyleColorID, col color.RGBA) *StyleSetter {
	ss.colors[colorID] = col
	return ss
}

// SetStyle sets styleVarID to width and height
func (ss *StyleSetter) SetStyle(varID StyleVarID, width, height float32) *StyleSetter {
	ss.styles[varID] = imgui.Vec2{X: width, Y: height}
	return ss
}

// SetFont sets font
func (ss *StyleSetter) SetFont(font *FontInfo) *StyleSetter {
	ss.font = font
	return ss
}

func (ss *StyleSetter) SetFontSize(size float32) *StyleSetter {
	var font FontInfo
	if ss.font != nil {
		font = *ss.font
	} else {
		font = defaultFonts[0]
	}

	font.size = size

	extraFonts = append(extraFonts, font)

	ss.font = &font

	return ss
}

// SetDisabled sets if items are disabled
func (ss *StyleSetter) SetDisabled(d bool) *StyleSetter {
	ss.disabled = d
	return ss
}

// To allows to specify a layout, StyleSetter should apply style for
func (ss *StyleSetter) To(widgets ...Widget) *StyleSetter {
	ss.layout = widgets
	return ss
}

// Build implements Widget
func (ss *StyleSetter) Build() {
	if ss.layout == nil || len(ss.layout) == 0 {
		return
	}

	for k, v := range ss.colors {
		imgui.PushStyleColor(imgui.StyleColorID(k), ToVec4Color(v))
	}

	for k, v := range ss.styles {
		imgui.PushStyleVarVec2(imgui.StyleVarID(k), v)
	}

	isFontPushed := false
	if ss.font != nil {
		isFontPushed = PushFont(ss.font)
	}

	imgui.BeginDisabled(ss.disabled)

	ss.layout.Build()

	imgui.EndDisabled()

	if isFontPushed {
		PopFont()
	}

	imgui.PopStyleColorV(len(ss.colors))
	imgui.PopStyleVarV(len(ss.styles))
}
