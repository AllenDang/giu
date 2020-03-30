package imgui

// #include "stdlib.h"
// #include "TextEditorWrapper.h"
import "C"
import "unsafe"

type TextEditor uintptr

func NewTextEditor() TextEditor {
	handle := C.IggNewTextEditor()
	return TextEditor(handle)
}

func (t TextEditor) handle() C.IggTextEditor {
	return C.IggTextEditor(t)
}

func (t TextEditor) Render(title string, size Vec2, border bool) {
	titleArg, titleFn := wrapString(title)
	defer titleFn()

	sizeArg, _ := size.wrapped()

	C.IggTextEditorRender(t.handle(), titleArg, sizeArg, castBool(border))
}

func (t TextEditor) SetShowWhitespaces(show bool) {
	C.IggTextEditorSetShowWhitespaces(t.handle(), castBool(show))
}

func (t TextEditor) SetTabSize(size int) {
	C.IggTextEditorSetTabSize(t.handle(), C.int(size))
}

func (t TextEditor) SetText(text string) {
	textArg, textFn := wrapString(text)
	defer textFn()

	C.IggTextEditorSetText(t.handle(), textArg)
}

func (t TextEditor) GetText() string {
	str := C.IggTextEditorGetText(t.handle())
	defer C.free(unsafe.Pointer(str))

	return C.GoString(str)
}

func (t TextEditor) HasSelection() bool {
	return C.IggTextEditorHasSelection(t.handle()) != 0
}

func (t TextEditor) GetSelectedText() string {
	str := C.IggTextEditorGetSelectedText(t.handle())
	defer C.free(unsafe.Pointer(str))

	return C.GoString(str)
}

func (t TextEditor) GetCurrentLineText() string {
	str := C.IggTextEditorGetCurrentLineText(t.handle())
	defer C.free(unsafe.Pointer(str))

	return C.GoString(str)
}

func (t TextEditor) IsTextChanged() bool {
	return C.IggTextEditorIsTextChanged(t.handle()) != 0
}

func (t TextEditor) GetCursorPos() (int, int) {
	var column int
	var line int

	C.IggTextEditorGetCursorPos(t.handle(), (*C.int)(unsafe.Pointer(&column)), (*C.int)(unsafe.Pointer(&line)))

	return column, line
}

func (t TextEditor) GetSelectionStart() (int, int) {
	var column int
	var line int

	C.IggTextEditorGetSelectionStart(t.handle(), (*C.int)(unsafe.Pointer(&column)), (*C.int)(unsafe.Pointer(&line)))

	return column, line
}

func (t TextEditor) SetLanguageDefinitionSQL() {
	C.IggTextEditorSetLanguageDefinitionSQL(t.handle())
}

func (t TextEditor) SetLanguageDefinitionCPP() {
	C.IggTextEditorSetLanguageDefinitionCPP(t.handle())
}

func (t TextEditor) SetLanguageDefinitionC() {
	C.IggTextEditorSetLanguageDefinitionC(t.handle())
}

func (t TextEditor) SetLanguageDefinitionLua() {
	C.IggTextEditorSetLanguageDefinitionLua(t.handle())
}

type ErrorMarkers uintptr

func NewErrorMarkers() ErrorMarkers {
	handle := C.IggTextEditorNewErrorMarkers()
	return ErrorMarkers(handle)
}

func (e ErrorMarkers) handle() C.IggTextEditorErrorMarkers {
	return C.IggTextEditorErrorMarkers(e)
}

func (e ErrorMarkers) Insert(pos int, errMsg string) {
	errMsgArg, errMsgFn := wrapString(errMsg)
	defer errMsgFn()

	C.IggTextEditorErrorMarkersInsert(e.handle(), C.int(pos), errMsgArg)
}

func (e ErrorMarkers) Clear() {
	C.IggTextEditorErrorMarkersClear(e.handle())
}

func (e ErrorMarkers) Size() uint {
	return uint(C.IggTextEditorErrorMarkersSize(e.handle()))
}

func (t TextEditor) SetErrorMarkers(markers ErrorMarkers) {
	C.IggTextEditorSetErrorMarkers(t.handle(), markers.handle())
}
