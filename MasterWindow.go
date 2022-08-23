package giu

import (
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

	Context = CreateContext(window)

	mw := &MasterWindow{
		title:  title,
		window: window,
	}

	mw.setTheme()

	return mw
}

func (w *MasterWindow) setTheme() {
	imgui.PushStyleVar_Float(imgui.ImGuiStyleVar_WindowRounding, 2)
	imgui.PushStyleVar_Float(imgui.ImGuiStyleVar_FrameRounding, 4)
	imgui.PushStyleVar_Float(imgui.ImGuiStyleVar_GrabRounding, 4)
	imgui.PushStyleVar_Float(imgui.ImGuiStyleVar_FrameBorderSize, 1)

	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_Text, imgui.ImVec4{X: 0.95, Y: 0.96, Z: 0.98, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TextDisabled, imgui.ImVec4{X: 0.36, Y: 0.42, Z: 0.47, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_WindowBg, imgui.ImVec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_ChildBg, imgui.ImVec4{X: 0.15, Y: 0.18, Z: 0.22, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_PopupBg, imgui.ImVec4{X: 0.08, Y: 0.08, Z: 0.08, W: 0.94})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_Border, imgui.ImVec4{X: 0.08, Y: 0.10, Z: 0.12, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_BorderShadow, imgui.ImVec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_FrameBg, imgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_FrameBgHovered, imgui.ImVec4{X: 0.12, Y: 0.20, Z: 0.28, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_FrameBgActive, imgui.ImVec4{X: 0.09, Y: 0.12, Z: 0.14, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TitleBg, imgui.ImVec4{X: 0.09, Y: 0.12, Z: 0.14, W: 0.65})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TitleBgActive, imgui.ImVec4{X: 0.08, Y: 0.10, Z: 0.12, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TitleBgCollapsed, imgui.ImVec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.51})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_MenuBarBg, imgui.ImVec4{X: 0.15, Y: 0.18, Z: 0.22, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_ScrollbarBg, imgui.ImVec4{X: 0.02, Y: 0.02, Z: 0.02, W: 0.39})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_ScrollbarGrab, imgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_ScrollbarGrabHovered, imgui.ImVec4{X: 0.18, Y: 0.22, Z: 0.25, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_ScrollbarGrabActive, imgui.ImVec4{X: 0.09, Y: 0.21, Z: 0.31, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_CheckMark, imgui.ImVec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_SliderGrab, imgui.ImVec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_SliderGrabActive, imgui.ImVec4{X: 0.37, Y: 0.61, Z: 1.00, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_Button, imgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_ButtonHovered, imgui.ImVec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_ButtonActive, imgui.ImVec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_Header, imgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 0.55})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_HeaderHovered, imgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_HeaderActive, imgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_Separator, imgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_SeparatorHovered, imgui.ImVec4{X: 0.10, Y: 0.40, Z: 0.75, W: 0.78})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_SeparatorActive, imgui.ImVec4{X: 0.10, Y: 0.40, Z: 0.75, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_ResizeGrip, imgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.25})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_ResizeGripHovered, imgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.67})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_ResizeGripActive, imgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.95})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_Tab, imgui.ImVec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TabHovered, imgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TabActive, imgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TabUnfocused, imgui.ImVec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TabUnfocusedActive, imgui.ImVec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_PlotLines, imgui.ImVec4{X: 0.61, Y: 0.61, Z: 0.61, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_PlotLinesHovered, imgui.ImVec4{X: 1.00, Y: 0.43, Z: 0.35, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_PlotHistogram, imgui.ImVec4{X: 0.90, Y: 0.70, Z: 0.00, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_PlotHistogramHovered, imgui.ImVec4{X: 1.00, Y: 0.60, Z: 0.00, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TextSelectedBg, imgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.35})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_DragDropTarget, imgui.ImVec4{X: 1.00, Y: 1.00, Z: 0.00, W: 0.90})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_NavHighlight, imgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_NavWindowingHighlight, imgui.ImVec4{X: 1.00, Y: 1.00, Z: 1.00, W: 0.70})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TableHeaderBg, imgui.ImVec4{X: 0.12, Y: 0.20, Z: 0.28, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TableBorderStrong, imgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	imgui.PushStyleColor_Vec4(imgui.ImGuiCol_TableBorderLight, imgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 0.70})
}

func (w *MasterWindow) beforeRender() {
	Context.invalidAllState()
	Context.FontAtlas.rebuildFontAtlas()
}

func (w *MasterWindow) afterRender() {
	Context.cleanState()
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

	w.window.Run(loopFunc, w.beforeRender, w.afterRender)
}
