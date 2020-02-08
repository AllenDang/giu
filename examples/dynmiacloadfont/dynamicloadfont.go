package main

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

func loadFont() {
	fonts := g.Context.IO().Fonts()

	ranges := imgui.NewGlyphRanges()

	builder := imgui.NewFontGlyphRangesBuilder()
	builder.AddText("铁憨憨你好！")
	// builder.AddRanges(fonts.GlyphRangesChineseFull())
	builder.BuildRanges(ranges)

	fontPath := "c:/Windows/Fonts/MSYHL.TTC"
	fonts.AddFontFromFileTTFV(fontPath, 12, imgui.DefaultFontConfig, ranges.Data())
}

func loop() {
	g.SingleWindow("dynamic load font", g.Layout{
		g.Label("你好啊世界！铁憨憨"),
	})
}

func main() {
	wnd := g.NewMasterWindow("Dynamic load font", 400, 200, false, loadFont)
	wnd.Main(loop)
}
