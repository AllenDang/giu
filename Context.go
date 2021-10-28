package giu

import (
	"sync"

	"github.com/AllenDang/imgui-go"
)

// Context represents a giu context.
var Context context

// Disposable should be implemented by all states stored in context.
// Dispose method is called when state is removed from context.
type Disposable interface {
	Dispose()
}

type state struct {
	valid bool
	data  Disposable
}

type context struct {
	renderer imgui.Renderer
	platform imgui.Platform

	widgetIndexCounter int

	// Indicate whether current application is running
	isAlive bool

	// States will used by custom widget to store data
	state sync.Map

	InputHandler InputHandler
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
	c.state.Range(func(k, v interface{}) bool {
		if s, ok := v.(*state); ok {
			s.valid = false
		}
		return true
	})
}

func (c *context) cleanState() {
	c.state.Range(func(k, v interface{}) bool {
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

func (c *context) SetState(id string, data Disposable) {
	c.state.Store(id, &state{valid: true, data: data})
}

func (c *context) GetState(id string) interface{} {
	if v, ok := c.state.Load(id); ok {
		if s, ok := v.(*state); ok {
			s.valid = true
			return s.data
		}
	}

	return nil
}

// Get widget index for current layout.
func (c *context) GetWidgetIndex() int {
	i := c.widgetIndexCounter
	c.widgetIndexCounter++
	return i
}
