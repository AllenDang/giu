package main

import (
	"image"
	"image/color"

	g "github.com/ianling/giu"
	"github.com/ianling/imgui-go"
)

func loop() {
	imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0)
	g.PushColorWindowBg(color.RGBA{50, 50, 50, 0})
	g.PushColorFrameBg(color.RGBA{10, 10, 10, 0})
	g.SingleWindow("transparent").Layout(
		g.Custom(func() {
			canvas := g.GetCanvas()
			pos := g.GetCursorScreenPos()
			color := color.RGBA{200, 75, 75, 255}
			canvas.AddLine(pos, pos.Add(image.Pt(100, 100)), color, 1)
			canvas.AddRect(pos.Add(image.Pt(110, 0)), pos.Add(image.Pt(200, 100)), color, 5, g.DrawFlags_RoundCornersAll, 1)
			canvas.AddRectFilled(pos.Add(image.Pt(220, 0)), pos.Add(image.Pt(320, 100)), color, 0, 0)

			pos0 := pos.Add(image.Pt(0, 110))
			cp0 := pos.Add(image.Pt(80, 110))
			cp1 := pos.Add(image.Pt(50, 210))
			pos1 := pos.Add(image.Pt(120, 210))
			canvas.AddBezierCubic(pos0, cp0, cp1, pos1, color, 1, 0)

			p1 := pos.Add(image.Pt(160, 110))
			p2 := pos.Add(image.Pt(120, 210))
			p3 := pos.Add(image.Pt(210, 210))
			p4 := pos.Add(image.Pt(210, 150))
			// canvas.AddTriangle(p1, p2, p3, color, 2)
			canvas.AddQuad(p1, p2, p3, p4, color, 1)

			p1 = p1.Add(image.Pt(120, 60))
			canvas.AddCircleFilled(p1, 50, color)
		}),
	).Build()
	g.PopStyleColor()
	g.PopStyleColor()
	imgui.PopStyleVar()
}

func main() {
	wnd := g.NewMasterWindow("transparent", 300, 200, g.MasterWindowFlagsNotResizable|g.MasterWindowFlagsFloating|g.MasterWindowFlagsFrameless|g.MasterWindowFlagsTransparent, nil)
	wnd.SetBgColor(color.RGBA{0, 0, 0, 0})
	wnd.Run(loop)
}
