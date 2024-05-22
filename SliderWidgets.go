package giu

import (
	"fmt"

	imgui "github.com/AllenDang/cimgui-go"
)

var _ Widget = &SliderIntWidget{}

// SliderIntWidget is a slider around int32 values.
type SliderIntWidget struct {
	label    ID
	value    *int32
	minValue int32
	maxValue int32
	format   string
	width    float32
	onChange func()
}

// SliderInt constructs new SliderIntWidget.
func SliderInt(value *int32, minValue, maxValue int32) *SliderIntWidget {
	return &SliderIntWidget{
		label:    GenAutoID("##SliderInt"),
		value:    value,
		minValue: minValue,
		maxValue: maxValue,
		format:   "%d",
		width:    0,
		onChange: nil,
	}
}

// Format sets data format displayed on the slider
// NOTE: on C side of imgui, it will be processed like:
// fmt.Sprintf(format, currentValue) so you can do e.g.
// SLiderInt(...).Format("My age is %d") and %d will be replaced with current value.
func (s *SliderIntWidget) Format(format string) *SliderIntWidget {
	s.format = format
	return s
}

// Size sets slider's width.
func (s *SliderIntWidget) Size(width float32) *SliderIntWidget {
	s.width = width
	return s
}

// OnChange sets callback when slider's position gets changed.
func (s *SliderIntWidget) OnChange(onChange func()) *SliderIntWidget {
	s.onChange = onChange
	return s
}

// Label sets slider label (id).
func (s *SliderIntWidget) Label(label string) *SliderIntWidget {
	s.label = GenAutoID(label)
	return s
}

// Labelf sets formatted label.
func (s *SliderIntWidget) Labelf(format string, args ...any) *SliderIntWidget {
	return s.Label(fmt.Sprintf(format, args...))
}

func (s *SliderIntWidget) ID(id ID) *SliderIntWidget {
	s.label = id
	return s
}

// Build implements Widget interface.
func (s *SliderIntWidget) Build() {
	if s.width != 0 {
		PushItemWidth(s.width)

		defer PopItemWidth()
	}

	if imgui.SliderIntV(Context.FontAtlas.RegisterString(s.label.String()), s.value, s.minValue, s.maxValue, s.format, 0) && s.onChange != nil {
		s.onChange()
	}
}

var _ Widget = &VSliderIntWidget{}

// VSliderIntWidget stands from Vertical SliderIntWidget.
type VSliderIntWidget struct {
	label    ID
	width    float32
	height   float32
	value    *int32
	minValue int32
	maxValue int32
	format   string
	flags    SliderFlags
	onChange func()
}

// VSliderInt creates new vslider int.
func VSliderInt(value *int32, minValue, maxValue int32) *VSliderIntWidget {
	return &VSliderIntWidget{
		label:    GenAutoID("##VSliderInt"),
		width:    18,
		height:   60,
		value:    value,
		minValue: minValue,
		maxValue: maxValue,
		format:   "%d",
		flags:    SliderFlagsNone,
	}
}

// Size sets slider's size.
func (vs *VSliderIntWidget) Size(width, height float32) *VSliderIntWidget {
	vs.width, vs.height = width, height
	return vs
}

// Flags sets flags.
func (vs *VSliderIntWidget) Flags(flags SliderFlags) *VSliderIntWidget {
	vs.flags = flags
	return vs
}

// Format sets format (see comment on (*SliderIntWidget).Format).
func (vs *VSliderIntWidget) Format(format string) *VSliderIntWidget {
	vs.format = format
	return vs
}

// OnChange sets callback called when slider's position gets changed.
func (vs *VSliderIntWidget) OnChange(onChange func()) *VSliderIntWidget {
	vs.onChange = onChange
	return vs
}

// Label sets slider's label (id).
func (vs *VSliderIntWidget) Label(label string) *VSliderIntWidget {
	vs.label = GenAutoID(label)
	return vs
}

// Labelf sets formatted label.
func (vs *VSliderIntWidget) Labelf(format string, args ...any) *VSliderIntWidget {
	return vs.Label(fmt.Sprintf(format, args...))
}

func (vs *VSliderIntWidget) ID(id ID) *VSliderIntWidget {
	vs.label = id
	return vs
}

// Build implements Widget interface.
func (vs *VSliderIntWidget) Build() {
	if imgui.VSliderIntV(
		Context.FontAtlas.RegisterString(vs.label.String()),
		imgui.Vec2{X: vs.width, Y: vs.height},
		vs.value,
		vs.minValue,
		vs.maxValue,
		vs.format,
		imgui.SliderFlags(vs.flags),
	) && vs.onChange != nil {
		vs.onChange()
	}
}

var _ Widget = &SliderFloatWidget{}

// SliderFloatWidget does similar to SliderIntWidget but slides around
// float32 values.
type SliderFloatWidget struct {
	label    ID
	value    *float32
	minValue float32
	maxValue float32
	format   string
	width    float32
	onChange func()
}

// SliderFloat creates new slider float widget.
func SliderFloat(value *float32, minValue, maxValue float32) *SliderFloatWidget {
	return &SliderFloatWidget{
		label:    GenAutoID("##SliderFloat"),
		value:    value,
		minValue: minValue,
		maxValue: maxValue,
		format:   "%.3f",
		width:    0,
		onChange: nil,
	}
}

// Format sets format of text displayed on the slider.
// default is %.3f.
func (sf *SliderFloatWidget) Format(format string) *SliderFloatWidget {
	sf.format = format
	return sf
}

// OnChange is callback called when slider's position gets changed.
func (sf *SliderFloatWidget) OnChange(onChange func()) *SliderFloatWidget {
	sf.onChange = onChange

	return sf
}

// Size sets slider's width.
func (sf *SliderFloatWidget) Size(width float32) *SliderFloatWidget {
	sf.width = width
	return sf
}

// Label sets slider's label (id).
func (sf *SliderFloatWidget) Label(label string) *SliderFloatWidget {
	sf.label = GenAutoID(Context.FontAtlas.RegisterString(label))
	return sf
}

// Labelf sets formatted label.
func (sf *SliderFloatWidget) Labelf(format string, args ...any) *SliderFloatWidget {
	return sf.Label(fmt.Sprintf(format, args...))
}

func (sf *SliderFloatWidget) ID(id ID) *SliderFloatWidget {
	sf.label = id
	return sf
}

// Build implements Widget interface.
func (sf *SliderFloatWidget) Build() {
	if sf.width != 0 {
		PushItemWidth(sf.width)

		defer PopItemWidth()
	}

	if imgui.SliderFloatV(Context.FontAtlas.RegisterString(sf.label.String()), sf.value, sf.minValue, sf.maxValue, sf.format, 1.0) && sf.onChange != nil {
		sf.onChange()
	}
}
