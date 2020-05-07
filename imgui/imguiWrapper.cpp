#include "imguiWrappedHeader.h"
#include "imguiWrapper.h"
#include "WrapperConverter.h"

IggContext iggCreateContext(IggFontAtlas sharedFontAtlas)
{
   ImGuiContext *context = ImGui::CreateContext(reinterpret_cast<ImFontAtlas *>(sharedFontAtlas));
   return reinterpret_cast<IggContext>(context);
}

void iggDestroyContext(IggContext context)
{
   ImGui::DestroyContext(reinterpret_cast<ImGuiContext *>(context));
}

IggContext iggGetCurrentContext()
{
   return reinterpret_cast<IggContext>(ImGui::GetCurrentContext());
}

void iggSetCurrentContext(IggContext context)
{
   ImGui::SetCurrentContext(reinterpret_cast<ImGuiContext *>(context));
}

void iggSetMaxWaitBeforeNextFrame(double time)
{
   ImGui::SetMaxWaitBeforeNextFrame(time);
}

IggIO iggGetCurrentIO()
{
   return reinterpret_cast<IggIO>(&ImGui::GetIO());
}

IggGuiStyle iggGetCurrentStyle()
{
   return reinterpret_cast<IggGuiStyle>(&ImGui::GetStyle());
}

void iggNewFrame()
{
   ImGui::NewFrame();
}

void iggRender()
{
   ImGui::Render();
}

IggDrawData iggGetDrawData()
{
   return reinterpret_cast<IggDrawData>(ImGui::GetDrawData());
}

void iggEndFrame()
{
   ImGui::EndFrame();
}

double iggGetEventWaitingTime()
{
   return ImGui::GetEventWaitingTime();
}

char const *iggGetVersion()
{
   return ImGui::GetVersion();
}

void iggShowDemoWindow(IggBool *open)
{
   BoolWrapper openArg(open);

   ImGui::ShowDemoWindow(openArg);
}

void iggShowUserGuide(void)
{
   ImGui::ShowUserGuide();
}

IggBool iggBegin(char const *id, IggBool *open, int flags)
{
   BoolWrapper openArg(open);
   return ImGui::Begin(id, openArg, flags) ? 1 : 0;
}

void iggEnd(void)
{
   ImGui::End();
}

IggBool iggBeginChild(char const *id, IggVec2 const *size, IggBool border, int flags)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::BeginChild(id, *sizeArg, border, flags) ? 1 : 0;
}

void iggEndChild(void)
{
   ImGui::EndChild();
}

void iggWindowPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetWindowPos());
}

void iggWindowSize(IggVec2 *size)
{
   exportValue(*size, ImGui::GetWindowSize());
}

float iggWindowWidth(void)
{
   return ImGui::GetWindowWidth();
}

float iggWindowHeight(void)
{
   return ImGui::GetWindowHeight();
}

void iggContentRegionAvail(IggVec2 *size)
{
   exportValue(*size, ImGui::GetContentRegionAvail());
}

IggBool iggIsWindowAppearing()
{
  return ImGui::IsWindowAppearing() ? 1: 0;
}

void iggSetNextWindowPos(IggVec2 const *pos, int cond, IggVec2 const *pivot)
{
   Vec2Wrapper posArg(pos);
   Vec2Wrapper pivotArg(pivot);
   ImGui::SetNextWindowPos(*posArg, cond, *pivotArg);
}

void iggSetNextWindowSize(IggVec2 const *size, int cond)
{
   Vec2Wrapper sizeArg(size);
   ImGui::SetNextWindowSize(*sizeArg, cond);
}

void iggSetNextWindowContentSize(IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   ImGui::SetNextWindowContentSize(*sizeArg);
}

void iggSetNextWindowFocus(void)
{
   ImGui::SetNextWindowFocus();
}

void iggSetNextWindowBgAlpha(float value)
{
   ImGui::SetNextWindowBgAlpha(value);
}

void iggPushFont(IggFont handle)
{
   ImFont *font = reinterpret_cast<ImFont *>(handle);
   ImGui::PushFont(font);
}

void iggPopFont(void)
{
   ImGui::PopFont();
}

void iggPushStyleColor(int index, IggVec4 const *col)
{
   Vec4Wrapper colArg(col);
   ImGui::PushStyleColor(index, *colArg);
}

void iggPopStyleColor(int count)
{
   ImGui::PopStyleColor(count);
}

void iggPushStyleVarFloat(int index, float value)
{
   ImGui::PushStyleVar(index, value);
}

void iggPushStyleVarVec2(int index, IggVec2 const *value)
{
   Vec2Wrapper valueArg(value);
   ImGui::PushStyleVar(index, *valueArg);
}

void iggPopStyleVar(int count)
{
   ImGui::PopStyleVar(count);
}

float iggGetFontSize()
{
   return ImGui::GetFontSize();
}

void iggCalcTextSize(const char *text, int length, IggBool hide_text_after_double_hash, float wrap_width, IggVec2 *value)
{
   exportValue(*value, ImGui::CalcTextSize(text, text + length, hide_text_after_double_hash, wrap_width));
}

unsigned int iggGetColorU32(const IggVec4 col)
{
   Vec4Wrapper colArg(&col);
   return ImGui::GetColorU32(*colArg);
}

void iggStyleColorsDark()
{
  ImGui::StyleColorsDark();
}

void iggStyleColorsClassic()
{
  ImGui::StyleColorsClassic();
}

void iggStyleColorsLight()
{
  ImGui::StyleColorsLight();
}

void iggPushItemWidth(float width)
{
   ImGui::PushItemWidth(width);
}

void iggPopItemWidth(void)
{
   ImGui::PopItemWidth();
}

float iggCalcItemWidth(void)
{
   return ImGui::CalcItemWidth();
}

void iggPushTextWrapPos(float wrapPosX)
{
   ImGui::PushTextWrapPos(wrapPosX);
}

void iggPopTextWrapPos(void)
{
   ImGui::PopTextWrapPos();
}

void iggPushID(char const *id)
{
   ImGui::PushID(id);
}
void iggPopID(void)
{
   ImGui::PopID();
}

void iggTextUnformatted(char const *text)
{
   ImGui::TextUnformatted(text);
}

void iggLabelText(char const *label, char const *text)
{
   ImGui::LabelText(label, "%s", text);
}

IggBool iggButton(char const *label, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::Button(label, *sizeArg) ? 1 : 0;
}

IggBool iggInvisibleButton(char const *label, IggVec2 const *size)
{
  Vec2Wrapper sizeArg(size);
  return ImGui::InvisibleButton(label, *sizeArg) ? 1 : 0;
}

void iggImage(IggTextureID textureID,
              IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
              IggVec4 const *tintCol, IggVec4 const *borderCol)
{
   Vec2Wrapper sizeArg(size);
   Vec2Wrapper uv0Arg(uv0);
   Vec2Wrapper uv1Arg(uv1);
   Vec4Wrapper tintColArg(tintCol);
   Vec4Wrapper borderColArg(borderCol);
   ImGui::Image(static_cast<ImTextureID>(textureID), *sizeArg, *uv0Arg, *uv1Arg, *tintColArg, *borderColArg);
}

IggBool iggImageButton(IggTextureID textureID,
                       IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
                       int framePadding, IggVec4 const *bgCol,
                       IggVec4 const *tintCol)
{
   Vec2Wrapper sizeArg(size);
   Vec2Wrapper uv0Arg(uv0);
   Vec2Wrapper uv1Arg(uv1);
   Vec4Wrapper bgColArg(bgCol);
   Vec4Wrapper tintColArg(tintCol);
   return ImGui::ImageButton(static_cast<ImTextureID>(textureID), *sizeArg, *uv0Arg, *uv1Arg, framePadding, *bgColArg, *tintColArg) ? 1 : 0;
}

IggBool iggCheckbox(char const *label, IggBool *selected)
{
   BoolWrapper selectedArg(selected);
   return ImGui::Checkbox(label, selectedArg) ? 1 : 0;
}

IggBool iggRadioButton(char const *label, IggBool active)
{
  return ImGui::RadioButton(label, active != 0) ? 1 : 0;
}

void iggProgressBar(float fraction, IggVec2 const *size, char const *overlay)
{
   Vec2Wrapper sizeArg(size);
   ImGui::ProgressBar(fraction, *sizeArg, overlay);
}

IggBool iggBeginCombo(char const *label, char const *previewValue, int flags)
{
   return ImGui::BeginCombo(label, previewValue, flags) ? 1 : 0;
}

void iggEndCombo(void)
{
   ImGui::EndCombo();
}

IggBool iggDragFloat(char const *label, float *value, float speed, float min, float max, char const *format, float power)
{
   return ImGui::DragFloat(label, value, speed, min, max, format, power) ? 1 : 0;
}

IggBool iggDragInt(char const *label, int *value, float speed, int min, int max, char const *format)
{
   return ImGui::DragInt(label, value, speed, min, max, format) ? 1 : 0;
}

IggBool iggSliderFloat(char const *label, float *value, float minValue, float maxValue, char const *format, float power)
{
   return ImGui::SliderFloat(label, value, minValue, maxValue, format, power) ? 1 : 0;
}

IggBool iggSliderFloatN(char const *label, float *value, int n, float minValue, float maxValue, char const *format, float power)
{
   return ImGui::SliderScalarN(label, ImGuiDataType_Float, (void *)value, n, &minValue, &maxValue, format, power) ? 1 : 0;
}

IggBool iggSliderInt(char const *label, int *value, int minValue, int maxValue, char const *format)
{
   return ImGui::SliderInt(label, value, minValue, maxValue, format) ? 1 : 0;
}

extern "C" int iggInputTextCallback(IggInputTextCallbackData data, int key);

static int iggInputTextCallbackWrapper(ImGuiInputTextCallbackData *data)
{
   return iggInputTextCallback(reinterpret_cast<IggInputTextCallbackData>(data), static_cast<int>(reinterpret_cast<size_t>(data->UserData)));
}

IggBool iggInputText(char const *label, char *buf, unsigned int bufSize, int flags, int callbackKey)
{
   return ImGui::InputText(label, buf, static_cast<size_t>(bufSize), flags,
                           iggInputTextCallbackWrapper, reinterpret_cast<void *>(callbackKey))
              ? 1
              : 0;
}

IggBool iggInputTextMultiline(char const *label, char *buf, unsigned int bufSize, IggVec2 const *size, int flags, int callbackKey)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::InputTextMultiline(label, buf, static_cast<size_t>(bufSize), *sizeArg, flags,
                                    iggInputTextCallbackWrapper, reinterpret_cast<void *>(callbackKey)) ? 1 : 0;
}

IggBool iggInputInt(char const *label, int* v, int step, int step_fast, int flags)
{
  return ImGui::InputInt(label, v, step, step_fast, flags) ? 1 : 0;
}

IggBool iggInputFloat(char const *label, float* v, float step, float step_fast, const char* format, int flags)
{
  return ImGui::InputFloat(label, v, step, step_fast, format, flags) ? 1 : 0;
}

IggBool iggColorEdit3(char const *label, float *col, int flags)
{
   return ImGui::ColorEdit3(label, col, flags) ? 1 : 0;
}

IggBool iggColorEdit4(char const *label, float *col, int flags)
{
   return ImGui::ColorEdit4(label, col, flags) ? 1 : 0;
}

IggBool iggColorPicker3(char const *label, float *col, int flags)
{
   return ImGui::ColorPicker3(label, col, flags) ? 1 : 0;
}

IggBool iggColorPicker4(char const *label, float *col, int flags)
{
   return ImGui::ColorPicker4(label, col, flags) ? 1 : 0;
}

void iggSeparator(void)
{
   ImGui::Separator();
}

void iggSameLine(float posX, float spacingW)
{
   ImGui::SameLine(posX, spacingW);
}

void iggSpacing(void)
{
   ImGui::Spacing();
}

void iggDummy(IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   ImGui::Dummy(*sizeArg);
}

void iggBeginGroup(void)
{
   ImGui::BeginGroup();
}

void iggEndGroup(void)
{
   ImGui::EndGroup();
}

void iggCursorPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetCursorPos());
}

float iggCursorPosX(void)
{
   return ImGui::GetCursorPosX();
}

float iggCursorPosY(void)
{
   return ImGui::GetCursorPosY();
}

void iggCursorStartPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetCursorStartPos());
}

void iggCursorScreenPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetCursorScreenPos());
}

void iggSetCursorPos(IggVec2 const *localPos)
{
   Vec2Wrapper localPosArg(localPos);
   ImGui::SetCursorPos(*localPosArg);
}

void iggSetCursorScreenPos(IggVec2 const *absPos)
{
   Vec2Wrapper absPosArg(absPos);
   ImGui::SetCursorScreenPos(*absPosArg);
}

void iggAlignTextToFramePadding()
{
   ImGui::AlignTextToFramePadding();
}

float iggGetTextLineHeight(void)
{
   return ImGui::GetTextLineHeight();
}

float iggGetTextLineHeightWithSpacing(void)
{
   return ImGui::GetTextLineHeightWithSpacing();
}

IggBool iggTreeNode(char const *label, int flags)
{
   return ImGui::TreeNodeEx(label, flags) ? 1 : 0;
}

void iggTreePop(void)
{
   ImGui::TreePop();
}

void iggSetNextItemOpen(IggBool open, int cond)
{
   ImGui::SetNextItemOpen(open != 0, cond);
}

float iggGetTreeNodeToLabelSpacing(void)
{
   return ImGui::GetTreeNodeToLabelSpacing();
}

IggBool iggSelectable(char const *label, IggBool selected, int flags, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::Selectable(label, selected != 0, flags, *sizeArg) ? 1 : 0;
}

IggBool iggListBoxV(char const *label, int *currentItem, char const *const items[], int itemsCount, int heightItems)
{
   return ImGui::ListBox(label, currentItem, items, itemsCount, heightItems) ? 1 : 0;
}

void iggPlotLines(char const *label, float const *values, int valuesCount, int valuesOffset, char const *overlayText, float scaleMin, float scaleMax, IggVec2 const *graphSize)
{
   Vec2Wrapper graphSizeArg(graphSize);
   ImGui::PlotLines(label, values, valuesCount, valuesOffset, overlayText, scaleMin, scaleMax, *graphSizeArg);
}

void iggPlotHistogram(char const *label, float const *values, int valuesCount, int valuesOffset, char const *overlayText, float scaleMin, float scaleMax, IggVec2 const *graphSize)
{
   Vec2Wrapper graphSizeArg(graphSize);
   ImGui::PlotHistogram(label, values, valuesCount, valuesOffset, overlayText, scaleMin, scaleMax, *graphSizeArg);
}

void iggSetTooltip(char const *text)
{
   ImGui::SetTooltip("%s", text);
}

void iggBeginTooltip(void)
{
   ImGui::BeginTooltip();
}

void iggEndTooltip(void)
{
   ImGui::EndTooltip();
}

IggBool iggBeginMainMenuBar(void)
{
   return ImGui::BeginMainMenuBar() ? 1 : 0;
}

void iggEndMainMenuBar(void)
{
   ImGui::EndMainMenuBar();
}

IggBool iggBeginMenuBar(void)
{
   return ImGui::BeginMenuBar() ? 1 : 0;
}

void iggEndMenuBar(void)
{
   ImGui::EndMenuBar();
}

IggBool iggBeginMenu(char const *label, IggBool enabled)
{
   return ImGui::BeginMenu(label, enabled != 0) ? 1 : 0;
}

void iggEndMenu(void)
{
   ImGui::EndMenu();
}

IggBool iggMenuItem(char const *label, char const *shortcut, IggBool selected, IggBool enabled)
{
   return ImGui::MenuItem(label, shortcut, selected != 0, enabled != 0) ? 1 : 0;
}

void iggOpenPopup(char const *id)
{
   ImGui::OpenPopup(id);
}

IggBool iggBeginPopupModal(char const *name, IggBool *open, int flags)
{
   BoolWrapper openArg(open);
   return ImGui::BeginPopupModal(name, openArg, flags) ? 1 : 0;
}

IggBool iggBeginPopupContextItem(char const *label, int mouseButton)
{
   return ImGui::BeginPopupContextItem(label, mouseButton) ? 1 : 0;
}

void iggEndPopup(void)
{
   ImGui::EndPopup();
}

void iggCloseCurrentPopup(void)
{
   ImGui::CloseCurrentPopup();
}

IggBool iggIsItemHovered(int flags)
{
   return ImGui::IsItemHovered(flags) ? 1 : 0;
}

IggBool iggIsItemActive()
{
  return ImGui::IsItemActive() ? 1 : 0;
}

IggBool iggIsAnyItemActive()
{
   return ImGui::IsAnyItemActive() ? 1 : 0;
}

IggBool iggIsKeyDown(int key)
{
   return ImGui::IsKeyDown(key);
}

IggBool iggIsKeyPressed(int key, IggBool repeat)
{
   return ImGui::IsKeyPressed(key, repeat);
}

IggBool iggIsKeyReleased(int key)
{
   return ImGui::IsKeyReleased(key);
}

IggBool iggIsMouseDown(int button)
{
   return ImGui::IsMouseDown(button);
}

IggBool iggIsAnyMouseDown()
{
   return ImGui::IsAnyMouseDown();
}

IggBool iggIsMouseClicked(int button, IggBool repeat)
{
   return ImGui::IsMouseClicked(button, repeat);
}

IggBool iggIsMouseReleased(int button)
{
   return ImGui::IsMouseReleased(button);
}

IggBool iggIsMouseDoubleClicked(int button)
{
   return ImGui::IsMouseDoubleClicked(button);
}

void iggColumns(int count, char const *label, IggBool border)
{
   ImGui::Columns(count, label, border);
}

void iggNextColumn()
{
   ImGui::NextColumn();
}

int iggGetColumnIndex()
{
   return ImGui::GetColumnIndex();
}

int iggGetColumnWidth(int index)
{
   return ImGui::GetColumnWidth(index);
}

void iggSetColumnWidth(int index, float width)
{
   ImGui::SetColumnWidth(index, width);
}

float iggGetColumnOffset(int index)
{
   return ImGui::GetColumnOffset(index);
}

void iggSetColumnOffset(int index, float offsetX)
{
   ImGui::SetColumnOffset(index, offsetX);
}

int iggGetColumnsCount()
{
   return ImGui::GetColumnsCount();
}

void iggSetScrollHereY(float centerYRatio)
{
   ImGui::SetScrollHereY(centerYRatio);
}

void iggSetItemDefaultFocus()
{
   ImGui::SetItemDefaultFocus();
}

IggBool iggIsItemFocused()
{
   return ImGui::IsItemFocused();
}

IggBool iggIsAnyItemFocused()
{
   return ImGui::IsAnyItemFocused();
}

int iggGetMouseCursor()
{
   return ImGui::GetMouseCursor();
}

void iggSetMouseCursor(int cursor)
{
   ImGui::SetMouseCursor(cursor);
}

void iggSetKeyboardFocusHere(int offset)
{
   ImGui::SetKeyboardFocusHere(offset);
}

IggBool iggBeginTabBar(char const *str_id, int flags)
{
   return ImGui::BeginTabBar(str_id, flags) ? 1 : 0;
}

void iggEndTabBar()
{
   ImGui::EndTabBar();
}

IggBool iggBeginTabItem(char const *label, IggBool *p_open, int flags)
{
   BoolWrapper openArg(p_open);
   return ImGui::BeginTabItem(label, openArg, flags) ? 1 : 0;
}

void iggEndTabItem()
{
   ImGui::EndTabItem();
}

void iggSetTabItemClosed(char const *tab_or_docked_window_label)
{
   ImGui::SetTabItemClosed(tab_or_docked_window_label);
}

IggDrawList iggGetWindowDrawList()
{
  ImDrawList* drawlist = ImGui::GetWindowDrawList();
  return static_cast<IggDrawList>(drawlist);
}

void iggGetItemRectMin(IggVec2 *size)
{
  exportValue(*size, ImGui::GetItemRectMin());
}

void iggGetItemRectMax(IggVec2 *size)
{
  exportValue(*size, ImGui::GetItemRectMax());
}

void iggGetItemRectSize(IggVec2 *size)
{
  exportValue(*size, ImGui::GetItemRectSize());
}
