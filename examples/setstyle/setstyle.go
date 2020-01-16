package main

import (
	"image/color"

	"github.com/AllenDang/giu"
)

func loop(w *giu.MasterWindow) {
	giu.SingleWindow(w, "set style", giu.Layout{
		giu.LabelV("I'm a styled label", &color.RGBA{0x36, 0x74, 0xD5, 255}, nil),
		giu.Label("I'm a normal label"),
	})
}

func main() {
	wnd := giu.NewMasterWindow("Set Style", 400, 200, false, nil)
	wnd.Main(loop)
}
