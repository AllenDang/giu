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

extern void iggImPlotLine(const char* label_id, const double* values, int count, double xscale, double x0, int offset);
extern void iggImPlotLineXY(const char* label_id, const double* xs, const double* ys, int count, int offset);

extern void iggImPlotScatter(const char* label_id, const double* values, int count, double xscale, double x0, int offset);
extern void iggImPlotScatterXY(const char* label_id, const double* xs, const double* ys, int count, int offset);

extern void iggImPlotBars(const char* label_id, const double* values, int count, double width, double shift, int offset);
extern void iggImPlotBarsXY(const char* label_id, const double* xs, const double* ys, int count, double width, int offset);

extern void iggImPlotBarsH(const char* label_id, const double* values, int count, double height, double shift, int offset);
extern void iggImPlotBarsHXY(const char* label_id, const double* xs, const double* ys, int count, double height, int offset);

extern void iggImPlotStairs(const char* label_id, const double* values, int count, double xscale, double x0, int offset);
extern void iggImPlotStairsXY(const char* label_id, const double* xs, const double* ys, int count, int offset);

extern void iggImPlotErrorBars(const char* label_id, const double* xs, const double* ys, const double* err, int count, int offset);
extern void iggImPlotErrorBarsH(const char* label_id, const double* xs, const double* ys, const double* err, int count, int offset);

extern void iggImPlotStems(const char* label_id, const double* values, int count, double y_ref, double xscale, double x0, int offset);
extern void iggImPlotStemsXY(const char* label_id, const double* xs, const double* ys, int count, double y_ref, int offset);

extern void iggImPlotVLines(const char* label_id, const double* xs, int count, int offset);
extern void iggImPlotHLines(const char* label_id, const double* ys, int count, int offset);

extern void iggImPlotPieChart(const char* const label_ids[], const double* values, int count, double x, double y, double radius, IggBool normalize, const char* label_fmt, double angle0);

extern void iggImPlotGetPlotPos(IggVec2 *pos);
extern void iggImPlotGetPlotSize(IggVec2 *size);
extern IggBool iggImPlotIsPlotHovered();
extern IggBool iggImPlotIsPlotXAxisHovered();
extern IggBool iggImPlotIsPlotYAxisHovered(int y_axis);

#ifdef __cplusplus
}
#endif
