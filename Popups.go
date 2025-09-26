package giu

import (
	"fmt"

	"github.com/AllenDang/cimgui-go/imgui"
)

// OpenPopup opens a popup with specified id.
func OpenPopup(name string) {
	SetState[popupState](Context, popupStateID(name), &popupState{open: true})
}

// CloseCurrentPopup closes currently opened popup.
// If no popups opened, no action will be taken.
func CloseCurrentPopup() {
	imgui.CloseCurrentPopup()
}

func popupStateID(name string) ID {
	return ID(fmt.Sprintf("%s##state", name))
}

func applyPopupState(name string) {
	var state *popupState
	if state = GetState[popupState](Context, popupStateID(name)); state == nil {
		state = &popupState{open: false}
	}

	if state.open {
		imgui.OpenPopupStr(name)
		SetState[popupState](Context, popupStateID(name), &popupState{open: false})
	}
}

var _ Disposable = &popupState{}

type popupState struct {
	open bool
}

func (s *popupState) Dispose() {
	// noop
}

var _ Widget = &PopupWidget{}

// PopupWidget  is a window which appears next to the mouse cursor.
// For instance it is used to display color palette in ColorSelectWidget.
type PopupWidget struct {
	name   string
	flags  WindowFlags
	layout Layout
}

// Popup creates new popup widget.
func Popup(name string) *PopupWidget {
	return &PopupWidget{
		name:   name,
		flags:  0,
		layout: nil,
	}
}

// Flags sets popup's flags.
func (p *PopupWidget) Flags(flags WindowFlags) *PopupWidget {
	p.flags = flags
	return p
}

// Layout sets popup's layout.
func (p *PopupWidget) Layout(widgets ...Widget) *PopupWidget {
	p.layout = Layout(widgets)
	return p
}

// Build implements Widget interface.
func (p *PopupWidget) Build() {
	applyPopupState(p.name)

	if imgui.BeginPopupV(p.name, imgui.WindowFlags(p.flags)) {
		p.layout.Build()
		imgui.EndPopup()
	}
}

var _ Widget = &PopupModalWidget{}

// PopupModalWidget is a popup window that block every interactions behind it, cannot be closed by
// user, adds a dimming background, has a title bar.
type PopupModalWidget struct {
	name   string
	open   *bool
	flags  WindowFlags
	layout Layout
}

// PopupModal creates new popup modal widget.
func PopupModal(name string) *PopupModalWidget {
	return &PopupModalWidget{
		name:   name,
		open:   nil,
		flags:  WindowFlagsNoResize,
		layout: nil,
	}
}

// IsOpen allows to control popup's state
// NOTE: changing opens' value will not result in changing popup's state
// if OpenPopup(...) wasn't called!
func (p *PopupModalWidget) IsOpen(open *bool) *PopupModalWidget {
	p.open = open
	return p
}

// Flags allows to specify popup's flags.
func (p *PopupModalWidget) Flags(flags WindowFlags) *PopupModalWidget {
	p.flags = flags
	return p
}

// Layout sets layout.
func (p *PopupModalWidget) Layout(widgets ...Widget) *PopupModalWidget {
	p.layout = Layout(widgets)
	return p
}

// Build implements Widget interface.
func (p *PopupModalWidget) Build() {
	applyPopupState(p.name)

	if imgui.BeginPopupModalV(p.name, p.open, imgui.WindowFlags(p.flags)) {
		p.layout.Build()
		imgui.EndPopup()
	}
}
