package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/AllenDang/cimgui-go/imgui"
	g "github.com/AllenDang/giu"
)

var (
	fromrgba     = &g.ReflectiveBoundTexture{}
	fromfile     = &g.ReflectiveBoundTexture{}
	fromurl      = &g.ReflectiveBoundTexture{}
	rgba         *image.RGBA
	sonicOffsetX = int32(1180)
	sonicOffsetY = int32(580)
)

func loop() {
	var start_pos image.Point

	g.SingleWindow().Layout(
		g.Custom(func() {
			start_pos = g.GetCursorScreenPos()
		}),
		g.Label("Display wich has size of contentAvaiable (stretch)"),
		fromfile.ToImageWidget().OnClick(func() {
			fmt.Println("contentAvailable image was clicked")
		}).Size(-1, -1),

		g.Label("Display image from preloaded rgba"),
		fromrgba.ToImageWidget().OnClick(func() {
			fmt.Println("rgba image was clicked")
		}),

		g.Label("Display image from file"),
		fromfile.ToImageWidget().OnClick(func() {
			fmt.Println("image from file was clicked")
		}),

		g.Label("Display image from url + 0.25 scale"),
		fromurl.ToImageWidget().OnClick(func() {
			fmt.Println("image from url clicked")
		}).Scale(0.25, 0.25),

		g.Separator(),
		g.Label("Advanced Drawing manipulation"),
		g.DragInt("Sonic Offset X", &sonicOffsetX, 0, 1280),
		g.DragInt("Sonic Offset Y", &sonicOffsetY, 0, 720),
		g.Custom(func() {
			size := fromurl.GetSurfaceSize()
			sonicOffset := image.Point{int(sonicOffsetX), int(sonicOffsetY)}
			pos_with_offset := start_pos.Add(sonicOffset)
			computed_posX := (float32(pos_with_offset.X)) + imgui.ScrollX()
			computed_posY := (float32(pos_with_offset.Y)) + imgui.ScrollY()
			scale := imgui.Vec2{X: 0.10, Y: 0.10}
			p_min := imgui.Vec2{X: computed_posX, Y: computed_posY}
			p_max := imgui.Vec2{X: computed_posX + float32(size.X)*scale.X, Y: computed_posY + float32(size.Y)*scale.Y}
			imgui.ForegroundDrawListViewportPtr().AddImage(fromurl.Texture().ID(), p_min, p_max)
		}),
		g.Separator(),
		g.Label("For more advanced image examples (async/statefull/dynamic) check the asyncimage example!"),
	)
}

func main() {
	var err error

	rgba, err = g.LoadImage("./fallback.png")
	if err != nil {
		log.Fatalf("Cannot loadIamge fallback.png")
	}

	fromfile.SetSurfaceFromFile("gopher.png", false)
	fromrgba.SetSurfaceFromRGBA(rgba, false)
	fromurl.SetSurfaceFromURL("https://static.wikia.nocookie.net/smashbros/images/0/0e/Art_Sonic_TSR.png/revision/latest?cb=20200210122913&path-prefix=fr", time.Second*5, false)

	wnd := g.NewMasterWindow("Load Image", 1280, 720, 0)
	wnd.Run(loop)
}
