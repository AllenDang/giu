#include "imguiWrappedHeader.h"
#include "DrawListWrapper.h"
#include "WrapperConverter.h"

int iggDrawListGetCommandCount(IggDrawList handle)
{
   ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
   return list->CmdBuffer.Size;
}

IggDrawCmd iggDrawListGetCommand(IggDrawList handle, int index)
{
   ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
   return reinterpret_cast<IggDrawCmd>(&list->CmdBuffer.Data[index]);
}

void iggDrawListGetRawIndexBuffer(IggDrawList handle, void **data, int *byteSize)
{
   ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
   *data = list->IdxBuffer.Data;
   *byteSize = static_cast<int>(sizeof(ImDrawIdx)) * list->IdxBuffer.Size;
}

void iggDrawListGetRawVertexBuffer(IggDrawList handle, void **data, int *byteSize)
{
   ImDrawList *list = reinterpret_cast<ImDrawList *>(handle);
   *data = list->VtxBuffer.Data;
   *byteSize = static_cast<int>(sizeof(ImDrawVert)) * list->VtxBuffer.Size;
}

void iggGetIndexBufferLayout(size_t *entrySize)
{
   *entrySize = sizeof(ImDrawIdx);
}

void iggGetVertexBufferLayout(size_t *entrySize, size_t *posOffset, size_t *uvOffset, size_t *colOffset)
{
   *entrySize = sizeof(ImDrawVert);
   *posOffset = IM_OFFSETOF(ImDrawVert, pos);
   *uvOffset = IM_OFFSETOF(ImDrawVert, uv);
   *colOffset = IM_OFFSETOF(ImDrawVert, col);
}
