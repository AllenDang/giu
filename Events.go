package giu

import "github.com/AllenDang/giu/imgui"

func IsItemHovered() bool {
	return imgui.IsItemHovered()
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

type MouseButton int

const (
	MouseButtonLeft   MouseButton = 0
	MouseButtonRight  MouseButton = 1
	MouseButtonMiddle MouseButton = 2
)

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

func IsWindowFocused(flags FocusedFlags) bool {
	return imgui.IsWindowFocused(int(flags))
}

func IsWindowHovered(flags HoveredFlags) bool {
	return imgui.IsWindowHovered(int(flags))
}
