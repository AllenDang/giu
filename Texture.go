package giu

import (
	"errors"
	"github.com/faiface/mainthread"
	"image"
	"runtime"

	"github.com/AllenDang/imgui-go"
)

type Texture struct {
	id imgui.TextureID
}

type loadImageResult struct {
	id  imgui.TextureID
	err error
}

// Create new texture from rgba.
// Note: this function has to be invokded in a go routine.
// If call this in mainthread will result in stuck.
func NewTextureFromRgba(rgba *image.RGBA) (*Texture, error) {
	Update()
	result := mainthread.CallVal(func() interface{} {
		texId, err := Context.renderer.LoadImage(rgba)
		return &loadImageResult{id: texId, err: err}
	})

	if tid, ok := result.(*loadImageResult); ok {
		texture := Texture{id: tid.id}
		if tid.err == nil {
			// Set finalizer
			runtime.SetFinalizer(&texture, (*Texture).release)
		}
		return &texture, tid.err
	}
	return nil, errors.New("Unknown error occurred")
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
