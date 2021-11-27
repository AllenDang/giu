package main

import (
	"fmt"

	"github.com/AllenDang/giu"
	"github.com/faiface/mainthread"
)

func loop() {
	go func() {
		giu.Update()
		mainthread.Call(func() {
			w, h := giu.GetAvailableRegion()
			fmt.Println(w, h)
		})
	}()

	giu.SingleWindow().Layout(
		giu.Label("Hello World"),
	)
}

func main() {
	wnd := giu.NewMasterWindow("issue 409 [bug]", 640, 480, 0)
	wnd.Run(loop)
}
