package main

import "github.com/AllenDang/giu"

var (
	checkbox1 = true
	checkbox2 = true
)

func loop() {
	giu.Window("Window 2").
		RegisterKeyboardShortcuts(
			giu.WindowShortcut{Key: giu.KeyZ, Modifier: giu.ModControl, Callback: func() { checkbox2 = !checkbox2 }},
		).
		Layout(
			giu.Checkbox("Press Ctrl+C to change my state - I'm a global shortcut", &checkbox1),
			giu.Checkbox("Press Ctrl+Z to change my state - I'm a local shortcut", &checkbox2),
		)

	giu.Window("Window 1").
		Layout(
			giu.Checkbox("Press Ctrl+C to change my state - I'm a global shortcut", &checkbox1),
		)
}

func main() {
	wnd := giu.NewMasterWindow("keyboard shortcuts", 640, 480, 0).
		RegisterKeyboardShortcuts(
			giu.WindowShortcut{
				Key:      giu.KeyC,
				Modifier: giu.ModControl,
				Callback: func() { checkbox1 = !checkbox1 }},
		)

	wnd.Run(loop)
}
