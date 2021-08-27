package giu

type mouseEvent struct {
	mouseButton MouseButton
	callback    func()
	cond        func(MouseButton) bool
}

type keyEvent struct {
	key      Key
	callback func()
	cond     func(Key) bool
}

// EventHandler is a universal event handler for giu widgets.
// put giu.Event()... after any widget to handle any event
type EventHandler struct {
	hover       func()
	mouseEvents []mouseEvent
	keyEvents   []keyEvent
}

func Event() *EventHandler {
	return &EventHandler{
		mouseEvents: make([]mouseEvent, 0),
		keyEvents:   make([]keyEvent, 0),
	}
}

func (eh *EventHandler) OnHover(onHover func()) *EventHandler {
	eh.hover = onHover
	return eh
}

// Key events

func (eh *EventHandler) OnKeyDown(key Key, cb func()) *EventHandler {
	eh.keyEvents = append(eh.keyEvents, keyEvent{key, cb, IsKeyDown})
	return eh
}

func (eh *EventHandler) OnKeyPressed(key Key, cb func()) *EventHandler {
	eh.keyEvents = append(eh.keyEvents, keyEvent{key, cb, IsKeyPressed})
	return eh
}

func (eh *EventHandler) OnKeyReleased(key Key, cb func()) *EventHandler {
	eh.keyEvents = append(eh.keyEvents, keyEvent{key, cb, IsKeyReleased})
	return eh
}

// Mouse events

func (eh *EventHandler) OnClick(mouseButton MouseButton, callback func()) *EventHandler {
	eh.mouseEvents = append(eh.mouseEvents, mouseEvent{mouseButton, callback, IsMouseClicked})
	return eh
}

func (eh *EventHandler) OnDClick(mouseButton MouseButton, callback func()) *EventHandler {
	eh.mouseEvents = append(eh.mouseEvents, mouseEvent{mouseButton, callback, IsMouseDoubleClicked})
	return eh
}

func (eh *EventHandler) OnMouseDown(mouseButton MouseButton, callback func()) *EventHandler {
	eh.mouseEvents = append(eh.mouseEvents, mouseEvent{mouseButton, callback, IsMouseDown})
	return eh
}

func (eh *EventHandler) OnMouseReleased(mouseButton MouseButton, callback func()) *EventHandler {
	eh.mouseEvents = append(eh.mouseEvents, mouseEvent{mouseButton, callback, IsMouseReleased})
	return eh
}

func (eh *EventHandler) Build() {
	if !IsItemHovered() {
		return
	}

	if eh.hover != nil {
		eh.hover()
	}

	if len(eh.keyEvents) > 0 {
		for _, event := range eh.keyEvents {
			if event.callback != nil && event.cond(event.key) {
				event.callback()
			}
		}
	}

	if len(eh.mouseEvents) > 0 {
		for _, event := range eh.mouseEvents {
			if event.callback != nil && event.cond(event.mouseButton) {
				event.callback()
			}
		}
	}
}
