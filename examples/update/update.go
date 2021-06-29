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

func refresh() {
	ticker := time.NewTicker(time.Second * 1)

	for {
		counter = rand.Intn(100)
		giu.Update()

		<-ticker.C
	}
}

func loop() {
	giu.SingleWindow().Layout(
		giu.Label("Below number is updated by a goroutine"),
		giu.Label(fmt.Sprintf("%d", counter)),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Update", 400, 200, giu.MasterWindowFlagsNotResizable)

	go refresh()

	wnd.Run(loop)
}
