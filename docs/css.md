# INTRO

GIU provides a special widget called "CSSWidget".
This widget allows to set App's style for your app.

# Usage

1. open your stylesheet (e.g. with go-embed)
2. Tell giu about your stylesheet using `giu.ParseCSSStyleSheet(...)`
3. Put css tags in your code - `giu.CSS("tag name")`

For simple use-case see [examples/CSS-styling](../examples/CSS-styling/)

# Styles list

## Colors

StyleColorText                  = StyleColorID(imgui.StyleColorText)                  // color
StyleColorTextDisabled          = StyleColorID(imgui.StyleColorTextDisabled)          // disabled-color
StyleColorWindowBg              = StyleColorID(imgui.StyleColorWindowBg)              // background-color
StyleColorChildBg               = StyleColorID(imgui.StyleColorChildBg)               // child-background-color
StyleColorPopupBg               = StyleColorID(imgui.StyleColorPopupBg)               // popup-background-color
StyleColorBorder                = StyleColorID(imgui.StyleColorBorder)                // border-color
StyleColorBorderShadow          = StyleColorID(imgui.StyleColorBorderShadow)          // border-shadow-color
StyleColorFrameBg               = StyleColorID(imgui.StyleColorFrameBg)               // frame-background-color
StyleColorFrameBgHovered        = StyleColorID(imgui.StyleColorFrameBgHovered)        // frame-background-hovered-color
StyleColorFrameBgActive         = StyleColorID(imgui.StyleColorFrameBgActive)         // frame-background-active-color
StyleColorTitleBg               = StyleColorID(imgui.StyleColorTitleBg)               // title-background-color
StyleColorTitleBgActive         = StyleColorID(imgui.StyleColorTitleBgActive)         // title-background-active-color
StyleColorTitleBgCollapsed      = StyleColorID(imgui.StyleColorTitleBgCollapsed)      // title-background-collapsed-color
StyleColorMenuBarBg             = StyleColorID(imgui.StyleColorMenuBarBg)             // menu-bar-background-color
StyleColorScrollbarBg           = StyleColorID(imgui.StyleColorScrollbarBg)           // scrollbar-background-color
StyleColorScrollbarGrab         = StyleColorID(imgui.StyleColorScrollbarGrab)         // scrollbar-grab-color
StyleColorScrollbarGrabHovered  = StyleColorID(imgui.StyleColorScrollbarGrabHovered)  // scrollbar-grab-hovered-color
StyleColorScrollbarGrabActive   = StyleColorID(imgui.StyleColorScrollbarGrabActive)   // scrollbar-grab-active-color
StyleColorCheckMark             = StyleColorID(imgui.StyleColorCheckMark)             // checkmark-color
StyleColorSliderGrab            = StyleColorID(imgui.StyleColorSliderGrab)            // slider-grab-color
StyleColorSliderGrabActive      = StyleColorID(imgui.StyleColorSliderGrabActive)      // slider-grab-active-color
StyleColorButton                = StyleColorID(imgui.StyleColorButton)                // button-color
StyleColorButtonHovered         = StyleColorID(imgui.StyleColorButtonHovered)         // button-hovered-color
StyleColorButtonActive          = StyleColorID(imgui.StyleColorButtonActive)          // button-active-color
StyleColorHeader                = StyleColorID(imgui.StyleColorHeader)                // header-color
StyleColorHeaderHovered         = StyleColorID(imgui.StyleColorHeaderHovered)         // header-hovered-color
StyleColorHeaderActive          = StyleColorID(imgui.StyleColorHeaderActive)          // header-active-color
StyleColorSeparator             = StyleColorID(imgui.StyleColorSeparator)             // separator-color
StyleColorSeparatorHovered      = StyleColorID(imgui.StyleColorSeparatorHovered)      // separator-hovered-color
StyleColorSeparatorActive       = StyleColorID(imgui.StyleColorSeparatorActive)       // separator-active-color
StyleColorResizeGrip            = StyleColorID(imgui.StyleColorResizeGrip)            // resize-grip-color
StyleColorResizeGripHovered     = StyleColorID(imgui.StyleColorResizeGripHovered)     // resize-grip-hovered-color
StyleColorResizeGripActive      = StyleColorID(imgui.StyleColorResizeGripActive)      // resize-grip-active-color
StyleColorTab                   = StyleColorID(imgui.StyleColorTab)                   // tab-color
StyleColorTabHovered            = StyleColorID(imgui.StyleColorTabHovered)            // tab-hovered-color
StyleColorTabActive             = StyleColorID(imgui.StyleColorTabActive)             // tab-active-color
StyleColorTabUnfocused          = StyleColorID(imgui.StyleColorTabUnfocused)          // tab-unfocused-color
StyleColorTabUnfocusedActive    = StyleColorID(imgui.StyleColorTabUnfocusedActive)    // tab-unfocused-active-color
StyleColorPlotLines             = StyleColorID(imgui.StyleColorPlotLines)             // plot-lines-color
StyleColorPlotLinesHovered      = StyleColorID(imgui.StyleColorPlotLinesHovered)      // plot-lines-hovered-color
StyleColorProgressBarActive     = StyleColorPlotLinesHovered                          // progress-bar-active-color
StyleColorPlotHistogram         = StyleColorID(imgui.StyleColorPlotHistogram)         // plot-histogram-color
StyleColorPlotHistogramHovered  = StyleColorID(imgui.StyleColorPlotHistogramHovered)  // plot-histogram-hovered-color
StyleColorTableHeaderBg         = StyleColorID(imgui.StyleColorTableHeaderBg)         // table-header-background-color
StyleColorTableBorderStrong     = StyleColorID(imgui.StyleColorTableBorderStrong)     // table-border-strong-color
StyleColorTableBorderLight      = StyleColorID(imgui.StyleColorTableBorderLight)      // table-border-light-color
StyleColorTableRowBg            = StyleColorID(imgui.StyleColorTableRowBg)            // table-row-background-color
StyleColorTableRowBgAlt         = StyleColorID(imgui.StyleColorTableRowBgAlt)         // table-row-alternate-background-color
StyleColorTextSelectedBg        = StyleColorID(imgui.StyleColorTextSelectedBg)        // text-selected-background-color
StyleColorDragDropTarget        = StyleColorID(imgui.StyleColorDragDropTarget)        // drag-drop-target-color
StyleColorNavHighlight          = StyleColorID(imgui.StyleColorNavHighlight)          // navigation-highlight-color
StyleColorNavWindowingHighlight = StyleColorID(imgui.StyleColorNavWindowingHighlight) // windowing-highlight-color
StyleColorNavWindowingDimBg     = StyleColorID(imgui.StyleColorNavWindowingDimBg)     // windowing-dim-background-color
StyleColorModalWindowDimBg      = StyleColorID(imgui.StyleColorModalWindowDimBg)      // modal-window-dim-background-color

## Style Variables

// StyleVarAlpha is a float.
StyleVarAlpha = StyleVarID(imgui.StyleVarAlpha) // alpha
// StyleVarDisabledAlpha is a float.
StyleVarDisabledAlpha = StyleVarID(imgui.StyleVarDisabledAlpha) // disabled-alpha
// StyleVarWindowPadding is a Vec2.
StyleVarWindowPadding = StyleVarID(imgui.StyleVarWindowPadding) // window-padding
// StyleVarWindowRounding is a float.
StyleVarWindowRounding = StyleVarID(imgui.StyleVarWindowRounding) // window-rounding
// StyleVarWindowBorderSize is a float.
StyleVarWindowBorderSize = StyleVarID(imgui.StyleVarWindowBorderSize) // window-border-size
// StyleVarWindowMinSize is a Vec2.
StyleVarWindowMinSize = StyleVarID(imgui.StyleVarWindowMinSize) // window-min-size
// StyleVarWindowTitleAlign is a Vec2.
StyleVarWindowTitleAlign = StyleVarID(imgui.StyleVarWindowTitleAlign) // window-title-align
// StyleVarChildRounding is a float.
StyleVarChildRounding = StyleVarID(imgui.StyleVarChildRounding) // child-rounding
// StyleVarChildBorderSize is a float.
StyleVarChildBorderSize = StyleVarID(imgui.StyleVarChildBorderSize) // child-border-size
// StyleVarPopupRounding is a float.
StyleVarPopupRounding = StyleVarID(imgui.StyleVarPopupRounding) // popup-rounding
// StyleVarPopupBorderSize is a float.
StyleVarPopupBorderSize = StyleVarID(imgui.StyleVarPopupBorderSize) // popup-border-size
// StyleVarFramePadding is a Vec2.
StyleVarFramePadding = StyleVarID(imgui.StyleVarFramePadding) // frame-padding
// StyleVarFrameRounding is a float.
StyleVarFrameRounding = StyleVarID(imgui.StyleVarFrameRounding) // frame-rounding
// StyleVarFrameBorderSize is a float.
StyleVarFrameBorderSize = StyleVarID(imgui.StyleVarFrameBorderSize) // frame-border-size
// StyleVarItemSpacing is a Vec2.
StyleVarItemSpacing = StyleVarID(imgui.StyleVarItemSpacing) // item-spacing
// StyleVarItemInnerSpacing is a Vec2.
StyleVarItemInnerSpacing = StyleVarID(imgui.StyleVarItemInnerSpacing) // item-inner-spacing
// StyleVarIndentSpacing is a float.
StyleVarIndentSpacing = StyleVarID(imgui.StyleVarIndentSpacing) // indent-spacing
// StyleVarScrollbarSize is a float.
StyleVarScrollbarSize = StyleVarID(imgui.StyleVarScrollbarSize) // scrollbar-size
// StyleVarScrollbarRounding is a float.
StyleVarScrollbarRounding = StyleVarID(imgui.StyleVarScrollbarRounding) // scrollbar-rounding
// StyleVarGrabMinSize is a float.
StyleVarGrabMinSize = StyleVarID(imgui.StyleVarGrabMinSize) // grab-min-size
// StyleVarGrabRounding is a float.
StyleVarGrabRounding = StyleVarID(imgui.StyleVarGrabRounding) // grab-rounding
// StyleVarTabRounding is a float.
StyleVarTabRounding = StyleVarID(imgui.StyleVarTabRounding) // tab-rounding
// StyleVarButtonTextAlign is a Vec2.
StyleVarButtonTextAlign = StyleVarID(imgui.StyleVarButtonTextAlign) // button-text-align
// StyleVarSelectableTextAlign is a Vec2.
StyleVarSelectableTextAlign = StyleVarID(imgui.StyleVarSelectableTextAlign) // selectable-text-align
