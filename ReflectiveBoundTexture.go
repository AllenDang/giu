package giu

import (
	"errors"
	"hash/crc32"
	"image"
	"image/color"
	"sync"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/imgui"
)

// ErrNilRGBA is an error that indicates the RGBA surface result is nil.
var ErrNilRGBA = errors.New("surface RGBA Result is nil")

// defaultSurface returns a default RGBA surface with a uniform color.
func defaultSurface() *image.RGBA {
	surface, _ := NewUniformLoader(128.0, 128.0, color.RGBA{255, 255, 255, 255}).ServeRGBA()
	return surface
}

// ReflectiveBoundTexture represents a texture that can be dynamically updated and bound to a GPU.
type ReflectiveBoundTexture struct {
	Surface *image.RGBA // Surface is the RGBA image data for the texture.
	tex     *Texture    // tex is the GPU texture resource.
	lastSum uint32      // lastSum is the checksum of the last surface data.
	mu      sync.Mutex  // mu is a mutex to protect concurrent access to the texture.
	fsroot  string      // fsroot is the root filesystem path for the texture when using file scheme in URL Surface Loader
}

/* Return a waranted:
 * - Initialized
 * - With proper resources bindings against gpu (free old, bound new)
 * - Up to date Texture.
 */

// commit updates the ReflectiveBoundTexture by checking if the surface has changed,
// and if so, rebinds the texture to the GPU. It returns the updated ReflectiveBoundTexture
// and a boolean indicating whether the texture has changed.
//
// The method locks the texture for concurrent access, calculates the checksum of the
// current surface, and compares it with the last stored checksum. If the checksums differ,
// it unbinds the old texture, binds the new one, and updates the checksum.
//
// Returns:
//   - *ReflectiveBoundTexture: The updated texture object.
//   - bool: True if the texture has changed, false otherwise.
func (i *ReflectiveBoundTexture) commit() (*ReflectiveBoundTexture, bool) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.Surface == nil {
		i.Surface = defaultSurface()
	}

	var hasChanged bool
	if sum := crc32.ChecksumIEEE(i.Surface.Pix); sum != i.lastSum {
		hasChanged = true

		i.unbind()
		i.bind()
		i.lastSum = sum
	}

	return i, hasChanged
}

// SetSurfaceFromRGBA sets the surface of the ReflectiveBoundTexture from the provided RGBA image.
// If the provided image is nil, it returns an error. If the commit flag is true, it commits the changes.
//
// Parameters:
//   - img: The RGBA image to set as the surface.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the provided image is nil, otherwise nil.
func (i *ReflectiveBoundTexture) SetSurfaceFromRGBA(img *image.RGBA, commit bool) error {
	if img != nil {
		i.Surface = img
	} else {
		return ErrNilRGBA
	}

	if commit {
		i.commit()
	}

	return nil
}

// ToImageWidget converts the ReflectiveBoundTexture to an ImageWidget.
//
// Returns:
//   - *ImageWidget: The ImageWidget representation of the ReflectiveBoundTexture.
func (i *ReflectiveBoundTexture) ToImageWidget() *ImageWidget {
	return Image(i.Texture())
}

// ImguiImageVOptionStruct represents the options for rendering an image in ImGui.
type ImguiImageVOptionStruct struct {
	Uv0       imgui.Vec2 // The UV coordinate of the top-left corner of the image.
	Uv1       imgui.Vec2 // The UV coordinate of the bottom-right corner of the image.
	TintCol   imgui.Vec4 // The tint color to apply to the image.
	BorderCol imgui.Vec4 // The border color to apply to the image.
}

// GetImGuiImageVDefaultOptionsStruct returns the default options for rendering an image in ImGui.
//
// Returns:
//   - ImguiImageVOptionStruct: The default options for rendering an image.
func (i *ReflectiveBoundTexture) GetImGuiImageVDefaultOptionsStruct() ImguiImageVOptionStruct {
	return ImguiImageVOptionStruct{
		Uv0:       imgui.Vec2{X: 0, Y: 0},
		Uv1:       imgui.Vec2{X: 1, Y: 1},
		TintCol:   imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1},
		BorderCol: imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0},
	}
}

// ImguiImage renders the ReflectiveBoundTexture as an image in ImGui.
//
// Parameters:
//   - width: The width of the image. If set to -1, it will use the available content region width.
//   - height: The height of the image. If set to -1, it will use the available content region height.
//   - options: The options for rendering the image.
func (i *ReflectiveBoundTexture) ImguiImage(width, height float32) {
	size := imgui.Vec2{X: width, Y: height}

	if size.X == -1 {
		rect := imgui.ContentRegionAvail()
		size.X = rect.X
	}

	if size.Y == -1 {
		rect := imgui.ContentRegionAvail()
		size.Y = rect.Y
	}

	imgui.Image(i.Texture().ID(), size)
}

// ImguiImageV renders the ReflectiveBoundTexture as an image in ImGui with additional options.
//
// Parameters:
//   - width: The width of the image. If set to -1, it will use the available content region width.
//   - height: The height of the image. If set to -1, it will use the available content region height.
//   - options: The options for rendering the image, including UV coordinates, tint color, and border color.
func (i *ReflectiveBoundTexture) ImguiImageV(width, height float32, options ImguiImageVOptionStruct) {
	size := imgui.Vec2{X: width, Y: height}

	if size.X == -1 {
		rect := imgui.ContentRegionAvail()
		size.X = rect.X
	}

	if size.Y == -1 {
		rect := imgui.ContentRegionAvail()
		size.Y = rect.Y
	}

	imgui.ImageWithBgV(i.Texture().ID(), size, options.Uv0, options.Uv1, options.BorderCol, options.TintCol)
}

// ImguiImageButtonV renders the ReflectiveBoundTexture as an image button in ImGui with additional options.
//
// Parameters:
//   - id: The ID of the image button.
//   - width: The width of the image button. If set to -1, it will use the available content region width.
//   - height: The height of the image button. If set to -1, it will use the available content region height.
//   - options: The options for rendering the image button, including UV coordinates, tint color, and border color.
func (i *ReflectiveBoundTexture) ImguiImageButtonV(id string, width, height float32, options ImguiImageVOptionStruct) {
	size := imgui.Vec2{X: width, Y: height}

	if size.X == -1 {
		rect := imgui.ContentRegionAvail()
		size.X = rect.X
	}

	if size.Y == -1 {
		rect := imgui.ContentRegionAvail()
		size.Y = rect.Y
	}

	imgui.ImageButtonV(id, i.Texture().ID(), size, options.Uv0, options.Uv1, options.TintCol, options.BorderCol)
}

// unbind releases the texture associated with the ReflectiveBoundTexture from the backend.
func (i *ReflectiveBoundTexture) unbind() {
	if i.tex != nil {
		Context.Backend().DeleteTexture(i.tex.ID())
		i.tex = nil
	}
}

// bind creates a new texture from the RGBA surface and assigns it to the ReflectiveBoundTexture.
// note it bypasses normal texture management up to cimgui-go to avoid double free from finalizers.
func (i *ReflectiveBoundTexture) bind() {
	img := ImageToRgba(i.Surface)
	i.tex = &Texture{
		&backend.Texture{
			ID:     backend.TextureManager(Context.backend).CreateTextureRgba(img, img.Bounds().Dx(), img.Bounds().Dy()),
			Width:  img.Bounds().Dx(),
			Height: img.Bounds().Dy(),
		},
	}
}

// GetSurfaceWidth returns the width of the RGBA surface.
func (i *ReflectiveBoundTexture) GetSurfaceWidth() int {
	return i.Surface.Bounds().Dx()
}

// GetSurfaceHeight returns the height of the RGBA surface.
func (i *ReflectiveBoundTexture) GetSurfaceHeight() int {
	return i.Surface.Bounds().Dy()
}

// GetSurfaceSize returns the size of the RGBA surface as an image.Point.
func (i *ReflectiveBoundTexture) GetSurfaceSize() image.Point {
	return i.Surface.Bounds().Size()
}

// Texture commits any pending changes to the RGBA surface and returns the associated texture.
func (i *ReflectiveBoundTexture) Texture() *Texture {
	i.commit()
	return i.tex
}

// TextureID commits any pending changes and returns the ImGui TextureID of the associated texture.
func (i *ReflectiveBoundTexture) TextureID() imgui.TextureID {
	return i.Texture().ID()
}

// GetRGBA returns the RGBA surface of the ReflectiveBoundTexture.
// If the commit parameter is true, it commits any pending changes before returning the surface.
//
// Parameters:
//   - commit: A boolean indicating whether to commit any pending changes.
//
// Returns:
//   - *image.RGBA: The RGBA surface of the ReflectiveBoundTexture.
func (i *ReflectiveBoundTexture) GetRGBA(commit bool) *image.RGBA {
	if commit {
		i.commit()
	}

	return i.Surface
}

// ForceRelease forces releasing resources against all finalizers,
// effectively losing the object but ensuring both RAM and VRAM
// are freed.
func (i *ReflectiveBoundTexture) ForceRelease() {
	i.unbind()
	i.Surface = nil

	var u uint32
	i.lastSum = u
}

// ForceCommit forces committing.
func (i *ReflectiveBoundTexture) ForceCommit() (*ReflectiveBoundTexture, bool) {
	return i.commit()
}
