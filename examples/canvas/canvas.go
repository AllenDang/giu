package main

import (
	"image"
	"image/color"
	"log"

	g "github.com/ianling/giu"
)

var (
	texture *g.Texture
)

func loop() {
	g.SingleWindow("canvas").Layout(g.Layout{
		g.Label("Canvas demo"),
		g.Custom(func() {
			canvas := g.GetCanvas()
			pos := g.GetCursorScreenPos()
			color := color.RGBA{200, 75, 75, 255}
			canvas.AddLine(pos, pos.Add(image.Pt(100, 100)), color, 1)
			canvas.AddRect(pos.Add(image.Pt(110, 0)), pos.Add(image.Pt(200, 100)), color, 5, g.CornerFlags_All, 1)
			canvas.AddRectFilled(pos.Add(image.Pt(220, 0)), pos.Add(image.Pt(320, 100)), color, 0, 0)

			pos0 := pos.Add(image.Pt(0, 110))
			cp0 := pos.Add(image.Pt(80, 110))
			cp1 := pos.Add(image.Pt(50, 210))
			pos1 := pos.Add(image.Pt(120, 210))
			canvas.AddBezierCurve(pos0, cp0, cp1, pos1, color, 1, 0)

			p1 := pos.Add(image.Pt(160, 110))
			p2 := pos.Add(image.Pt(120, 210))
			p3 := pos.Add(image.Pt(210, 210))
			p4 := pos.Add(image.Pt(210, 150))
			// canvas.AddTriangle(p1, p2, p3, color, 2)
			canvas.AddQuad(p1, p2, p3, p4, color, 1)

			p1 = p1.Add(image.Pt(120, 60))
			canvas.AddCircleFilled(p1, 50, color)

			p1 = pos.Add(image.Pt(10, 400))
			p2 = pos.Add(image.Pt(50, 440))
			p3 = pos.Add(image.Pt(200, 500))
			canvas.PathLineTo(p1)
			canvas.PathLineTo(p2)
			canvas.PathBezierCurveTo(p2.Add(image.Pt(40, 0)), p3.Add(image.Pt(-50, 0)), p3, 0)
			canvas.PathStroke(color, false, 1)

			if texture != nil {
				canvas.AddImage(texture, image.Pt(350, 25), image.Pt(500, 125))
			}
		}),
	}).Build()
}

func main() {
	wnd := g.NewMasterWindow("Canvas", 600, 600, g.MasterWindowFlagsNotResizable, nil)

	img, err := g.LoadImage("gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		texture, _ = g.NewTextureFromRgba(img)
	}()

	wnd.Run(loop)
}
