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

		<-ticker.C
	}
}

func loop(w *giu.MasterWindow) {
	giu.SingleWindow(w, "Update", giu.Layout{
		giu.Custom(func() {
			giu.Update(time.Second)
		}),
		giu.Label(fmt.Sprintf("%d", counter)),
	})
}

func main() {
	wnd := giu.NewMasterWindow("Update", 400, 200, false, nil)

	go refersh()

	wnd.Main(loop)
}
