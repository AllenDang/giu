package giu

import (
	"image/color"

	"github.com/AllenDang/imgui-go"
)

// PushFont sets font to "font"
// NOTE: PopFont has to be called
// NOTE: Don't use PushFont. use StyleSetter instead.
func PushFont(font *FontInfo) bool {
	if font == nil {
		return false
	}

	if f, ok := extraFontMap[font.String()]; ok {
		imgui.PushFont(*f)
		return true
	}

	return false
}

// PopFont pops the font (should be called after PushFont).
func PopFont() {
	imgui.PopFont()
}

// PushStyleColor wrapps imgui.PushStyleColor
// NOTE: don't forget to call PopStyleColor()!
func PushStyleColor(id StyleColorID, col color.Color) {
	imgui.PushStyleColor(imgui.StyleColorID(id), ToVec4Color(col))
}

// PushColorText calls PushStyleColor(StyleColorText,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorText(col color.Color) {
	imgui.PushStyleColor(imgui.StyleColorText, ToVec4Color(col))
}

// PushColorTextDisabled calls PushStyleColor(StyleColorTextDisabled,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorTextDisabled(col color.Color) {
	imgui.PushStyleColor(imgui.StyleColorTextDisabled, ToVec4Color(col))
}

// PushColorWindowBg calls PushStyleColor(StyleColorWindowBg,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorWindowBg(col color.Color) {
	imgui.PushStyleColor(imgui.StyleColorWindowBg, ToVec4Color(col))
}

// PushColorFrameBg calls PushStyleColor(StyleColorFrameBg,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorFrameBg(col color.Color) {
	imgui.PushStyleColor(imgui.StyleColorFrameBg, ToVec4Color(col))
}

// PushColorButton calls PushStyleColor(StyleColorButton,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButton(col color.Color) {
	imgui.PushStyleColor(imgui.StyleColorButton, ToVec4Color(col))
}

// PushColorButtonHovered calls PushStyleColor(StyleColorButtonHovered,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButtonHovered(col color.Color) {
	imgui.PushStyleColor(imgui.StyleColorButtonHovered, ToVec4Color(col))
}

// PushColorButtonActive calls PushStyleColor(StyleColorButtonActive,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButtonActive(col color.Color) {
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
// NOTE: If you don't call PopStyle imgui will panic.
func PopStyle() {
	imgui.PopStyleVar()
}

// PopStyleV does similarly to PopStyle, but allows to specify number
// of styles you're going to pop.
func PopStyleV(count int) {
	imgui.PopStyleVarV(count)
}

// PopStyleColor is used to stop applying colors styles.
// It should be called after each PushStyleColor... (for each push)
// If PopStyleColor wasn't called after PushColor... or was called
// inproperly, imgui will panic.
func PopStyleColor() {
	imgui.PopStyleColor()
}

// PopStyleColorV does similar to PopStyleColor, but allows to specify
// how much style colors would you like to pop.
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
// will panic.
func PushItemWidth(width float32) {
	imgui.PushItemWidth(width)
}

// PopItemWidth should be called to stop applying PushItemWidth effect
// If it isn't called imgui will panic.
func PopItemWidth() {
	imgui.PopItemWidth()
}

// PushTextWrapPos adds the position, where the text should be frapped.
// use PushTextWrapPos, render text. If text reaches frame end,
// rendering will be continued at the start pos in line below.
// NOTE: Don't forget to call PopWrapTextPos
// NOTE: it is done automatically in LabelWidget (see (*LabelWIdget).Wrapped()).
func PushTextWrapPos() {
	imgui.PushTextWrapPos()
}

// PopTextWrapPos should be caled as many times as PushTextWrapPos
// on each frame.
func PopTextWrapPos() {
	imgui.PopTextWrapPos()
}

// MouseCursorType represents a type (layout) of mouse cursor.
type MouseCursorType int

// cursor types.
const (
	// MouseCursorNone no mouse cursor.
	MouseCursorNone MouseCursorType = -1
	// MouseCursorArrow standard arrow mouse cursor.
	MouseCursorArrow MouseCursorType = 0
	// MouseCursorTextInput when hovering over InputText, etc.
	MouseCursorTextInput MouseCursorType = 1
	// MouseCursorResizeAll (Unused by imgui functions).
	MouseCursorResizeAll MouseCursorType = 2
	// MouseCursorResizeNS when hovering over an horizontal border.
	MouseCursorResizeNS MouseCursorType = 3
	// MouseCursorResizeEW when hovering over a vertical border or a column.
	MouseCursorResizeEW MouseCursorType = 4
	// MouseCursorResizeNESW when hovering over the bottom-left corner of a window.
	MouseCursorResizeNESW MouseCursorType = 5
	// MouseCursorResizeNWSE when hovering over the bottom-right corner of a window.
	MouseCursorResizeNWSE MouseCursorType = 6
	// MouseCursorHand (Unused by imgui functions. Use for e.g. hyperlinks).
	MouseCursorHand  MouseCursorType = 7
	MouseCursorCount MouseCursorType = 8
)

// SetMouseCursor sets mouse cursor layout.
func SetMouseCursor(cursor MouseCursorType) {
	imgui.SetMouseCursor(int(cursor))
}

// GetWindowPadding returns window padding.
func GetWindowPadding() (x, y float32) {
	vec2 := imgui.CurrentStyle().WindowPadding()
	return vec2.X, vec2.Y
}

// GetItemSpacing returns current item spacing.
func GetItemSpacing() (w, h float32) {
	vec2 := imgui.CurrentStyle().ItemSpacing()
	return vec2.X, vec2.Y
}

// GetItemInnerSpacing returns current item inner spacing.
func GetItemInnerSpacing() (w, h float32) {
	vec2 := imgui.CurrentStyle().ItemInnerSpacing()
	return vec2.X, vec2.Y
}

// GetFramePadding returns current frame padding.
func GetFramePadding() (x, y float32) {
	vec2 := imgui.CurrentStyle().FramePadding()
	return vec2.X, vec2.Y
}

// StyleColorID identifies a color in the UI style.
type StyleColorID imgui.StyleColorID

// StyleColor identifier.
const (
	StyleColorText                  StyleColorID = StyleColorID(imgui.StyleColorText)
	StyleColorTextDisabled          StyleColorID = StyleColorID(imgui.StyleColorTextDisabled)
	StyleColorWindowBg              StyleColorID = StyleColorID(imgui.StyleColorWindowBg)
	StyleColorChildBg               StyleColorID = StyleColorID(imgui.StyleColorChildBg)
	StyleColorPopupBg               StyleColorID = StyleColorID(imgui.StyleColorPopupBg)
	StyleColorBorder                StyleColorID = StyleColorID(imgui.StyleColorBorder)
	StyleColorBorderShadow          StyleColorID = StyleColorID(imgui.StyleColorBorderShadow)
	StyleColorFrameBg               StyleColorID = StyleColorID(imgui.StyleColorFrameBg)
	StyleColorFrameBgHovered        StyleColorID = StyleColorID(imgui.StyleColorFrameBgHovered)
	StyleColorFrameBgActive         StyleColorID = StyleColorID(imgui.StyleColorFrameBgActive)
	StyleColorTitleBg               StyleColorID = StyleColorID(imgui.StyleColorTitleBg)
	StyleColorTitleBgActive         StyleColorID = StyleColorID(imgui.StyleColorTitleBgActive)
	StyleColorTitleBgCollapsed      StyleColorID = StyleColorID(imgui.StyleColorTitleBgCollapsed)
	StyleColorMenuBarBg             StyleColorID = StyleColorID(imgui.StyleColorMenuBarBg)
	StyleColorScrollbarBg           StyleColorID = StyleColorID(imgui.StyleColorScrollbarBg)
	StyleColorScrollbarGrab         StyleColorID = StyleColorID(imgui.StyleColorScrollbarGrab)
	StyleColorScrollbarGrabHovered  StyleColorID = StyleColorID(imgui.StyleColorScrollbarGrabHovered)
	StyleColorScrollbarGrabActive   StyleColorID = StyleColorID(imgui.StyleColorScrollbarGrabActive)
	StyleColorCheckMark             StyleColorID = StyleColorID(imgui.StyleColorCheckMark)
	StyleColorSliderGrab            StyleColorID = StyleColorID(imgui.StyleColorSliderGrab)
	StyleColorSliderGrabActive      StyleColorID = StyleColorID(imgui.StyleColorSliderGrabActive)
	StyleColorButton                StyleColorID = StyleColorID(imgui.StyleColorButton)
	StyleColorButtonHovered         StyleColorID = StyleColorID(imgui.StyleColorButtonHovered)
	StyleColorButtonActive          StyleColorID = StyleColorID(imgui.StyleColorButtonActive)
	StyleColorHeader                StyleColorID = StyleColorID(imgui.StyleColorHeader)
	StyleColorHeaderHovered         StyleColorID = StyleColorID(imgui.StyleColorHeaderHovered)
	StyleColorHeaderActive          StyleColorID = StyleColorID(imgui.StyleColorHeaderActive)
	StyleColorSeparator             StyleColorID = StyleColorID(imgui.StyleColorSeparator)
	StyleColorSeparatorHovered      StyleColorID = StyleColorID(imgui.StyleColorSeparatorHovered)
	StyleColorSeparatorActive       StyleColorID = StyleColorID(imgui.StyleColorSeparatorActive)
	StyleColorResizeGrip            StyleColorID = StyleColorID(imgui.StyleColorResizeGrip)
	StyleColorResizeGripHovered     StyleColorID = StyleColorID(imgui.StyleColorResizeGripHovered)
	StyleColorResizeGripActive      StyleColorID = StyleColorID(imgui.StyleColorResizeGripActive)
	StyleColorTab                   StyleColorID = StyleColorID(imgui.StyleColorTab)
	StyleColorTabHovered            StyleColorID = StyleColorID(imgui.StyleColorTabHovered)
	StyleColorTabActive             StyleColorID = StyleColorID(imgui.StyleColorTabActive)
	StyleColorTabUnfocused          StyleColorID = StyleColorID(imgui.StyleColorTabUnfocused)
	StyleColorTabUnfocusedActive    StyleColorID = StyleColorID(imgui.StyleColorTabUnfocusedActive)
	StyleColorPlotLines             StyleColorID = StyleColorID(imgui.StyleColorPlotLines)
	StyleColorPlotLinesHovered      StyleColorID = StyleColorID(imgui.StyleColorPlotLinesHovered)
	StyleColorProgressBarActive     StyleColorID = StyleColorPlotLinesHovered
	StyleColorPlotHistogram         StyleColorID = StyleColorID(imgui.StyleColorPlotHistogram)
	StyleColorPlotHistogramHovered  StyleColorID = StyleColorID(imgui.StyleColorPlotHistogramHovered)
	StyleColorTableHeaderBg         StyleColorID = StyleColorID(imgui.StyleColorTableHeaderBg)
	StyleColorTableBorderStrong     StyleColorID = StyleColorID(imgui.StyleColorTableBorderStrong)
	StyleColorTableBorderLight      StyleColorID = StyleColorID(imgui.StyleColorTableBorderLight)
	StyleColorTableRowBg            StyleColorID = StyleColorID(imgui.StyleColorTableRowBg)
	StyleColorTableRowBgAlt         StyleColorID = StyleColorID(imgui.StyleColorTableRowBgAlt)
	StyleColorTextSelectedBg        StyleColorID = StyleColorID(imgui.StyleColorTextSelectedBg)
	StyleColorDragDropTarget        StyleColorID = StyleColorID(imgui.StyleColorDragDropTarget)
	StyleColorNavHighlight          StyleColorID = StyleColorID(imgui.StyleColorNavHighlight)
	StyleColorNavWindowingHighlight StyleColorID = StyleColorID(imgui.StyleColorNavWindowingHighlight)
	StyleColorNavWindowingDimBg     StyleColorID = StyleColorID(imgui.StyleColorNavWindowingDimBg)
	StyleColorModalWindowDimBg      StyleColorID = StyleColorID(imgui.StyleColorModalWindowDimBg)
)

// StyleVarID identifies a style variable in the UI style.
type StyleVarID imgui.StyleVarID

// Style IDs.
const (
	// StyleVarAlpha is a float.
	StyleVarAlpha StyleVarID = StyleVarID(imgui.StyleVarAlpha)
	// float     DisabledAlpha.
	StyleVarDisabledAlpha StyleVarID = StyleVarID(imgui.StyleVarDisabledAlpha)
	// StyleVarWindowPadding is a Vec2.
	StyleVarWindowPadding StyleVarID = StyleVarID(imgui.StyleVarWindowPadding)
	// StyleVarWindowRounding is a float.
	StyleVarWindowRounding StyleVarID = StyleVarID(imgui.StyleVarWindowRounding)
	// StyleVarWindowBorderSize is a float.
	StyleVarWindowBorderSize StyleVarID = StyleVarID(imgui.StyleVarWindowBorderSize)
	// StyleVarWindowMinSize is a Vec2.
	StyleVarWindowMinSize StyleVarID = StyleVarID(imgui.StyleVarWindowMinSize)
	// StyleVarWindowTitleAlign is a Vec2.
	StyleVarWindowTitleAlign StyleVarID = StyleVarID(imgui.StyleVarWindowTitleAlign)
	// StyleVarChildRounding is a float.
	StyleVarChildRounding StyleVarID = StyleVarID(imgui.StyleVarChildRounding)
	// StyleVarChildBorderSize is a float.
	StyleVarChildBorderSize StyleVarID = StyleVarID(imgui.StyleVarChildBorderSize)
	// StyleVarPopupRounding is a float.
	StyleVarPopupRounding StyleVarID = StyleVarID(imgui.StyleVarPopupRounding)
	// StyleVarPopupBorderSize is a float.
	StyleVarPopupBorderSize StyleVarID = StyleVarID(imgui.StyleVarPopupBorderSize)
	// StyleVarFramePadding is a Vec2.
	StyleVarFramePadding StyleVarID = StyleVarID(imgui.StyleVarFramePadding)
	// StyleVarFrameRounding is a float.
	StyleVarFrameRounding StyleVarID = StyleVarID(imgui.StyleVarFrameRounding)
	// StyleVarFrameBorderSize is a float.
	StyleVarFrameBorderSize StyleVarID = StyleVarID(imgui.StyleVarFrameBorderSize)
	// StyleVarItemSpacing is a Vec2.
	StyleVarItemSpacing StyleVarID = StyleVarID(imgui.StyleVarItemSpacing)
	// StyleVarItemInnerSpacing is a Vec2.
	StyleVarItemInnerSpacing StyleVarID = StyleVarID(imgui.StyleVarItemInnerSpacing)
	// StyleVarIndentSpacing is a float.
	StyleVarIndentSpacing StyleVarID = StyleVarID(imgui.StyleVarIndentSpacing)
	// StyleVarScrollbarSize is a float.
	StyleVarScrollbarSize StyleVarID = StyleVarID(imgui.StyleVarScrollbarSize)
	// StyleVarScrollbarRounding is a float.
	StyleVarScrollbarRounding StyleVarID = StyleVarID(imgui.StyleVarScrollbarRounding)
	// StyleVarGrabMinSize is a float.
	StyleVarGrabMinSize StyleVarID = StyleVarID(imgui.StyleVarGrabMinSize)
	// StyleVarGrabRounding is a float.
	StyleVarGrabRounding StyleVarID = StyleVarID(imgui.StyleVarGrabRounding)
	// StyleVarTabRounding is a float.
	StyleVarTabRounding StyleVarID = StyleVarID(imgui.StyleVarTabRounding)
	// StyleVarButtonTextAlign is a Vec2.
	StyleVarButtonTextAlign StyleVarID = StyleVarID(imgui.StyleVarButtonTextAlign)
	// StyleVarSelectableTextAlign is a Vec2.
	StyleVarSelectableTextAlign StyleVarID = StyleVarID(imgui.StyleVarSelectableTextAlign)
)

// IsVec2 returns true if the style var id should be processed as imgui.Vec2
// if not, it is interpreted as float32.
func (s StyleVarID) IsVec2() bool {
	lookup := map[StyleVarID]bool{
		// StyleVarWindowPadding is a Vec2.
		StyleVarWindowPadding:    true,
		StyleVarWindowMinSize:    true,
		StyleVarWindowTitleAlign: true,
		StyleVarFramePadding:     true,
		StyleVarItemSpacing:      true,
		// StyleVarItemInnerSpacing is a Vec2.
		StyleVarItemInnerSpacing:    true,
		StyleVarButtonTextAlign:     true,
		StyleVarSelectableTextAlign: true,
	}

	result, ok := lookup[s]

	return result && ok
}

var _ Widget = &StyleSetter{}

// StyleSetter is a user-friendly way to manage imgui styles.
type StyleSetter struct {
	colors   map[StyleColorID]color.Color
	styles   map[StyleVarID]any
	font     *FontInfo
	disabled bool
	layout   Layout
}

// Style initializes a style setter (see examples/setstyle).
func Style() *StyleSetter {
	var ss StyleSetter
	ss.colors = make(map[StyleColorID]color.Color)
	ss.styles = make(map[StyleVarID]any)

	return &ss
}

// SetColor sets colorID's color.
func (ss *StyleSetter) SetColor(colorID StyleColorID, col color.Color) *StyleSetter {
	ss.colors[colorID] = col
	return ss
}

// SetStyle sets styleVarID to width and height.
func (ss *StyleSetter) SetStyle(varID StyleVarID, width, height float32) *StyleSetter {
	ss.styles[varID] = imgui.Vec2{X: width, Y: height}
	return ss
}

// SetStyleFloat sets styleVarID to float value.
// NOTE: for float typed values see above in comments over
// StyleVarID's comments.
func (ss *StyleSetter) SetStyleFloat(varID StyleVarID, value float32) *StyleSetter {
	ss.styles[varID] = value
	return ss
}

// SetFont sets font.
func (ss *StyleSetter) SetFont(font *FontInfo) *StyleSetter {
	ss.font = font
	return ss
}

// SetFontSize sets size of the font.
// NOTE: Be aware, that StyleSetter needs to add a new font to font atlas for
// each font's size.
func (ss *StyleSetter) SetFontSize(size float32) *StyleSetter {
	var font FontInfo
	if ss.font != nil {
		font = *ss.font
	} else {
		font = defaultFonts[0]
	}

	ss.font = font.SetSize(size)

	return ss
}

// SetDisabled sets if items are disabled.
func (ss *StyleSetter) SetDisabled(d bool) *StyleSetter {
	ss.disabled = d
	return ss
}

// To allows to specify a layout, StyleSetter should apply style for.
func (ss *StyleSetter) To(widgets ...Widget) *StyleSetter {
	ss.layout = widgets
	return ss
}

// Build implements Widget.
func (ss *StyleSetter) Build() {
	if ss.layout == nil || len(ss.layout) == 0 {
		return
	}

	for k, v := range ss.colors {
		imgui.PushStyleColor(imgui.StyleColorID(k), ToVec4Color(v))
	}

	for k, v := range ss.styles {
		if k.IsVec2() {
			var value imgui.Vec2
			switch typed := v.(type) {
			case imgui.Vec2:
				value = typed
			case float32:
				value = imgui.Vec2{X: typed, Y: typed}
			}

			imgui.PushStyleVarVec2(imgui.StyleVarID(k), value)
		} else {
			var value float32
			switch typed := v.(type) {
			case float32:
				value = typed
			case imgui.Vec2:
				value = typed.X
			}

			imgui.PushStyleVarFloat(imgui.StyleVarID(k), value)
		}
	}

	if ss.font != nil {
		if PushFont(ss.font) {
			defer PopFont()
		}
	}

	imgui.BeginDisabled(ss.disabled)

	ss.layout.Build()

	imgui.EndDisabled()

	imgui.PopStyleColorV(len(ss.colors))
	imgui.PopStyleVarV(len(ss.styles))
}
