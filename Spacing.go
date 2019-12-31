package giu

import "github.com/AllenDang/giu/imgui"

type SpacingWidget struct {
	BaseWidget
}

func Spacing() *SpacingWidget {
	return &SpacingWidget{
		BaseWidget: BaseWidget{width: 0},
	}
}

func (s *SpacingWidget) Build() {
	imgui.Spacing()
}

type DummyWidget struct {
	BaseWidget
	size imgui.Vec2
}

func Dummy(size imgui.Vec2) *DummyWidget {
	return &DummyWidget{
		BaseWidget: BaseWidget{width: 0},
		size:       size,
	}
}

func (d *DummyWidget) Build() {
	imgui.Dummy(d.size)
}
