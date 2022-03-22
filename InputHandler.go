package giu

// input menager is used to register a keyboard shortcuts in an app.

// Shortcut represents a keyboard shortcut.
type Shortcut struct {
	Key      Key
	Modifier Modifier
	Callback func()
	IsGlobal ShortcutType
}

// WindowShortcut represents a window-level shortcut
// could be used as an argument to (*Window).RegisterKeyboardShortcuts.
type WindowShortcut struct {
	Key      Key
	Modifier Modifier
	Callback func()
}

// ShortcutType represens a type of shortcut (global or local).
type ShortcutType bool

const (
	// GlobalShortcut is registered for all the app.
	GlobalShortcut ShortcutType = true

	// LocalShortcut is registered for current window only.
	LocalShortcut ShortcutType = false
)

// InputHandler is an interface which needs to be implemented
// by user-definied input handlers.
type InputHandler interface {
	// RegisterKeyboardShortcuts adds a specified shortcuts into input handler
	RegisterKeyboardShortcuts(...Shortcut)
	// UnregisterKeyboardShortcuts removes iwndow shourtcuts from input handler
	UnregisterWindowShortcuts()
	// Handle handles a shortcut
	Handle(Key, Modifier)
}

// --- Default implementation of giu input manager ---

var _ InputHandler = &inputHandler{}

func newInputHandler() *inputHandler {
	return &inputHandler{
		shortcuts: make(map[keyCombo]*callbacks),
	}
}

type inputHandler struct {
	shortcuts map[keyCombo]*callbacks
}

func (i *inputHandler) RegisterKeyboardShortcuts(s ...Shortcut) {
	for _, shortcut := range s {
		combo := keyCombo{shortcut.Key, shortcut.Modifier}

		cb, isRegistered := i.shortcuts[combo]
		if !isRegistered {
			cb = &callbacks{}
		}

		if shortcut.IsGlobal {
			cb.global = shortcut.Callback
		} else {
			cb.window = shortcut.Callback
		}

		i.shortcuts[combo] = cb
	}
}

func (i *inputHandler) UnregisterWindowShortcuts() {
	for _, s := range i.shortcuts {
		s.window = nil
	}
}

func (i *inputHandler) Handle(key Key, mod Modifier) {
	for combo, cb := range i.shortcuts {
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

type keyCombo struct {
	key      Key
	modifier Modifier
}

type callbacks struct {
	global func()
	window func()
}
