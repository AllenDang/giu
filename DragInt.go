package giu

import "github.com/AllenDang/giu/imgui"

type DragIntWidget struct {
	BaseWidget
	label  string
	value  *int32
	speed  float32
	min    int32
	max    int32
	format string
}

func DragIntV(label string, value *int32, speed float32, min, max int32, format string, width float32) *DragIntWidget {
	return &DragIntWidget{
		BaseWidget: BaseWidget{width: width},
		label:      label,
		value:      value,
		speed:      speed,
		min:        min,
		max:        max,
		format:     format,
	}
}

func DragInt(label string, value *int32) *DragIntWidget {
	return DragIntV(label, value, 1.0, 0, 0, "%d", 0)
}

func (d *DragIntWidget) Build() {
	imgui.DragIntV(d.label, d.value, d.speed, d.min, d.max, d.format)
}
