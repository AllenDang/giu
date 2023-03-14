package giu

import imgui "github.com/AllenDang/cimgui-go"

// MouseButton represents imgui.MoseButton.
type MouseButton imgui.MouseButton

// mouse buttons.
const (
	MouseButtonLeft   MouseButton = imgui.MouseButtonLeft
	MouseButtonRight  MouseButton = imgui.MouseButtonRight
	MouseButtonMiddle MouseButton = imgui.MouseButtonMiddle
)

// IsItemHovered returns true if mouse is over the item.
func IsItemHovered() bool {
	return imgui.IsItemHovered()
}

// IsItemClicked returns true if mouse is clicked
// NOTE: if you're looking for clicking detection, see EventHandler.go.
func IsItemClicked(mouseButton MouseButton) bool {
	return imgui.IsItemClickedV(imgui.MouseButton(mouseButton))
}

// IsItemActive returns true if item is active.
func IsItemActive() bool {
	return imgui.IsItemActive()
}

// IsKeyDown returns true if key `key` is down.
func IsKeyDown(key Key) bool {
	return imgui.IsKeyDownNil(imgui.Key(key))
}

// IsKeyPressed returns true if key `key` is pressed.
func IsKeyPressed(key Key) bool {
	return imgui.IsKeyPressedBool(imgui.Key(key))
}

// IsKeyReleased returns true if key `key` is released.
func IsKeyReleased(key Key) bool {
	return imgui.IsKeyReleasedNil(imgui.Key(key))
}

// IsMouseDown returns true if mouse button `button` is down.
func IsMouseDown(button MouseButton) bool {
	return imgui.IsMouseDownNil(imgui.MouseButton(button))
}

// IsMouseClicked returns true if mouse button `button` is clicked
// NOTE: if you're looking for clicking detection, see EventHandler.go.
func IsMouseClicked(button MouseButton) bool {
	return imgui.IsMouseClickedBool(imgui.MouseButton(button))
}

// IsMouseReleased returns true if mouse button `button` is released.
func IsMouseReleased(button MouseButton) bool {
	return imgui.IsMouseReleasedNil(imgui.MouseButton(button))
}

// IsMouseDoubleClicked returns true if mouse button `button` is double clicked.
func IsMouseDoubleClicked(button MouseButton) bool {
	return imgui.IsMouseDoubleClicked(imgui.MouseButton(button))
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
	return imgui.IsWindowFocusedV(imgui.FocusedFlags(flags))
}

// IsWindowHovered returns true if the window is hovered.
func IsWindowHovered(flags HoveredFlags) bool {
	return imgui.IsWindowHoveredV(imgui.HoveredFlags(flags))
}
