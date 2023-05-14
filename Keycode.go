package giu

import (
	imgui "github.com/AllenDang/cimgui-go"
)

// Key represents a imgui key.
type Key imgui.Key

// These key codes are inspired by the USB HID Usage Tables v1.12 (p. 53-60),
// but re-arranged to map to 7-bit ASCII for printable keys (function keys are
// put in the 256+ range).
const (
	KeyNone         Key = Key(imgui.KeyNone)
	KeyUnknown          = KeyNone // DEPRECATED: since cimgui-go migration use KeyNone
	KeySpace        Key = Key(imgui.KeySpace)
	KeyApostrophe   Key = Key(imgui.KeyApostrophe)
	KeyComma        Key = Key(imgui.KeyComma)
	KeyMinus        Key = Key(imgui.KeyMinus)
	KeyPeriod       Key = Key(imgui.KeyPeriod)
	KeySlash        Key = Key(imgui.KeySlash)
	Key0            Key = Key(imgui.Key0)
	Key1            Key = Key(imgui.Key1)
	Key2            Key = Key(imgui.Key2)
	Key3            Key = Key(imgui.Key3)
	Key4            Key = Key(imgui.Key4)
	Key5            Key = Key(imgui.Key5)
	Key6            Key = Key(imgui.Key6)
	Key7            Key = Key(imgui.Key7)
	Key8            Key = Key(imgui.Key8)
	Key9            Key = Key(imgui.Key9)
	KeySemicolon    Key = Key(imgui.KeySemicolon)
	KeyEqual        Key = Key(imgui.KeyEqual)
	KeyA            Key = Key(imgui.KeyA)
	KeyB            Key = Key(imgui.KeyB)
	KeyC            Key = Key(imgui.KeyC)
	KeyD            Key = Key(imgui.KeyD)
	KeyE            Key = Key(imgui.KeyE)
	KeyF            Key = Key(imgui.KeyF)
	KeyG            Key = Key(imgui.KeyG)
	KeyH            Key = Key(imgui.KeyH)
	KeyI            Key = Key(imgui.KeyI)
	KeyJ            Key = Key(imgui.KeyJ)
	KeyK            Key = Key(imgui.KeyK)
	KeyL            Key = Key(imgui.KeyL)
	KeyM            Key = Key(imgui.KeyM)
	KeyN            Key = Key(imgui.KeyN)
	KeyO            Key = Key(imgui.KeyO)
	KeyP            Key = Key(imgui.KeyP)
	KeyQ            Key = Key(imgui.KeyQ)
	KeyR            Key = Key(imgui.KeyR)
	KeyS            Key = Key(imgui.KeyS)
	KeyT            Key = Key(imgui.KeyT)
	KeyU            Key = Key(imgui.KeyU)
	KeyV            Key = Key(imgui.KeyV)
	KeyW            Key = Key(imgui.KeyW)
	KeyX            Key = Key(imgui.KeyX)
	KeyY            Key = Key(imgui.KeyY)
	KeyZ            Key = Key(imgui.KeyZ)
	KeyLeftBracket  Key = Key(imgui.KeyLeftBracket)
	KeyBackslash    Key = Key(imgui.KeyBackslash)
	KeyRightBracket Key = Key(imgui.KeyRightBracket)
	KeyGraveAccent  Key = Key(imgui.KeyGraveAccent)
	//KeyWorld1       Key = Key(imgui.KeyWorld1)
	//KeyWorld2       Key = Key(imgui.KeyWorld2)
	KeyEscape      Key = Key(imgui.KeyEscape)
	KeyEnter       Key = Key(imgui.KeyEnter)
	KeyTab         Key = Key(imgui.KeyTab)
	KeyBackspace   Key = Key(imgui.KeyBackspace)
	KeyInsert      Key = Key(imgui.KeyInsert)
	KeyDelete      Key = Key(imgui.KeyDelete)
	KeyRight       Key = Key(imgui.KeyRightArrow)
	KeyLeft        Key = Key(imgui.KeyLeftArrow)
	KeyDown        Key = Key(imgui.KeyDownArrow)
	KeyUp          Key = Key(imgui.KeyUpArrow)
	KeyPageUp      Key = Key(imgui.KeyPageUp)
	KeyPageDown    Key = Key(imgui.KeyPageDown)
	KeyHome        Key = Key(imgui.KeyHome)
	KeyEnd         Key = Key(imgui.KeyEnd)
	KeyCapsLock    Key = Key(imgui.KeyCapsLock)
	KeyScrollLock  Key = Key(imgui.KeyScrollLock)
	KeyNumLock     Key = Key(imgui.KeyNumLock)
	KeyPrintScreen Key = Key(imgui.KeyPrintScreen)
	KeyPause       Key = Key(imgui.KeyPause)
	KeyF1          Key = Key(imgui.KeyF1)
	KeyF2          Key = Key(imgui.KeyF2)
	KeyF3          Key = Key(imgui.KeyF3)
	KeyF4          Key = Key(imgui.KeyF4)
	KeyF5          Key = Key(imgui.KeyF5)
	KeyF6          Key = Key(imgui.KeyF6)
	KeyF7          Key = Key(imgui.KeyF7)
	KeyF8          Key = Key(imgui.KeyF8)
	KeyF9          Key = Key(imgui.KeyF9)
	KeyF10         Key = Key(imgui.KeyF10)
	KeyF11         Key = Key(imgui.KeyF11)
	KeyF12         Key = Key(imgui.KeyF12)
	//KeyF13          Key = Key(imgui.KeyF13)
	//KeyF14          Key = Key(imgui.KeyF14)
	//KeyF15          Key = Key(imgui.KeyF15)
	//KeyF16          Key = Key(imgui.KeyF16)
	//KeyF17          Key = Key(imgui.KeyF17)
	//KeyF18          Key = Key(imgui.KeyF18)
	//KeyF19          Key = Key(imgui.KeyF19)
	//KeyF20          Key = Key(imgui.KeyF20)
	//KeyF21          Key = Key(imgui.KeyF21)
	//KeyF22          Key = Key(imgui.KeyF22)
	//KeyF23          Key = Key(imgui.KeyF23)
	//KeyF24          Key = Key(imgui.KeyF24)
	//KeyF25          Key = Key(imgui.KeyF25)
	//KeyKP0          Key = Key(imgui.KeyKP0)
	//KeyKP1          Key = Key(imgui.KeyKP1)
	//KeyKP2          Key = Key(imgui.KeyKP2)
	//KeyKP3          Key = Key(imgui.KeyKP3)
	//KeyKP4          Key = Key(imgui.KeyKP4)
	//KeyKP5          Key = Key(imgui.KeyKP5)
	//KeyKP6          Key = Key(imgui.KeyKP6)
	//KeyKP7          Key = Key(imgui.KeyKP7)
	//KeyKP8          Key = Key(imgui.KeyKP8)
	//KeyKP9          Key = Key(imgui.KeyKP9)
	//KeyKPDecimal    Key = Key(imgui.KeyKPDecimal)
	//KeyKPDivide     Key = Key(imgui.KeyKPDivide)
	//KeyKPMultiply   Key = Key(imgui.KeyKPMultiply)
	//KeyKPSubtract   Key = Key(imgui.KeyKPSubtract)
	//KeyKPAdd        Key = Key(imgui.KeyKPAdd)
	//KeyKPEnter      Key = Key(imgui.KeyKPEnter)
	//KeyKPEqual      Key = Key(imgui.KeyKPEqual)
	KeyLeftShift    Key = Key(imgui.KeyLeftShift)
	KeyLeftControl  Key = Key(imgui.KeyLeftCtrl)
	KeyLeftAlt      Key = Key(imgui.KeyLeftAlt)
	KeyLeftSuper    Key = Key(imgui.KeyLeftSuper)
	KeyRightShift   Key = Key(imgui.KeyRightShift)
	KeyRightControl Key = Key(imgui.KeyRightCtrl)
	KeyRightAlt     Key = Key(imgui.KeyRightAlt)
	KeyRightSuper   Key = Key(imgui.KeyRightSuper)
	KeyMenu         Key = Key(imgui.KeyMenu)
	//KeyLast         Key = Key(imgui.KeyLast)
)

// Modifier represents imgui.Modifier.
// TODO: consider if it is necessary and ad constants
type Modifier imgui.Key

// modifier keys.
const (
	ModNone Modifier = iota
	ModControl
	ModAlt
	ModSuper
	ModShift
	ModCapsLock
	ModNumLock
)

type Action int

const (
	Release Action = iota
	Press
	Repeat
)
