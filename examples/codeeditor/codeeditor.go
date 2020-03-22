package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

var (
	editor imgui.TextEditor
)

func loop() {
	g.SingleWindow("Code Editor", g.Layout{
		g.Line(
			g.Button("Get Text", func() {
				if editor.HasSelection() {
					fmt.Println(editor.GetSelectedText())
				} else {
					fmt.Println(editor.GetText())
				}

				fmt.Println("Current line is", editor.GetCurrentLineText())
			}),
			g.Button("Set Text", func() {
				editor.SetText("Set text")
			}),
		),
		g.Custom(func() {
			editor.Render("Hello", imgui.Vec2{X: 0, Y: 0}, true)
		}),
	})
}

func main() {
	editor = imgui.NewTextEditor()
	editor.SetShowWhitespaces(false)
	editor.SetTabSize(2)
	editor.SetText("select * from greeting\nwhere date > current_timestamp\norder by date")
	editor.SetLanguageDefinitionSQL()

	wnd := g.NewMasterWindow("Code Editor", 800, 600, 0, nil)
	wnd.Main(loop)
}
