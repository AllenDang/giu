package main

import (
	"bytes"
	"fmt"

	g "github.com/ianling/giu"
	"github.com/ianling/imgui-go"
)

var (
	dropTarget string = "Drop here"
)

func loop() {
	g.SingleWindow("Drag and Drop").Layout(
		g.Line(
			g.Button("Drag me: 9"),
			g.Custom(func() {
				if imgui.BeginDragDropSource(imgui.DragDropFlagsNone) {
					imgui.SetDragDropPayload("DND_DEMO", []byte("9"), imgui.ConditionNone)
					g.Label("9").Build()
					imgui.EndDragDropSource()
				}
			}),
			g.Button("Drag me: 10"),
			g.Custom(func() {
				if imgui.BeginDragDropSource(imgui.DragDropFlagsNone) {
					imgui.SetDragDropPayload("DND_DEMO", []byte("10"), imgui.ConditionNone)
					g.Label("10").Build()
					imgui.EndDragDropSource()
				}
			}),
		),
		g.InputTextMultiline("##DropTarget", &dropTarget).Size(-1, -1).Flags(g.InputTextFlags_ReadOnly),
		g.Custom(func() {
			if imgui.BeginDragDropTarget() {
				payload := imgui.AcceptDragDropPayload("DND_DEMO", imgui.DragDropFlagsNone)
				if !bytes.Equal(payload, []byte{}) {
					dropTarget = fmt.Sprintf("Dropped value: %s", string(payload))
				}
				imgui.EndDragDropTarget()
			}
		}),
	).Build()
}

func main() {
	wnd := g.NewMasterWindow("Drag and Drop", 600, 400, g.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}
