package imgui

// #include "StyleWrapper.h"
import "C"

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
	StyleColorTextSelectedBg        StyleColorID = 42
	StyleColorDragDropTarget        StyleColorID = 43
	StyleColorNavHighlight          StyleColorID = 44 // Gamepad/keyboard: current highlighted item
	StyleColorNavWindowingHighlight StyleColorID = 45 // Highlight window when using CTRL+TAB
	StyleColorNavWindowingDarkening StyleColorID = 46 // Darken/colorize entire screen behind the CTRL+TAB window list, when active
	StyleColorModalWindowDarkening  StyleColorID = 47 // Darken/colorize entire screen behind a modal window, when one is active
)

// Style describes the overall graphical representation of the user interface.
type Style uintptr

func (style Style) handle() C.IggGuiStyle {
	return C.IggGuiStyle(style)
}

// ItemInnerSpacing is the horizontal and vertical spacing between elements of
// a composed widget (e.g. a slider and its label).
func (style Style) ItemInnerSpacing() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetItemInnerSpacing(style.handle(), valueArg)
	valueFin()
	return value
}

func (style Style) WindowPadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetWindowPadding(style.handle(), valueArg)
	valueFin()
	return value
}

func (style Style) FramePadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetFramePadding(style.handle(), valueArg)
	valueFin()
	return value
}

// SetColor sets a color value of the UI style.
func (style Style) SetColor(id StyleColorID, value Vec4) {
	valueArg, _ := value.wrapped()
	C.iggStyleSetColor(style.handle(), C.int(id), valueArg)
}

func (style Style) GetColor(id StyleColorID) Vec4 {
	var col Vec4
	colArg, colFin := col.wrapped()
	C.iggStyleGetColor(style.handle(), C.int(id), colArg)
	colFin()
	return col
}

// ScaleAllSizes applies a scaling factor to all sizes.
// To scale your entire UI (e.g. if you want your app to use High DPI or generally be DPI aware) you may use this helper function.
// Scaling the fonts is done separately and is up to you.
//
// Important: This operation is lossy because all sizes are rounded to integer.
// If you need to change your scale multiples, call this over a freshly initialized style rather than scaling multiple times.
func (style Style) ScaleAllSizes(scale float32) {
	C.iggStyleScaleAllSizes(style.handle(), C.float(scale))
}
