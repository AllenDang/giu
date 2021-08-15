package giu

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
)

type LanguageDefinition byte

const (
	LanguageDefinitionSQL LanguageDefinition = iota
	LanguageDefinitionCPP
	LanguageDefinitionLua
	LanguageDefinitionC
)

type CodeEditorWidget struct {
	title string
	width,
	height float32
	border bool
	editor imgui.TextEditor
}

func CodeEditor(title string) *CodeEditorWidget {
	return &CodeEditorWidget{
		title:  title,
		editor: imgui.NewTextEditor(),
	}
}

func (ce *CodeEditorWidget) ShowWhitespaces(s bool) *CodeEditorWidget {
	ce.editor.SetShowWhitespaces(s)
	return ce
}

func (ce *CodeEditorWidget) TabSize(size int) *CodeEditorWidget {
	ce.editor.SetTabSize(size)
	return ce
}

func (ce *CodeEditorWidget) LanguageDefinition(definition LanguageDefinition) *CodeEditorWidget {
	lookup := map[LanguageDefinition]func(){
		LanguageDefinitionSQL: ce.editor.SetLanguageDefinitionSQL,
		LanguageDefinitionCPP: ce.editor.SetLanguageDefinitionCPP,
		LanguageDefinitionLua: ce.editor.SetLanguageDefinitionLua,
		LanguageDefinitionC:   ce.editor.SetLanguageDefinitionC,
	}

	setter, correctDefinition := lookup[definition]
	if !correctDefinition {
		panic(fmt.Sprintf("giu/CodeEditor.go: unknown language definition %d", definition))
	}

	setter()

	return ce
}

func (ce *CodeEditorWidget) Text(str string) *CodeEditorWidget {
	ce.editor.SetText(str)
	return ce
}

func (ce *CodeEditorWidget) ErrorMarkers(markers imgui.ErrorMarkers) *CodeEditorWidget {
	ce.editor.SetErrorMarkers(markers)
	return ce
}

func (ce *CodeEditorWidget) HandleKeyboardInputs(b bool) *CodeEditorWidget {
	ce.editor.SetHandleKeyboardInputs(b)
	return ce
}

func (ce *CodeEditorWidget) Size(w, h float32) *CodeEditorWidget {
	ce.width, ce.height = w, h
	return ce
}

func (ce *CodeEditorWidget) Border(border bool) *CodeEditorWidget {
	ce.border = border
	return ce
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

func (ce *CodeEditorWidget) Build() {
	// register text in font atlas
	tStr(ce.editor.GetText())

	// build editor
	ce.editor.Render(ce.title, imgui.Vec2{X: ce.width, Y: ce.height}, ce.border)
}
