package main

import "github.com/AllenDang/giu"

func loop() {
	giu.SingleWindow().Layout(
		giu.Label("hehehe"),
		giu.NodeEditor(),
	)
}

func main() {
	wnd := giu.NewMasterWindow("ImNodes demo", 1280, 720, 0)
	wnd.Run(loop)
}
