package main

import (
	"time"

	"github.com/AllenDang/giu"
)

var (
	text string
	date time.Time
)

func loop() {
	giu.Window("window").Layout(
		giu.Align(giu.AlignCenter).To(
			giu.Label("I'm a centered label"),
			giu.Button("I'm a centered button"),
		),

		giu.Align(giu.AlignRight).To(
			giu.Label("I'm an alined to right label"),
			giu.InputText(&text),
		),

		giu.Align(giu.AlignRight).To(
			giu.Label("I'm the label"),
			giu.Layout{
				giu.Label("I'm the other label embedded in another layout"),
				giu.Label("I'm the next label"),
			},
			giu.Label("I'm the last label"),
		),
		giu.Label("Buttons in row:"),
		giu.Align(giu.AlignCenter).To(
			giu.Row(
				giu.Button("button 1"),
				giu.Button("button 2"),
			),
		),

		giu.Label("manual alignment"),
		giu.AlignManually(
			giu.AlignCenter,
			giu.Button("I'm button with 200 width").
				Size(200, 30),
			200, false,
		),
		giu.AlignManually(
			giu.AlignCenter,
			giu.InputText(&text),
			100, true,
		),

		giu.Align(giu.AlignCenter).To(
			giu.DatePicker("<- date picker centered", &date),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Alignment demo", 640, 480, 0)
	wnd.Run(loop)
}
