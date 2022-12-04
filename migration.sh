git add migration.*
git commit --amend
git stash
files=`find . -iname \*go`

# switch to cimgui-go
sed -i -e 's/imgui/cimgui/g' $files
go get github.com/AllenDang/cimgui-go@158164eb30c79c00a3c393a1d6642609f2f2e206
go mod tidy

# mainly StyleIDs.go
sed -i -e 's/cimgui\.StyleColorID/cimgui\.ImGuiCol/g' $files
sed -i -e 's/cimgui\.StyleVarID/cimgui\.ImGuiStyleVar/g' $files
sed -i -e 's/\(cimgui\.StyleColor\)\(\w\+\)/cimgui\.ImGuiCol_\2/g' $files
sed -i -e 's/\(cimgui\.StyleVar\)\(\w\+\)/cimgui\.ImGuiStyleVar_\2/g' $files

# another types
sed -i -e 's/cimgui\.DrawList/cimgui\.ImDrawList/g' $files
sed -i -e 's/cimgui\.TextureID/cimgui\.ImTextureID/g' $files
sed -i -e 's/cimgui\.Vec2/cimgui\.ImVec2/g' $files
sed -i -e 's/cimgui\.Vec4/cimgui\.ImVec4/g' $files
sed -i -e 's/cimgui\.Font/cimgui\.ImFont/g' $files
sed -i -e 's/cimgui\.Condition/cimgui\.ImGuiCond/g' $files
sed -i -e 's/cimgui\.ImGuiCond\(\w\+\)/cimgui\.ImGuiCond_\1/g' $files
sed -i -e 's/cimgui\.InputTextCallback/cimgui\.ImGuiInputTextCallback/g' $files
sed -i -e 's/cimgui\.Context/cimgui\.ImGuiContext/g' $files

# Context; TODO - check if nothing else is changed
sed -i -e 's/cimgui\.IO()/cimgui\.GetIO()/g' $files
sed -i -e 's/cimgui\.IO/cimgui\.ImGuiIO/g' $files

# flags
#
# input text:
sed -i -e 's/cimgui\.InputTextFlags\(\w\+\)/cimgui\.ImGuiInputTextFlags_\1/g' $files
# API CHANGE!
sed -i -e 's/^.*cimgui\.ImGuiInputTextFlags_AlwaysInsertMode.*//g' $files

# window flags
sed -i -e 's/cimgui\.WindowFlags/cimgui\.ImGuiWindowFlags/g' $files
# type was int; change to cimgui.ImGuiWindowFlags
sed -i -e 's/\(type WindowFlags \)int/\1cimgui.GLFWWindowFlags/g' $files
sed -i -e 's/\(cimgui\.ImGuiWindowFlags\)\(\w\+\)/WindowFlags(\1_\2)/g' $files

# combo flags
sed -i -e 's/cimgui\.ComboFlags/cimgui\.ImGuiComboFlags/g' $files
sed -i -e 's/\(cimgui\.ImGuiComboFlags\)\(\w\+\)/\1_\2/g' $files

# selectable flags
sed -i -e 's/cimgui\.SelectableFlags/cimgui\.ImGuiSelectableFlags/g' $files
sed -i -e 's/\(cimgui\.ImGuiSelectableFlags\)\(\w\+\)/\1_\2/g' $files

# Tab Item Flags
sed -i -e 's/cimgui\.TabItemFlags/cimgui\.ImGuiTabItemFlags/g' $files
sed -i -e 's/\(cimgui\.ImGuiTabItemFlags\)\(\w\+\)/\1_\2/g' $files
# remove TabItemFlagsNoPushID
# API CHANGE!
sed -i -e 's/^.*cimgui\.ImGuiTabItemFlags_NoPushID.*//g' $files

# Tab Bar Flags
sed -i -e 's/cimgui\.TabBarFlags/cimgui\.ImGuiTabBarFlags/g' $files
sed -i -e 's/\(cimgui\.ImGuiTabBarFlags\)\(\w\+\)/\1_\2/g' $files
sed -i -e 's/\(cimgui\.ImGuiTabBarFlags_FittingPolicyDefault\)/\1_/g' $files
sed -i -e 's/\(cimgui\.ImGuiTabBarFlags_FittingPolicyMask\)/\1_/g' $files

# Tree Node Flags
sed -i -e 's/cimgui\.TreeNodeFlags/cimgui\.ImGuiTreeNodeFlags/g' $files
sed -i -e 's/\(cimgui\.ImGuiTreeNodeFlags\)\(\w\+\)/\1_\2/g' $files

# Focused Flags
sed -i -e 's/cimgui\.FocusedFlags/cimgui\.ImGuiFocusedFlags/g' $files
sed -i -e 's/\(cimgui\.ImGuiFocusedFlags\)\(\w\+\)/\1_\2/g' $files

# Hovered Flags
sed -i -e 's/cimgui\.HoveredFlags/cimgui\.ImGuiHoveredFlags/g' $files
sed -i -e 's/\(cimgui\.ImGuiHoveredFlags\)\(\w\+\)/\1_\2/g' $files

# Color Edit Flags
# TODO: COPY-PASTE them again (many things has changed
# API CHANGE!
sed -i -e 's/\(.*ColorEditFlags.*=.*\)/\/\/ \1/g' $files

# Table Flags
sed -i -e 's/cimgui\.TableFlags/cimgui\.ImGuiTableFlags/g' $files
sed -i -e 's/\(cimgui\.ImGuiTableFlags_NoBordersInBodyUntilResize\)TableFlags/\1/g' $files

# Table Row Flags
sed -i -e 's/cimgui\.TableRowFlags/cimgui\.ImGuiTableRowFlags/g' $files

# Table Column Flags
sed -i -e 's/cimgui\.TableColumnFlags/cimgui\.ImGuiTableColumnFlags/g' $files

# ImPlotFlags:
# disable flags that are not present:
# API CHANGE!
sed -i -e 's/\(.*cimgui\.ImPlotFlags_NoMousePos.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.ImPlotFlags_NoHighlight.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.ImPlotFlags_YAxis2.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.ImPlotFlags_YAxis3.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.ImPlotFlags_Query.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.ImPlotFlags_AntiAliased.*\)/\/\/ \1/g' $files

# Plot Axis Flags
# API CHANGE!
sed -i -e 's/\(.*cimgui\.ImPlotAxisFlags_LogScale.*\)/\/\/ \1/g' $files
sed -i -e 's/\(.*cimgui\.ImPlotAxisFlags_Time.*\)/\/\/ \1/g' $files

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
sed -i -e 's/\.AddLine/\.AddLineV/g' $files
sed -i -e 's/\.AddRect/\.AddRectV/g' $files

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

