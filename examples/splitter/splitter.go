package main

import (
	g "github.com/AllenDang/giu"
)

var (
	sashPos1 float32 = 200
	sashPos2 float32 = 200
	sashPos3 float32 = 200
	sashPos4 float32 = 100
)

func loop() {
	g.SingleWindow().Layout(
		g.SplitLayout(g.DirectionHorizontal, &sashPos1,
			g.Layout{
				g.Label("Left panel"),
				g.Row(g.Button("Button1"), g.Button("Button2")),
			},
			g.SplitLayout(g.DirectionVertical, &sashPos2,
				g.Layout{},
				g.SplitLayout(g.DirectionHorizontal, &sashPos3,
					g.Layout{},
					g.SplitLayout(g.DirectionVertical, &sashPos4,
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
