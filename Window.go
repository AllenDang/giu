package giu

import (
	"github.com/AllenDang/giu/imgui"
)

func Window(title string, x, y, width, height float32, builder Builder) {
	WindowV(
		title,
		nil,
		imgui.WindowFlagsNoCollapse|
			imgui.WindowFlagsNoMove|
			imgui.WindowFlagsNoResize,
		x, y,
		width, height,
		builder,
	)
}

func SingleWindow(w *MasterWindow, title string, builder Builder) {
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
		builder,
	)
}

func WindowV(title string, open *bool, flags int, x, y, width, height float32, builder Builder) {
	imgui.SetNextWindowPos(imgui.Vec2{X: x, Y: y})
	imgui.SetNextWindowSize(imgui.Vec2{X: width, Y: height})

	imgui.BeginV(title, open, flags)
	builder()
	imgui.End()
}
