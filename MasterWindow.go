package giu

import (
	"image/color"
	"time"

	"github.com/faiface/mainthread"

	"github.com/AllenDang/imgui-go"
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

func NewMasterWindow(title string, width, height int, flags MasterWindowFlags) *MasterWindow {
	context := imgui.CreateContext(nil)
	imgui.ImPlotCreateContext()
	imgui.ImNodesCreateContext()

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

	// Assign platform to contex
	Context.platform = p

	r, err := imgui.NewOpenGL3(io, scale)
	if err != nil {
		panic(err)
	}

	// Create context
	Context.renderer = r

	// Create font
	if len(defaultFonts) == 0 {
		io.Fonts().AddFontDefault()
		fontAtlas := io.Fonts().TextureDataRGBA32()
		r.SetFontTexture(fontAtlas)
	} else {
		shouldRebuildFontAtlas = true
		rebuildFontAtlas()
	}

	mw := &MasterWindow{
		clearColor: [4]float32{0, 0, 0, 1},
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
	imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 4)
	imgui.PushStyleVarFloat(imgui.StyleVarGrabRounding, 4)
	imgui.PushStyleVarFloat(imgui.StyleVarFrameBorderSize, 1)

	style.SetColor(imgui.StyleColorText, imgui.Vec4{X: 0.95, Y: 0.96, Z: 0.98, W: 1.00})
	style.SetColor(imgui.StyleColorTextDisabled, imgui.Vec4{X: 0.36, Y: 0.42, Z: 0.47, W: 1.00})
	style.SetColor(imgui.StyleColorWindowBg, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	style.SetColor(imgui.StyleColorChildBg, imgui.Vec4{X: 0.15, Y: 0.18, Z: 0.22, W: 1.00})
	style.SetColor(imgui.StyleColorPopupBg, imgui.Vec4{X: 0.08, Y: 0.08, Z: 0.08, W: 0.94})
	style.SetColor(imgui.StyleColorBorder, imgui.Vec4{X: 0.08, Y: 0.10, Z: 0.12, W: 1.00})
	style.SetColor(imgui.StyleColorBorderShadow, imgui.Vec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.00})
	style.SetColor(imgui.StyleColorFrameBg, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(imgui.StyleColorFrameBgHovered, imgui.Vec4{X: 0.12, Y: 0.20, Z: 0.28, W: 1.00})
	style.SetColor(imgui.StyleColorFrameBgActive, imgui.Vec4{X: 0.09, Y: 0.12, Z: 0.14, W: 1.00})
	style.SetColor(imgui.StyleColorTitleBg, imgui.Vec4{X: 0.09, Y: 0.12, Z: 0.14, W: 0.65})
	style.SetColor(imgui.StyleColorTitleBgActive, imgui.Vec4{X: 0.08, Y: 0.10, Z: 0.12, W: 1.00})
	style.SetColor(imgui.StyleColorTitleBgCollapsed, imgui.Vec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.51})
	style.SetColor(imgui.StyleColorMenuBarBg, imgui.Vec4{X: 0.15, Y: 0.18, Z: 0.22, W: 1.00})
	style.SetColor(imgui.StyleColorScrollbarBg, imgui.Vec4{X: 0.02, Y: 0.02, Z: 0.02, W: 0.39})
	style.SetColor(imgui.StyleColorScrollbarGrab, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(imgui.StyleColorScrollbarGrabHovered, imgui.Vec4{X: 0.18, Y: 0.22, Z: 0.25, W: 1.00})
	style.SetColor(imgui.StyleColorScrollbarGrabActive, imgui.Vec4{X: 0.09, Y: 0.21, Z: 0.31, W: 1.00})
	style.SetColor(imgui.StyleColorCheckMark, imgui.Vec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	style.SetColor(imgui.StyleColorSliderGrab, imgui.Vec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	style.SetColor(imgui.StyleColorSliderGrabActive, imgui.Vec4{X: 0.37, Y: 0.61, Z: 1.00, W: 1.00})
	style.SetColor(imgui.StyleColorButton, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(imgui.StyleColorButtonHovered, imgui.Vec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	style.SetColor(imgui.StyleColorButtonActive, imgui.Vec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00})
	style.SetColor(imgui.StyleColorHeader, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 0.55})
	style.SetColor(imgui.StyleColorHeaderHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80})
	style.SetColor(imgui.StyleColorHeaderActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00})
	style.SetColor(imgui.StyleColorSeparator, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(imgui.StyleColorSeparatorHovered, imgui.Vec4{X: 0.10, Y: 0.40, Z: 0.75, W: 0.78})
	style.SetColor(imgui.StyleColorSeparatorActive, imgui.Vec4{X: 0.10, Y: 0.40, Z: 0.75, W: 1.00})
	style.SetColor(imgui.StyleColorResizeGrip, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.25})
	style.SetColor(imgui.StyleColorResizeGripHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.67})
	style.SetColor(imgui.StyleColorResizeGripActive, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.95})
	style.SetColor(imgui.StyleColorTab, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	style.SetColor(imgui.StyleColorTabHovered, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80})
	style.SetColor(imgui.StyleColorTabActive, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(imgui.StyleColorTabUnfocused, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	style.SetColor(imgui.StyleColorTabUnfocusedActive, imgui.Vec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	style.SetColor(imgui.StyleColorPlotLines, imgui.Vec4{X: 0.61, Y: 0.61, Z: 0.61, W: 1.00})
	style.SetColor(imgui.StyleColorPlotLinesHovered, imgui.Vec4{X: 1.00, Y: 0.43, Z: 0.35, W: 1.00})
	style.SetColor(imgui.StyleColorPlotHistogram, imgui.Vec4{X: 0.90, Y: 0.70, Z: 0.00, W: 1.00})
	style.SetColor(imgui.StyleColorPlotHistogramHovered, imgui.Vec4{X: 1.00, Y: 0.60, Z: 0.00, W: 1.00})
	style.SetColor(imgui.StyleColorTextSelectedBg, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.35})
	style.SetColor(imgui.StyleColorDragDropTarget, imgui.Vec4{X: 1.00, Y: 1.00, Z: 0.00, W: 0.90})
	style.SetColor(imgui.StyleColorNavHighlight, imgui.Vec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00})
	style.SetColor(imgui.StyleColorNavWindowingHighlight, imgui.Vec4{X: 1.00, Y: 1.00, Z: 1.00, W: 0.70})
	style.SetColor(imgui.StyleColorTableHeaderBg, imgui.Vec4{X: 0.12, Y: 0.20, Z: 0.28, W: 1.00})
	style.SetColor(imgui.StyleColorTableBorderStrong, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(imgui.StyleColorTableBorderLight, imgui.Vec4{X: 0.20, Y: 0.25, Z: 0.29, W: 0.70})

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
	Context.invalidAllState()

	rebuildFontAtlas()

	p := w.platform
	r := w.renderer

	p.NewFrame()
	r.PreRender(w.clearColor)

	imgui.NewFrame()
	w.updateFunc()
	imgui.Render()

	r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())
	p.PostRender()

	Context.cleanState()
}

// Run the main loop to create new frame, process events and call update ui func.
func (w *MasterWindow) run() {
	p := w.platform

	ticker := time.NewTicker(time.Second / time.Duration(p.GetTPS()))
	shouldQuit := false
	for !shouldQuit {
		mainthread.Call(func() {
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

// Return position of master window.
func (w *MasterWindow) GetPos() (x, y int) {
	if w.platform != nil {
		if glfwPlatform, ok := w.platform.(*imgui.GLFW); ok {
			x, y = glfwPlatform.GetWindow().GetPos()
		}
	}

	return
}

// Set position of master window.
func (w *MasterWindow) SetPos(x, y int) {
	if w.platform != nil {
		if glfwPlatform, ok := w.platform.(*imgui.GLFW); ok {
			glfwPlatform.GetWindow().SetPos(x, y)
		}
	}
}

func (w *MasterWindow) SetSize(x, y int) {
	if w.platform != nil {
		if glfwPlatform, ok := w.platform.(*imgui.GLFW); ok {
			mainthread.CallNonBlock(func() {
				glfwPlatform.GetWindow().SetSize(x, y)
			})
		}
	}
}

func (w *MasterWindow) SetDropCallback(cb func([]string)) {
	w.platform.SetDropCallback(cb)
}

// Call the main loop.
// loopFunc will be used to construct the ui.
func (w *MasterWindow) Run(loopFunc func()) {
	mainthread.Run(func() {
		w.updateFunc = loopFunc

		Context.isAlive = true

		w.run()

		Context.isAlive = false

		mainthread.Call(func() {
			w.renderer.Dispose()
			w.platform.Dispose()

			imgui.ImNodesDestroyContext()
			imgui.ImPlotDestroyContext()
			w.context.Destroy()
		})
	})
}
