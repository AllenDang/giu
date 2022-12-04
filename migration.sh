git add migration.*
git commit --amend
git stash
files=`find . -iname \*go`

# switch to cimgui-go
sed -ie 's/imgui/cimgui/g' $files
go get github.com/AllenDang/cimgui-go@158164eb30c79c00a3c393a1d6642609f2f2e206
go mod tidy

# mainly StyleIDs.go
sed -ie 's/cimgui\.StyleColorID/cimgui\.ImGuiCol/g' $files
sed -ie 's/cimgui\.StyleVarID/cimgui\.ImGuiStyleVar/g' $files

# another types
sed -ie 's/cimgui\.DrawList/cimgui\.ImDrawList/g' $files
sed -ie 's/cimgui\.TextureID/cimgui\.ImTextureID/g' $files
sed -ie 's/cimgui\.Vec2/cimgui\.ImVec2/g' $files
sed -ie 's/cimgui\.Vec4/cimgui\.ImVec4/g' $files
sed -ie 's/cimgui\.Font/cimgui\.ImFont/g' $files
sed -ie 's/cimgui\.Condition/cimgui\.ImGuiCond/g' $files
sed -ie 's/cimgui\.InputTextCallback/cimgui\.ImGuiInputTextCallback/g' $files
sed -ie 's/cimgui\.Context/cimgui\.ImGuiContext/g' $files

# Context; TODO - check if nothing else is changed
sed -ie 's/cimgui\.IO()/cimgui\.GetIO()/g' $files
sed -ie 's/cimgui\.IO/cimgui\.ImGuiIO/g' $files

# flags
#
# input text:
sed -ie 's/cimgui\.InputTextFlags\(\w\+\)/cimgui\.ImGuiInputTextFlags_\1/g' $files
# API CHANGE!
sed -ie 's/^.*cimgui\.ImGuiInputTextFlags_AlwaysInsertMode.*//g' $files

# window flags
sed -ie 's/cimgui\.WindowFlags/cimgui\.ImGuiWindowFlags/g' $files
# type was int; change to cimgui.ImGuiWindowFlags
sed -ie 's/\(type WindowFlags \)int/\1cimgui.GLFWWindowFlags/g' $files
sed -ie 's/\(cimgui\.ImGuiWindowFlags\)\(\w\+\)/WindowFlags(\1_\2)/g' $files

sed -ie 's/^/\/\/ /g' Markdown.go
echo "package giu" >> Markdown.go
sed -ie 's/^/\/\/ /g' CodeEditor.go
echo "package giu" >> CodeEditor.go
sed -ie 's/^/\/\/ /g' MemoryEditor.go
echo "package giu" >> MemoryEditor.go
