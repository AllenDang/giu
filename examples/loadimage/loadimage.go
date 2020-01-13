package main

import (
	"bytes"
	"image"
	"image/draw"
	_ "image/png"
	"time"

	g "github.com/AllenDang/giu"
	resty "github.com/go-resty/resty/v2"
)

var (
	texture *g.Texture
	url     string
	client  *resty.Client
)

func loadImage(imageUrl string) {
	resp, err := client.R().Get(imageUrl)
	if err == nil {
		img, _, err := image.Decode(bytes.NewReader(resp.Body()))
		if err == nil {
			rgba := image.NewRGBA(img.Bounds())
			draw.Draw(rgba, img.Bounds(), img, image.Point{}, draw.Src)
			texture, _ = g.NewTextureFromRgba(rgba)
		}
	}
}

func btnLoadClicked() {
	go loadImage(url)
}

func loop(w *g.MasterWindow) {
	g.SingleWindow(w, "load image",
		g.InputText("Url", &url),
		g.Button("btnLoad", btnLoadClicked),
		g.Image(texture, -1, -1),
	)
}

func main() {
	client = resty.New()
	client.SetTimeout(10 * time.Second)

	wnd := g.NewMasterWindow("Load Image", 600, 400, false, nil)
	wnd.Main(loop)
}
