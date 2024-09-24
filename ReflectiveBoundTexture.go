package giu

import (
	"fmt"
	"hash/crc32"
	"image"
	"image/color"
	"sync"

	imgui "github.com/AllenDang/cimgui-go"
)

func defaultSurface() *image.RGBA {
	surface, _ := UniformLoader(REFLECTIVE_SURFACE_DEFAULT_WIDTH, REFLECTIVE_SURFACE_DEFAULT_HEIGHT, REFLECTIVE_SURFACE_DEFAULT_COLOR).ServeRGBA()
	return surface
}

const (
	REFLECTIVE_SURFACE_DEFAULT_WIDTH  = 128
	REFLECTIVE_SURFACE_DEFAULT_HEIGHT = 128
)

var (
	REFLECTIVE_SURFACE_DEFAULT_COLOR = color.RGBA{255, 255, 255, 255}
)

type ReflectiveBoundTexture struct {
	Surface *image.RGBA
	tex     *Texture
	lastSum uint32
	mu      sync.Mutex
}

/* Return a waranted:
 * Initialized
 * With proper resources bindings against gpu (free old, bound new)
 * Up to date Texture
 */
func (i *ReflectiveBoundTexture) commit() (*ReflectiveBoundTexture, bool) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.Surface == nil {
		i.Surface = defaultSurface()
	}

	var has_changed bool
	if sum := crc32.ChecksumIEEE(i.Surface.Pix); sum != i.lastSum {
		has_changed = true
		i.unbind()
		i.bind()
		i.lastSum = sum
	}
	return i, has_changed
}

func (i *ReflectiveBoundTexture) SetSurfaceFromRGBA(img *image.RGBA, commit bool) error {
	if img != nil {
		i.Surface = img
	} else {
		return fmt.Errorf("RGBA Result is nil")
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
	Uv0        imgui.Vec2
	Uv1        imgui.Vec2
	Tint_col   imgui.Vec4
	Border_col imgui.Vec4
}

func (i *ReflectiveBoundTexture) GetImGuiImageVDefaultOptionsStruct() ImguiImageVOptionStruct {
	return ImguiImageVOptionStruct{
		Uv0:        imgui.Vec2{0, 0},
		Uv1:        imgui.Vec2{1, 1},
		Tint_col:   imgui.Vec4{1, 1, 1, 1},
		Border_col: imgui.Vec4{0, 0, 0, 0},
	}
}
func (i *ReflectiveBoundTexture) ImguiImage(width float32, height float32, options ImguiImageVOptionStruct) {
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

func (i *ReflectiveBoundTexture) ImguiImageV(width float32, height float32, options ImguiImageVOptionStruct) {

	size := imgui.Vec2{X: width, Y: height}

	if size.X == -1 {
		rect := imgui.ContentRegionAvail()
		size.X = rect.X
	}

	if size.Y == -1 {
		rect := imgui.ContentRegionAvail()
		size.Y = rect.Y
	}

	imgui.ImageV(i.Texture().ID(), size, options.Uv0, options.Uv1, options.Tint_col, options.Border_col)
}

func (i *ReflectiveBoundTexture) ImguiImageButtonV(id string, width float32, height float32, options ImguiImageVOptionStruct) {
	size := imgui.Vec2{X: width, Y: height}

	if size.X == -1 {
		rect := imgui.ContentRegionAvail()
		size.X = rect.X
	}

	if size.Y == -1 {
		rect := imgui.ContentRegionAvail()
		size.Y = rect.Y
	}

	imgui.ImageButtonV(id, i.Texture().ID(), size, options.Uv0, options.Uv1, options.Tint_col, options.Border_col)
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

// Force releasing against all finalizers,
// effectively losing the object but ensuring both RAM and VRAM
// are freed.
func (i *ReflectiveBoundTexture) ForceRelease() {
	i.unbind()
	i.Surface = nil
	var u uint32
	i.lastSum = u
	i = nil
}

// Force Commiting
func (i *ReflectiveBoundTexture) ForceCommit() (*ReflectiveBoundTexture, bool) {
	return i.commit()
}
