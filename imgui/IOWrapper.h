#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

    extern IggBool iggWantCaptureMouse(IggIO handle);
    extern IggBool iggWantCaptureKeyboard(IggIO handle);
    extern IggBool iggWantTextInput(IggIO handle);

    extern IggFontAtlas iggIoGetFonts(IggIO handle);

    extern void iggIoSetDisplaySize(IggIO handle, IggVec2 const *value);
    extern void iggIoSetMousePosition(IggIO handle, IggVec2 const *value);
    extern void iggIoSetMouseButtonDown(IggIO handle, int index, IggBool value);
    extern void iggIoAddMouseWheelDelta(IggIO handle, float x, float y);
    extern void iggIoGetMouseDelta(IggIO handle, IggVec2 *delta);
    extern void iggIoSetDeltaTime(IggIO handle, float value);
    extern void iggIoSetFontGlobalScale(IggIO handle, float value);

    extern IggBool iggIoGetMouseDrawCursor(IggIO handle);
    extern void iggIoSetMouseDrawCursor(IggIO handle, IggBool value);

    extern void iggIoKeyPress(IggIO handle, int key);
    extern void iggIoKeyRelease(IggIO handle, int key);
    extern void iggIoKeyMap(IggIO handle, int imguiKey, int nativeKey);
    extern void iggIoKeyCtrl(IggIO handle, int leftCtrl, int rightCtrl);
    extern void iggIoKeyShift(IggIO handle, int leftShift, int rightShift);
    extern void iggIoKeyAlt(IggIO handle, int leftAlt, int rightAlt);
    extern void iggIoKeySuper(IggIO handle, int leftSuper, int rightSuper);
    extern void iggIoAddInputCharactersUTF8(IggIO handle, char const *utf8Chars);
    extern void iggIoSetIniFilename(IggIO handle, char const *value);
    extern void iggIoSetConfigFlags(IggIO handle, int flags);
    extern int iggIoGetConfigFlags(IggIO handle);
    extern void iggIoSetBackendFlags(IggIO handle, int flags);

    extern void iggIoRegisterClipboardFunctions(IggIO handle);
    extern void iggIoClearClipboardFunctions(IggIO handle);

    extern int iggGetFrameCountSinceLastInput(IggIO handle);
    extern void iggSetFrameCountSinceLastInput(IggIO handle, int count);

#ifdef __cplusplus
}
#endif
