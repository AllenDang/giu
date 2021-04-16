package imgui

// Enumeration for MouseCursor()
// User code may request binding to display given cursor by calling SetMouseCursor(), which is why we have some cursors that are marked unused here
const (
	// MouseCursorNone no mouse cursor
	MouseCursorNone = -1
	// MouseCursorArrow standard arrow mouse cursor
	MouseCursorArrow = 0
	// MouseCursorTextInput when hovering over InputText, etc.
	MouseCursorTextInput = 1
	// MouseCursorResizeAll (Unused by imgui functions)
	MouseCursorResizeAll = 2
	// MouseCursorResizeNS when hovering over an horizontal border
	MouseCursorResizeNS = 3
	// MouseCursorResizeEW when hovering over a vertical border or a column
	MouseCursorResizeEW = 4
	// MouseCursorResizeNESW when hovering over the bottom-left corner of a window
	MouseCursorResizeNESW = 5
	// MouseCursorResizeNWSE when hovering over the bottom-right corner of a window
	MouseCursorResizeNWSE = 6
	// MouseCursorHand (Unused by imgui functions. Use for e.g. hyperlinks)
	MouseCursorHand  = 7
	MouseCursorCount = 8
)
