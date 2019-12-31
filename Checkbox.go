package giu

import "github.com/AllenDang/giu/imgui"

type CheckboxWidget struct {
	BaseWidget
	text     string
	selected *bool
	changed  func()
}

func Checkbox(text string, selected *bool, changed func()) *CheckboxWidget {
	return CheckboxV(text, selected, changed, 0)
}

func CheckboxV(text string, selected *bool, changed func(), width float32) *CheckboxWidget {
	return &CheckboxWidget{
		BaseWidget: BaseWidget{width: width},
		text:       text,
		selected:   selected,
		changed:    changed,
	}
}

func (c *CheckboxWidget) Build() {
	if imgui.Checkbox(c.text, c.selected) && c.changed != nil {
		c.changed()
	}
}
