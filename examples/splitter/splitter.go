package main

import (
	g "github.com/ianling/giu"
)

func loop() {
	g.SingleWindow("splitter").Layout(
		g.SplitLayout("Split", g.DirectionHorizontal, true, 200,
			g.Layout{
				g.Label("Left panel"),
				g.Line(g.Button("Button1"), g.Button("Button2")),
			},
			g.SplitLayout("Right panel", g.DirectionVertical, true, 200,
				g.Layout{},
				g.SplitLayout("HSplit", g.DirectionHorizontal, true, 200,
					g.Layout{},
					g.SplitLayout("VSplit", g.DirectionVertical, true, 100,
						g.Layout{},
						g.Layout{},
					),
				),
			),
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("Splitter", 800, 600, 0, nil)
	wnd.Run(loop)
}
