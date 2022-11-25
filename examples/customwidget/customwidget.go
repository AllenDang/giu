package main

import (
	"fmt"
	"image"
	"image/color"

	g "github.com/AllenDang/giu"
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

	// Calculate the center point
	radius := int(width/2 + padding*2)

	// Place a invisible button to be a placeholder for events
	buttonWidth := float32(radius) * 2
	g.InvisibleButton().Size(buttonWidth, buttonWidth).OnClick(c.clicked).Build()

	// If button is hovered
	drawActive := g.IsItemHovered()

	// Draw circle
	center := pos.Add(image.Pt(radius, radius))

	canvas := g.GetCanvas()
	if drawActive {
		canvas.AddCircleFilled(center, float32(radius), color.RGBA{12, 12, 200, 255})
	}
	canvas.AddCircle(center, float32(radius), color.RGBA{200, 12, 12, 255}, radius, 2)

	// Draw text
	canvas.AddText(center.Sub(image.Pt(int((width)/2), int(height/2))), color.RGBA{255, 255, 255, 255}, c.id)
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
	g.SingleWindow().Layout(
		g.Row(CircleButton("Hello", onHello), CircleButton("World", onWorld)),
		CircleButton("Circle Button", onCircleButton),
	)
}

func main() {
	wnd := g.NewMasterWindow("Custom Widget", 400, 300, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
