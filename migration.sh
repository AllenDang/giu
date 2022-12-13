git add migration.*
git commit --amend
git stash
files=`find . -iname \*go`

# switch to cimgui-go
sed -i -e 's/imgui/cimgui/g' $files
go get github.com/AllenDang/cimgui-go@5d9a1f0b1f727e23d05bc0ecb1f381d548b21c34
go mod tidy

# Types and constants:
sed -i -e 's/cimgui\.StyleColorID/cimgui\.Col/g' $files
sed -i -e 's/\(cimgui\.StyleColor\)\(\w\+\)/cimgui\.Col_\2/g' $files

sed -i -e 's/cimgui\.StyleVarID/cimgui\.StyleVar/g' $files
sed -i -e 's/\(cimgui\.StyleVar\)\(\w\+\)/cimgui\.StyleVar_\2/g' $files

sed -i -e 's/\(type MouseCursorType\).*/\1 cimgui\.MouseCursor/g' $files
sed -i -e 's/\(MouseCursor\)\(\w\+\)\( \+MouseCursorType = \).*/\1\2\3 cimgui\.MouseCursor_\2/g' $files
sed -i -e 's/\(cimgui\.MouseCursor_\)Count/\1COUNT/g' $files
sed -i -e 's/\(int(cursor)\)/cimgui.MouseCursor(cursor)/g' $files

sed -i -e 's/\(type DrawFlags.*\)int/\1 cimgui\.DrawFlags/g' $files
sed -i -e 's/\(DrawFlags\)\(\w\+\).*=.*/\1\2 DrawFlags = DrawFlags(cimgui\.DrawFlags_\2)/g' $files
sed -i -e 's/int(\(roundingCorners\))/cimgui\.DrawFlags(\1)/g' Canvas.go
sed -i -e 's/\(closed\),/cimgui\.DrawFlags(flags),/g' Canvas.go
# TODO: 
sed -i -e 's/\(cimgui\.DrawFlags_RoundCornersDefault\)/\1_/g' Canvas.go
sed -i -e 's/\(cimgui\.DrawFlags_RoundCornersMask\)/\1_/g' Canvas.go
sed -i -e 's/		DrawFlagsRoundCornersBottomLeft | DrawFlagsRoundCornersBottomRight//g' Canvas.go

sed -i -e 's/\(type Direction\) uint8/\1 cimgui.Dir/g' $files
sed -i -e 's/\(Direction\)\(\w\+\).*/\1\2 Direction = cimgui.Dir_\2/g' Direction.go
sed -i -e 's/\(uint8(b\.dir)\)/cimgui.Dir(b\.dir)/g' $files

# another types
sed -i -e 's/cimgui\.DrawList/cimgui\.ImDrawList/g' $files
sed -i -e 's/cimgui\.TextureID/cimgui\.ImTextureID/g' $files
sed -i -e 's/cimgui\.Vec2/cimgui\.ImVec2/g' $files
sed -i -e 's/cimgui\.Vec4/cimgui\.ImVec4/g' $files
sed -i -e 's/cimgui\.Font/cimgui\.ImFont/g' $files
sed -i -e 's/cimgui\.Condition/cimgui\.ImGuiCond/g' $files
sed -i -e 's/cimgui\.ImGuiCond\(\w\+\)/cimgui\.ImGuiCond_\1/g' $files
sed -i -e 's/cimgui\.InputTextCallback/cimgui\.InputTextCallback/g' $files
sed -i -e 's/cimgui\.Context/cimgui\.ImGuiContext/g' $files

sed -i -e 's/\(type InputTextFlags \)int/\1cimgui.InputTextFlags/g' $files
sed -i -e 's/\(type ComboFlags \)int/\1cimgui.ComboFlags/g' $files
sed -i -e 's/\(type SelectableFlags \)int/\1cimgui.SelectableFlags/g' $files
sed -i -e 's/\(type TabItemFlags \)int/\1cimgui.TabItemFlags/g' $files
sed -i -e 's/\(type TabBarFlags \)int/\1cimgui.TabBarFlags/g' $files
sed -i -e 's/\(type TreeNodeFlags \)int/\1cimgui.TreeNodeFlags/g' $files
sed -i -e 's/\(type FocusedFlags \)int/\1cimgui.FocusedFlags/g' $files
sed -i -e 's/\(type HoveredFlags \)int/\1cimgui.HoveredFlags/g' $files
sed -i -e 's/\(cimgui\.\)\(HoveredFlags\)\(\w\+\)/\1\2_\3/g' $files
sed -i -e 's/\(type TableFlags \)int/\1cimgui.TableFlags/g' $files
sed -i -e 's/\(type TableRowFlags \)int/\1cimgui.TableRowFlags/g' $files
sed -i -e 's/\(type TableColumnFlags \)int/\1cimgui.TableColumnFlags/g' $files
sed -i -e 's/\(type SliderFlags \)int/\1cimgui.SliderFlags/g' $files
sed -i -e 's/\(SliderFlags\)\(\w\+\).*/\1\2 SliderFlags = cimgui.SliderFlags_\2/g' Flags.go
sed -i -e 's/\(cimgui\.SliderFlags_InvalidMask\)/\1_/g' $files
sed -i -e 's/\(type PlotFlags \)int/\1cimgui.PlotFlags/g' $files
sed -i -e 's/\(type PlotAxisFlags \)int/\1cimgui.ImPlotAxisFlags/g' $files
#sed -i -e 's/\(type \)\(.*Flags\) int/\1 \2 cimgui.ImGui\2/g' $files

# Context; TODO - check if nothing else is changed
sed -i -e 's/cimgui\.IO()/cimgui\.GetIO()/g' $files
sed -i -e 's/cimgui\.IO/cimgui\.ImGuiIO/g' $files

# flags
#
# input text:
sed -i -e 's/cimgui\.InputTextFlags\(\w\+\)/cimgui\.InputTextFlags_\1/g' $files
# API CHANGE!
sed -i -e 's/^.*cimgui\.InputTextFlags_AlwaysInsertMode.*//g' $files

# window flags
sed -i -e 's/cimgui\.WindowFlags/cimgui\.WindowFlags/g' $files
# type was int; change to cimgui.mGuiWindowFlags
sed -i -e 's/\(type WindowFlags \)int/\1cimgui.GLFWWindowFlags/g' $files
sed -i -e 's/\(cimgui\.WindowFlags\)\(\w\+\)/WindowFlags(\1_\2)/g' $files

# combo flags
sed -i -e 's/cimgui\.ComboFlags/cimgui\.ComboFlags/g' $files
sed -i -e 's/\(cimgui\.ComboFlags\)\(\w\+\)/\1_\2/g' $files

# selectable flags
sed -i -e 's/cimgui\.SelectableFlags/cimgui\.SelectableFlags/g' $files
sed -i -e 's/\(cimgui\.SelectableFlags\)\(\w\+\)/\1_\2/g' $files

# Tab Item Flags
sed -i -e 's/cimgui\.TabItemFlags/cimgui\.TabItemFlags/g' $files
sed -i -e 's/\(cimgui\.TabItemFlags\)\(\w\+\)/\1_\2/g' $files
# remove TabItemFlagsNoPushID
# API CHANGE!
sed -i -e 's/^.*cimgui\.TabItemFlags_NoPushID.*//g' $files

# Tab Bar Flags
sed -i -e 's/cimgui\.TabBarFlags/cimgui\.TabBarFlags/g' $files
sed -i -e 's/\(cimgui\.TabBarFlags\)\(\w\+\)/\1_\2/g' $files
sed -i -e 's/\(cimgui\.TabBarFlags_FittingPolicyDefault\)/\1_/g' $files
sed -i -e 's/\(cimgui\.TabBarFlags_FittingPolicyMask\)/\1_/g' $files

# Tree Node Flags
sed -i -e 's/cimgui\.TreeNodeFlags/cimgui\.TreeNodeFlags/g' $files
sed -i -e 's/\(cimgui\.TreeNodeFlags\)\(\w\+\)/\1_\2/g' $files

# Focused Flags
sed -i -e 's/cimgui\.FocusedFlags/cimgui\.FocusedFlags/g' $files
sed -i -e 's/\(cimgui\.FocusedFlags\)\(\w\+\)/\1_\2/g' $files

# Hovered Flags

# Color Edit Flags
# TODO: COPY-PASTE them again (many things has changed
# API CHANGE!
sed -i -e 's/\(.*ColorEditFlags.*=.*\)/\/\/ \1/g' $files

# Table Flags
sed -i -e 's/cimgui\.TableFlags/cimgui\.TableFlags/g' $files
sed -i -e 's/\(cimgui\.TableFlags_NoBordersInBodyUntilResize\)TableFlags/\1/g' $files

# Table Row Flags
sed -i -e 's/cimgui\.TableRowFlags/cimgui\.TableRowFlags/g' $files

# Table Column Flags
sed -i -e 's/cimgui\.TableColumnFlags/cimgui\.TableColumnFlags/g' $files

# ImPlotFlags:
# disable flags that are not present:
# API CHANGE!
sed -i -e 's/\(cimgui\.\)Im\(PlotFlags\)/\1\2/g' $files
sed -i -e 's/\(.*cimgui\.PlotFlags_NoMousePos.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.PlotFlags_NoHighlight.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.PlotFlags_YAxis2.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.PlotFlags_YAxis3.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.PlotFlags_Query.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.PlotFlags_AntiAliased.*\)/\/\/ \1/g' $files

# Plot Axis Flags
# API CHANGE!
sed -i -e 's/\(.*cimgui\.ImPlotAxisFlags_LogScale.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.ImPlotAxisFlags_Time.*\)/\/\/ \1/g' $files
sed -i -e 's/\(cimgui\.\)Im\(PlotAxisFlags\)/\1\2/g' $files

# master window
# API CHANGE!
sed -i -e 's/\(.*cimgui\.GlfwDontCare.*\)/\/\/ \1/g' $files

sed -i -e 's/^/\/\/ /g' Markdown.go
echo "package giu" >> Markdown.go
sed -i -e 's/^/\/\/ /g' CodeEditor.go
echo "package giu" >> CodeEditor.go
sed -i -e 's/^/\/\/ /g' MemoryEditor.go
echo "package giu" >> MemoryEditor.go

# methods:
#
sed -i -e 's/cimgui\.PushID/cimgui\.PushID_Str/g' $files
sed -i -e 's/cimgui\.PushStyleVarFloat/cimgui\.PushStyleVar_Float/g' $files
sed -i -e 's/cimgui\.PushStyleVarVec2/cimgui\.PushStyleVar_Vec2/g' $files
sed -i -e 's/cimgui\.PushStyleColor/cimgui\.PushStyleColor_Vec4/g' $files
sed -i -e 's/\(.*\):= \(cimgui\.\)\(ContentRegionAvail\)()/\1:= cimgui.ImVec2{};\2Get\3(\&\1)/g' ImageWidgets.go
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


sed -i -e 's/\(DrawList\.AddText\)/\1_Vec2/g' Canvas.go
sed -i -e 's/\(numSegments int\)/\132/g' Canvas.go
sed -i -e 's/\(segments int\)/\132/g' Canvas.go
sed -i -e 's/\(min12 int\)/\132/g' Canvas.go
sed -i -e 's/\(max12 int\)/\132/g' Canvas.go
sed -i -e 's/\(closed bool\)/flags DrawFlags/g' Canvas.go

# styles
sed -i -e 's/\(cimgui\.PopStyle.*V(\)\(.*\))/\1int32(\2))/g' $files
sed -i -e 's/\(cimgui\.BeginDisabled.*\)/if ss.disabled {\1}/g' StyleSetter.go
sed -i -e 's/\(cimgui\.EndDisabled.*\)/if ss.disabled {\1}/g' StyleSetter.go
sed -i -e 's/\(cimgui\.BeginDisabled\)(.*)/\1()/g' $files

# Style.go
## Mouse Cursor

sed -i -e 's/cimgui\.CurrentStyle/cimgui\.GetStyle/g' $files
sed -i -e 's/\(cimgui\.GetStyle()\.\)\(\w\+()\)/\1Get\2/g' $files

# split layout/style
sed -i -e 's/\(cimgui\.GetStyle().GetColor\)/cimgui\.GetStyleColorVec4/g' $files


# ClickableWidgets.go

sed -i -e 's/\(cimgui\.TreeNode\)V/\1Ex_StrV/g' $files
sed -i -e 's/\(cimgui\.TreeNodeEx_StrV.*\)int(\(.*\))/\1cimgui\.TreeNodeFlags(\2)/g' $files

sed -i -e 's/\(cimgui\.Selectable\)V/\1_BoolV/g' $files
sed -i -e 's/\(cimgui\.Selectable_BoolV.*\)int(\(.*\))/\1cimgui\.SelectableFlags(\2)/g' $files

sed -i -e 's/\(cimgui\.RadioButton\)/\1_Bool/g' $files

# Events.go
#
sed -i -e 's/\(type MouseButton \)int/\1cimgui.MouseButton/g' $files
sed -i -e 's/\(MouseButton\)\(\w\+\).*=.*/\1\2 MouseButton = MouseButton(cimgui\.MouseButton_\2)/g' Events.go

sed -i -e 's/\(int(mouseButton)\)/cimgui.MouseButton(mouseButton)/g' Events.go
sed -i -e 's/\(cimgui.IsItemClicked\)/\1V/g' $files

sed -i -e 's/\(int(button)\)/cimgui.MouseButton(button)/g' Events.go

sed -i -e 's/\(cimgui.IsWindowFocused\)/\1V/g' $files
sed -i -e 's/\(IsWindowFocusedV(\)int(flags)/\1cimgui.FocusedFlags(flags)/g' Events.go

sed -i -e 's/\(cimgui.IsWindowHovered\)/\1V/g' $files
sed -i -e 's/\(IsWindowHoveredV(\)int(flags)/\1cimgui.HoveredFlags(flags)/g' Events.go

# list clipper
#
sed -i -e 's/\(cimgui.NewListClipper\)/cimgui\.NewImGuiListClipper/g' $files
sed -i -e 's/\(clipper\.Delete\)/clipper\.Destroy/g' $files
sed -i -e 's/\(clipper\.Begin(\)\(.*\))/\1int32(\2))/g' $files
sed -i -e 's/\(clipper\.DisplayStart\)/clipper\.GetDisplayStart/g' $files
sed -i -e 's/\(clipper\.DisplayEnd\)/clipper\.GetDisplayEnd/g' $files

# ProgressIndicator.go
sed -i -e 's/int(p.radius)/int32(p\.radius)/g' ProgressIndicator.go

# Popups.go
#
sed -i -e 's/\(cimgui\.OpenPopup\)/\1_Str/g' Popups.go
sed -i -e 's/\(cimgui\.BeginPopup\)/\1V/g' Popups.go
sed -i -e 's/\(cimgui\.BeginPopupVModalV\)/cimgui\.BeginPopupModalV/g' Popups.go
sed -i -e 's/\(BeginPopup.*(.*\)int(\(p.flags\))/\1cimgui.WindowFlags(\2)/g' Popups.go


# Window.go
#
sed -i -e 's/\(.*\)= \(cimgui\.\)\(WindowPos\)()/\1= cimgui.ImVec2{};\2Get\3(\&\1)/g' Window.go
sed -i -e 's/\(.*\)= \(cimgui\.\)\(WindowSize\)()/\1= cimgui.ImVec2{};\2Get\3(\&\1)/g' Window.go
sed -i -e 's/\(cimgui\.Begin.*\)int\((.*flags).*\)/\1cimgui\.WindowFlags\2/g' Window.go

# FontAtlasProcessor.go
#
sed -i -e 's/\(IO()\.\)\(Fonts()\)/\1Get\2/g' $files
sed -i -e 's/\(io\.\)\(Fonts()\)/\1Get\2/g' $files

sed -i -e 's/\(cimgui\.\)NewGlyphRanges/\1NewGlyphRange/g' $files
sed -i -e 's/\(cimgui\.New\)\(FontGlyphRangesBuilder\)/\1Im\2/g' $files
sed -i -e 's/\(fonts\.\)\(GlyphRangesDefault\)/\1Get\2/g' $files
sed -i -e 's/\(cimgui\.New\)\(FontConfig\)/\1Im\2/g' $files

# ExtraWidgets.go
#
sed -i -e 's/\(style\)\(\.GetColor\)/cimgui\.GetStyleColorVec4/g' $files
sed -i -e 's/style := .*//g' ExtraWidgets.go
sed -i -e 's/cimgui\.CurrentIO/cimgui\.GetIO/g' $files
sed -i -e 's/cimgui\.\(MouseCursor\)\(\w\+\)/cimgui\.\1_\2/g' ExtraWidgets.go
sed -i -e 's/\(cimgui\.TableNextRow\)/\1V/g' ExtraWidgets.go
sed -i -e 's/\(cimgui\.BeginTable\)/\1V/g' ExtraWidgets.go
sed -i -e 's/\(cimgui\.BeginTable.*\)\(colCount\)/\1int32(\2)/g' ExtraWidgets.go
sed -i -e 's/\(cimgui\.TableSetupScrollFreeze(\)\(.*\), \(.*\))/\1int32(\2), int32(\3))/g' ExtraWidgets.go TableWidgets.go

# SliderWidget.go
#
sed -i -e 's/\(cimgui\.SliderIntV(.*\)\() \)/\1, 0\2/g' SliderWidgets.go
sed -i -e 's/int\((vs\.flags)\)/cimgui\.SliderFlags\1/g' SliderWidgets.go

# TableWidgets.go
#
sed -i -e 's/\(cimgui\.TableNextRow\)/\1V/g' TableWidgets.go
# TODO - converting to float32 is wrong idea since wi're converting from float64 :-)
sed -i -e 's/\(cimgui\.TableNextRowV(.*, \)\(.*\))/\1float32(\2))/g' TableWidgets.go
sed -i -e 's/\(cimgui\.TableSetBgColor\)/\1V/g' TableWidgets.go
sed -i -e 's/\(cimgui\.TableSetupColumn\)/\1V/g' TableWidgets.go
sed -i -e 's/\(cimgui\.BeginTable\)/\1V/g' TableWidgets.go
sed -i -e 's/\(cimgui\.\)\(TableBgTarget\)/\1ImGui\2/g' TableWidgets.go
sed -i -e 's/\(cimgui\.GetColorU32\)/\1_Vec4/g' TableWidgets.go
# TODO - converting to float32 is wrong idea since wi're converting from float64 :-)
sed -i -e 's/\(cimgui\.BeginTableV(\)\(.*\), \(.*\), \(.*\), \(.*\), \(.*\))/\1\2, int32(\3\), \4, \5, float32(\6))/g' TableWidgets.go

# Widgets.go
#
sed -i -e 's/\(cimgui\.BeginChild\)V/\1_StrV/g' Widgets.go
sed -i -e 's/\(cimgui\.BeginChild_StrV(.*, \)int\((.*)\))/\1cimgui.WindowFlags\2)/g' Widgets.go
sed -i -e 's/\(cimgui\.BeginComboV(.*, \)int\((.*)\))/\1cimgui.ComboFlags\2)/g' Widgets.go
sed -i -e 's/\(cimgui\.Selectable\)/\1_Bool/g' Widgets.go
# TODO: mouse button here is PopupFlags in fact - need to update stuff manually
sed -i -e 's/int(\(.*mouseButton\))/cimgui\.ImGuiPopupFlags(\1)/g' Widgets.go

# TODO: add possibility to specify flags here
sed -i -e 's/\(cimgui\.DragIntV(.*\))/\1, 0)/g' Widgets.go

sed -i -e 's/\(cimgui\.MenuItem\)V/\1_BoolV/g' $files
sed -i -e 's/\(cimgui\.BeginTabItemV(.*, \)int\((.*)\))/\1cimgui.TabItemFlags\2)/g' Widgets.go
sed -i -e 's/\(cimgui\.BeginTabBarV(.*, \)int\((.*)\))/\1cimgui.TabBarFlags\2)/g' Widgets.go

# TODO: color edit flags are disabled now
sed -i -e 's/\(flags: ColorEditFlagsNone\)/\/\/ \1/g' Widgets.go

# gofmt
gofmt -w -s $files
