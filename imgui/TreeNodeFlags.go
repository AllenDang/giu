package imgui

const (
	// TreeNodeFlagsNone default = 0
	TreeNodeFlagsNone = 0
	// TreeNodeFlagsSelected draws as selected.
	TreeNodeFlagsSelected = 1 << 0
	// TreeNodeFlagsFramed draws full colored frame (e.g. for CollapsingHeader).
	TreeNodeFlagsFramed = 1 << 1
	// TreeNodeFlagsAllowItemOverlap hit testing to allow subsequent widgets to overlap this one.
	TreeNodeFlagsAllowItemOverlap = 1 << 2
	// TreeNodeFlagsNoTreePushOnOpen doesn't do a TreePush() when open
	// (e.g. for CollapsingHeader) = no extra indent nor pushing on ID stack.
	TreeNodeFlagsNoTreePushOnOpen = 1 << 3
	// TreeNodeFlagsNoAutoOpenOnLog doesn't automatically and temporarily open node when Logging is active
	// (by default logging will automatically open tree nodes).
	TreeNodeFlagsNoAutoOpenOnLog = 1 << 4
	// TreeNodeFlagsDefaultOpen defaults node to be open.
	TreeNodeFlagsDefaultOpen = 1 << 5
	// TreeNodeFlagsOpenOnDoubleClick needs double-click to open node.
	TreeNodeFlagsOpenOnDoubleClick = 1 << 6
	// TreeNodeFlagsOpenOnArrow opens only when clicking on the arrow part.
	// If TreeNodeFlagsOpenOnDoubleClick is also set, single-click arrow or double-click all box to open.
	TreeNodeFlagsOpenOnArrow = 1 << 7
	// TreeNodeFlagsLeaf allows no collapsing, no arrow (use as a convenience for leaf nodes).
	TreeNodeFlagsLeaf = 1 << 8
	// TreeNodeFlagsBullet displays a bullet instead of an arrow.
	TreeNodeFlagsBullet = 1 << 9
	// TreeNodeFlagsFramePadding uses FramePadding (even for an unframed text node) to
	// vertically align text baseline to regular widget height. Equivalent to calling AlignTextToFramePadding().
	TreeNodeFlagsFramePadding = 1 << 10
	// TreeNodeFlagsSpanAvailWidth extends hit box to the right-most edge, even if not framed.
	// This is not the default in order to allow adding other items on the same line.
	// In the future we may refactor the hit system to be front-to-back, allowing natural overlaps
	// and then this can become the default.
	TreeNodeFlagsSpanAvailWidth = 1 << 11
	// TreeNodeFlagsSpanFullWidth extends hit box to the left-most and right-most edges (bypass the indented area).
	TreeNodeFlagsSpanFullWidth = 1 << 12
	// TreeNodeFlagsNavLeftJumpsBackHere (WIP) Nav: left direction may move to this TreeNode() from any of its child
	// (items submitted between TreeNode and TreePop)
	TreeNodeFlagsNavLeftJumpsBackHere = 1 << 13
	// TreeNodeFlagsCollapsingHeader combines TreeNodeFlagsFramed and TreeNodeFlagsNoAutoOpenOnLog.
	TreeNodeFlagsCollapsingHeader = TreeNodeFlagsFramed | TreeNodeFlagsNoTreePushOnOpen | TreeNodeFlagsNoAutoOpenOnLog
)
