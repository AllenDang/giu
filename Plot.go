package giu

import (
	"image"

	"github.com/AllenDang/giu/imgui"
)

type PlotWidget interface {
	Plot()
}

type Plots []PlotWidget

type ImPlotYAxis int

const (
	ImPlotYAxisLeft          ImPlotYAxis = 0 // left (default)
	ImPlotYAxisFirstOnRight  ImPlotYAxis = 1 // first on right side
	ImPlotYAxisSecondOnRight ImPlotYAxis = 2 // second on right side
)

type PlotTicker struct {
	Position float64
	Label    string
}

type PlotCanvasWidget struct {
	title                            string
	xLabel                           string
	yLabel                           string
	width                            int
	height                           int
	flags                            imgui.ImPlotFlags
	xFlags, yFlags, y2Flags, y3Flags imgui.ImPlotAxisFlags
	y2Label                          string
	y3Label                          string
	xMin, xMax, yMin, yMax           float64
	axisLimitCondition               ExecCondition
	xTicksValue, yTicksValue         []float64
	xTicksLabel, yTicksLabel         []string
	xTicksShowDefault                bool
	yTicksShowDefault                bool
	yTicksYAxis                      ImPlotYAxis
	plots                            Plots
}

func Plot(title string) *PlotCanvasWidget {
	return &PlotCanvasWidget{
		title:              title,
		xLabel:             "",
		yLabel:             "",
		width:              -1,
		height:             0,
		flags:              imgui.ImPlotFlags_None,
		xFlags:             imgui.ImPlotAxisFlags_None,
		yFlags:             imgui.ImPlotAxisFlags_None,
		y2Flags:            imgui.ImPlotAxisFlags_NoGridLines,
		y3Flags:            imgui.ImPlotAxisFlags_NoGridLines,
		y2Label:            "",
		y3Label:            "",
		xMin:               0,
		xMax:               10,
		yMin:               0,
		yMax:               10,
		xTicksShowDefault:  true,
		yTicksShowDefault:  true,
		yTicksYAxis:        0,
		axisLimitCondition: ConditionOnce,
	}
}

func (p *PlotCanvasWidget) AxisLimits(xmin, xmax, ymin, ymax float64, cond ExecCondition) *PlotCanvasWidget {
	p.xMin = xmin
	p.xMax = xmax
	p.yMin = ymin
	p.yMax = ymax
	p.axisLimitCondition = cond

	return p
}

func (p *PlotCanvasWidget) XTicks(ticks []PlotTicker, showDefault bool) *PlotCanvasWidget {
	length := len(ticks)
	if length == 0 {
		return p
	}

	values := make([]float64, length)
	labels := make([]string, length)

	for i, t := range ticks {
		values[i] = t.Position
		labels[i] = t.Label
	}

	p.xTicksValue = values
	p.xTicksLabel = labels
	p.xTicksShowDefault = showDefault
	return p
}

func (p *PlotCanvasWidget) YTicks(ticks []PlotTicker, showDefault bool, yAxis ImPlotYAxis) *PlotCanvasWidget {
	length := len(ticks)
	if length == 0 {
		return p
	}

	values := make([]float64, length)
	labels := make([]string, length)

	for i, t := range ticks {
		values[i] = t.Position
		labels[i] = t.Label
	}

	p.yTicksValue = values
	p.yTicksLabel = labels
	p.yTicksShowDefault = showDefault
	p.yTicksYAxis = yAxis
	return p
}

func (p *PlotCanvasWidget) Flags(flags imgui.ImPlotFlags) *PlotCanvasWidget {
	p.flags = flags
	return p
}

func (p *PlotCanvasWidget) XAxeFlags(flags imgui.ImPlotAxisFlags) *PlotCanvasWidget {
	p.xFlags = flags
	return p
}

func (p *PlotCanvasWidget) YAxeFlags(yFlags, y2Flags, y3Flags imgui.ImPlotAxisFlags) *PlotCanvasWidget {
	p.yFlags = yFlags
	p.y2Flags = y2Flags
	p.y3Flags = y3Flags
	return p
}

func (p *PlotCanvasWidget) Plots(plots Plots) *PlotCanvasWidget {
	p.plots = plots
	return p
}

func (p *PlotCanvasWidget) Size(width, height int) *PlotCanvasWidget {
	p.width = width
	p.height = height
	return p
}

func (p *PlotCanvasWidget) Build() {
	if len(p.plots) > 0 {
		imgui.ImPlotSetNextPlotLimits(p.xMin, p.xMax, p.yMin, p.yMax, imgui.Condition(p.axisLimitCondition))

		if len(p.xTicksValue) > 0 {
			imgui.ImPlotSetNextPlotTicksX(p.xTicksValue, p.xTicksLabel, p.xTicksShowDefault)
		}

		if len(p.yTicksValue) > 0 {
			imgui.ImPlotSetNextPlotTicksY(p.yTicksValue, p.yTicksLabel, p.yTicksShowDefault, int(p.yTicksYAxis))
		}

		if imgui.ImPlotBegin(p.title, p.xLabel, p.yLabel, ToVec2(image.Pt(p.width, p.height)), p.flags, p.xFlags, p.yFlags, p.y2Flags, p.y3Flags, p.y2Label, p.y3Label) {
			for _, plot := range p.plots {
				plot.Plot()
			}
			imgui.ImPlotEnd()
		}
	}
}

type PlotBarWidget struct {
	title  string
	data   []float64
	width  float64
	shift  float64
	offset int
}

func PlotBar(title string, data []float64) *PlotBarWidget {
	return &PlotBarWidget{
		title:  title,
		data:   data,
		width:  0.2,
		shift:  0,
		offset: 0,
	}
}

func (p *PlotBarWidget) Width(width float64) *PlotBarWidget {
	p.width = width
	return p
}

func (p *PlotBarWidget) Shift(shift float64) *PlotBarWidget {
	p.shift = shift
	return p
}

func (p *PlotBarWidget) Offset(offset int) *PlotBarWidget {
	p.offset = offset
	return p
}

func (p *PlotBarWidget) Plot() {
	imgui.ImPlotBars(p.title, p.data, p.width, p.shift, p.offset)
}

type PlotBarHWidget struct {
	title  string
	data   []float64
	height float64
	shift  float64
	offset int
}

func PlotBarH(title string, data []float64) *PlotBarHWidget {
	return &PlotBarHWidget{
		title:  title,
		data:   data,
		height: 0.2,
		shift:  0,
		offset: 0,
	}
}

func (p *PlotBarHWidget) Height(height float64) *PlotBarHWidget {
	p.height = height
	return p
}

func (p *PlotBarHWidget) Shift(shift float64) *PlotBarHWidget {
	p.shift = shift
	return p
}

func (p *PlotBarHWidget) Offset(offset int) *PlotBarHWidget {
	p.offset = offset
	return p
}

func (p *PlotBarHWidget) Plot() {
	imgui.ImPlotBarsH(p.title, p.data, p.height, p.shift, p.offset)
}

type PlotLineWidget struct {
	title      string
	data       []float64
	xScale, x0 float64
	offset     int
}

func PlotLine(title string, data []float64) *PlotLineWidget {
	return &PlotLineWidget{
		title:  title,
		data:   data,
		xScale: 1,
		x0:     0,
		offset: 0,
	}
}

func (p *PlotLineWidget) XScale(scale float64) *PlotLineWidget {
	p.xScale = scale
	return p
}

func (p *PlotLineWidget) X0(x0 float64) *PlotLineWidget {
	p.x0 = x0
	return p
}

func (p *PlotLineWidget) Offset(offset int) *PlotLineWidget {
	p.offset = offset
	return p
}

func (p *PlotLineWidget) Plot() {
	imgui.ImPlotLine(p.title, p.data, p.xScale, p.x0, p.offset)
}
