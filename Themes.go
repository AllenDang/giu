package giu

import "github.com/AllenDang/cimgui-go/imgui"

// DefaultTheme generates a default GIU theme's StyleSetter.
//
//nolint:dupl // This is something like "data file", so I'd prefer to keep it as simple as possible.
func DefaultTheme() *StyleSetter {
	return Style().
		SetStyleFloat(StyleVarWindowRounding, 2).
		SetStyleFloat(StyleVarFrameRounding, 4).
		SetStyleFloat(StyleVarGrabRounding, 4).
		SetStyleFloat(StyleVarFrameBorderSize, 1).
		SetColorVec4(StyleColorText, imgui.Vec4{X: 0.95, Y: 0.96, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorTextDisabled, imgui.Vec4{X: 0.36, Y: 0.42, Z: 0.47, W: 1.00}).
		SetColorVec4(StyleColorWindowBg, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00}).
		SetColorVec4(StyleColorChildBg, imgui.Vec4{X: 0.15, Y: 0.18, Z: 0.22, W: 1.00}).
		SetColorVec4(StyleColorPopupBg, imgui.Vec4{X: 0.08, Y: 0.08, Z: 0.08, W: 0.94}).
		SetColorVec4(StyleColorBorder, imgui.Vec4{X: 0.08, Y: 0.10, Z: 0.12, W: 1.00}).
		SetColorVec4(StyleColorBorderShadow, imgui.Vec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.00}).
		SetColorVec4(StyleColorFrameBg, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00}).
		SetColorVec4(StyleColorFrameBgHovered, imgui.Vec4{X: 0.12, Y: 0.20, Z: 0.28, W: 1.00}).
		SetColorVec4(StyleColorFrameBgActive, imgui.Vec4{X: 0.09, Y: 0.12, Z: 0.14, W: 1.00}).
		SetColorVec4(StyleColorTitleBg, imgui.Vec4{X: 0.09, Y: 0.12, Z: 0.14, W: 0.65}).
		SetColorVec4(StyleColorTitleBgActive, imgui.Vec4{X: 0.08, Y: 0.10, Z: 0.12, W: 1.00}).
		SetColorVec4(StyleColorTitleBgCollapsed, imgui.Vec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.51}).
		SetColorVec4(StyleColorMenuBarBg, imgui.Vec4{X: 0.15, Y: 0.18, Z: 0.22, W: 1.00}).
		SetColorVec4(StyleColorScrollbarBg, imgui.Vec4{X: 0.02, Y: 0.02, Z: 0.02, W: 0.39}).
		SetColorVec4(StyleColorScrollbarGrab, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00}).
		SetColorVec4(StyleColorScrollbarGrabHovered, imgui.Vec4{X: 0.18, Y: 0.22, Z: 0.25, W: 1.00}).
		SetColorVec4(StyleColorScrollbarGrabActive, imgui.Vec4{X: 0.09, Y: 0.21, Z: 0.31, W: 1.00}).
		SetColorVec4(StyleColorCheckMark, imgui.Vec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00}).
		SetColorVec4(StyleColorSliderGrab, imgui.Vec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00}).
		SetColorVec4(StyleColorSliderGrabActive, imgui.Vec4{X: 0.37, Y: 0.61, Z: 1.00, W: 1.00}).
		SetColorVec4(StyleColorButton, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00}).
		SetColorVec4(StyleColorButtonHovered, imgui.Vec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00}).
		SetColorVec4(StyleColorButtonActive, imgui.Vec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorHeader, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 0.55}).
		SetColorVec4(StyleColorHeaderHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80}).
		SetColorVec4(StyleColorHeaderActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorSeparator, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00}).
		SetColorVec4(StyleColorSeparatorHovered, imgui.Vec4{X: 0.10, Y: 0.40, Z: 0.75, W: 0.78}).
		SetColorVec4(StyleColorSeparatorActive, imgui.Vec4{X: 0.10, Y: 0.40, Z: 0.75, W: 1.00}).
		SetColorVec4(StyleColorResizeGrip, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.25}).
		SetColorVec4(StyleColorResizeGripHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.67}).
		SetColorVec4(StyleColorResizeGripActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.95}).
		SetColorVec4(StyleColorTab, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00}).
		SetColorVec4(StyleColorTabHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80}).
		SetColorVec4(StyleColorTabActive, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00}).
		SetColorVec4(StyleColorTabUnfocused, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00}).
		SetColorVec4(StyleColorTabUnfocusedActive, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00}).
		SetColorVec4(StyleColorPlotLines, imgui.Vec4{X: 0.61, Y: 0.61, Z: 0.61, W: 1.00}).
		SetColorVec4(StyleColorPlotLinesHovered, imgui.Vec4{X: 1.00, Y: 0.43, Z: 0.35, W: 1.00}).
		SetColorVec4(StyleColorPlotHistogram, imgui.Vec4{X: 0.90, Y: 0.70, Z: 0.00, W: 1.00}).
		SetColorVec4(StyleColorPlotHistogramHovered, imgui.Vec4{X: 1.00, Y: 0.60, Z: 0.00, W: 1.00}).
		SetColorVec4(StyleColorTextSelectedBg, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.35}).
		SetColorVec4(StyleColorDragDropTarget, imgui.Vec4{X: 1.00, Y: 1.00, Z: 0.00, W: 0.90}).
		SetColorVec4(StyleColorNavWindowingHighlight, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorNavWindowingHighlight, imgui.Vec4{X: 1.00, Y: 1.00, Z: 1.00, W: 0.70}).
		SetColorVec4(StyleColorTableHeaderBg, imgui.Vec4{X: 0.12, Y: 0.20, Z: 0.28, W: 1.00}).
		SetColorVec4(StyleColorTableBorderStrong, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00}).
		SetColorVec4(StyleColorTableBorderLight, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 0.70})
}

// LightTheme generates a default GIU theme's StyleSetter.
//
//nolint:dupl // This is something like "data file", so I'd prefer to keep it as simple as possible.
func LightTheme() *StyleSetter {
	return Style().
		SetStyleFloat(StyleVarWindowRounding, 2).
		SetStyleFloat(StyleVarFrameRounding, 4).
		SetStyleFloat(StyleVarGrabRounding, 4).
		SetStyleFloat(StyleVarFrameBorderSize, 1).
		SetColorVec4(StyleColorText, imgui.Vec4{X: 0.10, Y: 0.10, Z: 0.10, W: 1.00}).
		SetColorVec4(StyleColorTextDisabled, imgui.Vec4{X: 0.60, Y: 0.60, Z: 0.60, W: 1.00}).
		SetColorVec4(StyleColorWindowBg, imgui.Vec4{X: 0.94, Y: 0.94, Z: 0.94, W: 1.00}).
		SetColorVec4(StyleColorChildBg, imgui.Vec4{X: 0.97, Y: 0.97, Z: 0.97, W: 1.00}).
		SetColorVec4(StyleColorPopupBg, imgui.Vec4{X: 1.00, Y: 1.00, Z: 1.00, W: 0.98}).
		SetColorVec4(StyleColorBorder, imgui.Vec4{X: 0.70, Y: 0.70, Z: 0.70, W: 1.00}).
		SetColorVec4(StyleColorBorderShadow, imgui.Vec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.00}).
		SetColorVec4(StyleColorFrameBg, imgui.Vec4{X: 1.00, Y: 1.00, Z: 1.00, W: 1.00}).
		SetColorVec4(StyleColorFrameBgHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.40}).
		SetColorVec4(StyleColorFrameBgActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.67}).
		SetColorVec4(StyleColorTitleBg, imgui.Vec4{X: 0.96, Y: 0.96, Z: 0.96, W: 1.00}).
		SetColorVec4(StyleColorTitleBgActive, imgui.Vec4{X: 0.82, Y: 0.82, Z: 0.82, W: 1.00}).
		SetColorVec4(StyleColorTitleBgCollapsed, imgui.Vec4{X: 1.00, Y: 1.00, Z: 1.00, W: 0.51}).
		SetColorVec4(StyleColorMenuBarBg, imgui.Vec4{X: 0.86, Y: 0.86, Z: 0.86, W: 1.00}).
		SetColorVec4(StyleColorScrollbarBg, imgui.Vec4{X: 0.98, Y: 0.98, Z: 0.98, W: 0.53}).
		SetColorVec4(StyleColorScrollbarGrab, imgui.Vec4{X: 0.69, Y: 0.69, Z: 0.69, W: 1.00}).
		SetColorVec4(StyleColorScrollbarGrabHovered, imgui.Vec4{X: 0.59, Y: 0.59, Z: 0.59, W: 1.00}).
		SetColorVec4(StyleColorScrollbarGrabActive, imgui.Vec4{X: 0.49, Y: 0.49, Z: 0.49, W: 1.00}).
		SetColorVec4(StyleColorCheckMark, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorSliderGrab, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorSliderGrabActive, imgui.Vec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorButton, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.40}).
		SetColorVec4(StyleColorButtonHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorButtonActive, imgui.Vec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorHeader, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.31}).
		SetColorVec4(StyleColorHeaderHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80}).
		SetColorVec4(StyleColorHeaderActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorSeparator, imgui.Vec4{X: 0.39, Y: 0.39, Z: 0.39, W: 0.62}).
		SetColorVec4(StyleColorSeparatorHovered, imgui.Vec4{X: 0.14, Y: 0.44, Z: 0.80, W: 0.78}).
		SetColorVec4(StyleColorSeparatorActive, imgui.Vec4{X: 0.14, Y: 0.44, Z: 0.80, W: 1.00}).
		SetColorVec4(StyleColorResizeGrip, imgui.Vec4{X: 0.35, Y: 0.35, Z: 0.35, W: 0.17}).
		SetColorVec4(StyleColorResizeGripHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.67}).
		SetColorVec4(StyleColorResizeGripActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.95}).
		SetColorVec4(StyleColorTab, imgui.Vec4{X: 0.76, Y: 0.80, Z: 0.84, W: 0.93}).
		SetColorVec4(StyleColorTabHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80}).
		SetColorVec4(StyleColorTabActive, imgui.Vec4{X: 0.60, Y: 0.73, Z: 0.88, W: 1.00}).
		SetColorVec4(StyleColorTabUnfocused, imgui.Vec4{X: 0.92, Y: 0.93, Z: 0.94, W: 1.00}).
		SetColorVec4(StyleColorTabUnfocusedActive, imgui.Vec4{X: 0.74, Y: 0.82, Z: 0.91, W: 1.00}).
		SetColorVec4(StyleColorPlotLines, imgui.Vec4{X: 0.39, Y: 0.39, Z: 0.39, W: 1.00}).
		SetColorVec4(StyleColorPlotLinesHovered, imgui.Vec4{X: 1.00, Y: 0.43, Z: 0.35, W: 1.00}).
		SetColorVec4(StyleColorPlotHistogram, imgui.Vec4{X: 0.90, Y: 0.70, Z: 0.00, W: 1.00}).
		SetColorVec4(StyleColorPlotHistogramHovered, imgui.Vec4{X: 1.00, Y: 0.45, Z: 0.00, W: 1.00}).
		SetColorVec4(StyleColorTextSelectedBg, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.35}).
		SetColorVec4(StyleColorDragDropTarget, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.95}).
		SetColorVec4(StyleColorNavWindowingHighlight, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80}).
		SetColorVec4(StyleColorNavWindowingHighlight, imgui.Vec4{X: 0.70, Y: 0.70, Z: 0.70, W: 0.70}).
		SetColorVec4(StyleColorTableHeaderBg, imgui.Vec4{X: 0.78, Y: 0.87, Z: 0.98, W: 1.00}).
		SetColorVec4(StyleColorTableBorderStrong, imgui.Vec4{X: 0.57, Y: 0.57, Z: 0.64, W: 1.00}).
		SetColorVec4(StyleColorTableBorderLight, imgui.Vec4{X: 0.68, Y: 0.68, Z: 0.74, W: 1.00})
}
