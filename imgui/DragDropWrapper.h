#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

typedef void* IggPayload;

extern void* iggPayloadGetData(IggPayload handle);

extern IggBool iggBeginDragDropSource(int flags);
extern IggBool iggSetDragDropPayload(const char* type, const void* data, unsigned int sz, int cond);
extern void iggEndDragDropSource();
extern IggBool iggBeginDragDropTarget();
extern IggPayload iggAcceptDragDropPayload(const char *type, int flags);
extern void iggEndDragDropTarget();


#ifdef __cplusplus
}
#endif
