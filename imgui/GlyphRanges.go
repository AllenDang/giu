package imgui

// #include "imguiWrapperTypes.h"
import "C"
import "unsafe"

// GlyphRanges describes a list of Unicode ranges; 2 value per range, values are inclusive.
// Standard ranges can be queried from FontAtlas.GlyphRanges*() functions.
type GlyphRanges uintptr

// EmptyGlyphRanges is one that does not contain any ranges.
const EmptyGlyphRanges GlyphRanges = 0

func (glyphs GlyphRanges) handle() C.IggGlyphRanges {
	return C.IggGlyphRanges(glyphs)
}

type glyphRange struct{ from, to uint16 }

func (glyphs GlyphRanges) extract() (result []glyphRange) {
	if glyphs == 0 {
		return
	}
	rawSlice := (*[1 << 30]uint16)(unsafe.Pointer(glyphs.handle()))[:]
	index := 0
	// iterate until end of list or an arbitrary paranoia limit, should the list not be proper.
	for (rawSlice[index] != 0) && (index < 1000) {
		result = append(result, glyphRange{from: rawSlice[index+0], to: rawSlice[index+1]})
		index += 2
	}
	return
}
