package giu

import (
	"image/color"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imguizmo"
	"github.com/AllenDang/cimgui-go/utils"
)

// GizmoOperation specifies the operation of Gizmo (used by manipulate).
type GizmoOperation int

// Possible Operations
const (
	OperationTranslateX GizmoOperation = GizmoOperation(imguizmo.TRANSLATEX)
	OperationTranslateY GizmoOperation = GizmoOperation(imguizmo.TRANSLATEY)
	OperationTranslateZ GizmoOperation = GizmoOperation(imguizmo.TRANSLATEZ)
	OperationTranslate  GizmoOperation = GizmoOperation(imguizmo.TRANSLATE)

	OperationRotateX GizmoOperation = GizmoOperation(imguizmo.ROTATEX)
	OperationRotateY GizmoOperation = GizmoOperation(imguizmo.ROTATEY)
	OperationRotateZ GizmoOperation = GizmoOperation(imguizmo.ROTATEZ)
	OperationRotate  GizmoOperation = GizmoOperation(imguizmo.ROTATE)

	OperationScaleX GizmoOperation = GizmoOperation(imguizmo.SCALEX)
	OperationScaleY GizmoOperation = GizmoOperation(imguizmo.SCALEY)
	OperationScaleZ GizmoOperation = GizmoOperation(imguizmo.SCALEZ)
	OperationScale  GizmoOperation = GizmoOperation(imguizmo.SCALE)

	OperationScaleXU GizmoOperation = GizmoOperation(imguizmo.SCALEXU)
	OperationScaleYU GizmoOperation = GizmoOperation(imguizmo.SCALEYU)
	OperationScaleZU GizmoOperation = GizmoOperation(imguizmo.SCALEZU)
	OperationScaleU  GizmoOperation = GizmoOperation(imguizmo.SCALEU)

	OperationBounds GizmoOperation = GizmoOperation(imguizmo.BOUNDS)

	OperationRotateScreen GizmoOperation = GizmoOperation(imguizmo.ROTATESCREEN)

	OperationUniversal GizmoOperation = GizmoOperation(imguizmo.UNIVERSAL)
)

// GizmoMode specifies the mode of Gizmo (used by manipulate).
type GizmoMode int

// values are not explained in source code
const (
	ModeLocal GizmoMode = GizmoMode(imguizmo.LOCAL)
	ModeWorld GizmoMode = GizmoMode(imguizmo.WORLD)
)

// GizmoI should be implemented by every sub-element of GizmoWidget.
type GizmoI interface {
	Gizmo(view, projection *HumanReadableMatrix)
}

var _ Widget = &GizmoWidget{}

// GizmoWidget implements ImGuizmo features.
// It is designed just like PlotWidget.
// This structure provides an "area" where you can put Gizmos (see (*GizmoWidget).Gizmos).
// If you wnat to have more understanding about what is going on here, read this:
// https://www.opengl-tutorial.org/beginners-tutorials/tutorial-3-matrices/ (DISCLAIMER: giu authors are not responsible if you go mad or something!)
// TODO: ProjectionMatrix edition (see https://github.com/jbowtie/glm-go)
type GizmoWidget struct {
	gizmos           []GizmoI
	view, projection *HumanReadableMatrix
}

// Gizmo creates a new GizmoWidget.
func Gizmo(view, projection *HumanReadableMatrix) *GizmoWidget {
	return &GizmoWidget{
		gizmos:     []GizmoI{},
		view:       view,
		projection: projection,
	}
}

// Gizmos adds GizmoI elements to the GizmoWidget area.
func (g *GizmoWidget) Gizmos(gizmos ...GizmoI) *GizmoWidget {
	g.gizmos = append(g.gizmos, gizmos...)
	return g
}

// build is a local build function.
// Just to separate Global() and Build() methods.
func (g *GizmoWidget) build() {
	for _, gizmo := range g.gizmos {
		gizmo.Gizmo(g.view, g.projection)
	}
}

// Build implements Widget interface.
func (g *GizmoWidget) Build() {
	imguizmo.SetDrawlist()
	displaySize := imgui.ContentRegionAvail()
	pos0 := imgui.CursorScreenPos()
	imguizmo.SetRect(pos0.X, pos0.Y, displaySize.X, displaySize.Y)
	g.build()
}

// Global works like Build() but does not attach the gizmo to the current window.
func (g *GizmoWidget) Global() {
	displaySize := imgui.CurrentIO().DisplaySize()
	pos0 := imgui.MainViewport().Pos()
	imguizmo.SetRect(pos0.X, pos0.Y, displaySize.X, displaySize.Y)
	g.build()
}

// [Gizmos]

var _ GizmoI = &GridGizmo{}

// GridGizmo draws a grid in the gizmo area.
type GridGizmo struct {
	// default to Identity
	matrix    *HumanReadableMatrix
	thickness float32
}

// Grid creates a new GridGizmo.
func Grid() *GridGizmo {
	return &GridGizmo{
		matrix:    IdentityMatrix(),
		thickness: 10,
	}
}

// Thickness sets a thickness of grid lines.
func (g *GridGizmo) Thickness(t float32) *GridGizmo {
	g.thickness = t
	return g
}

// Gizmo implements GizmoI interface.
func (g *GridGizmo) Gizmo(view, projection *HumanReadableMatrix) {
	imguizmo.DrawGrid(
		view.Matrix(),
		projection.Matrix(),
		g.matrix.Matrix(),
		g.thickness,
	)
}

// CubeGizmo draws a 3D cube in the gizmo area.
// View and Projection matrices are provided by GizmoWidget.
type CubeGizmo struct {
	matrix     *HumanReadableMatrix
	manipulate bool
}

// Cube creates a new CubeGizmo.
func Cube(matrix *HumanReadableMatrix) *CubeGizmo {
	return &CubeGizmo{
		matrix: matrix,
	}
}

// Manipulate adds ManipulateGizmo to the CubeGizmo.
func (c *CubeGizmo) Manipulate() *CubeGizmo {
	c.manipulate = true
	return c
}

// Gizmo implements GizmoI interface.
func (c *CubeGizmo) Gizmo(view, projection *HumanReadableMatrix) {
	imguizmo.DrawCubes(
		view.Matrix(),
		projection.Matrix(),
		c.matrix.Matrix(),
		1,
	)

	if c.manipulate {
		Manipulate(c.matrix).Gizmo(view, projection)
	}
}

// ManipulateGizmo is a gizmo that allows you to "visually manipulate a matrix".
// It can be attached to another Gizmo (e.g. CubeGizmo) and will allow to move/rotate/scale it.
// See (*CubeGizmo).Manipulate() method.
type ManipulateGizmo struct {
	matrix    *HumanReadableMatrix
	operation GizmoOperation
	mode      GizmoMode
}

// Manipulate creates a new ManipulateGizmo.
func Manipulate(matrix *HumanReadableMatrix) *ManipulateGizmo {
	return &ManipulateGizmo{
		matrix:    matrix,
		mode:      GizmoMode(imguizmo.LOCAL),
		operation: GizmoOperation(imguizmo.TRANSLATE),
	}
}

// Gizmo implements GizmoI interface.
func (m *ManipulateGizmo) Gizmo(view, projection *HumanReadableMatrix) {
	imguizmo.ManipulateV(
		view.Matrix(),
		projection.Matrix(),
		imguizmo.OPERATION(m.operation),
		imguizmo.MODE(m.mode),
		m.matrix.Matrix(),
		nil, // this is deltaMatrix (Can't see usecase for now)
		nil, // snap idk what is this
		nil, // localBounds idk what is this
		nil) // boundsSnap idk what is this
}

type ViewManipulateGizmo struct {
	position imgui.Vec2
	size     imgui.Vec2
	color    uint32
}

func ViewManipulate() *ViewManipulateGizmo {
	return &ViewManipulateGizmo{
		position: imgui.Vec2{128, 128},
		size:     imgui.Vec2{128, 128},
	}
}

// Position sets the position of the gizmo.
func (v *ViewManipulateGizmo) Position(x, y float32) *ViewManipulateGizmo {
	v.position = imgui.Vec2{x, y}
	return v
}

// Size sets the size of the gizmo.
func (v *ViewManipulateGizmo) Size(x, y float32) *ViewManipulateGizmo {
	v.size = imgui.Vec2{x, y}
	return v
}

func (v *ViewManipulateGizmo) Color(c color.Color) *ViewManipulateGizmo {
	v.color = ColorToUint(c)
	return v
}

// Gizmo implements GizmoI interface.
func (v *ViewManipulateGizmo) Gizmo(view, projection *HumanReadableMatrix) {
	imguizmo.ViewManipulateFloat(
		view.Matrix(),
		1,
		v.position,
		v.size,
		v.color,
	)
}

// [Gizmo helpers]

// HumanReadableMatrix is a suitable thing here.
// It makes it even possible to use gizmos.
// Recommended use is presented in examples/gizmo
// NOTE: You are supposed to allocate this with NewHumanReadableMatrix (do not use zero value)!!!
type HumanReadableMatrix struct {
	transform []float32 // supposed len is 3
	rotation  []float32 // supposed len is 3
	scale     []float32 // supposed len is 3
	matrix    []float32 // supposed len is 16
	dirty     bool
}

func NewHumanReadableMatrix() *HumanReadableMatrix {
	return &HumanReadableMatrix{
		transform: make([]float32, 3),
		rotation:  make([]float32, 3),
		scale:     make([]float32, 3),
		matrix:    make([]float32, 16),
		dirty:     true,
	}
}

func IdentityMatrix() *HumanReadableMatrix {
	r := NewHumanReadableMatrix()
	r.matrix = []float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	r.Decompile()
	r.dirty = false

	return r
}

func (m *HumanReadableMatrix) Transform(x, y, z float32) *HumanReadableMatrix {
	m.transform[0] = x
	m.transform[1] = y
	m.transform[2] = z
	m.dirty = true
	return m
}

func (m *HumanReadableMatrix) Rotation(x, y, z float32) *HumanReadableMatrix {
	m.rotation[0] = x
	m.rotation[1] = y
	m.rotation[2] = z
	m.dirty = true
	return m
}

func (m *HumanReadableMatrix) Scale(x, y, z float32) *HumanReadableMatrix {
	m.scale[0] = x
	m.scale[1] = y
	m.scale[2] = z
	m.dirty = true
	return m
}

func (m *HumanReadableMatrix) SetMatrix(f []float32) *HumanReadableMatrix {
	m.matrix = f
	m.Decompile()
	m.dirty = false
	return m
}

// Compile updates m.Matrix
// NOTE: this supposes matrix was allocated correctly!
func (m *HumanReadableMatrix) Compile() {
	imguizmo.RecomposeMatrixFromComponents(
		utils.SliceToPtr(m.transform),
		utils.SliceToPtr(m.rotation),
		utils.SliceToPtr(m.scale),
		utils.SliceToPtr(m.matrix),
	)

	m.dirty = false
}

func (m *HumanReadableMatrix) Decompile() {
	imguizmo.DecomposeMatrixToComponents(
		utils.SliceToPtr(m.matrix),
		utils.SliceToPtr(m.transform),
		utils.SliceToPtr(m.rotation),
		utils.SliceToPtr(m.scale),
	)
}

func (m *HumanReadableMatrix) GetTransform() []float32 {
	if m.dirty {
		m.Compile()
	}

	return m.transform
}

func (m *HumanReadableMatrix) GetRotation() []float32 {
	if m.dirty {
		m.Compile()
	}

	return m.rotation
}

func (m *HumanReadableMatrix) GetScale() []float32 {
	if m.dirty {
		m.Compile()
	}

	return m.scale
}

// Matrix returns current matrix compatible with ImGuizmo (pointer to 4x4m).
// It recompiles as necessary.
func (m *HumanReadableMatrix) Matrix() *float32 {
	if m.dirty {
		m.Compile()
	}

	return utils.SliceToPtr(m.matrix)
}

func (m *HumanReadableMatrix) MatrixSlice() []float32 {
	if m.dirty {
		m.Compile()
	}

	return m.matrix
}
