package giu

import (
	"errors"
	"fmt"
	"hash/crc32"
	"image"
	"image/color"
	"sync"

	"github.com/AllenDang/cimgui-go/imgui"
)

var errNilRGBA = errors.New("surface RGBA Result is nil")

func defaultSurface() *image.RGBA {
	surface, _ := UniformLoader(ReflectiveSurfaceDefaultWidth, ReflectiveSurfaceDefaultHeight, ReflectiveSurfaceDefaultColor).ServeRGBA()
	return surface
}

const (
	ReflectiveSurfaceDefaultWidth  = 128
	ReflectiveSurfaceDefaultHeight = 128
)

var (
	ReflectiveSurfaceDefaultColor = color.RGBA{255, 255, 255, 255}
)

type ReflectiveBoundTexture struct {
	Surface *image.RGBA
	tex     *Texture
	lastSum uint32
	mu      sync.Mutex
}

/* Return a waranted:
 * - Initialized
 * - With proper resources bindings against gpu (free old, bound new)
 * - Up to date Texture.
 */
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

func (i *ReflectiveBoundTexture) SetSurfaceFromRGBA(img *image.RGBA, commit bool) error {
	if img != nil {
		i.Surface = img
	} else {
		return fmt.Errorf("%w", errNilRGBA)
	}

	if commit {
		i.commit()
	}

	return nil
}

func (i *ReflectiveBoundTexture) ToImageWidget() *ImageWidget {
	return Image(i.Texture())
}

type ImguiImageVOptionStruct struct {
	Uv0       imgui.Vec2
	Uv1       imgui.Vec2
	TintCol   imgui.Vec4
	BorderCol imgui.Vec4
}

func (i *ReflectiveBoundTexture) GetImGuiImageVDefaultOptionsStruct() ImguiImageVOptionStruct {
	return ImguiImageVOptionStruct{
		Uv0:       imgui.Vec2{X: 0, Y: 0},
		Uv1:       imgui.Vec2{X: 1, Y: 1},
		TintCol:   imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1},
		BorderCol: imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0},
	}
}

func (i *ReflectiveBoundTexture) ImguiImage(width, height float32, options ImguiImageVOptionStruct) {
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

	imgui.ImageV(i.Texture().ID(), size, options.Uv0, options.Uv1, options.TintCol, options.BorderCol)
}

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

func (i *ReflectiveBoundTexture) unbind() {
	if i.tex != nil {
		Context.Backend().DeleteTexture(i.tex.ID())
		i.tex = nil
	}
}
func (i *ReflectiveBoundTexture) bind() {
	NewTextureFromRgba(i.Surface, func(tex *Texture) {
		i.tex = tex
	})
}

func (i *ReflectiveBoundTexture) GetSurfaceWidth() int {
	return i.Surface.Bounds().Dx()
}

func (i *ReflectiveBoundTexture) GetSurfaceHeight() int {
	return i.Surface.Bounds().Dy()
}

func (i *ReflectiveBoundTexture) GetSurfaceSize() image.Point {
	return i.Surface.Bounds().Size()
}

func (i *ReflectiveBoundTexture) Texture() *Texture {
	i.commit()
	return i.tex
}

func (i *ReflectiveBoundTexture) ID() imgui.TextureID {
	return i.Texture().ID()
}

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
