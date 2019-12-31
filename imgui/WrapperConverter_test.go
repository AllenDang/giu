package imgui

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringBufferAllocation(t *testing.T) {
	tt := []struct {
		initialValue string
		expectedSize int
	}{
		{initialValue: "", expectedSize: 1},
		{initialValue: "a", expectedSize: 2},
		{initialValue: "123456789", expectedSize: 10},
	}
	for _, tc := range tt {
		td := tc
		t.Run(fmt.Sprintf("<%s>", td.initialValue), func(t *testing.T) {
			buf := newStringBuffer(td.initialValue)
			defer buf.free()
			assert.Equal(t, td.expectedSize, buf.size)
		})
	}
}

func TestStringBufferStorage(t *testing.T) {
	tt := []string{"", "a", "ab", "SomeLongerText"}

	for _, tc := range tt {
		td := tc
		t.Run("Value <"+td+">", func(t *testing.T) {
			buf := newStringBuffer(td)
			require.NotNil(t, buf, "buffer expected")
			defer buf.free()
			result := buf.toGo()
			assert.Equal(t, td, result)
		})
	}
}

func TestStringBufferResize(t *testing.T) {
	tt := []struct {
		initialValue  string
		newSize       int
		expectedValue string
	}{
		{initialValue: "", newSize: 10, expectedValue: ""},
		{initialValue: "abcd", newSize: 10, expectedValue: "abcd"},
		{initialValue: "abcd", newSize: 3, expectedValue: "ab"},
		{initialValue: "efgh", newSize: 0, expectedValue: ""},
	}
	for _, tc := range tt {
		td := tc
		t.Run(fmt.Sprintf("<%s> -> %d", td.initialValue, tc.newSize), func(t *testing.T) {
			buf := newStringBuffer(td.initialValue)
			defer buf.free()
			buf.resizeTo(td.newSize)
			assert.Equal(t, td.expectedValue, buf.toGo())
		})
	}
}
