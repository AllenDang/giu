package giu

import imgui "github.com/AllenDang/cimgui-go"

// HoveredFlags represents a hovered flags.
type HoveredFlags imgui.HoveredFlags

// hovered flags list.
const (
	// HoveredFlagsNone Return true if directly over the item/window, not obstructed by another window,
	// not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsNone HoveredFlags = imgui.HoveredFlagsNone
	// HoveredFlagsChildWindows IsWindowHovered() only: Return true if any children of the window is hovered.
	HoveredFlagsChildWindows HoveredFlags = imgui.HoveredFlagsChildWindows
	// HoveredFlagsRootWindow IsWindowHovered() only: Test from root window (top most parent of the current hierarchy).
	HoveredFlagsRootWindow HoveredFlags = imgui.HoveredFlagsRootWindow
	// HoveredFlagsAnyWindow IsWindowHovered() only: Return true if any window is hovered.
	HoveredFlagsAnyWindow HoveredFlags = imgui.HoveredFlagsAnyWindow
	// HoveredFlagsAllowWhenBlockedByPopup Return true even if a popup window is normally blocking access to this item/window.
	HoveredFlagsAllowWhenBlockedByPopup HoveredFlags = imgui.HoveredFlagsAllowWhenBlockedByPopup
	// HoveredFlagsAllowWhenBlockedByActiveItem Return true even if an active item is blocking access to this item/window.
	// Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenBlockedByActiveItem HoveredFlags = imgui.HoveredFlagsAllowWhenBlockedByActiveItem
	// HoveredFlagsAllowWhenOverlapped Return true even if the position is overlapped by another window.
	HoveredFlagsAllowWhenOverlapped HoveredFlags = imgui.HoveredFlagsAllowWhenOverlapped
	// HoveredFlagsAllowWhenDisabled Return true even if the item is disabled.
	HoveredFlagsAllowWhenDisabled HoveredFlags = imgui.HoveredFlagsAllowWhenDisabled
)

// FocusedFlags represents imgui.FocusedFlags.
type FocusedFlags imgui.FocusedFlags

// focused flags list.
const (
	FocusedFlagsNone             = imgui.FocusedFlagsNone
	FocusedFlagsChildWindows     = imgui.FocusedFlagsChildWindows     // Return true if any children of the window is focused
	FocusedFlagsRootWindow       = imgui.FocusedFlagsRootWindow       // Test from root window (top most parent of the current hierarchy)
	FocusedFlagsAnyWindow        = imgui.FocusedFlagsAnyWindow        // Return true if any window is focused. Important: If you are trying to tell how to
	FocusedFlagsNoPopupHierarchy = imgui.FocusedFlagsNoPopupHierarchy // Do not consider popup hierarchy (do not treat popup emitter as parent of popup) (
	// FocusedFlagsDockHierarchy               = 1 << 4   // Consider docking hierarchy (treat dockspace host as parent of docked window) (when used with
	FocusedFlagsRootAndChildWindows = imgui.FocusedFlagsRootAndChildWindows
)

// Key represents a glfw key.
type Key imgui.Key

// These key codes are inspired by the USB HID Usage Tables v1.12 (p. 53-60),
// but re-arranged to map to 7-bit ASCII for printable keys (function keys are
// put in the 256+ range).
const (
	// KeyUnknown      Key = Key(glfw.KeyUnknown)
	KeySpace        Key = imgui.KeySpace
	KeyApostrophe   Key = imgui.KeyApostrophe
	KeyComma        Key = imgui.KeyComma
	KeyMinus        Key = imgui.KeyMinus
	KeyPeriod       Key = imgui.KeyPeriod
	KeySlash        Key = imgui.KeySlash
	Key0            Key = imgui.Key0
	Key1            Key = imgui.Key1
	Key2            Key = imgui.Key2
	Key3            Key = imgui.Key3
	Key4            Key = imgui.Key4
	Key5            Key = imgui.Key5
	Key6            Key = imgui.Key6
	Key7            Key = imgui.Key7
	Key8            Key = imgui.Key8
	Key9            Key = imgui.Key9
	KeySemicolon    Key = imgui.KeySemicolon
	KeyEqual        Key = imgui.KeyEqual
	KeyA            Key = imgui.KeyA
	KeyB            Key = imgui.KeyB
	KeyC            Key = imgui.KeyC
	KeyD            Key = imgui.KeyD
	KeyE            Key = imgui.KeyE
	KeyF            Key = imgui.KeyF
	KeyG            Key = imgui.KeyG
	KeyH            Key = imgui.KeyH
	KeyI            Key = imgui.KeyI
	KeyJ            Key = imgui.KeyJ
	KeyK            Key = imgui.KeyK
	KeyL            Key = imgui.KeyL
	KeyM            Key = imgui.KeyM
	KeyN            Key = imgui.KeyN
	KeyO            Key = imgui.KeyO
	KeyP            Key = imgui.KeyP
	KeyQ            Key = imgui.KeyQ
	KeyR            Key = imgui.KeyR
	KeyS            Key = imgui.KeyS
	KeyT            Key = imgui.KeyT
	KeyU            Key = imgui.KeyU
	KeyV            Key = imgui.KeyV
	KeyW            Key = imgui.KeyW
	KeyX            Key = imgui.KeyX
	KeyY            Key = imgui.KeyY
	KeyZ            Key = imgui.KeyZ
	KeyLeftBracket  Key = imgui.KeyLeftBracket
	KeyBackslash    Key = imgui.KeyBackslash
	KeyRightBracket Key = imgui.KeyRightBracket
	KeyGraveAccent  Key = imgui.KeyGraveAccent
	// KeyWorld1       Key = imgui.KeyWorld1
	// KeyWorld2       Key = imgui.KeyWorld2
	KeyEscape      Key = imgui.KeyEscape
	KeyEnter       Key = imgui.KeyEnter
	KeyTab         Key = imgui.KeyTab
	KeyBackspace   Key = imgui.KeyBackspace
	KeyInsert      Key = imgui.KeyInsert
	KeyDelete      Key = imgui.KeyDelete
	KeyRight       Key = imgui.KeyRightArrow
	KeyLeft        Key = imgui.KeyLeftArrow
	KeyDown        Key = imgui.KeyDownArrow
	KeyUp          Key = imgui.KeyUpArrow
	KeyPageUp      Key = imgui.KeyPageUp
	KeyPageDown    Key = imgui.KeyPageDown
	KeyHome        Key = imgui.KeyHome
	KeyEnd         Key = imgui.KeyEnd
	KeyCapsLock    Key = imgui.KeyCapsLock
	KeyScrollLock  Key = imgui.KeyScrollLock
	KeyNumLock     Key = imgui.KeyNumLock
	KeyPrintScreen Key = imgui.KeyPrintScreen
	KeyPause       Key = imgui.KeyPause
	KeyF1          Key = imgui.KeyF1
	KeyF2          Key = imgui.KeyF2
	KeyF3          Key = imgui.KeyF3
	KeyF4          Key = imgui.KeyF4
	KeyF5          Key = imgui.KeyF5
	KeyF6          Key = imgui.KeyF6
	KeyF7          Key = imgui.KeyF7
	KeyF8          Key = imgui.KeyF8
	KeyF9          Key = imgui.KeyF9
	KeyF10         Key = imgui.KeyF10
	KeyF11         Key = imgui.KeyF11
	KeyF12         Key = imgui.KeyF12
	// KeyF13          Key = imgui.KeyF13
	// KeyF14          Key = imgui.KeyF14
	// KeyF15          Key = imgui.KeyF15
	// KeyF16          Key = imgui.KeyF16
	// KeyF17          Key = imgui.KeyF17
	// KeyF18          Key = imgui.KeyF18
	// KeyF19          Key = imgui.KeyF19
	// KeyF20          Key = imgui.KeyF20
	// KeyF21          Key = imgui.KeyF21
	// KeyF22          Key = imgui.KeyF22
	// KeyF23          Key = imgui.KeyF23
	// KeyF24          Key = imgui.KeyF24
	// KeyF25          Key = imgui.KeyF25
	// KeyKP0          Key = imgui.KeyKP0
	// KeyKP1          Key = imgui.KeyKP1
	// KeyKP2          Key = imgui.KeyKP2
	// KeyKP3          Key = imgui.KeyKP3
	// KeyKP4          Key = imgui.KeyKP4
	// KeyKP5          Key = imgui.KeyKP5
	// KeyKP6          Key = imgui.KeyKP6
	// KeyKP7          Key = imgui.KeyKP7
	// KeyKP8          Key = imgui.KeyKP8
	// KeyKP9          Key = imgui.KeyKP9
	// KeyKPDecimal    Key = imgui.KeyKPDecimal
	// KeyKPDivide     Key = imgui.KeyKPDivide
	// KeyKPMultiply   Key = imgui.KeyKPMultiply
	// KeyKPSubtract   Key = imgui.KeyKPSubtract
	// KeyKPAdd        Key = imgui.KeyKPAdd
	// KeyKPEnter      Key = imgui.KeyKPEnter
	// KeyKPEqual      Key = imgui.KeyKPEqual
	KeyLeftShift    Key = imgui.KeyLeftShift
	KeyLeftControl  Key = imgui.KeyLeftCtrl
	KeyLeftAlt      Key = imgui.KeyLeftAlt
	KeyLeftSuper    Key = imgui.KeyLeftSuper
	KeyRightShift   Key = imgui.KeyRightShift
	KeyRightControl Key = imgui.KeyRightCtrl
	KeyRightAlt     Key = imgui.KeyRightAlt
	KeyRightSuper   Key = imgui.KeyRightSuper
	KeyMenu         Key = imgui.KeyMenu
	// KeyLast         Key = imgui.KeyLast
)

// Modifier represents glfw.Modifier.
//type Modifier glfw.ModifierKey
//
//modifier keys.
//const (
//	ModNone     Modifier = iota
//	ModControl  Modifier = Modifier(glfw.ModControl)
//	ModAlt      Modifier = Modifier(glfw.ModAlt)
//	ModSuper    Modifier = Modifier(glfw.ModSuper)
//	ModShift    Modifier = Modifier(glfw.ModShift)
//	ModCapsLock Modifier = Modifier(glfw.ModCapsLock)
//	ModNumLock  Modifier = Modifier(glfw.ModNumLock)
//)

//type Action glfw.Action
//
//const (
//	Release Action = Action(glfw.Release)
//	Press   Action = Action(glfw.Press)
//	Repeat  Action = Action(glfw.Repeat)
//)
