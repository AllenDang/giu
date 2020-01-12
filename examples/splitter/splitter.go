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
	g.SingleWindow(w, "splitter",
		func() {
			width += deltaX
			height += deltaY
			imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0, Y: 0})
		},
		g.ChildV("left", true, width, height, 0),
		g.SameLine(),
		g.VSplitter("vsplitter", 8, height, &deltaX),
		g.SameLine(),
		g.ChildV("rightUp", true, 0, height, 0,
			g.Label("Drag splitter between panels to resize them"),
		),
		g.HSplitter("hsplitter", -1, 8, &deltaY),
		g.ChildV("rightDown", true, 0, 0, 0),
		func() {
			imgui.PopStyleVar()
		},
	)
}

func main() {
	wnd := g.NewMasterWindow("Splitter", 800, 600, true, nil)
	wnd.Main(loop)
}
