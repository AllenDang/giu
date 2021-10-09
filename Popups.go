package giu

import (
	"github.com/AllenDang/imgui-go"
)

// OpenPopup opens a popup with specified id.
// NOTE: you need to build this popup first (see Pop(Modal)Widget)
func OpenPopup(name string) {
	imgui.OpenPopup(name)
}

// CloseCurrentPopup closes currently opened popup.
// If no popups opened, no action will be taken.
func CloseCurrentPopup() {
	imgui.CloseCurrentPopup()
}

var _ Widget = &PopupWidget{}

type PopupWidget struct {
	name   string
	flags  WindowFlags
	layout Layout
}

func Popup(name string) *PopupWidget {
	return &PopupWidget{
		name:   tStr(name),
		flags:  0,
		layout: nil,
	}
}

func (p *PopupWidget) Flags(flags WindowFlags) *PopupWidget {
	p.flags = flags
	return p
}

func (p *PopupWidget) Layout(widgets ...Widget) *PopupWidget {
	p.layout = Layout(widgets)
	return p
}

// Build implements Widget interface.
func (p *PopupWidget) Build() {
	if imgui.BeginPopup(p.name, int(p.flags)) {
		p.layout.Build()
		imgui.EndPopup()
	}
}

var _ Widget = &PopupModalWidget{}

type PopupModalWidget struct {
	name   string
	open   *bool
	flags  WindowFlags
	layout Layout
}

func PopupModal(name string) *PopupModalWidget {
	return &PopupModalWidget{
		name:   tStr(name),
		open:   nil,
		flags:  WindowFlagsNoResize,
		layout: nil,
	}
}

func (p *PopupModalWidget) IsOpen(open *bool) *PopupModalWidget {
	p.open = open
	return p
}

func (p *PopupModalWidget) Flags(flags WindowFlags) *PopupModalWidget {
	p.flags = flags
	return p
}

func (p *PopupModalWidget) Layout(widgets ...Widget) *PopupModalWidget {
	p.layout = Layout(widgets)
	return p
}

// Build implements Widget interface.
func (p *PopupModalWidget) Build() {
	if imgui.BeginPopupModalV(p.name, p.open, int(p.flags)) {
		p.layout.Build()
		imgui.EndPopup()
	}
}
