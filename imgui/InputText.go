package imgui

// #include "InputTextCallbackDataWrapper.h"
import "C"
import (
	"sync"
	"unsafe"
)

const (
	// InputTextFlagsNone sets everything default.
	InputTextFlagsNone = 0
	// InputTextFlagsCharsDecimal allows 0123456789.+-
	InputTextFlagsCharsDecimal = 1 << 0
	// InputTextFlagsCharsHexadecimal allow 0123456789ABCDEFabcdef
	InputTextFlagsCharsHexadecimal = 1 << 1
	// InputTextFlagsCharsUppercase turns a..z into A..Z.
	InputTextFlagsCharsUppercase = 1 << 2
	// InputTextFlagsCharsNoBlank filters out spaces, tabs.
	InputTextFlagsCharsNoBlank = 1 << 3
	// InputTextFlagsAutoSelectAll selects entire text when first taking mouse focus.
	InputTextFlagsAutoSelectAll = 1 << 4
	// InputTextFlagsEnterReturnsTrue returns 'true' when Enter is pressed (as opposed to when the value was modified).
	InputTextFlagsEnterReturnsTrue = 1 << 5
	// InputTextFlagsCallbackCompletion for callback on pressing TAB (for completion handling).
	InputTextFlagsCallbackCompletion = 1 << 6
	// InputTextFlagsCallbackHistory for callback on pressing Up/Down arrows (for history handling).
	InputTextFlagsCallbackHistory = 1 << 7
	// InputTextFlagsCallbackAlways for callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackAlways = 1 << 8
	// InputTextFlagsCallbackCharFilter for callback on character inputs to replace or discard them.
	// Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsCallbackCharFilter = 1 << 9
	// InputTextFlagsAllowTabInput when pressing TAB to input a '\t' character into the text field.
	InputTextFlagsAllowTabInput = 1 << 10
	// InputTextFlagsCtrlEnterForNewLine in multi-line mode, unfocus with Enter, add new line with Ctrl+Enter
	// (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsCtrlEnterForNewLine = 1 << 11
	// InputTextFlagsNoHorizontalScroll disables following the cursor horizontally.
	InputTextFlagsNoHorizontalScroll = 1 << 12
	// InputTextFlagsAlwaysInsertMode sets insert mode.
	InputTextFlagsAlwaysInsertMode = 1 << 13
	// InputTextFlagsReadOnly sets read-only mode.
	InputTextFlagsReadOnly = 1 << 14
	// InputTextFlagsPassword sets password mode, display all characters as '*'.
	InputTextFlagsPassword = 1 << 15
	// InputTextFlagsNoUndoRedo disables undo/redo. Note that input text owns the text data while active,
	// if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsNoUndoRedo = 1 << 16
	// InputTextFlagsCharsScientific allows 0123456789.+-*/eE (Scientific notation input).
	InputTextFlagsCharsScientific = 1 << 17
	// inputTextFlagsCallbackResize for callback on buffer capacity change requests.
	inputTextFlagsCallbackResize = 1 << 18
)

// InputTextCallback is called for sharing state of an input field.
// By default, the callback should return 0.
type InputTextCallback func(InputTextCallbackData) int32

type inputTextState struct {
	buf *stringBuffer

	key      C.int
	callback InputTextCallback
}

var inputTextStates = make(map[C.int]*inputTextState)
var inputTextStatesMutex sync.Mutex

func newInputTextState(text string, cb InputTextCallback) *inputTextState {
	state := &inputTextState{}
	state.buf = newStringBuffer(text)
	state.callback = cb
	state.register()
	return state
}

func (state *inputTextState) register() {
	inputTextStatesMutex.Lock()
	defer inputTextStatesMutex.Unlock()
	key := C.int(len(inputTextStates) + 1)
	for _, existing := inputTextStates[key]; existing; _, existing = inputTextStates[key] {
		key++
	}
	state.key = key
	inputTextStates[key] = state
}

func (state *inputTextState) release() {
	state.buf.free()

	if state.key != 0 {
		inputTextStatesMutex.Lock()
		defer inputTextStatesMutex.Unlock()
		delete(inputTextStates, state.key)
	}
}

func (state *inputTextState) onCallback(handle C.IggInputTextCallbackData) C.int {
	data := InputTextCallbackData{state: state, handle: handle}
	if data.EventFlag() == inputTextFlagsCallbackResize {
		state.buf.resizeTo(data.bufSize())
		data.setBuf(state.buf.ptr, state.buf.size, data.bufTextLen())
		return 0
	}
	if state.callback == nil {
		return 0
	}
	return C.int(state.callback(data))
}

//export iggInputTextCallback
func iggInputTextCallback(handle C.IggInputTextCallbackData, key C.int) C.int {
	state := iggInputTextStateFor(key)
	return state.onCallback(handle)
}

func iggInputTextStateFor(key C.int) *inputTextState {
	inputTextStatesMutex.Lock()
	defer inputTextStatesMutex.Unlock()
	return inputTextStates[key]
}

// InputTextCallbackData represents the shared state of InputText(), passed as an argument to your callback.
type InputTextCallbackData struct {
	state  *inputTextState
	handle C.IggInputTextCallbackData
}

// EventFlag returns one of the InputTextFlagsCallback* constants to indicate the nature of the callback.
func (data InputTextCallbackData) EventFlag() int {
	return int(C.iggInputTextCallbackDataGetEventFlag(data.handle))
}

// Flags returns the set of flags that the user originally passed to InputText.
func (data InputTextCallbackData) Flags() int {
	return int(C.iggInputTextCallbackDataGetFlags(data.handle)) & ^inputTextFlagsCallbackResize
}

// EventChar returns the current character input. Only valid during CharFilter callback.
func (data InputTextCallbackData) EventChar() rune {
	return rune(C.iggInputTextCallbackDataGetEventChar(data.handle))
}

// SetEventChar overrides what the user entered. Set to zero do drop the current input.
// Returning 1 from the callback also drops the current input.
// Only valid during CharFilter callback.
//
// Note: The internal representation of characters is based on uint16, so less than rune would provide.
func (data InputTextCallbackData) SetEventChar(value rune) {
	C.iggInputTextCallbackDataSetEventChar(data.handle, C.ushort(value))
}

// EventKey returns the currently pressed key. Valid for completion and history callbacks.
func (data InputTextCallbackData) EventKey() int {
	return int(C.iggInputTextCallbackDataGetEventKey(data.handle))
}

// Buffer returns a view into the current UTF-8 buffer.
// Only during the callbacks of [Completion,History,Always] the current buffer is returned.
// The returned slice is a temporary view into the underlying raw buffer. Do not keep it!
// The underlying memory allocation may even change through a call to InsertBytes().
//
// You may change the buffer through the following ways:
// If the new text has a different (encoded) length, use the functions InsertBytes() and/or DeleteBytes().
// Otherwise you may keep the buffer as is and modify the bytes. If you change the buffer this way directly, mark the buffer
// as modified with MarkBufferModified().
func (data InputTextCallbackData) Buffer() []byte {
	ptr := C.iggInputTextCallbackDataGetBuf(data.handle)
	if ptr == nil {
		return nil
	}
	textLen := data.bufTextLen()
	return ((*[1 << 30]byte)(unsafe.Pointer(ptr)))[:textLen]
}

// MarkBufferModified indicates that the content of the buffer was modified during a callback.
// Only considered during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) MarkBufferModified() {
	C.iggInputTextCallbackDataMarkBufferModified(data.handle)
}

func (data InputTextCallbackData) setBuf(buf unsafe.Pointer, size, textLen int) {
	C.iggInputTextCallbackDataSetBuf(data.handle, (*C.char)(buf), C.int(size), C.int(textLen))
}

func (data InputTextCallbackData) bufSize() int {
	return int(C.iggInputTextCallbackDataGetBufSize(data.handle))
}

func (data InputTextCallbackData) bufTextLen() int {
	return int(C.iggInputTextCallbackDataGetBufTextLen(data.handle))
}

// DeleteBytes removes the given count of bytes starting at the specified byte offset within the buffer.
// This function can be called during the [Completion,History,Always] callbacks.
// Clears the current selection.
//
// This function ignores the deletion beyond the current buffer length.
// Calling with negative offset or count arguments will panic.
func (data InputTextCallbackData) DeleteBytes(offset, count int) {
	if offset < 0 {
		panic("invalid offset")
	}
	if count < 0 {
		panic("invalid count")
	}
	textLen := data.bufTextLen()
	if offset >= textLen {
		return
	}
	toRemove := count
	available := textLen - offset
	if toRemove > available {
		toRemove = available
	}
	C.iggInputTextCallbackDataDeleteBytes(data.handle, C.int(offset), C.int(toRemove))
}

// InsertBytes inserts the given bytes at given byte offset into the buffer.
// Calling this function may change the underlying buffer allocation.
//
// This function can be called during the [Completion,History,Always] callbacks.
// Clears the current selection.
//
// Calling with an offset outside of the range of the buffer will panic.
func (data InputTextCallbackData) InsertBytes(offset int, bytes []byte) {
	if (offset < 0) || (offset > data.bufTextLen()) {
		panic("invalid offset")
	}
	var bytesPtr *C.char
	byteCount := len(bytes)
	if byteCount > 0 {
		bytesPtr = (*C.char)(unsafe.Pointer(&bytes[0]))
		C.iggInputTextCallbackDataInsertBytes(data.handle, C.int(offset), bytesPtr, C.int(byteCount))
	}
}

// CursorPos returns the byte-offset of the cursor within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) CursorPos() int {
	return int(C.iggInputTextCallbackDataGetCursorPos(data.handle))
}

// SetCursorPos changes the current byte-offset of the cursor within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) SetCursorPos(value int) {
	C.iggInputTextCallbackDataSetCursorPos(data.handle, C.int(value))
}

// SelectionStart returns the byte-offset of the selection start within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) SelectionStart() int {
	return int(C.iggInputTextCallbackDataGetSelectionStart(data.handle))
}

// SetSelectionStart changes the current byte-offset of the selection start within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) SetSelectionStart(value int) {
	C.iggInputTextCallbackDataSetSelectionStart(data.handle, C.int(value))
}

// SelectionEnd returns the byte-offset of the selection end within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) SelectionEnd() int {
	return int(C.iggInputTextCallbackDataGetSelectionEnd(data.handle))
}

// SetSelectionEnd changes the current byte-offset of the selection end within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) SetSelectionEnd(value int) {
	C.iggInputTextCallbackDataSetSelectionEnd(data.handle, C.int(value))
}
