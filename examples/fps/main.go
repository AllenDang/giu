package main

import (
	"time"

	"github.com/AllenDang/giu"
)

var (
	currentTime time.Time
	timeDelta   time.Duration

	fpsTime    time.Time
	frames     int
	currentFPS int
)

func loop() {
	giu.SingleWindow().Layout(
		giu.Custom(func() {
			frames++
			timeDelta = time.Now().Sub(currentTime)
			currentTime = time.Now()
			fpsTime = fpsTime.Add(timeDelta)
			if fpsTime.Second() >= 1 {
				currentFPS = frames
				frames = 0
				fpsTime = time.Time{}
			}
		}),
		giu.Labelf("Current time delta %v", timeDelta),
		giu.Labelf("Current FPS: %d", currentFPS),
	)
}

func main() {
	wnd := giu.NewMasterWindow("FPS calculation [example]", 640, 480, 0)

	// make sure if amount of calls giu.Update doesn't affect max FPS (60)
	go func() {
		for range time.Tick(time.Millisecond) {
			giu.Update()
		}
	}()

	wnd.Run(loop)
}
