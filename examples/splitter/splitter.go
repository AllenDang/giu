package main

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

var (
	width  float32 = 200
	height float32 = 400
	deltaX float32
	deltaY float32
)

func loop(w *g.MasterWindow) {
	g.SingleWindow(w, "splitter", g.Layout{
		g.Custom(func() {
			width += deltaX
			height += deltaY
			imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0, Y: 0})
		}),
		g.SameLine(
			g.Child("left", width, height, true, 0, g.Layout{}),
			g.VSplitter("vsplitter", 8, height, &deltaX),
			g.Child("rightUp", 0, height, true, 0, g.Layout{}),
		),
		g.HSplitter("hsplitter", -1, 8, &deltaY),
		g.Child("rightDown", 0, 0, true, 0, g.Layout{}),
		g.Custom(func() {
			imgui.PopStyleVar()
		}),
	})
}

func main() {
	wnd := g.NewMasterWindow("Splitter", 800, 600, true, nil)
	wnd.Main(loop)
}
