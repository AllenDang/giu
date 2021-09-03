package main

import "github.com/AllenDang/giu"

var text string

func loop() {
	giu.SingleWindow().Layout(
		giu.Align(giu.AlignCenter).To(
			giu.Label("I'm a centered label"),
			giu.Button("I'm a centered button"),
		),

		giu.Align(giu.AlignRight).To(
			giu.Label("I'm a alined to right label"),
			giu.InputText(&text),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Alignment demo", 640, 480, 0)
	wnd.Run(loop)
}
