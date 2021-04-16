package imgui

// Constants to fill IO.KeyMap() lookup with indices into the IO.KeysDown[512] array.
// The mapped indices are then the ones reported to IO.KeyPress() and IO.KeyRelease().
const (
	KeyTab         = 0
	KeyLeftArrow   = 1
	KeyRightArrow  = 2
	KeyUpArrow     = 3
	KeyDownArrow   = 4
	KeyPageUp      = 5
	KeyPageDown    = 6
	KeyHome        = 7
	KeyEnd         = 8
	KeyInsert      = 9
	KeyDelete      = 10
	KeyBackspace   = 11
	KeySpace       = 12
	KeyEnter       = 13
	KeyEscape      = 14
	KeyKeyPadEnter = 15
	KeyA           = 16 // for text edit CTRL+A: select all
	KeyC           = 17 // for text edit CTRL+C: copy
	KeyV           = 18 // for text edit CTRL+V: paste
	KeyX           = 19 // for text edit CTRL+X: cut
	KeyY           = 20 // for text edit CTRL+Y: redo
	KeyZ           = 21 // for text edit CTRL+Z: undo
)
