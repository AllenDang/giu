package main

import (
	g "github.com/ianling/giu"
)

var (
	showWindow2 bool
	checked     bool
)

func onShowWindow2() {
	showWindow2 = true
}

func onHideWindow2() {
	showWindow2 = false
}

func loop() {
	g.MainMenuBar().Layout(
		g.Menu("File").Layout(
			g.MenuItem("Open"),
			g.Separator(),
			g.MenuItem("Exit"),
		),
		g.Menu("Misc").Layout(
			g.Checkbox("Enable Me", &checked),
			g.Button("Button"),
		),
	).Build()

	g.Window("Window 1").Pos(10, 30).Size(200, 100).Layout(
		g.Label("I'm a label in window 1"),
		g.Button("Show Window 2").OnClick(onShowWindow2),
	).Build()

	if showWindow2 {
		g.Window("Window 2").IsOpen(&showWindow2).Flags(g.WindowFlagsNone).Pos(250, 30).Size(200, 100).Layout(
			g.Label("I'm a label in window 2"),
			g.Button("Hide me").OnClick(onHideWindow2),
		).Build()
	}
}

func main() {
	wnd := g.NewMasterWindow("Multi sub window demo", 600, 400, 0, nil)
	wnd.Run(loop)
}
