//nolint:gocritic,govet,unused // this file is TODO. We don't want commentedOutCode linter issues here yet.
package giu

import (
	"fmt"
)

// LanguageDefinition represents code editor's language definition.
type LanguageDefinition byte

// language definitions:.
const (
	LanguageDefinitionSQL LanguageDefinition = iota
	LanguageDefinitionCPP
	LanguageDefinitionLua
	LanguageDefinitionC
)

var _ Disposable = &codeEditorState{}

type codeEditorState struct {
	// editor imgui.TextEditor
}

// Dispose implements Disposable interface.
func (s *codeEditorState) Dispose() {
	// noop
}

// static check if code editor implements Widget interface.
var _ Widget = &CodeEditorWidget{}

// CodeEditorWidget represents imgui.TextEditor.
type CodeEditorWidget struct {
	title ID
	width,
	height float32
	border bool
}

// CodeEditor creates new code editor widget.
func CodeEditor() *CodeEditorWidget {
	panic("Code Editor is not implemented yet!")

	return &CodeEditorWidget{
		title: GenAutoID("##CodeEditor"),
	}
}

// ID allows to manually set editor's ID.
// It isn't necessary to use it in a normal conditions.

func (ce *CodeEditorWidget) ID(id ID) *CodeEditorWidget {
	ce.title = id
	return ce
}

// ShowWhitespaces sets if whitespace is shown in code editor.
func (ce *CodeEditorWidget) ShowWhitespaces(s bool) *CodeEditorWidget {
	// ce.getState().editor.SetShowWhitespaces(s)
	return ce
}

// TabSize sets editor's tab size.
func (ce *CodeEditorWidget) TabSize(size int) *CodeEditorWidget {
	// ce.getState().editor.SetTabSize(size)
	return ce
}

// LanguageDefinition sets code editor language definition.

func (ce *CodeEditorWidget) LanguageDefinition(definition LanguageDefinition) *CodeEditorWidget {
	// s := ce.getState()
	lookup := map[LanguageDefinition]func(){
		// LanguageDefinitionSQL: s.editor.SetLanguageDefinitionSQL,
		// LanguageDefinitionCPP: s.editor.SetLanguageDefinitionCPP,
		// LanguageDefinitionLua: s.editor.SetLanguageDefinitionLua,
		// LanguageDefinitionC:   s.editor.SetLanguageDefinitionC,
	}

	setter, correctDefinition := lookup[definition]
	if !correctDefinition {
		panic(fmt.Sprintf("giu/CodeEditor.go: unknown language definition %d", definition))
	}

	setter()

	return ce
}

// Text sets editor's text.
func (ce *CodeEditorWidget) Text(str string) *CodeEditorWidget {
	// ce.getState().editor.SetText(str)
	return ce
}

// ErrorMarkers sets error markers.
// func (ce *CodeEditorWidget) ErrorMarkers(markers imgui.ErrorMarkers) *CodeEditorWidget {
//	ce.getState().editor.SetErrorMarkers(markers)
// return ce
//}

// HandleKeyboardInputs sets if editor should handle keyboard input.
func (ce *CodeEditorWidget) HandleKeyboardInputs(b bool) *CodeEditorWidget {
	// ce.getState().editor.SetHandleKeyboardInputs(b)
	return ce
}

// Size sets editor's size.
func (ce *CodeEditorWidget) Size(w, h float32) *CodeEditorWidget {
	ce.width, ce.height = w, h
	return ce
}

// Border sets editors borders.
func (ce *CodeEditorWidget) Border(border bool) *CodeEditorWidget {
	ce.border = border
	return ce
}

// HasSelection returns true if some text is selected.
func (ce *CodeEditorWidget) HasSelection() bool {
	// return ce.getState().editor.HasSelection()
	return false
}

// GetSelectedText returns selected text.
func (ce *CodeEditorWidget) GetSelectedText() string {
	// return ce.getState().editor.GetSelectedText()
	return ""
}

// GetText returns whole text from editor.
func (ce *CodeEditorWidget) GetText() string {
	//	return ce.getState().editor.GetText()
	return ""
}

// GetCurrentLineText returns current line.
func (ce *CodeEditorWidget) GetCurrentLineText() string {
	//	return ce.getState().editor.GetCurrentLineText()
	return ""
}

// GetCursorPos returns cursor position.
// (in characters).
func (ce *CodeEditorWidget) GetCursorPos() (x, y int) {
	//	return ce.getState().editor.GetCursorPos()
	return 0, 0
}

// GetSelectionStart returns star pos of selection.
func (ce *CodeEditorWidget) GetSelectionStart() (x, y int) {
	//	return ce.getState().editor.GetSelectionStart()
	return 0, 0
}

// InsertText inserts the `text`.
func (ce *CodeEditorWidget) InsertText(text string) {
	//	ce.getState().editor.InsertText(text)
}

// GetWordUnderCursor returns the word under the cursor.
func (ce *CodeEditorWidget) GetWordUnderCursor() string {
	//	return ce.getState().editor.GetWordUnderCursor()
	return ""
}

// SelectWordUnderCursor selects the word under cursor.
func (ce *CodeEditorWidget) SelectWordUnderCursor() {
	// ce.getState().editor.SelectWordUnderCursor()
}

// IsTextChanged returns true if the editable text was changed in the frame.
func (ce *CodeEditorWidget) IsTextChanged() bool {
	//	return ce.getState().editor.IsTextChanged()
	return false
}

// GetScreenCursorPos returns cursor position on the screen.
// (in pixels).
func (ce *CodeEditorWidget) GetScreenCursorPos() (x, y int) {
	//	return ce.getState().editor.GetScreenCursorPos()
	return 0, 0
}

// Copy copies selection.
func (ce *CodeEditorWidget) Copy() {
	// ce.getState().editor.Copy()
}

// Cut cuts selection.
func (ce *CodeEditorWidget) Cut() {
	// ce.getState().editor.Cut()
}

// Paste does the same as Ctrl+V.
func (ce *CodeEditorWidget) Paste() {
	// ce.getState().editor.Paste()
}

// Delete deletes the selection.
func (ce *CodeEditorWidget) Delete() {
	// ce.getState().editor.Delete()
}

// Build implements Widget interface.
func (ce *CodeEditorWidget) Build() {
	// s := ce.getState()

	// register text in font atlas
	// Context.FontAtlas.RegisterString(s.editor.GetText())

	// build editor
	// s.editor.Render(ce.title, imgui.Vec2{X: ce.width, Y: ce.height}, ce.border)
}

func (ce *CodeEditorWidget) getState() (state *codeEditorState) {
	if state = GetState[codeEditorState](Context, ce.title.String()); state == nil {
		state = &codeEditorState{
			// editor: imgui.NewTextEditor(),
		}

		SetState(Context, ce.title.String(), state)
	}

	return state
}
