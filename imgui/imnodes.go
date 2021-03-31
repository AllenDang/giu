package imgui

// #cgo CXXFLAGS: -std=c++11
// #include "imnodesWrapper.h"
import "C"

func ImNodesCreateContext() {
	C.iggImNodesCreateContext()
}

func ImNodesDestroyContext() {
	C.iggImNodesDestroyContext()
}

func ImNodesBeginNodeEditor() {
	C.iggImNodesBeginNodeEditor()
}

func ImNodesEndNodeEditor() {
	C.iggImNodesEndNodeEditor()
}

func ImNodesBeginNode(id int) {
	C.iggImNodesBeginNode(C.int(id))
}

func ImNodesEndNode() {
	C.iggImNodesEndNode()
}

func ImNodesBeginNodeTitleBar() {
	C.iggImNodesBeginNodeTitleBar()
}

func ImNodesEndNodeTitleBar() {
	C.iggImNodesEndNodeTitleBar()
}

func ImNodesBeginInputAttribute(id int) {
	C.iggImNodesBeginInputAttribute(C.int(id))
}

func ImNodesEndInputAttribute() {
	C.iggImNodesEndInputAttribute()
}

func ImNodesBeginOutputAttribute(id int) {
	C.iggImNodesBeginOutputAttribute(C.int(id))
}

func ImNodesEndOutputAttribute() {
	C.iggImNodesEndOutputAttribute()
}

func ImNodesLink(id, startAttributeId, endAttributeId int) {
	C.iggImNodesLink(C.int(id), C.int(startAttributeId), C.int(endAttributeId))
}

func ImNodesIsLinkCreated() (startedNodeId, startedAttributeId, endedNodeId, endedAttributeId int32, createdFromSnap, isLinkCreated bool) {
	sNodeIdArg, sNodeIdDeleter := wrapInt32(&startedNodeId)
	defer sNodeIdDeleter()

	sAttrIdArg, sAttrIdDeleter := wrapInt32(&startedAttributeId)
	defer sAttrIdDeleter()

	eNodeIdArg, eNodeIdDeleter := wrapInt32(&endedNodeId)
	defer eNodeIdDeleter()

	eAttrIdArg, eAttrIdDeleter := wrapInt32(&endedAttributeId)
	defer eAttrIdDeleter()

	bArg, bArgDeleter := wrapBool(&createdFromSnap)
	defer bArgDeleter()

	isLinkCreated = C.iggImNodesIsLinkCreated(sNodeIdArg, sAttrIdArg, eNodeIdArg, eAttrIdArg, bArg) != 0

	return
}

func ImNodesIsLinkDestroyed() (linkId int32, isDestroyed bool) {
	linkIdArg, linkIdDeleter := wrapInt32(&linkId)
	defer linkIdDeleter()

	isDestroyed = C.iggImNodesIsLinkDestroyed(linkIdArg) != 0

	return
}

type ImNodesAttributeFlags int

const (
	AttributeFlags_None ImNodesAttributeFlags = 0
	// Allow detaching a link by left-clicking and dragging the link at a pin it is connected to.
	// NOTE: the user has to actually delete the link for this to work. A deleted link can be
	// detected by calling IsLinkDestroyed() after EndNodeEditor().
	AttributeFlags_EnableLinkDetachWithDragClick ImNodesAttributeFlags = 1 << 0
	// Visual snapping of an in progress link will trigger IsLink Created/Destroyed events. Allows
	// for previewing the creation of a link while dragging it across attributes. See here for demo:
	// https://github.com/Nelarius/imnodes/issues/41#issuecomment-647132113 NOTE: the user has to
	// actually delete the link for this to work. A deleted link can be detected by calling
	// IsLinkDestroyed() after EndNodeEditor().
	AttributeFlags_EnableLinkCreationOnSnap ImNodesAttributeFlags = 1 << 1
)

func ImNodesPushAttributeFlag(flag ImNodesAttributeFlags) {
	C.iggImNodesPushAttributeFlag(C.int(flag))
}

func ImNodesPopAttributeFlag() {
	C.iggImNodesPopAttributeFlag()
}

func ImNodesEnableDetachWithCtrlClick() {
	C.iggImNodesEnableDetachWithCtrlClick()
}

func ImNodesSetNodeScreenSpacePos(nodeId int, pos Vec2) {
	posArg, _ := pos.wrapped()
	C.iggImNodesSetNodeScreenSpacePos(C.int(nodeId), posArg)
}

func ImNodesSetNodeEditorSpacePos(nodeId int, pos Vec2) {
	posArg, _ := pos.wrapped()
	C.iggImNodesSetNodeEditorSpacePos(C.int(nodeId), posArg)
}

func ImNodesSetNodeGridSpacePos(nodeId int, pos Vec2) {
	posArg, _ := pos.wrapped()
	C.iggImNodesSetNodeGridSpacePos(C.int(nodeId), posArg)
}

func ImNodesGetNodeScreenSpacePos(nodeId int) *Vec2 {
	var pos Vec2
	valueArg, valueDeleter := pos.wrapped()
	defer valueDeleter()

	C.iggImNodesGetNodeScreenSpacePos(C.int(nodeId), valueArg)
	return &pos
}

func ImNodesGetNodeEditorSpacePos(nodeId int) *Vec2 {
	var pos Vec2
	valueArg, valueDeleter := pos.wrapped()
	defer valueDeleter()

	C.iggImNodesGetNodeEditorSpacePos(C.int(nodeId), valueArg)
	return &pos
}

func ImNodesGetNodeGridSpacePos(nodeId int) *Vec2 {
	var pos Vec2
	valueArg, valueDeleter := pos.wrapped()
	defer valueDeleter()

	C.iggImNodesGetNodeGridSpacePos(C.int(nodeId), valueArg)
	return &pos
}
