package imgui

// #include "imguiWrapperTypes.h"
import "C"

// GlyphRanges describes a list of Unicode ranges; 2 value per range, values are inclusive.
// Standard ranges can be queried from FontAtlas.GlyphRanges*() functions.
type GlyphRanges uintptr

// EmptyGlyphRanges is one that does not contain any ranges.
const EmptyGlyphRanges GlyphRanges = 0

func (glyphs GlyphRanges) handle() C.IggGlyphRanges {
	return C.IggGlyphRanges(glyphs)
}
