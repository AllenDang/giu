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

type PlotAxisConfig struct {
	enabled            bool
	label              string
	scale              PlotScale
	flags              PlotAxisFlags
	limitMin, limitMax float64
	limitCond          ExecCondition
}

func (pac *PlotAxisConfig) Enable() *PlotAxisConfig {
	pac.enabled = true
	return pac
}

func (pac *PlotAxisConfig) Label(l string) *PlotAxisConfig {
	pac.label = l
	return pac
}

func (pac *PlotAxisConfig) Scale(s PlotScale) *PlotAxisConfig {
	pac.scale = s
	return pac
}

func (pac *PlotAxisConfig) Flags(f PlotAxisFlags) *PlotAxisConfig {
	pac.flags = f
	return pac
}

func (pac *PlotAxisConfig) Min(m float64) *PlotAxisConfig {
	pac.limitMin = m
	return pac
}

func (pac *PlotAxisConfig) Max(m float64) *PlotAxisConfig {
	pac.limitMax = m
	return pac
}

func (pac *PlotAxisConfig) LimitCond(c ExecCondition) *PlotAxisConfig {
	pac.limitCond = c
	return pac
}

// PlotCanvasWidget represents a giu plot widget.
type PlotCanvasWidget struct {
	title  string
	width  int
	height int
	flags  PlotFlags
	axes   map[implot.AxisEnum]*PlotAxisConfig

	axisLimitCondition       ExecCondition
	xTicksValue, yTicksValue []float64
	xTicksLabel, yTicksLabel []string
	xTicksShowDefault        bool
	yTicksShowDefault        bool
	yTicksYAxis              ImPlotYAxis
	plots                    []PlotWidget
}

// Plot adds creates a new plot widget.
func Plot(title string) *PlotCanvasWidget {
	axes := make(map[PlotYAxis]*PlotAxisConfig)
	axes[AxisY1] = (&PlotAxisConfig{}).Enable().Flags(PlotAxisFlagsNone)
	axes[AxisY2] = (&PlotAxisConfig{}).Flags(PlotAxisFlagsNoGridLines)
	axes[AxisY3] = (&PlotAxisConfig{}).Flags(PlotAxisFlagsNoGridLines)
	axes[AxisX1] = (&PlotAxisConfig{}).Enable().Flags(PlotAxisFlagsNone).Min(0).Max(10)
	axes[AxisX2] = (&PlotAxisConfig{}).Flags(PlotAxisFlagsNone)
	axes[AxisX3] = (&PlotAxisConfig{}).Flags(PlotAxisFlagsNone)
	return &PlotCanvasWidget{
		title: title,

		axes:               axes,
		width:              -1,
		height:             0,
		flags:              PlotFlagsNone,
		xTicksShowDefault:  true,
		yTicksShowDefault:  true,
		yTicksYAxis:        0,
		axisLimitCondition: ConditionOnce,
	}
}

// XLabel sets label for each x axis. If none specified, it will default to AxisX1.
func (p *PlotCanvasWidget) XLabel(label string, axes ...PlotXAxis) *PlotCanvasWidget {
	if len(axes) == 0 {
		axes = append(axes, AxisX1)
	}

	for _, axis := range axes {
		p.axes[axis].Enable().Label(label)
	}

	return p
}

// YLabel sets label for each y axis. If none specified, it will default to AxisY1.
func (p *PlotCanvasWidget) YLabel(label string, axes ...PlotYAxis) *PlotCanvasWidget {
	if len(axes) == 0 {
		axes = append(axes, AxisY1)
	}

	for _, axis := range axes {
		p.axes[axis].Enable().Label(label)
	}

	return p
}

// Limits sets X and Y axis limits for the default axis (AxisX1 and AxisY1).
func (p *PlotCanvasWidget) Limits(xmin, xmax, ymin, ymax float64, cond ExecCondition) *PlotCanvasWidget {
	return p.XLimits(xmin, xmax, cond).YLimits(ymin, ymax, cond)
}

// XLimits allows to set X axis limits.
func (p *PlotCanvasWidget) XLimits(xmin, xmax float64, cond ExecCondition, axes ...PlotXAxis) *PlotCanvasWidget {
	if len(axes) == 0 {
		axes = append(axes, AxisX1)
	}

	for _, axis := range axes {
		p.axes[axis].Enable().Min(xmin).Max(xmax).LimitCond(cond)
	}

	return p
}

// YLimits allows to set Y axis limits.
func (p *PlotCanvasWidget) YLimits(ymin, ymax float64, cond ExecCondition, axes ...PlotYAxis) *PlotCanvasWidget {
	if len(axes) == 0 {
		axes = append(axes, AxisY1)
	}

	for _, axis := range axes {
		p.axes[axis].Enable().Min(ymin).Max(ymax).LimitCond(cond)
	}

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

// XFlags sets x axis fags.
func (p *PlotCanvasWidget) XFlags(flags PlotAxisFlags, axes ...PlotXAxis) *PlotCanvasWidget {
	if len(axes) == 0 {
		axes = append(axes, AxisX1)
	}
	for _, axis := range axes {
		p.axes[axis].Enable().Flags(flags)
	}

	return p
}

// YFlags sets y axis flags. You can specify multiple axes to apply those flags. Default is Plot.
func (p *PlotCanvasWidget) YFlags(yFlags PlotAxisFlags, axes ...PlotYAxis) *PlotCanvasWidget {
	if len(axes) == 0 {
		axes = append(axes, AxisY1)
	}
	for _, axis := range axes {
		p.axes[axis].Enable().Flags(yFlags)
	}

	return p
}

// XScale sets the plot x axis scale.
func (p *PlotCanvasWidget) XScale(scale PlotScale, axes ...PlotXAxis) *PlotCanvasWidget {
	if len(axes) == 0 {
		axes = append(axes, AxisX1)
	}

	for _, axis := range axes {
		p.axes[axis].Enable().Scale(scale)
	}

	return p
}

// YScale sets the plot y axis scale.
func (p *PlotCanvasWidget) YScale(scale PlotScale, axes ...PlotYAxis) *PlotCanvasWidget {
	if len(axes) == 0 {
		axes = append(axes, AxisY1)
	}

	for _, axis := range axes {
		p.axes[axis].Enable().Scale(scale)
	}

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
		// set up y axes
		for axis, cfg := range p.axes {
			if !cfg.enabled {
				continue
			}

			implot.SetupAxisV(
				axis,
				cfg.label,
				implot.AxisFlags(cfg.flags),
			)

			implot.SetupAxisScalePlotScale(
				axis,
				implot.Scale(p.axes[axis].scale),
			)

			implot.SetupAxisLimitsV(
				axis,
				p.axes[axis].limitMin,
				p.axes[axis].limitMax,
				implot.Cond(p.axes[axis].limitCond),
			)
		}

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
