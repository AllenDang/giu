package main

import (
	g "github.com/ianling/giu"
)

var (
	showPD bool    = true
	radius float32 = 20
)

func loop() {
	g.SingleWindow("Extra Widgets").Layout(
		g.Checkbox("Show ProgressIndicator", &showPD),
		g.Condition(showPD, g.Layout{
			g.SliderFloat("Radius", &radius, 10, 100),
			g.Line(
				g.ProgressIndicator("pd1", "", 20+radius, 20+radius, radius),
				g.ProgressIndicator("pd2", "", 20+radius, 20+radius, radius),
			),
			g.ProgressIndicator("pd3", "Loading...", 0, 0, radius),
		}, nil),
	).Build()
}

func main() {
	wnd := g.NewMasterWindow("Extra Widgets", 800, 600, g.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}
