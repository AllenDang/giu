package giu

type Widget func()

func Layout(widgets ...Widget) Widget {
	return func() {
		for _, w := range widgets {
			w()
		}
	}
}
