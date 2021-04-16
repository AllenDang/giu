#include "imguiWrappedHeader.h" 
#include "FontGlyphRangesBuilderWrapper.h"

IggGlyphRanges IggNewGlyphRanges() {
  ImVector<ImWchar> *ranges = new ImVector<ImWchar>();
  return static_cast<IggGlyphRanges>(ranges);
}

IggGlyphRanges IggGlyphRangesData(IggGlyphRanges handle) {
  ImVector<ImWchar> *ranges = reinterpret_cast<ImVector<ImWchar>*>(handle); 
  return static_cast<IggGlyphRanges>(ranges->Data);
}

IggFontGlyphRangesBuilder IggNewFontGlyphRangesBuilder()
{
  ImFontGlyphRangesBuilder *builder = new ImFontGlyphRangesBuilder();
  return static_cast<IggFontGlyphRangesBuilder>(builder);
}

void IggFontGlyphRangesBuilderClear(IggFontGlyphRangesBuilder handle)
{
  ImFontGlyphRangesBuilder *builder = reinterpret_cast<ImFontGlyphRangesBuilder*>(handle);
  builder->Clear();
}

void IggFontGlyphRangesBuilderAddText(IggFontGlyphRangesBuilder handle, const char* text)
{
  ImFontGlyphRangesBuilder *builder = reinterpret_cast<ImFontGlyphRangesBuilder*>(handle);
  builder->AddText(text);
}

void IggFontGlyphRangesBuilderAddRanges(IggFontGlyphRangesBuilder handle, IggGlyphRanges ranges)
{
  ImFontGlyphRangesBuilder *builder = reinterpret_cast<ImFontGlyphRangesBuilder*>(handle);
  builder->AddRanges(reinterpret_cast<ImWchar*>(ranges));
}

void IggFontGlyphRangesBuilderBuildRanges(IggFontGlyphRangesBuilder handle, IggGlyphRanges ranges)
{
  ImFontGlyphRangesBuilder *builder = reinterpret_cast<ImFontGlyphRangesBuilder*>(handle);
  builder->BuildRanges(reinterpret_cast<ImVector<ImWchar>*>(ranges));
}

