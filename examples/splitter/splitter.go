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

func loop() {
	g.SingleWindow("splitter", g.Layout{
		g.Custom(func() {
			width += deltaX
			height += deltaY
			imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0, Y: 0})
		}),
		g.Line(
			g.Child("left", true, width, height, 0, nil),
			g.VSplitter("vsplitter", 8, height, &deltaX),
			g.Child("rightUp", true, 0, height, 0, g.Layout{
				g.Label("Drag splitter between panels to resize them"),
			},
			),
		),
		g.HSplitter("hsplitter", -1, 8, &deltaY),
		g.Child("rightDown", true, 0, 0, 0, nil),
		g.Custom(func() {
			imgui.PopStyleVar()
		}),
	})
}

func main() {
	wnd := g.NewMasterWindow("Splitter", 800, 600, 0, nil)
	wnd.Main(loop)
}
