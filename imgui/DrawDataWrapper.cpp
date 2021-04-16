#include "imguiWrappedHeader.h"
#include "DrawDataWrapper.h"
#include "WrapperConverter.h"

IggBool iggDrawDataValid(IggDrawData handle)
{
   ImDrawData *drawData = reinterpret_cast<ImDrawData *>(handle);
   IggBool result = 0;
   exportValue(result, drawData->Valid);
   return result;
}

void iggDrawDataGetCommandLists(IggDrawData handle, void **handles, int *count)
{
   ImDrawData *drawData = reinterpret_cast<ImDrawData *>(handle);
   *handles = reinterpret_cast<void **>(drawData->CmdLists);
   *count = drawData->CmdListsCount;
}

void iggDrawDataScaleClipRects(IggDrawData handle, IggVec2 const *scale)
{
   ImDrawData *drawData = reinterpret_cast<ImDrawData *>(handle);
   Vec2Wrapper wrappedScale(scale);
   drawData->ScaleClipRects(*wrappedScale);
}
