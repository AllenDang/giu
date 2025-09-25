// Package main shows how to use multiple fonts in giu.
package main

import (
	g "github.com/AllenDang/giu"
)

var (
	anotherFont *g.FontInfo

	content = "Hello world from giu!\n你好啊世界！"
)

func loop() {
	g.SingleWindow().Layout(
		g.Label("Title line").Font(anotherFont),
		g.Label("Content line"),

		g.Label("Change font for other widgets"),

		g.Style().SetFont(anotherFont).SetFontSize(28).To(
			g.Button("Button with big font"),
		),

		g.Button("Normal button"),

		// Show chinese characters
		g.Label("你好啊！世界"),

		// Change font for input area
		g.InputTextMultiline(&content).Size(g.Auto, g.Auto),
	)
}

func main() {
	wnd := g.NewMasterWindow("Multiple fonts", 600, 400, g.MasterWindowFlagsNotResizable)

	// Change the default font
	g.Context.FontAtlas.SetDefaultFont("Arial.ttf")

	// Add a new font and manually set it when needed
	anotherFont = g.Context.FontAtlas.AddFont("Menlo.ttc")

	wnd.Run(loop)
}
