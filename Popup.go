package giu

import "github.com/AllenDang/giu/imgui"

type PopupWidget struct {
	BaseWidget
	name   string
	open   *bool
	flags  int
	layout Layout
}

func PopupV(name string, open *bool, flags int, layout Layout) *PopupWidget {
	return &PopupWidget{
		BaseWidget: BaseWidget{width: 0},
		name:       name,
		open:       open,
		flags:      flags,
		layout:     layout,
	}
}

func Popup(name string, layout Layout) *PopupWidget {
	return PopupV(name, nil, imgui.WindowFlagsNoResize, layout)
}

func OpenPopup(name string) {
	imgui.OpenPopup(name)
}

func CloseCurrentPopup() {
	imgui.CloseCurrentPopup()
}

func (p *PopupWidget) Build() {
	if imgui.BeginPopupModalV(p.name, p.open, p.flags) {
		p.layout.Build()
		imgui.EndPopup()
	}
}
