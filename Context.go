package giu

import (
	"github.com/AllenDang/giu/imgui"
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
	renderer imgui.Renderer
	platform imgui.Platform

	// Indicate whether current application is running
	isAlive bool

	// States will used by custom widget to store data
	state map[string]*state
}

func (c context) GetRenderer() imgui.Renderer {
	return c.renderer
}

func (c context) GetPlatform() imgui.Platform {
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
	var invalidIds []string
	for id, s := range c.state {
		if !s.valid {
			invalidIds = append(invalidIds, id)
			s.data.Dispose()
		}
	}

	for _, id := range invalidIds {
		delete(c.state, id)
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
