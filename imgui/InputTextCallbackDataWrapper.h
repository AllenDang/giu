#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern int iggInputTextCallbackDataGetEventFlag(IggInputTextCallbackData handle);
extern int iggInputTextCallbackDataGetFlags(IggInputTextCallbackData handle);

extern unsigned short iggInputTextCallbackDataGetEventChar(IggInputTextCallbackData handle);
extern void iggInputTextCallbackDataSetEventChar(IggInputTextCallbackData handle, unsigned short value);
extern int iggInputTextCallbackDataGetEventKey(IggInputTextCallbackData handle);

extern char *iggInputTextCallbackDataGetBuf(IggInputTextCallbackData handle);
extern void iggInputTextCallbackDataSetBuf(IggInputTextCallbackData handle, char *buf, int size, int textLen);
extern void iggInputTextCallbackDataMarkBufferModified(IggInputTextCallbackData handle);
extern int iggInputTextCallbackDataGetBufSize(IggInputTextCallbackData handle);
extern int iggInputTextCallbackDataGetBufTextLen(IggInputTextCallbackData handle);
extern void iggInputTextCallbackDataDeleteBytes(IggInputTextCallbackData handle, int offset, int count);
extern void iggInputTextCallbackDataInsertBytes(IggInputTextCallbackData handle, int offset, char *bytes, int count);

extern int iggInputTextCallbackDataGetCursorPos(IggInputTextCallbackData handle);
extern void iggInputTextCallbackDataSetCursorPos(IggInputTextCallbackData handle, int value);
extern int iggInputTextCallbackDataGetSelectionStart(IggInputTextCallbackData handle);
extern void iggInputTextCallbackDataSetSelectionStart(IggInputTextCallbackData handle, int value);
extern int iggInputTextCallbackDataGetSelectionEnd(IggInputTextCallbackData handle);
extern void iggInputTextCallbackDataSetSelectionEnd(IggInputTextCallbackData handle, int value);

#ifdef __cplusplus
}
#endif
