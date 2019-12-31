package giu

import "github.com/AllenDang/giu/imgui"

type SliderIntWidget struct {
	BaseWidget
	label  string
	value  *int32
	min    int32
	max    int32
	format string
}

func SliderIntV(label string, value *int32, min, max int32, format string, width float32) *SliderIntWidget {
	return &SliderIntWidget{
		BaseWidget: BaseWidget{width: width},
		label:      label,
		value:      value,
		min:        min,
		max:        max,
		format:     format,
	}
}

func SliderInt(label string, value *int32, min, max int32) *SliderIntWidget {
	return SliderIntV(label, value, min, max, "%d", 0)
}

func (s *SliderIntWidget) Build() {
	imgui.SliderIntV(s.label, s.value, s.min, s.max, s.format)
}
