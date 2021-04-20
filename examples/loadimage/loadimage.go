package main

import (
	_ "image/jpeg"
	_ "image/png"
	"time"

	g "github.com/ianling/giu"
)

func loop() {
	g.SingleWindow("load image").Layout(
		g.Label("Display image from file"),
		g.ImageWithFile("gopher.png").Size(300, 200),

		g.Label("Display image from url (wait few seconds to download)"),
		g.ImageWithUrl("https://png.pngitem.com/pimgs/s/3-36108_gopher-golang-hd-png-download.png").Size(300, 200),

		g.Label("Display images from url with loading and fallback"),
		g.ImageWithUrl(
			"https://png.pngitem.com/pimgs/s/424-4241958_transparent-gopher-png-golang-gopher-png-png-download.png").
			Timeout(5*time.Second).
			Size(300, 200).
			LayoutForLoading(
				g.Child("Loading").Size(300, 200).Layout(g.Layout{
					g.Label("Loading..."),
				}),
			).
			LayoutForFailure(
				g.ImageWithFile("./fallback.png").Size(300, 200),
			),

		g.Label("Display image from url without placeholder (no size when loading)"),
		g.ImageWithUrl("https://www.pngitem.com/pimgs/m/424-4242405_go-lang-gopher-clipart-png-download-golang-gopher.png").Size(300, 200),

		g.Label("Footer"),
	)
}

func main() {
	wnd := g.NewMasterWindow("Load Image", 600, 500, g.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}
