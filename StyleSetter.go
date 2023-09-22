package giu

import (
	"image/color"

	imgui "github.com/AllenDang/cimgui-go"
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

func (ss *StyleSetter) Range(rangeFunc func(w Widget)) {
	if ss.layout != nil {
		var result Layout

		// need to consider the following cases:
		// if 0 - return
		// if 1 - joing push/render/pop in one step
		// else - join push with first widget, than render another
		//        widgets and in the end render last widget with pop func
		// it is, because Push/Pop don't move cursor so
		// they doesn't exist for imgui in fact.
		// It leads to problemms with RowWidget
		//
		// see: https://github.com/AllenDang/giu/issues/619
		layoutLen := len(ss.layout)
		switch layoutLen {
		case 0:
			return
		case 1:
			result = Layout{
				Custom(func() {
					ss.Push()
					ss.layout.Build()
					ss.Pop()
				}),
			}
		default:
			result = Layout{
				Custom(func() {
					ss.Push()
					ss.layout[0].Build()
				}),
				ss.layout[1 : len(ss.layout)-1],
				Custom(func() {
					ss.layout[layoutLen-1].Build()
					ss.Pop()
				}),
			}
		}

		result.Range(rangeFunc)
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

// Push allows to manually activate Styles written inside of StyleSetter
// it works like imgui.PushXXX() stuff, but for group of style variables,
// just like StyleSetter.
// NOTE: DO NOT ORGET to call ss.Pop() at the end of styled layout, because
// else you'll get ImGui exception!
func (ss *StyleSetter) Push() {
	// Push colors
	for k, v := range ss.colors {
		imgui.PushStyleColorVec4(imgui.Col(k), ToVec4Color(v))
	}

	// push style vars
	for k, v := range ss.styles {
		if k.IsVec2() {
			var value imgui.Vec2
			switch typed := v.(type) {
			case imgui.Vec2:
				value = typed
			case float32:
				value = imgui.Vec2{X: typed, Y: typed}
			}

			imgui.PushStyleVarVec2(imgui.StyleVar(k), value)
		} else {
			var value float32
			switch typed := v.(type) {
			case float32:
				value = typed
			case imgui.Vec2:
				value = typed.X
			}

			imgui.PushStyleVarFloat(imgui.StyleVar(k), value)
		}
	}

	// push font
	if ss.font != nil {
		ss.isFontPushed = PushFont(ss.font)
	}

	if ss.disabled {
		imgui.BeginDisabled()
	}
}

// Pop allows to manually pop the whole StyleSetter (use after Push!)
func (ss *StyleSetter) Pop() {
	if ss.isFontPushed {
		imgui.PopFont()
	}

	if ss.disabled {
		imgui.EndDisabled()
	}

	imgui.PopStyleColorV(int32(len(ss.colors)))
	imgui.PopStyleVarV(int32(len(ss.styles)))
}
