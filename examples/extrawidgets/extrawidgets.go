package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

var (
	showPD bool    = true
	radius float32 = 20
	stack  int32
)

func loop() {
	g.SingleWindow().Layout(
		g.Checkbox("Show ProgressIndicator", &showPD),
		g.Condition(showPD, g.Layout{
			g.SliderFloat(&radius, 10, 100).Label("Radius"),
			g.Row(
				g.ProgressIndicator("", 2*radius+20, 2*radius+20, radius),
				g.ProgressIndicator("", 2*radius+20, 2*radius+20, radius),
			),
			g.ProgressIndicator("Loading...", 0, 0, radius),
		}, nil),
		g.Separator(),
		g.Label("stack widget"),
		g.SliderInt(&stack, 0, 2),
		g.Stack(stack,
			g.Layout{
				g.Label("I'm label 1"),
				g.Button("I'm a button").OnClick(func() { fmt.Println("button 1") }),
			},
			g.Layout{
				g.Label("I'm a label 2"),
				g.Button("I'm a button").OnClick(func() { fmt.Println("button 2") }),
			},
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("Extra Widgets", 800, 600, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
