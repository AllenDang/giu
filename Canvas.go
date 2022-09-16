package giu

import (
	"image"
	"image/color"

	imgui "github.com/AllenDang/cimgui-go"
)

// Canvas represents imgui.DrawList
// from imgui.h:
//       A single draw command list (generally one per window,
//       conceptually you may see this as a dynamic "mesh" builder)
//
// for more details and use cases see examples/canvas.
type Canvas struct {
	drawlist imgui.ImDrawList
}

// GetCanvas returns current draw list (for current window).
// it will fail if called out of window's layout.
func GetCanvas() *Canvas {
	return &Canvas{
		drawlist: imgui.GetWindowDrawList(),
	}
}

// AddLine draws a line (from p1 to p2).
func (c *Canvas) AddLine(p1, p2 image.Point, col color.Color, thickness float32) {
	c.drawlist.AddLineV(ToVec2(p1), ToVec2(p2), ToU32(col), thickness)
}

// AddRect draws a rectangle.
func (c *Canvas) AddRect(pMin, pMax image.Point, col color.Color, rounding float32, flags imgui.ImDrawFlags, thickness float32) {
	c.drawlist.AddRectV(ToVec2(pMin), ToVec2(pMax), ToU32(col), rounding, flags, thickness)
}

// AddRectFilled draws a rectangle filled with `col`.
func (c *Canvas) AddRectFilled(pMin, pMax image.Point, col color.Color, rounding float32, flags imgui.ImDrawFlags) {
	c.drawlist.AddRectFilledV(ToVec2(pMin), ToVec2(pMax), ToU32(col), rounding, flags)
}

// AddText draws text.
func (c *Canvas) AddText(pos image.Point, col color.Color, text string) {
	c.drawlist.AddText_Vec2(ToVec2(pos), ToU32(col), Context.FontAtlas.RegisterString(text))
}

// AddBezierCubic draws bezier cubic.
func (c *Canvas) AddBezierCubic(pos0, cp0, cp1, pos1 image.Point, col color.Color, thickness float32, numSegments int32) {
	c.drawlist.AddBezierCubicV(ToVec2(pos0), ToVec2(cp0), ToVec2(cp1), ToVec2(pos1), ToU32(col), thickness, numSegments)
}

// AddTriangle draws a triangle.
func (c *Canvas) AddTriangle(p1, p2, p3 image.Point, col color.Color, thickness float32) {
	c.drawlist.AddTriangleV(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToU32(col), thickness)
}

// AddTriangleFilled draws a filled triangle.
func (c *Canvas) AddTriangleFilled(p1, p2, p3 image.Point, col color.Color) {
	c.drawlist.AddTriangleFilled(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToU32(col))
}

// AddCircle draws a circle.
func (c *Canvas) AddCircle(center image.Point, radius float32, col color.Color, segments int32, thickness float32) {
	c.drawlist.AddCircleV(ToVec2(center), radius, ToU32(col), segments, thickness)
}

// AddCircleFilled draws a filled circle.
func (c *Canvas) AddCircleFilled(center image.Point, radius float32, col color.Color, segments int32) {
	c.drawlist.AddCircleFilledV(ToVec2(center), radius, ToU32(col), segments)
}

// AddQuad draws a quad.
func (c *Canvas) AddQuad(p1, p2, p3, p4 image.Point, col color.Color, thickness float32) {
	c.drawlist.AddQuadV(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec2(p4), ToU32(col), thickness)
}

// AddQuadFilled draws a filled quad.
func (c *Canvas) AddQuadFilled(p1, p2, p3, p4 image.Point, col color.Color) {
	c.drawlist.AddQuadFilled(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec2(p4), ToU32(col))
}

// Stateful path API, add points then finish with PathFillConvex() or PathStroke().

func (c *Canvas) PathClear() {
	c.drawlist.PathClear()
}

func (c *Canvas) PathLineTo(pos image.Point) {
	c.drawlist.PathLineTo(ToVec2(pos))
}

func (c *Canvas) PathLineToMergeDuplicate(pos image.Point) {
	c.drawlist.PathLineToMergeDuplicate(ToVec2(pos))
}

func (c *Canvas) PathFillConvex(col color.Color) {
	c.drawlist.PathFillConvex(ToU32(col))
}

func (c *Canvas) PathStroke(col color.Color, flags imgui.ImDrawFlags, thickness float32) {
	c.drawlist.PathStrokeV(ToU32(col), flags, thickness)
}

func (c *Canvas) PathArcTo(center image.Point, radius, min, max float32, numSegments int32) {
	c.drawlist.PathArcToV(ToVec2(center), radius, min, max, numSegments)
}

func (c *Canvas) PathArcToFast(center image.Point, radius float32, min12, max12 int32) {
	c.drawlist.PathArcToFast(ToVec2(center), radius, min12, max12)
}

func (c *Canvas) PathBezierCubicCurveTo(p1, p2, p3 image.Point, numSegments int32) {
	c.drawlist.PathBezierCubicCurveToV(ToVec2(p1), ToVec2(p2), ToVec2(p3), numSegments)
}

func (c *Canvas) AddImage(texture imgui.ImTextureID, pMin, pMax, uvMin, uvMax image.Point, col color.Color) {
	c.drawlist.AddImageV(texture, ToVec2(pMin), ToVec2(pMax), ToVec2(uvMin), ToVec2(uvMax), ToU32(col))
}
