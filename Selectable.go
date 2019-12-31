package giu

import "github.com/AllenDang/giu/imgui"

type SelectableWidget struct {
	BaseWidget
	label    string
	selected bool
	flags    int
	size     imgui.Vec2
	clicked  func()
}

func SelectableV(label string, selected bool, flags int, size imgui.Vec2, clicked func()) *SelectableWidget {
	return &SelectableWidget{
		BaseWidget: BaseWidget{width: 0},
		label:      label,
		selected:   selected,
		flags:      flags,
		size:       size,
		clicked:    clicked,
	}
}

func Selectable(label string, clicked func()) *SelectableWidget {
	return SelectableV(label, false, 0, imgui.Vec2{}, clicked)
}

func (s *SelectableWidget) Build() {
	if imgui.SelectableV(s.label, s.selected, s.flags, s.size) && s.clicked != nil {
		s.clicked()
	}
}
