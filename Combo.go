package giu

import "github.com/AllenDang/giu/imgui"

type ComboWidget struct {
	BaseWidget
	label        string
	previewValue string
	selected     *int32
	items        []string
	changed      func()
}

func ComboV(label, previewValue string, items []string, selected *int32, width float32, changed func()) *ComboWidget {
	return &ComboWidget{
		BaseWidget:   BaseWidget{width: width},
		label:        label,
		previewValue: previewValue,
		selected:     selected,
		items:        items,
		changed:      changed,
	}
}

func Combo(label string, items []string, selected *int32, changed func()) *ComboWidget {
	previewValue := ""
	if len(items) > 0 {
		previewValue = items[0]
	}

	return ComboV(label, previewValue, items, selected, 0, changed)
}

func (c *ComboWidget) Build() {
	if imgui.BeginCombo(c.label, c.previewValue) {
		for i, item := range c.items {
			if imgui.Selectable(item) {
				*c.selected = int32(i)
				if c.changed != nil {
					c.changed()
				}
			}
		}

		imgui.EndCombo()
	}
}
