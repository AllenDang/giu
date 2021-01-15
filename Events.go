package giu

import "github.com/inkyblackness/imgui-go/v3"

type MouseButton int

const (
	MouseButtonLeft   MouseButton = 0
	MouseButtonRight  MouseButton = 1
	MouseButtonMiddle MouseButton = 2
)

func IsItemHovered() bool {
	return imgui.IsItemHovered()
}

func IsItemClicked() bool {
	return imgui.IsItemClicked()
}

func IsItemActive() bool {
	return imgui.IsItemActive()
}

func IsKeyDown(key Key) bool {
	return imgui.IsKeyDown(int(key))
}

func IsKeyPressed(key Key) bool {
	return imgui.IsKeyPressed(int(key))
}

func IsKeyReleased(key Key) bool {
	return imgui.IsKeyReleased(int(key))
}

func IsMouseDown(button MouseButton) bool {
	return imgui.IsMouseDown(int(button))
}

func IsMouseClicked(button MouseButton) bool {
	return imgui.IsMouseClicked(int(button))
}

func IsMouseReleased(button MouseButton) bool {
	return imgui.IsMouseReleased(int(button))
}

func IsMouseDoubleClicked(button MouseButton) bool {
	return imgui.IsMouseDoubleClicked(int(button))
}

func IsWindowAppearing() bool {
	return imgui.IsWindowAppearing()
}

func IsWindowCollapsed() bool {
	return imgui.IsWindowCollapsed()
}

func IsWindowFocusedV(flags FocusedFlags) bool {
	return imgui.IsWindowFocusedV(int(flags))
}

func IsWindowFocused() bool {
	return imgui.IsWindowFocused()
}

func IsWindowHoveredV(flags HoveredFlags) bool {
	return imgui.IsWindowHoveredV(int(flags))
}

func IsWindowHovered() bool {
	return imgui.IsWindowHovered()
}
