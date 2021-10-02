package giu

var _ Disposable = &eventHandlerState{}

type eventHandlerState struct {
	isActive bool
}

// Dispose implements Disposable interface.
func (s *eventHandlerState) Dispose() {
	// noop
}

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

var _ Widget = &EventHandler{}

// EventHandler is a universal event handler for giu widgets.
// put giu.Event()... after any widget to handle any event.
type EventHandler struct {
	hover       func()
	mouseEvents []mouseEvent
	keyEvents   []keyEvent
	onActivate,
	onDeactivate,
	onActive func()
}

// Event adds a new event to widget above.
func Event() *EventHandler {
	return &EventHandler{
		mouseEvents: make([]mouseEvent, 0),
		keyEvents:   make([]keyEvent, 0),
	}
}

// OnHover sets callback when item gets hovered.
func (eh *EventHandler) OnHover(onHover func()) *EventHandler {
	eh.hover = onHover
	return eh
}

// OnActive sets a callback when ite IS ACTIVE (not activated).
func (eh *EventHandler) OnActive(cb func()) *EventHandler {
	eh.onActive = cb
	return eh
}

// OnActivate sets callback when item gets activated.
func (eh *EventHandler) OnActivate(cb func()) *EventHandler {
	eh.onActivate = cb
	return eh
}

// OnDeactivate sets callback when item gets deactivated.
func (eh *EventHandler) OnDeactivate(cb func()) *EventHandler {
	eh.onDeactivate = cb
	return eh
}

// Key events

// OnKeyDown sets callback when key `key` is down.
func (eh *EventHandler) OnKeyDown(key Key, cb func()) *EventHandler {
	eh.keyEvents = append(eh.keyEvents, keyEvent{key, cb, IsKeyDown})
	return eh
}

// OnKeyPressed sets callback when key `key` is pressed.
func (eh *EventHandler) OnKeyPressed(key Key, cb func()) *EventHandler {
	eh.keyEvents = append(eh.keyEvents, keyEvent{key, cb, IsKeyPressed})
	return eh
}

// OnKeyReleased sets callback when key `key` is released.
func (eh *EventHandler) OnKeyReleased(key Key, cb func()) *EventHandler {
	eh.keyEvents = append(eh.keyEvents, keyEvent{key, cb, IsKeyReleased})
	return eh
}

// Mouse events

// OnClick sets callback when mouse button `mouseButton` is clicked.
func (eh *EventHandler) OnClick(mouseButton MouseButton, callback func()) *EventHandler {
	eh.mouseEvents = append(eh.mouseEvents, mouseEvent{mouseButton, callback, IsMouseClicked})
	return eh
}

// OnDClick sets callback when mouse button `mouseButton` is double-clicked.
func (eh *EventHandler) OnDClick(mouseButton MouseButton, callback func()) *EventHandler {
	eh.mouseEvents = append(eh.mouseEvents, mouseEvent{mouseButton, callback, IsMouseDoubleClicked})
	return eh
}

// OnMouseDown sets callback when mouse button `mouseButton` is down.
func (eh *EventHandler) OnMouseDown(mouseButton MouseButton, callback func()) *EventHandler {
	eh.mouseEvents = append(eh.mouseEvents, mouseEvent{mouseButton, callback, IsMouseDown})
	return eh
}

// OnMouseReleased sets callback when mouse button `mouseButton` is released.
func (eh *EventHandler) OnMouseReleased(mouseButton MouseButton, callback func()) *EventHandler {
	eh.mouseEvents = append(eh.mouseEvents, mouseEvent{mouseButton, callback, IsMouseReleased})
	return eh
}

// Build implements Widget interface
// nolint:gocognit,gocyclo // will fix later
func (eh *EventHandler) Build() {
	isActive := IsItemActive()

	if eh.onActivate != nil || eh.onDeactivate != nil {
		var state *eventHandlerState
		stateID := GenAutoID("eventHandlerState")
		if s := Context.GetState(stateID); s != nil {
			var isOk bool
			state, isOk = s.(*eventHandlerState)
			Assert(isOk, "EventHandler", "Build", "unexpected type of state received")
		} else {
			newState := &eventHandlerState{}
			Context.SetState(stateID, newState)
			state = newState
		}

		if eh.onActivate != nil && isActive && !state.isActive {
			state.isActive = true
			eh.onActivate()
		}

		if eh.onDeactivate != nil && !isActive && state.isActive {
			state.isActive = false
			eh.onDeactivate()
		}
	}

	if isActive && eh.onActive != nil {
		eh.onActive()
	}

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
