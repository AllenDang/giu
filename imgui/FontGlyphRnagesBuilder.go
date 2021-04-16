package imgui

// #include "FontGlyphRangesBuilderWrapper.h"
import "C"

type FontGlyphRangesBuilder uintptr

func NewGlyphRanges() GlyphRanges {
	handle := C.IggNewGlyphRanges()
	return GlyphRanges(handle)
}

func (ranges GlyphRanges) Data() GlyphRanges {
	return GlyphRanges(C.IggGlyphRangesData(ranges.handle()))
}

func NewFontGlyphRangesBuilder() FontGlyphRangesBuilder {
	handle := C.IggNewFontGlyphRangesBuilder()
	return FontGlyphRangesBuilder(handle)
}

func (builder FontGlyphRangesBuilder) handle() C.IggFontGlyphRangesBuilder {
	return C.IggFontGlyphRangesBuilder(builder)
}

func (builder FontGlyphRangesBuilder) AddText(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()

	C.IggFontGlyphRangesBuilderAddText(builder.handle(), textArg)
}

func (builder FontGlyphRangesBuilder) AddRanges(ranges GlyphRanges) {
	C.IggFontGlyphRangesBuilderAddRanges(builder.handle(), C.IggGlyphRanges(ranges))
}

func (builder FontGlyphRangesBuilder) Clear() {
	C.IggFontGlyphRangesBuilderClear(builder.handle())
}

func (builder FontGlyphRangesBuilder) BuildRanges(ranges GlyphRanges) {
	C.IggFontGlyphRangesBuilderBuildRanges(builder.handle(), C.IggGlyphRanges(ranges))
}
