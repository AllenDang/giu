package imgui

import (
	. "github.com/inkyblackness/imgui-go/v3"
	"image"
)

// Renderer covers rendering imgui draw data.
type Renderer interface {
	// PreRender causes the display buffer to be prepared for new output.
	PreRender(clearColor [4]float32)
	// Render draws the provided imgui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData DrawData)
	// Sets the texture minifying filter.
	SetTextureMinFilter(min uint) error
	// Sets the texture magnifying filter.
	SetTextureMagFilter(mag uint) error
	// Load image and return the TextureID
	LoadImage(image *image.RGBA) (TextureID, error)
	// Release image
	ReleaseImage(textureId TextureID)
	// Dispose
	Dispose()
}
