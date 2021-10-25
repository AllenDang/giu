package main

import (
	"fmt"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var markdown string = fmt.Sprint(
	"  * list\n",
	"[link](https://github.com)\n",
)

func loop() {
	giu.SingleWindow().Layout(
		giu.InputTextMultiline(&markdown),
		giu.Custom(func() {
			imgui.Markdown(&markdown)
		}),
	)
}

func main() {
	wnd := giu.NewMasterWindow("ImGui Markdown [Demo]", 640, 480, 0)
	wnd.Run(loop)
}
