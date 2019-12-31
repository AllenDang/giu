package giu

import "github.com/AllenDang/giu/imgui"

type TreeNodeWidget struct {
	BaseWidget
	label  string
	flags  int
	layout Layout
}

func TreeNodeV(label string, flags int, layout Layout) *TreeNodeWidget {
	return &TreeNodeWidget{
		BaseWidget: BaseWidget{width: 0},
		label:      label,
		flags:      flags,
		layout:     layout,
	}
}

func TreeNode(label string, layout Layout) *TreeNodeWidget {
	return TreeNodeV(label, 0, layout)
}

func (t *TreeNodeWidget) Build() {
	if imgui.TreeNodeV(t.label, t.flags) {
		t.layout.Build()

		if (t.flags & imgui.TreeNodeFlagsNoTreePushOnOpen) == 0 {
			imgui.TreePop()
		}
	}
}
