package giu

import (
	"image"

	"github.com/AllenDang/imgui-go"
)

type PlotWidget interface {
	Plot()
}

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
	plots                            []PlotWidget
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

func (p *PlotCanvasWidget) Plots(plots ...PlotWidget) *PlotCanvasWidget {
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

		if imgui.ImPlotBegin(tStr(p.title), tStr(p.xLabel), tStr(p.yLabel), ToVec2(image.Pt(p.width, p.height)), p.flags, p.xFlags, p.yFlags, p.y2Flags, p.y3Flags, tStr(p.y2Label), tStr(p.y3Label)) {
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
	imgui.ImPlotBarsH(tStr(p.title), p.data, p.height, p.shift, p.offset)
}

type PlotLineWidget struct {
	title      string
	values     []float64
	xScale, x0 float64
	offset     int
}

func PlotLine(title string, values []float64) *PlotLineWidget {
	return &PlotLineWidget{
		title:  title,
		values: values,
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
	imgui.ImPlotLine(tStr(p.title), p.values, p.xScale, p.x0, p.offset)
}

type PlotLineXYWidget struct {
	title  string
	xs, ys []float64
	offset int
}

func PlotLineXY(title string, xvalues, yvalues []float64) *PlotLineXYWidget {
	return &PlotLineXYWidget{
		title:  title,
		xs:     xvalues,
		ys:     yvalues,
		offset: 0,
	}
}

func (p *PlotLineXYWidget) Offset(offset int) *PlotLineXYWidget {
	p.offset = offset
	return p
}

func (p *PlotLineXYWidget) Plot() {
	imgui.ImPlotLineXY(tStr(p.title), p.xs, p.ys, p.offset)
}

type PlotPieChartWidget struct {
	labels       []string
	values       []float64
	x, y, radius float64
	normalize    bool
	labelFormat  string
	angle0       float64
}

func PlotPieChart(labels []string, values []float64, x, y, radius float64) *PlotPieChartWidget {
	return &PlotPieChartWidget{
		labels:      labels,
		values:      values,
		x:           x,
		y:           y,
		radius:      radius,
		normalize:   false,
		labelFormat: "%.1f",
		angle0:      90,
	}
}

func (p *PlotPieChartWidget) Normalize(n bool) *PlotPieChartWidget {
	p.normalize = n
	return p
}

func (p *PlotPieChartWidget) LabelFormat(fmtStr string) *PlotPieChartWidget {
	p.labelFormat = fmtStr
	return p
}

func (p *PlotPieChartWidget) Angle0(a float64) *PlotPieChartWidget {
	p.angle0 = a
	return p
}

func (p *PlotPieChartWidget) Plot() {
	for _, l := range p.labels {
		tStr(l)
	}
	imgui.ImPlotPieChart(p.labels, p.values, p.x, p.y, p.radius, p.normalize, p.labelFormat, p.angle0)
}

type PlotScatterWidget struct {
	label      string
	values     []float64
	xscale, x0 float64
	offset     int
}

func PlotScatter(label string, values []float64) *PlotScatterWidget {
	return &PlotScatterWidget{
		label:  label,
		values: values,
		xscale: 1,
		x0:     0,
		offset: 0,
	}
}

func (p *PlotScatterWidget) XScale(s float64) *PlotScatterWidget {
	p.xscale = s
	return p
}

func (p *PlotScatterWidget) X0(x float64) *PlotScatterWidget {
	p.x0 = x
	return p
}

func (p *PlotScatterWidget) Offset(offset int) *PlotScatterWidget {
	p.offset = offset
	return p
}

func (p *PlotScatterWidget) Plot() {
	imgui.ImPlotScatter(tStr(p.label), p.values, p.xscale, p.x0, p.offset)
}

type PlotScatterXYWidget struct {
	label  string
	xs, ys []float64
	offset int
}

func PlotScatterXY(label string, xs, ys []float64) *PlotScatterXYWidget {
	return &PlotScatterXYWidget{
		label:  label,
		xs:     xs,
		ys:     ys,
		offset: 0,
	}
}

func (p *PlotScatterXYWidget) Offset(offset int) *PlotScatterXYWidget {
	p.offset = offset
	return p
}

func (p *PlotScatterXYWidget) Plot() {
	imgui.ImPlotScatterXY(tStr(p.label), p.xs, p.ys, p.offset)
}
