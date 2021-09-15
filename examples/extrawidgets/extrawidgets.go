package main

import (
	g "github.com/AllenDang/giu"
)

var (
	showPD bool    = true
	radius float32 = 20
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
	)
}

func main() {
	wnd := g.NewMasterWindow("Extra Widgets", 800, 600, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
