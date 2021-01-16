package main

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

var (
	h2Font      imgui.Font
	chineseFont imgui.Font

	content = "Hello world from giu!\n你好啊世界！"
)

func loadFont() {
	fonts := g.Context.IO().Fonts()

	// Set font size
	fontSize := float32(12)

	// Font ttf/ttc file path
	fontPath := "/System/Library/Fonts/Menlo.ttc"
	chineseFontPath := "/Library/Fonts/Microsoft/SimHei.ttf"

	// Base font
	fonts.AddFontFromFileTTF(fontPath, fontSize)
	// Bold heading font h2
	h2Font = fonts.AddFontFromFileTTF(fontPath, fontSize*2)
	// Chinese font
	chineseFont = fonts.AddFontFromFileTTFV(chineseFontPath, fontSize, imgui.DefaultFontConfig, imgui.GlyphRangesAll())
}

func loop() {
	g.SingleWindow("Multiple fonts").Layout(g.Layout{
		g.Label("Title line").Font(&h2Font),
		g.Label("Content line"),

		g.Label("Change font for other widgets"),

		// Change font for other widgets
		g.Custom(func() { g.PushFont(h2Font) }),

		g.Button("Button with big font"),

		g.Custom(func() { g.PopFont() }),

		g.Button("Normal button"),

		// Show chinese characters
		g.Label("你好啊！世界").Font(&chineseFont),

		// Change font for input area
		g.Custom(func() { g.PushFont(chineseFont) }),
		g.InputTextMultiline("##content", &content).Size(-1, -1),
		g.Custom(func() { g.PopFont() }),
	})
}

func main() {
	wnd := g.NewMasterWindow("Multiple fonts", 600, 400, g.MasterWindowFlagsNotResizable, loadFont)
	wnd.Run(loop)
}
