package giu

import (
	"image/color"

	imgui "github.com/AllenDang/cimgui-go"
)

// MasterWindowFlags wrapps imgui.GLFWWindowFlags.
type MasterWindowFlags imgui.GLFWWindowFlags

// master window flags.
const (
	// Specifies the window will be fixed size.
	MasterWindowFlagsNotResizable MasterWindowFlags = MasterWindowFlags(imgui.GLFWWindowFlagsNotResizable)
	// Specifies whether the window is maximized.
	MasterWindowFlagsMaximized MasterWindowFlags = MasterWindowFlags(imgui.GLFWWindowFlagsMaximized)
	// Specifies whether the window will be always-on-top.
	MasterWindowFlagsFloating MasterWindowFlags = MasterWindowFlags(imgui.GLFWWindowFlagsFloating)
	// Specifies whether the window will be frameless.
	MasterWindowFlagsFrameless MasterWindowFlags = MasterWindowFlags(imgui.GLFWWindowFlagsFrameless)
	// Specifies whether the window will be transparent.
	MasterWindowFlagsTransparent MasterWindowFlags = MasterWindowFlags(imgui.GLFWWindowFlagsTransparent)
)

// MasterWindow represents a glfw master window
// It is a base for a windows (see Window.go).
type MasterWindow struct {
	title  string
	window imgui.GLFWwindow
}

// NewMasterWindow creates a new master window and initializes GLFW.
// it should be called in main function. For more details and use cases,
// see examples/helloworld/.
func NewMasterWindow(title string, width, height int, flags MasterWindowFlags) *MasterWindow {
	window := imgui.CreateGlfwWindow(title, width, height, imgui.GLFWWindowFlags(flags))

	Context = *CreateContext(window)

	mw := &MasterWindow{
		title:  title,
		window: window,
	}

	GiuTheme()

	imgui.SetBeforeRenderHook(mw.beforeRender)
	imgui.SetAfterRenderHook(mw.afterRender)

	return mw
}

func (w *MasterWindow) beforeRender() {
	Context.invalidAllState()
	Context.FontAtlas.rebuildFontAtlas()
}

func (w *MasterWindow) afterRender() {
	Context.cleanState()
}

func (w *MasterWindow) SetBgColor(col color.RGBA) {
	imgui.SetBgColor(ToVec4Color(col))
}

// Run runs the main loop.
// loopFunc will be used to construct the ui.
// Run should be called at the end of main function, after setting
// up the master window.
func (w *MasterWindow) Run(loopFunc func()) {
	Context.isAlive = true
	defer func() {
		Context.isAlive = false
	}()

	w.window.Run(loopFunc)
}
