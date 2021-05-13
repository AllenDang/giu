package main

import (
	g "github.com/AllenDang/giu"
)

var (
	content string
)

func loop() {
	g.SingleWindow("hello world").Layout(
		g.Label("Hello world from giu"),
		g.InputTextMultiline("##content", &content).Size(-1, -1),
	)
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, 0)
	wnd.Run(loop)
}
