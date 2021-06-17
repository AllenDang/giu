package main

import "github.com/AllenDang/giu"

func loop() {
	w1 := giu.Window("window 1")

	giu.Window("window 2").Layout(
		giu.Label("I'm a window"),
		giu.Button("Click me to focus the other window").OnClick(func() {
			w1.BringToFront()
		}),
	)

	w1.Layout(
		giu.Label("I'm window 1"),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Bring to front", 320, 240, 0)
	wnd.Run(loop)
}
