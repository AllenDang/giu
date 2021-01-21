#pragma once

#include "imguiWrapperTypes.h"
#include "implotWrapperTypes.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggImPlotSetNextPlotLimits(double xmin, double xmax, double ymin, double ymax, int cond);

extern void iggImPlotSetNextPlotTicksX(const double* values, int n_ticks, const char* const labels[], IggBool show_default);

extern void iggImPlotSetNextPlotTicksY(const double* values, int n_ticks, const char* const labels[], IggBool show_default, int y_axis);

extern void iggImPlotFitNextPlotAxes(IggBool x, IggBool y, IggBool y2, IggBool y3);

extern IggImPlotContext iggImPlotCreateContext();
extern void iggImPlotDestroyContext();

extern IggBool iggImPlotBeginPlot(const char* title_id,
                                  const char* x_label,
                                  const char* y_label,
                                  const IggVec2* size,
                                  int flags,
                                  int x_flags,
                                  int y_flags,
                                  int y2_flags,
                                  int y3_flags,
                                  const char* y2_label,
                                  const char* y3_label);
extern void iggImPlotEndPlot();

extern void iggImPlotBars(const char* label_id, const double* values, int count, double width, double shift, int offset);

extern void iggImPlotBarsH(const char* label_id, const double* values, int count, double height, double shift, int offset);

extern void iggImPlotLine(const char* label_id, const double* values, int count, double xscale, double x0, int offset);

#ifdef __cplusplus
}
#endif
