package main

import (
	"image/color"

	"github.com/ianling/giu"
)

func loop() {
	giu.SingleWindow("set style").Layout(giu.Layout{
		giu.Label("I'm a styled label").Color(&color.RGBA{0x36, 0x74, 0xD5, 255}),
		giu.Label("I'm a normal label"),
	}).Build()
}

func main() {
	wnd := giu.NewMasterWindow("Set Style", 400, 200, giu.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}
