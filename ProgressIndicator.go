package giu

import (
	"image"
	"math"
	"sync"
	"time"

	"github.com/AllenDang/imgui-go"
)

var _ Disposable = &progressIndicatorState{}

type progressIndicatorState struct {
	angle float64
	stop  bool
	m     *sync.Mutex
}

func (ps *progressIndicatorState) update() {
	ticker := time.NewTicker(time.Second / 60)

	for {
		ps.m.Lock()
		if ps.stop {
			ps.m.Unlock()
			break
		}

		if ps.angle > 6.2 {
			ps.angle = 0
		}

		ps.angle += 0.1

		ps.m.Unlock()

		Update()
		<-ticker.C
	}

	ticker.Stop()
}

// Dispose implements Disposable interface.
func (ps *progressIndicatorState) Dispose() {
	ps.m.Lock()
	ps.stop = true
	ps.m.Unlock()
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
	if state := GetState[progressIndicatorState](Context, p.internalID); state == nil {
		// Register state and start go routine
		ps := progressIndicatorState{
			angle: 0.0,
			stop:  false,
			m:     &sync.Mutex{},
		}

		SetState(Context, p.internalID, &ps)

		go ps.update()
	} else {
		child := Child().Border(false).Size(p.width, p.height).Layout(Layout{
			Custom(func() {
				// Process width and height
				width, height := GetAvailableRegion()

				canvas := GetCanvas()

				pos := GetCursorScreenPos()

				centerPt := pos.Add(image.Pt(int(width/2), int(height/2)))

				state.m.Lock()
				angle := state.angle
				state.m.Unlock()

				centerPt2 := image.Pt(
					int(float64(p.radius)*math.Sin(angle)+float64(centerPt.X)),
					int(float64(p.radius)*math.Cos(angle)+float64(centerPt.Y)),
				)

				color := imgui.CurrentStyle().GetColor(imgui.StyleColorText)
				rgba := Vec4ToRGBA(color)

				canvas.AddCircle(centerPt, p.radius, rgba, int(p.radius), p.radius/20.0)
				canvas.AddCircleFilled(centerPt2, p.radius/5, rgba)

				// Draw text
				if len(p.label) > 0 {
					labelWidth, _ := CalcTextSize(Context.FontAtlas.RegisterString(p.label))
					labelPos := centerPt.Add(image.Pt(-1*int(labelWidth/2), int(p.radius+p.radius/5+8)))
					canvas.AddText(labelPos, rgba, p.label)
				}
			}),
		})

		child.Build()
	}
}
