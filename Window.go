package giu

import (
	"github.com/AllenDang/imgui-go"
)

func SingleWindow(title string) *WindowWidget {
	size := Context.platform.DisplaySize()
	return Window(title).
		Flags(
			WindowFlagsNoTitleBar|
				WindowFlagsNoCollapse|
				WindowFlagsNoScrollbar|
				WindowFlagsNoMove|
				WindowFlagsNoResize).
		Size(size[0], size[1])
}

func SingleWindowWithMenuBar(title string) *WindowWidget {
	size := Context.platform.DisplaySize()
	return Window(title).
		Flags(
			WindowFlagsNoTitleBar|
				WindowFlagsNoCollapse|
				WindowFlagsNoScrollbar|
				WindowFlagsNoMove|
				WindowFlagsMenuBar|
				WindowFlagsNoResize).Size(size[0], size[1])
}

type WindowWidget struct {
	title         string
	open          *bool
	flags         WindowFlags
	x, y          float32
	width, height float32
	hasFocus      bool
	bringToFront  bool
	layout        Layout
	pos           imgui.Vec2
	size          imgui.Vec2
}

func Window(title string) *WindowWidget {
	open := true

	return &WindowWidget{
		title: title,
		open:  &open,
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

func (w *WindowWidget) Layout(widgets ...Widget) *WindowWidget {
	if widgets == nil {
		return w
	}

	if w.flags&WindowFlagsNoMove != 0 && w.flags&WindowFlagsNoResize != 0 {
		imgui.SetNextWindowPos(imgui.Vec2{X: w.x, Y: w.y})
		imgui.SetNextWindowSize(imgui.Vec2{X: w.width, Y: w.height})
	} else {
		imgui.SetNextWindowPosV(imgui.Vec2{X: w.x, Y: w.y}, imgui.ConditionFirstUseEver, imgui.Vec2{X: 0, Y: 0})
		imgui.SetNextWindowSizeV(imgui.Vec2{X: w.width, Y: w.height}, imgui.ConditionFirstUseEver)
	}

	widgets = append(widgets, Custom(func() {
		w.hasFocus = IsWindowFocused()
		w.pos = imgui.WindowPos()
		w.size = imgui.WindowSize()
	}))

	w.layout = widgets

	return w
}

func (w *WindowWidget) Build() {
	if w.layout == nil {
		return
	}

	if w.bringToFront {
		w.bringToFront = false
		imgui.SetNextWindowFocus()
	}

	showed := imgui.BeginV(w.title, w.open, imgui.WindowFlags(w.flags))

	if showed {
		w.layout.Build()
	}

	imgui.End()
}

func (w *WindowWidget) HasFocus() bool {
	return w.hasFocus
}

func (w *WindowWidget) CurrentPosition() (x, y float32) {
	return w.pos.X, w.pos.Y
}

func (w *WindowWidget) CurrentSize() (width, height float32) {
	return w.size.X, w.size.Y
}

func (w *WindowWidget) BringToFront() {
	w.bringToFront = true
}
