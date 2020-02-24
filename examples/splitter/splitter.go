package main

import (
	g "github.com/AllenDang/giu"
)

func loop() {
	g.SingleWindow("splitter", g.Layout{
		g.SplitLayout("Split", g.DirectionHorizontal, 200,
			g.Layout{
				g.Label("Left panel"),
				g.Line(g.Button("Button1", nil), g.Button("Button2", nil)),
			},
			g.SplitLayout("Right panel", g.DirectionVertical, 200,
				g.Layout{},
				g.SplitLayout("HSplit", g.DirectionHorizontal, 200,
					g.Layout{},
					g.SplitLayout("VSplit", g.DirectionVertical, 100,
						g.Layout{},
						g.Layout{},
					),
				),
			),
		),
	})
}

func main() {
	wnd := g.NewMasterWindow("Splitter", 800, 600, 0, nil)
	wnd.Main(loop)
}
