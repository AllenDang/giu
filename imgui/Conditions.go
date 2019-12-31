package imgui

// Condition for SetWindow***(), SetNextWindow***(), SetNextTreeNode***() functions.
// Important: Treat as a regular enum! Do NOT combine multiple values using binary operators!
// All the functions above treat 0 as a shortcut to ConditionAlways.
type Condition int

const (
	// ConditionAlways sets the variable.
	ConditionAlways Condition = 1 << 0
	// ConditionOnce sets the variable once per runtime session (only the first call with succeed).
	ConditionOnce = 1 << 1
	// ConditionFirstUseEver sets the variable if the object/window has no persistently saved data (no entry in .ini file).
	ConditionFirstUseEver = 1 << 2
	// ConditionAppearing sets the variable if the object/window is appearing after being hidden/inactive (or the first time).
	ConditionAppearing = 1 << 3
)
