#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

	extern IggContext iggCreateContext(IggFontAtlas sharedFontAtlas);
	extern void iggDestroyContext(IggContext context);
	extern IggContext iggGetCurrentContext();
	extern void iggSetCurrentContext(IggContext context);
	extern void iggSetMaxWaitBeforeNextFrame(double time);

	extern IggIO iggGetCurrentIO(void);
	extern IggGuiStyle iggGetCurrentStyle(void);
	extern void iggNewFrame(void);
	extern void iggRender(void);
	extern IggDrawData iggGetDrawData(void);
	extern void iggEndFrame(void);

	extern double iggGetEventWaitingTime(void);

	extern char const *iggGetVersion(void);
	extern void iggShowDemoWindow(IggBool *open);
	extern void iggShowUserGuide(void);

	extern IggBool iggBegin(char const *id, IggBool *open, int flags);
	extern void iggEnd(void);
	extern IggBool iggBeginChild(char const *id, IggVec2 const *size, IggBool border, int flags);
	extern void iggEndChild(void);

	extern void iggWindowPos(IggVec2 *pos);
	extern void iggWindowSize(IggVec2 *size);
	extern float iggWindowWidth(void);
	extern float iggWindowHeight(void);
	extern void iggContentRegionAvail(IggVec2 *size);

  extern IggBool iggIsWindowAppearing();

	extern void iggSetNextWindowPos(IggVec2 const *pos, int cond, IggVec2 const *pivot);
	extern void iggSetNextWindowSize(IggVec2 const *size, int cond);
	extern void iggSetNextWindowContentSize(IggVec2 const *size);
	extern void iggSetNextWindowFocus(void);
	extern void iggSetNextWindowBgAlpha(float value);

	extern void iggPushFont(IggFont handle);
	extern void iggPopFont(void);
	extern void iggPushStyleColor(int index, IggVec4 const *col);
	extern void iggPopStyleColor(int count);
	extern void iggPushStyleVarFloat(int index, float value);
	extern void iggPushStyleVarVec2(int index, IggVec2 const *value);
	extern void iggPopStyleVar(int count);
	extern void iggCalcTextSize(const char *text, int length, IggBool hide_text_after_double_hash, float wrap_width, IggVec2 *value);
	extern unsigned int iggGetColorU32(const IggVec4 col);

  extern void iggStyleColorsDark();
  extern void iggStyleColorsClassic();
  extern void iggStyleColorsLight();

	extern float iggGetFontSize();

	extern void iggPushItemWidth(float width);
	extern void iggPopItemWidth(void);
	extern float iggCalcItemWidth(void);
	extern void iggPushTextWrapPos(float wrapPosX);
	extern void iggPopTextWrapPos(void);

	extern void iggPushID(char const *id);
	extern void iggPopID(void);

	extern void iggTextUnformatted(char const *text);
	extern void iggLabelText(char const *label, char const *text);

	extern IggBool iggButton(char const *label, IggVec2 const *size);
  extern IggBool iggInvisibleButton(char const *label, IggVec2 const *size);
	extern void iggImage(IggTextureID textureID,
						 IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
						 IggVec4 const *tintCol, IggVec4 const *borderCol);
	extern IggBool iggImageButton(IggTextureID textureID,
								  IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
								  int framePadding, IggVec4 const *bgCol,
								  IggVec4 const *tintCol);
	extern IggBool iggCheckbox(char const *label, IggBool *selected);
  extern IggBool iggRadioButton(char const *label, IggBool active);
	extern void iggProgressBar(float fraction, IggVec2 const *size, char const *overlay);

	extern IggBool iggBeginCombo(char const *label, char const *previewValue, int flags);
	extern void iggEndCombo(void);

	extern IggBool iggDragFloat(char const *label, float *value, float speed, float min, float max, char const *format, float power);
	extern IggBool iggDragInt(char const *label, int *value, float speed, int min, int max, char const *format);

	extern IggBool iggSliderFloat(char const *label, float *value, float minValue, float maxValue, char const *format, float power);
	extern IggBool iggSliderFloatN(char const *label, float *value, int n, float minValue, float maxValue, char const *format, float power);

	extern IggBool iggSliderInt(char const *label, int *value, int minValue, int maxValue, char const *format);

  extern IggBool iggInputText(char const* label, char* buf, unsigned int bufSize, int flags, int callbackKey);
  extern IggBool iggInputTextMultiline(char const* label, char* buf, unsigned int bufSize, IggVec2 const *size, int flags, int callbackKey);
  
  extern IggBool iggInputInt(char const* label, int* v, int step, int step_fast, int flags);
  extern IggBool iggInputFloat(char const* label, float* v, float step, float step_fast, const char* format, int flats);

  extern IggBool iggColorEdit3(char const *label, float *col, int flags);
  extern IggBool iggColorEdit4(char const *label, float *col, int flags);
  extern IggBool iggColorPicker3(char const *label, float *col, int flags);
  extern IggBool iggColorPicker4(char const *label, float *col, int flags);

	extern void iggSeparator(void);
	extern void iggSameLine(float posX, float spacingW);
	extern void iggSpacing(void);
	extern void iggDummy(IggVec2 const *size);
	extern void iggBeginGroup(void);
	extern void iggEndGroup(void);

	extern void iggCursorPos(IggVec2 *pos);
	extern float iggCursorPosX(void);
	extern float iggCursorPosY(void);
	extern void iggCursorStartPos(IggVec2 *pos);
	extern void iggCursorScreenPos(IggVec2 *pos);

	extern void iggSetCursorPos(IggVec2 const *localPos);
	extern void iggSetCursorScreenPos(IggVec2 const *absPos);
	extern void iggAlignTextToFramePadding();
	extern float iggGetTextLineHeight(void);
	extern float iggGetTextLineHeightWithSpacing(void);

	extern IggBool iggTreeNode(char const *label, int flags);
	extern void iggTreePop(void);
	extern void iggSetNextItemOpen(IggBool open, int cond);
	extern float iggGetTreeNodeToLabelSpacing(void);

	extern IggBool iggSelectable(char const *label, IggBool selected, int flags, IggVec2 const *size);
	extern IggBool iggListBoxV(char const *label, int *currentItem, char const *const items[], int itemCount, int heightItems);

	extern void iggPlotLines(const char *label, const float *values, int valuesCount, int valuesOffset, const char *overlayText, float scaleMin, float scaleMax, IggVec2 const *graphSize);
	extern void iggPlotHistogram(const char *label, const float *values, int valuesCount, int valuesOffset, const char *overlayText, float scaleMin, float scaleMax, IggVec2 const *graphSize);

	extern void iggSetTooltip(char const *text);
	extern void iggBeginTooltip(void);
	extern void iggEndTooltip(void);

	extern IggBool iggBeginMainMenuBar(void);
	extern void iggEndMainMenuBar(void);
	extern IggBool iggBeginMenuBar(void);
	extern void iggEndMenuBar(void);
	extern IggBool iggBeginMenu(char const *label, IggBool enabled);
	extern void iggEndMenu(void);
	extern IggBool iggMenuItem(char const *label, char const *shortcut, IggBool selected, IggBool enabled);

	extern void iggOpenPopup(char const *id);
	extern IggBool iggBeginPopupModal(char const *name, IggBool *open, int flags);
	extern IggBool iggBeginPopupContextItem(char const *label, int mouseButton);
	extern void iggEndPopup(void);
	extern void iggCloseCurrentPopup(void);

	extern IggBool iggIsItemHovered(int flags);
  extern IggBool iggIsItemActive();
  extern IggBool iggIsAnyItemActive();

	extern IggBool iggIsKeyDown(int key);
	extern IggBool iggIsKeyPressed(int key, IggBool repeat);
	extern IggBool iggIsKeyReleased(int key);
	extern IggBool iggIsMouseDown(int button);
	extern IggBool iggIsAnyMouseDown();
	extern IggBool iggIsMouseClicked(int button, IggBool repeat);
	extern IggBool iggIsMouseReleased(int button);
	extern IggBool iggIsMouseDoubleClicked(int button);

	extern void iggColumns(int count, char const *label, IggBool border);
	extern void iggNextColumn();
	extern int iggGetColumnIndex();
	extern int iggGetColumnWidth(int index);
	extern void iggSetColumnWidth(int index, float width);
	extern float iggGetColumnOffset(int index);
	extern void iggSetColumnOffset(int index, float offsetX);
	extern int iggGetColumnsCount();
	extern void iggSetScrollHereY(float centerYRatio);

	extern void iggSetItemDefaultFocus();
	extern IggBool iggIsItemFocused();
	extern IggBool iggIsAnyItemFocused();
	extern int iggGetMouseCursor();
	extern void iggSetMouseCursor(int cursor);
  extern void iggSetKeyboardFocusHere(int offset);

	extern IggBool iggBeginTabBar(char const *str_id, int flags);
	extern void iggEndTabBar();
	extern IggBool iggBeginTabItem(char const *label, IggBool *p_open, int flags);
	extern void iggEndTabItem();
	extern void iggSetTabItemClosed(char const *tab_or_docked_window_label);

  extern IggDrawList iggGetWindowDrawList();

  extern void iggGetItemRectMin(IggVec2 *size);
  extern void iggGetItemRectMax(IggVec2 *size);
  extern void iggGetItemRectSize(IggVec2 *size);


#ifdef __cplusplus
}
#endif
