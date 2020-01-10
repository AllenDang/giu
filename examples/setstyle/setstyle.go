package main

import (
	"image/color"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

func loop(w *giu.MasterWindow) {
	giu.SingleWindow(w, "set style", giu.Layout{
		giu.Custom(func() {
			// #3674D5
			col := color.RGBA{0x36, 0x74, 0xD5, 255}
			imgui.PushStyleColor(imgui.StyleColorText, giu.ToVec4Color(col))
		}),
		giu.Label("I'm a styled label setting by imgui.PushStyleColor"),
		giu.Custom(func() {
			imgui.PopStyleColor()
		}),
		giu.Label("I'm a normal label"),
	})
}

func main() {
	wnd := giu.NewMasterWindow("Set Style", 400, 200, false, nil)
	wnd.Main(loop)
}
