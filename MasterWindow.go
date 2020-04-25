package giu

import (
	"image/color"
	"time"

	"github.com/AllenDang/giu/imgui"
)

type MasterWindowFlags imgui.GLFWWindowFlags

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

type MasterWindow struct {
	width      int
	height     int
	clearColor [4]float32
	title      string
	platform   imgui.Platform
	renderer   imgui.Renderer
	context    *imgui.Context
	io         *imgui.IO
	updateFunc func()
}

func NewMasterWindow(title string, width, height int, flags MasterWindowFlags, loadFontFunc func()) *MasterWindow {
	context := imgui.CreateContext(nil)

	io := imgui.CurrentIO()

	io.SetConfigFlags(imgui.ConfigFlagEnablePowerSavingMode)

	// Disable imgui.ini
	io.SetIniFilename("")

	p, err := imgui.NewGLFW(io, title, width, height, imgui.GLFWWindowFlags(flags))
	if err != nil {
		panic(err)
	}

	scale := p.GetContentScale()

	imgui.DPIScale = scale

	// Init Context.state
	Context.state = make(map[string]*state)
	// Assign platform to contex
	Context.platform = p

	if loadFontFunc != nil {
		loadFontFunc()
	}

	r, err := imgui.NewOpenGL3(io, scale)
	if err != nil {
		panic(err)
	}

	// Create context
	Context.renderer = r

	mw := &MasterWindow{
		clearColor: [4]float32{0.22, 0.26, 0.28, 1},
		width:      width,
		height:     height,
		title:      title,
		io:         &io,
		context:    context,
		platform:   p,
		renderer:   r,
	}

	p.SetSizeChangeCallback(mw.sizeChange)

	mw.setTheme()

	return mw
}

func (w *MasterWindow) setTheme() {
	style := imgui.CurrentStyle()

	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 2)
	imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 2)
	imgui.PushStyleVarFloat(imgui.StyleVarFrameBorderSize, 1)

	style.SetColor(imgui.StyleColorText, imgui.Vec4{X: 0.82, Y: 0.82, Z: 0.82, W: 1.00})
	// style.SetColor(imgui.StyleColorTextDisabled, imgui.Vec4{})
	style.SetColor(imgui.StyleColorWindowBg, imgui.Vec4{X: 0.22, Y: 0.26, Z: 0.28, W: 1.00})
	// style.SetColor(imgui.StyleColorChildBg, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorPopupBg, imgui.Vec4{})
	style.SetColor(imgui.StyleColorBorder, imgui.Vec4{X: 0.18, Y: 0.18, Z: 0.18, W: 1.00})
	// style.SetColor(imgui.StyleColorBorderShadow, imgui.Vec4{})
	style.SetColor(imgui.StyleColorFrameBg, imgui.Vec4{X: 0.20, Y: 0.23, Z: 0.24, W: 1.00})
	style.SetColor(imgui.StyleColorFrameBgHovered, imgui.Vec4{X: 0.18, Y: 0.21, Z: 0.22, W: 1.00})
	style.SetColor(imgui.StyleColorFrameBgActive, imgui.Vec4{X: 0.19, Y: 0.33, Z: 0.44, W: 1.00})
	// style.SetColor(imgui.StyleColorTitleBg, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorTitleBgActive, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorTitleBgCollapsed, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorMenuBarBg, imgui.Vec4{})
	style.SetColor(imgui.StyleColorScrollbarBg, imgui.Vec4{X: 0.20, Y: 0.23, Z: 0.24, W: 1.00})
	style.SetColor(imgui.StyleColorScrollbarGrab, imgui.Vec4{X: 0.19, Y: 0.33, Z: 0.44, W: 1.00})
	style.SetColor(imgui.StyleColorScrollbarGrabHovered, imgui.Vec4{X: 0.21, Y: 0.35, Z: 0.45, W: 1.00})
	style.SetColor(imgui.StyleColorScrollbarGrabActive, imgui.Vec4{X: 0.23, Y: 0.36, Z: 0.47, W: 1.00})
	// style.SetColor(imgui.StyleColorCheckMark, imgui.Vec4{})
	style.SetColor(imgui.StyleColorSliderGrab, imgui.Vec4{X: 0.19, Y: 0.33, Z: 0.44, W: 1.00})
	style.SetColor(imgui.StyleColorSliderGrabActive, imgui.Vec4{X: 0.23, Y: 0.36, Z: 0.47, W: 1.00})
	style.SetColor(imgui.StyleColorButton, imgui.Vec4{X: 0.19, Y: 0.33, Z: 0.44, W: 1.00})
	style.SetColor(imgui.StyleColorButtonHovered, imgui.Vec4{X: 0.23, Y: 0.36, Z: 0.47, W: 1.00})
	style.SetColor(imgui.StyleColorButtonActive, imgui.Vec4{X: 0.25, Y: 0.38, Z: 0.49, W: 1.00})
	// style.SetColor(imgui.StyleColorHeader, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorHeaderHovered, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorHeaderActive, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorSeparator, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorSeparatorHovered, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorSeparatorActive, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorResizeGrip, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorResizeGripHovered, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorResizeGripActive, imgui.Vec4{})
	style.SetColor(imgui.StyleColorTab, imgui.Vec4{X: 0.19, Y: 0.33, Z: 0.44, W: 1.00})
	style.SetColor(imgui.StyleColorTabHovered, imgui.Vec4{X: 0.23, Y: 0.36, Z: 0.47, W: 1.00})
	style.SetColor(imgui.StyleColorTabActive, imgui.Vec4{X: 0.25, Y: 0.38, Z: 0.49, W: 1.00})
	// style.SetColor(imgui.StyleColorTabUnfocused, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorTabUnfocusedActive, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorPlotLines, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorPlotLinesHovered, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorPlotHistogram, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorPlotHistogramHovered, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorTextSelectedBg, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorDragDropTarget, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorNavHighlight, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorNavWindowingHighlight, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorNavWindowingDarkening, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorModalWindowDarkening, imgui.Vec4{})

	scale := w.platform.GetContentScale()

	style.ScaleAllSizes(scale)
}

// Set background color of master window.
func (w *MasterWindow) SetBgColor(color color.RGBA) {
	w.clearColor = [4]float32{float32(color.R) / 255.0, float32(color.G) / 255.0, float32(color.B) / 255.0, float32(color.A) / 255.0}
}

func (w *MasterWindow) sizeChange(width, height int) {
	w.render()
}

func (w *MasterWindow) render() {
	p := w.platform
	r := w.renderer

	p.NewFrame()
	imgui.NewFrame()

	w.updateFunc()

	imgui.Render()
	r.PreRender(w.clearColor)
	r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())
	p.PostRender()
}

// Run the main loop to create new frame, process events and call update ui func.
func (w *MasterWindow) run() {
	p := w.platform

	ticker := time.NewTicker(time.Second / 60)
	shouldQuit := false
	for !shouldQuit {
		Call(func() {
			p.ProcessEvents()
			w.render()

			shouldQuit = p.ShouldStop()
		})

		<-ticker.C
	}
}

// Return size of master window.
func (w *MasterWindow) GetSize() (width, height int) {
	if w.platform != nil {
		if glfwPlatform, ok := w.platform.(*imgui.GLFW); ok {
			return glfwPlatform.GetWindow().GetSize()
		}
	}

	return w.width, w.height
}

func (w *MasterWindow) SetDropCallback(cb func([]string)) {
	w.platform.SetDropCallback(cb)
}

// Call the main loop.
// loopFunc will be used to construct the ui.
func (w *MasterWindow) Main(loopFunc func()) {
	Run(func() {
		w.updateFunc = loopFunc

		Context.isAlive = true

		w.run()

		Context.isAlive = false

		Call(func() {
			w.renderer.Dispose()
			w.platform.Dispose()
			w.context.Destroy()
		})
	})
}
