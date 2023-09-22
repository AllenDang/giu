package giu

import (
	"image"
	"image/color"
	"testing"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/stretchr/testify/assert"
)

func Test_ToVec4(t *testing.T) {
	tests := []struct {
		name     string
		source   color.Color
		expected imgui.Vec4
	}{
		{
			name:     "Red - RGBA",
			source:   &color.RGBA{R: 255, G: 0, B: 0, A: 255},
			expected: imgui.Vec4{X: 1, Y: 0, Z: 0, W: 1},
		},
		{
			name:     "Purple - RGBA",
			source:   &color.RGBA{R: 158, G: 0, B: 173, A: 255},
			expected: imgui.Vec4{X: 0.61960787, Y: 0, Z: 0.6784314, W: 1},
		},
		{
			name:     "Purple - CMYK",
			source:   &color.CMYK{C: 22, M: 255, Y: 0, K: 82},
			expected: imgui.Vec4{X: 0.6198978, Y: 0, Z: 0.6784314, W: 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt, test.expected, ToVec4Color(test.source), "Unexpected result")
		})
	}
}

func Test_ToVec2(t *testing.T) {
	tests := []struct {
		name     string
		source   image.Point
		expected imgui.Vec2
	}{
		{"Point 0,0", image.Pt(0, 0), imgui.Vec2{X: 0, Y: 0}},
		{"Random point 1", image.Pt(80, 209), imgui.Vec2{X: 80, Y: 209}},
		{"Random point 2", image.Pt(200, 128), imgui.Vec2{X: 200, Y: 128}},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt, test.expected, ToVec2(test.source), "Unexpected result")
		})
	}
}

func Test_Vec4ToRGBA(t *testing.T) {
	tests := []struct {
		name     string
		source   imgui.Vec4
		expected color.RGBA
	}{
		{
			name:     "Red",
			source:   imgui.Vec4{X: 1, Y: 0, Z: 0, W: 1},
			expected: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		},
		{
			name:     "Red - with 20% alpha",
			source:   imgui.Vec4{X: 1, Y: 0, Z: 0, W: 0.2},
			expected: color.RGBA{R: 255, G: 0, B: 0, A: 51},
		},
		{
			name:     "Purple - RGBA",
			source:   imgui.Vec4{X: 0.61960787, Y: 0, Z: 0.6784314, W: 1},
			expected: color.RGBA{R: 158, G: 0, B: 173, A: 255},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt, test.expected, Vec4ToRGBA(test.source), "Unexpected result")
		})
	}
}

func Test_Assert(t *testing.T) {
	tests := []struct {
		name        string
		condition   bool
		shouldPanic bool
	}{
		{"expected behavior - no panic", true, false},
		{"something happened - panic", false, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.shouldPanic {
				assert.Panics(tt, func() { Assert(test.condition, "somewidget", "somemethod", "panics") }, "unexpected behavior")
			} else {
				assert.NotPanics(tt, func() { Assert(test.condition, "somewidget", "somemethod", "panics") }, "unexpected behavior")
			}
		})
	}
}
