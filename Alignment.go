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

// Align sets widgets alignment.
// usage: see examples/align
//
// FIXME: all widgets will be build twice
// it means, that if you have e.g. CustomWidget it could do unexpected things.
// Example:
// Align(AlignToCenter).To(
//   Custom(func() { fmt.Println("running custom widget") }),
// )
// will print the message two times per frame.
//
// BUG: DatePickerWidget doesn't work properly
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
	a.layout.Range(func(item Widget) {
		var width float32
		if item != nil {
			item.Build()
			size := imgui.GetItemRectSize()
			width = size.X
		}

		widgetsWidths = append(widgetsWidths, width)
	})
	imgui.PopStyleVar()

	// reset cursor pos
	SetCursorPos(startPos)

	// ALIGN WIDGETS
	idx := 0o0
	a.layout.Range(func(item Widget) {
		if item == nil {
			return
		}

		w := widgetsWidths[idx]
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

		idx++
	})
}
