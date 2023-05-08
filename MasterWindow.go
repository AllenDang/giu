package giu

import (
	"image"
	"image/color"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/faiface/mainthread"
	"gopkg.in/eapache/queue.v1"
)

// MasterWindowFlags wraps imgui.GLFWWindowFlags.
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

// DontCare could be used as an argument to (*MasterWindow).SetSizeLimits.
// var DontCare int = imgui.GlfwDontCare

// MasterWindow represents a glfw master window
// It is a base for a windows (see Window.go).
type MasterWindow struct {
	backend imgui.Backend

	width      int
	height     int
	clearColor [4]float32
	title      string
	context    imgui.Context
	io         *imgui.IO
	updateFunc func()

	// possibility to expend InputHandler's stuff
	// See SetAdditionalInputHandler
	additionalInputCallback InputHandlerHandleCallback
}

// NewMasterWindow creates a new master window and initializes GLFW.
// it should be called in main function. For more details and use cases,
// see examples/helloworld/.
func NewMasterWindow(title string, width, height int, flags MasterWindowFlags) *MasterWindow {
	ctx := imgui.CreateContext()
	imgui.PlotCreateContext()
	// imgui.ImNodesCreateContext() // TODO after implementing ImNodes in cimgui

	io := imgui.CurrentIO()

	// TODO: removed ConfigFlagEnablePowerSavingMode
	io.SetConfigFlags(imgui.BackendFlagsRendererHasVtxOffset)

	// Disable imgui.ini
	io.SetIniFilename("")

	backend := imgui.CreateBackend()

	Context = CreateContext(backend)

	// init texture loading queue
	Context.textureLoadingQueue = queue.New()

	mw := &MasterWindow{
		clearColor: [4]float32{0, 0, 0, 1},
		width:      width,
		height:     height,
		title:      title,
		io:         &io,
		context:    ctx,
	}

	backend.SetBeforeRenderHook(mw.beforeRender)
	backend.SetAfterRenderHook(mw.afterRender)
	backend.SetBeforeDestroyContextHook(mw.beforeDestroy)
	backend.CreateWindow(title, width, height, imgui.GLFWWindowFlags(flags))

	mw.SetInputHandler(newInputHandler())

	// TODO
	// p.SetSizeChangeCallback(mw.sizeChange)

	mw.setTheme()

	return mw
}

func (w *MasterWindow) setTheme() {
	// style := imgui.CurrentStyle()

	// Scale DPI in windows
	// TODO
	//if runtime.GOOS == "windows" {
	//	style.ScaleAllSizes(Context.GetPlatform().GetContentScale())
	//}

	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 2)
	imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 4)
	imgui.PushStyleVarFloat(imgui.StyleVarGrabRounding, 4)
	imgui.PushStyleVarFloat(imgui.StyleVarFrameBorderSize, 1)

	// TODO: idk why but I can't find these functions...
	// style.SetColor(imgui.ColText, imgui.Vec4{X: 0.95, Y: 0.96, Z: 0.98, W: 1.00})
	// style.SetColor(imgui.ColTextDisabled, imgui.Vec4{X: 0.36, Y: 0.42, Z: 0.47, W: 1.00})
	// style.SetColor(imgui.ColWindowBg, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	// style.SetColor(imgui.ColChildBg, imgui.Vec4{X: 0.15, Y: 0.18, Z: 0.22, W: 1.00})
	// style.SetColor(imgui.ColPopupBg, imgui.Vec4{X: 0.08, Y: 0.08, Z: 0.08, W: 0.94})
	// style.SetColor(imgui.ColBorder, imgui.Vec4{X: 0.08, Y: 0.10, Z: 0.12, W: 1.00})
	// style.SetColor(imgui.ColBorderShadow, imgui.Vec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.00})
	// style.SetColor(imgui.ColFrameBg, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	// style.SetColor(imgui.ColFrameBgHovered, imgui.Vec4{X: 0.12, Y: 0.20, Z: 0.28, W: 1.00})
	// style.SetColor(imgui.ColFrameBgActive, imgui.Vec4{X: 0.09, Y: 0.12, Z: 0.14, W: 1.00})
	// style.SetColor(imgui.ColTitleBg, imgui.Vec4{X: 0.09, Y: 0.12, Z: 0.14, W: 0.65})
	// style.SetColor(imgui.ColTitleBgActive, imgui.Vec4{X: 0.08, Y: 0.10, Z: 0.12, W: 1.00})
	// style.SetColor(imgui.ColTitleBgCollapsed, imgui.Vec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.51})
	// style.SetColor(imgui.ColMenuBarBg, imgui.Vec4{X: 0.15, Y: 0.18, Z: 0.22, W: 1.00})
	// style.SetColor(imgui.ColScrollbarBg, imgui.Vec4{X: 0.02, Y: 0.02, Z: 0.02, W: 0.39})
	// style.SetColor(imgui.ColScrollbarGrab, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	// style.SetColor(imgui.ColScrollbarGrabHovered, imgui.Vec4{X: 0.18, Y: 0.22, Z: 0.25, W: 1.00})
	// style.SetColor(imgui.ColScrollbarGrabActive, imgui.Vec4{X: 0.09, Y: 0.21, Z: 0.31, W: 1.00})
	// style.SetColor(imgui.ColCheckMark, imgui.Vec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	// style.SetColor(imgui.ColSliderGrab, imgui.Vec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	// style.SetColor(imgui.ColSliderGrabActive, imgui.Vec4{X: 0.37, Y: 0.61, Z: 1.00, W: 1.00})
	// style.SetColor(imgui.ColButton, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	// style.SetColor(imgui.ColButtonHovered, imgui.Vec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	// style.SetColor(imgui.ColButtonActive, imgui.Vec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00})
	// style.SetColor(imgui.ColHeader, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 0.55})
	// style.SetColor(imgui.ColHeaderHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80})
	// style.SetColor(imgui.ColHeaderActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00})
	// style.SetColor(imgui.ColSeparator, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	// style.SetColor(imgui.ColSeparatorHovered, imgui.Vec4{X: 0.10, Y: 0.40, Z: 0.75, W: 0.78})
	// style.SetColor(imgui.ColSeparatorActive, imgui.Vec4{X: 0.10, Y: 0.40, Z: 0.75, W: 1.00})
	// style.SetColor(imgui.ColResizeGrip, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.25})
	// style.SetColor(imgui.ColResizeGripHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.67})
	// style.SetColor(imgui.ColResizeGripActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.95})
	// style.SetColor(imgui.ColTab, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	// style.SetColor(imgui.ColTabHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80})
	// style.SetColor(imgui.ColTabActive, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	// style.SetColor(imgui.ColTabUnfocused, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	// style.SetColor(imgui.ColTabUnfocusedActive, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	// style.SetColor(imgui.ColPlotLines, imgui.Vec4{X: 0.61, Y: 0.61, Z: 0.61, W: 1.00})
	// style.SetColor(imgui.ColPlotLinesHovered, imgui.Vec4{X: 1.00, Y: 0.43, Z: 0.35, W: 1.00})
	// style.SetColor(imgui.ColPlotHistogram, imgui.Vec4{X: 0.90, Y: 0.70, Z: 0.00, W: 1.00})
	// style.SetColor(imgui.ColPlotHistogramHovered, imgui.Vec4{X: 1.00, Y: 0.60, Z: 0.00, W: 1.00})
	// style.SetColor(imgui.ColTextSelectedBg, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.35})
	// style.SetColor(imgui.ColDragDropTarget, imgui.Vec4{X: 1.00, Y: 1.00, Z: 0.00, W: 0.90})
	// style.SetColor(imgui.ColNavHighlight, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00})
	// style.SetColor(imgui.ColNavWindowingHighlight, imgui.Vec4{X: 1.00, Y: 1.00, Z: 1.00, W: 0.70})
	// style.SetColor(imgui.ColTableHeaderBg, imgui.Vec4{X: 0.12, Y: 0.20, Z: 0.28, W: 1.00})
	// style.SetColor(imgui.ColTableBorderStrong, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	// style.SetColor(imgui.ColTableBorderLight, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 0.70})
}

// SetBgColor sets background color of master window.
func (w *MasterWindow) SetBgColor(bgColor color.Color) {
	const mask = 0xffff

	r, g, b, a := bgColor.RGBA()
	w.clearColor = [4]float32{
		float32(r) / mask,
		float32(g) / mask,
		float32(b) / mask,
		float32(a) / mask,
	}
}

func (w *MasterWindow) sizeChange(width, height int) {
	w.render()
}

func (w *MasterWindow) beforeRender() {
	Context.invalidAllState()
	Context.FontAtlas.rebuildFontAtlas()

	// process texture load requests
	if Context.textureLoadingQueue != nil && Context.textureLoadingQueue.Length() > 0 {
		for Context.textureLoadingQueue.Length() > 0 {
			request, ok := Context.textureLoadingQueue.Remove().(textureLoadRequest)
			Assert(ok, "MasterWindow", "Run", "processing texture requests: wrong type of texture request")
			NewTextureFromRgba(request.img, request.cb)
		}
	}

	// TODO InputHandler not re-implemented yet
	// p.ProcessEvents()
}

func (w *MasterWindow) afterRender() {
	Context.cleanState()
}

func (w *MasterWindow) beforeDestroy() {
	imgui.PlotDestroyContext()
	//w.context.Destroy() // TODO: check why it panics if it is here
	// imgui.ImNodesDestroyContext() // TODO: after adding ImNodes (https://github.com/AllenDang/cimgui-go/issues/137)
}

func (w *MasterWindow) render() {
	mainStylesheet := Style()
	if s, found := Context.cssStylesheet["main"]; found {
		mainStylesheet = s
	}

	mainStylesheet.Push()
	w.updateFunc()
	mainStylesheet.Pop()
}

// Run runs the main loop.
// loopFunc will be used to construct the ui.
// Run should be called at the end of main function, after setting
// up the master window.
func (w *MasterWindow) Run(loopFunc func()) {
	mainthread.Run(func() {
		Context.isRunning = true
		w.updateFunc = loopFunc

		Context.isAlive = true

		Context.backend.Run(w.render)

		Context.isAlive = false

		Context.isRunning = false
	})
}

// GetSize return size of master window.
func (w *MasterWindow) GetSize() (width, height int) {
	if w.backend != nil {
		// TODO
		//if glfwPlatform, ok := w.platform.(*imgui.GLFW); ok {
		//return glfwPlatform.GetWindow().GetSize()
		//}
		w, h := w.backend.DisplaySize()
		return int(w), int(h)
	}

	return w.width, w.height
}

// GetPos return position of master window.
func (w *MasterWindow) GetPos() (x, y int) {
	if w.backend != nil {
		// TODO
		//if glfwPlatform, ok := w.platform.(*imgui.GLFW); ok {
		//	x, y = glfwPlatform.GetWindow().GetPos()
		//}
	}

	return
}

// SetPos sets position of master window.
func (w *MasterWindow) SetPos(x, y int) {
	if w.backend != nil {
		w.backend.SetWindowPos(x, y)
	}
}

// SetSize sets size of master window.
func (w *MasterWindow) SetSize(x, y int) {
	if w.backend != nil {
		// TODO
		//if glfwPlatform, ok := w.platform.(*imgui.GLFW); ok {
		//	mainthread.CallNonBlock(func() {
		//		glfwPlatform.GetWindow().SetSize(x, y)
		//	})
		//}
	}
}

// SetCloseCallback sets the close callback of the window, which is called when
// the user attempts to close the window, for example by clicking the close
// widget in the title bar.
//
// The close flag is set before this callback is called, but you can modify it at
// any time with returned value of callback function.
//
// Mac OS X: Selecting Quit from the application menu will trigger the close
// callback for all windows.
func (w *MasterWindow) SetCloseCallback(cb func() bool) {
	// TODO
	// w.backend.SetCloseCallback(cb)
}

// SetDropCallback sets callback when file was dropped into the window.
func (w *MasterWindow) SetDropCallback(cb func([]string)) {
	// TODO: https://github.com/AllenDang/cimgui-go/pull/145
	// w.platform.SetDropCallback(cb)
}

// RegisterKeyboardShortcuts registers a global - master window - keyboard shortcuts.
func (w *MasterWindow) RegisterKeyboardShortcuts(s ...WindowShortcut) *MasterWindow {
	for _, shortcut := range s {
		Context.InputHandler.RegisterKeyboardShortcuts(Shortcut{
			Key:      shortcut.Key,
			Modifier: shortcut.Modifier,
			Callback: shortcut.Callback,
			IsGlobal: GlobalShortcut,
		})
	}

	return w
}

// SetIcon sets the icon of the specified window. If passed an array of candidate images,
// those of or closest to the sizes desired by the system are selected. If no images are
// specified, the window reverts to its default icon.
//
// The image is ideally provided in the form of *image.NRGBA.
// The pixels are 32-bit, little-endian, non-premultiplied RGBA, i.e. eight
// bits per channel with the red channel first. They are arranged canonically
// as packed sequential rows, starting from the top-left corner. If the image
// type is not *image.NRGBA, it will be converted to it.
//
// The desired image sizes varies depending on platform and system settings. The selected
// images will be rescaled as needed. Good sizes include 16x16, 32x32 and 48x48.
func (w *MasterWindow) SetIcon(icons []image.Image) {
	// TODO
	// w.platform.SetIcon(icons)
}

// SetSizeLimits sets the size limits of the client area of the specified window.
// If the window is full screen or not resizable, this function does nothing.
//
// The size limits are applied immediately and may cause the window to be resized.
// To specify only a minimum size or only a maximum one, set the other pair to giu.DontCare.
// To disable size limits for a window, set them all to giu.DontCare.
func (w *MasterWindow) SetSizeLimits(minw, minh, maxw, maxh int) {
	// TODO
	// w.platform.SetSizeLimits(minw, minh, maxw, maxh)
}

// SetTitle updates master window's title.
func (w *MasterWindow) SetTitle(title string) {
	// TODO
	// w.platform.SetTitle(title)
}

// Close will safely close the master window.
func (w *MasterWindow) Close() {
	w.SetShouldClose(true)
}

// SetShouldClose sets whether master window should be closed.
func (w *MasterWindow) SetShouldClose(v bool) {
	// TODO
	// w.platform.SetShouldStop(v)
}

// SetInputHandler allows to change default input handler.
// see InputHandler.go.
func (w *MasterWindow) SetInputHandler(handler InputHandler) {
	// TODO
	Context.InputHandler = handler

	//w.platform.SetInputCallback(func(key glfw.Key, modifier glfw.ModifierKey, action glfw.Action) {
	//	k, m, a := Key(key), Modifier(modifier), Action(action)
	//	handler.Handle(k, m, a)
	//	if w.additionalInputCallback != nil {
	//		w.additionalInputCallback(k, m, a)
	//	}
	//})
}

// SetAdditionalInputHandlerCallback allows to set an input callback to handle more events (not only these from giu.inputHandler).
// See examples/issue-501.
func (w *MasterWindow) SetAdditionalInputHandlerCallback(cb InputHandlerHandleCallback) {
	w.additionalInputCallback = cb
}
