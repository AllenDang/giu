// Package main demonstrates use of TableFlagsSortable and Sort function.
package main

import (
	"fmt"
	"sort"

	"github.com/AllenDang/giu"
)

var (
	data = []string{"A", "AA", "ABC", "CBA", "BBB"}
	cols = []*giu.TableRowWidget{}
)

func rebuildColumns() {
	cols = make([]*giu.TableRowWidget, 0)
	for _, d := range data {
		cols = append(cols, giu.TableRow(giu.Label(d)))
	}
}

func loop() {
	giu.SingleWindow().Layout(
		giu.Table().Flags(giu.TableFlagsSortable|giu.TableFlagsResizable).Columns(
			giu.TableColumn("Col 1").Sort(func(s giu.SortDirection) {
				fmt.Println("sorting col 1", s)
				switch s {
				case giu.SortAscending:
					sort.Strings(data)
				case giu.SortDescending:
					sort.Sort(sort.Reverse(sort.StringSlice(data)))
				}

				rebuildColumns()
			}),
			giu.TableColumn("Col 2"),
		).Rows(cols...),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Table sorting", 640, 480, 0)
	wnd.SetUserFile("giu.ini")

	rebuildColumns()

	wnd.Run(loop)
}
