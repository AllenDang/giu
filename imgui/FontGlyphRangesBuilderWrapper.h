#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern IggGlyphRanges IggNewGlyphRanges();
extern IggGlyphRanges IggGlyphRangesData(IggGlyphRanges handle);

extern IggFontGlyphRangesBuilder IggNewFontGlyphRangesBuilder();
extern void IggFontGlyphRangesBuilderClear(IggFontGlyphRangesBuilder handle);
extern void IggFontGlyphRangesBuilderAddRanges(IggFontGlyphRangesBuilder handle, IggGlyphRanges ranges);
extern void IggFontGlyphRangesBuilderAddText(IggFontGlyphRangesBuilder handle, const char* text);
extern void IggFontGlyphRangesBuilderBuildRanges(IggFontGlyphRangesBuilder handle, IggGlyphRanges ranges);

#ifdef __cplusplus 
}
#endif
