package giu

import "github.com/AllenDang/giu/imgui"

type ChildWidget struct {
	BaseWidget
	id     string
	size   imgui.Vec2
	border bool
	flags  int
	layout Layout
}

func Child(id string, width, height float32, border bool, flags int, layout Layout) *ChildWidget {
	return &ChildWidget{
		BaseWidget: BaseWidget{width: 0},
		id:         id,
		size:       imgui.Vec2{X: width, Y: height},
		border:     border,
		flags:      flags,
		layout:     layout,
	}
}

func (c *ChildWidget) Build() {
	imgui.BeginChildV(c.id, c.size, c.border, c.flags)
	c.layout.Build()
	imgui.EndChild()
}
