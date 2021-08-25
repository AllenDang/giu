package giu

type mouseEvent struct {
	mouseButton MouseButton
	callback    func()
	cond        func(MouseButton) bool
}

// EventHandler is a universal event handler for giu widgets.
// put giu.Event()... after any widget to handle any event
type EventHandler struct {
	hover       func()
	mouseEvents []mouseEvent
}

func Event() *EventHandler {
	return &EventHandler{
		mouseEvents: make([]mouseEvent, 0),
	}
}

func (eh *EventHandler) OnHover(onHover func()) *EventHandler {
	eh.hover = onHover
	return eh
}

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

func (eh *EventHandler) Build() {
	if !IsItemHovered() {
		return
	}

	if len(eh.mouseEvents) > 0 {
		for _, event := range eh.mouseEvents {
			if event.callback != nil && event.cond(event.mouseButton) {
				event.callback()
			}
		}
	}
}
