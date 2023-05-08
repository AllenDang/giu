package giu

import (
	"fmt"
	"sync"

	"github.com/AllenDang/imgui-go"
	"gopkg.in/eapache/queue.v1"
)

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
	// TODO: should be handled by mainthread tbh
	// see https://github.com/faiface/mainthread/pull/4
	isRunning bool

	renderer imgui.Renderer
	platform imgui.Platform

	widgetIndexCounter int

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

func CreateContext(p imgui.Platform, r imgui.Renderer) *context {
	result := context{
		platform:      p,
		renderer:      r,
		cssStylesheet: make(cssStylesheet),
		m:             &sync.Mutex{},
	}

	result.FontAtlas = newFontAtlas()

	// Create font
	if len(result.FontAtlas.defaultFonts) == 0 {
		io := result.IO()
		io.Fonts().AddFontDefault()
		fontAtlas := io.Fonts().TextureDataRGBA32()
		r.SetFontTexture(fontAtlas)
	} else {
		result.FontAtlas.shouldRebuildFontAtlas = true
	}

	return &result
}

func (c *context) GetRenderer() imgui.Renderer {
	return c.renderer
}

func (c *context) GetPlatform() imgui.Platform {
	return c.platform
}

func (c *context) IO() imgui.IO {
	return imgui.CurrentIO()
}

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

	// Reset widgetIndexCounter
	c.widgetIndexCounter = 0
}

func SetState[T any, PT genericDisposable[T]](c *context, id string, data PT) {
	c.state.Store(id, &state{valid: true, data: data})
}

func (c *context) SetState(id string, data Disposable) {
	c.state.Store(id, &state{valid: true, data: data})
}

func GetState[T any, PT genericDisposable[T]](c *context, id string) PT {
	if s, ok := c.load(id); ok {
		c.m.Lock()
		s.valid = true
		c.m.Unlock()

		data, isOk := s.data.(PT)
		Assert(isOk, "Context", "GetState", fmt.Sprintf("got state of unexpected type: expected %T, instead found %T", new(T), s.data))

		return data
	}

	return nil
}

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

// Get widget index for current layout.
func (c *context) GetWidgetIndex() int {
	i := c.widgetIndexCounter
	c.widgetIndexCounter++

	return i
}
