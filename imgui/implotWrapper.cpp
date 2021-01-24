#include "implot.h"
#include "implotWrapper.h"
#include "WrapperConverter.h"

void iggImPlotSetNextPlotLimits(double xmin, double xmax, double ymin, double ymax, int cond)
{
  ImPlot::SetNextPlotLimits(xmin, xmax, ymin, ymax, cond);
}

void iggImPlotSetNextPlotTicksX(const double* values, int n_ticks, const char* const labels[], IggBool show_default)
{
  ImPlot::SetNextPlotTicksX(values, n_ticks, labels, show_default);
}

void iggImPlotSetNextPlotTicksY(const double* values, int n_ticks, const char* const labels[], IggBool show_default, int y_axis)
{
  ImPlot::SetNextPlotTicksY(values, n_ticks, labels, show_default, y_axis);
}

void iggImPlotFitNextPlotAxes(IggBool x, IggBool y, IggBool y2, IggBool y3)
{
  ImPlot::FitNextPlotAxes(x, y, y2, y3);
}

IggImPlotContext iggImPlotCreateContext()
{
  ImPlotContext *context = ImPlot::CreateContext();
  ImPlot::GetStyle().AntiAliasedLines = true;
  return reinterpret_cast<IggImPlotContext>(context);
}

void iggImPlotDestroyContext()
{
  ImPlot::DestroyContext();
}

IggBool iggImPlotBeginPlot(const char* title_id,
                                  const char* x_label,
                                  const char* y_label,
                                  const IggVec2* size,
                                  int flags,
                                  int x_flags,
                                  int y_flags,
                                  int y2_flags,
                                  int y3_flags,
                                  const char* y2_label,
                                  const char* y3_label)
{
  Vec2Wrapper sizeArg(size);
  return ImPlot::BeginPlot(title_id, x_label, y_label, *sizeArg, flags, x_flags, y_flags, y2_flags, y3_flags, y2_label, y3_label);
}

void iggImPlotEndPlot()
{
  ImPlot::EndPlot();
}

void iggImPlotBars(const char* label_id, const double* values, int count, double width, double shift, int offset)
{
  ImPlot::PlotBars(label_id, values, count, width, shift, offset);
}

void iggImPlotBarsXY(const char* label_id, const double* xs, const double* ys, int count, double width, int offset)
{
  ImPlot::PlotBars(label_id, xs, ys, count, width, offset);
}

void iggImPlotBarsH(const char* label_id, const double* values, int count, double height, double shift, int offset)
{
  ImPlot::PlotBarsH(label_id, values, count, height, shift, offset);
}

void iggImPlotBarsHXY(const char* label_id, const double* xs, const double* ys, int count, double height, int offset)
{
  ImPlot::PlotBarsH(label_id, xs, ys, count, height, offset);
}

void iggImPlotErrorBars(const char* label_id, const double* xs, const double* ys, const double* err, int count, int offset)
{
  ImPlot::PlotErrorBars(label_id, xs, ys, err, count, offset);
}


void iggImPlotErrorBarsH(const char* label_id, const double* xs, const double* ys, const double* err, int count, int offset)
{
  ImPlot::PlotErrorBarsH(label_id, xs, ys, err, count, offset);
}

void iggImPlotLine(const char* label_id, const double* values, int count, double xscale, double x0, int offset)
{
  ImPlot::PlotLine(label_id, values, count, xscale, x0, offset);
}


void iggImPlotLineXY(const char* label_id, const double* xs, const double* ys, int count, int offset)
{
  ImPlot::PlotLine(label_id, xs, ys, count, offset);
}


void iggImPlotScatter(const char* label_id, const double* values, int count, double xscale, double x0, int offset)
{
  ImPlot::PlotScatter(label_id, values, count, xscale, x0, offset);
}

void iggImPlotScatterXY(const char* label_id, const double* xs, const double* ys, int count, int offset)
{
  ImPlot::PlotScatter(label_id, xs, ys, count, offset);
}

void iggImPlotStairs(const char* label_id, const double* values, int count, double xscale, double x0, int offset)
{
  ImPlot::PlotStairs(label_id, values, count, xscale, x0, offset);
}

void iggImPlotStairsXY(const char* label_id, const double* xs, const double* ys, int count, int offset)
{
  ImPlot::PlotStairs(label_id, xs, ys, count, offset);
}

void iggImPlotStems(const char* label_id, const double* values, int count, double y_ref, double xscale, double x0, int offset)
{
  ImPlot::PlotStems(label_id, values, count, y_ref, xscale, x0, offset);
}

void iggImPlotStemsXY(const char* label_id, const double* xs, const double* ys, int count, double y_ref, int offset)
{
  ImPlot::PlotStems(label_id, xs, ys, count, y_ref, offset);
}

void iggImPlotVLines(const char* label_id, const double* xs, int count, int offset)
{
  ImPlot::PlotVLines(label_id, xs, count, offset);
}

void iggImPlotHLines(const char* label_id, const double* ys, int count, int offset)
{
  ImPlot::PlotHLines(label_id, ys, count, offset);
}

void iggImPlotPieChart(const char* const label_ids[], const double* values, int count, double x, double y, double radius, IggBool normalize, const char* label_fmt, double angle0)
{
  ImPlot::PlotPieChart(label_ids, values, count, x, y, radius, normalize, label_fmt, angle0);
}

void iggImPlotGetPlotPos(IggVec2 *pos)
{
  exportValue(*pos, ImPlot::GetPlotPos());
}

void iggImPlotGetPlotSize(IggVec2 *size)
{
  exportValue(*size, ImPlot::GetPlotSize());
}

IggBool iggImPlotIsPlotHovered() {
  return ImPlot::IsPlotHovered() ? 1 : 0;
}

IggBool iggImPlotIsPlotXAxisHovered()
{
  return ImPlot::IsPlotXAxisHovered() ? 1 : 0;
}

IggBool iggImPlotIsPlotYAxisHovered(int y_axis)
{
  return ImPlot::IsPlotYAxisHovered(y_axis) ? 1 : 0;
}
