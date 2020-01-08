package giu

import (
	"image"
	"image/color"
	"time"

	"github.com/AllenDang/giu/imgui"
)

type MasterWindow struct {
	width      int
	height     int
	clearColor [4]float32
	title      string
	resizable  bool
	platform   imgui.Platform
	renderer   imgui.Renderer
	context    *imgui.Context
	io         *imgui.IO
	updateFunc func(w *MasterWindow)
}

// Create a master window.
func NewMasterWindow(title string, width, height int, resizable bool, loadFontFunc func()) *MasterWindow {
	return NewMasterWindowWithBgColor(title, width, height, resizable, loadFontFunc, nil)
}

func NewMasterWindowWithBgColor(title string, width, height int, resizable bool, loadFontFunc func(), bgColor *color.RGBA) *MasterWindow {
	context := imgui.CreateContext(nil)

	io := imgui.CurrentIO()

	io.SetConfigFlags(imgui.ConfigFlagEnablePowerSavingMode)

	// Disable imgui.ini
	io.SetIniFilename("")

	if loadFontFunc != nil {
		loadFontFunc()
	}

	p, err := imgui.NewGLFW(io, title, width, height, resizable)
	if err != nil {
		panic(err)
	}

	r, err := imgui.NewOpenGL3(io)
	if err != nil {
		panic(err)
	}

	col := [4]float32{0.22, 0.26, 0.28, 1}

	if bgColor != nil {
		vec4 := ToVec4Color(*bgColor)
		col = [4]float32{vec4.X, vec4.Y, vec4.Z, vec4.W}
	}

	mw := &MasterWindow{
		clearColor: col,
		width:      width,
		height:     height,
		title:      title,
		resizable:  resizable,
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
	style.SetColor(imgui.StyleColorWindowBg, imgui.Vec4{X: 0.22, Y: 0.26, Z: 0.28, W: 0.84})
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
	// style.SetColor(imgui.StyleColorTab, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorTabHovered, imgui.Vec4{})
	// style.SetColor(imgui.StyleColorTabActive, imgui.Vec4{})
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

	w.updateFunc(w)

	imgui.Render()
	r.PreRender(w.clearColor)
	r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())
	p.PostRender()

}

// Run the main loop to create new frame, process events and call update ui func.
func (w *MasterWindow) run() {
	p := w.platform

	ticker := time.NewTicker(time.Second / 60)
	for !p.ShouldStop() {
		p.ProcessEvents()

		w.render()

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

// Load a image by renderer and return the TextureID which will be used later by imgui.
func (w *MasterWindow) LoadImage(image *image.RGBA) (imgui.TextureID, error) {
	return w.renderer.LoadImage(image)
}

func (w *MasterWindow) ReleaseImage(textureId imgui.TextureID) {
	w.renderer.ReleaseImage(textureId)
}

func (w *MasterWindow) Update() {
	w.platform.Update()
}

// Call the main loop.
// loopFunc will be used to construct the ui.
func (w *MasterWindow) Main(loopFunc func(w *MasterWindow)) {
	w.updateFunc = loopFunc

	w.run()

	w.renderer.Dispose()
	w.platform.Dispose()
	w.context.Destroy()
}
