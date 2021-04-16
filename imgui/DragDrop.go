package imgui

// #include "DragDropWrapper.h"
import "C"
import (
	"unsafe"
)

type Payload uintptr

func (p Payload) handle() C.IggPayload {
	return C.IggPayload(p)
}

func (p Payload) Data() int {
	raw := C.iggPayloadGetData(p.handle())
	return *(*int)((unsafe.Pointer)(raw))
}

func BeginDragDropSource() bool {
	return BeginDragDropSourceV(0)
}

func BeginDragDropSourceV(flags int) bool {
	return C.iggBeginDragDropSource(C.int(flags)) != 0
}

func SetDragDropPayload(payloadType string, data int) bool {
	return SetDragDropPayloadV(payloadType, data, 0)
}

func SetDragDropPayloadV(payloadType string, data int, cond int) bool {
	typeArg, typeFn := wrapString(payloadType)
	defer typeFn()

	return C.iggSetDragDropPayload(typeArg, unsafe.Pointer(&data), C.uint(unsafe.Sizeof(&data)), C.int(cond)) != 0
}

func EndDragDropSource() {
	C.iggEndDragDropSource()
}

func BeginDragDropTarget() bool {
	return C.iggBeginDragDropTarget() != 0
}

func AcceptDragDropPayload(payloadType string) Payload {
	return AcceptDragDropPayloadV(payloadType, 0)
}

func AcceptDragDropPayloadV(payloadType string, flags int) Payload {
	typeArg, typeFn := wrapString(payloadType)
	defer typeFn()

	payload := C.iggAcceptDragDropPayload(typeArg, C.int(flags))
	return Payload(payload)
}

func EndDragDropTarget() {
	C.iggEndDragDropTarget()
}
