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

type codeEditorState struct {
	editor imgui.TextEditor
}

func (s *codeEditorState) Dispose() {
	// noop
}

type CodeEditorWidget struct {
	title string
	width,
	height float32
	border bool
}

func CodeEditor(title string) *CodeEditorWidget {
	return &CodeEditorWidget{
		title: title,
	}
}

func (ce *CodeEditorWidget) ShowWhitespaces(s bool) *CodeEditorWidget {
	ce.getState().editor.SetShowWhitespaces(s)
	return ce
}

func (ce *CodeEditorWidget) TabSize(size int) *CodeEditorWidget {
	ce.getState().editor.SetTabSize(size)
	return ce
}

func (ce *CodeEditorWidget) LanguageDefinition(definition LanguageDefinition) *CodeEditorWidget {
	s := ce.getState()
	lookup := map[LanguageDefinition]func(){
		LanguageDefinitionSQL: s.editor.SetLanguageDefinitionSQL,
		LanguageDefinitionCPP: s.editor.SetLanguageDefinitionCPP,
		LanguageDefinitionLua: s.editor.SetLanguageDefinitionLua,
		LanguageDefinitionC:   s.editor.SetLanguageDefinitionC,
	}

	setter, correctDefinition := lookup[definition]
	if !correctDefinition {
		panic(fmt.Sprintf("giu/CodeEditor.go: unknown language definition %d", definition))
	}

	setter()

	return ce
}

func (ce *CodeEditorWidget) Text(str string) *CodeEditorWidget {
	ce.getState().editor.SetText(str)
	return ce
}

func (ce *CodeEditorWidget) ErrorMarkers(markers imgui.ErrorMarkers) *CodeEditorWidget {
	ce.getState().editor.SetErrorMarkers(markers)
	return ce
}

func (ce *CodeEditorWidget) HandleKeyboardInputs(b bool) *CodeEditorWidget {
	ce.getState().editor.SetHandleKeyboardInputs(b)
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
	return ce.getState().editor.HasSelection()
}

func (ce *CodeEditorWidget) GetSelectedText() string {
	return ce.getState().editor.GetSelectedText()
}

func (ce *CodeEditorWidget) GetText() string {
	return ce.getState().editor.GetText()
}

func (ce *CodeEditorWidget) GetCurrentLineText() string {
	return ce.getState().editor.GetCurrentLineText()
}

func (ce *CodeEditorWidget) GetCursorPos() (int, int) {
	return ce.getState().editor.GetCursorPos()
}

func (ce *CodeEditorWidget) GetSelectionStart() (int, int) {
	return ce.getState().editor.GetSelectionStart()
}

func (ce *CodeEditorWidget) InsertText(text string) {
	ce.getState().editor.InsertText(text)
}

func (ce *CodeEditorWidget) GetWordUnderCursor() string {
	return ce.getState().editor.GetWordUnderCursor()
}

func (ce *CodeEditorWidget) SelectWordUnderCursor() {
	ce.getState().editor.SelectWordUnderCursor()
}

func (ce *CodeEditorWidget) IsTextChanged() bool {
	return ce.getState().editor.IsTextChanged()
}

func (ce *CodeEditorWidget) GetScreenCursorPos() (int, int) {
	return ce.getState().editor.GetScreenCursorPos()
}

func (ce *CodeEditorWidget) Copy() {
	ce.getState().editor.Copy()
}

func (ce *CodeEditorWidget) Cut() {
	ce.getState().editor.Cut()
}

func (ce *CodeEditorWidget) Paste() {
	ce.getState().editor.Paste()
}

func (ce *CodeEditorWidget) Delete() {
	ce.getState().editor.Delete()
}

func (ce *CodeEditorWidget) Build() {
	s := ce.getState()

	// register text in font atlas
	tStr(s.editor.GetText())

	// build editor
	s.editor.Render(ce.title, imgui.Vec2{X: ce.width, Y: ce.height}, ce.border)
}

func (ce *CodeEditorWidget) getState() (state *codeEditorState) {
	if s := Context.GetState(ce.title); s == nil {
		state = &codeEditorState{
			editor: imgui.NewTextEditor(),
		}

		Context.SetState(ce.title, state)
	} else {
		state = s.(*codeEditorState)
	}

	return state
}
