// Package main provides examples of loading and presenting images in advanced ways
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
	var startPos image.Point

	g.SingleWindow().Layout(
		g.Custom(func() {
			startPos = g.GetCursorScreenPos()
		}),
		g.Label("Display which has size of contentAvaiable (stretch)"),
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
		g.DragInt(&sonicOffsetX).Label("Sonic Offset X").MinValue(0).MaxValue(1280),
		g.DragInt(&sonicOffsetY).Label("Sonic Offset Y").MinValue(0).MaxValue(720),
		g.Custom(func() {
			size := fromurl.GetSurfaceSize()
			sonicOffset := image.Point{int(sonicOffsetX), int(sonicOffsetY)}
			posWithOffset := startPos.Add(sonicOffset)
			computedPosX := (float32(posWithOffset.X)) + imgui.ScrollX()
			computedPosY := (float32(posWithOffset.Y)) + imgui.ScrollY()
			scale := imgui.Vec2{X: 0.10, Y: 0.10}
			pMin := imgui.Vec2{X: computedPosX, Y: computedPosY}
			pMax := imgui.Vec2{X: computedPosX + float32(size.X)*scale.X, Y: computedPosY + float32(size.Y)*scale.Y}
			imgui.ForegroundDrawListViewportPtr().AddImage(fromurl.Texture().ID(), pMin, pMax)
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

	_ = fromfile.SetSurfaceFromFile("gopher.png", false)
	_ = fromrgba.SetSurfaceFromRGBA(rgba, false)
	_ = fromurl.SetSurfaceFromURL("https://static.wikia.nocookie.net/smashbros/images/0/0e/Art_Sonic_TSR.png/revision/latest?cb=20200210122913&path-prefix=fr", time.Second*5, false)

	wnd := g.NewMasterWindow("Load Image", 1280, 720, 0)
	wnd.Run(loop)
}
