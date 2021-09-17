package giu

const (
	// Auto is used to widget.Size to indicate height or width to occupy available spaces
	Auto float32 = -1
)

type Widget interface {
	Build()
}

var (
	_ Widget    = Layout{}
	_ Splitable = Layout{}
)

type Layout []Widget

func (l Layout) Build() {
	for _, w := range l {
		if w != nil {
			w.Build()
		}
	}
}

// Splitable is implemented by widgets, which can be split (ranged)
// Layout implements Splitable.
type Splitable interface {
	Range(func(w Widget))
}

// Range ranges ofer the Layout, calling rangeFunc
// on each loop iteration.
func (l Layout) Range(rangeFunc func(Widget)) {
	for _, w := range l {
		if splitable, canRange := w.(Splitable); canRange {
			splitable.Range(rangeFunc)
			continue
		}

		rangeFunc(w)
	}
}
