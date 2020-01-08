#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern int iggDrawListGetCommandCount(IggDrawList handle);
extern IggDrawCmd iggDrawListGetCommand(IggDrawList handle, int index);
extern void iggDrawListGetRawIndexBuffer(IggDrawList handle, void **data, int *byteSize);
extern void iggDrawListGetRawVertexBuffer(IggDrawList handle, void **data, int *byteSize);

extern void iggGetIndexBufferLayout(size_t *entrySize);
extern void iggGetVertexBufferLayout(size_t *entrySize, size_t *posOffset, size_t *uvOffset, size_t *colOffset);

extern void iggDrawListAddLine(IggDrawList handle, IggVec2 *p1, IggVec2 *p2, unsigned int col, float thickness);

#ifdef __cplusplus
}
#endif
