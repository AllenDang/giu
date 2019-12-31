package giu

import "github.com/AllenDang/giu/imgui"

type MenuItemWidget struct {
	BaseWidget
	label    string
	selected bool
	enabled  bool
	clicked  func()
}

func MenuItemV(label string, selected, enabled bool, clicked func()) *MenuItemWidget {
	return &MenuItemWidget{
		BaseWidget: BaseWidget{width: 0},
		label:      label,
		selected:   selected,
		enabled:    enabled,
		clicked:    clicked,
	}
}

func MenuItem(label string, clicked func()) *MenuItemWidget {
	return MenuItemV(label, false, true, clicked)
}

func (mi *MenuItemWidget) Build() {
	if imgui.MenuItemV(mi.label, "", mi.selected, mi.enabled) && mi.clicked != nil {
		mi.clicked()
	}
}

type MenuWidget struct {
	BaseWidget
	label   string
	enabled bool
	layout  Layout
}

func MenuV(label string, enabled bool, layout Layout) *MenuWidget {
	return &MenuWidget{
		BaseWidget: BaseWidget{width: 0},
		label:      label,
		enabled:    enabled,
		layout:     layout,
	}
}

func Menu(label string, layout Layout) *MenuWidget {
	return MenuV(label, true, layout)
}

func (m *MenuWidget) Build() {
	if imgui.BeginMenuV(m.label, m.enabled) {
		m.layout.Build()
		imgui.EndMenu()
	}
}

type MenuBarWidget struct {
	BaseWidget
	layout Layout
}

func MenuBar(layout Layout) *MenuBarWidget {
	return &MenuBarWidget{
		BaseWidget: BaseWidget{width: 0},
		layout:     layout,
	}
}

func (mb *MenuBarWidget) Build() {
	if imgui.BeginMenuBar() {
		mb.layout.Build()
		imgui.EndMenuBar()
	}
}

func MainMenuBar(layout Layout) {
	if imgui.BeginMainMenuBar() {
		layout.Build()
		imgui.EndMainMenuBar()
	}
}
