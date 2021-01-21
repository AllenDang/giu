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

void iggImPlotBarsH(const char* label_id, const double* values, int count, double height, double shift, int offset)
{
  ImPlot::PlotBarsH(label_id, values, count, height, shift, offset);
}

void iggImPlotLine(const char* label_id, const double* values, int count, double xscale, double x0, int offset)
{
  ImPlot::PlotLine(label_id, values, count, xscale, x0, offset);
}
