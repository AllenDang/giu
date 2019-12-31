package giu

import "github.com/AllenDang/giu/imgui"

type LabelWidget struct {
	BaseWidget
	caption string
}

func LabelV(caption string, width float32) *LabelWidget {
	return &LabelWidget{caption: caption, BaseWidget: BaseWidget{width}}
}

func Label(caption string) *LabelWidget {
	return LabelV(caption, 0)
}

func (l *LabelWidget) Build() {
	imgui.Text(l.caption)
}
