package giu

import imgui "github.com/AllenDang/cimgui-go"

// Direction represents a ArrowButton direction.
type Direction imgui.Dir

// directions.
const (
	DirectionLeft  = Direction(imgui.DirLeft)
	DirectionRight = Direction(imgui.DirRight)
	DirectionUp    = Direction(imgui.DirUp)
	DirectionDown  = Direction(imgui.DirDown)
)
