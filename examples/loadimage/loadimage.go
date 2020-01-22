package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"time"

	g "github.com/AllenDang/giu"
	resty "github.com/go-resty/resty/v2"
)

var (
	texture *g.Texture
	url     string
	client  *resty.Client
	loading bool
)

func loadImage(imageUrl string) {
	loading = true
	g.Update()

	resp, err := client.R().Get(imageUrl)
	if err == nil {
		img, _, err := image.Decode(bytes.NewReader(resp.Body()))
		if err == nil {
			rgba := image.NewRGBA(img.Bounds())
			draw.Draw(rgba, img.Bounds(), img, image.Point{}, draw.Src)
			texture, _ = g.NewTextureFromRgba(rgba)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

	loading = false
	g.Update()
}

func btnLoadClicked() {
	if len(url) > 0 {
		go loadImage(url)
	}
}

func loop() {
	g.SingleWindow("load image", g.Layout{
		g.Label("Paste image url and click download"),
		g.InputText("##Url", -1, &url),
		g.Button("Download and display", btnLoadClicked),
		g.Custom(func() {
			if loading {
				g.Label("Downloadig image ...").Build()
			} else {
				g.Image(texture, -1, -1).Build()
			}
		}),
	})
}

func main() {
	client = resty.New()
	client.SetTimeout(10 * time.Second)

	wnd := g.NewMasterWindow("Load Image", 600, 400, false, nil)
	wnd.Main(loop)
}
