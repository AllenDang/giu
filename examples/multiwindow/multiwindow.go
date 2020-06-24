package main

import (
	g "github.com/AllenDang/giu"
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
	g.MainMenuBar(g.Layout{
		g.Menu("File", g.Layout{
			g.MenuItem("Open", nil),
			g.Separator(),
			g.MenuItem("Exit", nil),
		}),
		g.Menu("Misc", g.Layout{
			g.Checkbox("Enable Me", &checked, nil),
			g.Button("Button", nil),
		}),
	}).Build()

	g.Window("Window 1", 10, 30, 200, 100, g.Layout{
		g.Label("I'm a label in window 1"),
		g.Button("Show Window 2", onShowWindow2),
	})

	if showWindow2 {
		g.WindowV("Window 2", &showWindow2, g.WindowFlagsNone, 250, 30, 200, 100, g.Layout{
			g.Label("I'm a label in window 2"),
			g.Button("Hide me", onHideWindow2),
		})
	}
}

func main() {
	wnd := g.NewMasterWindow("Multi sub window demo", 600, 400, 0, nil)
	wnd.Main(loop)
}
