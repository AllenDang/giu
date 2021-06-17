package main

import "github.com/AllenDang/giu"

var (
	checkbox1 = true
	checkbox2 = true
)

func loop() {
	giu.RegisterKeyboardShortcut([]giu.Shortcut{
		{giu.KeyC, giu.ModControl, func() { checkbox1 = !checkbox1 }, giu.GlobalShortcut},
	}...)

	giu.Window("Window 1").
		Layout(
			giu.Checkbox("Press Ctrl+C to change my state - I'm a global shortcut", &checkbox1),
		)

	giu.Window("Window 2").
		RegisterKeyboardShortcuts([]giu.Shortcut{
			{giu.KeyZ, giu.ModControl, func() { checkbox2 = !checkbox2 }, giu.LocalShortcut},
		}...).
		Layout(
			giu.Checkbox("Press Ctrl+C to change my state - I'm a global shortcut", &checkbox1),
			giu.Checkbox("Press Ctrl+Z to change my state - I'm a local shortcut", &checkbox2),
		)
}

func main() {
	wnd := giu.NewMasterWindow("keyboard shortcuts", 640, 480, 0)
	wnd.Run(loop)
}
