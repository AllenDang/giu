package giu

type idSetter interface {
	setId(index int)
}

type Widget interface {
	Build()
}

type Layout []Widget

func (l Layout) Build() {
	for _, w := range l {
		if w != nil {
			// Set id for children
			if ids, ok := w.(idSetter); ok {
				index := Context.getWidgetIndexAndIncr()
				ids.setId(index)
			}
			w.Build()
		}
	}
}
