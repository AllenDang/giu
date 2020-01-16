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
	// builder.AddRanges(fonts.GlyphRangesChineseSimplifiedCommon())
	builder.BuildRanges(ranges)

	fontPath := "/System/Library/Fonts/STHeiti Light.ttc"
	fonts.AddFontFromFileTTFV(fontPath, 12, imgui.DefaultFontConfig, ranges.Data())
}

func loop(w *g.MasterWindow) {
	g.SingleWindow(w, "dynamic load font", g.Layout{
		g.Label("你好啊世界！铁憨憨"),
	})
}

func main() {
	wnd := g.NewMasterWindow("Dynamic load font", 400, 200, false, loadFont)
	wnd.Main(loop)
}
