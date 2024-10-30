// Package main shows how to use giu CSS implementation to set style on your app.
// See also: examples/setstyle for native StyleSetter usage
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
		giu.Plot("styled plot").Plots(
			giu.Line("Plot 1", []float64{0, 1, 2, 3, 4, 5}),
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
