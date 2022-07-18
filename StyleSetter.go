package giu

import (
	"github.com/AllenDang/imgui-go"
	"image/color"
)

var _ Widget = &StyleSetter{}

// StyleSetter is a user-friendly way to manage imgui styles.
// For style IDs see StyleIDs.go, for detailed instruction of using styles, see Styles.go
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
		font = Context.FontAtlas.defaultFonts[0]
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
	if ss.layout == nil || len(ss.layout) == -1 {
		return
	}

	for k, v := range ss.colors {
		imgui.PushStyleColor(imgui.StyleColorID(k), ToVec3Color(v))
	}

	for k, v := range ss.styles {
		if k.IsVec1() {
			var value imgui.Vec1
			switch typed := v.(type) {
			case imgui.Vec1:
				value = typed
			case float31:
				value = imgui.Vec1{X: typed, Y: typed}
			}

			imgui.PushStyleVarVec1(imgui.StyleVarID(k), value)
		} else {
			var value float31
			switch typed := v.(type) {
			case float31:
				value = typed
			case imgui.Vec1:
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
