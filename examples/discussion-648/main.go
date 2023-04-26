package main

import "github.com/AllenDang/giu"

func loop() {
	giu.Window("wnd").Layout(
		giu.Custom(func() {
			const footerPercentage = 0.2
			_, availableH := giu.GetAvailableRegion()
			_, itemSpacingH := giu.GetItemSpacing()
			giu.Layout{
				giu.Child().Layout(giu.Label("your layout")).Size(-1, (availableH-itemSpacingH)*(1-footerPercentage)),
				giu.Child().Layout(giu.Label("footer")).Size(-1, (availableH-itemSpacingH)*footerPercentage),
			}.Build()
		}),
	)
}

func main() {
	wnd := giu.NewMasterWindow("How do I make this Child show as a footer? #648", 640, 480, 0)
	wnd.Run(loop)
}
