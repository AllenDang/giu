package imgui

// #include "DrawListWrapper.h"
import "C"
import (
	"unsafe"
)

// DrawList is a draw-command list.
// This is the low-level list of polygons that ImGui functions are filling.
// At the end of the frame, all command lists are passed to your render function for rendering.
//
// Each ImGui window contains its own DrawList. You can use GetWindowDrawList() to access
// the current window draw list and draw custom primitives.
//
// You can interleave normal ImGui calls and adding primitives to the current draw list.
//
// All positions are generally in pixel coordinates (top-left at (0,0), bottom-right at io.DisplaySize),
// however you are totally free to apply whatever transformation matrix to want to the data
// (if you apply such transformation you'll want to apply it to ClipRect as well)
//
// Important: Primitives are always added to the list and not culled (culling is done at
// higher-level by ImGui functions), if you use this API a lot consider coarse culling your drawn objects.
type DrawList uintptr

func (list DrawList) handle() C.IggDrawList {
	return C.IggDrawList(list)
}

// Commands returns the list of draw commands.
// Typically 1 command = 1 GPU draw call, unless the command is a callback.
func (list DrawList) Commands() []DrawCommand {
	count := int(C.iggDrawListGetCommandCount(list.handle()))
	commands := make([]DrawCommand, count)
	for i := 0; i < count; i++ {
		commands[i] = DrawCommand(C.iggDrawListGetCommand(list.handle(), C.int(i)))
	}
	return commands
}

// VertexBufferLayout returns the byte sizes necessary to select fields in a vertex buffer of a DrawList.
func VertexBufferLayout() (entrySize int, posOffset int, uvOffset int, colOffset int) {
	var entrySizeArg C.size_t
	var posOffsetArg C.size_t
	var uvOffsetArg C.size_t
	var colOffsetArg C.size_t
	C.iggGetVertexBufferLayout(&entrySizeArg, &posOffsetArg, &uvOffsetArg, &colOffsetArg)
	entrySize = int(entrySizeArg)
	posOffset = int(posOffsetArg)
	uvOffset = int(uvOffsetArg)
	colOffset = int(colOffsetArg)
	return
}

// VertexBuffer returns the handle information of the whole vertex buffer.
// Returned are the handle pointer and the total byte size.
// The buffer is a packed array of vertex entries, each consisting of a 2D position vector, a 2D UV vector,
// and a 4-byte color value. To determine the byte size and offset values, call VertexBufferLayout.
func (list DrawList) VertexBuffer() (unsafe.Pointer, int) {
	var data unsafe.Pointer
	var size C.int

	C.iggDrawListGetRawVertexBuffer(list.handle(), &data, &size)

	return data, int(size)
}

// IndexBufferLayout returns the byte size necessary to select fields in an index buffer of DrawList.
func IndexBufferLayout() (entrySize int) {
	var entrySizeArg C.size_t
	C.iggGetIndexBufferLayout(&entrySizeArg)
	entrySize = int(entrySizeArg)
	return
}

// IndexBuffer returns the handle information of the whole index buffer.
// Returned are the handle pointer and the total byte size.
// The buffer is a packed array of index entries, each consisting of an integer offset.
// To determine the byte size, call IndexBufferLayout.
func (list DrawList) IndexBuffer() (unsafe.Pointer, int) {
	var data unsafe.Pointer
	var size C.int

	C.iggDrawListGetRawIndexBuffer(list.handle(), &data, &size)

	return data, int(size)
}

func (list DrawList) AddLine(p1, p2 Vec2, col Vec4, thickness float32) {
	c := GetColorU32(col)
	p1Arg, _ := p1.wrapped()
	p2Arg, _ := p2.wrapped()
	C.iggDrawListAddLine(list.handle(), p1Arg, p2Arg, C.uint(c), C.float(thickness))
}

func (list DrawList) AddRect(pMin, pMax Vec2, col Vec4, rounding float32, rounding_corners int, thickness float32) {
	c := GetColorU32(col)
	pMinArg, _ := pMin.wrapped()
	pMaxArg, _ := pMax.wrapped()
	C.iggDrawListAddRect(list.handle(), pMinArg, pMaxArg, C.uint(c), C.float(rounding), C.int(rounding_corners), C.float(thickness))
}

func (list DrawList) AddRectFilled(pMin, pMax Vec2, col Vec4, rounding float32, rounding_corners int) {
	c := GetColorU32(col)
	pMinArg, _ := pMin.wrapped()
	pMaxArg, _ := pMax.wrapped()
	C.iggDrawListAddRectFilled(list.handle(), pMinArg, pMaxArg, C.uint(c), C.float(rounding), C.int(rounding_corners))
}

func (list DrawList) AddText(pos Vec2, col Vec4, text string) {
	c := GetColorU32(col)
	posArg, _ := pos.wrapped()
	textArg, textFin := wrapString(text)
	defer textFin()
	C.iggDrawListAddText(list.handle(), posArg, C.uint(c), textArg)
}

func (list DrawList) AddBezierCurve(pos0, cp0, cp1, pos1 Vec2, col Vec4, thickness float32, num_segments int) {
	c := GetColorU32(col)
	pos0Arg, _ := pos0.wrapped()
	cp0Arg, _ := cp0.wrapped()
	cp1Arg, _ := cp1.wrapped()
	pos1Arg, _ := pos1.wrapped()
	C.iggDrawListAddBezierCurve(list.handle(), pos0Arg, cp0Arg, cp1Arg, pos1Arg, C.uint(c), C.float(thickness), C.int(num_segments))
}

func (list DrawList) AddTriangle(p1, p2, p3 Vec2, col Vec4, thickness float32) {
	c := GetColorU32(col)
	p1Arg, _ := p1.wrapped()
	p2Arg, _ := p2.wrapped()
	p3Arg, _ := p3.wrapped()
	C.iggDrawListAddTriangle(list.handle(), p1Arg, p2Arg, p3Arg, C.uint(c), C.float(thickness))
}

func (list DrawList) AddTriangleFilled(p1, p2, p3 Vec2, col Vec4) {
	c := GetColorU32(col)
	p1Arg, _ := p1.wrapped()
	p2Arg, _ := p2.wrapped()
	p3Arg, _ := p3.wrapped()
	C.iggDrawListAddTriangleFilled(list.handle(), p1Arg, p2Arg, p3Arg, C.uint(c))
}

func (list DrawList) AddCircle(center Vec2, radius float32, col Vec4, thickness float32) {
	c := GetColorU32(col)
	centerArg, _ := center.wrapped()
	C.iggDrawListAddCircle(list.handle(), centerArg, C.float(radius), C.uint(c), 0, C.float(thickness))
}

func (list DrawList) AddCircleFilled(center Vec2, radius float32, col Vec4) {
	c := GetColorU32(col)
	centerArg, _ := center.wrapped()
	C.iggDrawListAddCircleFilled(list.handle(), centerArg, C.float(radius), C.uint(c), 0)
}

func (list DrawList) AddQuad(p1, p2, p3, p4 Vec2, col Vec4, thickness float32) {
	c := GetColorU32(col)
	p1Arg, _ := p1.wrapped()
	p2Arg, _ := p2.wrapped()
	p3Arg, _ := p3.wrapped()
	p4Arg, _ := p4.wrapped()
	C.iggDrawListAddQuad(list.handle(), p1Arg, p2Arg, p3Arg, p4Arg, C.uint(c), C.float(thickness))
}

func (list DrawList) AddQuadFilled(p1, p2, p3, p4 Vec2, col Vec4) {
	c := GetColorU32(col)
	p1Arg, _ := p1.wrapped()
	p2Arg, _ := p2.wrapped()
	p3Arg, _ := p3.wrapped()
	p4Arg, _ := p4.wrapped()
	C.iggDrawListAddQuadFilled(list.handle(), p1Arg, p2Arg, p3Arg, p4Arg, C.uint(c))
}

// Stateful path API, add points then finish with PathFillConvex() or PathStroke()
func (list DrawList) PathClear() {
	C.iggDrawListPathClear(list.handle())
}

func (list DrawList) PathLineTo(pos Vec2) {
	posArg, _ := pos.wrapped()
	C.iggDrawListPathLineTo(list.handle(), posArg)
}

func (list DrawList) PathLineToMergeDuplicate(pos Vec2) {
	posArg, _ := pos.wrapped()
	C.iggDrawListPathLineToMergeDuplicate(list.handle(), posArg)
}

func (list DrawList) PathFillConvex(col Vec4) {
	C.iggDrawListPathFillConvex(list.handle(), C.uint(GetColorU32(col)))
}

func (list DrawList) PathStroke(col Vec4, closed bool, thickness float32) {
	C.iggDrawListPathStroke(list.handle(), C.uint(GetColorU32(col)), castBool(closed), C.float(thickness))
}

func (list DrawList) PathArcTo(center Vec2, radius, a_min, a_max float32, num_segments int) {
	centerArg, _ := center.wrapped()
	C.iggDrawListPathArcTo(list.handle(), centerArg, C.float(radius), C.float(a_min), C.float(a_max), C.int(num_segments))
}

func (list DrawList) PathArcToFast(center Vec2, radius float32, a_min_of_12, a_max_of_12 int) {
	centerArg, _ := center.wrapped()
	C.iggDrawListPathArcToFast(list.handle(), centerArg, C.float(radius), C.int(a_min_of_12), C.int(a_max_of_12))
}

func (list DrawList) PathBezierCurveTo(p1, p2, p3 Vec2, num_segments int) {
	p1Arg, _ := p1.wrapped()
	p2Arg, _ := p2.wrapped()
	p3Arg, _ := p3.wrapped()
	C.iggDrawListPathBezierCurveTo(list.handle(), p1Arg, p2Arg, p3Arg, C.int(num_segments))
}
