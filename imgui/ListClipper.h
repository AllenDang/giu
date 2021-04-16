#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

typedef struct tagIggListClipper
{
   float StartPosY;
   float ItemsHeight;
   int ItemsCount;
   int StepNo;
   int DisplayStart;
   int DisplayEnd;
} IggListClipper;

extern IggBool iggListClipperStep(IggListClipper *clipper);
extern void iggListClipperBegin(IggListClipper *clipper, int items_count, float items_height);
extern void iggListClipperEnd(IggListClipper *clipper);

#ifdef __cplusplus
}
#endif
