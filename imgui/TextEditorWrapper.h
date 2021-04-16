#pragma once

#include "imguiWrapperTypes.h"
#ifdef __cplusplus
extern "C" {
#endif

typedef void *IggTextEditor;
typedef void *IggTextEditorErrorMarkers;

extern IggTextEditor IggNewTextEditor();
extern void IggTextEditorRender(IggTextEditor handle, const char *aTitle,
                                IggVec2 const *size, int aBorder);
extern void IggTextEditorSetShowWhitespaces(IggTextEditor handle, int aValue);
extern void IggTextEditorSetTabSize(IggTextEditor handle, int size);
extern void IggTextEditorSetText(IggTextEditor handle, const char *text);
extern void IggTextEditorInsertText(IggTextEditor handle, const char *text);
extern char const *IggTextEditorGetText(IggTextEditor handle);
extern char const *IggTextEditorGetWordUnderCursor(IggTextEditor handle);
extern IggBool IggTextEditorHasSelection(IggTextEditor handle);
extern char const *IggTextEditorGetSelectedText(IggTextEditor handle);
extern char const *IggTextEditorGetCurrentLineText(IggTextEditor handle);
extern IggBool IggTextEditorIsTextChanged(IggTextEditor handle);
extern void IggTextEditorGetScreenCursorPos(IggTextEditor handle, int *x,
                                            int *y);
extern void IggTextEditorGetCursorPos(IggTextEditor handle, int *column,
                                      int *line);
extern void IggTextEditorGetSelectionStart(IggTextEditor handle, int *column,
                                           int *line);

extern void IggTextEditorCopy(IggTextEditor handle);
extern void IggTextEditorCut(IggTextEditor handle);
extern void IggTextEditorPaste(IggTextEditor handle);
extern void IggTextEditorDelete(IggTextEditor handle);

extern void IggTextEditorSelectWordUnderCursor(IggTextEditor handle);

extern void IggTextEditorSetLanguageDefinitionSQL(IggTextEditor handle);
extern void IggTextEditorSetLanguageDefinitionCPP(IggTextEditor handle);
extern void IggTextEditorSetLanguageDefinitionC(IggTextEditor handle);
extern void IggTextEditorSetLanguageDefinitionLua(IggTextEditor handle);

extern IggTextEditorErrorMarkers IggTextEditorNewErrorMarkers();
extern void IggTextEditorErrorMarkersInsert(IggTextEditorErrorMarkers handle,
                                            int pos, const char *errMsg);
extern void IggTextEditorErrorMarkersClear(IggTextEditorErrorMarkers handle);
extern unsigned int
IggTextEditorErrorMarkersSize(IggTextEditorErrorMarkers handle);

extern void IggTextEditorSetErrorMarkers(IggTextEditor handle,
                                         IggTextEditorErrorMarkers marker);

extern void IggTextEditorSetHandleKeyboardInputs(IggTextEditor handle, int aValue);

#ifdef __cplusplus
}
#endif
