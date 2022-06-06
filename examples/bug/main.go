package main

import (
	"fmt"

	"github.com/AllenDang/giu"
)

const popupID = "somepopupmodal"

func loop() {
	items := make([]string, 200)
	for i := range items {
		items[i] = fmt.Sprintf("Item %d", i)
	}

	giu.SingleWindow().Layout(
		giu.Row(
			giu.ListBox("listbox", items).Size(300, 400),
			giu.ListBox("listbox2", items).Size(300, 400),
		),
		giu.Popup(popupID).Layout(
			giu.Row(
				giu.Button("Do something").OnClick(func() { fmt.Println("Something done") }),
				giu.Button("close me").OnClick(func() { giu.CloseCurrentPopup() }),
			),
		),
		giu.Button("Open popup").OnClick(func() { giu.OpenPopup(popupID) }),
	)
}

func main() {
	wnd := giu.NewMasterWindow("#390 [bug]", 640, 480, 0)
	wnd.Run(loop)
}
