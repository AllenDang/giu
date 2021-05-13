package giu

import (
	"github.com/AllenDang/imgui-go"
)

func SingleWindow(title string) *WindowWidget {
	size := Context.platform.DisplaySize()
	return Window(title).
		Flags(
			imgui.WindowFlagsNoTitleBar|
				imgui.WindowFlagsNoCollapse|
				imgui.WindowFlagsNoScrollbar|
				imgui.WindowFlagsNoMove|
				imgui.WindowFlagsNoResize).
		Size(size[0], size[1])
}

func SingleWindowWithMenuBar(title string) *WindowWidget {
	size := Context.platform.DisplaySize()
	return Window(title).
		Flags(
			imgui.WindowFlagsNoTitleBar|
				imgui.WindowFlagsNoCollapse|
				imgui.WindowFlagsNoScrollbar|
				imgui.WindowFlagsNoMove|
				imgui.WindowFlagsMenuBar|
				imgui.WindowFlagsNoResize).Size(size[0], size[1])
}

type WindowWidget struct {
	title         string
	open          *bool
	flags         WindowFlags
	x, y          float32
	width, height float32
}

func Window(title string) *WindowWidget {
	return &WindowWidget{
		title: title,
	}
}

func (w *WindowWidget) IsOpen(open *bool) *WindowWidget {
	w.open = open
	return w
}

func (w *WindowWidget) Flags(flags WindowFlags) *WindowWidget {
	w.flags = flags
	return w
}

func (w *WindowWidget) Size(width, height float32) *WindowWidget {
	w.width, w.height = width, height
	return w
}

func (w *WindowWidget) Pos(x, y float32) *WindowWidget {
	w.x, w.y = x, y
	return w
}

func (w *WindowWidget) Layout(widgets ...Widget) {
	if widgets == nil {
		return
	}

	if w.flags&imgui.WindowFlagsNoMove != 0 && w.flags&imgui.WindowFlagsNoResize != 0 {
		imgui.SetNextWindowPos(imgui.Vec2{X: w.x, Y: w.y})
		imgui.SetNextWindowSize(imgui.Vec2{X: w.width, Y: w.height})
	} else {
		imgui.SetNextWindowPosV(imgui.Vec2{X: w.x, Y: w.y}, imgui.ConditionFirstUseEver, imgui.Vec2{X: 0, Y: 0})
		imgui.SetNextWindowSizeV(imgui.Vec2{X: w.width, Y: w.height}, imgui.ConditionFirstUseEver)
	}

	showed := imgui.BeginV(tStr(w.title), w.open, int(w.flags))

	if showed {
		Layout(widgets).Build()
	}

	imgui.End()
}
