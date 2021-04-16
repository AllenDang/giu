package imgui

const (
	// ComboFlagNone default = 0
	ComboFlagNone = 0
	// ComboFlagPopupAlignLeft aligns the popup toward the left by default.
	ComboFlagPopupAlignLeft = 1 << 0
	// ComboFlagHeightSmall has max ~4 items visible.
	// Tip: If you want your combo popup to be a specific size you can use SetNextWindowSizeConstraints() prior to calling BeginCombo().
	ComboFlagHeightSmall = 1 << 1
	// ComboFlagHeightRegular has max ~8 items visible (default).
	ComboFlagHeightRegular = 1 << 2
	// ComboFlagHeightLarge has max ~20 items visible.
	ComboFlagHeightLarge = 1 << 3
	// ComboFlagHeightLargest has as many fitting items as possible.
	ComboFlagHeightLargest = 1 << 4
	// ComboFlagNoArrowButton displays on the preview box without the square arrow button.
	ComboFlagNoArrowButton = 1 << 5
	// ComboFlagNoPreview displays only a square arrow button.
	ComboFlagNoPreview = 1 << 6
)
