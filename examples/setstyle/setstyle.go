// Package main shows how to setyle your app by using giu.StyleSetter.
// See also: examples/CSS-styling
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

	// Setting a style for the entire window
	style := g.Style()
	style.SetColor(g.StyleColorWindowBg, color.RGBA{0x55, 0x55, 0x55, 255})
	wnd.SetStyle(style)

	wnd.Run(loop)
}
