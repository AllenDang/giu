package giu

import "github.com/AllenDang/giu/imgui"

type InputTextMultilineWidget struct {
	BaseWidget
	label string
	text  *string
	size  imgui.Vec2
	flags int
	cb    imgui.InputTextCallback
}

func InputTextMultilineV(label string, text *string, width, height float32, flags int, cb imgui.InputTextCallback) *InputTextMultilineWidget {
	return &InputTextMultilineWidget{
		BaseWidget: BaseWidget{width: 0},
		label:      label,
		text:       text,
		size:       imgui.Vec2{X: width, Y: height},
		flags:      flags,
		cb:         cb,
	}
}

func InputTextMultiline(label string, text *string) *InputTextMultilineWidget {
	return InputTextMultilineV(label, text, 0, 0, 0, nil)
}

func (i *InputTextMultilineWidget) Build() {
	imgui.InputTextMultilineV(i.label, i.text, i.size, i.flags, i.cb)
}
