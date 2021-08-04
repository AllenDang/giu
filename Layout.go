package giu

const (
	// Auto is used to widget.Size to indicate height or width to occupy available spaces
	Auto float32 = -1
)

type Widget interface {
	Build()
}

type Layout []Widget

func (l Layout) Build() {
	for _, w := range l {
		if w != nil {
			w.Build()
		}
	}
}
