package main

import (
	"bytes"
	_ "embed"
	"image/gif"
	"log"
	"time"

	"github.com/AllenDang/giu"
)

//go:embed golang.gif
var gifFileData []byte

const gifFilepath = "./golang.gif"

var (
	frames       []*giu.Texture
	gifImg       *gif.GIF
	currentFrame int
)

func loop() {
	// load textures
	if frames[0] == nil {
		for i, frame := range gifImg.Image {
			// lol, this is the most magic thing in go i've ever seen :D
			i := i
			giu.NewTextureFromRgba(giu.ImageToRgba(frame), func(t *giu.Texture) {
				frames[i] = t
			})
		}
	}

	giu.SingleWindow().Layout(
		giu.Image(frames[currentFrame]),
	)
}

func main() {
	var err error

	wnd := giu.NewMasterWindow("GIF renderer [example]", 640, 480, 0)

	gifImg, err = gif.DecodeAll(bytes.NewReader(gifFileData))
	if err != nil {
		log.Fatal(err)
	}

	frames = make([]*giu.Texture, len(gifImg.Image))

	go func() {
		for {
			time.Sleep(time.Duration(gifImg.Delay[currentFrame]*10) * time.Millisecond)
			giu.Update()
			currentFrame++
			if currentFrame == len(frames) {
				currentFrame = 0
			}
		}
	}()

	wnd.Run(loop)
}
