package giu

import (
	"fmt"
	"image"

	"github.com/AllenDang/imgui-go"
)

type AlignmentType byte

const (
	AlignLeft AlignmentType = iota
	AlignCenter
	AlignRight
)

type AlignmentSetter struct {
	alignType AlignmentType
	layout    Layout
	id        string
}

func Align(at AlignmentType) *AlignmentSetter {
	return &AlignmentSetter{
		alignType: at,
		id:        GenAutoID("alignSetter"),
	}
}

func (a *AlignmentSetter) To(widgets ...Widget) *AlignmentSetter {
	a.layout = Layout(widgets)
	return a
}

func (a *AlignmentSetter) ID(id string) *AlignmentSetter {
	a.id = id
	return a
}

func (a *AlignmentSetter) Build() {
	if a.layout == nil {
		return
	}

	// WORKAROUND: get widgets widths rendering them with 100% transparency
	// first save start cursor position
	startPos := GetCursorPos()

	widgetsWidths := make([]float32, 0)

	// render widgets with 0 alpha and store thems widths
	imgui.PushStyleVarFloat(imgui.StyleVarID(StyleVarAlpha), 0)
	for _, item := range a.layout {
		var width float32
		if item != nil {
			item.Build()
			size := imgui.GetItemRectSize()
			width = size.X
		}

		widgetsWidths = append(widgetsWidths, width)
	}
	imgui.PopStyleVar()

	// reset cursor pos
	SetCursorPos(startPos)

	// ALIGN WIDGETS
	for i, item := range a.layout {
		if item == nil {
			continue
		}

		w := widgetsWidths[i]
		currentPos := GetCursorPos()
		availableW, _ := GetAvailableRegion()
		switch a.alignType {
		case AlignLeft:
			// noop
		case AlignCenter:
			SetCursorPos(image.Pt(int(availableW/2-w/2), int(currentPos.Y)))
		case AlignRight:
			SetCursorPos(image.Pt(int(availableW-w), int(currentPos.Y)))
		default:
			panic(fmt.Sprintf("giu: (*AlignSetter).Build: unknown align type %d", a.alignType))
		}

		item.Build()
	}
}
