package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}

func loop(w *g.MasterWindow) {
	g.SingleWindow(w, "hello world", g.Layout{
		g.Label("Hello world from giu"),
		g.Line(
			g.Button("Click Me", onClickMe),
			g.Button("I'm so cute", onImSoCute)),
	})
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, false, nil)
	wnd.Main(loop)
}
