package giu

import "github.com/AllenDang/giu/imgui"

var (
	Context context
)

type context struct {
	renderer imgui.Renderer
	platform imgui.Platform
}

func (c context) GetRenderer() imgui.Renderer {
	return c.renderer
}

func (c context) GetPlatform() imgui.Platform {
	return c.platform
}

func (c context) IO() imgui.IO {
	return imgui.CurrentIO()
}
