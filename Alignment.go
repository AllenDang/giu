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

type alignSetterState struct {
	widgets []float32
}

func (s *alignSetterState) Dispose() {
	// noop
}

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

// BUG: currently layout cannot be changed
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

	stateID := fmt.Sprintf("%s_state", a.id)
	var state *alignSetterState

	// WORKAROUND: to get widgets width in further code, save them in state
	if s := Context.GetState(stateID); s == nil {
		newState := &alignSetterState{
			widgets: make([]float32, 0),
		}

		Context.SetState(stateID, newState)

		state = newState

		getItemWidth := func(i Widget) float32 {
			if i == nil {
				return 0
			}

			i.Build()
			w := imgui.GetItemRectSize()
			return w.X
		}

		for _, item := range a.layout {
			state.widgets = append(state.widgets, getItemWidth(item))
		}
	} else {
		state = s.(*alignSetterState)
	}

	for i, item := range a.layout {
		if item == nil {
			continue
		}
		w := state.widgets[i]
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
