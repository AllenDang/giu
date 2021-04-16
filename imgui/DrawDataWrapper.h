#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern IggBool iggDrawDataValid(IggDrawData handle);
extern void iggDrawDataGetCommandLists(IggDrawData handle, void **handles, int *count);
extern void iggDrawDataScaleClipRects(IggDrawData handle, IggVec2 const *scale);

#ifdef __cplusplus
}
#endif
