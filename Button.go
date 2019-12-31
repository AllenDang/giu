package giu

import "github.com/AllenDang/giu/imgui"

type ButtonWidget struct {
	BaseWidget
	caption string
	size    imgui.Vec2
	clicked func()
}

func ButtonV(caption string, width, height float32, clicked func()) *ButtonWidget {
	return &ButtonWidget{
		BaseWidget: BaseWidget{width: 0},
		caption:    caption,
		clicked:    clicked,
		size:       imgui.Vec2{X: width, Y: height},
	}
}

func Button(caption string, clicked func()) *ButtonWidget {
	return ButtonV(caption, 0, 0, clicked)
}

func (l *ButtonWidget) Build() {
	if imgui.ButtonV(l.caption, l.size) && l.clicked != nil {
		l.clicked()
	}
}

type ImageButtonWidget struct {
	BaseWidget
	id      imgui.TextureID
	size    imgui.Vec2
	clicked func()
}

func ImageButton(id imgui.TextureID, width, height float32, clicked func()) *ImageButtonWidget {
	return &ImageButtonWidget{
		BaseWidget: BaseWidget{width: 0},
		id:         id,
		size:       imgui.Vec2{X: width, Y: height},
		clicked:    clicked,
	}
}

func (ib *ImageButtonWidget) Build() {
	if imgui.ImageButton(ib.id, ib.size) && ib.clicked != nil {
		ib.clicked()
	}
}
