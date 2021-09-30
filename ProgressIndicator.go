package giu

import (
	"image"
	"math"
	"time"

	"github.com/AllenDang/imgui-go"
)

var _ Disposable = &progressIndicatorState{}

type progressIndicatorState struct {
	angle float64
	stop  bool
}

func (ps *progressIndicatorState) update() {
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

// Dispose implements Disposable interface.
func (ps *progressIndicatorState) Dispose() {
	ps.stop = true
}

// static check to ensure if ProgressIndicatorWidget implements Widget interface.
var _ Widget = &ProgressIndicatorWidget{}

// ProgressIndicatorWidget represents progress indicator widget
// see examples/extrawidgets/.
type ProgressIndicatorWidget struct {
	internalID string
	width      float32
	height     float32
	radius     float32
	label      string
}

// ProgressIndicator creates a new ProgressIndicatorWidget.
func ProgressIndicator(label string, width, height, radius float32) *ProgressIndicatorWidget {
	return &ProgressIndicatorWidget{
		internalID: "###giu-progress-indicator",
		width:      width,
		height:     height,
		radius:     radius,
		label:      label,
	}
}

// Build implements Widget interface.
func (p *ProgressIndicatorWidget) Build() {
	// State exists
	if s := Context.GetState(p.internalID); s == nil {
		// Register state and start go routine
		ps := progressIndicatorState{angle: 0.0, stop: false}
		Context.SetState(p.internalID, &ps)
		go ps.update()
	} else {
		var isOk bool
		state, isOk := s.(*progressIndicatorState)
		Assert(isOk, "ProgressIndicatorWidget", "Build", "got unexpected type of widget's sate")

		child := Child().Border(false).Size(p.width, p.height).Layout(Layout{
			Custom(func() {
				// Process width and height
				width, height := GetAvailableRegion()

				canvas := GetCanvas()

				pos := GetCursorScreenPos()

				centerPt := pos.Add(image.Pt(int(width/2), int(height/2)))
				centerPt2 := image.Pt(
					int(float64(p.radius)*math.Sin(state.angle)+float64(centerPt.X)),
					int(float64(p.radius)*math.Cos(state.angle)+float64(centerPt.Y)),
				)

				color := imgui.CurrentStyle().GetColor(imgui.StyleColorText)
				rgba := Vec4ToRGBA(color)

				canvas.AddCircle(centerPt, p.radius, rgba, int(p.radius), p.radius/20.0)
				canvas.AddCircleFilled(centerPt2, p.radius/5, rgba)

				// Draw text
				if len(p.label) > 0 {
					labelWidth, _ := CalcTextSize(tStr(p.label))
					labelPos := centerPt.Add(image.Pt(-1*int(labelWidth/2), int(p.radius+p.radius/5+8)))
					canvas.AddText(labelPos, rgba, p.label)
				}
			}),
		})

		child.Build()
	}
}
