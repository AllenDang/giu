#!/bin/bash
git add migration.*
git commit --amend
git stash
files=$(find . -iname \*go)

# switch to imgui-go
sed -i -e 's/\(AllenDang\/\)\(imgui-go\)/\1c\2/g' $files
go get github.com/AllenDang/cimgui-go@584fdd8e2a0b23ba815071183d03f3e755990e4f
go mod edit -replace github.com/AllenDang/cimgui-go=github.com/neclepsio/cimgui-go@e2590f5e36f74e2d2b285f15ff064286627c7e71
go mod tidy

# Types and constants:
sed -i -e 's/imgui\.StyleColorID/imgui\.Col/g' $files
sed -i -e 's/\(imgui\.StyleColor\)\(\w\+\)/imgui\.Col\2/g' $files

sed -i -e 's/imgui\.StyleVarID/imgui\.StyleVar/g' $files
sed -i -e 's/\(imgui\.StyleVar\)\(\w\+\)/imgui\.StyleVar\2/g' $files

sed -i -e 's/\(type MouseCursorType\).*/\1 imgui\.MouseCursor/g' $files
sed -i -e 's/\(MouseCursor\)\(\w\+\)\( \+MouseCursorType = \).*/\1\2\3 imgui\.MouseCursor\2/g' $files
sed -i -e 's/\(imgui\.MouseCursor\)Count/\1COUNT/g' $files
sed -i -e 's/\(int(cursor)\)/imgui.MouseCursor(cursor)/g' $files

sed -i -e 's/\(type DrawFlags.*\)int/\1 imgui\.DrawFlags/g' $files
sed -i -e 's/\(DrawFlags\)\(\w\+\).*=.*/\1\2 DrawFlags = DrawFlags(imgui\.DrawFlags\2)/g' $files
sed -i -e 's/int(\(roundingCorners\))/imgui\.DrawFlags(\1)/g' Canvas.go
sed -i -e 's/\(closed\),/imgui\.DrawFlags(flags),/g' Canvas.go
# TODO:
sed -i -e 's/		DrawFlagsRoundCornersBottomLeft | DrawFlagsRoundCornersBottomRight//g' Canvas.go

sed -i -e 's/\(type Direction\) uint8/\1 imgui.Dir/g' $files
sed -i -e 's/\(Direction\)\(\w\+\).*/\1\2 Direction = imgui.Dir\2/g' Direction.go
sed -i -e 's/\(uint8(b\.dir)\)/imgui.Dir(b\.dir)/g' $files

# another types
sed -i -e 's/imgui\.Condition/imgui\.Cond/g' $files
sed -i -e 's/imgui\.Cond\(\w\+\)/imgui\.Cond\1/g' $files
sed -i -e 's/imgui\.InputTextCallback/imgui\.InputTextCallback/g' $files

sed -i -e 's/\(type InputTextFlags \)int/\1imgui.InputTextFlags/g' $files
sed -i -e 's/\(type ComboFlags \)int/\1imgui.ComboFlags/g' $files
sed -i -e 's/\(type SelectableFlags \)int/\1imgui.SelectableFlags/g' $files
sed -i -e 's/\(type TabItemFlags \)int/\1imgui.TabItemFlags/g' $files
sed -i -e 's/\(type TabBarFlags \)int/\1imgui.TabBarFlags/g' $files
sed -i -e 's/\(type TreeNodeFlags \)int/\1imgui.TreeNodeFlags/g' $files
sed -i -e 's/\(type FocusedFlags \)int/\1imgui.FocusedFlags/g' $files
sed -i -e 's/\(type HoveredFlags \)int/\1imgui.HoveredFlags/g' $files
sed -i -e 's/\(imgui\.\)\(HoveredFlags\)\(\w\+\)/\1\2\3/g' $files
sed -i -e 's/\(type TableFlags \)int/\1imgui.TableFlags/g' $files
sed -i -e 's/\(type TableRowFlags \)int/\1imgui.TableRowFlags/g' $files
sed -i -e 's/\(type TableColumnFlags \)int/\1imgui.TableColumnFlags/g' $files
sed -i -e 's/\(type SliderFlags \)int/\1imgui.SliderFlags/g' $files
sed -i -e 's/\(SliderFlags\)\(\w\+\).*/\1\2 SliderFlags = imgui.SliderFlags\2/g' Flags.go
sed -i -e 's/\(type PlotFlags \)int/\1imgui.PlotFlags/g' $files
sed -i -e 's/\(type PlotAxisFlags \)int/\1imgui.PlotAxisFlags/g' $files
#sed -i -e 's/\(type \)\(.*Flags\) int/\1 \2 imgui.ImGui\2/g' $files

# Context; TODO - check if nothing else is changed
sed -i -e 's/imgui\.IO()/imgui\.GetIO()/g' $files

# flags
#
# input text:
sed -i -e 's/imgui\.InputTextFlags\(\w\+\)/imgui\.InputTextFlags\1/g' $files
# API CHANGE!
sed -i -e 's/^.*imgui\.InputTextFlagsAlwaysInsertMode.*//g' $files

# window flags
sed -i -e 's/imgui\.WindowFlags/imgui\.WindowFlags/g' $files
# type was int; change to imgui.mGuiWindowFlags
sed -i -e 's/\(type WindowFlags \)int/\1imgui.GLFWWindowFlags/g' $files
sed -i -e 's/\(imgui\.WindowFlags\)\(\w\+\)/WindowFlags(\1\2)/g' $files

# combo flags
sed -i -e 's/imgui\.ComboFlags/imgui\.ComboFlags/g' $files
sed -i -e 's/\(imgui\.ComboFlags\)\(\w\+\)/\1\2/g' $files

# selectable flags
sed -i -e 's/imgui\.SelectableFlags/imgui\.SelectableFlags/g' $files
sed -i -e 's/\(imgui\.SelectableFlags\)\(\w\+\)/\1\2/g' $files

# Tab Item Flags
sed -i -e 's/imgui\.TabItemFlags/imgui\.TabItemFlags/g' $files
sed -i -e 's/\(imgui\.TabItemFlags\)\(\w\+\)/\1\2/g' $files
# remove TabItemFlagsNoPushID
# API CHANGE!
sed -i -e 's/^.*imgui\.TabItemFlagsNoPushID.*//g' $files

# Tab Bar Flags
sed -i -e 's/imgui\.TabBarFlags/imgui\.TabBarFlags/g' $files
sed -i -e 's/\(imgui\.TabBarFlags\)\(\w\+\)/\1\2/g' $files

# Tree Node Flags
sed -i -e 's/imgui\.TreeNodeFlags/imgui\.TreeNodeFlags/g' $files
sed -i -e 's/\(imgui\.TreeNodeFlags\)\(\w\+\)/\1\2/g' $files

# Focused Flags
sed -i -e 's/imgui\.FocusedFlags/imgui\.FocusedFlags/g' $files
sed -i -e 's/\(imgui\.FocusedFlags\)\(\w\+\)/\1\2/g' $files

# Hovered Flags

# Color Edit Flags
# TODO: COPY-PASTE them again (many things has changed
# API CHANGE!
sed -i -e 's/\(.*ColorEditFlags.*=.*\)/\/\/ \1/g' $files

# Table Flags
sed -i -e 's/imgui\.TableFlags_/imgui\.TableFlags/g' $files
sed -i -e 's/\(imgui\.TableFlagsNoBordersInBodyUntilResize\)TableFlags/\1/g' $files
sed -i -e 's/\(imgui\.TableFlags.*\)_/\1/g' $files

# Table Row Flags
sed -i -e 's/imgui\.TableRowFlags_/imgui\.TableRowFlags/g' $files

# Table Column Flags
sed -i -e 's/imgui\.TableColumnFlags_/imgui\.TableColumnFlags/g' $files
sed -i -e 's/\(imgui\.TableColumnFlags.*\)_/\1/g' $files

# ImPlotFlags:
# disable flags that are not present:
# API CHANGE!
sed -i -e 's/\(imgui\.\)Im\(PlotFlags\)/\1\2/g' $files
sed -i -e 's/\(imgui\.PlotFlags\)_/\1/g' $files
sed -i -e 's/\(.*imgui\.PlotFlagsNoMousePos.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*imgui\.PlotFlagsNoHighlight.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*imgui\.PlotFlagsYAxis2.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*imgui\.PlotFlagsYAxis3.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*imgui\.PlotFlagsQuery.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*imgui\.PlotFlagsAntiAliased.*\)/\/\/ \1/g' $files

# Plot Axis Flags
# API CHANGE!
sed -i -e 's/\(imgui\.\)Im\(PlotAxisFlags\)/\1\2/g' $files
sed -i -e 's/\(.*imgui\.PlotAxisFlagsLogScale.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*imgui\.PlotAxisFlagsTime.*\)/\/\/ \1/g' $files
sed -i -e 's/\(imgui\.PlotAxisFlags\)_/\1/g' $files

# master window
# API CHANGE!
sed -i -e 's/\(.*imgui\.GlfwDontCare.*\)/\/\/ \1/g' $files

sed -i -e 's/^/\/\/ /g' Markdown.go
echo "package giu" >> Markdown.go
sed -i -e 's/^/\/\/ /g' CodeEditor.go
echo "package giu" >> CodeEditor.go
sed -i -e 's/^/\/\/ /g' MemoryEditor.go
echo "package giu" >> MemoryEditor.go

# methods:
#
sed -i -e 's/imgui\.PushID/imgui\.PushIDStr/g' $files
sed -i -e 's/imgui\.PushStyleVarFloat/imgui\.PushStyleVarFloat/g' $files
sed -i -e 's/imgui\.PushStyleVarVec2/imgui\.PushStyleVarVec2/g' $files
sed -i -e 's/imgui\.PushStyleColor/imgui\.PushStyleColorVec4/g' $files
sed -i -e 's/\(.*\):= \(imgui\.\)\(ContentRegionAvail\)()/\1:= imgui.ImVec2{};\2Get\3(\&\1)/g' ImageWidgets.go
sed -i -e 's/\.AddLine/\.AddLineV/g' $files
sed -i -e 's/\.AddRect/\.AddRectV/g' Canvas.go

echo "
// ColorToUint converts GO color into Uint32 color
// it is 0xRRGGBBAA
func ColorToUint(col color.Color) uint32 {
        r, g, b, a := col.RGBA()
        mask := uint32(0xff)
        return r&mask<<24 + g&mask<<16 + b&mask<<8 + a&mask
}

// UintToColor converts uint32 of form 0xRRGGBB into color.RGBA
func UintToColor(col uint32) *color.RGBA {
        mask := 0xff
        r := byte(col >> 24 & uint32(mask))
        g := byte(col >> 16 & uint32(mask))
        b := byte(col >> 8 & uint32(mask))
        a := byte(col >> 0 & uint32(mask))
        return &color.RGBA{
                R: r,
                G: g,
                B: b,
                A: a,
        }
}
" >> Utils.go

# so now, it seems that DrawList api changed so that it requires uint32 instead of Vec4
# as colors.
sed -i -e 's/ToVec4Color/ColorToUint/g' Canvas.go
sed -i -e 's/DrawList\.AddRectVFilled/DrawList\.AddRectFilledV/g' Canvas.go
sed -i -e 's/DrawList\.AddTriangle(/DrawList\.AddTriangleV(/g' Canvas.go
sed -i -e 's/DrawList\.AddCircle(/DrawList\.AddCircleV(/g' Canvas.go
sed -i -e 's/DrawList\.AddBezierCubic(/DrawList\.AddBezierCubicV(/g' Canvas.go
sed -i -e 's/DrawList\.AddQuad(/DrawList\.AddQuadV(/g' Canvas.go
sed -i -e 's/\(DrawList\.PathStroke\)(/\1V(/g' Canvas.go
sed -i -e 's/\(DrawList\.PathArcTo\)(/\1V(/g' Canvas.go
sed -i -e 's/\(DrawList\.PathBezierCubicCurveTo\)(/\1V(/g' Canvas.go


sed -i -e 's/\(DrawList\.AddText\)/\1Vec2/g' Canvas.go
sed -i -e 's/\(numSegments int\)/\132/g' Canvas.go
sed -i -e 's/\(segments int\)/\132/g' Canvas.go
sed -i -e 's/\(min12 int\)/\132/g' Canvas.go
sed -i -e 's/\(max12 int\)/\132/g' Canvas.go
sed -i -e 's/\(closed bool\)/flags DrawFlags/g' Canvas.go

# styles
sed -i -e 's/\(imgui\.PopStyle.*V(\)\(.*\))/\1int32(\2))/g' $files
sed -i -e 's/\(imgui\.BeginDisabled.*\)/if ss.disabled {\1}/g' StyleSetter.go
sed -i -e 's/\(imgui\.EndDisabled.*\)/if ss.disabled {\1}/g' StyleSetter.go
sed -i -e 's/\(imgui\.BeginDisabled\)(.*)/\1()/g' $files

# Style.go
## Mouse Cursor

sed -i -e 's/imgui\.CurrentStyle/imgui\.Style/g' $files
sed -i -e 's/\(imgui\.Style()\.\)\(\w\+()\)/\1Get\2/g' $files

# split layout/style
sed -i -e 's/\(imgui\.Style().GetColor\)/imgui\.StyleColorVec4/g' $files


# ClickableWidgets.go

sed -i -e 's/\(imgui\.TreeNode\)V/\1ExStrV/g' $files
sed -i -e 's/\(imgui\.TreeNodeExStrV.*\)int(\(.*\))/\1imgui\.TreeNodeFlags(\2)/g' $files

sed -i -e 's/\(imgui\.Selectable\)V/\1BoolV/g' $files
sed -i -e 's/\(imgui\.SelectableBoolV.*\)int(\(.*\))/\1imgui\.SelectableFlags(\2)/g' $files

sed -i -e 's/\(imgui\.RadioButton\)/\1Bool/g' $files

# Events.go
#
sed -i -e 's/\(type MouseButton \)int/\1imgui.MouseButton/g' $files
sed -i -e 's/\(MouseButton\)\(\w\+\).*=.*/\1\2 MouseButton = MouseButton(imgui\.MouseButton\2)/g' Events.go

sed -i -e 's/\(int(mouseButton)\)/imgui.MouseButton(mouseButton)/g' Events.go
sed -i -e 's/\(imgui.IsItemClicked\)/\1V/g' $files

sed -i -e 's/\(int(button)\)/imgui.MouseButton(button)/g' Events.go

sed -i -e 's/\(imgui.IsWindowFocused\)/\1V/g' $files
sed -i -e 's/\(IsWindowFocusedV(\)int(flags)/\1imgui.FocusedFlags(flags)/g' Events.go

sed -i -e 's/\(imgui.IsWindowHovered\)/\1V/g' $files
sed -i -e 's/\(IsWindowHoveredV(\)int(flags)/\1imgui.HoveredFlags(flags)/g' Events.go

# list clipper
#
sed -i -e 's/\(imgui.NewListClipper\)/imgui\.NewImGuiListClipper/g' $files
sed -i -e 's/\(clipper\.Delete\)/clipper\.Destroy/g' $files
sed -i -e 's/\(clipper\.Begin(\)\(.*\))/\1int32(\2))/g' $files
sed -i -e 's/\(clipper\.DisplayStart\)/clipper\.GetDisplayStart/g' $files
sed -i -e 's/\(clipper\.DisplayEnd\)/clipper\.GetDisplayEnd/g' $files

# ProgressIndicator.go
sed -i -e 's/int(p.radius)/int32(p\.radius)/g' ProgressIndicator.go

# Popups.go
#
sed -i -e 's/\(imgui\.OpenPopup\)/\1Str/g' Popups.go
sed -i -e 's/\(imgui\.BeginPopup\)/\1V/g' Popups.go
sed -i -e 's/\(imgui\.BeginPopupVModalV\)/imgui\.BeginPopupModalV/g' Popups.go
sed -i -e 's/\(BeginPopup.*(.*\)int(\(p.flags\))/\1imgui.WindowFlags(\2)/g' Popups.go


# Window.go
#
sed -i -e 's/\(imgui\.\)\(WindowPos\)/\1Get\2/g' Window.go
sed -i -e 's/\(imgui\.\)\(WindowSize\)/\1Get\2/g' Window.go
sed -i -e 's/\(imgui\.Begin.*\)int\((.*flags).*\)/\1imgui\.WindowFlags\2/g' Window.go

# FontAtlasProcessor.go
#
#sed -i -e 's/\(IO()\.\)\(Fonts()\)/\1Get\2/g' $files
sed -i -e 's/\(io\.\)\(Fonts()\)/\1Get\2/g' $files

sed -i -e 's/\(imgui\.\)NewGlyphRanges/\1NewGlyphRange/g' $files
sed -i -e 's/\(imgui\.New\)\(FontGlyphRangesBuilder\)/\1Im\2/g' $files
sed -i -e 's/\(fonts\.\)\(GlyphRangesDefault\)/\1Get\2/g' $files
sed -i -e 's/\(imgui\.New\)\(FontConfig\)/\1Im\2/g' $files

# ExtraWidgets.go
#
sed -i -e 's/\(style\)\(\.GetColor\)/imgui\.StyleColorVec4/g' $files
sed -i -e 's/style := .*//g' ExtraWidgets.go
sed -i -e 's/imgui\.CurrentIO/imgui\.GetIO/g' $files
sed -i -e 's/imgui\.\(MouseCursor\)\(\w\+\)/imgui\.\1\2/g' ExtraWidgets.go
sed -i -e 's/\(imgui\.TableNextRow\)/\1V/g' ExtraWidgets.go
sed -i -e 's/\(imgui\.BeginTable\)/\1V/g' ExtraWidgets.go
sed -i -e 's/\(imgui\.BeginTable.*\)\(colCount\)/\1int32(\2)/g' ExtraWidgets.go
sed -i -e 's/\(imgui\.TableSetupScrollFreeze(\)\(.*\), \(.*\))/\1int32(\2), int32(\3))/g' ExtraWidgets.go TableWidgets.go

# SliderWidget.go
#
sed -i -e 's/\(imgui\.SliderIntV(.*\)\() \)/\1, 0\2/g' SliderWidgets.go
sed -i -e 's/int\((vs\.flags)\)/imgui\.SliderFlags\1/g' SliderWidgets.go

# TableWidgets.go
#
sed -i -e 's/\(imgui\.TableNextRow\)/\1V/g' TableWidgets.go
# TODO - converting to float32 is wrong idea since wi're converting from float64 :-)
sed -i -e 's/\(imgui\.TableNextRowV(.*, \)\(.*\))/\1float32(\2))/g' TableWidgets.go
sed -i -e 's/\(imgui\.TableSetBgColor\)/\1V/g' TableWidgets.go
sed -i -e 's/\(imgui\.TableSetupColumn\)/\1V/g' TableWidgets.go
sed -i -e 's/\(imgui\.BeginTable\)/\1V/g' TableWidgets.go
sed -i -e 's/\(imgui\.\)\(TableBgTarget\)/\1ImGui\2/g' TableWidgets.go
sed -i -e 's/\(imgui\.GetColorU32\)/\1Vec4/g' TableWidgets.go
# TODO - converting to float32 is wrong idea since wi're converting from float64 :-)
sed -i -e 's/\(imgui\.BeginTableV(\)\(.*\), \(.*\), \(.*\), \(.*\), \(.*\))/\1\2, int32(\3\), \4, \5, float32(\6))/g' TableWidgets.go

# Widgets.go
#
sed -i -e 's/\(imgui\.BeginChild\)V/\1StrV/g' Widgets.go
sed -i -e 's/\(imgui\.BeginChildStrV(.*, \)int\((.*)\))/\1imgui.WindowFlags\2)/g' Widgets.go
sed -i -e 's/\(imgui\.BeginComboV(.*, \)int\((.*)\))/\1imgui.ComboFlags\2)/g' Widgets.go
sed -i -e 's/\(imgui\.Selectable\)/\1Bool/g' Widgets.go
# TODO: mouse button here is PopupFlags in fact - need to update stuff manually
sed -i -e 's/int(\(.*mouseButton\))/imgui\.ImGuiPopupFlags(\1)/g' Widgets.go

# TODO: add possibility to specify flags here
sed -i -e 's/\(imgui\.DragIntV(.*\))/\1, 0)/g' Widgets.go

sed -i -e 's/\(imgui\.MenuItem\)V/\1BoolV/g' $files
sed -i -e 's/\(imgui\.BeginTabItemV(.*, \)int\((.*)\))/\1imgui.TabItemFlags\2)/g' Widgets.go
sed -i -e 's/\(imgui\.BeginTabBarV(.*, \)int\((.*)\))/\1imgui.TabBarFlags\2)/g' Widgets.go

# TODO: color edit flags are disabled now
sed -i -e 's/\(flags: ColorEditFlagsNone\)/\/\/ \1/g' Widgets.go

sed -i -e 's/\(imgui\.\)Get\(WindowDrawList\)/\1\2/g' $files

# gofmt
gofmt -w -s $files
