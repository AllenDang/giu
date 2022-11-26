package giu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type teststate struct{}

func (s *teststate) Dispose() {
	// noop
}

type teststate2 struct{}

func (t *teststate2) Dispose() {
	// noop
}

func Test_SetGetState(t *testing.T) {
	tests := []struct {
		id   string
		data *teststate
	}{
		{"nil", nil},
		{"pointer", &teststate{}},
	}

	for _, tc := range tests {
		t.Run(tc.id, func(t *testing.T) {
			ctx := context{}
			SetState(&ctx, tc.id, tc.data)
			restored := GetState[teststate](&ctx, tc.id)
			assert.Equal(t, tc.data, restored, "unexpected state restored")
		})
	}
}

func Test_SetGetStateGeneric(t *testing.T) {
	tests := []struct {
		id   string
		data *teststate
	}{
		{"nil", nil},
		{"pointer", &teststate{}},
	}

	for _, tc := range tests {
		t.Run(tc.id, func(t *testing.T) {
			ctx := context{}
			SetState(&ctx, tc.id, tc.data)
			restored := GetState[teststate](&ctx, tc.id)
			assert.Equal(t, tc.data, restored, "unexpected state restored")
		})
	}
}

func Test_SetGetWrongStateGeneric(t *testing.T) {
	id := "id"
	data := &teststate{}
	ctx := context{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected code to assert to panic, but it didn't")
		}
	}()
	SetState(&ctx, id, data)
	GetState[teststate2](&ctx, id)
}

func Test_invalidState(t *testing.T) {
	ctx := context{}

	state1ID := "state1"
	state2ID := "state2"
	states := map[string]*teststate{
		state1ID: {},
		state2ID: {},
	}

	for i, s := range states {
		SetState(&ctx, i, s)
	}

	ctx.invalidAllState()

	_ = GetState[teststate](&ctx, state2ID)

	ctx.cleanState()

	assert.NotNil(t, GetState[teststate](&ctx, state2ID),
		"although state has been accessed during the frame, it has ben deleted by invalidAllState/cleanState")
	assert.Nil(t, GetState[teststate](&ctx, state1ID),
		"although state hasn't been accessed during the frame, it hasn't ben deleted by invalidAllState/cleanState")
}

func Test_GetWidgetIndex(t *testing.T) {
	ctx := context{}
	for i := 0; i <= 3; i++ {
		assert.Equal(t, i, ctx.GetWidgetIndex(), "widget index wasn't increased")
	}
}
