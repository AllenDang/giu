package main

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

func loadFont() {
	fonts := g.Context.IO().Fonts()

	ranges := imgui.NewGlyphRanges()

	builder := imgui.NewFontGlyphRangesBuilder()
	builder.AddText("铁憨憨你好！") // intentionally not loading all of the characters we use below
	// builder.AddRanges(fonts.GlyphRangesChineseFull())
	builder.BuildRanges(ranges)

	fontPath := "/Users/iling/Downloads/msyhl.ttc"
	fonts.AddFontFromFileTTFV(fontPath, 24, imgui.DefaultFontConfig, ranges.Data())
}

func loop() {
	g.SingleWindow("dynamic load font").Layout(
		g.Label("你好啊世界！铁憨憨"),
	).Build()
}

func main() {
	wnd := g.NewMasterWindow("Dynamic load font", 400, 200, g.MasterWindowFlagsNotResizable, loadFont)
	wnd.Run(loop)
}
