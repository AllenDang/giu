package giu

import "github.com/AllenDang/giu/imgui"

type GroupWidget struct {
	BaseWidget
	layout Layout
}

func GroupV(width float32, layout Layout) *GroupWidget {
	return &GroupWidget{
		BaseWidget: BaseWidget{width: width},
		layout:     layout,
	}
}

func Group(layout Layout) *GroupWidget {
	return GroupV(0, layout)
}

func (g *GroupWidget) Build() {
	imgui.BeginGroup()
	g.layout.Build()
	imgui.EndGroup()
}
