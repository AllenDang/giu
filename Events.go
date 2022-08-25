package giu

import imgui "github.com/AllenDang/cimgui-go"

// IsItemHovered returns true if mouse is over the item.
func IsItemHovered() bool {
	return imgui.IsItemHovered(0)
}

// IsItemClicked returns true if mouse is clicked
// NOTE: if you're looking for clicking detection, see EventHandler.go.
func IsItemClicked(mouseButton imgui.ImGuiMouseButton) bool {
	return imgui.IsItemClicked(mouseButton)
}

// IsItemActive returns true if item is active.
func IsItemActive() bool {
	return imgui.IsItemActive()
}

// IsKeyDown returns true if key `key` is down.
func IsKeyDown(key imgui.ImGuiKey) bool {
	return imgui.IsKeyDown(key)
}

// IsKeyPressed returns true if key `key` is pressed.
func IsKeyPressed(key imgui.ImGuiKey) bool {
	return imgui.IsKeyPressed(key, false)
}

// IsKeyReleased returns true if key `key` is released.
func IsKeyReleased(key imgui.ImGuiKey) bool {
	return imgui.IsKeyReleased(key)
}

// IsMouseDown returns true if mouse button `button` is down.
func IsMouseDown(button imgui.ImGuiMouseButton) bool {
	return imgui.IsMouseDown(button)
}

// IsMouseClicked returns true if mouse button `button` is clicked
// NOTE: if you're looking for clicking detection, see EventHandler.go.
func IsMouseClicked(button imgui.ImGuiMouseButton) bool {
	return imgui.IsMouseClicked(button, false)
}

// IsMouseReleased returns true if mouse button `button` is released.
func IsMouseReleased(button imgui.ImGuiMouseButton) bool {
	return imgui.IsMouseReleased(button)
}

// IsMouseDoubleClicked returns true if mouse button `button` is double clicked.
func IsMouseDoubleClicked(button imgui.ImGuiMouseButton) bool {
	return imgui.IsMouseDoubleClicked(button)
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
func IsWindowFocused(flags imgui.ImGuiFocusedFlags) bool {
	return imgui.IsWindowFocused(flags)
}

// IsWindowHovered returns true if the window is hovered.
func IsWindowHovered(flags imgui.ImGuiHoveredFlags) bool {
	return imgui.IsWindowHovered(flags)
}
