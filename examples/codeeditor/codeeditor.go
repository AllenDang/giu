package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	editor     *g.CodeEditorWidget
	errMarkers imgui.ErrorMarkers
)

func loop() {
	g.SingleWindow("Code Editor").Layout(
		g.Row(
			g.Button("Get Text").OnClick(func() {
				if editor.HasSelection() {
					fmt.Println(editor.GetSelectedText())
				} else {
					fmt.Println(editor.GetText())
				}

				column, line := editor.GetCursorPos()
				fmt.Println("Cursor pos:", column, line)

				column, line = editor.GetSelectionStart()
				fmt.Println("Selection start:", column, line)

				fmt.Println("Current line is", editor.GetCurrentLineText())
			}),
			g.Button("Set Text").OnClick(func() {
				editor.SetText("Set text")
			}),
			g.Button("Set Error Marker").OnClick(func() {
				errMarkers.Clear()
				errMarkers.Insert(1, "Error message")
				fmt.Println("ErrMarkers Size:", errMarkers.Size())

				editor.SetErrorMarkers(errMarkers)
			}),
		),
		g.Custom(func() {
			editor.Render()
		}),
	)
}

func main() {
	errMarkers = imgui.NewErrorMarkers()

	editor = g.CodeEditor("Code Editor", 0, 0, true)
	editor.SetShowWhitespaces(false)
	editor.SetTabSize(2)
	editor.SetText("select * from greeting\nwhere date > current_timestamp\norder by date")
	editor.SetLanguageDefinitionSQL()

	wnd := g.NewMasterWindow("Code Editor", 800, 600, 0)
	wnd.Run(loop)
}
