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

type textureLoadRequest struct {
	img image.Image
	cb  func(*Texture)
}

type loadImageResult struct {
	id  imgui.TextureID
	err error
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
	Assert(Context.isRunning, "", "NewTextureFromRgba", "cannot load texture befor (*MasterWindow).Run call!")
	loadTexture(rgba, loadCallback)
}

func loadTexture(rgba image.Image, loadCallback func(*Texture)) {
	go func() {
		Update()
		result := mainthread.CallVal(func() any {
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
