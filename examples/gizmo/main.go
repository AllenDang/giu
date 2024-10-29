package main

import (
	"fmt"

	"github.com/AllenDang/giu"
)

func loop() {
	giu.Gizmo().Gizmos(giu.Custom(func() {
		fmt.Println("Hello world from global gizmos")
	})).Global()

	giu.Window("Gizmo demo").Layout(
		giu.Gizmo().Gizmos(
			giu.Custom(func() {
				fmt.Println("Hello world from window gizmos")
			}),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Gizmo (ImGuizmo) demo", 1280, 720, 0)
	wnd.Run(loop)
}
