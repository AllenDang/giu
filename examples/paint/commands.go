package main

import (
	"image"
	"image/color"
)

type Undo struct {
	RevertIndex int
}

type DrawCommand struct {
	Tool      int
	Color     color.Color
	BrushSize float32
	From      image.Point
	To        image.Point
}

type Line struct {
	P1        image.Point
	P2        image.Point
	C         color.Color
	Thickness float32
}

type Fill struct {
	P1 image.Point
	C  color.Color
}

func (d *DrawCommand) ToLine() Line {
	return Line{P1: d.From, P2: d.To, C: d.Color, Thickness: d.BrushSize}
}

func (d *DrawCommand) ToFill() Fill {
	return Fill{P1: d.From, C: d.Color}
}


