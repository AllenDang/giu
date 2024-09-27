package main

import (
	"image"
	"image/color"

	"github.com/AllenDang/cimgui-go/imgui"
	g "github.com/AllenDang/giu"
)

var (
	canvasDetectedHeight      float32
	canvasComputedWidth       float32
	canvasMarginComputedWidth float32
	canvasInited              = false
	canvas                    *Canvas
	//buffer                    *image.RGBA
	buffer        = []DrawCommand{}
	current_color = color.RGBA{0, 0, 0, 255}
	current_tool  = 0
	brush_size    = float32(12.0)
	was_drawing   = false
	lastTo        image.Point
)

func FlushDrawCommands(c *Canvas) {
	bcopy := append(buffer)
	go c.AppendDrawCommands(&bcopy)
	buffer = nil
	buffer = []DrawCommand{}
}

type Canvas struct {
	DrawCommands     []DrawCommand
	Image            *image.RGBA
	Backend          *g.ReflectiveBoundTexture
	LastPaintedIndex int
	LastComputedLen  int
	UndoIndexes      []int
	inited           bool
}

func (c *Canvas) GetDrawCommands(since_index int) []DrawCommand {
	return c.DrawCommands[since_index:]
}

func (c *Canvas) PushImageToBackend(commit bool) error {
	return c.Backend.SetSurfaceFromRGBA(c.Image, commit)
}

func (c *Canvas) AppendDrawCommands(cmds *[]DrawCommand) {
	/*lockid, err := g.mu.Lock()
	if err != nil {
		panic(err)
	}
	defer g.mu.Unlock(lockid)*/
	c.DrawCommands = append(c.DrawCommands, *cmds...)
}

func (c *Canvas) Compute() {
	if !c.inited {
		Floodfill(c.Image, color.RGBA{255, 255, 254, 255}, 1, 1)
		Floodfill(c.Image, color.RGBA{255, 255, 255, 255}, 2, 2)
		err := c.PushImageToBackend(false)
		if err != nil {
			return
		}
		c.inited = true
		return
	}
	if len(c.DrawCommands) < 1 {
		return
	}
	if len(c.DrawCommands) <= c.LastComputedLen {
		return
	}
	draws := c.GetDrawCommands(c.LastComputedLen)
	for _, r := range draws {
		switch r.Tool {
		case 0:
			line := r.ToLine()
			DrawLine(line.P1.X, line.P1.Y, line.P2.X, line.P2.Y, line.C, line.Thickness, c.Image)
		case 1:
			f := r.ToFill()
			Floodfill(c.Image, f.C, f.P1.X, f.P1.Y)
		default:
		}
	}
	_ = c.PushImageToBackend(false)
	c.LastComputedLen = len(c.DrawCommands)
}

func NewCanvas(height float32) (*Canvas, error) {
	backend := &g.ReflectiveBoundTexture{}
	image := defaultSurface(height)
	err := backend.SetSurfaceFromRGBA(image, false)
	if err != nil {
		return nil, err
	}
	c := &Canvas{Image: image, Backend: backend}
	return c, nil
}

func FittingCanvasSize16By9(height float32) image.Point {
	width := height * (16.0 / 9.0)
	return image.Point{X: int(width), Y: int(height)}
}

func defaultSurface(height float32) *image.RGBA {
	p := FittingCanvasSize16By9(height)
	surface, _ := g.UniformLoader(p.X, p.Y, color.RGBA{255, 255, 255, 255}).ServeRGBA()
	return surface
}

var defaultColors = []color.RGBA{
	//	UPLINE
	{0, 0, 0, 255},
	{127, 127, 127, 255},
	{136, 0, 21, 255},
	{237, 28, 36, 255},
	{255, 127, 39, 255},
	{255, 242, 0, 255},
	{34, 177, 76, 255},
	{0, 162, 232, 255},
	{63, 72, 204, 255},
	{163, 73, 164, 255},
	// DOWNLINE
	{255, 255, 255, 255},
	{195, 195, 195, 255},
	{185, 122, 87, 255},
	{255, 174, 201, 255},
	{255, 201, 14, 255},
	{239, 228, 176, 255},
	{181, 230, 29, 255},
	{153, 217, 234, 255},
	{112, 146, 190, 255},
	{200, 191, 231, 255},
}

func computeCanvasBounds() {
	avail := imgui.ContentRegionAvail()
	canvasDetectedHeight = avail.Y
	canvasSize := FittingCanvasSize16By9(canvasDetectedHeight)
	canvasComputedWidth = float32(canvasSize.X)
	canvasMarginComputedWidth = (avail.X - canvasComputedWidth) / 2.0
}

func CanvasWidget() g.Widget {
	canvas.Compute()
	return g.Custom(func() {
		if was_drawing && !g.IsMouseDown(g.MouseButtonLeft) {
			was_drawing = false
			FlushDrawCommands(canvas)
			lastTo = image.Point{0, 0}
		}
		scr := g.GetCursorScreenPos()
		canvas.Backend.ToImageWidget().Build()
		if g.IsItemHovered() {
			mousepos := g.GetMousePos()
			if mousepos.X >= scr.X && mousepos.X <= scr.X+int(canvasComputedWidth) && mousepos.Y >= scr.Y && mousepos.Y <= scr.Y+int(canvasDetectedHeight) {
				if g.IsWindowFocused(0) {
					inpos := image.Point{mousepos.X - scr.X, mousepos.Y - scr.Y}
					if imgui.IsMouseClickedBool(imgui.MouseButtonLeft) {
						was_drawing = true
						lastTo = image.Point{0, 0}
						buffer = append(buffer, DrawCommand{Tool: current_tool, Color: current_color, BrushSize: brush_size, From: inpos, To: inpos})
						lastTo = inpos
						FlushDrawCommands(canvas)
					}
					if g.IsMouseDown(g.MouseButtonLeft) && was_drawing {
						delta := imgui.CurrentIO().MouseDelta()
						dx := int(delta.X)
						dy := int(delta.Y)
						if dx == 0 || dy == 0 {
							FlushDrawCommands(canvas)
						}
						buffer = append(buffer, DrawCommand{Tool: current_tool, Color: current_color, BrushSize: brush_size, From: lastTo, To: inpos})
						lastTo = inpos
						if len(buffer) >= 8 {
							FlushDrawCommands(canvas)
						}
					}
				}
			}
		}
	})
}

func CanvasRow() g.Widget {
	return g.Custom(func() {
		if !canvasInited {
			computeCanvasBounds()
			var err error
			canvas, err = NewCanvas(canvasDetectedHeight)
			if err != nil {
				return
			}
			canvasInited = true
			return
		}
		g.Row(
			g.Dummy(canvasMarginComputedWidth, canvasDetectedHeight),
			CanvasWidget(),
		).Build()
	})
}
