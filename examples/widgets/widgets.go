package main

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

func loop() {
	imgui.ShowDemoWindow(nil)
}

func main() {
	wnd := g.NewMasterWindow("Widgets", 1024, 768, true, nil)
	wnd.Main(loop)
}
