package main

import (
	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

func loop() {
	giu.SingleWindow().Layout(
		giu.Custom(func() {
			imgui.PushStyleColor(imgui.StyleColorChildBg, imgui.CurrentStyle().GetColor(imgui.StyleColorWindowBg))
			if imgui.BeginChild("child") {
				giu.Button("button").Build()
			}
			imgui.EndChild()
			imgui.PopStyleColor()
		}),
		giu.Label("label"),
	)
}

func main() {
	wnd := giu.NewMasterWindow("native child layout [demo]", 640, 480, 0)
	wnd.Run(loop)
}
