package imgui

const (
	// TabItemFlagsNone default = 0
	TabItemFlagsNone = 0
	// TabItemFlagsUnsavedDocument Append '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. Also: tab is selected on closure and closure is deferred by one frame to allow code to undo it
	// without flicker.
	TabItemFlagsUnsavedDocument = 1 << 0
	// TabItemFlagsSetSelected Trigger flag to programmatically make the tab selected when calling BeginTabItem()
	TabItemFlagsSetSelected = 1 << 1
	// TabItemFlagsNoCloseWithMiddleMouseButton  Disable behavior of closing tabs (that are submitted with
	// p_open != NULL) with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabItemFlagsNoCloseWithMiddleMouseButton = 1 << 2
	// TabItemFlagsNoPushID Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem()
	TabItemFlagsNoPushID = 1 << 3
)
