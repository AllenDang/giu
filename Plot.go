package giu

import (
	"image"

	imgui "github.com/AllenDang/cimgui-go"
)

type (
	PlotXAxis = imgui.PlotAxisEnum
	PlotYAxis = imgui.PlotAxisEnum
)

const (
	AxisX1 = imgui.AxisX1
	AxisX2 = imgui.AxisX2
	AxisX3 = imgui.AxisX3
	AxisY1 = imgui.AxisY1
	AxisY2 = imgui.AxisY2
	AxisY3 = imgui.AxisY3
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

func (p *PlotCanvasWidget) SetXAxisLabel(axis PlotXAxis, label string) *PlotCanvasWidget {
	switch axis {
	case AxisX1:
		p.xLabel = label
	case AxisX2:
		p.y2Label = label
	case AxisX3:
		p.y3Label = label
	}

	return p
}

func (p *PlotCanvasWidget) SetYAxisLabel(axis PlotYAxis, label string) *PlotCanvasWidget {
	switch axis {
	case AxisY1:
		p.yLabel = label
	case AxisY2:
		p.y2Label = label
	case AxisY3:
		p.y3Label = label
	}

	return p
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
	if len(p.plots) == 0 {
		return
	}

	if imgui.PlotBeginPlotV(
		Context.FontAtlas.RegisterString(p.title),
		ToVec2(image.Pt(p.width, p.height)),
		imgui.PlotFlags(p.flags),
	) {
		imgui.PlotSetupAxisLimitsV(
			imgui.AxisX1,
			p.xMin,
			p.xMax,
			imgui.PlotCond(p.axisLimitCondition),
		)
		imgui.PlotSetupAxisLimitsV(
			imgui.AxisY1,
			p.yMin,
			p.yMax,
			imgui.PlotCond(p.axisLimitCondition),
		)

		if len(p.xTicksValue) > 0 {
			imgui.PlotSetupAxisTicksdoublePtrV(
				imgui.AxisX1,
				&p.xTicksValue,
				int32(len(p.xTicksValue)),
				p.xTicksLabel,
				p.xTicksShowDefault,
			)
		}

		if len(p.yTicksValue) > 0 {
			imgui.PlotSetupAxisTicksdoublePtrV(
				imgui.AxisY1,
				&p.yTicksValue,
				int32(len(p.yTicksValue)),
				p.yTicksLabel,
				p.yTicksShowDefault,
			)
		}

		imgui.PlotSetupAxisV(
			imgui.AxisX1,
			Context.FontAtlas.RegisterString(p.xLabel),
			imgui.PlotAxisFlags(p.xFlags),
		)

		imgui.PlotSetupAxisV(
			imgui.AxisY1,
			Context.FontAtlas.RegisterString(p.yLabel),
			imgui.PlotAxisFlags(p.yFlags),
		)

		if p.y2Label != "" {
			imgui.PlotSetupAxisV(
				imgui.AxisY2,
				Context.FontAtlas.RegisterString(p.y2Label),
				imgui.PlotAxisFlags(p.y2Flags),
			)
		}

		if p.y3Label != "" {
			imgui.PlotSetupAxisV(
				imgui.AxisY3,
				Context.FontAtlas.RegisterString(p.y3Label),
				imgui.PlotAxisFlags(p.y3Flags),
			)
		}

		for _, plot := range p.plots {
			plot.Plot()
		}

		imgui.PlotEndPlot()
	}
}

func SwitchPlotAxes(x PlotXAxis, y PlotYAxis) PlotWidget {
	return Custom(func() {
		imgui.PlotSetAxes(x, y)
	})
}

// BarPlot adds bar plot (column chart) to the canvas.
type BarPlot struct {
	title  string
	data   []float64
	width  float64
	shift  float64
	offset int
}

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
	imgui.PlotPlotBarsdoublePtrIntV(
		p.title,
		&p.data,
		int32(len(p.data)),
		p.width,
		p.shift,
		0, // TODO: implement
		int32(p.offset),
		8, // in fact this is sizeof(double) = 8
	)
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
	imgui.PlotPlotBarsdoublePtrIntV(
		Context.FontAtlas.RegisterString(p.title),
		&p.data,
		int32(len(p.data)),
		p.height,
		p.shift,
		imgui.PlotBarsFlagsHorizontal,
		int32(p.offset),
		0,
	)
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
	imgui.PlotSetAxis(
		imgui.PlotAxisEnum(p.yAxis),
	)

	imgui.PlotPlotLinedoublePtrIntV(
		Context.FontAtlas.RegisterString(p.title),
		&p.values,
		int32(len(p.values)),
		p.xScale,
		p.x0,
		0, // flags
		int32(p.offset),
		8, // in fact this is sizeof(double) = 8
	)
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
	imgui.PlotSetAxis(imgui.PlotAxisEnum(p.yAxis))
	imgui.PlotPlotLinedoublePtrdoublePtrV(
		Context.FontAtlas.RegisterString(p.title),
		&p.xs,
		&p.ys,
		int32(len(p.xs)),
		0, // flags
		int32(p.offset),
		8, // in fact this is sizeof(double) = 8
	)
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
	// TODO: p.normalized not used anymore - replace with flags
	imgui.PlotPlotPieChartdoublePtrStrV(
		Context.FontAtlas.RegisterStringSlice(p.labels),
		&p.values,
		int32(len(p.values)),
		p.x,
		p.y,
		p.radius,
		p.labelFormat,
		p.angle0,
		imgui.PlotPieChartFlagsNormalize,
	)
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
	imgui.PlotPlotScatterdoublePtrIntV(
		Context.FontAtlas.RegisterString(p.label),
		&p.values,
		int32(len(p.values)),
		p.xscale,
		p.x0,
		0, // TODO: implement flags
		int32(p.offset),
		8, // in fact this is sizeof(double) = 8
	)
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
	imgui.PlotPlotScatterdoublePtrdoublePtrV(
		Context.FontAtlas.RegisterString(p.label),
		&p.xs,
		&p.ys,
		int32(len(p.xs)),
		0, // TODO: implement
		int32(p.offset),
		8, // in fact this is sizeof(double) = 8
	)
}
