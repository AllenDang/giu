// +build !imguifreetype

package imgui

func (atlas FontAtlas) buildWithFreeType(flags int) error {
	return ErrFreeTypeNotAvailable
}
