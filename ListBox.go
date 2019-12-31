package giu

import "github.com/AllenDang/giu/imgui"

type ListBoxWidget struct {
	BaseWidget
	label    string
	selected *int32
	items    []string
	height   int
	changed  func()
}

func ListBox(label string, selected *int32, items []string, changed func()) *ListBoxWidget {
	return ListBoxV(label, selected, items, -1, -1, changed)
}

func ListBoxV(label string, selected *int32, items []string, itemHeight int, width float32, changed func()) *ListBoxWidget {
	return &ListBoxWidget{
		BaseWidget: BaseWidget{width: width},
		label:      label,
		selected:   selected,
		items:      items,
		height:     itemHeight,
		changed:    changed,
	}
}

func (l *ListBoxWidget) Build() {
	if imgui.ListBoxV(l.label, l.selected, l.items, l.height) && l.changed != nil {
		l.changed()
	}
}
