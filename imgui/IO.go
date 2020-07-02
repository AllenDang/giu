package imgui

// #include "IOWrapper.h"
import "C"

// IO is where your app communicate with ImGui. Access via CurrentIO().
// Read 'Programmer guide' section in imgui.cpp file for general usage.
type IO struct {
	handle C.IggIO
}

// WantCaptureMouse returns true if imgui will use the mouse inputs.
// Do not dispatch them to your main game/application in this case.
// In either case, always pass on mouse inputs to imgui.
// (e.g. unclicked mouse is hovering over an imgui window, widget is active, mouse was clicked over an imgui window, etc.)
func (io IO) WantCaptureMouse() bool {
	return C.iggWantCaptureMouse(io.handle) != 0
}

// WantCaptureKeyboard returns true if imgui will use the keyboard inputs.
// Do not dispatch them to your main game/application (in both cases, always pass keyboard inputs to imgui).
// (e.g. InputText active, or an imgui window is focused and navigation is enabled, etc.).
func (io IO) WantCaptureKeyboard() bool {
	return C.iggWantCaptureKeyboard(io.handle) != 0
}

// WantTextInput is true, you may display an on-screen keyboard.
// This is set by ImGui when it wants textual keyboard input to happen (e.g. when a InputText widget is active).
func (io IO) WantTextInput() bool {
	return C.iggWantTextInput(io.handle) != 0
}

// SetDisplaySize sets the size in pixels.
func (io IO) SetDisplaySize(value Vec2) {
	out, _ := value.wrapped()
	C.iggIoSetDisplaySize(io.handle, out)
}

func (io IO) GetMouseDrawCursor() bool {
	return C.iggIoGetMouseDrawCursor(io.handle) != 0
}

func (io IO) SetMouseDrawCursor(b bool) {
	C.iggIoSetMouseDrawCursor(io.handle, castBool(b))
}

// Fonts returns the font atlas to load and assemble one or more fonts into a single tightly packed texture.
func (io IO) Fonts() FontAtlas {
	return FontAtlas(C.iggIoGetFonts(io.handle))
}

// SetMousePosition sets the mouse position, in pixels.
// Set to Vec2(-math.MaxFloat32,-mathMaxFloat32) if mouse is unavailable (on another screen, etc.)
func (io IO) SetMousePosition(value Vec2) {
	posArg, _ := value.wrapped()
	C.iggIoSetMousePosition(io.handle, posArg)
}

// SetMouseButtonDown sets whether a specific mouse button is currently pressed.
// Mouse buttons: left, right, middle + extras.
// ImGui itself mostly only uses left button (BeginPopupContext** are using right button).
// Other buttons allows us to track if the mouse is being used by your application +
// available to user as a convenience via IsMouse** API.
func (io IO) SetMouseButtonDown(index int, down bool) {
	var downArg C.IggBool
	if down {
		downArg = 1
	}
	C.iggIoSetMouseButtonDown(io.handle, C.int(index), downArg)
}

// AddMouseWheelDelta adds the given offsets to the current mouse wheel values.
// 1 vertical unit scrolls about 5 lines text.
// Most users don't have a mouse with an horizontal wheel, may not be provided by all back-ends.
func (io IO) AddMouseWheelDelta(horizontal, vertical float32) {
	C.iggIoAddMouseWheelDelta(io.handle, C.float(horizontal), C.float(vertical))
}

func (io IO) GetMouseDelta() Vec2 {
	var delta Vec2
	deltaArg, deltaFin := delta.wrapped()
	C.iggIoGetMouseDelta(io.handle, deltaArg)
	deltaFin()
	return delta
}

// SetDeltaTime sets the time elapsed since last frame, in seconds.
func (io IO) SetDeltaTime(value float32) {
	C.iggIoSetDeltaTime(io.handle, C.float(value))
}

// SetFontGlobalScale sets the global scaling factor for all fonts.
func (io IO) SetFontGlobalScale(value float32) {
	C.iggIoSetFontGlobalScale(io.handle, C.float(value))
}

// KeyPress sets the KeysDown flag.
func (io IO) KeyPress(key int) {
	C.iggIoKeyPress(io.handle, C.int(key))
}

// KeyRelease clears the KeysDown flag.
func (io IO) KeyRelease(key int) {
	C.iggIoKeyRelease(io.handle, C.int(key))
}

// KeyMap maps a key into the KeysDown array which represents your "native" keyboard state.
func (io IO) KeyMap(imguiKey int, nativeKey int) {
	C.iggIoKeyMap(io.handle, C.int(imguiKey), C.int(nativeKey))
}

// KeyCtrl sets the keyboard modifier control pressed.
func (io IO) KeyCtrl(leftCtrl int, rightCtrl int) {
	C.iggIoKeyCtrl(io.handle, C.int(leftCtrl), C.int(rightCtrl))
}

// KeyShift sets the keyboard modifier shift pressed.
func (io IO) KeyShift(leftShift int, rightShift int) {
	C.iggIoKeyShift(io.handle, C.int(leftShift), C.int(rightShift))
}

// KeyAlt sets the keyboard modifier alt pressed.
func (io IO) KeyAlt(leftAlt int, rightAlt int) {
	C.iggIoKeyAlt(io.handle, C.int(leftAlt), C.int(rightAlt))
}

// KeySuper sets the keyboard modifier super pressed.
func (io IO) KeySuper(leftSuper int, rightSuper int) {
	C.iggIoKeySuper(io.handle, C.int(leftSuper), C.int(rightSuper))
}

// AddInputCharacters adds a new character into InputCharacters[].
func (io IO) AddInputCharacters(chars string) {
	textArg, textFin := wrapString(chars)
	defer textFin()
	C.iggIoAddInputCharactersUTF8(io.handle, textArg)
}

// SetIniFilename changes the filename for the settings. Default: "imgui.ini".
// Use an empty string to disable the ini from being used.
func (io IO) SetIniFilename(value string) {
	valueArg, valueFin := wrapString(value)
	defer valueFin()
	C.iggIoSetIniFilename(io.handle, valueArg)
}

// SetConfigFlags sets the gamepad/keyboard navigation options, etc.
func (io IO) SetConfigFlags(flags int) {
	C.iggIoSetConfigFlags(io.handle, C.int(flags))
}

func (io IO) GetConfigFlags() int {
	return int(C.iggIoGetConfigFlags(io.handle))
}

// SetBackendFlags sets back-end capabilities.
func (io IO) SetBackendFlags(flags int) {
	C.iggIoSetBackendFlags(io.handle, C.int(flags))
}

func (io IO) GetFrameCountSinceLastInput() int {
	return int(C.iggGetFrameCountSinceLastInput(io.handle))
}

func (io IO) SetFrameCountSinceLastInput(count int) {
	C.iggSetFrameCountSinceLastInput(io.handle, C.int(count))
}

// Clipboard describes the access to the text clipboard of the window manager.
type Clipboard interface {
	// Text returns the current text from the clipboard, if available.
	Text() (string, error)
	// SetText sets the text as the current text on the clipboard.
	SetText(value string)
}

var clipboards = map[C.IggIO]Clipboard{}
var dropLastClipboardText = func() {}

// SetClipboard registers a clipboard for text copy/paste actions.
// If no clipboard is set, then a fallback implementation may be used, if available for the OS.
// To disable clipboard handling overall, pass nil as the Clipboard.
//
// Since ImGui queries the clipboard text via a return value, the wrapper has to hold the
// current clipboard text as a copy in memory. This memory will be freed at the next clipboard operation.
func (io IO) SetClipboard(board Clipboard) {
	dropLastClipboardText()

	if board != nil {
		clipboards[io.handle] = board
		C.iggIoRegisterClipboardFunctions(io.handle)
	} else {
		C.iggIoClearClipboardFunctions(io.handle)
		delete(clipboards, io.handle)
	}
}

//export iggIoGetClipboardText
func iggIoGetClipboardText(handle C.IggIO) *C.char {
	dropLastClipboardText()

	board := clipboards[handle]
	if board == nil {
		return nil
	}

	text, err := board.Text()
	if err != nil {
		return nil
	}
	textPtr, textFin := wrapString(text)
	dropLastClipboardText = func() {
		dropLastClipboardText = func() {}
		textFin()
	}
	return textPtr
}

//export iggIoSetClipboardText
func iggIoSetClipboardText(handle C.IggIO, text *C.char) {
	dropLastClipboardText()

	board := clipboards[handle]
	if board == nil {
		return
	}
	board.SetText(C.GoString(text))
}
