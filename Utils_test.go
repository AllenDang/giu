package giu

import (
	"image/color"
	"testing"

	"github.com/AllenDang/imgui-go"
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
			expected: imgui.Vec4{1, 0, 0, 1},
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
