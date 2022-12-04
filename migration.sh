files=`find . -iname \*go`
sed -ie 's/imgui/cimgui/g' $files
