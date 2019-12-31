package giu

import (
	"github.com/AllenDang/giu/imgui"
)

type InputTextWidget struct {
	BaseWidget
	label   string
	value   *string
	flags   int
	cb      imgui.InputTextCallback
	changed func()
}

func InputTextV(label string, value *string, width float32, flags int, cb imgui.InputTextCallback, changed func()) *InputTextWidget {
	return &InputTextWidget{
		BaseWidget: BaseWidget{width},
		label:      label,
		value:      value,
		flags:      flags,
		cb:         cb,
		changed:    changed,
	}
}

func InputText(label string, value *string, changed func()) *InputTextWidget {
	return InputTextV(label, value, 0, 0, nil, changed)
}

func (t *InputTextWidget) Build() {
	if imgui.InputTextV(t.label, t.value, t.flags, t.cb) && t.changed != nil {
		t.changed()
	}
}
