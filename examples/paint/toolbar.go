package main

import (
	"fmt"
	"image/color"
	"path"

	"github.com/AllenDang/cimgui-go/imgui"
	g "github.com/AllenDang/giu"
)

var (
	pickerRefColor color.RGBA
	penButtonImg   = &g.ReflectiveBoundTexture{}
	fillButtonImg  = &g.ReflectiveBoundTexture{}
	undoButtonImg  = &g.ReflectiveBoundTexture{}
	clearButtonImg = &g.ReflectiveBoundTexture{}
	openButtonImg  = &g.ReflectiveBoundTexture{}
	saveButtonImg  = &g.ReflectiveBoundTexture{}
	brushButtonImg = &g.ReflectiveBoundTexture{}
	toolbarInited  = false
)

func initToolbar() {
	p := func(img string) string { return path.Join("icons", img) }
	penButtonImg.SetSurfaceFromFile(p("pencil.png"), false)
	fillButtonImg.SetSurfaceFromFile(p("paint-bucket.png"), false)
	undoButtonImg.SetSurfaceFromFile(p("undo.png"), false)
	brushButtonImg.SetSurfaceFromFile(p("brush.png"), false)
	clearButtonImg.SetSurfaceFromFile(p("clear.png"), false)
	openButtonImg.SetSurfaceFromFile(p("open-folder.png"), false)
	saveButtonImg.SetSurfaceFromFile(p("floppy-disk.png"), false)
	toolbarInited = true
}

func ShowToolbar() g.Widget {
	if !toolbarInited {
		initToolbar()
	}
	return g.Child().Size(-1, TOOLBAR_H).Layout(
		ButtonColorMaker(),
	)
}

func colorPopup(ce *color.RGBA, fe g.ColorEditFlags) {

	p := g.ToVec4Color(pickerRefColor)
	pcol := []float32{p.X, p.Y, p.Z, p.W}

	if imgui.BeginPopup("Custom Color") {
		c := g.ToVec4Color(*ce)
		col := [4]float32{
			c.X,
			c.Y,
			c.Z,
			c.W,
		}
		refCol := pcol

		if imgui.ColorPicker4V(
			g.Context.FontAtlas.RegisterString("##COLOR_POPUP##me"),
			&col,
			imgui.ColorEditFlags(fe),
			refCol,
		) {
			*ce = g.Vec4ToRGBA(imgui.Vec4{
				X: col[0],
				Y: col[1],
				Z: col[2],
				W: col[3],
			})
		}
		imgui.EndPopup()
	}

}

func ButtonColorMaker() *g.RowWidget {
	start_ul := imgui.CursorPos()
	sz := imgui.Vec2{}
	return g.Row(g.Custom(func() {
		for i := range defaultColors {
			if i%2 == 0 {
				col := g.ToVec4Color(defaultColors[i])
				if imgui.ColorButtonV(fmt.Sprintf("%d##cur_color%d", i, i), col, 0, imgui.Vec2{X: 0, Y: 0}) {
					current_color = defaultColors[i]
				}
				sz = imgui.ItemRectSize()
				imgui.SameLineV(0, 0)
			}
		}
		col := g.ToVec4Color(current_color)
		if imgui.ColorButtonV(fmt.Sprintf("##CHOSENcur_color%d", current_color), col, 0, sz.Mul(2.0)) {
			pickerRefColor = current_color
			imgui.OpenPopupStr("Custom Color")
		}
		colorPopup(&current_color, g.ColorEditFlagsNoAlpha)
		imgui.SameLine()
		if imgui.ImageButton("##pen_tool", penButtonImg.Texture().ID(), sz.Mul(1.7)) {
			current_tool = 0
		}
		imgui.SameLine()
		if imgui.ImageButton("##fill_tool", fillButtonImg.Texture().ID(), sz.Mul(1.7)) {
			current_tool = 1
		}
		imgui.SameLine()
		imgui.ImageButton("##undo_tool", undoButtonImg.Texture().ID(), sz.Mul(1.7))
		imgui.SameLine()
		if imgui.ImageButton("##clear_tool", clearButtonImg.Texture().ID(), sz.Mul(1.7)) {
			canvas.Backend.ForceRelease()
			canvas, _ = NewCanvas(canvasDetectedHeight)
		}
		imgui.SameLine()
		imgui.ImageButton("##open_tool", openButtonImg.Texture().ID(), sz.Mul(1.7))
		imgui.SameLine()
		imgui.ImageButton("##save_tool", saveButtonImg.Texture().ID(), sz.Mul(1.7))
		imgui.SameLine()
		imgui.ImageButton("##brush_tool", brushButtonImg.Texture().ID(), sz.Mul(1.7))
		imgui.SetCursorPos(start_ul)
		for i := range defaultColors {
			if i%2 != 0 {
				col := g.ToVec4Color(defaultColors[i])
				if imgui.ColorButtonV(fmt.Sprintf("%d##cur_color%d", i, i), col, 0, imgui.Vec2{X: 0, Y: 0}) {
					current_color = defaultColors[i]
				}
				imgui.SameLineV(0, 0)
			}
		}
	},
	),
	)
}
