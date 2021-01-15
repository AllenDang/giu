package giu

import (
	"image"
	"image/color"

	"github.com/ianling/imgui-go"
)

type Canvas struct {
	drawlist imgui.DrawList
}

func GetCanvas() *Canvas {
	return &Canvas{
		drawlist: imgui.WindowDrawList(),
	}
}

func (c *Canvas) AddLine(p1, p2 image.Point, color color.RGBA, thickness float32) {
	c.drawlist.AddLineV(ToVec2(p1), ToVec2(p2), imgui.Packed(color), thickness)
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

func (c *Canvas) AddRect(pMin, pMax image.Point, color color.RGBA, rounding float32, roundingCorners CornerFlags, thickness float32) {
	c.drawlist.AddRectV(ToVec2(pMin), ToVec2(pMax), imgui.Packed(color), rounding, int(roundingCorners), thickness)
}

func (c *Canvas) AddRectFilled(pMin, pMax image.Point, color color.RGBA, rounding float32, roundingCorners CornerFlags) {
	c.drawlist.AddRectFilledV(ToVec2(pMin), ToVec2(pMax), imgui.Packed(color), rounding, int(roundingCorners))
}

func (c *Canvas) AddText(pos image.Point, color color.RGBA, text string) {
	c.drawlist.AddText(ToVec2(pos), imgui.Packed(color), text)
}

func (c *Canvas) AddBezierCurve(pos0, cp0, cp1, pos1 image.Point, color color.RGBA, thickness float32, numSegments int) {
	c.drawlist.AddBezierCurve(ToVec2(pos0), ToVec2(cp0), ToVec2(cp1), ToVec2(pos1), imgui.Packed(color), thickness, numSegments)
}

func (c *Canvas) AddTriangle(p1, p2, p3 image.Point, color color.RGBA, thickness float32) {
	c.drawlist.AddTriangleV(ToVec2(p1), ToVec2(p2), ToVec2(p3), imgui.Packed(color), thickness)
}

func (c *Canvas) AddTriangleFilled(p1, p2, p3 image.Point, color color.RGBA) {
	c.drawlist.AddTriangleFilled(ToVec2(p1), ToVec2(p2), ToVec2(p3), imgui.Packed(color))
}

func (c *Canvas) AddCircle(center image.Point, radius float32, color color.RGBA, thickness float32, numSegments int) {
	c.drawlist.AddCircleV(ToVec2(center), radius, imgui.Packed(color), numSegments, thickness)
}

func (c *Canvas) AddCircleFilled(center image.Point, radius float32, color color.RGBA) {
	c.drawlist.AddCircleFilled(ToVec2(center), radius, imgui.Packed(color))
}

func (c *Canvas) AddQuad(p1, p2, p3, p4 image.Point, color color.RGBA, thickness float32) {
	c.drawlist.AddQuad(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec2(p4), imgui.Packed(color), thickness)
}

func (c *Canvas) AddQuadFilled(p1, p2, p3, p4 image.Point, color color.RGBA) {
	c.drawlist.AddQuadFilled(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec2(p4), imgui.Packed(color))
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
	c.drawlist.PathFillConvex(imgui.Packed(color))
}

func (c *Canvas) PathStroke(color color.RGBA, closed bool, thickness float32) {
	c.drawlist.PathStroke(imgui.Packed(color), closed, thickness)
}

func (c *Canvas) PathArcTo(center image.Point, radius, aMin, aMax float32, numSegments int) {
	c.drawlist.PathArcTo(ToVec2(center), radius, aMin, aMax, numSegments)
}

func (c *Canvas) PathArcToFast(center image.Point, radius float32, aMinOf12, aMaxOf12 int) {
	c.drawlist.PathArcToFast(ToVec2(center), radius, aMinOf12, aMaxOf12)
}

func (c *Canvas) PathBezierCurveTo(p1, p2, p3 image.Point, numSegments int) {
	c.drawlist.PathBezierCurveTo(ToVec2(p1), ToVec2(p2), ToVec2(p3), numSegments)
}

func (c *Canvas) AddImage(texture *Texture, pMin, pMax image.Point) {
	c.drawlist.AddImage(texture.id, ToVec2(pMin), ToVec2(pMax))
}

func (c *Canvas) AddImageV(texture *Texture, pMin, pMax image.Point, uvMin, uvMax image.Point, color color.RGBA) {
	c.drawlist.AddImageV(texture.id, ToVec2(pMin), ToVec2(pMax), ToVec2(uvMin), ToVec2(uvMax), imgui.Packed(color))
}
