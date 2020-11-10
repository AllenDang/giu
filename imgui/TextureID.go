package imgui

// #include "imguiWrapperTypes.h"
import "C"

// TextureID is a user data to identify a texture.
//
// TextureID is a uintptr used to pass renderer-agnostic texture references around until it hits your render function.
// imgui knows nothing about what those bits represent, it just passes them around. It is up to you to decide what you want the value to carry!
//
// It could be an identifier to your OpenGL texture (cast as uint32), a key to your custom engine material, etc.
// At the end of the chain, your renderer takes this value to cast it back into whatever it needs to select a current texture to render.
//
// To display a custom image/texture within an imgui window, you may use functions such as imgui.Image().
// imgui will generate the geometry and draw calls using the TextureID that you passed and which your renderer can use.
// It is your responsibility to get textures uploaded to your GPU.
//
// Note: Internally, the value is based on a pointer type, so its size is dependent on your architecture.
// For the most part, this will be 64bits on current systems (in 2018).
// Beware: This value must never be a Go pointer, because the value escapes the runtime!
type TextureID uintptr

func (id TextureID) handle() C.IggTextureID {
	return C.IggTextureID(id)
}
