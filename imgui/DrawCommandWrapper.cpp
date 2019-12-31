#include "imguiWrappedHeader.h"
#include "DrawCommandWrapper.h"
#include "WrapperConverter.h"

void iggDrawCommandGetElementCount(IggDrawCmd handle, unsigned int *count)
{
   ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
   *count = cmd->ElemCount;
}

void iggDrawCommandGetClipRect(IggDrawCmd handle, IggVec4 *rect)
{
   ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
   exportValue(*rect, cmd->ClipRect);
}

void iggDrawCommandGetTextureID(IggDrawCmd handle, IggTextureID *id)
{
   ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
   *id = cmd->TextureId;
}

IggBool iggDrawCommandHasUserCallback(IggDrawCmd handle)
{
   ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
   return (cmd->UserCallback != 0) ? 1 : 0;
}

void iggDrawCommandCallUserCallback(IggDrawCmd handle, IggDrawList listHandle)
{
   ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
   ImDrawList *list = reinterpret_cast<ImDrawList *>(listHandle);
   cmd->UserCallback(list, cmd);
}

