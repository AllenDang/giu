package main

import (
	g "github.com/ianling/giu"
	"github.com/ianling/imgui-go"
)

func loop() {
	imgui.ShowDemoWindow(nil)
}

func main() {
	wnd := g.NewMasterWindow("Widgets", 1024, 768, 0, nil)
	wnd.Run(loop)
}
