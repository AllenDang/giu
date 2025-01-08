package giu

import (
	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/glfwbackend"
)

// sometimes we need to tell what we mean ;-)
// good example is:
// FlagsNotResizable in giu,
// but cimgui-go has only FlagsResizable. So SetFlags(Resizable, 0)
type flagValue[T ~int] struct {
	flag  T
	value int
}

// GIUBackend is an abstraction layer between cimgui-go's Backends.
type GIUBackend backend.Backend[MasterWindowFlags]

var _ GIUBackend = &GLFWBackend{}

// GLFWBackend is an implementation of glfbackend.GLFWBackend cimgui-go backend with respect to
// giu's MasterWIndowFlags.
type GLFWBackend struct {
	*glfwbackend.GLFWBackend
}

func NewGLFWBackend() *GLFWBackend {
	return &GLFWBackend{
		GLFWBackend: glfwbackend.NewGLFWBackend(),
	}
}

func (b *GLFWBackend) SetInputMode(mode MasterWindowFlags, value MasterWindowFlags) {
	flag := b.parseFlag(mode)
	b.GLFWBackend.SetInputMode(flag.flag, glfwbackend.GLFWWindowFlags(flag.value))
}

func (b *GLFWBackend) SetSwapInterval(interval MasterWindowFlags) error {
	intervalV := b.parseFlag(interval).flag
	b.GLFWBackend.SetSwapInterval(intervalV)
	return nil
}

func (b *GLFWBackend) SetWindowFlags(flags MasterWindowFlags, _ int) {
	flag := b.parseFlag(flags)
	b.GLFWBackend.SetWindowFlags(flag.flag, flag.value)
}

func (b *GLFWBackend) parseFlag(m MasterWindowFlags) flagValue[glfwbackend.GLFWWindowFlags] {
	data := map[MasterWindowFlags]flagValue[glfwbackend.GLFWWindowFlags]{
		MasterWindowFlagsNotResizable: {glfwbackend.GLFWWindowFlagsResizable, 0},
		MasterWindowFlagsMaximized:    {glfwbackend.GLFWWindowFlagsMaximized, 1},
		MasterWindowFlagsFloating:     {glfwbackend.GLFWWindowFlagsFloating, 1},
		MasterWindowFlagsFrameless:    {glfwbackend.GLFWWindowFlagsDecorated, 0},
		MasterWindowFlagsTransparent:  {glfwbackend.GLFWWindowFlagsTransparent, 1},
		MasterWindowFlagsHidden:       {glfwbackend.GLFWWindowFlagsVisible, 0},
	}

	d, ok := data[m]
	Assert(ok, "GLFWBackend", "parseFlag", "Unknown MasterWindowFlags")

	return d
}
