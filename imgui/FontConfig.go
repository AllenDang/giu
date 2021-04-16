package imgui

// #include "FontConfigWrapper.h"
import "C"

// FontConfig describes properties of a single font.
type FontConfig uintptr

// DefaultFontConfig lets ImGui take default properties as per implementation.
// The properties of the default configuration cannot be changed using the SetXXX functions.
const DefaultFontConfig FontConfig = 0

func (config FontConfig) handle() C.IggFontConfig {
	return C.IggFontConfig(config)
}

// NewFontConfig creates a new font configuration.
// Delete must be called on the returned config.
func NewFontConfig() FontConfig {
	configHandle := C.iggNewFontConfig()
	return FontConfig(configHandle)
}

// Delete removes the font configuration and resets it to the DefaultFontConfig.
func (config *FontConfig) Delete() {
	if *config != DefaultFontConfig {
		C.iggFontConfigDelete(config.handle())
		*config = DefaultFontConfig
	}
}

// SetSize sets the size in pixels for rasterizer (more or less maps to the
// resulting font height).
func (config FontConfig) SetSize(sizePixels float32) {
	if config != DefaultFontConfig {
		C.iggFontConfigSetSize(config.handle(), C.float(sizePixels))
	}
}

// SetOversampleH sets the oversampling amount for the X axis.
// Rasterize at higher quality for sub-pixel positioning.
// We don't use sub-pixel positions on the Y axis.
func (config FontConfig) SetOversampleH(value int) {
	if config != DefaultFontConfig {
		C.iggFontConfigSetOversampleH(config.handle(), C.int(value))
	}
}

// SetOversampleV sets the oversampling amount for the Y axis.
// Rasterize at higher quality for sub-pixel positioning.
// We don't use sub-pixel positions on the Y axis.
func (config FontConfig) SetOversampleV(value int) {
	if config != DefaultFontConfig {
		C.iggFontConfigSetOversampleV(config.handle(), C.int(value))
	}
}

// SetPixelSnapH aligns every glyph to pixel boundary if enabled. Useful e.g. if
// you are merging a non-pixel aligned font with the default font. If enabled,
// you can set OversampleH/V to 1.
func (config FontConfig) SetPixelSnapH(value bool) {
	if config != DefaultFontConfig {
		C.iggFontConfigSetPixelSnapH(config.handle(), castBool(value))
	}
}

// SetGlyphMinAdvanceX sets the minimum AdvanceX for glyphs.
// Set Min to align font icons, set both Min/Max to enforce mono-space font.
func (config FontConfig) SetGlyphMinAdvanceX(value float32) {
	if config != DefaultFontConfig {
		C.iggFontConfigSetGlyphMinAdvanceX(config.handle(), C.float(value))
	}
}

// SetGlyphMaxAdvanceX sets the maximum AdvanceX for glyphs.
// Set both Min/Max to enforce mono-space font.
func (config FontConfig) SetGlyphMaxAdvanceX(value float32) {
	if config != DefaultFontConfig {
		C.iggFontConfigSetGlyphMaxAdvanceX(config.handle(), C.float(value))
	}
}

// SetMergeMode merges the new fonts into the previous font if enabled. This way
// you can combine multiple input fonts into one (e.g. ASCII font + icons +
// Japanese glyphs). You may want to use GlyphOffset.y when merge font of
// different heights.
func (config FontConfig) SetMergeMode(value bool) {
	if config != DefaultFontConfig {
		C.iggFontConfigSetMergeMode(config.handle(), castBool(value))
	}
}

// getFontDataOwnedByAtlas gets the current ownership status of the font data.
func (config FontConfig) getFontDataOwnedByAtlas() bool {
	if config != DefaultFontConfig {
		return C.iggFontConfigGetFontDataOwnedByAtlas(config.handle()) != 0
	}

	return true
}
