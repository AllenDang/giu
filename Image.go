package giu

import "github.com/AllenDang/giu/imgui"

type ImageWidget struct {
	BaseWidget
	id   imgui.TextureID
	size imgui.Vec2
}

func Image(id imgui.TextureID, width, height float32) *ImageWidget {
	return &ImageWidget{
		BaseWidget: BaseWidget{width: 0},
		id:         id,
		size:       imgui.Vec2{X: width, Y: height},
	}
}

func (i *ImageWidget) Build() {
	if i.id != 0 {
		rect := imgui.ContentRegionAvail()
		if i.size.X == -1 {
			i.size.X = rect.X
		}
		if i.size.Y == -1 {
			i.size.Y = rect.Y
		}
		imgui.Image(i.id, i.size)
	}
}
