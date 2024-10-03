// Package main demonstrate use of advanced image object via paint clone
package main

import (
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/AllenDang/cimgui-go/imgui"

	g "github.com/AllenDang/giu"
)

var (
	showWindow = true
	wnd        *g.MasterWindow
)

const (
	windowWidth   = 1280
	windowHeight  = 720
	toolbarHeight = 100
)

func loop() {
	if !showWindow {
		wnd.SetShouldClose(true)
	}

	g.PushColorWindowBg(color.RGBA{30, 30, 30, 255})
	g.Window("GIU Paint").IsOpen(&showWindow).Pos(10, 30).Size(windowWidth, windowHeight).Flags(g.WindowFlagsNoResize).Layout(
		showToolbar(),
		g.Separator(),
		CanvasRow(),
	)
	g.PopStyleColor()
}

func noOSDecoratedWindowsConfig() g.MasterWindowFlags {
	imgui.CreateContext()
	io := imgui.CurrentIO()
	io.SetConfigViewportsNoAutoMerge(true)
	io.SetConfigViewportsNoDefaultParent(true)
	io.SetConfigWindowsMoveFromTitleBarOnly(true)

	return g.MasterWindowFlagsHidden | g.MasterWindowFlagsTransparent | g.MasterWindowFlagsFrameless
}

func main() {
	// This prepare creating a fully imgui window with no native decoration.
	// Flags are to be used with NewMasterWindow.
	// Should NOT use SingleLayoutWindow !
	mwFlags := noOSDecoratedWindowsConfig()

	wnd = g.NewMasterWindow("Paint Demo", 1280, 720, mwFlags)
	wnd.SetTargetFPS(60)
	wnd.Run(loop)
}
