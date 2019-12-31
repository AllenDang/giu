package giu

import "github.com/AllenDang/giu/imgui"

type TabItemWidget struct {
	BaseWidget
	label  string
	open   *bool
	flags  int
	layout Layout
}

func TabItemV(label string, open *bool, flags int, layout Layout) *TabItemWidget {
	return &TabItemWidget{
		BaseWidget: BaseWidget{width: 0},
		label:      label,
		open:       open,
		flags:      flags,
		layout:     layout,
	}
}

func TabItem(label string, layout Layout) *TabItemWidget {
	return TabItemV(label, nil, 0, layout)
}

func (ti *TabItemWidget) Build() {
	if imgui.BeginTabItemV(ti.label, ti.open, ti.flags) {
		ti.layout.Build()
		imgui.EndTabItem()
	}
}

type TabBarWidget struct {
	BaseWidget
	id    string
	flags int
	tabs  []*TabItemWidget
}

func TabBarV(id string, flags int, width float32, tabs []*TabItemWidget) *TabBarWidget {
	return &TabBarWidget{
		BaseWidget: BaseWidget{width: width},
		id:         id,
		flags:      flags,
		tabs:       tabs,
	}
}

func TabBar(id string, tabs []*TabItemWidget) *TabBarWidget {
	return TabBarV(id, 0, 0, tabs)
}

func (t *TabBarWidget) Build() {
	if imgui.BeginTabBarV(t.id, t.flags) {
		for _, tab := range t.tabs {
			tab.Build()
		}
		imgui.EndTabBar()
	}
}
