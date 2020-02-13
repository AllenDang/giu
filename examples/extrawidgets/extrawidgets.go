package main

import (
	g "github.com/AllenDang/giu"
)

var (
	showPD bool    = true
	radius float32 = 20
)

func loop() {
	g.SingleWindow("Extra Widgets", g.Layout{
		g.Checkbox("Show ProgressIndicator", &showPD, nil),
		g.Condition(showPD, g.Layout{
			g.SliderFloat("Radius", &radius, 10, 100, ""),
			g.ProgressIndicator("PD", radius),
			g.Line(
				g.ProgressIndicator("PD1", radius),
				g.ProgressIndicator("PD2", 20),
			),
		}),
	})
}

func main() {
	wnd := g.NewMasterWindow("Extra Widgets", 800, 600, false, nil)
	wnd.Main(loop)
}
