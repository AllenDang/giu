package giu

import (
	"image"
	"image/color"

	"github.com/AllenDang/imgui-go"
)

type Canvas struct {
	drawlist imgui.DrawList
}

func GetCanvas() *Canvas {
	return &Canvas{
		drawlist: imgui.GetWindowDrawList(),
	}
}

func (c *Canvas) AddLine(p1, p2 image.Point, color color.RGBA, thickness float32) {
	c.drawlist.AddLine(ToVec2(p1), ToVec2(p2), ToVec4Color(color), thickness)
}

type DrawFlags int

const (
	DrawFlags_None                    DrawFlags = 0
	DrawFlags_Closed                  DrawFlags = 1 << 0 // PathStroke(), AddPolyline(): specify that shape should be closed (portant: this is always == 1 for legacy reason)
	DrawFlags_RoundCornersTopLeft     DrawFlags = 1 << 4 // AddRect(), AddRectFilled(), PathRect(): enable rounding top-left corner only (when rounding > 0.0f, we default to all corners). Was 0x01.
	DrawFlags_RoundCornersTopRight    DrawFlags = 1 << 5 // AddRect(), AddRectFilled(), PathRect(): enable rounding top-right corner only (when rounding > 0.0f, we default to all corners). Was 0x02.
	DrawFlags_RoundCornersBottomLeft  DrawFlags = 1 << 6 // AddRect(), AddRectFilled(), PathRect(): enable rounding bottom-left corner only (when rounding > 0.0f, we default to all corners). Was 0x04.
	DrawFlags_RoundCornersBottomRight DrawFlags = 1 << 7 // AddRect(), AddRectFilled(), PathRect(): enable rounding bottom-right corner only (when rounding > 0.0f, we default to all corners). Wax 0x08.
	DrawFlags_RoundCornersNone        DrawFlags = 1 << 8 // AddRect(), AddRectFilled(), PathRect(): disable rounding on all corners (when rounding > 0.0f). This is NOT zero, NOT an implicit flag!
	DrawFlags_RoundCornersTop         DrawFlags = DrawFlags_RoundCornersTopLeft | DrawFlags_RoundCornersTopRight
	DrawFlags_RoundCornersBottom      DrawFlags = DrawFlags_RoundCornersBottomLeft | DrawFlags_RoundCornersBottomRight
	DrawFlags_RoundCornersLeft        DrawFlags = DrawFlags_RoundCornersBottomLeft | DrawFlags_RoundCornersTopLeft
	DrawFlags_RoundCornersRight       DrawFlags = DrawFlags_RoundCornersBottomRight | DrawFlags_RoundCornersTopRight
	DrawFlags_RoundCornersAll         DrawFlags = DrawFlags_RoundCornersTopLeft | DrawFlags_RoundCornersTopRight | DrawFlags_RoundCornersBottomLeft | DrawFlags_RoundCornersBottomRight
	DrawFlags_RoundCornersDefault_    DrawFlags = DrawFlags_RoundCornersAll // Default to ALL corners if none of the _RoundCornersXX flags are specified.
	DrawFlags_RoundCornersMask_       DrawFlags = DrawFlags_RoundCornersAll | DrawFlags_RoundCornersNone
)

func (c *Canvas) AddRect(pMin, pMax image.Point, color color.RGBA, rounding float32, rounding_corners DrawFlags, thickness float32) {
	c.drawlist.AddRect(ToVec2(pMin), ToVec2(pMax), ToVec4Color(color), rounding, int(rounding_corners), thickness)
}

func (c *Canvas) AddRectFilled(pMin, pMax image.Point, color color.RGBA, rounding float32, rounding_corners DrawFlags) {
	c.drawlist.AddRectFilled(ToVec2(pMin), ToVec2(pMax), ToVec4Color(color), rounding, int(rounding_corners))
}

func (c *Canvas) AddText(pos image.Point, color color.RGBA, text string) {
	c.drawlist.AddText(ToVec2(pos), ToVec4Color(color), text)
}

func (c *Canvas) AddBezierCubic(pos0, cp0, cp1, pos1 image.Point, color color.RGBA, thickness float32, num_segments int) {
	c.drawlist.AddBezierCubic(ToVec2(pos0), ToVec2(cp0), ToVec2(cp1), ToVec2(pos1), ToVec4Color(color), thickness, num_segments)
}

func (c *Canvas) AddTriangle(p1, p2, p3 image.Point, color color.RGBA, thickness float32) {
	c.drawlist.AddTriangle(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec4Color(color), thickness)
}

func (c *Canvas) AddTriangleFilled(p1, p2, p3 image.Point, color color.RGBA) {
	c.drawlist.AddTriangleFilled(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec4Color(color))
}

func (c *Canvas) AddCircle(center image.Point, radius float32, color color.RGBA, thickness float32) {
	c.drawlist.AddCircle(ToVec2(center), radius, ToVec4Color(color), thickness)
}

func (c *Canvas) AddCircleFilled(center image.Point, radius float32, color color.RGBA) {
	c.drawlist.AddCircleFilled(ToVec2(center), radius, ToVec4Color(color))
}

func (c *Canvas) AddQuad(p1, p2, p3, p4 image.Point, color color.RGBA, thickness float32) {
	c.drawlist.AddQuad(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec2(p4), ToVec4Color(color), thickness)
}

func (c *Canvas) AddQuadFilled(p1, p2, p3, p4 image.Point, color color.RGBA) {
	c.drawlist.AddQuadFilled(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec2(p4), ToVec4Color(color))
}

// Stateful path API, add points then finish with PathFillConvex() or PathStroke()

func (c *Canvas) PathClear() {
	c.drawlist.PathClear()
}

func (c *Canvas) PathLineTo(pos image.Point) {
	c.drawlist.PathLineTo(ToVec2(pos))
}

func (c *Canvas) PathLineToMergeDuplicate(pos image.Point) {
	c.drawlist.PathLineToMergeDuplicate(ToVec2(pos))
}

func (c *Canvas) PathFillConvex(color color.RGBA) {
	c.drawlist.PathFillConvex(ToVec4Color(color))
}

func (c *Canvas) PathStroke(color color.RGBA, closed bool, thickness float32) {
	c.drawlist.PathStroke(ToVec4Color(color), closed, thickness)
}

func (c *Canvas) PathArcTo(center image.Point, radius, a_min, a_max float32, num_segments int) {
	c.drawlist.PathArcTo(ToVec2(center), radius, a_min, a_max, num_segments)
}

func (c *Canvas) PathArcToFast(center image.Point, radius float32, a_min_of_12, a_max_of_12 int) {
	c.drawlist.PathArcToFast(ToVec2(center), radius, a_min_of_12, a_max_of_12)
}

func (c *Canvas) PathBezierCubicCurveTo(p1, p2, p3 image.Point, num_segments int) {
	c.drawlist.PathBezierCubicCurveTo(ToVec2(p1), ToVec2(p2), ToVec2(p3), num_segments)
}

func (c *Canvas) AddImage(texture *Texture, pMin, pMax image.Point) {
	c.drawlist.AddImage(texture.id, ToVec2(pMin), ToVec2(pMax))
}

func (c *Canvas) AddImageV(texture *Texture, pMin, pMax image.Point, uvMin, uvMax image.Point, color color.RGBA) {
	c.drawlist.AddImageV(texture.id, ToVec2(pMin), ToVec2(pMax), ToVec2(uvMin), ToVec2(uvMax), ToVec4Color(color))
}
