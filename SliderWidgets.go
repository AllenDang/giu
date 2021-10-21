package giu

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
)

var _ Widget = &SliderIntWidget{}

type SliderIntWidget struct {
	label    string
	value    *int32
	min      int32
	max      int32
	format   string
	width    float32
	onChange func()
}

func SliderInt(value *int32, min, max int32) *SliderIntWidget {
	return &SliderIntWidget{
		label:    GenAutoID("##SliderInt"),
		value:    value,
		min:      min,
		max:      max,
		format:   "%d",
		width:    0,
		onChange: nil,
	}
}

func (s *SliderIntWidget) Format(format string) *SliderIntWidget {
	s.format = format
	return s
}

func (s *SliderIntWidget) Size(width float32) *SliderIntWidget {
	s.width = width
	return s
}

func (s *SliderIntWidget) OnChange(onChange func()) *SliderIntWidget {
	s.onChange = onChange

	return s
}

func (s *SliderIntWidget) Label(label string) *SliderIntWidget {
	s.label = Context.FontAtlas.tStr(label)
	return s
}

func (s *SliderIntWidget) Labelf(format string, args ...interface{}) *SliderIntWidget {
	return s.Label(fmt.Sprintf(format, args...))
}

// Build implements Widget interface.
func (s *SliderIntWidget) Build() {
	if s.width != 0 {
		PushItemWidth(s.width)
		defer PopItemWidth()
	}

	if imgui.SliderIntV(Context.FontAtlas.tStr(s.label), s.value, s.min, s.max, s.format) && s.onChange != nil {
		s.onChange()
	}
}

var _ Widget = &VSliderIntWidget{}

type VSliderIntWidget struct {
	label    string
	width    float32
	height   float32
	value    *int32
	min      int32
	max      int32
	format   string
	flags    SliderFlags
	onChange func()
}

func VSliderInt(value *int32, min, max int32) *VSliderIntWidget {
	return &VSliderIntWidget{
		label:  GenAutoID("##VSliderInt"),
		width:  18,
		height: 60,
		value:  value,
		min:    min,
		max:    max,
		format: "%d",
		flags:  SliderFlagsNone,
	}
}

func (vs *VSliderIntWidget) Size(width, height float32) *VSliderIntWidget {
	vs.width, vs.height = width, height
	return vs
}

func (vs *VSliderIntWidget) Flags(flags SliderFlags) *VSliderIntWidget {
	vs.flags = flags
	return vs
}

func (vs *VSliderIntWidget) Format(format string) *VSliderIntWidget {
	vs.format = format
	return vs
}

func (vs *VSliderIntWidget) OnChange(onChange func()) *VSliderIntWidget {
	vs.onChange = onChange
	return vs
}

func (vs *VSliderIntWidget) Label(label string) *VSliderIntWidget {
	vs.label = Context.FontAtlas.tStr(label)
	return vs
}

func (vs *VSliderIntWidget) Labelf(format string, args ...interface{}) *VSliderIntWidget {
	return vs.Label(fmt.Sprintf(format, args...))
}

// Build implements Widget interface.
func (vs *VSliderIntWidget) Build() {
	if imgui.VSliderIntV(
		Context.FontAtlas.tStr(vs.label),
		imgui.Vec2{X: vs.width, Y: vs.height},
		vs.value,
		vs.min,
		vs.max,
		vs.format,
		int(vs.flags),
	) && vs.onChange != nil {
		vs.onChange()
	}
}

var _ Widget = &SliderFloatWidget{}

type SliderFloatWidget struct {
	label    string
	value    *float32
	min      float32
	max      float32
	format   string
	width    float32
	onChange func()
}

func SliderFloat(value *float32, min, max float32) *SliderFloatWidget {
	return &SliderFloatWidget{
		label:    GenAutoID("##SliderFloat"),
		value:    value,
		min:      min,
		max:      max,
		format:   "%.3f",
		width:    0,
		onChange: nil,
	}
}

func (sf *SliderFloatWidget) Format(format string) *SliderFloatWidget {
	sf.format = format
	return sf
}

func (sf *SliderFloatWidget) OnChange(onChange func()) *SliderFloatWidget {
	sf.onChange = onChange

	return sf
}

func (sf *SliderFloatWidget) Size(width float32) *SliderFloatWidget {
	sf.width = width
	return sf
}

func (sf *SliderFloatWidget) Label(label string) *SliderFloatWidget {
	sf.label = Context.FontAtlas.tStr(label)
	return sf
}

func (sf *SliderFloatWidget) Labelf(format string, args ...interface{}) *SliderFloatWidget {
	return sf.Label(fmt.Sprintf(format, args...))
}

// Build implements Widget interface.
func (sf *SliderFloatWidget) Build() {
	if sf.width != 0 {
		PushItemWidth(sf.width)
		defer PopItemWidth()
	}

	if imgui.SliderFloatV(Context.FontAtlas.tStr(sf.label), sf.value, sf.min, sf.max, sf.format, 1.0) && sf.onChange != nil {
		sf.onChange()
	}
}
