package imgui

// #include "FontAtlasWrapper.h"
import "C"
import "unsafe"

// Alpha8Image represents a imgui backed 8-bit alpha value image.
type Alpha8Image struct {
	Width, Height int
	Pixels        unsafe.Pointer
}

// RGBA32Image represents a imgui backed 32-bit RGBA (8 bits per channel) value image.
type RGBA32Image struct {
	Width, Height int
	Pixels        unsafe.Pointer
}

// FontAtlas contains runtime data for multiple fonts,
// bake multiple fonts into a single texture, TTF/OTF font loader
type FontAtlas uintptr

func (atlas FontAtlas) handle() C.IggFontAtlas {
	return C.IggFontAtlas(atlas)
}

// GlyphRangesDefault describes Basic Latin, Extended Latin.
func (atlas FontAtlas) GlyphRangesDefault() GlyphRanges {
	return GlyphRanges(C.iggGetGlyphRangesDefault(atlas.handle()))
}

// GlyphRangesKorean describes Default + Korean characters.
func (atlas FontAtlas) GlyphRangesKorean() GlyphRanges {
	return GlyphRanges(C.iggGetGlyphRangesKorean(atlas.handle()))
}

// GlyphRangesJapanese describes Default + Hiragana, Katakana, Half-Width, Selection of 1946 Ideographs.
func (atlas FontAtlas) GlyphRangesJapanese() GlyphRanges {
	return GlyphRanges(C.iggGetGlyphRangesJapanese(atlas.handle()))
}

// GlyphRangesChineseFull describes Default + Half-Width + Japanese Hiragana/Katakana + full set of about 21000 CJK
// Unified Ideographs.
func (atlas FontAtlas) GlyphRangesChineseFull() GlyphRanges {
	return GlyphRanges(C.iggGetGlyphRangesChineseFull(atlas.handle()))
}

// GlyphRangesChineseSimplifiedCommon describes Default + Half-Width + Japanese Hiragana/Katakana + set of 2500 CJK
// Unified Ideographs for common simplified Chinese.
func (atlas FontAtlas) GlyphRangesChineseSimplifiedCommon() GlyphRanges {
	return GlyphRanges(C.iggGetGlyphRangesChineseSimplifiedCommon(atlas.handle()))
}

// GlyphRangesCyrillic describes Default + about 400 Cyrillic characters.
func (atlas FontAtlas) GlyphRangesCyrillic() GlyphRanges {
	return GlyphRanges(C.iggGetGlyphRangesCyrillic(atlas.handle()))
}

// GlyphRangesThai describes Default + Thai characters.
func (atlas FontAtlas) GlyphRangesThai() GlyphRanges {
	return GlyphRanges(C.iggGetGlyphRangesThai(atlas.handle()))
}

// AddFontDefault adds the default font to the atlas. This is done by default if you do not call any
// of the AddFont* methods before retrieving the texture data.
func (atlas FontAtlas) AddFontDefault() Font {
	fontHandle := C.iggAddFontDefault(atlas.handle())
	return Font(fontHandle)
}

// AddFontDefaultV adds the default font to the atlas using the specified FontConfig.
func (atlas FontAtlas) AddFontDefaultV(cfg FontConfig) Font {
	fontHandle := C.iggAddFontDefaultV(atlas.handle(), cfg.handle())
	return Font(fontHandle)
}

// AddFontFromFileTTFV attempts to load a font from given TTF file.
func (atlas FontAtlas) AddFontFromFileTTFV(filename string, sizePixels float32,
	config FontConfig, glyphRange GlyphRanges) Font {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()

	// Multiply sizePixels with DPI scale
	if DPIScale > 0 {
		sizePixels *= DPIScale
	}

	fontHandle := C.iggAddFontFromFileTTF(atlas.handle(), filenameArg, C.float(sizePixels),
		config.handle(), glyphRange.handle())
	return Font(fontHandle)
}

// AddFontFromFileTTF calls AddFontFromFileTTFV(filename, sizePixels, DefaultFontConfig, EmptyGlyphRanges).
func (atlas FontAtlas) AddFontFromFileTTF(filename string, sizePixels float32) Font {
	return atlas.AddFontFromFileTTFV(filename, sizePixels, DefaultFontConfig, EmptyGlyphRanges)
}

// AddFontFromMemoryTTFV attempts to load a font from given TTF byte array.
func (atlas FontAtlas) AddFontFromMemoryTTFV(
	fontData []byte, sizePixels float32,
	config FontConfig,
	glyphRange GlyphRanges,
) Font {
	// NOTE: We never free the fontDataC array because IMGUI's AddFontFromMemoryTTF takes ownership if
	// FontConfig.FontDataOwnedByAtlas == true (which it is by default). We do not expose this flag in Go
	// so we can assume in most cases it is true.
	if !config.getFontDataOwnedByAtlas() {
		panic("Only ImFontConfig.FontDataOwnedByAtlas == true is supported.")
	}

	fontDataC := C.malloc(C.size_t(len(fontData)))
	cBuf := (*[1 << 30]byte)(fontDataC)

	copy(cBuf[:], fontData)

	// Multiply sizePixels with DPI scale
	if DPIScale > 0 {
		sizePixels *= DPIScale
	}

	fontHandle := C.iggAddFontFromMemoryTTF(atlas.handle(), (*C.char)(fontDataC), C.int(len(fontData)), C.float(sizePixels),
		config.handle(), glyphRange.handle())

	return Font(fontHandle)
}

// AddFontFromMemoryTTF calls AddFontFromMemoryTTFV(fontData, sizePixels, DefaultFontConfig, EmptyGlyphRanges).
func (atlas FontAtlas) AddFontFromMemoryTTF(fontData []byte, sizePixels float32) Font {
	return atlas.AddFontFromMemoryTTFV(fontData, sizePixels, DefaultFontConfig, EmptyGlyphRanges)
}

// SetTexDesiredWidth registers the width desired by user before building the image. Must be a power-of-two.
// If have many glyphs your graphics API have texture size restrictions you may want to increase texture width to decrease height.
// Set to 0 by default, causing auto-calculation.
func (atlas FontAtlas) SetTexDesiredWidth(value int) {
	C.iggFontAtlasSetTexDesiredWidth(atlas.handle(), C.int(value))
}

// TextureDataAlpha8 returns the image in 8-bit alpha values for the font atlas.
// The returned image is valid as long as the font atlas is.
func (atlas FontAtlas) TextureDataAlpha8() *Alpha8Image {
	var pixels *C.uchar
	var width C.int
	var height C.int
	var bytesPerPixel C.int
	C.iggFontAtlasGetTexDataAsAlpha8(atlas.handle(), &pixels, &width, &height, &bytesPerPixel)

	return &Alpha8Image{
		Width:  int(width),
		Height: int(height),
		Pixels: unsafe.Pointer(pixels), // nolint: gas
	}
}

// TextureDataRGBA32 returns the image in 32-bit RGBA values for the font atlas.
// The returned image is valid as long as the font atlas is.
func (atlas FontAtlas) TextureDataRGBA32() *RGBA32Image {
	var pixels *C.uchar
	var width C.int
	var height C.int
	var bytesPerPixel C.int
	C.iggFontAtlasGetTexDataAsRGBA32(atlas.handle(), &pixels, &width, &height, &bytesPerPixel)

	return &RGBA32Image{
		Width:  int(width),
		Height: int(height),
		Pixels: unsafe.Pointer(pixels), // nolint: gas
	}
}

// SetTextureID sets user data to refer to the texture once it has been uploaded to user's graphic systems.
// It is passed back to you during rendering via the DrawCommand.
func (atlas FontAtlas) SetTextureID(id TextureID) {
	C.iggFontAtlasSetTextureID(atlas.handle(), id.handle())
}
