package giu

import "github.com/AllenDang/imgui-go"

type CodeEditorWidget struct {
	title  string
	editor imgui.TextEditor
}

func CodeEditor(title string) *CodeEditorWidget {
	return &CodeEditorWidget{
		title:  title,
		editor: imgui.NewTextEditor(),
	}
}

func (ce *CodeEditorWidget) SetShowWhitespaces(s bool) {
	ce.editor.SetShowWhitespaces(s)
}

func (ce *CodeEditorWidget) SetTabSize(size int) {
	ce.editor.SetTabSize(size)
}

func (ce *CodeEditorWidget) SetLanguageDefinitionSQL() {
	ce.editor.SetLanguageDefinitionSQL()
}

func (ce *CodeEditorWidget) SetLanguageDefinitionCPP() {
	ce.editor.SetLanguageDefinitionCPP()
}

func (ce *CodeEditorWidget) SetLanguageDefinitionLua() {
	ce.editor.SetLanguageDefinitionLua()
}

func (ce *CodeEditorWidget) SetLanguageDefinitionC() {
	ce.editor.SetLanguageDefinitionC()
}

func (ce *CodeEditorWidget) SetText(str string) {
	ce.editor.SetText(str)
}

func (ce *CodeEditorWidget) HasSelection() bool {
	return ce.editor.HasSelection()
}

func (ce *CodeEditorWidget) GetSelectedText() string {
	return ce.editor.GetSelectedText()
}

func (ce *CodeEditorWidget) GetText() string {
	return ce.editor.GetText()
}

func (ce *CodeEditorWidget) GetCurrentLineText() string {
	return ce.editor.GetCurrentLineText()
}

func (ce *CodeEditorWidget) GetCursorPos() (int, int) {
	return ce.editor.GetCursorPos()
}

func (ce *CodeEditorWidget) GetSelectionStart() (int, int) {
	return ce.editor.GetSelectionStart()
}

func (ce *CodeEditorWidget) SetErrorMarkers(markers imgui.ErrorMarkers) {
	ce.editor.SetErrorMarkers(markers)
}

func (ce *CodeEditorWidget) InsertText(text string) {
	ce.editor.InsertText(text)
}

func (ce *CodeEditorWidget) GetWordUnderCursor() string {
	return ce.editor.GetWordUnderCursor()
}

func (ce *CodeEditorWidget) SelectWordUnderCursor() {
	ce.editor.SelectWordUnderCursor()
}

func (ce *CodeEditorWidget) IsTextChanged() bool {
	return ce.editor.IsTextChanged()
}

func (ce *CodeEditorWidget) SetHandleKeyboardInputs(b bool) {
	ce.editor.SetHandleKeyboardInputs(b)
}

func (ce *CodeEditorWidget) GetScreenCursorPos() (int, int) {
	return ce.editor.GetScreenCursorPos()
}

func (ce *CodeEditorWidget) Copy() {
	ce.editor.Copy()
}

func (ce *CodeEditorWidget) Cut() {
	ce.editor.Cut()
}

func (ce *CodeEditorWidget) Paste() {
	ce.editor.Paste()
}

func (ce *CodeEditorWidget) Delete() {
	ce.editor.Delete()
}

func (ce *CodeEditorWidget) Render(width, height float32, border bool) {
	tStr(ce.editor.GetText())
	ce.editor.Render(ce.title, imgui.Vec2{X: width, Y: height}, border)
}
