package main

import (
	"image"
	"image/color"
)

// DrawCommand represents a generic drawing command with attributes for tool type, color, brush size, and points.
type DrawCommand struct {
	Tool      int         // Tool indicates the type of drawing tool used (e.g., line, fill).
	Color     color.Color // Color specifies the color used for the drawing command.
	BrushSize float32     // BrushSize defines the size of the brush for the drawing command.
	From      image.Point // From is the starting point of the drawing command.
	To        image.Point // To is the ending point of the drawing command.
}

// LineCommand represents a command to draw a line with specific attributes.
type LineCommand struct {
	P1        image.Point // P1 is the starting point of the line.
	P2        image.Point // P2 is the ending point of the line.
	C         color.Color // C specifies the color of the line.
	Thickness float32     // Thickness defines the thickness of the line.
}

// FillCommand represents a command to fill an area with a specific color.
type FillCommand struct {
	P1 image.Point // P1 is the starting point for the fill operation.
	C  color.Color // C specifies the color used for the fill.
}

// ToLine converts a DrawCommand into a LineCommand.
func (d *DrawCommand) ToLine() LineCommand {
	return LineCommand{P1: d.From, P2: d.To, C: d.Color, Thickness: d.BrushSize}
}

// ToFill converts a DrawCommand into a FillCommand.
func (d *DrawCommand) ToFill() FillCommand {
	return FillCommand{P1: d.From, C: d.Color}
}
