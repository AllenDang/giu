package main

import (
	"github.com/AllenDang/giu"
)

func loop() {
	giu.SingleWindow().Layout(
		giu.Label("hehehe"),
		giu.NodeEditor().Nodes(
			giu.Node().
				Static(
					giu.Label("Main content"),
				).Output(
				giu.Label("Output attribute"),
			).Input(
				giu.Label("Input attribute"),
			),
			giu.Node().Static(
				giu.Label("Main content"),
			).TitleBar(
				giu.Label("This is a title bar"),
			).Input(
				giu.Label("Iput attribute"),
			),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("ImNodes demo", 1280, 720, 0)
	wnd.Run(loop)
}
