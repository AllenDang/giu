package giu

import "fmt"

type Direction uint8

const (
	DirectionHorizontal Direction = 1 << iota
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
	id                 string
	direction          Direction
	layout1            Widget
	layout2            Widget
	originItemSpacingX float32
	originItemSpacingY float32
	sashPos            float32
}

func SplitLayout(id string, direction Direction, sashPos float32, layout1, layout2 Widget) *SplitLayoutWidget {
	return &SplitLayoutWidget{
		id:        id,
		direction: direction,
		sashPos:   sashPos,
		layout1:   layout1,
		layout2:   layout2,
	}
}

func (s *SplitLayoutWidget) restoreItemSpacing(layout Widget) Layout {
	return Layout{
		Custom(func() {
			PushItemSpacing(s.originItemSpacingX, s.originItemSpacingY)
		}),
		layout,
		Custom(func() {
			PopStyle()
		}),
	}
}

// Build Child panel. If layout is a SplitLayout, set the frame padding to zero.
func (s *SplitLayoutWidget) buildChild(id string, width, height float32, layout Widget) Widget {
	_, isSplitLayoutWidget := layout.(*SplitLayoutWidget)

	return Layout{
		Custom(func() {
			if isSplitLayoutWidget {
				PushFramePadding(0, 0)
			}
		}),
		Child(id, !isSplitLayoutWidget, width, height, 0, s.restoreItemSpacing(layout)),
		Custom(func() {
			if isSplitLayoutWidget {
				PopStyle()
			}
		}),
	}
}

func (s *SplitLayoutWidget) Build() {
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

	var layout Layout

	splitLayoutState.sashPos += splitLayoutState.delta

	if s.direction == DirectionHorizontal {
		layout = Layout{
			Line(
				s.buildChild(fmt.Sprintf("%s_layout1", stateId), splitLayoutState.sashPos, 0, s.layout1),
				VSplitter(fmt.Sprintf("%s_vsplitter", stateId), itemSpacingX, 0, &(splitLayoutState.delta)),
				s.buildChild(fmt.Sprintf("%s_layout2", stateId), 0, 0, s.layout2),
			),
		}
	} else {
		layout = Layout{
			s.buildChild(fmt.Sprintf("%s_layout1", stateId), 0, splitLayoutState.sashPos, s.layout1),
			HSplitter(fmt.Sprintf("%s_hsplitter", stateId), 0, itemSpacingY, &(splitLayoutState.delta)),
			s.buildChild(fmt.Sprintf("%s_layout2", stateId), 0, 0, s.layout2),
		}
	}

	PushItemSpacing(0, 0)
	layout.Build()
	PopStyle()
}
