package giu

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
)

func SingleWindow() *WindowWidget {
	size := Context.platform.DisplaySize()
	title := fmt.Sprintf("SingleWindow_%d", Context.GetWidgetIndex())
	return Window(title).
		Flags(
			imgui.WindowFlagsNoTitleBar|
				imgui.WindowFlagsNoCollapse|
				imgui.WindowFlagsNoScrollbar|
				imgui.WindowFlagsNoMove|
				imgui.WindowFlagsNoResize).
		Size(size[0], size[1])
}

func SingleWindowWithMenuBar() *WindowWidget {
	size := Context.platform.DisplaySize()
	title := fmt.Sprintf("SingleWindow_%d", Context.GetWidgetIndex())
	return Window(title).
		Flags(
			imgui.WindowFlagsNoTitleBar|
				imgui.WindowFlagsNoCollapse|
				imgui.WindowFlagsNoScrollbar|
				imgui.WindowFlagsNoMove|
				imgui.WindowFlagsMenuBar|
				imgui.WindowFlagsNoResize).Size(size[0], size[1])
}

type windowState struct {
	hasFocus bool
	currentPosition,
	currentSize imgui.Vec2
}

func (s *windowState) Dispose() {
	// noop
}

type WindowWidget struct {
	title         string
	open          *bool
	flags         WindowFlags
	x, y          float32
	width, height float32
	bringToFront  bool
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

	ws := w.getState()

	if w.flags&imgui.WindowFlagsNoMove != 0 && w.flags&imgui.WindowFlagsNoResize != 0 {
		imgui.SetNextWindowPos(imgui.Vec2{X: w.x, Y: w.y})
		imgui.SetNextWindowSize(imgui.Vec2{X: w.width, Y: w.height})
	} else {
		imgui.SetNextWindowPosV(imgui.Vec2{X: w.x, Y: w.y}, imgui.ConditionFirstUseEver, imgui.Vec2{X: 0, Y: 0})
		imgui.SetNextWindowSizeV(imgui.Vec2{X: w.width, Y: w.height}, imgui.ConditionFirstUseEver)
	}

	if w.bringToFront {
		imgui.SetNextWindowFocus()
		w.bringToFront = false
	}

	widgets = append(widgets,
		Custom(func() {
			hasFocus := IsWindowFocused(0)
			if !hasFocus && ws.hasFocus {
				unregisterWindowShortcuts()
			}

			ws.hasFocus = hasFocus

			ws.currentPosition = imgui.WindowPos()
			ws.currentSize = imgui.WindowSize()
		}),
	)

	showed := imgui.BeginV(tStr(w.title), w.open, int(w.flags))

	if showed {
		Layout(widgets).Build()
	}

	imgui.End()
}

func (w *WindowWidget) CurrentPosition() (x, y float32) {
	pos := w.getState().currentPosition
	return pos.X, pos.Y
}

func (w *WindowWidget) CurrentSize() (width, height float32) {
	size := w.getState().currentSize
	return size.X, size.Y
}

func (w *WindowWidget) BringToFront() {
	w.bringToFront = true
}

func (w *WindowWidget) HasFocus() bool {
	return w.getState().hasFocus
}

func (w *WindowWidget) RegisterKeyboardShortcuts(s ...WindowShortcut) *WindowWidget {
	if w.HasFocus() {
		for _, shortcut := range s {
			RegisterKeyboardShortcuts(Shortcut{
				Key:      shortcut.Key,
				Modifier: shortcut.Modifier,
				Callback: shortcut.Callback,
				IsGlobal: LocalShortcut,
			})
		}
	}

	return w
}

func (w *WindowWidget) getStateID() string {
	return fmt.Sprintf("%s_windowState", w.title)
}

// returns window state
func (w *WindowWidget) getState() (state *windowState) {
	s := Context.GetState(w.getStateID())

	if s != nil {
		state = s.(*windowState)
	} else {
		state = &windowState{}

		Context.SetState(w.getStateID(), state)
	}

	return state
}
