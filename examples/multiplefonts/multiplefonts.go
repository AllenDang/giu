package main

import (
	g "github.com/AllenDang/giu"
)

var (
	bigFont *g.FontInfo

	content = "Hello world from giu!\n你好啊世界！"
)

func loop() {
	g.SingleWindow("Multiple fonts").Layout(
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
		g.InputTextMultiline("##content", &content).Size(-1, -1),
	)
}

func main() {
	bigFont = g.AddFont("Menlo.ttc", 24)

	wnd := g.NewMasterWindow("Multiple fonts", 600, 400, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
