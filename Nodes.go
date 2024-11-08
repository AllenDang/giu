package giu

import (
	"fmt"

	"github.com/AllenDang/cimgui-go/imnodes"
)

type nodeEditorState struct {
	context      *imnodes.EditorContext
	links        []*LinkWidget
	linksCounter int32
}

func newNodeEditorState() *nodeEditorState {
	return &nodeEditorState{
		context: imnodes.EditorContextCreate(),
		links:   make([]*LinkWidget, 0),
	}
}

// Dispose implements Disposable interface.
func (n *nodeEditorState) Dispose() {
	imnodes.EditorContextFree(n.context)
}

func (n *NodeEditorWidget) getState() *nodeEditorState {
	if state := GetState[nodeEditorState](Context, n.id); state != nil {
		return state
	}

	state := newNodeEditorState()
	SetState[nodeEditorState](Context, n.id, state)

	return state
}

// NodeElementType represents a type of node element.
type NodeElementType int

const (
	// NodeElementInput describes a lyout associated with an input "point".
	NodeElementInput NodeElementType = iota
	// NodeElementOutput describes a lyout associated with an output "point".
	NodeElementOutput
	// NodeElementBody describes just a plain layout.
	NodeElementBody
	// NodeElementTitleBar should be rendered before any other element.
	// It describes a layout of a node title bar.
	NodeElementTitleBar
)

type nodeElement struct {
	elementType NodeElementType
	layout      Layout
	id          int32
}

type NodeEditorWidget struct {
	nodes     []*NodeWidget
	idCounter int32
	id        ID
}

func NodeEditor() *NodeEditorWidget {
	return &NodeEditorWidget{
		idCounter: 0,
		id:        GenAutoID("NodeEditor"),
	}
}

func (n *NodeEditorWidget) Nodes(nodes ...*NodeWidget) *NodeEditorWidget {
	n.nodes = nodes

	return n
}

func (n *NodeEditorWidget) Build() {
	n.idCounter = 0
	state := n.getState()
	imnodes.EditorContextSet(state.context)

	imnodes.BeginNodeEditor()
	for _, node := range n.nodes {
		node.BuildNode(&n.idCounter)
	}

	for _, link := range state.links {
		imnodes.Link(link.linkID, link.startID, link.endID)
	}

	imnodes.EndNodeEditor()

	potentialNewLink := Link(state.linksCounter, 0, 0)
	if imnodes.IsLinkCreatedBoolPtr(&potentialNewLink.startID, &potentialNewLink.endID) {
		state.links = append(state.links, potentialNewLink)
		state.linksCounter++
	}

	var id int32
	if imnodes.IsLinkDestroyed(&id) {
		for i, link := range state.links {
			if link.linkID == id {
				state.links = append(state.links[:i], state.links[i+1:]...)
				break
			}
		}
	}
}

type NodeWidget struct {
	elements []nodeElement
}

func Node() *NodeWidget {
	return &NodeWidget{}
}

func (n *NodeWidget) Static(widgets ...Widget) *NodeWidget {
	n.elements = append(n.elements, nodeElement{NodeElementBody, Layout(widgets)})

	return n
}

func (n *NodeWidget) TitleBar(widgets ...Widget) *NodeWidget {
	for i := range n.elements {
		if n.elements[i].elementType != NodeElementTitleBar {
			n.elements = append(n.elements[:i], append([]nodeElement{{NodeElementTitleBar, Layout(widgets)}}, n.elements[i:]...)...)
		}
	}

	return n
}

func (n *NodeWidget) Input(widgets ...Widget) *NodeWidget {
	n.elements = append(n.elements, nodeElement{NodeElementInput, Layout(widgets)})

	return n
}

func (n *NodeWidget) Output(widgets ...Widget) *NodeWidget {
	n.elements = append(n.elements, nodeElement{NodeElementOutput, Layout(widgets)})

	return n
}

// ElementID allows you to manually specify an ID for the last added element.
// Appliable only for NodeElementInput and NodeElementOutput.
func (n *NodeWidget) ElementID(id int32) *NodeWidget {
	n.elements[len(n.elements)-1].id = id

	return n
}

func (n *NodeWidget) BuildNode(idCounter *int32) {
	fMap := map[NodeElementType]struct {
		begin func(int32)
		end   func()
	}{
		NodeElementInput:    {imnodes.BeginInputAttribute, imnodes.EndInputAttribute},
		NodeElementOutput:   {imnodes.BeginOutputAttribute, imnodes.EndOutputAttribute},
		NodeElementBody:     {imnodes.BeginStaticAttribute, imnodes.EndStaticAttribute},
		NodeElementTitleBar: {func(int32) { imnodes.BeginNodeTitleBar() }, imnodes.EndNodeTitleBar},
	}

	// Assert(n.layout != nil && len(n.layout) > 0, "NodeWidget", "BuildNode", "Node layout is required")
	imnodes.BeginNode(*idCounter)
	*idCounter++

	for _, element := range n.elements {
		if element.layout != nil {
			f, ok := fMap[element.elementType]
			if !ok {
				panic(fmt.Sprintf("NodeWidget:BuildNode: Unknown node element type", element.elementType))
			}
			f.begin(*idCounter)
			*idCounter++
			element.layout.Build()
			f.end()
		}
	}

	imnodes.EndNode()
}

type LinkWidget struct {
	linkID  int32
	startID int32
	endID   int32
}

func Link(linkID, startID, endID int32) *LinkWidget {
	return &LinkWidget{
		startID: startID,
		endID:   endID,
		linkID:  linkID,
	}
}
