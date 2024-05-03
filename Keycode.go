package giu

import "C"
import (
	"log"

	imgui "github.com/AllenDang/cimgui-go"
)

// Key represents a imgui key.
type Key imgui.Key

// These key codes are inspired by the USB HID Usage Tables v1.12 (p. 53-60),
// but re-arranged to map to 7-bit ASCII for printable keys (function keys are
// put in the 256+ range).
const (
	KeyNone           = Key(imgui.KeyNone)
	KeySpace          = Key(imgui.KeySpace)
	KeyApostrophe     = Key(imgui.KeyApostrophe)
	KeyComma          = Key(imgui.KeyComma)
	KeyMinus          = Key(imgui.KeyMinus)
	KeyPeriod         = Key(imgui.KeyPeriod)
	KeySlash          = Key(imgui.KeySlash)
	Key0              = Key(imgui.Key0)
	Key1              = Key(imgui.Key1)
	Key2              = Key(imgui.Key2)
	Key3              = Key(imgui.Key3)
	Key4              = Key(imgui.Key4)
	Key5              = Key(imgui.Key5)
	Key6              = Key(imgui.Key6)
	Key7              = Key(imgui.Key7)
	Key8              = Key(imgui.Key8)
	Key9              = Key(imgui.Key9)
	KeySemicolon      = Key(imgui.KeySemicolon)
	KeyEqual          = Key(imgui.KeyEqual)
	KeyA              = Key(imgui.KeyA)
	KeyB              = Key(imgui.KeyB)
	KeyC              = Key(imgui.KeyC)
	KeyD              = Key(imgui.KeyD)
	KeyE              = Key(imgui.KeyE)
	KeyF              = Key(imgui.KeyF)
	KeyG              = Key(imgui.KeyG)
	KeyH              = Key(imgui.KeyH)
	KeyI              = Key(imgui.KeyI)
	KeyJ              = Key(imgui.KeyJ)
	KeyK              = Key(imgui.KeyK)
	KeyL              = Key(imgui.KeyL)
	KeyM              = Key(imgui.KeyM)
	KeyN              = Key(imgui.KeyN)
	KeyO              = Key(imgui.KeyO)
	KeyP              = Key(imgui.KeyP)
	KeyQ              = Key(imgui.KeyQ)
	KeyR              = Key(imgui.KeyR)
	KeyS              = Key(imgui.KeyS)
	KeyT              = Key(imgui.KeyT)
	KeyU              = Key(imgui.KeyU)
	KeyV              = Key(imgui.KeyV)
	KeyW              = Key(imgui.KeyW)
	KeyX              = Key(imgui.KeyX)
	KeyY              = Key(imgui.KeyY)
	KeyZ              = Key(imgui.KeyZ)
	KeyLeftBracket    = Key(imgui.KeyLeftBracket)
	KeyBackslash      = Key(imgui.KeyBackslash)
	KeyRightBracket   = Key(imgui.KeyRightBracket)
	KeyGraveAccent    = Key(imgui.KeyGraveAccent)
	KeyEscape         = Key(imgui.KeyEscape)
	KeyEnter          = Key(imgui.KeyEnter)
	KeyTab            = Key(imgui.KeyTab)
	KeyBackspace      = Key(imgui.KeyBackspace)
	KeyInsert         = Key(imgui.KeyInsert)
	KeyDelete         = Key(imgui.KeyDelete)
	KeyRight          = Key(imgui.KeyRightArrow)
	KeyLeft           = Key(imgui.KeyLeftArrow)
	KeyDown           = Key(imgui.KeyDownArrow)
	KeyUp             = Key(imgui.KeyUpArrow)
	KeyPageUp         = Key(imgui.KeyPageUp)
	KeyPageDown       = Key(imgui.KeyPageDown)
	KeyHome           = Key(imgui.KeyHome)
	KeyEnd            = Key(imgui.KeyEnd)
	KeyCapsLock       = Key(imgui.KeyCapsLock)
	KeyScrollLock     = Key(imgui.KeyScrollLock)
	KeyNumLock        = Key(imgui.KeyNumLock)
	KeyPrintScreen    = Key(imgui.KeyPrintScreen)
	KeyPause          = Key(imgui.KeyPause)
	KeyF1             = Key(imgui.KeyF1)
	KeyF2             = Key(imgui.KeyF2)
	KeyF3             = Key(imgui.KeyF3)
	KeyF4             = Key(imgui.KeyF4)
	KeyF5             = Key(imgui.KeyF5)
	KeyF6             = Key(imgui.KeyF6)
	KeyF7             = Key(imgui.KeyF7)
	KeyF8             = Key(imgui.KeyF8)
	KeyF9             = Key(imgui.KeyF9)
	KeyF10            = Key(imgui.KeyF10)
	KeyF11            = Key(imgui.KeyF11)
	KeyF12            = Key(imgui.KeyF12)
	KeyNumPad0        = Key(imgui.GLFWKeyKp0)
	KeyNumPad1        = Key(imgui.GLFWKeyKp1)
	KeyNumPad2        = Key(imgui.GLFWKeyKp2)
	KeyNumPad3        = Key(imgui.GLFWKeyKp3)
	KeyNumPad4        = Key(imgui.GLFWKeyKp4)
	KeyNumPad5        = Key(imgui.GLFWKeyKp5)
	KeyNumPad6        = Key(imgui.GLFWKeyKp6)
	KeyNumPad7        = Key(imgui.GLFWKeyKp7)
	KeyNumPad8        = Key(imgui.GLFWKeyKp8)
	KeyNumPad9        = Key(imgui.GLFWKeyKp9)
	KeyNumPadDecimal  = Key(imgui.GLFWKeyKpDecimal)
	KeyNumPadDivide   = Key(imgui.GLFWKeyKpDivide)
	KeyNumPadMultiply = Key(imgui.GLFWKeyKpMultiply)
	KeyNumPadSubtract = Key(imgui.GLFWKeyKpSubtract)
	KeyNumPadAdd      = Key(imgui.GLFWKeyKpAdd)
	KeyNumPadEnter    = Key(imgui.GLFWKeyKpEnter)
	KeyNumPadEqual    = Key(imgui.GLFWKeyKpEqual)
	KeyLeftShift      = Key(imgui.KeyLeftShift)
	KeyLeftControl    = Key(imgui.KeyLeftCtrl)
	KeyLeftAlt        = Key(imgui.KeyLeftAlt)
	KeyLeftSuper      = Key(imgui.KeyLeftSuper)
	KeyRightShift     = Key(imgui.KeyRightShift)
	KeyRightControl   = Key(imgui.KeyRightCtrl)
	KeyRightAlt       = Key(imgui.KeyRightAlt)
	KeyRightSuper     = Key(imgui.KeyRightSuper)
	KeyMenu           = Key(imgui.KeyMenu)
	KeyWorld1         = Key(imgui.GLFWKeyWorld1)
	KeyWorld2         = Key(imgui.GLFWKeyWorld2)
	KeyUnknown        = Key(-1)
)

// refer glfw3.h.
func keyFromGLFWKey(k imgui.GLFWKey) Key {
	data := map[imgui.GLFWKey]Key{
		imgui.GLFWKeySpace:        KeySpace,
		imgui.GLFWKeyApostrophe:   KeyApostrophe,
		imgui.GLFWKeyComma:        KeyComma,
		imgui.GLFWKeyMinus:        KeyMinus,
		imgui.GLFWKeyPeriod:       KeyPeriod,
		imgui.GLFWKeySlash:        KeySlash,
		imgui.GLFWKey0:            Key0,
		imgui.GLFWKey1:            Key1,
		imgui.GLFWKey2:            Key2,
		imgui.GLFWKey3:            Key3,
		imgui.GLFWKey4:            Key4,
		imgui.GLFWKey5:            Key5,
		imgui.GLFWKey6:            Key6,
		imgui.GLFWKey7:            Key7,
		imgui.GLFWKey8:            Key8,
		imgui.GLFWKey9:            Key9,
		imgui.GLFWKeySemicolon:    KeySemicolon,
		imgui.GLFWKeyEqual:        KeyEqual,
		imgui.GLFWKeyA:            KeyA,
		imgui.GLFWKeyB:            KeyB,
		imgui.GLFWKeyC:            KeyC,
		imgui.GLFWKeyD:            KeyD,
		imgui.GLFWKeyE:            KeyE,
		imgui.GLFWKeyF:            KeyF,
		imgui.GLFWKeyG:            KeyG,
		imgui.GLFWKeyH:            KeyH,
		imgui.GLFWKeyI:            KeyI,
		imgui.GLFWKeyJ:            KeyJ,
		imgui.GLFWKeyK:            KeyK,
		imgui.GLFWKeyL:            KeyL,
		imgui.GLFWKeyM:            KeyM,
		imgui.GLFWKeyN:            KeyN,
		imgui.GLFWKeyO:            KeyO,
		imgui.GLFWKeyP:            KeyP,
		imgui.GLFWKeyQ:            KeyQ,
		imgui.GLFWKeyR:            KeyR,
		imgui.GLFWKeyS:            KeyS,
		imgui.GLFWKeyT:            KeyT,
		imgui.GLFWKeyU:            KeyU,
		imgui.GLFWKeyV:            KeyV,
		imgui.GLFWKeyW:            KeyW,
		imgui.GLFWKeyX:            KeyX,
		imgui.GLFWKeyY:            KeyY,
		imgui.GLFWKeyZ:            KeyZ,
		imgui.GLFWKeyLeftBracket:  KeyLeftBracket,
		imgui.GLFWKeyBackslash:    KeyBackslash,
		imgui.GLFWKeyRightBracket: KeyRightBracket,
		imgui.GLFWKeyGraveAccent:  KeyGraveAccent,
		imgui.GLFWKeyEscape:       KeyEscape,
		imgui.GLFWKeyEnter:        KeyEnter,
		imgui.GLFWKeyTab:          KeyTab,
		imgui.GLFWKeyBackspace:    KeyBackspace,
		imgui.GLFWKeyInsert:       KeyInsert,
		imgui.GLFWKeyDelete:       KeyDelete,
		imgui.GLFWKeyRight:        KeyRight,
		imgui.GLFWKeyLeft:         KeyLeft,
		imgui.GLFWKeyDown:         KeyDown,
		imgui.GLFWKeyUp:           KeyUp,
		imgui.GLFWKeyPageUp:       KeyPageUp,
		imgui.GLFWKeyPageDown:     KeyPageDown,
		imgui.GLFWKeyHome:         KeyHome,
		imgui.GLFWKeyEnd:          KeyEnd,
		imgui.GLFWKeyCapsLock:     KeyCapsLock,
		imgui.GLFWKeyScrollLock:   KeyScrollLock,
		imgui.GLFWKeyNumLock:      KeyNumLock,
		imgui.GLFWKeyPrintScreen:  KeyPrintScreen,
		imgui.GLFWKeyPause:        KeyPause,
		imgui.GLFWKeyF1:           KeyF1,
		imgui.GLFWKeyF2:           KeyF2,
		imgui.GLFWKeyF3:           KeyF3,
		imgui.GLFWKeyF4:           KeyF4,
		imgui.GLFWKeyF5:           KeyF5,
		imgui.GLFWKeyF6:           KeyF6,
		imgui.GLFWKeyF7:           KeyF7,
		imgui.GLFWKeyF8:           KeyF8,
		imgui.GLFWKeyF9:           KeyF9,
		imgui.GLFWKeyF10:          KeyF10,
		imgui.GLFWKeyF11:          KeyF11,
		imgui.GLFWKeyF12:          KeyF12,
		imgui.GLFWKeyKp0:          KeyNumPad0,
		imgui.GLFWKeyKp1:          KeyNumPad1,
		imgui.GLFWKeyKp2:          KeyNumPad2,
		imgui.GLFWKeyKp3:          KeyNumPad3,
		imgui.GLFWKeyKp4:          KeyNumPad4,
		imgui.GLFWKeyKp5:          KeyNumPad5,
		imgui.GLFWKeyKp6:          KeyNumPad6,
		imgui.GLFWKeyKp7:          KeyNumPad7,
		imgui.GLFWKeyKp8:          KeyNumPad8,
		imgui.GLFWKeyKp9:          KeyNumPad9,
		imgui.GLFWKeyKpDecimal:    KeyNumPadDecimal,
		imgui.GLFWKeyKpDivide:     KeyNumPadDivide,
		imgui.GLFWKeyKpMultiply:   KeyNumPadMultiply,
		imgui.GLFWKeyKpSubtract:   KeyNumPadSubtract,
		imgui.GLFWKeyKpAdd:        KeyNumPadAdd,
		imgui.GLFWKeyKpEnter:      KeyNumPadEnter,
		imgui.GLFWKeyKpEqual:      KeyNumPadEqual,
		imgui.GLFWKeyLeftShift:    KeyLeftShift,
		imgui.GLFWKeyLeftControl:  KeyLeftControl,
		imgui.GLFWKeyLeftAlt:      KeyLeftAlt,
		imgui.GLFWKeyLeftSuper:    KeyLeftSuper,
		imgui.GLFWKeyRightShift:   KeyRightShift,
		imgui.GLFWKeyRightControl: KeyRightControl,
		imgui.GLFWKeyRightAlt:     KeyRightAlt,
		imgui.GLFWKeyRightSuper:   KeyRightSuper,
		imgui.GLFWKeyMenu:         KeyMenu,
		imgui.GLFWKeyWorld1:       KeyWorld1,
		imgui.GLFWKeyWorld2:       KeyWorld2,
		-1:                        KeyUnknown,
	}

	if v, ok := data[k]; ok {
		return v
	}

	log.Panicf("Unknown key: %v", k)

	return 0
}

// Modifier represents imgui.Modifier.
type Modifier imgui.Key

// modifier keys.
const (
	ModNone     Modifier = 0
	ModControl           = Modifier(imgui.GLFWModControl)
	ModAlt               = Modifier(imgui.GLFWModAlt)
	ModSuper             = Modifier(imgui.GLFWModSuper)
	ModShift             = Modifier(imgui.GLFWModShift)
	ModCapsLock          = Modifier(imgui.GLFWModCapsLock)
	ModNumLock           = Modifier(imgui.GLFWModNumLock)
)

type Action int

const (
	Release Action = iota
	Press
	Repeat
)
