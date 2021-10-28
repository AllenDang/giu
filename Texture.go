package giu

import (
	"fmt"
	"image"
	"runtime"

	"github.com/AllenDang/imgui-go"
	"github.com/faiface/mainthread"
)

// Texture represents imgui.TextureID.
// It is base unit of images in imgui.
type Texture struct {
	id imgui.TextureID
}

type loadImageResult struct {
	id  imgui.TextureID
	err error
}

// NewTextureFromRgba creates a new texture from image.Image and, when it is done, calls loadCallback(loadedTexture).
func NewTextureFromRgba(rgba image.Image, loadCallback func(*Texture)) {
	go func() {
		Update()
		result := mainthread.CallVal(func() interface{} {
			texID, err := Context.renderer.LoadImage(ImageToRgba(rgba))
			return &loadImageResult{id: texID, err: err}
		})

		tid, ok := result.(*loadImageResult)
		switch {
		case !ok:
			panic("giu: NewTextureFromRgba: unexpected error occurred")
		case tid.err != nil:
			panic(fmt.Sprintf("giu: NewTextureFromRgba: error loading texture: %v", tid.err))
		}

		texture := Texture{id: tid.id}

		// Set finalizer
		runtime.SetFinalizer(&texture, (*Texture).release)

		// execute callback
		loadCallback(&texture)
	}()
}

// ToTexture converts imgui.TextureID to Texture.
func ToTexture(textureID imgui.TextureID) *Texture {
	return &Texture{id: textureID}
}

func (t *Texture) release() {
	Update()
	mainthread.Call(func() {
		Context.renderer.ReleaseImage(t.id)
	})
}
