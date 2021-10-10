package giu

import "github.com/AllenDang/imgui-go"

var _ Widget = &StackWidget{}

// StackWidget is used to ensure, that the build methods of all
// the widgets (layouts field) was called, but only the selected
// (visible field) layout is rendered (visible) in app.
type StackWidget struct {
	visible int32
	layouts []Widget
}

// Stack creates a new StackWidget.
func Stack(visible int32, layouts ...Widget) *StackWidget {
	return &StackWidget{
		visible: visible,
		layouts: layouts,
	}
}

// Build implements widget interface.
func (s *StackWidget) Build() {
	// save visible cursor position
	visiblePos := GetCursorScreenPos()

	// build visible layout
	// NOTE: it is important to build the visiblely showed layout before
	// building another ones, because the interactive layout widgets
	// (e.g. buttons) should be rendered on top of `stack`
	layouts := s.layouts

	if s.visible >= 0 && s.visible < int32(len(s.layouts)) {
		s.layouts[s.visible].Build()
		// remove visible layout from layouts list
		// nolint:gocritic // remove visible widget
		layouts = append(s.layouts[:s.visible], s.layouts[:s.visible+1]...)
	}

	// build invisible layouts with 0 alpha
	imgui.PushStyleVarFloat(imgui.StyleVarAlpha, 0)
	for _, l := range layouts {
		SetCursorScreenPos(visiblePos)
		l.Build()
	}
	imgui.PopStyleVar()
}
