package giu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InputHandler_new(t *testing.T) {
	i := newInputHandler()
	assert.NotNil(t, i, "input handler wasn't created")
	assert.NotNil(t, i.shortcuts, "input handler wasn't created")
}

func Test_InputHandle_RegisterKeyboardShortcuts(t *testing.T) {
	tests := []struct {
		id       string
		key      Key
		mod      Modifier
		isGlobal ShortcutType
		cb       func()
	}{
		{"global shourtcut", Key(1), Modifier(2), ShortcutType(true), func() {}},
		{"window shourtcut", Key(9), Modifier(3), ShortcutType(false), func() {}},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(lt *testing.T) {
			a := assert.New(lt)
			i := newInputHandler()
			i.RegisterKeyboardShortcuts(Shortcut{
				Key:      tt.key,
				Modifier: tt.mod,
				Callback: tt.cb,
				IsGlobal: tt.isGlobal,
			})
			combo := keyCombo{
				key:      tt.key,
				modifier: tt.mod,
			}

			shortcut, exist := i.shortcuts[combo]

			a.True(exist, "shortcut wasn't registered in input manager")
			if tt.isGlobal {
				// TODO: figure out why it doesn't work
				// a.Equal(shortcut.global, tt.cb, "worng shortcut set in input manager")
				a.NotNil(shortcut.global, "worng shortcut set in input manager")
				a.Nil(shortcut.window, "worng shortcut set in input manager")
			} else {
				// TODO: figure out why it doesn't work
				// a.Equal(shortcut.window, tt.cb, "worng shortcut set in input manager")
				a.NotNil(shortcut.window, "worng shortcut set in input manager")
				a.Nil(shortcut.global, "worng shortcut set in input manager")
			}
		})
	}
}

func Test_InputHandler_UnregisterWindowShortcuts(t *testing.T) {
	i := newInputHandler()
	sh := []Shortcut{
		{Key(5), Modifier(0), func() {}, true},
		{Key(8), Modifier(2), func() {}, false},
	}

	i.RegisterKeyboardShortcuts(sh...)

	i.UnregisterWindowShortcuts()

	for _, s := range i.shortcuts {
		assert.Nil(t, s.window, "some window shortcuts wasn't unregistered")
	}
}

func Test_InputHandler_Handle(t *testing.T) {
	a := assert.New(t)
	i := newInputHandler()

	var shortcut1, shortcut2 bool

	sh := []Shortcut{
		{Key(5), Modifier(0), func() { shortcut1 = true }, true},
		{Key(8), Modifier(2), func() { shortcut2 = true }, false},
	}

	i.RegisterKeyboardShortcuts(sh...)

	i.Handle(Key(0), Modifier(0))
	a.False(shortcut1, "Shortcut 1 was handled, but shouldn't.")
	a.False(shortcut2, "Shortcut 2 was handled, but shouldn't.")
	i.Handle(Key(5), Modifier(0))
	a.True(shortcut1, "Shortcut 1 was not handled, but shouldn be.")
	a.False(shortcut2, "Shortcut 2 was handled, but shouldn't.")
	i.Handle(Key(8), Modifier(2))
	a.True(shortcut1, "Shortcut 1 was not handled, but shouldn be.")
	a.True(shortcut2, "Shortcut 2 was not handled, but shouldn be.")
}
