package imgui

// FreeTypeError describes a problem with FreeType font rendering.
type FreeTypeError string

// Error returns the readable text presentation of the error.
func (err FreeTypeError) Error() string {
	return string(err)
}

const (
	// ErrFreeTypeNotAvailable is used if the implementation of freetype is not available in this build.
	ErrFreeTypeNotAvailable = FreeTypeError("Not available for this build")
	// ErrFreeTypeFailed is used if building a font atlas was not possible.
	ErrFreeTypeFailed = FreeTypeError("Failed to build FontAtlas with FreeType")
)
