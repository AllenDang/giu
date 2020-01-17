package imgui_test

import (
	"testing"

	"github.com/AllenDang/giu/imgui"

	"github.com/stretchr/testify/assert"
)

func TestDrawData(t *testing.T) {
	context := imgui.CreateContext(nil)
	defer context.Destroy()

	io := imgui.CurrentIO()
	io.SetDisplaySize(imgui.Vec2{X: 800, Y: 600})
	io.Fonts().TextureDataAlpha8()

	for i := 0; i < 2; i++ {
		imgui.NewFrame()
		imgui.ShowDemoWindow(nil)
		imgui.Render()
	}
	drawData := imgui.RenderedDrawData()

	assert.True(t, drawData.Valid(), "Draw data should be valid")
	list := drawData.CommandLists()
	assert.True(t, len(list) > 0, "At least one draw data list expected")
}
