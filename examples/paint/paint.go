package main

import (
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/AllenDang/cimgui-go/imgui"
	g "github.com/AllenDang/giu"
)

var (
	imageScaleX = float32(1.0)
	imageScaleY = float32(1.0)
	linkedScale = true
	showWindow  = true
	wnd         *g.MasterWindow
)

const (
	WINDOW_W  = 1280
	WINDOW_H  = 720
	TOOLBAR_H = 100
)

func loop() {
	if !showWindow {
		wnd.SetShouldClose(true)
	}

	g.PushColorWindowBg(color.RGBA{30, 30, 30, 255})
	g.Window("GIU Paint").IsOpen(&showWindow).Pos(10, 30).Size(WINDOW_W, WINDOW_H).Flags(g.WindowFlagsNoResize).Layout(
		ShowToolbar(),
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
