package giu

import (
	"image"
	"runtime"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/imgui"
)

// Texture represents imgui.TextureID.
// It is base unit of images in imgui.
type Texture struct {
	tex *backend.Texture
}

type textureLoadRequest struct {
	img image.Image
	cb  func(*Texture)
}

type textureFreeRequest struct {
	tex *Texture
}

// EnqueueNewTextureFromRgba adds loading texture request to loading queue
// it allows us to run this method in main loop
// NOTE: remember to call it after NewMasterWindow!
func EnqueueNewTextureFromRgba(rgba image.Image, loadCb func(t *Texture)) {
	Assert((Context.textureLoadingQueue != nil), "", "EnqueueNewTextureFromRgba", "you need to call EnqueueNewTextureFromRgba after giu.NewMasterWindow call!")
	Context.textureLoadingQueue.Add(textureLoadRequest{rgba, loadCb})
}

// NewTextureFromRgba creates a new texture from image.Image and, when it is done, calls loadCallback(loadedTexture).
func NewTextureFromRgba(rgba image.Image, loadCallback func(*Texture)) {
	tex := backend.NewTextureFromRgba(ImageToRgba(rgba))
	giuTex := &Texture{
		tex,
	}

	runtime.SetFinalizer(giuTex, func(tex *Texture) {
		Context.textureFreeingQueue.Add(textureFreeRequest{tex})
	})

	loadCallback(giuTex)
}

// ToTexture converts backend.Texture to Texture.
func ToTexture(texture *backend.Texture) *Texture {
	return &Texture{tex: texture}
}

// ID returns imgui.TextureID of the texture.
func (t *Texture) ID() imgui.TextureID {
	if t.tex != nil {
		return t.tex.ID
	}

	return 0
}
