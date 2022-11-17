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

<!--
and regex here are:
81,128s/\/\/.*\(float\|Vec2\).*\nStyleVar\(\w\+\).*\/\/ \(.*\)/- `\3` - \2 (\1)/g
87,110s/\([A-Z][a-z]\+\)/\1 /g
-->

- `alpha` - Alpha  (float)
- `disabled-alpha` - Disabled Alpha  (float)
- `window-padding` - Window Padding  (Vec 2)
- `window-rounding` - Window Rounding  (float)
- `window-border-size` - Window Border Size  (float)
- `window-min-size` - Window Min Size  (Vec 2)
- `window-title-align` - Window Title Align  (Vec 2)
- `child-rounding` - Child Rounding  (float)
- `child-border-size` - Child Border Size  (float)
- `popup-rounding` - Popup Rounding  (float)
- `popup-border-size` - Popup Border Size  (float)
- `frame-padding` - Frame Padding  (Vec 2)
- `frame-rounding` - Frame Rounding  (float)
- `frame-border-size` - Frame Border Size  (float)
- `item-spacing` - Item Spacing  (Vec 2)
- `item-inner-spacing` - Item Inner Spacing  (Vec 2)
- `indent-spacing` - Indent Spacing  (float)
- `scrollbar-size` - Scrollbar Size  (float)
- `scrollbar-rounding` - Scrollbar Rounding  (float)
- `grab-min-size` - Grab Min Size  (float)
- `grab-rounding` - Grab Rounding  (float)
- `tab-rounding` - Tab Rounding  (float)
- `button-text-align` - Button Text Align  (Vec 2)
- `selectable-text-align` - Selectable Text Align  (Vec 2)

# Data types

- color - supported types are:
	* Named colors (e.g. red, yellow, e.t.c.)
	* `rgb()` and `rgba()`
	* `hsl()` and `hsla()`
	* `hwb()` and `hwba()`
	* `hsv()` and `hsva()`
	* for more details about colors parsing visit [this repository](https://github.com/mazznoer/csscolorparser)
- float in form of plain number
- Vec2 - set of **exactly two** numbers, first for X and second for Y

## example

```css
main {
	color: rgba(50, 100, 150, 255);
	background-color: yellow;
	alpha: 100;
	item-spacing: 80, 20;
}
```

# special tags

CSS widget supports a **special tag** called `main`.
This tag is automatically applied for the whole app and you don't need
to perform any additional actions to add it.
<ins>There is **no** need to call `giu.CSS("main")`</ins>

# limitations

- be careful with CSSS comments since they may not be parsed correctly :smile:
  e.g. comments inside tags may not be supported, so if you need
  comment out the whole tag, However feel free to play with that.
- more complex rules are not supported, just the simpliest one (like ruleName {...})
