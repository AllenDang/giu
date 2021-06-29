package main

import (
	g "github.com/AllenDang/giu"
)

var (
	content string
)

func loop() {
	g.SingleWindow().Layout(
		g.Table().Columns(
			g.TableColumn("Col"),
			g.TableColumn("Col"),
		).Rows(
			g.TableRow(g.Label("c1"), g.Label("c2")),
			g.TableRow(g.Label("c1"), g.Label("c2")),
			g.TableRow(g.Label("c1"), g.Label("c2")),
		),
		g.Label("Hello world from giu"),
		g.InputTextMultiline("##content", &content).Size(-1, -1),
	)
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, 0)
	wnd.Run(loop)
}
