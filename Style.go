package giu

import (
	"image/color"

	"github.com/AllenDang/imgui-go"
)

func PushFont(font *FontInfo) bool {
	if f, ok := extraFontMap[font.String()]; ok {
		imgui.PushFont(*f)
		return true
	}
	return false
}

func PopFont() {
	imgui.PopFont()
}

func PushStyleColor(id StyleColorID, col color.RGBA) {
	imgui.PushStyleColor(imgui.StyleColorID(id), ToVec4Color(col))
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

func GetItemSpacing() (float32, float32) {
	vec2 := imgui.CurrentStyle().ItemSpacing()
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

const (
	// StyleVarAlpha is a float
	StyleVarAlpha StyleVarID = 0
	// StyleVarWindowPadding is a Vec2
	StyleVarWindowPadding StyleVarID = 1
	// StyleVarWindowRounding is a float
	StyleVarWindowRounding StyleVarID = 2
	// StyleVarWindowBorderSize is a float
	StyleVarWindowBorderSize StyleVarID = 3
	// StyleVarWindowMinSize is a Vec2
	StyleVarWindowMinSize StyleVarID = 4
	// StyleVarWindowTitleAlign is a Vec2
	StyleVarWindowTitleAlign StyleVarID = 5
	// StyleVarChildRounding is a float
	StyleVarChildRounding StyleVarID = 6
	// StyleVarChildBorderSize is a float
	StyleVarChildBorderSize StyleVarID = 7
	// StyleVarPopupRounding is a float
	StyleVarPopupRounding StyleVarID = 8
	// StyleVarPopupBorderSize is a float
	StyleVarPopupBorderSize StyleVarID = 9
	// StyleVarFramePadding is a Vec2
	StyleVarFramePadding StyleVarID = 10
	// StyleVarFrameRounding is a float
	StyleVarFrameRounding StyleVarID = 11
	// StyleVarFrameBorderSize is a float
	StyleVarFrameBorderSize StyleVarID = 12
	// StyleVarItemSpacing is a Vec2
	StyleVarItemSpacing StyleVarID = 13
	// StyleVarItemInnerSpacing is a Vec2
	StyleVarItemInnerSpacing StyleVarID = 14
	// StyleVarIndentSpacing is a float
	StyleVarIndentSpacing StyleVarID = 15
	// StyleVarScrollbarSize is a float
	StyleVarScrollbarSize StyleVarID = 16
	// StyleVarScrollbarRounding is a float
	StyleVarScrollbarRounding StyleVarID = 17
	// StyleVarGrabMinSize is a float
	StyleVarGrabMinSize StyleVarID = 18
	// StyleVarGrabRounding is a float
	StyleVarGrabRounding StyleVarID = 19
	// StyleVarTabRounding is a float
	StyleVarTabRounding StyleVarID = 20
	// StyleVarButtonTextAlign is a Vec2
	StyleVarButtonTextAlign StyleVarID = 21
	// StyleVarSelectableTextAlign is a Vec2
	StyleVarSelectableTextAlign StyleVarID = 22
)

type StyleSetter struct {
	colors   map[StyleColorID]color.RGBA
	styles   map[StyleVarID]imgui.Vec2
	font     *FontInfo
	disabled bool
	layout   Layout
}

func Style() *StyleSetter {
	var ss StyleSetter
	ss.colors = make(map[StyleColorID]color.RGBA)
	ss.styles = make(map[StyleVarID]imgui.Vec2)

	return &ss
}

func (ss *StyleSetter) SetColor(colorId StyleColorID, col color.RGBA) *StyleSetter {
	ss.colors[colorId] = col
	return ss
}

func (ss *StyleSetter) SetStyle(varId StyleVarID, width, height float32) *StyleSetter {
	ss.styles[varId] = imgui.Vec2{X: width, Y: height}
	return ss
}

func (ss *StyleSetter) SetFont(font *FontInfo) *StyleSetter {
	ss.font = font
	return ss
}

func (ss *StyleSetter) SetDisabled(d bool) *StyleSetter {
	ss.disabled = d
	return ss
}

func (ss *StyleSetter) To(widgets ...Widget) *StyleSetter {
	ss.layout = widgets
	return ss
}

func (ss *StyleSetter) Build() {
	if len(ss.layout) > 0 {
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

		if ss.disabled {
			imgui.PushDisabled()
		}

		ss.layout.Build()

		if ss.disabled {
			imgui.PopDisabled()
		}

		if isFontPushed {
			PopFont()
		}

		imgui.PopStyleColorV(len(ss.colors))
		imgui.PopStyleVarV(len(ss.styles))
	}
}
