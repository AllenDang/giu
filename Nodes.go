package giu

import "github.com/AllenDang/cimgui-go/imnodes"

type NodeEditorWidget struct{}

func NodeEditor() *NodeEditorWidget {
	return &NodeEditorWidget{}
}

func (n *NodeEditorWidget) Build() {
	imnodes.BeginNodeEditor()
	imnodes.EndNodeEditor()
}
