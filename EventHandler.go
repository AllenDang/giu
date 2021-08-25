package giu

type mouseEvent struct {
	mouseButton MouseButton
	callback    func()
}

// EventHandler is a universal event handler for giu widgets.
// put giu.Event()... after any widget to handle any event
type EventHandler struct {
	hover  func()
	click  []mouseEvent
	dClick []mouseEvent
}

func Event() *EventHandler {
	return &EventHandler{
		click:  make([]mouseEvent, 0),
		dClick: make([]mouseEvent, 0),
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

func (eh *EventHandler) Build() {
	isHovered := IsItemHovered()

	if len(eh.click) > 0 {
		for _, event := range eh.click {
			if event.callback != nil && isHovered && IsMouseClicked(event.mouseButton) {
				event.callback()
			}
		}
	}

	if len(eh.dClick) > 0 {
		for _, event := range eh.dClick {
			if event.callback != nil && isHovered && IsMouseDoubleClicked(event.mouseButton) {
				event.callback()
			}
		}
	}
}
