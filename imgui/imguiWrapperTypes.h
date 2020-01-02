#pragma once

#ifdef __cplusplus
extern "C"
{
#endif

typedef int IggBool;
typedef void *IggTextureID;

typedef void *IggContext;
typedef void *IggDrawCmd;
typedef void *IggDrawData;
typedef void *IggDrawList;
typedef void *IggFontAtlas;
typedef void *IggFontConfig;
typedef void *IggFont;
typedef void *IggGlyphRanges;
typedef void *IggGuiStyle;
typedef void *IggInputTextCallbackData;
typedef void *IggIO;
typedef void *IggFontGlyphRangesBuilder;

typedef struct tagIggVec2
{
   float x;
   float y;
} IggVec2;

typedef struct tagIggVec4
{
   float x;
   float y;
   float z;
   float w;
} IggVec4;

#ifdef __cplusplus
}
#endif
