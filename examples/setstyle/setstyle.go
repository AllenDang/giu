package main

import (
	"image/color"

	"github.com/AllenDang/giu"
)

func loop(w *giu.MasterWindow) {
	giu.SingleWindow(w, "set style", func() {
		// #3674D5
		giu.PushColorText(color.RGBA{0x36, 0x74, 0xD5, 255})
		giu.Label("I'm a styled label setting by imgui.PushStyleColor")
		giu.PopStyleColor()
		giu.Label("I'm a normal label")
	})
}

func main() {
	wnd := giu.NewMasterWindow("Set Style", 400, 200, false, nil)
	wnd.Main(loop)
}
