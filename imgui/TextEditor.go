package imgui

// #include "TextEditorWrapper.h"
import "C"

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
	return C.GoString(C.IggTextEditorGetText(t.handle()))
}

func (t TextEditor) HasSelection() bool {
	return C.IggTextEditorHasSelection(t.handle()) != 0
}

func (t TextEditor) GetSelectedText() string {
	return C.GoString(C.IggTextEditorGetSelectedText(t.handle()))
}

func (t TextEditor) GetCurrentLineText() string {
	return C.GoString(C.IggTextEditorGetCurrentLineText(t.handle()))
}

func (t TextEditor) IsTextChanged() bool {
	return C.IggTextEditorIsTextChanged(t.handle()) != 0
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
