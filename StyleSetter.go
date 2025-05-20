package giu

import (
	"image/color"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/implot"
)

var _ Widget = &StyleSetter{}

// StyleSetter is a user-friendly way to manage imgui styles.
// For style IDs see StyleIDs.go, for detailed instruction of using styles, see Styles.go.
type StyleSetter struct {
	colors     map[StyleColorID]imgui.Vec4
	styles     map[StyleVarID]any
	plotColors map[StylePlotColorID]color.Color
	plotStyles map[StylePlotVarID]any
	font       *FontInfo
	disabled   bool

	layout Layout
	plots  []PlotWidget

	// set by imgui.PushFont inside ss.Push() method
	isFontPushed bool
}

// Style initializes a style setter (see examples/setstyle).
func Style() *StyleSetter {
	var ss StyleSetter
	ss.colors = make(map[StyleColorID]imgui.Vec4)
	ss.plotColors = make(map[StylePlotColorID]color.Color)
	ss.styles = make(map[StyleVarID]any)
	ss.plotStyles = make(map[StylePlotVarID]any)

	return &ss
}

// Add merges two StyleSetters.
// Add puts other "on top" of ss, meaning, "other" is applied after "ss".
// e.g. if both StyleSetters set imgui.StyleVarAlpha, the value from "other" will be used.
// NOTE: font value "nil" is treated as "not set" and will not be changed if declared by other.
// NOTE: true is preffered over false for disabled field.
// NOTE: layout field will be reset.
func (ss *StyleSetter) Add(other *StyleSetter) *StyleSetter {
	if other == nil {
		return ss
	}

	for k, v := range other.colors {
		ss.colors[k] = v
	}

	for k, v := range other.styles {
		ss.styles[k] = v
	}

	for k, v := range other.plotColors {
		ss.plotColors[k] = v
	}

	for k, v := range other.plotStyles {
		ss.plotStyles[k] = v
	}

	if other.font != nil {
		ss.font = other.font
	}

	if other.disabled {
		ss.disabled = true
	}

	ss.layout = nil

	return ss
}

// SetColor sets colorID's color.
func (ss *StyleSetter) SetColor(colorID StyleColorID, col color.Color) *StyleSetter {
	ss.colors[colorID] = ToVec4Color(col)
	return ss
}

// SetColorVec4 is a lower-level function to set colors.
// It omits color conversion for e.g. better performance/compatibility.
// Historically was introduced to easily convert DefaulutTheme from using imgui api to StyleSetter.
func (ss *StyleSetter) SetColorVec4(colorID StyleColorID, col imgui.Vec4) *StyleSetter {
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

// SetPlotColor sets colorID's color.
func (ss *StyleSetter) SetPlotColor(colorID StylePlotColorID, col color.Color) *StyleSetter {
	ss.plotColors[colorID] = col
	return ss
}

// SetPlotStyle sets stylePlotVarID to width and height.
func (ss *StyleSetter) SetPlotStyle(varID StylePlotVarID, width, height float32) *StyleSetter {
	ss.plotStyles[varID] = imgui.Vec2{X: width, Y: height}
	return ss
}

// SetPlotStyleFloat sets StylePlotVarID to float value.
// NOTE: for float typed values see above in comments over
// StyleVarID's comments.
func (ss *StyleSetter) SetPlotStyleFloat(varID StylePlotVarID, value float32) *StyleSetter {
	ss.plotStyles[varID] = value
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

// Plots allows to set plots to apply style for.
func (ss *StyleSetter) Plots(widgets ...PlotWidget) *StyleSetter {
	ss.plots = widgets
	return ss
}

// Range implements Splitable interface.
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
	if len(ss.layout) == 0 {
		return
	}

	ss.Push()

	ss.layout.Build()

	ss.Pop()
}

// Plot implements PlotWidget.
func (ss *StyleSetter) Plot() {
	if len(ss.plots) == 0 {
		return
	}

	ss.Push()

	for _, plot := range ss.plots {
		plot.Plot()
	}

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
		imgui.PushStyleColorVec4(imgui.Col(k), v)
	}

	// Push plot colors
	for k, v := range ss.plotColors {
		implot.PushStyleColorVec4(implot.Col(k), ToVec4Color(v))
	}

	// push style vars
	for k, v := range ss.styles {
		pushVarID(k.IsVec2(), v, func(value float32) {
			imgui.PushStyleVarFloat(imgui.StyleVar(k), value)
		}, func(value imgui.Vec2) {
			imgui.PushStyleVarVec2(imgui.StyleVar(k), value)
		})
	}

	// Push plot colors
	for k, v := range ss.plotStyles {
		pushVarID(k.IsVec2(), v, func(value float32) {
			implot.PushStyleVarFloat(implot.StyleVar(k), value)
		}, func(value imgui.Vec2) {
			implot.PushStyleVarVec2(implot.StyleVar(k), value)
		})
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
	implot.PopStyleColorV(int32(len(ss.plotColors)))
	imgui.PopStyleVarV(int32(len(ss.styles)))
	implot.PopStyleVarV(int32(len(ss.plotStyles)))
}

func pushVarID(isVec2 bool, v any, pushFloat func(float32), pushVec2 func(imgui.Vec2)) {
	if isVec2 {
		var value imgui.Vec2
		switch typed := v.(type) {
		case imgui.Vec2:
			value = typed
		case float32:
			value = imgui.Vec2{X: typed, Y: typed}
		}

		pushVec2(value)
	} else {
		var value float32
		switch typed := v.(type) {
		case float32:
			value = typed
		case imgui.Vec2:
			value = typed.X
		}

		pushFloat(value)
	}
}
