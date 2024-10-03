package main

import (
	"embed"
	"fmt"
	"image/color"
	"io/fs"

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

//go:embed all:icons
var icons embed.FS

type assetLoadingInfo struct {
	file    string
	backend *g.ReflectiveBoundTexture
}

func assets() (fs.FS, error) {
	f, err := fs.Sub(icons, "icons")
	if err != nil {
		return nil, fmt.Errorf("error in assets: %w", err)
	}

	return f, nil
}

func loadAsset(path string, backend *g.ReflectiveBoundTexture) error {
	assets, _ := assets()

	file, err := assets.Open(path)
	if err != nil {
		return fmt.Errorf("LoadAsset: error opening image file %s: %w", file, err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("error closing image file: %s", file))
		}
	}()

	err = backend.SetSurfaceFromFsFile(file, false)
	if err != nil {
		return fmt.Errorf("error in loadAsset: %w", err)
	}

	return nil
}

var loadableAssets = []assetLoadingInfo{
	{file: "pencil.png", backend: penButtonImg},
	{file: "paint-bucket.png", backend: fillButtonImg},
	{file: "undo.png", backend: undoButtonImg},
	{file: "brush.png", backend: brushButtonImg},
	{file: "clear.png", backend: clearButtonImg},
	{file: "open-folder.png", backend: openButtonImg},
	{file: "floppy-disk.png", backend: saveButtonImg},
}

func initToolbar() error {
	for _, info := range loadableAssets {
		if err := loadAsset(info.file, info.backend); err != nil {
			return err
		}
	}

	toolbarInited = true

	return nil
}

func showToolbar() g.Widget {
	if !toolbarInited {
		_ = initToolbar()
	}

	return g.Child().Size(-1, toolbarHeight).Layout(
		buttonColorMaker(),
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

func buttonColorMaker() *g.RowWidget {
	startUl := imgui.CursorPos()
	sz := imgui.Vec2{}

	return g.Row(g.Custom(func() {
		for i := range defaultColors {
			if i%2 == 0 {
				col := g.ToVec4Color(defaultColors[i])
				if imgui.ColorButtonV(fmt.Sprintf("%d##cur_color%d", i, i), col, 0, imgui.Vec2{X: 0, Y: 0}) {
					currentColor = defaultColors[i]
				}

				sz = imgui.ItemRectSize()
				imgui.SameLineV(0, 0)
			}
		}

		col := g.ToVec4Color(currentColor)
		if imgui.ColorButtonV(fmt.Sprintf("##CHOSENcur_color%d", currentColor), col, 0, sz.Mul(2.0)) {
			pickerRefColor = currentColor

			imgui.OpenPopupStr("Custom Color")
		}

		colorPopup(&currentColor, g.ColorEditFlagsNoAlpha)
		imgui.SameLine()

		if imgui.ImageButton("##pen_tool", penButtonImg.Texture().ID(), sz.Mul(1.7)) {
			currentTool = 0
		}

		imgui.SameLine()

		if imgui.ImageButton("##fill_tool", fillButtonImg.Texture().ID(), sz.Mul(1.7)) {
			currentTool = 1
		}

		imgui.SameLine()

		if imgui.ImageButton("##undo_tool", undoButtonImg.Texture().ID(), sz.Mul(1.7)) {
			undoCanvas()
		}

		imgui.SameLine()

		if imgui.ImageButton("##clear_tool", clearButtonImg.Texture().ID(), sz.Mul(1.7)) {
			_ = clearCanvas()
		}

		imgui.SameLine()
		imgui.ImageButton("##open_tool", openButtonImg.Texture().ID(), sz.Mul(1.7))
		imgui.SameLine()
		imgui.ImageButton("##save_tool", saveButtonImg.Texture().ID(), sz.Mul(1.7))

		if imgui.ImageButton("##brush_tool", brushButtonImg.Texture().ID(), sz.Mul(0.9)) {
			brushSize = 12.0
		}

		imgui.SameLine()
		imgui.PushItemWidth(225.0)
		imgui.SliderFloat("##Brush Size", &brushSize, float32(0.1), float32(72.0))
		imgui.PopItemWidth()
		imgui.SetCursorPos(startUl)

		for i := range defaultColors {
			if i%2 != 0 {
				col := g.ToVec4Color(defaultColors[i])
				if imgui.ColorButtonV(fmt.Sprintf("%d##cur_color%d", i, i), col, 0, imgui.Vec2{X: 0, Y: 0}) {
					currentColor = defaultColors[i]
				}

				imgui.SameLineV(0, 0)
			}
		}
	},
	),
	)
}
