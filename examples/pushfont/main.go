package main

import (
	"fmt"

	"github.com/AllenDang/giu"
)

var font *giu.FontInfo

func loop() {
	fontPushed := false
	giu.Window("example").Layout(
		giu.Custom(func() {
			fontPushed = giu.PushFont(font)
		}),
		giu.Label("x"),
		giu.Custom(func() {
			if fontPushed {
				giu.PopFont()
			}
		}),
	)
}

func main() {
	wnd := giu.NewMasterWindow("example", 640, 480, 0)
	font = giu.Context.FontAtlas.AddFont("NotoSansSinhala-Regular.ttf", 20)
	fmt.Println(font)
	wnd.Run(loop)
}
