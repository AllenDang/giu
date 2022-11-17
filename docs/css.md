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

<!--
Here is my regex, I used to convert variables from StyleIDs.go
18,71s/StyleColor\(\w\+\) \+= .*\/\/ \(.*\)/- `\2` - \1/g
18,71s/\([A-Z][a-z]\+\)/\1 /g
-->

- `color` - Text 
- `disabled-color` - Text Disabled 
- `background-color` - Window Bg 
- `child-background-color` - Child Bg 
- `popup-background-color` - Popup Bg 
- `border-color` - Border 
- `border-shadow-color` - Border Shadow 
- `frame-background-color` - Frame Bg 
- `frame-background-hovered-color` - Frame Bg Hovered 
- `frame-background-active-color` - Frame Bg Active 
- `title-background-color` - Title Bg 
- `title-background-active-color` - Title Bg Active 
- `title-background-collapsed-color` - Title Bg Collapsed 
- `menu-bar-background-color` - Menu Bar Bg 
- `scrollbar-background-color` - Scrollbar Bg 
- `scrollbar-grab-color` - Scrollbar Grab 
- `scrollbar-grab-hovered-color` - Scrollbar Grab Hovered 
- `scrollbar-grab-active-color` - Scrollbar Grab Active 
- `checkmark-color` - Check Mark 
- `slider-grab-color` - Slider Grab 
- `slider-grab-active-color` - Slider Grab Active 
- `button-color` - Button 
- `button-hovered-color` - Button Hovered 
- `button-active-color` - Button Active 
- `header-color` - Header 
- `header-hovered-color` - Header Hovered 
- `header-active-color` - Header Active 
- `separator-color` - Separator 
- `separator-hovered-color` - Separator Hovered 
- `separator-active-color` - Separator Active 
- `resize-grip-color` - Resize Grip 
- `resize-grip-hovered-color` - Resize Grip Hovered 
- `resize-grip-active-color` - Resize Grip Active 
- `tab-color` - Tab 
- `tab-hovered-color` - Tab Hovered 
- `tab-active-color` - Tab Active 
- `tab-unfocused-color` - Tab Unfocused 
- `tab-unfocused-active-color` - Tab Unfocused Active 
- `plot-lines-color` - Plot Lines 
- `plot-lines-hovered-color` - Plot Lines Hovered 
- `progress-bar-active-color` - Progress Bar Active 
- `plot-histogram-color` - Plot Histogram 
- `plot-histogram-hovered-color` - Plot Histogram Hovered 
- `table-header-background-color` - Table Header Bg 
- `table-border-strong-color` - Table Border Strong 
- `table-border-light-color` - Table Border Light 
- `table-row-background-color` - Table Row Bg 
- `table-row-alternate-background-color` - Table Row Bg Alt 
- `text-selected-background-color` - Text Selected Bg 
- `drag-drop-target-color` - Drag Drop Target 
- `navigation-highlight-color` - Nav Highlight 
- `windowing-highlight-color` - Nav Windowing Highlight 
- `windowing-dim-background-color` - Nav Windowing Dim Bg 
- `modal-window-dim-background-color` - Modal Window Dim Bg 

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
