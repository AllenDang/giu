package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
)

var (
	showWindow2 bool
	checked     bool

	x1, y1 = float32(30), float32(30)
	width1, height1 = float32(240), float32(125)
	window1 = g.Window("Window 1").Pos(x1, y1).Size(width1, height1)

	x2, y2 = float32(250), float32(30)
	width2, height2 = float32(200), float32(100)
	window2 = g.Window("Window 2").IsOpen(&showWindow2).Flags(g.WindowFlagsNone).Pos(x2, y2).Size(width2, height2)
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
	)

	// every frame, we simply update the text inside the windows
	var focusText string
	if window1.HasFocus() {
		focusText = "This window has focus"
	} else {
		focusText = "This window does not have focus"
	}
	x1, y1 = window1.CurrentPosition()
	width1, height1 = window1.CurrentSize()
	window1.Layout(
		g.Label("I'm a label in window 1"),
		g.Label(fmt.Sprintf("X: %.0f, Y: %.0f", x1, y1)),
		g.Label(fmt.Sprintf("W: %.0f, H: %.0f", width1, height1)),
		g.Button("Show Window 2").OnClick(onShowWindow2),
		g.Label(focusText),
	)

	if showWindow2 {
		x2, y2 = window2.CurrentPosition()
		width2, height2 = window2.CurrentSize()
		window2.Layout(
				g.Label("I'm a label in window 2"),
				g.Label(fmt.Sprintf("X: %.0f, Y: %.0f", x2, y2)),
				g.Label(fmt.Sprintf("W: %.0f, H: %.0f", width2, height2)),
				g.Button("Hide me").OnClick(onHideWindow2),
		)
	}
}

func main() {
	wnd := g.NewMasterWindow("Multi sub window demo", 600, 400, 0, nil)
	wnd.Run(loop)
}
