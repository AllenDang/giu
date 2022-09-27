package main

import (
	_ "embed"

	"github.com/AllenDang/giu"
)

//go:embed style.css
var cssStyle []byte

func loop() {
	giu.Window("Window").Layout(
		giu.CSSTag("button").To(
			giu.Button("HI! I'm a button styled with CSS"),
		),
		giu.CSSTag("label").To(
			giu.Label("I'ma  normal label"),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("CSS Style [example]", 640, 480, 0)
	if err := giu.ParseCSSStyleSheet(cssStyle); err != nil {
		panic(err)
	}
	wnd.Run(loop)
}
