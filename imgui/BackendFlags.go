package imgui

const (
	// BackendFlagNone default = 0
	BackendFlagNone = 0
	// BackendFlagHasGamepad back-end Platform supports gamepad and currently has one connected.
	BackendFlagHasGamepad = 1 << 0
	// BackendFlagHasMouseCursors back-end Platform supports honoring GetMouseCursor() value to change the OS cursor
	// shape.
	BackendFlagHasMouseCursors = 1 << 1
	// BackendFlagHasSetMousePos back-end Platform supports io.WantSetMousePos requests to reposition the OS mouse
	// position (only used if ImGuiConfigFlags_NavEnableSetMousePos is set).
	BackendFlagHasSetMousePos = 1 << 2
	// BackendFlagsRendererHasVtxOffset back-end Renderer supports ImDrawCmd::VtxOffset. This enables output of large
	// meshes (64K+ vertices) while still using 16-bits indices.
	BackendFlagsRendererHasVtxOffset = 1 << 3
)
