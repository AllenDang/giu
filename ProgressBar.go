package giu

import "github.com/AllenDang/giu/imgui"

type ProgressBarWidget struct {
	BaseWidget
	fraction float32
	size     imgui.Vec2
	overlay  string
}

func ProgressBarV(fraction float32, width, height float32, overlay string) *ProgressBarWidget {
	return &ProgressBarWidget{
		BaseWidget: BaseWidget{width: 0},
		fraction:   fraction,
		size:       imgui.Vec2{X: width, Y: height},
		overlay:    overlay,
	}
}

func ProgressBar(fraction float32) *ProgressBarWidget {
	return ProgressBarV(fraction, -1, 0, "")
}

func (p *ProgressBarWidget) Build() {
	imgui.ProgressBarV(p.fraction, p.size, p.overlay)
}
