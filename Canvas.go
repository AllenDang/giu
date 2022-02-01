package giu

import (
	"image"
	"image/color"

	"github.com/AllenDang/imgui-go"
)

// Canvas represents imgui.DrawList
// from imgui.h:
//       A single draw command list (generally one per window,
//       conceptually you may see this as a dynamic "mesh" builder)
//
// for more details and use cases see examples/canvas.
type Canvas struct {
	drawlist imgui.DrawList
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
	c.drawlist.AddLine(ToVec2(p1), ToVec2(p2), ToVec4Color(col), thickness)
}

// DrawFlags represents imgui.DrawFlags.
type DrawFlags int

// draw flags enum:.
const (
	DrawFlagsNone DrawFlags = 0
	// PathStroke(), AddPolyline(): specify that shape should be closed (portant: this is always == 1 for legacy reason).
	DrawFlagsClosed DrawFlags = 1 << 0
	// AddRect(), AddRectFilled(), PathRect(): enable rounding top-left corner only (when rounding > 0.0f, we default to all corners).
	// Was 0x01.
	DrawFlagsRoundCornersTopLeft DrawFlags = 1 << 4
	// AddRect(), AddRectFilled(), PathRect(): enable rounding top-right corner only (when rounding > 0.0f, we default to all corners).
	// Was 0x02.
	DrawFlagsRoundCornersTopRight DrawFlags = 1 << 5
	// AddRect(), AddRectFilled(), PathRect(): enable rounding bottom-left corner only (when rounding > 0.0f, we default to all corners).
	// Was 0x04.
	DrawFlagsRoundCornersBottomLeft DrawFlags = 1 << 6
	// AddRect(), AddRectFilled(), PathRect(): enable rounding bottom-right corner only (when rounding > 0.0f,
	// we default to all corners). Wax 0x08.
	DrawFlagsRoundCornersBottomRight DrawFlags = 1 << 7
	// AddRect(), AddRectFilled(), PathRect(): disable rounding on all corners (when rounding > 0.0f). This is NOT zero, NOT an implicit flag!
	DrawFlagsRoundCornersNone   DrawFlags = 1 << 8
	DrawFlagsRoundCornersTop    DrawFlags = DrawFlagsRoundCornersTopLeft | DrawFlagsRoundCornersTopRight
	DrawFlagsRoundCornersBottom DrawFlags = DrawFlagsRoundCornersBottomLeft | DrawFlagsRoundCornersBottomRight
	DrawFlagsRoundCornersLeft   DrawFlags = DrawFlagsRoundCornersBottomLeft | DrawFlagsRoundCornersTopLeft
	DrawFlagsRoundCornersRight  DrawFlags = DrawFlagsRoundCornersBottomRight | DrawFlagsRoundCornersTopRight
	DrawFlagsRoundCornersAll    DrawFlags = DrawFlagsRoundCornersTopLeft | DrawFlagsRoundCornersTopRight |
		DrawFlagsRoundCornersBottomLeft | DrawFlagsRoundCornersBottomRight
	// Default to ALL corners if none of the RoundCornersXX flags are specified.
	DrawFlagsRoundCornersDefault DrawFlags = DrawFlagsRoundCornersAll
	DrawFlagsRoundCornersMask    DrawFlags = DrawFlagsRoundCornersAll | DrawFlagsRoundCornersNone
)

// AddRect draws a rectangle.
func (c *Canvas) AddRect(pMin, pMax image.Point, col color.Color, rounding float32, roundingCorners DrawFlags, thickness float32) {
	c.drawlist.AddRect(ToVec2(pMin), ToVec2(pMax), ToVec4Color(col), rounding, int(roundingCorners), thickness)
}

// AddRectFilled draws a rectangle filled with `col`.
func (c *Canvas) AddRectFilled(pMin, pMax image.Point, col color.Color, rounding float32, roundingCorners DrawFlags) {
	c.drawlist.AddRectFilled(ToVec2(pMin), ToVec2(pMax), ToVec4Color(col), rounding, int(roundingCorners))
}

// AddText draws text.
func (c *Canvas) AddText(pos image.Point, col color.Color, text string) {
	c.drawlist.AddText(ToVec2(pos), ToVec4Color(col), tStr(text))
}

// AddBezierCubic draws bezier cubic.
func (c *Canvas) AddBezierCubic(pos0, cp0, cp1, pos1 image.Point, col color.Color, thickness float32, numSegments int) {
	c.drawlist.AddBezierCubic(ToVec2(pos0), ToVec2(cp0), ToVec2(cp1), ToVec2(pos1), ToVec4Color(col), thickness, numSegments)
}

// AddTriangle draws a triangle.
func (c *Canvas) AddTriangle(p1, p2, p3 image.Point, col color.Color, thickness float32) {
	c.drawlist.AddTriangle(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec4Color(col), thickness)
}

// AddTriangleFilled draws a filled triangle.
func (c *Canvas) AddTriangleFilled(p1, p2, p3 image.Point, col color.Color) {
	c.drawlist.AddTriangleFilled(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec4Color(col))
}

// AddCircle draws a circle.
func (c *Canvas) AddCircle(center image.Point, radius float32, col color.Color, segments int, thickness float32) {
	c.drawlist.AddCircle(ToVec2(center), radius, ToVec4Color(col), segments, thickness)
}

// AddCircleFilled draws a filled circle.
func (c *Canvas) AddCircleFilled(center image.Point, radius float32, col color.Color) {
	c.drawlist.AddCircleFilled(ToVec2(center), radius, ToVec4Color(col))
}

// AddQuad draws a quad.
func (c *Canvas) AddQuad(p1, p2, p3, p4 image.Point, col color.Color, thickness float32) {
	c.drawlist.AddQuad(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec2(p4), ToVec4Color(col), thickness)
}

// AddQuadFilled draws a filled quad.
func (c *Canvas) AddQuadFilled(p1, p2, p3, p4 image.Point, col color.Color) {
	c.drawlist.AddQuadFilled(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec2(p4), ToVec4Color(col))
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
	c.drawlist.PathFillConvex(ToVec4Color(col))
}

func (c *Canvas) PathStroke(col color.Color, closed bool, thickness float32) {
	c.drawlist.PathStroke(ToVec4Color(col), closed, thickness)
}

func (c *Canvas) PathArcTo(center image.Point, radius, min, max float32, numSegments int) {
	c.drawlist.PathArcTo(ToVec2(center), radius, min, max, numSegments)
}

func (c *Canvas) PathArcToFast(center image.Point, radius float32, min12, max12 int) {
	c.drawlist.PathArcToFast(ToVec2(center), radius, min12, max12)
}

func (c *Canvas) PathBezierCubicCurveTo(p1, p2, p3 image.Point, numSegments int) {
	c.drawlist.PathBezierCubicCurveTo(ToVec2(p1), ToVec2(p2), ToVec2(p3), numSegments)
}

func (c *Canvas) AddImage(texture *Texture, pMin, pMax image.Point) {
	c.drawlist.AddImage(texture.id, ToVec2(pMin), ToVec2(pMax))
}

func (c *Canvas) AddImageV(texture *Texture, pMin, pMax, uvMin, uvMax image.Point, col color.Color) {
	c.drawlist.AddImageV(texture.id, ToVec2(pMin), ToVec2(pMax), ToVec2(uvMin), ToVec2(uvMax), ToVec4Color(col))
}
