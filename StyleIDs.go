package giu

import "github.com/AllenDang/imgui-go"

// Here are the style IDs for styling imgui apps.
// For details about each of attributes read comment above them.
// You have two ways to apply style to a widget:
// 1. Use the StyleSetter e.g.:
//   giu.Style().
//  	SetStyle(giu.StyleVarWindowPadding, imgui.Vec2{10, 10})
//  	SetStyleFloat(giu.StyleVarGrabRounding, 5)
//  	SetColor(giu.StyleColorButton, colornames.Red).
// 		To(/*your widgets here*/),
// NOTE/TODO: style variables could be Vec2 or float32 for details see comments
// 2. use PushStyle/PushStyleColor in giu.Custom widget
//    NOTE: remember about calling PopStyle/PopStyleColor at the end of styled section!
//    example:
//    ```golang
// 	  	giu.Custom(func() {
// 		  	imgui.PushStyleVarFlot(giu.StyleVarFrameRounding, 2)
//    	}),
// 		/*your widgets here*/
//   	giu.Custom(func() {
//   		imgui.PopStyleVar()
//   	}),
//    ```
//
// for more details, see examples/setstyle

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
