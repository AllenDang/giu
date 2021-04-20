package main

import (
	"fmt"
	"image"
	"image/color"

	g "github.com/ianling/giu"
)

type CircleButtonWidget struct {
	id      string
	clicked func()
}

func CircleButton(id string, clicked func()) *CircleButtonWidget {
	return &CircleButtonWidget{
		id:      id,
		clicked: clicked,
	}
}

func (c *CircleButtonWidget) Build() {
	width, height := g.CalcTextSize(c.id)
	var padding float32 = 8.0

	pos := g.GetCursorPos()

	// Calcuate the center point
	radius := int(width/2 + padding)

	// Place a invisible button to be a placeholder for events
	buttonWidth := float32(radius) * 2
	g.InvisibleButton(c.id).Size(buttonWidth, buttonWidth).OnClick(c.clicked).Build()

	// If button is hovered
	drawActive := g.IsItemHovered()

	// Draw circle
	center := pos.Add(image.Pt(radius, radius))

	canvas := g.GetCanvas()
	if drawActive {
		canvas.AddCircleFilled(center, float32(radius), color.RGBA{R: 12, G: 12, B: 200, A: 255})
	}
	canvas.AddCircle(center, float32(radius), color.RGBA{R: 200, G: 12, B: 12, A: 255}, 2)

	// Draw text
	canvas.AddText(center.Sub(image.Pt(int((width-padding)/2), int(height/2))),
		color.RGBA{R: 255, G: 255, B: 255, A: 255}, c.id)
}

func onHello() {
	fmt.Println("Hello")
}

func onWorld() {
	fmt.Println("World")
}

func onCircleButton() {
	fmt.Println("Circle Button")
}

func loop() {
	g.SingleWindow("custom widget").Layout(
		g.Line(CircleButton("Hello", onHello), CircleButton("World", onWorld)),
		CircleButton("Circle Button", onCircleButton),
	)
}

func main() {
	wnd := g.NewMasterWindow("Custom Widget", 400, 300, g.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}
