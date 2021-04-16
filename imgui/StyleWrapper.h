#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern void iggStyleGetItemSpacing(IggGuiStyle handle, IggVec2 *value);

extern void iggStyleGetItemInnerSpacing(IggGuiStyle handle, IggVec2 *value);

extern void iggStyleGetWindowPadding(IggGuiStyle handle, IggVec2 *value);

extern void iggStyleGetFramePadding(IggGuiStyle handle, IggVec2 *value);

extern void iggStyleSetColor(IggGuiStyle handle, int index, IggVec4 const *color);

extern void iggStyleGetColor(IggGuiStyle handle, int index, IggVec4 *color);

extern void iggStyleScaleAllSizes(IggGuiStyle handle, float scale);

#ifdef __cplusplus
}
#endif
