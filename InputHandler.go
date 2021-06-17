// input menager is used to register a keyboard shortcuts in an app
package giu

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

// store keyboard shortcuts
var shortcuts map[keyCombo]*callbacks

func init() {
	shortcuts = make(map[keyCombo]*callbacks)
}

// ShortcutType represens a type of shortcut (global or local)
type ShortcutType bool

const (
	// GlobalShortcut is registered for all the app
	GlobalShortcut ShortcutType = true

	// LocLShortcut is registered for current window only
	LocalShortcut ShortcutType = false
)

type keyCombo struct {
	key      glfw.Key
	modifier glfw.ModifierKey
}

type callbacks struct {
	global func()
	window func()
}

type Shortcut struct {
	Key      Key
	Modifier Modifier
	Callback func()
	IsGlobal ShortcutType
}

func RegisterKeyboardShortcuts(s ...Shortcut) {
	for _, shortcut := range s {
		combo := keyCombo{glfw.Key(shortcut.Key), glfw.ModifierKey(shortcut.Modifier)}

		cb, isRegistered := shortcuts[combo]
		if !isRegistered {
			cb = &callbacks{}
		}

		if shortcut.IsGlobal {
			cb.global = shortcut.Callback
		} else {
			cb.window = shortcut.Callback
		}

		shortcuts[combo] = cb
	}
}

func unregisterWindowShortcuts() {
	for _, s := range shortcuts {
		s.window = nil
	}
}

func handler(key glfw.Key, mod glfw.ModifierKey, action glfw.Action) {
	if action != glfw.Press {
		return
	}

	for combo, cb := range shortcuts {
		if combo.key != key || combo.modifier != mod {
			continue
		}

		if cb.window != nil {
			cb.window()
		} else if cb.global != nil {
			cb.global()
		}
	}
}

type WindowShortcut struct {
	Key      Key
	Modifier Modifier
	Callback func()
}
