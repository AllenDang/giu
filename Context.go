package giu

import (
	"fmt"
	"sync"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/glfwbackend"
	"github.com/AllenDang/cimgui-go/imgui"
	"gopkg.in/eapache/queue.v1"
)

// GenAutoID automatically generates widget's ID.
// It returns an unique value each time it is called.
func GenAutoID(id string) ID {
	idx, ok := Context.widgetIndex[id]
	if !ok {
		Context.widgetIndex[id] = 0
		return ID(id)
	}

	Context.widgetIndex[id]++

	return ID(fmt.Sprintf("%s##%d", id, idx))
}

// Context represents a giu context.
var Context *context

// Disposable should be implemented by all states stored in context.
// Dispose method is called when state is removed from context.
type Disposable interface {
	Dispose()
}

type genericDisposable[T any] interface {
	Disposable
	*T
}

type state struct {
	valid bool
	data  Disposable
}

type context struct {
	backend backend.Backend[glfwbackend.GLFWWindowFlags]

	isRunning bool

	widgetIndex map[string]int

	// Indicate whether current application is running
	isAlive bool

	// States will used by custom widget to store data
	state sync.Map

	InputHandler InputHandler
	FontAtlas    *FontAtlas

	textureLoadingQueue *queue.Queue

	cssStylesheet cssStylesheet

	m *sync.Mutex
}

// CreateContext creates a new giu context.
func CreateContext(b backend.Backend[glfwbackend.GLFWWindowFlags]) *context {
	result := context{
		cssStylesheet:       make(cssStylesheet),
		backend:             b,
		FontAtlas:           newFontAtlas(),
		textureLoadingQueue: queue.New(),
		m:                   &sync.Mutex{},
		widgetIndex:         make(map[string]int),
	}

	// Create font
	if len(result.FontAtlas.defaultFonts) == 0 {
		fonts := result.IO().Fonts()
		fonts.AddFontDefault()
		fontTextureImg, w, h, _ := fonts.GetTextureDataAsRGBA32()
		tex := result.backend.CreateTexture(fontTextureImg, int(w), int(h))
		fonts.SetTexID(tex)
		fonts.SetTexReady(true)
	} else {
		result.FontAtlas.shouldRebuildFontAtlas = true
	}

	return &result
}

// IO returns the imgui.IO object.
func (c *context) IO() *imgui.IO {
	return imgui.CurrentIO()
}

// invalidAllState should be called before rendering.
// it ensures all states are marked as invalid for that moment.
func (c *context) invalidAllState() {
	c.state.Range(func(k, v any) bool {
		if s, ok := v.(*state); ok {
			c.m.Lock()
			s.valid = false
			c.m.Unlock()
		}

		return true
	})
}

// cleanState removes all states that were not marked as valid during rendering.
// should be called after rendering.
func (c *context) cleanState() {
	c.state.Range(func(k, v any) bool {
		if s, ok := v.(*state); ok {
			c.m.Lock()
			valid := s.valid
			c.m.Unlock()

			if !valid {
				c.state.Delete(k)
				s.data.Dispose()
			}
		}

		return true
	})

	// Reset widgetIndex
	c.widgetIndex = make(map[string]int)
}

// Backend returns the imgui.backend used by the context.
func (c *context) Backend() backend.Backend[glfwbackend.GLFWWindowFlags] {
	return c.backend
}

// SetState is a generic version of Context.SetState.
func SetState[T any, PT genericDisposable[T]](c *context, id string, data PT) {
	c.state.Store(id, &state{valid: true, data: data})
}

// SetState stores data in context by id.
func (c *context) SetState(id string, data Disposable) {
	c.state.Store(id, &state{valid: true, data: data})
}

// GetState is a generic version of Context.GetState.
func GetState[T any, PT genericDisposable[T]](c *context, id string) PT {
	if s, ok := c.load(id); ok {
		c.m.Lock()
		s.valid = true
		c.m.Unlock()

		data, isOk := s.data.(PT)
		Assert(isOk, "Context", "GetState", "got state of unexpected type: expected %T, instead found %T", new(T), s.data)

		return data
	}

	return nil
}

// GetState returns previously stored state by id.
func (c *context) GetState(id string) any {
	if s, ok := c.load(id); ok {
		c.m.Lock()
		s.valid = true
		c.m.Unlock()

		return s.data
	}

	return nil
}

func (c *context) load(id any) (*state, bool) {
	if v, ok := c.state.Load(id); ok {
		if s, ok := v.(*state); ok {
			return s, true
		}
	}

	return nil, false
}
