package giu

import (
	"errors"
	"fmt"
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

func NewTextureFromRgba(rgba *image.RGBA) (*Texture, error) {
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
	Call(func() {
		Context.renderer.ReleaseImage(t.id)
		fmt.Printf("Texture %v is released", t.id)
	})
}
