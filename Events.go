package giu

import (
	"github.com/AllenDang/cimgui-go"
)

// MouseButton represents cimgui.MoseButton.
type MouseButton cimgui.ImGuiMouseButton

// mouse buttons.
const (
	MouseButtonLeft   MouseButton = cimgui.ImGuiMouseButton_Left
	MouseButtonRight  MouseButton = cimgui.ImGuiMouseButton_Right
	MouseButtonMiddle MouseButton = cimgui.ImGuiMouseButton_Middle
)

// IsItemHovered returns true if mouse is over the item.
func IsItemHovered() bool {
	return cimgui.IsItemHovered()
}

// IsItemClicked returns true if mouse is clicked
// NOTE: if you're looking for clicking detection, see EventHandler.go.
func IsItemClicked(mouseButton MouseButton) bool {
	return cimgui.IsItemClickedV(cimgui.ImGuiMouseButton(mouseButton))
}

// IsItemActive returns true if item is active.
func IsItemActive() bool {
	return cimgui.IsItemActive()
}

// IsKeyDown returns true if key `key` is down.
func IsKeyDown(key Key) bool {
	return cimgui.IsKeyDown(cimgui.ImGuiKey(key))
}

// IsKeyPressed returns true if key `key` is pressed.
func IsKeyPressed(key Key) bool {
	return cimgui.IsKeyPressed(cimgui.ImGuiKey(key))
}

// IsKeyReleased returns true if key `key` is released.
func IsKeyReleased(key Key) bool {
	return cimgui.IsKeyReleased(cimgui.ImGuiKey(key))
}

// IsMouseDown returns true if mouse button `button` is down.
func IsMouseDown(button MouseButton) bool {
	return cimgui.IsMouseDown(cimgui.ImGuiMouseButton(button))
}

// IsMouseClicked returns true if mouse button `button` is clicked
// NOTE: if you're looking for clicking detection, see EventHandler.go.
func IsMouseClicked(button MouseButton) bool {
	return cimgui.IsMouseClicked(cimgui.ImGuiMouseButton(button))
}

// IsMouseReleased returns true if mouse button `button` is released.
func IsMouseReleased(button MouseButton) bool {
	return cimgui.IsMouseReleased(cimgui.ImGuiMouseButton(button))
}

// IsMouseDoubleClicked returns true if mouse button `button` is double clicked.
func IsMouseDoubleClicked(button MouseButton) bool {
	return cimgui.IsMouseDoubleClicked(cimgui.ImGuiMouseButton(button))
}

// IsWindowAppearing returns true if window is appearing.
func IsWindowAppearing() bool {
	return cimgui.IsWindowAppearing()
}

// IsWindowCollapsed returns true if window is disappearing.
func IsWindowCollapsed() bool {
	return cimgui.IsWindowCollapsed()
}

// IsWindowFocused returns true if window is focused
// NOTE: see also (*Window).HasFocus and (*Window).BringToFront.
func IsWindowFocused(flags FocusedFlags) bool {
	return cimgui.IsWindowFocusedV(cimgui.ImGuiFocusedFlags(flags))
}

// IsWindowHovered returns true if the window is hovered.
func IsWindowHovered(flags HoveredFlags) bool {
	return cimgui.IsWindowHoveredV(cimgui.ImGuiHoveredFlags(flags))
}
