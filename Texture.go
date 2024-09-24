package giu

import (
	"image"

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
	loadCallback(&Texture{
		tex,
	})
}

// ToTexture converts backend.Texture to Texture.
func ToTexture(texture *backend.Texture) *Texture {
	return &Texture{tex: texture}
}

func (t *Texture) ID() imgui.TextureID {
	if t.tex != nil {
		return t.tex.ID
	}

	return imgui.TextureID{}
}
