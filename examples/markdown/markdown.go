package main

import (
	"fmt"

	"github.com/pkg/browser"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var markdown string = fmt.Sprint(
	"  * list\n",
	"Here is [a link to some cool website!](https://github.com/AllenDang/giu) you must click it!\n",
)

func loop() {
	giu.SingleWindow().Layout(
		giu.InputTextMultiline(&markdown),
		giu.Custom(func() {
			imgui.Markdown(&markdown, func(s string) {
				browser.OpenURL(s)
			})
		}),
	)
}

func main() {
	wnd := giu.NewMasterWindow("ImGui Markdown [Demo]", 640, 480, 0)
	wnd.Run(loop)
}
