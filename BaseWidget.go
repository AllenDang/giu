package giu

type BaseWidget struct {
	width float32
}

func (w *BaseWidget) Width() float32 {
	return w.width
}
