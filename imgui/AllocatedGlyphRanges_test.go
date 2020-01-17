package imgui_test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/AllenDang/giu/imgui"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAllocatedGlyphRangesFreeSetsToZero(t *testing.T) {
	var builder imgui.GlyphRangesBuilder
	ranges := builder.Build()
	require.NotEqual(t, uintptr(0), uintptr(ranges.GlyphRanges), "Should have allocated")
	ranges.Free()
	assert.Equal(t, uintptr(0), uintptr(ranges.GlyphRanges), "Should have reset")
}

func TestGlyphRangesBuilderAdd(t *testing.T) {
	type singleRange struct {
		from rune
		to   rune
	}
	tt := []struct {
		name     string
		input    []singleRange
		expected []uint16
	}{
		{name: "No input should contain terminator only", input: nil, expected: []uint16{0}},
		{name: "adding of single rune range", input: []singleRange{{from: 'A', to: 'A'}}, expected: []uint16{65, 65, 0}},
		{name: "adding of larger range", input: []singleRange{{from: 'A', to: 'C'}}, expected: []uint16{65, 67, 0}},
		{name: "adding two ranges", input: []singleRange{{from: 'A', to: 'C'}, {from: 'E', to: 'F'}}, expected: []uint16{65, 67, 69, 70, 0}},
		{name: "wrong order is ignored", input: []singleRange{{from: 'B', to: 'A'}}, expected: []uint16{0}},
	}

	for _, tc := range tt {
		var builder imgui.GlyphRangesBuilder
		for _, add := range tc.input {
			builder.Add(add.from, add.to)
		}
		result := builder.Build()
		require.NotEqual(t, uintptr(0), uintptr(result.GlyphRanges))
		for index, expected := range tc.expected {
			resultPtr := unsafe.Pointer(uintptr(result.GlyphRanges) + uintptr(2*index))
			resultShort := (*uint16)(resultPtr)
			assert.Equal(t, expected, *resultShort, fmt.Sprintf("%s: Index %d mismatch", tc.name, index))
		}
		result.Free()
	}
}

func TestGlyphRangesBuilderAddExisting(t *testing.T) {
	aRange := func(from, to rune) imgui.AllocatedGlyphRanges {
		var builder imgui.GlyphRangesBuilder
		builder.Add(from, to)
		return builder.Build()
	}
	baseRangeAA := aRange('A', 'A')
	defer baseRangeAA.Free()
	baseRangeEF := aRange('E', 'F')
	defer baseRangeEF.Free()

	tt := []struct {
		name     string
		input    []imgui.GlyphRanges
		expected []uint16
	}{
		{name: "adding of single range", input: []imgui.GlyphRanges{baseRangeAA.GlyphRanges}, expected: []uint16{65, 65, 0}},
		{name: "adding of two single ranges", input: []imgui.GlyphRanges{baseRangeAA.GlyphRanges, baseRangeEF.GlyphRanges}, expected: []uint16{65, 65, 69, 70, 0}},
	}

	for _, tc := range tt {
		var builder imgui.GlyphRangesBuilder
		for _, add := range tc.input {
			builder.AddExisting(add)
		}
		result := builder.Build()
		require.NotEqual(t, uintptr(0), uintptr(result.GlyphRanges))
		for index, expected := range tc.expected {
			resultPtr := unsafe.Pointer(uintptr(result.GlyphRanges) + uintptr(2*index))
			resultShort := (*uint16)(resultPtr)
			assert.Equal(t, expected, *resultShort, fmt.Sprintf("%s: Index %d mismatch", tc.name, index))
		}
		result.Free()
	}
}

func TestGlyphRangesBuilderCompactsRanges(t *testing.T) {
	type singleRange struct {
		from rune
		to   rune
	}
	tt := []struct {
		name     string
		input    []singleRange
		expected []uint16
	}{
		{name: "Removing duplications", input: []singleRange{{from: 'A', to: 'B'}, {from: 'A', to: 'B'}}, expected: []uint16{65, 66}},
		{name: "combining ranges A", input: []singleRange{{from: 'A', to: 'D'}, {from: 'B', to: 'E'}}, expected: []uint16{65, 69, 0}},
		{name: "combining ranges B", input: []singleRange{{from: 'C', to: 'E'}, {from: 'A', to: 'D'}}, expected: []uint16{65, 69, 0}},
		{name: "combining ranges C", input: []singleRange{{from: 'A', to: 'E'}, {from: 'B', to: 'C'}}, expected: []uint16{65, 69, 0}},
		{name: "combining ranges of whole set", input: []singleRange{
			{from: 'A', to: 'B'}, {from: 'E', to: 'F'}, {from: 'A', to: 'C'}}, expected: []uint16{65, 67, 69, 70, 0}},
	}

	for _, tc := range tt {
		var builder imgui.GlyphRangesBuilder
		for _, add := range tc.input {
			builder.Add(add.from, add.to)
		}
		result := builder.Build()
		require.NotEqual(t, uintptr(0), uintptr(result.GlyphRanges))
		for index, expected := range tc.expected {
			resultPtr := unsafe.Pointer(uintptr(result.GlyphRanges) + uintptr(2*index))
			resultShort := (*uint16)(resultPtr)
			assert.Equal(t, expected, *resultShort, fmt.Sprintf("%s: Index %d mismatch", tc.name, index))
		}
		result.Free()
	}
}
