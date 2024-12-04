package giu

import (
	"image/color"

	"github.com/AllenDang/cimgui-go/imgui"
)

// SplitDirection represents a direction (vertical/horizontal) of splitting layout.
type SplitDirection uint8

const (
	// DirectionHorizontal is a horizontal line.
	DirectionHorizontal SplitDirection = 1 << iota
	// DirectionVertical is a vertical line.
	DirectionVertical
)

// SplitRefType describes how sashPos argument to the SplitLayout should be interpreted.
type SplitRefType byte

const (
	// SplitRefLeft is the default. Splitter placed counting from left/top layout's edge.
	SplitRefLeft SplitRefType = iota
	// SplitRefRight splitter placed counting from right/bottom layout's edge.
	SplitRefRight
	// SplitRefProc sashPos will be clamped in range [0, 1]. Then the position is considered a percent of GetAvailableRegion.
	SplitRefProc
)

var _ Disposable = &splitLayoutState{}

type splitLayoutState struct {
	delta float32
}

// Dispose implements disposable interface.
func (s *splitLayoutState) Dispose() {
	// noop
}

// SplitLayoutWidget creates two children with a line between them.
// This line can be moved by the user to adjust child sizes.
type SplitLayoutWidget struct {
	id                  ID
	direction           SplitDirection
	layout1             Widget
	layout2             Widget
	originItemSpacingX  float32
	originItemSpacingY  float32
	originFramePaddingX float32
	originFramePaddingY float32
	sashPos             *float32
	border              bool
	splitRefType        SplitRefType
}

// SplitLayout creates split layout widget.
func SplitLayout(direction SplitDirection, sashPos *float32, layout1, layout2 Widget) *SplitLayoutWidget {
	return &SplitLayoutWidget{
		direction: direction,
		sashPos:   sashPos,
		layout1:   layout1,
		layout2:   layout2,
		border:    true,
		id:        GenAutoID("SplitLayout"),
	}
}

// Border sets if children should have borders.
func (s *SplitLayoutWidget) Border(b bool) *SplitLayoutWidget {
	s.border = b
	return s
}

// ID allows to manually set splitter's id.
func (s *SplitLayoutWidget) ID(id ID) *SplitLayoutWidget {
	s.id = id
	return s
}

// SplitRefType allows to set how sashPos should be interpreted.
// Default is counting from left/top layout's edge in px.
func (s *SplitLayoutWidget) SplitRefType(refType SplitRefType) *SplitLayoutWidget {
	s.splitRefType = refType
	return s
}

// Build implements widget interface.
func (s *SplitLayoutWidget) Build() {
	splitLayoutState := s.getState()
	s.originItemSpacingX, s.originItemSpacingY = GetItemInnerSpacing()
	s.originFramePaddingX, s.originFramePaddingY = GetFramePadding()
	availableW, availableH := GetAvailableRegion()

	var layout Layout

	var sashPos float32

	switch s.splitRefType {
	case SplitRefLeft:
		sashPos = *s.sashPos
	case SplitRefRight:
		switch s.direction {
		case DirectionHorizontal:
			sashPos = availableH - *s.sashPos
		case DirectionVertical:
			sashPos = availableW - *s.sashPos
		}
	case SplitRefProc:
		if *s.sashPos < 0 {
			*s.sashPos = 0
		} else if *s.sashPos > 1 {
			*s.sashPos = 1
		}

		switch s.direction {
		case DirectionHorizontal:
			sashPos = availableH * *s.sashPos
		case DirectionVertical:
			sashPos = availableW * *s.sashPos
		}
	}

	sashPos += splitLayoutState.delta
	if sashPos < 1 {
		sashPos = 1
	}

	switch s.direction {
	case DirectionHorizontal:
		if sashPos >= availableH {
			sashPos = availableH
		}

		layout = Layout{
			Column(
				s.buildChild(Auto, sashPos, s.layout1),
				Splitter(DirectionHorizontal, &(splitLayoutState.delta)).Size(0, s.originItemSpacingY),
				s.buildChild(Auto, Auto, s.layout2),
			),
		}
	case DirectionVertical:
		if sashPos >= availableW {
			sashPos = availableW
		}

		layout = Layout{
			Row(
				s.buildChild(sashPos, Auto, s.layout1),
				Splitter(DirectionVertical, &(splitLayoutState.delta)).Size(s.originItemSpacingX, 0),
				s.buildChild(Auto, Auto, s.layout2),
			),
		}
	}

	PushItemSpacing(0, 0)
	layout.Build()
	PopStyle()

	s.encodeSashPos(sashPos, availableW, availableH)
}

func (s *SplitLayoutWidget) encodeSashPos(sashPos, availableW, availableH float32) {
	switch s.splitRefType {
	case SplitRefLeft:
		*s.sashPos = sashPos
	case SplitRefRight:
		switch s.direction {
		case DirectionHorizontal:
			*s.sashPos = availableH - sashPos
		case DirectionVertical:
			*s.sashPos = availableW - sashPos
		}
	case SplitRefProc:
		switch s.direction {
		case DirectionHorizontal:
			*s.sashPos = sashPos / availableH
		case DirectionVertical:
			*s.sashPos = sashPos / availableW
		}
	}
}

func (s *SplitLayoutWidget) restoreItemSpacing(layout Widget) Layout {
	return Layout{
		Custom(func() {
			PushItemSpacing(s.originItemSpacingX, s.originItemSpacingY)
			PushFramePadding(s.originFramePaddingX, s.originFramePaddingY)
			// Restore Child bg color
			bgColor := imgui.StyleColorVec4(imgui.ColChildBg)
			PushStyleColor(StyleColorChildBg, Vec4ToRGBA(*bgColor))
		}),
		layout,
		Custom(func() {
			PopStyleColor()
			PopStyleV(2)
		}),
	}
}

// Build Child panel. If layout is a SplitLayout, set the frame padding to zero.
func (s *SplitLayoutWidget) buildChild(width, height float32, layout Widget) Widget {
	return Layout{
		Custom(func() {
			_, isSplitLayoutWidget := layout.(*SplitLayoutWidget)
			hasFramePadding := isSplitLayoutWidget || !s.border
			hasBorder := !isSplitLayoutWidget && s.border

			if hasFramePadding {
				PushFramePadding(0, 0)
			}

			PushStyleColor(StyleColorChildBg, color.RGBA{R: 0, G: 0, B: 0, A: 0})

			Child().
				Border(hasBorder).
				Size(width, height).
				Layout(s.restoreItemSpacing(layout)).
				Build()

			PopStyleColor()

			if hasFramePadding {
				PopStyle()
			}
		}),
	}
}

func (s *SplitLayoutWidget) getState() (state *splitLayoutState) {
	if state = GetState[splitLayoutState](Context, s.id); state == nil {
		state = &splitLayoutState{delta: 0.0}
		SetState(Context, s.id, state)
	}

	return state
}
