package imgui

// Flags for FreeType rasterizer. By default, hinting is enabled and the font's native hinter is preferred over the auto-hinter.
const (
	// FreeTypeRasterizerFlagsNoHinting disables hinting.
	// This generally generates 'blurrier' bitmap glyphs when the glyph are rendered in any of the anti-aliased modes.
	FreeTypeRasterizerFlagsNoHinting = 1 << 0
	// FreeTypeRasterizerFlagsNoAutoHint disables auto-hinter.
	FreeTypeRasterizerFlagsNoAutoHint = 1 << 1
	// FreeTypeRasterizerFlagsForceAutoHint indicates that the auto-hinter is preferred over the font's native hinter.
	FreeTypeRasterizerFlagsForceAutoHint = 1 << 2
	// FreeTypeRasterizerFlagsLightHinting is a lighter hinting algorithm for gray-level modes.
	// Many generated glyphs are fuzzier but better resemble their original shape.
	// This is achieved by snapping glyphs to the pixel grid only vertically (Y-axis),
	// as is done by Microsoft's ClearType and Adobe's proprietary font renderer.
	// This preserves inter-glyph spacing in horizontal text.
	FreeTypeRasterizerFlagsLightHinting = 1 << 3
	// FreeTypeRasterizerFlagsMonoHinting is a strong hinting algorithm that should only be used for monochrome output.
	FreeTypeRasterizerFlagsMonoHinting = 1 << 4
	// FreeTypeRasterizerFlagsBold is for styling: Should we artificially embolden the font?
	FreeTypeRasterizerFlagsBold = 1 << 5
	// FreeTypeRasterizerFlagsOblique is for styling: Should we slant the font, emulating italic style?
	FreeTypeRasterizerFlagsOblique = 1 << 6
	// FreeTypeRasterizerFlagsMonochrome disables anti-aliasing. Combine this with MonoHinting for best results!
	FreeTypeRasterizerFlagsMonochrome = 1 << 7
)

var (
	EnableFreeType = false
)

// BuildWithFreeTypeV builds the FontAtlas using FreeType instead of the default rasterizer.
// FreeType renders small fonts better.
// Call this function instead of FontAtlas.Build() . As with FontAtlas.Build(), this function
// needs to be called before retrieving the texture data.
//
// FreeType support must be enabled with the build tag "imguifreetype".
func (atlas FontAtlas) BuildWithFreeTypeV(flags int) error {
	return atlas.buildWithFreeType(flags)
}

// BuildWithFreeType calls BuildWithFreeTypeV(0).
func (atlas FontAtlas) BuildWithFreeType() error {
	return atlas.BuildWithFreeTypeV(0)
}
