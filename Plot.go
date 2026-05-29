package giu

import (
	"image"

	"github.com/AllenDang/cimgui-go/implot"
	"github.com/AllenDang/cimgui-go/utils"
)

type (
	// PlotXAxis allows to chose X axis.
	PlotXAxis = implot.AxisEnum
	// PlotYAxis allows to chose Y axis.
	PlotYAxis = implot.AxisEnum
)

// Available axes.
const (
	AxisX1 = implot.AxisX1
	AxisX2 = implot.AxisX2
	AxisX3 = implot.AxisX3
	AxisY1 = implot.AxisY1
	AxisY2 = implot.AxisY2
	AxisY3 = implot.AxisY3
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
	xScale, yScale, y2Scale, y3Scale PlotScale
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
		xScale:             PlotScaleLinear,
		yScale:             PlotScaleLinear,
		y2Scale:            PlotScaleLinear,
		y3Scale:            PlotScaleLinear,
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

// XLabel sets x axis label.
func (p *PlotCanvasWidget) XLabel(label string, axes ...PlotXAxis) *PlotCanvasWidget {
	if len(axes) == 0 {
		p.xLabel = label
	} else {
		for _, axis := range axes {
			switch axis {
			case AxisX1:
				p.xLabel = label
			case AxisX2:
				panic("TODO: X2Axis not implemented in giu yet.")
			case AxisX3:
				panic("TODO: X3Axis not implemented in giu yet.")
			}
		}
	}

	return p
}

// SetYAxisLabel sets y axis label.
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

// XScale sets the plot x axis scale.
func (p *PlotCanvasWidget) XScale(scale PlotScale) *PlotCanvasWidget {
	p.xScale = scale
	return p
}

// YScale sets the plot y axis scale.
func (p *PlotCanvasWidget) YScale(scale PlotScale) *PlotCanvasWidget {
	p.yScale = scale
	return p
}

// Y2Scale sets the plot y2 axis scale.
func (p *PlotCanvasWidget) Y2Scale(scale PlotScale) *PlotCanvasWidget {
	p.y2Scale = scale
	return p
}

// Y3Scale sets the plot y3 axis scale.
func (p *PlotCanvasWidget) Y3Scale(scale PlotScale) *PlotCanvasWidget {
	p.y3Scale = scale
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

	if implot.BeginPlotV(
		Context.PrepareString(p.title),
		ToVec2(image.Pt(p.width, p.height)),
		implot.Flags(p.flags),
	) {
		implot.SetupAxisScalePlotScale(
			implot.AxisX1,
			implot.Scale(p.xScale),
		)
		implot.SetupAxisScalePlotScale(
			implot.AxisY1,
			implot.Scale(p.yScale),
		)

		if p.y2Label != "" {
			implot.SetupAxisScalePlotScale(
				implot.AxisY2,
				implot.Scale(p.y2Scale),
			)
		}

		if p.y3Label != "" {
			implot.SetupAxisScalePlotScale(
				implot.AxisY3,
				implot.Scale(p.y3Scale),
			)
		}

		implot.SetupAxisLimitsV(
			implot.AxisX1,
			p.xMin,
			p.xMax,
			implot.Cond(p.axisLimitCondition),
		)
		implot.SetupAxisLimitsV(
			implot.AxisY1,
			p.yMin,
			p.yMax,
			implot.Cond(p.axisLimitCondition),
		)

		if len(p.xTicksValue) > 0 {
			implot.SetupAxisTicksdoublePtrV(
				implot.AxisX1,
				utils.SliceToPtr(p.xTicksValue),
				int32(len(p.xTicksValue)),
				p.xTicksLabel,
				p.xTicksShowDefault,
			)
		}

		if len(p.yTicksValue) > 0 {
			implot.SetupAxisTicksdoublePtrV(
				implot.AxisY1,
				utils.SliceToPtr(p.yTicksValue),
				int32(len(p.yTicksValue)),
				p.yTicksLabel,
				p.yTicksShowDefault,
			)
		}

		implot.SetupAxisV(
			implot.AxisX1,
			Context.PrepareString(p.xLabel),
			implot.AxisFlags(p.xFlags),
		)

		implot.SetupAxisV(
			implot.AxisY1,
			Context.PrepareString(p.yLabel),
			implot.AxisFlags(p.yFlags),
		)

		if p.y2Label != "" {
			implot.SetupAxisV(
				implot.AxisY2,
				Context.PrepareString(p.y2Label),
				implot.AxisFlags(p.y2Flags),
			)
		}

		if p.y3Label != "" {
			implot.SetupAxisV(
				implot.AxisY3,
				Context.PrepareString(p.y3Label),
				implot.AxisFlags(p.y3Flags),
			)
		}

		for _, plot := range p.plots {
			plot.Plot()
		}

		implot.EndPlot()
	}
}

// SwitchPlotAxes switches plot axes.
func SwitchPlotAxes(x PlotXAxis, y PlotYAxis) PlotWidget {
	return Custom(func() {
		implot.SetAxes(x, y)
	})
}

// BarPlot adds bar plot (column chart) to the canvas.
type BarPlot struct {
	title string
	data  []float64
	width float64
	shift float64
	spec  *PlotSpec
}

// Bar adds plot bars to the canvas.
func Bar(title string, data []float64) *BarPlot {
	return &BarPlot{
		title: title,
		data:  data,
		width: 0.2,
		shift: 0,
		spec:  NewPlotSpec(),
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
	p.spec.SetProperty(PlotPropertyOffset, offset)
	return p
}

// Plot implements Plot interface.
func (p *BarPlot) Plot() {
	p.spec.SetProperty(PlotPropertyStride, int(8))
	implot.PlotBarsdoublePtrIntV(
		p.title,
		utils.SliceToPtr(p.data),
		int32(len(p.data)),
		p.width,
		p.shift,
		*p.spec.GetSpec(),
	)
}

// BarHPlot represents a column chart on Y axis.
type BarHPlot struct {
	title  string
	data   []float64
	height float64
	shift  float64
	spec   *PlotSpec
}

// BarH adds plot bars on y axis.
func BarH(title string, data []float64) *BarHPlot {
	return &BarHPlot{
		title:  title,
		data:   data,
		height: 0.2,
		shift:  0,
		spec:   NewPlotSpec(),
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
	p.spec.SetProperty(PlotPropertyOffset, offset)
	return p
}

// Plot implements plot interface.
func (p *BarHPlot) Plot() {
	p.spec.SetProperty(PlotPropertyFlags, int(implot.BarsFlagsHorizontal))
	implot.PlotBarsdoublePtrIntV(
		Context.PrepareString(p.title),
		utils.SliceToPtr(p.data),
		int32(len(p.data)),
		p.height,
		p.shift,
		*p.spec.GetSpec(),
	)
}

// LinePlot represents a plot line (linear chart).
type LinePlot struct {
	title      string
	values     []float64
	xScale, x0 float64
	yAxis      ImPlotYAxis
	spec       *PlotSpec
}

// Line adds a new plot line to the canvas.
func Line(title string, values []float64) *LinePlot {
	return &LinePlot{
		title:  title,
		values: values,
		xScale: 1,
		x0:     0,
		spec:   NewPlotSpec(),
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
	p.spec.SetProperty(PlotPropertyOffset, offset)
	return p
}

// Plot implements Plot interface.
func (p *LinePlot) Plot() {
	implot.SetAxis(
		implot.AxisEnum(p.yAxis),
	)

	p.spec.SetProperty(PlotPropertyStride, int(8))

	implot.PlotLinedoublePtrIntV(
		Context.PrepareString(p.title),
		utils.SliceToPtr(p.values),
		int32(len(p.values)),
		p.xScale,
		p.x0,
		*p.spec.GetSpec(),
	)
}

// LineXYPlot adds XY plot line.
type LineXYPlot struct {
	title  string
	xs, ys []float64
	yAxis  ImPlotYAxis
	spec   *PlotSpec
}

// LineXY adds XY plot line to canvas.
func LineXY(title string, xvalues, yvalues []float64) *LineXYPlot {
	return &LineXYPlot{
		title: title,
		xs:    xvalues,
		ys:    yvalues,
		spec:  NewPlotSpec(),
	}
}

// SetPlotYAxis sets yAxis parameters.
func (p *LineXYPlot) SetPlotYAxis(yAxis ImPlotYAxis) *LineXYPlot {
	p.yAxis = yAxis
	return p
}

// Offset sets chart's offset.
func (p *LineXYPlot) Offset(offset int) *LineXYPlot {
	p.spec.SetProperty(PlotPropertyOffset, offset)
	return p
}

// Plot implements Plot interface.
func (p *LineXYPlot) Plot() {
	implot.SetAxis(implot.AxisEnum(p.yAxis))
	p.spec.SetProperty(PlotPropertyStride, int(8))
	implot.PlotLinedoublePtrdoublePtrV(
		Context.PrepareString(p.title),
		utils.SliceToPtr(p.xs),
		utils.SliceToPtr(p.ys),
		int32(len(p.xs)),
		*p.spec.GetSpec(),
	)
}

// PieChartPlot represents a pie chart.
// TODO: support PlotPieChartFlags.
type PieChartPlot struct {
	labels       []string
	values       []float64
	x, y, radius float64
	normalize    bool
	labelFormat  string
	angle0       float64
	spec         *PlotSpec
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
		spec:        NewPlotSpec(),
	}
}

// Normalize sets normalize flag.
func (p *PieChartPlot) Normalize(n bool) *PieChartPlot {
	p.normalize = n
	return p
}

// LabelFormat sets format of labels.
func (p *PieChartPlot) LabelFormat(fmtStr string) *PieChartPlot {
	p.labelFormat = fmtStr
	return p
}

// Angle0 sets start angle.
func (p *PieChartPlot) Angle0(a float64) *PieChartPlot {
	p.angle0 = a
	return p
}

// Plot implements Plot interface.
func (p *PieChartPlot) Plot() {
	var flags implot.PieChartFlags
	if p.normalize {
		flags |= implot.PieChartFlagsNormalize
	}

	p.spec.SetProperty(PlotPropertyFlags, int(flags))
	implot.PlotPieChartdoublePtrStrV(
		Context.PrepareStringSlice(p.labels),
		utils.SliceToPtr(p.values),
		int32(len(p.values)),
		p.x,
		p.y,
		p.radius,
		p.labelFormat,
		p.angle0,
		*p.spec.GetSpec(),
	)
}

// ScatterPlot represents a scatter plot.
type ScatterPlot struct {
	label      string
	values     []float64
	xscale, x0 float64
	spec       *PlotSpec
}

// Scatter adds scatter plot to the canvas.
func Scatter(label string, values []float64) *ScatterPlot {
	return &ScatterPlot{
		label:  label,
		values: values,
		xscale: 1,
		x0:     0,
		spec:   NewPlotSpec(),
	}
}

// XScale sets x-axis scale.
func (p *ScatterPlot) XScale(s float64) *ScatterPlot {
	p.xscale = s
	return p
}

// X0 sets start position on x axis.
func (p *ScatterPlot) X0(x float64) *ScatterPlot {
	p.x0 = x
	return p
}

// Offset sets chart offset.
func (p *ScatterPlot) Offset(offset int) *ScatterPlot {
	p.spec.SetProperty(PlotPropertyOffset, offset)
	return p
}

// Plot implements Plot interface.
func (p *ScatterPlot) Plot() {
	p.spec.SetProperty(PlotPropertyStride, int(8))
	implot.PlotScatterdoublePtrIntV(
		Context.PrepareString(p.label),
		utils.SliceToPtr(p.values),
		int32(len(p.values)),
		p.xscale,
		p.x0,
		*p.spec.GetSpec(),
	)
}

// ScatterXYPlot represents a scatter plot with possibility to set x and y values.
type ScatterXYPlot struct {
	label  string
	xs, ys []float64
	spec   *PlotSpec
}

// ScatterXY adds scatter plot with x and y values.
func ScatterXY(label string, xs, ys []float64) *ScatterXYPlot {
	return &ScatterXYPlot{
		label: label,
		xs:    xs,
		ys:    ys,
		spec:  NewPlotSpec(),
	}
}

// Offset sets chart offset.
func (p *ScatterXYPlot) Offset(offset int) *ScatterXYPlot {
	p.spec.SetProperty(PlotPropertyOffset, offset)
	return p
}

// Plot implements Plot interface.
func (p *ScatterXYPlot) Plot() {
	p.spec.SetProperty(PlotPropertyStride, int(8))
	implot.PlotScatterdoublePtrdoublePtrV(
		Context.PrepareString(p.label),
		utils.SliceToPtr(p.xs),
		utils.SliceToPtr(p.ys),
		int32(len(p.xs)),
		*p.spec.GetSpec(),
	)
}
