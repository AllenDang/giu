package main

import (
	"image/color"

	"github.com/AllenDang/giu"
)

func loop() {
	giu.SingleWindow("set style", giu.Layout{
		giu.LabelV("I'm a styled label", false, &color.RGBA{0x36, 0x74, 0xD5, 255}, nil),
		giu.Label("I'm a normal label"),
	})
}

func main() {
	wnd := giu.NewMasterWindow("Set Style", 400, 200, giu.MasterWindowFlagsNotResizable, nil)
	wnd.Main(loop)
}
