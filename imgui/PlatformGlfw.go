package imgui

import (
	"fmt"
	"math"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type GLFWClipboard struct {
	window *glfw.Window
}

func NewGLFWClipboard(w *glfw.Window) *GLFWClipboard {
	return &GLFWClipboard{window: w}
}

func (c *GLFWClipboard) Text() (string, error) {
	return c.window.GetClipboardString(), nil
}

func (c *GLFWClipboard) SetText(text string) {
	c.window.SetClipboardString(text)
}

// GLFW implements a platform based on github.com/go-gl/glfw (v3.2).
type GLFW struct {
	imguiIO IO

	window *glfw.Window

	time             float64
	mouseJustPressed [3]bool

	mouseCursors map[int]*glfw.Cursor

	sizeChangeCallback func(int, int)
	dropCallback       func([]string)
}

// NewGLFW attempts to initialize a GLFW context.
func NewGLFW(io IO, title string, width, height int, resizable bool) (*GLFW, error) {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize glfw: %v", err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)
	glfw.WindowHint(glfw.ScaleToMonitor, glfw.True)

	if !resizable {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		glfw.Terminate()
		return nil, fmt.Errorf("failed to create window: %v", err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	platform := &GLFW{
		imguiIO: io,
		window:  window,
	}
	platform.setKeyMapping()
	platform.installCallbacks()

	// Create mosue cursors
	platform.mouseCursors = make(map[int]*glfw.Cursor)
	platform.mouseCursors[MouseCursorArrow] = glfw.CreateStandardCursor(glfw.ArrowCursor)
	platform.mouseCursors[MouseCursorTextInput] = glfw.CreateStandardCursor(glfw.IBeamCursor)
	platform.mouseCursors[MouseCursorResizeAll] = glfw.CreateStandardCursor(glfw.CrosshairCursor)
	platform.mouseCursors[MouseCursorHand] = glfw.CreateStandardCursor(glfw.HandCursor)
	platform.mouseCursors[MouseCursorResizeEW] = glfw.CreateStandardCursor(glfw.HResizeCursor)
	platform.mouseCursors[MouseCursorResizeNS] = glfw.CreateStandardCursor(glfw.VResizeCursor)

	io.SetClipboard(NewGLFWClipboard(window))

	return platform, nil
}

// Dispose cleans up the resources.
func (platform *GLFW) Dispose() {
	platform.window.Destroy()
	glfw.Terminate()
}

func (platform *GLFW) GetContentScale() float32 {
	x, _ := platform.window.GetContentScale()
	return x
}

func (platform *GLFW) GetWindow() *glfw.Window {
	return platform.window
}

// ShouldStop returns true if the window is to be closed.
func (platform *GLFW) ShouldStop() bool {
	return platform.window.ShouldClose()
}

func (platform *GLFW) WaitForEvent() {
	if platform.imguiIO.GetConfigFlags()&ConfigFlagEnablePowerSavingMode == 0 {
		return
	}

	windowIsHidden := platform.window.GetAttrib(glfw.Visible) == glfw.False || platform.window.GetAttrib(glfw.Iconified) == glfw.True

	waitingTime := math.Inf(0)

	if !windowIsHidden {
		waitingTime = GetEventWaitingTime()
	}

	if waitingTime > 0 {
		if math.IsInf(waitingTime, 0) {
			glfw.WaitEvents()
		} else {
			glfw.WaitEventsTimeout(waitingTime)
		}
	}
}

// ProcessEvents handles all pending window events.
func (platform *GLFW) ProcessEvents() {
	platform.WaitForEvent()
	glfw.PollEvents()
}

// DisplaySize returns the dimension of the display.
func (platform *GLFW) DisplaySize() [2]float32 {
	w, h := platform.window.GetSize()
	return [2]float32{float32(w), float32(h)}
}

// FramebufferSize returns the dimension of the framebuffer.
func (platform *GLFW) FramebufferSize() [2]float32 {
	w, h := platform.window.GetFramebufferSize()
	return [2]float32{float32(w), float32(h)}
}

// NewFrame marks the begin of a render pass. It forwards all current state to imgui IO.
func (platform *GLFW) NewFrame() {
	// Setup display size (every frame to accommodate for window resizing)
	displaySize := platform.DisplaySize()
	platform.imguiIO.SetDisplaySize(Vec2{X: displaySize[0], Y: displaySize[1]})

	// Setup time step
	currentTime := glfw.GetTime()
	if platform.time > 0 {
		platform.imguiIO.SetDeltaTime(float32(currentTime - platform.time))
	}
	platform.time = currentTime

	// Setup inputs
	if platform.window.GetAttrib(glfw.Focused) != 0 {
		x, y := platform.window.GetCursorPos()
		platform.imguiIO.SetMousePosition(Vec2{X: float32(x), Y: float32(y)})
	} else {
		platform.imguiIO.SetMousePosition(Vec2{X: -math.MaxFloat32, Y: -math.MaxFloat32})
	}

	for i := 0; i < len(platform.mouseJustPressed); i++ {
		down := platform.mouseJustPressed[i] || (platform.window.GetMouseButton(glfwButtonIDByIndex[i]) == glfw.Press)
		platform.imguiIO.SetMouseButtonDown(i, down)
		platform.mouseJustPressed[i] = false
	}

	platform.updateMouseCursor()
}

// PostRender performs a buffer swap.
func (platform *GLFW) PostRender() {
	platform.window.SwapBuffers()
}

func (platform *GLFW) SetSizeChangeCallback(cb func(int, int)) {
	platform.sizeChangeCallback = cb
}

func (platform *GLFW) Update() {
	glfw.PostEmptyEvent()
}

func (platform *GLFW) SetDropCallback(cb func(names []string)) {
	platform.dropCallback = cb
}

func (platform *GLFW) updateMouseCursor() {
	io := platform.imguiIO
	if (io.GetConfigFlags()&ConfigFlagNoMouseCursorChange) == 1 || platform.window.GetInputMode(glfw.CursorMode) == glfw.CursorDisabled {
		return
	}

	cursor := MouseCursor()
	if cursor == MouseCursorNone || io.GetMouseDrawCursor() {
		platform.window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	} else {
		gCursor := platform.mouseCursors[MouseCursorArrow]
		if c, ok := platform.mouseCursors[cursor]; ok {
			gCursor = c
		}
		platform.window.SetCursor(gCursor)
		platform.window.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	}
}

func (platform *GLFW) setKeyMapping() {
	// Keyboard mapping. ImGui will use those indices to peek into the io.KeysDown[] array.
	platform.imguiIO.KeyMap(KeyTab, int(glfw.KeyTab))
	platform.imguiIO.KeyMap(KeyLeftArrow, int(glfw.KeyLeft))
	platform.imguiIO.KeyMap(KeyRightArrow, int(glfw.KeyRight))
	platform.imguiIO.KeyMap(KeyUpArrow, int(glfw.KeyUp))
	platform.imguiIO.KeyMap(KeyDownArrow, int(glfw.KeyDown))
	platform.imguiIO.KeyMap(KeyPageUp, int(glfw.KeyPageUp))
	platform.imguiIO.KeyMap(KeyPageDown, int(glfw.KeyPageDown))
	platform.imguiIO.KeyMap(KeyHome, int(glfw.KeyHome))
	platform.imguiIO.KeyMap(KeyEnd, int(glfw.KeyEnd))
	platform.imguiIO.KeyMap(KeyInsert, int(glfw.KeyInsert))
	platform.imguiIO.KeyMap(KeyDelete, int(glfw.KeyDelete))
	platform.imguiIO.KeyMap(KeyBackspace, int(glfw.KeyBackspace))
	platform.imguiIO.KeyMap(KeySpace, int(glfw.KeySpace))
	platform.imguiIO.KeyMap(KeyEnter, int(glfw.KeyEnter))
	platform.imguiIO.KeyMap(KeyEscape, int(glfw.KeyEscape))
	platform.imguiIO.KeyMap(KeyA, int(glfw.KeyA))
	platform.imguiIO.KeyMap(KeyC, int(glfw.KeyC))
	platform.imguiIO.KeyMap(KeyV, int(glfw.KeyV))
	platform.imguiIO.KeyMap(KeyX, int(glfw.KeyX))
	platform.imguiIO.KeyMap(KeyY, int(glfw.KeyY))
	platform.imguiIO.KeyMap(KeyZ, int(glfw.KeyZ))
}

func (platform *GLFW) installCallbacks() {
	platform.window.SetMouseButtonCallback(platform.mouseButtonChange)
	platform.window.SetScrollCallback(platform.mouseScrollChange)
	platform.window.SetKeyCallback(platform.keyChange)
	platform.window.SetCharCallback(platform.charChange)
	platform.window.SetSizeCallback(platform.sizeChange)
	platform.window.SetDropCallback(platform.onDrop)
}

var glfwButtonIndexByID = map[glfw.MouseButton]int{
	glfw.MouseButton1: 0,
	glfw.MouseButton2: 1,
	glfw.MouseButton3: 2,
}

var glfwButtonIDByIndex = map[int]glfw.MouseButton{
	0: glfw.MouseButton1,
	1: glfw.MouseButton2,
	2: glfw.MouseButton3,
}

func (platform *GLFW) onDrop(window *glfw.Window, names []string) {
	window.Focus()

	if platform.dropCallback != nil {
		platform.dropCallback(names)
	}
}

func (platform *GLFW) sizeChange(window *glfw.Window, width, height int) {
	platform.imguiIO.SetFrameCountSinceLastInput(0)

	// Notify size changed and redraw.
	if platform.sizeChangeCallback != nil {
		platform.sizeChangeCallback(width, height)
	}
}

func (platform *GLFW) mouseButtonChange(window *glfw.Window, rawButton glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	platform.imguiIO.SetFrameCountSinceLastInput(0)

	buttonIndex, known := glfwButtonIndexByID[rawButton]

	if known && (action == glfw.Press) {
		platform.mouseJustPressed[buttonIndex] = true
	}
}

func (platform *GLFW) mouseScrollChange(window *glfw.Window, x, y float64) {
	platform.imguiIO.SetFrameCountSinceLastInput(0)
	platform.imguiIO.AddMouseWheelDelta(float32(x), float32(y))
}

func (platform *GLFW) keyChange(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	platform.imguiIO.SetFrameCountSinceLastInput(0)

	if action == glfw.Press {
		platform.imguiIO.KeyPress(int(key))
	}
	if action == glfw.Release {
		platform.imguiIO.KeyRelease(int(key))
	}

	// Modifiers are not reliable across systems
	platform.imguiIO.KeyCtrl(int(glfw.KeyLeftControl), int(glfw.KeyRightControl))
	platform.imguiIO.KeyShift(int(glfw.KeyLeftShift), int(glfw.KeyRightShift))
	platform.imguiIO.KeyAlt(int(glfw.KeyLeftAlt), int(glfw.KeyRightAlt))
	platform.imguiIO.KeySuper(int(glfw.KeyLeftSuper), int(glfw.KeyRightSuper))
}

func (platform *GLFW) charChange(window *glfw.Window, char rune) {
	platform.imguiIO.SetFrameCountSinceLastInput(0)
	platform.imguiIO.AddInputCharacters(string(char))
}
