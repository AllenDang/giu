package giu

import (
	"image"

	"github.com/AllenDang/imgui-go"
)

// PlotWidget is implemented by all the particular plots, which can be used
// in (*PlotCanvasWidget).Plots.
type PlotWidget interface {
	Plot()
}

// ImPlotYAxis represents y axis settings.
type ImPlotYAxis int

// ImPlotYAxis enum:.
const (
	ImPlotYAxisLeft          ImPlotYAxis = 0 // left (default)
	ImPlotYAxisFirstOnRight  ImPlotYAxis = 1 // first on right side
	ImPlotYAxisSecondOnRight ImPlotYAxis = 2 // second on right side
)

// PlotTicker represents axis ticks.
type PlotTicker struct {
	Position float64
	Label    string
}

// PlotCanvasWidget represents a giu plot widget.
type PlotCanvasWidget struct {
	title                            string
	xLabel                           string
	yLabel                           string
	width                            int
	height                           int
	flags                            PlotFlags
	xFlags, yFlags, y2Flags, y3Flags PlotAxisFlags
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

// Plot adds creates a new plot widget.
func Plot(title string) *PlotCanvasWidget {
	return &PlotCanvasWidget{
		title:              title,
		xLabel:             "",
		yLabel:             "",
		width:              -1,
		height:             0,
		flags:              PlotFlagsNone,
		xFlags:             PlotAxisFlagsNone,
		yFlags:             PlotAxisFlagsNone,
		y2Flags:            PlotAxisFlagsNoGridLines,
		y3Flags:            PlotAxisFlagsNoGridLines,
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

// AxisLimits sets X and Y axis limits.
func (p *PlotCanvasWidget) AxisLimits(xmin, xmax, ymin, ymax float64, cond ExecCondition) *PlotCanvasWidget {
	p.xMin = xmin
	p.xMax = xmax
	p.yMin = ymin
	p.yMax = ymax
	p.axisLimitCondition = cond

	return p
}

// XTicks sets x axis ticks.
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

// YTicks sets y axis ticks.
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

// Flags sets plot canvas flags.
func (p *PlotCanvasWidget) Flags(flags PlotFlags) *PlotCanvasWidget {
	p.flags = flags
	return p
}

// XAxeFlags sets x axis fags.
func (p *PlotCanvasWidget) XAxeFlags(flags PlotAxisFlags) *PlotCanvasWidget {
	p.xFlags = flags
	return p
}

// YAxeFlags sets y axis flags.
func (p *PlotCanvasWidget) YAxeFlags(yFlags, y2Flags, y3Flags PlotAxisFlags) *PlotCanvasWidget {
	p.yFlags = yFlags
	p.y2Flags = y2Flags
	p.y3Flags = y3Flags

	return p
}

// Plots adds plots to plot canvas.
func (p *PlotCanvasWidget) Plots(plots ...PlotWidget) *PlotCanvasWidget {
	p.plots = plots
	return p
}

// Size set canvas size.
func (p *PlotCanvasWidget) Size(width, height int) *PlotCanvasWidget {
	p.width = width
	p.height = height

	return p
}

// Build implements Widget interface.
func (p *PlotCanvasWidget) Build() {
	if len(p.plots) > 0 {
		imgui.ImPlotSetNextPlotLimits(p.xMin, p.xMax, p.yMin, p.yMax, imgui.Condition(p.axisLimitCondition))

		if len(p.xTicksValue) > 0 {
			imgui.ImPlotSetNextPlotTicksX(p.xTicksValue, p.xTicksLabel, p.xTicksShowDefault)
		}

		if len(p.yTicksValue) > 0 {
			imgui.ImPlotSetNextPlotTicksY(p.yTicksValue, p.yTicksLabel, p.yTicksShowDefault, int(p.yTicksYAxis))
		}

		if imgui.ImPlotBegin(
			Context.FontAtlas.RegisterString(p.title), Context.FontAtlas.RegisterString(p.xLabel),
			Context.FontAtlas.RegisterString(p.yLabel), ToVec2(image.Pt(p.width, p.height)),
			imgui.ImPlotFlags(p.flags), imgui.ImPlotAxisFlags(p.xFlags),
			imgui.ImPlotAxisFlags(p.yFlags), imgui.ImPlotAxisFlags(p.y2Flags),
			imgui.ImPlotAxisFlags(p.y3Flags), Context.FontAtlas.RegisterString(p.y2Label), Context.FontAtlas.RegisterString(p.y3Label),
		) {
			for _, plot := range p.plots {
				plot.Plot()
			}

			imgui.ImPlotEnd()
		}
	}
}

// BarPlot adds bar plot (column chart) to the canvas.
type BarPlot struct {
	title  string
	data   []float64
	width  float64
	shift  float64
	offset int
}

// Bar adds a plot bar (column chart).
func Bar(title string, data []float64) *BarPlot {
	return &BarPlot{
		title:  title,
		data:   data,
		width:  0.2,
		shift:  0,
		offset: 0,
	}
}

// Width sets bar width.
func (p *BarPlot) Width(width float64) *BarPlot {
	p.width = width
	return p
}

// Shift sets shift of the bar.
func (p *BarPlot) Shift(shift float64) *BarPlot {
	p.shift = shift
	return p
}

// Offset sets bar's offset.
func (p *BarPlot) Offset(offset int) *BarPlot {
	p.offset = offset
	return p
}

// Plot implements Plot interface.
func (p *BarPlot) Plot() {
	imgui.ImPlotBars(p.title, p.data, p.width, p.shift, p.offset)
}

// BarHPlot represents a column chart on Y axis.
type BarHPlot struct {
	title  string
	data   []float64
	height float64
	shift  float64
	offset int
}

// BarH adds plot bars on y axis.
func BarH(title string, data []float64) *BarHPlot {
	return &BarHPlot{
		title:  title,
		data:   data,
		height: 0.2,
		shift:  0,
		offset: 0,
	}
}

// Height sets bar height (in fact bars' width).
func (p *BarHPlot) Height(height float64) *BarHPlot {
	p.height = height
	return p
}

// Shift sets shift.
func (p *BarHPlot) Shift(shift float64) *BarHPlot {
	p.shift = shift
	return p
}

// Offset sets offset.
func (p *BarHPlot) Offset(offset int) *BarHPlot {
	p.offset = offset
	return p
}

// Plot implements plot interface.
func (p *BarHPlot) Plot() {
	imgui.ImPlotBarsH(Context.FontAtlas.RegisterString(p.title), p.data, p.height, p.shift, p.offset)
}

// LinePlot represents a plot line (linear chart).
type LinePlot struct {
	title      string
	values     []float64
	xScale, x0 float64
	offset     int
	yAxis      ImPlotYAxis
}

// Line adds a new plot line to the canvas.
func Line(title string, values []float64) *LinePlot {
	return &LinePlot{
		title:  title,
		values: values,
		xScale: 1,
		x0:     0,
		offset: 0,
	}
}

// SetPlotYAxis sets yAxis parameters.
func (p *LinePlot) SetPlotYAxis(yAxis ImPlotYAxis) *LinePlot {
	p.yAxis = yAxis
	return p
}

// XScale sets x-axis-scale.
func (p *LinePlot) XScale(scale float64) *LinePlot {
	p.xScale = scale
	return p
}

// X0 sets a start position on x axis.
func (p *LinePlot) X0(x0 float64) *LinePlot {
	p.x0 = x0
	return p
}

// Offset sets chart offset.
func (p *LinePlot) Offset(offset int) *LinePlot {
	p.offset = offset
	return p
}

// Plot implements Plot interface.
func (p *LinePlot) Plot() {
	imgui.ImPlotSetPlotYAxis(imgui.ImPlotYAxis(p.yAxis))
	imgui.ImPlotLine(Context.FontAtlas.RegisterString(p.title), p.values, p.xScale, p.x0, p.offset)
}

// LineXYPlot adds XY plot line.
type LineXYPlot struct {
	title  string
	xs, ys []float64
	offset int
	yAxis  ImPlotYAxis
}

// LineXY adds XY plot line to canvas.
func LineXY(title string, xvalues, yvalues []float64) *LineXYPlot {
	return &LineXYPlot{
		title:  title,
		xs:     xvalues,
		ys:     yvalues,
		offset: 0,
	}
}

// SetPlotYAxis sets yAxis parameters.
func (p *LineXYPlot) SetPlotYAxis(yAxis ImPlotYAxis) *LineXYPlot {
	p.yAxis = yAxis
	return p
}

// Offset sets chart's offset.
func (p *LineXYPlot) Offset(offset int) *LineXYPlot {
	p.offset = offset
	return p
}

// Plot implements Plot interface.
func (p *LineXYPlot) Plot() {
	imgui.ImPlotSetPlotYAxis(imgui.ImPlotYAxis(p.yAxis))
	imgui.ImPlotLineXY(Context.FontAtlas.RegisterString(p.title), p.xs, p.ys, p.offset)
}

// PieChartPlot represents a pie chart.
type PieChartPlot struct {
	labels       []string
	values       []float64
	x, y, radius float64
	normalize    bool
	labelFormat  string
	angle0       float64
}

// PieChart adds pie chart to the canvas.
func PieChart(labels []string, values []float64, x, y, radius float64) *PieChartPlot {
	return &PieChartPlot{
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

func (p *PieChartPlot) Normalize(n bool) *PieChartPlot {
	p.normalize = n
	return p
}

// LabelFormat sets format of labels.
func (p *PieChartPlot) LabelFormat(fmtStr string) *PieChartPlot {
	p.labelFormat = fmtStr
	return p
}

func (p *PieChartPlot) Angle0(a float64) *PieChartPlot {
	p.angle0 = a
	return p
}

func (p *PieChartPlot) Plot() {
	imgui.ImPlotPieChart(Context.FontAtlas.RegisterStringSlice(p.labels), p.values, p.x, p.y, p.radius, p.normalize, p.labelFormat, p.angle0)
}

type ScatterPlot struct {
	label      string
	values     []float64
	xscale, x0 float64
	offset     int
}

func Scatter(label string, values []float64) *ScatterPlot {
	return &ScatterPlot{
		label:  label,
		values: values,
		xscale: 1,
		x0:     0,
		offset: 0,
	}
}

func (p *ScatterPlot) XScale(s float64) *ScatterPlot {
	p.xscale = s
	return p
}

func (p *ScatterPlot) X0(x float64) *ScatterPlot {
	p.x0 = x
	return p
}

func (p *ScatterPlot) Offset(offset int) *ScatterPlot {
	p.offset = offset
	return p
}

func (p *ScatterPlot) Plot() {
	imgui.ImPlotScatter(Context.FontAtlas.RegisterString(p.label), p.values, p.xscale, p.x0, p.offset)
}

type ScatterXYPlot struct {
	label  string
	xs, ys []float64
	offset int
}

func ScatterXY(label string, xs, ys []float64) *ScatterXYPlot {
	return &ScatterXYPlot{
		label:  label,
		xs:     xs,
		ys:     ys,
		offset: 0,
	}
}

func (p *ScatterXYPlot) Offset(offset int) *ScatterXYPlot {
	p.offset = offset
	return p
}

func (p *ScatterXYPlot) Plot() {
	imgui.ImPlotScatterXY(Context.FontAtlas.RegisterString(p.label), p.xs, p.ys, p.offset)
}
