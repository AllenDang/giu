package main

import (
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/AllenDang/giu"
)

func loop() {
	giu.SingleWindow().Layout(
		giu.Custom(func() {
			imgui.Text("hi")
		}),
	)
}

func main() {
	wnd := giu.NewMasterWindow("test", 640, 480, 0)
	wnd.Run(loop)
}
