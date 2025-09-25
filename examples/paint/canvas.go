package main

import (
	"fmt"
	"image"
	"image/color"
	"sync"

	"github.com/AllenDang/cimgui-go/imgui"

	g "github.com/AllenDang/giu"
)

var (
	canvasDetectedHeight      float32
	canvasComputedWidth       float32
	canvasMarginComputedWidth float32
	canvasInited              = false
	canvas                    *Canvas
	buffer                    = []DrawCommand{}
	currentColor              = color.RGBA{0, 0, 0, 255}
	currentTool               = 0
	brushSize                 = float32(12.0)
	wasDrawing                = false
	lastTo                    image.Point
)

func flushDrawCommands(c *Canvas) {
	var bcopy []DrawCommand

	bcopy = append(bcopy, buffer...)

	go c.AppendDrawCommands(&bcopy)

	buffer = nil
	buffer = []DrawCommand{}
}

// Canvas represents a drawable surface where draw commands are executed.
// It holds the image data, the backend texture, and manages the state of drawing operations.
type Canvas struct {
	// DrawCommands is a slice of drawCommand that records all drawing actions performed on the canvas.
	DrawCommands []DrawCommand

	// Image is the RGBA image that represents the current state of the canvas.
	Image *image.RGBA

	// Backend is the texture that interfaces with the graphical backend to display the canvas.
	Backend *g.ReflectiveBoundTexture

	// LastPaintedIndex is the index of the last draw command that was painted on the canvas.
	LastPaintedIndex int

	// LastComputedLen is the length of the draw commands that have been processed.
	LastComputedLen int

	// UndoIndexes is a slice of integers that keeps track of the indexes of draw commands for undo operations.
	UndoIndexes []int

	// inited is a boolean flag indicating whether the canvas has been initialized.
	inited bool

	appendMu sync.Mutex
}

// getDrawCommands returns a slice of drawCommand starting from the specified index.
// It allows retrieval of draw commands that have been added since a given point in time.
func (c *Canvas) getDrawCommands(sinceIndex int) []DrawCommand {
	return c.DrawCommands[sinceIndex:]
}

// PushImageToBackend updates the backend texture with the current state of the canvas image.
// The commit parameter determines whether the changes should be committed immediately.
func (c *Canvas) PushImageToBackend(commit bool) error {
	err := c.Backend.SetSurfaceFromRGBA(c.Image, commit)
	if err != nil {
		return fmt.Errorf("failed to push image to backend: %w", err)
	}

	return nil
}

// AppendDrawCommands adds a slice of drawCommand to the canvas's existing draw commands.
// It appends the provided commands to the DrawCommands slice.
func (c *Canvas) AppendDrawCommands(cmds *[]DrawCommand) {
	c.appendMu.Lock()
	defer c.appendMu.Unlock()

	c.DrawCommands = append(c.DrawCommands, *cmds...)
}

// Compute processes the draw commands on the canvas and updates the image accordingly.
// It initializes the canvas if it hasn't been initialized yet, and then processes any
// new draw commands that have been added since the last computation.
func (c *Canvas) Compute() {
	// Initialize the canvas if it hasn't been initialized
	if !c.inited {
		// Perform initial flood fill operations to set up the canvas
		Floodfill(c.Image, color.RGBA{255, 255, 254, 255}, 1, 1)
		Floodfill(c.Image, color.RGBA{255, 255, 255, 255}, 2, 2)

		// Push the initial image state to the backend
		err := c.PushImageToBackend(false)
		if err != nil {
			return
		}

		// Mark the canvas as initialized
		c.inited = true

		return
	}

	var draws []DrawCommand

	if func() bool {
		c.appendMu.Lock()
		defer c.appendMu.Unlock()

		// Return if there are no draw commands to process
		if len(c.DrawCommands) < 1 {
			return true
		}

		// Return if all draw commands have already been processed
		if len(c.DrawCommands) <= c.LastComputedLen {
			return true
		}

		// Get the new draw commands that need to be processed
		draws = c.getDrawCommands(c.LastComputedLen)
		return false
	}() {
		return
	}

	for _, r := range draws {
		switch r.Tool {
		case 0:
			// Process line drawing commands
			line := r.ToLine()
			DrawLine(line.P1.X, line.P1.Y, line.P2.X, line.P2.Y, line.C, line.Thickness, c.Image)
		case 1:
			// Process fill commands
			f := r.ToFill()
			Floodfill(c.Image, f.C, f.P1.X, f.P1.Y)
		}
	}

	// Update the backend with the new image state
	_ = c.PushImageToBackend(false)

	// Update the last computed length to the current number of draw commands
	c.LastComputedLen = len(c.DrawCommands)
}

func undoCanvas() {
	if len(canvas.UndoIndexes) > 0 {
		lastUndoIndex := canvas.UndoIndexes[len(canvas.UndoIndexes)-1]
		uind := canvas.UndoIndexes[:len(canvas.UndoIndexes)-1]
		dc := canvas.DrawCommands[:lastUndoIndex]
		canvas.Backend.ForceRelease()
		canvas, _ = NewCanvas(canvasDetectedHeight)
		canvas.UndoIndexes = uind
		canvas.DrawCommands = dc
		canvas.Compute()
	}
}

func clearCanvas() error {
	var err error

	canvas.Backend.ForceRelease()
	canvas, err = NewCanvas(canvasDetectedHeight)

	return err
}

// NewCanvas creates a new Canvas with a specified height.
// It initializes the canvas with a default surface and binds it to a ReflectiveBoundTexture backend.
// Returns a pointer to the Canvas and an error if the surface cannot be set.
func NewCanvas(height float32) (*Canvas, error) {
	backend := &g.ReflectiveBoundTexture{}
	img := defaultSurface(height)

	err := backend.SetSurfaceFromRGBA(img, false)
	if err != nil {
		return nil, fmt.Errorf("failed to set surface from RGBA: %w", err)
	}

	c := &Canvas{Image: img, Backend: backend}

	return c, nil
}

func fittingCanvasSize16By9(height float32) image.Point {
	width := height * (16.0 / 9.0)
	return image.Point{X: int(width), Y: int(height)}
}

func defaultSurface(height float32) *image.RGBA {
	p := fittingCanvasSize16By9(height)
	surface, _ := g.NewUniformLoader(p.X, p.Y, color.RGBA{255, 255, 255, 255}).ServeRGBA()

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
	canvasSize := fittingCanvasSize16By9(canvasDetectedHeight)
	canvasComputedWidth = float32(canvasSize.X)
	canvasMarginComputedWidth = (avail.X - canvasComputedWidth) / 2.0
}

// CanvasWidget creates a widget for the canvas, handling drawing operations and user interactions.
// It manages mouse events to draw on the canvas and updates the drawing buffer accordingly.
func CanvasWidget() g.Widget {
	canvas.Compute()

	return g.Custom(func() {
		// Check if the user has stopped drawing
		if wasDrawing && !g.IsMouseDown(g.MouseButtonLeft) {
			wasDrawing = false

			flushDrawCommands(canvas)

			lastTo = image.Point{0, 0}
		}

		// Get the current screen position of the cursor
		scr := g.GetCursorScreenPos()

		// Render the canvas image
		canvas.Backend.ToImageWidget().Build()

		// Check if the canvas is hovered by the mouse
		if g.IsItemHovered() {
			mousepos := g.GetMousePos()
			if mousepos.X >= scr.X && mousepos.X <= scr.X+int(canvasComputedWidth) && mousepos.Y >= scr.Y && mousepos.Y <= scr.Y+int(canvasDetectedHeight) {
				inpos := image.Point{mousepos.X - scr.X, mousepos.Y - scr.Y}

				// Start drawing on mouse click
				if imgui.IsMouseClickedBool(imgui.MouseButtonLeft) {
					wasDrawing = true

					canvas.UndoIndexes = append(canvas.UndoIndexes, len(canvas.DrawCommands))
					lastTo = image.Point{0, 0}

					buffer = append(buffer, DrawCommand{Tool: currentTool, Color: currentColor, BrushSize: brushSize, From: inpos, To: inpos})
					lastTo = inpos

					flushDrawCommands(canvas)
				}

				// Continue drawing while the mouse is held down
				if g.IsMouseDown(g.MouseButtonLeft) && wasDrawing {
					delta := imgui.CurrentIO().MouseDelta()
					dx := int(delta.X)
					dy := int(delta.Y)

					if dx == 0 || dy == 0 {
						flushDrawCommands(canvas)
					}

					buffer = append(buffer, DrawCommand{Tool: currentTool, Color: currentColor, BrushSize: brushSize, From: lastTo, To: inpos})
					lastTo = inpos

					if len(buffer) >= 8 {
						flushDrawCommands(canvas)
					}
				}
			}
		}
	})
}

// CanvasRow creates a row layout for the canvas widget, initializing the canvas if necessary.
// It ensures the canvas is properly sized and positioned within the GUI.
func CanvasRow() g.Widget {
	return g.Custom(func() {
		// Initialize the canvas if it hasn't been initialized yet
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

		// Build the row layout with the canvas widget
		g.Row(
			g.Dummy(canvasMarginComputedWidth, canvasDetectedHeight),
			CanvasWidget(),
		).Build()
	})
}
