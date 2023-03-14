package giu

import (
	"fmt"

	imgui "github.com/AllenDang/cimgui-go"
)

// SingleWindow creates one window filling all available space
// in MasterWindow. If SingleWindow is set up, no other windows can't be
// definied.
func SingleWindow() *WindowWidget {
	width, height := Context.window.DisplaySize()
	title := fmt.Sprintf("SingleWindow_%d", Context.GetWidgetIndex())
	return Window(title).
		Flags(
			imgui.WindowFlagsNoTitleBar|
				imgui.WindowFlagsNoCollapse|
				imgui.WindowFlagsNoScrollbar|
				imgui.WindowFlagsNoMove|
				imgui.WindowFlagsNoResize).
		Size(float32(width), float32(height))
}

// SingleWindowWithMenuBar creates a SingleWindow and allows to add menubar on its top.
func SingleWindowWithMenuBar() *WindowWidget {
	width, height := Context.window.DisplaySize()
	title := fmt.Sprintf("SingleWindow_%d", Context.GetWidgetIndex())
	return Window(title).
		Flags(
			imgui.WindowFlagsNoTitleBar|
				imgui.WindowFlagsNoCollapse|
				imgui.WindowFlagsNoScrollbar|
				imgui.WindowFlagsNoMove|
				imgui.WindowFlagsMenuBar|
				imgui.WindowFlagsNoResize).Size(float32(width), float32(height))
}

// WindowWidget represents imgui.Window
// Windows are used to display ui widgets.
// They are in second place in the giu hierarchy (after the MasterWindow)
// NOTE: to disable multiple window, use SingleWindow.
type WindowWidget struct {
	title         string
	open          *bool
	flags         imgui.WindowFlags
	x, y          float32
	width, height float32
	bringToFront  bool
}

// Window creates a WindowWidget.
func Window(title string) *WindowWidget {
	return &WindowWidget{
		title: title,
	}
}

// IsOpen sets if window widget is `opened` (minimalized).
func (w *WindowWidget) IsOpen(open *bool) *WindowWidget {
	w.open = open
	return w
}

// Flags sets window flags.
func (w *WindowWidget) Flags(flags imgui.WindowFlags) *WindowWidget {
	w.flags = flags
	return w
}

// Size sets window size
// NOTE: size can be changed by user, if you want to prevent
// user from changing window size, use NoResize flag.
func (w *WindowWidget) Size(width, height float32) *WindowWidget {
	w.width, w.height = width, height
	return w
}

// Pos sets the window start position
// NOTE: The position could be changed by user later.
// To prevent user from changin window position use
// WIndowFlagsNoMove.
func (w *WindowWidget) Pos(x, y float32) *WindowWidget {
	w.x, w.y = x, y
	return w
}

// Layout is a final step of the window setup.
// it should be called to add a layout to the window and build it.
func (w *WindowWidget) Layout(widgets ...Widget) {
	if widgets == nil {
		return
	}

	viewport := imgui.MainViewport()
	basePos := viewport.Pos()

	if w.flags&imgui.WindowFlagsNoMove != 0 && w.flags&imgui.WindowFlagsNoResize != 0 {
		imgui.SetNextWindowPos(imgui.Vec2{X: basePos.X + w.x, Y: basePos.Y + w.y})
		imgui.SetNextWindowSize(imgui.Vec2{X: w.width, Y: w.height})
	} else {
		imgui.SetNextWindowPosV(imgui.Vec2{X: basePos.X + w.x, Y: basePos.Y + w.y}, imgui.CondFirstUseEver, imgui.Vec2{X: 0, Y: 0})
		imgui.SetNextWindowSizeV(imgui.Vec2{X: w.width, Y: w.height}, imgui.CondFirstUseEver)
	}

	if w.bringToFront {
		imgui.SetNextWindowFocus()
		w.bringToFront = false
	}

	showed := imgui.BeginV(Context.FontAtlas.RegisterString(w.title), w.open, w.flags)

	if showed {
		Layout(widgets).Build()
	}

	imgui.End()
}

// CurrentPosition returns a current position of the window.
func (w *WindowWidget) CurrentPosition() (x, y float32) {
	pos := imgui.WindowPos()
	return pos.X, pos.Y
}

// CurrentSize returns current size of the window.
func (w *WindowWidget) CurrentSize() (width, height float32) {
	size := imgui.WindowPos()
	return size.X, size.Y
}

// BringToFront sets window focused.
func (w *WindowWidget) BringToFront() {
	w.bringToFront = true
}
