package giu

import (
	"fmt"
	"image"
	"math"
	"time"

	"github.com/AllenDang/imgui-go"
)

type ProgressIndicatorState struct {
	angle float64
	stop  bool
}

func (ps *ProgressIndicatorState) Update() {
	ticker := time.NewTicker(time.Second / 60)
	for !ps.stop {
		if ps.angle > 6.2 {
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
	id         string
	internalId string
	width      float32
	height     float32
	radius     float32
	label      string
}

func ProgressIndicator(id, label string, width, height, radius float32) *ProgressIndicatorWidget {
	return &ProgressIndicatorWidget{
		id:         id,
		internalId: "###giu-progress-indicator",
		width:      width * Context.GetPlatform().GetContentScale(),
		height:     height * Context.GetPlatform().GetContentScale(),
		radius:     radius * Context.GetPlatform().GetContentScale(),
		label:      label,
	}
}

func (p *ProgressIndicatorWidget) Build() {
	// State exists
	if s := Context.GetState(p.internalId); s == nil {
		// Register state and start go routine
		ps := ProgressIndicatorState{angle: 0.0, stop: false}
		Context.SetState(p.internalId, &ps)
		go ps.Update()
	} else {
		state := s.(*ProgressIndicatorState)

		child := Child(fmt.Sprintf("%s-container", p.id)).Border(false).Size(p.width, p.height).Layout(Layout{
			Custom(func() {
				// Process width and height
				width, height := GetAvaiableRegion()

				canvas := GetCanvas()

				pos := GetCursorScreenPos()

				centerPt := pos.Add(image.Pt(int(width/2), int(height/2)))
				centerPt2 := image.Pt(
					int(float64(p.radius)*math.Sin(state.angle)+float64(centerPt.X)),
					int(float64(p.radius)*math.Cos(state.angle)+float64(centerPt.Y)),
				)

				color := imgui.CurrentStyle().GetColor(imgui.StyleColorText)
				rgba := Vec4ToRGBA(color)

				canvas.AddCircle(centerPt, float32(p.radius), rgba, float32(p.radius/20.0))
				canvas.AddCircleFilled(centerPt2, float32(p.radius/5), rgba)

				// Draw text
				if len(p.label) > 0 {
					labelWidth, _ := CalcTextSize(p.label)
					labelPos := centerPt.Add(image.Pt(-1*int(labelWidth/2), int(p.radius+p.radius/5+8)))
					canvas.AddText(labelPos, rgba, p.label)
				}
			}),
		})

		child.Build()
	}
}
