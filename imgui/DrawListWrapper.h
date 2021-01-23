#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C" {
#endif

extern int iggDrawListGetCommandCount(IggDrawList handle);
extern IggDrawCmd iggDrawListGetCommand(IggDrawList handle, int index);
extern void iggDrawListGetRawIndexBuffer(IggDrawList handle, void **data,
                                         int *byteSize);
extern void iggDrawListGetRawVertexBuffer(IggDrawList handle, void **data,
                                          int *byteSize);

extern void iggGetIndexBufferLayout(size_t *entrySize);
extern void iggGetVertexBufferLayout(size_t *entrySize, size_t *posOffset,
                                     size_t *uvOffset, size_t *colOffset);

extern void iggDrawListAddLine(IggDrawList handle, IggVec2 *p1, IggVec2 *p2,
                               unsigned int col, float thickness);
extern void iggDrawListAddRect(IggDrawList handle, IggVec2 *p_min,
                               IggVec2 *p_max, unsigned int col, float rounding,
                               int rounding_corners, float thickness);
extern void iggDrawListAddRectFilled(IggDrawList handle, IggVec2 *p_min,
                                     IggVec2 *p_max, unsigned int col,
                                     float rounding, int rounding_corners);
extern void iggDrawListAddText(IggDrawList handle, IggVec2 *pos,
                               unsigned int col, const char *text);
extern void iggDrawListAddBezierCubic(IggDrawList handle, IggVec2 *pos0,
                                      IggVec2 *cp0, IggVec2 *cp1, IggVec2 *pos1,
                                      unsigned int col, float thickness,
                                      int num_segments);
extern void iggDrawListAddTriangle(IggDrawList handle, IggVec2 *p1, IggVec2 *p2,
                                   IggVec2 *p3, unsigned int col,
                                   float thickness);
extern void iggDrawListAddTriangleFilled(IggDrawList handle, IggVec2 *p1,
                                         IggVec2 *p2, IggVec2 *p3,
                                         unsigned int col);
extern void iggDrawListAddCircle(IggDrawList handle, IggVec2 *center,
                                 float radius, unsigned int col,
                                 int num_segments, float thickness);
extern void iggDrawListAddCircleFilled(IggDrawList handle, IggVec2 *center,
                                       float radius, unsigned int col,
                                       int num_segments);
extern void iggDrawListAddQuad(IggDrawList handle, IggVec2 *p1, IggVec2 *p2,
                               IggVec2 *p3, IggVec2 *p4, unsigned int col,
                               float thickness);
extern void iggDrawListAddQuadFilled(IggDrawList handle, IggVec2 *p1,
                                     IggVec2 *p2, IggVec2 *p3, IggVec2 *p4,
                                     unsigned int col);

extern void iggDrawListPathClear(IggDrawList handle);
extern void iggDrawListPathLineTo(IggDrawList handle, IggVec2 *pos);
extern void iggDrawListPathLineToMergeDuplicate(IggDrawList handle,
                                                IggVec2 *pos);
extern void iggDrawListPathFillConvex(IggDrawList handle, unsigned int col);
extern void iggDrawListPathStroke(IggDrawList handle, unsigned int col,
                                  IggBool closed, float thickness);
extern void iggDrawListPathArcTo(IggDrawList handle, IggVec2 *center,
                                 float radius, float a_min, float a_max,
                                 int num_segments);
extern void iggDrawListPathArcToFast(IggDrawList handle, IggVec2 *center,
                                     float radius, int a_min_of_12,
                                     int a_max_of_12);
extern void iggDrawListPathBezierCubicCurveTo(IggDrawList handle, IggVec2 *p1,
                                         IggVec2 *p2, IggVec2 *p3,
                                         int num_segments);
extern void iggDrawListAddImage(IggDrawList handle, IggTextureID id,
                                IggVec2 *p_min, IggVec2 *p_max);
extern void iggDrawListAddImageV(IggDrawList handle, IggTextureID id,
                                IggVec2 *p_min, IggVec2 *p_max,
                                IggVec2 *uv_min, IggVec2 *uv_max,
                                unsigned int color);

#ifdef __cplusplus
}
#endif
