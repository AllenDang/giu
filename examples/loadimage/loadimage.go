package main

import (
	_ "image/jpeg"
	_ "image/png"
	"time"

	g "github.com/AllenDang/giu"
)

func loop() {
	g.SingleWindow("load image", g.Layout{
		g.Label("Display image from file"),
		g.ImageWithFile("gopher.png", 300, 200),

		g.Label("Display image from url (wait few seconds to download)"),
		g.ImageWithUrl("https://png.pngitem.com/pimgs/s/3-36108_gopher-golang-hd-png-download.png", 10*time.Second, 300, 200),

		g.Label("Display images from url with loading and fallback"),
		g.ImageWithUrlV("https://png.pngitem.com/pimgs/s/424-4241958_transparent-gopher-png-golang-gopher-png-png-download.png",
			5*time.Second, 300, 200,
			g.Layout{
				g.Child("Loading", true, 300, 200, 0, g.Layout{
					g.Label("Loading..."),
				}),
			},
			g.Layout{
				g.ImageWithFile("./fallback.png", 300, 200),
			},
		),

		g.Label("Display image from url without placeholder (no size when loading)"),
		g.ImageWithUrlV("https://www.pngitem.com/pimgs/m/424-4242405_go-lang-gopher-clipart-png-download-golang-gopher.png",
			10*time.Second, 300, 200,
			nil,
			nil,
		),

		g.Label("Footer"),
	})
}

func main() {
	wnd := g.NewMasterWindow("Load Image", 600, 500, g.MasterWindowFlagsNotResizable, nil)
	wnd.Main(loop)
}
