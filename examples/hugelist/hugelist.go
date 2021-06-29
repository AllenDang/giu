package main

import (
	"fmt"
	"image/color"

	g "github.com/AllenDang/giu"
)

var (
	names []string
)

func buildRows() []*g.TableRowWidget {
	rows := make([]*g.TableRowWidget, len(names))

	for i := range rows {
		rows[i] = g.TableRow(
			g.Label(fmt.Sprintf("%d", i)),
			g.Label(names[i]),
		)
	}

	rows[0].BgColor(&(color.RGBA{200, 100, 100, 255}))

	return rows
}

func loop() {
	g.SingleWindow().Layout(
		g.Label("Note: FastTable only works if all rows have same height"),
		g.Table().Freeze(0, 1).FastMode(true).Rows(buildRows()...),
	)
}

func main() {
	names = make([]string, 10000)
	for i := range names {
		names[i] = fmt.Sprintf("Huge list name demo 范例 %d", i)
	}

	wnd := g.NewMasterWindow("Huge list demo", 800, 600, 0)
	wnd.Run(loop)
}
