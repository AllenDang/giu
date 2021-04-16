package imgui

import "C"
import (
	"errors"
	"fmt"
)

// AssertHandler is a handler for an assertion that happened in the native part of ImGui.
type AssertHandler func(expression string, file string, line int)

var assertHandler AssertHandler = func(expression string, file string, line int) {
	message := fmt.Sprintf(`Assertion failed!
File: %s, Line %d

Expression: %s
`, file, line, expression)
	panic(errors.New(message))
}

// SetAssertHandler registers a handler function for all future assertions.
// Setting nil will disable special handling.
// The default handler panics.
func SetAssertHandler(handler AssertHandler) {
	assertHandler = handler
}

//export iggAssert
func iggAssert(result C.int, expression *C.char, file *C.char, line C.int) {
	if (result == 0) && (assertHandler != nil) {
		assertHandler(C.GoString(expression), C.GoString(file), int(line))
	}
}
