package giu

import "github.com/AllenDang/giu/imgui"

type ContextMenuWidget struct {
	BaseWidget
	label       string
	mouseButton int
	layout      Layout
}

func ContextMenuV(label string, mouseButton int, layout Layout) *ContextMenuWidget {
	return &ContextMenuWidget{
		BaseWidget:  BaseWidget{width: 0},
		label:       label,
		mouseButton: mouseButton,
		layout:      layout,
	}
}

func ContextMenu(layout Layout) *ContextMenuWidget {
	return ContextMenuV("", 1, layout)
}

func (c *ContextMenuWidget) Build() {
	if imgui.BeginPopupContextItemV(c.label, c.mouseButton) {
		c.layout.Build()
		imgui.EndPopup()
	}
}
