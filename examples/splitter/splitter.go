package main

import (
	g "github.com/AllenDang/giu"
)

func loop() {
	g.SingleWindow().Layout(
		g.SplitLayout(g.DirectionHorizontal, true, 200,
			g.Layout{
				g.Label("Left panel"),
				g.Row(g.Button("Button1"), g.Button("Button2")),
			},
			g.SplitLayout(g.DirectionVertical, true, 200,
				g.Layout{},
				g.SplitLayout(g.DirectionHorizontal, true, 200,
					g.Layout{},
					g.SplitLayout(g.DirectionVertical, true, 100,
						g.Layout{},
						g.Layout{},
					),
				),
			),
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("Splitter", 800, 600, 0)
	wnd.Run(loop)
}
