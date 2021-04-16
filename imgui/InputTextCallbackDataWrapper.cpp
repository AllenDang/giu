#include "imguiWrappedHeader.h"
#include "InputTextCallbackDataWrapper.h"

int iggInputTextCallbackDataGetEventFlag(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->EventFlag;
}

int iggInputTextCallbackDataGetFlags(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->Flags;
}

unsigned short iggInputTextCallbackDataGetEventChar(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->EventChar;
}

void iggInputTextCallbackDataSetEventChar(IggInputTextCallbackData handle, unsigned short value)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   data->EventChar = value;
}

int iggInputTextCallbackDataGetEventKey(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->EventKey;
}

char *iggInputTextCallbackDataGetBuf(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->Buf;
}

void iggInputTextCallbackDataSetBuf(IggInputTextCallbackData handle, char *buf, int size, int textLen)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   data->Buf = buf;
   data->BufSize = size;
   data->BufTextLen = textLen;
   data->BufDirty = true;
}

void iggInputTextCallbackDataMarkBufferModified(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   data->BufDirty = true;
}

int iggInputTextCallbackDataGetBufSize(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->BufSize;
}

int iggInputTextCallbackDataGetBufTextLen(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->BufTextLen;
}

void iggInputTextCallbackDataDeleteBytes(IggInputTextCallbackData handle, int offset, int count)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   data->DeleteChars(offset, count);
}

void iggInputTextCallbackDataInsertBytes(IggInputTextCallbackData handle, int offset, char *bytes, int count)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   data->InsertChars(offset, bytes, bytes+count);
}

int iggInputTextCallbackDataGetCursorPos(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->CursorPos;
}

void iggInputTextCallbackDataSetCursorPos(IggInputTextCallbackData handle, int value)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   data->CursorPos = value;
}

int iggInputTextCallbackDataGetSelectionStart(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->SelectionStart;
}

void iggInputTextCallbackDataSetSelectionStart(IggInputTextCallbackData handle, int value)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   data->SelectionStart = value;
}

int iggInputTextCallbackDataGetSelectionEnd(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->SelectionEnd;
}

void iggInputTextCallbackDataSetSelectionEnd(IggInputTextCallbackData handle, int value)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   data->SelectionEnd = value;
}
