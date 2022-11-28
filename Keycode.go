package giu

import (
	"github.com/AllenDang/cimgui-go"
)

// HoveredFlags represents a hovered flags.
type HoveredFlags cimgui.ImGuiHoveredFlags

// hovered flags list.
const (
	// HoveredFlagsNone Return true if directly over the item/window, not obstructed by another window,
	// not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsNone HoveredFlags = cimgui.ImGuiHoveredFlags_None
	// HoveredFlagsChildWindows IsWindowHovered() only: Return true if any children of the window is hovered.
	HoveredFlagsChildWindows HoveredFlags = cimgui.ImGuiHoveredFlags_ChildWindows
	// HoveredFlagsRootWindow IsWindowHovered() only: Test from root window (top most parent of the current hierarchy).
	HoveredFlagsRootWindow HoveredFlags = cimgui.ImGuiHoveredFlags_RootWindow
	// HoveredFlagsAnyWindow IsWindowHovered() only: Return true if any window is hovered.
	HoveredFlagsAnyWindow HoveredFlags = cimgui.ImGuiHoveredFlags_AnyWindow
	// HoveredFlagsAllowWhenBlockedByPopup Return true even if a popup window is normally blocking access to this item/window.
	HoveredFlagsAllowWhenBlockedByPopup HoveredFlags = cimgui.ImGuiHoveredFlags_AllowWhenBlockedByPopup
	// HoveredFlagsAllowWhenBlockedByActiveItem Return true even if an active item is blocking access to this item/window.
	// Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenBlockedByActiveItem HoveredFlags = cimgui.ImGuiHoveredFlags_AllowWhenBlockedByActiveItem
	// HoveredFlagsAllowWhenOverlapped Return true even if the position is overlapped by another window.
	HoveredFlagsAllowWhenOverlapped HoveredFlags = cimgui.ImGuiHoveredFlags_AllowWhenOverlapped
	// HoveredFlagsAllowWhenDisabled Return true even if the item is disabled.
	HoveredFlagsAllowWhenDisabled HoveredFlags = cimgui.ImGuiHoveredFlags_AllowWhenDisabled
)

// FocusedFlags represents imgui.FocusedFlags.
type FocusedFlags cimgui.ImGuiFocusedFlags

// focused flags list.
const (
	FocusedFlagsNone             = cimgui.ImGuiFocusedFlags_None
	FocusedFlagsChildWindows     = cimgui.ImGuiFocusedFlags_ChildWindows     // Return true if any children of the window is focused
	FocusedFlagsRootWindow       = cimgui.ImGuiFocusedFlags_RootWindow       // Test from root window (top most parent of the current hierarchy)
	FocusedFlagsAnyWindow        = cimgui.ImGuiFocusedFlags_AnyWindow        // Return true if any window is focused. Important: If you are trying to tell how to
	FocusedFlagsNoPopupHierarchy = cimgui.ImGuiFocusedFlags_NoPopupHierarchy // Do not consider popup hierarchy (do not treat popup emitter as parent of popup) (
	// FocusedFlagsDockHierarchy               = 1 << 4   // Consider docking hierarchy (treat dockspace host as parent of docked window) (when used with
	FocusedFlagsRootAndChildWindows = cimgui.ImGuiFocusedFlags_RootAndChildWindows
)

// Key represents a glfw key.
type Key cimgui.ImGuiKey

// These key codes are inspired by the USB HID Usage Tables v1.12 (p. 53-60),
// but re-arranged to map to 7-bit ASCII for printable keys (function keys are
// put in the 256+ range).
const (
	// KeyUnknown      Key = Key(glfw.KeyUnknown)
	KeySpace        Key = cimgui.ImGuiKey_Space
	KeyApostrophe   Key = cimgui.ImGuiKey_Apostrophe
	KeyComma        Key = cimgui.ImGuiKey_Comma
	KeyMinus        Key = cimgui.ImGuiKey_Minus
	KeyPeriod       Key = cimgui.ImGuiKey_Period
	KeySlash        Key = cimgui.ImGuiKey_Slash
	Key0            Key = cimgui.ImGuiKey_0
	Key1            Key = cimgui.ImGuiKey_1
	Key2            Key = cimgui.ImGuiKey_2
	Key3            Key = cimgui.ImGuiKey_3
	Key4            Key = cimgui.ImGuiKey_4
	Key5            Key = cimgui.ImGuiKey_5
	Key6            Key = cimgui.ImGuiKey_6
	Key7            Key = cimgui.ImGuiKey_7
	Key8            Key = cimgui.ImGuiKey_8
	Key9            Key = cimgui.ImGuiKey_9
	KeySemicolon    Key = cimgui.ImGuiKey_Semicolon
	KeyEqual        Key = cimgui.ImGuiKey_Equal
	KeyA            Key = cimgui.ImGuiKey_A
	KeyB            Key = cimgui.ImGuiKey_B
	KeyC            Key = cimgui.ImGuiKey_C
	KeyD            Key = cimgui.ImGuiKey_D
	KeyE            Key = cimgui.ImGuiKey_E
	KeyF            Key = cimgui.ImGuiKey_F
	KeyG            Key = cimgui.ImGuiKey_G
	KeyH            Key = cimgui.ImGuiKey_H
	KeyI            Key = cimgui.ImGuiKey_I
	KeyJ            Key = cimgui.ImGuiKey_J
	KeyK            Key = cimgui.ImGuiKey_K
	KeyL            Key = cimgui.ImGuiKey_L
	KeyM            Key = cimgui.ImGuiKey_M
	KeyN            Key = cimgui.ImGuiKey_N
	KeyO            Key = cimgui.ImGuiKey_O
	KeyP            Key = cimgui.ImGuiKey_P
	KeyQ            Key = cimgui.ImGuiKey_Q
	KeyR            Key = cimgui.ImGuiKey_R
	KeyS            Key = cimgui.ImGuiKey_S
	KeyT            Key = cimgui.ImGuiKey_T
	KeyU            Key = cimgui.ImGuiKey_U
	KeyV            Key = cimgui.ImGuiKey_V
	KeyW            Key = cimgui.ImGuiKey_W
	KeyX            Key = cimgui.ImGuiKey_X
	KeyY            Key = cimgui.ImGuiKey_Y
	KeyZ            Key = cimgui.ImGuiKey_Z
	KeyLeftBracket  Key = cimgui.ImGuiKey_LeftBracket
	KeyBackslash    Key = cimgui.ImGuiKey_Backslash
	KeyRightBracket Key = cimgui.ImGuiKey_RightBracket
	KeyGraveAccent  Key = cimgui.ImGuiKey_GraveAccent
	// KeyWorld1       Key = cimgui.ImGuiKey_World1
	// KeyWorld2       Key = cimgui.ImGuiKey_World2
	KeyEscape      Key = cimgui.ImGuiKey_Escape
	KeyEnter       Key = cimgui.ImGuiKey_Enter
	KeyTab         Key = cimgui.ImGuiKey_Tab
	KeyBackspace   Key = cimgui.ImGuiKey_Backspace
	KeyInsert      Key = cimgui.ImGuiKey_Insert
	KeyDelete      Key = cimgui.ImGuiKey_Delete
	KeyRight       Key = cimgui.ImGuiKey_RightArrow
	KeyLeft        Key = cimgui.ImGuiKey_LeftArrow
	KeyDown        Key = cimgui.ImGuiKey_DownArrow
	KeyUp          Key = cimgui.ImGuiKey_UpArrow
	KeyPageUp      Key = cimgui.ImGuiKey_PageUp
	KeyPageDown    Key = cimgui.ImGuiKey_PageDown
	KeyHome        Key = cimgui.ImGuiKey_Home
	KeyEnd         Key = cimgui.ImGuiKey_End
	KeyCapsLock    Key = cimgui.ImGuiKey_CapsLock
	KeyScrollLock  Key = cimgui.ImGuiKey_ScrollLock
	KeyNumLock     Key = cimgui.ImGuiKey_NumLock
	KeyPrintScreen Key = cimgui.ImGuiKey_PrintScreen
	KeyPause       Key = cimgui.ImGuiKey_Pause
	KeyF1          Key = cimgui.ImGuiKey_F1
	KeyF2          Key = cimgui.ImGuiKey_F2
	KeyF3          Key = cimgui.ImGuiKey_F3
	KeyF4          Key = cimgui.ImGuiKey_F4
	KeyF5          Key = cimgui.ImGuiKey_F5
	KeyF6          Key = cimgui.ImGuiKey_F6
	KeyF7          Key = cimgui.ImGuiKey_F7
	KeyF8          Key = cimgui.ImGuiKey_F8
	KeyF9          Key = cimgui.ImGuiKey_F9
	KeyF10         Key = cimgui.ImGuiKey_F10
	KeyF11         Key = cimgui.ImGuiKey_F11
	KeyF12         Key = cimgui.ImGuiKey_F12
	// KeyF13          Key = cimgui.ImGuiKey_F13
	// KeyF14          Key = cimgui.ImGuiKey_F14
	// KeyF15          Key = cimgui.ImGuiKey_F15
	// KeyF16          Key = cimgui.ImGuiKey_F16
	// KeyF17          Key = cimgui.ImGuiKey_F17
	// KeyF18          Key = cimgui.ImGuiKey_F18
	// KeyF19          Key = cimgui.ImGuiKey_F19
	// KeyF20          Key = cimgui.ImGuiKey_F20
	// KeyF21          Key = cimgui.ImGuiKey_F21
	// KeyF22          Key = cimgui.ImGuiKey_F22
	// KeyF23          Key = cimgui.ImGuiKey_F23
	// KeyF24          Key = cimgui.ImGuiKey_F24
	// KeyF25          Key = cimgui.ImGuiKey_F25
	// KeyKP0          Key = cimgui.ImGuiKey_KP0
	// KeyKP1          Key = cimgui.ImGuiKey_KP1
	// KeyKP2          Key = cimgui.ImGuiKey_KP2
	// KeyKP3          Key = cimgui.ImGuiKey_KP3
	// KeyKP4          Key = cimgui.ImGuiKey_KP4
	// KeyKP5          Key = cimgui.ImGuiKey_KP5
	// KeyKP6          Key = cimgui.ImGuiKey_KP6
	// KeyKP7          Key = cimgui.ImGuiKey_KP7
	// KeyKP8          Key = cimgui.ImGuiKey_KP8
	// KeyKP9          Key = cimgui.ImGuiKey_KP9
	// KeyKPDecimal    Key = cimgui.ImGuiKey_KPDecimal
	// KeyKPDivide     Key = cimgui.ImGuiKey_KPDivide
	// KeyKPMultiply   Key = cimgui.ImGuiKey_KPMultiply
	// KeyKPSubtract   Key = cimgui.ImGuiKey_KPSubtract
	// KeyKPAdd        Key = cimgui.ImGuiKey_KPAdd
	// KeyKPEnter      Key = cimgui.ImGuiKey_KPEnter
	// KeyKPEqual      Key = cimgui.ImGuiKey_KPEqual
	KeyLeftShift    Key = cimgui.ImGuiKey_LeftShift
	KeyLeftControl  Key = cimgui.ImGuiKey_LeftCtrl
	KeyLeftAlt      Key = cimgui.ImGuiKey_LeftAlt
	KeyLeftSuper    Key = cimgui.ImGuiKey_LeftSuper
	KeyRightShift   Key = cimgui.ImGuiKey_RightShift
	KeyRightControl Key = cimgui.ImGuiKey_RightCtrl
	KeyRightAlt     Key = cimgui.ImGuiKey_RightAlt
	KeyRightSuper   Key = cimgui.ImGuiKey_RightSuper
	KeyMenu         Key = cimgui.ImGuiKey_Menu
	// KeyLast         Key = cimgui.ImGuiKey_Last
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
