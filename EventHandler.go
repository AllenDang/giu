package giu

type mouseEvent struct {
	mouseButton MouseButton
	callback    func()
}

// EventHandler is a universal event handler for giu widgets.
// put giu.Event()... after any widget to handle any event
type EventHandler struct {
	hover     func()
	click     []mouseEvent
	dClick    []mouseEvent
	mouseDown []mouseEvent
}

func Event() *EventHandler {
	return &EventHandler{
		click:     make([]mouseEvent, 0),
		dClick:    make([]mouseEvent, 0),
		mouseDown: make([]mouseEvent, 0),
	}
}

func (eh *EventHandler) OnHover(onHover func()) *EventHandler {
	eh.hover = onHover
	return eh
}

func (eh *EventHandler) OnClick(mouseButton MouseButton, callback func()) *EventHandler {
	eh.click = append(eh.click, mouseEvent{mouseButton, callback})
	return eh
}

func (eh *EventHandler) OnDClick(mouseButton MouseButton, callback func()) *EventHandler {
	eh.dClick = append(eh.dClick, mouseEvent{mouseButton, callback})
	return eh
}

func (eh *EventHandler) OnMouseDown(mouseButton MouseButton, callback func()) *EventHandler {
	eh.mouseDown = append(eh.mouseDown, mouseEvent{mouseButton, callback})
	return eh
}

func (eh *EventHandler) Build() {
	if !IsItemHovered() {
		return
	}

	if len(eh.click) > 0 {
		for _, event := range eh.click {
			if event.callback != nil && IsMouseClicked(event.mouseButton) {
				event.callback()
			}
		}
	}

	if len(eh.dClick) > 0 {
		for _, event := range eh.dClick {
			if event.callback != nil && IsMouseDoubleClicked(event.mouseButton) {
				event.callback()
			}
		}
	}

	if len(eh.mouseDown) > 0 {
		for _, event := range eh.mouseDown {
			if event.callback != nil && IsMouseDoubleClicked(event.mouseButton) {
				event.callback()
			}
		}
	}
}
