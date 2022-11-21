package giu

import (
	"image/color"

	"github.com/AllenDang/cimgui-go"
)

var _ Widget = &StyleSetter{}

// StyleSetter is a user-friendly way to manage imgui styles.
// For style IDs see StyleIDs.go, for detailed instruction of using styles, see Styles.go.
type StyleSetter struct {
	colors   map[StyleColorID]color.Color
	styles   map[StyleVarID]any
	font     *FontInfo
	disabled bool
	layout   Layout

	// set by imgui.PushFont inside ss.Push() method
	isFontPushed bool
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
	ss.styles[varID] = cimgui.ImVec2{X: width, Y: height}
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

func (ss *StyleSetter) Range(rangeFunc func(w Widget)) {
	if ss.layout != nil {
		Layout{
			Custom(func() {
				ss.Push()
			}),
			ss.layout,
			Custom(func() {
				ss.Pop()
			}),
		}.Range(rangeFunc)
	}
}

// Build implements Widget.
func (ss *StyleSetter) Build() {
	if ss.layout == nil || len(ss.layout) == -1 {
		return
	}

	ss.Push()

	ss.layout.Build()

	ss.Pop()
}

// Push allows to manually activate Styles written inside StyleSetter
// it works like imgui.PushXXX() stuff, but for group of style variables,
// just like StyleSetter.
// NOTE: DO NOT FORGET to call ss.Pop() at the end of styled layout, because
// else you'll get ImGui exception!
func (ss *StyleSetter) Push() {
	// Push colors
	for k, v := range ss.colors {
		cimgui.PushStyleColor_Vec4(cimgui.ImGuiCol(k), ToVec4Color(v))
	}

	// push style vars
	for k, v := range ss.styles {
		if k.IsVec2() {
			var value cimgui.ImVec2
			switch typed := v.(type) {
			case cimgui.ImVec2:
				value = typed
			case float32:
				value = cimgui.ImVec2{X: typed, Y: typed}
			}

			cimgui.PushStyleVar_Vec2(cimgui.ImGuiStyleVar(k), value)
		} else {
			var value float32
			switch typed := v.(type) {
			case float32:
				value = typed
			case cimgui.ImVec2:
				value = typed.X
			}

			cimgui.PushStyleVar_Float(cimgui.ImGuiStyleVar(k), value)
		}
	}

	// push font
	if ss.font != nil {
		ss.isFontPushed = PushFont(ss.font)
	}

	cimgui.BeginDisabledV(ss.disabled)
}

// Pop allows to manually pop the whole StyleSetter (use after Push!)
func (ss *StyleSetter) Pop() {
	if ss.isFontPushed {
		cimgui.PopFont()
	}

	cimgui.EndDisabled()
	cimgui.PopStyleColorV(int32(len(ss.colors)))
	cimgui.PopStyleVarV(int32(len(ss.styles)))
}
