package main

import "github.com/AllenDang/giu"

func loop() {
	giu.SingleWindow().Layout(
		giu.Table().Columns(
			giu.TableColumn("State").InnerWidthOrWeight(20).Flags(giu.TableColumnFlagsWidthFixed),
		).Rows(
			giu.TableRow(
				giu.Label("1"),
			),
			giu.TableRow(
				giu.Label("2"),
			),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Width of table column [discussion 652]", 640, 480, 0)
	wnd.Run(loop)
}
