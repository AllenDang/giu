package giu

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/image/colornames"
)

func TestStyleSetter_Add(t *testing.T) {
	cases := []struct {
		name     string
		setter   *StyleSetter
		other    *StyleSetter
		expected *StyleSetter
	}{
		{
			name: "Adding nil",
			setter: Style().SetColor(StyleColorText, colornames.Red).
				SetStyle(StyleVarWindowPadding, 10, 10),
			other: nil,
			expected: Style().SetColor(StyleColorText, colornames.Red).
				SetStyle(StyleVarWindowPadding, 10, 10),
		},
		{
			name: "Adding some styles/colors",
			setter: Style().SetColor(StyleColorText, colornames.Red).
				SetStyle(StyleVarWindowPadding, 10, 10),
			other: Style().SetColor(StyleColorWindowBg, colornames.Red).
				SetStyle(StyleVarFramePadding, 10, 10),
			expected: Style().SetColor(StyleColorText, colornames.Red).
				SetStyle(StyleVarWindowPadding, 10, 10).
				SetColor(StyleColorWindowBg, colornames.Red).
				SetStyle(StyleVarFramePadding, 10, 10),
		},
		{
			name: "Overwrite",
			setter: Style().SetColor(StyleColorText, colornames.Red).
				SetStyle(StyleVarWindowPadding, 10, 10),
			other: Style().SetColor(StyleColorText, colornames.Blue).
				SetStyle(StyleVarWindowPadding, 11, 11),
			expected: Style().SetColor(StyleColorText, colornames.Blue).
				SetStyle(StyleVarWindowPadding, 11, 11),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.setter.Add(c.other)
			assert.True(t, assert.ObjectsAreEqual(c.setter, c.expected), "Expected: %v, got: %v", c.expected, c.setter)
		})
	}
}
