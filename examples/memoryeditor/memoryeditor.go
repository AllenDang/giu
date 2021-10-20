package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

var (
	buf     []uint8
	content string
)

func loop() {
	g.SingleWindow().Layout(
		g.Button("Print data value").OnClick(func() {
			fmt.Println(buf)
		}),
		g.MemoryEditor().Contents(buf),
	)
}

func main() {
	buf = []uint8{1, 2, 3, 4, 5, 6, 7}
	wnd := g.NewMasterWindow("Memory Editor", 800, 600, 0)
	wnd.Run(loop)
}
