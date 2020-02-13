package giu

import (
	"image"
	"image/color"
	"math"
	"time"
)

type ProgressIndicatorState struct {
	angle float64
	stop  bool
}

func (ps *ProgressIndicatorState) Update() {
	ticker := time.NewTicker(time.Second / 60)
	for !ps.stop {
		if ps.angle > 360 {
			ps.angle = 0
		}
		ps.angle += 0.1

		Update()
		<-ticker.C
	}

	ticker.Stop()
}

func (ps *ProgressIndicatorState) Dispose() {
	ps.stop = true
}

type ProgressIndicatorWidget struct {
	id     string
	radius float32
}

func ProgressIndicator(radius float32) *ProgressIndicatorWidget {
	return &ProgressIndicatorWidget{
		id:     "giu-progress-indicator",
		radius: radius * Context.GetPlatform().GetContentScale(),
	}
}

func (p *ProgressIndicatorWidget) Build() {
	// State exists
	if s := Context.GetState(p.id); s == nil {
		// Register state and start go routine
		ps := ProgressIndicatorState{angle: 0.0, stop: false}
		Context.SetState(p.id, &ps)
		go ps.Update()
	} else {
		state := s.(*ProgressIndicatorState)

		canvas := GetCanvas()

		pos := GetCursorScreenPos()

		centerPt := pos.Add(image.Pt(int(p.radius+p.radius/5), int(p.radius+p.radius/5)))
		centerPt2 := image.Pt(
			int(float64(p.radius)*math.Sin(state.angle)+float64(centerPt.X)),
			int(float64(p.radius)*math.Cos(state.angle)+float64(centerPt.Y)),
		)

		canvas.AddCircle(centerPt, float32(p.radius), color.RGBA{255, 255, 255, 255}, int(p.radius), float32(p.radius/20.0))
		canvas.AddCircleFilled(centerPt2, float32(p.radius/5), color.RGBA{255, 255, 255, 255}, int(p.radius/5))

		width := float32(p.radius + p.radius/5)
		InvisibleButton("pd", width, width, nil).Build()
	}
}
