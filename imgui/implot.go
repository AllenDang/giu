package imgui

// #cgo CXXFLAGS: -std=c++11
// #include "implotWrapper.h"
import "C"
import "unsafe"

// The following functions MUST be called BEFORE BeginPlot!
// Set the axes range limits of the next plot. Call right before BeginPlot(). If ImGuiCond_Always is used, the axes limits will be locked.
func ImPlotSetNextPlotLimits(xmin, xmax, ymin, ymax float64, cond Condition) {
	C.iggImPlotSetNextPlotLimits(C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.int(cond))
}

func ImPlotSetNextPlotTicksX(values []float64, labels []string, showDefault bool) {
	if len(values) == 0 || len(labels) == 0 {
		return
	}

	labelsArg := make([]*C.char, len(labels))
	for i, l := range labels {
		la, lf := wrapString(l)
		defer lf()
		labelsArg[i] = la
	}

	C.iggImPlotSetNextPlotTicksX(
		(*C.double)(unsafe.Pointer(&values[0])),
		C.int(len(values)),
		&labelsArg[0],
		castBool(showDefault),
	)
}

func ImPlotSetNextPlotTicksY(values []float64, labels []string, showDefault bool, yAxis int) {
	if len(values) == 0 || len(labels) == 0 {
		return
	}

	labelsArg := make([]*C.char, len(labels))
	for i, l := range labels {
		la, lf := wrapString(l)
		defer lf()
		labelsArg[i] = la
	}

	C.iggImPlotSetNextPlotTicksY(
		(*C.double)(unsafe.Pointer(&values[0])),
		C.int(len(values)),
		&labelsArg[0],
		castBool(showDefault),
		C.int(yAxis),
	)
}

func ImPlotFitNextPlotAxis(x, y, y2, y3 bool) {
	C.iggImPlotFitNextPlotAxes(
		castBool(x),
		castBool(y),
		castBool(y2),
		castBool(y3),
	)
}

type ImPlotContext struct {
	handle C.IggImPlotContext
}

// Creates a new ImPlot context. Call this after ImGui::CreateContext.
func ImPlotCreateContext() *ImPlotContext {
	return &ImPlotContext{handle: C.iggImPlotCreateContext()}
}

// Destroys an ImPlot context. Call this before ImGui::DestroyContext. NULL = destroy current context
func ImPlotDestroyContext() {
	C.iggImPlotDestroyContext()
}

type ImPlotFlags int

const (
	ImPlotFlags_None        ImPlotFlags = 0       // default
	ImPlotFlags_NoTitle     ImPlotFlags = 1 << 0  // the plot title will not be displayed (titles are also hidden if preceeded by double hashes, e.g. "##MyPlot")
	ImPlotFlags_NoLegend    ImPlotFlags = 1 << 1  // the legend will not be displayed
	ImPlotFlags_NoMenus     ImPlotFlags = 1 << 2  // the user will not be able to open context menus with right-click
	ImPlotFlags_NoBoxSelect ImPlotFlags = 1 << 3  // the user will not be able to box-select with right-click drag
	ImPlotFlags_NoMousePos  ImPlotFlags = 1 << 4  // the mouse position, in plot coordinates, will not be displayed inside of the plot
	ImPlotFlags_NoHighlight ImPlotFlags = 1 << 5  // plot items will not be highlighted when their legend entry is hovered
	ImPlotFlags_NoChild     ImPlotFlags = 1 << 6  // a child window region will not be used to capture mouse scroll (can boost performance for single ImGui window applications)
	ImPlotFlags_Equal       ImPlotFlags = 1 << 7  // primary x and y axes will be constrained to have the same units/pixel (does not apply to auxiliary y-axes)
	ImPlotFlags_YAxis2      ImPlotFlags = 1 << 8  // enable a 2nd y-axis on the right side
	ImPlotFlags_YAxis3      ImPlotFlags = 1 << 9  // enable a 3rd y-axis on the right side
	ImPlotFlags_Query       ImPlotFlags = 1 << 10 // the user will be able to draw query rects with middle-mouse or CTRL + right-click drag
	ImPlotFlags_Crosshairs  ImPlotFlags = 1 << 11 // the default mouse cursor will be replaced with a crosshair when hovered
	ImPlotFlags_AntiAliased ImPlotFlags = 1 << 12 // plot lines will be software anti-aliased (not recommended for high density plots, prefer MSAA)
	ImPlotFlags_CanvasOnly  ImPlotFlags = ImPlotFlags_NoTitle | ImPlotFlags_NoLegend | ImPlotFlags_NoMenus | ImPlotFlags_NoBoxSelect | ImPlotFlags_NoMousePos
)

type ImPlotAxisFlags int

const (
	ImPlotAxisFlags_None          ImPlotAxisFlags = 0      // default
	ImPlotAxisFlags_NoLabel       ImPlotAxisFlags = 1 << 0 // the axis label will not be displayed (axis labels also hidden if the supplied string name is NULL)
	ImPlotAxisFlags_NoGridLines   ImPlotAxisFlags = 1 << 1 // the axis grid lines will not be displayed
	ImPlotAxisFlags_NoTickMarks   ImPlotAxisFlags = 1 << 2 // the axis tick marks will not be displayed
	ImPlotAxisFlags_NoTickLabels  ImPlotAxisFlags = 1 << 3 // the axis tick labels will not be displayed
	ImPlotAxisFlags_LogScale      ImPlotAxisFlags = 1 << 4 // a logartithmic (base 10) axis scale will be used (mutually exclusive with ImPlotAxisFlags_Time)
	ImPlotAxisFlags_Time          ImPlotAxisFlags = 1 << 5 // axis will display date/time formatted labels (mutually exclusive with ImPlotAxisFlags_LogScale)
	ImPlotAxisFlags_Invert        ImPlotAxisFlags = 1 << 6 // the axis will be inverted
	ImPlotAxisFlags_LockMin       ImPlotAxisFlags = 1 << 7 // the axis minimum value will be locked when panning/zooming
	ImPlotAxisFlags_LockMax       ImPlotAxisFlags = 1 << 8 // the axis maximum value will be locked when panning/zooming
	ImPlotAxisFlags_Lock          ImPlotAxisFlags = ImPlotAxisFlags_LockMin | ImPlotAxisFlags_LockMax
	ImPlotAxisFlags_NoDecorations ImPlotAxisFlags = ImPlotAxisFlags_NoLabel | ImPlotAxisFlags_NoGridLines | ImPlotAxisFlags_NoTickMarks | ImPlotAxisFlags_NoTickLabels
)

//-----------------------------------------------------------------------------
// Begin/End Plot
//-----------------------------------------------------------------------------

// Starts a 2D plotting context. If this function returns true, EndPlot() must
// be called, e.g. "if (BeginPlot(...)) { ... EndPlot(); }". #title_id must
// be unique. If you need to avoid ID collisions or don't want to display a
// title in the plot, use double hashes (e.g. "MyPlot##Hidden" or "##NoTitle").
// If #x_label and/or #y_label are provided, axes labels will be displayed.
func ImPlotBegin(title string, xLabel, yLabel string, size Vec2, flags ImPlotFlags, xFlags, yFlags, y2Flags, y3Flags ImPlotAxisFlags, y2Label, y3Label string) bool {
	titleArg, titleFin := wrapString(title)
	defer titleFin()

	xLabelArg, xLabelFin := wrapString(xLabel)
	defer xLabelFin()

	yLabelArg, yLabelFin := wrapString(yLabel)
	defer yLabelFin()

	sizeArg, _ := size.wrapped()

	y2LabelArg, y2LabelFin := wrapString(y2Label)
	defer y2LabelFin()

	y3LabelArg, y3LabelFin := wrapString(y3Label)
	defer y3LabelFin()

	return C.iggImPlotBeginPlot(
		titleArg,
		xLabelArg,
		yLabelArg,
		sizeArg,
		C.int(flags),
		C.int(xFlags),
		C.int(yFlags),
		C.int(y2Flags),
		C.int(y3Flags),
		y2LabelArg,
		y3LabelArg) != 0
}

// Only call EndPlot() if BeginPlot() returns true! Typically called at the end
// of an if statement conditioned on BeginPlot().
func ImPlotEnd() {
	C.iggImPlotEndPlot()
}

// Plots a vertical bar graph. #width and #shift are in X units.
func ImPlotBars(label string, values []float64, width, shift float64, offset int) {
	if len(values) == 0 {
		return
	}

	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.iggImPlotBars(labelArg, (*C.double)(unsafe.Pointer(&values[0])), C.int(len(values)), C.double(width), C.double(shift), C.int(offset))
}

func ImPlotBarsH(label string, values []float64, height, shift float64, offset int) {
	if len(values) == 0 {
		return
	}

	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.iggImPlotBarsH(labelArg, (*C.double)(unsafe.Pointer(&values[0])), C.int(len(values)), C.double(height), C.double(shift), C.int(offset))
}

func ImPlotLine(label string, values []float64, xscale, x0 float64, offset int) {
	if len(values) == 0 {
		return
	}

	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.iggImPlotLine(labelArg, (*C.double)(unsafe.Pointer(&values[0])), C.int(len(values)), C.double(xscale), C.double(x0), C.int(offset))
}

func ImPlotLineXY(label string, xs, ys []float64, offset int) {
	if len(xs) == 0 || len(ys) == 0 || (len(xs) != len(ys)) {
		return
	}

	labelArg, labelFin := wrapString(label)
	defer labelFin()

	C.iggImPlotLineXY(
		labelArg,
		(*C.double)(unsafe.Pointer(&xs[0])),
		(*C.double)(unsafe.Pointer(&ys[0])),
		C.int(len(xs)),
		C.int(offset),
	)

}
