package imgui

const (
	// SelectableFlagsNone default = 0
	SelectableFlagsNone = 0
	// SelectableFlagsDontClosePopups makes clicking the selectable not close any parent popup windows.
	SelectableFlagsDontClosePopups = 1 << 0
	// SelectableFlagsSpanAllColumns allows the selectable frame to span all columns (text will still fit in current column).
	SelectableFlagsSpanAllColumns = 1 << 1
	// SelectableFlagsAllowDoubleClick generates press events on double clicks too.
	SelectableFlagsAllowDoubleClick = 1 << 2
	// SelectableFlagsDisabled disallows selection and displays text in a greyed out color.
	SelectableFlagsDisabled = 1 << 3
)
