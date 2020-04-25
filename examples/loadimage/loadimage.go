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
		g.ImageWithUrl("https://www.pngitem.com/pimgs/m/424-4242405_go-lang-gopher-clipart-png-download-golang-gopher.png", 10*time.Second, 300, 200),
	})
}

func main() {
	wnd := g.NewMasterWindow("Load Image", 600, 500, g.MasterWindowFlagsNotResizable, nil)
	wnd.Main(loop)
}
