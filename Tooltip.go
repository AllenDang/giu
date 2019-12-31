package giu

import "github.com/AllenDang/giu/imgui"

type TooltipWidget struct {
	BaseWidget
	text string
}

func Tooltip(text string) *TooltipWidget {
	return &TooltipWidget{
		BaseWidget: BaseWidget{width: 0},
		text:       text,
	}
}

func (t *TooltipWidget) Build() {
	if imgui.IsItemHovered() {
		imgui.SetTooltip(t.text)
	}
}
