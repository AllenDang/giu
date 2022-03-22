package giu

import "github.com/AllenDang/imgui-go"

// MouseButton represents imgui.MoseButton.
type MouseButton int

// mouse buttons.
const (
	MouseButtonLeft   MouseButton = 0
	MouseButtonRight  MouseButton = 1
	MouseButtonMiddle MouseButton = 2
)

// IsItemHovered returns true if mouse is over the item.
func IsItemHovered() bool {
	return imgui.IsItemHovered()
}

// IsItemClicked returns true if mouse is clicked
// NOTE: if you're looking for clicking detection, see EventHandler.go.
func IsItemClicked(mouseButton MouseButton) bool {
	return imgui.IsItemClicked(int(mouseButton))
}

// IsItemActive returns true if item is active.
func IsItemActive() bool {
	return imgui.IsItemActive()
}

// IsKeyDown returns true if key `key` is down.
func IsKeyDown(key Key) bool {
	return imgui.IsKeyDown(int(key))
}

// IsKeyPressed returns true if key `key` is pressed.
func IsKeyPressed(key Key) bool {
	return imgui.IsKeyPressed(int(key))
}

// IsKeyReleased returns true if key `key` is released.
func IsKeyReleased(key Key) bool {
	return imgui.IsKeyReleased(int(key))
}

// IsMouseDown returns true if mouse button `button` is down.
func IsMouseDown(button MouseButton) bool {
	return imgui.IsMouseDown(int(button))
}

// IsMouseClicked returns true if mouse button `button` is clicked
// NOTE: if you're looking for clicking detection, see EventHandler.go.
func IsMouseClicked(button MouseButton) bool {
	return imgui.IsMouseClicked(int(button))
}

// IsMouseReleased returns true if mouse button `button` is released.
func IsMouseReleased(button MouseButton) bool {
	return imgui.IsMouseReleased(int(button))
}

// IsMouseDoubleClicked returns true if mouse button `button` is double clicked.
func IsMouseDoubleClicked(button MouseButton) bool {
	return imgui.IsMouseDoubleClicked(int(button))
}

// IsWindowAppearing returns true if window is appearing.
func IsWindowAppearing() bool {
	return imgui.IsWindowAppearing()
}

// IsWindowCollapsed returns true if window is disappearing.
func IsWindowCollapsed() bool {
	return imgui.IsWindowCollapsed()
}

// IsWindowFocused returns true if window is focused
// NOTE: see also (*Window).HasFocus and (*Window).BringToFront.
func IsWindowFocused(flags FocusedFlags) bool {
	return imgui.IsWindowFocused(int(flags))
}

// IsWindowHovered returns true if the window is hovered.
func IsWindowHovered(flags HoveredFlags) bool {
	return imgui.IsWindowHovered(int(flags))
}
