package giu

type IdSetter interface {
	SetId(index int)
}

type Widget interface {
	Build()
}

type Layout []Widget

func (l Layout) Build() {
	for _, w := range l {
		if w != nil {
			// Set id for children
			if ids, ok := w.(IdSetter); ok {
				index := Context.GetWidgetIndex()
				ids.SetId(index)
			}
			w.Build()
		}
	}
}
