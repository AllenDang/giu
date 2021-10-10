package giu

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
)

// SingleWindow creates one window filling all available space
// in MasterWindow. If SingleWindow is set up, no other windows can't be
// definied.
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

// SingleWindowWithMenuBar creates a SingleWindow and allows to add menubar on its top.
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

var _ Disposable = &windowState{}

type windowState struct {
	hasFocus bool
	currentPosition,
	currentSize imgui.Vec2
}

// Dispose implements Disposable interface.
func (s *windowState) Dispose() {
	// noop
}

// WindowWidget represents imgui.Window
// Windows are used to display ui widgets.
// They are in second place in the giu hierarchy (after the MasterWindow)
// NOTE: to disable multiple window, use SingleWindow.
type WindowWidget struct {
	title         string
	open          *bool
	flags         WindowFlags
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
func (w *WindowWidget) Flags(flags WindowFlags) *WindowWidget {
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
				Context.InputHandler.UnregisterWindowShortcuts()
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

// CurrentPosition returns a current position of the window.
func (w *WindowWidget) CurrentPosition() (x, y float32) {
	pos := w.getState().currentPosition
	return pos.X, pos.Y
}

// CurrentSize returns current size of the window.
func (w *WindowWidget) CurrentSize() (width, height float32) {
	size := w.getState().currentSize
	return size.X, size.Y
}

// BringToFront sets window focused.
func (w *WindowWidget) BringToFront() {
	w.bringToFront = true
}

// HasFocus returns true if window is focused.
func (w *WindowWidget) HasFocus() bool {
	return w.getState().hasFocus
}

// RegisterKeyboardShortcuts adds local (window-level) keyboard shortcuts
// see InputHandler.go.
func (w *WindowWidget) RegisterKeyboardShortcuts(s ...WindowShortcut) *WindowWidget {
	if w.HasFocus() {
		for _, shortcut := range s {
			Context.InputHandler.RegisterKeyboardShortcuts(Shortcut{
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

// returns window state.
func (w *WindowWidget) getState() (state *windowState) {
	if s := Context.GetState(w.getStateID()); s != nil {
		var isOk bool
		state, isOk = s.(*windowState)
		Assert(isOk, "WindowWidget", "getState", "unexpected state recovered.")
	} else {
		state = &windowState{}

		Context.SetState(w.getStateID(), state)
	}

	return state
}
