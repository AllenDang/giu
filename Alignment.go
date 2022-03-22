package giu

import (
	"fmt"
	"image"

	"github.com/AllenDang/imgui-go"
)

// AlignmentType represents a bype of alignment to use with AlignSetter.
type AlignmentType byte

const (
	// AlignLeft is here just for clearity.
	// if set, no action is taken so don't use it.
	AlignLeft AlignmentType = iota
	// AlignCenter centers widget.
	AlignCenter
	// AlignRight aligns a widget to right side of window.
	AlignRight
)

// AlignManually allows to apply alignment manually.
// As long as AlignSetter is really EXPERIMENTAL feature
// and may fail randomly, the following method is supposed to
// always work, as long as you set it up correctly.
// To use it just pass a single widget with its exact width.
// be sure to apply widget's size by using "Size" method!
// forceApplyWidth argument allows you to ask giu to force-set width
// of `widget`
// NOTE that forcing width doesn't work for each widget type! For example
// Button won't work because its size is set by argument to imgui call
// not PushWidth api.
func AlignManually(alignmentType AlignmentType, widget Widget, widgetWidth float32, forceApplyWidth bool) Widget {
	return Custom(func() {
		spacingX, _ := GetItemSpacing()
		availableW, _ := GetAvailableRegion()

		var dummyX float32

		switch alignmentType {
		case AlignLeft:
			widget.Build()
			return
		case AlignCenter:
			dummyX = (availableW-widgetWidth)/2 - spacingX
		case AlignRight:
			dummyX = availableW - widgetWidth - spacingX
		}

		Dummy(dummyX, 0).Build()

		if forceApplyWidth {
			PushItemWidth(widgetWidth)
			defer PopItemWidth()
		}

		imgui.SameLine()
		widget.Build()
	})
}

var _ Widget = &AlignmentSetter{}

// AlignmentSetter allows to align to right / center a widget or widgets group.
// NOTE: Because of AlignSetter uses experimental GetWidgetWidth,
// it is experimental too.
// usage: see examples/align
//
// list of known bugs:
// - BUG: DatePickerWidget doesn't work properly
// - BUG: there is some bug with SelectableWidget
// - BUG: ComboWidget and ComboCustomWidgets doesn't work properly.
type AlignmentSetter struct {
	alignType AlignmentType
	layout    Layout
	id        string
}

// Align sets widgets alignment.
func Align(at AlignmentType) *AlignmentSetter {
	return &AlignmentSetter{
		alignType: at,
		id:        GenAutoID("alignSetter"),
	}
}

// To sets a layout, alignment should be applied to.
func (a *AlignmentSetter) To(widgets ...Widget) *AlignmentSetter {
	a.layout = Layout(widgets)
	return a
}

// ID allows to manually set AlignmentSetter ID
// NOTE: there isn't any known reason to use this method, however
// it is here for some random cases. YOU DON'T NEED TO USE IT
// in normal conditions.
func (a *AlignmentSetter) ID(id string) *AlignmentSetter {
	a.id = id
	return a
}

// Build implements Widget interface.
func (a *AlignmentSetter) Build() {
	if a.layout == nil {
		return
	}

	a.layout.Range(func(item Widget) {
		// if item is inil, just skip it
		if item == nil {
			return
		}

		switch item.(type) {
		// ok, it doesn't make sense to align again :-)
		case *AlignmentSetter:
			item.Build()
			return
		// there is a bug with selectables and combos, so skip them for now
		case *SelectableWidget, *ComboWidget, *ComboCustomWidget:
			item.Build()
			return
		}

		currentPos := GetCursorPos()
		w := GetWidgetWidth(item)
		availableW, _ := GetAvailableRegion()
		// we need to increase available region by 2 * window padding (X),
		// because GetCursorPos considers it
		paddingW, _ := GetWindowPadding()
		availableW += 2 * paddingW

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

// GetWidgetWidth returns a width of widget
// NOTE: THIS IS A BETA SOLUTION and may contain bugs
// in most cases, you may want to use supported by imgui GetItemRectSize.
// There is an upstream issue for this problem:
// https://github.com/ocornut/imgui/issues/3714
//
// This function is just a workaround used in giu.
//
// NOTE: user-definied widgets, which contains more than one
// giu widget will be processed incorrectly (only width of the last built
// widget will be processed)
//
// here is a list of known bugs:
// - BUG: user can interact with invisible widget (created by GetWidgetWidth)
//   - https://github.com/AllenDang/giu/issues/341
//   - https://github.com/ocornut/imgui/issues/4588
//
// if you find anything else, please report it on
// https://github.com/AllenDang/giu Any contribution is appreciated!
func GetWidgetWidth(w Widget) (result float32) {
	imgui.PushID(GenAutoID("GetWIdgetWidthMeasurement"))
	defer imgui.PopID()

	// save cursor position before rendering
	currentPos := GetCursorPos()

	// render widget in `dry` mode
	imgui.PushStyleVarFloat(imgui.StyleVarAlpha, 0)
	w.Build()
	imgui.PopStyleVar()

	// save widget's width
	// check cursor position
	imgui.SameLine()
	spacingW, _ := GetItemSpacing()
	result = float32(GetCursorPos().X-currentPos.X) - spacingW

	SetCursorPos(currentPos)

	return result
}
