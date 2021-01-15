package main

import (
	g "github.com/AllenDang/giu"
	"github.com/inkyblackness/imgui-go/v3"
)

func loop() {
	imgui.ShowDemoWindow(nil)
}

func main() {
	wnd := g.NewMasterWindow("Widgets", 1024, 768, 0, nil)
	wnd.Run(loop)
}
