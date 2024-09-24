package giu

import "github.com/AllenDang/cimgui-go/imgui"

// Direction represents a ArrowButton direction.
type Direction imgui.Dir

// directions.
const (
	DirectionLeft  = Direction(imgui.DirLeft)
	DirectionRight = Direction(imgui.DirRight)
	DirectionUp    = Direction(imgui.DirUp)
	DirectionDown  = Direction(imgui.DirDown)
)
