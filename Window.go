package giu

import (
	"fmt"

	imgui "github.com/AllenDang/cimgui-go"
)

// SingleWindow creates one window filling all available space
// in MasterWindow. If SingleWindow is set up, no other windows may be
// defined.
func SingleWindow() *WindowWidget {
	pos := imgui.MainViewport().Pos()
	sizeX, sizeY := Context.backend.DisplaySize()
	title := GenAutoID("SingleWindow")

	return Window(title.String()). // TODO: maybe we should implement auto id in Window too?
					Flags(
			WindowFlags(imgui.WindowFlagsNoTitleBar)|
				WindowFlags(imgui.WindowFlagsNoCollapse)|
				WindowFlags(imgui.WindowFlagsNoScrollbar)|
				WindowFlags(imgui.WindowFlagsNoMove)|
				WindowFlags(imgui.WindowFlagsNoResize),
		).
		Pos(pos.X, pos.Y).Size(float32(sizeX), float32(sizeY))
}

// SingleWindowWithMenuBar creates a SingleWindow and allows to add menubar on its top.
func SingleWindowWithMenuBar() *WindowWidget {
	pos := imgui.MainViewport().Pos()
	sizeX, sizeY := Context.backend.DisplaySize()
	title := GenAutoID("SingleWindowWithMenuBar")

	return Window(title.String()). // TODO: maybe we should implement auto id in Window too?
					Flags(
			WindowFlags(imgui.WindowFlagsNoTitleBar)|
				WindowFlags(imgui.WindowFlagsNoCollapse)|
				WindowFlags(imgui.WindowFlagsNoScrollbar)|
				WindowFlags(imgui.WindowFlagsNoMove)|
				WindowFlags(imgui.WindowFlagsMenuBar)|
				WindowFlags(imgui.WindowFlagsNoResize),
		).Size(float32(sizeX), float32(sizeY)).Pos(pos.X, pos.Y)
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
	defaultPos := imgui.MainViewport().Pos()

	return (&WindowWidget{
		title: title,
	}).Pos(defaultPos.X, defaultPos.Y)
}

// IsOpen sets if window widget is `opened` (minimized).
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
// To prevent user from changing window position use
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

	if w.flags&WindowFlags(imgui.WindowFlagsNoMove) != 0 && w.flags&WindowFlags(imgui.WindowFlagsNoResize) != 0 {
		imgui.SetNextWindowPos(imgui.Vec2{X: w.x, Y: w.y})
		imgui.SetNextWindowSize(imgui.Vec2{X: w.width, Y: w.height})
	} else {
		imgui.SetNextWindowPosV(imgui.Vec2{X: w.x, Y: w.y}, imgui.CondFirstUseEver, imgui.Vec2{X: 0, Y: 0})
		imgui.SetNextWindowSizeV(imgui.Vec2{X: w.width, Y: w.height}, imgui.CondFirstUseEver)
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

	showed := imgui.BeginV(Context.FontAtlas.RegisterString(w.title), w.open, imgui.WindowFlags(w.flags))

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
	if state = GetState[windowState](Context, w.getStateID()); state == nil {
		state = &windowState{}
		SetState(Context, w.getStateID(), state)
	}

	return state
}
