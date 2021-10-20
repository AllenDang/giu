package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	buf       []uint8
	memEditor imgui.MemoryEditor
	content   string
)

func loop() {
	g.SingleWindow().Layout(
		g.Button("Print data value").OnClick(func() {
			fmt.Println(buf)
		}),
		g.Custom(func() {
			memEditor.DrawContents(buf)
		}),
	)
}

func main() {
	buf = []uint8{1, 2, 3, 4, 5, 6, 7}
	memEditor = imgui.NewMemoryEditor()

	wnd := g.NewMasterWindow("Memory Editor", 800, 600, 0)
	wnd.Run(loop)
}
