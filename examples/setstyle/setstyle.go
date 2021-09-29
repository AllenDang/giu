package main

import (
	"image/color"

	g "github.com/AllenDang/giu"
)

func loop() {
	g.SingleWindow().Layout(
		g.Style().
			SetColor(g.StyleColorText, color.RGBA{0x36, 0x74, 0xD5, 255}).
			To(
				g.Label("I'm a styled label"),
			),
		g.Style().
			SetColor(g.StyleColorBorder, color.RGBA{0x36, 0x74, 0xD5, 255}).
			SetStyle(g.StyleVarFramePadding, 10, 10).
			To(
				g.Button("I'm a styled button"),
			),
		g.Button("I'm a normal button"),
		g.Style().
			SetFontSize(60).To(
			g.Label("large label"),
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("Set Style", 400, 200, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
