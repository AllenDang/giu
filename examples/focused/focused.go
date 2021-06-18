package main

import (
	"fmt"

	"github.com/AllenDang/giu"
)

func loop() {
	giu.Window("window 1").Layout(
		giu.Label("I'm a label"),
		giu.Custom(func() {
			fmt.Println(giu.IsWindowFocused(2))
		}),
	)
	giu.Window("window 2").Layout(giu.Label("I'm a label in window 2"))
}

func main() {
	wnd := giu.NewMasterWindow("focused test", 640, 480, 0)
	wnd.Run(loop)
}
