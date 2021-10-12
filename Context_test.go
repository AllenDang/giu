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

func Test_invalidState(t *testing.T) {
	ctx := context{}

	state1ID := "state1"
	state2ID := "state2"
	states := map[string]Disposable{
		state1ID: &teststate{},
		state2ID: &teststate{},
	}

	for i, s := range states {
		ctx.SetState(i, s)
	}

	ctx.invalidAllState()

	_ = ctx.GetState(state2ID)

	ctx.cleanState()

	assert.NotNil(t, ctx.GetState(state2ID),
		"altought state has been accessed during the frame, it has ben deleted by invalidAllState/cleanState")
	assert.Nil(t, ctx.GetState(state1ID),
		"altought state hasn't been accessed during the frame, it hasn't ben deleted by invalidAllState/cleanState")
}

func Test_GetWidgetIndex(t *testing.T) {
	ctx := context{}
	for i := 0; i <= 3; i++ {
		assert.Equal(t, i, ctx.GetWidgetIndex(), "widget index wasn't increased")
	}
}
