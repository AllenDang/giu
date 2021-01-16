package main

import (
	"math"

	g "github.com/AllenDang/giu"
)

func loop() {
	var plotdata []float32
	const delta = 0.01
	for x := 0.0; len(plotdata) < 1000; x += delta {
		plotdata = append(plotdata, float32(math.Sin(x)))
	}
	g.SingleWindow("hello world").Layout(g.Layout{
		g.Label("Hello world from giu"),
		g.Label("Simple sin(x) plot:"),
		g.PlotLines("testplot", plotdata),
		g.Label("sin(x) plot with overlay text, and size:"),
		g.PlotLines("plot label", plotdata).OverlayText("overlay text").Size(500, 200),
	})
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 600, 400, g.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}
