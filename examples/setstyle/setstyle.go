package main

import (
	"image/color"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

func loop() {
	giu.SingleWindow("set style").Layout(
		giu.Style().
			SetColor(imgui.StyleColorText, color.RGBA{0x36, 0x74, 0xD5, 255}).
			To(
				giu.Label("I'm a styled label"),
			),
		giu.Style().
			SetColor(imgui.StyleColorBorder, color.RGBA{0x36, 0x74, 0xD5, 255}).
			SetStyle(imgui.StyleVarFramePadding, 10, 10).
			To(
				giu.Button("I'm a styled button"),
			),
		giu.Button("I'm a normal button"),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Set Style", 400, 200, giu.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}
