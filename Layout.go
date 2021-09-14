package giu

const (
	// Auto is used to widget.Size to indicate height or width to occupy available spaces
	Auto float32 = -1
)

// Widget is a base unit of giu rendering system.
// each widget just needs to implement Build method which is called,
// when widget needs to be rendered.
type Widget interface {
	Build()
}

var _ Widget = Layout{}

// Layout is a set of widgets. It implements Widget interface so
// Layout can be used as a widget.
type Layout []Widget

// Build implements Widget interface
func (l Layout) Build() {
	for _, w := range l {
		if w != nil {
			w.Build()
		}
	}
}
