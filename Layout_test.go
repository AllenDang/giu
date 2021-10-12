package giu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testwidget struct {
	counter *int
}

func (w *testwidget) Build() {
	if w.counter == nil {
		return
	}

	*w.counter++
}

type splitablewidget struct {
	w1, w2 *testwidget
}

func (w *splitablewidget) Build() {
	w.w1.Build()
	w.w2.Build()
}

func (w *splitablewidget) Range(r func(w Widget)) {
	r(w.w1)
	r(w.w2)
}

func Test_Layout_Range(t *testing.T) {
	tests := []struct {
		name                     string
		expectedTestWidgetsCount int
		layout                   Layout
	}{
		{"standard layout", 3, Layout{
			&testwidget{},
			&testwidget{},
			&testwidget{},
		}},
		{"layout with splitable widgets", 4, Layout{
			&testwidget{},
			&splitablewidget{&testwidget{}, &testwidget{}},
			&testwidget{},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			counter := 0
			test.layout.Range(func(w Widget) {
				if _, isTestwidget := w.(*testwidget); isTestwidget {
					counter++
				}
			})

			assert.Equal(tt, test.expectedTestWidgetsCount, counter, "Layout wasn't ranged correctly")
		})
	}
}

func Test_Layout_Build(t *testing.T) {
	tests := []struct {
		name                        string
		expectedNumTestWidgetsBuilt int
		layout                      Layout
	}{
		{"standard layout", 2, Layout{
			&testwidget{},
			&testwidget{},
		}},
		{"layout with nil widgets", 2, Layout{
			&testwidget{},
			nil,
			&testwidget{},
		}},
		{"layout with nested layouts", 5, Layout{
			&testwidget{},
			Layout{
				&testwidget{},
				&testwidget{},
				Layout{
					&testwidget{},
				},
			},
			&testwidget{},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			counter := 0
			test.layout.Range(func(w Widget) {
				if tw, isTestwidget := w.(*testwidget); isTestwidget {
					tw.counter = &counter
				}
			})

			test.layout.Build()

			assert.Equal(tt, test.expectedNumTestWidgetsBuilt, counter, "layout wasn't built correctly")
		})
	}
}
