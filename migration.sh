git stash
files=`find . -iname \*go`
sed -ie 's/imgui/cimgui/g' $files
go get github.com/AllenDang/cimgui-go@158164eb30c79c00a3c393a1d6642609f2f2e206
go mod tidy
sed -ie 's/cimgui\.StyleColorID/cimgui\.ImGuiCol/g' $files
sed -ie 's/cimgui\.StyleVarID/cimgui\.ImGuiStyleVar/g' $files
