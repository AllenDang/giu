package giu

import (
	"fmt"
	"sync"

	"github.com/AllenDang/cimgui-go/imgui"
	"gopkg.in/eapache/queue.v1"
)

// GenAutoID automatically generates widget's ID.
// It returns an unique value each time it is called.
func GenAutoID(id string) ID {
	idx := int(0)
	idxAny, ok := Context.widgetIndex.Load(id)

	if ok {
		idx, ok = idxAny.(int)
		Assert(ok, "Context", "GenAutoID", "unexpected type of widgetIndex value: expected int, instead found %T", idxAny)

		idx++
	}

	Context.widgetIndex.Store(id, idx)

	return ID(fmt.Sprintf("%s##%d", id, idx))
}

// Context represents a giu context.
var Context *GIUContext

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

// GIUContext represents a giu context. (Current context is giu.Context.
//
//nolint:revive // I WANT TO CALL THIS GIUContext!
type GIUContext struct {
	backend GIUBackend

	isRunning bool

	widgetIndex sync.Map

	// Indicate whether current application is running
	isAlive bool

	// when dirty is true, flushStates must be called before any GetState use
	// when it is false, calling flushStates is noop
	dirty bool

	// States will used by custom widget to store data
	state sync.Map

	InputHandler InputHandler
	FontAtlas    *FontAtlas
	Translator   Translator

	textureLoadingQueue *queue.Queue
	textureFreeingQueue *queue.Queue

	cssStylesheet *CSSStylesheet

	m *sync.Mutex
}

// CreateContext creates a new giu context.
func CreateContext(b GIUBackend) *GIUContext {
	result := GIUContext{
		cssStylesheet:       CSS(),
		backend:             b,
		FontAtlas:           newFontAtlas(),
		textureLoadingQueue: queue.New(),
		textureFreeingQueue: queue.New(),
		m:                   &sync.Mutex{},
		Translator:          &EmptyTranslator{},
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
func (c *GIUContext) IO() *imgui.IO {
	return imgui.CurrentIO()
}

// SetDirty permits MasterWindow defering setting dirty states after it's render().
func (c *GIUContext) SetDirty() {
	c.dirty = true
}

// PrepareString prepares string to be displayed by imgui.
// It does the following:
// - adds a string to the FontAtlas
// - translates the string with the Translator set in the context
// Not all widgets will use this. Text with user-defined input (e.g. InputText will still use FontAtlas.RegisterString).
func (c *GIUContext) PrepareString(str string) string {
	str = c.Translator.Translate(str)
	return c.FontAtlas.RegisterString(str)
}

// SetCSSStylesheet sets the "main" stylesheet for the context.
// Setting it gives you 2 benefits:
// - MasterWindow uses MainTag of this for its default theme.
func (c *GIUContext) SetCSSStylesheet(css *CSSStylesheet) {
	c.cssStylesheet = css
}

// cleanStates removes all states that were not marked as valid during rendering,
// then reset said flag before new usage
// should always be called before first Get/Set state use in renderloop
// since afterRender() and beforeRender() are not waranted to run (see glfw_window_refresh_callback)
// we call it at the very start of our render()
// but just in case something happened, we also use the "dirty" flag to enforce (or avoid) flushing
// on critical path.
func (c *GIUContext) cleanStates() {
	if !c.dirty {
		return
	}

	c.state.Range(func(k, v any) bool {
		if s, ok := v.(*state); ok {
			c.m.Lock()
			valid := s.valid
			c.m.Unlock()

			if valid {
				s.valid = false
			} else {
				c.state.Delete(k)
				s.data.Dispose()
			}
		}

		return true
	})

	c.widgetIndex.Clear()
	c.dirty = false
}

// Backend returns the imgui.backend used by the context.
func (c *GIUContext) Backend() GIUBackend {
	return c.backend
}

// SetState is a generic version of Context.SetState.
func SetState[T any, PT genericDisposable[T]](c *GIUContext, id ID, data PT) {
	c.cleanStates()
	c.state.Store(id, &state{valid: true, data: data})
}

// SetState stores data in context by id.
func (c *GIUContext) SetState(id ID, data Disposable) {
	c.cleanStates()
	c.state.Store(id, &state{valid: true, data: data})
}

// GetState is a generic version of Context.GetState.
func GetState[T any, PT genericDisposable[T]](c *GIUContext, id ID) PT {
	c.cleanStates()

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
func (c *GIUContext) GetState(id ID) any {
	c.cleanStates()

	if s, ok := c.load(id); ok {
		c.m.Lock()
		s.valid = true
		c.m.Unlock()

		return s.data
	}

	return nil
}

func (c *GIUContext) load(id any) (*state, bool) {
	if v, ok := c.state.Load(id); ok {
		if s, ok := v.(*state); ok {
			return s, true
		}
	}

	return nil, false
}
