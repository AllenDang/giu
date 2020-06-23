package main

import (
	g "github.com/AllenDang/giu"
)

var (
	showWindow2 bool
)

func onShowWindow2() {
	showWindow2 = true
}

func onHideWindow2() {
	showWindow2 = false
}

func loop() {
	g.Window("Window 1", 10, 10, 200, 100, g.Layout{
		g.Label("I'm a label in window 1"),
		g.Button("Show Window 2", onShowWindow2),
	})

	if showWindow2 {
		g.Window("Window 2", 250, 10, 200, 100, g.Layout{
			g.Label("I'm a label in window 2"),
			g.Button("Hide me", onHideWindow2),
		})
	}
}

func main() {
	wnd := g.NewMasterWindow("Multi sub window demo", 600, 400, 0, nil)
	wnd.Main(loop)
}
