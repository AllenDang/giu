package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AllenDang/giu"
)

var (
	counter int
)

func refersh() {
	ticker := time.NewTicker(time.Second * 1)

	for {
		counter = rand.Intn(100)
		giu.Update()

		<-ticker.C
	}
}

func loop(w *giu.MasterWindow) {
	giu.SingleWindow(w, "Update", giu.Layout{
		giu.Label(fmt.Sprintf("%d", counter)),
		giu.Line(giu.Button("Button 1", nil), giu.Button("Button 2", nil)),
		giu.Table("table", true, giu.Rows{
			giu.Row(giu.Label("Column1"), giu.Label("Column2")),
			giu.Row(giu.Label("Column1"), giu.Label("Column2")),
			giu.Row(giu.Label("Column1"), giu.Label("Column2")),
			giu.Row(giu.Label("Column1"), giu.Label("Column2")),
			giu.Row(giu.Label("Column1"), giu.Label("Column2")),
		}),
	})
}

func main() {
	wnd := giu.NewMasterWindow("Update", 400, 200, false, nil)

	go refersh()

	wnd.Main(loop)
}
