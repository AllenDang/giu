package main

import (
	g "github.com/AllenDang/giu"
)

var (
	bigFont *g.FontInfo

	content = "Hello world from giu!\n你好啊世界！"
)

func loop() {
	g.SingleWindow().Layout(
		g.Label("Title line").Font(bigFont),
		g.Label("Content line"),

		g.Label("Change font for other widgets"),

		g.Style().SetFont(bigFont).To(
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
	g.Context.FontAtlas.SetDefaultFont("Arial.ttf", 12)

	// Add a new font and manually set it when needed
	bigFont = g.Context.FontAtlas.AddFont("Menlo.ttc", 24)

	wnd.Run(loop)
}
