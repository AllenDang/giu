package giu

import (
	"sync"

	imgui "github.com/AllenDang/cimgui-go"
	"gopkg.in/eapache/queue.v1"
)

// Context represents a giu context.
var Context context

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
	widgetIndexCounter int

	// Indicate whether current application is running
	isAlive bool

	// States will used by custom widget to store data
	state sync.Map

	FontAtlas FontAtlas

	textureLoadingQueue *queue.Queue

	window imgui.GLFWwindow
}

func CreateContext(window imgui.GLFWwindow) context {
	result := context{
		window: window,
	}

	result.FontAtlas = newFontAtlas()

	// Create font
	if len(result.FontAtlas.defaultFonts) == 0 {
		io := result.IO()
		fonts := io.GetFonts()
		fonts.AddFontDefault(0)
		fontAtlas, width, height, _ := fonts.GetTextureDataAsRGBA32()
		texId := imgui.CreateTexture(fontAtlas, int(width), int(height))
		fonts.SetTexID(texId)
	} else {
		result.FontAtlas.shouldRebuildFontAtlas = true
		result.FontAtlas.rebuildFontAtlas()
	}

	return result
}

func (c *context) IO() imgui.ImGuiIO {
	return imgui.GetIO()
}

func (c *context) invalidAllState() {
	c.state.Range(func(k, v any) bool {
		if s, ok := v.(*state); ok {
			s.valid = false
		}
		return true
	})
}

func (c *context) cleanState() {
	c.state.Range(func(k, v any) bool {
		if s, ok := v.(*state); ok {
			if !s.valid {
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

func GetState[T any, PT genericDisposable[T]](c context, id string) PT {
	if s, ok := c.load(id); ok {
		s.valid = true
		data, _ := s.data.(PT)
		return data

	}
	return nil
}

func (c *context) GetState(id string) any {
	if s, ok := c.load(id); ok {
		s.valid = true
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
