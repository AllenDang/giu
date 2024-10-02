package giu

import "C"
import (
	"log"

	"github.com/AllenDang/cimgui-go/backend/glfwbackend"
	"github.com/AllenDang/cimgui-go/imgui"
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
	KeyNumPad0        = Key(glfwbackend.GLFWKeyKp0)
	KeyNumPad1        = Key(glfwbackend.GLFWKeyKp1)
	KeyNumPad2        = Key(glfwbackend.GLFWKeyKp2)
	KeyNumPad3        = Key(glfwbackend.GLFWKeyKp3)
	KeyNumPad4        = Key(glfwbackend.GLFWKeyKp4)
	KeyNumPad5        = Key(glfwbackend.GLFWKeyKp5)
	KeyNumPad6        = Key(glfwbackend.GLFWKeyKp6)
	KeyNumPad7        = Key(glfwbackend.GLFWKeyKp7)
	KeyNumPad8        = Key(glfwbackend.GLFWKeyKp8)
	KeyNumPad9        = Key(glfwbackend.GLFWKeyKp9)
	KeyNumPadDecimal  = Key(glfwbackend.GLFWKeyKpDecimal)
	KeyNumPadDivide   = Key(glfwbackend.GLFWKeyKpDivide)
	KeyNumPadMultiply = Key(glfwbackend.GLFWKeyKpMultiply)
	KeyNumPadSubtract = Key(glfwbackend.GLFWKeyKpSubtract)
	KeyNumPadAdd      = Key(glfwbackend.GLFWKeyKpAdd)
	KeyNumPadEnter    = Key(glfwbackend.GLFWKeyKpEnter)
	KeyNumPadEqual    = Key(glfwbackend.GLFWKeyKpEqual)
	KeyLeftShift      = Key(imgui.KeyLeftShift)
	KeyLeftControl    = Key(imgui.KeyLeftCtrl)
	KeyLeftAlt        = Key(imgui.KeyLeftAlt)
	KeyLeftSuper      = Key(imgui.KeyLeftSuper)
	KeyRightShift     = Key(imgui.KeyRightShift)
	KeyRightControl   = Key(imgui.KeyRightCtrl)
	KeyRightAlt       = Key(imgui.KeyRightAlt)
	KeyRightSuper     = Key(imgui.KeyRightSuper)
	KeyMenu           = Key(imgui.KeyMenu)
	KeyWorld1         = Key(glfwbackend.GLFWKeyWorld1)
	KeyWorld2         = Key(glfwbackend.GLFWKeyWorld2)
	KeyUnknown        = Key(-1)
)

// refer glfw3.h.
func keyFromGLFWKey(k glfwbackend.GLFWKey) Key {
	data := map[glfwbackend.GLFWKey]Key{
		glfwbackend.GLFWKeySpace:        KeySpace,
		glfwbackend.GLFWKeyApostrophe:   KeyApostrophe,
		glfwbackend.GLFWKeyComma:        KeyComma,
		glfwbackend.GLFWKeyMinus:        KeyMinus,
		glfwbackend.GLFWKeyPeriod:       KeyPeriod,
		glfwbackend.GLFWKeySlash:        KeySlash,
		glfwbackend.GLFWKey0:            Key0,
		glfwbackend.GLFWKey1:            Key1,
		glfwbackend.GLFWKey2:            Key2,
		glfwbackend.GLFWKey3:            Key3,
		glfwbackend.GLFWKey4:            Key4,
		glfwbackend.GLFWKey5:            Key5,
		glfwbackend.GLFWKey6:            Key6,
		glfwbackend.GLFWKey7:            Key7,
		glfwbackend.GLFWKey8:            Key8,
		glfwbackend.GLFWKey9:            Key9,
		glfwbackend.GLFWKeySemicolon:    KeySemicolon,
		glfwbackend.GLFWKeyEqual:        KeyEqual,
		glfwbackend.GLFWKeyA:            KeyA,
		glfwbackend.GLFWKeyB:            KeyB,
		glfwbackend.GLFWKeyC:            KeyC,
		glfwbackend.GLFWKeyD:            KeyD,
		glfwbackend.GLFWKeyE:            KeyE,
		glfwbackend.GLFWKeyF:            KeyF,
		glfwbackend.GLFWKeyG:            KeyG,
		glfwbackend.GLFWKeyH:            KeyH,
		glfwbackend.GLFWKeyI:            KeyI,
		glfwbackend.GLFWKeyJ:            KeyJ,
		glfwbackend.GLFWKeyK:            KeyK,
		glfwbackend.GLFWKeyL:            KeyL,
		glfwbackend.GLFWKeyM:            KeyM,
		glfwbackend.GLFWKeyN:            KeyN,
		glfwbackend.GLFWKeyO:            KeyO,
		glfwbackend.GLFWKeyP:            KeyP,
		glfwbackend.GLFWKeyQ:            KeyQ,
		glfwbackend.GLFWKeyR:            KeyR,
		glfwbackend.GLFWKeyS:            KeyS,
		glfwbackend.GLFWKeyT:            KeyT,
		glfwbackend.GLFWKeyU:            KeyU,
		glfwbackend.GLFWKeyV:            KeyV,
		glfwbackend.GLFWKeyW:            KeyW,
		glfwbackend.GLFWKeyX:            KeyX,
		glfwbackend.GLFWKeyY:            KeyY,
		glfwbackend.GLFWKeyZ:            KeyZ,
		glfwbackend.GLFWKeyLeftBracket:  KeyLeftBracket,
		glfwbackend.GLFWKeyBackslash:    KeyBackslash,
		glfwbackend.GLFWKeyRightBracket: KeyRightBracket,
		glfwbackend.GLFWKeyGraveAccent:  KeyGraveAccent,
		glfwbackend.GLFWKeyEscape:       KeyEscape,
		glfwbackend.GLFWKeyEnter:        KeyEnter,
		glfwbackend.GLFWKeyTab:          KeyTab,
		glfwbackend.GLFWKeyBackspace:    KeyBackspace,
		glfwbackend.GLFWKeyInsert:       KeyInsert,
		glfwbackend.GLFWKeyDelete:       KeyDelete,
		glfwbackend.GLFWKeyRight:        KeyRight,
		glfwbackend.GLFWKeyLeft:         KeyLeft,
		glfwbackend.GLFWKeyDown:         KeyDown,
		glfwbackend.GLFWKeyUp:           KeyUp,
		glfwbackend.GLFWKeyPageUp:       KeyPageUp,
		glfwbackend.GLFWKeyPageDown:     KeyPageDown,
		glfwbackend.GLFWKeyHome:         KeyHome,
		glfwbackend.GLFWKeyEnd:          KeyEnd,
		glfwbackend.GLFWKeyCapsLock:     KeyCapsLock,
		glfwbackend.GLFWKeyScrollLock:   KeyScrollLock,
		glfwbackend.GLFWKeyNumLock:      KeyNumLock,
		glfwbackend.GLFWKeyPrintScreen:  KeyPrintScreen,
		glfwbackend.GLFWKeyPause:        KeyPause,
		glfwbackend.GLFWKeyF1:           KeyF1,
		glfwbackend.GLFWKeyF2:           KeyF2,
		glfwbackend.GLFWKeyF3:           KeyF3,
		glfwbackend.GLFWKeyF4:           KeyF4,
		glfwbackend.GLFWKeyF5:           KeyF5,
		glfwbackend.GLFWKeyF6:           KeyF6,
		glfwbackend.GLFWKeyF7:           KeyF7,
		glfwbackend.GLFWKeyF8:           KeyF8,
		glfwbackend.GLFWKeyF9:           KeyF9,
		glfwbackend.GLFWKeyF10:          KeyF10,
		glfwbackend.GLFWKeyF11:          KeyF11,
		glfwbackend.GLFWKeyF12:          KeyF12,
		glfwbackend.GLFWKeyKp0:          KeyNumPad0,
		glfwbackend.GLFWKeyKp1:          KeyNumPad1,
		glfwbackend.GLFWKeyKp2:          KeyNumPad2,
		glfwbackend.GLFWKeyKp3:          KeyNumPad3,
		glfwbackend.GLFWKeyKp4:          KeyNumPad4,
		glfwbackend.GLFWKeyKp5:          KeyNumPad5,
		glfwbackend.GLFWKeyKp6:          KeyNumPad6,
		glfwbackend.GLFWKeyKp7:          KeyNumPad7,
		glfwbackend.GLFWKeyKp8:          KeyNumPad8,
		glfwbackend.GLFWKeyKp9:          KeyNumPad9,
		glfwbackend.GLFWKeyKpDecimal:    KeyNumPadDecimal,
		glfwbackend.GLFWKeyKpDivide:     KeyNumPadDivide,
		glfwbackend.GLFWKeyKpMultiply:   KeyNumPadMultiply,
		glfwbackend.GLFWKeyKpSubtract:   KeyNumPadSubtract,
		glfwbackend.GLFWKeyKpAdd:        KeyNumPadAdd,
		glfwbackend.GLFWKeyKpEnter:      KeyNumPadEnter,
		glfwbackend.GLFWKeyKpEqual:      KeyNumPadEqual,
		glfwbackend.GLFWKeyLeftShift:    KeyLeftShift,
		glfwbackend.GLFWKeyLeftControl:  KeyLeftControl,
		glfwbackend.GLFWKeyLeftAlt:      KeyLeftAlt,
		glfwbackend.GLFWKeyLeftSuper:    KeyLeftSuper,
		glfwbackend.GLFWKeyRightShift:   KeyRightShift,
		glfwbackend.GLFWKeyRightControl: KeyRightControl,
		glfwbackend.GLFWKeyRightAlt:     KeyRightAlt,
		glfwbackend.GLFWKeyRightSuper:   KeyRightSuper,
		glfwbackend.GLFWKeyMenu:         KeyMenu,
		glfwbackend.GLFWKeyWorld1:       KeyWorld1,
		glfwbackend.GLFWKeyWorld2:       KeyWorld2,
		-1:                              KeyUnknown,
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
	ModControl           = Modifier(glfwbackend.GLFWModControl)
	ModAlt               = Modifier(glfwbackend.GLFWModAlt)
	ModSuper             = Modifier(glfwbackend.GLFWModSuper)
	ModShift             = Modifier(glfwbackend.GLFWModShift)
	ModCapsLock          = Modifier(glfwbackend.GLFWModCapsLock)
	ModNumLock           = Modifier(glfwbackend.GLFWModNumLock)
)

// Action represents key status change type.
type Action int

// Actions.
const (
	Release Action = iota
	Press
	Repeat
)
