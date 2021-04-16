package imgui_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AllenDang/giu/imgui"
)

func TestVec2Addition(t *testing.T) {
	tt := []struct {
		first    imgui.Vec2
		second   imgui.Vec2
		expected imgui.Vec2
	}{
		{first: imgui.Vec2{X: 0, Y: 0}, second: imgui.Vec2{X: 10, Y: 10}, expected: imgui.Vec2{X: 10, Y: 10}},
		{first: imgui.Vec2{X: 10, Y: 10}, second: imgui.Vec2{X: -20, Y: 100.5}, expected: imgui.Vec2{X: -10, Y: 110.5}},
		{first: imgui.Vec2{X: 2, Y: 4}, second: imgui.Vec2{X: -2, Y: -4}, expected: imgui.Vec2{X: 0, Y: 0}},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, tc.first.Plus(tc.second), "Failed for %v + %v", tc.first, tc.second)
	}
}

func TestVec2Subtraction(t *testing.T) {
	tt := []struct {
		first    imgui.Vec2
		second   imgui.Vec2
		expected imgui.Vec2
	}{
		{first: imgui.Vec2{X: 0, Y: 0}, second: imgui.Vec2{X: 10, Y: 10}, expected: imgui.Vec2{X: -10, Y: -10}},
		{first: imgui.Vec2{X: 10, Y: 10}, second: imgui.Vec2{X: -20, Y: 100.5}, expected: imgui.Vec2{X: 30, Y: -90.5}},
		{first: imgui.Vec2{X: 2, Y: 4}, second: imgui.Vec2{X: -2, Y: -4}, expected: imgui.Vec2{X: 4, Y: 8}},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, tc.first.Minus(tc.second), "Failed for %v - %v", tc.first, tc.second)
	}
}

func TestVec2Scaling(t *testing.T) {
	tt := []struct {
		first    imgui.Vec2
		scale    float32
		expected imgui.Vec2
	}{
		{first: imgui.Vec2{X: 1, Y: 2}, scale: 1.0, expected: imgui.Vec2{X: 1, Y: 2}},
		{first: imgui.Vec2{X: 10, Y: 10}, scale: 0.0, expected: imgui.Vec2{X: 0, Y: 0}},
		{first: imgui.Vec2{X: 2, Y: 4}, scale: -2.0, expected: imgui.Vec2{X: -4, Y: -8}},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, tc.first.Times(tc.scale), "Failed for %v * %v", tc.first, tc.scale)
	}
}

func TestVec4Addition(t *testing.T) {
	vec := func(x, y, z, w float32) imgui.Vec4 {
		return imgui.Vec4{X: x, Y: y, Z: z, W: w}
	}
	tt := []struct {
		first    imgui.Vec4
		second   imgui.Vec4
		expected imgui.Vec4
	}{
		{first: vec(0, 0, 0, 0), second: vec(10, 20, 30, 40), expected: vec(10, 20, 30, 40)},
		{first: vec(10, 20, 30, 40), second: vec(-5, -5, -5, -5), expected: vec(5, 15, 25, 35)},
		{first: vec(1, 2, 3, 4), second: vec(-10, -10, -10, -10), expected: vec(-9, -8, -7, -6)},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, tc.first.Plus(tc.second), "Failed for %v + %v", tc.first, tc.second)
	}
}

func TestVec4Subtraction(t *testing.T) {
	vec := func(x, y, z, w float32) imgui.Vec4 {
		return imgui.Vec4{X: x, Y: y, Z: z, W: w}
	}
	tt := []struct {
		first    imgui.Vec4
		second   imgui.Vec4
		expected imgui.Vec4
	}{
		{first: vec(0, 0, 0, 0), second: vec(10, 20, 30, 40), expected: vec(-10, -20, -30, -40)},
		{first: vec(10, 20, 30, 40), second: vec(-5, -5, -5, -5), expected: vec(15, 25, 35, 45)},
		{first: vec(1, 2, 3, 4), second: vec(-10, -10, -10, -10), expected: vec(11, 12, 13, 14)},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, tc.first.Minus(tc.second), "Failed for %v - %v", tc.first, tc.second)
	}
}

func TestVec4Scaling(t *testing.T) {
	vec := func(x, y, z, w float32) imgui.Vec4 {
		return imgui.Vec4{X: x, Y: y, Z: z, W: w}
	}
	tt := []struct {
		first    imgui.Vec4
		scale    float32
		expected imgui.Vec4
	}{
		{first: vec(10, 20, 30, 40), scale: 1.0, expected: vec(10, 20, 30, 40)},
		{first: vec(10, 20, 30, 40), scale: -0.5, expected: vec(-5, -10, -15, -20)},
		{first: vec(1, 2, 3, 4), scale: 2.0, expected: vec(2, 4, 6, 8)},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, tc.first.Times(tc.scale), "Failed for %v * %v", tc.first, tc.scale)
	}
}
