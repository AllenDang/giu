package giu

import (
	"image/color"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imguizmo"
	"github.com/AllenDang/cimgui-go/utils"
	glm "github.com/gucio321/glm-go"
)

// GizmoOperation specifies the operation of Gizmo (used by manipulate).
type GizmoOperation int

// Possible Operations.
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

// values are not explained in source code.
const (
	ModeLocal GizmoMode = GizmoMode(imguizmo.LOCAL)
	ModeWorld GizmoMode = GizmoMode(imguizmo.WORLD)
)

// GizmoI should be implemented by every sub-element of GizmoWidget.
type GizmoI interface {
	Gizmo(view *ViewMatrix, projection *ProjectionMatrix)
}

var _ Widget = &GizmoWidget{}

// GizmoWidget implements ImGuizmo features.
// It is designed just like PlotWidget.
// This structure provides an "area" where you can put Gizmos (see (*GizmoWidget).Gizmos).
// If you want to have more understanding about what is going on here, read this:
// https://www.opengl-tutorial.org/beginners-tutorials/tutorial-3-matrices/ (DISCLAIMER: giu authors are not responsible if you go mad or something!)
type GizmoWidget struct {
	gizmos       []GizmoI
	view         *ViewMatrix
	projection   *ProjectionMatrix
	id           ID
	disabled     bool
	orthographic bool
}

// Gizmo creates a new GizmoWidget.
func Gizmo(view *ViewMatrix, projection *ProjectionMatrix) *GizmoWidget {
	return &GizmoWidget{
		gizmos:       []GizmoI{},
		view:         view,
		projection:   projection,
		id:           GenAutoID("gizmo"),
		disabled:     false,
		orthographic: false,
	}
}

// ID sets the ID of the GizmoWidget. (useful if you use multiple gizmo widgets. It is set by AutoID anyway).
func (g *GizmoWidget) ID(id ID) *GizmoWidget {
	g.id = id
	return g
}

// Disabled sets GizmoWidget's disabled state.
func (g *GizmoWidget) Disabled(b bool) *GizmoWidget {
	g.disabled = b
	return g
}

// Orthographic sets the projection matrix to orthographic.
func (g *GizmoWidget) Orthographic(b bool) *GizmoWidget {
	g.orthographic = b
	return g
}

// Gizmos adds GizmoI elements to the GizmoWidget area.
func (g *GizmoWidget) Gizmos(gizmos ...GizmoI) *GizmoWidget {
	g.gizmos = append(g.gizmos, gizmos...)
	return g
}

// build is a local build function.
// Just to separate Global() and Build() methods.
func (g *GizmoWidget) build() {
	imguizmo.PushIDStr(string(g.id))
	imguizmo.Enable(!g.disabled)
	imguizmo.SetOrthographic(g.orthographic)

	for _, gizmo := range g.gizmos {
		gizmo.Gizmo(g.view, g.projection)
	}

	imguizmo.PopID()
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
	matrix    *ViewMatrix
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

// Matrix allows to set grid matrix. Default to IdentityMatrix.
func (g *GridGizmo) Matrix(matrix *ViewMatrix) *GridGizmo {
	g.matrix = matrix
	return g
}

// Gizmo implements GizmoI interface.
func (g *GridGizmo) Gizmo(view *ViewMatrix, projection *ProjectionMatrix) {
	imguizmo.DrawGrid(
		view.getMatrix(),
		projection.getMatrix(),
		g.matrix.getMatrix(),
		g.thickness,
	)
}

var _ GizmoI = &CubeGizmo{}

// CubeGizmo draws a 3D cube in the gizmo area.
// View and Projection matrices are provided by GizmoWidget.
type CubeGizmo struct {
	matrix     *ViewMatrix
	manipulate bool
}

// Cube creates a new CubeGizmo.
func Cube(matrix *ViewMatrix) *CubeGizmo {
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
func (c *CubeGizmo) Gizmo(view *ViewMatrix, projection *ProjectionMatrix) {
	imguizmo.DrawCubes(
		view.getMatrix(),
		projection.getMatrix(),
		c.matrix.getMatrix(),
		1,
	)

	if c.manipulate {
		Manipulate(c.matrix).Gizmo(view, projection)
	}
}

var _ GizmoI = &ManipulateGizmo{}

// ManipulateGizmo is a gizmo that allows you to "visually manipulate a matrix".
// It can be attached to another Gizmo (e.g. CubeGizmo) and will allow to move/rotate/scale it.
// See (*CubeGizmo).Manipulate() method.
type ManipulateGizmo struct {
	matrix    *ViewMatrix
	operation GizmoOperation
	mode      GizmoMode
}

// Manipulate creates a new ManipulateGizmo.
func Manipulate(matrix *ViewMatrix) *ManipulateGizmo {
	return &ManipulateGizmo{
		matrix:    matrix,
		mode:      GizmoMode(imguizmo.LOCAL),
		operation: GizmoOperation(imguizmo.TRANSLATE),
	}
}

// Gizmo implements GizmoI interface.
func (m *ManipulateGizmo) Gizmo(view *ViewMatrix, projection *ProjectionMatrix) {
	imguizmo.ManipulateV(
		view.getMatrix(),
		projection.getMatrix(),
		imguizmo.OPERATION(m.operation),
		imguizmo.MODE(m.mode),
		m.matrix.getMatrix(),
		nil, // this is deltaMatrix (Can't see usecase for now)
		nil, // snap idk what is this
		nil, // localBounds idk what is this
		nil) // boundsSnap idk what is this
}

var _ GizmoI = &ViewManipulateGizmo{}

// ViewManipulateGizmo is a gizmo that allows you to manipulate world rotation visualy.
type ViewManipulateGizmo struct {
	position imgui.Vec2
	size     imgui.Vec2
	color    uint32
}

// ViewManipulate creates a new ViewManipulateGizmo.
func ViewManipulate() *ViewManipulateGizmo {
	return &ViewManipulateGizmo{
		position: imgui.Vec2{X: 128, Y: 128},
		size:     imgui.Vec2{X: 128, Y: 128},
	}
}

// Position sets the position of the gizmo.
func (v *ViewManipulateGizmo) Position(x, y float32) *ViewManipulateGizmo {
	v.position = imgui.Vec2{X: x, Y: y}

	return v
}

// Size sets the size of the gizmo.
func (v *ViewManipulateGizmo) Size(x, y float32) *ViewManipulateGizmo {
	v.size = imgui.Vec2{X: x, Y: y}

	return v
}

// Color sets the color of the gizmo.
func (v *ViewManipulateGizmo) Color(c color.Color) *ViewManipulateGizmo {
	v.color = ColorToUint(c)

	return v
}

// Gizmo implements GizmoI interface.
func (v *ViewManipulateGizmo) Gizmo(view *ViewMatrix, _ *ProjectionMatrix) {
	imguizmo.ViewManipulateFloat(
		view.getMatrix(),
		1,
		v.position,
		v.size,
		v.color,
	)
}

// [Gizmo helpers]

// ViewMatrix allows to generate a "view" matrix:
// - position
// - rotation
// - scale
// NOTE: You are supposed to allocate this with NewViewMatrix (do not use zero value)!!!
type ViewMatrix struct {
	transform []float32 // supposed len is 3
	rotation  []float32 // supposed len is 3
	scale     []float32 // supposed len is 3
	matrix    []float32 // supposed len is 16
	dirty     bool
}

// NewViewMatrix creates a new ViewMatrix.
func NewViewMatrix() *ViewMatrix {
	return &ViewMatrix{
		transform: make([]float32, 3),
		rotation:  make([]float32, 3),
		scale:     make([]float32, 3),
		matrix:    make([]float32, 16),
		dirty:     true,
	}
}

// IdentityMatrix creates a new ViewMatrix with identity matrix.
func IdentityMatrix() *ViewMatrix {
	r := NewViewMatrix()
	r.matrix = []float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	r.decompile()
	r.dirty = false

	return r
}

// Transform sets the position of the matrix.
func (m *ViewMatrix) Transform(x, y, z float32) *ViewMatrix {
	m.transform[0] = x
	m.transform[1] = y
	m.transform[2] = z
	m.dirty = true

	return m
}

// Rotation sets the rotation of the matrix.
func (m *ViewMatrix) Rotation(x, y, z float32) *ViewMatrix {
	m.rotation[0] = x
	m.rotation[1] = y
	m.rotation[2] = z
	m.dirty = true

	return m
}

// Scale sets the scale of the matrix.
func (m *ViewMatrix) Scale(x, y, z float32) *ViewMatrix {
	m.scale[0] = x
	m.scale[1] = y
	m.scale[2] = z
	m.dirty = true

	return m
}

// SetMatrix allows you to set the matrix directly.
// NOTE: f is supposed to be 16 elements long.
// NOTE: it is not recommended - use components functions.
func (m *ViewMatrix) SetMatrix(f []float32) *ViewMatrix {
	m.matrix = f
	m.decompile()
	m.dirty = false

	return m
}

// Copy copies returns a copy of the matrix.
// Useful e.g. in exaples/ to duplicate the matrix.
func (m *ViewMatrix) Copy() *ViewMatrix {
	return NewViewMatrix().
		Transform(m.transform[0], m.transform[1], m.transform[2]).
		Rotation(m.rotation[0], m.rotation[1], m.rotation[2]).
		Scale(m.scale[0], m.scale[1], m.scale[2])
}

// Compile updates m.matrix
// NOTE: this supposes matrix was allocated correctly!
func (m *ViewMatrix) compile() {
	imguizmo.RecomposeMatrixFromComponents(
		utils.SliceToPtr(m.transform),
		utils.SliceToPtr(m.rotation),
		utils.SliceToPtr(m.scale),
		utils.SliceToPtr(m.matrix),
	)

	m.dirty = false
}

// decompile updates m.transform, m.rotation, m.scale from m.matrix.
func (m *ViewMatrix) decompile() {
	imguizmo.DecomposeMatrixToComponents(
		utils.SliceToPtr(m.matrix),
		utils.SliceToPtr(m.transform),
		utils.SliceToPtr(m.rotation),
		utils.SliceToPtr(m.scale),
	)
}

// matrix returns current matrix compatible with ImGuizmo (pointer to 4x4m).
// It recompiles as necessary.
func (m *ViewMatrix) getMatrix() *float32 {
	if m.dirty {
		m.compile()
	}

	return utils.SliceToPtr(m.matrix)
}

// MatrixSlice returns ViewMatrix as a slice (for debugging purposes).
func (m *ViewMatrix) MatrixSlice() []float32 {
	if m.dirty {
		m.compile()
	}

	return m.matrix
}

// ProjectionMatrix represents a matrix for Gizmo projection.
// ref: https://www.opengl-tutorial.org/beginners-tutorials/tutorial-3-matrices/#the-projection-matrix
type ProjectionMatrix struct {
	// The vertical Field of View, in radians: the amount of "zoom". Think "camera lens". Usually between 90° (extra wide) and 30° (quite zoomed in)
	// This value is in radians! See Deg2Rad.
	fov float32
	// Aspect Ratio. Depends on the size of your window. Notice that 4/3 == 800/600 == 1280/960, sounds familiar?
	aspect float32
	// Near clipping plane. Keep as big as possible, or you'll get precision issues.
	nearClipping float32
	// Far clipping plane. Keep as little as possible.
	farClipping float32

	dirty  bool
	matrix []float32
}

// NewProjectionMatrix creates a new ProjectionMatrix.
func NewProjectionMatrix() *ProjectionMatrix {
	return &ProjectionMatrix{
		fov:          Deg2Rad(45),
		aspect:       3.0 / 4.0,
		nearClipping: 0.1,
		farClipping:  100.0,
		dirty:        true,
		matrix:       make([]float32, 16),
	}
}

// FOV sets the Field of View.
func (p *ProjectionMatrix) FOV(fov float32) *ProjectionMatrix {
	p.fov = fov
	p.dirty = true

	return p
}

// Aspect sets the Aspect Ratio.
func (p *ProjectionMatrix) Aspect(aspect float32) *ProjectionMatrix {
	p.aspect = aspect
	p.dirty = true

	return p
}

// NearClipping sets the Near Clipping plane.
func (p *ProjectionMatrix) NearClipping(near float32) *ProjectionMatrix {
	p.nearClipping = near
	p.dirty = true

	return p
}

// FarClipping sets the Far Clipping plane.
func (p *ProjectionMatrix) FarClipping(far float32) *ProjectionMatrix {
	p.farClipping = far
	p.dirty = true

	return p
}

// Copy returns a copy of the matrix.
func (p *ProjectionMatrix) Copy() *ProjectionMatrix {
	return NewProjectionMatrix().
		FOV(p.fov).
		Aspect(p.aspect).
		NearClipping(p.nearClipping).
		FarClipping(p.farClipping)
}

// getMatrix returns the matrix compatible with ImGuizmo (pointer to 4x4m).
func (p *ProjectionMatrix) getMatrix() *float32 {
	if p.dirty {
		p.compile()
	}

	return utils.SliceToPtr(p.matrix)
}

// compile updates p.matrix.
func (p *ProjectionMatrix) compile() {
	p.matrix = glm.MatrixPerspective(p.fov, p.aspect, p.nearClipping, p.farClipping)
	p.dirty = false
}
