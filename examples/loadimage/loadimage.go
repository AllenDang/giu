package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"time"

	g "github.com/AllenDang/giu"
)

var (
	rgba *image.RGBA
	tex  *g.Texture
)

func loop() {
	g.SingleWindow().Layout(
		g.Label("Display image from texture"),
		g.Image(tex),
		g.Label("Display image from rgba"),
		g.ImageWithRgba(rgba).OnClick(func() {
			fmt.Println("rgba image was clicked")
		}).Size(200, 100),

		g.Label("Display image from file"),
		g.ImageWithFile("gopher.png").OnClick(func() {
			fmt.Println("image from file was clicked")
		}).Size(300, 200),

		g.Label("Display image from url (wait few seconds to download)"),
		g.ImageWithURL("https://png.pngitem.com/pimgs/s/3-36108_gopher-golang-hd-png-download.png").OnClick(func() {
			fmt.Println("image from url clicked")
		}).Size(300, 200),

		g.Label("Display images from url with loading and fallback"),
		g.ImageWithURL(
			"https://png.pngitem.com/pimgs/s/424-4241958_transparent-gopher-png-golang-gopher-png-png-download.png").
			Timeout(5*time.Second).
			Size(300, 200).
			LayoutForLoading(
				g.Child().Size(300, 200).Layout(g.Layout{
					g.Label("Loading..."),
				}),
			).
			LayoutForFailure(
				g.ImageWithFile("./fallback.png").Size(300, 200),
			).
			OnReady(func() {
				fmt.Println("Image is downloaded.")
			}),

		g.Label("Handle failure event"),
		g.ImageWithURL("http://x.y/z.jpg").Timeout(2*time.Second).OnFailure(func(err error) {
			fmt.Printf("Failed to download image, Error msg is %s\n", err.Error())
		}),

		g.Label("Display image from url without placeholder (no size when loading)"),
		g.ImageWithURL("https://www.pngitem.com/pimgs/m/424-4242405_go-lang-gopher-clipart-png-download-golang-gopher.png").Size(300, 200),

		g.Label("Footer"),
	)
}

func main() {
	rgba, _ = g.LoadImage("./fallback.png")

	wnd := g.NewMasterWindow("Load Image", 600, 500, g.MasterWindowFlagsNotResizable)
	g.EnqueueNewTextureFromRgba(rgba, func(t *g.Texture) {
		tex = t
	})
	wnd.SetIcon(rgba)
	wnd.Run(loop)
}
