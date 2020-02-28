package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

var (
	names []string
)

func buildRows() []*g.RowWidget {
	rows := make([]*g.RowWidget, len(names))

	for i := range rows {
		rows[i] = g.Row(
			g.Label(fmt.Sprintf("%d", i)),
			g.Label(names[i]),
		)
	}

	return rows
}

func loop() {
	g.SingleWindow("Huge list demo", g.Layout{
		g.Label("Use FastTable to display huge amount of rows"),
		g.Label("Note: FastTable only works if all rows have same height"),
		g.Child("Container", true, 0, 0, 0, g.Layout{
			g.FastTable("Fast table", true, buildRows()),
		}),
	})
}

func main() {
	names = make([]string, 10000)
	for i := range names {
		names[i] = fmt.Sprintf("Huge list name demo %d", i)
	}

	wnd := g.NewMasterWindow("Huge list demo", 800, 600, 0, nil)
	wnd.Main(loop)
}
