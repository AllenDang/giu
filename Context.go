package giu

import (
	imguiLocal "github.com/ianling/giu/imgui"
	"github.com/ianling/imgui-go"
)

var (
	Context context
)

type Disposable interface {
	Dispose()
}

type state struct {
	valid bool
	data  Disposable
}

type context struct {
	renderer imguiLocal.Renderer
	platform imguiLocal.Platform

	// Indicate whether current application is running
	isAlive bool

	// States will used by custom widget to store data
	state map[string]*state
}

func (c context) GetRenderer() imguiLocal.Renderer {
	return c.renderer
}

func (c context) GetPlatform() imguiLocal.Platform {
	return c.platform
}

func (c context) IO() imgui.IO {
	return imgui.CurrentIO()
}

func (c context) invalidAllState() {
	for _, s := range c.state {
		s.valid = false
	}
}

func (c context) cleanState() {
	for id, s := range c.state {
		if !s.valid {
			delete(c.state, id)
			s.data.Dispose()
		}
	}
}

func (c context) SetState(id string, data Disposable) {
	c.state[id] = &state{valid: true, data: data}
}

func (c context) GetState(id string) interface{} {
	if s, ok := c.state[id]; ok {
		s.valid = true
		return s.data
	}

	return nil
}
