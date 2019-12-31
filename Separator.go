package giu

import "github.com/AllenDang/giu/imgui"

type SeparatorWidget struct {
	BaseWidget
}

func Separator() *SeparatorWidget {
	return &SeparatorWidget{
		BaseWidget: BaseWidget{width: 0},
	}
}

func (s *SeparatorWidget) Build() {
	imgui.Separator()
}
