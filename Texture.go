package giu

import (
	"errors"
	"image"
	"runtime"

	"github.com/AllenDang/giu/imgui"
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
	result := CallVal(func() interface{} {
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

func (t *Texture) release() {
	Update()
	Call(func() {
		Context.renderer.ReleaseImage(t.id)
	})
}
