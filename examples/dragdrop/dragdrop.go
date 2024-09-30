package main

import (
	"fmt"
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"

	g "github.com/AllenDang/giu"
)

var dropTarget string = "Drop here"

func loop() {
	g.SingleWindow().Layout(
		g.Row(
			g.Custom(func() {
				g.Button("Drag me: 9").Build()

				if imgui.BeginDragDropSource() {
					data := int(9)
					imgui.SetDragDropPayload(
						"DND_DEMO",
						uintptr(unsafe.Pointer(&data)),
						uint64(unsafe.Sizeof(data)),
					)
					g.Label("9").Build()
					imgui.EndDragDropSource()
				}
			}),
			g.Custom(func() {
				g.Button("Drag me: 10").Build()

				if imgui.BeginDragDropSource() {
					data := 10
					imgui.SetDragDropPayload(
						"DND_DEMO",
						uintptr(unsafe.Pointer(&data)),
						uint64(unsafe.Sizeof(data)),
					)
					g.Label("10").Build()
					imgui.EndDragDropSource()
				}
			}),
		),
		g.InputTextMultiline(&dropTarget).Size(g.Auto, g.Auto).Flags(g.InputTextFlagsReadOnly),
		g.Custom(func() {
			if imgui.BeginDragDropTarget() {
				payload := imgui.AcceptDragDropPayload("DND_DEMO")
				if payload != nil && payload.CData != nil {
					dropTarget = fmt.Sprintf("Dropped value: %d", *(*int)(unsafe.Pointer(payload.Data()))) //nolint:govet // its cimgui-go design
				}

				imgui.EndDragDropTarget()
			}
		}),
	)
}

func main() {
	wnd := g.NewMasterWindow("Drag and Drop", 600, 400, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
