package giu

import "github.com/AllenDang/giu/imgui"

func IsItemHovered() bool {
	return imgui.IsItemHovered()
}

func IsItemActive() bool {
	return imgui.IsItemActive()
}

func IsKeyDown(key int) bool {
	return imgui.IsKeyDown(key)
}

func IsKeyPressed(key int) bool {
	return imgui.IsKeyPressed(key)
}

func IsKeyReleased(key int) bool {
	return imgui.IsKeyReleased(key)
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
