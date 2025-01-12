// Package main shows how to create a transparent window with custom drawing.
package main

import (
	"image"
	"image/color"

	"github.com/AllenDang/cimgui-go/imgui"

	g "github.com/AllenDang/giu"
)

var (
	wnd           *g.MasterWindow
	isMovingFrame = false
)

func FramelessMovableWidget(widget g.Widget) *g.CustomWidget {
	return g.Custom(func() {
		if isMovingFrame && !g.IsMouseDown(g.MouseButtonLeft) {
			isMovingFrame = false
			return
		}
		widget.Build()
		if g.IsItemHovered() {
			if g.IsMouseDown(g.MouseButtonLeft) {
				isMovingFrame = true
			}
		}
		if isMovingFrame {
			delta := imgui.CurrentIO().MouseDelta()
			dx := int(delta.X)
			dy := int(delta.Y)
			if dx != 0 || dy != 0 {
				ox, oy := wnd.GetPos()
				wnd.SetPos(ox+dx, oy+dy)
			}
		}
	})
}

func loop() {
	imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0)
	g.PushColorWindowBg(color.RGBA{50, 50, 70, 130})
	g.PushColorFrameBg(color.RGBA{30, 30, 60, 110})
	g.SingleWindow().Layout(
		FramelessMovableWidget(
			g.Label("Maintain Left-click on me to move the frameless window !"),
		),
		g.Custom(func() {
			canvas := g.GetCanvas()
			pos := g.GetCursorScreenPos()
			col := color.RGBA{200, 75, 75, 255}

			canvas.AddLine(pos, pos.Add(image.Pt(100, 100)), col, 1)
			canvas.AddRect(pos.Add(image.Pt(110, 0)), pos.Add(image.Pt(200, 100)), col, 5, g.DrawFlagsRoundCornersAll, 1)
			canvas.AddRectFilled(pos.Add(image.Pt(220, 0)), pos.Add(image.Pt(320, 100)), col, 0, 0)

			pos0 := pos.Add(image.Pt(0, 110))
			cp0 := pos.Add(image.Pt(80, 110))
			cp1 := pos.Add(image.Pt(50, 210))
			pos1 := pos.Add(image.Pt(120, 210))
			canvas.AddBezierCubic(pos0, cp0, cp1, pos1, col, 1, 0)

			p1 := pos.Add(image.Pt(160, 110))
			p2 := pos.Add(image.Pt(120, 210))
			p3 := pos.Add(image.Pt(210, 210))
			p4 := pos.Add(image.Pt(210, 150))

			canvas.AddTriangle(p1, p2, p3, col, 2)
			canvas.AddQuad(p1, p2, p3, p4, col, 1)

			p1 = p1.Add(image.Pt(120, 60))
			canvas.AddCircleFilled(p1, 50, col)
		}),
	)
	g.PopStyleColor()
	g.PopStyleColor()
	imgui.PopStyleVar()
}

func main() {
	wnd = g.NewMasterWindow("transparent", 300, 200, g.MasterWindowFlagsNotResizable|g.MasterWindowFlagsFloating|g.MasterWindowFlagsFrameless|g.MasterWindowFlagsTransparent)
	wnd.SetBgColor(color.RGBA{0, 0, 0, 0})
	wnd.SetPos(50, 50)
	wnd.Run(loop)
}
