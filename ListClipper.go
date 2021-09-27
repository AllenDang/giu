package giu

import (
	"github.com/AllenDang/imgui-go"
)

var _ Widget = &ListClipperWrapper{}

// ListClipperWrapper is a ImGuiListClipper implementation.
// it can be used to diplay a large, vertical list of items and
// avoid rendering them.
type ListClipperWrapper struct {
	layout Layout
}

func ListClipper() *ListClipperWrapper {
	return &ListClipperWrapper{}
}

func (l *ListClipperWrapper) Layout(layout ...Widget) *ListClipperWrapper {
	l.layout = layout
	return l
}

func (l *ListClipperWrapper) Build() {
	// read all the layout widgets and (eventually) split nested layouts
	var layout Layout
	l.layout.Range(func(w Widget) {
		layout = append(layout, w)
	})

	var clipper imgui.ListClipper
	clipper.Begin(len(layout))

	for clipper.Step() {
		for i := clipper.DisplayStart; i < clipper.DisplayEnd; i++ {
			layout[i].Build()
		}
	}

	clipper.End()
}
