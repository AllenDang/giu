package giu

import (
	"image"
	"image/color"

	"github.com/AllenDang/giu/imgui"
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

type CornerFlags int

const (
	CornerFlags_None     CornerFlags = 0
	CornerFlags_TopLeft  CornerFlags = 1 << 0                                      // 0x1
	CornerFlags_TopRight CornerFlags = 1 << 1                                      // 0x2
	CornerFlags_BotLeft  CornerFlags = 1 << 2                                      // 0x4
	CornerFlags_BotRight CornerFlags = 1 << 3                                      // 0x8
	CornerFlags_Top      CornerFlags = CornerFlags_TopLeft | CornerFlags_TopRight  // 0x3
	CornerFlags_Bot      CornerFlags = CornerFlags_BotLeft | CornerFlags_BotRight  // 0xC
	CornerFlags_Left     CornerFlags = CornerFlags_TopLeft | CornerFlags_BotLeft   // 0x5
	CornerFlags_Right    CornerFlags = CornerFlags_TopRight | CornerFlags_BotRight // 0xA
	CornerFlags_All      CornerFlags = 0xF                                         // In your function calls you may use ~0 (= all bits sets) instead of ImDrawCornerFlags_All, as a convenience

)

func (c *Canvas) AddRect(pMin, pMax image.Point, color color.RGBA, rounding float32, rounding_corners CornerFlags, thickness float32) {
	c.drawlist.AddRect(ToVec2(pMin), ToVec2(pMax), ToVec4Color(color), rounding, int(rounding_corners), thickness)
}

func (c *Canvas) AddRectFilled(pMin, pMax image.Point, color color.RGBA, rounding float32, rounding_corners CornerFlags) {
	c.drawlist.AddRectFilled(ToVec2(pMin), ToVec2(pMax), ToVec4Color(color), rounding, int(rounding_corners))
}

func (c *Canvas) AddText(pos image.Point, color color.RGBA, text string) {
	c.drawlist.AddText(ToVec2(pos), ToVec4Color(color), text)
}

func (c *Canvas) AddBezierCurve(pos0, cp0, cp1, pos1 image.Point, color color.RGBA, thickness float32, num_segments int) {
	c.drawlist.AddBezierCurve(ToVec2(pos0), ToVec2(cp0), ToVec2(cp1), ToVec2(pos1), ToVec4Color(color), thickness, num_segments)
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

func (c *Canvas) PathBezierCurveTo(p1, p2, p3 image.Point, num_segments int) {
	c.drawlist.PathBezierCurveTo(ToVec2(p1), ToVec2(p2), ToVec2(p3), num_segments)
}
