package giu

import (
    platform2 "github.com/AllenDang/giu/platform"
    "sync"

    "github.com/AllenDang/giu/renderer"
	"github.com/AllenDang/imgui-go"
)

var Context context

type Disposable interface {
	Dispose()
}

type state struct {
	valid bool
	data  Disposable
}

type context struct {
	renderer renderer.Renderer
	platform platform2.Platform

	// Indicate whether current application is running
	isAlive bool

	// States will used by custom widget to store data
	state sync.Map
}

func (c *context) GetRenderer() renderer.Renderer {
	return c.renderer
}

func (c *context) GetPlatform() platform2.Platform {
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
