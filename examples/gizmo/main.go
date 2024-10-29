package main

import (
	"fmt"

	"github.com/AllenDang/cimgui-go/imguizmo"
	"github.com/AllenDang/giu"
)

var (
	view = giu.NewHumanReadableMatrix().
		Transform(0, 0, -7).
		Rotation(10, 5, 0).
		Scale(1, 1, 1)

	projection = giu.NewHumanReadableMatrix().SetMatrix([]float32{
		2.3787, 0, 0, 0,
		0, 3.1716, 0, 0,
		0, 0, -1.0002, -1,
		0, 0, -0.2, 0,
	})

	cube = giu.NewHumanReadableMatrix().
		Transform(0.5, 0.5, 0.5).
		Rotation(0, 0, 0).
		Scale(1, 1, 1)

	zmoOP   = int32(imguizmo.TRANSLATE)
	zmoMODE = int32(imguizmo.LOCAL)

	Bounds = []float32{-0.5, -0.5, -0.5, 0.5, 0.5, 0.5}

	// GizmoControls can be enabled in editor and allows mouse control over gizmo
	GizmoControls bool
	UsingGizmo    bool
)

func gizmos() []giu.GizmoI {
	return []giu.GizmoI{
		giu.Grid(),
		giu.Cube(cube), //.Manipulate(),
		giu.Manipulate(cube),
		giu.Custom(func() {
			fmt.Println(cube)
			/*
				imguizmo.ViewManipulateFloat(
					view.Matrix(),
					1,
					imgui.Vec2{128, 128},
					imgui.Vec2{128, 128},
					0x01010101,
				)
			*/
		}),
	}
}

func loop() {
	giu.Gizmo(view, projection).Gizmos(gizmos()...).Global()

	giu.Window("Gizmo demo").Layout(
		giu.Gizmo(view, projection).Gizmos(
			gizmos()...,
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Gizmo (ImGuizmo) demo", 1280, 720, 0)
	wnd.Run(loop)
}
