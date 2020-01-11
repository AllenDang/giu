package giu

import "github.com/AllenDang/giu/imgui"

type HSplitterWidget struct {
	BaseWidget
	id     string
	width  float32
	height float32
	delta  *float32
}

func HSplitter(id string, width, height float32, delta *float32) *HSplitterWidget {
	return &HSplitterWidget{
		BaseWidget: BaseWidget{width: -1},
		id:         id,
		width:      width,
		height:     height,
		delta:      delta,
	}
}

func (h *HSplitterWidget) Build() {
	imgui.InvisibleButton(h.id, imgui.Vec2{X: h.width, Y: h.height})
	if imgui.IsItemActive() {
		*(h.delta) = imgui.CurrentIO().GetMouseDelta().Y
	} else {
		*(h.delta) = 0
	}
}

type VSplitterWidget struct {
	BaseWidget
	id     string
	width  float32
	height float32
	delta  *float32
}

func VSplitter(id string, width, height float32, delta *float32) *VSplitterWidget {
	return &VSplitterWidget{
		BaseWidget: BaseWidget{width: -1},
		id:         id,
		width:      width,
		height:     height,
		delta:      delta,
	}
}

func (v *VSplitterWidget) Build() {
	imgui.InvisibleButton(v.id, imgui.Vec2{X: v.width, Y: v.height})
	if imgui.IsItemActive() {
		*(v.delta) = imgui.CurrentIO().GetMouseDelta().X
	} else {
		*(v.delta) = 0
	}
}
