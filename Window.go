package giu

import (
	"github.com/AllenDang/giu/imgui"
)

func Window(title string, x, y, width, height float32, layout Layout) {
	WindowV(
		title,
		nil,
		0,
		x, y,
		width, height,
		layout,
	)
}

func SingleWindow(title string, layout Layout) {
	size := Context.platform.DisplaySize()
	WindowV(
		title,
		nil,
		imgui.WindowFlagsNoTitleBar|
			imgui.WindowFlagsNoCollapse|
			imgui.WindowFlagsNoScrollbar|
			imgui.WindowFlagsNoMove|
			imgui.WindowFlagsNoResize,
		0, 0,
		size[0], size[1],
		layout,
	)
}

func SingleWindowWithMenuBar(title string, layout Layout) {
	size := Context.platform.DisplaySize()
	WindowV(
		title,
		nil,
		imgui.WindowFlagsNoTitleBar|
			imgui.WindowFlagsNoCollapse|
			imgui.WindowFlagsNoScrollbar|
			imgui.WindowFlagsNoMove|
			imgui.WindowFlagsMenuBar|
			imgui.WindowFlagsNoResize,
		0, 0,
		size[0], size[1],
		layout,
	)
}

func WindowV(title string, open *bool, flags WindowFlags, x, y, width, height float32, layout Layout) {
	if flags&imgui.WindowFlagsNoMove != 0 && flags&imgui.WindowFlagsNoResize != 0 {
		imgui.SetNextWindowPos(imgui.Vec2{X: x, Y: y})
		imgui.SetNextWindowSize(imgui.Vec2{X: width, Y: height})
	} else {
		imgui.SetNextWindowPosV(imgui.Vec2{X: x, Y: y}, imgui.ConditionFirstUseEver, imgui.Vec2{X: 0, Y: 0})
		imgui.SetNextWindowSizeV(imgui.Vec2{X: width, Y: height}, imgui.ConditionFirstUseEver)
	}

	imgui.BeginV(title, open, int(flags))

	// Mark all state as invalid.
	Context.invalidAllState()

	layout.Build()

	// Clean remaining invalid states
	Context.cleanState()

	imgui.End()
}
