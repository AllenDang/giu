package giu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type teststate struct{}

func (s *teststate) Dispose() {
	// noop
}

func Test_SetGetState(t *testing.T) {
	tests := []struct {
		id   string
		data Disposable
	}{
		{"nil", nil},
		{"pointer", &teststate{}},
	}

	for _, tc := range tests {
		t.Run(tc.id, func(t *testing.T) {
			ctx := context{}
			ctx.SetState(tc.id, tc.data)
			restored := ctx.GetState(tc.id)
			assert.Equal(t, tc.data, restored, "unexpected state restored")
		})
	}
}

func Test_GetWidgetIndex(t *testing.T) {
	ctx := context{}
	for i := 0; i <= 3; i++ {
		assert.Equal(t, i, ctx.GetWidgetIndex(), "widget index wasn't increased")
	}
}
