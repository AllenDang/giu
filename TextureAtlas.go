package giu

import (
	"errors"
	"image"
	"log"
	"sync"
)

var textures = &sync.Map{}

type textureField struct {
	texture *Texture
	isValid bool
	forever bool
}

type TextureLoader struct {
	img     *image.RGBA
	forever bool

	tex *Texture
}

func AddTexture(img *image.RGBA) *TextureLoader {
	return &TextureLoader{
		img: img,
	}
}

// Tex could be used to get texture pointer.
// it will be filled with the texture when Build called
// NOTE: please remember about intializing tex argument in your code
// by setting it to &giu.Texture{}, because else it will not be set.
func (t *TextureLoader) Tex(tex *Texture) *TextureLoader {
	if tex != nil {
		t.tex = tex
	}
	return t
}

// Forever if set, the texture will not be lost.
func (t *TextureLoader) Forever() *TextureLoader {
	t.forever = true
	return t
}

func (t *TextureLoader) Build() {
	// if already loaded, valid texture
	if tex, ok := textures.Load(t.img); ok && tex != nil {
		texture := tex.(*textureField)
		if texture != nil {
			*t.tex = *texture.texture
			texture.isValid = true
			return
		}
	}

	// if not, store it to prevent app from invoking load function
	// more than 1 time
	textures.Store(t.img, nil)
	go func() {
		if t.img == nil {
			return
		}

		texture, err := NewTextureFromRgba(t.img)

		switch {
		case err != nil:
			log.Print(err)
			textures.Delete(t.img)
		case texture == nil:
			log.Print(errors.New("giu: texture atlas: loaded texture is nil"))
			textures.Delete(t.img)
		}

		textures.Store(t.img, &textureField{texture, true, t.forever})
	}()
}

func invalidTextures() {
	textures.Range(func(_, value interface{}) bool {
		if value == nil {
			return true
		}

		tex := value.(*textureField)
		tex.isValid = false
		return true
	})
}

func cleanTextures() {
	textures.Range(func(k, v interface{}) bool {
		if v == nil {
			return true
		}

		if tf := v.(*textureField); !tf.isValid {
			textures.Delete(k)
		}

		return true
	})
}
