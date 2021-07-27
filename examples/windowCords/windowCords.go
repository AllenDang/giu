package main

import (
	"fmt"

	"github.com/AllenDang/giu"
)

func loop() {
	w := giu.Window("window 1")

	posX, posY := w.CurrentPosition()
	width, height := w.CurrentSize()
	w.Layout(
		giu.Label(fmt.Sprintf("Position: %v, %v", posX, posY)),
		giu.Label(fmt.Sprintf("Size: %v, %v", width, height)),
	)
}

func main() {
	wnd := giu.NewMasterWindow("window size/position", 640, 480, 0)

	wnd.Run(loop)
}
