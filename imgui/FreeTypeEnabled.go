// +build imguifreetype

package imgui

// #cgo pkg-config: freetype2
// #cgo CXXFLAGS: -DIMGUI_FREETYPE_ENABLED
// #cgo CFLAGS: -DIMGUI_FREETYPE_ENABLED
// #cgo CPPFLAGS: -DIMGUI_FREETYPE_ENABLED
// #include "FreeTypeWrapper.h"
import "C"

func (atlas FontAtlas) buildWithFreeType(flags int) error {
	if C.iggFreeTypeBuildFontAtlas(atlas.handle(), C.uint(flags)) == 0 {
		return ErrFreeTypeFailed
	}
	return nil
}
