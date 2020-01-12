package main

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

func loop(w *g.MasterWindow) {
	g.SingleWindow(w, "canvas",
		g.Label("Canvas demo"),
		func() {
			drawlist := imgui.GetWindowDrawList()

			pos := imgui.CursorScreenPos()

			drawlist.AddLine(
				imgui.Vec2{X: pos.X, Y: pos.Y},
				imgui.Vec2{X: pos.X + 100, Y: pos.Y + 100},
				imgui.Vec4{X: 1, Y: 1, Z: 0.4, W: 1}, 1.0)
		},
	)
}

func main() {
	wnd := g.NewMasterWindow("Canvas", 400, 300, false, nil)
	wnd.Main(loop)
}
