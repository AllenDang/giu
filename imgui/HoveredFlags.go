package imgui

const (
	// HoveredFlagsNone Return true if directly over the item/window, not obstructed by another window,
	// not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsNone = 0
	// HoveredFlagsChildWindows IsWindowHovered() only: Return true if any children of the window is hovered.
	HoveredFlagsChildWindows = 1 << 0
	// HoveredFlagsRootWindow IsWindowHovered() only: Test from root window (top most parent of the current hierarchy).
	HoveredFlagsRootWindow = 1 << 1
	// HoveredFlagsAnyWindow IsWindowHovered() only: Return true if any window is hovered.
	HoveredFlagsAnyWindow = 1 << 2
	// HoveredFlagsAllowWhenBlockedByPopup Return true even if a popup window is normally blocking access to this item/window.
	HoveredFlagsAllowWhenBlockedByPopup = 1 << 3
	// HoveredFlagsAllowWhenBlockedByActiveItem Return true even if an active item is blocking access to this item/window.
	// Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenBlockedByActiveItem = 1 << 5
	// HoveredFlagsAllowWhenOverlapped Return true even if the position is overlapped by another window
	HoveredFlagsAllowWhenOverlapped = 1 << 6
	// HoveredFlagsAllowWhenDisabled Return true even if the item is disabled
	HoveredFlagsAllowWhenDisabled = 1 << 7
)

// HoveredFlags combinations
const (
	HoveredFlagsRectOnly            = HoveredFlagsAllowWhenBlockedByPopup | HoveredFlagsAllowWhenBlockedByActiveItem | HoveredFlagsAllowWhenOverlapped
	HoveredFlagsRootAndChildWindows = HoveredFlagsRootWindow | HoveredFlagsChildWindows
)
