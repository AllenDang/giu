package main

import "github.com/AllenDang/giu"

var text string

func loop() {
	giu.SingleWindow().Layout(
		giu.Label("Here are widgets with the same IDs in the code, but they IDs are diffrent for ImGui"),
		giu.Button("Button##1"),
		giu.Button("Button##1"),
		giu.InputText(&text).Label("input"),
		giu.InputText(&text).Label("input"),
		giu.InputTextMultiline(&text),
		giu.InputTextMultiline(&text),
	)
}

func main() {
	wnd := giu.NewMasterWindow("AutoID test", 640, 480, 0)
	wnd.Run(loop)
}
