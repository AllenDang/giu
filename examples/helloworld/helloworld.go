package main

import (
	"fmt"
	"os"

	g "github.com/ianling/giu"
)

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}

func onQuit() {
	os.Exit(0)
}

func loop() {
	g.SingleWindow("hello world").Layout(
		g.Label("Hello world from giu"),
		g.Line(
			g.Button("Click Me").OnClick(onClickMe),
			g.Button("I'm so cute").OnClick(onImSoCute),
			g.Button("Quit").OnClick(onQuit),
		),
	).Build()
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}
