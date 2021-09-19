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
// FIXME: DONOT put giu widgets inside of CustomWidget function in
// Align setter. CustomWidgets will be skip in alignment process
//
// BUG: DatePickerWidget doesn't work properly
func Align(at AlignmentType) *AlignmentSetter {
	return &AlignmentSetter{
		alignType: at,
		id:        GenAutoID("alignSetter"),
	}
}

// To sets a layout, alignment should be applied to
func (a *AlignmentSetter) To(widgets ...Widget) *AlignmentSetter {
	a.layout = Layout(widgets)
	return a
}

// ID allows to manually set AlignmentSetter ID (it shouldn't be used
// in a normal conditions)
func (a *AlignmentSetter) ID(id string) *AlignmentSetter {
	a.id = id
	return a
}

func (a *AlignmentSetter) Build() {
	if a.layout == nil {
		return
	}

	// WORKAROUND: get widgets widths rendering them with 100% transparency
	// to align them later
	a.layout.Range(func(item Widget) {
		// if item is inil, just skip it
		if item == nil {
			return
		}

		// exclude some widgets from alignment process
		switch item.(type) {
		case *CustomWidget:
			item.Build()
			return
		}

		// save cursor position before rendering
		currentPos := GetCursorPos()

		// render widget in `dry` mode
		imgui.PushStyleVarFloat(imgui.StyleVarAlpha, 0)
		item.Build()
		imgui.PopStyleVar()

		// save widget's width
		size := imgui.GetItemRectSize()
		w := size.X

		availableW, _ := GetAvailableRegion()

		// set cursor position to align the widget
		switch a.alignType {
		case AlignLeft:
			SetCursorPos(currentPos)
		case AlignCenter:
			SetCursorPos(image.Pt(int(availableW/2-w/2), currentPos.Y))
		case AlignRight:
			SetCursorPos(image.Pt(int(availableW-w), currentPos.Y))
		default:
			panic(fmt.Sprintf("giu: (*AlignSetter).Build: unknown align type %d", a.alignType))
		}

		// build aligned widget
		item.Build()
	})
}
