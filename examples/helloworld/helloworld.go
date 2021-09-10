package main

import (
	g "github.com/AllenDang/giu"
)

var content string

func loop() {
	g.SingleWindow().Layout(
		g.Label("Hello world from giu"),
		g.InputTextMultiline(&content).Size(g.Auto, g.Auto),
	)
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, 0)
	wnd.Run(loop)
}
