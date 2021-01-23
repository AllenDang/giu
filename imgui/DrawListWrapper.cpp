#include "imguiWrappedHeader.h"
#include "DrawListWrapper.h"
#include "WrapperConverter.h"

int iggDrawListGetCommandCount(IggDrawList handle) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  return list->CmdBuffer.Size;
}

IggDrawCmd iggDrawListGetCommand(IggDrawList handle, int index) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  return reinterpret_cast<IggDrawCmd>(&list->CmdBuffer.Data[index]);
}

void iggDrawListGetRawIndexBuffer(IggDrawList handle, void **data,
                                  int *byteSize) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  *data = list->IdxBuffer.Data;
  *byteSize = static_cast<int>(sizeof(ImDrawIdx)) * list->IdxBuffer.Size;
}

void iggDrawListGetRawVertexBuffer(IggDrawList handle, void **data,
                                   int *byteSize) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  *data = list->VtxBuffer.Data;
  *byteSize = static_cast<int>(sizeof(ImDrawVert)) * list->VtxBuffer.Size;
}

void iggGetIndexBufferLayout(size_t *entrySize) {
  *entrySize = sizeof(ImDrawIdx);
}

void iggGetVertexBufferLayout(size_t *entrySize, size_t *posOffset,
                              size_t *uvOffset, size_t *colOffset) {
  *entrySize = sizeof(ImDrawVert);
  *posOffset = IM_OFFSETOF(ImDrawVert, pos);
  *uvOffset = IM_OFFSETOF(ImDrawVert, uv);
  *colOffset = IM_OFFSETOF(ImDrawVert, col);
}

void iggDrawListAddLine(IggDrawList handle, IggVec2 *p1, IggVec2 *p2,
                        unsigned int col, float thickness) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper p1Arg(p1);
  Vec2Wrapper p2Arg(p2);
  list->AddLine(*p1Arg, *p2Arg, col, thickness);
}

void iggDrawListAddRect(IggDrawList handle, IggVec2 *p_min, IggVec2 *p_max,
                        unsigned int col, float rounding, int rounding_corners,
                        float thickness) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper p1Arg(p_min);
  Vec2Wrapper p2Arg(p_max);
  list->AddRect(*p1Arg, *p2Arg, col, rounding, rounding_corners, thickness);
}

void iggDrawListAddRectFilled(IggDrawList handle, IggVec2 *p_min,
                              IggVec2 *p_max, unsigned int col, float rounding,
                              int rounding_corners) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper p1Arg(p_min);
  Vec2Wrapper p2Arg(p_max);
  list->AddRectFilled(*p1Arg, *p2Arg, col, rounding, rounding_corners);
}

void iggDrawListAddText(IggDrawList handle, IggVec2 *pos, unsigned int col,
                        const char *text) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper posArg(pos);
  list->AddText(*posArg, col, text);
}

void iggDrawListAddBezierCubic(IggDrawList handle, IggVec2 *pos0, IggVec2 *cp0,
                               IggVec2 *cp1, IggVec2 *pos1, unsigned int col,
                               float thickness, int num_segments) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper pos0Arg(pos0);
  Vec2Wrapper pos1Arg(pos1);
  Vec2Wrapper cp0Arg(cp0);
  Vec2Wrapper cp1Arg(cp1);
  list->AddBezierCubic(*pos0Arg, *cp0Arg, *cp1Arg, *pos1Arg, col, thickness,
                       num_segments);
}

void iggDrawListAddTriangle(IggDrawList handle, IggVec2 *p1, IggVec2 *p2,
                            IggVec2 *p3, unsigned int col, float thickness) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper p1Arg(p1);
  Vec2Wrapper p2Arg(p2);
  Vec2Wrapper p3Arg(p3);
  list->AddTriangle(*p1Arg, *p2Arg, *p3Arg, col, thickness);
}

void iggDrawListAddTriangleFilled(IggDrawList handle, IggVec2 *p1, IggVec2 *p2,
                                  IggVec2 *p3, unsigned int col) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper p1Arg(p1);
  Vec2Wrapper p2Arg(p2);
  Vec2Wrapper p3Arg(p3);
  list->AddTriangleFilled(*p1Arg, *p2Arg, *p3Arg, col);
}

void iggDrawListAddCircle(IggDrawList handle, IggVec2 *center, float radius,
                          unsigned int col, int num_segments, float thickness) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper centerArg(center);
  list->AddCircle(*centerArg, radius, col, num_segments, thickness);
}

void iggDrawListAddCircleFilled(IggDrawList handle, IggVec2 *center,
                                float radius, unsigned int col,
                                int num_segments) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper centerArg(center);
  list->AddCircleFilled(*centerArg, radius, col, num_segments);
}

void iggDrawListAddQuad(IggDrawList handle, IggVec2 *p1, IggVec2 *p2,
                        IggVec2 *p3, IggVec2 *p4, unsigned int col,
                        float thickness) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper p1Arg(p1);
  Vec2Wrapper p2Arg(p2);
  Vec2Wrapper p3Arg(p3);
  Vec2Wrapper p4Arg(p4);
  list->AddQuad(*p1Arg, *p2Arg, *p3Arg, *p4Arg, col, thickness);
}

void iggDrawListAddQuadFilled(IggDrawList handle, IggVec2 *p1, IggVec2 *p2,
                              IggVec2 *p3, IggVec2 *p4, unsigned int col) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper p1Arg(p1);
  Vec2Wrapper p2Arg(p2);
  Vec2Wrapper p3Arg(p3);
  Vec2Wrapper p4Arg(p4);
  list->AddQuadFilled(*p1Arg, *p2Arg, *p3Arg, *p4Arg, col);
}

void iggDrawListPathClear(IggDrawList handle) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  list->PathClear();
}

void iggDrawListPathLineTo(IggDrawList handle, IggVec2 *pos) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper posArg(pos);
  list->PathLineTo(*posArg);
}

void iggDrawListPathLineToMergeDuplicate(IggDrawList handle, IggVec2 *pos) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper posArg(pos);
  list->PathLineToMergeDuplicate(*posArg);
}

void iggDrawListPathFillConvex(IggDrawList handle, unsigned int col) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  list->PathFillConvex(col);
}

void iggDrawListPathStroke(IggDrawList handle, unsigned int col, IggBool closed,
                           float thickness) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  list->PathStroke(col, closed != 0, thickness);
}

void iggDrawListPathArcTo(IggDrawList handle, IggVec2 *center, float radius,
                          float a_min, float a_max, int num_segments) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper centerArg(center);
  list->PathArcTo(*centerArg, radius, a_min, a_max, num_segments);
}

void iggDrawListPathArcToFast(IggDrawList handle, IggVec2 *center, float radius,
                              int a_min_of_12, int a_max_of_12) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper centerArg(center);
  list->PathArcToFast(*centerArg, radius, a_min_of_12, a_max_of_12);
}

void iggDrawListPathBezierCubicCurveTo(IggDrawList handle, IggVec2 *p1, IggVec2 *p2,
                                  IggVec2 *p3, int num_segments) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper p1Arg(p1);
  Vec2Wrapper p2Arg(p2);
  Vec2Wrapper p3Arg(p3);
  list->PathBezierCubicCurveTo(*p1Arg, *p2Arg, *p3Arg, num_segments);
}

void iggDrawListAddImage(IggDrawList handle, IggTextureID id, IggVec2 *p_min,
                         IggVec2 *p_max) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper pMinArg(p_min);
  Vec2Wrapper pMaxArg(p_max);
  list->AddImage(ImTextureID(id), *pMinArg, *pMaxArg);
}
void iggDrawListAddImageV(IggDrawList handle, IggTextureID id, IggVec2 *p_min,
                         IggVec2 *p_max, IggVec2 *uv_min, IggVec2 *uv_max, unsigned int color) {
  ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
  Vec2Wrapper pMinArg(p_min);
  Vec2Wrapper pMaxArg(p_max);
  Vec2Wrapper uvMinArg(uv_min);
  Vec2Wrapper uvMaxArg(uv_max);
  list->AddImage(ImTextureID(id), *pMinArg, *pMaxArg, *uvMinArg, *uvMaxArg, color);
}
