package giu

import (
	"github.com/AllenDang/giu/imgui"
)

type Widget interface {
	Width() float32
	Build()
}

type SameLineWidget struct {
	BaseWidget
	widgets []Widget
}

func SameLine(widgets ...Widget) *SameLineWidget {
	return &SameLineWidget{
		BaseWidget: BaseWidget{width: 0},
		widgets:    widgets,
	}
}

func (l *SameLineWidget) Build() {
	for i, widget := range l.widgets {
		_, isTooltip := widget.(*TooltipWidget)
		_, isContextMenu := widget.(*ContextMenuWidget)
		_, isPopup := widget.(*PopupWidget)
		_, isTabItem := widget.(*TabItemWidget)

		if i > 0 && !isTooltip && !isContextMenu && !isPopup && !isTabItem {
			imgui.SameLine()
		}
		if widget.Width() != 0 {
			imgui.PushItemWidth(widget.Width())
		}

		widget.Build()

		if widget.Width() != 0 {
			imgui.PopItemWidth()
		}
	}

}

type Layout []Widget

func (l *Layout) Build() {
	for _, w := range *l {
		if w.Width() != 0 {
			imgui.PushItemWidth(w.Width())
		}

		w.Build()

		if w.Width() != 0 {
			imgui.PopItemWidth()
		}
	}
}

func (l *Layout) Width() float32 {
	return 0
}
