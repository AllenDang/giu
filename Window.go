package giu

import (
	"github.com/AllenDang/giu/imgui"
)

func Window(title string, x, y, width, height float32, widgets ...Widget) {
	WindowV(
		title,
		nil,
		imgui.WindowFlagsNoCollapse|
			imgui.WindowFlagsNoMove|
			imgui.WindowFlagsNoResize,
		x, y,
		width, height,
		widgets...,
	)
}

func SingleWindow(w *MasterWindow, title string, widgets ...Widget) {
	width, height := w.GetSize()
	WindowV(
		title,
		nil,
		imgui.WindowFlagsNoTitleBar|
			imgui.WindowFlagsNoBackground|
			imgui.WindowFlagsNoCollapse|
			imgui.WindowFlagsNoScrollbar|
			imgui.WindowFlagsNoMove|
			imgui.WindowFlagsNoResize,
		0, 0,
		float32(width), float32(height),
		widgets...,
	)
}

func WindowV(title string, open *bool, flags int, x, y, width, height float32, widgets ...Widget) {
	imgui.SetNextWindowPos(imgui.Vec2{X: x, Y: y})
	imgui.SetNextWindowSize(imgui.Vec2{X: width, Y: height})

	imgui.BeginV(title, open, flags)
	Layout(widgets...)()
	imgui.End()
}
