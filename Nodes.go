package giu

import (
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
	titleBar Layout
	layout   Layout
	input    Layout
	output   Layout
}

func Node(staticLayout ...Widget) *NodeWidget {
	return &NodeWidget{
		layout: Layout(staticLayout),
	}
}

func (n *NodeWidget) TitleBar(widgets ...Widget) *NodeWidget {
	n.titleBar = Layout(widgets)

	return n
}

func (n *NodeWidget) Input(widgets ...Widget) *NodeWidget {
	n.input = Layout(widgets)

	return n
}

func (n *NodeWidget) Output(widgets ...Widget) *NodeWidget {
	n.output = Layout(widgets)

	return n
}

func (n *NodeWidget) BuildNode(idCounter *int32) {
	Assert(n.layout != nil && len(n.layout) > 0, "NodeWidget", "BuildNode", "Node layout is required")
	imnodes.BeginNode(*idCounter)
	*idCounter++

	if n.titleBar != nil {
		imnodes.BeginNodeTitleBar()
		n.titleBar.Build()
		imnodes.EndNodeTitleBar()
	}

	if n.input != nil {
		imnodes.PushAttributeFlag(imnodes.AttributeFlagsEnableLinkDetachWithDragClick)
		imnodes.BeginInputAttribute(*idCounter)
		*idCounter++
		n.input.Build()
		imnodes.EndInputAttribute()
		imnodes.PopAttributeFlag()
	}

	imnodes.BeginStaticAttribute(*idCounter)
	*idCounter++
	n.layout.Build()
	imnodes.EndStaticAttribute()

	if n.output != nil {
		imnodes.BeginOutputAttribute(*idCounter)
		*idCounter++
		n.output.Build()
		imnodes.EndOutputAttribute()
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
