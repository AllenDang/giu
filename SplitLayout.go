package giu

import (
	"fmt"
	"image/color"
)

type SplitDirection uint8

const (
	DirectionHorizontal SplitDirection = 1 << iota
	DirectionVertical
)

type SplitLayoutState struct {
	delta   float32
	sashPos float32
}

func (s *SplitLayoutState) Dispose() {
	// Nothing to do here.
}

type SplitLayoutWidget struct {
	id                  string
	direction           SplitDirection
	layout1             Widget
	layout2             Widget
	originItemSpacingX  float32
	originItemSpacingY  float32
	originFramePaddingX float32
	originFramePaddingY float32
	sashPos             float32
	border              bool
}

func SplitLayout(direction SplitDirection, border bool, sashPos float32, layout1, layout2 Widget) *SplitLayoutWidget {
	return &SplitLayoutWidget{
		direction: direction,
		sashPos:   sashPos,
		layout1:   layout1,
		layout2:   layout2,
		border:    border,
	}
}

func (s *SplitLayoutWidget) restoreItemSpacing(layout Widget) Layout {
	return Layout{
		Custom(func() {
			PushItemSpacing(s.originItemSpacingX, s.originItemSpacingY)
			PushFramePadding(s.originFramePaddingX, s.originFramePaddingY)
			// Restore Child bg color
			PushStyleColor(StyleColorChildBg, color.RGBA{R: 0x26, G: 0x2e, B: 0x38, A: 0xff})
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
	_, isSplitLayoutWidget := layout.(*SplitLayoutWidget)

	return Layout{
		Custom(func() {
			if isSplitLayoutWidget || !s.border {
				PushFramePadding(0, 0)
			}
		}),
		Style().SetColor(StyleColorChildBg, color.RGBA{R: 0x1c, G: 0x26, B: 0x2b, A: 0xff}).To(
			Child().Border(!isSplitLayoutWidget && s.border).Size(width, height).Layout(s.restoreItemSpacing(layout)),
		),
		Custom(func() {
			if isSplitLayoutWidget || !s.border {
				PopStyle()
			}
		}),
	}
}

func (s *SplitLayoutWidget) Build() {
	s.id = GenAutoID("SplitLayout")

	var splitLayoutState *SplitLayoutState
	// Register state
	stateId := fmt.Sprintf("SplitLayout_%s", s.id)
	if state := Context.GetState(stateId); state == nil {
		splitLayoutState = &SplitLayoutState{delta: 0.0, sashPos: s.sashPos}
		Context.SetState(stateId, splitLayoutState)
	} else {
		splitLayoutState = state.(*SplitLayoutState)
	}

	itemSpacingX, itemSpacingY := GetItemInnerSpacing()
	s.originItemSpacingX, s.originItemSpacingY = itemSpacingX, itemSpacingY

	s.originFramePaddingX, s.originFramePaddingY = GetFramePadding()
	s.originFramePaddingX /= Context.GetPlatform().GetContentScale()
	s.originFramePaddingY /= Context.GetPlatform().GetContentScale()

	var layout Layout

	splitLayoutState.sashPos += splitLayoutState.delta

	if s.direction == DirectionHorizontal {
		layout = Layout{
			Row(
				s.buildChild(splitLayoutState.sashPos, 0, s.layout1),
				VSplitter(&(splitLayoutState.delta)).Size(itemSpacingX, 0),
				s.buildChild(0, 0, s.layout2),
			),
		}
	} else {
		layout = Layout{
			s.buildChild(0, splitLayoutState.sashPos, s.layout1),
			HSplitter(&(splitLayoutState.delta)).Size(0, itemSpacingY),
			s.buildChild(0, 0, s.layout2),
		}
	}

	PushItemSpacing(0, 0)
	layout.Build()
	PopStyle()
}
