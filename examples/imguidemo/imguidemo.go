package main

import (
	imgui "github.com/AllenDang/cimgui-go"
	g "github.com/AllenDang/giu"
)

func loop() {
	imgui.ShowDemoWindow(nil)
}

func main() {
	wnd := g.NewMasterWindow("Widgets", 1024, 768, 0)
	wnd.Run(loop)
}
