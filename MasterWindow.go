package giu

import (
	"errors"
	"image"
	"image/color"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/glfwbackend"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imguizmo"
	"github.com/AllenDang/cimgui-go/immarkdown"
	"github.com/AllenDang/cimgui-go/imnodes"
	"github.com/AllenDang/cimgui-go/implot"
	"golang.org/x/image/colornames"
)

// MasterWindowFlags implements BackendWindowFlags.
type MasterWindowFlags int

// master window flags.
const (
	// Specifies the window will be fixed size.
	MasterWindowFlagsNotResizable MasterWindowFlags = 1 << iota
	// Specifies whether the window is maximized.
	MasterWindowFlagsMaximized
	// Specifies whether the window will be always-on-top.
	MasterWindowFlagsFloating
	// Specifies whether the window will be frameless.
	MasterWindowFlagsFrameless
	// Specifies whether the window will be transparent.
	MasterWindowFlagsTransparent
	// Specifies whether the window will be hidden (for use with multiple windows).
	MasterWindowFlagsHidden
)

// parseAndApply converts MasterWindowFlags to appropriate glfwbackend.GLFWWindowFlags.

// TODO(gucio321) implement this in cimgui-go
// DontCare could be used as an argument to (*MasterWindow).SetSizeLimits.
// var DontCare int = imgui.GlfwDontCare

// MasterWindow represents a glfw master window
// It is a base for a windows (see Window.go).
type MasterWindow struct {
	// generally Context should be used instead but as I don't like global
	// variables, I prefer to keep a pointer here and refer it as possible.
	ctx *GIUContext

	width      int
	height     int
	clearColor imgui.Vec4
	title      string
	context    *imgui.Context
	io         *imgui.IO
	updateFunc func()
	theme      *StyleSetter

	// possibility to expend InputHandler's stuff
	// See SetAdditionalInputHandler
	additionalInputCallback InputHandlerHandleCallback
}

// NewMasterWindow creates a new master window and initializes GLFW.
// it should be called in main function. For more details and use cases,
// see examples/helloworld/.
func NewMasterWindow(title string, width, height int, flags MasterWindowFlags) *MasterWindow {
	imGuiContext := imgui.CreateContext()

	implot.CreateContext()
	imnodes.CreateContext()

	io := imgui.CurrentIO()

	// TODO: removed ConfigFlagEnablePowerSavingMode
	// TODO: removed io.SetConfigFlags(imgui.BackendFlagsRendererHasVtxOffset)
	io.SetBackendFlags(imgui.BackendFlagsRendererHasVtxOffset)

	currentBackend, err := backend.CreateBackend(NewGLFWBackend())
	if err != nil && !errors.Is(err, backend.CExposerError) {
		panic(err)
	}

	// Create GIU context
	Context = CreateContext(currentBackend)

	mw := &MasterWindow{
		clearColor: imgui.Vec4{X: 0, Y: 0, Z: 0, W: 1},
		width:      width,
		height:     height,
		title:      title,
		io:         io,
		context:    imGuiContext,
		ctx:        Context,
		theme:      DefaultTheme(),
	}

	// Disable imgui.ini
	mw.SetUserFile("")

	currentBackend.SetBeforeRenderHook(mw.beforeRender)
	currentBackend.SetAfterRenderHook(mw.afterRender)
	currentBackend.SetBeforeDestroyContextHook(mw.beforeDestroy)

	for f := MasterWindowFlagsNotResizable; f <= MasterWindowFlagsHidden; f <<= 1 {
		if f&flags != 0 {
			currentBackend.SetWindowFlags(f, 0) // 0 because it is not used anyway (flag values are determined by giu
		}
	}

	currentBackend.CreateWindow(title, width, height)

	mw.SetInputHandler(newInputHandler())

	mw.ctx.backend.SetSizeChangeCallback(mw.sizeChange)

	mw.SetBgColor(colornames.Black)

	mw.SetScale(0) // set content scale

	return mw
}

func (w *MasterWindow) setTheme() (fin func()) {
	if w.theme == nil {
		return func() {}
	}

	w.theme.Push()

	return w.theme.Pop
}

func (w *MasterWindow) sizeChange(_, _ int) {
	// noop
}

func (w *MasterWindow) beforeRender() {
	// Clean callbacks
	// see https://github.com/AllenDang/cimgui-go?tab=readme-ov-file#callbacks
	immarkdown.ClearMarkdownLinkCallbackPool()

	Context.FontAtlas.rebuildFontAtlas()

	Context.WithLock(func() {
		// process texture load requests
		if Context.textureLoadingQueue != nil && Context.textureLoadingQueue.Length() > 0 {
			for Context.textureLoadingQueue.Length() > 0 {
				request, ok := Context.textureLoadingQueue.Remove().(textureLoadRequest)
				Assert(ok, "MasterWindow", "Run", "processing texture requests: wrong type of texture request")
				NewTextureFromRgba(request.img, request.cb)
			}
		}

		// process texture free requests
		if Context.textureFreeingQueue != nil && Context.textureFreeingQueue.Length() > 0 {
			for Context.textureFreeingQueue.Length() > 0 {
				request, ok := Context.textureFreeingQueue.Remove().(textureFreeRequest)
				Assert(ok, "MasterWindow", "Run", "processing texture requests: wrong type of texture request")
				request.tex.tex.Release()
			}
		}
	})
}

func (w *MasterWindow) afterRender() {
}

func (w *MasterWindow) beforeDestroy() {
	implot.DestroyContext()
	imnodes.DestroyContext()
}

func (w *MasterWindow) render() {
	imguizmo.BeginFrame()

	Context.cleanStates()
	defer Context.SetDirty()

	fin := w.setTheme()
	defer fin()

	mainStylesheet := Context.cssStylesheet.GetTag(MainTag)

	mainStylesheet.Push()
	w.updateFunc()
	mainStylesheet.Pop()
}

// Run runs the main loop.
// loopFunc will be used to construct the ui.
// Run should be called at the end of main function, after setting
// up the master window.
func (w *MasterWindow) Run(loopFunc func()) {
	mainthreadCallPlatform(func() {
		Context.isRunning = true
		w.updateFunc = loopFunc

		Context.m.Lock()
		Context.isAlive = true
		Context.m.Unlock()

		Context.backend.Run(w.render)

		Context.m.Lock()
		Context.isAlive = false

		Context.isRunning = false
		Context.m.Unlock()
	})
}

// GetSize return size of master window.
func (w *MasterWindow) GetSize() (width, height int) {
	if w.ctx.backend != nil {
		w, h := w.ctx.backend.DisplaySize()
		return int(w), int(h)
	}

	return w.width, w.height
}

// SetStyle sets the style for the master window. Default is DefaultTheme().
func (w *MasterWindow) SetStyle(ss *StyleSetter) {
	w.theme = ss
}

// GetStyle returns the style for the master window.
func (w *MasterWindow) GetStyle() *StyleSetter {
	return w.theme.Add(Context.cssStylesheet.GetTag(MainTag))
}

// SetBgColor sets background color of master window.
func (w *MasterWindow) SetBgColor(bgColor color.Color) {
	const mask = 0xffff

	r, g, b, a := bgColor.RGBA()
	w.clearColor = imgui.Vec4{
		X: float32(r) / mask,
		Y: float32(g) / mask,
		Z: float32(b) / mask,
		W: float32(a) / mask,
	}

	w.ctx.backend.SetBgColor(w.clearColor)
}

// SetTargetFPS sets target FPS of master window.
// Default for GLFW is 30.
func (w *MasterWindow) SetTargetFPS(fps uint) {
	w.ctx.backend.SetTargetFPS(fps)
}

// GetPos return position of master window.
func (w *MasterWindow) GetPos() (x, y int) {
	var xResult, yResult int32
	if w.ctx.backend != nil {
		xResult, yResult = w.ctx.backend.GetWindowPos()
	}

	return int(xResult), int(yResult)
}

// SetPos sets position of master window.
func (w *MasterWindow) SetPos(x, y int) {
	if w.ctx.backend != nil {
		w.ctx.backend.SetWindowPos(x, y)
	}
}

// SetSize sets size of master window.
func (w *MasterWindow) SetSize(x, y int) {
	if w.ctx.backend != nil {
		w.ctx.backend.SetWindowSize(x, y)
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
	w.ctx.backend.SetCloseCallback(func() {
		w.ctx.backend.SetShouldClose(cb())
	})
}

// SetDropCallback sets callback when file was dropped into the window.
func (w *MasterWindow) SetDropCallback(cb func([]string)) {
	w.ctx.backend.SetDropCallback(cb)
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
func (w *MasterWindow) SetIcon(icons ...image.Image) {
	w.ctx.backend.SetIcons(icons...)
}

// SetSizeLimits sets the size limits of the client area of the specified window.
// If the window is full screen or not resizable, this function does nothing.
//
// The size limits are applied immediately and may cause the window to be resized.
// To specify only a minimum size or only a maximum one, set the other pair to giu.DontCare.
// To disable size limits for a window, set them all to giu.DontCare.
func (w *MasterWindow) SetSizeLimits(minw, minh, maxw, maxh int) {
	w.ctx.backend.SetWindowSizeLimits(minw, minh, maxw, maxh)
}

// SetTitle updates master window's title.
func (w *MasterWindow) SetTitle(title string) {
	w.ctx.backend.SetWindowTitle(title)
}

// Close will safely close the master window.
func (w *MasterWindow) Close() {
	w.ctx.backend.SetShouldClose(true)
}

// SetShouldClose sets whether master window should be closed.
func (w *MasterWindow) SetShouldClose(v bool) {
	w.ctx.backend.SetShouldClose(v)
}

// SetInputHandler allows to change default input handler.
// see InputHandler.go.
func (w *MasterWindow) SetInputHandler(handler InputHandler) {
	Context.InputHandler = handler

	w.ctx.backend.SetKeyCallback(func(key, _, action, modifier int) {
		k, m, a := keyFromGLFWKey(glfwbackend.GLFWKey(key)), Modifier(modifier), Action(action)
		handler.Handle(k, m, a)

		if w.additionalInputCallback != nil {
			w.additionalInputCallback(k, m, a)
		}
	})
}

// SetAdditionalInputHandlerCallback allows to set an input callback to handle more events (not only these from giu.inputHandler).
// See examples/issue-501.
func (w *MasterWindow) SetAdditionalInputHandlerCallback(cb InputHandlerHandleCallback) {
	w.additionalInputCallback = cb
}

// SetUserFile sets the path to the .ini file which saves the user preferences (e.g. subwindow positions, table column widths, etc.).
// Provide an empty string to disable it. This is the default.
// See examples/sortable-table.
func (w *MasterWindow) SetUserFile(path string) {
	w.io.SetIniFilename(path)
}

// SetScale is executed internally by NewMasterWindow with the default content scale factor.
// You can call it if you relly want.
// If 0 passed, ContentScale will be re-applied.
func (w *MasterWindow) SetScale(scale float32) {
	if scale == 0 {
		scale, _ = Context.backend.ContentScale()
	}

	s := imgui.CurrentStyle()
	s.ScaleAllSizes(scale)
	s.SetFontScaleDpi(scale)
}
