package main

import "github.com/AllenDang/giu"

func main() {
	wnd := giu.NewMasterWindow("test", 640, 480, 0)
	wnd.Run(loop)
}

func loop() {
	giu.SingleWindow().Layout(
		giu.TabItem("tabitem"),
	)
}
