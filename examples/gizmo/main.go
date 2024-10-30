package main

import (
	"image/color"

	"github.com/AllenDang/giu"
)

var (
	view = giu.NewViewMatrix().
		Transform(0, 0, -7).
		Rotation(10, 5, 0).
		Scale(1, 1, 1)

	projection = giu.NewProjectionMatrix().
			FOV(giu.Deg2Rad(30)).
			Aspect(1280.0 / 720.0)

	cube = giu.NewViewMatrix().
		Transform(0.5, 0.5, 0.5).
		Rotation(0, 0, 0).
		Scale(1, 2, 1)
)

func gizmos() []giu.GizmoI {
	return []giu.GizmoI{
		giu.Grid(),
		giu.Cube(cube).Manipulate(),
		giu.ViewManipulate().Color(
			color.RGBA{
				R: 45,
				G: 15,
				B: 121,
				A: 255,
			},
		),
	}
}

func loop() {
	giu.Gizmo(view, projection).Gizmos(gizmos()...).Global()

	giu.Window("Gizmo demo").
		Size(300, 300).
		Pos(50, 300).
		Layout(
			giu.Gizmo(view, projection).Gizmos(
				gizmos()...,
			).Disabled(true),
		)
}

func main() {
	wnd := giu.NewMasterWindow("Gizmo (ImGuizmo) demo", 1280, 720, 0)
	wnd.Run(loop)
}
