// Package main demonstrates how to use the Code Editor widget.
package main

import (
	"fmt"

	"github.com/AllenDang/giu"
)

var (
	editor         *giu.CodeEditorWidget
	palettes             = []string{"Dark", "Light", "Mariana", "Retro Blue"}
	currentPalette int32 = 0
)

// errMarkers imgui.ErrorMarkers

func loop() {
	giu.SingleWindow().Layout(
		giu.Row(
			giu.Button("Get Text").OnClick(func() {
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
			giu.Button("Set Text").OnClick(func() {
				editor.Text("Set text")
			}),
			//nolint:gocritic,wsl // this should be here for documentation and as a reminder.
			giu.Button("Set Error Marker").OnClick(func() {
				panic("implement me!")
				// errMarkers.Clear()
				// errMarkers.Insert(1, "Error message")
				// fmt.Println("ErrMarkers Size:", errMarkers.Size())
				// editor.ErrorMarkers(errMarkers)
			}),
			giu.Combo("Palette", palettes[currentPalette], palettes, &currentPalette).OnChange(func() {
				editor.Palette(giu.CodeEditorPalette(currentPalette))
			}),
		),
		editor,
	)
}

func main() {
	wnd := giu.NewMasterWindow("Code Editor", 800, 600, 0)

	//nolint:gocritic // should be here for doc
	// errMarkers = imgui.NewErrorMarkers()

	editor = giu.CodeEditor().
		ShowWhitespaces(false).
		TabSize(2).
		Text("select * from greeting\nwhere date > current_timestamp\norder by date").
		LanguageDefinition(giu.LanguageDefinitionSQL).
		Border(true)

	wnd.Run(loop)
}
